package main

// import (
// 	"fmt"
// 	"os/exec"
// 	"runtime"
// )

// // openBrowser tries to open the URL in a browser and returns an error if it fails.
// func openBrowser(url string) error {
// 	var cmd string
// 	var args []string

// 	switch runtime.GOOS {
// 	case "windows":
// 		cmd = "rundll32"
// 		args = []string{"url.dll,FileProtocolHandler", url}
// 	case "darwin":
// 		cmd = "open"
// 		args = []string{url}
// 	case "linux":
// 		cmd = "xdg-open"
// 		args = []string{url}
// 	default:
// 		return fmt.Errorf("unsupported platform")
// 	}

// 	return exec.Command(cmd, args...).Start()
// }

// func main() {
// 	url := "http://example.com"
// 	if err := openBrowser(url); err != nil {
// 		fmt.Printf("Failed to open browser: %v\n", err)
// 	} else {
// 		fmt.Println("Browser opened successfully")
// 	}
// }
