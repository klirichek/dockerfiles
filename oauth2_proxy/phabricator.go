package providers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
        "io/ioutil"
	"log"
	"net/http"
        "net/url"
        "crypto/tls"
        "github.com/bitly/go-simplejson"
)

type PhabricatorProvider struct {
	*ProviderData
}

func NewPhabricatorProvider(p *ProviderData) *PhabricatorProvider {
        phabricatorHost := p.LoginUrl.String();

        p.ProviderName = "Phabricator"
        p.LoginUrl = &url.URL{Scheme: "https",
			Host: phabricatorHost,
                        Path: "/oauthserver/auth"}
	if p.RedeemUrl.String() == "" {
                p.RedeemUrl = &url.URL{Scheme: "https",
			Host: phabricatorHost,
                        Path: "/oauthserver/token/"}
	}
	if p.ProfileUrl.String() == "" {
		p.ProfileUrl = &url.URL{Scheme: "https",
			Host: phabricatorHost,
                        Path: "/api/user.whoami"}
	}
	if p.ValidateUrl.String() == "" {
		p.ValidateUrl = &url.URL{Scheme: "https",
			Host: phabricatorHost,
                        Path: "/api/user.whoami"}
	}
	if p.Scope == "" {
                p.Scope = "whoami"
	}
	return &PhabricatorProvider{ProviderData: p}
}

func (p *PhabricatorProvider) GetEmailAddress(s *SessionState) (string, error) {
	req, err := http.NewRequest("GET",
		p.ProfileUrl.String()+"?access_token="+s.AccessToken, nil)
	if err != nil {
		log.Printf("failed building request %s", err)
                return "", err
        }
        var resp *http.Response
        var client *http.Client
        // Temporarily adding insecure certificate bypass for testing
        allow_insecure := true
        if allow_insecure {
          tr := &http.Transport{
            TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
          }
          client = &http.Client{Transport: tr}
        } else {
          client = http.DefaultClient
        }
        resp, err = client.Do(req)
	if err != nil {
                log.Printf("%s %s %s", req.Method, req.URL, err)
                return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	log.Printf("%d %s %s %s", resp.StatusCode, req.Method, req.URL, body)
        if err != nil {
                log.Printf("%s %s %s", req.Method, req.URL, err)
                return "", err
	}
        if resp.StatusCode != 200 {
                log.Printf("%s %s %s", req.Method, req.URL, err)
                return "", err
	}
        json, err := simplejson.NewJson(body)
        if err != nil {
                log.Printf("%s %s %s", req.Method, req.URL, err)
                return "", err
        }
        //	json, err := api.Request(req)
	if err != nil {
		log.Printf("failed making request %s", err)
                return "", err
	}
        return json.Get("result").Get("primaryEmail").String()
}

func (p *PhabricatorProvider) Redeem(redirectUrl, code string) (s *SessionState, err error) {
	if code == "" {
		err = errors.New("missing code")
		return
	}

	params := url.Values{}
	params.Add("redirect_uri", redirectUrl)
	params.Add("client_id", p.ClientID)
	params.Add("client_secret", p.ClientSecret)
	params.Add("code", code)
	params.Add("grant_type", "authorization_code")
	var req *http.Request
	req, err = http.NewRequest("POST", p.RedeemUrl.String(), bytes.NewBufferString(params.Encode()))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

        var resp *http.Response
        var client *http.Client
        // Temporarily adding insecure certificate bypass for testing
        allow_insecure := true
        if allow_insecure {
          tr := &http.Transport{
            TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
          }
          client = &http.Client{Transport: tr}
        } else {
          client = http.DefaultClient
        }
        resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("got %d from %q %s", resp.StatusCode, p.RedeemUrl.String(), body)
		return
	}

	// blindly try json and x-www-form-urlencoded
	var jsonResponse struct {
		AccessToken string `json:"access_token"`
	}
	err = json.Unmarshal(body, &jsonResponse)
	if err == nil {
		s = &SessionState{
			AccessToken: jsonResponse.AccessToken,
		}
		return
	}

	var v url.Values
	v, err = url.ParseQuery(string(body))
	if err != nil {
		return
	}
	if a := v.Get("access_token"); a != "" {
		s = &SessionState{AccessToken: a}
	} else {
		err = fmt.Errorf("no access token found %s", body)
	}
	return
}
