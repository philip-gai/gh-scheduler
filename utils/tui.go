package utils

import (
	"github.com/gizak/termui/v3/widgets"
)

func PushListRow(text string, list *widgets.List) {
	list.Rows = append(list.Rows, text)
	list.ScrollBottom()
}

func ConcatListRow(text string, list *widgets.List) {
	list.Rows[len(list.Rows)-1] += text
	list.ScrollBottom()
}

func BackspaceListRow(list *widgets.List) {
	currentRow := list.Rows[len(list.Rows)-1]
	list.Rows[len(list.Rows)-1] = currentRow[:len(currentRow)-1]
	list.ScrollBottom()
}
