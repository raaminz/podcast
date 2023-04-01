package cmd

import (
	"bytes"
	_ "embed"
	"encoding/xml"
	"io"
	"net/url"
	"strings"
	"text/template"
	"time"

	"github.com/BurntSushi/toml"
)

//go:embed podcast.toml
var tomlData string

//go:embed template_feed.gorss
var templateData string

func GenerateRSS(writer io.Writer) error {
	data, err := readTomlFile(tomlData)
	if err != nil {
		return err
	}
	err = formatRss(writer, data)
	if err != nil {
		return err
	}
	return nil
}

func readTomlFile(tomlData string) (map[string]any, error) {
	var data map[string]any
	if _, err := toml.Decode(tomlData, &data); err != nil {
		return nil, err
	}
	return data, nil
}

type RssData struct {
	BuildDate string
	Data      map[string]any
}

func (r *RssData) Get(key string) any {
	return r.Data[key]
}

func (r *RssData) GetEp(data map[string]any, key string) any {
	if v, ok := data[key]; ok {
		return v
	}

	if defaults, ok := r.Data["episodeDefaults"].(map[string]any); ok {
		return defaults[key]
	}
	return nil
}

func (r *RssData) GetEpPath(data map[string]any, key string) string {
	path := url.PathEscape(r.GetEp(data, "path").(string))
	return strings.Replace(r.GetEp(data, key).(string), "{path}", path, 1)
}

func (r *RssData) GetEpEscape(data map[string]any, key string) string {
	var buf bytes.Buffer
	err := xml.EscapeText(&buf, []byte(r.GetEp(data, key).(string)))
	if err != nil {
		return ""
	}
	return buf.String()
}

func (r *RssData) GetR(group, key string) any {
	return r.Data[group].(map[string]any)[key]
}

func formatRss(writer io.Writer, data map[string]any) error {
	rssData := RssData{
		// BuildDate format like Tue, 19 Apr 2022 09:32:49 GMT
		BuildDate: time.Now().Format("Mon, 02 Jan 2006 15:04:05 GMT"),
		Data:      data,
	}
	tmpl, err := template.New("rss").Parse(templateData)
	if err != nil {
		return err
	}
	err = tmpl.Execute(writer, &rssData)
	if err != nil {
		return err
	}
	return nil
}
