package main

import (
	"fmt"

	"github.com/rivo/tview"
)

type toolbarView struct {
	element *tview.Flex
	main    *mainView
}

func newToolbarView(m *mainView) *toolbarView {
	t := &toolbarView{main: m}
	flex := tview.NewFlex().SetDirection(tview.FlexRowCSS)

	btnHelp := tview.NewButton(t.fmtBtn("F1", "Help"))
	flex.AddItem(btnHelp, 0, 1, false)
	flex.AddItem(tview.NewBox(), 1, 0, false)

	btnCopy := tview.NewButton(t.fmtBtn("F5", "Copy"))

	btnMove := tview.NewButton(t.fmtBtn("F6", "Move"))

	btnRename := tview.NewButton("Rename")
	btnRename.SetSelectedFunc(m.showRenameWin)

	btnQuit := tview.NewButton(t.fmtBtn("F10", "Quit"))
	btnQuit.SetSelectedFunc(func() { m.pages.ShowPage("modal") })

	flex.AddItem(btnCopy, 0, 1, false)
	flex.AddItem(tview.NewBox(), 1, 0, false)
	flex.AddItem(btnMove, 0, 1, false)
	flex.AddItem(tview.NewBox(), 1, 0, false)
	flex.AddItem(btnRename, 0, 1, false)
	flex.AddItem(tview.NewBox(), 1, 0, false)
	flex.AddItem(btnQuit, 0, 1, false)

	t.element = flex

	return t
}

func (t *toolbarView) fmtBtn(key string, text string) string {
	return fmt.Sprintf("[black:orange]%s[-:-] %s", key, text)
}