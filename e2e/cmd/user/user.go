package user

import (
	"fmt"
	"os"

	"github.com/opd-ai/toxcore"
	"github.com/sirupsen/logrus"
)

const (
	SaveDataFile = "toxav_integration_profile.dat"
)

var logger = logrus.New()

func CreateNewProfile() (*toxcore.Tox, error) {
	options := toxcore.NewOptions()
	options.UDPEnabled = true

	tox, err := toxcore.New(options)
	if err != nil {
		return nil, err
	}

	if err := tox.SelfSetName("ToxAV Integration Demo"); err != nil {
		logger.WithError(err).Warn("Failed to set name")
	}

	if err := tox.SelfSetStatusMessage("Integrated Tox client with AV calling"); err != nil {
		logger.WithError(err).Warn("Failed to set status")
	}

	return tox, nil
}

func LoadOrCreateProfile() (*toxcore.Tox, error) {
	savedata, err := os.ReadFile(SaveDataFile)
	if err == nil {
		return LoadExistingProfile(savedata)
	}

	fmt.Println("📝 Creating new profile...")
	return CreateNewProfile()
}

func LoadExistingProfile(savedata []byte) (*toxcore.Tox, error) {
	fmt.Printf("📁 Loading existing profile (%d bytes)\n", len(savedata))
	tox, err := toxcore.NewFromSavedata(nil, savedata)
	if err != nil {
		fmt.Printf("⚠️  Failed to restore existing profile, creating new one: %v\n", err)
		return CreateNewProfile()
	}

	fmt.Println("✅ Profile restored successfully")
	return tox, nil
}

func SaveProfile(tox *toxcore.Tox) error {
	if tox == nil {
		return fmt.Errorf("tox profile is nil")
	}

	savedata := tox.GetSavedata()

	return os.WriteFile(SaveDataFile, savedata, 0o600)
}
