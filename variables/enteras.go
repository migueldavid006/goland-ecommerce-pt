package variables

import "fmt"

func MuestroEnteros() {
	var intcomun int
	intde32 := int32(10)
	intde64 := int64(100)
	
	fmt.Println("intcomun = ", intcomun)
	fmt.Println("intcomun32 = ", intde32)
	fmt.Println("intcomun64 = ", intde64)
}