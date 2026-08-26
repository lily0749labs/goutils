package slice

import (
	"reflect"
	"testing"
)

func TestDeleteFirst(t *testing.T) {
	t.Parallel()

	values := []int{1, 2, 3, 2}
	if got, want := DeleteFirst(values, 2), []int{1, 3, 2}; !reflect.DeepEqual(got, want) {
		t.Fatalf("DeleteFirst() = %#v, want %#v", got, want)
	}
	if got := DeleteFirst(values, 9); !reflect.DeepEqual(got, values) {
		t.Fatalf("DeleteFirst(missing) = %#v, want %#v", got, values)
	}
}

func TestUniqueSorted(t *testing.T) {
	t.Parallel()

	if got := UniqueSorted([]int(nil)); got != nil {
		t.Fatalf("UniqueSorted(nil) = %#v", got)
	}
	if got, want := UniqueSorted([]int{3, 1, 2, 3, 1}), []int{1, 2, 3}; !reflect.DeepEqual(got, want) {
		t.Fatalf("UniqueSorted() = %#v, want %#v", got, want)
	}
}
