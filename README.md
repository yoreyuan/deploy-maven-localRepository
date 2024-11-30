tools-maven-localRepository
============================

A maven localRepository tools, the main functions are as follows:
* Support **publishing local packages** to Maven repository
* Support **clearing files** with specified suffix in local repository
* Supports specifying values through parameters (**Deprecated**)
* Supports specifying parameters through yaml files


# 1 dev
```bash
# encoding

# build
go build -mod=vendor -ldflags="-s -w" -v -o tools_localRepo ./main.go

# package
# It will generate the 'tools_maven_localRepository-bin.zip' file
./build.sh

```

The 'tools_maven_localRepository-bin.zip' directories are described as follows：
```
tools_maven_localRepository
├── config.yaml.template    # APP configuration file template
├── darwin-amd64    # Mac OS
│   └── tools_localRepo
├── linux-amd64     # linux OS
│   └── tools_localRepo
├── readme.md
├── settings.xml.template       # maven settings.xml template
└── windows-amd64   # Windows OS
    └── tools_localRepo

```


# 2 Support
| packaging      | classifier      | example                                               |
|----------------|-----------------|-------------------------------------------------------|
| `${packaging}` | `${classifier}` | `${artifactId}-${version}-${classifier}.${packaging}` |
| pom            | -               | airbase-190.pom                                       |
| - / jar        | -               | slice-2.3.jar                                         |
| jar            | executable      | trino-verifier-463-executable.jar                     |
| jar            | sources         | trino-verifier-463-sources.jar                        |
| jar            | tests           | trino-verifier-463-tests.jar                          |
| jar            | test-sources    | trino-verifier-463-test-sources.jar                   |
| uexe           | -               | credValidator-2.3.0.uexe                              |
| exe            | linux-x86_64    | protoc-3.21.7-linux-x86_64.exe                        |


Folders with the following suffixes are ignored
* `.DS_Store`
* `.lastUpdated`
* `.md5`
* `.repositories`
* `.sha1`
* `.sha256`
* `.sha512`
* `.xml`


# 3 Getting Started

Prerequisites:
* mvn
* JDK

## 3.1 By parameters (Deprecated)
```bash
chmod +x tools_localRepo

# Execute the deploy command, 
# and the supported parameters are shown in the table below
./tools_localRepo -s /opt/tmp/settings.xml -repo /opt/repo

```

| Parameter    | Default value                                  | Example               | Explain                              |
|--------------|------------------------------------------------|-----------------------|--------------------------------------|
| `-help`/`-h` | -                                              |                       | Display help information             |
| `-s`         | ${HOME}/.m2/settings.xml                       | /opt/tmp/settings.xml |                                      |
| `-repo`      | ${HOME}/.m2/repository                         | /opt/repo             |                                      |
| `-url`       | http://nexus.yore.cn/repository/maven-releases |                       |                                      |
| `-repoId`    | yore_nexus                                     |                       |                                      |
| `-verbose`   | -                                              |                       | Output project debugging information |
| `-X`         | -                                              |                       | Output maven debugging information   |


## 3.2 By yaml file
Or through a configuration file with the parameter `--config`
```bash
./tools_localRepo --config ./config.yaml
```

For more configuration information, see configuration template file [`config.yaml.template`](static/config.yaml.template)
```yaml
# Output project debugging information
verbose: false
logLevel: "INFO"

# The path to the local repository maven will use to store artifacts
localRepository: "~/.m2/repository"

deploy:
  enable: true
  commandName: "/usr/local/bin/mvn"
  # Alternate path for the user settings file
  settingXml: "~/.m2/settings.xml"
  # The ID of the repository
  id: "yore_nexus"
  # The URL of the repository maven
  url: "http://nexus.yore.cn/repository/maven-releases"
  # Output maven debugging information
  debug: false

  # The list of file suffixes to be ignored
  excludeSuffixs:
    - ".DS_Store"
    - ".asc"
    - ".lastUpdated"
    - ".md5"
    - ".repositories"
    - ".sha1"
    - ".sha256"
    - ".sha512"
    - ".xml"

clean:
  # If it is in cleaning mode, the list of file suffixes that will be cleaned up
  enable: false
  suffixs:
    - ".DS_Store"
    - ".lastUpdated"

```
