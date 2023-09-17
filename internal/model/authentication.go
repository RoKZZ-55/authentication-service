package model

type UserAuthentication struct {
	UUID         string `json:"uuid" bson:"user_uuid"`
	TokenUUID    string `json:"token-uuid" bson:"token_uuid"`
	AccessToken  string `json:"access-token,omitempty"`
	RefreshToken string `json:"refresh-token,omitempty" bson:"refresh_token_hash"`
}
