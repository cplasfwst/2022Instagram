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
	"strconv"
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

		//1.1检查是否违规了
		CheckWeigui(data),

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
		//因为账号密码登陆，再次检测是否违规，因为不存在保存按钮
		//CheckWeigui(data),
		//点击保存
		chromedp.Click(`button[class="sqdOP  L3NKy   y3zKF     "]`, chromedp.ByQuery),
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
			//随机延迟一下账号
			time.Sleep(time.Second * time.Duration(ins.GetRandNum(10)))
			err := chromedp.Run(ctx, Renwu_Fatie(data))
			if err != nil {
				fmt.Println("check检查完是否有cookies评论错误", err)
			}
			renwushu++
			yanchi := ins.GetRandLMNum(ins.InsyanchiMin, ins.InsyanchiMax)
			ins.CountTime(yanchi, data, renwushu)
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
			data.Store("INSzhuangtai", "不存在缓存文件")
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
		log.Println("进来了检查登陆")
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
		//添加违规检测在这里试试

		if strings.EqualFold(url, "https://www.instagram.com/") {
			log.Println("已经使用cookies登陆lele")
			//data["INSzhuangtai"] = "已经使用cookies登陆"
			data.Store("INSzhuangtai", "已经使用cookies登陆")
			//开始循环任务，先定义一个账号个人完成任务数
			renwushu := 0
			for true {
				//随机延迟一下账号
				data.Store("INSzhuangtai", "准备开始第二轮任务")
				//修改为err2看看是否卡住在完成一轮任务
				err2 := chromedp.Run(ctx, Renwu_Fatie(data))
				if err2 != nil {
					fmt.Println("check检查完是否有cookies评论错误", err)
				}
				renwushu++
				yanchi := ins.GetRandLMNum(ins.InsyanchiMin, ins.InsyanchiMax)
				ins.CountTime(yanchi, data, renwushu)
			}
		}
		return
	}
}

//检查是否违规了
// 加载Cookies
func CheckWeigui(data sync.Map) chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		log.Println(ins.MapRead(data, "INSzhanghao"), "进来了违规提示")
		//chromedp.WaitVisible(`<body>`,chromedp.NodeVisible).Do(ctx)  注意NodeVisible有问题单body不行要加<>，可能要全名
		//time.Sleep(time.Second * 3)
		//fmt.Println("等待完成")
		//time.Sleep(time.Second * 15)
		//chromedp.WaitNotPresent()
		//处理违反事件开始-------------------------------------------------------------------------------------------------
		//判断是否提示违规
		var url string
		if err = chromedp.Evaluate(`window.location.href`, &url).Do(ctx); err != nil {
			log.Println("出错了")
			return
		}
		//fmt.Println(url)https://www.instagram.com/challenge/?next=/
		if strings.EqualFold(url, "https://www.instagram.com/challenge/?next=/accounts/login/") || strings.EqualFold(url, "https://www.instagram.com/challenge/?next=/") {
			//data["INSzhuangtai"] = "已经使用cookies登陆"
			data.Store("INSzhuangtai", "提示违规了，准备点击")
			//放在这里截图测试一下
			Capsrc(data).Do(ctx)
			read := ins.MapRead(data, "INSweigui")
			atoi, errwg := strconv.Atoi(read)
			if errwg != nil {
				data.Store("INSzhuangtai", "违规转化失败")
			}
			weigui := atoi + 1
			data.Store("INSweigui", weigui)
			chromedp.Click(`button[class="sqdOP  L3NKy   y3zKF   cB_4K  "]`, chromedp.ByQuery).Do(ctx)
			time.Sleep(time.Second * 3)
		}
		//处理违规事件结束----------------------------------------------------------------------------------------------
		//上面的处理比较糟糕，未发现应该如何处理，先加一个卡住发帖的方法:strings.Contains ,Contains是URL里面如果包含有第二个参数 则返回真1
		if strings.Contains(url, "https://www.instagram.com/challenge/?next=/accounts/onetap/") {
			//log.Println("提示违规")
			//data["INSzhuangtai"] = "已经使用cookies登陆"
			//放在这里截图测试一下
			Capsrc(data).Do(ctx)
			data.Store("INSzhuangtai", "发帖违规了2，准备点击")
			read2 := ins.MapRead(data, "INSweigui")
			atoi, errwg := strconv.Atoi(read2)
			if errwg != nil {
				data.Store("INSzhuangtai", "违规转化失败")
			}
			weigui := atoi + 1
			data.Store("INSweigui", weigui)
			chromedp.Click(`button[class="sqdOP  L3NKy   y3zKF   cB_4K  "]`, chromedp.ByQuery).Do(ctx)
			time.Sleep(time.Second * 3)
		}

		return
	}
}

//整个网页截图
func Capsrc(data sync.Map) chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		//先创建一个文件缓存
		var b1 []byte
		wd, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}

		//截图开始
		chromedp.CaptureScreenshot(&b1).Do(ctx)

		//获取当前账号的名字
		insName := ins.MapRead(data, "INSzhanghao")
		myPng := wd + "/data/capImg/" + insName + ".png"

		if err2 := ioutil.WriteFile(myPng, b1, 0777); err != nil {
			log.Println(err2)
		}

		return
	}
}
