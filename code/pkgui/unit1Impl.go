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

	go renwu.InsInit(ins.UserData[0])
}

func (f *TForm1) OnHostDailiChange(sender vcl.IObject) {

}

//测试读取富文本
func (f *TForm1) OnButton2Click(sender vcl.IObject) {
	path := f.Tiezi_huashu.Text()
	ins.Tiezi_huashu = ins.ReadTiezi(path)

	fmt.Println(ins.Tiezi_huashu[ins.GetRandNum(len(ins.Tiezi_huashu))])
}

func (f *TForm1) OnButton3Click(sender vcl.IObject) {
	changdu := len(ins.UserData)
	f.INSzhuangtai.SetRowCount(int32(changdu) + 1)
	for i := 0; i < changdu; i++ {
		f.INSzhuangtai.SetCells(0, int32(i+1), ins.UserData[i]["INSzhanghao"])
		f.INSzhuangtai.SetCells(1, int32(i+1), ins.UserData[i]["INSzhuangtai"])
	}
}

func (f *TForm1) OnZhuangTaiTimer(sender vcl.IObject) {
	changdu := len(ins.UserData)
	f.INSzhuangtai.SetRowCount(int32(changdu) + 1)
	for i := 0; i < changdu; i++ {
		f.INSzhuangtai.SetCells(0, int32(i+1), ins.UserData[i]["INSzhanghao"])
		f.INSzhuangtai.SetCells(1, int32(i+1), ins.UserData[i]["INSzhuangtai"])
	}

}
