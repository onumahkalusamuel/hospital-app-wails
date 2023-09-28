package pkg

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Pagination struct {
	Limit      int         `json:"limit,omitempty;query:limit"`
	Page       int         `json:"page,omitempty;query:page"`
	SortKey    string      `json:"sort_key,omitempty;query:sort_key"`
	SortOrder  string      `json:"sort_order,omitempty;query:sort_order"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	Query      string      `json:"query"`
	Rows       interface{} `json:"rows"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}
func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	if p.SortKey == "" {
		p.SortKey = "created_at"
	}
	if p.SortOrder == "" {
		p.SortOrder = "desc"
	}

	return fmt.Sprintf("%s %s", p.SortKey, p.SortOrder)
}

func GrabFromContext(p Pagination, c echo.Context) Pagination {
	if c.QueryParam("limit") != "" {
		p.Limit, _ = strconv.Atoi(c.QueryParam("limit"))
	}

	if c.QueryParam("page") != "" {
		p.Page, _ = strconv.Atoi(c.QueryParam("page"))
	}

	if c.QueryParam("query") != "" {
		p.Query = c.QueryParam("query")
	}

	return p
}
