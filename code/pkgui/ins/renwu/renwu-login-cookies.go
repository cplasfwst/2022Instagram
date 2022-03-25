package renwu

import (
	"2022Instagram-Qunkong/code/pkgui/ins"
	"context"
	"fmt"
	"github.com/chromedp/cdproto/browser"
	"github.com/chromedp/cdproto/fetch"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

//任务入口方法有两个：1：是已经保存cookies，2：是保存cookies后准备开始任务
//1,checkLoginStatus
//2,saveCookies

func Login_cookies(ctx context.Context, data sync.Map) {

	err := chromedp.Run(ctx, loginIns(data))
	if err != nil {
		fmt.Println("登陆错误", err)
	}

	fmt.Println("登陆CTX结束")

}

func loginIns(data sync.Map) chromedp.Tasks {
	login := chromedp.Tasks{
		//加载代理
		fetch.Enable().WithHandleAuthRequests(true),
		//加载通知权限
		browser.GrantPermissions([]browser.PermissionType{
			browser.PermissionTypeNotifications}),

		//0,加载cookies
		loadCookies(data),

		//1,打开INS
		chromedp.Navigate("https://www.instagram.com/accounts/login/"),

		//2,检查是否登陆
		checkLoginStatus(data),

		//3,如是没有登陆就输入账号密码并且登陆
		chromedp.SendKeys("#loginForm > div > div:nth-child(1) > div > label > input", ins.MapRead(data, "INSzhanghao"), chromedp.ByID),
		chromedp.SendKeys("#loginForm > div > div:nth-child(2) > div > label > input", ins.MapRead(data, "INSmima"), chromedp.ByID),
		chromedp.Click("#loginForm > div > div:nth-child(3) > button", chromedp.ByID),

		//点击保存cookies
		//chromedp.Click(`#react-root > section > main > div > div > div > section > div > button`, chromedp.ByID),
		//chromedp.Click(`#react-root > div > div > section > main > div > div > div > section > div > button`, chromedp.ByID),
		//chromedp.Click(`.sqdOP  L3NKy   y3zKF     `, chromedp.NodeVisible),
		//点击保存
		chromedp.Click(`button[class="sqdOP  L3NKy   y3zKF     "]`, chromedp.ByQuery),
		//等待进入主页
		//等待进入主页
		//chromedp.WaitVisible(`a[class="gmFkV"]`,chromedp.ByQuery).Do(ctx)
		//允许通知权限这个位置导致读取cookies出错
		//browser.GrantPermissions([]browser.PermissionType{
		//	browser.PermissionTypeNotifications}),

		//打开通知
		//chromedp.Click(`body > div.RnEpo.Yx5HN > div > div > div > div.mt3GC > button.aOOlW.bIiDR`,chromedp.ByQuery),

		//保存登陆信息
		saveCookies(data),
	}

	return login
}

// 保存Cookies
func saveCookies(data sync.Map) chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		//var gebierenwu int
		fmt.Println("进来了存储cookies")
		//读取临时cookies自加
		wd, err := os.Getwd()
		//需要等待页面加载完成
		time.Sleep(time.Second * 3)

		// 等待二维码登陆
		//#react-root > div > div > section > nav > div._8MQSO.Cx7Bp > div > div > div.ctQZg.KtFt3 > div > div:nth-child(1) > div > a > svg
		if err = chromedp.WaitVisible(`div[class="J5g42"]`, chromedp.NodeVisible).Do(ctx); err != nil {
			fmt.Println("等待二维码这里出错（保存cookies的前置条件）")
			return
		}
		fmt.Println("检测到了主页按钮")
		// cookies的获取对应是在devTools的network面板中
		// 1. 获取cookies
		cookies, err := network.GetAllCookies().Do(ctx)
		if err != nil {
			return
		}

		// 2. 序列化
		cookiesData, err := network.GetAllCookiesReturns{Cookies: cookies}.MarshalJSON()
		if err != nil {
			return
		}

		// 3. 存储到临时文件
		if err = ioutil.WriteFile(wd+"/data/cookies/"+ins.MapRead(data, "cookies"), cookiesData, 0755); err != nil {
			return
		}
		fmt.Println("保存cookies成功")

		// 等待元素加载完成
		//chromedp.WaitVisible("#react-root", chromedp.NodeVisible).Do(ctx)

		//开始启动循环任务
		renwushu := 0
		for true {
			err := chromedp.Run(ctx, Renwu_Fatie(data))
			if err != nil {
				fmt.Println("check检查完是否有cookies评论错误", err)
			}
			renwushu++
			ins.CountTime(ins.Insyanchi, data, renwushu)

		}

		return
	}
}

// 加载Cookies
func loadCookies(data sync.Map) chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		//读取临时cookies自加
		wd, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		// 如果cookies临时文件不存在则直接跳过
		if _, _err := os.Stat(wd + "/data/cookies/" + ins.MapRead(data, "cookies")); os.IsNotExist(_err) {
			fmt.Println("不存在缓存文件")
			return
		}

		// 如果存在则读取cookies的数据
		cookiesData, err := ioutil.ReadFile(wd + "/data/cookies/" + ins.MapRead(data, "cookies"))
		if err != nil {
			fmt.Println("读取缓存文件是吧")
			return
		}

		// 反序列化
		cookiesParams := network.SetCookiesParams{}
		if err = cookiesParams.UnmarshalJSON(cookiesData); err != nil {
			fmt.Println("反序列化失败")
			return
		}
		//data["INSzhuangtai"] = "读取cookies成功"
		data.Store("INSzhuangtai", "读取cookies成功")
		fmt.Println("读取cookies成功")
		// 设置cookies
		return network.SetCookies(cookiesParams.Cookies).Do(ctx)
	}
}

func checkLoginStatus(data sync.Map) chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		var url string
		if err = chromedp.Evaluate(`window.location.href`, &url).Do(ctx); err != nil {
			return
		}
		//利用网址判断
		//if strings.Contains(url, "https://www.instagram.com/accounts/onetap/?next=%2F") {
		//	log.Println("已经使用cookies登陆lele")
		//	chromedp.Stop()
		//}
		//利用网址等值判断
		//fmt.Println(url)
		if strings.EqualFold(url, "https://www.instagram.com/") {
			log.Println("已经使用cookies登陆lele")
			//data["INSzhuangtai"] = "已经使用cookies登陆"
			data.Store("INSzhuangtai", "已经使用cookies登陆")
			//开始循环任务，先定义一个账号个人完成任务数
			renwushu := 0
			for true {
				err := chromedp.Run(ctx, Renwu_Fatie(data))
				if err != nil {
					fmt.Println("check检查完是否有cookies评论错误", err)
				}
				renwushu++
				ins.CountTime(ins.Insyanchi, data, renwushu)

			}
		}

		return
	}
}
