package ssesimulator

import (
	"github.com/labstack/gommon/color"
)

func generateAssets(addCh chan<- Asset) {
	count := 8
	color.Println(color.Green("generateAssets: started"))

	if len(Assets) > count {
		color.Println(color.Yellow("generateAssets: resending"))
		for _, a := range Assets {
			addCh <- a
		}
		return
	}

	for i := 0; i < count; i++ {
		a := MakeAsset()
		Assets = append(Assets, a)
		addCh <- a
	}

	color.Println(color.Green("generateAssets: Completed"))
}
