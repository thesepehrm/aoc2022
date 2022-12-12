package game

type RightColumnCode struct {
	encryptedValue string
}

type LeftColumnCode struct {
	encryptedValue string
}

func NewRightColumnCode(encryptedValue string) RightColumnCode {
	return RightColumnCode{
		encryptedValue: encryptedValue,
	}
}

func NewLeftColumnCode(encryptedValue string) LeftColumnCode {
	return LeftColumnCode{
		encryptedValue: encryptedValue,
	}
}

func (c LeftColumnCode) Value() RPSMove {
	return leftColumnSideBinding[c.encryptedValue]
}

func (c RightColumnCode) Value() RPSMove {
	return rightColumnSideBinding[c.encryptedValue]
}

func (c RightColumnCode) ValuePart2() RPSResult {
	return rightColumnSideBindingPart2[c.encryptedValue]
}
