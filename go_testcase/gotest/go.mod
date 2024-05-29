module gotest

go 1.22.2

replace (
	add_sub => ../add_sub
	multi_div => ../multi_div
)

require (
	add_sub v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.9.0
	multi_div v0.0.0-00010101000000-000000000000
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
