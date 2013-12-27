package wow

import (
	"fmt"
	"strings"
	"errors"
	"net/http"
	"net/url"
	"time"
	"encoding/base64"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/json"
	"io/ioutil"
)

type ApiClient struct {
	Host string
	Locale string
	Secret string
	PublicKey string
}

func NewApiClient(region string, locale string) (*ApiClient, error) {
	var host string
	var validLocales []string
	switch region {
	case "US", "United States":
		host = "us.battle.net"
		validLocales = []string{"en_US", "es_MX", "pt_BR"}
	case "EU", "Europe":
		host = "eu.battle.net"
		validLocales = []string{"en_GB", "es_ES", "fr_FR", "ru_RU", "de_DE", "pt_PT", "it_IT"}
	case "KR", "Korea":
		host = "kr.battle.net"
		validLocales = []string{"ko_KR"}
	case "TW", "Taiwan":
		host = "tw.battle.net"
		validLocales = []string{"zh_TW"}
	case "ZH", "CN", "China":
		host = "www.battle.com.cn"
		validLocales = []string{"zh_CN"}
	default:
		return nil, errors.New(fmt.Sprintf("Region '%s' is not valid", region))
	}

	if locale == "" {
		return &ApiClient{Host: host, Locale: validLocales[0]}, nil
	} else {
		for _, valid := range validLocales {
			if valid == locale {
				return &ApiClient{Host: host, Locale: locale}, nil
			}
		}
	}
		
	return nil, errors.New(fmt.Sprintf("Locale '%s' is not valid for region '%s'", locale, region))
}

func (a *ApiClient) GetAchievement(id int) (*Achievement, error) {
	jsonBlob, err := a.get(fmt.Sprintf("achievement/%d", id))
	if err != nil {
		return nil, err
	}
	achieve := &Achievement{}
	err = json.Unmarshal(jsonBlob, achieve)
	if err != nil {
		return nil, err
	}
	return achieve, nil
}

func (a *ApiClient) get(path string) ([]byte, error) {
	url := a.url(path)
	client := &http.Client{}

	request, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return make([]byte, 0), err
	}

	if len(a.Secret) > 0 {
		request.Header.Add("Authorization", a.authorizationString(a.signature("GET", path)))
	}

	response, err := client.Do(request)
	if err != nil {
		return make([]byte, 0), err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return make([]byte, 0), err
	}
	
	return body, nil
}

func (a *ApiClient) url(path string) *url.URL {
	return &url.URL{
		Scheme: "http",
		Host: a.Host,
		Path: "/api/wow/" + path,
		RawQuery: "locale=" + a.Locale,
	}
}

func (a *ApiClient) authorizationString(signature string) string {
	return fmt.Sprintf(" BNET %s:%s", a.PublicKey, signature)
}

func (a *ApiClient) signature(verb string, path string) string {
	url := a.url(path)
	toBeSigned := []byte(strings.Join([]string{verb, time.Now().String(), url.Path, ""}, "\n"))
	mac := hmac.New(sha1.New, []byte(a.Secret))
	_, err := mac.Write(toBeSigned) // FIXME _ = signed
	if err != nil {
		handleError(err)
	}
	return base64.StdEncoding.EncodeToString([]byte("hi")) //FIXME Figure out crypto
}

func handleError(err error) {
	panic(err)
}

