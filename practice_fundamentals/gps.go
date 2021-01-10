// Experiment: gps.go

// As a final step, create a rover structure that embeds the gps and write a main function to test everything out.
// Initialize a GPS for Mars with a current location of Bradbury Landing (-4.5895, 137.4417) and a destination of
// Elysium Planitia (4.5, 135.9). Then create a curiosity rover and print out its message (which forwards to the
// gps).

package practicefundamentals

import "fmt"

// gps structure for a Global Positioning System (GPS) composed of a current location, destination location, and
// a world.
type gps struct {
	current      location
	desitination location
	world
}

// description for the location type that returns a string containing the name, latitude, and
// longitude.
func (l location) description() string {
	return fmt.Sprintf("%+v\n", l)
}

// distance method for gps structure that finds the distance between the current and destination locations.
func (gps gps) distance() float64 {
	return gps.world.distance(gps.current, gps.desitination)
}

// message method fpr gps structure returns a string describing how many kilometers remain to the destination.
func (gps gps) message() string {
	return fmt.Sprintf("We are %v kms awaw from our destination\n", gps.distance())
}

// As a final step, create a rover structure that embeds the gps and write a main function to test everything out.
type rover struct {
	name string
	gps
}
