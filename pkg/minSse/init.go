package minSse

// sprig "github.com/Masterminds/sprig/v3"

var Assets []Asset
var monitorID string

var tasks = []string{"Transcoding", "Captioning", "Audio QC", "Video QC", "Dubbing", "Subbing"}

func init() {
	Assets = make([]Asset, 0)
}
