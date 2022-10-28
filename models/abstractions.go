package models

// Settings ...
type Settings struct {
	Class        string
	Id           string
	Col1         *Div
	Col3         *Div
	SettingsForm *Form
	Category     string
}

// NewSettings instantiates a default Settings Abstract
func NewSettings(class string, id string, form *Form, col1 *Div, col3 *Div) *Settings {
	return &Settings{
		Class:        class,
		Id:           id,
		Col1:         col1,
		Col3:         col3,
		SettingsForm: form,
		Category:     "default",
	}
}

// Overview ...
type Overview struct {
	Class           string
	Id              string
	primaryFilter   *InputField // i.e. Group Filter
	secondaryFilter *InputField // i.e. User Filter
	Col1            *Div
	Col2            *Div
	Col3            *Div
	Category        string
}

// NewOverview instantiates a default Overview Abstract
func NewOverview(class string, id string, col1 *Div, col2 *Div, col3 *Div) *Overview {
	return &Overview{
		Class:    class,
		Id:       id,
		Col1:     col1,
		Col2:     col2,
		Col3:     col3,
		Category: "default",
	}
}

// NewTasksOverview instantiates a task Overview Abstract	// TODO NEXT - START HERE, called in TASK CONTROLLER
func NewTasksOverview(class string, id string, userFilter *InputField, groupFilter *InputField, col1 *Div, col2 *Div, col3 *Div) *Overview {
	return &Overview{
		Class:           class,
		Id:              id,
		primaryFilter:   groupFilter,
		secondaryFilter: userFilter,
		Col1:            col1,
		Col2:            col2,
		Col3:            col3,
		Category:        "tasks",
	}
}
