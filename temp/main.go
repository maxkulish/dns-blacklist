package main

import (
	"time"

	"github.com/maxkulish/dns-blacklist/pbar"
)

func main() {
	var bar pbar.Bar
	bar.NewOptionWithGraph(0, 100, "#")

	for i := 0; i <= 100; i++ {
		time.Sleep(100 * time.Millisecond)
		bar.Play(i)
	}
	bar.Finish()
}
