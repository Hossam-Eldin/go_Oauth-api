package accesstoken

import "time"

const (
	expirationTime =24
)

//AccessToken : type struct
type AccessToken struct{
	AccessToken string`json:"accest_token"`
	UserID  	int64`json:"user_id"`
	ClientID 	int64`json:"client_id"`
	Expires      int64`json:"expires"`
}

//GetNewAccessToken : this will return the accesstoken
func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}
//IsExpired : this check for access token or expired
func (at AccessToken) IsExpired() bool  {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}