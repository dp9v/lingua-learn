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
	NextBtn          *widget.Button
	Input            *OnKeyEntry
	TranslationLabel *widget.Label
	CorrectWordLabel *widget.Label
}

func (a *ShowWordsActivity) GetContent() fyne.CanvasObject {
	a.TranslationLabel.SetText(a.RoundWords[a.currentWord].Translation)
	return container.NewBorder(
		a.TranslationLabel,
		container.NewHBox(layout.NewSpacer(), a.NextBtn),
		nil,
		nil,
		container.NewVBox(a.Input, a.CorrectWordLabel),
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
		TranslationLabel: widget.NewLabel(""),
		CorrectWordLabel: widget.NewLabel(""),
	}
	activity.Input = NewOnKeyEntry(activity.onNextBtnClick)
	activity.NextBtn = widget.NewButton("Next", activity.onNextBtnClick)
	return &activity
}

func (a *ShowWordsActivity) onNextBtnClick() {
	var answer = strings.ToLower(a.Input.Text)
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
	a.CorrectWordLabel.SetText(a.RoundWords[a.currentWord].Original)
}

func (a *ShowWordsActivity) nextWord() {
	if len(a.CorrectWordLabel.Text) != 0 {
		a.RoundWords = append(a.RoundWords, a.RoundWords[a.currentWord])
	}
	a.currentWord++
	a.CorrectWordLabel.SetText("")
	a.TranslationLabel.SetText(a.RoundWords[a.currentWord].Translation)
	a.Input.SetText("")
}

func (a *ShowWordsActivity) focusInput() {
	a.app.w.Canvas().Focus(a.Input)
}
