package main

import (
	"interfaces/src/generated"
)

func main() {
	engine := generated.GinEngine()
	err := engine.Run(":8080")
	if err != nil {
		print(err)
	}
}
