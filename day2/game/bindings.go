package game

var leftColumnSideBinding = map[string]RPSMove{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
}

var rightColumnSideBinding = map[string]RPSMove{
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

var rightColumnSideBindingPart2 = map[string]RPSResult{
	"X": Lose,
	"Y": Tie,
	"Z": Win,
}

var rpsMoveScores = map[RPSMove]int{
	Rock:     1,
	Paper:    2,
	Scissors: 3,
}

var rpsResultScores = map[RPSResult]int{
	Win:  6,
	Tie:  3,
	Lose: 0,
}
