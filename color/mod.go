package color

type Color string

const (
	RESET  Color = "\033[0m"
	RED    Color = "\033[31m"
	GREEN  Color = "\033[32m"
	YELLOW Color = "\033[33m"
	BLUE   Color = "\033[34m"
	PURPLE Color = "\033[35m"
	CYAN   Color = "\033[36m"
	WHITE  Color = "\033[37m"
	GRAY   Color = "\033[38;5;7m"
)
