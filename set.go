package main

import (
	"fmt"
	"math/rand"
	"time"
)

type student struct {
	id   int
	name string
}

type person struct {
	id   int
	name string
}

var n [3]student
var m [2]person
var r [5]person

func main() {

	rand.Seed(time.Now().Unix()) // Try changing this number!
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
		"Yes",
		" oyes",
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

	for i := 0; i < 3; i++ {
		n[i].id = ids[rand.Intn(len(ids))]
		n[i].name = answers[rand.Intn(len(answers))]

	}
	fmt.Println(n)
	for i := 0; i < 2; i++ {
		m[i].id = ids[rand.Intn(len(ids))]
		m[i].name = answers[rand.Intn(len(answers))]

	}
	fmt.Println(m)
	//************intersection********

	k := 0
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(n); j++ {
			if m[i].id == n[j].id {
				r[k].id = n[j].id
				r[k].name = n[j].name
				k += 1
			}

		}
	}
	// k = 0
	// while r[k].id!=0 {
	// 	fmt.Println(r[k].id)
	// 	fmt.Println(r[k].name)
	// }

	fmt.Println("Set_Intersection\n ", r)

	//*************union*********
	for j := 0; j < len(n); j++ {

		r[j].id = n[j].id
		r[j].name = n[j].name
	}

	k = len(n)
	var j int
	for i := 0; i < len(m); i++ {
		for j = 0; j < len(n); j++ {
			if m[i].id == r[j].id {
				break
			}
		}

		if j == len(n) {
			r[k].id = m[i].id
			r[k].name = m[i].name
			k += 1
		}
	}

	fmt.Println("Union")
	fmt.Println(r)
}
