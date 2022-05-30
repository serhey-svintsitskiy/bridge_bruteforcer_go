package main

import (
	"bridge_bruteforcer/src/bruteforcer"
	"fmt"
)

func main() {
	app := new(bruteforcer.Bruteforcer)
	app.Bruteforce()
	bestResult := app.GetBestResult()

	fmt.Println("Best result: ")
	fmt.Println("Time: ", bestResult.Time)
	fmt.Println("MovesLog: ")
	for i, moveLog := range bestResult.MovesLog {
		fmt.Println(i, "|", moveLog)
	}
}
