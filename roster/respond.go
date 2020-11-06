package roster

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"

	"github.com/pkg/errors"

	"github.com/slack-go/slack"
)

/*
Day.Monday
Day.Tuesday
Day.Wednesday
Day.Thursday

*/
func (s Server) Respond(ctx context.Context, f *firestore.Client) error {
	var wg sync.WaitGroup
	resp, err := http.Get("https://alex.github.io/nyt-2020-election-scraper/battleground-state-changes.csv")
	if err != nil {
		return err
	}
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	oldread, err := download("joshcarp-election", "old")
	old, err := ioutil.ReadAll(oldread)
	d := diff(string(old), string(contents))
	if len(d) == 0 {
		return nil
	}
	upload("joshcarp-election", "old", bytes.NewReader(contents))

	resp, _ = http.Get("https://alex.github.io/nyt-2020-election-scraper/battleground-state-changes.txt")
	contents, _ = ioutil.ReadAll(resp.Body)
	var re = regexp.MustCompile(`(?m)\S* \(EV: \d*\) Total Votes: .*`)
	res := []string{}
	for _, b := range re.FindAllStringSubmatch(string(contents), -1){
		for _, c := range b{
			res = append(res, c)
		}
	}
	a := f.Collection("webhooks")
	docs := a.Documents(ctx)
	docs1, _ := docs.GetAll()
	for _, sub := range docs1 {
		wg.Add(1)
		go func(sub1 firestore.DocumentSnapshot) {
			var secret slack.OAuthV2Response
			sub1.DataTo(&secret)
			PostWebhookCustomHTTPContext(
				ctx,
				secret.IncomingWebhook.URL,
				s.Client,
				&slack.WebhookMessage{
					Text: "New Election update: https://alex.github.io/nyt-2020-election-scraper/battleground-state-changes.html \n"+ strings.Join(res, "\n"),
				})
			wg.Done()
		}(*sub)
	}
	wg.Wait()
	return nil
}

func PostWebhookCustomHTTPContext(ctx context.Context, url string, httpClient HttpClient, msg *slack.WebhookMessage) error {
	raw, err := json.Marshal(msg)
	if err != nil {
		return errors.Wrap(err, "marshal failed")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(raw))
	if err != nil {
		return errors.Wrap(err, "failed new request")
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to post webhook")
	}
	defer resp.Body.Close()

	return checkStatusCode(resp)
}

func checkStatusCode(resp *http.Response) error {
	if resp.StatusCode == http.StatusTooManyRequests {
		retry, err := strconv.ParseInt(resp.Header.Get("Retry-After"), 10, 64)
		if err != nil {
			return err
		}
		return &slack.RateLimitedError{time.Duration(retry) * time.Second}
	}

	// Slack seems to send an HTML body along with 5xx error codes. Don't parse it.
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%d: %s", resp.StatusCode, resp.Status)
	}

	return nil
}

func download(bucket, object string) (io.Reader, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()
	return client.Bucket(bucket).Object(object).NewReader(ctx)
}

func upload(bucket string, object string, r io.Reader) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	wc := client.Bucket(bucket).Object(object).NewWriter(ctx)
	if _, err = io.Copy(wc, r); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}
	return nil
}

func diff(text1, text2 string) []string {
	a := strings.Split(text1, "\n")
	amap := make(map[string]bool)
	b := strings.Split(text2, "\n")
	bmap := make(map[string]bool)
	for _, a2 := range a {
		amap[a2] = true
	}
	for _, b2 := range b {
		bmap[b2] = true
	}
	final := []string{}
	if len(bmap) < len(amap) {
		bmap, amap = amap, bmap
	}
	for key, _ := range bmap {
		if amap[key] == false {
			final = append(final, key)
		}
	}
	return final
}