package main

import (
	"fmt"
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"strconv"
)

func calculateTravelCost(days int, countryIndex, seasonIndex int, isLux bool, isGuide bool) float64 {
	var pricePerDay float64

	switch countryIndex {
	case 0: // Болгарія
		if seasonIndex == 0 {
			pricePerDay = 100.0 
		} else if seasonIndex == 1 {
			pricePerDay = 150.0
		}
	case 1: // Німеччина
		if seasonIndex == 0 {
			pricePerDay = 160.0
		} else if seasonIndex == 1 {
			pricePerDay = 200.0
		}
	case 2: // Польща
		if seasonIndex == 0 {
			pricePerDay = 120.0 
		} else if seasonIndex == 1 {
			pricePerDay = 180.0 
		}
	}

	totalCost := float64(days) * pricePerDay

	if isGuide {
		totalCost += 50.0 * float64(days)
	}

	if isLux {
		totalCost += 50.0 * float64(days)
	}

	return totalCost
}

func displayPriceList() {
	window := ui.NewWindow("Ціни на тури", 0, 0, false)
	window.SetMargined(true)

	box := ui.NewVerticalBox()
	box.SetPadded(true)

	priceLabels := []string{
		"Болгарія, літо – $100/день",
		"Болгарія, зима – $150/день",
		"Німеччина, літо – $160/день",
		"Німеччина, зима – $200/день",
		"Польща, літо – $120/день",
		"Польща, зима – $180/день",
		"Вартість індивідуального гіда – $50/день",
		"Вартість націнки за люксовий номер – $50/день",
	}

	for _, label := range priceLabels {
		box.Append(ui.NewLabel(label), false)
	}

	okButton := ui.NewButton("OK")
	box.Append(okButton, false)

	group := ui.NewGroup("Ціни на подорож:")
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

func travelCalculator() {
	mainWindow := ui.NewWindow("Калькулятор вартості подорожі", 0, 0, false)
	mainWindow.SetMargined(true)

	box := ui.NewVerticalBox()
	box.SetPadded(true)

	paramBox := ui.NewGroup("Параметри подорожі:")
	paramBox.SetMargined(true)
	paramInputBox := ui.NewHorizontalBox()
	paramInputBox.SetPadded(true)

	form := ui.NewForm()
	form.SetPadded(true)

	daysEntry := ui.NewEntry()
	countryCombobox := ui.NewCombobox()
	countryCombobox.Append("Болгарія")
	countryCombobox.Append("Німеччина")
	countryCombobox.Append("Польща")

	seasonCombobox := ui.NewCombobox()
	seasonCombobox.Append("Літо")
	seasonCombobox.Append("Зима")

	form.Append("Кількість днів", daysEntry, false)
	form.Append("Країна", countryCombobox, false)
	form.Append("Сезон", seasonCombobox, false)

	paramInputBox.Append(form, false)
	paramBox.SetChild(paramInputBox)

	additionalBox := ui.NewGroup("Додаткові параметри:")
	additionalBox.SetMargined(true)

	additionalParamsBox := ui.NewVerticalBox()
	additionalParamsBox.SetPadded(true)

	luxCheckbox := ui.NewCheckbox("Номер люкс")
	guideCheckbox := ui.NewCheckbox("Індивідуальний гід")
	additionalParamsBox.Append(luxCheckbox, false)
	additionalParamsBox.Append(guideCheckbox, false)

	additionalBox.SetChild(additionalParamsBox)

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
		countryIndex := countryCombobox.Selected()
		seasonIndex := seasonCombobox.Selected()

		daysText := daysEntry.Text()
		days, err := strconv.Atoi(daysText)

		if err == nil {
			totalCost := calculateTravelCost(days, countryIndex, seasonIndex, luxCheckbox.Checked(), guideCheckbox.Checked())
			resultLabel.SetText(fmt.Sprintf("Вартість: %.2f USD", totalCost))
		} else {
			ui.MsgBoxError(mainWindow, "Помилка!", "Будь ласка, введіть коректні значення.")
		}
	})

	box.Append(paramBox, false)
	box.Append(additionalBox, false)
	box.Append(resultBox, false)
	box.Append(buttonBox, false)

	mainWindow.SetChild(box)
	mainWindow.Show()

	mainWindow.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
}

func main() {
	err := ui.Main(travelCalculator)
	if err != nil {
		fmt.Println("Помилка: ", err)
	}
}
