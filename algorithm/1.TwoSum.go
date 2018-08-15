func twoSum(nums []int, target int) []int {
    var maps = make(map[int]int)
    for i,v :=range nums{        
        n:=target-v
        if j,ok := maps[n];ok{
            return []int{j,i}
        }else{
            maps[v] = i
        }
    }
    return nil
}