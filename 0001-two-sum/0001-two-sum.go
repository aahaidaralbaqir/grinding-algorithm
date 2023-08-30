func twoSum(nums []int, target int) []int {
    for x, v := range nums {
        fmt.Println(v)
        for z, y := range nums {
            if  x != z && v + y == target {
                return []int{x, z}
            }
        }
    }
    return []int{}
}