package game

const (
	Rock RPSMove = iota
	Paper
	Scissors
)

type RPSMove uint8

func (m RPSMove) String() string {
	switch m {
	case Rock:
		return "rock"
	case Paper:
		return "paper"
	case Scissors:
		return "scissors"
	}

	return ""
}

func (m RPSMove) Points() int {
	return rpsMoveScores[m]
}

func (m RPSMove) Match(op RPSMove) RPSResult {
	if m == op {
		return Tie
	}

	switch m {
	case Rock:
		if op == Paper {
			return Lose
		}
	case Paper:
		if op == Scissors {
			return Lose
		}
	case Scissors:
		if op == Rock {
			return Lose
		}
	}

	return Win
}

func FindMoveFor(op RPSMove, outcome RPSResult) RPSMove {
	if outcome == Tie {
		return op
	}

	moveForOutcome := map[RPSMove](map[RPSResult]RPSMove){
		Rock: {
			Win:  Paper,
			Lose: Scissors,
		},

		Paper: {
			Win:  Scissors,
			Lose: Rock,
		},
		Scissors: {
			Win:  Rock,
			Lose: Paper,
		},
	}
	return moveForOutcome[op][outcome]
}
