// 生成DGA域名
package main

import (
	"crypto/md5"
	"fmt"
	"strings"
	"time"
)

var ch = make(chan string)

func createDomain(year, month, day, count int, secret string, tld string) {
	if !strings.HasPrefix(tld, ".") {
		tld = "." + tld
	}

	md5Hash := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%v%v%v%v%v", year, month, day, count,secret))))
	domain := fmt.Sprintf("%v%v", md5Hash[12:23], tld)
	ch <- domain
}

func main() {
	t := time.Now()
	y, m, d := t.Year(), int(t.Month()), t.Day()

	for i:=1; i<10; i++ {
		go createDomain(y, m, d, i,"Chelsea", "com")
		
		domain := <-ch
		fmt.Println(domain)
	}
}
