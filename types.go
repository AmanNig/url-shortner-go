package main

import (
	"crypto/rand"
	"encoding/base64"
)

type ShortUrl struct {
	ID int `ID`
    LongURL string `json:"long_url"`
    Alias   string `json:"alias"`
}
type ShortUrlReq struct{
	LongURL string `json:"long_url"`
}
type ShortUrlResp struct{
	Alias   string `json:"alias"`
}

func NewUrl(long string) *ShortUrl{
	ali:= GenerateShortURL(8)
	return &ShortUrl{
		LongURL: long,
		Alias: ali,
	}

}
func GenerateShortURL(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	// Encode to base64 and trim unwanted characters
	return base64.URLEncoding.EncodeToString(bytes)[:length]
}