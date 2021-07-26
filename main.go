package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/linkedin/goavro/v2"
)

const WIDTH float32 = 600
const HEIGHT float32 = 300

func verifyMessage(schema string, message string, resultLabel *widget.Label) {
	codec, err := goavro.NewCodec(schema)
	if err != nil {
		resultLabel.SetText(err.Error())
		return
	}
	decoded, _, err := codec.NativeFromTextual([]byte(message))
	if err != nil {
		resultLabel.SetText(err.Error())
		return
	}
	resultLabel.SetText("Message matches schema!")
	fmt.Println(decoded)
}

func main() {
	a := app.New()
	w := a.NewWindow("Avro Verif Tool")
	w.CenterOnScreen()
	w.SetMainMenu(fyne.NewMainMenu(fyne.NewMenu("Theme", fyne.NewMenuItem("Dark", func() {
		a.Settings().SetTheme(theme.DarkTheme())
	}), fyne.NewMenuItem("Light", func() {
		a.Settings().SetTheme(theme.LightTheme())
	}))))

	w.Resize(fyne.NewSize(WIDTH, HEIGHT))
	avroLabel := widget.NewLabel("Avro schema")
	schemaEntry := widget.NewMultiLineEntry()
	schemaEntry.SetText("AVRO SCHEMA")
	messageLabel := widget.NewLabel("Message")
	messageEntry := widget.NewMultiLineEntry()
	messageEntry.SetText("MESSAGE")
	messageEntry.Resize(fyne.NewSize(100, 300))
	resultLabel := widget.NewLabel("")

	w.SetContent(container.NewVBox(
		avroLabel,
		schemaEntry,
		messageLabel,
		messageEntry,
		resultLabel,
		widget.NewButton("Verify", func() {
			verifyMessage(schemaEntry.Text, messageEntry.Text, resultLabel)
		}),
	))
	w.ShowAndRun()

}
