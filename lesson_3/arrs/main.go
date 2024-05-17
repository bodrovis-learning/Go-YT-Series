package main

func main() {
	var a [3]int

	a[0] = 1
	a[1] = 2
	a[2] = 3

	modifyMe(&a)

	for index, el := range a {
		println(index, el)
	}

	q := [...]int{1, 2, 3}
	println(q[0])

	r := [...]rune{9: 'a'}
	println(r[9], r[8])

	sl := r[3:10]
	println(sl[6], len(sl))

	sl = append(sl, 'b')
	println(sl[7], len(sl))
}

func modifyMe(arr *[3]int) {
	arr[1] = 100
}
