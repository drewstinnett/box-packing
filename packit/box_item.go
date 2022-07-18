package packit

type BoxItem interface {
	GetSize() int64
}

// Sorted BoxItems by size, smallest first
type BoxItemsBySize []BoxItem

func (a BoxItemsBySize) Len() int           { return len(a) }
func (a BoxItemsBySize) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BoxItemsBySize) Less(i, j int) bool { return a[i].GetSize() < a[j].GetSize() }
