package pkg

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/smirkcat/go-google-translate/pkg/client"
)

type TranslateRequest struct {
	SourceLang string
	TargetLang string
	Text       string
}

func Translate(req *TranslateRequest) (string, error) {
	config := &client.Config{
		Source: req.SourceLang,
		Target: req.TargetLang,
	}
	resp := client.NewClient(config).Translate(req.Text).Get().Do()
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Failed.StatusCode %d Status:%s", resp.StatusCode, resp.Status)
	}
	respHtml := string(resp.ResponseBody)
	re := regexp.MustCompile(`class="t0">(.*?)<`)
	match := re.FindStringSubmatch(respHtml)
	if len(match) != 2 {
		return "", errors.New("Failed to translate")
	}
	translated := strings.Replace(match[1], "&quot;", "", -1)
	translated = strings.Replace(translated, "&#39;", `'`, -1)
	return translated, nil
}
