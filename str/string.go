package str

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func Utf8ToGBK(utf8str string) string {
	result, _, err := transform.String(simplifiedchinese.GBK.NewEncoder(), utf8str)
	if err != nil {
		fmt.Println(err.Error())
	}
	return result
}

func RandString(length int) string {
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, length)
	scope := [][]int{{48, 10}, {97, 26}, {65, 26}}
	for i := 0; i < length; i++ {
		k := scope[rand.Intn(3)]
		b := k[0] + rand.Intn(k[1])
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func ip4_ntoa(ipnr uint32) string {
	var bytes [4]byte
	bytes[0] = byte(ipnr & 0xFF)
	bytes[1] = byte((ipnr >> 8) & 0xFF)
	bytes[2] = byte((ipnr >> 16) & 0xFF)
	bytes[3] = byte((ipnr >> 24) & 0xFF)

	return fmt.Sprintf("%d.%d.%d.%d",bytes[3],bytes[2],bytes[1],bytes[0])
}

// Convert net.IP to int64 ,  http://www.outofmemory.cn
func ip4_aton(ipnr string) uint32 {
	bits := strings.Split(ipnr, ".")

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum uint32

	sum += uint32(b0) << 24
	sum += uint32(b1) << 16
	sum += uint32(b2) << 8
	sum += uint32(b3)

	return sum
}

func RandIP4Int() uint32 {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Uint32()
}

func RandIP4Str() string {
	ip4int := RandIP4Int()
	return ip4_ntoa(ip4int)
}

func Str2Int(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func Int2Str(i int) string {
	return strconv.Itoa(i)
}

func Base64Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Base64Decode(s string) ([]byte, error) {
	ds, err := base64.StdEncoding.DecodeString(s)
	return ds, err
}

func Base64UrlEncode(b []byte) string {
	return base64.URLEncoding.EncodeToString(b)
}

func Base64UrlDecode(s string) ([]byte, error) {
	ds, err := base64.URLEncoding.DecodeString(s)
	return ds, err
}

func JsonEncode(obj interface{}) []byte {
	b, _ := json.Marshal(obj)
	return b
}

func JsonDecode(data []byte, obj interface{}) {
	json.Unmarshal(data, obj)
}

func Md5(s string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(s))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}
