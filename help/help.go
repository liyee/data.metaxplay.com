package help

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"data.metaxplay.com/common"
	"github.com/oschwald/geoip2-golang"
)

func GetClient(ipClient string) *geoip2.Country {

	db, err := geoip2.Open(common.CONFIG.System.GeoDir + "/GeoLite2-Country.mmdb")
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

func LogFile(content, from, test string, dir string) {
	var testName string
	if test == "1" {
		testName = "test"
	} else {
		testName = "release"
	}

	fromMap := map[string]int{"web": 1, "app": 1}

	d := time.Now().Format("2006010215")
	filePath := dir + "/" + from + "/ob/" + d + ".log"
	if fromMap[from] == 1 {
		filePath = dir + "/" + from + "/ob/" + testName + "/" + d + ".log"
	}

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
