#!/usr/bin/env bash

export CGO_ENABLED=0
export GOARCH=amd64

OS_ARR=(
"linux"
"windows"
"darwin"
)

dirName="deploy_maven_localRepository"
if [ -d "${dirName}" ]; then
  rm -rf ${dirName}
fi
mkdir ${dirName}
if [ -f "${dirName}-bin.zip" ]; then
  rm ${dirName}-bin.zip
fi

for osName in ${OS_ARR[@]}; do
  GOOS=${osName} go build -mod=vendor -ldflags="-s -w" -v -o nexus_deploy-${osName}-${GOARCH} ./main.go
  mv nexus_deploy-${osName}-${GOARCH} ${dirName}/
done

cp readme.md ${dirName}/
cp settings.xml ${dirName}/

zip -r ${dirName}-bin.zip ${dirName}/
rm -rf ${dirName}
