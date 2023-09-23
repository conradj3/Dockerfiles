package main

import (
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/require"
)

const (
	failedToGetExitCode       = "Failed to get exit code: %v"
	expectedError             = "Expected error but got nil"
	expectedExitCode          = "Expected exit code 1 but got %d"
	dockerConnectionError     = "could not connect to Docker"
	dockerImage               = "conradj3/nginx-distroless-unprivileged"
	dockerContainerStartError = "could not start container"
	dockerPoolPurgeError      = "failed to remove container"
)

func TestNginxDefaultHtml(t *testing.T) {

	pool, err := dockertest.NewPool("")

	require.NoError(t, err, dockerConnectionError)

	resource, err := pool.Run(dockerImage, "latest", []string{})
	require.NoError(t, err, dockerContainerStartError)

	t.Cleanup(func() {
		require.NoError(t, pool.Purge(resource), dockerPoolPurgeError)
	})

	var resp *http.Response

	err = pool.Retry(func() error {
		resp, err = http.Get(fmt.Sprint("http://localhost:", resource.GetPort("80/tcp"), "/"))
		if err != nil {
			t.Log("container not ready, waiting...")
			return err
		}
		return nil
	})
	require.NoError(t, err, "HTTP error")
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode, "HTTP status code")

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "failed to read HTTP body")

	require.Contains(t, string(body), "Welcome to nginx!", "Default HTML not found")
}

func TestDockerExecSh(t *testing.T) {
	pool, err := dockertest.NewPool("")
	require.NoError(t, err, dockerConnectionError)

	resource, err := pool.Run(dockerImage, "latest", []string{})
	require.NoError(t, err, dockerContainerStartError)

	t.Cleanup(func() {
		require.NoError(t, pool.Purge(resource), dockerPoolPurgeError)
	})
	execCmd := exec.Command("docker", "exec", "-it", resource.Container.Name, "sh")
	err = execCmd.Run()

	if err == nil {
		t.Fatal(expectedError)
	}

	exitErr, ok := err.(*exec.ExitError)
	if !ok {
		t.Fatal("Failed to get exit code:", err)
	}
	if exitErr.ExitCode() != 1 {
		t.Fatalf(expectedExitCode, exitErr.ExitCode())
	}
}

func TestDockerExecAsh(t *testing.T) {
	pool, err := dockertest.NewPool("")
	require.NoError(t, err, dockerConnectionError)

	resource, err := pool.Run(dockerImage, "latest", []string{})
	require.NoError(t, err, dockerContainerStartError)

	t.Cleanup(func() {
		require.NoError(t, pool.Purge(resource), dockerPoolPurgeError)
	})

	execCmd := exec.Command("docker", "exec", "-it", resource.Container.Name, "ash")
	err = execCmd.Run()

	if err == nil {
		t.Fatalf(expectedError)
	}

	exitErr, ok := err.(*exec.ExitError)
	if !ok {
		t.Fatalf(failedToGetExitCode, err)
	}
	if exitErr.ExitCode() != 1 {
		t.Fatalf(expectedExitCode, exitErr.ExitCode())
	}
}

func TestDockerExecLs(t *testing.T) {
	pool, err := dockertest.NewPool("")
	require.NoError(t, err, dockerConnectionError)

	resource, err := pool.Run(dockerImage, "latest", []string{})
	require.NoError(t, err, dockerContainerStartError)

	t.Cleanup(func() {
		require.NoError(t, pool.Purge(resource), dockerPoolPurgeError)
	})

	execCmd := exec.Command("docker", "exec", "-it", resource.Container.Name, "ls")
	err = execCmd.Run()

	if err == nil {
		t.Fatalf(expectedError)
	}

	exitErr, ok := err.(*exec.ExitError)
	if !ok {
		t.Fatalf(failedToGetExitCode, err)
	}
	if exitErr.ExitCode() != 1 {
		t.Fatalf(expectedExitCode, exitErr.ExitCode())
	}
}
