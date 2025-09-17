package github

import (
	"context"

	"github.com/github/github-mcp-server/pkg/translations"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// CreatePermissionCheckedTool creates a wrapper tool that checks repository permissions
// before calling the original tool handler
func CreatePermissionCheckedTool(originalTool mcp.Tool, originalHandler server.ToolHandlerFunc, repoChecker *RepoPermissionChecker) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// If no permission checker is provided, skip permission checking (e.g., for docs generation)
		if repoChecker == nil {
			return originalHandler(ctx, request)
		}

		// Extract owner and repo parameters for permission check
		owner, err := RequiredParam[string](request, "owner")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		repo, err := RequiredParam[string](request, "repo")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		// Check repository access permission
		if err := repoChecker.IsRepoAllowed(ctx, owner, repo); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		// Call the original handler if permission check passes
		return originalHandler(ctx, request)
	}
}

// ListBranchesWithPermissionCheck creates a tool to list branches with permission checking
func ListBranchesWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := ListBranches(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// ListTagsWithPermissionCheck creates a tool to list tags with permission checking
func ListTagsWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := ListTags(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// GetTagWithPermissionCheck creates a tool to get tag details with permission checking
func GetTagWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := GetTag(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// ListReleasesWithPermissionCheck creates a tool to list releases with permission checking
func ListReleasesWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := ListReleases(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// GetLatestReleaseWithPermissionCheck creates a tool to get latest release with permission checking
func GetLatestReleaseWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := GetLatestRelease(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// GetReleaseByTagWithPermissionCheck creates a tool to get release by tag with permission checking
func GetReleaseByTagWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := GetReleaseByTag(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// ForkRepositoryWithPermissionCheck creates a tool to fork repository with permission checking
func ForkRepositoryWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := ForkRepository(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// RenameRepositoryWithPermissionCheck creates a tool to rename repository with permission checking
func RenameRepositoryWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := RenameRepository(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// CreateBranchWithPermissionCheck creates a tool to create branch with permission checking
func CreateBranchWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := CreateBranch(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// PushFilesWithPermissionCheck creates a tool to push files with permission checking
func PushFilesWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := PushFiles(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}

// DeleteFileWithPermissionCheck creates a tool to delete file with permission checking
func DeleteFileWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := DeleteFile(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}