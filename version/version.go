package version

import (
	"fmt"
	"runtime"
)

var (
	Version    string // 版本
	BuildTime  string // 编译时间
	CommitHash string // commit hash
)

// Print 打印版本信息
func Print() string {
	kv := [][2]string{
		{"Go", runtime.Version()},
		{"BuildTime", BuildTime},
		{"CommitHash", CommitHash},
		{"Version", Version},
	}
	str := ""
	for _, v := range kv {
		str += fmt.Sprintf("%10s: %s\n", v[0], v[1])
	}
	return str
}
