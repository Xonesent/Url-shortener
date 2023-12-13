package url_shortener

import (
	"fmt"
	"net/url"
	"regexp"
)

type Link struct {
	Base_URL  string `json:"base_url"`
	Short_URL string `json:"short_url"`
}

func Validate_Base_URL(link *Link) error {
	if link == nil {
		return fmt.Errorf("nil url")
	}

	if link.Base_URL == "" {
		return fmt.Errorf("empty url")
	}

	parsed, err := url.Parse(link.Base_URL)
	if err != nil {
		return fmt.Errorf("%v is an invalid URL: %v", link.Base_URL, err)
	}

	if !(parsed != nil && parsed.Scheme != "" && parsed.Host != "") {
		return fmt.Errorf("%v is an invalid URL: %v", link.Base_URL, err)
	}

	return nil
}

func Validate_Short_URL(link *Link) error {
	if link == nil {
		return fmt.Errorf("pass nil pointer")
	}

	regex := regexp.MustCompile(`^[a-zA-Z0-9_]{10}$`)
	if !regex.MatchString(link.Short_URL) {
		return fmt.Errorf("%v is an invalid Short_URL", link.Short_URL)
	}

	return nil
}
