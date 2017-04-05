//Package greeting implements the HelloWorld function.
package greeting

const testVersion = 3

// HelloWorld greets user with name if given, else it returns "Hello World!"
func HelloWorld(name string) string {
	if name != "" {
		return "Hello, " + name + "!"
	}

	return "Hello, World!"
}
