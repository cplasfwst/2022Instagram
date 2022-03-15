package ins

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"time"
)

func Renwu_Sou_PL() chromedp.Tasks {
	pinglun := chromedp.Tasks{

		//1,寻找关于足球的话题
		First(),

		//2,开始评论

	}

	return pinglun
}

func First() chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		chromedp.SendKeys("#react-root > section > nav > div._8MQSO.Cx7Bp > div > div > div.QY4Ed.P0xOK > input", "足球", chromedp.ByID).Do(ctx)

		chromedp.Click("#react-root > section > nav > div._8MQSO.Cx7Bp > div > div > div.QY4Ed.P0xOK > div.yPP5B > div > div._01UL2 > div > div:nth-child(1) > a", chromedp.ByID).Do(ctx)

		chromedp.WaitVisible("#react-root > section > main > article > div.EZdmt > div > div > div:nth-child(1) > div:nth-child(1)", chromedp.ByID).Do(ctx)

		chromedp.Click("#react-root > section > main > article > div.EZdmt > div > div > div:nth-child(1) > div:nth-child(1)", chromedp.ByID).Do(ctx)

		for true {
			time.Sleep(time.Second * 3600)
			fmt.Println("进行点击")
			chromedp.Click("body > div.RnEpo._Yhr4 > div.Z2Inc._7c9RR > div > div > button", chromedp.ByQuery).Do(ctx)

			//pinglun().Do(ctx)
		}

		return
	}
}

func pinglun() chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {

		chromedp.SendKeys("body > div.RnEpo._Yhr4 > div.pbNvD.QZZGH.bW6vo > div > article > div > div.HP0qD > div > div > div.eo2As > section.sH9wk._JgwE > div > form > textarea", "新group:足球交友交流，足球神秘貼士，內幕球賽資料:xdoctorwho.com", chromedp.ByQuery).Do(ctx)

		chromedp.Click("body > div.RnEpo._Yhr4 > div.pbNvD.QZZGH.bW6vo > div > article > div > div.HP0qD > div > div > div.eo2As > section.sH9wk._JgwE > div > form > button", chromedp.ByQuery).Do(ctx)

		fmt.Println("评论完成")

		return
	}
}
