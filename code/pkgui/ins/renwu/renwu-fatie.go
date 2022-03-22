package renwu

import (
	"2022Instagram-Qunkong/code/pkgui/ins"
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"os"
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
		//寻找发动态按钮
		chromedp.Click("#react-root > section > nav > div._8MQSO.Cx7Bp > div > div > div.ctQZg.KtFt3 > div > div:nth-child(3) > div > button > div > svg > line:nth-child(3)", chromedp.ByID).Do(ctx)
		//获取本机路径
		wd, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		filepath := []string{wd + `/data/jpg/03.jpg`}
		fmt.Println("正在选择图片2", filepath)
		//上传图片(总结：WaitVisible是界面看到的东西)
		chromedp.WaitVisible(`body > div.RnEpo.gpWnf.Yx5HN > div.pbNvD > div > div > div > div.uYzeu > div._C8iK > div > div > div.qF0y9.Igw0E.rBNOH.eGOV_.ybXk5._4EzTm.kEKum > div > button`, chromedp.ByQuery).Do(ctx)
		chromedp.SetUploadFiles(`body > div.RnEpo.gpWnf.Yx5HN > div.pbNvD > div > div > div > div.uYzeu > div._C8iK > form > input`, filepath, chromedp.ByQuery).Do(ctx)
		fmt.Println("正在选择图片3")
		chromedp.WaitEnabled(`body > div.RnEpo.gpWnf.Yx5HN > div.pbNvD > div > div > div > div.uYzeu > div._C8iK > form > input`, chromedp.ByQuery).Do(ctx)
		fmt.Println("上传已经消失")
		//点击（原图）缩放图片
		chromedp.WaitVisible(`body > div.RnEpo.gpWnf.Yx5HN > div.pbNvD > div > div > div > div.uYzeu > div._C8iK > div > div > div > div.qF0y9.Igw0E.IwRSH.eGOV_._4EzTm.bkEs3.soMvl.JI_ht.DhRcB.O1flK.D8xaz.fm1AK > div > div:nth-child(2) > div > button`, chromedp.ByQuery).Do(ctx)
		chromedp.Click(`body > div.RnEpo.gpWnf.Yx5HN > div.pbNvD > div > div > div > div.uYzeu > div._C8iK > div > div > div > div.qF0y9.Igw0E.IwRSH.eGOV_._4EzTm.bkEs3.soMvl.JI_ht.DhRcB.O1flK.D8xaz.fm1AK > div > div:nth-child(2) > div > button`, chromedp.ByQuery).Do(ctx)
		chromedp.Click(`body > div.RnEpo.gpWnf.Yx5HN > div.pbNvD > div > div > div > div.uYzeu > div._C8iK > div > div > div > div.qF0y9.Igw0E.IwRSH.eGOV_._4EzTm.bkEs3.soMvl.JI_ht.DhRcB.O1flK.D8xaz.fm1AK > div > div.qF0y9.Igw0E.IwRSH.eGOV_._4EzTm.lC6p0.HVWg4 > div > button:nth-child(1)`, chromedp.ByQuery).Do(ctx)
		fmt.Println("已经完成点击1")
		//点击继续按钮
		chromedp.WaitVisible(`body > div.RnEpo.gpWnf.Yx5HN > div.pbNvD > div > div > div > div.qF0y9.Igw0E.IwRSH.eGOV_._4EzTm > div > div > div.qF0y9.Igw0E.rBNOH.YBx95._4EzTm.fm1AK > h1 > div`, chromedp.ByQuery).Do(ctx)
		chromedp.Click(`body > div.RnEpo.gpWnf.Yx5HN > div.pbNvD > div > div > div > div.qF0y9.Igw0E.IwRSH.eGOV_._4EzTm > div > div > div.WaOAr._8E02J > div > button`, chromedp.ByQueryAll).Do(ctx)
		fmt.Println("已经完成点击2")
		//点击原版图像
		chromedp.Click(`body > div.RnEpo.gpWnf.Yx5HN > div.pbNvD > div > div > div > div.uYzeu.gIMwG > div._83r9B > div > div > div > div.qF0y9.Igw0E.IwRSH.eGOV_.vwCYk.lDRO1 > div > div:nth-child(1) > button`, chromedp.ByQuery).Do(ctx)
		//点击继续按钮
		chromedp.Click(`body > div.RnEpo.gpWnf.Yx5HN > div.pbNvD > div > div > div > div.qF0y9.Igw0E.IwRSH.eGOV_._4EzTm > div > div > div.WaOAr._8E02J > div > button`, chromedp.ByQuery).Do(ctx)
		fmt.Println("已经完成所有点击")
		chromedp.SendKeys(`body > div.RnEpo.gpWnf.Yx5HN > div.pbNvD > div > div > div > div.uYzeu.gIMwG > div._83r9B > div > div > div > div:nth-child(2) > div.qF0y9.Igw0E.IwRSH.eGOV_._4EzTm > textarea`,
			ins.Tiezi[ins.GetRandNum()], chromedp.ByQuery).Do(ctx)

		fmt.Println("输入完成")
		//点击分享按钮
		//chromedp.Click(`body > div.RnEpo.gpWnf.Yx5HN > div.pbNvD > div > div > div > div.qF0y9.Igw0E.IwRSH.eGOV_._4EzTm > div > div > div.WaOAr._8E02J > div > button`, chromedp.ByQuery).Do(ctx)

		//等待完成
		//chromedp.WaitVisible(`body > div.RnEpo.gpWnf.Yx5HN > div.pbNvD > div > div > div > div.uYzeu > div._C8iK > div > div > div > h2`, chromedp.ByQuery).Do(ctx)
		//点击关闭
		//chromedp.Click(`body > div.RnEpo.gpWnf.Yx5HN > div.NOTWr > button`, chromedp.ByQuery).Do(ctx)
		return
	}
}
