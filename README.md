# Boxes
Go library for creating boxes in terminal

## Usage
```go
package main

import (
	"fmt"

	"github.com/nikochiko/boxes/boxes"
)

func main() {
	content := `This is a multi-line string
spanning multiple lines
with multiple \n characters.
This content will be centered
inside the box by default.`

	box := boxes.BoxWithContent{
		Box: boxes.Box{
			Width: 50,
			Height: 10,
		},
		Content: content,
	}

	fmt.Println(box.String())
}
```
