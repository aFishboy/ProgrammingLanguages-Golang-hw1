package min

// Min returns the minimum value in the arr,
// and 0 if arr is nil.
func Min(arr []int) int {
	/* check if the array is empty and if so return 0 */
	if  arrLen := len(arr); arrLen == 0 {
		return 0
	} else {
		/* array is not 0 */
		var min int = arr[0]; // make the first elem min
		for  index := 1; index < arrLen; index++ { // loop through each elem
			/* if current elem is smaller update min */
			if arr[index] < min {
				min = arr[index]
			}
		}
		return min
	}
}
