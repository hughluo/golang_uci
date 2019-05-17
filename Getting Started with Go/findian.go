// Write a program which prompts the user to enter a string. The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’. The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’. The program should print “Not Found!” otherwise. The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.
package main
import "fmt"
import "strings"

func main() {
	var ipt string
	fmt.Scan(&ipt)
        lower_ipt := strings.ToLower(ipt)
	if lower_ipt[0] == byte('i') && lower_ipt[len(lower_ipt)-1] == byte('n') && strings.Contains(lower_ipt, "a"){
		fmt.Println("Found!")

	} else {
		fmt.Println("Not Found!")
	}

}
