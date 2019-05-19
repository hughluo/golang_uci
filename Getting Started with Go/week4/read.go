/* Write a program which reads information from a file and represents it in a slice of structs. 
 * Assume that there is a text file which contains a series of names.
 * Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

 * Your program will define a name struct which has two fields, fname for the first name, and lname for the last name. 
 * Each field will be a string of size 20 (characters).

 * Your program should prompt the user for the name of the text file. 
 * Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file. 
 * Each struct created will be added to a slice, and after all lines have been read from the file, 
 * your program will have a slice containing one struct for each line in the file. 
 * After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.
 */
package main
import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
)

func main() {
	type name struct {
		fname string
		lname string
	}

	name_slice := make([]name, 0)

	var file_path string
	fmt.Println("Please enter the name of the text file...")
	fmt.Scan(&file_path)
	f, err := os.Open(file_path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		t_s := strings.Split(t, " ")
		n := name{t_s[0], t_s[1]}
		name_slice = append(name_slice, n)
	}

	if err := scanner.Err(); err != nil {
        log.Fatal(err)
	}

	for _, name := range name_slice {
		fmt.Printf("%v %v \n", name.fname, name.lname)
	}
}
