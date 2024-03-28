package sorting

// BubbleSort -- sorts input arr in place using bubble sort
func BubbleSort(arr []int) {
	n := len(arr)
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
}

// ShellSort -- sorts input arr in place using Shell's sort
func ShellSort(v []int) {
	n := len(v)
	for gap := n / 2; gap > 0; gap /= 2 {
		for i := gap; i < n; i++ {
			for j := i - gap; j >= 0; j = j - gap {
				jq := j + gap
				if v[j] <= v[jq] {
					break
				}
				tmp := v[j]
				v[j] = v[jq]
				v[jq] = tmp
			}
		}
	}
}

func QuickSort(v []int) {
	rquick(v, 0, len(v)-1)
}

func rquick(v []int, lo int, hi int) {
	if lo >= hi {
		return
	}
	i := lo
	j := hi
	piv := v[j]
	for i < j {
		for i < j && v[i] <= piv {
			i++
		}
		for i < j && v[j] >= piv {
			j--
		}
		if i < j {
			tmp := v[i]
			v[i] = v[j]
			v[j] = tmp
		}
	}
	//set pivot into "center"
	tmp := v[i]
	v[i] = v[hi]
	v[hi] = tmp
	rquick(v, lo, i-1)
	rquick(v, i+1, hi)

}
func qsort(v []int, l int, r int) {
	if l >= r {
		return
	}

	k := partition(v, l, r)
	qsort(v, l, k-1)
	qsort(v, k+1, r)
}

func partition(v []int, l int, r int) int {
	i := l - 1
	for j := l; j < r; j++ {
		if v[j] > v[r] {
			// do nothing
		} else {
			i++
			tmp := v[i]
			v[i] = v[j]
			v[j] = tmp
		}
	}
	i += 1
	tmp := v[i]
	v[i] = v[r]
	v[r] = tmp
	return i
}
