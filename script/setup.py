import os

arch = [
        'darwin:386',
        'darwin:amd64',
        'dragonfly:amd64',
        'freebsd:386',
        'freebsd:amd64',
        'freebsd:arm',
        'linux:386',
        'linux:amd64',
        'linux:arm',
        'linux:arm64',
        'linux:ppc64',
        'linux:ppc64le',
        'linux:mips',
        'linux:mipsle',
        'linux:mips64',
        'linux:mips64le',
        'netbsd:386',
        'netbsd:amd64',
        'netbsd:arm',
        'openbsd:386',
        'openbsd:amd64',
        'openbsd:arm',
        'plan9:386',
        'plan9:amd64',
        'solaris:amd64',
        'windows:386',
        'windows:amd64'
    ]

__PAYLOAD_CODE__ = "1337"

def cross_compile():
    def compile(go_arch: str, go_os: str):
        cmd = f'GOOS={go_os} GOARCH={go_arch} go build -o ./dist/gorianet_{__PAYLOAD_CODE__}_{go_arch}_{go_os}{".exe" if go_os == "windows" else ""} ../bot/'
        os.system(cmd)
        print(f'* {go_os} --> {go_arch}')

    
    for proc in arch:
        parsed = proc.split(':')
        go_os, go_arch = parsed[0], parsed[1]

        compile(go_arch, go_os)
    
    """
        Note: thread made crash vps..
        
        thl = []
        thl.append(threading.Thread(target= compile, args= (go_arch, go_os,)))

        for t in thl:
            t.start()

        for t in thl:
            t.join()
    """

os.system('rm -rf ./dist && mkdir ./dist && clear')
print('''
>> GoriaNet cross compiler. made by github.com/Its-Vichy/GoriaNet
''')
cross_compile()