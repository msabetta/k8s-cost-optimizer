package unit

import (
	"testing"

	"k8s-cost-optimizer/pkg/logger"
)

func TestLoggerInit(t *testing.T) {
	logger.Init()
	if logger.Info == nil || logger.Error == nil {
		t.Fatal("Logger not initialized")
	}
}
