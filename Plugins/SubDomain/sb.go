package SubDomain

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

//go:embed data/temp.txt
var subdomains string

func getSubdomainInFile() []string {
	reader := bufio.NewScanner(strings.NewReader(subdomains))
	reader.Split(bufio.ScanLines) //bufio.ScanLines其实就是文本的换行符，将文本的数据逐行分割成切片
	var ret []string
	for reader.Scan() { //逐行读取
		ret = append(ret, reader.Text()) //将读取的数据添加到切片中
	}
	return ret
}

func GetSubDomain(domain string) []string {

	var result []string

	if !isWildCard(domain) {
		log.Println("进入爆破方法")
		var wg sync.WaitGroup
		wg.Add(len(getSubdomainInFile()))
		for _, subdomain := range getSubdomainInFile() {
			go func(subdomain string) {
				defer wg.Done()

				//log.Println("正在尝试解析====>" + subdomain + "." + domain)
				response := lookupSubDomain(subdomain, domain)
				if response != "" {
					result = append(result, response)
				}
			}(subdomain)
		}
		wg.Wait()
		return result
	}
	return result
}

//这里直接查询域名，不查询子域名

func lookupSubDomain(subdomain, domain string) string {

	fqdn := fmt.Sprintf("%s.%s", subdomain, domain)
	_, err := net.LookupIP(fqdn)
	if err != nil {
		return ""
	}
	log.Println("成功解析到" + fqdn)
	return fqdn
}

func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func isWildCard(domain string) bool {
	rand.Seed(time.Now().UnixNano())
	flag := false
	// 生成 5 位随机数
	min := 10000
	max := 99999
	num := rand.Intn(max-min+1) + min
	str := strconv.Itoa(num)

	s := lookupSubDomain(str, domain)
	if s != "" {
		log.Println("域名使用了泛解析")
		flag = true
		return flag
	}
	return flag
}
