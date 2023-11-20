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
	Token = req.Token

	//Check for errors fetching folders
	if err != nil {
		return nil, err
	}

	//Determine Start Position from Supplied Token
	TokenFound = false
	if req.Token != "" {
		for i := 0; i < len(Folders); i++ {
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
	EndPos = StartPos + req.PageSize - 1

	//Check EndPos within Bounds
	if EndPos >= len(Folders) {
		EndPos = len(Folders) - 1
	}

	//Generate Token
	NewToken = Folders[EndPos].Id.String()

	//Generate empty token if Final Folder displayed
	if EndPos == len(Folders)-1 {
		NewToken = ""
	}

	//Create Paged Folders
	PagedFolders = Folders[StartPos : EndPos+1]

	//Create and Populate Response Structure
	pagination_folder_fetch_response = &FetchFolderPaginationResponse{
		Folders: PagedFolders,
		Token:   NewToken,
	}

	//Return Response
	return pagination_folder_fetch_response, err
}

/*
Pagination Implementation Explanation:

I chose to implement pagination by calling the existing 'FetchAllFoldersByOrgID' as this function has already been built and tested.
I implemented new request and response types to enable the token to become part of the folder request.

I chose to make the token the id of the last folder of the request, as this is a unique string that is unlikely to be known by an unauthorised source.
This allows the function to locate the folder with the matching id (token), and start the request from the next folder.

I also chose to make 'PageSize' part of the request, to allow multiple different numbers of folders to be returned.

*/
