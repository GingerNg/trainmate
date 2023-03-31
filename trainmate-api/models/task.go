package models

type Task struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Version  string `json:"version"`
	Desp     string `json:"desp"`
	MetaInfo string `json:"meta_info"`
}

// var taskTblName = "Task"
// func (m *Task) IsValid() bool {
// 	return m.validateLanguage() && m.validateExtension()
// }

// func (m *Task) validateLanguage() bool {
// 	return len(m.Language) >= 1
// }

// func (m *Task) validateExtension() bool {
// 	return len(m.Extension) >= 1
// }
