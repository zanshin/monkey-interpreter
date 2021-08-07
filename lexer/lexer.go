package lexer

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

// Return a pointer to a Lexer, built from passed input
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// Get next character, and advance position in input
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}
