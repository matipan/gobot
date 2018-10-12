package ollie

import (
	"testing"

	"github.com/matipan/gobot"
	"github.com/matipan/gobot/gobottest"
)

var _ gobot.Driver = (*Driver)(nil)

func initTestOllieDriver() *Driver {
	d := NewDriver(NewBleTestAdaptor())
	return d
}

func TestOllieDriver(t *testing.T) {
	d := initTestOllieDriver()
	d.SetName("NewName")
	gobottest.Assert(t, d.Name(), "NewName")
}

func TestOllieDriverStartAndHalt(t *testing.T) {
	d := initTestOllieDriver()
	gobottest.Assert(t, d.Start(), nil)
	gobottest.Assert(t, d.Halt(), nil)
}
