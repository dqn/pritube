package chvideos

import (
	"fmt"
	"os"
	"testing"
)

func TestFetchAll(t *testing.T) {
	videos, err := FetchAll(os.Getenv("CHANNEL_ID"))
	if err != nil {
		t.Fatal(err)
	}

	for _, video := range videos {
		fmt.Println(video.VideoID)
	}

	fmt.Println(len(videos))
}

func TestFetchNext(t *testing.T) {
	c := New(os.Getenv("CHANNEL_ID"))

	cnt := 0
	for !c.IsCompleted {
		gvr, err := c.FetchNext()
		if err != nil {
			t.Fatal(err)
		}

		for _, g := range gvr {
			fmt.Println(g.VideoID)
		}

		cnt += len(gvr)
	}

	fmt.Println(cnt)
}
