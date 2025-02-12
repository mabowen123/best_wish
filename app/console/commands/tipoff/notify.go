package tipoff

import (
	tipoffdao "best_wish/app/dao/tipoff"
	"best_wish/lib/wxpusher"
	"best_wish/until"
	"fmt"
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/str"
	"net/http"
	"strings"
	"time"
	"unicode/utf8"
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

func checkUrl(url string) (string, bool) {
	if !strings.HasPrefix(url, "http") {
		url = until.JoinDomain("http://new.xianbao.fun", url)
	}

	resp, err := http.Head(url)

	if err != nil {
		facades.Log().Infof("请求目标链接异常", url)
		return url, false
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		facades.Log().Infof("请求目标链接404", url)
		return url, false
	}

	return url, true
}

// Handle Execute the console command.
func (receiver *Notify) Handle(ctx console.Context) error {
	list, err := tipoffdao.GetNeedNoticeList()
	if err != nil {
		facades.Log().Errorf("查询通知列表出错%s", err)
		return err
	}

	noticeIds := []uint{}
	var content strings.Builder
	var summary strings.Builder
	for _, tipoff := range list {
		url, isOk := checkUrl(tipoff.Url)

		if isOk {
			content.WriteString(fmt.Sprintf("<div style='margin-bottom: 20px';><h1>%s</h1>\n\t<p>%s</p>\n   <a href=\"%s\">🔗查看详情</a></div>", tipoff.Title, tipoff.Content, url))
			summary.WriteString(fmt.Sprintf("%s;", str.Of(until.ReplaceAllCharAndEmojiToBlank(tipoff.Title, []string{"!", "@", "#", "$", "%", " ", "|", "｜", ",", "，", "/", "~"})).Substr(0, 19)))
			noticeIds = append(noticeIds, tipoff.ID)
		}

		if utf8.RuneCountInString(summary.String()) < 40 {
			continue
		}

		hour := time.Now().Hour()
		content.WriteString(fmt.Sprintf("<img src=\"https://cdn.weipaitang.com/sky/yzlzs/imagecb/image/20250212/24ac16e1218c48bc8fd859ebaca8275a-W750H1350\" alt=\"加载失败\" width=\"600px\">"))
		isNotice := true
		if hour < 2 || hour > 6 {
			isNotice = wxpusher.SendMsg(&wxpusher.SendTongzhiParams{
				AppToken:    "AT_AAixJoECoUTJMyoN0ELrATDYHHu34qLy",
				Content:     content.String(),
				ContentType: 2,
				Summary:     str.Of(summary.String()).ReplaceLast(";", "").String(),
				TopicIds: []int{
					25804,
				},
				VerifyPay: 0,
			})
		}

		if isNotice {
			tipoffdao.UpdateIsNotice(noticeIds)
			time.Sleep(time.Second)
		}

		noticeIds = []uint{}
		content.Reset()
		summary.Reset()
	}

	return nil
}
