
package ecosim

import (
	"time"
	"fmt"
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
func (e *Engine) Run(t time.Duration) {
	for actor,_ := range e.actors {
		actor.run(t)
	}
	e.market.Trade()
}



/*
 * The immediate and eventual output of a process.
 * Eventual output is defined as actual output + the fraction of not-yet-produced output for the next cycle.
 */
type Output struct {
	Process Process
	Immediate Resources
	Eventual Resources
}
func (o *Output) MergeIn(other *Output) {
	o.Immediate.MergeIn(other.Immediate)
	o.Eventual.MergeIn(other.Eventual)
}
func (o *Output) Profit(market *Market) *Profit {
	return &Profit{o, market.Value(o.Immediate), market.Value(o.Eventual)}
}

type Update struct {
	Profit *Profit
	OpportunityCost *Profit
	Time time.Duration
}

/*
 * The Output of a process along with its immedate and eventual profit (or, the avoidance costs along with their loss)
 */
type Profit struct {
	Output *Output
	Immediate float64
	Eventual float64
}
func (p *Profit) MergeIn(o *Profit) {
	p.Output.MergeIn(o.Output)
	p.Immediate += o.Immediate
	p.Eventual += p.Eventual
}

type Process interface {
	/*
	 * The results when running this process for a time.
	 */
	Run(time.Duration) *Output
	/*
	 * The results when avoiding this process for a time.
	 */
	Avoid(time.Duration) *Output
}

type Skill interface{}

type Skills map[Skill]float64

type Actor struct {
	StandardTrader
	processes []Process
	skills Skills
	resources Resources
}
func (s *Actor) AddProcess(factory *StandardProcessFactory) {
	s.processes = append(s.processes, factory.produce(s))
}
func (s *Actor) run(t time.Duration) {
	/*
	 * Find the most profitable process P to run with present resources.
	 */
	/*
	 * Run P.
	 */
	/*
	 * Apply the results of running P.
	 */
	/*
	 * For each available process p.
	 */
		/*
		 * Find out the profit of p regardless of present resources.
		 */
		/*
		 * Set the internal value v(r) of all by p used resources 
		 * proportional to their market price so that the sum of
		 * their value is equal to the market value of their 
		 * produced resources.
		 * Remember if p is the process M with the highest profit (delta of value and cost).
		 */
	/*
	 * For each resource r that we have or is required by M.
	 */
		/*
		 * Put in an Ask for v(r)+1 for r for whatever amount we have.
		 */
		/*
		 * Put in a Bid for v(r) for the delta between what M requires and what we have.
		 */
}
