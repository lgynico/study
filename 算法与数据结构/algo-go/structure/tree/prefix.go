package tree

type (
	prefixNode struct {
		pass int
		end  int
		path map[byte]*prefixNode
	}

	Prefix struct {
		root prefixNode
	}
)

func NewPrefix() Prefix {
	return Prefix{
		root: prefixNode{
			path: make(map[byte]*prefixNode),
		},
	}
}

func (p *Prefix) Insert(value string) {
	var node = &p.root
	for i := 0; i < len(value); i++ {
		next, ok := node.path[value[i]-'0']
		if !ok {
			next = &prefixNode{path: make(map[byte]*prefixNode)}
			node.path[value[i]-'0'] = next
		}
		node = next
		node.pass++
		if i == len(value)-1 {
			node.end++
		}
	}
}

func (p *Prefix) Remove(value string) {
	var node = &p.root
	for i := 0; i < len(value); i++ {
		node = node.path[value[i]-'0']
		if node == nil {
			return
		}
		node.pass--
		if i == len(value)-1 {
			node.end--
		}

		if node.end == 0 {
			delete(node.path, value[i]-'0')
			return
		}
	}
}

func (p *Prefix) Contains(value string) bool {
	var node = &p.root
	for i := 0; i < len(value); i++ {
		node = node.path[value[i]-'0']
		if node == nil || node.pass == 0 {
			return false
		}
		if i == len(value)-1 {
			return node.end > 0
		}
	}
	return true
}

func (p *Prefix) ContainsPrefix(prefix string) bool {
	var node = &p.root
	for i := 0; i < len(prefix); i++ {
		node = node.path[prefix[i]-'0']
		if node == nil || node.pass == 0 {
			return false
		}
	}
	return true
}
