// Loops

func main() {
	i := 0

	for i < 5 {
		fmt.Println(i)
		i++
	}
}

// -------------------
thisListIsGood := []int{"hi", "hello"}

for i, v := range thisListIsGood {
	fmt.Printf("%d, %d", i, v)
}