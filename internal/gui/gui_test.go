package gui

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
	"testing"
)

func TestGui(t *testing.T) {
	myApp := app.New()
	myWindow := myApp.NewWindow("工具栏控件")

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			log.Println("新建文档")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("显示帮助")
		}),
	)

	content := container.NewBorder(toolbar, nil, nil, nil, widget.NewLabel("内容"))
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
