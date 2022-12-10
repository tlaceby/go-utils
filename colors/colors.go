package colors

import "fmt"

type COLOR_STRING string

const (
	KReset  COLOR_STRING = "\x1b[0m"    // Text Reset
	KBlack  COLOR_STRING = "\x1b[0;30m" // Black
	KRed    COLOR_STRING = "\x1b[0;31m" // Red
	KGreen  COLOR_STRING = "\x1b[0;32m" // Green
	KYellow COLOR_STRING = "\x1b[0;33m" // Yellow
	KBlue   COLOR_STRING = "\x1b[0;34m" // Blue
	KPurple COLOR_STRING = "\x1b[0;35m" // Purple
	KCyan   COLOR_STRING = "\x1b[0;36m" // Cyan
	KWhite  COLOR_STRING = "\x1b[0;37m" // Wite

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
	return fmt.Sprintf("%s%s%s", KReset, s, KReset)
}

func Black(s string) string {
	return fmt.Sprintf("%s%s%s", KBlack, s, KReset)
}

func Red(s string) string {
	return fmt.Sprintf("%s%s%s", KRed, s, KReset)
}

func Green(s string) string {
	return fmt.Sprintf("%s%s%s", KGreen, s, KReset)
}

func Yellow(s string) string {
	return fmt.Sprintf("%s%s%s", KYellow, s, KReset)
}

func Blue(s string) string {
	return fmt.Sprintf("%s%s%s", KBlue, s, KReset)
}

func Purple(s string) string {
	return fmt.Sprintf("%s%s%s", KPurple, s, KReset)
}

func FCyan(s string) string {
	return fmt.Sprintf("%s%s%s", KCyan, s, KReset)
}

func White(s string) string {
	return fmt.Sprintf("%s%s%s", KWhite, s, KReset)
}

// ------------------
// BOLD PRINT METHODS
// ------------------

func BoldBlack(s string) string {
	return fmt.Sprintf("%s%s%s", Bold_Black, s, KReset)
}

func BoldRed(s string) string {
	return fmt.Sprintf("%s%s%s", Bold_Red, s, KReset)
}

func BoldGreen(s string) string {
	return fmt.Sprintf("%s%s%s", Bold_Green, s, KReset)
}

func BoldYellow(s string) string {
	return fmt.Sprintf("%s%s%s", Bold_Yellow, s, KReset)
}

func BoldBlue(s string) string {
	return fmt.Sprintf("%s%s%s", Bold_Blue, s, KReset)
}

func BoldPurple(s string) string {
	return fmt.Sprintf("%s%s%s", Bold_Purple, s, KReset)
}

func BoldCyan(s string) string {
	return fmt.Sprintf("%s%s%s", Bold_Cyan, s, KReset)
}

func BoldWhite(s string) string {
	return fmt.Sprintf("%s%s%s", Bold_White, s, KReset)
}
