package game

func RPS(leftColumn string, rightColumn string) int {
	op := NewLeftColumnCode(leftColumn).Value()
	self := NewRightColumnCode(rightColumn).Value()

	return self.Points() + self.Match(op).Points()
}

func RPSPart2(leftColumn string, rightColumn string) int {
	op := NewLeftColumnCode(leftColumn).Value()
	desiredOutcome := NewRightColumnCode(rightColumn).ValuePart2()

	selfMove := FindMoveFor(op, desiredOutcome)

	return selfMove.Points() + desiredOutcome.Points()
}
