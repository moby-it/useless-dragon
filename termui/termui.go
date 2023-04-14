package termui

import (
	"fmt"
	"useless_dragon/combat"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	boxWidth        = 32
	enemyBoxWidth   = 46
	boxHeight       = 8
	red             = lipgloss.Color("9")
	blue            = lipgloss.Color("27")
	gray            = lipgloss.Color("255")
	hover           = lipgloss.Color("228")
	updateCombatUI  = "updateCombatUI"
	ActionsView     = "actions"
	PlayerStatsView = "playerStats"
	EnemyStatsView  = "enemyStats"
)

var (
	bold            = lipgloss.NewStyle().Bold(true)
	headerStyle     = bold.Copy().Foreground(gray).MarginBottom(1)
	hoverSyle       = lipgloss.NewStyle().Foreground(hover)
	boxStyle        = lipgloss.NewStyle().MarginBottom(1).MarginTop(1).BorderStyle(lipgloss.NormalBorder()).Width(boxWidth).Height(boxHeight)
	focusedBoxStyle = boxStyle.Copy().BorderForeground(blue).Width(boxWidth).Height(boxHeight)
	statsStyle      = lipgloss.NewStyle().PaddingLeft(1).PaddingRight(1).Width(boxWidth).Height(boxHeight)
	actionsStyle    = lipgloss.NewStyle().PaddingLeft(1).PaddingRight(1).Width(boxWidth).Height(5)
)

type model struct {
	combat         *combat.Combat
	cursor         int
	selectedAction int
	combatChan     chan *combat.Combat
}

func InitialModel(cUpdater chan *combat.Combat) model {
	return model{
		cursor:         0,
		selectedAction: -1,
		combatChan:     cUpdater,
	}
}
func ListenForCombatChanges(m model) tea.Cmd {
	return func() tea.Msg {
		<-m.combat.UpdateUi
		return updateCombatUI
	}
}
func ListenForCombatCompletion(m model) func() tea.Msg {
	return func() tea.Msg {
		c := <-m.combatChan
		return c
	}
}
func (m model) Init() tea.Cmd {
	return ListenForCombatCompletion(m)
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Is it a key press?
	case *combat.Combat:
		m = model{
			combat:         msg,
			cursor:         0,
			selectedAction: -1,
			combatChan:     m.combatChan,
		}
	case string:
		if msg == updateCombatUI {
			m.cursor = 0
			m.selectedAction = -1
		}
	case tea.KeyMsg:
		if m.combat == nil {
			return m, nil
		}
		key := msg.String()
		// Cool, what was the actual key pressed?
		switch key {
		case "down":
			if m.selectedAction < 0 && m.cursor < 2 {
				m.cursor++
			}
			if m.selectedAction >= 0 && m.cursor < len(m.combat.Enemies)-1 {
				m.cursor++
			}
		case "up":
			if m.selectedAction < 0 && m.cursor > 0 {
				m.cursor--
			}
			if m.selectedAction >= 0 && m.cursor > 0 {
				m.cursor--
			}
		case "enter":
			if m.selectedAction < 0 {
				m.selectedAction = m.cursor
				// guard has no target
				if m.cursor == 2 {
					playerAction := combat.PlayerAction{
						Action: combat.CreateGuard(),
					}
					m.combat.PlayerActionChan <- playerAction
				}
				m.cursor = 0
			} else {
				playerAction, err := m.extractPlayerAction()
				if err != nil {
					return m, nil
				}
				m.combat.PlayerActionChan <- playerAction
			}
		case "esc":
			if m.selectedAction >= 0 {
				m.selectedAction = -1
				m.cursor = 0
			}
		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	if m.combat != nil {
		if m.combat.Status != combat.Playing {
			m.combat = nil
			return m, ListenForCombatCompletion(m)
		} else {
			return m, ListenForCombatChanges(m)
		}
	}
	return m, nil
}
func (m model) View() string {
	if m.combat == nil {
		return ""
	}
	s := boxStyle.Render(renderCombatant(m.combat.Player))
	s += renderActions(m)
	s = lipgloss.JoinHorizontal(lipgloss.Top, s, renderEnemies(m))
	return s
}

func (m model) extractPlayerAction() (combat.PlayerAction, error) {
	var playerAction combat.PlayerAction
	enemy := m.combat.Enemies[m.cursor]
	if enemy.Health <= 0 {
		return playerAction, fmt.Errorf("enemy has no hp")
	}
	switch m.selectedAction {
	case 0:
		playerAction = combat.PlayerAction{
			Action: combat.CreateBasicAttack(),
			Target: enemy,
		}
	case 1:
		playerAction = combat.PlayerAction{
			Action: combat.CreatePowerAttack(),
			Target: enemy,
		}
	case 2:
		playerAction = combat.PlayerAction{
			Action: combat.CreateGuard(),
			Target: enemy,
		}
	}
	return playerAction, nil
}
