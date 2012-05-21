
package ecosim

import (
	"time"
	"../gomarket"
)

type Actor interface {
	gomarket.Trader
	gomarket.Carrier
	Processes() []Process
}

type UpdatableActor interface {
	Actor
	Update(time.Duration, map[gomarket.Resource]float64)
}

type ProcessFunc func(Actor, time.Duration) map[gomarket.Resource]float64

type ProcessWrapper struct {
	actor UpdatableActor
	process ProcessFunc
}
func (p *ProcessWrapper) Project(t time.Duration) map[gomarket.Resource]float64 {
	return p.process(p.actor, t)
}
func (p *ProcessWrapper) Run(t time.Duration) {
	p.actor.Update(t, p.process(p.actor, t))
}



type StandardActor struct {
	*gomarket.StandardTrader
	Resources map[gomarket.Resource]float64
	processes []Process
}
func NewStandardActor() *StandardActor {
	s := &StandardActor{nil, make(map[gomarket.Resource]float64), nil}
	s.StandardTrader = gomarket.NewStandardTrader(s)
	return s
}
func (s *StandardActor) Buy(bid, ask *gomarket.Order, price float64) {
}
func (s *StandardActor) Deliver(bid, ask *gomarket.Order, price float64) {
}
func (s *StandardActor) AddProcess(p Process) {
	s.processes = append(s.processes, p)
}
func (s *StandardActor) AddProcessFunc(p ProcessFunc) {
	s.processes = append(s.processes, &ProcessWrapper{s, p})
}
func (s *StandardActor) Processes() []Process {
	return s.processes
}
func (s *StandardActor) Update(t time.Duration, delta map[gomarket.Resource]float64) {
	for resource, unit_delta := range delta {
		if _, ok := s.Resources[resource]; ok {
			s.Resources[resource] = s.Resources[resource] + unit_delta
		} else {
			s.Resources[resource] = unit_delta
		}
	}
}

