package main

import (
	"radon-helpers/jsonwidget"
	"radon-helpers/shared"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type Radon struct {
	jsonWidget jsonwidget.Widget
}

func main() {
	a := app.New()
	w := a.NewWindow("Radon Helpers")
	w.Resize(fyne.NewSize(shared.DEFAULT_WINDOW_WIDTH, shared.DEFAULT_WINDOW_HEIGHT))

	r := Radon{
		jsonWidget: jsonwidget.NewJsonWidget(),
	}

	w.SetContent(r.jsonWidget.GetCanvas())
	w.ShowAndRun()
}
