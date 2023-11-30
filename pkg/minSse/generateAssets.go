package minsse

func generateAssets(count int) {
	for i := 0; i < count; i++ {
		a := MakeAsset()
		Assets = append(Assets, a)
		addCh <- a
	}
}
