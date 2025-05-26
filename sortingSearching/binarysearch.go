package sortingsearching

func BinarySearch(arr []int, target int) int {
	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := start + (end-start)/2

		if target == arr[mid] {
			return mid
		}

		if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return -1
}
