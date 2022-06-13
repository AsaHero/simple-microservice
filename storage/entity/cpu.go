package entity

type CPU struct {
	Id            uint32     `db:"id"`
	Brand         string  `db:"brand"`
	Name          string  `db:"name"`
	NumberCores   int32   `db:"number_cores"`
	NumberThreads int32   `db:"number_threads"`
	MinGhz        float32 `db:"min_ghz"`
	MaxGhz        float32 `db:"max_ghz"`
}
