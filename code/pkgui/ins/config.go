package ins

import (
	"context"
	"github.com/chromedp/cdproto/fetch"
	"github.com/chromedp/chromedp"
	"log"
)

func InsInit() {
	ctx, _ := chromedp.NewExecAllocator(
		context.Background(),

		// 以默认配置的数组为基础，覆写headless参数
		// 当然也可以根据自己的需要进行修改，这个flag是浏览器的设置
		append(
			chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless", false),
			chromedp.ProxyServer("gate1.rola.info:1009"),
			chromedp.Flag("proxy-bypass-list", "<-loopback>"),
			//chromedp.Flag("disable-web-security", true),
			//本机：Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36
			//chromedp.UserAgent("Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Mobile Safari/537.36"),
			chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36"),
		)...,
	)

	ctx, _ = chromedp.NewContext(
		ctx,
		// 设置日志方法
		chromedp.WithLogf(log.Printf),
	)

	//设置代理
	lctx, lcancel := context.WithCancel(ctx)
	chromedp.ListenTarget(lctx, func(ev interface{}) {
		switch ev := ev.(type) {
		case *fetch.EventRequestPaused:
			go func() {
				_ = chromedp.Run(ctx, fetch.ContinueRequest(ev.RequestID))
			}()
		case *fetch.EventAuthRequired:
			if ev.AuthChallenge.Source == fetch.AuthChallengeSourceProxy {
				go func() {
					_ = chromedp.Run(ctx,
						fetch.ContinueWithAuth(ev.RequestID, &fetch.AuthChallengeResponse{
							Response: fetch.AuthChallengeResponseResponseProvideCredentials,
							Username: "cplasfwst_dc_1",
							Password: "xz5201314789",
						}),
						// Chrome will remember the credential for the current instance,
						// so we can disable the fetch domain once credential is provided.
						// Please file an issue if Chrome does not work in this way.
						fetch.Disable(),
					)
					// and cancel the event handler too.
					lcancel()
				}()
			}
		}
	})

	MyTest(ctx)
}
