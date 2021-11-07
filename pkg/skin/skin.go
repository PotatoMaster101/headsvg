package skin

import (
	"encoding/base64"
	"regexp"

	"github.com/Lukaesebrot/mojango"
)

var client = mojango.New()

// GetSkinURL returns the skin URL for the given `username`.
func GetSkinURL(username string) (string, error) {
	texture, err := GetTexture(username)
	if err != nil {
		return "", err
	}

	reg := regexp.MustCompile("\"url\" : \"(http://textures.minecraft.net/texture/.*)\"")
	match := reg.FindStringSubmatch(texture)
	return match[1], nil
}

// GetTexture returns the texture string for the specified `username`.
func GetTexture(username string) (string, error) {
	uuid, _ := client.FetchUUID(username)
	profile, err := client.FetchProfile(uuid, false)
	if err != nil {
		return "", err
	}

	texture := ""
	for i := range profile.Properties {
		if profile.Properties[i].Name == "textures" {
			texture = profile.Properties[i].Value
			break
		}
	}

	dec, _ := base64.StdEncoding.DecodeString(texture)
	return string(dec), nil
}
