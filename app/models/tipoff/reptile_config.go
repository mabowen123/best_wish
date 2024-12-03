package tipoff

import (
	"github.com/goravel/framework/database/orm"
)

type (
	ReptileConfigType uint
)

const ReptileConfigTypeTipOff ReptileConfigType = 1

type ReptileConfig struct {
	orm.Model
	Url          string
	IntervalTime uint64
	Type         ReptileConfigType
	NextTime     uint64
	Remark       string
}

func (r *ReptileConfig) TableName() string {
	return "reptile_config"
}
