package tipoff

import (
	tipoffdao "best_wish/app/dao/tipoff"
	tipoffmodel "best_wish/app/models/tipoff"
	"best_wish/until"
	"encoding/json"
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
	"github.com/goravel/framework/facades"
	"io"
	"math/rand"
	"net/http"
)

type Reptile struct {
}

type ReptileUrlData struct {
	ID            any    `json:"id,omitempty"`
	Title         string `json:"title,omitempty"`
	Content       string `json:"content,omitempty"`
	DateTime      string `json:"datetime,omitempty"`
	ShortTime     string `json:"shorttime,omitempty"`
	ShiJianChuo   any    `json:"shijianchuo,omitempty"`
	CateId        string `json:"cateid,omitempty"`
	CateName      string `json:"catename,omitempty"`
	Comments      any    `json:"comments,omitempty"`
	LouZhu        string `json:"louzhu,omitempty"`
	LouZhuRegTime string `json:"louzhuregtime,omitempty"`
	Url           string `json:"url,omitempty"`
	YuanUrl       string `json:"yuanurl,omitempty"`
}
type HotReptileUrlData struct {
	Remen6  []ReptileUrlData `json:"remen6,omitempty"`
	Remen24 []ReptileUrlData `json:"remen24,omitempty"`
	Remen48 []ReptileUrlData `json:"remen48,omitempty"`
}

// Signature The name and signature of the console command.
func (receiver *Reptile) Signature() string {
	return "tip:off:reptile"
}

// Description The console command description.
func (receiver *Reptile) Description() string {
	return "线报数据爬虫"
}

// Extend The console command extend.
func (receiver *Reptile) Extend() command.Extend {
	return command.Extend{}
}

// Handle Execute the console command.
func (receiver *Reptile) Handle(ctx console.Context) error {
	facades.Log().Info("开始爬取线报数据")
	list, err := tipoffdao.GetNeedReptileConfigList()

	if err != nil {
		facades.Log().Errorf("查询线报出错%s", err)
		return err
	}

	for _, item := range list {
		facades.Log().Infof("请求%s", item.Url)
		sendRequestToReptile(item.Url, item.ID)
		tipoffdao.UpdateNextTime(item.ID, item.IntervalTime+uint64(rand.Intn(120)))
	}
	facades.Log().Info("结束爬取线报数据")
	return nil
}

func sendRequestToReptile(url string, configId uint) {
	resp, err := http.Get(url)
	if err != nil {
		facades.Log().Errorf("请求错误%s", err)
		return
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var responseData []ReptileUrlData
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		facades.Log().Info("json解析错误,尝试解析为热门数据")

		var hotData HotReptileUrlData
		err = json.Unmarshal(body, &hotData)
		if err != nil {
			facades.Log().Info("二次解析数据错误")
			return
		}
		for _, item6 := range hotData.Remen6 {
			responseData = append(responseData, item6)
		}
		for _, item24 := range hotData.Remen24 {
			responseData = append(responseData, item24)
		}
		for _, item48 := range hotData.Remen48 {
			responseData = append(responseData, item48)
		}
	}

	for _, item := range responseData {
		tipoffdao.FirstOrCreateByOrigId(tipoffmodel.TipOffNoticeData{
			OrigId:        until.ToBeInt64(item.ID),
			Title:         item.Title,
			Content:       item.Content,
			DateTime:      item.DateTime,
			ShortTime:     item.ShortTime,
			ShiJianChuo:   until.ToBeInt64(item.ShiJianChuo),
			CateId:        item.CateId,
			CateName:      item.CateName,
			Comments:      until.ToBeInt64(item.Comments),
			LouZhu:        item.LouZhu,
			LouZhuRegTime: item.LouZhuRegTime,
			Url:           item.Url,
			PachongId:     int64(configId),
			YuanUrl:       item.YuanUrl,
			IsNotice:      0,
		})
	}
}
