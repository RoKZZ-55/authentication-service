package model

type UserAuthentication struct {
	GUID         string `json:"guid" bson:"user_guid"`
	TokenGUID    string `json:"token-guid" bson:"token_guid"`
	AccessToken  string `json:"access-token,omitempty"`
	RefreshToken string `json:"refresh-token,omitempty" bson:"refresh_token_hash"`
}
