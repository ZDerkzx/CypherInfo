package models

type Money struct {
	Dollars int64 `json:"dollars"`
	Cents   int32 `json:"cents"`
}
