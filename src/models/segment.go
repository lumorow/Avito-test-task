package models

type Segment struct {
	Id          int
	SegmentName string `json:"segment_name"`
}

type Segments struct {
	SegmentsName []string `json:"segments_name"`
}
