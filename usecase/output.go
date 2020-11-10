package usecase

type Output interface {
	Out() error
}

type output struct {
}

func NewOutput() Output {
	return &output{}
}

func (o *output) Out() error {
	return nil
}
