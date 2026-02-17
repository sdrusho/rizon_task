package config

import (
	"crypto/hmac"
	"crypto/sha256"
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"
)

type TokenHeader struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type TokenBody struct {
	User   string `json:"user"`
	Expiry int64  `json:"exp"`
	Iss    string `json:"iss"`
	Type   string `json:"type"`
}

// CreateToken - Create a JWT access and refresh token from the UserName and roles passed
// both tokens will be signed with the same signing key TODO: Need to read this from config or DB and it needs to be environment specific

func CreateToken(tokenKey string, tokenType, userId string, expTime time.Time) (string, error) {

	hostName, _ := os.Hostname()

	tokenClaims := TokenBody{
		User:   userId,
		Expiry: expTime.Unix(),
		Iss:    hostName,
		Type:   tokenType,
	}

	baseToken, err := assemblePayload(tokenClaims)
	if err != nil {
		return "", err
	}

	// ** Ok, now sign it **
	signed := signToken(tokenKey, baseToken)

	baseToken += "." + signed
	return baseToken, nil
}

func ParseToken(signerKey string, jwtToken string) (*TokenBody, error) {

	parts := strings.Split(jwtToken, ".")
	if len(parts) < 2 {
		return nil, errors.New("invalid token")
	}

	signature := signToken(signerKey, parts[0]+"."+parts[1])
	if signature != parts[2] {
		return nil, errors.New("invalid token")
	}

	s := parts[1]
	if i := len(s) % 4; i != 0 {
		s += strings.Repeat("=", 4-i)
	}
	decoded, _ := b64.StdEncoding.DecodeString(s)

	tokenBody := &TokenBody{}
	err := json.Unmarshal(decoded, &tokenBody)
	return tokenBody, err
}

func StringToPointer(s string) *string {
	return &s
}

func assemblePayload(claims TokenBody) (string, error) {
	var err error
	header := TokenHeader{
		Alg: "HS256",
		Typ: "JWT",
	}

	parts := make([]string, 2)
	for i := range parts {
		var jsonValue []byte
		if i == 0 {
			if jsonValue, err = json.Marshal(header); err != nil {
				return "", err
			}
		} else {
			if jsonValue, err = json.Marshal(claims); err != nil {
				return "", err
			}
		}

		parts[i] = strings.TrimRight(b64.URLEncoding.EncodeToString(jsonValue), "=")
	}
	return strings.Join(parts, "."), nil
}

func signToken(tokenKey string, token string) string {
	signer := hmac.New(sha256.New, []byte(tokenKey))
	signer.Write([]byte(token))
	hmacResult := signer.Sum(nil)

	return strings.TrimRight(b64.URLEncoding.EncodeToString(hmacResult), "=")
}
