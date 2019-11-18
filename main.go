package main

import (
	"flag"
	"log"
	"time"

	"github.com/muzavan/github-contrib-write/util"
)

func main() {
	folderFlag := flag.String("folder", "./gen", "Folder where generated file/commits stored")
	dateFlag := flag.String("date", time.Now().Format(util.DateFormat), "Current date in yyyy-mm-dd format using now as default")
	nameFlag := flag.String("name", "Gen Bot", "Name of the committer")
	emailFlag := flag.String("email", "genbot@genbot.com", "Email of the committer")
	textFlag := flag.String("textFile", "file.txt", "Path to file with 1s to define your github contrib panel style")

	flag.Parse()
	points, err := util.ReadFromFile(*textFlag)

	if err != nil {
		log.Fatalf("Can't read the panel file. Reason: %s", err)
	}

	if err = util.GitInit(*folderFlag); err != nil {
		log.Fatalf("Can't execute git init in folder. Reason: %s", err)
	}

	if err = util.GitConfig(*folderFlag, "user.name", *nameFlag); err != nil {
		log.Fatalf("Can't set git config for user.name. Reason: %s", err)
	}

	if err = util.GitConfig(*folderFlag, "user.email", *emailFlag); err != nil {
		log.Fatalf("Can't set git config for user.email. Reason: %s", err)
	}

	currDate, err := time.Parse("2006-01-02", *dateFlag)
	if err != nil {
		log.Fatalf("Can't parse date. Reason: %s", err)
	}

	for p := range points {
		x, y := points[p].X, points[p].Y
		date := util.PanelToDate(currDate, x, y)
		_, err := util.GenerateFile(*folderFlag, p)

		if err != nil {
			log.Printf("Can't generate file for panel %d, %d (%s). Reason: %s", x, y, date.Format(util.DateFormat), err)
			continue
		}

		// TODO: For now it's safe to use ".", should used generated file name instead
		if err = util.GitAdd(*folderFlag, "."); err != nil {
			log.Printf("Can't add file for panel %d, %d (%s). Reason: %s", x, y, date.Format(util.DateFormat), err)
			continue
		}

		if err = util.GitCommit(*folderFlag, date); err != nil {
			log.Printf("Can't commit file for panel %d, %d (%s). Reason: %s", x, y, date.Format(util.DateFormat), err)
		}
	}

	log.Printf("It's done. Please review it and don't forget to push your code.")
}
