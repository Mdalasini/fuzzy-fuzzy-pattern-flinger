package matcher

type matcherToStatePair struct {
	m matcher
	s *state
}

type state struct {
	name string
	transition []matcherToStatePair
}

func newState(name string) *state {
	return &state{name: name}
}

func (s *state) addTransition(toState *state, matcher matcher) {
	s.transition = append(s.transition, matcherToStatePair{matcher, toState})
}

func (s *state) unshiftTransitions(toState *state, matcher matcher) {
	s.transition = append([]matcherToStatePair{{matcher, toState}}, s.transition...)
}