package test

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/gcp"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Give the resources an environment to operate as a part of, for the purposes of resource tagging
// Give it a random string so we're sure it's created this test run
var expectedProject string
var testPreq *testing.T
var terraformOptions *terraform.Options
var tmpSaUserEmail string
var blacklistRegions []string
var projectId string

func TestMain(m *testing.M) {
	expectedProject = fmt.Sprintf("terratest %s", strings.ToLower(random.UniqueId()))
	blacklistRegions = []string{"asia-east2"}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func(){
		<-c
		TestCleanup(testPreq)
		os.Exit(1)
	}()

	result := m.Run()

	Clean()

	os.Exit(result)
}

// -------------------------------------------------------------------------------------------------------- //
// Utility functions
// -------------------------------------------------------------------------------------------------------- //
func setTerraformOptions(dir string) {
	terraformOptions = &terraform.Options {
		TerraformDir: dir,
		// Pass the expectedProject for tagging
		Vars: map[string]interface{}{
			"project": expectedProject,
			"sa_user_email": tmpSaUserEmail,
		},
		EnvVars: map[string]string{
			"GOOGLE_CLOUD_PROJECT": projectId,
		},
	}
}

// A build step that removes temporary build and test files
func Clean() error {
	fmt.Println("Cleaning...")

	return filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && info.Name() == "vendor" {
			return filepath.SkipDir
		}
		if info.IsDir() && info.Name() == ".terraform" {
			os.RemoveAll(path)
			fmt.Printf("Removed \"%v\"\n", path)
			return filepath.SkipDir
		}
		if !info.IsDir() && (info.Name() == "terraform.tfstate" ||
		info.Name() == "terraform.tfplan" ||
		info.Name() == "terraform.tfstate.backup") {
			os.Remove(path)
			fmt.Printf("Removed \"%v\"\n", path)
		}
		return nil
	})
}

func Test_Prereq(t *testing.T) {
	projectId = gcp.GetGoogleProjectIDFromEnvVar(t)

	setTerraformOptions(".")
	testPreq = t

	terraform.InitAndApply(t, terraformOptions)

	tmpSaUserEmail = terraform.OutputRequired(t, terraformOptions, "sa_user_email")
}

// -------------------------------------------------------------------------------------------------------- //
// Unit Tests
// -------------------------------------------------------------------------------------------------------- //
func TestUT_Assertions(t *testing.T) {
	expectedAssertNameTooLong := "'s generated name is too long:"
	expectedAssertNameInvalidChars := "does not match regex"

	setTerraformOptions("assertions")

	out, err := terraform.InitAndPlanE(t, terraformOptions)

	require.Error(t, err)
	assert.Contains(t, out, expectedAssertNameTooLong)
	assert.Contains(t, out, expectedAssertNameInvalidChars)
}

func TestUT_Defaults(t *testing.T) {
	setTerraformOptions("defaults")
	terraform.InitAndPlan(t, terraformOptions)
}

func TestUT_Overrides(t *testing.T) {
	setTerraformOptions("overrides")
	terraform.InitAndPlan(t, terraformOptions)
}

// -------------------------------------------------------------------------------------------------------- //
// Integration Tests
// -------------------------------------------------------------------------------------------------------- //

func TestIT_Defaults(t *testing.T) {
	setTerraformOptions("defaults")

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	outputs := terraform.OutputAll(t, terraformOptions)

	// Ugly typecasting because Go....
	saMap := outputs["map"].(map[string]interface{})
	sa := saMap["bucket_reader"].(map[string]interface{})
	saId := sa["id"].(string)

	// Make sure our SA is created
	fmt.Printf("Checking Service Account %s...\n", saId)
	assert.Contains(t, saId, "terratest")
}

func TestIT_Overrides(t *testing.T) {
	setTerraformOptions("overrides")

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	outputs := terraform.OutputAll(t, terraformOptions)

	// Ugly typecasting because Go....
	saMap := outputs["map"].(map[string]interface{})
	sa := saMap["bucket_reader"].(map[string]interface{})
	saId := sa["id"].(string)

	// Make sure our SA is created
	fmt.Printf("Checking Service Account %s...\n", saId)
	assert.Contains(t, saId, "terratest")
}

// -------------------------------------------------------------------------------------------------------- //
// Cleanup Tests
// -------------------------------------------------------------------------------------------------------- //
func TestCleanup(t *testing.T) {
	fmt.Println("Cleaning possible lingering resources..")
	terraform.Destroy(t, terraformOptions)

	// Also clean up prereq. resources
	fmt.Println("Cleaning our prereq resources...")
	setTerraformOptions(".")
	terraform.Destroy(t, terraformOptions)
}
