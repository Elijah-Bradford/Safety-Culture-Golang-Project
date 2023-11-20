package folders

import (
	"errors"

	"github.com/gofrs/uuid"
)

// Pagination Types
type FetchFolderPaginationRequest struct {
	OrgID    uuid.UUID
	Token    string
	PageSize int
}

type FetchFolderPaginationResponse struct {
	Folders []*Folder
	Token   string
}

// Pagination Functions
func RequestFoldersPaginated(req *FetchFolderPaginationRequest) (*FetchFolderPaginationResponse, error) {
	//Initialise Local Variables
	var (
		err                              error
		StartPos                         int
		EndPos                           int
		Folders                          []*Folder
		Token                            string
		TokenFound                       bool
		NewToken                         string
		PagedFolders                     []*Folder
		pagination_folder_fetch_response *FetchFolderPaginationResponse
	)

	//Populate Local Variables
	Folders, err = FetchAllFoldersByOrgID(req.OrgID)
	StartPos = 0

	//Determine Start Position from Supplied Token
	TokenFound = false
	if req.Token != "" {
		for i := 0; i <= len(Folders); i++ {
			if Folders[i].Id.String() == Token {
				StartPos = i + 1
				TokenFound = true
				break
			}
		}
		//Invalid Token
		if !TokenFound {
			return nil, errors.New("Invalid Token Supplied.")
		}
	}

	//Determine End Position from Page Size
	EndPos = StartPos + req.PageSize

	//Generate Token
	NewToken = Folders[EndPos].Id.String()

	//Create Paged Folders
	PagedFolders = Folders[StartPos:EndPos]

	//Create and Populate Response Structure
	pagination_folder_fetch_response = &FetchFolderPaginationResponse{
		Folders: PagedFolders,
		Token:   NewToken,
	}

	//Return Response
	return pagination_folder_fetch_response, err

}
