//go:build darwin && !linux && !freebsd && !netbsd && !openbsd && !windows && !js
// +build darwin,!linux,!freebsd,!netbsd,!openbsd,!windows,!js

package beeep

import "os/exec"

// Alert displays a desktop notification and plays a default system sound.
func Alert(title, message, appIcon string) error {
	tn, err := exec.LookPath("terminal-notifier")
	if err == nil {
		cmd := exec.Command(tn, "-title", title, "-message", message, "-sound", "default", "-appIcon", appIcon)
		return cmd.Run()
	}

	osa, err := exec.LookPath("osascript")
	if err != nil {
		return err
	}

	cmd := exec.Command(osa, "-e", `tell application "System Events" to display notification "`+message+`" with title "`+title+`" sound name "default"`)
	return cmd.Run()
}
