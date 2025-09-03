package github

import (
	"github.com/github/github-mcp-server/pkg/translations"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// ===== ISSUES TOOLS WITH PERMISSION CHECK =====

// GetIssueWithPermissionCheck creates a tool to get issue with permission checking
func GetIssueWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := GetIssue(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// CreateIssueWithPermissionCheck creates a tool to create issue with permission checking
func CreateIssueWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := CreateIssue(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// AddIssueCommentWithPermissionCheck creates a tool to add issue comment with permission checking
func AddIssueCommentWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := AddIssueComment(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// UpdateIssueWithPermissionCheck creates a tool to update issue with permission checking
func UpdateIssueWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := UpdateIssue(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// GetIssueCommentsWithPermissionCheck creates a tool to get issue comments with permission checking
func GetIssueCommentsWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := GetIssueComments(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// ListIssueTypesWithPermissionCheck creates a tool to list issue types with permission checking
func ListIssueTypesWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := ListIssueTypes(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// AddSubIssueWithPermissionCheck creates a tool to add sub issue with permission checking
func AddSubIssueWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := AddSubIssue(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// ListSubIssuesWithPermissionCheck creates a tool to list sub issues with permission checking
func ListSubIssuesWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := ListSubIssues(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// RemoveSubIssueWithPermissionCheck creates a tool to remove sub issue with permission checking
func RemoveSubIssueWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := RemoveSubIssue(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// ReprioritizeSubIssueWithPermissionCheck creates a tool to reprioritize sub issue with permission checking
func ReprioritizeSubIssueWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := ReprioritizeSubIssue(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// ===== PULL REQUESTS TOOLS WITH PERMISSION CHECK =====

// GetPullRequestWithPermissionCheck creates a tool to get pull request with permission checking
func GetPullRequestWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := GetPullRequest(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// ListPullRequestsWithPermissionCheck creates a tool to list pull requests with permission checking
func ListPullRequestsWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := ListPullRequests(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// CreatePullRequestWithPermissionCheck creates a tool to create pull request with permission checking
func CreatePullRequestWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := CreatePullRequest(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// UpdatePullRequestWithPermissionCheck creates a tool to update pull request with permission checking
func UpdatePullRequestWithPermissionCheck(getClient GetClientFn, getGQLClient GetGQLClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := UpdatePullRequest(getClient, getGQLClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// MergePullRequestWithPermissionCheck creates a tool to merge pull request with permission checking
func MergePullRequestWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := MergePullRequest(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// GetPullRequestFilesWithPermissionCheck creates a tool to get pull request files with permission checking
func GetPullRequestFilesWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := GetPullRequestFiles(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// GetPullRequestStatusWithPermissionCheck creates a tool to get pull request status with permission checking
func GetPullRequestStatusWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := GetPullRequestStatus(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// UpdatePullRequestBranchWithPermissionCheck creates a tool to update pull request branch with permission checking
func UpdatePullRequestBranchWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := UpdatePullRequestBranch(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// GetPullRequestCommentsWithPermissionCheck creates a tool to get pull request comments with permission checking
func GetPullRequestCommentsWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := GetPullRequestComments(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// GetPullRequestReviewsWithPermissionCheck creates a tool to get pull request reviews with permission checking
func GetPullRequestReviewsWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := GetPullRequestReviews(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// GetPullRequestDiffWithPermissionCheck creates a tool to get pull request diff with permission checking
func GetPullRequestDiffWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := GetPullRequestDiff(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// RequestCopilotReviewWithPermissionCheck creates a tool to request copilot review with permission checking
func RequestCopilotReviewWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := RequestCopilotReview(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// ===== PR REVIEW TOOLS WITH PERMISSION CHECK =====

// CreateAndSubmitPullRequestReviewWithPermissionCheck creates a tool with permission checking
func CreateAndSubmitPullRequestReviewWithPermissionCheck(getGQLClient GetGQLClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := CreateAndSubmitPullRequestReview(getGQLClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// CreatePendingPullRequestReviewWithPermissionCheck creates a tool with permission checking
func CreatePendingPullRequestReviewWithPermissionCheck(getGQLClient GetGQLClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := CreatePendingPullRequestReview(getGQLClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// AddCommentToPendingReviewWithPermissionCheck creates a tool with permission checking
func AddCommentToPendingReviewWithPermissionCheck(getGQLClient GetGQLClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := AddCommentToPendingReview(getGQLClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// SubmitPendingPullRequestReviewWithPermissionCheck creates a tool with permission checking
func SubmitPendingPullRequestReviewWithPermissionCheck(getGQLClient GetGQLClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := SubmitPendingPullRequestReview(getGQLClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// DeletePendingPullRequestReviewWithPermissionCheck creates a tool with permission checking
func DeletePendingPullRequestReviewWithPermissionCheck(getGQLClient GetGQLClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := DeletePendingPullRequestReview(getGQLClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}