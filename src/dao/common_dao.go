package dao

import (
	"CyberAssetMapper/src/service/dto"
	"gorm.io/gorm"
)

func Paginate(p dto.Paginate) func(orm *gorm.DB) *gorm.DB {
	return func(orm *gorm.DB) *gorm.DB {
		return orm.Offset((p.GetPage() - 1) * p.GetLimit()).Limit(p.GetLimit())
	}
}
