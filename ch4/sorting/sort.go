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
