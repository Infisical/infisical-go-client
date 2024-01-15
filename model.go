package infisicalclient

import "time"

type GetEncryptedSecretsV3Request struct {
	Environment    string
	WorkspaceId    string
	SecretPath     string
	IncludeImports bool
}
type GetEncryptedSecretsV3Response struct {
	Secrets []EncryptedSecret `json:"secrets"`
	Imports []interface{}     `json:"imports"`
}

type GetServiceTokenDetailsResponse struct {
	Id        string `json:"_id"`
	Name      string `json:"name"`
	Workspace string `json:"workspace"`
	Scopes    []struct {
		Environment string `json:"environment"`
		SecretPath  string `json:"secretPath"`
		Id          string `json:"_id"`
	} `json:"scopes"`
	User struct {
		Id           string        `json:"_id"`
		AuthMethods  []string      `json:"authMethods"`
		Email        string        `json:"email"`
		FirstName    string        `json:"firstName"`
		LastName     string        `json:"lastName"`
		IsMfaEnabled bool          `json:"isMfaEnabled"`
		MfaMethods   []interface{} `json:"mfaMethods"`
		CreatedAt    time.Time     `json:"createdAt"`
		UpdatedAt    time.Time     `json:"updatedAt"`
		V            int           `json:"__v"`
	} `json:"user"`
	LastUsed     time.Time `json:"lastUsed"`
	EncryptedKey string    `json:"encryptedKey"`
	Iv           string    `json:"iv"`
	Tag          string    `json:"tag"`
	Permissions  []string  `json:"permissions"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	V            int       `json:"__v"`
}

type EncryptedSecret struct {
	Id                      string        `json:"_id"`
	Version                 int           `json:"version"`
	Workspace               string        `json:"workspace"`
	Type                    string        `json:"type"`
	Tags                    []interface{} `json:"tags"`
	Environment             string        `json:"environment"`
	SecretKeyCiphertext     string        `json:"secretKeyCiphertext"`
	SecretKeyIV             string        `json:"secretKeyIV"`
	SecretKeyTag            string        `json:"secretKeyTag"`
	SecretValueCiphertext   string        `json:"secretValueCiphertext"`
	SecretValueIV           string        `json:"secretValueIV"`
	SecretValueTag          string        `json:"secretValueTag"`
	SecretCommentCiphertext string        `json:"secretCommentCiphertext"`
	SecretCommentIV         string        `json:"secretCommentIV"`
	SecretCommentTag        string        `json:"secretCommentTag"`
	Algorithm               string        `json:"algorithm"`
	KeyEncoding             string        `json:"keyEncoding"`
	Folder                  string        `json:"folder"`
	V                       int           `json:"__v"`
	CreatedAt               time.Time     `json:"createdAt"`
	UpdatedAt               time.Time     `json:"updatedAt"`
}
