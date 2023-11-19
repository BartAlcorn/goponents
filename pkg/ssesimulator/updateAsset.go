package ssesimulator

import "fmt"

func UpdateAsset() {
	l := len(Assets)

	if l > 0 {
		fmt.Printf("found %v", l)
	}
}
