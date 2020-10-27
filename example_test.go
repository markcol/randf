package randf_test

import (
	"fmt"

	"github.com/markcol/randf"
)

func Example() {
	r := randf.New()
	r.Seed(42)
	f := r.Float32()
	fmt.Println(f)
	// Output: 0.94806355
}
