package entity

type Units string

const (
	Unknown  Units = "UNKNOWN"
	Bit      Units = "BIT"
	Byte     Units = "BYTE"
	Kilobyte Units = "KILOBYTE"
	Megabyte Units = "MEGABAYTE"
	Gigabyte Units = "GIGABYTE"
	Terabyte Units = "TERABYTE"
)

type Memory struct {
	Id    uint32   `db:"id"`
	Value int64 `db:"value"`
	Unit  Units `db:"unit"`
}
