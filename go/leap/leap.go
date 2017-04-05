// Package leap supplies a function to calculate a leap year.
/*8on every year that is evenly divisible by 4
  except every year that is evenly divisible by 100
    unless the year is also evenly divisible by 400
*/
package leap

// testVersion should match the targetTestVersion in the test file.
const testVersion = 3

// IsLeapYear determines if a year is a leap year or not.
func IsLeapYear(year int) bool {
	if year%4 != 0 {
		return false
	}
	if year%100 == 0 && year%400 != 0 {
		return false
	}
	return true
}
