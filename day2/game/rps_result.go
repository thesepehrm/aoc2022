package game

type RPSResult uint8

func (r RPSResult) Points() int {
	return rpsResultScores[r]
}

const (
	Win RPSResult = iota
	Tie
	Lose
)
