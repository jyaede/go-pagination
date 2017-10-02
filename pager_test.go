package pagination

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testItems = []interface{}{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
}

func TestPager(t *testing.T) {
	assert := assert.New(t)
	p := NewPager(3, 10, 400, testItems...)
	assert.Equal(3, p.Page)
	assert.Equal(10, p.PerPage)
	assert.Equal(21, p.From)
	assert.Equal(30, p.To)
	assert.Equal(10, p.Count)
	assert.Equal(400, p.Total)
	assert.Equal(40, p.Pages)
	assert.Nil(p.Items)

	p = NewPager(1, 10, 400, testItems...)
	assert.Equal(1, p.Page)
	assert.Equal(10, p.PerPage)
	assert.Equal(1, p.From)
	assert.Equal(10, p.To)
	assert.Equal(10, p.Count)
	assert.Equal(400, p.Total)
	assert.Equal(40, p.Pages)
	assert.Nil(p.Items)

	p.SetItems(1, 2)
	assert.Len(p.Items, 2)
	p.AddItems(2, 4)
	assert.Len(p.Items, 4)
	p.AddItems(5)
	assert.Len(p.Items, 5)
}
