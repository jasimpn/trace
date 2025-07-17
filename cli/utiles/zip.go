package utiles

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func ZipTo(output string, inputs []string) error {
	outFile, _ := os.Create(output)
	defer outFile.Close()
	zipWriter := zip.NewWriter(outFile)
	defer zipWriter.Close()

	for _, input := range inputs {
		filepath.Walk(input, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			file, _ := os.Open(path)
			defer file.Close()

			relPath, _ := filepath.Rel(filepath.Dir(input), path)
			writer, _ := zipWriter.Create(relPath)
			_, err = io.Copy(writer, file)
			return err
		})
	}
	return nil
}
