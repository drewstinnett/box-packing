package packit

type Box struct {
	Capacity int64
	Used     int64
	Items    []BoxItem
}

func (b *Box) Remaining() int64 {
	return b.Capacity - b.Used
}

func (b *Box) AddItem(i BoxItem) {
	b.Items = append(b.Items, i)
	b.Used += i.GetSize()
}
