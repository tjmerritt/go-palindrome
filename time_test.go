
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

func Benchmark_TimeLinearInt(t *testing.B) {
        st := time.Now()
	for _, ti := range timeTests {
		c, x := TimeLinearInt(ti.min, ti.max)

		if c != ti.count || x >= ti.maxtime {
			t.Error(fmt.Sprintf("TimeLinearInt error, %d(%d) %v(%v)", c, ti.count, x, ti.maxtime))
		}
	}
        fmt.Printf("TimeLinearInt: %v\n", time.Now().Sub(st))
}

func Benchmark_TimeLinearStr(t *testing.B) {
	for _, ti := range timeTests {
		c, x := TimeLinearStr(ti.min, ti.max)

		if c != ti.count || x >= ti.maxtime {
			t.Error(fmt.Sprintf("TimeLinearStr error, %d(%d) %v(%v)", c, ti.count, x, ti.maxtime))
		}
	}
}

func Benchmark_TimeNextGenerator(t *testing.B) {
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

func Benchmark_TimePrevGenerator(t *testing.B) {
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

func Test_Benchmarks(t *testing.T) {
        c, _ := TimeLinearInt(1, 10000000)
        if c != 10998 {
                t.Errorf("Test_Benchmarks, TimeLinearInt failed, got %d, expected 10998", c)
        }
        c, _ = TimeLinearStr(1, 10000000)
        if c != 10998 {
                t.Errorf("Test_Benchmarks, TimeLinearStr failed, got %d, expected 10998", c)
        }
        c, _ = TimeNextGenerator(1, 10000000)
        if c != 10998 {
                t.Errorf("Test_Benchmarks, TimeNextGenerator failed, got %d, expected 10998", c)
        }
        c, _ = TimeNextGenerator(10, 10000000)
        if c != 10989 {
                t.Errorf("Test_Benchmarks, TimeNextGenerator failed, got %d, expected 10989", c)
        }
        c, _ = TimePrevGenerator(1, 10000000)
        if c != 10998 {
                t.Errorf("Test_Benchmarks, TimePrevGenerator failed, got %d, expected 10998", c)
        }
}

