
package ecosim

import (
	"time"
	"fmt"
	. "../gomarket"
)

/*
 * Helps us to calculate profits.
 */
type ProfitCalculator struct {
	market *Market
	avoidanceCosts map[Process]*Profit
	time time.Duration
}
/*
 * Create a new calculator. 
 * Will go through all the processes of the actor and calculate the losses incurred 
 * by NOT doing them.
 */
func NewProfitCalculator(market *Market, actor Actor, t time.Duration) *ProfitCalculator {
	calculator := &ProfitCalculator{market, make(map[Process]*Profit), t}
	for _,process := range actor.Processes() {
		calculator.avoidanceCosts[process] = process.Avoid(t).Profit(market)
	}
	return calculator
}
/*
 * Calculate the profit (or loss) of running the given process.
 * 
 * Calculated by adding all the costs of avoiding OTHER processes to the profit of running
 * the given process.
 */
func (e *ProfitCalculator) processProfit(process Process) *Profit {
	if _, ok := e.avoidanceCosts[process]; ok {
		profit := process.Run(e.time).Profit(e.market)
		for avoidedProcess, cost := range e.avoidanceCosts {
			if process != avoidedProcess {
				profit.MergeIn(cost)
			}
		}
		return profit
	} else {
		panic(fmt.Sprint(process,"is not in",e.avoidanceCosts))
	}
	return (*Profit)(nil)
}
