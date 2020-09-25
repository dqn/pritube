package metadata

import (
	"fmt"
	"os"
	"testing"
)

func TestFetch(t *testing.T) {
	m := New()

	// m.Language = "ja"
	metadata, err := m.Fetch(os.Getenv("VIDEO_ID"))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("ViewCount: %s\n", metadata.ViewCount)
	fmt.Printf("ShortViewCount: %s\n", metadata.ShortViewCount)
	fmt.Printf("ViewCount: %d\n", metadata.ViewCountInt)
	fmt.Printf("IsLive: %t\n", metadata.IsLive)
	fmt.Printf("LikeCount: %s\n", metadata.LikeCount)
	fmt.Printf("DislikeCount: %s\n", metadata.DislikeCount)
	fmt.Printf("Date: %s\n", metadata.Date)
	fmt.Printf("Title: %s\n", metadata.Title)
	fmt.Printf("Description: %s\n", metadata.Description)
}
