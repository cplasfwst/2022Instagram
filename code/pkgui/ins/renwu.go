package ins

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
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
		chromedp.SendKeys("#react-root > section > nav > div._8MQSO.Cx7Bp > div > div > div.QY4Ed.P0xOK > input", Guanjianci, chromedp.ByID).Do(ctx)

		chromedp.Click("#react-root > section > nav > div._8MQSO.Cx7Bp > div > div > div.QY4Ed.P0xOK > div.yPP5B > div > div._01UL2 > div > div:nth-child(1) > a", chromedp.ByID).Do(ctx)

		chromedp.WaitVisible("#react-root > section > main > article > div.EZdmt > div > div > div:nth-child(1) > div:nth-child(1)", chromedp.ByID).Do(ctx)

		chromedp.Click("#react-root > section > main > article > div.EZdmt > div > div > div:nth-child(1) > div:nth-child(2)", chromedp.ByID).Do(ctx)

		for true {
			CountTime(PinglunCD)
			//time.Sleep(time.Minute * time.Duration(pkgui.PinglunCD))
			fmt.Println("进行点击")
			//点击下一个
			chromedp.Click("body > div.RnEpo._Yhr4 > div.Z2Inc._7c9RR > div > div.l8mY4.feth3 > button > div > span > svg", chromedp.ByQuery).Do(ctx)

			pinglun().Do(ctx)
		}

		return
	}
}

func pinglun() chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {

		chromedp.SendKeys("body > div.RnEpo._Yhr4 > div.pbNvD.QZZGH.bW6vo > div > article > div > div.HP0qD > div > div > div.eo2As > section.sH9wk._JgwE > div > form > textarea", Huashu, chromedp.ByQuery).Do(ctx)

		chromedp.Click("body > div.RnEpo._Yhr4 > div.pbNvD.QZZGH.bW6vo > div > article > div > div.HP0qD > div > div > div.eo2As > section.sH9wk._JgwE > div > form > button", chromedp.ByQuery).Do(ctx)

		fmt.Println("评论完成")

		return
	}
}
