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

With the output:

<img src="https://user-images.githubusercontent.com/37668193/135746601-63e4a612-b4e3-4103-b18d-6b0b4e0f2af1.png" width=600 />
