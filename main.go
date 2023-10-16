package main

import (
	"crypto/md5"
	"crypto/sha256"
	_ "embed"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"image/color"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)
func Base64Encode(str string)string{ //base64编码
	base64str:=base64.StdEncoding.EncodeToString([]byte(str))
	return base64str
}
func Base64Decode(str string)string{ //base64解码
	dbase64str,_:=base64.StdEncoding.DecodeString(str)
  	return string(dbase64str)
}

func UrlEncode(str string)string{ //url编码
	urlencode:=url.QueryEscape(str)
	return urlencode
}
func UrlDecode(str string)string{ //url解码
	urldecode,_:=url.QueryUnescape(str)
  	return urldecode
}

func HexEncode(str string)string{ //转十六进制
	hexstr:=hex.EncodeToString([]byte(str))
	return hexstr
}
func HexDecode(str string)string{ //解密十六进制
	dhexstr,_:=hex.DecodeString(str)
  	return string(dhexstr)
}

func Md5Encode(str string)string{ //MD5加密
	md5byte:=md5.Sum([]byte(str))  //得到的是字节流
  	md5str:=hex.EncodeToString(md5byte[:])  //转为十六进制
  	return md5str  //或直接使用%x转为16进制：fmt.Printf("%x",md5_str)
}
func Sha256Encode(str string)string{ //sha256加密
	sha256byte:=sha256.Sum256([]byte(str))
  	sha256str:=fmt.Sprintf("%x",sha256byte) //转为十六进制
  	return sha256str
}

func SetElement(title string)(*widget.Form,*widget.Entry){
	entry:=widget.NewEntry()
	form:=widget.NewForm(widget.NewFormItem(title,entry))
	return form,entry
}
//go:embed 1.jpeg
var jpg []byte

//go:embed kaiti.TTF
var font []byte
var myfont=&fyne.StaticResource{
	StaticName: "font",
	StaticContent: font,
}
type theme1 struct{}
var _ fyne.Theme=(*theme1)(nil)
func(*theme1)Font(s fyne.TextStyle)fyne.Resource{
	return myfont
}
func (*theme1) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color { 
    return theme.DefaultTheme().Color(n, v)
 }
func (*theme1) Icon(n fyne.ThemeIconName) fyne.Resource { 
    return theme.DefaultTheme().Icon(n)
 } 
func (*theme1) Size(n fyne.ThemeSizeName) float32{
     return theme.DefaultTheme().Size(n) 
}
func main() {
	// os.Setenv("FYNE_FONT",a.Name()) //设置中文，必须放到最前面写
	myApp := app.New()
	myApp.Settings().SetTheme(&theme1{})

	myWindow := myApp.NewWindow("大王LDL专用编码工具    Author By 李栋良良")
	form1,entry1:=SetElement("输入内容   ")
	form2,entry2:=SetElement("base64编码 ")
	form3,entry3:=SetElement("url 编码      ")
	form4,entry4:=SetElement("16进制编码")
	form5,entry5:=SetElement("MD5     ")
	form6,entry6:=SetElement("SHA256")
	draft:=widget.NewMultiLineEntry()  //多行文本
	draft.Wrapping=fyne.TextWrapBreak  //自动换行
	draft.SetMinRowsVisible(6)         //默认6行高度
	draft_form:=widget.NewForm(widget.NewFormItem("草稿",draft))

	bt1:=widget.NewButton("加密",func() { //加密按钮
		entry2.SetText(Base64Encode(entry1.Text))
		entry3.SetText(UrlEncode(entry1.Text))
		entry4.SetText(HexEncode(entry1.Text))
		entry5.SetText(Md5Encode(entry1.Text))
		entry6.SetText(Sha256Encode(entry1.Text))
	})
	bt2:=widget.NewButton("解码",func() {
		entry2.SetText(Base64Decode(entry2.Text))
	})
	bt3:=widget.NewButton("解码",func() {
		entry3.SetText(UrlDecode(entry3.Text))
	})
	bt4:=widget.NewButton("解码",func() {
		entry4.SetText(HexDecode(entry4.Text))
	})

	content := container.NewVBox(
		container.NewBorder(nil,nil,nil,bt1,form1),
		container.NewBorder(nil,nil,nil,bt2,form2),
		container.NewBorder(nil,nil,nil,bt3,form3),
		container.NewBorder(nil,nil,nil,bt4,form4),
		container.NewBorder(nil,nil,nil,nil,form5),
		container.NewBorder(nil,nil,nil,nil,form6),
		container.NewBorder(nil,nil,nil,nil,draft_form),
	)
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(600, 400))
	
	// logo,_:=fyne.LoadResourceFromPath("1.jpg")
	logo:=fyne.NewStaticResource("1.jpeg",jpg)
	myWindow.SetIcon(logo)    //设置logo图标
	myWindow.ShowAndRun()
}
