package cmd

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestParseIntegration(t *testing.T) {
	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Error getting current directory: %v", err)
	}

	// change the current directory to the root of the project
	err = os.Chdir(currentDir + "/..")
	if err != nil {
		t.Fatalf("Error changing directory: %v", err)
	}
	tempOutputDir := "temp_output"
	err = os.Mkdir(tempOutputDir, os.ModePerm)
	if err != nil {
		t.Fatalf("error creating temp directory for test: %v", err)
	}
	defer func() {
		err = os.RemoveAll(tempOutputDir)
		if err != nil {
			t.Fatalf("error removing temp directory for test: %v", err)
		}
	}()

	rootCmd.SetArgs([]string{"parse", "input/data1.csv", "--config=config/config.json", "--output=" + tempOutputDir})

	err = rootCmd.Execute()
	if err != nil {
		t.Fatalf("Error executing parse command: %v", err)
	}
	validFiles, validFileErr := filepath.Glob(filepath.Join(tempOutputDir, "valid_*.csv"))
	invalidFiles, invalidFileErr := filepath.Glob(filepath.Join(tempOutputDir, "invalid_*.csv"))
	if validFileErr != nil || invalidFileErr != nil || len(validFiles) != 1 || len(invalidFiles) != 1 {
		t.Fatalf("Expected valid and invalid CSV files to be created")
	}
	expectedValidContent := []byte("id,name,email,salary\n1,John Doe,doe@test.com,$10.00\n2,Mary Jane,Mary@tes.com,$15\n3,Max Topperson,max@test.com,$11\n")
	assertFileContent(t, err, validFiles[0], string(expectedValidContent))
	expectedInvalidContent := []byte("name,email,wage,number,errors\nAlfred Donald,,$11.5,4,Email is required\n")
	assertFileContent(t, err, invalidFiles[0], string(expectedInvalidContent))
}

func assertFileContent(t *testing.T, err error, file string, expectedContent string) {
	t.Helper()
	actualContent, err := os.ReadFile(file)
	if err != nil {
		t.Errorf("Error reading actual %v CSV file: %v", file, err)
	}
	assert.Equal(t, expectedContent, string(actualContent), "Content of invalid CSV file doesn't match expected")
}
