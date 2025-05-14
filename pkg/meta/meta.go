package meta

import (
	"os"
	"strconv"
)

type Meta struct {
	TotalCount int `json:"total_count"`
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
}

func New(page, perPage, total int) (*Meta, error) {
	if perPage <= 0 {
		var err error
		perPage, err = strconv.Atoi(os.Getenv("DEFAULT_PAGE_SIZE"))
		if err != nil {
			return nil, err
		}
	}

	return &Meta{TotalCount: total}, nil
}
