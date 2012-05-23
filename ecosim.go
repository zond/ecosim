
package ecosim

import (
	"time"
	. "../gomarket"
)

type Engine struct {
	actors map[Actor]bool
	market *Market
}
func NewEngine() *Engine {
	return &Engine{make(map[Actor]bool), NewMarket()}
}
func (e *Engine) Add(a Actor) {
	e.actors[a] = true
	e.market.Add(a)
}
func (e *Engine) Del(a Actor) {
	delete(e.actors, a)
	e.market.Del(a)
}
func (e *Engine) profit(output Output) {
	return_value := 0.0
	for resource, units := range output.output {
		if price, ok := e.market.Price(resource); ok {
			return_value = return_value + price * units * output.cycles
		} else {
			return_value = return_value + units * output.cycles
		}
	}
	return return_value
}
func (e *Engine) Run(t time.Duration) {
	for actor,_ := range e.actors {
		best_process, best_profit, best_output := Process(nil), 0.0, Output(nil)
		next_best_profit := 0.0
		for _,process := range actor.Processes() {
			output := process.Run(t)
			profit := e.Profit(output)
			if profit > best_profit {
				next_best_profit = best_profit
				best_process, best_profit, best_output = process, profit, output
			} else if profit > next_best_profit {
				next_best_profit = profit
			}
		}
		if best_output != nil {
			actor.Update(&Update{best_output, t, next_best_profit})
		}
	}
	e.market.Trade()
}

type Update struct {
	Output Output
	time time.Duration
	opportunityCost float64
}

type Output struct {
	Actual map[Resource]float64
	Potential map[Resource]float64
}

type Process interface {
	Run(time.Duration) Output
}

type Actor interface {
	Trader
	Carrier
	Processes() []Process
	Update(Update)
}

type Skill interface{}

type StandardActor struct {
	StandardTrader
	processes []Process
	skills map[Skill]float64
	resources map[Resource]float64
}
func (s *StandardActor) AddProcess(factory *StandardProcessFactory) {
	s.processes = append(s.processes, factory.produce(s))
}
func (s *StandardActor) Processes() []Process {
	return s.processes
}
func (s *StandardActor) Update(update Update) {
	fmt.Println(s,"updated with",update)
}
