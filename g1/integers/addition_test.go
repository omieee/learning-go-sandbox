// calc_test.go
package addition

import (
	"fmt"
	"testing"
) //

func TestAdd(t *testing.T) {
	got := Add(2, 3)
	want := 5

	if got != want {
		t.Errorf("Add(2, 3) = %d; want %d", got, want)
	}
}

func ExampleAdd() {
	sum := Add(1, 2)
	fmt.Println(sum)
	//Output: 3
}
