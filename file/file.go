package file

import (
    "os"
    "path"
    "strings"
    "log"
)

func FileInDir(dir, fname string) string {
	var fileBase string = strings.ToLower(path.Base(fname))

    dirEntry, err := os.ReadDir(dir)
	if err != nil {
        log.Println("Error:", err)
		return ""
	}

    for _, f := range dirEntry {
		if fileBase == strings.ToLower(f.Name()) {
			// case insensitive
			return f.Name()
		}
	}

    return ""
}

// fname/path with extension removed
func FileExtensionRemove(fname string) string {
	return strings.TrimSuffix(fname, path.Ext(fname))
}

func FileSimplifyName(fname string) string {
	fname = FileExtensionRemove(fname)
	fname = strings.ToLower(fname)
	fname = strings.ReplaceAll(fname, "-", "")
	fname = strings.ReplaceAll(fname, "_", "")
	return fname
}
