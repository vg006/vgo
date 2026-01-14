package license

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

type GitHubLicense struct {
	Key  string `json:"key"`
	Name string `json:"name"`
	Body string `json:"body"`
}

// GetGitUsername retrieves the git user.name from global config
func GetGitUsername() string {
	cmd := exec.Command("git", "config", "--global", "--get", "user.name")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(output))
}

// FetchLicenses retrieves the list of available licenses from GitHub API
func FetchLicenses() ([]GitHubLicense, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get("https://api.github.com/licenses")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch licenses: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("github API returned status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var licenses []GitHubLicense
	err = json.Unmarshal(body, &licenses)
	if err != nil {
		return nil, fmt.Errorf("failed to parse licenses: %w", err)
	}

	return licenses, nil
}

// GetLicenseContent retrieves the full license content from GitHub API
func GetLicenseContent(key string) (string, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	url := fmt.Sprintf("https://api.github.com/licenses/%s", key)
	resp, err := client.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch license content: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("github API returned status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	var license GitHubLicense
	err = json.Unmarshal(body, &license)
	if err != nil {
		return "", fmt.Errorf("failed to parse license: %w", err)
	}

	return license.Body, nil
}

// RenderLicense replaces placeholders in license content with actual values
func RenderLicense(content, author, year string) string {
	replacements := map[string]string{
		"[year]":                      year,
		"[fullname]":                  author,
		"[yyyy]":                      year,
		"[name of copyright owner]":   author,
		"[NAME OF COPYRIGHT OWNER]":   author,
	}

	result := content
	for placeholder, value := range replacements {
		result = strings.ReplaceAll(result, placeholder, value)
	}

	return result
}
