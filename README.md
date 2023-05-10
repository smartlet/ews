# ews

EWS 是一种全面的服务，应用程序可以使用 EWS 来访问几乎所有存储在 Exchange Online、作为 Office 365 一部分的 Exchange Online 或 Exchange 本地邮箱中的信息。 EWS 使用标准 Web 协议，对 Exchange 服务器提供访问权限；诸如 EWS 托管 API 等库会环绕 EWS 操作，提供面向对象的界面。 运行本文中的示例后，你将基本了解可以使用 EWS 执行的操作。

## 支持操作

完整参考: https://learn.microsoft.com/zh-cn/exchange/client-developer/web-service-reference/ews-operations-in-exchange

### 电子数据展示操作
- GetDiscoverySearchConfiguration 操作
- GetHoldOnMailboxes 操作
- GetNonIndexableItemDetails 操作
- GetNonIndexableItemStatistics 操作
- GetSearchableMailboxes 操作
- SearchMailboxes 操作
- SetHoldOnMailboxes 操作

### Exchange 邮箱数据操作

Exchange 邮箱数据操作
- ArchiveItem 操作
- CreateItem 操作
- CopyItem 操作
- DeleteItem 操作
- FindItem 操作
- GetItem 操作
- MarkAllItemsAsRead 操作
- MoveItem 操作
- SendItem 操作
- UpdateItem 操作

Exchange 邮箱数据文件夹操作
- CreateFolder 操作
- CreateFolderPath 操作
- CreateManagedFolder 操作
- CopyFolder 操作
- DeleteFolder 操作
- EmptyFolder 操作
- FindFolder 操作
- GetFolder 操作
- MoveFolder 操作
- UpdateFolder 操作

Exchange 邮箱数据附件操作
- CreateAttachment 操作
- GetAttachment 操作
- DeleteAttachment 操作

Exchange 邮箱提醒操作
- GetReminders 操作
- PerformReminderAction 操作

Exchange 邮箱数据会话操作
- ApplyConversationAction 操作
- FindConversation 操作
- GetConversationItems 操作

Exchange 邮箱数据实用程序操作
- ConvertId 操作
- ExpandDL 操作
- GetUserPhoto 操作
- MarkAsJunk 操作
- ResolveNames 操作
- GetPasswordExpirationDate 操作

### 可用性操作
- GetUserAvailability 操作
- GetRoomLists 操作
- GetRooms 操作
- GetUserOofSettings 操作
- SetUserOofSettings 操作

### 批量传输操作
- UploadItems 操作
- ExportItems 操作

### 代理管理操作
- AddDelegate 操作
- GetDelegate 操作
- UpdateDelegate 操作
- RemoveDelegate 操作

### 收件箱规则操作
- GetInboxRules 操作
- UpdateInboxRules 操作

### 邮件应用管理操作
- DisableApp 操作
- GetAppManifests 操作
- GetAppMarketplaceUrl 操作
- GetClientAccessToken 操作
- InstallApp 操作
- UninstallApp 操作

### 邮件提示操作
- GetMailTips 操作

### 邮件跟踪操作
- FindMessageTrackingReport 操作
- GetMessageTrackingReport 操作

### 通知操作
- GetEvents 操作
- GetStreamingEvents 操作
- Subscribe 操作
- Unsubscribe 操作

### 角色操作
- FindPeople 操作
- GetPersona 操作

### 保留策略操作
- GetUserRetentionPolicyTags 操作

### 服务配置操作
- GetServiceConfiguration 操作

### 共享操作
- CreateItem 操作
- GetSharingFolder 操作
- GetSharingMetadata 操作
- RefreshSharingFolder 操作

### 同步操作
- SyncFolderHierarchy 操作
- SyncFolderItems 操作

### 时区操作
- GetServerTimeZones 操作

### 统一消息操作
- DisconnectPhoneCall 操作
- GetPhoneCallInformation 操作
- PlayOnPhone 操作

