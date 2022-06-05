package model

type Unit string

const (
	UnitSilver Unit = "Silver"
	UnitGold   Unit = "Gold"
	UnitIron   Unit = "Iron"
)

func (u Unit) Valid() bool {
	switch u {
	case UnitSilver, UnitGold, UnitIron:
		return true
	}
	return false
}
