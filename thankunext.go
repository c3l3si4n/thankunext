package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func contains(elements []string, element string) bool {
	for _, v := range elements {
		if v == element {
			return true
		}
	}

	return false
}

func getBuildManifestPath(bodyContent string) string {
	re, _ := regexp.Compile(`(?m)/_next/static/[\w-]+/_buildManifest\.js`)

	return re.FindString(bodyContent)
}

func getPageContent(url string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error on NewRequest.", err)
		os.Exit(1)
	}

	req.Header.Add("user-agent", "thankunext/1.0")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error when accessing the url.", err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error when loading page content.", err)
		os.Exit(1)
	}

	return string(body)
}

func getBuildManifestContent(buildManifestPath string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", buildManifestPath, nil)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error on NewRequest.", err)
		os.Exit(1)
	}

	req.Header.Add("user-agent", "thankunext/1.0")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error accessing buildManifest content.", err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error when loading buildManifest content.", err)
		os.Exit(1)
	}

	return string(body)

}

func parseBuildManifestContent(buildManifestContent string) []string {
	re, _ := regexp.Compile(`"(/[a-zA-Z0-9_/\[\]\.-]+)"`)
	var paths []string

	for _, match := range re.FindAllStringSubmatch(buildManifestContent, -1) {
		path := match[1]

		if contains(paths, path) {
			continue
		}

		paths = append(paths, path)
	}

	return paths

}

func main() {
	fmt.Fprintln(os.Stderr, `<!-- thankunext v0.01, made by @c3l3si4n -->`)

	if len(os.Args[1:]) > 0 {
		target := strings.TrimSuffix(os.Args[1], "/")
		pageContent := getPageContent(target)
		buildManifestPath := getBuildManifestPath(pageContent)
		buildManifestContent := getBuildManifestContent(target + buildManifestPath)
		paths := parseBuildManifestContent(buildManifestContent)

		fmt.Println(strings.Join(paths, "\n"))

	} else {
		fmt.Printf("[thankunext] %s <url>", os.Args[0])
	}
}
