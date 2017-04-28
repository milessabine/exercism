// Clock stub file

// Package clock includes functions to create and .
package clock

import (
	"fmt"
	"math"
)

// The value of testVersion here must match `targetTestVersion` in the file
// clock_test.go.
const testVersion = 4

// Clock API as stub definitions.  No, it doesn't compile yet.
// More details and hints are in clock_test.go.
// Complete the type definition.  Pick a suitable data type.
type Clock struct {
	Hour   int
	Minute int
}

// New creates a new instance of a type Clock
func New(hour, minute int) Clock {
	t := Clock{Hour: hour, Minute: minute}
	t.adjust()
	return t
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minute)
}

func (c Clock) Add(minutes int) Clock {
	c.Minute = c.Minute + minutes
	c.adjust()
	return c
}

func (c *Clock) adjust() {
	if c.Minute >= 60 {
		newhour := c.Minute / 60
		c.Minute = c.Minute % 60
		if newhour >= 1 {
			c.Hour = c.Hour + newhour
		}
	}
	if c.Minute < 0 {
		newhour := float64(c.Minute) / 60.0
		c.Minute = c.Minute%60 + 60
		if newhour < 0 {
			c.Hour = c.Hour + int(math.Floor(newhour))
		}
	}
	if c.Hour < 0 {
		c.Hour = c.Hour%24 + 24

	}
	if c.Hour >= 24 {
		c.Hour = c.Hour % 24
	}

}

// Remember to delete all of the stub comments.
// They are just noise, and reviewers will complain.
