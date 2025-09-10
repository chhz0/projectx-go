package main

import (
	"fmt"

	"github.com/chhz0/projectx-go/pkg/version"
)

func main() {
	fmt.Printf("--> string \n %s\n", version.String())
	fmt.Printf("--> text \n %s\n", version.Text())
	if str, err := version.JSON(); err == nil {
		fmt.Printf("--> json \n %s\n", str)
	}
}
