package config

import (
	"testing"
)

func TestInitialization(t *testing.T) {
	var empty Conf

	if Configuration == empty {
		t.Errorf("Configuration is Empty : %v", Configuration)
	}
}
