package handlers

import "CyberAssetMapper/model"

func InsertSubdomain() {
	model.Insert_SubDomain("www.baidu.com", 2)
}
