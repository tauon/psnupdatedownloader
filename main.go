package main

import (
	"crypto/tls"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type Package struct {
	URL string `xml:"url,attr"`
}

type Tag struct {
	Packages []Package `xml:"package"`
}

type TitlePatch struct {
	Tags []Tag `xml:"tag"`
}

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Println("Usage: psn_downloader <gameID>")
		os.Exit(1)
	}

	gameID := flag.Arg(0)

	titlePatch, err := fetchAndParseXML(gameID)
	if err != nil {
		fmt.Printf("Error fetching XML: %v\n", err)
		os.Exit(1)
	}

	destDir := gameID
	if err := os.MkdirAll(destDir, 0755); err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		os.Exit(1)
	}

	for _, tag := range titlePatch.Tags {
		for _, pkg := range tag.Packages {
			fmt.Printf("Downloading: %s\n", pkg.URL)
			if err := downloadPackage(pkg.URL, destDir); err != nil {
				fmt.Printf("Failed to download package: %v\n", err)
			}
		}
	}

	fmt.Println("Download completed.")
}

func fetchAndParseXML(gameID string) (*TitlePatch, error) {
	// Disable cert verification due to playstation.net TLS error
	customTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Transport: customTransport}

	url := fmt.Sprintf("https://a0.ww.np.dl.playstation.net/tpl/np/%s/%s-ver.xml", gameID, gameID)
	response, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var titlePatch TitlePatch
	err = xml.Unmarshal(body, &titlePatch)
	if err != nil {
		return nil, err
	}

	return &titlePatch, nil
}

func downloadPackage(url, destDir string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath.Join(destDir, filepath.Base(url)))
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
