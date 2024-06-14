package deletelog

import (
	"os"
	"strings"

	log "github.com/hellden/delete-docker-log/internal/logging"
)

func DeleteLog(LogPath string) {
	splitLogPath := strings.Trim(LogPath, "'")
	e := os.Remove(splitLogPath)

	if e != nil {
		log.Error.Fatal(e)
		return
	}
	log.Info.Println("Remove file is ok !")
}
