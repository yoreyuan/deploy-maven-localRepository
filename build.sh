#!/usr/bin/env bash

export CGO_ENABLED=0
export GOARCH=amd64

OS_ARR=(
"linux"
"windows"
"darwin"
)

dirName="tools_maven_localRepository"
if [ -d "${dirName}" ]; then
  rm -rf ${dirName}
fi
mkdir ${dirName}
if [ -f "${dirName}-bin.zip" ]; then
  rm ${dirName}-bin.zip
fi

for osName in ${OS_ARR[@]}; do
  GOOS=${osName} go build -mod=vendor -ldflags="-s -w" -v -o tools_localRepo ./main.go
  mkdir ${dirName}/${osName}-${GOARCH}
  mv tools_localRepo ${dirName}/${osName}-${GOARCH}/
done

cp readme.md ${dirName}/
cp static/config.yaml.template ${dirName}/
cp static/settings.xml.template ${dirName}/

zip -r ${dirName}-bin.zip ${dirName}/
rm -rf ${dirName}
