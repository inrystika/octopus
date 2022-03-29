package server

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
	api "server/openai-server/api/v1"
	"server/openai-server/internal/conf"
	"server/openai-server/internal/service"
	"strings"
)

type UserInfo struct {
	UserId   string `json:"uid"`
	UserName string `json:"loginName"`
}

func getConfByPlatform(oauth *conf.Oauth, platform string) (string, *oauth2.Config) {
	if platform == "pcl" {
		return oauth.Pcl.AuthServerURL, &oauth2.Config{
			ClientID:     oauth.Pcl.ClientID,
			ClientSecret: oauth.Pcl.ClientSecret,
			RedirectURL:  oauth.Pcl.RedirectURL,
			Endpoint: oauth2.Endpoint{
				AuthURL:  oauth.Pcl.AuthServerURL + "/authorize",
				TokenURL: oauth.Pcl.AuthServerURL + "/getToken",
			},
		}
	} else {
		return "", nil
	}
}

func NewOauthHandler(conf *conf.Server, ctx context.Context, service *service.Service) http.Handler {
	var AuthServerURL string
	var config *oauth2.Config

	r := mux.NewRouter()
	r.HandleFunc("/v1/oauth2/{platform}/authorize", func(w http.ResponseWriter, r *http.Request) {
		AuthServerURL, config = getConfByPlatform(conf.GetOauth(), strings.Split(r.URL.Path, "/")[3])
		u := config.AuthCodeURL("xyz",
			oauth2.SetAuthURLParam("response_type", "code"))
		http.Redirect(w, r, u, http.StatusFound)
	}).Methods("GET")

	r.HandleFunc("/v1/oauth2/{platform}/callback", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		state := r.Form.Get("state")
		if state != "xyz" {
			http.Error(w, "State invalid", http.StatusBadRequest)
			return
		}
		code := r.Form.Get("code")
		if code == "" {
			http.Error(w, "Code not found", http.StatusBadRequest)
			return
		}
		token, err := config.Exchange(context.Background(), code)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		resp, err := http.Get(fmt.Sprintf("%s/getUserInfo?access_token=%s&client_id=%s", AuthServerURL, token.AccessToken, config.ClientID))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}
		var user UserInfo
		if err := json.Unmarshal(body, &user); err != nil {
			return
		}
		reqBind := &api.Bind{
			Platform: strings.Split(r.URL.Path, "/")[3],
			UserId:   user.UserId,
			UserName: user.UserName,
		}
		reply, err := service.AuthService.GetTokenByBind(ctx, &api.GetTokenRequest{
			Bind: reqBind,
		})
		if err != nil {
			return
		}
		url := fmt.Sprintf("%s?token=%s&thirdUserId=%s&thirdUserName=%s", conf.Oauth.RegisterURL, reply.Token, base64.StdEncoding.EncodeToString([]byte(user.UserId)), user.UserName)
		http.Redirect(w, r, url, http.StatusFound)

	}).Methods("GET")
	return r
}
