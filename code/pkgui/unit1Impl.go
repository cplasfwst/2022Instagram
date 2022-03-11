package pkgui

import (
	"2022Instagram-Qunkong/code/pkgui/ins"
	"github.com/ying32/govcl/vcl"
)

//::private::
type TForm1Fields struct {
}

func (f *TForm1) OnButton1Click(sender vcl.IObject) {
	ins.InsInit()
}
