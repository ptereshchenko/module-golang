package stack

type Stack struct {
  nodes []int
  count int
}

func New() *Stack {
  return &Stack{}
}

func (s *Stack) Push(n int) {
  s.nodes = append(s.nodes[:s.count], n)
  s.count++
}

func (s *Stack) Pop() int {
  s.count--
  return s.nodes[s.count]
}
