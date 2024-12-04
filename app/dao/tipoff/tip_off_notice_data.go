package tipoff

import (
	tipoffmodel "best_wish/app/models/tipoff"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/carbon"
)

func FirstOrCreateByOrigId(data tipoffmodel.TipOffNoticeData) {
	facades.Orm().Query().
		FirstOrCreate(&data, tipoffmodel.TipOffNoticeData{OrigId: data.OrigId})
}

func GetNeedNoticeList() ([]tipoffmodel.TipOffNoticeData, error) {
	var needNoticeListList []tipoffmodel.TipOffNoticeData
	err := facades.Orm().Query().
		Select("id,url,title,content").
		Where("is_notice = ?", tipoffmodel.IsNoticeTypeNo).
		OrderByDesc("id").
		Where("created_at >= ?", carbon.Now().SubHours(3).ToDateTimeMicroString()).
		Get(&needNoticeListList)

	return needNoticeListList, err
}

func UpdateIsNotice(ids []uint) {
	facades.Orm().Query().Model(&tipoffmodel.TipOffNoticeData{}).
		Where("is_notice  = ?", tipoffmodel.IsNoticeTypeNo).
		Where("id IN ?", ids).
		Update("is_notice", tipoffmodel.IsNoticeTypeYes)
}
func DelOldData() {
	facades.Orm().Query().Where("created_at < ?", carbon.Now().SubMonth().ToDateTimeMicroString()).Delete(&tipoffmodel.TipOffNoticeData{})
}
