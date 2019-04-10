#!/bin/bash
#############################################################################
# 使用帮助
if [ $# -ne 1 ] || [ "-h" = "$1" ] || [ "--help" = "$1" ]
then
    echo "介绍: 初始化项目"
    echo "用法: sh $0 gin-toolkit"
    echo "说明：gin-toolkit 是你项目名称"
    echo "项目目录$GOPATH/src/github.com/TonyDoen/go_code_review/gin-toolkit"
    exit
fi
#############################################################################
export LANGUAGE="utf-8"

CATALOG="$GOPATH/src/github.com/TonyDoen/go_code_review/$1"
mkdir -p "${CATALOG}"
rsync -rv --exclude=.git --exclude=.idea ./ "${CATALOG}"
cd "${CATALOG}"
rm $0
grep -rl 'gin-toolkit' ./  | xargs sed -i "" "s/gin-toolkit/$1/g"