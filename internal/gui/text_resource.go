package gui

type TextResource struct {
	Text string
}

func NewStrResource(text string) *TextResource {
	return &TextResource{
		Text: text,
	}
}

func (s TextResource) Name() string {
	return s.Text
}

func (s TextResource) Content() []byte {
	return []byte{}
}
