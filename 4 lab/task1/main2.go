package main

import (
	"fmt"
	"strconv"
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -L. -lcalculator
#include "calculator.c"
*/
import "C"

func calculateCost(height, width int, materialIndex, glassIndex int, withSill bool) float64 {
	result := C.calculate_cost(
		C.int(height),
		C.int(width),
		C.int(materialIndex),
		C.int(glassIndex),
		C.int(boolToInt(withSill)),
	)
	return float64(result)
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func windowCalculator() {
	mainWindow := ui.NewWindow("Калькулятор вартості склопакета", 0, 0, false)
	mainWindow.SetMargined(true)

	box := ui.NewVerticalBox()
	box.SetPadded(true)

	paramBox := ui.NewGroup("Параметри вікна:")
	paramBox.SetMargined(true)
	paramInputBox := ui.NewHorizontalBox()
	paramInputBox.SetPadded(true)

	form := ui.NewForm()
	form.SetPadded(true)

	heightEntry := ui.NewEntry()
	widthEntry := ui.NewEntry()

	materialCombobox := ui.NewCombobox()
	materialCombobox.Append("Дерево")
	materialCombobox.Append("Метал")
	materialCombobox.Append("Металопластик")

	form.Append("Ширина (см)", heightEntry, false)
	form.Append("Висота (см)", widthEntry, false)
	form.Append("Матеріал", materialCombobox, false)

	paramInputBox.Append(form, false)
	paramBox.SetChild(paramInputBox)

	glassBox := ui.NewGroup("Тип склопакета:")
	glassBox.SetMargined(true)

	glassParamsBox := ui.NewVerticalBox()
	glassParamsBox.SetPadded(true)

	glassCombobox := ui.NewCombobox()
	glassCombobox.Append("Однокамерний")
	glassCombobox.Append("Двокамерний")

	glassParamsBox.Append(glassCombobox, false)

	windowSillCheckBox := ui.NewCheckbox("Підвіконня")
	glassParamsBox.Append(windowSillCheckBox, false)

	glassBox.SetChild(glassParamsBox)

	resultBox := ui.NewHorizontalBox()
	resultLabel := ui.NewLabel("")
	resultBox.Append(resultLabel, false)

	buttonBox := ui.NewVerticalBox()
	calculateButton := ui.NewButton("Розрахувати")
	priceListButton := ui.NewButton("Список цін")

	buttonBox.Append(calculateButton, false)
	buttonBox.Append(priceListButton, false)

	priceListButton.OnClicked(func(*ui.Button) {
		displayPriceList()
	})

	calculateButton.OnClicked(func(*ui.Button) {
		materialIndex := materialCombobox.Selected()
		glassIndex := glassCombobox.Selected()

		heightText := heightEntry.Text()
		widthText := widthEntry.Text()

		height, err := strconv.Atoi(heightText)
		width, err2 := strconv.Atoi(widthText)

		if err == nil && err2 == nil {
			totalCost := calculateCost(height, width, materialIndex, glassIndex, windowSillCheckBox.Checked())
			resultLabel.SetText(fmt.Sprintf("Вартість: %.2f грн", totalCost))
		} else {
			ui.MsgBoxError(mainWindow, "Помилка!", "Будь ласка, введіть коректні значення.")
		}
	})

	box.Append(paramBox, false)
	box.Append(glassBox, false)
	box.Append(resultBox, false)
	box.Append(buttonBox, false)

	mainWindow.SetChild(box)
	mainWindow.Show()

	mainWindow.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
}
func displayPriceList() {
	window := ui.NewWindow("Ціни на склопакети", 0, 0, false)
	window.SetMargined(true)

	box := ui.NewVerticalBox()
	box.SetPadded(true)

	priceLabels := []string{
		"Однокамерний, дерев'яний – 2.5 грн/см²",
		"Двокамерний, дерев'яний – 3 грн/см²",
		"Однокамерний, металевий – 0.5 грн/см²",
		"Двокамерний, металевий – 1 грн/см²",
		"Однокамерний, металопластиковий – 1.5 грн/см²",
		"Двокамерний, металопластиковий – 2 грн/см²",
		"Вартість підвіконня – 350 грн",
	}

	for _, label := range priceLabels {
		box.Append(ui.NewLabel(label), false)
	}

	okButton := ui.NewButton("OK")
	box.Append(okButton, false)

	group := ui.NewGroup("Ціни за 1 см² склопакета:")
	group.SetChild(box)

	window.SetChild(group)
	window.Show()

	okButton.OnClicked(func(*ui.Button) {
		window.Destroy()
	})

	window.OnClosing(func(*ui.Window) bool {
		window.Hide()
		return true
	})
}

func main() {
	err := ui.Main(windowCalculator)
	if err != nil {
		fmt.Println("Помилка: ", err)
	}
}
