/* Write a program which prompts the user to first enter a name, and then enter an address. 
 * Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively. 
 * Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
 */
package main
import(
	"fmt"
	"encoding/json"
)

func main() {
	var m = map[string]string{}
	var name string
	var address string
	fmt.Println("Please enter a name...")
	fmt.Scan(&name)
	fmt.Println("Please enter an address...")
	fmt.Scan(&address)
	m["name"] = name
	m["address"] = address
	barr, err := json.Marshal(m)
	if err != nil {
		return
	}
	fmt.Println(string(barr))
}
