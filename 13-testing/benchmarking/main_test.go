package benchmarking

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("josh")
	if s != "welcome sir!, josh" {
		t.Errorf("expected: 'welcome sir!, josh' | got: %v", s)
	}
}

func ExampleGreet() {
	fmt.Println(Greet("josh"))
	// Output:
	// welcome sir!, josh
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("josh")
	}
}

// $ go test -cover :: analysis of how many statements are covered by the test suite
// $ go test -coverprofile=c.out :: identifies coverage holes
// $ go tool cover -html=c.out :: opens up a visual for full coverage analysis
