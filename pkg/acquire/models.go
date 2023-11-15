package acquire

import "time"

type Acquire struct {
	ID                string       `json:"id"`
	T                 string       `json:"__t,omitempty"`
	V                 int          `json:"__v,omitempty"`
	Asset             Asset        `json:"asset,omitempty"`
	Audit             Audit        `json:"audit,omitempty"`
	IsDeleted         bool         `json:"isDeleted,omitempty"`
	Type              string       `json:"type,omitempty"`
	Tasks             []Task       `json:"tasks,omitempty"`
	Title             string       `json:"title,omitempty"`
	FootprintConfigID string       `json:"footprintConfigId,omitempty"`
	Milestones        []Milestones `json:"milestones,omitempty"`
	Status            Status       `json:"status,omitempty"`
}

type Codes struct {
	Code   string `json:"code,omitempty"`
	Detail string `json:"detail,omitempty"`
	Data   any    `json:"data,omitempty"`
}

type NamedID struct {
	ID        string `json:"id,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

type Title struct {
	Full     string `json:"full,omitempty"`
	Language string `json:"language,omitempty"`
	Type     string `json:"type,omitempty"`
}
type Details struct {
	ContentType                 string    `json:"contentType,omitempty"`
	ControllingNetwork          string    `json:"controllingNetwork,omitempty"`
	CorrelationMetadata         []NamedID `json:"correlationMetadata,omitempty"`
	Distributor                 string    `json:"distributor,omitempty"`
	EditScrID                   NamedID   `json:"editScrId,omitempty"`
	EditUUID                    NamedID   `json:"editUuid,omitempty"`
	EntityType                  string    `json:"entityType,omitempty"`
	EpisodeNumberInSeason       string    `json:"episodeNumberInSeason,omitempty"`
	MediaID                     NamedID   `json:"mediaId,omitempty"`
	PaidID                      NamedID   `json:"paidId,omitempty"`
	PropertyID                  NamedID   `json:"propertyId,omitempty"`
	ScrID                       NamedID   `json:"scrId,omitempty"`
	SeasonID                    NamedID   `json:"seasonId,omitempty"`
	SeasonNumber                string    `json:"seasonNumber,omitempty"`
	SeasonTitles                []Title   `json:"seasonTitles,omitempty"`
	TargetManifestationLocation string    `json:"targetManifestationLocation,omitempty"`
	SeriesTitles                []Title   `json:"seriesTitles,omitempty"`
	Titles                      []Title   `json:"titles,omitempty"`
	ManifestID                  string    `json:"ManifestID,omitempty"`
	WorkflowPriority            string    `json:"workflowPriority,omitempty"`
}
type Asset struct {
	Type                  string      `json:"type,omitempty"`
	Status                TasksStatus `json:"status,omitempty"`
	ID                    string      `json:"id,omitempty"`
	EventSource           string      `json:"eventSource,omitempty"`
	EventType             string      `json:"eventType,omitempty"`
	T                     string      `json:"__t,omitempty"`
	V                     int         `json:"__v,omitempty"`
	Audit                 Audit       `json:"audit,omitempty"`
	CorrelationID         string      `json:"correlationId,omitempty"`
	Details               Details     `json:"details,omitempty"`
	EventMetaID           string      `json:"eventMetaId,omitempty"`
	IsOverall             bool        `json:"isOverall,omitempty"`
	PaidID                NamedID     `json:"paidId,omitempty"`
	EditScrID             NamedID     `json:"editScrId,omitempty"`
	ScrID                 NamedID     `json:"scrId,omitempty"`
	PropertyID            NamedID     `json:"propertyId,omitempty"`
	EntityType            string      `json:"entityType,omitempty"`
	SeasonID              NamedID     `json:"seasonId,omitempty"`
	NamedID               NamedID     `json:"NamedID,omitempty"`
	SeasonNumber          string      `json:"seasonNumber,omitempty"`
	ControllingNetwork    string      `json:"controllingNetwork,omitempty"`
	EpisodeNumberInSeason string      `json:"episodeNumberInSeason,omitempty"`
	Titles                []Title     `json:"titles,omitempty"`
	SeriesTitles          []Title     `json:"seriesTitles,omitempty"`
	SeasonTitles          []Title     `json:"seasonTitles,omitempty"`
	EditUUID              NamedID     `json:"editUuid,omitempty"`
	Distributor           string      `json:"distributor,omitempty"`
	ContentType           string      `json:"contentType,omitempty"`
	MediaID               NamedID     `json:"mediaId,omitempty"`
}
type AuditTasks struct {
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
	ID          string    `json:"id,omitempty"`
	Type        string    `json:"type,omitempty"`
	EventSource string    `json:"eventSource,omitempty"`
}
type Audit struct {
	CreatedAt   time.Time    `json:"createdAt,omitempty"`
	GeneratedAt time.Time    `json:"generatedAt,omitempty"`
	GeneratedBy string       `json:"generatedBy,omitempty"`
	UpdatedAt   time.Time    `json:"updatedAt,omitempty"`
	UpdatedBy   string       `json:"updatedBy,omitempty"`
	Tasks       []AuditTasks `json:"tasks,omitempty"`
}

type TasksStatus struct {
	Codes []Codes `json:"codes,omitempty"`
	Type  string  `json:"type,omitempty"`
}

type Task struct {
	ID            string      `json:"id,omitempty"`
	T             string      `json:"__t,omitempty"`
	Type          string      `json:"type,omitempty"`
	V             int         `json:"__v,omitempty"`
	Audit         Audit       `json:"audit,omitempty"`
	CorrelationID string      `json:"correlationId,omitempty"`
	Details       Details     `json:"details,omitempty"`
	EventMetaID   string      `json:"eventMetaId,omitempty"`
	EventSource   string      `json:"eventSource,omitempty"`
	EventType     string      `json:"eventType,omitempty"`
	IsOverall     bool        `json:"isOverall,omitempty"`
	Status        TasksStatus `json:"status,omitempty"`
}

type Milestones struct {
	Name   string  `json:"name,omitempty"`
	Type   string  `json:"type,omitempty"`
	Status string  `json:"status,omitempty"`
	Codes  []Codes `json:"codes,omitempty"`
	Audit  Audit   `json:"audit,omitempty"`
}
type PrepAssetAcquirePrep struct {
	Name   string  `json:"name,omitempty"`
	Type   string  `json:"type,omitempty"`
	Status string  `json:"status,omitempty"`
	Codes  []Codes `json:"codes,omitempty"`
	Audit  Audit   `json:"audit,omitempty"`
}

type PrepAsset struct {
	Name   string  `json:"name,omitempty"`
	Type   string  `json:"type,omitempty"`
	Status string  `json:"status,omitempty"`
	Audit  Audit   `json:"audit,omitempty"`
	Codes  []Codes `json:"codes,omitempty"`
}
type Prep struct {
	Name   string  `json:"name,omitempty"`
	Type   string  `json:"type,omitempty"`
	Status string  `json:"status,omitempty"`
	Audit  Audit   `json:"audit,omitempty"`
	Codes  []Codes `json:"codes,omitempty"`
}

type TransformAssetAcquireTransform struct {
	Name   string  `json:"name,omitempty"`
	Type   string  `json:"type,omitempty"`
	Status string  `json:"status,omitempty"`
	Codes  []Codes `json:"codes,omitempty"`
	Audit  Audit   `json:"audit,omitempty"`
}

type TransformAsset struct {
	Name   string  `json:"name,omitempty"`
	Type   string  `json:"type,omitempty"`
	Status string  `json:"status,omitempty"`
	Audit  Audit   `json:"audit,omitempty"`
	Codes  []Codes `json:"codes,omitempty"`
}
type Transform struct {
	Name   string  `json:"name,omitempty"`
	Type   string  `json:"type,omitempty"`
	Status string  `json:"status,omitempty"`
	Audit  Audit   `json:"audit,omitempty"`
	Codes  []Codes `json:"codes,omitempty"`
}
type ManifestationAsset struct {
	Name   string `json:"name,omitempty"`
	Type   string `json:"type,omitempty"`
	Status string `json:"status,omitempty"`
	Audit  any    `json:"audit,omitempty"`
	Codes  any    `json:"codes,omitempty"`
}
type Manifestation struct {
	Name   string `json:"name,omitempty"`
	Type   string `json:"type,omitempty"`
	Status string `json:"status,omitempty"`
	Audit  any    `json:"audit,omitempty"`
	Codes  any    `json:"codes,omitempty"`
}
type AcquireFootprint struct {
	Name   string  `json:"name,omitempty"`
	Type   string  `json:"type,omitempty"`
	Status string  `json:"status,omitempty"`
	Audit  Audit   `json:"audit,omitempty"`
	Codes  []Codes `json:"codes,omitempty"`
}
type Status struct {
	PrepAssetAcquirePrep           PrepAssetAcquirePrep           `json:"Prep|Asset|acquirePrep,omitempty"`
	PrepAsset                      PrepAsset                      `json:"Prep|Asset,omitempty"`
	Prep                           Prep                           `json:"Prep,omitempty"`
	TransformAssetAcquireTransform TransformAssetAcquireTransform `json:"Transform|Asset|acquireTransform,omitempty"`
	TransformAsset                 TransformAsset                 `json:"Transform|Asset,omitempty"`
	Transform                      Transform                      `json:"Transform,omitempty"`
	ManifestationAsset             ManifestationAsset             `json:"Manifestation|Asset,omitempty"`
	Manifestation                  Manifestation                  `json:"Manifestation,omitempty"`
	AcquireFootprint               AcquireFootprint               `json:"acquire-footprint,omitempty"`
}
