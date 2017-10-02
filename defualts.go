package pagination

//defaults
var (
	DefaultPage    = 1
	DefaultPerPage = 10
	MaxPerPage     = 100
)

//SetDefaultPage ...
func SetDefaultPage(page int) {
	DefaultPage = page
}

//SetDefaultPerPage ...
func SetDefaultPerPage(perPage int) {
	DefaultPerPage = perPage
}

//SetMaxPerPage ...
func SetMaxPerPage(max int) {
	MaxPerPage = max
}
