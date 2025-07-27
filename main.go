package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"ziphasher/hasher"
	"ziphasher/zip"
)

var (
	zipPath = flag.String("zip", "", "Path to the ZIP file")
	algo    = flag.String("algo", "sha256", "Hash algorithm: md5, sha1, sha256, sha512")
	jsonOut = flag.Bool("json", false, "Enable JSON output")
)

func main() {
	flag.Parse()

	if *zipPath == "" {
		log.Fatal("❌ Please provide a ZIP file path using -zip")
	}

	files, err := zip.ExtractZipToTemp(*zipPath)
	if err != nil {
		log.Fatalf("❌ Error while extracting ZIP file: %v", err)
	}

	result := make(map[string]string)

	for _, file := range files {
		hash, err := hasher.HashFile(file, *algo)
		if err == nil {
			result[file] = hash
		}
	}

	if *jsonOut {
		json.NewEncoder(log.Writer()).Encode(result)
	} else {
		for f, h := range result {
			fmt.Printf("%s: %s\n", f, h)
		}
	}
}