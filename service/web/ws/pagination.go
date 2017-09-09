package ws

import (
	"net/http"
	"strconv"

	"upper.io/db.v3"
)

type Page struct {
	Entries uint64      `json:"total_entries"`
	Pages   uint        `json:"total_pages"`
	Page    uint        `json:"current_page"`
	PerPage uint        `json:"entries_per_page"`
	Data    interface{} `json:"data"`
}

func Paginate(r *http.Request, res db.Result, dest interface{}) (*Page, error) {
	var goNext bool
	if val := r.URL.Query().Get("next"); val != "" {
		goNext = true
	}

	var goPrev bool
	if val := r.URL.Query().Get("prev"); val != "" {
		goPrev = true
	}

	var page = uint(1)
	if val := r.URL.Query().Get("page"); val != "" {
		n, err := strconv.Atoi(val)
		if err == nil && n > 0 {
			page = uint(n)
		}
	}

	var cursor int
	if val := r.URL.Query().Get("cursor"); val != "" {
		n, err := strconv.Atoi(val)
		if err == nil && n > 0 {
			cursor = n
		}
	}

	var pageSize = uint(20)
	if val := r.URL.Query().Get("limit"); val != "" {
		n, err := strconv.Atoi(val)
		if err == nil && n > 0 && n <= 100 {
			pageSize = uint(n)
		}
	}

	res = res.Paginate(pageSize)

	pages, err := res.TotalPages()
	if err != nil {
		return nil, err
	}

	entries, err := res.TotalEntries()
	if err != nil {
		return nil, err
	}

	switch {
	case goNext && cursor >= 0:
		res = res.Cursor("id")
		if err := res.NextPage(cursor).All(dest); err != nil {
			return nil, err
		}
	case goPrev && cursor > 0:
		res = res.Cursor("id")
		if err := res.PrevPage(cursor).All(dest); err != nil {
			return nil, err
		}
	default:
		if err := res.Page(page).All(dest); err != nil {
			return nil, err
		}
	}

	return &Page{
		Page:    page,
		PerPage: pageSize,
		Pages:   pages,
		Entries: entries,
		Data:    dest,
	}, nil
}
