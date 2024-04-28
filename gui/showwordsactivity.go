package gui

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"learn_words/common"
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
	var answer = strings.ToLower(a.input.Text)
	var correctAnswer = strings.ToLower(a.RoundWords[a.currentWord].Original)
	if answer != correctAnswer {
		if common.Normalize(answer) == common.Normalize(correctAnswer) {
			dialog.NewConfirm(
				"Is correct",
				fmt.Sprintf("You answer is not correct because of diacritic."+
					"Do you want to accept it? \n Your answer is: %s \n Correct answer is: %s", answer, correctAnswer),
				a.onAnswer,
				a.app.w,
			).Show()
		} else {
			a.onAnswer(false)
		}
	} else {
		a.onAnswer(true)
	}
}

func (a *ShowWordsActivity) onAnswer(isCorrect bool) {
	defer a.focusInput()
	if isCorrect {
		a.onCorrectAnswer()
	} else {
		a.onWrongAnswer()
	}
}

func (a *ShowWordsActivity) onCorrectAnswer() {
	defer a.focusInput()
	if a.currentWord+1 == len(a.RoundWords) {
		dialog.ShowError(errors.New("no more words"), a.app.w)
		return
	}
	a.nextWord()
}

func (a *ShowWordsActivity) onWrongAnswer() {
	a.correctWordLabel.SetText(a.RoundWords[a.currentWord].Original)
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
