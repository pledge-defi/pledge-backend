package utils

import (
	"errors"
	"os"
	"path/filepath"
)

// IsDir Determine whether the given path is a folder
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile Determine whether the given path is a file
func IsFile(path string) bool {
	return !IsDir(path)
}

// Create a directory
//func MkDir(path string) error {
//	return os.MkdirAll(path, os.ModePerm)
//}

func MkDir(dir string) error {
	rootDir, err := filepath.Abs("./")
	if err != nil {
		return errors.New("获取根目录失败，" + err.Error())
	}
	dir = rootDir + "/" + dir
	_, err = os.Stat(dir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
