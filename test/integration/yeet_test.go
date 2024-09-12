package integration

import (
	"os/exec"
	"testing"
	"time"
)

func TestYeetRun(t *testing.T) {
	// Start the Docker Compose environment
	startCmd := exec.Command("docker-compose", "up", "-d")
	startCmd.Dir = "."
	err := startCmd.Run()
	if err != nil {
		t.Fatalf("Failed to start Docker Compose: %v", err)
	}
	defer func() {
		stopCmd := exec.Command("docker-compose", "down")
		stopCmd.Dir = "."
		_ = stopCmd.Run()
	}()

	// Wait for services to be ready
	time.Sleep(10 * time.Second)

	// Run yeet command inside the client container
	yeetCmd := exec.Command("docker-compose", "exec", "client", "/yeet/yeet", "run")
	yeetCmd.Dir = "."
	output, err := yeetCmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to run yeet command: %v\nOutput: %s", err, output)
	}

	// Add assertions based on the expected output
	expectedOutput := "run called"
	if string(output) != expectedOutput {
		t.Errorf("Unexpected output. Got: %s, Want: %s", output, expectedOutput)
	}
}