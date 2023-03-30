package dto

type SearchEntry struct {
	Filters  string
	Includes string
	Sort     string
	Page     int `binding:"min=0"`
	PageSize int `binding:"min=0"`
}
