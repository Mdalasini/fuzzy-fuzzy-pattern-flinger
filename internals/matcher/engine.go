package matcher

import "fmt"

type engineNFA struct {
	states       map[string]*state
	initialState string
	endingStates []string
}

func NewEngineNFA() *engineNFA {
	return &engineNFA{states: make(map[string]*state)}
}

func (nfa *engineNFA) AddState(name string) {
	nfa.states[name] = newState(name)
}

func (nfa *engineNFA) DeclareStates(names ...string) {
	for _, n := range names {
		nfa.AddState(n)
	}
}

func (nfa *engineNFA) SetInitialState(name string) {
	nfa.initialState = name
}

func (nfa *engineNFA) SetEndingStates(names []string) {
	nfa.endingStates = names
}

func (nfa *engineNFA) AddTransition(fromState, toState string, matcher matcher) {
	nfa.states[fromState].addTransition(nfa.states[toState], matcher)
}

func (nfa *engineNFA) UnshiftTransitions(fromState, toState string, matcher matcher) {
	nfa.states[fromState].unshiftTransitions(nfa.states[toState], matcher)
}

type stackItem struct {
	currentState    *state
	epsilionVisited []string
}

type stack struct {
	items []stackItem
}

func (s *stack) pop() stackItem {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func (nfa *engineNFA) Compute(input string) bool {
	stack := &stack{}
	stack.items = append(stack.items, stackItem{nfa.states[nfa.initialState], make([]string, 0)})

	scanner := newScanner(input)

	for {
		if len(stack.items) == 0 {
			break
		}
		item := stack.pop()
		if contains(nfa.endingStates, item.currentState.name) {
			return true
		}

		for i := len(item.currentState.transition) - 1; i >= 0; i-- {
			matcherToStatePair := item.currentState.transition[i]
			if matcherToStatePair.m.matches(scanner) {
				if matcherToStatePair.m.isEpsilon() {
					// * Don't follow the transition. We already have been in that state
					if contains(item.epsilionVisited, matcherToStatePair.s.name) {
						continue
					}
					// * Record that you've made this transition
					item.epsilionVisited = append(item.epsilionVisited, item.currentState.name)
				} else {
					// * we are transversing a non-epsilon transition, so reset the visited counter
					item.epsilionVisited = make([]string, 0)
				}
				stack.items = append(stack.items, stackItem{matcherToStatePair.s, item.epsilionVisited})
			} else {
				// * If doesn't match and scanner isn't empty restart NFA
				if _, err := scanner.r.ReadByte(); err == nil {
					fmt.Println("didn't match")
					stack.items = append(stack.items, stackItem{nfa.states[nfa.initialState], make([]string, 0)})
				}								
			}
		}
	}
	return false
}