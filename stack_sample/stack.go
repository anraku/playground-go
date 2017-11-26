package main

type Stack struct {
	Value []string
	limit int
}

func (s *Stack) Push(ss string) {
	s.Value = append(s.Value, ss)
	if len(s.Value) > s.limit {
		s.Value = s.Value[1:]
	}
}

func (s *Stack) Pop() string {
	if len(s.Value) == 0 {
		return ""
	}
	v := s.Value[len(s.Value)-1]
	s.Value = s.Value[:len(s.Value)-1]
	return v
}

func (s *Stack) ChangeLimit(i int) {
	s.limit = i
}
