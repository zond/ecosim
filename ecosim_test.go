
package ecosim

import (
	"testing"
	"time"
	. "../gomarket"
)

type TestActor struct {
	Trader
	ran map[Process]bool
	processes []Process
}
func (t *TestActor) Buy(bid, ask *Order, price float64) {
}
func (t *TestActor) Deliver(bid, ask *Order, price float64) {
}
func (t *TestActor) Processes() []Process {
	return t.processes
}
func (t *TestActor) Update(update *Update) {
	t.ran[update.Profit.Output.Process] = true
}

type TestProcess struct {
	output Resources 
}
func (t *TestProcess) Run(time time.Duration) *Output {
	return &Output{t, t.output, t.output}
}
func (t *TestProcess) Avoid(time time.Duration) *Output {
	return &Output{t, make(Resources), make(Resources)}
}


func TestInitialRun(t *testing.T) {
	e := NewEngine()
	actor1 := &TestActor{nil, make(map[Process]bool), nil}
	actor1.Trader = NewStandardTrader(actor1)
	farming := &TestProcess{Resources{"rice": 10.0, "tools": -1.0}}
	smithing := &TestProcess{Resources{"tools": 3.0, "ore": -1.0}}
	actor1.processes = append(actor1.processes, farming)
	actor1.processes = append(actor1.processes, smithing)
	e.Add(actor1)
	e.Run(time.Second)
	if !actor1.ran[farming] {
		t.Error("should have farmed")
	}
	if actor1.ran[smithing] {
		t.Error("should not have smithed")
	}
}