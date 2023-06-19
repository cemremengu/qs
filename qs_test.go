package qs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	qs := "sort[0]=title:asc&filters[title][$eq]=hello&filters[col][$in][0]=hello&filters[col][$in][1]=world&fields[0]=title&pagination[limit]=10&pagination[offset]=1"

	q := Parse(qs)

	assert.Equal(t, Query{
		Sort: []string{"title:asc"},
		Filters: []filter{
			{
				Value: "hello",
				Param: "title",
				Op:    "$eq",
			},
			{
				Value: []interface{}{"hello", "world"},
				Param: "col",
				Op:    "$in",
			},
		},
		Fields: []string{"title"},
		Pagination: pagination{
			Limit:  10,
			Offset: 1,
		},
	}, q)
}

func TestOnlyPagination(t *testing.T) {
	qs := "pagination[limit]=10&pagination[offset]=1"

	q := Parse(qs)

	assert.Equal(t, Query{
		Pagination: pagination{
			Limit:  10,
			Offset: 1,
		},
	}, q)
}

func TestOnlySort(t *testing.T) {
	qs := "sort[0]=title:asc&filters[title][$eq]=hello&filters[col][$in][0]=hello&filters[col][$in][1]=world&fields[0]=title&pagination[limit]=10&pagination[offset]=1"

	q := Parse(qs)

	assert.Equal(t, Query{
		Sort: []string{"title:asc"},
		Filters: []filter{
			{
				Value: "hello",
				Param: "title",
				Op:    "$eq",
			},
			{
				Value: []interface{}{"hello", "world"},
				Param: "col",
				Op:    "$in",
			},
		},
		Fields: []string{"title"},
		Pagination: pagination{
			Limit:  10,
			Offset: 1,
		},
	}, q)
}
