package group

type Group[E any] []E

// Your Len function goes here!
func Len[E any](v []E) int {
	return len(v)
}
