package window_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager/internal/window"
)

func TestSlidingWindow_Order(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	out, err := window.SlidingWindow(ctx, 5, func(windowLocation int) (interface{}, error) {
		if windowLocation < 50 {
			return windowLocation, nil
		}
		return nil, nil
	})

	// Test the print vs the output comment
	for o := range out {
		fmt.Print(o, " ")
	}
	fmt.Println("error:", <-err)

	// Output:
	// 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 error: <nil>
}

func TestSlidingWindow_Errors(t *testing.T) {
	const msg = "error at location 10"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	out, err := window.SlidingWindow(ctx, 5, func(windowLocation int) (interface{}, error) {
		if windowLocation == 10 {
			return nil, fmt.Errorf(msg)
		}
		return windowLocation, nil
	})

	num := 0
	for range out {
		num++
	}

	// Error stops iteration and returns an error
	if num > 10 {
		t.Errorf("kept iterating beyond 10 times: actual %d", num)
	}
	if e := <-err; e == nil {
		t.Fatal("Error was expected")
	}
}

func TestSlidingWindow_WithoutBuffer(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	out, err := window.SlidingWindow(ctx, 1, func(windowLocation int) (interface{}, error) {
		if windowLocation < 50 {
			return windowLocation, nil
		}
		return nil, nil
	})

	// Test the print vs the output comment
	for o := range out {
		fmt.Print(o, " ")
	}
	fmt.Println("error:", <-err)

	// Output:
	// 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 error: <nil>
}

func BenchmarkSlidingWindow_Order(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		out, err := window.SlidingWindow(ctx, 15, func(windowLocation int) (interface{}, error) {
			if windowLocation < 50 {
				return windowLocation, nil
			}
			return nil, nil
		})

		// Test the print vs the output comment
		n := 0
		for o := range out {
			if on, ok := o.(int); !ok || on != n {
				b.Errorf("incorrect element expected %v got %v", n, o)
			}
			n++
		}
		if e := <-err; e != nil {
			b.Error(e)
		}
	}
}
