package main

import (

	"encoding/json"
	"os"
	//"strconv"
	

	log "github.com/sirupsen/logrus"

)



/* ----------------------------
Structure to hold the contents of config.json
---------------------------- */
type Settings struct {

	Locations struct {
		// remote server - copy settings
		Raw_dir         string `json:"raw_dir"`          // where to look for files appearing
		Processed_dir   string `json:"processed_dir"`    // where to put the processed image files (directory tree)
		Quarantine_dir  string `json:"quarantine_dir"`   // where to put the files that are not images (subdir by date)
		Log_dir         string `json:"log_dir"` 	     // where to place the log files (managed to the dates)
		Archive_dir     string `json:"archive_dir"`      // where to put the files that we processed, as backup (subdir by type)

		//Loglevel    log.Level `json:"loglevel"`
	}
	
	Options struct {
		Depth string `json:"depth"`     				// For the processed files, the depth that we will go in naming
		Chars string `json:"chars"`            			// number of characters for the file name for each level of directory
		Cleanup string `json:"cleanup"`                 // number of days of quarantine directories before we delete those contents
		ImageWidth string `json:"imagewidth"`           // width of the compressed image
		ImageHeight string `json:"imageheight"`          // height of the compressed image
	}

}

/* ----------------------------
Go does not give us a "standard" way to read config files.  Instead, we are encouraged to use
user contributed libraries.  Unfortunately, those libraries can be buggy.  Viper, a popular library
for config files, failed on me.  So I switched to JSON  -- Captain Nick Silver, 3/7/21
---------------------------- */

func ReadConfig() Settings {
	var s Settings
	var configfilename string
	

	// Set the path to look for the configurations file
	mycwd, _ := os.Getwd()
	configfilename = mycwd +  "\\config.json"

	file, err := os.Open(configfilename)
	if err != nil {
		log.Fatal("Error opening config file", configfilename, err.Error())
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&s)
	if err != nil {
		log.Fatal("can't decode config JSON: ", err)
	}
	//log.SetLevel(s.Locations.Loglevel)


	return s
}
