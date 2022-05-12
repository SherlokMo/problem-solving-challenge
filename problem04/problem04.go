package problem04

func Compose(f, g func(x int) int) func(x int) int {

	return func(x int) int {
		x = g(x)
		x = f(x)
		return x
	}
}
