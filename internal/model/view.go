package model

func (m Model) View() string {
	switch m.Mode {
	case LANDING:
		return m.viewLanding()
	case QUICK:
		return m.viewQuick()
	case CREATE:
		return m.viewCreate()
	case SEARCH, BROWSE:
		return m.viewBrowse()
	}

	return ""
}
