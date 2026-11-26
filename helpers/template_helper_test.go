package helpers

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestLoadTemplates(t *testing.T) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("impossible de trouver le dossier du test")
	}

	previousDirectory, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(previousDirectory); err != nil {
			t.Errorf("restauration du dossier de travail : %v", err)
		}
	})

	if err := os.Chdir(filepath.Dir(filepath.Dir(filename))); err != nil {
		t.Fatal(err)
	}

	LoadTemplates()
	if templates == nil {
		t.Fatal("les templates n'ont pas été chargés")
	}
}

func TestWeatherDisplay(t *testing.T) {
	tests := []struct {
		code  int
		icon  string
		label string
	}{
		{code: 0, icon: "☀️", label: "Ciel dégagé"},
		{code: 63, icon: "🌧️", label: "Pluie"},
		{code: 75, icon: "🌨️", label: "Neige"},
		{code: 95, icon: "⛈️", label: "Orages"},
	}

	for _, test := range tests {
		if got := weatherIcon(test.code); got != test.icon {
			t.Errorf("weatherIcon(%d) = %q, attendu %q", test.code, got, test.icon)
		}
		if got := weatherLabel(test.code); got != test.label {
			t.Errorf("weatherLabel(%d) = %q, attendu %q", test.code, got, test.label)
		}
	}
}
