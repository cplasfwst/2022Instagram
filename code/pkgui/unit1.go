// Automatically generated by the res2go IDE plug-in, do not edit.
package pkgui

import (
	"github.com/ying32/govcl/vcl"
)

type TForm1 struct {
	*vcl.TForm
	PageControl1 *vcl.TPageControl
	TabSheet1    *vcl.TTabSheet
	Button1      *vcl.TButton
	INSzhuangtai *vcl.TStringGrid
	Button3      *vcl.TButton
	Label2       *vcl.TLabel
	InsyanchiMin *vcl.TEdit
	Label3       *vcl.TLabel
	Label5       *vcl.TLabel
	Label6       *vcl.TLabel
	Zhidingnum   *vcl.TEdit
	Label7       *vcl.TLabel
	Zhiding      *vcl.TButton
	Label10      *vcl.TLabel
	Zhiding1     *vcl.TEdit
	Label11      *vcl.TLabel
	Zhiding2     *vcl.TEdit
	Zhiding12Btn *vcl.TButton
	TabSheet2    *vcl.TTabSheet
	Label1       *vcl.TLabel
	Isheadless   *vcl.TCheckBox
	Label4       *vcl.TLabel
	CheckBox1    *vcl.TCheckBox
	IPqiehuan    *vcl.TButton
	TabSheet3    *vcl.TTabSheet
	Button2      *vcl.TButton
	Tiezi_huashu *vcl.TEdit
	Label8       *vcl.TLabel
	INSgudinghot *vcl.TEdit
	Label9       *vcl.TLabel
	INSsuijihot  *vcl.TEdit
	ZhuangTai    *vcl.TTimer
	Label12      *vcl.TLabel
	InsyanchiMax *vcl.TEdit

	//::private::
	TForm1Fields
}

var Form1 *TForm1

// vcl.Application.CreateForm(&Form1)

func NewForm1(owner vcl.IComponent) (root *TForm1) {
	vcl.CreateResForm(owner, &root)
	return
}

