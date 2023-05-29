package androidgo

var DEVICES = []Device{
	{ // Android Simulator (Pixel 6) Android 12
		DisplayMetrics: DisplayMetrics{
			HeightPixels: 2274,
			WidthPixels:  1080,
		},
		Build: Build{
			BOARD:          "goldfish_x86_64",
			BOOTLOADER:     "unknown",
			BRAND:          "google",
			DEVICE:         "emulator64_x86_64_arm64",
			DISPLAY:        "", /* Generated */
			FINGERPRINT:    "", /* Generated */
			HARDWARE:       "ranchu",
			HOST:           "abfarm-release-2004-0065",
			ID:             "SE1A.220826.006",
			MANUFACTURER:   "Google",
			MODEL:          "sdk_gphone64_x86_64",
			PRODUCT:        "sdk_gphone64_x86_64",
			RADIO:          "1.0.0.0",
			SUPPORTED_ABIS: []string{"x86", "armeabi-v7a", "armeabi"},
			TAGS:           "dev-keys",
			TYPE:           "userdebug",
			USER:           "android-build",
			VERSION: Version{
				CODENAME:    "REL",
				INCREMENTAL: "9534912",
				RELEASE:     "12",
				SDK_INT:     31,
			},
		},
	},
}
