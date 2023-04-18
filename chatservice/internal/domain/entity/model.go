package entity

type Model struct {
	Name      string
	MaxTokens int
}

// *Model é um ponteiro
func NewModel(name string, maxTokens int) *Model {
	return &Model{
		Name:      name,
		MaxTokens: maxTokens,
	}
}

func (m *Model) GetMaxTokens() int {
	return m.MaxTokens
}

func (m *Model) GetModelName() string {
	return m.Name
}
