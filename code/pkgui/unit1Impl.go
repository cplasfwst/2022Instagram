package pkgui

import (
	"2022Instagram-Qunkong/code/pkgui/ins"
	"fmt"
	"github.com/ying32/govcl/vcl"
	"strconv"
)

//::private::
type TForm1Fields struct {
}

func (f *TForm1) getGlobal() {
	//账号1的关键词
	if f.Guanjianci.Text() != "" {
		ins.Guanjianci = f.Guanjianci.Text()
	} else {
		fmt.Println("关键词不能为空")
	}

	//代理配置读取
	if f.HostDaili.Text() != "" {
		ins.HostDaili = f.HostDaili.Text()
	} else {
		fmt.Println("代理不能为空")
	}
	if f.DLzhanghao.Text() != "" {
		ins.DLzhanghao = f.DLzhanghao.Text()
	} else {
		fmt.Println("代理账号不能为空")
	}
	if f.DLmima.Text() != "" {
		ins.DLmima = f.DLmima.Text()
	} else {
		fmt.Println("代理密码不能为空")
	}

	//话术配置
	if f.Huashu.Text() != "" {
		ins.Huashu = f.Huashu.Text()
	} else {
		fmt.Println("话术不能为空")
	}
	if f.PinglunCD.Text() != "" {
		atoi, err := strconv.Atoi(f.PinglunCD.Text())
		ins.PinglunCD = atoi
		if err != nil {
			fmt.Println("转换延迟出错")
		}
	} else {
		fmt.Println("延迟不能为空")
	}
}

func (f *TForm1) OnButton1Click(sender vcl.IObject) {
	f.getGlobal()

	ins.InsInit()
}
