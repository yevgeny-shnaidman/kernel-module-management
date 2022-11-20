package utilexecute

import (
	bytes2 "bytes"
	"os/exec"
)

const TimeoutExitCode = 124

func getExitCode(err error) int {
	if err == nil {
		return 0
	}
	switch value := err.(type) {
	case *exec.ExitError:
		return value.ExitCode()
	default:
		return -1
	}
}

func getErrorStr(err error, stderr *bytes2.Buffer) string {
	b := stderr.Bytes()
	if len(b) > 0 {
		return string(b)
	} else if err != nil {
		return err.Error()
	}
	return ""
}

func Execute(command string, args ...string) (stdout string, stderr string, exitCode int) {
	cmd := exec.Command(command, args...)
	var stdoutBytes, stderrBytes bytes2.Buffer
	cmd.Stdout = &stdoutBytes
	cmd.Stderr = &stderrBytes
	err := cmd.Run()
	return stdoutBytes.String(), getErrorStr(err, &stderrBytes), getExitCode(err)
}

func ExecutePrivileged(command string, args ...string) (stdout string, stderr string, exitCode int) {
	// nsenter is used here to launch processes inside the container in a way that makes said processes feel
	// and behave as if they're running on the host directly rather than inside the container
	commandBase := "nsenter"

	arguments := []string{
		"--target", "1",
		// Entering the cgroup namespace is not required for podman on CoreOS (where the
		// agent typically runs), but it's needed on some Fedora versions and
		// some other systemd based systems. Those systems are used to run dry-mode
		// agents for load testing. If this flag is not used, Podman will sometimes
		// have trouble creating a systemd cgroup slice for new containers.
		"--cgroup",
		// The mount namespace is required for podman to access the host's container
		// storage
		"--mount",
		// TODO: Document why we need the IPC namespace
		"--ipc",
		// Network namespace is needed for accessing host networking information
		// during inventory collection
		"--net",
		"--",
		command,
	}

	arguments = append(arguments, args...)
	return Execute(commandBase, arguments...)
}
