package constant

import "path/filepath"

var (
	LocalRepository = ""
	SettingXml      = ""
	RepoUrl         = "http://nexus.yore.cn/repository/maven-releases"
	RepoId          = "yore_nexus"
	Verbose         = false
	MvnDebug        = false
	ExcludeSuffixs  = []string{
		".DS_Store",
		".asc",
		".lastUpdated",
		".md5",
		".repositories",
		".sha1", ".sha256", ".sha512",
		".xml"}
	Separator = string(filepath.Separator)
)
