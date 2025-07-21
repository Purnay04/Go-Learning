package main

import "fmt"

func main()  {
    
    /*
        for(int num: nums){...} // num=nums[i]
    */

    nums:=[] int{1,2,3,4,5}
    fmt.Println(nums)

    for index, num := range nums {
        num *= 2
        fmt.Printf("%d, %d",index, num)
        fmt.Println()
    }
    fmt.Println(nums)

    // skip index and similarly can skip value as well keeping the index
    
    for _, num:= range nums {
        fmt.Printf("%d", num)
        fmt.Println()
    }

}

