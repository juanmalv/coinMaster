package models

//Income model details the data included on a basic income saving
type Income struct {
	Amount   *float64 `json:"amount"`
	Date     string   `json:"date"`
	Reason   string   `json:"reason"`
	Category string   `json:"category"`
}
