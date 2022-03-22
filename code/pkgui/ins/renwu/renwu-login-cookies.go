package renwu

import (
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
)

func Login_cookies(ctx context.Context, data map[string]string) {

	err := chromedp.Run(ctx, loginIns(data))
	if err != nil {
		fmt.Println("登陆错误", err)
	}

	fmt.Println("任務完成了")

}

func loginIns(data map[string]string) chromedp.Tasks {
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
		chromedp.SendKeys("#loginForm > div > div:nth-child(1) > div > label > input", data["INSzhanghao"], chromedp.ByID),
		chromedp.SendKeys("#loginForm > div > div:nth-child(2) > div > label > input", data["INSmima"], chromedp.ByID),
		chromedp.Click("#loginForm > div > div:nth-child(3) > button", chromedp.ByID),

		//点击保存cookies
		chromedp.Click(`#react-root > section > main > div > div > div > section > div > button`, chromedp.ByID),

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
func saveCookies(data map[string]string) chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		fmt.Println("进来了存储cookies")
		//读取临时cookies自加
		wd, err := os.Getwd()
		// 等待二维码登陆
		if err = chromedp.WaitVisible(`#react-root > section > nav > div._8MQSO.Cx7Bp > div > div > div.ctQZg.KtFt3 > div > div:nth-child(1) > div > a > svg`, chromedp.ByID).Do(ctx); err != nil {
			fmt.Println("等待二维码这里出错（保存cookies的前置条件）")
			return
		}

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
		if err = ioutil.WriteFile(wd+"/data/cookies/"+data["cookies"], cookiesData, 0755); err != nil {
			return
		}
		fmt.Println("保存cookies成功")

		err = chromedp.Run(ctx, Renwu_Fatie())
		if err != nil {
			fmt.Println("存cookies后评论错误", err)
		}
		return
	}
}

// 加载Cookies
func loadCookies(data map[string]string) chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		//读取临时cookies自加
		wd, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		// 如果cookies临时文件不存在则直接跳过
		if _, _err := os.Stat(wd + "/data/cookies/" + data["cookies"]); os.IsNotExist(_err) {
			fmt.Println("不存在缓存文件")
			return
		}

		// 如果存在则读取cookies的数据
		cookiesData, err := ioutil.ReadFile(wd + "/data/cookies/" + data["cookies"])
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
		data["INSzhuangtai"] = "读取cookies成功"
		fmt.Println("读取cookies成功")
		// 设置cookies
		return network.SetCookies(cookiesParams.Cookies).Do(ctx)
	}
}

func checkLoginStatus(data map[string]string) chromedp.ActionFunc {
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
			data["INSzhuangtai"] = "已经使用cookies登陆"
			//chromedp.Stop()
			err := chromedp.Run(ctx, Renwu_Fatie())
			if err != nil {
				fmt.Println("check检查完是否有cookies评论错误", err)
			}
		}

		return
	}
}
