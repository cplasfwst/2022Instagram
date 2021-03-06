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
	atoi, _ := strconv.Atoi(f.InsyanchiMin.Text())
	ins.InsyanchiMin = atoi
	atoi2, _ := strconv.Atoi(f.InsyanchiMax.Text())
	ins.InsyanchiMax = atoi2
	//是否隐藏浏览器
	ins.Isheadless = f.Isheadless.Checked()
	//随机帖子话术
	path := f.Tiezi_huashu.Text()
	ins.Tiezi_huashu = ins.ReadTiezi(path)
	//预先加载账号密码文本(添加了只初始化一次的控制)
	if ins.IsUserData {
		ins.UserData, _ = ins.ImportuserMap("users.txt")
		ins.IsUserData = false
	}
	//获取图片数量
	ins.PictureCount = ins.GetPictureCount()
	//开启状态更新器
	f.ZhuangTai.SetEnabled(true)
	//获取固定热门关键词和关键词数量
	ins.INSgudinghot = f.INSgudinghot.Text()
	shuliang, _ := strconv.Atoi(f.INSsuijihot.Text())
	ins.INSsuijihot = shuliang
}

func (f *TForm1) OnButton1Click(sender vcl.IObject) {
	f.getGlobal()

	//fmt.Println(len(ins.UserData))
	//fmt.Println(ins.UserData)
	fmt.Println("图片的数量是", ins.PictureCount)

	//循环查看多少个账号然后启动
	for i := 0; i < len(ins.UserData); i++ {
		go renwu.InsInit(ins.UserData[i])
		//time.Sleep(time.Second*5)
		//if i >= 19 && i < 43 {
		//	time.Sleep(time.Second * 2)
		//	go renwu.InsInit(ins.UserData[i])
		//}
		//if i >= 43 && i < 63 {
		//	time.Sleep(time.Second * 2)
		//	go renwu.InsInit(ins.UserData[i])
		//}
	}
}

func (f *TForm1) OnHostDailiChange(sender vcl.IObject) {

}

//测试读取随机话术(任意测试)
func (f *TForm1) OnButton2Click(sender vcl.IObject) {
	path := f.Tiezi_huashu.Text()
	ins.Tiezi_huashu = ins.ReadTiezi(path)

	fmt.Println(ins.Tiezi_huashu[ins.GetRandNum(len(ins.Tiezi_huashu))])
}

func (f *TForm1) OnButton3Click(sender vcl.IObject) {
	//测试状态
	//changdu := len(ins.UserData)
	//fmt.Println(changdu)
	ins.UserData, _ = ins.ImportuserMap("users.txt")
	fmt.Println(ins.MapRead(ins.UserData[0], "UserAgent"))
	fmt.Println(ins.MapRead(ins.UserData[16], "UserAgent"))
	fmt.Println(ins.MapRead(ins.UserData[54], "UserAgent"))

}

func (f *TForm1) OnZhuangTaiTimer(sender vcl.IObject) {
	changdu := len(ins.UserData)
	f.INSzhuangtai.SetRowCount(int32(changdu) + 1)
	for i := 0; i < changdu; i++ {
		f.INSzhuangtai.EditorTextChanged(0, int32(i+1), strconv.Itoa(i))
		f.INSzhuangtai.EditorTextChanged(1, int32(i+1), ins.MapRead(ins.UserData[i], "INSzhanghao"))
		f.INSzhuangtai.EditorTextChanged(2, int32(i+1), ins.MapRead(ins.UserData[i], "INSzhuangtai"))
	}

}

//切换所有账号的IP
func (f *TForm1) OnIPqiehuanClick(sender vcl.IObject) {
	//预先加载账号密码文本
	ins.UserData, _ = ins.ImportuserMap("users.txt")
	fmt.Println(ins.UserData)
	fmt.Println(len(ins.UserData))

	for i := 0; i < len(ins.UserData); i++ {
		ins.ChangeIP(ins.MapRead(ins.UserData[i], "DLzhanghao"))
	}

}

//指定运行第几个账号
func (f *TForm1) OnZhidingClick(sender vcl.IObject) {
	f.getGlobal()
	//获取账号文本
	num := f.Zhidingnum.Text()
	atoi, _ := strconv.Atoi(num)

	fmt.Println("这里是！", ins.UserData[7])
	go renwu.InsInit(ins.UserData[atoi])

}

//指定运行第X到X个账号
func (f *TForm1) OnZhiding12BtnClick(sender vcl.IObject) {
	f.getGlobal()
	num1 := f.Zhiding1.Text()
	num2 := f.Zhiding2.Text()
	atoi1, _ := strconv.Atoi(num1)
	atoi2, _ := strconv.Atoi(num2)
	//geshu := atoi2 - atoi1

	for i := atoi1; i <= atoi2; i++ {
		go renwu.InsInit(ins.UserData[i])
	}

}
