package github

import (
	"context"

	"github.com/github/github-mcp-server/pkg/raw"
	"github.com/github/github-mcp-server/pkg/toolsets"
	"github.com/github/github-mcp-server/pkg/translations"
	"github.com/google/go-github/v74/github"
	"github.com/mark3labs/mcp-go/server"
	"github.com/shurcooL/githubv4"
)

type GetClientFn func(context.Context) (*github.Client, error)
type GetGQLClientFn func(context.Context) (*githubv4.Client, error)

var DefaultTools = []string{"all"}

func DefaultToolsetGroup(readOnly bool, getClient GetClientFn, getGQLClient GetGQLClientFn, getRawClient raw.GetRawClientFn, t translations.TranslationHelperFunc, contentWindowSize int, repoChecker *RepoPermissionChecker) *toolsets.ToolsetGroup {
	tsg := toolsets.NewToolsetGroup(readOnly)

	// Create toolsets - all tools use permission checking (null-safe when repoChecker is nil)
	repos := toolsets.NewToolset("repos", "GitHub Repository related tools").
		AddReadTools(
			toolsets.NewServerTool(SearchRepositories(getClient, t)),
			toolsets.NewServerTool(GetFileContentsWithPermissionCheck(getClient, getRawClient, t, repoChecker)),
			toolsets.NewServerTool(ListCommitsWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(SearchCode(getClient, t)),
			toolsets.NewServerTool(GetCommitWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(ListBranchesWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(ListTagsWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(GetTagWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(ListReleasesWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(GetLatestReleaseWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(GetReleaseByTagWithPermissionCheck(getClient, t, repoChecker)),
		).
		AddWriteTools(
			toolsets.NewServerTool(CreateOrUpdateFileWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(CreateRepositoryWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(ForkRepositoryWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(RenameRepositoryWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(CreateBranchWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(PushFilesWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(DeleteFileWithPermissionCheck(getClient, t, repoChecker)),
		).
		AddResourceTemplates(
			toolsets.NewServerResourceTemplate(GetRepositoryResourceContent(getClient, getRawClient, t)),
			toolsets.NewServerResourceTemplate(GetRepositoryResourceBranchContent(getClient, getRawClient, t)),
			toolsets.NewServerResourceTemplate(GetRepositoryResourceCommitContent(getClient, getRawClient, t)),
			toolsets.NewServerResourceTemplate(GetRepositoryResourceTagContent(getClient, getRawClient, t)),
			toolsets.NewServerResourceTemplate(GetRepositoryResourcePrContent(getClient, getRawClient, t)),
		)
	issues := toolsets.NewToolset("issues", "GitHub Issues related tools").
		AddReadTools(
			toolsets.NewServerTool(GetIssueWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(SearchIssues(getClient, t)),
			toolsets.NewServerTool(ListIssues(getGQLClient, t)),
			toolsets.NewServerTool(GetIssueCommentsWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(ListIssueTypesWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(ListSubIssuesWithPermissionCheck(getClient, t, repoChecker)),
		).
		AddWriteTools(
			toolsets.NewServerTool(CreateIssueWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(AddIssueCommentWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(UpdateIssueWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(AssignCopilotToIssue(getGQLClient, t)),
			toolsets.NewServerTool(AddSubIssueWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(RemoveSubIssueWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(ReprioritizeSubIssueWithPermissionCheck(getClient, t, repoChecker)),
		).AddPrompts(
		toolsets.NewServerPrompt(AssignCodingAgentPrompt(t)),
		toolsets.NewServerPrompt(IssueToFixWorkflowPrompt(t)),
	)
	users := toolsets.NewToolset("users", "GitHub User related tools").
		AddReadTools(
			toolsets.NewServerTool(SearchUsers(getClient, t)),
		)
	orgs := toolsets.NewToolset("orgs", "GitHub Organization related tools").
		AddReadTools(
			toolsets.NewServerTool(SearchOrgs(getClient, t)),
		)
	pullRequests := toolsets.NewToolset("pull_requests", "GitHub Pull Request related tools").
		AddReadTools(
			toolsets.NewServerTool(GetPullRequestWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(ListPullRequestsWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(GetPullRequestFilesWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(SearchPullRequests(getClient, t)),
			toolsets.NewServerTool(GetPullRequestStatusWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(GetPullRequestCommentsWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(GetPullRequestReviewsWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(GetPullRequestDiffWithPermissionCheck(getClient, t, repoChecker)),
		).
		AddWriteTools(
			toolsets.NewServerTool(MergePullRequestWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(UpdatePullRequestBranchWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(CreatePullRequestWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(UpdatePullRequestWithPermissionCheck(getClient, getGQLClient, t, repoChecker)),
			toolsets.NewServerTool(RequestCopilotReviewWithPermissionCheck(getClient, t, repoChecker)),

			// Reviews
			toolsets.NewServerTool(CreateAndSubmitPullRequestReviewWithPermissionCheck(getGQLClient, t, repoChecker)),
			toolsets.NewServerTool(CreatePendingPullRequestReviewWithPermissionCheck(getGQLClient, t, repoChecker)),
			toolsets.NewServerTool(AddCommentToPendingReviewWithPermissionCheck(getGQLClient, t, repoChecker)),
			toolsets.NewServerTool(SubmitPendingPullRequestReviewWithPermissionCheck(getGQLClient, t, repoChecker)),
			toolsets.NewServerTool(DeletePendingPullRequestReviewWithPermissionCheck(getGQLClient, t, repoChecker)),
		)
	codeSecurity := toolsets.NewToolset("code_security", "Code security related tools, such as GitHub Code Scanning").
		AddReadTools(
			toolsets.NewServerTool(GetCodeScanningAlertWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(ListCodeScanningAlertsWithPermissionCheck(getClient, t, repoChecker)),
		)
	secretProtection := toolsets.NewToolset("secret_protection", "Secret protection related tools, such as GitHub Secret Scanning").
		AddReadTools(
			toolsets.NewServerTool(GetSecretScanningAlertWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(ListSecretScanningAlertsWithPermissionCheck(getClient, t, repoChecker)),
		)
	dependabot := toolsets.NewToolset("dependabot", "Dependabot tools").
		AddReadTools(
			toolsets.NewServerTool(GetDependabotAlertWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(ListDependabotAlertsWithPermissionCheck(getClient, t, repoChecker)),
		)

	notifications := toolsets.NewToolset("notifications", "GitHub Notifications related tools").
		AddReadTools(
			toolsets.NewServerTool(ListNotifications(getClient, t)),
			toolsets.NewServerTool(GetNotificationDetails(getClient, t)),
		).
		AddWriteTools(
			toolsets.NewServerTool(DismissNotification(getClient, t)),
			toolsets.NewServerTool(MarkAllNotificationsRead(getClient, t)),
			toolsets.NewServerTool(ManageNotificationSubscription(getClient, t)),
			toolsets.NewServerTool(ManageRepositoryNotificationSubscriptionWithPermissionCheck(getClient, t, repoChecker)),
		)

	discussions := toolsets.NewToolset("discussions", "GitHub Discussions related tools").
		AddReadTools(
			toolsets.NewServerTool(ListDiscussions(getGQLClient, t)),
			toolsets.NewServerTool(GetDiscussion(getGQLClient, t)),
			toolsets.NewServerTool(GetDiscussionComments(getGQLClient, t)),
			toolsets.NewServerTool(ListDiscussionCategories(getGQLClient, t)),
		)

	actions := toolsets.NewToolset("actions", "GitHub Actions workflows and CI/CD operations").
		AddReadTools(
			toolsets.NewServerTool(ListWorkflowsWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(ListWorkflowRunsWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(GetWorkflowRunWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(GetWorkflowRunLogsWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(ListWorkflowJobsWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(GetJobLogsWithPermissionCheck(getClient, t, contentWindowSize, repoChecker)),
			toolsets.NewServerTool(ListWorkflowRunArtifactsWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(DownloadWorkflowRunArtifactWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(GetWorkflowRunUsageWithPermissionCheck(getClient, t, repoChecker)),
		).
		AddWriteTools(
			toolsets.NewServerTool(RunWorkflowWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(RerunWorkflowRunWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(RerunFailedJobsWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(CancelWorkflowRunWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(DeleteWorkflowRunLogsWithPermissionCheck(getClient, t, repoChecker)),
		)

	securityAdvisories := toolsets.NewToolset("security_advisories", "Security advisories related tools").
		AddReadTools(
			toolsets.NewServerTool(ListGlobalSecurityAdvisories(getClient, t)),
			toolsets.NewServerTool(GetGlobalSecurityAdvisory(getClient, t)),
			toolsets.NewServerTool(ListRepositorySecurityAdvisoriesWithPermissionCheck(getClient, t, repoChecker)),
			toolsets.NewServerTool(ListOrgRepositorySecurityAdvisories(getClient, t)),
		)

	// Keep experiments alive so the system doesn't error out when it's always enabled
	experiments := toolsets.NewToolset("experiments", "Experimental features that are not considered stable yet")

	contextTools := toolsets.NewToolset("context", "Tools that provide context about the current user and GitHub context you are operating in").
		AddReadTools(
			toolsets.NewServerTool(GetMe(getClient, t)),
			toolsets.NewServerTool(GetTeams(getClient, getGQLClient, t)),
			toolsets.NewServerTool(GetTeamMembers(getGQLClient, t)),
		)

	gists := toolsets.NewToolset("gists", "GitHub Gist related tools").
		AddReadTools(
			toolsets.NewServerTool(ListGists(getClient, t)),
		).
		AddWriteTools(
			toolsets.NewServerTool(CreateGist(getClient, t)),
			toolsets.NewServerTool(UpdateGist(getClient, t)),
		)

	// Add toolsets to the group
	tsg.AddToolset(contextTools)
	tsg.AddToolset(repos)
	tsg.AddToolset(issues)
	tsg.AddToolset(orgs)
	tsg.AddToolset(users)
	tsg.AddToolset(pullRequests)
	tsg.AddToolset(actions)
	tsg.AddToolset(codeSecurity)
	tsg.AddToolset(secretProtection)
	tsg.AddToolset(dependabot)
	tsg.AddToolset(notifications)
	tsg.AddToolset(experiments)
	tsg.AddToolset(discussions)
	tsg.AddToolset(gists)
	tsg.AddToolset(securityAdvisories)

	return tsg
}

// InitDynamicToolset creates a dynamic toolset that can be used to enable other toolsets, and so requires the server and toolset group as arguments
func InitDynamicToolset(s *server.MCPServer, tsg *toolsets.ToolsetGroup, t translations.TranslationHelperFunc) *toolsets.Toolset {
	// Create a new dynamic toolset
	// Need to add the dynamic toolset last so it can be used to enable other toolsets
	dynamicToolSelection := toolsets.NewToolset("dynamic", "Discover GitHub MCP tools that can help achieve tasks by enabling additional sets of tools, you can control the enablement of any toolset to access its tools when this toolset is enabled.").
		AddReadTools(
			toolsets.NewServerTool(ListAvailableToolsets(tsg, t)),
			toolsets.NewServerTool(GetToolsetsTools(tsg, t)),
			toolsets.NewServerTool(EnableToolset(s, tsg, t)),
		)

	dynamicToolSelection.Enabled = true
	return dynamicToolSelection
}

// ToBoolPtr converts a bool to a *bool pointer.
func ToBoolPtr(b bool) *bool {
	return &b
}

// ToStringPtr converts a string to a *string pointer.
// Returns nil if the string is empty.
func ToStringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
