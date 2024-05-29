module main

go 1.22.2

replace (
	add_sub => ./add_sub
	multi_div => ./multi_div
)

require (
	add_sub v0.0.0-00010101000000-000000000000
	multi_div v0.0.0-00010101000000-000000000000
)
