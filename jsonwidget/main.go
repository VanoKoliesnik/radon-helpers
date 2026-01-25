package jsonwidget

import (
	"radon-helpers/event"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type Widget struct {
	eventBus event.EventBus

	menu   *Menu
	editor *Editor
	viewer *Viewer
}

func NewJsonWidget() Widget {
	w := Widget{
		eventBus: *event.NewEventBus(),
	}

	w.menu = newMenu(&w.eventBus)
	w.editor = newEditor(&w.eventBus)
	w.viewer = newViewer(&w.eventBus)

	return w
}

func (w Widget) GetCanvas() fyne.CanvasObject {
	return container.NewBorder(
		w.menu.getLayout(), // top

		nil, // bottom

		nil, // left

		nil, // right

		container.NewHSplit(
			w.editor.getLayout(),
			w.viewer.getLayout(),
		),
	)
}
