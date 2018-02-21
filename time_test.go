
package palindrome

import (
	"fmt"
	"time"
	"testing"
)

type TestInfo struct {
	min int
	max int
	count int
	maxtime time.Duration
}

var timeTests []TestInfo = []TestInfo{
	TestInfo{1, 1000, 108, time.Second},
	TestInfo{1, 10000, 198, time.Second},
	TestInfo{1, 100000, 1098, time.Second},
	TestInfo{1, 1000000, 1998, time.Second},
	TestInfo{1, 10000000, 10998, 2 * time.Second},
}

var timeGenTests []TestInfo = []TestInfo{
	TestInfo{1, 100000000, 19998, 5 * time.Second},
	TestInfo{1, 1000000000, 109998, 5 * time.Second},
	TestInfo{1, 10000000000, 199998, 5 * time.Second},
	TestInfo{1, 100000000000, 1099998, 5 * time.Second},
	TestInfo{1, 1000000000000, 1999998, 5 * time.Second},
	TestInfo{1, 10000000000000, 10999998, 5 * time.Second},
	TestInfo{1, 100000000000000, 19999998, 5 * time.Second},
}

func Test_TimeLinearInt(t *testing.T) {
	for _, ti := range timeTests {
		c, x := TimeLinearInt(ti.min, ti.max)

		if c != ti.count || x >= ti.maxtime {
			t.Error(fmt.Sprintf("TimeLinearInt error, %d(%d) %v(%v)", c, ti.count, x, ti.maxtime))
		}
	}
}

func Test_TimeLinearStr(t *testing.T) {
	for _, ti := range timeTests {
		c, x := TimeLinearStr(ti.min, ti.max)

		if c != ti.count || x >= ti.maxtime {
			t.Error(fmt.Sprintf("TimeLinearStr error, %d(%d) %v(%v)", c, ti.count, x, ti.maxtime))
		}
	}
}

func Test_TimeNextGenerator(t *testing.T) {
	for _, ti := range timeTests {
		c, x := TimeNextGenerator(ti.min, ti.max)

		if c != ti.count || x >= ti.maxtime {
			t.Error(fmt.Sprintf("TimeNextGenerator error, %d(%d) %v(%v)", c, ti.count, x, ti.maxtime))
		}
	}
	for _, ti := range timeGenTests {
		c, x := TimeNextGenerator(ti.min, ti.max)

		if c != ti.count || x >= ti.maxtime {
			t.Error(fmt.Sprintf("TimeNextGenerator error, %d(%d) %v(%v)", c, ti.count, x, ti.maxtime))
		}
	}
}

func Test_TimePrevGenerator(t *testing.T) {
	for _, ti := range timeTests {
		c, x := TimePrevGenerator(ti.min, ti.max)

		if c != ti.count || x >= ti.maxtime {
			t.Error(fmt.Sprintf("TimePrevGenerator error, %d(%d) %v(%v)", c, ti.count, x, ti.maxtime))
		}
	}
	for _, ti := range timeGenTests {
		c, x := TimePrevGenerator(ti.min, ti.max)

		if c != ti.count || x >= ti.maxtime {
			t.Error(fmt.Sprintf("TimePrevGenerator error, %d(%d) %v(%v)", c, ti.count, x, ti.maxtime))
		}
	}
}
