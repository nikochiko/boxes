package main

import (
	"fmt"

	"github.com/nikochiko/boxes/boxes"
)

func main() {
	content := `Content-Type: application/json
Accept: */*
X-API-KEY: kyquasdfoiqweozx123asdf923
`

	b := boxes.BoxWithContent{
		Box: boxes.Box{
			Width:  100,
			Height: 10,
		},
		Content: content,
	}

	fmt.Print(b.String())
}
