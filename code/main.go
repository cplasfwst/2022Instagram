// Automatically generated by the res2go IDE plug-in.
package main

import (
	"2022Instagram-Qunkong/code/pkgui"
	"github.com/ying32/govcl/vcl"
)

//func init()  {
//	fmt.Println("开始")
//	if true {
//		log.Fatal("??")
//	}
//}

func main() {
	vcl.Application.SetScaled(true)
	vcl.Application.SetTitle("project1")
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(&pkgui.Form1)
	vcl.Application.Run()
}
