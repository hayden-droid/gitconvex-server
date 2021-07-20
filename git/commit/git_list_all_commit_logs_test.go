package commit

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	git2go "github.com/libgit2/git2go/v31"
	"github.com/neel1996/gitconvex/git/commit/stub"
	"github.com/neel1996/gitconvex/git/middleware"
	"github.com/neel1996/gitconvex/mocks"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type ListAllCommitLogsTestSuite struct {
	suite.Suite
	mockController    *gomock.Controller
	mockRepo          *mocks.MockRepository
	mockRevWalk       *mocks.MockRevWalk
	repo              middleware.Repository
	listAllCommitLogs ListAllLogs
}

func TestListAllCommitLogsTestSuite(t *testing.T) {
	suite.Run(t, new(ListAllCommitLogsTestSuite))
}

func (suite *ListAllCommitLogsTestSuite) SetupTest() {
	suite.mockController = gomock.NewController(suite.T())
	suite.mockRepo = mocks.NewMockRepository(suite.mockController)
	suite.mockRevWalk = mocks.NewMockRevWalk(suite.mockController)

	r, err := git2go.OpenRepository(os.Getenv("GITCONVEX_TEST_REPO"))
	if err != nil {
		fmt.Println(err)
	}

	suite.repo = middleware.NewRepository(r)
	suite.listAllCommitLogs = NewListAllLogs(suite.mockRepo)
}

func (suite *ListAllCommitLogsTestSuite) TestGet_WhenLogsAreAvailable_ShouldReturnCommitLogs() {
	suite.listAllCommitLogs = NewListAllLogs(suite.repo)
	got, err := suite.listAllCommitLogs.Get()

	suite.NotZero(len(got))
	suite.Nil(err)
}

func (suite *ListAllCommitLogsTestSuite) TestGet_WhenRepoWalkFails_ShouldReturnError() {
	suite.mockRepo.EXPECT().Walk().Return(nil, errors.New("WALKER_ERROR"))

	_, err := suite.listAllCommitLogs.Get()

	suite.NotNil(err)
}

func (suite *ListAllCommitLogsTestSuite) TestGet_WhenRepoHasNoCommits_ShouldReturnNoCommitLogs() {
	suite.mockRepo.EXPECT().Walk().Return(stub.NewRevWalkStub(false), nil)

	got, err := suite.listAllCommitLogs.Get()

	suite.Zero(len(got))
	suite.Nil(err)
}

func (suite *ListAllCommitLogsTestSuite) TestGet_WhenIteratorReturnsError_ShouldReturnError() {
	suite.mockRepo.EXPECT().Walk().Return(stub.NewRevWalkStub(true), nil)

	_, err := suite.listAllCommitLogs.Get()

	suite.NotNil(err)
}
