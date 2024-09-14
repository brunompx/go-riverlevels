package retriever

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
)

func GetDataFromWeb(url string) []byte {

	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var data string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Text(`body`, &data, chromedp.NodeVisible),
	)
	if err != nil {
		log.Fatal(err)
	}
	return []byte(data)
}
