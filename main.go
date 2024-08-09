package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	err := filepath.Walk("C:/Users/Ricardo/Videos/compress", func(path string, info os.FileInfo, err error) error {
		var input, output string

		if err != nil {
			fmt.Println(err)
			return err
		}

		if !info.IsDir() {
			fmt.Printf("Compress filename: %s\n", info.Name())

			input = "C:/Users/Ricardo/Videos/compress/" + info.Name()
			output = "C:/Users/Ricardo/Videos/compressed/" + info.Name()

			cmd := exec.Command("ffmpeg",
				"-i",
				input,
				"-c:v",
				"libx264",
				"-pix_fmt",
				"yuv420p",
				output,
			)

			// pipe the commands output to the applications
			// standard output
			cmd.Stdout = os.Stdout

			fmt.Println("Run command: " + cmd.String())

			// Run still runs the command and waits for completion
			// but the output is instantly piped to Stdout
			if err := cmd.Run(); err != nil {
				fmt.Println("could not run command: ", err)
			}

		}

		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
