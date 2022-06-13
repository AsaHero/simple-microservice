package entity

type LayoutType string

const (
	Qwerty LayoutType = "QWERTY"
	Qwertz LayoutType = "QWERTZ"
	Azerty LayoutType = "AZERTY"
)

type Keyboard struct {
	Id uint32 `db:"id"`
	Layout LayoutType `db:"layout"`
	Backlit bool `db:"backlit"`
}
