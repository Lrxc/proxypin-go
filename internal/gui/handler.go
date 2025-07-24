package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"proxypin-go/assets"
	"proxypin-go/internal/config"
	"proxypin-go/internal/system"
	"proxypin-go/internal/util"
)

func asyncTask(myWindow fyne.Window) {
	file, err := assets.ReadFile("server.crt")
	if err != nil {
		dialog.ShowError(err, myWindow)
		return
	}
	b, err := system.CheckExistCert(file)
	if b {
		return
	}

	dialog.ShowConfirm("警告", "证书未安装,请先安装", func(b bool) {
		if !b {
			dialog.ShowInformation("警告", "请先安装证书", myWindow)
			return
		}

		err = system.InstallCert(file)
		if err != nil {
			dialog.ShowError(err, myWindow)
		} else {
			dialog.ShowInformation("", "安装成功", myWindow)
		}
	}, myWindow)
}

func btnOnClick(myWindow fyne.Window, btn *widget.Button) func() {
	return func() {
		text := btn.Text
		if text == PROXY_BTN_START {
			err := system.SysProxyOn()
			if err != nil {
				dialog.ShowError(err, myWindow)
				return
			}
			Proxy_Status.Set(PROXY_STATUS_RUNNING)
			btn.Text = PROXY_BTN_STOP
			btn.Importance = widget.HighImportance //高亮颜色
		}
		if text == PROXY_BTN_STOP {
			err := system.SysProxyOff() //关闭代理
			if err != nil {
				dialog.ShowError(err, myWindow)
				return
			}
			Proxy_Status.Set(PROXY_STATUS_OFF)
			btn.Text = PROXY_BTN_START
			btn.Importance = widget.MediumImportance
		}
	}
}

func settingClick(myWindow fyne.Window, itme *widget.ToolbarAction) func() {
	return func() {
		// 创建子菜单项
		menuItems := []*fyne.MenuItem{
			fyne.NewMenuItem("安装证书", settingInstallCa(myWindow)),
			fyne.NewMenuItem("关于", func() {
				dialog.ShowInformation("关于", "v1.0.1", myWindow)
			}),
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

func settingInstallCa(myWindow fyne.Window) func() {
	return func() {
		file, err := assets.ReadFile("server.crt")
		if err != nil {
			dialog.ShowError(err, myWindow)
			return
		}

		err = system.InstallCert(file)
		if err != nil {
			dialog.ShowError(err, myWindow)
			return
		}
		dialog.ShowInformation("success", "", myWindow)
	}
}

func editRuleClick(myApp fyne.App) func() {
	return func() {
		//打开一个新窗口
		newWin := myApp.NewWindow(AppName)
		newWin.Resize(fyne.NewSize(400, 500))
		newWin.CenterOnScreen() //居中显示

		entry := widget.NewMultiLineEntry()
		entry.SetText(util.PrettyJSON(config.Conf.Rule))
		entry.Wrapping = fyne.TextWrapWord //启用自动换行

		saveBtn := widget.NewButton("保存", func() {
			err := config.WriteJson(entry.Text)
			if err != nil {
				dialog.ShowError(err, newWin)
			} else {
				dialog.ShowInformation("success", "", newWin)
			}
		})
		refreshBtn := widget.NewButton("格式化", func() {
			entry.SetText(util.PrettyJSON(entry.Text))
		})

		btn := container.NewHBox(saveBtn, refreshBtn)

		content := container.NewBorder(btn, nil, nil, nil, entry)
		newWin.SetContent(content)
		newWin.Show()
	}
}
