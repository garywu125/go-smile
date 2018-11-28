# release differnet platform execution file
# 處理步驟:

# tag via git :
#   add tag : commit/push all/add new tag name
#            - git tag -a <tag-name> -m "message"
#   checkout tag: git checkout <tag-name>
#   sync remote repository: git push origin <tag_name>

# build on shell:
#   switch to powershell
#   batch process: PowerShell.exe -Command "./release.ps1" <tag-name>

# upload tag release buid to github.com

$Env:GOOS = "linux"; $Env:GOARCH = "amd64";go build -o ../release/test1/v0.1.1/test1_linux_amd64
$Env:GOOS = "windows"; $Env:GOARCH = "amd64";go build -o ../release/test1/v0.1.1/test1_windows_amd64.exe