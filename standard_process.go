
package ecosim

import (
	"time"
	. "../gomarket"
	"math"
)

/*
 * How much something has been used, and what effect it had.
 */
type Usage struct {
	Effect float64
	Amount float64
}

/*
 * A required resource (skill, actual resource, whatever)
 */
type Requirement struct {
	/*
	 * The amount of this required resource that gives an effect of 1.0.
	 */
	norm float64
	/*
	 * The maximum applicable amount of this required resource.
	 */
	max float64
}
func (r *Requirement) used(available float64) float64 {
	return math.Max(0.0, math.Min(r.max, available))
}
/*
 * Apply this requirement to an amount of available resource.
 * Will (for this calculation only) modify the norm with the given factor.
 */
func (r *Requirement) Apply(available float64, factor float64) Usage {
	amount := r.used(available)
	effect := r.impact * (used / (r.norm * factor))
	return &Usage{effect, amount}
}

/*
 * In effect the configuration of a Process.
 *
 * The output will be the standardOutput multiplied with a production factor.
 * The production factor is the product of the required factor and the useful factor.
 * The required factor is the product of the effects of the required skills and resources.
 * The useful factor is the sum of the effects of the useful skills and resources.
 *
 * The resource cost is the sum of the used up required and useful resources.
 * The amount of resources used is the max usage of each resource requirement.
 * The effect of the used resources is multiplied 
 */
type StandardProcessFactory struct {
	requiredSkills map[Skill]Requirement
	usefulSkills map[Skill]Requirement
	requiredResources map[Resource]Requirement
	usefulResources map[Resource]Requirement
	standardOutput map[Resource]float64
	requiredTime time.Duration
}
func (s *StandardProcessFactory) produce(a *StandardActor) *StandardProcess {
	return &StandardProcess{s, a, 0.0}
}

type StandardProcess struct {
	*StandardProcessFactory
	actor *StandardActor
	progress float64
}
/*
 * Calculate the effect of the skills of our actor.
 */
func (s *StandardProcess) skillFactor() float64 {
	/*
	 * Factor is the product of the relevant skills and their impact
	 */
	requiredScale := 1.0
	for skill, requirement := range s.requiredSkills {
		requiredScale = requiredScale * requirement.Apply(s.actor.skills[skill], 1.0).Effect
	}
	/*
	 * Factor is the sum of the relevant skills and their impact
	 */
	usefulScale := 0.0
	for skill, requirement := range s.usefulSkills {
		usefulScale = usefulScale + requirement.Apply(s.actor.skills[skill], 1.0).Effect
	}
	return requiredScale * usefulScale
}
/*
 * Consume the resources for one cycle of this process and return the production factor effect.
 */
func (s *StandardProcess) consume(mirror *ResourceMirror, costFactor float64) float64 {
	/*
	 * Factor is the product of the availability of the required resources
	 */
	requiredFactor := 1.0
	for resource, requirement := range s.requiredResources {
		usage := requirement.Apply(mirror.Left(resource), costFactor)
		requiredFactor = requiredFactor * usage.Effect
		mirror.Consume(resource, usage.Amount)
	}
	/*
	 * Factor is the sum of the availability of the required resources
	 */
	usefulFactor := 1.0
	for resource, requirement := range s.usefulResources {
		usage := requirement.Apply(mirror.Left(resource), costFactor)
		requiredFactor = usefulFactor + usage.Effect
		mirror.Consume(resource, usage.Amount)
	}
	return requiredFactor * usefulFactor
}
/*
 * Append the output of this process to the given mirror
 */
func (s *StandardProcess) appendOutput(productionFactor float64, mirror *ResourceMirror) {
	for resource, units := range s.output {
		mirror.Produce(resource, productionFactor * units)
	}
}
/*
* Append the costs and production of one cycle to the mirror, taking a skill factor into account.
 */
func (s *StandardProcess) resources(mirror *ResourceMirror, skillFactor float64) {
	/*
	 * The cost is ameliorated by the square root of the skill.
	 */
	costFactor := 1.0 / math.Sqrt(skillFactor)
	/*
	 * How much does lack of resources hinder production?
	 */
	resourceFactor := s.consume(mirror, costFactor)
	if resourceFactor > 0.0 {
		/*
		 * What we actually produce is proportional to how competent we are
		 * and any lack of resources.
		 */
		actualProduction := skillFactor * resourceFactor
		
		s.appendOutput(actualProduction, mirror)
		return result
	}
}
func (s *StandardProcess) Run(t time.Duration) Output {
	/*
	 * How skilled is the actor? Will affect both costs and cycle length.
	 */
	skillFactor := s.skillFactor()
	/*
	 * How long will our cycle be, based on our skill?
	 */
	cycleLength := s.requiredTime / skillFactor;
	/*
	 * How many cycles are contained within (time spent on unfinished cycle so far + t)?
	 */
	cycles := (s.progress + t) / cycleLength
	/*
	 * Make a mirror of our resources for this experiment.
	 */
	mirror := NewResourceMirror(s.actor.resources)
	/*
	 * Iterate int64(cycles) times to see how much we will produce during t.
	 */
	
	/*
	 * Iterate one more time to see how much we would have produced if we had one more cycle to go,
	 * Then multiply that with the fraction of a cycle that we managed to finish to calculate the
	 * eventual results.
	 */

	
	return &Output{...}
}

