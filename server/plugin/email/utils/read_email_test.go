package utils

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"regexp"
	"testing"
)

func TestUsage(t *testing.T) {
	global.GVA_CONFIG.Email.From = "enroutejewelry.dev@outlook.com"
	global.GVA_CONFIG.Email.Host = "outlook.office365.com"
	global.GVA_CONFIG.Email.Secret = "tQohnD5eUzJY5H7QyvdC"
	global.GVA_CONFIG.Email.Nickname = "enroutejewelry.dev"
	global.GVA_CONFIG.Email.Port = 993
	reg, _ := regexp.Compile(`https://shipbobattachments.shipbob.com/api/attachment/[\w-]+`)

	SearchEmail("Your Product Catalog Export", reg)
}
