package github

import (
	"context"
	"fmt"
	"strings"
	"sync"
)

// RepoPermissionChecker handles repository access permissions
type RepoPermissionChecker struct {
	allowedRepos []string
	getClient    GetClientFn
	currentUser  string
	userMutex    sync.RWMutex
}

// NewRepoPermissionChecker creates a new repository permission checker
func NewRepoPermissionChecker(allowedRepos []string, getClient GetClientFn) *RepoPermissionChecker {
	return &RepoPermissionChecker{
		allowedRepos: allowedRepos,
		getClient:    getClient,
	}
}

// getCurrentUser gets the authenticated user's login name
func (r *RepoPermissionChecker) getCurrentUser(ctx context.Context) (string, error) {
	r.userMutex.RLock()
	if r.currentUser != "" {
		username := r.currentUser
		r.userMutex.RUnlock()
		return username, nil
	}
	r.userMutex.RUnlock()

	r.userMutex.Lock()
	defer r.userMutex.Unlock()

	// Double-check pattern to avoid race conditions
	if r.currentUser != "" {
		return r.currentUser, nil
	}

	client, err := r.getClient(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get GitHub client: %w", err)
	}

	user, _, err := client.Users.Get(ctx, "")
	if err != nil {
		return "", fmt.Errorf("failed to get current user: %w", err)
	}

	r.currentUser = user.GetLogin()
	return r.currentUser, nil
}

// IsRepoAllowed checks if access to the given repository is allowed
func (r *RepoPermissionChecker) IsRepoAllowed(ctx context.Context, owner, repo string) error {
	// If no restrictions are set, allow all access
	if len(r.allowedRepos) == 0 {
		return nil
	}

	currentUser, err := r.getCurrentUser(ctx)
	if err != nil {
		return fmt.Errorf("failed to get current user for permission check: %w", err)
	}

	// IMPORTANT: Only apply restrictions to current user's personal repositories
	// Allow unrestricted access to other users' and organizations' repositories
	if owner != currentUser {
		return nil // Allow access to all non-personal repositories
	}

	// For personal repositories, check against allowed list
	targetRepo := owner + "/" + repo

	for _, allowedRepo := range r.allowedRepos {
		allowedRepo = strings.TrimSpace(allowedRepo)
		
		if strings.Contains(allowedRepo, "/") {
			// Complete format "owner/repo"
			if allowedRepo == targetRepo {
				return nil
			}
		} else {
			// Shorthand format "repo" - defaults to current user's repository
			if allowedRepo == repo && owner == currentUser {
				return nil
			}
		}
	}

	return fmt.Errorf("access to personal repository '%s/%s' is not allowed by GITHUB_ALLOWED_REPOS configuration", owner, repo)
}