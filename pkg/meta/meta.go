package meta

import (
	"os"
	"strconv"
)

type Meta struct {
	TotalCount int `json:"total_count"`
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	PageCount  int `json:"page_count"`
}

func New(page, pageSize, total int) (*Meta, error) {
	if pageSize <= 0 {
		var err error
		pageSize, err = strconv.Atoi(os.Getenv("DEFAULT_PAGE_SIZE"))
		if err != nil {
			return nil, err
		}
	}

	//Obtenemos el total de paginas
	pageCount := 0
	if total > 0 {
		pageCount = (total + pageSize - 1) / pageSize
		if page > pageCount {
			page = pageCount
		}
	}
	// Validamos que no nos pasen un page menor a 1
	if page < 1 {
		page = 1
	}
	return &Meta{TotalCount: total,
		Page:      page,
		PageSize:  pageSize,
		PageCount: pageCount,
	}, nil
}

func (p *Meta) Offset() int {
	return (p.Page - 1) * p.PageSize
}

func (p *Meta) Limit() int {
	return p.PageSize
}
