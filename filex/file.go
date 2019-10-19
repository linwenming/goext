package filex

import (
	"os"
	"path/filepath"
)

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err == nil {
		return true
	}
	if os.IsExist(err) {
		return true
	}
	return false
}

func CreateDir(file string) error {
	_, err := os.Stat(file)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		dir := filepath.Dir(file)
		err := os.MkdirAll(dir, os.ModePerm)
		return err
	}
	return nil
}

func CreateFile(file string) (*os.File, error) {
	var err error
	if err = CreateDir(file); err != nil {
		return nil, err
	}
	return os.Create(file)
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}
