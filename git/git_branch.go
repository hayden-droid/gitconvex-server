package git

import (
	"fmt"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/neel1996/gitconvex-server/global"
	"go/types"
)

type Branch struct {
	CurrentBranch string
	BranchList    []string
}

func GetBranchList(repo *git.Repository) *Branch {
	var branchList *Branch
	var branches []string
	var currentBranch string

	logger := global.Logger{}

	if repo != nil {
		head, _ := repo.Head()
		currentBranch = head.Name().String()

		bIter, _ := repo.Branches()
		_ = bIter.ForEach(func(reference *plumbing.Reference) error {
			if reference != nil {
				branches = append(branches, reference.String())
				return nil
			} else {
				return types.Error{Msg: "Empty reference"}
			}
		})
		bIter.Close()
	}

	branchList = &Branch{
		BranchList:    branches,
		CurrentBranch: currentBranch,
	}

	logger.Log(fmt.Sprintf("Obtained branch info -- \n%v\n", branchList), global.StatusInfo)
	return branchList
}
