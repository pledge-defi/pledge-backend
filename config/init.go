package config

import (
	"github.com/BurntSushi/toml"
	"path"
	"path/filepath"
	"runtime"
)

func init() {
	currentAbPath := getCurrentAbPathByCaller()
	tomlFile, err := filepath.Abs(currentAbPath + "/configV21.toml")
	//tomlFile, err := filepath.Abs(currentAbPath + "/configV22.toml")
	if err != nil {
		panic("read toml file err: " + err.Error())
		return
	}
	if _, err := toml.DecodeFile(tomlFile, &Config); err != nil {
		panic("read toml file err: " + err.Error())
		return
	}
}

func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
