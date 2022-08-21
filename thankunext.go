package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-rod/rod"
)

func handleError(err error) {

	var evalErr *rod.ErrEval

	if errors.Is(err, context.DeadlineExceeded) { // timeout error
		fmt.Println("timeout err")
	} else if errors.As(err, &evalErr) { // eval error
		fmt.Println(evalErr.LineNumber)
	} else if err != nil {
		fmt.Println("can't handle", err)
	}
}

func dumpBuildManifestData(url string) string {
	val := ""

	type buildManifestStruct struct {
		SortedPages []string `json:sortedPages`
	}

	if strings.HasSuffix(url, "buildManifest.js") {
		page := rod.New()
		err := rod.Try(func() {
			loaded_page := page.Timeout(10 * time.Second).MustConnect().MustPage(url).MustWaitLoad()
			loaded_page.MustEval("() => eval(document.documentElement.innerText)")
			val = loaded_page.MustEval("() => JSON.stringify(self.__BUILD_MANIFEST)").Str()

		})

		if err != nil {
			return ""
		}

	} else {
		page := rod.New()
		err := rod.Try(func() {
			loaded_page := page.Timeout(10 * time.Second).MustConnect().MustPage(url).MustWaitLoad()
			val = loaded_page.MustEval("() => JSON.stringify(self.__BUILD_MANIFEST)").Str()
		})

		if err != nil {
			return ""
		}
	}

	if val != "<nil>" {
		var buildManifestUnmarshal buildManifestStruct
		if err := json.Unmarshal([]byte(val), &buildManifestUnmarshal); err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			return ""
		}
		sortedPages := buildManifestUnmarshal.SortedPages

		for _, element := range sortedPages {
			fmt.Println(element)
		}
		return val
	} else {
		return ""
	}

}

func main() {
	fmt.Fprintln(os.Stderr, `<!-- thankunext v0.01, made by @c3l3si4n -->`)
	if len(os.Args[1:]) > 0 {

		dumpBuildManifestData(os.Args[1])
	} else {
		fmt.Printf("[thankunext] %s <_buildManifest.js url>", os.Args[0])
	}
	//buildManifestUrl := os.Args
}
