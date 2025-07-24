package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"proxypin-go/assets"
	"proxypin-go/internal/system"
)

var AppName = "ProxyPin-Go"

func Gui() {
	myApp := app.New()
	myWindow := myApp.NewWindow(AppName) //主窗口

	content := initView(myApp, myWindow)
	initTray(myApp, myWindow)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(800, 600)) //窗口大小
	myWindow.CenterOnScreen()               //窗口居中
	myWindow.ShowAndRun()
}

func initTray(myApp fyne.App, myWindow fyne.Window) {
	// 创建自定义托盘图标
	icon := &fyne.StaticResource{
		StaticName:    "tray.jpg",
		StaticContent: assets.Read("tray.jpg"),
	}

	if desk, ok := myApp.(desktop.App); ok {
		desk.SetSystemTrayIcon(icon)

		menu := fyne.NewMenu("我的应用",
			fyne.NewMenuItem("显示", func() { myWindow.Show() }),
			fyne.NewMenuItem("退出", func() { myApp.Quit() }),
		)
		desk.SetSystemTrayMenu(menu)
	}

	// 拦截关闭事件
	myWindow.SetCloseIntercept(func() {
		myWindow.Hide()
	})
}

func initView(myApp fyne.App, myWindow fyne.Window) *fyne.Container {
	// 创建一个标签
	selectedFileTitle := widget.NewLabel("")

	//创建一个按钮
	makeButton := widget.NewButton("Start", nil)
	makeButton.OnTapped = btnOnClick(makeButton)

	// 创建一个容器，包含标签和按钮
	content := container.NewVBox(
		selectedFileTitle,
		makeButton,
	)
	return content
}

func btnOnClick(btn *widget.Button) func() {
	return func() {
		text := btn.Text
		if text == "Start" {
			system.SysProxyOn()
			btn.Text = "Stop"
		}
		if text == "Stop" {
			system.SysProxyOff() //关闭代理
			btn.Text = "Start"
		}
	}
}
