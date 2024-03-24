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
	gap := n / 2
	for gap > 0 {
		for i := gap + 1; i < n; i++ {
			j := i - gap
			for j >= 0 {
				jq := j + gap
				if v[j] <= v[jq] {
					j = -1 // why not break?
				} else {
					k := v[j]
					v[j] = v[jq]
					v[jq] = k
				}
				j = j - gap
			}
		}
		gap /= 2
	}
}
