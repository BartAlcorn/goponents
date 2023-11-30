package minsse

// sprig "github.com/Masterminds/sprig/v3"

var addCh chan Asset
var updateCh chan Asset
var Assets []Asset
var monitorID string
var Stats map[string]StatBlock

var tasks = []string{"Transcoding", "Captioning", "Audio QC", "Video QC", "Dubbing", "Subbing"}

func init() {
	Assets = make([]Asset, 0)
	addCh = make(chan Asset)
	updateCh = make(chan Asset)
	Stats = make(map[string]StatBlock, 0)
	Stats["counts"] = StatBlock{Name: "counts"}
	Stats["metrics"] = StatBlock{Name: "metrics"}
}
