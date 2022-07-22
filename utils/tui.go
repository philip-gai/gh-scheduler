// These are shared utils to avoid circular imports
package utils

import (
	"strings"

	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func PushListRow(text string, list *widgets.List) {
	textRows := strings.Split(text, "\n")
	list.Rows = append(list.Rows, textRows...)
	list.ScrollBottom()
	termui.Render(list)
}

func ConcatListRow(text string, list *widgets.List) {
	list.Rows[len(list.Rows)-1] += text
}

func BackspaceListRow(list *widgets.List) {
	currentRow := list.Rows[len(list.Rows)-1]
	list.Rows[len(list.Rows)-1] = currentRow[:len(currentRow)-1]
}
