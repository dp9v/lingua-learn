package gui

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"learn_words/datasources/v2/models"
	"strings"
)

type ShowWordsActivity struct {
	RoundWords       models.WordList
	app              *Application
	currentWord      int
	nextBtn          *widget.Button
	input            *OnKeyEntry
	translationLabel *widget.Label
	correctWordLabel *widget.Label
}

func (a *ShowWordsActivity) GetContent() fyne.CanvasObject {
	a.translationLabel.SetText(a.RoundWords[a.currentWord].Translation)
	return container.NewBorder(
		a.translationLabel,
		container.NewHBox(layout.NewSpacer(), a.nextBtn),
		nil,
		nil,
		container.NewVBox(a.input, a.correctWordLabel),
	)
}

func (a *ShowWordsActivity) GetTitle() string {
	return "Show words"
}

func NewShowWordsActivity(app *Application, words models.WordList) *ShowWordsActivity {
	activity := ShowWordsActivity{
		app:              app,
		currentWord:      0,
		RoundWords:       words,
		translationLabel: widget.NewLabel(""),
		correctWordLabel: widget.NewLabel(""),
	}
	activity.input = NewOnKeyEntry(activity.onNextBtnClick)
	activity.nextBtn = widget.NewButton("Next", activity.onNextBtnClick)
	return &activity
}

func (a *ShowWordsActivity) onNextBtnClick() {
	defer a.focusInput()
	if strings.ToUpper(a.input.Text) != strings.ToUpper(a.RoundWords[a.currentWord].Original) {
		a.correctWordLabel.SetText(a.RoundWords[a.currentWord].Original)
		return
	}
	if a.currentWord+1 == len(a.RoundWords) {
		dialog.ShowError(errors.New("Слова закончились"), a.app.w)
		return
	}

	a.nextWord()
}

func (a *ShowWordsActivity) nextWord() {
	if len(a.correctWordLabel.Text) != 0 {
		a.RoundWords = append(a.RoundWords, a.RoundWords[a.currentWord])
	}
	a.currentWord++
	a.correctWordLabel.SetText("")
	a.translationLabel.SetText(a.RoundWords[a.currentWord].Translation)
	a.input.SetText("")
}

func (a *ShowWordsActivity) focusInput() {
	a.app.w.Canvas().Focus(a.input)
}
