package email

import (
	"crypto/tls"
	"github.com/go-gomail/gomail"
	"scheduler/pkg/applet/utils"
	"time"
	"encoding/json"
	"net/http"
	"go.uber.org/zap"
)

type EmailInfo struct {
	ServerHost string
	ServerPort int
	FromEmail  string
	FromPasswd string
	Recipient  []string
}

type UserCenterAccessToken struct {
	AccessToken 	string	`json:"access_token"`
	TokenType		string	`json:"token_type"`
	ExpireIn		float64	`json:"expires_in"`
	ExpireAt    	float64	`json:"expires_at"`
}

var emailMessage *gomail.Message
var userCenterToken *UserCenterAccessToken

func GetUserEmailAddr(logger *zap.Logger, userId, userCenterAddr, ClientId, ClientSecret, GrantType, Scope string) string{
    
    var userToken = GetUserCenterToken(logger, userCenterAddr, ClientId, ClientSecret, GrantType, Scope)
	var token = "Bearer " + userToken.AccessToken;

	headers := make(map[string]string)
	headers["Authorization"] = token

	bytes, _, err := utils.DoRequest( "GET", userCenterAddr + "/api/user/list?uid=" + userId, "", headers)

	var result map[string]interface{}
	err = json.Unmarshal([]byte(bytes), &result)

    if err != nil {
		logger.Error("GetUserEmailAddr Failed: " + err.Error())
	}

	data := result["data"].(map[string]interface{})
	rows := data["rows"].([]interface{})
	if len(rows) == 0 {
		logger.Error("GetUserEmailAddr Failed!!!!")
		return ""
	}
	row := (data["rows"].([]interface{})[0]).(map[string]interface{})
	email := row["email"].(string)

	return email
  }

func GetUserCenterToken(logger *zap.Logger, userCenterAddr, ClientId, ClientSecret, GrantType, Scope string) *UserCenterAccessToken {
    if userCenterToken == nil {
		userCenterToken = getAccessTokenByClient(logger, userCenterAddr, ClientId, ClientSecret, GrantType, Scope);
      	formatAccessToken(userCenterToken);
    } else if float64(time.Now().Unix()) >= userCenterToken.ExpireAt {
        userCenterToken = getAccessTokenByClient(logger, userCenterAddr, ClientId, ClientSecret, GrantType, Scope);
		formatAccessToken(userCenterToken);
	}
    return userCenterToken
}

func getAccessTokenByClient(logger *zap.Logger, userCenterAddr, ClientId, ClientSecret, GrantType, Scope string) *UserCenterAccessToken{

	var addr = userCenterAddr + "/oauth/token"
	
	form := map[string]string{
		"client_id": ClientId,
		"client_secret": ClientSecret,
		"grant_type": GrantType,
		"scope": Scope,
	}
	
	bytes, code, err := utils.PostForm(addr, form)

	if err != nil {
		logger.Error("getAccessTokenByClient Failed: " + err.Error())
	}

	if code != http.StatusOK {
		logger.Error("getAccessTokenByClient Failed, " + string(bytes))
	}

	var result *UserCenterAccessToken = &UserCenterAccessToken{}

	json.Unmarshal(bytes, result)

	return result
}

func formatAccessToken(accessToken *UserCenterAccessToken) (*UserCenterAccessToken) {
	advanceTimeOut := 5 * 60;
    accessToken.ExpireAt = float64(time.Now().Unix()) - float64(advanceTimeOut) + accessToken.ExpireIn;
    return accessToken;
}

func SendEmail(logger *zap.Logger, subject, body string, emailInfo *EmailInfo) {

	if len(emailInfo.Recipient) == 0 {
		logger.Error("Email receipient can not be empty!")
		return
	}

	emailMessage = gomail.NewMessage()
	//emailMessage.SetHeader("To", "hackmong@163.com")
	emailMessage.SetHeader("To", emailInfo.Recipient...)
	emailMessage.SetAddressHeader("From", emailInfo.FromEmail, "")
	emailMessage.SetHeader("Subject", subject)
	emailMessage.SetBody("text/html", body)

	d := gomail.NewPlainDialer(emailInfo.ServerHost, emailInfo.ServerPort,
		emailInfo.FromEmail, emailInfo.FromPasswd)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := d.DialAndSend(emailMessage)
	if err != nil {
		logger.Error("Email Send Failed: " + err.Error())
	}else{
		logger.Info("Email Send Success, recipient: " +  emailInfo.Recipient[0])
	}
}