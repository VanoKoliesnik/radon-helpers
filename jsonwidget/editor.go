package jsonwidget

import (
	"radon-helpers/event"
	"radon-helpers/formatter"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type Editor struct {
	eventBus *event.EventBus

	widget *widget.Entry
}

func newEditor(eb *event.EventBus) *Editor {
	e := &Editor{
		eventBus: eb,
		widget:   widget.NewMultiLineEntry(),
	}

	e.widget.Wrapping = fyne.TextWrapOff
	e.widget.SetPlaceHolder("Insert JSON here")

	e.setListeners()

	return e
}

func (e *Editor) setListeners() {
	e.widget.OnChanged = func(text string) {
		e.eventBus.Emit(event_TextEdited, text)
	}

	e.eventBus.On(event_PasteButtonClicked, func(_ any) {
		e.paste()
	})
	e.eventBus.On(event_FormatButtonClicked, func(_ any) {
		e.format()
	})
	e.eventBus.On(event_ClearButtonClicked, func(_ any) {
		e.clear()
	})
	e.eventBus.On(event_FormattingError, func(err any) {
		if err, ok := err.(error); ok {
			e.showError(err)
		}
	})
}

func (e *Editor) paste() {
	e.setText(fyne.CurrentApp().Clipboard().Content())
}

func (e *Editor) format() {
	formattedJson, err := formatter.FormatJson(e.widget.Text)
	if err != nil {
		e.eventBus.Emit(event_FormattingError, err)
		return
	}

	e.setText(formattedJson)
}

func (e *Editor) clear() {
	fyne.Do(func() { e.widget.SetText("") })
	e.eventBus.Emit(event_TextEdited, "")
}

func (e *Editor) showError(err error) {
	fyne.Do(func() {
		windows := fyne.CurrentApp().Driver().AllWindows()
		if len(windows) > 0 {
			dialog.ShowError(err, windows[0])
		}
	})
}

func (e *Editor) getLayout() *container.Scroll {
	return container.NewScroll(e.widget)
}

func (e *Editor) setText(text string) {
	fyne.Do(func() { e.widget.SetText(text) })
	e.eventBus.Emit(event_TextEdited, text)
}
