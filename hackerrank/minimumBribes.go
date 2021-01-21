// Complete the minimumBribes function below.
package hackerrank

import "fmt"

// Complete the minimumBribes function below.
func minimumBribes(q []int32) {

	minNumOfBribes := 0
	var index int32
	var unsortedQueueLength int32 = int32(len(q))
	willNeedAnotherGo := true

	for index = 0; index < unsortedQueueLength; index++ {
		newValue := q[index]
		oldIndex := newValue - 1

		if index < oldIndex-2 {
			fmt.Println("Too chaotic")
			return
		}
	}

	for willNeedAnotherGo {
		willNeedAnotherGo = false
		for index = 0; index < unsortedQueueLength-1; index++ {
			if q[index] > q[index+1] {
				tmp := q[index]
				q[index] = q[index+1]
				q[index+1] = tmp
				willNeedAnotherGo = true
				minNumOfBribes++
			}
		}
		unsortedQueueLength--
	}

	fmt.Println(minNumOfBribes)
}
