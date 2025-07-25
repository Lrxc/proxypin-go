package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"proxypin-go/assets"
	"proxypin-go/internal/gui/cus"
)

var AppName = "ProxyPin-Go"

func Gui() {
	myApp := app.New()
	myWindow := myApp.NewWindow(AppName)    //主窗口
	myWindow.Resize(fyne.NewSize(400, 500)) //窗口大小
	myWindow.CenterOnScreen()               //窗口居中

	content := initView(myApp, myWindow)
	initTray(myApp, myWindow)
	go initTask(myWindow)

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
	setItem := widget.NewToolbarAction(theme.SettingsIcon(), nil)
	setItem.OnActivated = setOnClick(myWindow, setItem)

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.FileTextIcon(), editRuleOnClick(myApp)),
		widget.NewToolbarSpacer(),
		setItem,
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			dialog.ShowInformation("关于", "v1.0.1", myWindow)
		}),
	)

	caBtn := widget.NewButton(CA_STATUS, nil)
	caBtn.OnTapped = caOnClick(myWindow, caBtn)
	caBtn.Importance = widget.WarningImportance

	// 创建一个标签
	statusTitle := widget.NewLabel(PROXY_TITLE)
	Proxy_Status.Set(PROXY_STATUS_OFF)
	// 创建一个标签(绑定数据,自动更新)
	statusLabel := widget.NewLabelWithData(Proxy_Status)
	statusLabel.Importance = widget.WarningImportance

	//开始按钮
	startBtn := widget.NewButton(PROXY_BTN_START, nil)
	startBtn.OnTapped = btnOnClick(myWindow, startBtn)

	//横线
	thickLine := canvas.NewRectangle(color.NRGBA{R: 128, G: 128, B: 128, A: 255})

	top := container.NewVBox(
		toolbar,
		thickLine,
		cus.NewLayout(0, 50),
		container.NewHBox(layout.NewSpacer(), caBtn, layout.NewSpacer()),
		container.NewHBox(cus.NewLayout(120, 0), statusTitle, statusLabel), //嵌套一个水平布局,并且居中
		cus.NewLayout(0, 50),
	)
	content := container.NewBorder(
		top,
		cus.NewLayout(200, 100), cus.NewLayout(100, 100), cus.NewLayout(100, 100),
		startBtn,
	)
	return content
}
