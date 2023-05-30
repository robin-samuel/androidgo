package androidgo

import (
	crand "crypto/rand"
	"encoding/hex"
	"math/rand"

	"github.com/google/uuid"
	"github.com/robin-samuel/androidgo/locale"
	"github.com/robin-samuel/androidgo/os/systemclock"
)

type Device struct {
	UUID           string
	AndroidID      string
	MobileUID      string
	Fingerprint    string
	DisplayMetrics DisplayMetrics
	Build          Build
	SystemClock    systemclock.SystemClock
	Locale         locale.Locale
	IPv6           string
	IPv4           string
	Proxy          string
}

func RandomDevice() Device {
	device := DEVICES[rand.Intn(len(DEVICES))]
	device.generateUUID()
	device.generateAndroidID()
	device.generateMobileUID()
	device.generateFingerprint()
	device.SystemClock = *systemclock.NewSystemClock()
	device.Locale = locale.RandomLocale()
	device.Build.bDISPLAY()
	device.Build.bFINGERPRINT()
	device.Build.bTIME()
	device.Build.VERSION = RandomVersion()
	return device
}

func RandomVersion() Version {
	return VERSIONS[rand.Intn(len(VERSIONS))]
}

func (d *Device) OsType() string {
	return "Android"
}

func (d *Device) generateUUID() {
	// Generate Random UUID
	d.UUID = uuid.New().String()
}

func (d *Device) generateAndroidID() {
	// Generate Random Android ID (16 characters)
	id := make([]byte, 16)
	crand.Read(id)

	d.AndroidID = hex.EncodeToString(id)
}

func (d *Device) generateMobileUID() {
	// Generate Random Mobile UID (41 characters)
	uid := make([]byte, 41)
	crand.Read(uid)

	d.MobileUID = hex.EncodeToString(uid)
}

func (d *Device) generateFingerprint() {
	// Generate Random Fingerprint (64 characters)
	fingerprint := make([]byte, 64)
	crand.Read(fingerprint)

	d.Fingerprint = hex.EncodeToString(fingerprint)
}
