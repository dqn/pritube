package video

import (
	"fmt"
	"os"
	"testing"
)

func TestGetVideoInfo(t *testing.T) {
	pr, err := FetchInfo(os.Getenv("VIDEO_ID"))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(pr.VideoDetails.Title)
}
