package pagination

//Query includes struct tags as helpers for different data formats
type Query struct {
	Page    int `json:"page" schema:"page" bson:"page" query:"page"`
	PerPage int `json:"per_page" schema:"per_page" bson:"per_page" query:"per_page"`
}

//NewQuery ...
func NewQuery(page, perPage int) Query {
	q := Query{
		Page:    page,
		PerPage: perPage,
	}
	q.SetDefaults()
	return q
}

//SetDefaults ...
func (q *Query) SetDefaults() {
	if q.Page <= 0 {
		q.Page = DefaultPage
	}
	if q.PerPage <= 0 || (MaxPerPage > 0 && q.PerPage > MaxPerPage) {
		q.PerPage = DefaultPerPage
	}
}
