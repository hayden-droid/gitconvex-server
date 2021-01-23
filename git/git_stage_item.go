package git

import (
	"fmt"
	git2go "github.com/libgit2/git2go/v31"
	"github.com/neel1996/gitconvex-server/global"
	"io/ioutil"
)

type StageItemInterface interface {
	StageItem() string
	addError(errMsg string) string
}

type StageItemStruct struct {
	Repo     *git2go.Repository
	FileItem string
}

func (s StageItemStruct) addError(errMsg string) string {
	logger := global.Logger{}
	logger.Log(fmt.Sprintf("Error occurred while staging %s -> %s", s.FileItem, errMsg), global.StatusError)
	return global.StageItemError
}

// StageItem stages a selected file from the target repo
// The function relies on the native git client to stage an item, as go-git staging is time consuming for huge repos
func (s StageItemStruct) StageItem() string {
	repo := s.Repo
	fileItem := s.FileItem
	repoPath := repo.Workdir()

	fileByte, _ := ioutil.ReadFile(repoPath + "/" + fileItem)
	fileId, fileIdErr := repo.CreateBlobFromBuffer(fileByte)

	if fileIdErr != nil {
		return s.addError(fileIdErr.Error())
	}

	indexEntry := git2go.IndexEntry{
		Mode: git2go.FilemodeBlob,
		Id:   fileId,
		Path: fileItem,
	}

	repoIndex, repoIndexErr := repo.Index()
	if repoIndexErr != nil {
		return s.addError(repoIndexErr.Error())
	}

	stageErr := repoIndex.Add(&indexEntry)
	if stageErr != nil {
		return s.addError(stageErr.Error())
	} else {
		indexWriteErr := repoIndex.Write()
		if indexWriteErr != nil {
			return s.addError(indexWriteErr.Error())
		}

		logger.Log(fmt.Sprintf("File -> %s staged", fileItem), global.StatusInfo)
		return global.StageItemSuccess
	}
}
