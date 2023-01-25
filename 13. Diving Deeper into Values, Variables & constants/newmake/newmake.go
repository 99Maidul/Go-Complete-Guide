package newmake

import "fmt"

func main() {

	hobbies := make([]string, 2, 10)
	hobbies[0] = "Sports"
	hobbies[1] = "Reading"

	hobbies = append(hobbies, "Cooking", "Dancing")

	fmt.Println(hobbies)

	morehobbies := new([]string)
	*morehobbies = append(*morehobbies, "Sports", "Reading", "Cooking", "Dancing")
	fmt.Println(*morehobbies)
}
