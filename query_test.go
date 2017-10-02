package pagination

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuery(t *testing.T) {
	q := NewQuery(0, 0)
	assert.Equal(t, DefaultPage, q.Page)
	assert.Equal(t, DefaultPerPage, q.PerPage)

	q = NewQuery(2, 5)
	assert.Equal(t, 2, q.Page)
	assert.Equal(t, 5, q.PerPage)
}
