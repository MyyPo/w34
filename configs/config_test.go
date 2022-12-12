package configs

import "testing"

func TestLoadConfig(t *testing.T) {
	t.Run("Load config files in folder", func(t *testing.T) {
		got, err := NewConfig("$APP/app/configs")
		if err != nil {
			t.Errorf("error loading config: %v", err)
		}

		t.Logf("Got config: %v", got)
	})
}
