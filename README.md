deploy-maven-localRepository
============================

A tool that can publish packages in the Maven local repository

# 1 dev
```bash
# encoding

# build
go build -mod=vendor -ldflags="-s -w" -v -o deploy_local_repo ./main.go

# package
./build.sh

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


# 3 Use

```bash
chmod +x deploy_local_repo

# Execute the deploy command, 
# and the supported parameters are shown in the table below
./deploy_local_repo -s /opt/tmp/settings.xml -repo /opt/repo

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

