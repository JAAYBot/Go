

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	lenght := len(nums)
	finalArray := [][]int{}

	for i := 0; i < lenght-2 && nums[i]<=0; i++ {
			if  i>0 && nums[i] == nums[i-1]{ continue }
			for j := i+1;  (j < lenght-1) && nums[i]+nums[j]<=0; j++{
				if j-i>1 && nums[j] == nums[j-1] { continue }
					k := sort.SearchInts(nums[j+1:], (-nums[i]-nums[j]))
					k = k+j+1
				if k != lenght && k > j && nums[k] == (-nums[i]-nums[j]){ 
					temp :=[]int{nums[i], nums[j], nums[k]}
					finalArray = append(finalArray, temp)
				}
			}
	}
	return finalArray
}