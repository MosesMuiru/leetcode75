package main

func Inter(list1, list2 []int) []int {
	in := []int{}

	for i := range list1 {
		for j := range list2 {
			if list1[i] == list2[j] {
				in = append(in, list1[i])
				continue
			}

		}
	}

	for i := range list2 {
		for j := range list1 {
			if list2[i] == list1[j] {
				in = append(in, list1[j])
				continue
			}

		}
	}

	return in

}
