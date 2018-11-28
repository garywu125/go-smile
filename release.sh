#!/bin/bash
# release differnet platform execution file
# 處理步驟:

# tag via git :
#   add tag : commit/push all/add new tag name
#            - git tag -a <tag-name> -m "message"
#   check out tag: git checkout <tag-name>
#   sync remote repository: git push origin <tag_name>

#  build local tag release on local shell:
#   switch to git shell
#   batch process: sh ./release.sh <app-name> <tag-name>

# upload local tag release build to update github.com release

 app="$1"
 tag="$2"
 win_build="../release/$app/$tag/${app}_windows_amd64.exe"
 linux_build="../release/$app/$tag/${app}_linux_amd64"
 GOOS=linux GOARCH=amd64 go build -o $linux_build
 GOOS=windows GOARCH=amd64 go build -o $win_build





# echo Script Name: "$0"
# echo Total Number of Argument Passed: "$#"
# echo Arguments List -
# echo 1. $1
# echo 2. $2
# echo 3. $3
# echo tag. $app
# echo tag. $tag
# echo  windows $win_build
# echo  linux $linux_build
# echo All Arguments are: "$*"