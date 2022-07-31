package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/stashapp/stash/pkg/ffmpeg"
	"github.com/stashapp/stash/pkg/hash/videophash"
)

type Instance struct {
	ffprobe ffmpeg.FFProbe
	ffmpeg  ffmpeg.FFMpeg
}

func main() {
	var (
		instance = Instance{
			ffprobe: ffmpeg.FFProbe("ffprobe"),
			ffmpeg:  ffmpeg.FFMpeg("ffmpeg"),
		}
		filePath string
	)

	flag.StringVar(&filePath, "f", "", "Path to video file")
	flag.Parse()

	if len(strings.TrimSpace(filePath)) == 0 {
		flag.PrintDefaults()
		os.Exit(0)
	}

	videoFile, err := instance.ffprobe.NewVideoFile(filePath)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	phash, err := videophash.Generate(instance.ffmpeg, videoFile)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(2) //nolint:gomnd
	}

	phashString := strconv.FormatUint(*phash, 16) //nolint:gomnd
	_, _ = fmt.Fprintln(os.Stdout, phashString)
}
