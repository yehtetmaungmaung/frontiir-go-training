package main

import "fmt"

func main() {

	var data = map[string]map[string]int{
		"myanmar": {
			"100": 63,
			"101": 55,
			"102": 42,
			"103": 80,
			"104": 67,
			"105": 54,
			"106": 88,
		},
		"english": {
			"100": 92,
			"101": 68,
			"102": 75,
			"103": 88,
			"104": 55,
			"105": 72,
			"106": 80,
		},
		"mathematics": {
			"100": 63,
			"101": 77,
			"102": 54,
			"103": 89,
			"104": 70,
			"105": 42,
			"106": 95,
		},
		"physics": {
			"100": 81,
			"101": 58,
			"102": 67,
			"103": 73,
			"104": 90,
			"105": 64,
			"106": 87,
		},
		"chemistry": {
			"100": 46,
			"101": 83,
			"102": 71,
			"103": 59,
			"104": 78,
			"105": 93,
			"106": 66,
		},
		"biology": {
			"100": 84,
			"101": 57,
			"102": 74,
			"103": 65,
			"104": 82,
			"105": 50,
			"106": 96,
		},
	}

	answers := make(map[string]map[string]int)

	for subject, sid := range data {
		for sid, mark := range sid {
			if _, ok := answers[sid]; !ok {
				answers[sid] = map[string]int{}
			}
			answers[sid][subject] = mark
		}
	}

	for k, v := range answers {
		fmt.Println(k, v)
	}
}
