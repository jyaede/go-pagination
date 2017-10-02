package pagination

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaults(t *testing.T) {

	SetDefaultPage(2)
	SetDefaultPerPage(20)
	SetMaxPerPage(50)

	assert.Equal(t, 2, DefaultPage)
	assert.Equal(t, 20, DefaultPerPage)
	assert.Equal(t, 50, MaxPerPage)
}
