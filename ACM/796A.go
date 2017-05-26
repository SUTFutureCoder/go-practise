package main

import (
	"fmt"
)

func main(){
	var n, tmpM, k int
//	n := firstLineSlice[0] //总共的数量
	fmt.Scanf("%d %d %d\n", &n, &tmpM, &k)
	m := tmpM - 1

	secondLineSlice := make(map[int]int)
	for i := 0; i < n; i++ {
		var temp int
		fmt.Scanf("%d", &temp)
		secondLineSlice[i] =  temp
	}

	//开始处理
	distance := 0
	for {
		if (m + distance >= 0 && secondLineSlice[m + distance] != 0){
			//进行判断
			currentHousePrice := secondLineSlice[m + distance]
			if (currentHousePrice <= k){
				break;
			}
		}
		if (m - distance >= 0 && secondLineSlice[m - distance] != 0){
			//再次进行判断
			currentHousePrice := secondLineSlice[m - distance]
			if (currentHousePrice <= k){
				break;
			}
		}
		distance++
	}
	fmt.Println(distance * 10)
	return ;
}

