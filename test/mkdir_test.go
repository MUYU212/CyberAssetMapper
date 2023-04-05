package test

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var stRootDir string
var stSeparator string
var iJsonData map[string]any

const stJsonFileName = "dir.json"

func loadJson() {
	stSeparator = string(filepath.Separator)
	stWorkDir, _ := os.Getwd()
	stRootDir = stWorkDir[:strings.LastIndex(stWorkDir, stSeparator)] //获取到网站根目录

	gnJsonBytes, _ := os.ReadFile(stWorkDir + stSeparator + stJsonFileName)
	err := json.Unmarshal(gnJsonBytes, &iJsonData)
	if err != nil {
		fmt.Println(err)
	}
}

func parseMap(mapData map[string]any, stParentDir string) {
	for k, v := range mapData {
		switch v.(type) {
		case string:
			{
				path, _ := v.(string)
				if path == "" {
					continue
				}
				if stParentDir != "" {
					//fmt.Println(path)
					path = stParentDir + stSeparator + path
					if k == "text" {
						stParentDir = path
					}
				} else {
					stParentDir = path
				}
				createDir(path)
			}
		case []any:
			{
				parseArray(v.([]any), stParentDir)
			}

		}
	}
}

func parseArray(giJsonData []any, stParentDir string) {
	for _, v := range giJsonData {
		mapV, _ := v.(map[string]any)
		parseMap(mapV, stParentDir)
	}
}

func createDir(path string) {
	if path == "" {
		return
	}

	err := os.MkdirAll(stRootDir+stSeparator+path, fs.ModePerm)
	if err != nil {
		panic("Create dir error: " + err.Error())
	}
}

func TestGenerateDir(t *testing.T) {
	fmt.Println("根据json文件创建项目文件夹测试...")
	loadJson()
	parseMap(iJsonData, "")
}
