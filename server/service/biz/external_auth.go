package biz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz"
)

type ExternalAuthService struct {
}

func (externalAuthService *ExternalAuthService) GetAuthInfo(source string) (auth string, err error) {
	db := global.GVA_DB.Model(&biz.ExternalAuth{})
	db.Where("source = ?", source)

	res := biz.ExternalAuth{}

	err = db.Find(&res).Error
	return  res.Auth.String(), err
}