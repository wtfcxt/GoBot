package logger

import (
	"fmt"
	"github.com/fatih/color"
)

var (
	TypeInfo = 0
	TypeWarn = 1
	TypeError = 2
	TypeDebug = 3
)

func LogCrash(err error) {
	fmt.Println(color.RedString("[CRASH] A critical error occurred and the bot crashed. Please report this incident to cxt#1234 on discord."))
	panic(err)
}

func LogLogo() {
	fmt.Fprintln(color.Output, color.YellowString("  _____     ___       __    _  __\n / ___/__  / _ )___  / /_  | |/_/\n/ (_ / _ \\/ _  / _ \\/ __/ _>  <  \n\\___/\\___/____/\\___/\\__/ /_/|_|  \n                                 "))
}

func LogModule(logtype int, module string, str string) {
	switch logtype {
	case TypeInfo:
		fmt.Fprintln(color.Output, color.HiBlueString("[INFO] <" + module + "> " + color.HiWhiteString(str)))
	case TypeWarn:
		fmt.Fprintln(color.Output, color.YellowString("[WARN] <" + module + "> " + color.HiWhiteString(str)))
	case TypeError:
		fmt.Fprintln(color.Output, color.RedString("[ERROR] <" + module + "> " + color.HiWhiteString(str)))
	case TypeDebug:
		fmt.Fprintln(color.Output, color.HiMagentaString("[DEBUG] <" + module + "> " + color.HiWhiteString(str)))
	default:
		fmt.Fprintln(color.Output, color.RedString("[ERROR] <GoLog> " + color.HiWhiteString("Invalid LoggerType. Original message: " + str)))
	}
}