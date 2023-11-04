package main

import (
	//From go standard library and provides input/output functionality.
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folders"

	//Provides functionality for working with UUIDs
	"github.com/gofrs/uuid"
)

func main() {
	//Initialises a pointer to an instance of the 'FetchFolderRequest' struct
	req := &folders.FetchFolderRequest{
		//Sets Origin ID to the parsed UUID of the Default UUID string
		OrgID: uuid.FromStringOrNil(folders.DefaultOrgID),
	}

	//Initialises and populates variables to store the result and errors from the GetAllFolders request
	res, err := folders.GetAllFolders(req)

	//If an error occurs
	if err != nil {
		//Print the error
		fmt.Printf("%v", err)
		return
	}

	//Print the results of the 'get all folders' request
	folders.PrettyPrint(res)
}
