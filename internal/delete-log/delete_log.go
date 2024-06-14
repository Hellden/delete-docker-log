package deletelog

import (
	"os"

	log "github.com/hellden/delete-docker-log/internal/logging"
)

func DeleteLog(LogPath string) {
	e := os.Remove(LogPath)
	if e != nil {
		log.Error.Fatal(e)
	}
	log.Info.Println("Remove file is ok !")
}
