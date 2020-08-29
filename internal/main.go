package main

import (
	"github.com/therecipe/qt/widgets"
	"os"
)

func main()  {
	app := widgets.NewQApplication(len(os.Args), os.Args)
	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(800, 600)
	window.SetWindowTitle("Miranda - The Evernote Client")

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQVBoxLayout())
	window.SetCentralWidget(widget)

	input := widgets.NewQLineEdit(nil)
	input.SetPlaceholderText("...")
	widget.Layout().AddWidget(input)

	window.Show()
	app.Exec()
}