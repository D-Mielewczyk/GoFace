package utils

import "path/filepath"

func ConvertPath(input_path string, output_dir string) string {
	fileName := filepath.Base(input_path)
	return filepath.Join(output_dir, fileName)
}
