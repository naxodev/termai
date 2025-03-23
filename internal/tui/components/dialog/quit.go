package dialog

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kujtimiihoxha/termai/internal/tui/components/core"
	"github.com/kujtimiihoxha/termai/internal/tui/layout"
	"github.com/kujtimiihoxha/termai/internal/tui/styles"
	"github.com/kujtimiihoxha/termai/internal/tui/util"

	"github.com/charmbracelet/huh"
)

const question = "Are you sure you want to quit?"

var (
	width  = lipgloss.Width(question) + 6
	height = 3
)

type QuitDialog interface {
	tea.Model
	layout.Sizeable
	layout.Bindings
}

type quitDialogCmp struct {
	form   *huh.Form
	width  int
	height int
}

func (q *quitDialogCmp) Init() tea.Cmd {
	return nil
}

func (q *quitDialogCmp) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	// Process the form
	form, cmd := q.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		q.form = f
		cmds = append(cmds, cmd)
	}

	if q.form.State == huh.StateCompleted {
		v := q.form.GetBool("quit")
		if v {
			return q, tea.Quit
		}
		cmds = append(cmds, util.CmdHandler(core.DialogCloseMsg{}))
	}

	return q, tea.Batch(cmds...)
}

func (q *quitDialogCmp) View() string {
	return q.form.View()
}

func (q *quitDialogCmp) GetSize() (int, int) {
	return q.width, q.height
}

func (q *quitDialogCmp) SetSize(width int, height int) {
	q.width = width
	q.height = height
}

func (q *quitDialogCmp) BindingKeys() []key.Binding {
	return q.form.KeyBinds()
}

func newQuitDialogCmp() QuitDialog {
	confirm := huh.NewConfirm().
		Title(question).
		Affirmative("Yes!").
		Key("quit").
		Negative("No.")

	theme := styles.HuhTheme()
	theme.Focused.FocusedButton = theme.Focused.FocusedButton.Background(styles.Warning)
	theme.Blurred.FocusedButton = theme.Blurred.FocusedButton.Background(styles.Warning)
	form := huh.NewForm(huh.NewGroup(confirm)).
		WithWidth(width).
		WithHeight(height).
		WithShowHelp(false).
		WithTheme(theme).
		WithShowErrors(false)
	confirm.Focus()
	return &quitDialogCmp{
		form:  form,
		width: width,
	}
}

func NewQuitDialogCmd() tea.Cmd {
	content := layout.NewSinglePane(
		newQuitDialogCmp().(*quitDialogCmp),
		layout.WithSignlePaneSize(width+2, height+2),
		layout.WithSinglePaneBordered(true),
		layout.WithSinglePaneFocusable(true),
		layout.WithSinglePaneActiveColor(styles.Warning),
	)
	content.Focus()
	return util.CmdHandler(core.DialogMsg{
		Content: content,
	})
}
