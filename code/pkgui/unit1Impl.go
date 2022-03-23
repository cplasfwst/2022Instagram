package pkgui

import (
	"2022Instagram-Qunkong/code/pkgui/ins"
	"2022Instagram-Qunkong/code/pkgui/ins/renwu"
	"fmt"
	"github.com/ying32/govcl/vcl"
	"strconv"
)

//::private::
type TForm1Fields struct {
}

func (f *TForm1) getGlobal() {
	//获取INS任务延迟数
	atoi, _ := strconv.Atoi(f.Insyanchi.Text())
	ins.Insyanchi = atoi
	//是否隐藏浏览器
	ins.Isheadless = f.Isheadless.Checked()
	//随机帖子话术
	path := f.Tiezi_huashu.Text()
	ins.Tiezi_huashu = ins.ReadTiezi(path)
	//预先加载账号密码文本
	ins.UserData, _ = ins.ImportuserMap("users.txt")
	//获取图片数量
	ins.PictureCount = ins.GetPictureCount()
	//开启状态更新器
	f.ZhuangTai.SetEnabled(true)
}

func (f *TForm1) OnButton1Click(sender vcl.IObject) {
	f.getGlobal()

	//fmt.Println(len(ins.UserData))
	//fmt.Println(ins.UserData)
	fmt.Println("图片的数量是", ins.PictureCount)

	//循环查看多少个账号然后启动
	for i := 0; i < len(ins.UserData); i++ {
		go renwu.InsInit(ins.UserData[i])
	}
}

func (f *TForm1) OnHostDailiChange(sender vcl.IObject) {

}

//测试读取随机话术
func (f *TForm1) OnButton2Click(sender vcl.IObject) {
	path := f.Tiezi_huashu.Text()
	ins.Tiezi_huashu = ins.ReadTiezi(path)

	fmt.Println(ins.Tiezi_huashu[ins.GetRandNum(len(ins.Tiezi_huashu))])
}

func (f *TForm1) OnButton3Click(sender vcl.IObject) {
	//测试状态
	changdu := len(ins.UserData)
	fmt.Println(changdu)
	//f.INSzhuangtai.SetRowCount(int32(changdu) + 1)
	//for i := 0; i < changdu; i++ {
	//	f.INSzhuangtai.SetCells(0, int32(i+1), ins.UserData[i]["INSzhanghao"])
	//	f.INSzhuangtai.SetCells(1, int32(i+1), ins.UserData[i]["INSzhuangtai"])
	//}
	//测试其他
	//fmt.Println(f.Isheadless.Checked())
}

func (f *TForm1) OnZhuangTaiTimer(sender vcl.IObject) {
	changdu := len(ins.UserData)
	f.INSzhuangtai.SetRowCount(int32(changdu) + 1)
	for i := 0; i < changdu; i++ {
		f.INSzhuangtai.SetCells(0, int32(i+1), ins.UserData[i]["INSzhanghao"])
		f.INSzhuangtai.SetCells(1, int32(i+1), ins.UserData[i]["INSzhuangtai"])
	}

}

func (f *TForm1) OnButton4Click(sender vcl.IObject) {

}

//切换所有账号的IP
func (f *TForm1) OnIPqiehuanClick(sender vcl.IObject) {
	//预先加载账号密码文本
	ins.UserData, _ = ins.ImportuserMap("users.txt")
	fmt.Println(ins.UserData)
	fmt.Println(len(ins.UserData))

	for i := 0; i < len(ins.UserData); i++ {
		ins.ChangeIP(ins.UserData[i]["DLzhanghao"])
	}

}
