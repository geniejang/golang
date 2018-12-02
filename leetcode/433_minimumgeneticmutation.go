package leetcode

type gene struct {
	Shape     string
	IsEnd     bool
	Visited   bool
	Adjacents []*gene
}

func areAdjacent(a, b string) bool {
	diff := 0
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] != b[i] {
			diff++
		}
	}
	return diff == 1
}

func newGene(shape, end string) gene {
	return gene{Shape: shape, IsEnd: shape == end, Visited: false, Adjacents: []*gene{}}
}

func minMutation(start string, end string, bank []string) int {
	current := newGene(start, end)
	genes := map[string]*gene{}
	for _, shape := range bank {
		gene := newGene(shape, end)
		genes[shape] = &gene
		if areAdjacent(current.Shape, gene.Shape) {
			current.Adjacents = append(current.Adjacents, &gene)
		}
	}
	for i := len(bank) - 1; i >= 1; i-- {
		for j := 0; j < i; j++ {
			a, b := genes[bank[i]], genes[bank[j]]
			if areAdjacent(a.Shape, b.Shape) {
				a.Adjacents = append(a.Adjacents, b)
				b.Adjacents = append(b.Adjacents, a)
			}
		}
	}

	for mutations, q := 0, []*gene{&current}; len(q) > 0; mutations++ {
		nexts := len(q)
		for i := 0; i < nexts; i++ {
			next := q[i]
			if next.IsEnd {
				return mutations
			}
			for _, adj := range next.Adjacents {
				if !adj.Visited {
					adj.Visited = true
					q = append(q, adj)
				}
			}
		}
		q = q[nexts:]
	}
	return -1
}
