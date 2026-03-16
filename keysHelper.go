package main

const (
	firstPrintableASCII = 32
	lastPrintableASCII  = 126
)

var keysMap map[string]string = map[string]string{
	"3":  "KeyCtrlC",
	"13": "KeyEnter",
	"19": "KeyCtrlS",

	"127":          "KeyBackspace",
	"27-91-51-126": "KeyDelete",

	"27-91-65": "KeyUp",
	"27-91-66": "KeyDown",
	"27-91-67": "KeyRight",
	"27-91-68": "KeyLeft",

	"27-91-70": "KeyEnd",
	"27-91-72": "KeyHome",
}

func DetermineKey(keyBytes []byte) string {
	if len(keyBytes) == 0 {
		return ""
	}

	if len(keyBytes) == 1 {
		b := keyBytes[0]
		if firstPrintableASCII <= b && b <= lastPrintableASCII {
			return string(b)
		}
	}

	joinedBytes := JoinBytes(keyBytes, "-")
	return keysMap[joinedBytes]
}
