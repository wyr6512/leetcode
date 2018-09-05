var numDict = map[string]int{
    "I":1,
    "IV":4,
    "V":5,
    "IX":9,
    "X":10,
    "XL":40,
    "L":50,
    "XC":90,
    "C":100,
    "CD":400,
    "D":500,
    "CM":900,
    "M":1000,
}

func romanToInt(s string) int {
    sum:=0
    b := []byte(s)
    l :=len(b)
    if l<=1{
        return numDict[s]
    }
    for i:=0;i<l;{
    	if l-i == 1{
    		s = string(b[i:i+1])
    		sum+=numDict[s]
    		break
    	}
        s = string(b[i:i+2])
        if n,ok:=numDict[s];ok{
            sum+=n
            i+=2
        }else{
            s=string(b[i:i+1])
            sum+=numDict[s]
            i+=1
        }
    }
    return sum
}
