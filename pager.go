package pagination

import (
	"math"
)

//Pager ...
type Pager struct {
	Page    int           `json:"page"`
	PerPage int           `json:"per_page"`
	From    int           `json:"from"`
	To      int           `json:"to"`
	Count   int           `json:"count"`
	Total   int           `json:"total,omitempty"`
	Pages   int           `json:"pages"`
	Items   []interface{} `json:"items"`
}

//NewPager ...
func NewPager(page, perPage, total int, items ...interface{}) *Pager {
	p := &Pager{
		Page:    page,
		PerPage: perPage,
		Count:   len(items),
		Total:   total,
		Pages:   calcPages(total, perPage),
	}
	p.From = p.from()
	p.To = p.to()
	if p.Count == 0 {
		p.Items = []interface{}{}
	}
	return p
}

//SetItems replaces any existing items and sets the new ones
func (p *Pager) SetItems(items ...interface{}) {
	p.Items = items
}

//AddItem ...
func (p *Pager) AddItem(item interface{}) {
	p.AddItems(item)
}

//AddItems ...
func (p *Pager) AddItems(items ...interface{}) {
	p.Items = append(p.Items, items...)
}

func (p Pager) from() int {
	if p.Total == 0 {
		return 0
	}
	return (p.Page-1)*p.PerPage + 1
}

func (p Pager) to() int {
	if p.Total == 0 {
		return 0
	}
	return p.from() + p.Count - 1
}

func calcPages(total, perpage int) int {
	if total == 0 {
		return 0
	}
	d := float64(total) / float64(perpage)
	return int(math.Ceil(d))
}
