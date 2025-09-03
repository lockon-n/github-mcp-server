package github

import (
	"context"
	
	"github.com/github/github-mcp-server/pkg/translations"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// CreateRepositoryWithPermissionCheck creates a tool to create repository with permission checking
func CreateRepositoryWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := CreateRepository(getClient, t)
	
	return originalTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// If repository restrictions are enabled, block repository creation entirely
		if repoChecker != nil && len(repoChecker.allowedRepos) > 0 {
			return mcp.NewToolResultError("repository creation is disabled when GITHUB_ALLOWED_REPOS is configured for security reasons"), nil
		}

		// Call the original handler if no restrictions are set
		return originalHandler(ctx, request)
	}
}