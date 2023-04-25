package test

import (
	"CyberAssetMapper/src/utils"
	"bytes"
	"fmt"
	"github.com/projectdiscovery/subfinder/v2/pkg/resolve"
	"github.com/projectdiscovery/subfinder/v2/pkg/runner"
	"io"
	"log"
)

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
	lines := utils.SplitLines(string(data))
	for _, value := range lines {
		fmt.Println(value)
	}
}
