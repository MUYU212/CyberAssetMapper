package SubDomain

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

func getSubdomainInFile() []string {
	dir, _ := os.Getwd()
	filename := dir + "/Plugins/data/subdomain.txt"

	lines, err := readLines(filename)

	if err != nil {
		fmt.Println("读取文件时出错")
		fmt.Println(err)
		return nil
	}
	return lines

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
				//log.Println("正在尝试解析" + subdomain + "." + domain)
				domain = subdomain + "." + domain
				response := lookupDomain(domain)
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

func lookupDomain(domain string) string {

	fqdn := fmt.Sprintf("%s", domain)
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
	randomDomain := str + "." + domain
	s := lookupDomain(randomDomain)
	if s != "" {
		log.Println("域名使用了泛解析")
		flag = true
		return flag
	}
	return flag
}
