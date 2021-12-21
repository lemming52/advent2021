package diracdie

import "fmt"

type State struct {
	one, two           int
	oneScore, twoScore int
	count              int
}

func combineState(a, b *State) *State {
	a.count += b.count
	return a
}

func PlayQuantum(one, two int) int {
	state := map[string]*State{
		fmt.Sprintf("%d-%d:%d-%d", one, two, 0, 0): {
			one:      one,
			two:      two,
			oneScore: 0,
			twoScore: 0,
			count:    1}}
	turn := 0
	oneWins, twoWins := 0, 0
	for len(state) > 0 {
		if turn%2 == 0 {
			state = takeQuantumTurn(state, true)
			victors := countVictories(state)
			oneWins += victors
		} else {
			state = takeQuantumTurn(state, false)
			victors := countVictories(state)
			twoWins += victors
		}
		turn++
	}
	if oneWins > twoWins {
		return oneWins
	}
	return twoWins
}

func takeQuantumTurn(states map[string]*State, playerOne bool) map[string]*State {
	newStates := map[string]*State{}
	for _, state := range states {
		individualPositions := positionIncrements(state, playerOne)
		for k, state := range individualPositions {
			v, ok := newStates[k]
			if !ok {
				newStates[k] = state
			} else {
				newStates[k] = combineState(v, state)
			}
		}
	}
	return newStates
}

func countVictories(state map[string]*State) int {
	victors := 0
	for k, v := range state {
		if v.oneScore >= 21 || v.twoScore >= 21 {
			victors += v.count
			delete(state, k)
		}
	}
	return victors
}

func positionIncrements(state *State, playerOne bool) map[string]*State {
	var states []*State
	if playerOne {
		states = []*State{
			{
				one:      (state.one + 3) % 10,
				oneScore: state.oneScore + (state.one+3)%10,
				two:      state.two,
				twoScore: state.twoScore,
				count:    state.count * 1,
			}, {
				one:      (state.one + 4) % 10,
				oneScore: state.oneScore + (state.one+4)%10,
				two:      state.two,
				twoScore: state.twoScore,
				count:    state.count * 3,
			}, {
				one:      (state.one + 5) % 10,
				oneScore: state.oneScore + (state.one+5)%10,
				two:      state.two,
				twoScore: state.twoScore,
				count:    state.count * 6,
			}, {
				one:      (state.one + 6) % 10,
				oneScore: state.oneScore + (state.one+6)%10,
				two:      state.two,
				twoScore: state.twoScore,
				count:    state.count * 7,
			}, {
				one:      (state.one + 7) % 10,
				oneScore: state.oneScore + (state.one+7)%10,
				two:      state.two,
				twoScore: state.twoScore,
				count:    state.count * 6,
			}, {
				one:      (state.one + 8) % 10,
				oneScore: state.oneScore + (state.one+8)%10,
				two:      state.two,
				twoScore: state.twoScore,
				count:    state.count * 3,
			}, {
				one:      (state.one + 9) % 10,
				oneScore: state.oneScore + (state.one+9)%10,
				two:      state.two,
				twoScore: state.twoScore,
				count:    state.count * 1,
			},
		}
	} else {
		states = []*State{
			{
				two:      (state.two + 3) % 10,
				twoScore: state.twoScore + (state.two+3)%10,
				one:      state.one,
				oneScore: state.oneScore,
				count:    state.count * 1,
			}, {
				two:      (state.two + 4) % 10,
				twoScore: state.twoScore + (state.two+4)%10,
				one:      state.one,
				oneScore: state.oneScore,
				count:    state.count * 3,
			}, {
				two:      (state.two + 5) % 10,
				twoScore: state.twoScore + (state.two+5)%10,
				one:      state.one,
				oneScore: state.oneScore,
				count:    state.count * 6,
			}, {
				two:      (state.two + 6) % 10,
				twoScore: state.twoScore + (state.two+6)%10,
				one:      state.one,
				oneScore: state.oneScore,
				count:    state.count * 7,
			}, {
				two:      (state.two + 7) % 10,
				twoScore: state.twoScore + (state.two+7)%10,
				one:      state.one,
				oneScore: state.oneScore,
				count:    state.count * 6,
			}, {
				two:      (state.two + 8) % 10,
				twoScore: state.twoScore + (state.two+8)%10,
				one:      state.one,
				oneScore: state.oneScore,
				count:    state.count * 3,
			}, {
				two:      (state.two + 9) % 10,
				twoScore: state.twoScore + (state.two+9)%10,
				one:      state.one,
				oneScore: state.oneScore,
				count:    state.count * 1,
			},
		}
	}
	newStates := map[string]*State{}
	for _, s := range states {
		if s.one == 0 {
			s.one = 10
			s.oneScore += 10
		}
		if s.two == 0 {
			s.two = 10
			s.twoScore += 10
		}
		newStates[fmt.Sprintf("%d-%d:%d-%d", s.one, s.two, s.oneScore, s.twoScore)] = s
	}
	return newStates
}
