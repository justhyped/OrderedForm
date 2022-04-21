# OrderedForm
## Installation
Run the following command in your project folder:
``go get github.com/justhyped/OrderedForm``

## Usage
```go
package main

import (
	"github.com/justhyped/OrderedForm"
	"net/http"
	"strings"
)

func main() {
	// create the form
	form := new(OrderedForm.OrderedForm)
	form.Set("key", "value")
	form.Set("key1", "value1")

	// create a post request
	req, _ := http.NewRequest("POST", "url here", strings.NewReader(form.URLEncode()))

	// form.URLEncode will return key=value&key1=value1. the values are properly query escapep
	// and it's order is maintained.	
}
```
## License
[MIT](https://choosealicense.com/licenses/mit/)