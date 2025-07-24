package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
)

type MySpacer struct {
	layout.Spacer
}

func NewMySpacer(w, h float32) *MySpacer {
	spacer := MySpacer{}
	spacer.Resize(fyne.NewSize(w, h))
	return &spacer
}
