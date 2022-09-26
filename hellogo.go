package studygit

type Hello struct {
	Name string
	Say  string
}

func New() *Hello {
	return &Hello{}
}

func (h *Hello) GetName() string {
	return h.Name
}

func (h *Hello) GetSay() string {
	return h.Say
}

func (h *Hello) Set(name, say string) {
	h.Name = name
	h.Say = say
}
