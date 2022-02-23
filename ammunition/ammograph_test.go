package ammunition

import (
	"fmt"
	"os"
	"testing"
)

func TestNewAmmoGraph(t *testing.T) {
	LoadAmmunition("../data/ammunition.json")
	caliber := "9x19PARA"
	ammos := GetAmmosByCaliber(caliber)
	fmt.Printf("NewAmmoGraph(): %d ammos loaded from %s caliber\n", len(ammos), caliber)

	NewAmmoGraph(ammos)
	fi, err := os.Stat("/tmp/example.png")
	if err != nil {
		t.Fatalf("NewAmmoGraph(...) returned an error: %v, error\n", err)
	}

	if fi.Size() <= 0 {
		t.Fatalf("NewAmmoGraph().Size() = %d, want X > 0\n", fi.Size())
	}
	fmt.Printf("NewAmmoGraph().Size() = %d\n", fi.Size())
}
