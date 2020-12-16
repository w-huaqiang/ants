package main

import (
	"fmt"

	"bjzdgt.com/ants/utils"
)

func main() {
	r1 := "845@qq.com"
	r2 := "huaqiang@163.com"
	r3 := "12gg@ggcom"

	fmt.Println(utils.EmailFormatCheck(r1))
	fmt.Println(utils.EmailFormatCheck(r2))
	fmt.Println(utils.EmailFormatCheck(r3))
}
