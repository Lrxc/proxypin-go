package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"proxypin-go/assets"
	"proxypin-go/internal/config"
)

var AppName = "ProxyPin-Go"

func Gui() {
	myApp := app.New()
	myWindow := myApp.NewWindow(AppName)    //主窗口
	myWindow.Resize(fyne.NewSize(400, 500)) //窗口大小
	myWindow.CenterOnScreen()               //窗口居中

	content := initView(myApp, myWindow)
	initTray(myApp, myWindow)

	myWindow.SetContent(content)
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

// 选择的打包路径
var (
	Proxy_Status = binding.NewString()
)

func initView(myApp fyne.App, myWindow fyne.Window) *fyne.Container {
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), editRuleClick(myApp)),
		widget.NewToolbarAction(theme.SettingsIcon(), settingClick(myApp)),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), helpClick(myWindow)),
	)

	// 创建一个标签
	statusTitle := widget.NewLabel(config.PROXY_TITLE)
	Proxy_Status.Set(config.PROXY_STATUS_OFF)
	// 创建一个标签(绑定数据,自动更新)
	statusLabel := widget.NewLabelWithData(Proxy_Status)

	startBtn := widget.NewButton(config.PROXY_BTN_START, nil)
	startBtn.OnTapped = btnOnClick(startBtn)
	// 设置按钮的最小和固定大小
	startBtn.Resize(fyne.NewSize(200, 200))

	content := container.NewVBox(
		toolbar,
		//嵌套一个水平布局
		container.NewHBox(statusTitle, statusLabel),
	)
	content = container.NewBorder(
		content,
		NewMySpacer(100, 100), layout.NewSpacer(), layout.NewSpacer(),
		startBtn,
	)
	return content
}
