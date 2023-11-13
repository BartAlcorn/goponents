package acquire

import "time"

type Acquire struct {
	T     string `json:"__t"`
	ID    string `json:"id"`
	V     int    `json:"__v"`
	Asset struct {
		ID          string `json:"id"`
		EventSource string `json:"eventSource"`
		Type        string `json:"type"`
		Status      struct {
			Codes []struct {
				Code   string `json:"code"`
				Detail string `json:"detail"`
				Data   string `json:"data"`
			} `json:"codes"`
			Type string `json:"type"`
		} `json:"status"`
		EventType string `json:"eventType"`
		ID0       struct {
			Oid string `json:"$oid"`
		} `json:"_id"`
		T     string `json:"__t"`
		V     int    `json:"__v"`
		Audit struct {
			CreatedAt   time.Time `json:"createdAt"`
			GeneratedAt time.Time `json:"generatedAt"`
			UpdatedAt   time.Time `json:"updatedAt"`
			UpdatedBy   string    `json:"updatedBy"`
		} `json:"audit"`
		CorrelationID string `json:"correlationId"`
		Details       struct {
			CorrelationMetadata []struct {
				ID          string `json:"id"`
				IDNamespace string `json:"idNamespace"`
			} `json:"correlationMetadata"`
			ManifestID        string `json:"manifestId"`
			ManifestNamespace string `json:"manifestNamespace"`
			MediaFilename     string `json:"mediaFilename"`
			MmcFilePath       string `json:"mmcFilePath"`
		} `json:"details"`
		EventMetaID  string `json:"eventMetaId"`
		IsOverall    bool   `json:"isOverall"`
		SeasonTitles []any  `json:"seasonTitles"`
		SeriesTitles []any  `json:"seriesTitles"`
		Titles       []any  `json:"titles"`
	} `json:"asset"`
	Audit struct {
		Asset struct {
			UpdatedAt time.Time `json:"updatedAt"`
		} `json:"asset"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
		UpdatedBy string    `json:"updatedBy"`
		Tasks     []struct {
			UpdatedAt   time.Time `json:"updatedAt"`
			ID          string    `json:"id"`
			Type        string    `json:"type"`
			EventSource string    `json:"eventSource"`
		} `json:"tasks"`
	} `json:"audit"`
	IsDeleted bool   `json:"isDeleted"`
	Type      string `json:"type"`
	Tasks     []struct {
		Type   string `json:"type"`
		Status struct {
			Codes []struct {
				Code   string `json:"code"`
				Detail string `json:"detail"`
				Data   string `json:"data"`
			} `json:"codes"`
			Type string `json:"type"`
		} `json:"status"`
		ID          string `json:"id"`
		EventSource string `json:"eventSource"`
		EventType   string `json:"eventType"`
		ID0         struct {
			Oid string `json:"$oid"`
		} `json:"_id"`
		T     string `json:"__t"`
		V     int    `json:"__v"`
		Audit struct {
			CreatedAt   time.Time `json:"createdAt"`
			GeneratedAt time.Time `json:"generatedAt"`
			UpdatedAt   time.Time `json:"updatedAt"`
			UpdatedBy   string    `json:"updatedBy"`
		} `json:"audit"`
		CorrelationID string `json:"correlationId"`
		Details       struct {
			CorrelationMetadata []struct {
				ID          string `json:"id"`
				IDNamespace string `json:"idNamespace"`
			} `json:"correlationMetadata"`
			ManifestID        string `json:"manifestId"`
			ManifestNamespace string `json:"manifestNamespace"`
			MediaFilename     string `json:"mediaFilename"`
			MmcFilePath       string `json:"mmcFilePath"`
		} `json:"details"`
		EventMetaID  string `json:"eventMetaId"`
		IsOverall    bool   `json:"isOverall"`
		SeasonTitles []any  `json:"seasonTitles"`
		SeriesTitles []any  `json:"seriesTitles"`
		Titles       []any  `json:"titles"`
	} `json:"tasks"`
	FootprintConfigID string `json:"footprintConfigId"`
	Milestones        []struct {
		Name   string `json:"name"`
		Type   string `json:"type"`
		Status string `json:"status"`
		Codes  []struct {
			Code   string `json:"code"`
			Detail string `json:"detail"`
			Data   string `json:"data"`
		} `json:"codes"`
		Audit struct {
			CreatedAt   time.Time `json:"createdAt"`
			GeneratedAt time.Time `json:"generatedAt"`
			UpdatedAt   time.Time `json:"updatedAt"`
			UpdatedBy   string    `json:"updatedBy"`
		} `json:"audit"`
	} `json:"milestones"`
	Status struct {
		PrepAssetUpload struct {
			Name   string `json:"name"`
			Type   string `json:"type"`
			Status string `json:"status"`
			Codes  []struct {
				Code   string `json:"code"`
				Detail string `json:"detail"`
				Data   string `json:"data"`
			} `json:"codes"`
			Audit struct {
				CreatedAt   time.Time `json:"createdAt"`
				GeneratedAt time.Time `json:"generatedAt"`
				UpdatedAt   time.Time `json:"updatedAt"`
				UpdatedBy   string    `json:"updatedBy"`
			} `json:"audit"`
		} `json:"Prep|Asset|upload"`
		PrepAssetAcquirePrep                        Event `json:"Prep|Asset|acquirePrep"`
		PrepAsset                                   Event `json:"Prep|Asset"`
		Prep                                        Event `json:"Prep"`
		TransformAssetAcquireTransform              Event `json:"Transform|Asset|acquireTransform"`
		TransformAsset                              Event `json:"Transform|Asset"`
		Transform                                   Event `json:"Transform"`
		ManifestationAssetManifestationRegistration Event `json:"Manifestation|Asset|manifestationRegistration"`
		ManifestationAsset                          Event `json:"Manifestation|Asset"`
		Manifestation                               Event `json:"Manifestation"`
		AcquireFootprint                            Event `json:"acquire-footprint"`
	} `json:"status"`
}

type Event struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Status string `json:"status"`
	Audit  any    `json:"audit"`
	Codes  any    `json:"codes"`
}
