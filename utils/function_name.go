package utils

import (
	"fmt"
	"runtime"
	"strings"
)

func PrintFuncName() {
	pc := make([]uintptr, 10)  // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	funcNameParts := strings.Split(f.Name(), ".")
	fmt.Printf("=== %s ===\n", funcNameParts[len(funcNameParts)-1])
}
