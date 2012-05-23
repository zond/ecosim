
package ecosim

import (
	. "../gomarket"
	"math"
)

type ResourceMirror struct {
	original Resources
	delta Resources
}
func NewResourceMirror(original Resources) {
	return &ResourceMirror{original, make(Resources)}
}
func (r *ResourceMirror) Left(resource Resource) {
	return math.Max(0.0, r.original[resource] + r.delta[resource])
}
func (r *ResourceMirror) Produce(resource Resource, units float64) {
	r.delta[resource] = r.delta[resource] + units
}
func (r *ResourceMirror) Consume(resource Resource, units float64) float64 {
	left := r.Left(resource)
	returnValue := math.Min(left, units)
	r.delta[resource] = r.delta[resource] - returnValue
	return returnValue
}
func (r *ResourceMirror) Delta() Resources {
	return delta
}
