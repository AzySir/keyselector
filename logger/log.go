package logger

type Verbose interface {
	Verbose() string
	SetVerbose(s string)
}

type Enabled struct {
	name string
}

func (v *Enabled) GetVerbose() string {
	return v.name
}

func (v *Enabled) SetVerbose(s string) {
	v.name = s
}
