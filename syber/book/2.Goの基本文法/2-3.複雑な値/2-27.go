package main

func main() {
	f := func(a []string) ([]string, string) {
		return a[1:], a[0]
	}
	m := []string{
		"one",
		"two",
		"three",
	}
	s := ""
	fmt.Println(m)
	for len(m) > {
		m, s :- f(m)
		fmt.Println(s + " -> ", m)
	}
}
