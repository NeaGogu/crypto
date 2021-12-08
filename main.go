package main

import "fmt"

func initS(KEY [5]int, S [256]int) [256]int {
	j := 0
	for i := 0; i < 256; i++ {
		j = (j + S[i] + KEY[i%len(KEY)]) % 256
		S[i], S[j] = S[j], S[i]
	}

	return S
}

func generateOutput(n int, S [256]int) []int {
	i := 0
	j := 0
	var output []int

	for t := 0; t < n; t++ {
		i = (i + 1) % 256
		j = (j + S[i]) % 256
		S[i], S[j] = S[j], S[i]
		output = append(output, S[(S[i]+S[j])%256])
	}

	return output

}

func compareOutput(output []int, outputToMatch []int) bool {
	for i, val := range output {
		if outputToMatch[i] != val {
			return false
		}
	}
	return true
}

func breakAlgo(start int, stop int) {
	expectedOutput := []int{130, 189, 254, 192, 238, 132, 216, 132, 82, 173}
	var S [256]int
	var S_copy [256]int
	var output []int

	KEY := [5]int{80, 0, 0, 0, 0}

	for i := 0; i < 256; i++ {
		S[i] = i
	}

out:
	for k2 := start; k2 < stop; k2++ {
		for k3 := 0; k3 < 256; k3++ {
			for k4 := 0; k4 < 256; k4++ {
				for k5 := 0; k5 < 256; k5++ {
					KEY[1] = k2
					KEY[2] = k3
					KEY[3] = k4
					KEY[4] = k5

					S_copy = S
					S_copy = initS(KEY, S_copy)
					output = generateOutput(10, S_copy)
					if compareOutput(output, expectedOutput) {
						fmt.Println("FOUND KEY")
						fmt.Printf("KEY: %d", KEY)
						fmt.Printf("SEQUENCE: %d", output)
						break out
					}
				}
			}
		}
	}

}

func testKey(key [5]int) {
	var S [256]int
	var output []int

	for i := 0; i < 256; i++ {
		S[i] = i
	}
	S = initS(key, S)
	output = generateOutput(10, S)
	fmt.Print(output)

}

func main() {

	go breakAlgo(0, 50)
	fmt.Println("Started 1")
	go breakAlgo(50, 100)
	fmt.Println("Started 2")
	go breakAlgo(100, 150)
	fmt.Println("Started 3")
	go breakAlgo(150, 200)
	fmt.Println("Started 4")
	breakAlgo(200, 258)
}
