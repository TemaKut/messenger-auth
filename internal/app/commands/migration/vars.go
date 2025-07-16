package migration

type Direction = string

const (
	DirectionUp     Direction = "up"
	DirectionDown   Direction = "down"
	DirectionDownTo Direction = "down-to"
)
