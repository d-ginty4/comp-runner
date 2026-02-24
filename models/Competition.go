package models

type Competition struct {
	Id              int         `json:"id"`
	Name            string      `json:"name"`
	EntryFee        interface{} `json:"entry_fee"`
	Location        interface{} `json:"location"`
	NumberOfEntries int         `json:"numberOfEntries"`
	Date            string      `json:"date"`
	Time            string      `json:"time"`
	Status          string      `json:"status"` // e.g. "upcoming", "live", "finished"
	Published       bool        `json:"published"`
}
