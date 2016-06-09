package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	cases := []int{2, 3, 4}
	sel := []int{1, 2}
	ids := []int{
		131, 213, 4323, 521, 234, 523, 52343, 3521, 4324}

	answers := []string{
		"certain",
		"decidedly",
		"Without",
		"definitely",
		"it",
		"yes",
		"likely",
		"Ogood",
		"Yes"}
	questions := []string{"oyes",
		"Reply",
		"Ask",
		"Better",
		"Cannot",
		"Concentrate",
		"count",
		"replyo",
		"sources",
		"Outlook",
		"doubtful",
	}
	take := sel[rand.Intn(len(sel))]
	ncase := cases[rand.Intn(len(cases))]

	m, mkeys := initialise(take, ids, questions, answers)
	fmt.Println(m)
	n, nkeys := initialise(take, ids, questions, answers)
	fmt.Println(n)

	result, rkeys := intersection(m, n, mkeys, nkeys)

	//********for more than 2 sets**********
	if ncase > 2 {
		for i := 2; i < ncase; i++ {
			next, nextkeys := initialise(take, ids, questions, answers)
			fmt.Println(next)
			result, rkeys = intersection(next, result, nextkeys, rkeys)
		}
	}
	fmt.Println("\n\n########### INTERSECTION OF RANDOMLY GENERATED SETS ###########\n\n")
	fmt.Println(result)
}

//***********function to randomly generate map,it returns map and array(keys) of generated map******
func initialise(take int, ids []int, questions []string, answers []string) (map[int][]string, []int) {
	var tmap = make(map[int][]string)

	for i := 0; i < 3; i++ {
		tid := ids[rand.Intn(len(ids))]
		tquestion := questions[rand.Intn(len(questions))]
		tanswer := answers[rand.Intn(len(answers))]
		var s = make([]string, take)

		if take == 1 {
			s[0] = tquestion
		}
		if take == 2 {
			s[0] = tquestion
			s[1] = tanswer
		}

		tmap[tid] = s
	}
	var keys = make([]int, 0, len(tmap))
	for k := range tmap {
		keys = append(keys, k)
	}
	return tmap, keys
}

//*********function to generate to intersection of two sets*************
func intersection(m map[int][]string, n map[int][]string, mkeys []int, nkeys []int) (map[int][]string, []int) {
	var tmap = make(map[int][]string)
	k := 0
	for i := 0; i < len(mkeys); i++ {
		for j := 0; j < len(nkeys); j++ {
			if mkeys[i] == nkeys[j] {
				tmap[mkeys[i]] = m[mkeys[i]]
				k++

			}
		}
	}
	var keys = make([]int, 0, len(tmap))
	for k := range tmap {
		keys = append(keys, k)
	}
	return tmap, keys
}
