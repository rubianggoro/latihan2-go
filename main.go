package main

import ("fmt"
		"time")

func main()  {

	// SOAL NO. 1

	for i :=1; i <=100; i++ {
		if i % 15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i % 3 == 0 {
			fmt.Println("Fizz")
		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}


	// SOAL No. 2
	
	firstDate := time.Date(2020, time.January, 3, 23, 0, 0, 0, time.UTC)

	secondDate := time.Date(2020, time.February, 10, 23, 0, 0, 0, time.UTC)

	diff := secondDate.Sub(firstDate)

	days := int(diff.Hours() / 24)

	fmt.Println (days, "hari")



}