package briq

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/tomnomnom/linkheader"
)

var (
	DefaultPage  = 1
	DefaultLimit = 100
)

type (
	Pagination struct {
		Page    int
		PerPage int
	}

	Link struct {
		First    *Pagination
		Last     *Pagination
		Previous *Pagination
		Next     *Pagination
	}
)

func Page(i ...int) Pagination {
	var page, limit = DefaultPage, DefaultLimit

	if len(i) > 0 {
		page = i[0]
	}

	if len(i) > 1 {
		limit = i[1]
	}

	return Pagination{
		Page:    page,
		PerPage: limit,
	}
}

func (p Pagination) query() string {
	return fmt.Sprintf("?page=%d&per_page=%d", p.Page, p.PerPage)
}

func (l Link) HasNext() bool {
	if l.Next == nil {
		return false
	}

	return true
}

func (l Link) HasPrevious() bool {
	if l.Previous == nil {
		return false
	}

	return true
}

func linkFromResponse(r *http.Response) *Link {
	link := &Link{}

	ll := linkheader.Parse(r.Header.Get("Link"))

	setPagination := func(p **Pagination, rel string) {
		if l := ll.FilterByRel(rel); len(l) > 0 {

			u, err := url.Parse(l[0].URL)
			if err != nil {
				return
			}

			page, _ := strconv.Atoi(u.Query().Get("page"))
			pp, _ := strconv.Atoi(u.Query().Get("per_page"))

			*p = &Pagination{
				Page:    page,
				PerPage: pp,
			}
		}
	}

	for rel, pagination := range map[string]**Pagination{
		"first":    &link.First,
		"last":     &link.Last,
		"previous": &link.Previous,
		"next":     &link.Next,
	} {
		setPagination(pagination, rel)
	}

	return link
}
