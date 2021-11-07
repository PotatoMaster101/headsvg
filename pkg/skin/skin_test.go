package skin

import (
	"strings"
	"testing"
)

func TestGetSkinURL(t *testing.T) {
	// test good username
	url, err := GetSkinURL("MHF_Pig")
	if err != nil {
		t.Errorf("Error should not occur")
	}
	if url != "http://textures.minecraft.net/texture/a562a37b871f964bfc3e1311ea672aaa03984a5dc472154a34dc25af157e382b" {
		t.Errorf("Invalid URL, got: %s, want: %s", url, "http://textures.minecraft.net/texture/a562a37b871f964bfc3e1311ea672aaa03984a5dc472154a34dc25af157e382b")
	}

	// test bad username
	url, err = GetSkinURL("")
	if err == nil {
		t.Error("Error should occur")
	}
	if url != "" {
		t.Error("URL should be empty")
	}
}

func TestGetTexture(t *testing.T) {
	// test good username
	texture, err := GetTexture("MHF_Pig")
	if err != nil {
		t.Error("Error should not occur")
	}
	if !strings.Contains(texture, "\"url\" : \"http://textures.minecraft.net/texture/a562a37b871f964bfc3e1311ea672aaa03984a5dc472154a34dc25af157e382b\"") {
		t.Error("Skin URL not found")
	}

	// test bad username
	texture, err = GetTexture("")
	if err == nil {
		t.Error("Error should occur")
	}
	if texture != "" {
		t.Error("Texture should be empty")
	}
}
