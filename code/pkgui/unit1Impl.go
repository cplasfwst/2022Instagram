package pkgui

import (
	"2022Instagram-Qunkong/code/pkgui/ins"
	"2022Instagram-Qunkong/code/pkgui/ins/renwu"
	"fmt"
	"github.com/ying32/govcl/vcl"
)

//::private::
type TForm1Fields struct {
}

func (f *TForm1) getGlobal() {

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

	//随机帖子话术
	path := f.Tiezi_huashu.Text()
	ins.Tiezi = ins.ReadTiezi(path)
}

func (f *TForm1) OnButton1Click(sender vcl.IObject) {
	f.getGlobal()

	//预先加载账号密码文本
	ins.UserData, _ = ins.ImportuserMap("users.txt")
	fmt.Println(len(ins.UserData))
	fmt.Println(ins.UserData)

	go renwu.InsInit(ins.UserData[0])
}

func (f *TForm1) OnHostDailiChange(sender vcl.IObject) {

}

//测试读取富文本
func (f *TForm1) OnButton2Click(sender vcl.IObject) {
	path := f.Tiezi_huashu.Text()
	ins.Tiezi = ins.ReadTiezi(path)

	fmt.Println(ins.Tiezi[ins.GetRandNum()])
}
