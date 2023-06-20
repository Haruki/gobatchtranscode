package main

import (
	"log"
	"os"
)

func main() {
	sourceDir := os.Args[1]
	// targetDir := os.Args[2]
	file, err := os.Open(sourceDir)
	if err != nil {
		log.Panic("cannot open file")
	}
	filnames, _ := file.Readdirnames(0)
	for _, file := range filnames {
		log.Printf("File detected: %s", file)
	}

}

//ffmpeg -i input.mp4 -c:v libx265 -vtag hvc1 output.mp4

/*
args := "-i test.mp4 -acodec copy -vcodec copy -f flv rtmp://aaa/bbb"
    cmd := exec.Command("ffmpeg", strings.Split(args, " ")...)

    stderr, _ := cmd.StderrPipe()
    cmd.Start()

    scanner := bufio.NewScanner(stderr)
    scanner.Split(bufio.ScanWords)
    for scanner.Scan() {
        m := scanner.Text()
        fmt.Println(m)
    }
    cmd.Wait()
*/
