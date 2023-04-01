package cmd

import (
	"bytes"
	"fmt"
	"testing"

	_ "embed"
)

func TestReadTomlFile(t *testing.T) {
	data, err := readTomlFile(tomlData)

	if err != nil {
		t.Fatal(err)
	}
	if _, ok := data["title"]; !ok {
		t.Fatal("title not found")
	}
	if _, ok := data["links"]; !ok {
		t.Fatal("links not found")
	}
	if _, ok := data["episodes"]; !ok {
		t.Fatal("episodes not found")
	}

	switch data["episodes"].(type) {
	case []map[string]any:
		// ok
	default:
		t.Fatal("episodes is not a slice")
	}

	// for k, v := range data {
	// log.Printf("%s: %v\n", k, v)
	// }
}

func TestReadTomlError(t *testing.T) {
	_, err := readTomlFile("invalid toml data")
	if err == nil {
		t.Fatal("error should be returned")
	}
}

func TestGetKey(t *testing.T) {
	data := make(map[string]any)
	data["author"] = "Ramin Zare"
	data["URL"] = "https://raminzare.com"
	links := make(map[string]any)
	data["links"] = links
	links["cover"] = "/cover.jpg"

	defaults := make(map[string]any)
	data["episodeDefaults"] = defaults
	defaults["content"] = "/content.mp3"

	episodes := make([]map[string]any, 0, 0)

	ep1 := make(map[string]any)
	ep1["title"] = "Episode 1"
	episodes = append(episodes, ep1)

	data["episodes"] = episodes

	obj := RssData{
		BuildDate: "2018-01-01",
		Data:      data,
	}
	title := obj.Get("author")
	if title != "Ramin Zare" {
		t.Fatal("title is not correct")
	}
	cover := obj.GetR("links", "cover")
	if cover != "/cover.jpg" {
		t.Fatal("cover is not correct")
	}
	content := obj.GetEp(data["episodes"].([]map[string]any)[0], "content")
	if content != "/content.mp3" {
		t.Fatal("content is not correct")
	}
}

func TestFormatRss(t *testing.T) {
	data, err := readTomlFile(tomlData)
	if err != nil {
		t.Fatal(err)
	}
	var buf bytes.Buffer
	err = formatRss(&buf, data)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(buf.String())
}
