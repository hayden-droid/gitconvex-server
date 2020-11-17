package tests

import (
	"github.com/go-git/go-git/v5"
	"github.com/neel1996/gitconvex-server/api"
	"github.com/neel1996/gitconvex-server/graph/model"
	"testing"
)

func TestCodeFileView(t *testing.T) {
	r, _ := git.PlainOpen("..")
	w, _ := r.Worktree()
	repoPath := w.Filesystem.Root()
	expectedLine := "# gitconvex GoLang project"

	type args struct {
		repo     *git.Repository
		repoPath string
		fileName string
	}
	tests := []struct {
		name string
		args args
		want *model.CodeFileType
	}{
		{name: "Code view API test case", args: struct {
			repo     *git.Repository
			repoPath string
			fileName string
		}{repo: r, repoPath: repoPath, fileName: "README.md"}, want: &model.CodeFileType{FileData: []*string{&expectedLine}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := api.CodeFileView(tt.args.repoPath, tt.args.fileName); *got.FileData[0] != *tt.want.FileData[0] {
				t.Errorf("CodeFileView() = %v, want %v", got, tt.want)
			}
		})
	}
}
