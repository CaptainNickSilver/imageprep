package main

import (
	"os"
	"io/ioutil"

	log "github.com/sirupsen/logrus"

	"errors"
)

/* ------------------------------------
   CheckFileExists (full path) - bool
      Returns true if the file exists.
      Note that it only return 'false' if we are certain that the file does not exist

	  golang does not actually have a method to determine if a file actually exists
------------------------------------ */

func CheckFileExists(fname string) bool {

	if _, err := os.Stat(fname); err == nil {
		log.Debug("file exists:", fname)
		return true

	} else if errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does *not* exist
		log.Debug("file does not exist on working machine:", fname)
		return false
	} else {
		log.Debug("file may exist, assuming it does:", fname)
		return true
	}

}

func GetFileSize(fname string) int64 {
	if finfo, err := os.Stat(fname); err == nil {
		log.Debug("file exists:", fname)
		return finfo.Size()

	} else if errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does *not* exist
		log.Debug("file should exist but does not exist on working machine:", fname)
		return int64(0)
	} else {
		log.Debug("file may exist but cannot pull stats, assuming a zero file size:", fname)
		return int64(0)
	}

}
func DeleteFile(fname string) {
	if _, err := os.Stat(fname); err == nil {
		err = os.Remove(fname)
		if err == nil {
			log.Debug("deleting file:", fname)
		} else {
			log.Error("error deleting file:", fname, err)
		}
		return

	} else if errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does *not* exist
		log.Debug("Cannot delete file that does not exist:", fname)
		return
	} else {
		log.Debug("Error attempting to stat a file, choosing not to delete:", fname)
		return
	}

}

func CreateFile(contents string, fname string) bool {
	d1 := []byte(contents)
	
    err := ioutil.WriteFile(fname, d1, 0644)
    if err != nil {
		return false
	}
	return true 
}
