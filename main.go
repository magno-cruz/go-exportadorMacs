package main

import (
	//"fmt"
	//"os"
	//"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	a := app.New()
	w := a.NewWindow("Exportador de Macs")
	w.Resize(fyne.NewSize(300, 500))

	codigo := widget.NewMultiLineEntry()
	//texto := widget.NewMultiLineEntry()

	/*f, err := os.Create("interfocus.txt")
	if err != nil {
		fmt.Println("Arquivo já existe")
	} else {
		fmt.Println("Arquivo não existe")
	}*/
	content := container.NewVBox(codigo, widget.NewButton("Código", func() {
		return
	}))
	/*content2 := container.NewVBox(texto, widget.NewButton("Formatar", func() {
		macs := strings.Split(texto.Text, "\n")
		for i := 0; i < len(macs); i++ {
			macs[i] = fmt.Sprintf("0000000000000%v%v000001", codigo, macs[i])
		}

		juntar := strings.Join(macs, "\n")
		exportar := []byte(juntar)

		f.Write(exportar)
	}))*/

	w.SetContent(widget.NewLabel("Formatar os macs escaneados para o padrão de importação do InterFocus."))

	w.SetContent(content)
	//w.SetContent(content2)
	w.ShowAndRun()

}
