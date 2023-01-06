package main

import (
	"os/exec"

	lmmp3 "github.com/paij0se/lmmp3"
)

func main() {
	// download the video and convert it to a mp3 file
	// output: 'Vine Boom Sound Effect.mp3'
	lmmp3.DownloadAndConvert("https://www.youtube.com/watch?v=829pvBHyG6I")
	// delete the original mpeg file (only for windows)

	del := exec.Command("cmd", "/C", "del", "*.mpeg")
	if del.Run() != nil {
		panic(any("failed to delete files"))
	}
}
