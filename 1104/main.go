func pathInZigZagTree(label int) []int {
    nums := []int{}
    for label >=1 {
        nums = append(nums, label)
        if label % 2 == 1 {
            label--
        }
        label = label/2
    }
    result := []int{}
    for i:=len(nums)-1;i>=0;i--{
        result = append(result,nums[i])
    }
    var start,val int
    if len(result) %2 ==1 {
        start = 1
        val = 2
    } else {
        start = 2
        val = 4
    }
    for start < len(nums) {
        result[start] = val*2 - 1 - result[start] + val
        val = val*4
        start=start+2
    }
    return result
}
