package parser

import (
	"fmt"
	"strings"
)

var deferredPrints []func()

var traceLevel int = 0

const traceIdentPlaceholder string = "\t"

func identLevel() string {
	return strings.Repeat(traceIdentPlaceholder, traceLevel-1)
}

func deferPrint(statement string) {
	deferredPrints = append(deferredPrints, func() {
		fmt.Print(statement)
	})
}

func printDeferred() {
	incIdent()
	fmt.Printf("%s", identLevel())
	for _, f := range deferredPrints {
		f()
	}
	fmt.Print("\n")
	deferredPrints = []func(){}
	decIdent()
}

func tracePrintln(fs string) {
	fmt.Printf("%s%s\n", identLevel(), fs)
}

func incIdent() { traceLevel = traceLevel + 1 }
func decIdent() { traceLevel = traceLevel - 1 }

func trace(msg string) string {
	incIdent()
	tracePrintln("BEGIN " + msg)
	return msg
}

func untrace(msg string) {
	tracePrintln("END " + msg)
	decIdent()
}
