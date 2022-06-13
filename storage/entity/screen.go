package entity

type PanelType string

const (
	Ips  PanelType = "IPS"
	Oled PanelType = "OLED"
)

type Screen struct {
	Id          uint32    `db:"id"`
	Width       int       `db:"width"`
	Height      int       `db:"height"`
	SizeInch    float32   `db:"size_inch"`
	Panel       PanelType `db:"panel"`
	MultitTouch bool      `db:"multitouch"`
}
