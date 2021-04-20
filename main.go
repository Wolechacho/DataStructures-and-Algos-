package main

import (
	"DataStructures-and-Algos/DS"
	"fmt"
)

func main() {
	// result := DS.FindMissingRange([]int{0, 1, 3, 50, 75}, 0, 99)
	// fmt.Println(result)

	// s := DS.ShortestDistance([]string{"practice", "makes", "perfect", "coding", "makes"}, "makes", "coding")
	// fmt.Println(s)

	//f := DS.IsStrobogrammatic("688889")
	//fmt.Println(f)

	// v := DS.PermutatePalindrome("carerac")
	// fmt.Println(v)

	// h := DS.ValidWordAbbreviation("word", "1o2")
	// fmt.Println(h)

	// j1 := []string{"great", "fine"}
	// j2 := []string{"acting", "drama"}
	// j3 := []string{"skills", "talent"}
	// p := [][]string{j1, j2, j3}

	// booly := DS.Sentencesimilarity([]string{"great", "acting", "skills"}, []string{"fine", "drama", "talent"}, p)
	// fmt.Println(booly)

	// arr := DS.ValidAnagramMapping([]int{12, 28, 46, 32, 50}, []int{50, 12, 32, 46, 28})
	// fmt.Println(arr)

	// c := DS.ConfusingNumber(0)
	// fmt.Println(c)

	// indexpairs := DS.IndexPairs("thestoryofleetcodeandme", []string{"story", "fleet", "leetcode"})
	// fmt.Println(indexpairs)

	// x := DS.SumOfDigits([]int{99, 77, 33, 66, 55})
	// fmt.Println(x)

	// m := []int{1, 91}
	// n := []int{1, 92}
	// b := []int{2, 93}
	// c := []int{2, 97}
	// d := []int{1, 60}
	// g := []int{2, 77}
	// s := []int{1, 65}
	// u := []int{1, 87}
	// r := []int{1, 100}
	// t := []int{2, 100}
	// y := []int{2, 76}

	// items := [][]int{m, n, b, c, d, g, s, u, r, t, y}
	// fv := DS.HighFive(items)
	//fmt.Println(fv)

	// v := DS.TwosumlessathanK([]int{34, 23, 1, 24, 75, 33, 54, 8}, 60)
	// fmt.Println(v)

	// bv := DS.CountSubString("aaaaaaaaaa")
	// fmt.Println(bv)

	// ms := DS.MissingNumberProgression([]int{2, 4, 8, 10, 12})
	// fmt.Println(ms)

	// lk := DS.LargestSubArrayK([]int{1, 4, 5, 2, 3}, 4)
	// fmt.Println(lk)

	// a := []int{0, 1}
	// b := []int{1, 2}
	// ss := DS.ShiftString("abc", [][]int{a, b})
	// fmt.Println(ss)

	// arr := DS.ArrayTransformation([]int{6,2,3,4})
	// fmt.Println(arr)

	arr := DS.CountingElement([]int{1,1,2})
	fmt.Println(arr)

}

func printSlices() []string {
	result := []string{}

	a := 1
	b := 2

	fmt.Println(a, b)

	setSlices(result)
	return result
}

func setSlices(res []string) {
	//res[0] = "tunji"
	res = append(res, "sade")
	fmt.Println("from res", res)
}
