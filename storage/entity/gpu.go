package entity

type GPU struct {
	Id       uint32  `db:"id"`
	Brand    string  `db:"brand"`
	Name     string  `db:"name"`
	MinGhz   float32 `db:"min_ghz"`
	MaxGhz   float32 `db:"max_ghz"`
	MemoryId uint32  `db:"memory_id"`
}
