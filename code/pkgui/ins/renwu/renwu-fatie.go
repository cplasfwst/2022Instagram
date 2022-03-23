package renwu

import (
	"2022Instagram-Qunkong/code/pkgui/ins"
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"os"
	"strconv"
	"time"
)

func Renwu_Fatie() chromedp.Tasks {
	fatie := chromedp.Tasks{
		//第一步
		Fa_First(),
	}
	return fatie
}

func Fa_First() chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		//寻找发动态按钮#react-root > section > nav > div._8MQSO.Cx7Bp > div > div > div.ctQZg.KtFt3 > div > div:nth-child(3)
		chromedp.Click(`div[class="J5g42"] > div:nth-child(3)`, chromedp.ByQuery).Do(ctx)
		//获取本机路径
		wd, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		//itoa将INT转string，获取图片文件的总数量，然后获取一个随机数，对应图片名字
		filepath := []string{wd + `/data/jpg/` + strconv.Itoa(ins.GetRandNum(ins.PictureCount)) + ".jpg"}
		fmt.Println("正在选择图片2", filepath)
		//上传图片(总结：WaitVisible是界面看到的东西)
		chromedp.WaitVisible(`h2[class="_7UhW9      x-6xq  yUEEX    KV-D4          uL8Hv     l4b0S    "]`, chromedp.NodeVisible).Do(ctx)
		//注意，这里上传图片一定要Byquery
		chromedp.SetUploadFiles(`input[accept="image/jpeg,image/png,image/heic,image/heif,video/mp4,video/quicktime"]`, filepath, chromedp.ByQuery).Do(ctx)
		fmt.Println("正在选择图片3")
		//等待【將相片和影片拖曳到這裡】这几个字消失
		chromedp.WaitEnabled(`h2[class="_7UhW9      x-6xq  yUEEX    KV-D4          uL8Hv     l4b0S    `, chromedp.ByQuery).Do(ctx)
		fmt.Println("上传已经消失")
		time.Sleep(time.Second * 1)
		//点击缩放图片(不缩放会造成继续无法点击) !!!注意，点击事件要配合chromedp.ByQuery，不然无法点击
		chromedp.WaitVisible(`div.qF0y9.Igw0E.IwRSH.eGOV_._4EzTm.bkEs3.soMvl.JI_ht.DhRcB.O1flK.D8xaz.fm1AK`,
			chromedp.NodeVisible).Do(ctx)
		chromedp.Click(`div.qF0y9.Igw0E.IwRSH.eGOV_._4EzTm.bkEs3.soMvl.JI_ht.DhRcB.O1flK.D8xaz.fm1AK > div > div:nth-child(2) > div > button`,
			chromedp.ByQuery).Do(ctx)
		chromedp.Click(`div[class="YAPUk  gdFG_"] > button:nth-child(3)`, chromedp.ByQuery).Do(ctx)
		fmt.Println("已经点击完缩放图片")
		//点击继续(不等待按钮出现会造成继续无法点击)
		//总结，等待的sel不能和点击的sel一样，不然卡住了
		chromedp.WaitVisible(`div[class="WaOAr _8E02J"]`, chromedp.NodeVisible).Do(ctx)
		chromedp.Click(`div[class="WaOAr _8E02J"] > div > button`, chromedp.ByQueryAll).Do(ctx)
		fmt.Println("已经点击继续")
		//选择渲染方式
		//暂时留空用原图

		//点击继续(不等待按钮出现会造成继续无法点击)
		chromedp.WaitVisible(`img[class="sQuyK atPR5"]`, chromedp.NodeVisible).Do(ctx)
		fmt.Println("已经完成等待22222222222")
		chromedp.Click(`div[class="WaOAr _8E02J"] > div > button`, chromedp.ByQueryAll).Do(ctx)
		fmt.Println("已经点击继续2222222222222222222")
		//输入文本
		chromedp.SendKeys(`textarea[class="PUqUI lFzco"]`, ins.Tiezi_huashu[ins.GetRandNum(len(ins.Tiezi_huashu))], chromedp.ByQuery).Do(ctx)
		//等待定位条出现点击分享按钮
		chromedp.WaitVisible(`input[name="creation-location-input"]`, chromedp.NodeVisible).Do(ctx)
		//chromedp.Click(`div.WaOAr._8E02J > div > button`, chromedp.ByQueryAll).Do(ctx)
		fmt.Println("已经点击分享按钮")
		//等待分享完成
		chromedp.WaitVisible(`h2[class="_7UhW9      x-6xq  yUEEX    KV-D4          uL8Hv     l4b0S    "]`, chromedp.NodeVisible).Do(ctx)
		chromedp.Click(`div.NOTWr > button`, chromedp.ByQuery).Do(ctx)
		fmt.Println("完成一轮发帖，准备等待任务")

		return
	}
}
