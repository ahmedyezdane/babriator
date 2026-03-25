package screen

const (
	// Move cursor to position (row, col) - both 1-based
	ansiMoveCursor = "\033[%d;%dH"

	ansiEnterAlternateScreen = "\033[?1049h"
	ansiExitAlternateScreen  = "\033[?1049l"

	ansiEraseLine  = "\033[2K"
	ansiShowCursor = "\033[?25h"
	ansiHideCursor = "\033[?25l"
	ansiCreateLine = "\033[L"
	ansiClearLine  = "\033[2K"
	ansiDeleteLine = "\033[M"
)
