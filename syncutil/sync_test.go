package syncutil

import (
	"sync"
	"testing"
)

func TestMapLen(t *testing.T) {
	t.Parallel()

	if got := Sync.MapLen(nil); got != 0 {
		t.Fatalf("Sync.MapLen(nil) = %d", got)
	}

	var values sync.Map
	values.Store("a", 1)
	values.Store("b", 2)
	values.Store("c", 3)
	values.Delete("b")
	if got := Sync.MapLen(&values); got != 2 {
		t.Fatalf("Sync.MapLen() = %d, want 2", got)
	}
}
