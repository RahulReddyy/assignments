// slice = {1,2,3,3,4,5}
package Problem2

func Duplicates(slice []int) map[int]string {

	var mapData = make(map[int]string)
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] == slice[j] {
				mapData[slice[i]] = ""
			}
		}
	}
	return mapData
}
