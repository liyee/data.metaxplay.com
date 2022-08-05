package help

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/oschwald/geoip2-golang"
)

func GetClient(ipClient string) *geoip2.Country {
	db, err := geoip2.Open("GeoLite2-Country.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// If you are using strings that may be invalid, check that ip is not nil
	ip := net.ParseIP(ipClient)
	record, err := db.Country(ip)
	if err != nil {
		log.Fatal(err)
	}

	return record
}

func LogFile(content, from, test string) {
	var testName string
	if test == "1" {
		testName = "test"
	} else {
		testName = "release"
	}

	d := time.Now().Format("20060102")
	filePath := from + "/" + testName + "/" + d + ".log"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()

	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	write.WriteString(content + "\n")
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
}

func Regroup(data string) map[string]interface{} {
	var res interface{}
	buf := make(map[string]interface{})
	err := json.Unmarshal([]byte(data), &res)

	if err == nil {
		m, ok := res.(map[string]interface{})
		if ok {
			for k, v := range m {
				buf[k] = v
			}
		}
	}

	return buf
}
