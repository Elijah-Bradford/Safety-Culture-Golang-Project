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
