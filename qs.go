package qs

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

var (
	refilter = regexp.MustCompile(`\[(\w+)\]\[(\$\w+)\]`)
	reindex  = regexp.MustCompile(`\[(\d+)\]$`)
)

func parseFilter(key string) (string, string) {
	matches := refilter.FindStringSubmatch(key)

	return matches[1], matches[2]
}

type filter struct {
	Value interface{}
	Param string
	Op    string
}

type pagination struct {
	Limit  int
	Offset int
}

type Query struct {
	Sort       []string
	Filters    []filter
	Fields     []string
	Pagination pagination
}

func Parse(qs string) Query {
	u, _ := url.ParseQuery(qs)

	temp := map[string][]interface{}{}

	for key, value := range u {
		k := reindex.ReplaceAllString(key, "")
		if _, ok := temp[k]; !ok {
			temp[k] = []interface{}{value[0]}
		} else {
			temp[k] = append(temp[k], value[0])
		}
	}

	q := Query{Pagination: pagination{Limit: -1, Offset: 0}}

	for key, value := range temp {
		if strings.HasPrefix(key, "filter") {
			param, op := parseFilter(key)
			if op == "$in" {
				q.Filters = append(q.Filters, filter{value, param, op})
			} else {
				q.Filters = append(q.Filters, filter{value[0], param, op})
			}
		} else if strings.HasPrefix(key, "sort") {
			q.Sort = append(q.Sort, value[0].(string))
		} else if strings.HasPrefix(key, "fields") {
			q.Fields = append(q.Fields, value[0].(string))
		} else if strings.HasPrefix(key, "pagination[offset]") {
			q.Pagination.Offset, _ = strconv.Atoi(value[0].(string))
		} else if strings.HasPrefix(key, "pagination[limit]") {
			q.Pagination.Limit, _ = strconv.Atoi(value[0].(string))
		}
	}

	return q
}
