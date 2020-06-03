package measurement

import "time"

//Probe is
type Probe struct {
	Sensor string    `json:"sensor"`
	Unit   string    `json:"unit"`
	Date   time.Time `json:"date"`
	Value  float32   `json:"value"`
}
