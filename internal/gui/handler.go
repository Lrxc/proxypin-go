package gui

import (
	"encoding/json"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"proxypin-go/internal/config"
	"proxypin-go/internal/system"
)

func btnOnClick(btn *widget.Button) func() {
	return func() {
		text := btn.Text
		if text == config.PROXY_BTN_START {
			Proxy_Status.Set(config.PROXY_STATUS_RUNNING)
			system.SysProxyOn()
			btn.Text = config.PROXY_BTN_STOP
			btn.Importance = widget.SuccessImportance //高亮颜色
		}
		if text == config.PROXY_BTN_STOP {
			Proxy_Status.Set(config.PROXY_STATUS_OFF)
			system.SysProxyOff() //关闭代理
			btn.Text = config.PROXY_BTN_START
			btn.Importance = widget.MediumImportance
		}
	}
}

func helpClick(myWindow fyne.Window) func() {
	return func() {
		dialog.ShowInformation("success", "", myWindow)
	}
}

func settingClick(myApp fyne.App) func() {
	return func() {

	}
}

func editRuleClick(myApp fyne.App) func() {
	return func() {
		//打开一个新窗口
		newWin := myApp.NewWindow(AppName)
		newWin.Resize(fyne.NewSize(800, 600))
		newWin.CenterOnScreen() //居中显示

		entry := widget.NewMultiLineEntry()
		entry.SetText(prettyJSON(config.Conf.Rule))
		entry.Wrapping = fyne.TextWrapWord //启用自动换行

		saveBtn := widget.NewButton("Save", func() {
			err := config.WriteJson(entry.Text)
			if err != nil {
				dialog.ShowError(err, newWin)
			} else {
				dialog.ShowInformation("success", "", newWin)
			}
		})
		refreshBtn := widget.NewButton("Format", func() {
			entry.SetText(prettyJSON(entry.Text))
		})

		btn := container.NewHBox(saveBtn, refreshBtn)

		content := container.NewBorder(btn, nil, nil, nil, entry)
		newWin.SetContent(content)
		newWin.Show()
	}
}

func prettyJSON(raw any) string {
	if s, ok := raw.(string); ok {
		json.Unmarshal([]byte(s), &raw)
	}

	pretty, _ := json.MarshalIndent(raw, "", "  ")
	return string(pretty)
}
