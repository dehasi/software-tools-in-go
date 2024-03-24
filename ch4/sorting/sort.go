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
	qsort(v, 0, len(v)-1)
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
	pi := (l + r) / 2 // L + (R-L)/2
	p := v[pi]
	for l < r {
		for (v[l] < p) && (l < r) {
			l++
		}
		for p < v[r] && (l < r) {
			r--
		}

		if (v[l] > v[r]) && (l < r) {
			tmp := v[l]
			v[l] = v[r]
			v[r] = tmp
			l += 1
			r -= 1
		}
	}
	return l
}
