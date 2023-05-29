package androidgo

type Version struct {
	CODENAME    string
	INCREMENTAL string
	RELEASE     string
	SDK_INT     int
}

var VERSIONS = []Version{
	{ // Android 12
		CODENAME:    "REL",
		INCREMENTAL: "9612669",
		RELEASE:     "12",
		SDK_INT:     31,
	},
	{ // Android 11
		CODENAME:    "REL",
		INCREMENTAL: "7380771",
		RELEASE:     "11",
		SDK_INT:     30,
	},
	{ // Android 10
		CODENAME:    "REL",
		INCREMENTAL: "6578210",
		RELEASE:     "10",
		SDK_INT:     29,
	},
	{ // Android 9
		CODENAME:    "REL",
		INCREMENTAL: "5081123",
		RELEASE:     "9",
		SDK_INT:     28,
	},
	{ // Android 8.1
		CODENAME:    "REL",
		INCREMENTAL: "4458749",
		RELEASE:     "8.1",
		SDK_INT:     27,
	},
	{ // Android 8.0
		CODENAME:    "REL",
		INCREMENTAL: "5903476",
		RELEASE:     "8.0",
		SDK_INT:     26,
	},
	{ // Android 7.1
		CODENAME:    "REL",
		INCREMENTAL: "5903476",
		RELEASE:     "7.1",
		SDK_INT:     25,
	},
	{ // Android 7.0
		CODENAME:    "REL",
		INCREMENTAL: "5903476",
		RELEASE:     "7.0",
		SDK_INT:     24,
	},
}
