package httpx

import (
	"fmt"
	"testing"
)

func TestProxyThorn(t *testing.T) {
	verification("163.125.114.33:8088")
}

// ip验证
func verification(ipAddr string) {
	var ip, status = ProxyThorn(ipAddr)
	//判断是否有返回ip，并且请求状态为200
	if status == 200 && ip != "" {
		fmt.Println(ipAddr + " 请求 http://icanhazip.com 返回ip:【" + ip + "】-【检测结果：可用】")
	} else {
		fmt.Println(ipAddr + " 请求 http://icanhazip.com 返回ip:【" + ip + "】-【检测结果：不可用】")
	}

}