var form1Bytes = []byte("\x54\x50\x46\x30\x06\x54\x46\x6F\x72\x6D\x31\x05\x46\x6F\x72\x6D\x31\x04\x4C\x65\x66\x74\x03\x31\x01\x06\x48\x65\x69\x67\x68\x74\x03\xA4\x02\x03\x54\x6F\x70\x03\x9D\x00\x05\x57\x69\x64\x74\x68\x03\x25\x04\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x18\x32\x30\x32\x32\x2D\x58\x64\x6F\x63\x74\x6F\x72\xE7\xBE\xA4\xE6\x8E\xA7\xE7\xA5\x9E\xE5\x99\xA8\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\xA4\x02\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x25\x04\x0A\x4C\x43\x4C\x56\x65\x72\x73\x69\x6F\x6E\x06\x08\x32\x2E\x30\x2E\x31\x32\x2E\x30\x00\x0C\x54\x50\x61\x67\x65\x43\x6F\x6E\x74\x72\x6F\x6C\x0C\x50\x61\x67\x65\x43\x6F\x6E\x74\x72\x6F\x6C\x31\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x03\x5C\x02\x03\x54\x6F\x70\x02\x48\x05\x57\x69\x64\x74\x68\x03\x25\x04\x0A\x41\x63\x74\x69\x76\x65\x50\x61\x67\x65\x07\x09\x54\x61\x62\x53\x68\x65\x65\x74\x31\x08\x54\x61\x62\x49\x6E\x64\x65\x78\x02\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x09\x54\x54\x61\x62\x53\x68\x65\x65\x74\x09\x54\x61\x62\x53\x68\x65\x65\x74\x31\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x03\x49\x4E\x53\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\x3E\x02\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x1D\x04\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x31\x04\x4C\x65\x66\x74\x03\xA8\x03\x06\x48\x65\x69\x67\x68\x74\x02\x36\x03\x54\x6F\x70\x03\x08\x02\x05\x57\x69\x64\x74\x68\x02\x6D\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x12\xE5\xBC\x80\xE5\xA7\x8B\xE6\x89\x80\xE6\x9C\x89\xE8\xB4\xA6\xE5\x8F\xB7\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x0B\x54\x53\x74\x72\x69\x6E\x67\x47\x72\x69\x64\x0C\x49\x4E\x53\x7A\x68\x75\x61\x6E\x67\x74\x61\x69\x04\x4C\x65\x66\x74\x02\xFC\x06\x48\x65\x69\x67\x68\x74\x03\xD8\x01\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\x7C\x02\x08\x43\x6F\x6C\x43\x6F\x75\x6E\x74\x02\x03\x08\x52\x6F\x77\x43\x6F\x75\x6E\x74\x02\x01\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x09\x43\x6F\x6C\x57\x69\x64\x74\x68\x73\x01\x02\x2A\x03\xC5\x00\x03\x88\x01\x00\x05\x43\x65\x6C\x6C\x73\x01\x02\x03\x02\x00\x02\x00\x06\x06\xE5\xBA\x8F\xE5\x8F\xB7\x02\x01\x02\x00\x06\x06\xE8\xB4\xA6\xE5\x8F\xB7\x02\x02\x02\x00\x06\x0C\xE8\xB4\xA6\xE5\x8F\xB7\xE7\x8A\xB6\xE6\x80\x81\x00\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x33\x04\x4C\x65\x66\x74\x03\x9A\x01\x06\x48\x65\x69\x67\x68\x74\x02\x2D\x03\x54\x6F\x70\x03\xF5\x01\x05\x57\x69\x64\x74\x68\x02\x50\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE6\xB5\x8B\xE8\xAF\x95\xE6\x8C\x89\xE9\x92\xAE\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x32\x04\x4C\x65\x66\x74\x03\x8F\x02\x06\x48\x65\x69\x67\x68\x74\x02\x11\x03\x54\x6F\x70\x02\x0F\x05\x57\x69\x64\x74\x68\x02\x78\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x1E\xE6\xAF\x8F\xE4\xB8\xAA\xE4\xBB\xBB\xE5\x8A\xA1\xE7\x9A\x84\xE5\xBB\xB6\xE8\xBF\x9F\xE6\x97\xB6\xE9\x97\xB4\xEF\xBC\x9A\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x05\x54\x45\x64\x69\x74\x0C\x49\x6E\x73\x79\x61\x6E\x63\x68\x69\x4D\x69\x6E\x04\x4C\x65\x66\x74\x03\x08\x03\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x0D\x05\x57\x69\x64\x74\x68\x02\x30\x0B\x4E\x75\x6D\x62\x65\x72\x73\x4F\x6E\x6C\x79\x09\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x03\x04\x54\x65\x78\x74\x06\x04\x33\x36\x30\x30\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x33\x04\x4C\x65\x66\x74\x03\x88\x03\x06\x48\x65\x69\x67\x68\x74\x02\x11\x03\x54\x6F\x70\x02\x0F\x05\x57\x69\x64\x74\x68\x02\x0C\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x03\xE7\xA7\x92\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x35\x04\x4C\x65\x66\x74\x03\x8F\x02\x06\x48\x65\x69\x67\x68\x74\x02\x11\x03\x54\x6F\x70\x03\x90\x00\x05\x57\x69\x64\x74\x68\x02\x48\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x12\xE5\xBC\x80\xE5\x8F\x91\xE8\x80\x85\xE5\xB7\xA5\xE5\x85\xB7\xEF\xBC\x9A\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x36\x04\x4C\x65\x66\x74\x03\x98\x02\x06\x48\x65\x69\x67\x68\x74\x02\x11\x03\x54\x6F\x70\x03\xB8\x00\x05\x57\x69\x64\x74\x68\x02\x3C\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0F\xE6\x8C\x87\xE5\xAE\x9A\xE8\xBF\x90\xE8\xA1\x8C\xE7\xAC\xAC\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x05\x54\x45\x64\x69\x74\x0A\x5A\x68\x69\x64\x69\x6E\x67\x6E\x75\x6D\x04\x4C\x65\x66\x74\x03\xE2\x02\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x03\xB4\x00\x05\x57\x69\x64\x74\x68\x02\x30\x0B\x4E\x75\x6D\x62\x65\x72\x73\x4F\x6E\x6C\x79\x09\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x04\x04\x54\x65\x78\x74\x06\x01\x30\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x37\x04\x4C\x65\x66\x74\x03\x20\x03\x06\x48\x65\x69\x67\x68\x74\x02\x11\x03\x54\x6F\x70\x03\xB8\x00\x05\x57\x69\x64\x74\x68\x02\x24\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x09\xE4\xB8\xAA\xE8\xB4\xA6\xE5\x8F\xB7\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x5A\x68\x69\x64\x69\x6E\x67\x04\x4C\x65\x66\x74\x03\xA8\x03\x06\x48\x65\x69\x67\x68\x74\x02\x28\x03\x54\x6F\x70\x03\xAF\x00\x05\x57\x69\x64\x74\x68\x02\x59\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE6\x8C\x87\xE5\xAE\x9A\xE5\xBC\x80\xE5\xA7\x8B\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x05\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x07\x4C\x61\x62\x65\x6C\x31\x30\x04\x4C\x65\x66\x74\x03\x95\x02\x06\x48\x65\x69\x67\x68\x74\x02\x11\x03\x54\x6F\x70\x03\xFD\x00\x05\x57\x69\x64\x74\x68\x02\x24\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x09\xE8\xBF\x90\xE8\xA1\x8C\xEF\xBC\x9A\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x05\x54\x45\x64\x69\x74\x08\x5A\x68\x69\x64\x69\x6E\x67\x31\x04\x4C\x65\x66\x74\x03\xB8\x02\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x03\xF8\x00\x05\x57\x69\x64\x74\x68\x02\x38\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x06\x04\x54\x65\x78\x74\x06\x01\x30\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x07\x4C\x61\x62\x65\x6C\x31\x31\x04\x4C\x65\x66\x74\x03\xF8\x02\x06\x48\x65\x69\x67\x68\x74\x02\x11\x03\x54\x6F\x70\x03\xFE\x00\x05\x57\x69\x64\x74\x68\x02\x0C\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x03\xE5\x88\xB0\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x05\x54\x45\x64\x69\x74\x08\x5A\x68\x69\x64\x69\x6E\x67\x32\x04\x4C\x65\x66\x74\x03\x13\x03\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x03\xF8\x00\x05\x57\x69\x64\x74\x68\x02\x41\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x07\x04\x54\x65\x78\x74\x06\x02\x31\x30\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x0C\x5A\x68\x69\x64\x69\x6E\x67\x31\x32\x42\x74\x6E\x04\x4C\x65\x66\x74\x03\xA8\x03\x06\x48\x65\x69\x67\x68\x74\x02\x25\x03\x54\x6F\x70\x03\xF4\x00\x05\x57\x69\x64\x74\x68\x02\x5C\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x12\xE6\x8C\x87\xE5\xAE\x9A\xE6\x89\xB9\xE6\x95\xB0\xE5\xBC\x80\xE5\xA7\x8B\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x08\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x07\x4C\x61\x62\x65\x6C\x31\x32\x04\x4C\x65\x66\x74\x03\x3B\x03\x06\x48\x65\x69\x67\x68\x74\x02\x11\x03\x54\x6F\x70\x02\x10\x05\x57\x69\x64\x74\x68\x02\x0C\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x03\xE5\x88\xB0\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x05\x54\x45\x64\x69\x74\x0C\x49\x6E\x73\x79\x61\x6E\x63\x68\x69\x4D\x61\x78\x04\x4C\x65\x66\x74\x03\x4F\x03\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x0C\x05\x57\x69\x64\x74\x68\x02\x30\x0B\x4E\x75\x6D\x62\x65\x72\x73\x4F\x6E\x6C\x79\x09\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x09\x04\x54\x65\x78\x74\x06\x04\x37\x32\x30\x30\x00\x00\x00\x09\x54\x54\x61\x62\x53\x68\x65\x65\x74\x09\x54\x61\x62\x53\x68\x65\x65\x74\x32\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE8\xAE\xBE\xE7\xBD\xAE\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\x3E\x02\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x1D\x04\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x31\x04\x4C\x65\x66\x74\x02\x10\x06\x48\x65\x69\x67\x68\x74\x02\x11\x03\x54\x6F\x70\x02\x18\x05\x57\x69\x64\x74\x68\x02\x6C\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x1B\xE6\x98\xAF\xE5\x90\xA6\xE9\x9A\x90\xE8\x97\x8F\xE8\xBF\x90\xE8\xA1\x8C\xE9\xA1\xB5\xE9\x9D\xA2\xEF\xBC\x9A\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x09\x54\x43\x68\x65\x63\x6B\x42\x6F\x78\x0A\x49\x73\x68\x65\x61\x64\x6C\x65\x73\x73\x04\x4C\x65\x66\x74\x03\x8B\x00\x06\x48\x65\x69\x67\x68\x74\x02\x15\x03\x54\x6F\x70\x02\x16\x05\x57\x69\x64\x74\x68\x02\x22\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x03\xE6\x98\xAF\x07\x43\x68\x65\x63\x6B\x65\x64\x09\x05\x53\x74\x61\x74\x65\x07\x09\x63\x62\x43\x68\x65\x63\x6B\x65\x64\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x34\x04\x4C\x65\x66\x74\x02\x0F\x06\x48\x65\x69\x67\x68\x74\x02\x11\x03\x54\x6F\x70\x02\x4D\x05\x57\x69\x64\x74\x68\x02\x6B\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x1A\xE6\x89\x80\xE6\x9C\x89\xE8\xB4\xA6\xE5\x8F\xB7\x49\x50\xE5\x88\x87\xE6\x8D\xA2\xE5\x88\xB0\xEF\xBC\x9A\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x09\x54\x43\x68\x65\x63\x6B\x42\x6F\x78\x09\x43\x68\x65\x63\x6B\x42\x6F\x78\x31\x04\x4C\x65\x66\x74\x03\x80\x00\x06\x48\x65\x69\x67\x68\x74\x02\x15\x03\x54\x6F\x70\x02\x49\x05\x57\x69\x64\x74\x68\x02\x3A\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x09\xE4\xBF\x84\xE7\xBD\x97\xE6\x96\xAF\x07\x43\x68\x65\x63\x6B\x65\x64\x09\x05\x53\x74\x61\x74\x65\x07\x09\x63\x62\x43\x68\x65\x63\x6B\x65\x64\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x09\x49\x50\x71\x69\x65\x68\x75\x61\x6E\x04\x4C\x65\x66\x74\x02\x48\x06\x48\x65\x69\x67\x68\x74\x02\x29\x03\x54\x6F\x70\x02\x68\x05\x57\x69\x64\x74\x68\x02\x66\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE5\xBC\x80\xE5\xA7\x8B\xE5\x88\x87\xE6\x8D\xA2\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x00\x00\x00\x09\x54\x54\x61\x62\x53\x68\x65\x65\x74\x09\x54\x61\x62\x53\x68\x65\x65\x74\x33\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE9\x9A\x8F\xE6\x9C\xBA\xE8\xAF\x9D\xE6\x9C\xAF\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\x3E\x02\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x1D\x04\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x32\x04\x4C\x65\x66\x74\x03\xA8\x00\x06\x48\x65\x69\x67\x68\x74\x02\x2A\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x02\x47\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x07\x42\x75\x74\x74\x6F\x6E\x32\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x05\x54\x45\x64\x69\x74\x0C\x54\x69\x65\x7A\x69\x5F\x68\x75\x61\x73\x68\x75\x04\x4C\x65\x66\x74\x02\x28\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x0D\x05\x57\x69\x64\x74\x68\x02\x63\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x04\x54\x65\x78\x74\x06\x10\x74\x69\x65\x7A\x69\x5F\x68\x75\x61\x73\x68\x75\x2E\x74\x78\x74\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x38\x04\x4C\x65\x66\x74\x02\x09\x06\x48\x65\x69\x67\x68\x74\x02\x11\x03\x54\x6F\x70\x02\x69\x05\x57\x69\x64\x74\x68\x02\x60\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x18\xE5\x9B\xBA\xE5\xAE\x9A\xE7\x83\xAD\xE9\x97\xA8\xE5\x85\xB3\xE9\x94\xAE\xE5\xAD\x97\xEF\xBC\x9A\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x05\x54\x45\x64\x69\x74\x0C\x49\x4E\x53\x67\x75\x64\x69\x6E\x67\x68\x6F\x74\x04\x4C\x65\x66\x74\x02\x70\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x65\x05\x57\x69\x64\x74\x68\x03\x92\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x39\x04\x4C\x65\x66\x74\x02\x09\x06\x48\x65\x69\x67\x68\x74\x02\x11\x03\x54\x6F\x70\x03\x9E\x00\x05\x57\x69\x64\x74\x68\x02\x60\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x18\xE7\x83\xAD\xE9\x97\xA8\xE5\x85\xB3\xE9\x94\xAE\xE5\xAD\x97\xE6\x95\xB0\xE9\x87\x8F\xEF\xBC\x9A\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x05\x54\x45\x64\x69\x74\x0B\x49\x4E\x53\x73\x75\x69\x6A\x69\x68\x6F\x74\x04\x4C\x65\x66\x74\x02\x70\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x03\x98\x00\x05\x57\x69\x64\x74\x68\x03\x92\x00\x0B\x4E\x75\x6D\x62\x65\x72\x73\x4F\x6E\x6C\x79\x09\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x03\x04\x54\x65\x78\x74\x06\x02\x33\x30\x00\x00\x00\x00\x06\x54\x54\x69\x6D\x65\x72\x09\x5A\x68\x75\x61\x6E\x67\x54\x61\x69\x07\x45\x6E\x61\x62\x6C\x65\x64\x08\x04\x4C\x65\x66\x74\x03\xC3\x03\x03\x54\x6F\x70\x02\x71\x00\x00\x00")

// Register Form Resources
var _ = vcl.RegisterFormResource(Form1, &form1Bytes)
