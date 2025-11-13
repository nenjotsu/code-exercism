package greeting

// HelloWorld greets the world.
func HelloWorld() string {
	return "Hello, World!"
}

func Gcd(p int, q int) int {
	if q == 0 {
		return p
	}

	r := p % q
	return Gcd(q, r)
}

// fmt.Println()
