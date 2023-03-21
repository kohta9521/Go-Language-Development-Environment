package main

func main() {
	modify := func(a []string, f func([]string) []string) {
		return f(a)
	}

	m := []string{
		"1st", "2nd", "3rd",
	}
}
