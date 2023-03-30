package model

import (
	"CyberAssetMapper/model/db"
	"fmt"
	"gorm.io/gorm"
)

type SubDomain struct {
	gorm.Model
	Domain string
	TaskID uint `gorm:"not null;index"`
}

func Create_subDomainTable() {
	db.GLOBAL_DB.AutoMigrate(&SubDomain{})
	//创建SubDomain表
}

//插入子域名
func Insert_SubDomain(Domain string, taskID uint) (*SubDomain, error) {

	subDomain := &SubDomain{
		Domain: Domain,
		TaskID: taskID,
	}

	dbres := db.GLOBAL_DB.Create(subDomain)
	if dbres.Error != nil {
		panic(dbres.Error)
	} else {
		fmt.Println("插入成功")
		return subDomain, nil
	}

}
