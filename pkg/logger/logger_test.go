package logger

import "testing"

func TestXxx(t *testing.T) {
	InitLogger(false)
	Debug("debug")
	Debugf("debugf: %s", "debugf")
	NDebug("LOGGER", "name debug")
	NDebugf("LOGGER", "name debug: %s", "ndebugf")
}
