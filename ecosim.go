
package ecosim

import (
	"time"
	"../gomarket"
)

type Engine struct {
	actors map[Actor]bool
	market *gomarket.Market
}
func NewEngine() *Engine {
	return &Engine{make(map[Actor]bool), gomarket.NewMarket()}
}
func (e *Engine) Add(a Actor) {
	e.actors[a] = true
}
func (e *Engine) Del(a Actor) {
	delete(e.actors, a)
}
func (e *Engine) Profit(amounts map[gomarket.Resource]float64) float64 {
	return_value := 0.0
	for resource, units := range amounts {
		if price, ok := e.market.Price(resource); ok {
			return_value = return_value + price * units
		} else {
			return_value = return_value + units
		}
	}
	return return_value
}
 func (e *Engine) Run(t time.Duration) {
	for actor,_ := range e.actors {
		best_process, best_profit := Process(nil), 0.0
		next_best_profit := 0.0
		for _,process := range actor.Processes() {
			profit := e.Profit(process.Project(t))
			if profit > best_profit {
				next_best_profit = best_profit
				best_process, best_profit = process, profit
			} else if profit > next_best_profit {
				next_best_profit = profit
			}
		}
		if best_process != nil {
			best_process.Run(t)
		}
	}
	e.market.Trade()
}

type Process interface {
	Project(t time.Duration) map[gomarket.Resource]float64
	Run(t time.Duration) map[gomarket.Resource]float64
}


type Actor interface {
	Processes() []Process
}
