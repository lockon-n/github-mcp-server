package github

import (
	"github.com/github/github-mcp-server/pkg/translations"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// ===== ACTIONS TOOLS WITH PERMISSION CHECK =====

// ListWorkflowsWithPermissionCheck creates a tool to list workflows with permission checking
func ListWorkflowsWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := ListWorkflows(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// ListWorkflowRunsWithPermissionCheck creates a tool to list workflow runs with permission checking
func ListWorkflowRunsWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := ListWorkflowRuns(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// RunWorkflowWithPermissionCheck creates a tool to run workflow with permission checking
func RunWorkflowWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := RunWorkflow(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// GetWorkflowRunWithPermissionCheck creates a tool to get workflow run with permission checking
func GetWorkflowRunWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := GetWorkflowRun(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// GetWorkflowRunLogsWithPermissionCheck creates a tool to get workflow run logs with permission checking
func GetWorkflowRunLogsWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := GetWorkflowRunLogs(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// ListWorkflowJobsWithPermissionCheck creates a tool to list workflow jobs with permission checking
func ListWorkflowJobsWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := ListWorkflowJobs(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// GetJobLogsWithPermissionCheck creates a tool to get job logs with permission checking
func GetJobLogsWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, contentWindowSize int, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := GetJobLogs(getClient, t, contentWindowSize)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// RerunWorkflowRunWithPermissionCheck creates a tool to rerun workflow run with permission checking
func RerunWorkflowRunWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := RerunWorkflowRun(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// RerunFailedJobsWithPermissionCheck creates a tool to rerun failed jobs with permission checking
func RerunFailedJobsWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := RerunFailedJobs(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// CancelWorkflowRunWithPermissionCheck creates a tool to cancel workflow run with permission checking
func CancelWorkflowRunWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := CancelWorkflowRun(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// ListWorkflowRunArtifactsWithPermissionCheck creates a tool to list workflow run artifacts with permission checking
func ListWorkflowRunArtifactsWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := ListWorkflowRunArtifacts(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// DownloadWorkflowRunArtifactWithPermissionCheck creates a tool to download workflow run artifact with permission checking
func DownloadWorkflowRunArtifactWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := DownloadWorkflowRunArtifact(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// DeleteWorkflowRunLogsWithPermissionCheck creates a tool to delete workflow run logs with permission checking
func DeleteWorkflowRunLogsWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := DeleteWorkflowRunLogs(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// GetWorkflowRunUsageWithPermissionCheck creates a tool to get workflow run usage with permission checking
func GetWorkflowRunUsageWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := GetWorkflowRunUsage(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// ===== SECURITY TOOLS WITH PERMISSION CHECK =====

// GetCodeScanningAlertWithPermissionCheck creates a tool to get code scanning alert with permission checking
func GetCodeScanningAlertWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := GetCodeScanningAlert(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// ListCodeScanningAlertsWithPermissionCheck creates a tool to list code scanning alerts with permission checking
func ListCodeScanningAlertsWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := ListCodeScanningAlerts(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// GetSecretScanningAlertWithPermissionCheck creates a tool to get secret scanning alert with permission checking
func GetSecretScanningAlertWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := GetSecretScanningAlert(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// ListSecretScanningAlertsWithPermissionCheck creates a tool to list secret scanning alerts with permission checking
func ListSecretScanningAlertsWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := ListSecretScanningAlerts(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// GetDependabotAlertWithPermissionCheck creates a tool to get dependabot alert with permission checking
func GetDependabotAlertWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := GetDependabotAlert(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// ListDependabotAlertsWithPermissionCheck creates a tool to list dependabot alerts with permission checking
func ListDependabotAlertsWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := ListDependabotAlerts(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// ===== NOTIFICATIONS TOOLS WITH PERMISSION CHECK =====

// ManageRepositoryNotificationSubscriptionWithPermissionCheck creates a tool with permission checking
func ManageRepositoryNotificationSubscriptionWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := ManageRepositoryNotificationSubscription(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}