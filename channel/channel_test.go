package channel

import (
	"fmt"
	"os"
	"testing"
)

func TestFetchAll(t *testing.T) {
	c, err := FetchInfo(os.Getenv("CHANNEL_ID"))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(c.ID)
	fmt.Println(c.Title)
	fmt.Println(c.Description)
	fmt.Println(c.Avatars)
	fmt.Println(c.Banners)
	fmt.Println(c.ChannelURL)
	fmt.Println(c.SubscriberCount)
	fmt.Println(c.ViewCount)
	fmt.Println(c.Country)
	fmt.Println(c.JoinedDate)
}
