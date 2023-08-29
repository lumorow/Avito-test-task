package tests

import (
	"Avito-test-task/models"
	"github.com/go-resty/resty/v2"
	"testing"
)

func Test_POST_AddSegment_StatusCodeShouldEqual200(t *testing.T) {
	client := resty.New()

	segment := models.Segment{
		SegmentName: "AVITO_SUPER_SALE_50",
	}
	resp, _ := client.R().SetBody(&segment).Post("http://localhost:8080/segment")
	if resp.StatusCode() != 200 {
		t.Errorf("Unexpected status code, expected %d, got %d instead", 200, resp.StatusCode())
	}

	segment = models.Segment{
		SegmentName: "AVITO_WOW",
	}
	resp, _ = client.R().SetBody(&segment).Post("http://localhost:8080/segment")
	if resp.StatusCode() != 200 {
		t.Errorf("Unexpected status code, expected %d, got %d instead", 200, resp.StatusCode())
	}

	segment = models.Segment{
		SegmentName: "AVITO_CLASSIC_15",
	}
	resp, _ = client.R().SetBody(&segment).Post("http://localhost:8080/segment")
	if resp.StatusCode() != 200 {
		t.Errorf("Unexpected status code, expected %d, got %d instead", 200, resp.StatusCode())
	}

	segment = models.Segment{
		SegmentName: "AVITO_AWESOME_90",
	}
	resp, _ = client.R().SetBody(&segment).Post("http://localhost:8080/segment")
	if resp.StatusCode() != 200 {
		t.Errorf("Unexpected status code, expected %d, got %d instead", 200, resp.StatusCode())
	}

	segment = models.Segment{
		SegmentName: "AVITO_HOME",
	}
	resp, _ = client.R().SetBody(&segment).Post("http://localhost:8080/segment")
	if resp.StatusCode() != 200 {
		t.Errorf("Unexpected status code, expected %d, got %d instead", 200, resp.StatusCode())
	}

	segment = models.Segment{
		SegmentName: "AVITO_WORK",
	}
	resp, _ = client.R().SetBody(&segment).Post("http://localhost:8080/segment")
	if resp.StatusCode() != 200 {
		t.Errorf("Unexpected status code, expected %d, got %d instead", 200, resp.StatusCode())
	}
}

func Test_POST_AddSegment_StatusCodeShouldEqual500(t *testing.T) {
	client := resty.New()

	segment := models.Segment{
		SegmentName: "AVITO_SUPER_SALE_50",
	}
	resp, _ := client.R().SetBody(&segment).Post("http://localhost:8080/segment")
	if resp.StatusCode() != 500 {
		t.Errorf("Unexpected status code, expected %d, got %d instead", 500, resp.StatusCode())
	}

	segment = models.Segment{
		SegmentName: "AVITO_WOW",
	}
	resp, _ = client.R().SetBody(&segment).Post("http://localhost:8080/segment")
	if resp.StatusCode() != 500 {
		t.Errorf("Unexpected status code, expected %d, got %d instead", 500, resp.StatusCode())
	}

	segment = models.Segment{
		SegmentName: "AVITO_CLASSIC_15",
	}
	resp, _ = client.R().SetBody(&segment).Post("http://localhost:8080/segment")
	if resp.StatusCode() != 500 {
		t.Errorf("Unexpected status code, expected %d, got %d instead", 500, resp.StatusCode())
	}

	segment = models.Segment{
		SegmentName: "AVITO_AWESOME_90",
	}
	resp, _ = client.R().SetBody(&segment).Post("http://localhost:8080/segment")
	if resp.StatusCode() != 500 {
		t.Errorf("Unexpected status code, expected %d, got %d instead", 500, resp.StatusCode())
	}

	segment = models.Segment{
		SegmentName: "AVITO_HOME",
	}
	resp, _ = client.R().SetBody(&segment).Post("http://localhost:8080/segment")
	if resp.StatusCode() != 500 {
		t.Errorf("Unexpected status code, expected %d, got %d instead", 500, resp.StatusCode())
	}

	segment = models.Segment{
		SegmentName: "AVITO_WORK",
	}
	resp, _ = client.R().SetBody(&segment).Post("http://localhost:8080/segment")
	if resp.StatusCode() != 500 {
		t.Errorf("Unexpected status code, expected %d, got %d instead", 500, resp.StatusCode())
	}
}

