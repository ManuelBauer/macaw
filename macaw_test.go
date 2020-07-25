package macaw_test

import (
	"testing"

	"github.com/ManuelBauer/macaw"
)

func TestDebug(t *testing.T) {
	logger := macaw.CreateLogger("Example")
	logger.Debug("Debug")
}

func TestInfo(t *testing.T) {
	logger := macaw.CreateLogger("Example")
	logger.Info("Info")
}

func TestWarning(t *testing.T) {
	logger := macaw.CreateLogger("Example")
	logger.Warning("Warning")
}

func TestError(t *testing.T) {
	logger := macaw.CreateLogger("Example")
	logger.Error("Error")
}

func TestSubLogger(t *testing.T) {
	logger := macaw.CreateLogger("Example")
	sub := logger.CreateSubLogger("Sub logger")
	sub.Info("Message from the sub logger")
}

func TestSubLoggerOfSubLogger(t *testing.T) {
	logger := macaw.CreateLogger("Example")
	sub := logger.CreateSubLogger("Sub logger")
	sub2 := sub.CreateSubLogger("Additional sub logger")
	sub.Info("Message from the sub logger")
	sub2.Info("Message from additional sub logger")
}

func TestRaceCondition(t *testing.T) {
	logger := macaw.CreateLogger("Race condition")
	logger2 := macaw.CreateLogger("Race condition 2")

	debugString := "Debug message"

	for i := 0; i < 1000; i++ {
		logger.Debug(debugString, i)
		logger2.Info("Info message", i)
	}
}
