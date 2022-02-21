package common

import (
	"path"
	"pledge-backend/config"
	"runtime"
)

var Config *config.Conf

func GetCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
