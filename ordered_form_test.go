package OrderedForm

import "testing"

const correctOutput = `foo=bar&example=test`

func TestOrderedForm(t *testing.T) {
	form := new(OrderedForm)
	form.Set("foo", "bar")
	form.Set("example", `test`)

	v := form.URLEncode()

	if v != correctOutput {
		t.Fatalf("expected:\n%v, got:\n%v", correctOutput, v)
	}
}
