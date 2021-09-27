package main

import (
//	"bufio"
//	"bytes"
//	"encoding/csv"
	"fmt"
//	"log"
	"math"
//	"math/rand"
//	"net/url"
//	"os"
//	"regexp"
	"strconv"
	"strings"
//	"time"
//	"os/exec"

 "github.com/fsnotify/fsnotify"
//	"golang.org/x/text/language"
//	"golang.org/x/text/message"
)


var s Settings // global config file contents


/*       ----
* Image Prep -- this app will be running as a service on a Linux machine, so naturally I'm writing it on Windows.  (I'm a masochist)
* Business problem: we are creating a social media site where people can upload images and short videos (Tiktok and Instagram style).  
*        Of course, there's no guarantee that it is actually a video.  It could be a virus or malware.  Or, more likely, a REALLY HUGE
*        video or image.  We are going to serve it out to a million people.  I'm not sending out a 10 GB photo of your dog two million times. 
* Surrounding situation:  We will implement a simple upload capability that will drop files in a "raw" directory.  That's where this app 
*         comes in.  
* Functionality: This app will read the file, check to make sure it is actually a graphic file, and if so, the app will resize
*         images to a specific resolution (probably a lot smaller than the source).  (If the image is already small, don't process it at all).
*         the processed file will be written to the 'processed' directory.  
*         The original file will be moved to the archive directory.
*         Any file that is not a graphic image will be moved to the quarantine directory. 
*         Once the file is moved, we will call a routine that will record the new file in a database.  That part is in another module that 
*         will be private.  I'm sharing this bit in open source.  The rest... not so much.
* Notes:
*    - we will clean up after ourselves.  The quarantine directory and the archive directory will be a directory tree by date.  When we create a new
*      directory (because it's a new date), we will delete any directories that are older than two weeks.  
*    - The processed directory will be a tree, with the first two letters of each filename being the directory name in the tree, five levels deep. 
*    - we log everything in a log file, by date, in the log directory .  Log files older than two weeks will be deleted. 
*
* Design notes:
*    - I'll use multithreading capabilities of go to monitor the directory.  When a file appears, I'll capture information about the file and insert
*      it into a queue.  The main thread just reads the queue and opens the file, compresses it, and writes it to the right location (or moves it)
*      to quarantine. 
*    - when we create a directory for a date, we create all the directories for that date in both quarantine and archive.  We will immediately look
*      for subdirectories that are too old to keep.  Directory names will be YYYYMMDD format so that we convert them to big integers and compare them
*      easily and efficiently. 
*/

// we are going to create a list of users and then, for each user, analyze their behavior

//this is used to enqueue the images uploaded to the site
type FileQueueItem struct {
	sourcename string
	targetname string      // the target filename.  This is always either quarantine or archive.  This is not the processed name.
	processedname string   // the name of the processed version of this file
	qt_flag    bool        // set to true if the file should be quarantined
	imagefmt   string      // format of the image file.  Valid: PNG, JPG, GIF, or NOT (for not an image)
}



func main() {
	var filequeue []FileQueueItem // when a file is detected, it is investigated and added to the queue

	s = ReadConfig()

	MonitorFilesystem(s)

	for {
		GetFileEvents(filequeue)
		for _,fqi := range filequeue {
			fmt.Println( fqi.sourcename , " found")
		}
	}
}



// set up the file system monitoring based on configuration
func MonitorFilesystem(s Settings) {

}

func GetFileEvents(filequeue []FileQueueItem){

}


