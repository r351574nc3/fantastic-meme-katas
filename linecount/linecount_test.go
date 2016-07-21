package linecount

import (
	"fmt"
	"log"
	"testing"
)

func TestCountLinesIn(t *testing.T) {
	for _, c := range []struct {
		in1, want string
	}{
		{"test.html", "100"},
		{"", ""},
	} {
		got, err := CountLinesIn(c.in1)
		if err != nil {
			log.Fatal(err)
			fmt.Println("%d", len(got))
		}
	}
}
