package test

import (
	"fmt"
	"github.com/smartlet/ews"
	"testing"
)

func TestUpdateFolder(t *testing.T) {
	TestFindFolder(t)
	rsp, err := service.UpdateFolder(ews.MakeContext(testSess), &ews.UpdateFolderSoapIn{
		UpdateFolder: &ews.UpdateFolderType{
			FolderChanges: &ews.NonEmptyArrayOfFolderChangesType{
				FolderChange: []*ews.FolderChangeType{
					{
						FolderId: &ews.FolderIdType{
							Id: MyFirstLevelFolderId,
						},
						Updates: &ews.NonEmptyArrayOfFolderChangeDescriptionsType{
							SetFolderField: []*ews.SetFolderFieldType{
								{
									FieldURI: &ews.PathToUnindexedFieldType{
										FieldURI: ews.UnindexedFieldURITypeFolderDisplayName,
									},
									Folder: &ews.FolderType{
										DisplayName: "MyFirstLevelFolder2",
									},
								},
							},
						},
					},
				},
			},
		},
	}, nil)

	if err != nil {
		panic(err)
	}

	fmt.Println(rsp.UpdateFolderResponse.ResponseMessages.UpdateFolderResponseMessage[0].Folders.Folder[0].FolderId)
}
