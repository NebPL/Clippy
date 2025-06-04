package main

import (
	"fmt"
)

func (a *App) ParseInput(input string) string {
	return fmt.Sprintf("Hello %s, It's show time!", input)
}
