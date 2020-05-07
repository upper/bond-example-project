package app

import (
	"context"

	"github.com/upper/db"
)

type Page struct {
	Entries uint64      `json:"total_entries"`
	Pages   uint        `json:"total_pages"`
	Page    uint        `json:"current_page"`
	PerPage uint        `json:"entries_per_page"`
	Data    interface{} `json:"data"`
}

func PaginateQuery(ctx context.Context, res db.Result, dest interface{}) (*Page, error) {
	goNext := ctx.Value(ContextPaginationNext).(bool)
	goPrev := ctx.Value(ContextPaginationPrev).(bool)

	pageNumber := ctx.Value(ContextPaginationPage).(uint)
	pageCursor := ctx.Value(ContextPaginationCursor).(uint)
	pageSize := ctx.Value(ContextPaginationSize).(uint)

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
	case goNext && pageCursor >= 0:
		res = res.Cursor("id")
		if err := res.NextPage(pageCursor).All(dest); err != nil {
			return nil, err
		}
	case goPrev && pageCursor > 0:
		res = res.Cursor("id")
		if err := res.PrevPage(pageCursor).All(dest); err != nil {
			return nil, err
		}
	default:
		if err := res.Page(pageNumber).All(dest); err != nil {
			return nil, err
		}
	}

	return &Page{
		Page:    pageNumber,
		PerPage: pageSize,
		Pages:   pages,
		Entries: entries,
		Data:    dest,
	}, nil
}