func Test_PUT_AddSegmentForUser2000(t *testing.T) {
	client := resty.New()

	var response models.UserSegmentsResponse

	segmentsMap := map[string]bool{
		"AVITO_SUPER_SALE_50": true,
		"AVITO_HOME":          true,
		"AVITO_AWESOME_90":    true,
	}

	segments := models.Segments{
		SegmentsName: []string{"AVITO_SUPER_SALE_50", "AVITO_HOME", "AVITO_AWESOME_90"},
	}
	resp, _ := client.R().SetBody(&segments).SetResult(&response).Put("http://localhost:8080/user/2000/segments")

	if resp.StatusCode() != 200 {
		t.Errorf("Unexpected status code, expected %d, got %d instead", 200, resp.StatusCode())
	}

	if response.UserID != 2000 {
		t.Errorf("Unexpected UserID, expected %d, got %d instead", 2000, response.UserID)
	}

	if len(response.Segments) != len(segmentsMap) {
		t.Errorf("Unexpected segments, expected %d, got %d instead", len(segmentsMap), len(response.Segments))
	}

	for i := 0; i < len(response.Segments); i++ {
		if ok := segmentsMap[response.Segments[i]]; !ok {
			t.Errorf("Unexpected segment: %s", response.Segments[i])
		}
	}
}

func Test_PUT_AddSegmentForUser2001_1(t *testing.T) {
	client := resty.New()

	var response models.UserSegmentsResponse

	segmentsMap := map[string]bool{
		"AVITO_CLASSIC_15": true,
		"AVITO_HOME":       true,
	}

	segments := models.Segments{
		SegmentsName: []string{"AVITO_CLASSIC_15", "AVITO_HOME"},
	}

	resp, _ := client.R().SetBody(&segments).SetResult(&response).Put("http://localhost:8080/user/2001/segments")

	if resp.StatusCode() != 200 {
		t.Errorf("Unexpected status code, expected %d, got %d instead", 200, resp.StatusCode())
	}

	if response.UserID != 2001 {
		t.Errorf("Unexpected UserID, expected %d, got %d instead", 2001, response.UserID)
	}

	for i := 0; i < len(response.Segments); i++ {
		if ok := segmentsMap[response.Segments[i]]; !ok {
			t.Errorf("Unexpected segment: %s", response.Segments[i])
		}
	}
}

func Test_PUT_AddSegmentForUser2001_2(t *testing.T) {
	client := resty.New()

	var response models.UserSegmentsResponse

	segmentsMap := map[string]bool{
		"AVITO_WORK": true,
	}

	segments := models.Segments{
		SegmentsName: []string{"AVITO_WORK"},
	}

	resp, _ := client.R().SetBody(&segments).SetResult(&response).Put("http://localhost:8080/user/2001/segments")

	if resp.StatusCode() != 200 {
		t.Errorf("Unexpected status code, expected %d, got %d instead", 200, resp.StatusCode())
	}

	if response.UserID != 2001 {
		t.Errorf("Unexpected UserID, expected %d, got %d instead", 2001, response.UserID)
	}

	if len(response.Segments) != len(segmentsMap) {
		t.Errorf("Unexpected segments, expected %d, got %d instead", len(segmentsMap), len(response.Segments))
	}

	for i := 0; i < len(response.Segments); i++ {
		if ok := segmentsMap[response.Segments[i]]; !ok {
			t.Errorf("Unexpected segment: %s", response.Segments[i])
		}
	}
}

func Test_PUT_AddSegmentForUser_StatusCodeShouldEqual500(t *testing.T) {
	client := resty.New()

	segments := models.Segments{
		SegmentsName: []string{"AVITO_SUPER_SALE_50", "AVITO_HOME", "AVITO_AWESOME_90"},
	}

	resp, _ := client.R().SetBody(&segments).Put("http://localhost:8080/user/2000/segments")

	if resp.StatusCode() != 400 {
		t.Errorf("Unexpected status code, expected %d, got %d instead", 400, resp.StatusCode())
	}
}

func Test_GET_GetUserSegments2001(t *testing.T) {
	client := resty.New()

	var response models.UserSegmentsResponse

	segmentsMap := map[string]bool{
		"AVITO_CLASSIC_15": true,
		"AVITO_HOME":       true,
		"AVITO_WORK":       true,
	}

	resp, _ := client.R().SetResult(&response).Get("http://localhost:8080/user/2001/segments")

	if resp.StatusCode() != 200 {
		t.Errorf("Unexpected status code, expected %d, got %d instead", 200, resp.StatusCode())
	}

	if response.UserID != 2001 {
		t.Errorf("Unexpected UserID, expected %d, got %d instead", 2001, response.UserID)
	}

	if len(response.Segments) != len(segmentsMap) {
		t.Errorf("Unexpected segments, expected %d, got %d instead", len(segmentsMap), len(response.Segments))
	}

	for i := 0; i < len(response.Segments); i++ {
		if ok := segmentsMap[response.Segments[i]]; !ok {
			t.Errorf("Unexpected segment: %s", response.Segments[i])
		}
	}
}
