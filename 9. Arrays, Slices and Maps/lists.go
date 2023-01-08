package main

import "fmt"

type Product struct {
	id          string
	title       string
	description string
	price       float64
}

func main() {
	//1 Make an array of hobby
	hobbies := []string{"Programming", "Reading", "Writing", "Drawing", "Swimming"}
	fmt.Println(hobbies)
	//2 Print the first element of the array
	//2nd & 3rd element combined
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])
	//3 Create a slice based on 1st element that contains 1st and 2nd element
	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)
	// 4 Reslice from 3 and change it to contation 3rd and second element from original array
	fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)
	//5 Create dynamic array of course goals
	courseGoals := []string{"Learn Go", "Learn Python", "Learn Java"}
	fmt.Println(courseGoals)
	//6 Set 2nd element to "Learn C++" and "Learn python" to 4th element
	courseGoals[1] = "Learn C++"
	courseGoals = append(courseGoals, "Learn Python")
	fmt.Println(courseGoals)
	//7 Create Product struct and add dynamic list of structs
	products := []Product{
		{"first-product",
			"A first Product",
			"This is the first product",
			12.99},
		{"second-product",
			"A second Product",
			"This is the second product",
			13.99},
	}
	fmt.Println(products)
}
