package tipoff

import (
	"github.com/goravel/framework/database/orm"
)

type (
	IsNoticeType uint
)

const IsNoticeTypeYes IsNoticeType = 1
const IsNoticeTypeNo IsNoticeType = 0

type TipOffNoticeData struct {
	orm.Model
	OrigId        int64
	Title         string
	Content       string
	DateTime      string
	ShortTime     string
	ShiJianChuo   int64
	CateId        string
	CateName      string
	Comments      int64
	LouZhu        string
	LouZhuRegTime string
	Url           string
	PachongId     int64
	YuanUrl       string
	IsNotice      int64
}

func (r *TipOffNoticeData) TableName() string {
	return "tip_off_notice_data"
}
