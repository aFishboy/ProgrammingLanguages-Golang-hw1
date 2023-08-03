package nfa

// A state in the NFA is labeled by a single integer.
type state uint

// TransitionFunction tells us, given a current state and some symbol, which
// other states the NFA can move to.
//
// Deterministic automata have only one possible destination state,
// but we're working with non-deterministic automata.
type TransitionFunction func(st state, act rune) []state

func Reachable(
	// `transitions` tells us what our NFA looks like
	transitions TransitionFunction,
	// `start` and `final` tell us where to start, and where we want to end up
	currentState, final state,
	// `input` is a (possible empty) list of symbols to apply.
	input []rune,
) bool {
	/* check to see if there are no symbols given if so then check
		if currentState is desired final state  */
	if len(input) == 0 {
		return currentState == final
	}
	/* try each possible transition from the current state */
	for _, index := range transitions(currentState, input[0]) {
		/* recursively call Reachable with the next and remaining input symbols */
		if Reachable(transitions, index, final, input[1:]) == true {
			return true
		}
	}
	return false
}
