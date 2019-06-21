package filter

type Criteria interface {
	doFilter(string)
}

type Filter struct {
	name     string
	criteria Criteria
}

type CriteriaValue struct {
	key   string
	value interface{}
}

type Equals struct {
}

func (equals Equals) doFilter(value CriteriaValue) string {

	return "{\"" + value.key + "\":" + "{\"eq\": \"" + value.value.(string) + "\"}}"
}
