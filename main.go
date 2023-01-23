package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Exportador de Macs")
	w.Resize(fyne.NewSize(300, 700))

	macs := widget.NewMultiLineEntry()

	//fmt.Printf("%v", macs)

	content := container.NewVBox(macs, widget.NewButton("Save", func() {
		fmt.Println("Content was:", macs.Text)
	}))

	w.SetContent(widget.NewLabel("Formatas os macs escaneados para o padrão de importação do InterFocus."))

	w.SetContent(content)
	w.ShowAndRun()
}
