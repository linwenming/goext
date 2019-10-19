package logger

import (
	"goext/filex"
	"os"
	"path/filepath"
)

type LogWriter struct {
	file *os.File
}

func New(logPath string) (*LogWriter, error) {

	dir :=  filepath.Dir(logPath)
	if filex.Exists(dir) {
		if err:= os.MkdirAll(dir, os.ModePerm) ; err != nil {
			return nil, err
		}
	}

	file, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return &LogWriter{file: file}, nil
}

func (m *LogWriter) Write(b []byte) (n int, err error) {
	//_, _ = os.Stdout.Write(b)
	return m.file.Write(b)
}
