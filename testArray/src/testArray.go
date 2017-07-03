package main

import "fmt"

func main() {
	a := []int{0,1,2,3,4,5,6,7,8,9}
	s := a[0:5]

	Print(a, "a")
	Print(s, "s")

	s1 := append(s, 0)
	fmt.Printf("=================\n")

	Print(a, "a")
	Print(s, "s")
	Print(s1, "s1")

	fmt.Printf("=================\n")

	var ages [4]int = [4]int{1, 2, 3, 4}
	var arr = ages

	var aa1 = arr[0:2]
	var aa2 = arr[2:]

	ss2 := append(aa2,0)
	ss1 := append(aa1,0)
	var ss3 = ss1;
	ss3 = append(ss1,0)

	fmt.Println(ages)
	fmt.Println(arr)
	//Print(ages, "ages")
	//Print(arr, "arr")
	Print(ss1, "ss1")
	Print(ss2, "ss2")
	Print(ss3, "ss3")

}

func Print(o []int, name string) {
	fmt.Printf("%s: ", name)

	for _, v := range o {
		fmt.Printf("%d ", v)
	}

	fmt.Printf("\n")
}