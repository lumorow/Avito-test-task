package models

type UserSegments struct {
	UserUID   int `json:"user_id"`
	SegmentId int `json:"segment_id"`
}

type UserSegmentsResponse struct {
	UserID   int      `json:"user_id"`
	Segments []string `json:"segments"`
}
