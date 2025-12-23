package tipoff

import (
	tipoffdao "best_wish/app/dao/tipoff"
	"best_wish/lib/wxpusher"
	"best_wish/until"
	"fmt"
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
	"github.com/goravel/framework/facades"
	"net/http"
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

	for _, tipoff := range list {
		url, isOk := checkUrl(tipoff.Url)
		if !isOk {
			continue
		}

		hour := time.Now().Hour()
		if hour >= 2 && hour <= 6 {
			// 凌晨2点到6点之间不发送
			continue
		}

		// 使用 Markdown 格式构建消息内容
		content := fmt.Sprintf("### %s\n%s\n[🔗查看详情](%s)", tipoff.Title, tipoff.Content, url)

		// 发送到企业微信
		isNotice := wxpusher.SendWorkWechat(content)

		if isNotice {
			// 更新单条记录为已通知
			tipoffdao.UpdateIsNotice([]uint{tipoff.ID})
			// 每条消息之间间隔1秒，避免频率限制
			time.Sleep(time.Second)
		}
	}

	return nil
}
