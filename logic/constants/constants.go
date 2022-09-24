package constants

const (
	FfmpegJSON  = "ffprobe -v quiet -print_format json -show_format -show_streams %s"
	FfmpegCatch = "ffmpeg -ss 00:00:%s -v quiet -i %s %s.jpg -r 1 -vframes 1 -an -vcodec mjpeg"
)
