package httpx

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"goext/filex"
	"goext/str"
	"io"
	"net/http"
	"os"
)

const (
	Image = "jpg"
	Video = "mp4"
)

func Download(urls []string, dir, suffix string) ([]string, bool) {

	var files []string

	for _, url := range urls {

		name := str.Md5(url)
		path := fmt.Sprintf("%s%s.%s", dir, name, suffix)
		if filex.Exists(path) {
			files = append(files, path)
			continue
		}

		resp, err := http.Get(url)
		if err != nil {
			logrus.Error("[download]", err)
			return files, false
		}
		var out *os.File
		if out, err = filex.CreateFile(path); err != nil {
			logrus.Error("[download]", err)
			return files, false
		}
		if _, err := io.Copy(out, resp.Body); err != nil {
			logrus.Error("[download]", err)
			return files, false
		}
		files = append(files, path)
	}

	return files, len(files) > 0
}
