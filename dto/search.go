package dto

type SearchEntry struct {
	Filters  string
	Includes string
	Sort     string
	Page     int `binding:"min=0"`
	PageSize int `binding:"required,min=1,max=1000"`
}

type SearchFilters map[string]map[string]map[string]string
