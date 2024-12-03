package tipoff

import (
	tipoffmodel "best_wish/app/models/tipoff"
	"github.com/goravel/framework/facades"
	"time"
)

func GetNeedReptileConfigList() ([]tipoffmodel.ReptileConfig, error) {
	var reptileConfigList []tipoffmodel.ReptileConfig

	err := facades.Orm().Query().
		Select("id,url,next_time,interval_time").
		Where("type = ?", tipoffmodel.ReptileConfigTypeTipOff).
		Where("next_time <= ?", time.Now().Unix()).
		Get(&reptileConfigList)

	return reptileConfigList, err
}

func UpdateNextTime(id uint, interval uint64) {
	facades.Orm().Query().Model(&tipoffmodel.ReptileConfig{}).
		Where("id", id).
		Update("next_time", uint64(time.Now().Unix())+interval)
}
