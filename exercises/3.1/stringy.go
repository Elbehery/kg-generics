package stringy

import (
	"fmt"
	"io"
)

// Your StringifyTo function goes here!
func StringifyTo[S fmt.Stringer](w io.Writer, v S) {
	fmt.Fprintln(w, v.String())
}
