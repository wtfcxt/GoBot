package logger

import (
	"fmt"
	"github.com/fatih/color"
)

var (
	TypeInfo  = 0
	TypeWarn  = 1
	TypeError = 2
)

func LogNormalNoNewline(logtype int, str string) {
	switch logtype {
	case TypeInfo:
		fmt.Fprint(color.Output, color.HiBlueString("[INFO] "+color.HiWhiteString(str)))
	case TypeWarn:
		fmt.Fprint(color.Output, color.YellowString("[WARN] "+color.HiWhiteString(str)))
	case TypeError:
		fmt.Fprint(color.Output, color.RedString("[ERROR] "+color.HiWhiteString(str)))
	default:
		fmt.Fprint(color.Output, color.RedString("[ERROR] "+color.HiWhiteString("Invalid LoggerType. Original message: "+str)))
	}
}

func LogModuleNoNewline(logtype int, module string, str string) {
	switch logtype {
	case TypeInfo:
		fmt.Fprint(color.Output, color.HiBlueString("[INFO] <"+module+"> "+color.HiWhiteString(str)))
	case TypeWarn:
		fmt.Fprint(color.Output, color.YellowString("[WARN] <"+module+"> "+color.HiWhiteString(str)))
	case TypeError:
		fmt.Fprint(color.Output, color.RedString("[ERROR] <"+module+"> "+color.HiWhiteString(str)))
	default:
		fmt.Fprint(color.Output, color.RedString("[ERROR] <GoLog> "+color.HiWhiteString("Invalid LoggerType. Original message: "+str)))
	}
}

func AppendDone() {
	fmt.Fprintln(color.Output, color.GreenString(" Done"))
}

func AppendFail() {
	fmt.Fprintln(color.Output, color.RedString(" Failure"))
}

func LogModule(logtype int, module string, str string) {
	switch logtype {
	case TypeInfo:
		fmt.Fprintln(color.Output, color.HiBlueString("[INFO] <"+module+"> "+color.HiWhiteString(str)))
	case TypeWarn:
		fmt.Fprintln(color.Output, color.YellowString("[WARN] <"+module+"> "+color.HiWhiteString(str)))
	case TypeError:
		fmt.Fprintln(color.Output, color.RedString("[ERROR] <"+module+"> "+color.HiWhiteString(str)))
	default:
		fmt.Fprintln(color.Output, color.RedString("[ERROR] <GoLog> "+color.HiWhiteString("Invalid LoggerType. Original message: "+str)))
	}
}
