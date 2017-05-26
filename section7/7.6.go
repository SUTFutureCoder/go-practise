package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Append系列函数将整数等转换为字符串后，添加到现有的字节数组中
	str     := make([]byte, 0)
	str     = strconv.AppendInt(str, 4567, 10)
	str     = strconv.AppendBool(str, false)
	str     = strconv.AppendQuote(str, "abcdefg")
	str     = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str))

	//Format系列函数吧其他类型的转换为字符串
	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1024)
	fmt.Println(a, b, c, d, e)

	//Parse系列函数把字符串转换为其他类型
	f, err := strconv.ParseBool("false")
	if err != nil {
		fmt.Println(err)
	}
	g, err := strconv.ParseFloat("123.23", 64)
	if err != nil {
		fmt.Println(err)
	}

	h, err := strconv.ParseInt("1234", 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	i, err := strconv.ParseUint("12345", 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	j, err := strconv.Atoi("1023")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(f, g, h, i, j)
}
