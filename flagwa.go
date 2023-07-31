package main

import (
	"fmt"
	"github.com/Liyihwa/logwa"
	"os"
	"strconv"
	"strings"
)

var flagArgs map[string]bool     //存储从os.Args获取的标志参数
var PlaceArgs []string           //存储从os.Args获取的位置参数
var optionArgs map[string]string //存储从os.Args获取的选项参数
var placeIndex int
var logger *logwa.Logger

func isShortName(name string) bool {
	return !isLongName(name) && strings.Index(name, "-") == 0
}

func isLongName(name string) bool {
	return strings.Index(name, "--") == 0
}

func init() {
	logger = logwa.NewLogger(logwa.Config{
		Level:      logwa.WARNING,
		UseColor:   true,
		Target:     os.Stdout,
		LogMethods: logwa.LevelOnlyMethods(),
	})
	flagArgs = map[string]bool{}
	PlaceArgs = []string{}
	placeIndex = 0
	optionArgs = map[string]string{}
	rawArgs := os.Args
	for i := 1; i < len(rawArgs); i++ {
		arg := rawArgs[i]
		if isShortName(arg) {
			if i+1 < len(rawArgs) && !isShortName(rawArgs[i+1]) && !isLongName(rawArgs[i+1]) { //后有值
				if len(arg) > 2 { //不允许参数合并又带值的情况
					logger.Erro("Illegal parameter %s %s", arg, rawArgs[i+1])
				}
				optionArgs[arg[1:]] = rawArgs[i+1]
				i++
			} else { //后无值
				for j := 1; j < len(arg); j++ {
					flagArgs[string(arg[j])] = true
				}
			}
		} else if isLongName(arg) {
			if i+1 < len(rawArgs) && !isShortName(rawArgs[i+1]) && !isLongName(rawArgs[i+1]) { //后有值
				optionArgs[arg[1:]] = rawArgs[i+1]
				i++
			} else { //后无值
				for j := 1; j < len(arg); j++ {
					flagArgs[string(arg[j])] = true
				}
			}
		} else {
			PlaceArgs = append(PlaceArgs, arg)
		}
	}
}
func lengthCheck(long string) {
	if len(long) <= 2 {
		logger.Erro("The length of long flagArgs %s should longer than 2", long)
	}
}

func Int(short byte, long string, defaultValue int, sever bool) int {
	lengthCheck(long)
	arg, ok := optionArgs[string(short)]
	if !ok {
		arg, ok = optionArgs[long]
	}

	if !ok {
		if sever {
			logger.Erro("The arg named %c or %s not found.", short, long)
		} else {
			return defaultValue
		}
	}

	res, err := strconv.Atoi(arg)
	if err != nil {
		logger.Erro("{_rx}Parse Int {;}: %s", err)
	}
	return res
}

func Str(short byte, long string, defaultValue string, sever bool) string {
	lengthCheck(long)
	arg, ok := optionArgs[string(short)]
	if !ok {
		arg, ok = optionArgs[long]
	}

	if !ok {
		if sever {
			logger.Erro("The arg named %c or %s not found.", short, long)
		} else {
			return defaultValue
		}
	}

	return arg
}
func Float(short byte, long string, defaultValue float32, sever bool) float32 {
	lengthCheck(long)
	arg, ok := optionArgs[string(short)]
	if !ok {
		arg, ok = optionArgs[long]
	}

	if !ok {
		if sever {
			logger.Erro("The arg named %c or %s not found.", short, long)
		} else {
			return defaultValue
		}
	}

	res, err := strconv.ParseFloat(arg, 32)
	if err != nil {
		logger.Erro("{_rx}Parse Float {;} %s", err)
	}
	return float32(res)
}

func Bool(short byte, long string, defaultValue bool, sever bool) bool {
	lengthCheck(long)
	arg, ok := optionArgs[string(short)]
	if !ok {
		arg, ok = optionArgs[long]
	}

	if !ok {
		if sever {
			logger.Erro("The arg named %c or %s not found.", short, long)
		} else {
			return defaultValue
		}
	}

	res, err := strconv.ParseBool(arg)
	if err != nil {
		logger.Erro("{rx}Parse Bool {;} %s", err)
	}
	return res
}

func HasNext() bool {
	return placeIndex < len(PlaceArgs)
}

func NextInt() int {
	intValue, err := strconv.Atoi(NextStr())
	if err != nil {
		logger.Erro("{rx}Parse Int {;} %s", err)
	}
	return intValue
}

func NextBool() bool {
	boolValue, err := strconv.ParseBool(NextStr())
	if err != nil {
		logger.Erro("{rx}Parse Bool {;} %s", err)
	}
	return boolValue
}

func NextStr() string {
	if !HasNext() {
		logger.Erro("{rx}No more arg{;}")
	}
	res := PlaceArgs[placeIndex]
	placeIndex++
	return res
}

func NextFloat() float32 {
	floatValue, err := strconv.ParseFloat(NextStr(), 32)
	if err != nil {
		logger.Erro("{rx}Parse float {;}: %s", err)
	}
	return float32(floatValue)
}

func main() {
	fmt.Println(NextFloat())
}
