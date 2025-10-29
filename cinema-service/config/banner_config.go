package config

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func ClearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	case "linux", "darwin":
		cmd := exec.Command("/usr/bin/clear")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("OS not detected")
	}
}

func getOSVersion() string {
	switch runtime.GOOS {
	case "linux":
		if out, err := os.ReadFile("/proc/version"); err == nil {
			return strings.Fields(string(out))[0] + " " + strings.Fields(string(out))[2]
		}
	case "windows":
		if out, err := exec.Command("cmd", "/C", "ver").Output(); err == nil {
			return strings.TrimSpace(string(out))
		}
	case "darwin":
		if out, err := exec.Command("sw_vers", "-productVersion").Output(); err == nil {
			return "macOS " + strings.TrimSpace(string(out))
		}
	}
	return "Unknown OS"
}

func getTerminalWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 80
	}
	parts := strings.Fields(string(out))
	if len(parts) < 2 {
		return 80
	}
	width, err := strconv.Atoi(parts[1])
	if err != nil {
		return 80
	}
	return width
}

func PrintLine() {
	width := getTerminalWidth()
	line := strings.Repeat("─", width)
	fmt.Println(line)
}

func Banner() {
	var appconfig, errs = GetServerConfig()
	if errs != nil {
		log.Fatal(errs)
	}

	ClearScreen()

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Print(`
    __    __  __       ___ 
   / /_  / /_/ /_____ |__ \
  / __ \/ __/ __/ __ \__/ /
 / / / / /_/ /_/ /_/ / __/ 
/_/ /_/\__/\__/ .___/____/ 
             /_/            
`)

	fmt.Printf(
		"App Name     : %s\nHost Server  : %s\nPID          : %d\nRuntime      : %s\nStartup Time : %s\nSys Memory   : %.2f MB\nAlloc Memory : %.2f MB\n",
		appconfig.AppName,
		getOSVersion(),
		os.Getpid(),
		runtime.Version(),
		time.Now().Format("2006-01-02 15:04:05"),
		float64(m.Sys)/1024/1024,
		float64(m.Alloc)/1024/1024,
	)
}
