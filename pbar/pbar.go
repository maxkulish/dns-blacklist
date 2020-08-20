package pbar

import "fmt"

type Bar struct {
	percent int    // progress percentage
	cur     int    // current progress
	total   int    // total value of progress
	rate    string // the actual progress bar
	graph   string // the fill value for the progress bar
}

func (b *Bar) NewOption(start, total int) {
	b.cur = start
	b.total = total

	if b.graph == "" {
		b.graph = "█"
	}

	b.percent = b.getPercent()

	for i := 0; i < b.percent; i += 2 {
		b.rate += b.graph // initial progress position
	}
}

func (b *Bar) getPercent() int {
	return int(float32(b.cur) / float32(b.total) * 100)
}

func (b *Bar) NewOptionWithGraph(start, total int, graph string) {
	b.graph = graph
	b.NewOption(start, total)
}

func (b *Bar) Play(cur int) {
	b.cur = cur
	last := b.percent
	b.percent = b.getPercent()

	if b.percent != last && b.percent%2 == 0 {
		b.rate += b.graph
	}

	fmt.Printf("\r[%-50s]%3d%% %8d/%d", b.rate, b.percent, b.cur, b.total)
}

func (b *Bar) Finish() {
	fmt.Println()
}
