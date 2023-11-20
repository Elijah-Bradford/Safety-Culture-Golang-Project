package folders_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllFolders(t *testing.T) {
	t.Run("Test with valid OrgID", func(t *testing.T) {
		// Create a sample FetchFolderRequest
		req := &folders.FetchFolderRequest{
			OrgID: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
		}

		// Call the GetAllFolders function
		response, err := folders.GetAllFolders(req)

		//Assert No Errors Returned
		assert.NoError(t, err)
		//Assert Response is not nil
		assert.NotNil(t, response)
		assert.NotNil(t, response.Folders)
		//Assert all 666 Folders returned
		assert.Equal(t, 666, len(response.Folders))
	})

	t.Run("Test with invalid OrgID", func(t *testing.T) {
		// Create a FetchFolderRequest with an invalid OrgID
		invalidReq := &folders.FetchFolderRequest{
			OrgID: uuid.FromStringOrNil("invalid_org_id"),
		}

		// Call the GetAllFolders function with an invalid OrgID
		response, _ := folders.GetAllFolders(invalidReq)

		//Assert that no Folders are returned
		assert.Empty(t, response.Folders)
	})

	t.Run("Test with no OrgID", func(t *testing.T) {
		//Create invalid request with no uuid
		invalidReq := &folders.FetchFolderRequest{}

		//Call the GetAllFolders function with invalid request
		response, err := folders.GetAllFolders(invalidReq)

		//Assert that an Error has been returned
		assert.Error(t, err)
		//Assert that the response is empty
		assert.Empty(t, response)
	})

	t.Run("Test with no Request", func(t *testing.T) {
		var nilReq *folders.FetchFolderRequest

		//Call the GetAllFolders function with nil request
		response, err := folders.GetAllFolders(nilReq)

		//Assert that an Error has been returned
		assert.Error(t, err)
		//Assert that the response is empty
		assert.Empty(t, response)

	})
}

func Test_RequestFoldersPaginated(t *testing.T) {
	t.Run("Test 1st Folder, empty Token", func(t *testing.T) {
		//Create Request
		req := &folders.FetchFolderPaginationRequest{
			OrgID:    uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
			Token:    "",
			PageSize: 1,
		}

		//Call RequestFoldersPaginated function
		response, err := folders.RequestFoldersPaginated(req)

		//Assert that the Correct Folder was output with no errors
		assert.NoError(t, err)
		assert.Equal(t, 1, len(response.Folders))
		assert.Equal(t, "current-excalibur", response.Folders[0].Name)

		//Assert Correct Token Returned
		assert.Equal(t, "7ee73e98-b5a7-4ff5-a710-bfd8077ac0a9", response.Token)
	})

	t.Run("Test 2nd Folder from token", func(t *testing.T) {
		//Create Request to get next folder
		NewReq := &folders.FetchFolderPaginationRequest{
			OrgID:    uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
			Token:    "7ee73e98-b5a7-4ff5-a710-bfd8077ac0a9",
			PageSize: 1,
		}

		//Make New Request
		response, err := folders.RequestFoldersPaginated(NewReq)

		//Assert Next Folder Returned with no errors
		assert.NoError(t, err)
		assert.Equal(t, 1, len(response.Folders))
		assert.Equal(t, "amusing-boom", response.Folders[0].Name)
	})

	t.Run("Test PageLength greater than amount of folders", func(t *testing.T) {
		//Create Request
		req := &folders.FetchFolderPaginationRequest{
			OrgID:    uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
			Token:    "",
			PageSize: 1000,
		}

		//Call RequestFoldersPaginated function
		response, err := folders.RequestFoldersPaginated(req)

		//Assert that the correct number of folders was output with no errors
		assert.NoError(t, err)
		assert.Equal(t, 666, len(response.Folders))

		//Assert Correct (empty) Token Returned
		assert.Equal(t, "", response.Token)
	})

	t.Run("Test Invalid Token", func(t *testing.T) {
		//Create Request
		req := &folders.FetchFolderPaginationRequest{
			OrgID:    uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
			Token:    "invalid token",
			PageSize: 1000,
		}

		//Call RequestFoldersPaginated function
		response, err := folders.RequestFoldersPaginated(req)

		//Assert Error Returned
		assert.Error(t, err)

		//Assert No Folders Returned
		assert.Empty(t, response)
	})

	t.Run("Token points to last folder", func(t *testing.T) {
		//Create Request
		req := &folders.FetchFolderPaginationRequest{
			OrgID:    uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
			Token:    "28c2f7ea-f1f3-4235-aa28-04dc94a59a67",
			PageSize: 1,
		}

		//Call RequestFoldersPaginated function
		response, err := folders.RequestFoldersPaginated(req)

		//Assert No Error Returned
		assert.NoError(t, err)

		//Assert No Folders Returned
		assert.Empty(t, response.Folders)

		//Assert Empty Token Returned
		assert.Empty(t, response.Token)
	})
}
