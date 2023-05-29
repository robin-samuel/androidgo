package androidgo

import "math/rand"

type Build struct {
	BOARD          string
	BOOTLOADER     string
	BRAND          string
	DEVICE         string
	DISPLAY        string
	FINGERPRINT    string
	HARDWARE       string
	HOST           string
	ID             string
	MANUFACTURER   string
	MODEL          string
	PRODUCT        string
	RADIO          string
	SUPPORTED_ABIS []string
	TAGS           string
	TIME           int64
	TYPE           string
	USER           string
	VERSION        Version
}

func (b *Build) bDISPLAY() string {
	b.DISPLAY = b.PRODUCT + "-" + b.TYPE + " " + b.VERSION.RELEASE + " " + b.ID + " " + b.VERSION.INCREMENTAL + " " + b.TAGS
	return b.DISPLAY
}

func (b *Build) bFINGERPRINT() string {
	b.FINGERPRINT = b.BRAND + "/" + b.PRODUCT + "/" + b.DEVICE + ":" + b.VERSION.RELEASE + "/" + b.ID + "/" + b.VERSION.INCREMENTAL + ":" + b.TYPE + "/" + b.TAGS
	return b.FINGERPRINT
}

func (b *Build) bTIME() int64 {
	// Generate a random time between 2019-01-01 and 2021-01-01
	b.TIME = 1546300800 + rand.Int63n(631152000)
	return b.TIME
}
