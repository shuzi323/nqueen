package main

import (
	"os"
	"strconv"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/quick"
)

var toQml chan string

func main() {
	toQml = make(chan string)
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)
	//create an application
	gui.NewQGuiApplication(len(os.Args), os.Args)
	//create a view
	var view = quick.NewQQuickView(nil)
	view.SetTitle("nQueen")

	//creat a bridge
	var qmlBridge = NewQmlBridge(nil)
	qmlBridge.configureBridge()
	//config the view to know about the bridge
	view.RootContext().SetContextProperty("QmlBridge", qmlBridge)

	view.SetSource(core.NewQUrl3("qrc:/qml/ui.qml", 0))
	go func() {
		for strToQml := range toQml {
			qmlBridge.SendToQml(strToQml)
		}
	}()
	//Complete setup, and start the UI
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.Show()
	gui.QGuiApplication_Exec()
}

type QmlBridge struct {
	core.QObject

	//message to qml
	_ func(data string) `signal:"sendToQml"`

	//request from qml
	_ func(nStr string) `slot:"calculate`
}

func (q *QmlBridge) configureBridge() {
	q.ConnectCalculate(func(nStr string) {
		n, err := strconv.Atoi(nStr)
		if err != nil || n < 0 || n > 22 {
			toQml <- "Please enter a number between [0, 18]"
			return
		}
		toQml <- nqueen(n)
	})
}
