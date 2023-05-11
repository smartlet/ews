package test

import (
	"fmt"
	"github.com/smartlet/ews"
	"github.com/smartlet/ews/wsdl"
	"testing"
)

func TestGetFolder(t *testing.T) {
	defer dumpFile.Sync()
	rsp, err := service.GetFolder(ews.MakeContext(acc), &wsdl.GetFolderSoapIn{
		GetFolder: &wsdl.GetFolderType{
			FolderShape: &wsdl.FolderResponseShapeType{
				BaseShape: wsdl.DefaultShapeNamesTypeDefault,
			},
			FolderIds: &wsdl.NonEmptyArrayOfBaseFolderIdsType{
				DistinguishedFolderId: []*wsdl.DistinguishedFolderIdType{
					{Id: wsdl.DistinguishedFolderIdNameTypeIinbox},
				},
			},
		},
	}, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.ServerVersionInfo)
	fmt.Println(rsp.GetFolderResponse)
}

func TestGetFolder_error(t *testing.T) {
	defer dumpFile.Sync()

	detail := new(ews.FaultDetail)
	rsp, err := service.GetFolder(ews.MakeContext(acc), &wsdl.GetFolderSoapIn{
		GetFolder: &wsdl.GetFolderType{
			FolderShape: &wsdl.FolderResponseShapeType{
				BaseShape: wsdl.DefaultShapeNamesTypeDefault,
			},
			FolderIds: &wsdl.NonEmptyArrayOfBaseFolderIdsType{
				DistinguishedFolderId: []*wsdl.DistinguishedFolderIdType{
					{Id: wsdl.DistinguishedFolderIdNameTypeIinbox},
				},
			},
		},
	}, detail)
	if err != nil {
		fmt.Println(detail)
		fmt.Println(err)
	} else {
		fmt.Println(rsp.ServerVersionInfo)
		fmt.Println(rsp.GetFolderResponse)
	}

}
