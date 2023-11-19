package utils

import "os"

func RemoveFromSlice(slice []os.DirEntry, item os.DirEntry) []os.DirEntry {
	for i, v := range slice {
		if v.Name() == item.Name() {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
