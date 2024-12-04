package tipoff

import (
	tipoffdao "best_wish/app/dao/tipoff"
	"best_wish/lib/wxpusher"
	"best_wish/until"
	"fmt"
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
	"github.com/goravel/framework/facades"
	"strings"
	"time"
)

type Notify struct {
}

// Signature The name and signature of the console command.
func (receiver *Notify) Signature() string {
	return "tip:off:notice"
}

// Description The console command description.
func (receiver *Notify) Description() string {
	return "线报通知"
}

// Extend The console command extend.
func (receiver *Notify) Extend() command.Extend {
	return command.Extend{}
}

// Handle Execute the console command.
func (receiver *Notify) Handle(ctx console.Context) error {
	list, err := tipoffdao.GetNeedNoticeList()
	if err != nil {
		facades.Log().Errorf("查询通知列表出错%s", err)
		return err
	}
	for _, tipoff := range list {
		time.Sleep(time.Second)

		url := tipoff.Url

		if !strings.HasPrefix(url, "http") {
			url = until.JoinDomain("http://new.xianbao.fun", tipoff.Url)
		}

		now := time.Now()
		nowHour := now.Hour()
		nowWeekday := now.Weekday()

		isNotice := true

		if (nowHour < 2 || nowHour > 6) || (nowWeekday == time.Saturday || nowWeekday == time.Sunday) {
			isNotice = wxpusher.SendMsg(&wxpusher.SendTongzhiParams{
				AppToken:    "AT_AAixJoECoUTJMyoN0ELrATDYHHu34qLy",
				Content:     fmt.Sprintf("<h1>%s</h1>\n\t<p>%s</p>\n   <a href=\"%s\">🔗查看详情</a>", tipoff.Title, tipoff.Content, url),
				Summary:     tipoff.Title,
				ContentType: 2,
				TopicIds: []int{
					25804,
				},
				Url:       url,
				VerifyPay: 0,
			})
		}

		if isNotice {
			tipoffdao.UpdateIsNotice(tipoff.ID)
		}
	}

	return nil
}
