package jsonwidget

import (
	"radon-helpers/event"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Viewer struct {
	eventBus *event.EventBus

	widget *widget.TextGrid
}

func newViewer(eb *event.EventBus) *Viewer {
	v := &Viewer{
		eventBus: eb,

		widget: widget.NewTextGrid(),
	}

	v.widget.ShowLineNumbers = true
	v.widget.ShowWhitespace = true

	v.setListeners()

	return v
}

func (v *Viewer) setListeners() {
	v.eventBus.On(event_CopyButtonClicked, func(_ any) {
		fyne.CurrentApp().Clipboard().SetContent(v.widget.Text())
	})
	v.eventBus.On(event_TextEdited, func(text any) {
		if text, ok := text.(string); ok {
			v.setText(text)
		}
	})
}

func (v *Viewer) setText(text string) {
	highlightJSON(v.widget, text)
}

func (v *Viewer) getLayout() *container.Scroll {
	return container.NewScroll(v.widget)
}
