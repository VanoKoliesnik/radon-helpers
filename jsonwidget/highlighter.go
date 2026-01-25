package jsonwidget

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/lexers"
)

func highlightJSON(grid *widget.TextGrid, text string) {
	fyne.Do(func() { grid.SetText(text) })

	lexer := lexers.Get("json")
	if lexer == nil {
		log.Fatal("No JSON lexer found")
	}

	iterator, err := lexer.Tokenise(nil, text)
	if err != nil {
		log.Fatal(err)
	}

	colorMap := map[chroma.TokenType]color.RGBA{
		chroma.Punctuation:          {64, 64, 64, 255},  // green
		chroma.NameTag:              {0, 120, 255, 255}, // blue
		chroma.LiteralNumber:        {220, 140, 0, 255}, // orange
		chroma.LiteralNumberInteger: {220, 140, 0, 255}, // orange
		chroma.StringDouble:         {0, 180, 0, 255},   // green
		chroma.KeywordConstant:      {180, 0, 180, 255}, // purple
	}

	row, column := 0, 0
	for token := iterator(); token != chroma.EOF; token = iterator() {
		runes := []rune(token.Value)
		for _, tokenRune := range runes {
			style := &widget.CustomTextGridStyle{
				FGColor: color.Gray{},
			}

			if currentColor, ok := colorMap[token.Type]; ok {
				style.FGColor = currentColor
			}

			currentRow, currentCol := row, column
			currentStyle := style

			fyne.Do(func() { grid.SetStyle(currentRow, currentCol, currentStyle) })

			column++

			if tokenRune == '\n' {
				row++
				column = 0
			}
		}
	}
}
