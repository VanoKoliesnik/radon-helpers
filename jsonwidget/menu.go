package jsonwidget

import (
	"radon-helpers/event"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Menu struct {
	eventBus *event.EventBus

	formatButton *widget.Button
	copyButton   *widget.Button
	pasteButton  *widget.Button
	clearButton  *widget.Button
}

func newMenu(eb *event.EventBus) *Menu {
	m := &Menu{
		eventBus: eb,

		formatButton: widget.NewButton("Format", func() {
			eb.Emit(event_FormatButtonClicked, nil)
		}),
		copyButton: widget.NewButton("Copy", func() {
			eb.Emit(event_CopyButtonClicked, nil)
		}),
		pasteButton: widget.NewButton("Paste", func() {
			eb.Emit(event_PasteButtonClicked, nil)
		}),
		clearButton: widget.NewButton("Clear", func() {
			eb.Emit(event_ClearButtonClicked, nil)
		}),
	}

	return m
}

func (m *Menu) getLayout() *fyne.Container {
	return container.NewHBox(
		m.formatButton,
		m.copyButton,
		m.pasteButton,
		m.clearButton,
	)
}
