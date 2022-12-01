package colors

import "fmt"

type COLOR_STRING string

const (
	Reset  COLOR_STRING = "\x1b[0m"    // Text Reset
	Black  COLOR_STRING = "\x1b[0;30m" // Black
	Red    COLOR_STRING = "\x1b[0;31m" // Red
	Green  COLOR_STRING = "\x1b[0;32m" // Green
	Yellow COLOR_STRING = "\x1b[0;33m" // Yellow
	Blue   COLOR_STRING = "\x1b[0;34m" // Blue
	Purple COLOR_STRING = "\x1b[0;35m" // Purple
	Cyan   COLOR_STRING = "\x1b[0;36m" // Cyan
	White  COLOR_STRING = "\x1b[0;37m" // Wite

	Bold_Black  COLOR_STRING = "\x1b[1;30m" // Bold Black
	Bold_Red    COLOR_STRING = "\x1b[1;31m" // Bold Red
	Bold_Green  COLOR_STRING = "\x1b[1;32m" // Bold Green
	Bold_Yellow COLOR_STRING = "\x1b[1;33m" // Bold Yellow
	Bold_Blue   COLOR_STRING = "\x1b[1;34m" // Bold Blue
	Bold_Purple COLOR_STRING = "\x1b[1;35m" // Bold Purple
	Bold_Cyan   COLOR_STRING = "\x1b[1;36m" // Bold Cyan
	Bold_White  COLOR_STRING = "\x1b[1;37m" // Bold White

	BrightRed    COLOR_STRING = "\x1b[0;91m" // Bright Red
	BrightGreen  COLOR_STRING = "\x1b[0;92m" // Bright Green
	BrightYellow COLOR_STRING = "\x1b[0;93m" // Bright Yellow
	BrightBlue   COLOR_STRING = "\x1b[0;94m" // Bright Blue
	BrightPurple COLOR_STRING = "\x1b[0;95m" // Bright Purple
	BrightCyan   COLOR_STRING = "\x1b[0;96m" // Bright Cyan
	BrightWhite  COLOR_STRING = "\x1b[0;97m" // Bright White

	UnderlineBlack  COLOR_STRING = "\x1b[4;30m" // Underline Black
	UnderlineRed    COLOR_STRING = "\x1b[4;31m" // Underline Red
	UnderlineGreen  COLOR_STRING = "\x1b[4;32m" // Underline Green
	UnderlineYellow COLOR_STRING = "\x1b[4;33m" // Underline Yellow
	UnderlineBlue   COLOR_STRING = "\x1b[4;34m" // Underline Blue
	UnderlinePurple COLOR_STRING = "\x1b[4;35m" // Underline Purple
	UnderlineCyan   COLOR_STRING = "\x1b[4;36m" // Underline Cyan
	UnderlineWhite  COLOR_STRING = "\x1b[4;37m" // Underline White
)

func Default(s string) string {
	return fmt.Sprintf("%s%s%s", Reset, s, Reset)
}

func FBlack(s string) string {
	return fmt.Sprintf("%s%s%s", Black, s, Reset)
}

func FRed(s string) string {
	return fmt.Sprintf("%s%s%s", Red, s, Reset)
}

func FGreen(s string) string {
	return fmt.Sprintf("%s%s%s", Green, s, Reset)
}

func FYellow(s string) string {
	return fmt.Sprintf("%s%s%s", Yellow, s, Reset)
}

func FBlue(s string) string {
	return fmt.Sprintf("%s%s%s", Blue, s, Reset)
}

func FPurple(s string) string {
	return fmt.Sprintf("%s%s%s", Purple, s, Reset)
}

func FCyan(s string) string {
	return fmt.Sprintf("%s%s%s", Cyan, s, Reset)
}

func FWhite(s string) string {
	return fmt.Sprintf("%s%s%s", White, s, Reset)
}

// ------------------
// BOLD PRINT METHODS
// ------------------

func BoldBlack(s string) string {
	return fmt.Sprintf("%s%s%s", Bold_Black, s, Reset)
}

func BoldRed(s string) string {
	return fmt.Sprintf("%s%s%s", Bold_Red, s, Reset)
}

func BoldGreen(s string) string {
	return fmt.Sprintf("%s%s%s", Bold_Green, s, Reset)
}

func BoldYellow(s string) string {
	return fmt.Sprintf("%s%s%s", Bold_Yellow, s, Reset)
}

func BoldBlue(s string) string {
	return fmt.Sprintf("%s%s%s", Bold_Blue, s, Reset)
}

func BoldPurple(s string) string {
	return fmt.Sprintf("%s%s%s", Bold_Purple, s, Reset)
}

func BoldCyan(s string) string {
	return fmt.Sprintf("%s%s%s", Bold_Cyan, s, Reset)
}

func BoldWhite(s string) string {
	return fmt.Sprintf("%s%s%s", Bold_White, s, Reset)
}
