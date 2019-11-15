package httpx

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"goext/filex"
	"goext/str"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
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

/*
验证代理ip是否可用
通过传入一个代理ip，然后使用它去访问一个url看看是否访问成功，以此为依据进行判断当前代理ip是否有效。
参数：proxy_addr 要验证的ip
返回：ip 验证通过的ip、status 状态（200表示成功）
*/
func ProxyThorn(proxyAddr string) (ip string, status int) {
	httpUrl := "http://icanhazip.com"
	proxy, err := url.Parse(proxyAddr)

	netTransport := &http.Transport{
		Proxy:                 http.ProxyURL(proxy),
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second * time.Duration(5),
	}
	httpClient := &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}
	res, err := httpClient.Get(httpUrl)
	if err != nil {
		//fmt.Println("错误信息：",err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		//log.Println(err)
		return
	}
	c, _ := ioutil.ReadAll(res.Body)
	return string(c), res.StatusCode
}
