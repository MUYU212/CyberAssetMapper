package test

import (
	"bytes"
	"fmt"
	"github.com/projectdiscovery/subfinder/v2/pkg/resolve"
	"github.com/projectdiscovery/subfinder/v2/pkg/runner"
	"io"
	"log"
	"testing"
)

func Test(t *testing.T) {
	subfinderTest()
}

func subfinderTest() {
	runnerInstance, err := runner.NewRunner(&runner.Options{
		Threads:            10,                       //使用的线程数量
		Timeout:            30,                       //超时时间
		MaxEnumerationTime: 10,                       //最大枚举时间
		Resolvers:          resolve.DefaultResolvers, //默认的解析器
		ResultCallback: func(s *resolve.HostEntry) { //结果回调函数
		},
	})

	buf := bytes.Buffer{}
	err = runnerInstance.EnumerateSingleDomain("hbu.cn", []io.Writer{&buf})
	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(&buf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", data)
}
