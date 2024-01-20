package inmem

type VisitorStore struct {
	visitors int
}

func NewVisitorStore() *VisitorStore {
	return &VisitorStore{0}
}

func (vs *VisitorStore) Increment() int {
	vs.visitors++
	return vs.Visitors()
}

func (vs *VisitorStore) Visitors() int {
	return vs.visitors
}
