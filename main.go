package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mmcdole/gofeed"
)

const (
	u = ""
)

func main() {
	feed, err := gofeed.NewParser().ParseURL(u)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(feed.Title)
	fmt.Println(feed.FeedType, feed.FeedVersion)
	for _, item := range feed.Items {
		if item == nil {
			break
		}
		fmt.Println(item.Title)
		fmt.Println("\t->", item.Link)
		fmt.Println("\t->", item.PublishedParsed.Format(time.RFC3339))
	}
}
