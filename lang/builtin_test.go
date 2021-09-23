package lang

import (
	"testing"
	"time"
)

func TestCopy(t *testing.T) {
	src := []int{1, 2, 3}
	dst := make([]int, 3)
	copied := copy(dst, src)
	t.Logf("copied: %d", copied)

	for _, v := range src {
		t.Log(v)
	}
}

func TestTimeFormat(t *testing.T) {
	// this function returns the present time
	current_time := time.Now()

	// individual elements of time can
	// also be called to print accordingly
	t.Logf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		current_time.Year(), current_time.Month(), current_time.Day(),
		current_time.Hour(), current_time.Minute(), current_time.Second())

	// formatting time using
	// custom formats
	t.Log(current_time.Format("2006-01-02 15:04:05"))
	t.Log(current_time.Format("2006-Jan-02"))
	t.Log(current_time.Format("2006-01-02 3:4:5 pm"))

	t.Log(current_time.Format(time.RFC3339))
}

