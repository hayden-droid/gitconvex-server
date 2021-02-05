package git

import (
	git2go "github.com/libgit2/git2go/v31"
	"github.com/neel1996/gitconvex-server/global"
	"github.com/neel1996/gitconvex-server/graph/model"
)

type UnPushedCommitInterface interface {
	UnPushedCommits() []*model.GitCommits
}

type UnPushedCommitStruct struct {
	Repo      *git2go.Repository
	RemoteRef string
}

// commitModel function generates a pipe separated string with the required commit details
// The resultant string will be sent to the client
func commitModel(commit *git2go.Commit) *model.GitCommits {
	commitHash := commit.Id().String()[:7]
	commitAuthor := commit.Author().Name
	commitMessage := commit.Message()
	commitDate := commit.Author().When.String()

	return &model.GitCommits{
		Hash:          &commitHash,
		Author:        &commitAuthor,
		CommitTime:    &commitDate,
		CommitMessage: &commitMessage,
	}
}

// UnPushedCommits compares the local branch and the remote branch to extract the commits which are not pushed to the remote
func (u UnPushedCommitStruct) UnPushedCommits() []*model.GitCommits {
	repo := u.Repo
	remoteRef := u.RemoteRef
	var commitArray []*model.GitCommits

	// Returning nil commit response if repo has no HEAD
	head, _ := repo.Head()
	if head == nil {
		return []*model.GitCommits{}
	}

	remoteBranch, remoteBranchErr := repo.LookupBranch(remoteRef, git2go.BranchRemote)
	if remoteBranchErr != nil {
		logger.Log(remoteBranchErr.Error(), global.StatusError)
		return commitArray
	}

	// Checking if both branches have any varying commits
	diff := head.Cmp(remoteBranch.Reference)
	if diff == 0 {
		return commitArray
	}

	localCommit, _ := repo.LookupCommit(head.Target())
	remoteCommit, _ := repo.LookupCommit(remoteBranch.Target())

	if localCommit != nil && remoteCommit != nil {
		commonAncestor, _ := repo.MergeBase(localCommit.Id(), remoteCommit.Id())
		if commonAncestor != nil {
			commitArray = append(commitArray, commitModel(localCommit))
			// Return if there is only one new commit to be pushed
			if diff == 1 {
				return commitArray
			}

			n := localCommit.ParentCount()
			var i uint
			for i = 0; i < n; i++ {
				currentCommit := localCommit.Parent(i)
				if currentCommit != nil && currentCommit.Id() != commonAncestor {
					commitArray = append(commitArray, commitModel(localCommit))
				} else {
					break
				}
			}
		} else {
			logger.Log("No new commits available to push", global.StatusWarning)
			return []*model.GitCommits{}
		}
		return commitArray
	} else {
		logger.Log("No new commits available to push", global.StatusWarning)
		return []*model.GitCommits{}
	}
}
