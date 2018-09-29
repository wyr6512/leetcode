func lengthOfLongestSubstring(s string) int {
    dict := map[int32]int{}
    ans:=0
    i:=0
    for j,val :=range s{
        if p,ok:=dict[val];ok{
            i = max(p,i)
        }
        ans = max(ans,j-i+1)
        dict[val] = j+1
    }
    return ans
}

func max(a int,b int) int{
    if a>=b {
        return a
    }
    return b
}
