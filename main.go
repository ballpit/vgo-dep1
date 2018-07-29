package dep1

import (
	dep2 "github.com/ballpit/vgo-dep2/v2"
)

// TestFunc returns some data.
func TestFunc() string {
	return dep2.BreakingChange()
}
