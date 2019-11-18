package bench2

import (
	"strings"
	"testing"
)

const s = `lkhjsdflgkjnlk sdm;fjallsdc nlksjdhfglks jndflksjndflkjnsgfd
lkhjsdflgkjnlksdm;fjallsdcnlksjdhfglks jndflksjndflkjnsgfd
lkhjsdflgkjnlksdm ;fjallsdcnlksjdhfglksjndflksjndflkjnsgfd
lkhjsdflgkjnlksdm;fjallsdc nlksjdhfglksjndflksjndflkjnsgfd
lkhjsdflgkjnlk sdm;fjallsdcnlksjdhfglksjndflksjndflkjnsgfd
lkhjsdflgkjn lksd m;fjallsdcnlksjdhfglk sjndfl ksjn dflkjn sgfd`

var xs []string

func BenchmarkCat(b *testing.B) {
	xs = strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Cat(xs)
	}

}
func BenchmarkJoin(b *testing.B) {
	xs = strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(xs)
	}
}
