package folders

import (
	"github.com/gofrs/uuid"
)

func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	//Initialise Local Variables
	var (
		err error
	)

	//Call function, assigning result and errors returned
	r, err := FetchAllFoldersByOrgID(req.OrgID)

	//Populate fetch folder response pointer initialised with 'Folder' field to r results slice
	var folder_fetch_response *FetchFolderResponse
	folder_fetch_response = &FetchFolderResponse{Folders: r}

	//Return fetch folder response pointer, and nil error
	return folder_fetch_response, err
}

func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	//Fetches Sample folder data
	folders := GetSampleData()

	//Initialise empty slice to store folder pointers
	resFolder := []*Folder{}

	//Iterate through fetched folders
	for _, folder := range folders {
		//Find folders with matching OrgID
		if folder.OrgId == orgID {
			//Append matching OrgID folders to results slice
			resFolder = append(resFolder, folder)
		}
	}

	//Return folder results slice, and nil error
	return resFolder, nil
}
