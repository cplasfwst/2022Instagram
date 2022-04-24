package renwu

import (
	"2022Instagram-Qunkong/code/pkgui/ins"
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"sync"
	"time"
)

//养号功能

func Renwu_Yanghao(data sync.Map) chromedp.Tasks {
	fatie := chromedp.Tasks{
		//第一步
		Ya_First(data),
	}
	return fatie
}

func Ya_First(data sync.Map) chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		log.Println(ins.MapRead(data, "INSzhanghao"), "进来了养号")
		data.Store("INSzhuangtai", "进来了养号")
		//先去主页
		chromedp.Navigate("https://www.instagram.com/explore/").Do(ctx)
		//检查是否违规
		//BUG处理！！！！！！！！！！chromedp.ActionFunc里面不能运行chromedp.ActionFunc！！！！！！！（试试加个.do(ctx)！！）
		CheckWeigui(data).Do(ctx)
		//点击第二张图片
		chromedp.Click(`#react-root > section > main > div > div.K6yM_ > div > div:nth-child(1) > div:nth-child(3)`, chromedp.ByQuery).Do(ctx)
		//等待一定的时间后

		//点击下一步
		xiayibu := 0
		for true {
			time.Sleep(time.Second * 5)
			//点击下一步按钮
			chromedp.Click(`body > div.RnEpo._Yhr4 > div.Z2Inc._7c9RR > div > div.l8mY4.feth3 > button`, chromedp.ByQuery).Do(ctx)
			xiayibu++
			if xiayibu == 3 {
				fmt.Println("足够3次了，退出")
				chromedp.Click(`body > div.RnEpo._Yhr4 > div.NOTWr > button`, chromedp.ByQuery).Do(ctx)
				break
			}
		}

		fmt.Println("退出养号")

		return
	}
}
