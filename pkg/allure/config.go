package allure

const (
	// DefaultVersion - allure-testify current Version
	DefaultVersion = "Allure-Testify@v0.5.0"

	resultsPathEnvKey     = "ALLURE_OUTPUT_PATH"      // Indicates the path to the results print folder
	outputFolderEnvKey    = "ALLURE_OUTPUT_FOLDER"    // Indicates the name of the folder to print the results.
	issuePatternEnvKey    = "ALLURE_ISSUE_PATTERN"    // Indicates the URL pattern for Issue. It must contain exactly one `%s`
	testCasePatternEnvKey = "ALLURE_TESTCASE_PATTERN" // Indicates the URL pattern for TestCase. It must contain exactly one `%s`

	defaultTagsEnvKey = "ALLURE_LAUNCH_TAGS" // Indicates the default tags that will mark all tests in the run. The tags must be specified separated by commas.

	fileSystemPermissionCode = 0644 // Attachment permission
)
