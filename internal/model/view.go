package model

func (m Model) View() string {
	switch m.Mode {
	case LANDING:
		return m.viewLanding()
	case CREATE:
		// TODO:
		return "TODO: create note view"
	case SEARCH:
		return m.viewSearch()
	}

	return ""
}
