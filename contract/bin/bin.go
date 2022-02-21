package binfile

import (
	"os"
	"path"
	"runtime"
)

func GetBinByToken(token string) (string, error) {
	currentAbPath := getCurrentAbPathByCaller()
	by, err := os.ReadFile(currentAbPath + "/" + token + "abi")
	if err != nil {
		return "", err
	}
	return string(by), nil
}

func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
