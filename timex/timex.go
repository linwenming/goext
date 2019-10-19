package timex

import (
	"math/rand"
	"strings"
	"time"
)

const FORMAT string = "2006-01-02 15:04:05"
const FORMATDATA string = "2006-01-02 "
const FORMATSTR string = "2006/01/02/15"

//StrToTime 字符串转换成Time时间表示
func Parse(value string) time.Time {
	if value == "" {
		return time.Time{}
	}
	layouts := []string{
		"2006-01-02 15:04:05 -0700 MST",
		"2006-01-02 15:04:05 -0700",
		"2006-01-02 15:04:05",
		"2006/01/02 15:04:05 -0700 MST",
		"2006/01/02 15:04:05 -0700",
		"2006/01/02 15:04:05",
		"2006-01-02 -0700 MST",
		"2006-01-02 -0700",
		"2006-01-02",
		"2006/01/02 -0700 MST",
		"2006/01/02 -0700",
		"2006/01/02",
		"2006-01-02 15:04:05 -0700 -0700",
		"2006/01/02 15:04:05 -0700 -0700",
		"2006-01-02 -0700 -0700",
		"2006/01/02 -0700 -0700",
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	}

	var t time.Time
	var err error
	for _, layout := range layouts {
		t, err = time.Parse(layout, value)
		if err == nil {
			return t
		}
	}
	//panic(err)
	return time.Time{}
}

func String(t time.Time) string {
	return Format(FORMAT, t)
}

// https://github.com/piaohua
// 如果 ts 没传递，则使用当前时间
func Format(format string, ts ...time.Time) string {
	patterns := []string{
		// 年
		"Y", "2006", // 4 位数字完整表示的年份
		"y", "06",   // 2 位数字表示的年份

		// 月
		"m", "01",      // 数字表示的月份，有前导零
		"n", "1",       // 数字表示的月份，没有前导零
		"M", "Jan",     // 三个字母缩写表示的月份
		"F", "January", // 月份，完整的文本格式，例如 January 或者 March

		// 日
		"d", "02", // 月份中的第几天，有前导零的 2 位数字
		"j", "2",  // 月份中的第几天，没有前导零

		"D", "Mon",    // 星期几，文本表示，3 个字母
		"l", "Monday", // 星期几，完整的文本格式;L的小写字母

		// 时间
		"g", "3",  // 小时，12 小时格式，没有前导零
		"G", "15", // 小时，24 小时格式，没有前导零
		"h", "03", // 小时，12 小时格式，有前导零
		"H", "15", // 小时，24 小时格式，有前导零

		"a", "pm", // 小写的上午和下午值
		"A", "PM", // 小写的上午和下午值

		"i", "04", // 有前导零的分钟数
		"s", "05", // 秒数，有前导零
	}
	replacer := strings.NewReplacer(patterns...)
	format = replacer.Replace(format)

	t := time.Now()
	if len(ts) > 0 {
		t = ts[0]
	}
	return t.Format(format)
}

func RandSecond(sec int) time.Duration {
	if sec < 4 {
		return time.Duration(sec) * time.Second
	}
	rand.New(rand.NewSource(time.Now().UnixNano()))
	seed := sec / 2
	return time.Duration(seed + rand.Intn(seed+1)) * time.Second
}

func RandSleep(d time.Duration) {
	sec := d.Seconds()
	if sec < 4 {
		time.Sleep(d)
		return
	}
	rand.New(rand.NewSource(time.Now().UnixNano()))
	seed := int64(sec / 2)
	seed = seed + rand.Int63n(seed+1)
	time.Sleep(time.Duration(seed) * time.Second)
}