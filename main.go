package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"regexp"
)

func main() {
	regstr := `\d+\.\d+\.\d+\.\d+`
	reg, _ := regexp.Compile(regstr)
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		match := reg.FindAllString(sc.Text(), -1)

		if match != nil {

			ipAddress := net.ParseIP(match[0]) //将stdin 转成ip类型
			if ipAddress.IsPrivate() {         // 判断是不是内网地址
				fmt.Println(sc.Text() + ",private") //是就再后面加一个 private 表示内网
			} else {
				fmt.Println(sc.Text() + ",public") // 否则就是公网
			}
		}

	}

}
