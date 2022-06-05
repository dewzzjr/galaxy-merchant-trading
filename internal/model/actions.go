package model

type Action int

const (
	ActionUnknown Action = iota
	ActionDefine
	ActionStatement
	ActionQuestion
)

func (a Action) String() string {
	return [...]string{
		"unknown",
		"define",
		"statement",
		"question",
	}[a]
}
