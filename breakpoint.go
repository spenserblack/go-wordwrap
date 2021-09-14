package wordwrap

// Breakpoint represents where a string should be broken to wrap.
type breakpoint interface {
	// Start is where the break should start.
	Start() int
	// End is where the break should end.
	End() int
	// Len is the length of the break point. It should be 0 if it doesn't add or
	// remove any text, greater than 0 if it adds any characters, and less than 0
	// if it removes any characters.
	Len() int
}

// Spacepoint is a breakpoint triggered by whitespace.
type spacepoint int

func (p spacepoint) Start() int {
	return int(p)
}
func (p spacepoint) End() int {
	return int(p) + 1
}
func (p spacepoint) Len() int {
	return -1
}

// Hyphenpoint is a breakpoint triggered by a hyphen.
type hyphenpoint int

func (p hyphenpoint) Start() int {
	return int(p) + 1
}
func (p hyphenpoint) End() int {
	return int(p) + 1
}
func (p hyphenpoint) Len() int {
	return 0
}

// Defaultpoint occurs when there is no available trigger for breakage.
type defaultpoint int

func (p defaultpoint) Start() int {
	return int(p) + 1
}
func (p defaultpoint) End() int {
	return int(p) + 1
}
func (p defaultpoint) Len() int {
	return 0
}