### 统一联系人存储操作
- AddNewImContactToGroup 操作
- AddImContactToGroup 操作
- AddImGroup 操作
- AddNewTelUriContactToGroup 操作
- AddDistributionGroupToImList 操作
- GetImItemList 操作
- GetImItems 操作
- RemoveContactFromImList 操作
- RemoveImContactFromGroup 操作
- RemoveDistributionGroupFromImList 操作
- RemoveImGroup 操作
- SetImGroup 操作

### 用户配置操作
- CreateUserConfiguration 操作
- DeleteUserConfiguration 操作
- GetUserConfiguration 操作
- UpdateUserConfiguration 操作


## 完整API
```

type ExchangeServicePortType interface {
	ResolveNames(ctx context.Context, in *ResolveNamesSoapIn, detail error) (*ResolveNamesSoapOut, error)
	ExpandDL(ctx context.Context, in *ExpandDLSoapIn, detail error) (*ExpandDLSoapOut, error)
	GetServerTimeZones(ctx context.Context, in *GetServerTimeZonesSoapIn, detail error) (*GetServerTimeZonesSoapOut, error)
	FindFolder(ctx context.Context, in *FindFolderSoapIn, detail error) (*FindFolderSoapOut, error)
	FindItem(ctx context.Context, in *FindItemSoapIn, detail error) (*FindItemSoapOut, error)
	GetFolder(ctx context.Context, in *GetFolderSoapIn, detail error) (*GetFolderSoapOut, error)
	UploadItems(ctx context.Context, in *UploadItemsSoapIn, detail error) (*UploadItemsSoapOut, error)
	ExportItems(ctx context.Context, in *ExportItemsSoapIn, detail error) (*ExportItemsSoapOut, error)
	ConvertId(ctx context.Context, in *ConvertIdSoapIn, detail error) (*ConvertIdSoapOut, error)
	CreateFolder(ctx context.Context, in *CreateFolderSoapIn, detail error) (*CreateFolderSoapOut, error)
	CreateFolderPath(ctx context.Context, in *CreateFolderPathSoapIn, detail error) (*CreateFolderPathSoapOut, error)
	DeleteFolder(ctx context.Context, in *DeleteFolderSoapIn, detail error) (*DeleteFolderSoapOut, error)
	EmptyFolder(ctx context.Context, in *EmptyFolderSoapIn, detail error) (*EmptyFolderSoapOut, error)
	UpdateFolder(ctx context.Context, in *UpdateFolderSoapIn, detail error) (*UpdateFolderSoapOut, error)
	MoveFolder(ctx context.Context, in *MoveFolderSoapIn, detail error) (*MoveFolderSoapOut, error)
	CopyFolder(ctx context.Context, in *CopyFolderSoapIn, detail error) (*CopyFolderSoapOut, error)
	Subscribe(ctx context.Context, in *SubscribeSoapIn, detail error) (*SubscribeSoapOut, error)
	Unsubscribe(ctx context.Context, in *UnsubscribeSoapIn, detail error) (*UnsubscribeSoapOut, error)
	GetEvents(ctx context.Context, in *GetEventsSoapIn, detail error) (*GetEventsSoapOut, error)
	GetStreamingEvents(ctx context.Context, in *GetStreamingEventsSoapIn, detail error) (*GetStreamingEventsSoapOut, error)
	SyncFolderHierarchy(ctx context.Context, in *SyncFolderHierarchySoapIn, detail error) (*SyncFolderHierarchySoapOut, error)
	SyncFolderItems(ctx context.Context, in *SyncFolderItemsSoapIn, detail error) (*SyncFolderItemsSoapOut, error)
	CreateManagedFolder(ctx context.Context, in *CreateManagedFolderSoapIn, detail error) (*CreateManagedFolderSoapOut, error)
	GetItem(ctx context.Context, in *GetItemSoapIn, detail error) (*GetItemSoapOut, error)
	CreateItem(ctx context.Context, in *CreateItemSoapIn, detail error) (*CreateItemSoapOut, error)
	DeleteItem(ctx context.Context, in *DeleteItemSoapIn, detail error) (*DeleteItemSoapOut, error)
	UpdateItem(ctx context.Context, in *UpdateItemSoapIn, detail error) (*UpdateItemSoapOut, error)
	UpdateItemInRecoverableItems(ctx context.Context, in *UpdateItemInRecoverableItemsSoapIn, detail error) (*UpdateItemInRecoverableItemsSoapOut, error)
	SendItem(ctx context.Context, in *SendItemSoapIn, detail error) (*SendItemSoapOut, error)
	MoveItem(ctx context.Context, in *MoveItemSoapIn, detail error) (*MoveItemSoapOut, error)
	CopyItem(ctx context.Context, in *CopyItemSoapIn, detail error) (*CopyItemSoapOut, error)
	ArchiveItem(ctx context.Context, in *ArchiveItemSoapIn, detail error) (*ArchiveItemSoapOut, error)
	CreateAttachment(ctx context.Context, in *CreateAttachmentSoapIn, detail error) (*CreateAttachmentSoapOut, error)
	DeleteAttachment(ctx context.Context, in *DeleteAttachmentSoapIn, detail error) (*DeleteAttachmentSoapOut, error)
	GetAttachment(ctx context.Context, in *GetAttachmentSoapIn, detail error) (*GetAttachmentSoapOut, error)
	GetClientAccessToken(ctx context.Context, in *GetClientAccessTokenSoapIn, detail error) (*GetClientAccessTokenSoapOut, error)
	GetDelegate(ctx context.Context, in *GetDelegateSoapIn, detail error) (*GetDelegateSoapOut, error)
	AddDelegate(ctx context.Context, in *AddDelegateSoapIn, detail error) (*AddDelegateSoapOut, error)
	RemoveDelegate(ctx context.Context, in *RemoveDelegateSoapIn, detail error) (*RemoveDelegateSoapOut, error)
	UpdateDelegate(ctx context.Context, in *UpdateDelegateSoapIn, detail error) (*UpdateDelegateSoapOut, error)
	CreateUserConfiguration(ctx context.Context, in *CreateUserConfigurationSoapIn, detail error) (*CreateUserConfigurationSoapOut, error)
	DeleteUserConfiguration(ctx context.Context, in *DeleteUserConfigurationSoapIn, detail error) (*DeleteUserConfigurationSoapOut, error)
	GetUserConfiguration(ctx context.Context, in *GetUserConfigurationSoapIn, detail error) (*GetUserConfigurationSoapOut, error)
	GetSpecificUserConfiguration(ctx context.Context, in *GetSpecificUserConfigurationSoapIn, detail error) (*GetSpecificUserConfigurationSoapOut, error)
	UpdateUserConfiguration(ctx context.Context, in *UpdateUserConfigurationSoapIn, detail error) (*UpdateUserConfigurationSoapOut, error)
	GetUserAvailability(ctx context.Context, in *GetUserAvailabilitySoapIn, detail error) (*GetUserAvailabilitySoapOut, error)
	GetUserOofSettings(ctx context.Context, in *GetUserOofSettingsSoapIn, detail error) (*GetUserOofSettingsSoapOut, error)
	SetUserOofSettings(ctx context.Context, in *SetUserOofSettingsSoapIn, detail error) (*SetUserOofSettingsSoapOut, error)
	GetServiceConfiguration(ctx context.Context, in *GetServiceConfigurationSoapIn, detail error) (*GetServiceConfigurationSoapOut, error)
	GetMailTips(ctx context.Context, in *GetMailTipsSoapIn, detail error) (*GetMailTipsSoapOut, error)
	PlayOnPhone(ctx context.Context, in *PlayOnPhoneSoapIn, detail error) (*PlayOnPhoneSoapOut, error)
	GetPhoneCallInformation(ctx context.Context, in *GetPhoneCallInformationSoapIn, detail error) (*GetPhoneCallInformationSoapOut, error)
	DisconnectPhoneCall(ctx context.Context, in *DisconnectPhoneCallSoapIn, detail error) (*DisconnectPhoneCallSoapOut, error)
	GetSharingMetadata(ctx context.Context, in *GetSharingMetadataSoapIn, detail error) (*GetSharingMetadataSoapOut, error)
	RefreshSharingFolder(ctx context.Context, in *RefreshSharingFolderSoapIn, detail error) (*RefreshSharingFolderSoapOut, error)
	GetSharingFolder(ctx context.Context, in *GetSharingFolderSoapIn, detail error) (*GetSharingFolderSoapOut, error)
	SetTeamMailbox(ctx context.Context, in *SetTeamMailboxSoapIn, detail error) (*SetTeamMailboxSoapOut, error)
	UnpinTeamMailbox(ctx context.Context, in *UnpinTeamMailboxSoapIn, detail error) (*UnpinTeamMailboxSoapOut, error)
	GetRoomLists(ctx context.Context, in *GetRoomListsSoapIn, detail error) (*GetRoomListsSoapOut, error)
	GetRooms(ctx context.Context, in *GetRoomsSoapIn, detail error) (*GetRoomsSoapOut, error)
	FindMessageTrackingReport(ctx context.Context, in *FindMessageTrackingReportSoapIn, detail error) (*FindMessageTrackingReportSoapOut, error)
	GetMessageTrackingReport(ctx context.Context, in *GetMessageTrackingReportSoapIn, detail error) (*GetMessageTrackingReportSoapOut, error)
	FindConversation(ctx context.Context, in *FindConversationSoapIn, detail error) (*FindConversationSoapOut, error)
	ApplyConversationAction(ctx context.Context, in *ApplyConversationActionSoapIn, detail error) (*ApplyConversationActionSoapOut, error)
	GetConversationItems(ctx context.Context, in *GetConversationItemsSoapIn, detail error) (*GetConversationItemsSoapOut, error)
	FindPeople(ctx context.Context, in *FindPeopleSoapIn, detail error) (*FindPeopleSoapOut, error)
	FindTags(ctx context.Context, in *FindTagsSoapIn, detail error) (*FindTagsSoapOut, error)
	AddTag(ctx context.Context, in *AddTagSoapIn, detail error) (*AddTagSoapOut, error)
	HideTag(ctx context.Context, in *HideTagSoapIn, detail error) (*HideTagSoapOut, error)
	GetPersona(ctx context.Context, in *GetPersonaSoapIn, detail error) (*GetPersonaSoapOut, error)
	GetInboxRules(ctx context.Context, in *GetInboxRulesSoapIn, detail error) (*GetInboxRulesSoapOut, error)
	UpdateInboxRules(ctx context.Context, in *UpdateInboxRulesSoapIn, detail error) (*UpdateInboxRulesSoapOut, error)
	GetPasswordExpirationDate(ctx context.Context, in *GetPasswordExpirationDateSoapIn, detail error) (*GetPasswordExpirationDateSoapOut, error)
	GetSearchableMailboxes(ctx context.Context, in *GetSearchableMailboxesSoapIn, detail error) (*GetSearchableMailboxesSoapOut, error)
	SearchMailboxes(ctx context.Context, in *SearchMailboxesSoapIn, detail error) (*SearchMailboxesSoapOut, error)
	GetDiscoverySearchConfiguration(ctx context.Context, in *GetDiscoverySearchConfigurationSoapIn, detail error) (*GetDiscoverySearchConfigurationSoapOut, error)
	GetHoldOnMailboxes(ctx context.Context, in *GetHoldOnMailboxesSoapIn, detail error) (*GetHoldOnMailboxesSoapOut, error)
	SetHoldOnMailboxes(ctx context.Context, in *SetHoldOnMailboxesSoapIn, detail error) (*SetHoldOnMailboxesSoapOut, error)
	GetNonIndexableItemStatistics(ctx context.Context, in *GetNonIndexableItemStatisticsSoapIn, detail error) (*GetNonIndexableItemStatisticsSoapOut, error)
	GetNonIndexableItemDetails(ctx context.Context, in *GetNonIndexableItemDetailsSoapIn, detail error) (*GetNonIndexableItemDetailsSoapOut, error)
	MarkAllItemsAsRead(ctx context.Context, in *MarkAllItemsAsReadSoapIn, detail error) (*MarkAllItemsAsReadSoapOut, error)
	MarkAsJunk(ctx context.Context, in *MarkAsJunkSoapIn, detail error) (*MarkAsJunkSoapOut, error)
	ReportMessage(ctx context.Context, in *ReportMessageSoapIn, detail error) (*ReportMessageSoapOut, error)
	GetAppManifests(ctx context.Context, in *GetAppManifestsSoapIn, detail error) (*GetAppManifestsSoapOut, error)
	AddNewImContactToGroup(ctx context.Context, in *AddNewImContactToGroupSoapIn, detail error) (*AddNewImContactToGroupSoapOut, error)
	AddNewTelUriContactToGroup(ctx context.Context, in *AddNewTelUriContactToGroupSoapIn, detail error) (*AddNewTelUriContactToGroupSoapOut, error)
	AddImContactToGroup(ctx context.Context, in *AddImContactToGroupSoapIn, detail error) (*AddImContactToGroupSoapOut, error)
	RemoveImContactFromGroup(ctx context.Context, in *RemoveImContactFromGroupSoapIn, detail error) (*RemoveImContactFromGroupSoapOut, error)
	AddImGroup(ctx context.Context, in *AddImGroupSoapIn, detail error) (*AddImGroupSoapOut, error)
	AddDistributionGroupToImList(ctx context.Context, in *AddDistributionGroupToImListSoapIn, detail error) (*AddDistributionGroupToImListSoapOut, error)
	GetImItemList(ctx context.Context, in *GetImItemListSoapIn, detail error) (*GetImItemListSoapOut, error)
	GetImItems(ctx context.Context, in *GetImItemsSoapIn, detail error) (*GetImItemsSoapOut, error)
	RemoveContactFromImList(ctx context.Context, in *RemoveContactFromImListSoapIn, detail error) (*RemoveContactFromImListSoapOut, error)
	RemoveDistributionGroupFromImList(ctx context.Context, in *RemoveDistributionGroupFromImListSoapIn, detail error) (*RemoveDistributionGroupFromImListSoapOut, error)
	RemoveImGroup(ctx context.Context, in *RemoveImGroupSoapIn, detail error) (*RemoveImGroupSoapOut, error)
	SetImGroup(ctx context.Context, in *SetImGroupSoapIn, detail error) (*SetImGroupSoapOut, error)
	SetImListMigrationCompleted(ctx context.Context, in *SetImListMigrationCompletedSoapIn, detail error) (*SetImListMigrationCompletedSoapOut, error)
	GetUserRetentionPolicyTags(ctx context.Context, in *GetUserRetentionPolicyTagsSoapIn, detail error) (*GetUserRetentionPolicyTagsSoapOut, error)
	InstallApp(ctx context.Context, in *InstallAppSoapIn, detail error) (*InstallAppSoapOut, error)
	UpdateExtensionUsage(ctx context.Context, in *UpdateExtensionUsageSoapIn, detail error) (*UpdateExtensionUsageSoapOut, error)
	UninstallApp(ctx context.Context, in *UninstallAppSoapIn, detail error) (*UninstallAppSoapOut, error)
	DisableApp(ctx context.Context, in *DisableAppSoapIn, detail error) (*DisableAppSoapOut, error)
	GetAppMarketplaceUrl(ctx context.Context, in *GetAppMarketplaceUrlSoapIn, detail error) (*GetAppMarketplaceUrlSoapOut, error)
	FindAvailableMeetingTimes(ctx context.Context, in *FindAvailableMeetingTimesSoapIn, detail error) (*FindAvailableMeetingTimesSoapOut, error)
	FindMeetingTimeCandidates(ctx context.Context, in *FindMeetingTimeCandidatesSoapIn, detail error) (*FindMeetingTimeCandidatesSoapOut, error)
	GetUserPhoto(ctx context.Context, in *GetUserPhotoSoapIn, detail error) (*GetUserPhotoSoapOut, error)
	SetUserPhoto(ctx context.Context, in *SetUserPhotoSoapIn, detail error) (*SetUserPhotoSoapOut, error)
	GetMeetingSpace(ctx context.Context, in *GetMeetingSpaceSoapIn, detail error) (*GetMeetingSpaceSoapOut, error)
	DeleteMeetingSpace(ctx context.Context, in *DeleteMeetingSpaceSoapIn, detail error) (*DeleteMeetingSpaceSoapOut, error)
	UpdateMeetingSpace(ctx context.Context, in *UpdateMeetingSpaceSoapIn, detail error) (*UpdateMeetingSpaceSoapOut, error)
	CreateMeetingSpace(ctx context.Context, in *CreateMeetingSpaceSoapIn, detail error) (*CreateMeetingSpaceSoapOut, error)
	FindMeetingSpaceByJoinUrl(ctx context.Context, in *FindMeetingSpaceByJoinUrlSoapIn, detail error) (*FindMeetingSpaceByJoinUrlSoapOut, error)
	GetMeetingInstance(ctx context.Context, in *GetMeetingInstanceSoapIn, detail error) (*GetMeetingInstanceSoapOut, error)
	DeleteMeetingInstance(ctx context.Context, in *DeleteMeetingInstanceSoapIn, detail error) (*DeleteMeetingInstanceSoapOut, error)
	UpdateMeetingInstance(ctx context.Context, in *UpdateMeetingInstanceSoapIn, detail error) (*UpdateMeetingInstanceSoapOut, error)
	CreateMeetingInstance(ctx context.Context, in *CreateMeetingInstanceSoapIn, detail error) (*CreateMeetingInstanceSoapOut, error)
	StartSearchSession(ctx context.Context, in *StartSearchSessionSoapIn, detail error) (*StartSearchSessionSoapOut, error)
	GetSearchSuggestions(ctx context.Context, in *GetSearchSuggestionsSoapIn, detail error) (*GetSearchSuggestionsSoapOut, error)
	DeleteSearchSuggestion(ctx context.Context, in *DeleteSearchSuggestionSoapIn, detail error) (*DeleteSearchSuggestionSoapOut, error)
	ExecuteSearch(ctx context.Context, in *ExecuteSearchSoapIn, detail error) (*ExecuteSearchSoapOut, error)
	EndSearchSession(ctx context.Context, in *EndSearchSessionSoapIn, detail error) (*EndSearchSessionSoapOut, error)
	GetLastPrivateCatalogUpdate(ctx context.Context, in *GetLastPrivateCatalogUpdateSoapIn, detail error) (*GetLastPrivateCatalogUpdateSoapOut, error)
	GetPrivateCatalogAddIns(ctx context.Context, in *GetPrivateCatalogAddInsSoapIn, detail error) (*GetPrivateCatalogAddInsSoapOut, error)
}

```

## 使用示例
```
var acc = ews.NewAccountSession(
	"赵大海",
	"哇哈哈",
	"zhaodahai",
	"https://wwww.wahaha.com/EWS/Exchange.asmx",
)

var trace, _ = os.Create(`/tmp/trace.log`)

var httpCli = http.NewHTTPRoundTripper(new(http.Config))
var ntlmCli = http.NewNTLMRoundTripper(
	httpCli,
	kits.NewMemoryAuthorizer(),
	kits.NewMemoryCredential(map[string][2]string{
		acc.GetId(): {
			"zhaodahai",
			os.Getenv("PASSWORD"),
		},
	}))
var dumpCli = http.NewDumpRoundTripper(ntlmCli, trace)
var soapCli = soap.NewSOAPClient(dumpCli)
var service = wsdl.NewExchangeServicePortType(soapCli)

func TestGetFolder(t *testing.T) {
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
	data, err := xml.MarshalIndent(rsp, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", data)
}

```