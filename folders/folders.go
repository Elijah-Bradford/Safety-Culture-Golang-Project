package folders

import (
	"github.com/gofrs/uuid"
)

func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	//Initialise Local Variables
	var (
		err error
		f1  Folder
		//Creates a Slice of pointers to Folders
		fs []*Folder
	)

	//Creates a slice of folders
	f := []Folder{}

	//Call function, assigning result and discarding errors
	r, _ := FetchAllFoldersByOrgID(req.OrgID)

	//Iterate through the results slice
	for k, v := range r {
		//Append the Folder instances to 'f' slice
		f = append(f, *v)
	}

	//Creates a slice of folders
	var fp []*Folder

	//Iterate through the copied results slice
	for k1, v1 := range f {
		//append 'v1' folder numerous times
		fp = append(fp, &v1)
	}

	//Populate fetch folder response pointer initialised with 'Folder' field to fp results slice
	var ffr *FetchFolderResponse
	ffr = &FetchFolderResponse{Folders: fp}

	//Return fetch folder response pointer, and nil error
	return ffr, nil
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
