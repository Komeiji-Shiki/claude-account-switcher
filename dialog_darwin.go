package main

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// The titles come in as arguments, so they need no AppleScript quoting.
const resumeScript = `on run argv
	set chatTitles to items 5 thru -1 of argv
	set r to choose from list chatTitles with title (item 1 of argv) with prompt (item 2 of argv) default items chatTitles OK button name (item 3 of argv) cancel button name (item 4 of argv) with multiple selections allowed
	if r is false then return ""
	set out to ""
	repeat with i in r
		set out to out & i & linefeed
	end repeat
	return out
end run`

// The title, prompt and value come in as arguments, like the resume titles.
const numberScript = `on run argv
	set cancelLabel to item 4 of argv
	set confirmLabel to item 5 of argv
	set r to display dialog (item 2 of argv) with title (item 1 of argv) default answer (item 3 of argv) buttons {cancelLabel, confirmLabel} default button confirmLabel cancel button cancelLabel
	return text returned of r
end run`

// nativeAskNumber asks for a number under prompt, with value filled in, and
// returns what was typed. It is false when the dialog was cancelled or could
// not open.
func nativeAskNumber(l language, title, prompt, value string) (string, bool) {
	// Cancel makes osascript fail with "User canceled".
	out, err := exec.Command("osascript", "-e", numberScript, title, prompt, value, l.text("Cancel"), l.text("Set")).Output()
	if err != nil {
		return "", false
	}
	return strings.TrimSpace(string(out)), true
}

// nativeResume asks which of the chats titled titles should go on and
// returns one answer per title. After wait without an answer, or if the
// dialog cannot open, all of them go on.
func nativeResume(l language, titles []string, wait time.Duration) []bool {
	keep := make([]bool, len(titles))
	for i := range keep {
		keep[i] = true
	}
	// Numbered, so two chats with the same title stay apart.
	args := []string{"-e", resumeScript, l.text("Continue chats?"),
		l.text("These chats were working when Claude closed for the switch. Claude continues the selected ones when it starts again."),
		l.text("Continue"), l.text("Don't continue")}
	for i, t := range titles {
		args = append(args, fmt.Sprintf("%d. %s", i+1, t))
	}
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	out, err := exec.CommandContext(ctx, "osascript", args...).Output()
	var exit *exec.ExitError
	switch {
	case ctx.Err() != nil:
		return keep
	case errors.As(err, &exit):
		// Don't continue makes osascript fail with "User canceled".
		return make([]bool, len(titles))
	case err != nil:
		return keep
	}
	keep = make([]bool, len(titles))
	for _, line := range strings.Split(string(out), "\n") {
		n, _, _ := strings.Cut(line, ". ")
		if i, err := strconv.Atoi(n); err == nil && i >= 1 && i <= len(titles) {
			keep[i-1] = true
		}
	}
	return keep
}
