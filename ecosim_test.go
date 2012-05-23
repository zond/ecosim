
package ecosim

import (
	"testing"
	"time"
	"../gomarket"
)

type TestActor struct {
	gomarket.Trader
	ran map[Process]bool
	processes []Process
}
func (t *TestActor) Buy(bid, ask *gomarket.Order, price float64) {
}
func (t *TestActor) Deliver(bid, ask *gomarket.Order, price float64) {
}
func (t *TestActor) Processes() []Process {
	return processes
}
func (t *TestActor) Update(process Process, t time.Duration) {
	t.ran[process] = true
}


func TestInitialRun(t *testing.T) {
	e := NewEngine()
	actor1 := &TestActor{gomarket.NewStandardTrader(), make(map[Process]bool)}
	farming := func(t time.Duration) map[gomarket.Resource]float64 {
		return map[gomarket.Resoure]float64{"rice": 10.0, "tools": -1.0}
	}
	smithing := func(t time.Duration) map[gomarket.Resource]float64 {
		return map[gomarket.Resoure]float64{"tools": 3.0, "ore": -1.0}
	}
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