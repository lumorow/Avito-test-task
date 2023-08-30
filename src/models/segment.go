package models

type SegmentRequest struct {
	SegmentName string `json:"segment_name"`
}

type SegmentResponse struct {
	Id          int
	SegmentName string `json:"segment_name"`
}

type Segments struct {
	SegmentsName []string `json:"segments_name"`
}
