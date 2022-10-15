package main

import (
	"fmt"

	"github.com/monojitbarua/go-util-lib/error"
	"github.com/monojitbarua/go-util-lib/logger"
)

func main() {

	// checking error implementation
	x, err := errorCheck(0)
	if err != nil {
		fmt.Println("Value is zero")
	} else {
		fmt.Println("Value is not zero: ", x)
	}

	//
	// logger implementation
	logger.Info("Application logging started.")
	logger.Warn("Application logging sample warn message.")
	logger.Error("Application logging sample error message.")
}

func errorCheck(x int) (int, *error.ApplicationError) {
	if x == 0 {
		return 0, error.UnexpectedError("Zero value not expected")
	}
	return x, nil
}
