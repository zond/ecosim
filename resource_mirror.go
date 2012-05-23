
package ecosim

import (
	. "../gomarket"
	"math"
)

type ResourceMirror struct {
	original map[Resource]float64
	delta map[Resource]float64
}
func NewResourceMirror(original map[Resource]float64) {
	return &ResourceMirror{original, make(map[Resource]float64)}
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
func (r *ResourceMirror) Delta() map[Resource]float64 {
	return delta
}
