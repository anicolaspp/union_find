package unionfind

// UnionFind is a data structure to keep track of disjoint groups.
type UnionFind[T comparable] struct {
	rank   map[T]int
	parent map[T]T
}

func New[T comparable]() *UnionFind[T] {
	return &UnionFind[T]{
		rank:   make(map[T]int),
		parent: make(map[T]T),
	}
}

// Sets returns the list of disjoint sets.
func (ds *UnionFind[T]) Sets() [][]T {
	sets := map[T][]T{}

	for k := range ds.parent {
		kp := ds.find(k)
		sets[kp] = append(sets[kp], k)
	}

	var results [][]T
	for _, v := range sets {
		results = append(results, v)
	}

	return results
}

// Add adds a to the structure without any association.
func (ds *UnionFind[T]) Add(a T) {
	ds.Union(a, a)
}

// Union joins the groups of the given elements.
func (ds *UnionFind[T]) Union(a, b T) {
	ap := ds.find(a)
	bp := ds.find(b)

	if ap == bp {
		return
	}

	if ds.rank[ap] < ds.rank[bp] {
		ds.parent[ap] = bp
	} else if ds.rank[ap] > ds.rank[bp] {
		ds.parent[bp] = ap
	} else {
		ds.parent[ap] = bp
		ds.rank[ap] = ds.rank[bp] + 1
	}

	ds.parent[ap] = bp
}

func (ds *UnionFind[T]) find(item T) T {
	r, ok := ds.parent[item]
	if !ok {
		ds.parent[item] = item
		return item
	}

	if r != item {
		ds.parent[item] = ds.find(ds.parent[item])
	}

	return ds.parent[item]
}
