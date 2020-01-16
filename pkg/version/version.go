package version

import (
	"fmt"
	"runtime"
)

// Info contains versioning information.
type Info struct {
	GitTag       string `json:"gitTag"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

var (
	gitTag       string = ""
	gitCommit    string = "$Format:%H$"          // sha1 from git, output of $(git rev-parse HEAD)
	gitTreeState string = "not a git tree"       // state of git tree, either "clean" or "dirty"
	buildDate    string = "1970-01-01T00:00:00Z" // build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
)

// String returns info as a human-friendly version string.
func (info Info) String() string {
	return info.GitTag
}

func Get() Info {
	return Info{
		GitTag:       gitTag,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH), // 运行的设备和设备结构
	}
}
