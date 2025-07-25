package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"proxypin-go/assets"
	"proxypin-go/internal/config"
	"proxypin-go/internal/core"
	"proxypin-go/internal/system"
)

func initTask(myWindow fyne.Window) {
	//自动启动代理任务
	if config.Conf.System.GlobalProxy {
		err := system.SysProxyOn()
		if err != nil {
			Proxy_Status.Set(PROXY_STATUS_ABNORMAL)
		} else {
			Proxy_Status.Set(PROXY_STATUS_RUNNING)
		}
	}

	if config.Conf.System.Https {
		Https_Status.Set(PROXY_STATUS_RUNNING)
	}

	if config.Conf.Proxy.AutoEnable {
		go core.StartServer(config.Conf.System.Https)

		startBtn.Text = PROXY_BTN_STOP
		startBtn.Importance = widget.WarningImportance //高亮颜色
	}

	checkCert(myWindow)
}

func checkCert(myWindow fyne.Window) {
	file, err := assets.ReadFile("server.crt")
	if err != nil {
		dialog.ShowError(err, myWindow)
		return
	}
	b, err := system.CheckExistCert(file)
	if !b {
		caBtn.Show()
	}
}

func settingOnClick(myWindow fyne.Window, itme *widget.ToolbarAction) func() {
	return func() {
		proxyItme := fyne.NewMenuItem("开启系统代理", nil)
		proxyItme.Checked = config.Conf.System.GlobalProxy
		proxyItme.Action = settingProxyOnClick(myWindow, proxyItme)

		httpsItme := fyne.NewMenuItem("启用HTTPS", nil)
		httpsItme.Checked = config.Conf.System.Https
		httpsItme.Action = settingHttpsOnClick(myWindow, httpsItme)

		exitItme := fyne.NewMenuItem("最小化退出", nil)
		exitItme.Checked = config.Conf.System.MinExit
		exitItme.Action = settingExitOnClick(myWindow, exitItme)

		// 创建子菜单项
		menuItems := []*fyne.MenuItem{
			proxyItme,
			httpsItme,
			exitItme,
		}

		// 创建弹出菜单
		popUp := widget.NewPopUpMenu(
			fyne.NewMenu("", menuItems...),
			myWindow.Canvas(),
		)

		// 计算按钮位置
		object := itme.ToolbarObject()
		popUp.ShowAtPosition(object.Position().Add(fyne.NewPos(0, object.Size().Height+5)))
	}
}

func helpOnClick(myWindow fyne.Window, itme *widget.ToolbarAction) func() {
	return func() {
		caItme := fyne.NewMenuItem("安装证书", settingInstallCa(myWindow))
		aboutItem := fyne.NewMenuItem("关于", func() {
			dialog.ShowInformation("关于", "v1.0.1", myWindow)
		})

		// 创建子菜单项
		menuItems := []*fyne.MenuItem{
			caItme,
			aboutItem,
		}

		// 创建弹出菜单
		popUp := widget.NewPopUpMenu(
			fyne.NewMenu("", menuItems...),
			myWindow.Canvas(),
		)

		// 计算按钮位置
		object := itme.ToolbarObject()
		popUp.ShowAtPosition(object.Position().Add(fyne.NewPos(0, object.Size().Height+5)))
	}
}

func caInsOnClick(myWindow fyne.Window, btn *widget.Button) func() {
	return func() {
		file, err := assets.ReadFile("server.crt")
		if err != nil {
			dialog.ShowError(err, myWindow)
			return
		}
		err = system.InstallCert(file)
		if err != nil {
			dialog.ShowError(err, myWindow)
		} else {
			btn.Hide()
		}
	}
}

func startOnClick(myWindow fyne.Window, btn *widget.Button) func() {
	return func() {
		text := btn.Text
		if text == PROXY_BTN_START {
			go core.StartServer(config.Conf.System.Https)

			btn.Text = PROXY_BTN_STOP
			btn.Importance = widget.WarningImportance //高亮颜色
		}
		if text == PROXY_BTN_STOP {
			err := core.StopServer() //关闭代理
			if err != nil {
				dialog.ShowError(err, myWindow)
				return
			}
			btn.Text = PROXY_BTN_START
			btn.Importance = widget.MediumImportance
		}
	}
}
