package minsse

// generateAssets creates n number of Assets
// and pushes them onto the addCh channel
func generateAssets(count int) {
	for i := 0; i < count; i++ {
		a := MakeAsset()
		Assets = append(Assets, a)
		addCh <- a
	}
}
