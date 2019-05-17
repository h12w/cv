module github.com/h12w/cv

go 1.12

require (
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	gopkg.in/russross/blackfriday.v2 v2.0.1
	h12.io/cv v0.0.0-20181124220903-66dfd6721900
)

replace gopkg.in/russross/blackfriday.v2 v2.0.1 => github.com/russross/blackfriday/v2 v2.0.1

replace h12.io/cv => ./
