package head

import "testing"

func TestGetHeadFromNet(t *testing.T) {
	head, err := GetHeadFromNet("Dinnerbone", true)
	if err != nil {
		t.Error("Error should not occur")
	}
	for i := 0; i < len(head); i++ {
		for j := 0; j < len(head[0]); j++ {
			_, _, _, a := head[i][j].RGBA()
			if a == 0 {
				t.Error("100% transparent pixel on head")
			}
		}
	}

	head, err = GetHeadFromNet("", true)
	if err == nil {
		t.Error("Error should occur")
	}
}
