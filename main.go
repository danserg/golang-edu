// package main contains an example on how to use the ReadForm, but with the same way you can do the ReadJSON & ReadJSON
// you may see on other frameworks binders and all that, Iris has simplicity on it's mind, so I think this way is more simple than others.
package main

import (
	"fmt"

	"github.com/flosch/pongo2"

	"github.com/kataras/iris"
)

type Visitor struct {
	Username string
	Mail     string
	Data     []string `form:"mydata"`
}

func main() {

	iris.Get("/", func(ctx *iris.Context) {
		ctx.SetContentType("text/html ;charset=UTF-8")
		formTemplate.ExecuteWriter(nil, ctx.Response.BodyWriter())
	})

	iris.Post("/form_action", func(ctx *iris.Context) {
		visitor := &Visitor{}
		err := ctx.ReadForm(visitor)
		if err != nil {
			fmt.Println("Error when reading form: " + err.Error())
		}
		fmt.Printf("\n Visitor: %v", visitor.Data)
	})

	fmt.Println("Server is running at :8080")
	iris.Listen(":8080")
}

var formTemplate, err = pongo2.FromString(`
<!DOCTYPE html>
<head>
<meta charset="utf-8">
</head>
<body>
<form action="/form_action" method="post">
<input type="text" name="Username" />
<br/>
<input type="text" name="Mail" /><br/>
<select multiple="multiple" name="mydata">
<option value='one'>One</option>
<option value='two'>Two</option>
<option value='three'>Three</option>
<option value='four'>Four</option>
</select>
<hr/>
<input type="submit" value="Send data" />

</form>
</body>
</html>
`)
