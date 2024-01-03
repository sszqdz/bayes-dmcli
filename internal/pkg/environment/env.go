package environment

import (
	"bayes/dmcli/internal/pkg/dirr"
	"os"
	"path/filepath"
	"sync"

	"github.com/spf13/viper"
)

type ENV string

const (
	EnvDebug   ENV = "debug"
	EnvRelease ENV = "release"

	key      string = "ENV"
	fileName string = ".env"
)

var (
	Deepth  = 3
	loadEnv = sync.OnceValue[string](extractEnv)
)

// do not auto inject, use when runtime
func LoadEnv() ENV {
	return ENV(loadEnv())
}

func (e ENV) Is(env ENV) bool {
	return e == env
}

func (e ENV) String() string {
	return string(e)
}

func extractEnv() string {
	// 获取环境变量
	env := os.Getenv(key)
	if len(env) > 0 {
		return env
	}
	// 向上递归查找 .env 文件
	curPath, err := os.Executable()
	panicErr(err)
	curDir := filepath.Dir(curPath)
	envFilePath, err := dirr.FindFileInParentDirs(fileName, curDir, Deepth)
	if os.IsNotExist(err) {
		return ""
	}
	panicErr(err)

	vEnv := viper.New()
	vEnv.SetConfigFile(envFilePath)
	panicErr(vEnv.ReadInConfig())
	return vEnv.GetString(key)
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}
