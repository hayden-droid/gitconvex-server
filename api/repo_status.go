package api

import (
	"github.com/neel1996/gitconvex-server/git"
	"github.com/neel1996/gitconvex-server/global"
	"github.com/neel1996/gitconvex-server/graph/model"
	"strings"
)

func RepoStatus(repoId string) *model.GitRepoStatusResults {
	logger := global.Logger{}
	logger.Log("Collecting repo status information", global.StatusInfo)

	repoChan := make(chan git.RepoDetails)
	remoteChan := make(chan git.RemoteDataModel)
	branchChan := make(chan git.Branch)
	commitChan := make(chan git.AllCommitData)
	trackedFileCountChan := make(chan int)

	go git.Repo(repoId, repoChan)

	var repoName *string
	r := <-repoChan
	repo := r.GitRepo

	remote := ""
	var remoteURL *string
	remoteURL = &remote
	go git.RemoteData(repo, remoteChan)
	remoteData := <-remoteChan
	remotes := remoteData.RemoteURL
	sRemote := strings.Split(*remotes[0], "/")
	repoName = &sRemote[len(sRemote)-1]

	if len(remotes) > 1 {
		var tempRemoteArray []string
		for _, ptrRemote := range remotes {
			tempRemoteArray = append(tempRemoteArray, *ptrRemote)
		}
		*remoteURL = strings.Join(tempRemoteArray, "||")
	} else {
		*remoteURL = *remotes[0]
	}

	go git.GetBranchList(repo, branchChan)
	branchList := <-branchChan
	currentBranch := &branchList.CurrentBranch
	branches := branchList.BranchList
	allBranches := branchList.AllBranchList

	var latestCommit *string

	go git.AllCommits(repo, commitChan)
	commitData := <-commitChan
	latestCommit = &commitData.LatestCommit
	totalCommits := commitData.TotalCommits
	totalCommitsPtr := &totalCommits

	go git.TrackedFileCount(repo, trackedFileCountChan)
	trackedFileCount := <-trackedFileCountChan
	trackedFilePtr := &trackedFileCount

	return &model.GitRepoStatusResults{
		GitRemoteData:        remoteURL,
		GitRepoName:          repoName,
		GitBranchList:        branches,
		GitAllBranchList:     allBranches,
		GitCurrentBranch:     currentBranch,
		GitRemoteHost:        remoteData.RemoteHost,
		GitTotalCommits:      totalCommitsPtr,
		GitLatestCommit:      latestCommit,
		GitTotalTrackedFiles: trackedFilePtr,
	}
}
