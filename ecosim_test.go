
package ecosim

import (
	"testing"
	"time"
	"../gomarket"
)

type TestProcess struct {
	projection map[gomarket.Resource]float64
	ran bool
}
func NewTestProcess(projection map[gomarket.Resource]float64) *TestProcess {
	return &TestProcess{projection, false}
}
func (t *TestProcess) Project(ti time.Duration) map[gomarket.Resource]float64 {
	return t.projection
}
func (t *TestProcess) Run(ti time.Duration) map[gomarket.Resource]float64 {
	t.ran = true
	return t.projection
}

type TestActor struct {
	processes []Process
}
func (a *TestActor) Processes() []Process {
	return a.processes
}


func TestInitialRun(t *testing.T) {
	e := NewEngine()
	farming := NewTestProcess(map[gomarket.Resource]float64{"rice": 10.0, "tools": -1.0})
	smithing := NewTestProcess(map[gomarket.Resource]float64{"tools": 3.0, "ore": -1.0})
	actor1 := &TestActor{[]Process{farming, smithing}}
	e.Add(actor1)
	e.Run(time.Second)
	if !farming.ran {
		t.Error("should have farmed")
	}
	if smithing.ran {
		t.Error("should not have smithed")
	}
}