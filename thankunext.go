package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-rod/rod"
)

func dumpBuildManifestData(url string) string {
	if strings.HasSuffix(url, "buildManifest.js") {
		page := rod.New().MustConnect().MustPage(url)
		page.MustWaitLoad().MustEval("() => {eval(document.documentElement.innerText)}")
		val := page.MustEval("() => JSON.stringify(self.__BUILD_MANIFEST)").Str()
		return val
	} else {
		page := rod.New().MustConnect().MustPage(url)
		val := page.MustWaitLoad().MustEval("() => JSON.stringify(self.__BUILD_MANIFEST)").Str()
		return val
	}

}

func main() {
	fmt.Fprintln(os.Stderr, `<!-- thankunext v0.01, made by @c3l3si4n -->`)
	if len(os.Args[1:]) > 0 {
		fmt.Println(dumpBuildManifestData(os.Args[1]))
	} else {
		fmt.Printf("[thankunext] %s <_buildManifest.js url>", os.Args[0])
	}
	//buildManifestUrl := os.Args
}
