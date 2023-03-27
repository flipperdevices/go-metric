//go:build ignore

//go:generate go run generate.go ../proto-src

package main

import (
	"errors"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var packageRe = regexp.MustCompile(`(?m)package metric\.(.*?);`)
var protofileRe = regexp.MustCompile(`.*\.proto$`)

const baseImportPath = "github.com/flipperdevices/go-metric/internal/proto"

func main() {
	for _, path := range os.Args[1:] {
		args := []string{
			"--go_out=module=" + filepath.Dir(baseImportPath) + ":..",
		}
		var protoFiles []string

		err := filepath.WalkDir(path, func(fp string, d fs.DirEntry, err error) error {
			if d.IsDir() {
				args = append(args, "--proto_path", fp)
			} else if protofileRe.MatchString(d.Name()) {
				p, err := composeImportPath(fp)
				if err != nil {
					log.Fatalln("Can't compose import path", fp, err)
				}

				protoFiles = append(protoFiles, fp)
				rp, err := filepath.Rel(path, fp)
				if err != nil {
					log.Fatalln(err)
				}
				args = append(args, "--go_opt=M"+rp+"="+baseImportPath+p)
			}
			return nil
		})
		if err != nil {
			log.Fatalln(err)
		}
		args = append(args, protoFiles...)
		log.Print(args)
		cmd := exec.Command("protoc", args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Fatal("Failed generating", path, err)
		}
	}
}

func composeImportPath(file string) (string, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	matches := packageRe.FindStringSubmatch(string(b))
	switch len(matches) {
	case 2:
		return "/" + strings.ToLower(matches[1]), nil
	case 0:
		return "", nil
	}
	return "", errors.New("regex failed")
}
