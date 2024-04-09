import(
    "fmt"
    "math"
)

func containsNearbyDuplicate(nums []int, k int) bool {
   /* for i := 0; i < len(nums); i++ {
        for j := 0; j < i; j++ {
            if((nums[i] == nums[j]) && (math.Abs(float64(i) - float64(j)) <= float64(k))){
                return true;
            }
        }
    }
    return false;
*/
    if len(nums) < 2 {
        return false
    }

    mymap := make(map[int]int)
    repeat := 0 
    for _, number := range nums {
        mymap[number]++
        if mymap[number] == 2{
            repeat = number
            break;
        }
    }

    var output[]int

    for l := 0; l < len(nums); l++{
       if nums[l] == repeat{
            output = append(output, l)
        }
    }

    if len(output) < 2 {
        return false
    }

    temp := output[0]
    
    for i := 1; i < len(output); i++{
        if math.Abs(float64(output[i]) - float64(temp)) <= float64(k) {
            return true
        }
            temp = output[i]
    }
    return false
}
