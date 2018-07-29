package dep1

import (
	dep2 "github.com/v2/ballpit/vgo-dep2"
)

// TestFunc returns some data.
func TestFunc() string {
	return dep2.BreakingChange()
}
