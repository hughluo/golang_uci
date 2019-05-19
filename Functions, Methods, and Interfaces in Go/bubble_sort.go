/* Write a Bubble Sort program in Go. The program should prompt the user to type in a sequence of up to 10 integers. 
 * The program should print the integers out on one line, in sorted order, from least to greatest. 
 * Use your favorite search tool to find a description of how the bubble sort algorithm works.

 * As part of this program, you should write a function called BubbleSort() which takes a slice of integers as an argument and returns nothing. 
 * The BubbleSort() function should modify the slice so that the elements are in sorted order.

 * A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position of two adjacent elements in the slice. 
 * You should write a Swap() function which performs this operation. 
 * Your Swap() function should take two arguments, a slice of integers and an index value i which indicates a position in the slice. 
 * The Swap() function should return nothing, but it should swap the contents of the slice in position i with the contents in position i+1.
 */
package main

import (

	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main() {
	
	fmt.Println("Please type in sequence of integer seperated by whitespace...")
	reader := bufio.NewReader(os.Stdin)
	ipt_raw, _ := reader.ReadString('\n')
	ipt := strings.TrimSuffix(ipt_raw, "\n")
	sli := StringToSlice(ipt)
	
	BubbleSort(sli)
	fmt.Println(sli)
}

func StringToSlice(s string) []int {
	str_sli := strings.Split(s, " ")
	int_sli := make([]int, 0)
	for _, e := range str_sli {
		num, err := strconv.Atoi(e)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Non-Integer chracter detected. Abort...")
			break
		}
		int_sli = append(int_sli, num)
	}
	return int_sli
}

func BubbleSort(sli []int) {
	for i := range sli {
		for j := 0 ; j < len(sli) - i; j++ {
			if j+1 < len(sli) && sli[j+1] < sli[j] {
				Swap(sli, j)
			}
		} 
	}
}

func Swap(sli []int, i int) {
	tmp := sli[i]
	sli[i] = sli[i+1]
	sli[i+1] = tmp
}
