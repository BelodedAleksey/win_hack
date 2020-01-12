package win

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"unsafe"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

//Utf2Ansi Change Encoging from UTF-8 to ANSI
func Utf2Ansi(str string) (string, error) {
	var windows1251 *charmap.Charmap = charmap.Windows1251
	bs := []byte(str)
	readerBs := bytes.NewReader(bs)
	readerWin := transform.NewReader(readerBs, windows1251.NewEncoder())
	bWin, err := ioutil.ReadAll(readerWin)
	if err != nil {
		return "", err
	}
	return string(bWin), nil
}

//ToInt func
func toInt(s string, max int) int32 {
	if strings.HasSuffix(s, "%") {
		i, err := strconv.Atoi(s[:len(s)-1])
		if err != nil {
			return 0
		}
		if i > 100 {
			i = 100
		}
		return int32(math.Trunc(float64(max*i) / 100))
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return int32(i)
}

//Unicode2utf8 func
func Unicode2utf8(source string) string {
	var res = []string{""}
	sUnicode := strings.Split(source, "\\u")
	var context = ""
	for _, v := range sUnicode {
		var additional = ""
		if len(v) < 1 {
			continue
		}
		if len(v) > 4 {
			rs := []rune(v)
			v = string(rs[:4])
			additional = string(rs[4:])
		}
		temp, err := strconv.ParseInt(v, 16, 32)
		if err != nil {
			context += v
		}
		context += fmt.Sprintf("%c", temp)
		context += additional
	}
	res = append(res, context)
	return strings.Join(res, "")
}

//StringToUintptr func
func StringToUintptr(str string) uintptr {
	if str == "" {
		return 0
	}
	//enc := mahonia.NewEncoder("GBK")
	//str1 := enc.ConvertString(str)
	a1 := []byte(str)
	p1 := &a1[0]
	return uintptr(unsafe.Pointer(p1))
}
