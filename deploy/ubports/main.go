package main

import (
	"os"

	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/quick"
)

func main() {
	EnableHighDPI()

	gui.NewQGuiApplication(len(os.Args), os.Args)

	view := quick.NewQQuickView(nil)
	view.SetSource(NewQUrl("qrc:/qml/main.qml"))
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.Show()

	gui.QGuiApplication_Exec()
}
