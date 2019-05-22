/* Write a program to sort an array of integers.
 * The program should partition the array into 4 parts,
 * each of which is sorted by a different goroutine.
 * Each partition should be of approximately equal size.
 * Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

 * The program should prompt the user to input a series of integers.
 * Each goroutine which sorts ¼ of the array should print the subarray that it will sort.
 * When sorting is complete, the main goroutine should print the entire sorted list.
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	fmt.Println("Please type in sequence of integer (more than four) seperated by whitespace...")
	reader := bufio.NewReader(os.Stdin)
	ipt_raw, _ := reader.ReadString('\n')
	//fmt.Printf("Your raw input is %v\n", ipt_raw)
	ipt := strings.TrimSuffix(ipt_raw, "\n")
	//fmt.Printf("Your trimmed input is %v\n", ipt)
	sli := StringToSlice(ipt)
	s1, s2, s3, s4 := DivideToFourSlices(sli)
	wg.Add(4)
	fmt.Println("Start to sort...")
	go MySort(s1, &wg)
	go MySort(s2, &wg)
	go MySort(s3, &wg)
	go MySort(s4, &wg)
	wg.Wait()
	fmt.Println("Sort Done")
	//fmt.Println(sli)

	fmt.Println("Start to merge...")
	sa := Merge(s1, s2)
	sb := Merge(s3, s4)
	result := Merge(sa, sb)
	fmt.Println("Merge Done")
	fmt.Println(result)

}

func StringToSlice(s string) []int {
	str_sli := strings.Split(s, " ")
	int_sli := make([]int, 0)
	//fmt.Printf("Str sli is %v \n", str_sli)
	for _, e := range str_sli {
		//fmt.Printf("e is %v \n", e)
		num, err := strconv.Atoi(e)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Non-Integer chracter detected. Abort...")
			os.Exit(1)
		}
		int_sli = append(int_sli, num)

	}
	if len(int_sli) < 4 {
		fmt.Printf("Your input is %v \n", int_sli)
		fmt.Println("Less than four integers, Abort...")
		os.Exit(1)
	}
	return int_sli
}

func DivideToFourSlices(sli []int) ([]int, []int, []int, []int) {
	p := len(sli) / 4
	//fmt.Printf("p is %v \n", p)
	//fmt.Printf("sli is %v \n", sli)
	return sli[0:p], sli[p : 2*p], sli[2*p : 3*p], sli[3*p : len(sli)]
}

func MySort(s []int, wg *sync.WaitGroup) {
	sort.Ints(s)
	fmt.Println(s)
	wg.Done()
}

func Merge(si []int, sj []int) []int {
	sr := make([]int, 0)
	i := 0
	j := 0
	for {
		switch {
		case i < len(si) && j < len(sj):
			if si[i] < sj[j] {
				//fmt.Printf(" si.%v : %v < si.%v : %v\n", i, si[i], j, sj[j])
				sr = append(sr, si[i])
				i++
			} else {
				sr = append(sr, sj[j])
				//fmt.Printf(" si.%v : %v >= si.%v : %v\n", i, si[i], j, sj[j])
				j++
			}
		case j == len(sj) && i < len(si):
			for r := 0; r < len(si)-i; r++ {
				sr = append(sr, si[i])
				i++
			}

		case i == len(si) && j < len(sj):
			for r := 0; r < len(sj)-j; r++ {
				sr = append(sr, sj[j])
				j++
			}

		case i == len(si) && j == len(sj):
			return sr

		default:
			fmt.Println("This should never shown")
			os.Exit(1)
		}
	}
}
