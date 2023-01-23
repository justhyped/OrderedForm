package OrderedForm

import (
	"net/url"
	"strings"
)

// OrderedForm is a slice of [2]string, in which the first
// element is the key, and the second element is the value.
// Both the key and value are query escaped when using Set().
type OrderedForm [][2]string

func (o *OrderedForm) Set(k, v string) {
	// Set the key, and url encode the value
	m := [2]string{url.QueryEscape(k), url.QueryEscape(v)}
	*o = append(*o, m)
}

func (o *OrderedForm) URLEncode() string {
	var b strings.Builder
	for _, v := range *o {
		if b.Len() > 0 {
			b.WriteString("&")
		}
		b.WriteString(v[0])
		b.WriteString("=")
		b.WriteString(v[1])
	}

	return b.String()
}
