package github

import (
	"github.com/github/github-mcp-server/pkg/translations"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// ===== SECURITY ADVISORIES TOOLS WITH PERMISSION CHECK =====

// ListRepositorySecurityAdvisoriesWithPermissionCheck creates a tool with permission checking
func ListRepositorySecurityAdvisoriesWithPermissionCheck(getClient GetClientFn, t translations.TranslationHelperFunc, repoChecker *RepoPermissionChecker) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	originalTool, originalHandler := ListRepositorySecurityAdvisories(getClient, t)
	return originalTool, CreatePermissionCheckedTool(originalTool, originalHandler, repoChecker)
}