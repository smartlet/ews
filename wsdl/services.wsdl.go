package wsdl

import (
	"context"
	"encoding/xml"
)

// XsDuration https://www.w3.org/TR/xmlschema-2/#duration
type XsDuration string

// XsDateTime https://www.w3.org/TR/xmlschema-2/#dateTime
type XsDateTime string

// XsTime https://www.w3.org/TR/xmlschema-2/#time
type XsTime string

// XsDate https://www.w3.org/TR/xmlschema-2/#date
type XsDate string

// XsGYearMonth https://www.w3.org/TR/xmlschema-2/#gYearMonth
type XsGYearMonth string

// XsGYear https://www.w3.org/TR/xmlschema-2/#gYear
type XsGYear string

// XsGMonthDay https://www.w3.org/TR/xmlschema-2/#gMonthDay
type XsGMonthDay string

// XsGDay https://www.w3.org/TR/xmlschema-2/#gDay
type XsGDay string

// XsGMonth https://www.w3.org/TR/xmlschema-2/#gMonth
type XsGMonth string

// XsBoolean https://www.w3.org/TR/xmlschema-2/#boolean
type XsBoolean bool

// XsBase64Binary https://www.w3.org/TR/xmlschema-2/#base64Binary
type XsBase64Binary []byte

// XsHexBinary https://www.w3.org/TR/xmlschema-2/#hexBinary
type XsHexBinary []byte

// XsFloat https://www.w3.org/TR/xmlschema-2/#float
type XsFloat float32

// XsDouble https://www.w3.org/TR/xmlschema-2/#double
type XsDouble float64

// XsAnyURI https://www.w3.org/TR/xmlschema-2/#anyURI
type XsAnyURI string

// XsQName https://www.w3.org/TR/xmlschema-2/#QName
type XsQName string

// XsNOTATION https://www.w3.org/TR/xmlschema-2/#NOTATION
type XsNOTATION string

// XsString https://www.w3.org/TR/xmlschema-2/#string
type XsString string

// XsNormalizedString https://www.w3.org/TR/xmlschema-2/#normalizedString
type XsNormalizedString string

// XsToken https://www.w3.org/TR/xmlschema-2/#token
type XsToken string

// XsLanguage https://www.w3.org/TR/xmlschema-2/#language
type XsLanguage string

// XsName https://www.w3.org/TR/xmlschema-2/#Name
type XsName string

// XsNCName https://www.w3.org/TR/xmlschema-2/#NCName
type XsNCName string

// XsID https://www.w3.org/TR/xmlschema-2/#ID
type XsID string

// XsIDREF https://www.w3.org/TR/xmlschema-2/#IDREF
type XsIDREF string

// XsIDREFS https://www.w3.org/TR/xmlschema-2/#IDREFS
type XsIDREFS string

// XsENTITY https://www.w3.org/TR/xmlschema-2/#ENTITY
type XsENTITY string

// XsENTITIES https://www.w3.org/TR/xmlschema-2/#ENTITIES
type XsENTITIES string

// XsNMTOKEN https://www.w3.org/TR/xmlschema-2/#NMTOKEN
type XsNMTOKEN string

// XsNMTOKENS https://www.w3.org/TR/xmlschema-2/#NMTOKENS
type XsNMTOKENS string

// XsDecimal https://www.w3.org/TR/xmlschema-2/#decimal
type XsDecimal float64

// XsInteger https://www.w3.org/TR/xmlschema-2/#integer
type XsInteger int64

// XsNonPositiveInteger https://www.w3.org/TR/xmlschema-2/#nonPositiveInteger
type XsNonPositiveInteger int64

// XsNegativeInteger https://www.w3.org/TR/xmlschema-2/#negativeInteger
type XsNegativeInteger int64

// XsLong https://www.w3.org/TR/xmlschema-2/#long
type XsLong int64

// XsInt https://www.w3.org/TR/xmlschema-2/#int
type XsInt int32

// XsShort https://www.w3.org/TR/xmlschema-2/#short
type XsShort int16

// XsByte https://www.w3.org/TR/xmlschema-2/#byte
type XsByte int8

// XsNonNegativeInteger https://www.w3.org/TR/xmlschema-2/#nonNegativeInteger
type XsNonNegativeInteger int64

// XsUnsignedLong https://www.w3.org/TR/xmlschema-2/#unsignedLong
type XsUnsignedLong uint64

// XsUnsignedInt https://www.w3.org/TR/xmlschema-2/#unsignedInt
type XsUnsignedInt uint32

// XsUnsignedShort https://www.w3.org/TR/xmlschema-2/#unsignedShort
type XsUnsignedShort uint16

// XsUnsignedByte https://www.w3.org/TR/xmlschema-2/#unsignedByte
type XsUnsignedByte uint8

// XsPositiveInteger https://www.w3.org/TR/xmlschema-2/#positiveInteger
type XsPositiveInteger int64

// SOAPClient soap client interface
type SOAPClient interface {
	Call(ctx context.Context, soapAction string, inputHeader, inputBody, outputHeader, outputBody any) error
}

const (
	XmlnsS = "http://schemas.xmlsoap.org/soap/envelope/"
	XmlnsT = "https://schemas.microsoft.com/exchange/services/2006/types"
	XmlnsM = "https://schemas.microsoft.com/exchange/services/2006/messages"
)

var XmlnsPrefix = map[string]string{
	XmlnsS: "s",
	XmlnsM: "m",
	XmlnsT: "t",
}

type Envelope struct {
	XMLName xml.Name `xml:"s:Envelope"`
	XmlnsS  string   `xml:"xmlns:s,attr"`
	XmlnsM  string   `xml:"xmlns:m,attr"`
	XmlnsT  string   `xml:"xmlns:t,attr"`
	Header  any      `xml:"s:Header,omitempty"` // 必须没有XMLName. 在binding使用匿名struct实现!
	Body    any      `xml:"s:Body,omitempty"`   // 必须没有XMLName. 在binding使用匿名struct实现!
}

type Fault struct {
	FaultCode   string `xml:"faultcode,omitempty"`
	FaultString string `xml:"faultstring,omitempty"`
	FaultActor  string `xml:"faultactor,omitempty"`
	Detail      any    `xml:"detail,omitempty"`
}

func (f Fault) Error() string {
	return f.FaultCode + ":" + f.FaultString
}

type ReminderMinutesBeforeStartTypeUnion0 XsInt

type ReminderMinutesBeforeStartTypeUnion1 XsInt

type PropertyTagTypeUnion0 XsString

type UserConfigurationPropertyTypeItem XsString

const (
	UserConfigurationPropertyTypeItemId         = "Id"
	UserConfigurationPropertyTypeItemDictionary = "Dictionary"
	UserConfigurationPropertyTypeItemXmlData    = "XmlData"
	UserConfigurationPropertyTypeItemBinaryData = "BinaryData"
	UserConfigurationPropertyTypeItemAll        = "All"
)

type FreeBusyViewTypeItem XsString

const (
	FreeBusyViewTypeItemNone           = "None"
	FreeBusyViewTypeItemMergedOnly     = "MergedOnly"
	FreeBusyViewTypeItemFreeBusy       = "FreeBusy"
	FreeBusyViewTypeItemFreeBusyMerged = "FreeBusyMerged"
	FreeBusyViewTypeItemDetailed       = "Detailed"
	FreeBusyViewTypeItemDetailedMerged = "DetailedMerged"
)

type ServiceConfigurationTypeItem XsString

const (
	ServiceConfigurationTypeItemMailTips                       = "MailTips"
	ServiceConfigurationTypeItemUnifiedMessagingConfiguration  = "UnifiedMessagingConfiguration"
	ServiceConfigurationTypeItemProtectionRules                = "ProtectionRules"
	ServiceConfigurationTypeItemPolicyNudges                   = "PolicyNudges"
	ServiceConfigurationTypeItemSharePointURLs                 = "SharePointURLs"
	ServiceConfigurationTypeItemOfficeIntegrationConfiguration = "OfficeIntegrationConfiguration"
)

type ValueType XsString

type NameType XsString

type PriorityType XsInt

type RefreshIntervalType XsInt

type MailTipTypesItem XsString

const (
	MailTipTypesItemAll                     = "All"
	MailTipTypesItemOutOfOfficeMessage      = "OutOfOfficeMessage"
	MailTipTypesItemMailboxFullStatus       = "MailboxFullStatus"
	MailTipTypesItemCustomMailTip           = "CustomMailTip"
	MailTipTypesItemExternalMemberCount     = "ExternalMemberCount"
	MailTipTypesItemTotalMemberCount        = "TotalMemberCount"
	MailTipTypesItemMaxMessageSize          = "MaxMessageSize"
	MailTipTypesItemDeliveryRestriction     = "DeliveryRestriction"
	MailTipTypesItemModerationStatus        = "ModerationStatus"
	MailTipTypesItemInvalidRecipient        = "InvalidRecipient"
	MailTipTypesItemScope                   = "Scope"
	MailTipTypesItemRecipientSuggestions    = "RecipientSuggestions"
	MailTipTypesItemPreferAccessibleContent = "PreferAccessibleContent"
)

type WarmupOptionsTypeItem XsString

const (
	WarmupOptionsTypeItemNone        = "None"
	WarmupOptionsTypeItemSuggestions = "Suggestions"
	WarmupOptionsTypeItemResults     = "Results"
	WarmupOptionsTypeItemAll         = "All"
)

type SuggestionKindTypeItem XsString

const (
	SuggestionKindTypeItemNone         = "None"
	SuggestionKindTypeItemKeywords     = "Keywords"
	SuggestionKindTypeItemPeople       = "People"
	SuggestionKindTypeItemHashtags     = "Hashtags"
	SuggestionKindTypeItemQueryHistory = "QueryHistory"
	SuggestionKindTypeItemAll          = "All"
)

type SearchScopeArchivesTypeItem XsString

const (
	SearchScopeArchivesTypeItemMainArchive = "MainArchive"
	SearchScopeArchivesTypeItemAuxArchive  = "AuxArchive"
	SearchScopeArchivesTypeItemAll         = "All"
)

type SearchScopeGroupsTypeItem XsString

const (
	SearchScopeGroupsTypeItemMyGroups = "MyGroups"
)

type OneDriveViewTypeItem XsString

const (
	OneDriveViewTypeItemNone         = "None"
	OneDriveViewTypeItemSharedWithMe = "SharedWithMe"
	OneDriveViewTypeItemMyDocuments  = "MyDocuments"
	OneDriveViewTypeItemRecycleBin   = "RecycleBin"
)

type DelveViewTypeItem XsString

const (
	DelveViewTypeItemNone  = "None"
	DelveViewTypeItemFiles = "Files"
)

type ItemTypesFilterTypeItem XsString

const (
	ItemTypesFilterTypeItemNone              = "None"
	ItemTypesFilterTypeItemMailItems         = "MailItems"
	ItemTypesFilterTypeItemMailConversations = "MailConversations"
	ItemTypesFilterTypeItemCalendarItems     = "CalendarItems"
	ItemTypesFilterTypeItemContacts          = "Contacts"
	ItemTypesFilterTypeItemOneDriveItems     = "OneDriveItems"
	ItemTypesFilterTypeItemFileItems         = "FileItems"
	ItemTypesFilterTypeItemDelveItems        = "DelveItems"
	ItemTypesFilterTypeItemMessageItems      = "MessageItems"
)

type NonEmptyStringType XsString

type MailboxTypeType XsString

const (
	MailboxTypeTypeUnknown         = "Unknown"
	MailboxTypeTypeOneOff          = "OneOff"
	MailboxTypeTypeMailbox         = "Mailbox"
	MailboxTypeTypePublicDL        = "PublicDL"
	MailboxTypeTypePrivateDL       = "PrivateDL"
	MailboxTypeTypeContact         = "Contact"
	MailboxTypeTypePublicFolder    = "PublicFolder"
	MailboxTypeTypeGroupMailbox    = "GroupMailbox"
	MailboxTypeTypeImplicitContact = "ImplicitContact"
	MailboxTypeTypeUser            = "User"
)

type DistinguishedFolderIdNameType XsString

const (
	DistinguishedFolderIdNameTypeCcalendar                              = "calendar"
	DistinguishedFolderIdNameTypeCcontacts                              = "contacts"
	DistinguishedFolderIdNameTypeDdeleteditems                          = "deleteditems"
	DistinguishedFolderIdNameTypeDdrafts                                = "drafts"
	DistinguishedFolderIdNameTypeIinbox                                 = "inbox"
	DistinguishedFolderIdNameTypeJjournal                               = "journal"
	DistinguishedFolderIdNameTypeNnotes                                 = "notes"
	DistinguishedFolderIdNameTypeOoutbox                                = "outbox"
	DistinguishedFolderIdNameTypeSsentitems                             = "sentitems"
	DistinguishedFolderIdNameTypeTtasks                                 = "tasks"
	DistinguishedFolderIdNameTypeMmsgfolderroot                         = "msgfolderroot"
	DistinguishedFolderIdNameTypePpublicfoldersroot                     = "publicfoldersroot"
	DistinguishedFolderIdNameTypeRroot                                  = "root"
	DistinguishedFolderIdNameTypeJjunkemail                             = "junkemail"
	DistinguishedFolderIdNameTypeSsearchfolders                         = "searchfolders"
	DistinguishedFolderIdNameTypeVvoicemail                             = "voicemail"
	DistinguishedFolderIdNameTypeRrecoverableitemsroot                  = "recoverableitemsroot"
	DistinguishedFolderIdNameTypeRrecoverableitemsdeletions             = "recoverableitemsdeletions"
	DistinguishedFolderIdNameTypeRrecoverableitemsversions              = "recoverableitemsversions"
	DistinguishedFolderIdNameTypeRrecoverableitemspurges                = "recoverableitemspurges"
	DistinguishedFolderIdNameTypeRrecoverableitemsdiscoveryholds        = "recoverableitemsdiscoveryholds"
	DistinguishedFolderIdNameTypeAarchiveroot                           = "archiveroot"
	DistinguishedFolderIdNameTypeAarchivemsgfolderroot                  = "archivemsgfolderroot"
	DistinguishedFolderIdNameTypeAarchivedeleteditems                   = "archivedeleteditems"
	DistinguishedFolderIdNameTypeAarchiveinbox                          = "archiveinbox"
	DistinguishedFolderIdNameTypeAarchiverecoverableitemsroot           = "archiverecoverableitemsroot"
	DistinguishedFolderIdNameTypeAarchiverecoverableitemsdeletions      = "archiverecoverableitemsdeletions"
	DistinguishedFolderIdNameTypeAarchiverecoverableitemsversions       = "archiverecoverableitemsversions"
	DistinguishedFolderIdNameTypeAarchiverecoverableitemspurges         = "archiverecoverableitemspurges"
	DistinguishedFolderIdNameTypeAarchiverecoverableitemsdiscoveryholds = "archiverecoverableitemsdiscoveryholds"
	DistinguishedFolderIdNameTypeSsyncissues                            = "syncissues"
	DistinguishedFolderIdNameTypeCconflicts                             = "conflicts"
	DistinguishedFolderIdNameTypeLlocalfailures                         = "localfailures"
	DistinguishedFolderIdNameTypeSserverfailures                        = "serverfailures"
	DistinguishedFolderIdNameTypeRrecipientcache                        = "recipientcache"
	DistinguishedFolderIdNameTypeQquickcontacts                         = "quickcontacts"
	DistinguishedFolderIdNameTypeCconversationhistory                   = "conversationhistory"
	DistinguishedFolderIdNameTypeAadminauditlogs                        = "adminauditlogs"
	DistinguishedFolderIdNameTypeTtodosearch                            = "todosearch"
	DistinguishedFolderIdNameTypeMmycontacts                            = "mycontacts"
	DistinguishedFolderIdNameTypeDdirectory                             = "directory"
	DistinguishedFolderIdNameTypeIimcontactlist                         = "imcontactlist"
	DistinguishedFolderIdNameTypePpeopleconnect                         = "peopleconnect"
	DistinguishedFolderIdNameTypeFfavorites                             = "favorites"
	DistinguishedFolderIdNameTypeMmecontact                             = "mecontact"
	DistinguishedFolderIdNameTypePpersonmetadata                        = "personmetadata"
	DistinguishedFolderIdNameTypeTteamspaceactivity                     = "teamspaceactivity"
	DistinguishedFolderIdNameTypeTteamspacemessaging                    = "teamspacemessaging"
	DistinguishedFolderIdNameTypeTteamspaceworkitems                    = "teamspaceworkitems"
	DistinguishedFolderIdNameTypeSscheduled                             = "scheduled"
	DistinguishedFolderIdNameTypeOorionnotes                            = "orionnotes"
	DistinguishedFolderIdNameTypeTtagitems                              = "tagitems"
	DistinguishedFolderIdNameTypeAalltaggeditems                        = "alltaggeditems"
	DistinguishedFolderIdNameTypeAallcategorizeditems                   = "allcategorizeditems"
	DistinguishedFolderIdNameTypeEexternalcontacts                      = "externalcontacts"
	DistinguishedFolderIdNameTypeTteamchat                              = "teamchat"
	DistinguishedFolderIdNameTypeTteamchathistory                       = "teamchathistory"
	DistinguishedFolderIdNameTypeYyammerdata                            = "yammerdata"
	DistinguishedFolderIdNameTypeYyammerroot                            = "yammerroot"
	DistinguishedFolderIdNameTypeYyammerinbound                         = "yammerinbound"
	DistinguishedFolderIdNameTypeYyammeroutbound                        = "yammeroutbound"
	DistinguishedFolderIdNameTypeYyammerfeeds                           = "yammerfeeds"
	DistinguishedFolderIdNameTypeKkaizaladata                           = "kaizaladata"
	DistinguishedFolderIdNameTypeMmessageingestion                      = "messageingestion"
	DistinguishedFolderIdNameTypeOonedriveroot                          = "onedriveroot"
	DistinguishedFolderIdNameTypeOonedriverecylebin                     = "onedriverecylebin"
	DistinguishedFolderIdNameTypeOonedrivesystem                        = "onedrivesystem"
	DistinguishedFolderIdNameTypeOonedrivevolume                        = "onedrivevolume"
	DistinguishedFolderIdNameTypeIimportant                             = "important"
	DistinguishedFolderIdNameTypeSstarred                               = "starred"
	DistinguishedFolderIdNameTypeAarchive                               = "archive"
)

type ResolveNamesSearchScopeType XsString

const (
	ResolveNamesSearchScopeTypeActiveDirectory         = "ActiveDirectory"
	ResolveNamesSearchScopeTypeActiveDirectoryContacts = "ActiveDirectoryContacts"
	ResolveNamesSearchScopeTypeContacts                = "Contacts"
	ResolveNamesSearchScopeTypeContactsActiveDirectory = "ContactsActiveDirectory"
)

type DefaultShapeNamesType XsString

const (
	DefaultShapeNamesTypeIdOnly          = "IdOnly"
	DefaultShapeNamesTypeDefault         = "Default"
	DefaultShapeNamesTypeAllProperties   = "AllProperties"
	DefaultShapeNamesTypePcxPeopleSearch = "PcxPeopleSearch"
)

type ExchangeVersionType XsString

const (
	ExchangeVersionTypeExchange2007     = "Exchange2007"
	ExchangeVersionTypeExchange2007_SP1 = "Exchange2007_SP1"
	ExchangeVersionTypeExchange2009     = "Exchange2009"
	ExchangeVersionTypeExchange2010     = "Exchange2010"
	ExchangeVersionTypeExchange2010_SP1 = "Exchange2010_SP1"
	ExchangeVersionTypeExchange2010_SP2 = "Exchange2010_SP2"
	ExchangeVersionTypeExchange2012     = "Exchange2012"
	ExchangeVersionTypeExchange2013     = "Exchange2013"
	ExchangeVersionTypeExchange2013_SP1 = "Exchange2013_SP1"
	ExchangeVersionTypeExchange2015     = "Exchange2015"
	ExchangeVersionTypeExchange2016     = "Exchange2016"
	ExchangeVersionTypeV2015_10_05      = "V2015_10_05"
	ExchangeVersionTypeV2016_01_06      = "V2016_01_06"
	ExchangeVersionTypeV2016_04_13      = "V2016_04_13"
	ExchangeVersionTypeV2016_07_13      = "V2016_07_13"
	ExchangeVersionTypeV2016_10_10      = "V2016_10_10"
	ExchangeVersionTypeV2017_01_07      = "V2017_01_07"
	ExchangeVersionTypeV2017_04_14      = "V2017_04_14"
	ExchangeVersionTypeV2017_07_11      = "V2017_07_11"
	ExchangeVersionTypeV2017_10_09      = "V2017_10_09"
	ExchangeVersionTypeV2018_01_08      = "V2018_01_08"
)

type ResponseClassType XsString

const (
	ResponseClassTypeSuccess = "Success"
	ResponseClassTypeWarning = "Warning"
	ResponseClassTypeError   = "Error"
)

type ItemClassType XsString

type SensitivityChoicesType XsString

const (
	SensitivityChoicesTypeNormal       = "Normal"
	SensitivityChoicesTypePersonal     = "Personal"
	SensitivityChoicesTypePrivate      = "Private"
	SensitivityChoicesTypeConfidential = "Confidential"
)

type BodyTypeType XsString

const (
	BodyTypeTypeHTML = "HTML"
	BodyTypeTypeText = "Text"
)

type SendPromptType XsString

const (
	SendPromptTypeNone         = "None"
	SendPromptTypeSend         = "Send"
	SendPromptTypeVotingOption = "VotingOption"
)

type GuidType XsString

type SharingActionImportance XsString

const (
	SharingActionImportancePrimary   = "Primary"
	SharingActionImportanceSecondary = "Secondary"
)

type SharingActionType XsString

const (
	SharingActionTypeAccept = "Accept"
)

type SharingAction XsString

const (
	SharingActionAcceptAndViewCalendar = "AcceptAndViewCalendar"
	SharingActionViewCalendar          = "ViewCalendar"
	SharingActionAddThisCalendar       = "AddThisCalendar"
	SharingActionAccept                = "Accept"
)

type LegacyFreeBusyType XsString

const (
	LegacyFreeBusyTypeFree             = "Free"
	LegacyFreeBusyTypeTentative        = "Tentative"
	LegacyFreeBusyTypeBusy             = "Busy"
	LegacyFreeBusyTypeOOF              = "OOF"
	LegacyFreeBusyTypeWorkingElsewhere = "WorkingElsewhere"
	LegacyFreeBusyTypeNoData           = "NoData"
)

type CalendarItemTypeType XsString

const (
	CalendarItemTypeTypeSingle          = "Single"
	CalendarItemTypeTypeOccurrence      = "Occurrence"
	CalendarItemTypeTypeException       = "Exception"
	CalendarItemTypeTypeRecurringMaster = "RecurringMaster"
)

type ResponseTypeType XsString

const (
	ResponseTypeTypeUnknown            = "Unknown"
	ResponseTypeTypeOrganizer          = "Organizer"
	ResponseTypeTypeTentative          = "Tentative"
	ResponseTypeTypeAccept             = "Accept"
	ResponseTypeTypeDecline            = "Decline"
	ResponseTypeTypeNoResponseReceived = "NoResponseReceived"
)

type EmailReminderChangeType XsString

const (
	EmailReminderChangeTypeNone     = "None"
	EmailReminderChangeTypeAdded    = "Added"
	EmailReminderChangeTypeOverride = "Override"
	EmailReminderChangeTypeDeleted  = "Deleted"
)

type EmailReminderSendOption XsString

const (
	EmailReminderSendOptionNotSet       = "NotSet"
	EmailReminderSendOptionUser         = "User"
	EmailReminderSendOptionAllAttendees = "AllAttendees"
	EmailReminderSendOptionStaff        = "Staff"
	EmailReminderSendOptionCustomer     = "Customer"
)

type FileAsMappingType XsString

const (
	FileAsMappingTypeNone                  = "None"
	FileAsMappingTypeLastCommaFirst        = "LastCommaFirst"
	FileAsMappingTypeFirstSpaceLast        = "FirstSpaceLast"
	FileAsMappingTypeCompany               = "Company"
	FileAsMappingTypeLastCommaFirstCompany = "LastCommaFirstCompany"
	FileAsMappingTypeCompanyLastFirst      = "CompanyLastFirst"
	FileAsMappingTypeLastFirst             = "LastFirst"
	FileAsMappingTypeLastFirstCompany      = "LastFirstCompany"
	FileAsMappingTypeCompanyLastCommaFirst = "CompanyLastCommaFirst"
	FileAsMappingTypeLastFirstSuffix       = "LastFirstSuffix"
	FileAsMappingTypeLastSpaceFirstCompany = "LastSpaceFirstCompany"
	FileAsMappingTypeCompanyLastSpaceFirst = "CompanyLastSpaceFirst"
	FileAsMappingTypeLastSpaceFirst        = "LastSpaceFirst"
	FileAsMappingTypeDisplayName           = "DisplayName"
	FileAsMappingTypeFirstName             = "FirstName"
	FileAsMappingTypeLastFirstMiddleSuffix = "LastFirstMiddleSuffix"
	FileAsMappingTypeLastName              = "LastName"
	FileAsMappingTypeEmpty                 = "Empty"
)

type EmailAddressKeyType XsString

const (
	EmailAddressKeyTypeEmailAddress1 = "EmailAddress1"
	EmailAddressKeyTypeEmailAddress2 = "EmailAddress2"
	EmailAddressKeyTypeEmailAddress3 = "EmailAddress3"
)

type AbchEmailAddressTypeType XsString

const (
	AbchEmailAddressTypeTypePersonal = "Personal"
	AbchEmailAddressTypeTypeBusiness = "Business"
	AbchEmailAddressTypeTypeOther    = "Other"
	AbchEmailAddressTypeTypePassport = "Passport"
)

type PhysicalAddressKeyType XsString

const (
	PhysicalAddressKeyTypeHome     = "Home"
	PhysicalAddressKeyTypeBusiness = "Business"
	PhysicalAddressKeyTypeOther    = "Other"
)

type PhoneNumberKeyType XsString

const (
	PhoneNumberKeyTypeAssistantPhone   = "AssistantPhone"
	PhoneNumberKeyTypeBusinessFax      = "BusinessFax"
	PhoneNumberKeyTypeBusinessPhone    = "BusinessPhone"
	PhoneNumberKeyTypeBusinessPhone2   = "BusinessPhone2"
	PhoneNumberKeyTypeCallback         = "Callback"
	PhoneNumberKeyTypeCarPhone         = "CarPhone"
	PhoneNumberKeyTypeCompanyMainPhone = "CompanyMainPhone"
	PhoneNumberKeyTypeHomeFax          = "HomeFax"
	PhoneNumberKeyTypeHomePhone        = "HomePhone"
	PhoneNumberKeyTypeHomePhone2       = "HomePhone2"
	PhoneNumberKeyTypeIsdn             = "Isdn"
	PhoneNumberKeyTypeMobilePhone      = "MobilePhone"
	PhoneNumberKeyTypeOtherFax         = "OtherFax"
	PhoneNumberKeyTypeOtherTelephone   = "OtherTelephone"
	PhoneNumberKeyTypePager            = "Pager"
	PhoneNumberKeyTypePrimaryPhone     = "PrimaryPhone"
	PhoneNumberKeyTypeRadioPhone       = "RadioPhone"
	PhoneNumberKeyTypeTelex            = "Telex"
	PhoneNumberKeyTypeTtyTddPhone      = "TtyTddPhone"
	PhoneNumberKeyTypeBusinessMobile   = "BusinessMobile"
	PhoneNumberKeyTypeIPPhone          = "IPPhone"
	PhoneNumberKeyTypeMms              = "Mms"
	PhoneNumberKeyTypeMsn              = "Msn"
)

type ContactSourceType XsString

const (
	ContactSourceTypeActiveDirectory = "ActiveDirectory"
	ContactSourceTypeStore           = "Store"
)

type ImAddressKeyType XsString

const (
	ImAddressKeyTypeImAddress1 = "ImAddress1"
	ImAddressKeyTypeImAddress2 = "ImAddress2"
	ImAddressKeyTypeImAddress3 = "ImAddress3"
)

type PhysicalAddressIndexType XsString

const (
	PhysicalAddressIndexTypeNone     = "None"
	PhysicalAddressIndexTypeHome     = "Home"
	PhysicalAddressIndexTypeBusiness = "Business"
	PhysicalAddressIndexTypeOther    = "Other"
)

type ContactUrlKeyType XsString

const (
	ContactUrlKeyTypePersonal             = "Personal"
	ContactUrlKeyTypeBusiness             = "Business"
	ContactUrlKeyTypeAttachment           = "Attachment"
	ContactUrlKeyTypeEbcDisplayDefinition = "EbcDisplayDefinition"
	ContactUrlKeyTypeEbcFinalImage        = "EbcFinalImage"
	ContactUrlKeyTypeEbcLogo              = "EbcLogo"
	ContactUrlKeyTypeFeed                 = "Feed"
	ContactUrlKeyTypeImage                = "Image"
	ContactUrlKeyTypeInternalMarker       = "InternalMarker"
	ContactUrlKeyTypeOther                = "Other"
)

type MemberStatusType XsString

const (
	MemberStatusTypeUnrecognized = "Unrecognized"
	MemberStatusTypeNormal       = "Normal"
	MemberStatusTypeDemoted      = "Demoted"
)

type MeetingRequestTypeType XsString

const (
	MeetingRequestTypeTypeNone                = "None"
	MeetingRequestTypeTypeFullUpdate          = "FullUpdate"
	MeetingRequestTypeTypeInformationalUpdate = "InformationalUpdate"
	MeetingRequestTypeTypeNewMeetingRequest   = "NewMeetingRequest"
	MeetingRequestTypeTypeOutdated            = "Outdated"
	MeetingRequestTypeTypeSilentUpdate        = "SilentUpdate"
	MeetingRequestTypeTypePrincipalWantsCopy  = "PrincipalWantsCopy"
)

type DayOfWeekType XsString

const (
	DayOfWeekTypeSunday     = "Sunday"
	DayOfWeekTypeMonday     = "Monday"
	DayOfWeekTypeTuesday    = "Tuesday"
	DayOfWeekTypeWednesday  = "Wednesday"
	DayOfWeekTypeThursday   = "Thursday"
	DayOfWeekTypeFriday     = "Friday"
	DayOfWeekTypeSaturday   = "Saturday"
	DayOfWeekTypeDay        = "Day"
	DayOfWeekTypeWeekday    = "Weekday"
	DayOfWeekTypeWeekendDay = "WeekendDay"
)

type DayOfWeekIndexType XsString

const (
	DayOfWeekIndexTypeFirst  = "First"
	DayOfWeekIndexTypeSecond = "Second"
	DayOfWeekIndexTypeThird  = "Third"
	DayOfWeekIndexTypeFourth = "Fourth"
	DayOfWeekIndexTypeLast   = "Last"
)

type MonthNamesType XsString

const (
	MonthNamesTypeJanuary   = "January"
	MonthNamesTypeFebruary  = "February"
	MonthNamesTypeMarch     = "March"
	MonthNamesTypeApril     = "April"
	MonthNamesTypeMay       = "May"
	MonthNamesTypeJune      = "June"
	MonthNamesTypeJuly      = "July"
	MonthNamesTypeAugust    = "August"
	MonthNamesTypeSeptember = "September"
	MonthNamesTypeOctober   = "October"
	MonthNamesTypeNovember  = "November"
	MonthNamesTypeDecember  = "December"
)

type DaysOfWeekType []DayOfWeekType

type TransitionTargetKindType XsString

const (
	TransitionTargetKindTypePeriod = "Period"
	TransitionTargetKindTypeGroup  = "Group"
)

type LocationSourceType XsString

const (
	LocationSourceTypeNone              = "None"
	LocationSourceTypeLocationServices  = "LocationServices"
	LocationSourceTypePhonebookServices = "PhonebookServices"
	LocationSourceTypeDevice            = "Device"
	LocationSourceTypeContact           = "Contact"
	LocationSourceTypeResource          = "Resource"
)

type TaskDelegateStateType XsString

const (
	TaskDelegateStateTypeNoMatch  = "NoMatch"
	TaskDelegateStateTypeOwnNew   = "OwnNew"
	TaskDelegateStateTypeOwned    = "Owned"
	TaskDelegateStateTypeAccepted = "Accepted"
	TaskDelegateStateTypeDeclined = "Declined"
	TaskDelegateStateTypeMax      = "Max"
)

type TaskStatusType XsString

const (
	TaskStatusTypeNotStarted      = "NotStarted"
	TaskStatusTypeInProgress      = "InProgress"
	TaskStatusTypeCompleted       = "Completed"
	TaskStatusTypeWaitingOnOthers = "WaitingOnOthers"
	TaskStatusTypeDeferred        = "Deferred"
)

type RoleMemberTypeType XsString

const (
	RoleMemberTypeTypeNone       = "None"
	RoleMemberTypeTypePassport   = "Passport"
	RoleMemberTypeTypeEveryone   = "Everyone"
	RoleMemberTypeTypeEmail      = "Email"
	RoleMemberTypeTypePhone      = "Phone"
	RoleMemberTypeTypeSkypeId    = "SkypeId"
	RoleMemberTypeTypeExternalId = "ExternalId"
	RoleMemberTypeTypeGroup      = "Group"
	RoleMemberTypeTypeGuid       = "Guid"
	RoleMemberTypeTypeRole       = "Role"
	RoleMemberTypeTypeService    = "Service"
	RoleMemberTypeTypeCircle     = "Circle"
	RoleMemberTypeTypeDomain     = "Domain"
	RoleMemberTypeTypePartner    = "Partner"
)

type LobbyBypassType XsString

const (
	LobbyBypassTypeDisabled                      = "Disabled"
	LobbyBypassTypeEnabledForGatewayParticipants = "EnabledForGatewayParticipants"
)

type OnlineMeetingAccessLevelType XsString

const (
	OnlineMeetingAccessLevelTypeLocked   = "Locked"
	OnlineMeetingAccessLevelTypeInvited  = "Invited"
	OnlineMeetingAccessLevelTypeInternal = "Internal"
	OnlineMeetingAccessLevelTypeEveryone = "Everyone"
)

type PresentersType XsString

const (
	PresentersTypeDisabled = "Disabled"
	PresentersTypeInternal = "Internal"
	PresentersTypeEveryone = "Everyone"
)

type ImportanceChoicesType XsString

const (
	ImportanceChoicesTypeLow    = "Low"
	ImportanceChoicesTypeNormal = "Normal"
	ImportanceChoicesTypeHigh   = "High"
)

type ReminderMinutesBeforeStartType any // union(ReminderMinutesBeforeStartTypeUnion0|ReminderMinutesBeforeStartTypeUnion1)
type DistinguishedPropertySetType XsString

const (
	DistinguishedPropertySetTypeMeeting           = "Meeting"
	DistinguishedPropertySetTypeAppointment       = "Appointment"
	DistinguishedPropertySetTypeCommon            = "Common"
	DistinguishedPropertySetTypePublicStrings     = "PublicStrings"
	DistinguishedPropertySetTypeAddress           = "Address"
	DistinguishedPropertySetTypeInternetHeaders   = "InternetHeaders"
	DistinguishedPropertySetTypeCalendarAssistant = "CalendarAssistant"
	DistinguishedPropertySetTypeUnifiedMessaging  = "UnifiedMessaging"
	DistinguishedPropertySetTypeTask              = "Task"
	DistinguishedPropertySetTypeSharing           = "Sharing"
)

type PropertyTagType any // union(XsUnsignedShort|PropertyTagTypeUnion0)
type MapiPropertyTypeType XsString

const (
	MapiPropertyTypeTypeApplicationTime      = "ApplicationTime"
	MapiPropertyTypeTypeApplicationTimeArray = "ApplicationTimeArray"
	MapiPropertyTypeTypeBinary               = "Binary"
	MapiPropertyTypeTypeBinaryArray          = "BinaryArray"
	MapiPropertyTypeTypeBoolean              = "Boolean"
	MapiPropertyTypeTypeCLSID                = "CLSID"
	MapiPropertyTypeTypeCLSIDArray           = "CLSIDArray"
	MapiPropertyTypeTypeCurrency             = "Currency"
	MapiPropertyTypeTypeCurrencyArray        = "CurrencyArray"
	MapiPropertyTypeTypeDouble               = "Double"
	MapiPropertyTypeTypeDoubleArray          = "DoubleArray"
	MapiPropertyTypeTypeError                = "Error"
	MapiPropertyTypeTypeFloat                = "Float"
	MapiPropertyTypeTypeFloatArray           = "FloatArray"
	MapiPropertyTypeTypeInteger              = "Integer"
	MapiPropertyTypeTypeIntegerArray         = "IntegerArray"
	MapiPropertyTypeTypeLong                 = "Long"
	MapiPropertyTypeTypeLongArray            = "LongArray"
	MapiPropertyTypeTypeNull                 = "Null"
	MapiPropertyTypeTypeObject               = "Object"
	MapiPropertyTypeTypeObjectArray          = "ObjectArray"
	MapiPropertyTypeTypeShort                = "Short"
	MapiPropertyTypeTypeShortArray           = "ShortArray"
	MapiPropertyTypeTypeSystemTime           = "SystemTime"
	MapiPropertyTypeTypeSystemTimeArray      = "SystemTimeArray"
	MapiPropertyTypeTypeString               = "String"
	MapiPropertyTypeTypeStringArray          = "StringArray"
)

type FlagStatusType XsString

const (
	FlagStatusTypeNotFlagged = "NotFlagged"
	FlagStatusTypeFlagged    = "Flagged"
	FlagStatusTypeComplete   = "Complete"
)

type EmailPositionType XsString

const (
	EmailPositionTypeLatestReply = "LatestReply"
	EmailPositionTypeOther       = "Other"
	EmailPositionTypeSubject     = "Subject"
	EmailPositionTypeSignature   = "Signature"
)

type PredictedActionReasonType XsString

const (
	PredictedActionReasonTypeNone                          = "None"
	PredictedActionReasonTypeConversationStarterIsYou      = "ConversationStarterIsYou"
	PredictedActionReasonTypeOnlyRecipient                 = "OnlyRecipient"
	PredictedActionReasonTypeConversationContributions     = "ConversationContributions"
	PredictedActionReasonTypeMarkedImportantBySender       = "MarkedImportantBySender"
	PredictedActionReasonTypeSenderIsManager               = "SenderIsManager"
	PredictedActionReasonTypeSenderIsInManagementChain     = "SenderIsInManagementChain"
	PredictedActionReasonTypeSenderIsDirectReport          = "SenderIsDirectReport"
	PredictedActionReasonTypeActionBasedOnSender           = "ActionBasedOnSender"
	PredictedActionReasonTypeNameOnToLine                  = "NameOnToLine"
	PredictedActionReasonTypeNameOnCcLine                  = "NameOnCcLine"
	PredictedActionReasonTypeManagerPosition               = "ManagerPosition"
	PredictedActionReasonTypeReplyToAMessageFromMe         = "ReplyToAMessageFromMe"
	PredictedActionReasonTypePreviouslyFlagged             = "PreviouslyFlagged"
	PredictedActionReasonTypeActionBasedOnRecipients       = "ActionBasedOnRecipients"
	PredictedActionReasonTypeActionBasedOnSubjectWords     = "ActionBasedOnSubjectWords"
	PredictedActionReasonTypeActionBasedOnBasedOnBodyWords = "ActionBasedOnBasedOnBodyWords"
)

type IconIndexType XsString

const (
	IconIndexTypeDefault                            = "Default"
	IconIndexTypePostItem                           = "PostItem"
	IconIndexTypeMailRead                           = "MailRead"
	IconIndexTypeMailUnread                         = "MailUnread"
	IconIndexTypeMailReplied                        = "MailReplied"
	IconIndexTypeMailForwarded                      = "MailForwarded"
	IconIndexTypeMailEncrypted                      = "MailEncrypted"
	IconIndexTypeMailSmimeSigned                    = "MailSmimeSigned"
	IconIndexTypeMailEncryptedReplied               = "MailEncryptedReplied"
	IconIndexTypeMailSmimeSignedReplied             = "MailSmimeSignedReplied"
	IconIndexTypeMailEncryptedForwarded             = "MailEncryptedForwarded"
	IconIndexTypeMailSmimeSignedForwarded           = "MailSmimeSignedForwarded"
	IconIndexTypeMailEncryptedRead                  = "MailEncryptedRead"
	IconIndexTypeMailSmimeSignedRead                = "MailSmimeSignedRead"
	IconIndexTypeMailIrm                            = "MailIrm"
	IconIndexTypeMailIrmForwarded                   = "MailIrmForwarded"
	IconIndexTypeMailIrmReplied                     = "MailIrmReplied"
	IconIndexTypeSmsSubmitted                       = "SmsSubmitted"
	IconIndexTypeSmsRoutedToDeliveryPoint           = "SmsRoutedToDeliveryPoint"
	IconIndexTypeSmsRoutedToExternalMessagingSystem = "SmsRoutedToExternalMessagingSystem"
	IconIndexTypeSmsDelivered                       = "SmsDelivered"
	IconIndexTypeOutlookDefaultForContacts          = "OutlookDefaultForContacts"
	IconIndexTypeAppointmentItem                    = "AppointmentItem"
	IconIndexTypeAppointmentRecur                   = "AppointmentRecur"
	IconIndexTypeAppointmentMeet                    = "AppointmentMeet"
	IconIndexTypeAppointmentMeetRecur               = "AppointmentMeetRecur"
	IconIndexTypeAppointmentMeetNY                  = "AppointmentMeetNY"
	IconIndexTypeAppointmentMeetYes                 = "AppointmentMeetYes"
	IconIndexTypeAppointmentMeetNo                  = "AppointmentMeetNo"
	IconIndexTypeAppointmentMeetMaybe               = "AppointmentMeetMaybe"
	IconIndexTypeAppointmentMeetCancel              = "AppointmentMeetCancel"
	IconIndexTypeAppointmentMeetInfo                = "AppointmentMeetInfo"
	IconIndexTypeTaskItem                           = "TaskItem"
	IconIndexTypeTaskRecur                          = "TaskRecur"
	IconIndexTypeTaskOwned                          = "TaskOwned"
	IconIndexTypeTaskDelegated                      = "TaskDelegated"
)

type InferenceClassificationType XsString

const (
	InferenceClassificationTypeFocused = "Focused"
	InferenceClassificationTypeOther   = "Other"
)

type DistinguishedUserType XsString

const (
	DistinguishedUserTypeDefault   = "Default"
	DistinguishedUserTypeAnonymous = "Anonymous"
)

type PermissionActionType XsString

const (
	PermissionActionTypeNone  = "None"
	PermissionActionTypeOwned = "Owned"
	PermissionActionTypeAll   = "All"
)

type PermissionReadAccessType XsString

const (
	PermissionReadAccessTypeNone        = "None"
	PermissionReadAccessTypeFullDetails = "FullDetails"
)

type PermissionLevelType XsString

const (
	PermissionLevelTypeNone             = "None"
	PermissionLevelTypeOwner            = "Owner"
	PermissionLevelTypePublishingEditor = "PublishingEditor"
	PermissionLevelTypeEditor           = "Editor"
	PermissionLevelTypePublishingAuthor = "PublishingAuthor"
	PermissionLevelTypeAuthor           = "Author"
	PermissionLevelTypeNoneditingAuthor = "NoneditingAuthor"
	PermissionLevelTypeReviewer         = "Reviewer"
	PermissionLevelTypeContributor      = "Contributor"
	PermissionLevelTypeCustom           = "Custom"
)

type CalendarPermissionReadAccessType XsString

const (
	CalendarPermissionReadAccessTypeNone                      = "None"
	CalendarPermissionReadAccessTypeTimeOnly                  = "TimeOnly"
	CalendarPermissionReadAccessTypeTimeAndSubjectAndLocation = "TimeAndSubjectAndLocation"
	CalendarPermissionReadAccessTypeFullDetails               = "FullDetails"
)

type CalendarPermissionLevelType XsString

const (
	CalendarPermissionLevelTypeNone                              = "None"
	CalendarPermissionLevelTypeOwner                             = "Owner"
	CalendarPermissionLevelTypePublishingEditor                  = "PublishingEditor"
	CalendarPermissionLevelTypeEditor                            = "Editor"
	CalendarPermissionLevelTypePublishingAuthor                  = "PublishingAuthor"
	CalendarPermissionLevelTypeAuthor                            = "Author"
	CalendarPermissionLevelTypeNoneditingAuthor                  = "NoneditingAuthor"
	CalendarPermissionLevelTypeReviewer                          = "Reviewer"
	CalendarPermissionLevelTypeContributor                       = "Contributor"
	CalendarPermissionLevelTypeFreeBusyTimeOnly                  = "FreeBusyTimeOnly"
	CalendarPermissionLevelTypeFreeBusyTimeAndSubjectAndLocation = "FreeBusyTimeAndSubjectAndLocation"
	CalendarPermissionLevelTypeCustom                            = "Custom"
)

type SearchFolderTraversalType XsString

const (
	SearchFolderTraversalTypeShallow = "Shallow"
	SearchFolderTraversalTypeDeep    = "Deep"
)

type ClientAccessTokenTypeType XsString

const (
	ClientAccessTokenTypeTypeCallerIdentity           = "CallerIdentity"
	ClientAccessTokenTypeTypeExtensionCallback        = "ExtensionCallback"
	ClientAccessTokenTypeTypeScopedToken              = "ScopedToken"
	ClientAccessTokenTypeTypeExtensionRestApiCallback = "ExtensionRestApiCallback"
	ClientAccessTokenTypeTypeConnectors               = "Connectors"
	ClientAccessTokenTypeTypeLoki                     = "Loki"
	ClientAccessTokenTypeTypeDesktopOutlook           = "DesktopOutlook"
)

type SubscriptionIdType NonEmptyStringType

type WatermarkType NonEmptyStringType

type ConnectionStatusType XsString

const (
	ConnectionStatusTypeOK     = "OK"
	ConnectionStatusTypeClosed = "Closed"
)

type IdFormatType XsString

const (
	IdFormatTypeEwsLegacyId = "EwsLegacyId"
	IdFormatTypeEwsId       = "EwsId"
	IdFormatTypeEntryId     = "EntryId"
	IdFormatTypeHexEntryId  = "HexEntryId"
	IdFormatTypeStoreId     = "StoreId"
	IdFormatTypeOwaId       = "OwaId"
)

type InvalidRecipientResponseCodeType XsString

const (
	InvalidRecipientResponseCodeTypeOtherError                                           = "OtherError"
	InvalidRecipientResponseCodeTypeRecipientOrganizationNotFederated                    = "RecipientOrganizationNotFederated"
	InvalidRecipientResponseCodeTypeCannotObtainTokenFromSTS                             = "CannotObtainTokenFromSTS"
	InvalidRecipientResponseCodeTypeSystemPolicyBlocksSharingWithThisRecipient           = "SystemPolicyBlocksSharingWithThisRecipient"
	InvalidRecipientResponseCodeTypeRecipientOrganizationFederatedWithUnknownTokenIssuer = "RecipientOrganizationFederatedWithUnknownTokenIssuer"
)

type UserConfigurationDictionaryObjectTypesType XsString

const (
	UserConfigurationDictionaryObjectTypesTypeDateTime          = "DateTime"
	UserConfigurationDictionaryObjectTypesTypeBoolean           = "Boolean"
	UserConfigurationDictionaryObjectTypesTypeByte              = "Byte"
	UserConfigurationDictionaryObjectTypesTypeString            = "String"
	UserConfigurationDictionaryObjectTypesTypeInteger32         = "Integer32"
	UserConfigurationDictionaryObjectTypesTypeUnsignedInteger32 = "UnsignedInteger32"
	UserConfigurationDictionaryObjectTypesTypeInteger64         = "Integer64"
	UserConfigurationDictionaryObjectTypesTypeUnsignedInteger64 = "UnsignedInteger64"
	UserConfigurationDictionaryObjectTypesTypeStringArray       = "StringArray"
	UserConfigurationDictionaryObjectTypesTypeByteArray         = "ByteArray"
)

type ReminderGroupType XsString

const (
	ReminderGroupTypeCalendar = "Calendar"
	ReminderGroupTypeTask     = "Task"
)

type MailboxSearchLocationType XsString

const (
	MailboxSearchLocationTypePrimaryOnly = "PrimaryOnly"
	MailboxSearchLocationTypeArchiveOnly = "ArchiveOnly"
	MailboxSearchLocationTypeAll         = "All"
)

type SearchResultType XsString

const (
	SearchResultTypeStatisticsOnly = "StatisticsOnly"
	SearchResultTypePreviewOnly    = "PreviewOnly"
)

type HoldStatusType XsString

const (
	HoldStatusTypeNotOnHold   = "NotOnHold"
	HoldStatusTypePending     = "Pending"
	HoldStatusTypeOnHold      = "OnHold"
	HoldStatusTypePartialHold = "PartialHold"
	HoldStatusTypeFailed      = "Failed"
)

type ItemIndexErrorType XsString

const (
	ItemIndexErrorTypeNone                   = "None"
	ItemIndexErrorTypeGenericError           = "GenericError"
	ItemIndexErrorTypeTimeout                = "Timeout"
	ItemIndexErrorTypeStaleEvent             = "StaleEvent"
	ItemIndexErrorTypeMailboxOffline         = "MailboxOffline"
	ItemIndexErrorTypeAttachmentLimitReached = "AttachmentLimitReached"
	ItemIndexErrorTypeMarsWriterTruncation   = "MarsWriterTruncation"
	ItemIndexErrorTypeDocumentParserFailure  = "DocumentParserFailure"
)

type ElcFolderType XsString

const (
	ElcFolderTypeCalendar            = "Calendar"
	ElcFolderTypeContacts            = "Contacts"
	ElcFolderTypeDeletedItems        = "DeletedItems"
	ElcFolderTypeDrafts              = "Drafts"
	ElcFolderTypeInbox               = "Inbox"
	ElcFolderTypeJunkEmail           = "JunkEmail"
	ElcFolderTypeJournal             = "Journal"
	ElcFolderTypeNotes               = "Notes"
	ElcFolderTypeOutbox              = "Outbox"
	ElcFolderTypeSentItems           = "SentItems"
	ElcFolderTypeTasks               = "Tasks"
	ElcFolderTypeAll                 = "All"
	ElcFolderTypeManagedCustomFolder = "ManagedCustomFolder"
	ElcFolderTypeRssSubscriptions    = "RssSubscriptions"
	ElcFolderTypeSyncIssues          = "SyncIssues"
	ElcFolderTypeConversationHistory = "ConversationHistory"
	ElcFolderTypePersonal            = "Personal"
	ElcFolderTypeRecoverableItems    = "RecoverableItems"
	ElcFolderTypeNonIpmRoot          = "NonIpmRoot"
)

type RetentionActionType XsString

const (
	RetentionActionTypeNone                     = "None"
	RetentionActionTypeMoveToDeletedItems       = "MoveToDeletedItems"
	RetentionActionTypeMoveToFolder             = "MoveToFolder"
	RetentionActionTypeDeleteAndAllowRecovery   = "DeleteAndAllowRecovery"
	RetentionActionTypePermanentlyDelete        = "PermanentlyDelete"
	RetentionActionTypeMarkAsPastRetentionLimit = "MarkAsPastRetentionLimit"
	RetentionActionTypeMoveToArchive            = "MoveToArchive"
)

type IndexBasePointType XsString

const (
	IndexBasePointTypeBeginning = "Beginning"
	IndexBasePointTypeEnd       = "End"
)

type FolderQueryTraversalType XsString

const (
	FolderQueryTraversalTypeShallow     = "Shallow"
	FolderQueryTraversalTypeDeep        = "Deep"
	FolderQueryTraversalTypeSoftDeleted = "SoftDeleted"
)

type BodyTypeResponseType XsString

const (
	BodyTypeResponseTypeBest = "Best"
	BodyTypeResponseTypeHTML = "HTML"
	BodyTypeResponseTypeText = "Text"
)

type SortDirectionType XsString

const (
	SortDirectionTypeAscending  = "Ascending"
	SortDirectionTypeDescending = "Descending"
)

type UnindexedFieldURIType XsString

const (
	UnindexedFieldURITypeFfolderFolderId                          = "folder:FolderId"
	UnindexedFieldURITypeFfolderParentFolderId                    = "folder:ParentFolderId"
	UnindexedFieldURITypeFfolderDisplayName                       = "folder:DisplayName"
	UnindexedFieldURITypeFfolderUnreadCount                       = "folder:UnreadCount"
	UnindexedFieldURITypeFfolderTotalCount                        = "folder:TotalCount"
	UnindexedFieldURITypeFfolderChildFolderCount                  = "folder:ChildFolderCount"
	UnindexedFieldURITypeFfolderFolderClass                       = "folder:FolderClass"
	UnindexedFieldURITypeFfolderSearchParameters                  = "folder:SearchParameters"
	UnindexedFieldURITypeFfolderManagedFolderInformation          = "folder:ManagedFolderInformation"
	UnindexedFieldURITypeFfolderPermissionSet                     = "folder:PermissionSet"
	UnindexedFieldURITypeFfolderEffectiveRights                   = "folder:EffectiveRights"
	UnindexedFieldURITypeFfolderSharingEffectiveRights            = "folder:SharingEffectiveRights"
	UnindexedFieldURITypeFfolderDistinguishedFolderId             = "folder:DistinguishedFolderId"
	UnindexedFieldURITypeFfolderPolicyTag                         = "folder:PolicyTag"
	UnindexedFieldURITypeFfolderArchiveTag                        = "folder:ArchiveTag"
	UnindexedFieldURITypeFfolderReplicaList                       = "folder:ReplicaList"
	UnindexedFieldURITypeIitemItemId                              = "item:ItemId"
	UnindexedFieldURITypeIitemParentFolderId                      = "item:ParentFolderId"
	UnindexedFieldURITypeIitemItemClass                           = "item:ItemClass"
	UnindexedFieldURITypeIitemMimeContent                         = "item:MimeContent"
	UnindexedFieldURITypeIitemAttachments                         = "item:Attachments"
	UnindexedFieldURITypeIitemSubject                             = "item:Subject"
	UnindexedFieldURITypeIitemDateTimeReceived                    = "item:DateTimeReceived"
	UnindexedFieldURITypeIitemSize                                = "item:Size"
	UnindexedFieldURITypeIitemCategories                          = "item:Categories"
	UnindexedFieldURITypeIitemHasAttachments                      = "item:HasAttachments"
	UnindexedFieldURITypeIitemImportance                          = "item:Importance"
	UnindexedFieldURITypeIitemInReplyTo                           = "item:InReplyTo"
	UnindexedFieldURITypeIitemInternetMessageHeaders              = "item:InternetMessageHeaders"
	UnindexedFieldURITypeIitemIsAssociated                        = "item:IsAssociated"
	UnindexedFieldURITypeIitemIsDraft                             = "item:IsDraft"
	UnindexedFieldURITypeIitemIsFromMe                            = "item:IsFromMe"
	UnindexedFieldURITypeIitemIsResend                            = "item:IsResend"
	UnindexedFieldURITypeIitemIsSubmitted                         = "item:IsSubmitted"
	UnindexedFieldURITypeIitemIsUnmodified                        = "item:IsUnmodified"
	UnindexedFieldURITypeIitemDateTimeSent                        = "item:DateTimeSent"
	UnindexedFieldURITypeIitemDateTimeCreated                     = "item:DateTimeCreated"
	UnindexedFieldURITypeIitemBody                                = "item:Body"
	UnindexedFieldURITypeIitemResponseObjects                     = "item:ResponseObjects"
	UnindexedFieldURITypeIitemSensitivity                         = "item:Sensitivity"
	UnindexedFieldURITypeIitemReminderDueBy                       = "item:ReminderDueBy"
	UnindexedFieldURITypeIitemReminderIsSet                       = "item:ReminderIsSet"
	UnindexedFieldURITypeIitemReminderNextTime                    = "item:ReminderNextTime"
	UnindexedFieldURITypeIitemReminderMinutesBeforeStart          = "item:ReminderMinutesBeforeStart"
	UnindexedFieldURITypeIitemDisplayTo                           = "item:DisplayTo"
	UnindexedFieldURITypeIitemDisplayCc                           = "item:DisplayCc"
	UnindexedFieldURITypeIitemDisplayBcc                          = "item:DisplayBcc"
	UnindexedFieldURITypeIitemCulture                             = "item:Culture"
	UnindexedFieldURITypeIitemEffectiveRights                     = "item:EffectiveRights"
	UnindexedFieldURITypeIitemLastModifiedName                    = "item:LastModifiedName"
	UnindexedFieldURITypeIitemLastModifiedTime                    = "item:LastModifiedTime"
	UnindexedFieldURITypeIitemConversationId                      = "item:ConversationId"
	UnindexedFieldURITypeIitemUniqueBody                          = "item:UniqueBody"
	UnindexedFieldURITypeIitemFlag                                = "item:Flag"
	UnindexedFieldURITypeIitemStoreEntryId                        = "item:StoreEntryId"
	UnindexedFieldURITypeIitemInstanceKey                         = "item:InstanceKey"
	UnindexedFieldURITypeIitemNormalizedBody                      = "item:NormalizedBody"
	UnindexedFieldURITypeIitemEntityExtractionResult              = "item:EntityExtractionResult"
	UnindexedFieldURITypeIitemPolicyTag                           = "item:PolicyTag"
	UnindexedFieldURITypeIitemArchiveTag                          = "item:ArchiveTag"
	UnindexedFieldURITypeIitemRetentionDate                       = "item:RetentionDate"
	UnindexedFieldURITypeIitemPreview                             = "item:Preview"
	UnindexedFieldURITypeIitemPredictedActionReasons              = "item:PredictedActionReasons"
	UnindexedFieldURITypeIitemIsClutter                           = "item:IsClutter"
	UnindexedFieldURITypeIitemRightsManagementLicenseData         = "item:RightsManagementLicenseData"
	UnindexedFieldURITypeIitemBlockStatus                         = "item:BlockStatus"
	UnindexedFieldURITypeIitemHasBlockedImages                    = "item:HasBlockedImages"
	UnindexedFieldURITypeIitemWebClientReadFormQueryString        = "item:WebClientReadFormQueryString"
	UnindexedFieldURITypeIitemWebClientEditFormQueryString        = "item:WebClientEditFormQueryString"
	UnindexedFieldURITypeIitemTextBody                            = "item:TextBody"
	UnindexedFieldURITypeIitemIconIndex                           = "item:IconIndex"
	UnindexedFieldURITypeIitemMimeContentUTF8                     = "item:MimeContentUTF8"
	UnindexedFieldURITypeIitemMentions                            = "item:Mentions"
	UnindexedFieldURITypeIitemMentionedMe                         = "item:MentionedMe"
	UnindexedFieldURITypeIitemMentionsPreview                     = "item:MentionsPreview"
	UnindexedFieldURITypeIitemMentionsEx                          = "item:MentionsEx"
	UnindexedFieldURITypeIitemHashtags                            = "item:Hashtags"
	UnindexedFieldURITypeIitemAppliedHashtags                     = "item:AppliedHashtags"
	UnindexedFieldURITypeIitemAppliedHashtagsPreview              = "item:AppliedHashtagsPreview"
	UnindexedFieldURITypeIitemLikes                               = "item:Likes"
	UnindexedFieldURITypeIitemLikesPreview                        = "item:LikesPreview"
	UnindexedFieldURITypeIitemPendingSocialActivityTagIds         = "item:PendingSocialActivityTagIds"
	UnindexedFieldURITypeIitemAtAllMention                        = "item:AtAllMention"
	UnindexedFieldURITypeIitemCanDelete                           = "item:CanDelete"
	UnindexedFieldURITypeIitemInferenceClassification             = "item:InferenceClassification"
	UnindexedFieldURITypeIitemFirstBody                           = "item:FirstBody"
	UnindexedFieldURITypeMmessageConversationIndex                = "message:ConversationIndex"
	UnindexedFieldURITypeMmessageConversationTopic                = "message:ConversationTopic"
	UnindexedFieldURITypeMmessageInternetMessageId                = "message:InternetMessageId"
	UnindexedFieldURITypeMmessageIsRead                           = "message:IsRead"
	UnindexedFieldURITypeMmessageIsResponseRequested              = "message:IsResponseRequested"
	UnindexedFieldURITypeMmessageIsReadReceiptRequested           = "message:IsReadReceiptRequested"
	UnindexedFieldURITypeMmessageIsDeliveryReceiptRequested       = "message:IsDeliveryReceiptRequested"
	UnindexedFieldURITypeMmessageReceivedBy                       = "message:ReceivedBy"
	UnindexedFieldURITypeMmessageReceivedRepresenting             = "message:ReceivedRepresenting"
	UnindexedFieldURITypeMmessageReferences                       = "message:References"
	UnindexedFieldURITypeMmessageReplyTo                          = "message:ReplyTo"
	UnindexedFieldURITypeMmessageFrom                             = "message:From"
	UnindexedFieldURITypeMmessageSender                           = "message:Sender"
	UnindexedFieldURITypeMmessageToRecipients                     = "message:ToRecipients"
	UnindexedFieldURITypeMmessageCcRecipients                     = "message:CcRecipients"
	UnindexedFieldURITypeMmessageBccRecipients                    = "message:BccRecipients"
	UnindexedFieldURITypeMmessageApprovalRequestData              = "message:ApprovalRequestData"
	UnindexedFieldURITypeMmessageVotingInformation                = "message:VotingInformation"
	UnindexedFieldURITypeMmessageReminderMessageData              = "message:ReminderMessageData"
	UnindexedFieldURITypeMmessagePublishedCalendarItemIcs         = "message:PublishedCalendarItemIcs"
	UnindexedFieldURITypeMmessagePublishedCalendarItemName        = "message:PublishedCalendarItemName"
	UnindexedFieldURITypeMmessageMessageSafety                    = "message:MessageSafety"
	UnindexedFieldURITypeSsharingMessageSharingMessageAction      = "sharingMessage:SharingMessageAction"
	UnindexedFieldURITypeMmeetingAssociatedCalendarItemId         = "meeting:AssociatedCalendarItemId"
	UnindexedFieldURITypeMmeetingIsDelegated                      = "meeting:IsDelegated"
	UnindexedFieldURITypeMmeetingIsOutOfDate                      = "meeting:IsOutOfDate"
	UnindexedFieldURITypeMmeetingHasBeenProcessed                 = "meeting:HasBeenProcessed"
	UnindexedFieldURITypeMmeetingResponseType                     = "meeting:ResponseType"
	UnindexedFieldURITypeMmeetingProposedStart                    = "meeting:ProposedStart"
	UnindexedFieldURITypeMmeetingProposedEnd                      = "meeting:ProposedEnd"
	UnindexedFieldURITypeMmeetingDoNotForwardMeeting              = "meeting:DoNotForwardMeeting"
	UnindexedFieldURITypeMmeetingRequestMeetingRequestType        = "meetingRequest:MeetingRequestType"
	UnindexedFieldURITypeMmeetingRequestIntendedFreeBusyStatus    = "meetingRequest:IntendedFreeBusyStatus"
	UnindexedFieldURITypeMmeetingRequestChangeHighlights          = "meetingRequest:ChangeHighlights"
	UnindexedFieldURITypeCcalendarStart                           = "calendar:Start"
	UnindexedFieldURITypeCcalendarEnd                             = "calendar:End"
	UnindexedFieldURITypeCcalendarOriginalStart                   = "calendar:OriginalStart"
	UnindexedFieldURITypeCcalendarStartWallClock                  = "calendar:StartWallClock"
	UnindexedFieldURITypeCcalendarEndWallClock                    = "calendar:EndWallClock"
	UnindexedFieldURITypeCcalendarStartTimeZoneId                 = "calendar:StartTimeZoneId"
	UnindexedFieldURITypeCcalendarEndTimeZoneId                   = "calendar:EndTimeZoneId"
	UnindexedFieldURITypeCcalendarIsAllDayEvent                   = "calendar:IsAllDayEvent"
	UnindexedFieldURITypeCcalendarLegacyFreeBusyStatus            = "calendar:LegacyFreeBusyStatus"
	UnindexedFieldURITypeCcalendarLocation                        = "calendar:Location"
	UnindexedFieldURITypeCcalendarEnhancedLocation                = "calendar:EnhancedLocation"
	UnindexedFieldURITypeCcalendarWhen                            = "calendar:When"
	UnindexedFieldURITypeCcalendarIsMeeting                       = "calendar:IsMeeting"
	UnindexedFieldURITypeCcalendarIsCancelled                     = "calendar:IsCancelled"
	UnindexedFieldURITypeCcalendarIsRecurring                     = "calendar:IsRecurring"
	UnindexedFieldURITypeCcalendarMeetingRequestWasSent           = "calendar:MeetingRequestWasSent"
	UnindexedFieldURITypeCcalendarIsResponseRequested             = "calendar:IsResponseRequested"
	UnindexedFieldURITypeCcalendarCalendarItemType                = "calendar:CalendarItemType"
	UnindexedFieldURITypeCcalendarMyResponseType                  = "calendar:MyResponseType"
	UnindexedFieldURITypeCcalendarOrganizer                       = "calendar:Organizer"
	UnindexedFieldURITypeCcalendarRequiredAttendees               = "calendar:RequiredAttendees"
	UnindexedFieldURITypeCcalendarOptionalAttendees               = "calendar:OptionalAttendees"
	UnindexedFieldURITypeCcalendarResources                       = "calendar:Resources"
	UnindexedFieldURITypeCcalendarConflictingMeetingCount         = "calendar:ConflictingMeetingCount"
	UnindexedFieldURITypeCcalendarAdjacentMeetingCount            = "calendar:AdjacentMeetingCount"
	UnindexedFieldURITypeCcalendarConflictingMeetings             = "calendar:ConflictingMeetings"
	UnindexedFieldURITypeCcalendarAdjacentMeetings                = "calendar:AdjacentMeetings"
	UnindexedFieldURITypeCcalendarInboxReminders                  = "calendar:InboxReminders"
	UnindexedFieldURITypeCcalendarDuration                        = "calendar:Duration"
	UnindexedFieldURITypeCcalendarTimeZone                        = "calendar:TimeZone"
	UnindexedFieldURITypeCcalendarAppointmentReplyTime            = "calendar:AppointmentReplyTime"
	UnindexedFieldURITypeCcalendarAppointmentSequenceNumber       = "calendar:AppointmentSequenceNumber"
	UnindexedFieldURITypeCcalendarAppointmentState                = "calendar:AppointmentState"
	UnindexedFieldURITypeCcalendarRecurrence                      = "calendar:Recurrence"
	UnindexedFieldURITypeCcalendarFirstOccurrence                 = "calendar:FirstOccurrence"
	UnindexedFieldURITypeCcalendarLastOccurrence                  = "calendar:LastOccurrence"
	UnindexedFieldURITypeCcalendarModifiedOccurrences             = "calendar:ModifiedOccurrences"
	UnindexedFieldURITypeCcalendarDeletedOccurrences              = "calendar:DeletedOccurrences"
	UnindexedFieldURITypeCcalendarMeetingTimeZone                 = "calendar:MeetingTimeZone"
	UnindexedFieldURITypeCcalendarConferenceType                  = "calendar:ConferenceType"
	UnindexedFieldURITypeCcalendarAllowNewTimeProposal            = "calendar:AllowNewTimeProposal"
	UnindexedFieldURITypeCcalendarIsOnlineMeeting                 = "calendar:IsOnlineMeeting"
	UnindexedFieldURITypeCcalendarMeetingWorkspaceUrl             = "calendar:MeetingWorkspaceUrl"
	UnindexedFieldURITypeCcalendarNetShowUrl                      = "calendar:NetShowUrl"
	UnindexedFieldURITypeCcalendarUID                             = "calendar:UID"
	UnindexedFieldURITypeCcalendarRecurrenceId                    = "calendar:RecurrenceId"
	UnindexedFieldURITypeCcalendarDateTimeStamp                   = "calendar:DateTimeStamp"
	UnindexedFieldURITypeCcalendarStartTimeZone                   = "calendar:StartTimeZone"
	UnindexedFieldURITypeCcalendarEndTimeZone                     = "calendar:EndTimeZone"
	UnindexedFieldURITypeCcalendarJoinOnlineMeetingUrl            = "calendar:JoinOnlineMeetingUrl"
	UnindexedFieldURITypeCcalendarOnlineMeetingSettings           = "calendar:OnlineMeetingSettings"
	UnindexedFieldURITypeCcalendarIsOrganizer                     = "calendar:IsOrganizer"
	UnindexedFieldURITypeCcalendarCalendarActivityData            = "calendar:CalendarActivityData"
	UnindexedFieldURITypeCcalendarDoNotForwardMeeting             = "calendar:DoNotForwardMeeting"
	UnindexedFieldURITypeTtaskActualWork                          = "task:ActualWork"
	UnindexedFieldURITypeTtaskAssignedTime                        = "task:AssignedTime"
	UnindexedFieldURITypeTtaskBillingInformation                  = "task:BillingInformation"
	UnindexedFieldURITypeTtaskChangeCount                         = "task:ChangeCount"
	UnindexedFieldURITypeTtaskCompanies                           = "task:Companies"
	UnindexedFieldURITypeTtaskCompleteDate                        = "task:CompleteDate"
	UnindexedFieldURITypeTtaskContacts                            = "task:Contacts"
	UnindexedFieldURITypeTtaskDelegationState                     = "task:DelegationState"
	UnindexedFieldURITypeTtaskDelegator                           = "task:Delegator"
	UnindexedFieldURITypeTtaskDueDate                             = "task:DueDate"
	UnindexedFieldURITypeTtaskIsAssignmentEditable                = "task:IsAssignmentEditable"
	UnindexedFieldURITypeTtaskIsComplete                          = "task:IsComplete"
	UnindexedFieldURITypeTtaskIsRecurring                         = "task:IsRecurring"
	UnindexedFieldURITypeTtaskIsTeamTask                          = "task:IsTeamTask"
	UnindexedFieldURITypeTtaskMileage                             = "task:Mileage"
	UnindexedFieldURITypeTtaskOwner                               = "task:Owner"
	UnindexedFieldURITypeTtaskPercentComplete                     = "task:PercentComplete"
	UnindexedFieldURITypeTtaskRecurrence                          = "task:Recurrence"
	UnindexedFieldURITypeTtaskStartDate                           = "task:StartDate"
	UnindexedFieldURITypeTtaskStatus                              = "task:Status"
	UnindexedFieldURITypeTtaskStatusDescription                   = "task:StatusDescription"
	UnindexedFieldURITypeTtaskTotalWork                           = "task:TotalWork"
	UnindexedFieldURITypeCcontactsAlias                           = "contacts:Alias"
	UnindexedFieldURITypeCcontactsAssistantName                   = "contacts:AssistantName"
	UnindexedFieldURITypeCcontactsBirthday                        = "contacts:Birthday"
	UnindexedFieldURITypeCcontactsBusinessHomePage                = "contacts:BusinessHomePage"
	UnindexedFieldURITypeCcontactsChildren                        = "contacts:Children"
	UnindexedFieldURITypeCcontactsCompanies                       = "contacts:Companies"
	UnindexedFieldURITypeCcontactsCompanyName                     = "contacts:CompanyName"
	UnindexedFieldURITypeCcontactsCompleteName                    = "contacts:CompleteName"
	UnindexedFieldURITypeCcontactsContactSource                   = "contacts:ContactSource"
	UnindexedFieldURITypeCcontactsCulture                         = "contacts:Culture"
	UnindexedFieldURITypeCcontactsDepartment                      = "contacts:Department"
	UnindexedFieldURITypeCcontactsDisplayName                     = "contacts:DisplayName"
	UnindexedFieldURITypeCcontactsDirectoryId                     = "contacts:DirectoryId"
	UnindexedFieldURITypeCcontactsDirectReports                   = "contacts:DirectReports"
	UnindexedFieldURITypeCcontactsEmailAddresses                  = "contacts:EmailAddresses"
	UnindexedFieldURITypeCcontactsAbchEmailAddresses              = "contacts:AbchEmailAddresses"
	UnindexedFieldURITypeCcontactsFileAs                          = "contacts:FileAs"
	UnindexedFieldURITypeCcontactsFileAsMapping                   = "contacts:FileAsMapping"
	UnindexedFieldURITypeCcontactsGeneration                      = "contacts:Generation"
	UnindexedFieldURITypeCcontactsGivenName                       = "contacts:GivenName"
	UnindexedFieldURITypeCcontactsImAddresses                     = "contacts:ImAddresses"
	UnindexedFieldURITypeCcontactsInitials                        = "contacts:Initials"
	UnindexedFieldURITypeCcontactsJobTitle                        = "contacts:JobTitle"
	UnindexedFieldURITypeCcontactsManager                         = "contacts:Manager"
	UnindexedFieldURITypeCcontactsManagerMailbox                  = "contacts:ManagerMailbox"
	UnindexedFieldURITypeCcontactsMiddleName                      = "contacts:MiddleName"
	UnindexedFieldURITypeCcontactsMileage                         = "contacts:Mileage"
	UnindexedFieldURITypeCcontactsMSExchangeCertificate           = "contacts:MSExchangeCertificate"
	UnindexedFieldURITypeCcontactsNickname                        = "contacts:Nickname"
	UnindexedFieldURITypeCcontactsNotes                           = "contacts:Notes"
	UnindexedFieldURITypeCcontactsOfficeLocation                  = "contacts:OfficeLocation"
	UnindexedFieldURITypeCcontactsPhoneNumbers                    = "contacts:PhoneNumbers"
	UnindexedFieldURITypeCcontactsPhoneticFullName                = "contacts:PhoneticFullName"
	UnindexedFieldURITypeCcontactsPhoneticFirstName               = "contacts:PhoneticFirstName"
	UnindexedFieldURITypeCcontactsPhoneticLastName                = "contacts:PhoneticLastName"
	UnindexedFieldURITypeCcontactsPhoto                           = "contacts:Photo"
	UnindexedFieldURITypeCcontactsPhysicalAddresses               = "contacts:PhysicalAddresses"
	UnindexedFieldURITypeCcontactsPostalAddressIndex              = "contacts:PostalAddressIndex"
	UnindexedFieldURITypeCcontactsProfession                      = "contacts:Profession"
	UnindexedFieldURITypeCcontactsSpouseName                      = "contacts:SpouseName"
	UnindexedFieldURITypeCcontactsSurname                         = "contacts:Surname"
	UnindexedFieldURITypeCcontactsWeddingAnniversary              = "contacts:WeddingAnniversary"
	UnindexedFieldURITypeCcontactsUserSMIMECertificate            = "contacts:UserSMIMECertificate"
	UnindexedFieldURITypeCcontactsHasPicture                      = "contacts:HasPicture"
	UnindexedFieldURITypeCcontactsAccountName                     = "contacts:AccountName"
	UnindexedFieldURITypeCcontactsIsAutoUpdateDisabled            = "contacts:IsAutoUpdateDisabled"
	UnindexedFieldURITypeCcontactsIsMessengerEnabled              = "contacts:IsMessengerEnabled"
	UnindexedFieldURITypeCcontactsComment                         = "contacts:Comment"
	UnindexedFieldURITypeCcontactsContactShortId                  = "contacts:ContactShortId"
	UnindexedFieldURITypeCcontactsContactType                     = "contacts:ContactType"
	UnindexedFieldURITypeCcontactsCreatedBy                       = "contacts:CreatedBy"
	UnindexedFieldURITypeCcontactsGender                          = "contacts:Gender"
	UnindexedFieldURITypeCcontactsIsHidden                        = "contacts:IsHidden"
	UnindexedFieldURITypeCcontactsObjectId                        = "contacts:ObjectId"
	UnindexedFieldURITypeCcontactsPassportId                      = "contacts:PassportId"
	UnindexedFieldURITypeCcontactsIsPrivate                       = "contacts:IsPrivate"
	UnindexedFieldURITypeCcontactsSourceId                        = "contacts:SourceId"
	UnindexedFieldURITypeCcontactsTrustLevel                      = "contacts:TrustLevel"
	UnindexedFieldURITypeCcontactsUrls                            = "contacts:Urls"
	UnindexedFieldURITypeCcontactsCid                             = "contacts:Cid"
	UnindexedFieldURITypeCcontactsSkypeAuthCertificate            = "contacts:SkypeAuthCertificate"
	UnindexedFieldURITypeCcontactsSkypeContext                    = "contacts:SkypeContext"
	UnindexedFieldURITypeCcontactsSkypeId                         = "contacts:SkypeId"
	UnindexedFieldURITypeCcontactsXboxLiveTag                     = "contacts:XboxLiveTag"
	UnindexedFieldURITypeCcontactsSkypeRelationship               = "contacts:SkypeRelationship"
	UnindexedFieldURITypeCcontactsYomiNickname                    = "contacts:YomiNickname"
	UnindexedFieldURITypeCcontactsInviteFree                      = "contacts:InviteFree"
	UnindexedFieldURITypeCcontactsHidePresenceAndProfile          = "contacts:HidePresenceAndProfile"
	UnindexedFieldURITypeCcontactsIsPendingOutbound               = "contacts:IsPendingOutbound"
	UnindexedFieldURITypeCcontactsSupportGroupFeeds               = "contacts:SupportGroupFeeds"
	UnindexedFieldURITypeCcontactsUserTileHash                    = "contacts:UserTileHash"
	UnindexedFieldURITypeCcontactsUnifiedInbox                    = "contacts:UnifiedInbox"
	UnindexedFieldURITypeCcontactsMris                            = "contacts:Mris"
	UnindexedFieldURITypeCcontactsWlid                            = "contacts:Wlid"
	UnindexedFieldURITypeCcontactsAbchContactId                   = "contacts:AbchContactId"
	UnindexedFieldURITypeCcontactsNotInBirthdayCalendar           = "contacts:NotInBirthdayCalendar"
	UnindexedFieldURITypeCcontactsShellContactType                = "contacts:ShellContactType"
	UnindexedFieldURITypeCcontactsImMri                           = "contacts:ImMri"
	UnindexedFieldURITypeCcontactsPresenceTrustLevel              = "contacts:PresenceTrustLevel"
	UnindexedFieldURITypeCcontactsOtherMri                        = "contacts:OtherMri"
	UnindexedFieldURITypeCcontactsProfileLastChanged              = "contacts:ProfileLastChanged"
	UnindexedFieldURITypeCcontactsMobileIMEnabled                 = "contacts:MobileIMEnabled"
	UnindexedFieldURITypeDdistributionlistMembers                 = "distributionlist:Members"
	UnindexedFieldURITypeCcontactsPartnerNetworkProfilePhotoUrl   = "contacts:PartnerNetworkProfilePhotoUrl"
	UnindexedFieldURITypeCcontactsPartnerNetworkThumbnailPhotoUrl = "contacts:PartnerNetworkThumbnailPhotoUrl"
	UnindexedFieldURITypeCcontactsPersonId                        = "contacts:PersonId"
	UnindexedFieldURITypeCcontactsConversationGuid                = "contacts:ConversationGuid"
	UnindexedFieldURITypePpostitemPostedTime                      = "postitem:PostedTime"
	UnindexedFieldURITypeCconversationConversationId              = "conversation:ConversationId"
	UnindexedFieldURITypeCconversationConversationTopic           = "conversation:ConversationTopic"
	UnindexedFieldURITypeCconversationUniqueRecipients            = "conversation:UniqueRecipients"
	UnindexedFieldURITypeCconversationGlobalUniqueRecipients      = "conversation:GlobalUniqueRecipients"
	UnindexedFieldURITypeCconversationUniqueUnreadSenders         = "conversation:UniqueUnreadSenders"
	UnindexedFieldURITypeCconversationGlobalUniqueUnreadSenders   = "conversation:GlobalUniqueUnreadSenders"
	UnindexedFieldURITypeCconversationUniqueSenders               = "conversation:UniqueSenders"
	UnindexedFieldURITypeCconversationGlobalUniqueSenders         = "conversation:GlobalUniqueSenders"
	UnindexedFieldURITypeCconversationLastDeliveryTime            = "conversation:LastDeliveryTime"
	UnindexedFieldURITypeCconversationGlobalLastDeliveryTime      = "conversation:GlobalLastDeliveryTime"
	UnindexedFieldURITypeCconversationCategories                  = "conversation:Categories"
	UnindexedFieldURITypeCconversationGlobalCategories            = "conversation:GlobalCategories"
	UnindexedFieldURITypeCconversationFlagStatus                  = "conversation:FlagStatus"
	UnindexedFieldURITypeCconversationGlobalFlagStatus            = "conversation:GlobalFlagStatus"
	UnindexedFieldURITypeCconversationHasAttachments              = "conversation:HasAttachments"
	UnindexedFieldURITypeCconversationGlobalHasAttachments        = "conversation:GlobalHasAttachments"
	UnindexedFieldURITypeCconversationHasIrm                      = "conversation:HasIrm"
	UnindexedFieldURITypeCconversationGlobalHasIrm                = "conversation:GlobalHasIrm"
	UnindexedFieldURITypeCconversationMessageCount                = "conversation:MessageCount"
	UnindexedFieldURITypeCconversationGlobalMessageCount          = "conversation:GlobalMessageCount"
	UnindexedFieldURITypeCconversationUnreadCount                 = "conversation:UnreadCount"
	UnindexedFieldURITypeCconversationGlobalUnreadCount           = "conversation:GlobalUnreadCount"
	UnindexedFieldURITypeCconversationSize                        = "conversation:Size"
	UnindexedFieldURITypeCconversationGlobalSize                  = "conversation:GlobalSize"
	UnindexedFieldURITypeCconversationItemClasses                 = "conversation:ItemClasses"
	UnindexedFieldURITypeCconversationGlobalItemClasses           = "conversation:GlobalItemClasses"
	UnindexedFieldURITypeCconversationImportance                  = "conversation:Importance"
	UnindexedFieldURITypeCconversationGlobalImportance            = "conversation:GlobalImportance"
	UnindexedFieldURITypeCconversationItemIds                     = "conversation:ItemIds"
	UnindexedFieldURITypeCconversationGlobalItemIds               = "conversation:GlobalItemIds"
	UnindexedFieldURITypeCconversationLastModifiedTime            = "conversation:LastModifiedTime"
	UnindexedFieldURITypeCconversationInstanceKey                 = "conversation:InstanceKey"
	UnindexedFieldURITypeCconversationPreview                     = "conversation:Preview"
	UnindexedFieldURITypeCconversationIconIndex                   = "conversation:IconIndex"
	UnindexedFieldURITypeCconversationGlobalIconIndex             = "conversation:GlobalIconIndex"
	UnindexedFieldURITypeCconversationDraftItemIds                = "conversation:DraftItemIds"
	UnindexedFieldURITypeCconversationHasClutter                  = "conversation:HasClutter"
	UnindexedFieldURITypeCconversationMentionedMe                 = "conversation:MentionedMe"
	UnindexedFieldURITypeCconversationGlobalMentionedMe           = "conversation:GlobalMentionedMe"
	UnindexedFieldURITypeCconversationAtAllMention                = "conversation:AtAllMention"
	UnindexedFieldURITypeCconversationGlobalAtAllMention          = "conversation:GlobalAtAllMention"
	UnindexedFieldURITypePpersonFullName                          = "person:FullName"
	UnindexedFieldURITypePpersonGivenName                         = "person:GivenName"
	UnindexedFieldURITypePpersonSurname                           = "person:Surname"
	UnindexedFieldURITypePpersonPhoneNumber                       = "person:PhoneNumber"
	UnindexedFieldURITypePpersonSMSNumber                         = "person:SMSNumber"
	UnindexedFieldURITypePpersonEmailAddress                      = "person:EmailAddress"
	UnindexedFieldURITypePpersonAlias                             = "person:Alias"
	UnindexedFieldURITypePpersonDepartment                        = "person:Department"
	UnindexedFieldURITypePpersonLinkedInProfileLink               = "person:LinkedInProfileLink"
	UnindexedFieldURITypePpersonSkills                            = "person:Skills"
	UnindexedFieldURITypePpersonProfessionalBiography             = "person:ProfessionalBiography"
	UnindexedFieldURITypePpersonManagementChain                   = "person:ManagementChain"
	UnindexedFieldURITypePpersonDirectReports                     = "person:DirectReports"
	UnindexedFieldURITypePpersonPeers                             = "person:Peers"
	UnindexedFieldURITypePpersonTeamSize                          = "person:TeamSize"
	UnindexedFieldURITypePpersonCurrentJob                        = "person:CurrentJob"
	UnindexedFieldURITypePpersonBirthday                          = "person:Birthday"
	UnindexedFieldURITypePpersonHometown                          = "person:Hometown"
	UnindexedFieldURITypePpersonCurrentLocation                   = "person:CurrentLocation"
	UnindexedFieldURITypePpersonCompanyProfile                    = "person:CompanyProfile"
	UnindexedFieldURITypePpersonOffice                            = "person:Office"
	UnindexedFieldURITypePpersonHeadline                          = "person:Headline"
	UnindexedFieldURITypePpersonMutualConnections                 = "person:MutualConnections"
	UnindexedFieldURITypePpersonTitle                             = "person:Title"
	UnindexedFieldURITypePpersonMutualManager                     = "person:MutualManager"
	UnindexedFieldURITypePpersonInsights                          = "person:Insights"
	UnindexedFieldURITypePpersonUserProfilePicture                = "person:UserProfilePicture"
	UnindexedFieldURITypePpersonaPersonaId                        = "persona:PersonaId"
	UnindexedFieldURITypePpersonaPersonaType                      = "persona:PersonaType"
	UnindexedFieldURITypePpersonaGivenName                        = "persona:GivenName"
	UnindexedFieldURITypePpersonaCompanyName                      = "persona:CompanyName"
	UnindexedFieldURITypePpersonaSurname                          = "persona:Surname"
	UnindexedFieldURITypePpersonaDisplayName                      = "persona:DisplayName"
	UnindexedFieldURITypePpersonaEmailAddress                     = "persona:EmailAddress"
	UnindexedFieldURITypePpersonaFileAs                           = "persona:FileAs"
	UnindexedFieldURITypePpersonaHomeCity                         = "persona:HomeCity"
	UnindexedFieldURITypePpersonaCreationTime                     = "persona:CreationTime"
	UnindexedFieldURITypePpersonaRelevanceScore                   = "persona:RelevanceScore"
	UnindexedFieldURITypePpersonaRankingWeight                    = "persona:RankingWeight"
	UnindexedFieldURITypePpersonaWorkCity                         = "persona:WorkCity"
	UnindexedFieldURITypePpersonaPersonaObjectStatus              = "persona:PersonaObjectStatus"
	UnindexedFieldURITypePpersonaFileAsId                         = "persona:FileAsId"
	UnindexedFieldURITypePpersonaDisplayNamePrefix                = "persona:DisplayNamePrefix"
	UnindexedFieldURITypePpersonaYomiCompanyName                  = "persona:YomiCompanyName"
	UnindexedFieldURITypePpersonaYomiFirstName                    = "persona:YomiFirstName"
	UnindexedFieldURITypePpersonaYomiLastName                     = "persona:YomiLastName"
	UnindexedFieldURITypePpersonaTitle                            = "persona:Title"
	UnindexedFieldURITypePpersonaEmailAddresses                   = "persona:EmailAddresses"
	UnindexedFieldURITypePpersonaPhoneNumber                      = "persona:PhoneNumber"
	UnindexedFieldURITypePpersonaImAddress                        = "persona:ImAddress"
	UnindexedFieldURITypePpersonaImAddresses                      = "persona:ImAddresses"
	UnindexedFieldURITypePpersonaImAddresses2                     = "persona:ImAddresses2"
	UnindexedFieldURITypePpersonaImAddresses3                     = "persona:ImAddresses3"
	UnindexedFieldURITypePpersonaFolderIds                        = "persona:FolderIds"
	UnindexedFieldURITypePpersonaAttributions                     = "persona:Attributions"
	UnindexedFieldURITypePpersonaDisplayNames                     = "persona:DisplayNames"
	UnindexedFieldURITypePpersonaInitials                         = "persona:Initials"
	UnindexedFieldURITypePpersonaFileAses                         = "persona:FileAses"
	UnindexedFieldURITypePpersonaFileAsIds                        = "persona:FileAsIds"
	UnindexedFieldURITypePpersonaDisplayNamePrefixes              = "persona:DisplayNamePrefixes"
	UnindexedFieldURITypePpersonaGivenNames                       = "persona:GivenNames"
	UnindexedFieldURITypePpersonaMiddleNames                      = "persona:MiddleNames"
	UnindexedFieldURITypePpersonaSurnames                         = "persona:Surnames"
	UnindexedFieldURITypePpersonaGenerations                      = "persona:Generations"
	UnindexedFieldURITypePpersonaNicknames                        = "persona:Nicknames"
	UnindexedFieldURITypePpersonaYomiCompanyNames                 = "persona:YomiCompanyNames"
	UnindexedFieldURITypePpersonaYomiFirstNames                   = "persona:YomiFirstNames"
	UnindexedFieldURITypePpersonaYomiLastNames                    = "persona:YomiLastNames"
	UnindexedFieldURITypePpersonaBusinessPhoneNumbers             = "persona:BusinessPhoneNumbers"
	UnindexedFieldURITypePpersonaBusinessPhoneNumbers2            = "persona:BusinessPhoneNumbers2"
	UnindexedFieldURITypePpersonaHomePhones                       = "persona:HomePhones"
	UnindexedFieldURITypePpersonaHomePhones2                      = "persona:HomePhones2"
	UnindexedFieldURITypePpersonaMobilePhones                     = "persona:MobilePhones"
	UnindexedFieldURITypePpersonaMobilePhones2                    = "persona:MobilePhones2"
	UnindexedFieldURITypePpersonaAssistantPhoneNumbers            = "persona:AssistantPhoneNumbers"
	UnindexedFieldURITypePpersonaCallbackPhones                   = "persona:CallbackPhones"
	UnindexedFieldURITypePpersonaCarPhones                        = "persona:CarPhones"
	UnindexedFieldURITypePpersonaHomeFaxes                        = "persona:HomeFaxes"
	UnindexedFieldURITypePpersonaOrganizationMainPhones           = "persona:OrganizationMainPhones"
	UnindexedFieldURITypePpersonaOtherFaxes                       = "persona:OtherFaxes"
	UnindexedFieldURITypePpersonaOtherTelephones                  = "persona:OtherTelephones"
	UnindexedFieldURITypePpersonaOtherPhones2                     = "persona:OtherPhones2"
	UnindexedFieldURITypePpersonaPagers                           = "persona:Pagers"
	UnindexedFieldURITypePpersonaRadioPhones                      = "persona:RadioPhones"
	UnindexedFieldURITypePpersonaTelexNumbers                     = "persona:TelexNumbers"
	UnindexedFieldURITypePpersonaWorkFaxes                        = "persona:WorkFaxes"
	UnindexedFieldURITypePpersonaEmails1                          = "persona:Emails1"
	UnindexedFieldURITypePpersonaEmails2                          = "persona:Emails2"
	UnindexedFieldURITypePpersonaEmails3                          = "persona:Emails3"
	UnindexedFieldURITypePpersonaBusinessHomePages                = "persona:BusinessHomePages"
	UnindexedFieldURITypePpersonaSchool                           = "persona:School"
	UnindexedFieldURITypePpersonaPersonalHomePages                = "persona:PersonalHomePages"
	UnindexedFieldURITypePpersonaOfficeLocations                  = "persona:OfficeLocations"
	UnindexedFieldURITypePpersonaBusinessAddresses                = "persona:BusinessAddresses"
	UnindexedFieldURITypePpersonaHomeAddresses                    = "persona:HomeAddresses"
	UnindexedFieldURITypePpersonaOtherAddresses                   = "persona:OtherAddresses"
	UnindexedFieldURITypePpersonaTitles                           = "persona:Titles"
	UnindexedFieldURITypePpersonaDepartments                      = "persona:Departments"
	UnindexedFieldURITypePpersonaCompanyNames                     = "persona:CompanyNames"
	UnindexedFieldURITypePpersonaManagers                         = "persona:Managers"
	UnindexedFieldURITypePpersonaAssistantNames                   = "persona:AssistantNames"
	UnindexedFieldURITypePpersonaProfessions                      = "persona:Professions"
	UnindexedFieldURITypePpersonaSpouseNames                      = "persona:SpouseNames"
	UnindexedFieldURITypePpersonaHobbies                          = "persona:Hobbies"
	UnindexedFieldURITypePpersonaWeddingAnniversaries             = "persona:WeddingAnniversaries"
	UnindexedFieldURITypePpersonaBirthdays                        = "persona:Birthdays"
	UnindexedFieldURITypePpersonaChildren                         = "persona:Children"
	UnindexedFieldURITypePpersonaLocations                        = "persona:Locations"
	UnindexedFieldURITypePpersonaExtendedProperties               = "persona:ExtendedProperties"
	UnindexedFieldURITypePpersonaPostalAddress                    = "persona:PostalAddress"
	UnindexedFieldURITypePpersonaBodies                           = "persona:Bodies"
	UnindexedFieldURITypePpersonaIsFavorite                       = "persona:IsFavorite"
	UnindexedFieldURITypePpersonaInlineLinks                      = "persona:InlineLinks"
	UnindexedFieldURITypePpersonaItemLinkIds                      = "persona:ItemLinkIds"
	UnindexedFieldURITypePpersonaHasActiveDeals                   = "persona:HasActiveDeals"
	UnindexedFieldURITypePpersonaIsBusinessContact                = "persona:IsBusinessContact"
	UnindexedFieldURITypePpersonaAttributedHasActiveDeals         = "persona:AttributedHasActiveDeals"
	UnindexedFieldURITypePpersonaAttributedIsBusinessContact      = "persona:AttributedIsBusinessContact"
	UnindexedFieldURITypePpersonaSourceMailboxGuids               = "persona:SourceMailboxGuids"
	UnindexedFieldURITypePpersonaLastContactedDate                = "persona:LastContactedDate"
	UnindexedFieldURITypePpersonaExternalDirectoryObjectId        = "persona:ExternalDirectoryObjectId"
	UnindexedFieldURITypePpersonaMapiEntryId                      = "persona:MapiEntryId"
	UnindexedFieldURITypePpersonaMapiEmailAddress                 = "persona:MapiEmailAddress"
	UnindexedFieldURITypePpersonaMapiAddressType                  = "persona:MapiAddressType"
	UnindexedFieldURITypePpersonaMapiSearchKey                    = "persona:MapiSearchKey"
	UnindexedFieldURITypePpersonaMapiTransmittableDisplayName     = "persona:MapiTransmittableDisplayName"
	UnindexedFieldURITypePpersonaMapiSendRichInfo                 = "persona:MapiSendRichInfo"
	UnindexedFieldURITypeRrolememberMemberType                    = "rolemember:MemberType"
	UnindexedFieldURITypeRrolememberMemberId                      = "rolemember:MemberId"
	UnindexedFieldURITypeRrolememberDisplayName                   = "rolemember:DisplayName"
	UnindexedFieldURITypeNnetworkTokenRefreshLastCompleted        = "network:TokenRefreshLastCompleted"
	UnindexedFieldURITypeNnetworkTokenRefreshLastAttempted        = "network:TokenRefreshLastAttempted"
	UnindexedFieldURITypeNnetworkSyncEnabled                      = "network:SyncEnabled"
	UnindexedFieldURITypeNnetworkRejectedOffers                   = "network:RejectedOffers"
	UnindexedFieldURITypeNnetworkSessionHandle                    = "network:SessionHandle"
	UnindexedFieldURITypeNnetworkRefreshTokenExpiry2              = "network:RefreshTokenExpiry2"
	UnindexedFieldURITypeNnetworkRefreshToken2                    = "network:RefreshToken2"
	UnindexedFieldURITypeNnetworkPsaLastChanged                   = "network:PsaLastChanged"
	UnindexedFieldURITypeNnetworkOffers                           = "network:Offers"
	UnindexedFieldURITypeNnetworkLastWelcomeContact               = "network:LastWelcomeContact"
	UnindexedFieldURITypeNnetworkLastVersionSaved                 = "network:LastVersionSaved"
	UnindexedFieldURITypeNnetworkDomainTag                        = "network:DomainTag"
	UnindexedFieldURITypeNnetworkFirstAuthErrorDates              = "network:FirstAuthErrorDates"
	UnindexedFieldURITypeNnetworkErrorOffers                      = "network:ErrorOffers"
	UnindexedFieldURITypeNnetworkContactSyncSuccess               = "network:ContactSyncSuccess"
	UnindexedFieldURITypeNnetworkContactSyncError                 = "network:ContactSyncError"
	UnindexedFieldURITypeNnetworkClientToken2                     = "network:ClientToken2"
	UnindexedFieldURITypeNnetworkClientToken                      = "network:ClientToken"
	UnindexedFieldURITypeNnetworkClientPublishSecret              = "network:ClientPublishSecret"
	UnindexedFieldURITypeNnetworkUserEmail                        = "network:UserEmail"
	UnindexedFieldURITypeNnetworkAutoLinkSuccess                  = "network:AutoLinkSuccess"
	UnindexedFieldURITypeNnetworkAutoLinkError                    = "network:AutoLinkError"
	UnindexedFieldURITypeNnetworkIsDefault                        = "network:IsDefault"
	UnindexedFieldURITypeNnetworkSettings                         = "network:Settings"
	UnindexedFieldURITypeNnetworkProfileUrl                       = "network:ProfileUrl"
	UnindexedFieldURITypeNnetworkUserTileUrl                      = "network:UserTileUrl"
	UnindexedFieldURITypeNnetworkDomainId                         = "network:DomainId"
	UnindexedFieldURITypeNnetworkDisplayName                      = "network:DisplayName"
	UnindexedFieldURITypeNnetworkAccountName                      = "network:AccountName"
	UnindexedFieldURITypeNnetworkSourceEntryID                    = "network:SourceEntryID"
	UnindexedFieldURITypeAabchpersonFavoriteOrder                 = "abchperson:FavoriteOrder"
	UnindexedFieldURITypeAabchpersonPersonId                      = "abchperson:PersonId"
	UnindexedFieldURITypeAabchpersonExchangePersonIdGuid          = "abchperson:ExchangePersonIdGuid"
	UnindexedFieldURITypeAabchpersonAntiLinkInfo                  = "abchperson:AntiLinkInfo"
	UnindexedFieldURITypeAabchpersonRelevanceOrder1               = "abchperson:RelevanceOrder1"
	UnindexedFieldURITypeAabchpersonRelevanceOrder2               = "abchperson:RelevanceOrder2"
	UnindexedFieldURITypeAabchpersonContactHandles                = "abchperson:ContactHandles"
	UnindexedFieldURITypeAabchpersonCategories                    = "abchperson:Categories"
	UnindexedFieldURITypeBbookingServiceIds                       = "booking:ServiceIds"
	UnindexedFieldURITypeBbookingStaffIds                         = "booking:StaffIds"
	UnindexedFieldURITypeBbookingStaffInitials                    = "booking:StaffInitials"
	UnindexedFieldURITypeBbookingCustomerName                     = "booking:CustomerName"
	UnindexedFieldURITypeBbookingCustomerEmail                    = "booking:CustomerEmail"
	UnindexedFieldURITypeBbookingCustomerPhone                    = "booking:CustomerPhone"
	UnindexedFieldURITypeBbookingCustomerId                       = "booking:CustomerId"
	UnindexedFieldURITypeIinsightInsightId                        = "insight:InsightId"
	UnindexedFieldURITypeIinsightType                             = "insight:Type"
	UnindexedFieldURITypeIinsightStartTimeUtc                     = "insight:StartTimeUtc"
	UnindexedFieldURITypeIinsightEndTimeUtc                       = "insight:EndTimeUtc"
	UnindexedFieldURITypeIinsightStatus                           = "insight:Status"
	UnindexedFieldURITypeIinsightVersion                          = "insight:Version"
	UnindexedFieldURITypeIinsightApplicationsIds                  = "insight:ApplicationsIds"
	UnindexedFieldURITypeIinsightText                             = "insight:Text"
	UnindexedFieldURITypeIinsightSuggestedActions                 = "insight:SuggestedActions"
	UnindexedFieldURITypeIinsightAppContexts                      = "insight:AppContexts"
)

type DictionaryURIType XsString

const (
	DictionaryURITypeIitemInternetMessageHeader              = "item:InternetMessageHeader"
	DictionaryURITypeCcontactsImAddress                      = "contacts:ImAddress"
	DictionaryURITypeCcontactsPhysicalAddressStreet          = "contacts:PhysicalAddress:Street"
	DictionaryURITypeCcontactsPhysicalAddressCity            = "contacts:PhysicalAddress:City"
	DictionaryURITypeCcontactsPhysicalAddressState           = "contacts:PhysicalAddress:State"
	DictionaryURITypeCcontactsPhysicalAddressCountryOrRegion = "contacts:PhysicalAddress:CountryOrRegion"
	DictionaryURITypeCcontactsPhysicalAddressPostalCode      = "contacts:PhysicalAddress:PostalCode"
	DictionaryURITypeCcontactsPhoneNumber                    = "contacts:PhoneNumber"
	DictionaryURITypeCcontactsEmailAddress                   = "contacts:EmailAddress"
	DictionaryURITypeDdistributionlistMembersMember          = "distributionlist:Members:Member"
)

type AggregateType XsString

const (
	AggregateTypeMinimum = "Minimum"
	AggregateTypeMaximum = "Maximum"
)

type StandardGroupByType XsString

const (
	StandardGroupByTypeConversationTopic = "ConversationTopic"
)

type ItemQueryTraversalType XsString

const (
	ItemQueryTraversalTypeShallow     = "Shallow"
	ItemQueryTraversalTypeSoftDeleted = "SoftDeleted"
	ItemQueryTraversalTypeAssociated  = "Associated"
)

type DateTimePrecisionType XsString

const (
	DateTimePrecisionTypeSeconds      = "Seconds"
	DateTimePrecisionTypeMilliseconds = "Milliseconds"
)

type CreateActionType XsString

const (
	CreateActionTypeCreateNew      = "CreateNew"
	CreateActionTypeUpdate         = "Update"
	CreateActionTypeUpdateOrCreate = "UpdateOrCreate"
)

type DisposalType XsString

const (
	DisposalTypeHardDelete         = "HardDelete"
	DisposalTypeSoftDelete         = "SoftDelete"
	DisposalTypeMoveToDeletedItems = "MoveToDeletedItems"
)

type NotificationEventTypeType XsString

const (
	NotificationEventTypeTypeCopiedEvent          = "CopiedEvent"
	NotificationEventTypeTypeCreatedEvent         = "CreatedEvent"
	NotificationEventTypeTypeDeletedEvent         = "DeletedEvent"
	NotificationEventTypeTypeModifiedEvent        = "ModifiedEvent"
	NotificationEventTypeTypeMovedEvent           = "MovedEvent"
	NotificationEventTypeTypeNewMailEvent         = "NewMailEvent"
	NotificationEventTypeTypeFreeBusyChangedEvent = "FreeBusyChangedEvent"
)

type SubscriptionTimeoutType XsInt

type SubscriptionStatusFrequencyType XsInt

type StreamingSubscriptionConnectionTimeoutType XsInt

type MaxSyncChangesReturnedType XsInt

type SyncFolderItemsScopeType XsString

const (
	SyncFolderItemsScopeTypeNormalItems              = "NormalItems"
	SyncFolderItemsScopeTypeNormalAndAssociatedItems = "NormalAndAssociatedItems"
)

type DerivedItemIdType XsString

type MessageDispositionType XsString

const (
	MessageDispositionTypeSaveOnly        = "SaveOnly"
	MessageDispositionTypeSendOnly        = "SendOnly"
	MessageDispositionTypeSendAndSaveCopy = "SendAndSaveCopy"
)

type CalendarItemCreateOrDeleteOperationType XsString

const (
	CalendarItemCreateOrDeleteOperationTypeSendToNone           = "SendToNone"
	CalendarItemCreateOrDeleteOperationTypeSendOnlyToAll        = "SendOnlyToAll"
	CalendarItemCreateOrDeleteOperationTypeSendToAllAndSaveCopy = "SendToAllAndSaveCopy"
)

type AffectedTaskOccurrencesType XsString

const (
	AffectedTaskOccurrencesTypeAllOccurrences          = "AllOccurrences"
	AffectedTaskOccurrencesTypeSpecifiedOccurrenceOnly = "SpecifiedOccurrenceOnly"
)

type ConflictResolutionType XsString

const (
	ConflictResolutionTypeNeverOverwrite  = "NeverOverwrite"
	ConflictResolutionTypeAutoResolve     = "AutoResolve"
	ConflictResolutionTypeAlwaysOverwrite = "AlwaysOverwrite"
)

type CalendarItemUpdateOperationType XsString

const (
	CalendarItemUpdateOperationTypeSendToNone               = "SendToNone"
	CalendarItemUpdateOperationTypeSendOnlyToAll            = "SendOnlyToAll"
	CalendarItemUpdateOperationTypeSendOnlyToChanged        = "SendOnlyToChanged"
	CalendarItemUpdateOperationTypeSendToAllAndSaveCopy     = "SendToAllAndSaveCopy"
	CalendarItemUpdateOperationTypeSendToChangedAndSaveCopy = "SendToChangedAndSaveCopy"
)

type DelegateFolderPermissionLevelType XsString

const (
	DelegateFolderPermissionLevelTypeNone     = "None"
	DelegateFolderPermissionLevelTypeEditor   = "Editor"
	DelegateFolderPermissionLevelTypeReviewer = "Reviewer"
	DelegateFolderPermissionLevelTypeAuthor   = "Author"
	DelegateFolderPermissionLevelTypeCustom   = "Custom"
)

type DeliverMeetingRequestsType XsString

const (
	DeliverMeetingRequestsTypeDelegatesOnly                   = "DelegatesOnly"
	DeliverMeetingRequestsTypeDelegatesAndMe                  = "DelegatesAndMe"
	DeliverMeetingRequestsTypeDelegatesAndSendInformationToMe = "DelegatesAndSendInformationToMe"
	DeliverMeetingRequestsTypeNoForward                       = "NoForward"
)

type UserConfigurationPropertyType []UserConfigurationPropertyTypeItem

type MeetingAttendeeType XsString

const (
	MeetingAttendeeTypeOrganizer = "Organizer"
	MeetingAttendeeTypeRequired  = "Required"
	MeetingAttendeeTypeOptional  = "Optional"
	MeetingAttendeeTypeRoom      = "Room"
	MeetingAttendeeTypeResource  = "Resource"
)

type FreeBusyViewType []FreeBusyViewTypeItem

type SuggestionQuality XsString

const (
	SuggestionQualityExcellent = "Excellent"
	SuggestionQualityGood      = "Good"
	SuggestionQualityFair      = "Fair"
	SuggestionQualityPoor      = "Poor"
)

type OofState XsString

const (
	OofStateDisabled  = "Disabled"
	OofStateEnabled   = "Enabled"
	OofStateScheduled = "Scheduled"
)

type ExternalAudience XsString

const (
	ExternalAudienceNone  = "None"
	ExternalAudienceKnown = "Known"
	ExternalAudienceAll   = "All"
)

type ServiceConfigurationType []ServiceConfigurationTypeItem

type ProtectionRuleAllInternalType XsString

type ProtectionRuleValueType XsString

type ProtectionRuleTrueType XsString

type ProtectionRuleActionKindType XsString

const (
	ProtectionRuleActionKindTypeRightsProtectMessage = "RightsProtectMessage"
)

type MailTipTypes []MailTipTypesItem

type PhoneCallStateType XsString

const (
	PhoneCallStateTypeIdle         = "Idle"
	PhoneCallStateTypeConnecting   = "Connecting"
	PhoneCallStateTypeAlerted      = "Alerted"
	PhoneCallStateTypeConnected    = "Connected"
	PhoneCallStateTypeDisconnected = "Disconnected"
	PhoneCallStateTypeIncoming     = "Incoming"
	PhoneCallStateTypeTransferring = "Transferring"
	PhoneCallStateTypeForwarding   = "Forwarding"
)

type ConnectionFailureCauseType XsString

const (
	ConnectionFailureCauseTypeNone        = "None"
	ConnectionFailureCauseTypeUserBusy    = "UserBusy"
	ConnectionFailureCauseTypeNoAnswer    = "NoAnswer"
	ConnectionFailureCauseTypeUnavailable = "Unavailable"
	ConnectionFailureCauseTypeOther       = "Other"
)

type SharingDataType XsString

const (
	SharingDataTypeCalendar = "Calendar"
	SharingDataTypeContacts = "Contacts"
)

type TeamMailboxLifecycleStateType XsString

const (
	TeamMailboxLifecycleStateTypeActive        = "Active"
	TeamMailboxLifecycleStateTypeClosed        = "Closed"
	TeamMailboxLifecycleStateTypeUnlinked      = "Unlinked"
	TeamMailboxLifecycleStateTypePendingDelete = "PendingDelete"
)

type ConversationQueryTraversalType XsString

const (
	ConversationQueryTraversalTypeShallow = "Shallow"
	ConversationQueryTraversalTypeDeep    = "Deep"
)

type ViewFilterType XsString

const (
	ViewFilterTypeAll           = "All"
	ViewFilterTypeFlagged       = "Flagged"
	ViewFilterTypeHasAttachment = "HasAttachment"
	ViewFilterTypeToOrCcMe      = "ToOrCcMe"
	ViewFilterTypeUnread        = "Unread"
	ViewFilterTypeTaskActive    = "TaskActive"
	ViewFilterTypeTaskOverdue   = "TaskOverdue"
	ViewFilterTypeTaskCompleted = "TaskCompleted"
	ViewFilterTypeNoClutter     = "NoClutter"
	ViewFilterTypeClutter       = "Clutter"
)

type ConversationActionTypeType XsString

const (
	ConversationActionTypeTypeAlwaysCategorize   = "AlwaysCategorize"
	ConversationActionTypeTypeAlwaysDelete       = "AlwaysDelete"
	ConversationActionTypeTypeAlwaysMove         = "AlwaysMove"
	ConversationActionTypeTypeDelete             = "Delete"
	ConversationActionTypeTypeMove               = "Move"
	ConversationActionTypeTypeCopy               = "Copy"
	ConversationActionTypeTypeSetReadState       = "SetReadState"
	ConversationActionTypeTypeSetRetentionPolicy = "SetRetentionPolicy"
	ConversationActionTypeTypeFlag               = "Flag"
)

type RetentionType XsString

const (
	RetentionTypeDelete  = "Delete"
	RetentionTypeArchive = "Archive"
)

type ConversationNodeSortOrder XsString

const (
	ConversationNodeSortOrderTreeOrderAscending  = "TreeOrderAscending"
	ConversationNodeSortOrderTreeOrderDescending = "TreeOrderDescending"
	ConversationNodeSortOrderDateOrderAscending  = "DateOrderAscending"
	ConversationNodeSortOrderDateOrderDescending = "DateOrderDescending"
)

type FlaggedForActionType XsString

const (
	FlaggedForActionTypeAny                 = "Any"
	FlaggedForActionTypeCall                = "Call"
	FlaggedForActionTypeDoNotForward        = "DoNotForward"
	FlaggedForActionTypeFollowUp            = "FollowUp"
	FlaggedForActionTypeFYI                 = "FYI"
	FlaggedForActionTypeForward             = "Forward"
	FlaggedForActionTypeNoResponseNecessary = "NoResponseNecessary"
	FlaggedForActionTypeRead                = "Read"
	FlaggedForActionTypeReply               = "Reply"
	FlaggedForActionTypeReplyToAll          = "ReplyToAll"
	FlaggedForActionTypeReview              = "Review"
)

type RuleFieldURIType XsString

const (
	RuleFieldURITypeRuleId                                = "RuleId"
	RuleFieldURITypeDisplayName                           = "DisplayName"
	RuleFieldURITypePriority                              = "Priority"
	RuleFieldURITypeIsNotSupported                        = "IsNotSupported"
	RuleFieldURITypeActions                               = "Actions"
	RuleFieldURITypeConditionCategories                   = "Condition:Categories"
	RuleFieldURITypeConditionContainsBodyStrings          = "Condition:ContainsBodyStrings"
	RuleFieldURITypeConditionContainsHeaderStrings        = "Condition:ContainsHeaderStrings"
	RuleFieldURITypeConditionContainsRecipientStrings     = "Condition:ContainsRecipientStrings"
	RuleFieldURITypeConditionContainsSenderStrings        = "Condition:ContainsSenderStrings"
	RuleFieldURITypeConditionContainsSubjectOrBodyStrings = "Condition:ContainsSubjectOrBodyStrings"
	RuleFieldURITypeConditionContainsSubjectStrings       = "Condition:ContainsSubjectStrings"
	RuleFieldURITypeConditionFlaggedForAction             = "Condition:FlaggedForAction"
	RuleFieldURITypeConditionFromAddresses                = "Condition:FromAddresses"
	RuleFieldURITypeConditionFromConnectedAccounts        = "Condition:FromConnectedAccounts"
	RuleFieldURITypeConditionHasAttachments               = "Condition:HasAttachments"
	RuleFieldURITypeConditionImportance                   = "Condition:Importance"
	RuleFieldURITypeConditionIsApprovalRequest            = "Condition:IsApprovalRequest"
	RuleFieldURITypeConditionIsAutomaticForward           = "Condition:IsAutomaticForward"
	RuleFieldURITypeConditionIsAutomaticReply             = "Condition:IsAutomaticReply"
	RuleFieldURITypeConditionIsEncrypted                  = "Condition:IsEncrypted"
	RuleFieldURITypeConditionIsMeetingRequest             = "Condition:IsMeetingRequest"
	RuleFieldURITypeConditionIsMeetingResponse            = "Condition:IsMeetingResponse"
	RuleFieldURITypeConditionIsNDR                        = "Condition:IsNDR"
	RuleFieldURITypeConditionIsPermissionControlled       = "Condition:IsPermissionControlled"
	RuleFieldURITypeConditionIsReadReceipt                = "Condition:IsReadReceipt"
	RuleFieldURITypeConditionIsSigned                     = "Condition:IsSigned"
	RuleFieldURITypeConditionIsVoicemail                  = "Condition:IsVoicemail"
	RuleFieldURITypeConditionItemClasses                  = "Condition:ItemClasses"
	RuleFieldURITypeConditionMessageClassifications       = "Condition:MessageClassifications"
	RuleFieldURITypeConditionNotSentToMe                  = "Condition:NotSentToMe"
	RuleFieldURITypeConditionSentCcMe                     = "Condition:SentCcMe"
	RuleFieldURITypeConditionSentOnlyToMe                 = "Condition:SentOnlyToMe"
	RuleFieldURITypeConditionSentToAddresses              = "Condition:SentToAddresses"
	RuleFieldURITypeConditionSentToMe                     = "Condition:SentToMe"
	RuleFieldURITypeConditionSentToOrCcMe                 = "Condition:SentToOrCcMe"
	RuleFieldURITypeConditionSensitivity                  = "Condition:Sensitivity"
	RuleFieldURITypeConditionWithinDateRange              = "Condition:WithinDateRange"
	RuleFieldURITypeConditionWithinSizeRange              = "Condition:WithinSizeRange"
	RuleFieldURITypeExceptionCategories                   = "Exception:Categories"
	RuleFieldURITypeExceptionContainsBodyStrings          = "Exception:ContainsBodyStrings"
	RuleFieldURITypeExceptionContainsHeaderStrings        = "Exception:ContainsHeaderStrings"
	RuleFieldURITypeExceptionContainsRecipientStrings     = "Exception:ContainsRecipientStrings"
	RuleFieldURITypeExceptionContainsSenderStrings        = "Exception:ContainsSenderStrings"
	RuleFieldURITypeExceptionContainsSubjectOrBodyStrings = "Exception:ContainsSubjectOrBodyStrings"
	RuleFieldURITypeExceptionContainsSubjectStrings       = "Exception:ContainsSubjectStrings"
	RuleFieldURITypeExceptionFlaggedForAction             = "Exception:FlaggedForAction"
	RuleFieldURITypeExceptionFromAddresses                = "Exception:FromAddresses"
	RuleFieldURITypeExceptionFromConnectedAccounts        = "Exception:FromConnectedAccounts"
	RuleFieldURITypeExceptionHasAttachments               = "Exception:HasAttachments"
	RuleFieldURITypeExceptionImportance                   = "Exception:Importance"
	RuleFieldURITypeExceptionIsApprovalRequest            = "Exception:IsApprovalRequest"
	RuleFieldURITypeExceptionIsAutomaticForward           = "Exception:IsAutomaticForward"
	RuleFieldURITypeExceptionIsAutomaticReply             = "Exception:IsAutomaticReply"
	RuleFieldURITypeExceptionIsEncrypted                  = "Exception:IsEncrypted"
	RuleFieldURITypeExceptionIsMeetingRequest             = "Exception:IsMeetingRequest"
	RuleFieldURITypeExceptionIsMeetingResponse            = "Exception:IsMeetingResponse"
	RuleFieldURITypeExceptionIsNDR                        = "Exception:IsNDR"
	RuleFieldURITypeExceptionIsPermissionControlled       = "Exception:IsPermissionControlled"
	RuleFieldURITypeExceptionIsReadReceipt                = "Exception:IsReadReceipt"
	RuleFieldURITypeExceptionIsSigned                     = "Exception:IsSigned"
	RuleFieldURITypeExceptionIsVoicemail                  = "Exception:IsVoicemail"
	RuleFieldURITypeExceptionItemClasses                  = "Exception:ItemClasses"
	RuleFieldURITypeExceptionMessageClassifications       = "Exception:MessageClassifications"
	RuleFieldURITypeExceptionNotSentToMe                  = "Exception:NotSentToMe"
	RuleFieldURITypeExceptionSentCcMe                     = "Exception:SentCcMe"
	RuleFieldURITypeExceptionSentOnlyToMe                 = "Exception:SentOnlyToMe"
	RuleFieldURITypeExceptionSentToAddresses              = "Exception:SentToAddresses"
	RuleFieldURITypeExceptionSentToMe                     = "Exception:SentToMe"
	RuleFieldURITypeExceptionSentToOrCcMe                 = "Exception:SentToOrCcMe"
	RuleFieldURITypeExceptionSensitivity                  = "Exception:Sensitivity"
	RuleFieldURITypeExceptionWithinDateRange              = "Exception:WithinDateRange"
	RuleFieldURITypeExceptionWithinSizeRange              = "Exception:WithinSizeRange"
	RuleFieldURITypeActionAssignCategories                = "Action:AssignCategories"
	RuleFieldURITypeActionCopyToFolder                    = "Action:CopyToFolder"
	RuleFieldURITypeActionDelete                          = "Action:Delete"
	RuleFieldURITypeActionForwardAsAttachmentToRecipients = "Action:ForwardAsAttachmentToRecipients"
	RuleFieldURITypeActionForwardToRecipients             = "Action:ForwardToRecipients"
	RuleFieldURITypeActionMarkImportance                  = "Action:MarkImportance"
	RuleFieldURITypeActionMarkAsRead                      = "Action:MarkAsRead"
	RuleFieldURITypeActionMoveToFolder                    = "Action:MoveToFolder"
	RuleFieldURITypeActionPermanentDelete                 = "Action:PermanentDelete"
	RuleFieldURITypeActionRedirectToRecipients            = "Action:RedirectToRecipients"
	RuleFieldURITypeActionSendSMSAlertToRecipients        = "Action:SendSMSAlertToRecipients"
	RuleFieldURITypeActionServerReplyWithMessage          = "Action:ServerReplyWithMessage"
	RuleFieldURITypeActionStopProcessingRules             = "Action:StopProcessingRules"
	RuleFieldURITypeIsEnabled                             = "IsEnabled"
	RuleFieldURITypeIsInError                             = "IsInError"
	RuleFieldURITypeConditions                            = "Conditions"
	RuleFieldURITypeExceptions                            = "Exceptions"
)

type RuleValidationErrorCodeType XsString

const (
	RuleValidationErrorCodeTypeADOperationFailure               = "ADOperationFailure"
	RuleValidationErrorCodeTypeConnectedAccountNotFound         = "ConnectedAccountNotFound"
	RuleValidationErrorCodeTypeCreateWithRuleId                 = "CreateWithRuleId"
	RuleValidationErrorCodeTypeEmptyValueFound                  = "EmptyValueFound"
	RuleValidationErrorCodeTypeDuplicatedPriority               = "DuplicatedPriority"
	RuleValidationErrorCodeTypeDuplicatedOperationOnTheSameRule = "DuplicatedOperationOnTheSameRule"
	RuleValidationErrorCodeTypeFolderDoesNotExist               = "FolderDoesNotExist"
	RuleValidationErrorCodeTypeInvalidAddress                   = "InvalidAddress"
	RuleValidationErrorCodeTypeInvalidDateRange                 = "InvalidDateRange"
	RuleValidationErrorCodeTypeInvalidFolderId                  = "InvalidFolderId"
	RuleValidationErrorCodeTypeInvalidSizeRange                 = "InvalidSizeRange"
	RuleValidationErrorCodeTypeInvalidValue                     = "InvalidValue"
	RuleValidationErrorCodeTypeMessageClassificationNotFound    = "MessageClassificationNotFound"
	RuleValidationErrorCodeTypeMissingAction                    = "MissingAction"
	RuleValidationErrorCodeTypeMissingParameter                 = "MissingParameter"
	RuleValidationErrorCodeTypeMissingRangeValue                = "MissingRangeValue"
	RuleValidationErrorCodeTypeNotSettable                      = "NotSettable"
	RuleValidationErrorCodeTypeRecipientDoesNotExist            = "RecipientDoesNotExist"
	RuleValidationErrorCodeTypeRuleNotFound                     = "RuleNotFound"
	RuleValidationErrorCodeTypeSizeLessThanZero                 = "SizeLessThanZero"
	RuleValidationErrorCodeTypeStringValueTooBig                = "StringValueTooBig"
	RuleValidationErrorCodeTypeUnsupportedAddress               = "UnsupportedAddress"
	RuleValidationErrorCodeTypeUnexpectedError                  = "UnexpectedError"
	RuleValidationErrorCodeTypeUnsupportedRule                  = "UnsupportedRule"
)

type PreviewItemBaseShapeType XsString

const (
	PreviewItemBaseShapeTypeDefault = "Default"
	PreviewItemBaseShapeTypeCompact = "Compact"
)

type SearchPageDirectionType XsString

const (
	SearchPageDirectionTypePrevious = "Previous"
	SearchPageDirectionTypeNext     = "Next"
)

type HoldActionType XsString

const (
	HoldActionTypeCreate = "Create"
	HoldActionTypeUpdate = "Update"
	HoldActionTypeRemove = "Remove"
)

type ReportMessagePlatformType XsString

const (
	ReportMessagePlatformTypeUnknown      = "Unknown"
	ReportMessagePlatformTypeAndroid      = "Android"
	ReportMessagePlatformTypeIiOS         = "iOS"
	ReportMessagePlatformTypeMac          = "Mac"
	ReportMessagePlatformTypeOfficeOnline = "OfficeOnline"
	ReportMessagePlatformTypePC           = "PC"
	ReportMessagePlatformTypeUniversal    = "Universal"
)

type ReportMessageActionType XsString

const (
	ReportMessageActionTypeJunk        = "Junk"
	ReportMessageActionTypeNotJunk     = "NotJunk"
	ReportMessageActionTypePhish       = "Phish"
	ReportMessageActionTypeUnsubscribe = "Unsubscribe"
	ReportMessageActionTypeGetPolicy   = "GetPolicy"
)

type AddInStateType XsString

const (
	AddInStateTypeFlagged         = "Flagged"
	AddInStateTypeOK              = "OK"
	AddInStateTypeRemoved         = "Removed"
	AddInStateTypeUndefined       = "Undefined"
	AddInStateTypeWithdrawingSoon = "WithdrawingSoon"
	AddInStateTypeWithdrawn       = "Withdrawn"
)

type VersionType XsString

type AADOfficeExtensionStatusType XsString

const (
	AADOfficeExtensionStatusTypeOptionalDisabled = "OptionalDisabled"
	AADOfficeExtensionStatusTypeOptionalEnabled  = "OptionalEnabled"
	AADOfficeExtensionStatusTypeMandatory        = "Mandatory"
	AADOfficeExtensionStatusTypeUndefined        = "Undefined"
)

type DisableReasonType XsString

const (
	DisableReasonTypeNoReason                 = "NoReason"
	DisableReasonTypeOutlookClientPerformance = "OutlookClientPerformance"
	DisableReasonTypeOWAClientPerformance     = "OWAClientPerformance"
	DisableReasonTypeMobileClientPerformance  = "MobileClientPerformance"
)

type ActivityDomainType XsString

const (
	ActivityDomainTypeUnknown  = "Unknown"
	ActivityDomainTypePersonal = "Personal"
	ActivityDomainTypeWork     = "Work"
)

type AvailabilityStatusType XsString

const (
	AvailabilityStatusTypeUnknown          = "Unknown"
	AvailabilityStatusTypeFree             = "Free"
	AvailabilityStatusTypeTentative        = "Tentative"
	AvailabilityStatusTypeBusy             = "Busy"
	AvailabilityStatusTypeOof              = "Oof"
	AvailabilityStatusTypeWorkingElsewhere = "WorkingElsewhere"
)

type EmptySuggestionReason XsString

const (
	EmptySuggestionReasonUnknown                       = "Unknown"
	EmptySuggestionReasonAttendeesUnavailable          = "AttendeesUnavailable"
	EmptySuggestionReasonLocationsUnavailable          = "LocationsUnavailable"
	EmptySuggestionReasonOrganizerUnavailable          = "OrganizerUnavailable"
	EmptySuggestionReasonAttendeesUnavailableOrUnknown = "AttendeesUnavailableOrUnknown"
)

type UserPhotoSizeType XsString

const (
	UserPhotoSizeTypeHR48x48   = "HR48x48"
	UserPhotoSizeTypeHR64x64   = "HR64x64"
	UserPhotoSizeTypeHR96x96   = "HR96x96"
	UserPhotoSizeTypeHR120x120 = "HR120x120"
	UserPhotoSizeTypeHR240x240 = "HR240x240"
	UserPhotoSizeTypeHR360x360 = "HR360x360"
	UserPhotoSizeTypeHR432x432 = "HR432x432"
	UserPhotoSizeTypeHR504x504 = "HR504x504"
	UserPhotoSizeTypeHR648x648 = "HR648x648"
	UserPhotoSizeTypeHR1024xN  = "HR1024xN"
	UserPhotoSizeTypeHR1920xN  = "HR1920xN"
)

type UserPhotoTypeType XsString

const (
	UserPhotoTypeTypeUserPhoto          = "UserPhoto"
	UserPhotoTypeTypeProfileHeaderPhoto = "ProfileHeaderPhoto"
)

type MeetingSpaceTypeEnum XsString

const (
	MeetingSpaceTypeEnumPublic  = "Public"
	MeetingSpaceTypeEnumPrivate = "Private"
)

type ParticipantActivityRole XsString

const (
	ParticipantActivityRoleOrganizer = "Organizer"
	ParticipantActivityRoleAttendee  = "Attendee"
	ParticipantActivityRolePresenter = "Presenter"
)

type ParticipantActivityMediaType XsString

const (
	ParticipantActivityMediaTypeAppSharing  = "AppSharing"
	ParticipantActivityMediaTypeAudioVideo  = "AudioVideo"
	ParticipantActivityMediaTypeChat        = "Chat"
	ParticipantActivityMediaTypeDataConf    = "DataConf"
	ParticipantActivityMediaTypeMeeting     = "Meeting"
	ParticipantActivityMediaTypeMeetingConf = "MeetingConf"
	ParticipantActivityMediaTypePhoneConf   = "PhoneConf"
	ParticipantActivityMediaTypeFocus       = "Focus"
)

type ContentActivityType XsString

const (
	ContentActivityTypePoll        = "Poll"
	ContentActivityTypeWhiteBoard  = "WhiteBoard"
	ContentActivityTypeQAndA       = "QAndA"
	ContentActivityTypeChat        = "Chat"
	ContentActivityTypeMeeting     = "Meeting"
	ContentActivityTypeAnnotations = "Annotations"
	ContentActivityTypeSharedNotes = "SharedNotes"
)

type ContentActivityMediaType XsString

const (
	ContentActivityMediaTypeAppSharing  = "AppSharing"
	ContentActivityMediaTypeAudioVideo  = "AudioVideo"
	ContentActivityMediaTypeChat        = "Chat"
	ContentActivityMediaTypeDataConf    = "DataConf"
	ContentActivityMediaTypeMeeting     = "Meeting"
	ContentActivityMediaTypeMeetingConf = "MeetingConf"
	ContentActivityMediaTypePhoneConf   = "PhoneConf"
)

type ContentActivityAcl XsString

const (
	ContentActivityAclOrganizer = "Organizer"
	ContentActivityAclPresenter = "Presenter"
	ContentActivityAclEveryone  = "Everyone"
)

type WarmupOptionsType []WarmupOptionsTypeItem

type SuggestionKindType []SuggestionKindTypeItem

type SearchScopeArchivesType []SearchScopeArchivesTypeItem

type SearchScopeGroupsType []SearchScopeGroupsTypeItem

type OneDriveViewType []OneDriveViewTypeItem

type DelveViewType []DelveViewTypeItem

type SearchApplicationIdType XsString

const (
	SearchApplicationIdTypeOutlook   = "Outlook"
	SearchApplicationIdTypeOwa       = "Owa"
	SearchApplicationIdTypePaw       = "Paw"
	SearchApplicationIdTypeTeamspace = "Teamspace"
	SearchApplicationIdTypeOneDrive  = "OneDrive"
	SearchApplicationIdTypeOther     = "Other"
)

type ItemTypesFilterType []ItemTypesFilterTypeItem

type SearchResultsPropertySetNameType XsString

const (
	SearchResultsPropertySetNameTypeDefault   = "Default"
	SearchResultsPropertySetNameTypeOwa16     = "Owa16"
	SearchResultsPropertySetNameTypeOutlook16 = "Outlook16"
)

type ExecuteSearchSortOrderType XsString

const (
	ExecuteSearchSortOrderTypeDateTime  = "DateTime"
	ExecuteSearchSortOrderTypeRelevance = "Relevance"
)

type MatchOptionsType XsString

const (
	MatchOptionsTypeFullString = "FullString"
	MatchOptionsTypePrefix     = "Prefix"
)

type RefinerTypeType XsString

const (
	RefinerTypeTypeNone          = "None"
	RefinerTypeTypeTo            = "To"
	RefinerTypeTypeFrom          = "From"
	RefinerTypeTypeFolder        = "Folder"
	RefinerTypeTypeHasAttachment = "HasAttachment"
	RefinerTypeTypeMailboxSource = "MailboxSource"
)

type OfficeClientCodeType XsString

type ResponseCodeType XsString

const (
	ResponseCodeTypeNoError                                                  = "NoError"
	ResponseCodeTypeErrorAccessDenied                                        = "ErrorAccessDenied"
	ResponseCodeTypeErrorAccessModeSpecified                                 = "ErrorAccessModeSpecified"
	ResponseCodeTypeErrorAccountDisabled                                     = "ErrorAccountDisabled"
	ResponseCodeTypeErrorAddDelegatesFailed                                  = "ErrorAddDelegatesFailed"
	ResponseCodeTypeErrorAddressSpaceNotFound                                = "ErrorAddressSpaceNotFound"
	ResponseCodeTypeErrorADOperation                                         = "ErrorADOperation"
	ResponseCodeTypeErrorADSessionFilter                                     = "ErrorADSessionFilter"
	ResponseCodeTypeErrorADUnavailable                                       = "ErrorADUnavailable"
	ResponseCodeTypeErrorServiceUnavailable                                  = "ErrorServiceUnavailable"
	ResponseCodeTypeErrorAutoDiscoverFailed                                  = "ErrorAutoDiscoverFailed"
	ResponseCodeTypeErrorAffectedTaskOccurrencesRequired                     = "ErrorAffectedTaskOccurrencesRequired"
	ResponseCodeTypeErrorAttachmentNestLevelLimitExceeded                    = "ErrorAttachmentNestLevelLimitExceeded"
	ResponseCodeTypeErrorAttachmentSizeLimitExceeded                         = "ErrorAttachmentSizeLimitExceeded"
	ResponseCodeTypeErrorArchiveFolderPathCreation                           = "ErrorArchiveFolderPathCreation"
	ResponseCodeTypeErrorArchiveMailboxNotEnabled                            = "ErrorArchiveMailboxNotEnabled"
	ResponseCodeTypeErrorArchiveMailboxServiceDiscoveryFailed                = "ErrorArchiveMailboxServiceDiscoveryFailed"
	ResponseCodeTypeErrorAvailabilityConfigNotFound                          = "ErrorAvailabilityConfigNotFound"
	ResponseCodeTypeErrorBatchProcessingStopped                              = "ErrorBatchProcessingStopped"
	ResponseCodeTypeErrorCalendarCannotMoveOrCopyOccurrence                  = "ErrorCalendarCannotMoveOrCopyOccurrence"
	ResponseCodeTypeErrorCalendarCannotUpdateDeletedItem                     = "ErrorCalendarCannotUpdateDeletedItem"
	ResponseCodeTypeErrorCalendarCannotUseIdForOccurrenceId                  = "ErrorCalendarCannotUseIdForOccurrenceId"
	ResponseCodeTypeErrorCalendarCannotUseIdForRecurringMasterId             = "ErrorCalendarCannotUseIdForRecurringMasterId"
	ResponseCodeTypeErrorCalendarDurationIsTooLong                           = "ErrorCalendarDurationIsTooLong"
	ResponseCodeTypeErrorCalendarEndDateIsEarlierThanStartDate               = "ErrorCalendarEndDateIsEarlierThanStartDate"
	ResponseCodeTypeErrorCalendarFolderIsInvalidForCalendarView              = "ErrorCalendarFolderIsInvalidForCalendarView"
	ResponseCodeTypeErrorCalendarInvalidAttributeValue                       = "ErrorCalendarInvalidAttributeValue"
	ResponseCodeTypeErrorCalendarInvalidDayForTimeChangePattern              = "ErrorCalendarInvalidDayForTimeChangePattern"
	ResponseCodeTypeErrorCalendarInvalidDayForWeeklyRecurrence               = "ErrorCalendarInvalidDayForWeeklyRecurrence"
	ResponseCodeTypeErrorCalendarInvalidPropertyState                        = "ErrorCalendarInvalidPropertyState"
	ResponseCodeTypeErrorCalendarInvalidPropertyValue                        = "ErrorCalendarInvalidPropertyValue"
	ResponseCodeTypeErrorCalendarInvalidRecurrence                           = "ErrorCalendarInvalidRecurrence"
	ResponseCodeTypeErrorCalendarInvalidTimeZone                             = "ErrorCalendarInvalidTimeZone"
	ResponseCodeTypeErrorCalendarIsCancelledForAccept                        = "ErrorCalendarIsCancelledForAccept"
	ResponseCodeTypeErrorCalendarIsCancelledForDecline                       = "ErrorCalendarIsCancelledForDecline"
	ResponseCodeTypeErrorCalendarIsCancelledForRemove                        = "ErrorCalendarIsCancelledForRemove"
	ResponseCodeTypeErrorCalendarIsCancelledForTentative                     = "ErrorCalendarIsCancelledForTentative"
	ResponseCodeTypeErrorCalendarIsDelegatedForAccept                        = "ErrorCalendarIsDelegatedForAccept"
	ResponseCodeTypeErrorCalendarIsDelegatedForDecline                       = "ErrorCalendarIsDelegatedForDecline"
	ResponseCodeTypeErrorCalendarIsDelegatedForRemove                        = "ErrorCalendarIsDelegatedForRemove"
	ResponseCodeTypeErrorCalendarIsDelegatedForTentative                     = "ErrorCalendarIsDelegatedForTentative"
	ResponseCodeTypeErrorCalendarIsNotOrganizer                              = "ErrorCalendarIsNotOrganizer"
	ResponseCodeTypeErrorCalendarIsOrganizerForAccept                        = "ErrorCalendarIsOrganizerForAccept"
	ResponseCodeTypeErrorCalendarIsOrganizerForDecline                       = "ErrorCalendarIsOrganizerForDecline"
	ResponseCodeTypeErrorCalendarIsOrganizerForRemove                        = "ErrorCalendarIsOrganizerForRemove"
	ResponseCodeTypeErrorCalendarIsOrganizerForTentative                     = "ErrorCalendarIsOrganizerForTentative"
	ResponseCodeTypeErrorCalendarOccurrenceIndexIsOutOfRecurrenceRange       = "ErrorCalendarOccurrenceIndexIsOutOfRecurrenceRange"
	ResponseCodeTypeErrorCalendarOccurrenceIsDeletedFromRecurrence           = "ErrorCalendarOccurrenceIsDeletedFromRecurrence"
	ResponseCodeTypeErrorCalendarOutOfRange                                  = "ErrorCalendarOutOfRange"
	ResponseCodeTypeErrorCalendarMeetingRequestIsOutOfDate                   = "ErrorCalendarMeetingRequestIsOutOfDate"
	ResponseCodeTypeErrorCalendarViewRangeTooBig                             = "ErrorCalendarViewRangeTooBig"
	ResponseCodeTypeErrorCallerIsInvalidADAccount                            = "ErrorCallerIsInvalidADAccount"
	ResponseCodeTypeErrorCannotAccessDeletedPublicFolder                     = "ErrorCannotAccessDeletedPublicFolder"
	ResponseCodeTypeErrorCannotArchiveCalendarContactTaskFolderException     = "ErrorCannotArchiveCalendarContactTaskFolderException"
	ResponseCodeTypeErrorCannotArchiveItemsInPublicFolders                   = "ErrorCannotArchiveItemsInPublicFolders"
	ResponseCodeTypeErrorCannotArchiveItemsInArchiveMailbox                  = "ErrorCannotArchiveItemsInArchiveMailbox"
	ResponseCodeTypeErrorCannotCreateCalendarItemInNonCalendarFolder         = "ErrorCannotCreateCalendarItemInNonCalendarFolder"
	ResponseCodeTypeErrorCannotCreateContactInNonContactFolder               = "ErrorCannotCreateContactInNonContactFolder"
	ResponseCodeTypeErrorCannotCreatePostItemInNonMailFolder                 = "ErrorCannotCreatePostItemInNonMailFolder"
	ResponseCodeTypeErrorCannotCreateTaskInNonTaskFolder                     = "ErrorCannotCreateTaskInNonTaskFolder"
	ResponseCodeTypeErrorCannotDeleteObject                                  = "ErrorCannotDeleteObject"
	ResponseCodeTypeErrorCannotDisableMandatoryExtension                     = "ErrorCannotDisableMandatoryExtension"
	ResponseCodeTypeErrorCannotFindUser                                      = "ErrorCannotFindUser"
	ResponseCodeTypeErrorCannotGetSourceFolderPath                           = "ErrorCannotGetSourceFolderPath"
	ResponseCodeTypeErrorCannotGetExternalEcpUrl                             = "ErrorCannotGetExternalEcpUrl"
	ResponseCodeTypeErrorCannotOpenFileAttachment                            = "ErrorCannotOpenFileAttachment"
	ResponseCodeTypeErrorCannotDeleteTaskOccurrence                          = "ErrorCannotDeleteTaskOccurrence"
	ResponseCodeTypeErrorCannotEmptyFolder                                   = "ErrorCannotEmptyFolder"
	ResponseCodeTypeErrorCannotSetCalendarPermissionOnNonCalendarFolder      = "ErrorCannotSetCalendarPermissionOnNonCalendarFolder"
	ResponseCodeTypeErrorCannotSetNonCalendarPermissionOnCalendarFolder      = "ErrorCannotSetNonCalendarPermissionOnCalendarFolder"
	ResponseCodeTypeErrorCannotSetPermissionUnknownEntries                   = "ErrorCannotSetPermissionUnknownEntries"
	ResponseCodeTypeErrorCannotSpecifySearchFolderAsSourceFolder             = "ErrorCannotSpecifySearchFolderAsSourceFolder"
	ResponseCodeTypeErrorCannotUseFolderIdForItemId                          = "ErrorCannotUseFolderIdForItemId"
	ResponseCodeTypeErrorCannotUseItemIdForFolderId                          = "ErrorCannotUseItemIdForFolderId"
	ResponseCodeTypeErrorChangeKeyRequired                                   = "ErrorChangeKeyRequired"
	ResponseCodeTypeErrorChangeKeyRequiredForWriteOperations                 = "ErrorChangeKeyRequiredForWriteOperations"
	ResponseCodeTypeErrorClientDisconnected                                  = "ErrorClientDisconnected"
	ResponseCodeTypeErrorClientIntentInvalidStateDefinition                  = "ErrorClientIntentInvalidStateDefinition"
	ResponseCodeTypeErrorClientIntentNotFound                                = "ErrorClientIntentNotFound"
	ResponseCodeTypeErrorConnectionFailed                                    = "ErrorConnectionFailed"
	ResponseCodeTypeErrorContainsFilterWrongType                             = "ErrorContainsFilterWrongType"
	ResponseCodeTypeErrorContentConversionFailed                             = "ErrorContentConversionFailed"
	ResponseCodeTypeErrorContentIndexingNotEnabled                           = "ErrorContentIndexingNotEnabled"
	ResponseCodeTypeErrorCorruptData                                         = "ErrorCorruptData"
	ResponseCodeTypeErrorCreateItemAccessDenied                              = "ErrorCreateItemAccessDenied"
	ResponseCodeTypeErrorCreateManagedFolderPartialCompletion                = "ErrorCreateManagedFolderPartialCompletion"
	ResponseCodeTypeErrorCreateSubfolderAccessDenied                         = "ErrorCreateSubfolderAccessDenied"
	ResponseCodeTypeErrorCrossMailboxMoveCopy                                = "ErrorCrossMailboxMoveCopy"
	ResponseCodeTypeErrorCrossSiteRequest                                    = "ErrorCrossSiteRequest"
	ResponseCodeTypeErrorDataSizeLimitExceeded                               = "ErrorDataSizeLimitExceeded"
	ResponseCodeTypeErrorDataSourceOperation                                 = "ErrorDataSourceOperation"
	ResponseCodeTypeErrorDelegateAlreadyExists                               = "ErrorDelegateAlreadyExists"
	ResponseCodeTypeErrorDelegateCannotAddOwner                              = "ErrorDelegateCannotAddOwner"
	ResponseCodeTypeErrorDelegateMissingConfiguration                        = "ErrorDelegateMissingConfiguration"
	ResponseCodeTypeErrorDelegateNoUser                                      = "ErrorDelegateNoUser"
	ResponseCodeTypeErrorDelegateValidationFailed                            = "ErrorDelegateValidationFailed"
	ResponseCodeTypeErrorDeleteDistinguishedFolder                           = "ErrorDeleteDistinguishedFolder"
	ResponseCodeTypeErrorDeleteItemsFailed                                   = "ErrorDeleteItemsFailed"
	ResponseCodeTypeErrorDeleteUnifiedMessagingPromptFailed                  = "ErrorDeleteUnifiedMessagingPromptFailed"
	ResponseCodeTypeErrorDistinguishedUserNotSupported                       = "ErrorDistinguishedUserNotSupported"
	ResponseCodeTypeErrorDistributionListMemberNotExist                      = "ErrorDistributionListMemberNotExist"
	ResponseCodeTypeErrorDuplicateInputFolderNames                           = "ErrorDuplicateInputFolderNames"
	ResponseCodeTypeErrorDuplicateUserIdsSpecified                           = "ErrorDuplicateUserIdsSpecified"
	ResponseCodeTypeErrorDuplicateTransactionId                              = "ErrorDuplicateTransactionId"
	ResponseCodeTypeErrorEmailAddressMismatch                                = "ErrorEmailAddressMismatch"
	ResponseCodeTypeErrorEventNotFound                                       = "ErrorEventNotFound"
	ResponseCodeTypeErrorExceededConnectionCount                             = "ErrorExceededConnectionCount"
	ResponseCodeTypeErrorExceededSubscriptionCount                           = "ErrorExceededSubscriptionCount"
	ResponseCodeTypeErrorExceededFindCountLimit                              = "ErrorExceededFindCountLimit"
	ResponseCodeTypeErrorExpiredSubscription                                 = "ErrorExpiredSubscription"
	ResponseCodeTypeErrorExtensionNotFound                                   = "ErrorExtensionNotFound"
	ResponseCodeTypeErrorExtensionsNotAuthorized                             = "ErrorExtensionsNotAuthorized"
	ResponseCodeTypeErrorFolderCorrupt                                       = "ErrorFolderCorrupt"
	ResponseCodeTypeErrorFolderNotFound                                      = "ErrorFolderNotFound"
	ResponseCodeTypeErrorFolderPropertRequestFailed                          = "ErrorFolderPropertRequestFailed"
	ResponseCodeTypeErrorFolderSave                                          = "ErrorFolderSave"
	ResponseCodeTypeErrorFolderSaveFailed                                    = "ErrorFolderSaveFailed"
	ResponseCodeTypeErrorFolderSavePropertyError                             = "ErrorFolderSavePropertyError"
	ResponseCodeTypeErrorFolderExists                                        = "ErrorFolderExists"
	ResponseCodeTypeErrorFreeBusyGenerationFailed                            = "ErrorFreeBusyGenerationFailed"
	ResponseCodeTypeErrorGetServerSecurityDescriptorFailed                   = "ErrorGetServerSecurityDescriptorFailed"
	ResponseCodeTypeErrorImContactLimitReached                               = "ErrorImContactLimitReached"
	ResponseCodeTypeErrorImGroupDisplayNameAlreadyExists                     = "ErrorImGroupDisplayNameAlreadyExists"
	ResponseCodeTypeErrorImGroupLimitReached                                 = "ErrorImGroupLimitReached"
	ResponseCodeTypeErrorImpersonateUserDenied                               = "ErrorImpersonateUserDenied"
	ResponseCodeTypeErrorImpersonationDenied                                 = "ErrorImpersonationDenied"
	ResponseCodeTypeErrorImpersonationFailed                                 = "ErrorImpersonationFailed"
	ResponseCodeTypeErrorIncorrectSchemaVersion                              = "ErrorIncorrectSchemaVersion"
	ResponseCodeTypeErrorIncorrectUpdatePropertyCount                        = "ErrorIncorrectUpdatePropertyCount"
	ResponseCodeTypeErrorIndividualMailboxLimitReached                       = "ErrorIndividualMailboxLimitReached"
	ResponseCodeTypeErrorInsufficientResources                               = "ErrorInsufficientResources"
	ResponseCodeTypeErrorInternalServerError                                 = "ErrorInternalServerError"
	ResponseCodeTypeErrorInternalServerTransientError                        = "ErrorInternalServerTransientError"
	ResponseCodeTypeErrorInvalidAccessLevel                                  = "ErrorInvalidAccessLevel"
	ResponseCodeTypeErrorInvalidArgument                                     = "ErrorInvalidArgument"
	ResponseCodeTypeErrorInvalidAttachmentId                                 = "ErrorInvalidAttachmentId"
	ResponseCodeTypeErrorInvalidAttachmentSubfilter                          = "ErrorInvalidAttachmentSubfilter"
	ResponseCodeTypeErrorInvalidAttachmentSubfilterTextFilter                = "ErrorInvalidAttachmentSubfilterTextFilter"
	ResponseCodeTypeErrorInvalidAuthorizationContext                         = "ErrorInvalidAuthorizationContext"
	ResponseCodeTypeErrorInvalidChangeKey                                    = "ErrorInvalidChangeKey"
	ResponseCodeTypeErrorInvalidClientSecurityContext                        = "ErrorInvalidClientSecurityContext"
	ResponseCodeTypeErrorInvalidCompleteDate                                 = "ErrorInvalidCompleteDate"
	ResponseCodeTypeErrorInvalidContactEmailAddress                          = "ErrorInvalidContactEmailAddress"
	ResponseCodeTypeErrorInvalidContactEmailIndex                            = "ErrorInvalidContactEmailIndex"
	ResponseCodeTypeErrorInvalidCrossForestCredentials                       = "ErrorInvalidCrossForestCredentials"
	ResponseCodeTypeErrorInvalidDelegatePermission                           = "ErrorInvalidDelegatePermission"
	ResponseCodeTypeErrorInvalidDelegateUserId                               = "ErrorInvalidDelegateUserId"
	ResponseCodeTypeErrorInvalidExcludesRestriction                          = "ErrorInvalidExcludesRestriction"
	ResponseCodeTypeErrorInvalidExpressionTypeForSubFilter                   = "ErrorInvalidExpressionTypeForSubFilter"
	ResponseCodeTypeErrorInvalidExtendedProperty                             = "ErrorInvalidExtendedProperty"
	ResponseCodeTypeErrorInvalidExtendedPropertyValue                        = "ErrorInvalidExtendedPropertyValue"
	ResponseCodeTypeErrorInvalidFolderId                                     = "ErrorInvalidFolderId"
	ResponseCodeTypeErrorInvalidFolderTypeForOperation                       = "ErrorInvalidFolderTypeForOperation"
	ResponseCodeTypeErrorInvalidFractionalPagingParameters                   = "ErrorInvalidFractionalPagingParameters"
	ResponseCodeTypeErrorInvalidFreeBusyViewType                             = "ErrorInvalidFreeBusyViewType"
	ResponseCodeTypeErrorInvalidId                                           = "ErrorInvalidId"
	ResponseCodeTypeErrorInvalidIdEmpty                                      = "ErrorInvalidIdEmpty"
	ResponseCodeTypeErrorInvalidIdMalformed                                  = "ErrorInvalidIdMalformed"
	ResponseCodeTypeErrorInvalidIdMalformedEwsLegacyIdFormat                 = "ErrorInvalidIdMalformedEwsLegacyIdFormat"
	ResponseCodeTypeErrorInvalidIdMonikerTooLong                             = "ErrorInvalidIdMonikerTooLong"
	ResponseCodeTypeErrorInvalidIdNotAnItemAttachmentId                      = "ErrorInvalidIdNotAnItemAttachmentId"
	ResponseCodeTypeErrorInvalidIdReturnedByResolveNames                     = "ErrorInvalidIdReturnedByResolveNames"
	ResponseCodeTypeErrorInvalidIdStoreObjectIdTooLong                       = "ErrorInvalidIdStoreObjectIdTooLong"
	ResponseCodeTypeErrorInvalidIdTooManyAttachmentLevels                    = "ErrorInvalidIdTooManyAttachmentLevels"
	ResponseCodeTypeErrorInvalidIdXml                                        = "ErrorInvalidIdXml"
	ResponseCodeTypeErrorInvalidImContactId                                  = "ErrorInvalidImContactId"
	ResponseCodeTypeErrorInvalidImDistributionGroupSmtpAddress               = "ErrorInvalidImDistributionGroupSmtpAddress"
	ResponseCodeTypeErrorInvalidImGroupId                                    = "ErrorInvalidImGroupId"
	ResponseCodeTypeErrorInvalidIndexedPagingParameters                      = "ErrorInvalidIndexedPagingParameters"
	ResponseCodeTypeErrorInvalidInternetHeaderChildNodes                     = "ErrorInvalidInternetHeaderChildNodes"
	ResponseCodeTypeErrorInvalidItemForOperationArchiveItem                  = "ErrorInvalidItemForOperationArchiveItem"
	ResponseCodeTypeErrorInvalidItemForOperationCreateItemAttachment         = "ErrorInvalidItemForOperationCreateItemAttachment"
	ResponseCodeTypeErrorInvalidItemForOperationCreateItem                   = "ErrorInvalidItemForOperationCreateItem"
	ResponseCodeTypeErrorInvalidItemForOperationAcceptItem                   = "ErrorInvalidItemForOperationAcceptItem"
	ResponseCodeTypeErrorInvalidItemForOperationDeclineItem                  = "ErrorInvalidItemForOperationDeclineItem"
	ResponseCodeTypeErrorInvalidItemForOperationCancelItem                   = "ErrorInvalidItemForOperationCancelItem"
	ResponseCodeTypeErrorInvalidItemForOperationExpandDL                     = "ErrorInvalidItemForOperationExpandDL"
	ResponseCodeTypeErrorInvalidItemForOperationRemoveItem                   = "ErrorInvalidItemForOperationRemoveItem"
	ResponseCodeTypeErrorInvalidItemForOperationSendItem                     = "ErrorInvalidItemForOperationSendItem"
	ResponseCodeTypeErrorInvalidItemForOperationTentative                    = "ErrorInvalidItemForOperationTentative"
	ResponseCodeTypeErrorInvalidLogonType                                    = "ErrorInvalidLogonType"
	ResponseCodeTypeErrorInvalidLikeRequest                                  = "ErrorInvalidLikeRequest"
	ResponseCodeTypeErrorInvalidMailbox                                      = "ErrorInvalidMailbox"
	ResponseCodeTypeErrorInvalidManagedFolderProperty                        = "ErrorInvalidManagedFolderProperty"
	ResponseCodeTypeErrorInvalidManagedFolderQuota                           = "ErrorInvalidManagedFolderQuota"
	ResponseCodeTypeErrorInvalidManagedFolderSize                            = "ErrorInvalidManagedFolderSize"
	ResponseCodeTypeErrorInvalidMergedFreeBusyInterval                       = "ErrorInvalidMergedFreeBusyInterval"
	ResponseCodeTypeErrorInvalidNameForNameResolution                        = "ErrorInvalidNameForNameResolution"
	ResponseCodeTypeErrorInvalidOperation                                    = "ErrorInvalidOperation"
	ResponseCodeTypeErrorInvalidNetworkServiceContext                        = "ErrorInvalidNetworkServiceContext"
	ResponseCodeTypeErrorInvalidOofParameter                                 = "ErrorInvalidOofParameter"
	ResponseCodeTypeErrorInvalidPagingMaxRows                                = "ErrorInvalidPagingMaxRows"
	ResponseCodeTypeErrorInvalidParentFolder                                 = "ErrorInvalidParentFolder"
	ResponseCodeTypeErrorInvalidPercentCompleteValue                         = "ErrorInvalidPercentCompleteValue"
	ResponseCodeTypeErrorInvalidPermissionSettings                           = "ErrorInvalidPermissionSettings"
	ResponseCodeTypeErrorInvalidPhoneCallId                                  = "ErrorInvalidPhoneCallId"
	ResponseCodeTypeErrorInvalidPhoneNumber                                  = "ErrorInvalidPhoneNumber"
	ResponseCodeTypeErrorInvalidUserInfo                                     = "ErrorInvalidUserInfo"
	ResponseCodeTypeErrorInvalidPropertyAppend                               = "ErrorInvalidPropertyAppend"
	ResponseCodeTypeErrorInvalidPropertyDelete                               = "ErrorInvalidPropertyDelete"
	ResponseCodeTypeErrorInvalidPropertyForExists                            = "ErrorInvalidPropertyForExists"
	ResponseCodeTypeErrorInvalidPropertyForOperation                         = "ErrorInvalidPropertyForOperation"
	ResponseCodeTypeErrorInvalidPropertyRequest                              = "ErrorInvalidPropertyRequest"
	ResponseCodeTypeErrorInvalidPropertySet                                  = "ErrorInvalidPropertySet"
	ResponseCodeTypeErrorInvalidPropertyUpdateSentMessage                    = "ErrorInvalidPropertyUpdateSentMessage"
	ResponseCodeTypeErrorInvalidProxySecurityContext                         = "ErrorInvalidProxySecurityContext"
	ResponseCodeTypeErrorInvalidPullSubscriptionId                           = "ErrorInvalidPullSubscriptionId"
	ResponseCodeTypeErrorInvalidPushSubscriptionUrl                          = "ErrorInvalidPushSubscriptionUrl"
	ResponseCodeTypeErrorInvalidRecipients                                   = "ErrorInvalidRecipients"
	ResponseCodeTypeErrorInvalidRecipientSubfilter                           = "ErrorInvalidRecipientSubfilter"
	ResponseCodeTypeErrorInvalidRecipientSubfilterComparison                 = "ErrorInvalidRecipientSubfilterComparison"
	ResponseCodeTypeErrorInvalidRecipientSubfilterOrder                      = "ErrorInvalidRecipientSubfilterOrder"
	ResponseCodeTypeErrorInvalidRecipientSubfilterTextFilter                 = "ErrorInvalidRecipientSubfilterTextFilter"
	ResponseCodeTypeErrorInvalidReferenceItem                                = "ErrorInvalidReferenceItem"
	ResponseCodeTypeErrorInvalidRequest                                      = "ErrorInvalidRequest"
	ResponseCodeTypeErrorInvalidRestriction                                  = "ErrorInvalidRestriction"
	ResponseCodeTypeErrorInvalidRetentionTagTypeMismatch                     = "ErrorInvalidRetentionTagTypeMismatch"
	ResponseCodeTypeErrorInvalidRetentionTagInvisible                        = "ErrorInvalidRetentionTagInvisible"
	ResponseCodeTypeErrorInvalidRetentionTagInheritance                      = "ErrorInvalidRetentionTagInheritance"
	ResponseCodeTypeErrorInvalidRetentionTagIdGuid                           = "ErrorInvalidRetentionTagIdGuid"
	ResponseCodeTypeErrorInvalidRoutingType                                  = "ErrorInvalidRoutingType"
	ResponseCodeTypeErrorInvalidScheduledOofDuration                         = "ErrorInvalidScheduledOofDuration"
	ResponseCodeTypeErrorInvalidSchemaVersionForMailboxVersion               = "ErrorInvalidSchemaVersionForMailboxVersion"
	ResponseCodeTypeErrorInvalidSecurityDescriptor                           = "ErrorInvalidSecurityDescriptor"
	ResponseCodeTypeErrorInvalidSendItemSaveSettings                         = "ErrorInvalidSendItemSaveSettings"
	ResponseCodeTypeErrorInvalidSerializedAccessToken                        = "ErrorInvalidSerializedAccessToken"
	ResponseCodeTypeErrorInvalidServerVersion                                = "ErrorInvalidServerVersion"
	ResponseCodeTypeErrorInvalidSid                                          = "ErrorInvalidSid"
	ResponseCodeTypeErrorInvalidSIPUri                                       = "ErrorInvalidSIPUri"
	ResponseCodeTypeErrorInvalidSmtpAddress                                  = "ErrorInvalidSmtpAddress"
	ResponseCodeTypeErrorInvalidSubfilterType                                = "ErrorInvalidSubfilterType"
	ResponseCodeTypeErrorInvalidSubfilterTypeNotAttendeeType                 = "ErrorInvalidSubfilterTypeNotAttendeeType"
	ResponseCodeTypeErrorInvalidSubfilterTypeNotRecipientType                = "ErrorInvalidSubfilterTypeNotRecipientType"
	ResponseCodeTypeErrorInvalidSubscription                                 = "ErrorInvalidSubscription"
	ResponseCodeTypeErrorInvalidSubscriptionRequest                          = "ErrorInvalidSubscriptionRequest"
	ResponseCodeTypeErrorInvalidSyncStateData                                = "ErrorInvalidSyncStateData"
	ResponseCodeTypeErrorInvalidTimeInterval                                 = "ErrorInvalidTimeInterval"
	ResponseCodeTypeErrorInvalidUserOofSettings                              = "ErrorInvalidUserOofSettings"
	ResponseCodeTypeErrorInvalidUserPrincipalName                            = "ErrorInvalidUserPrincipalName"
	ResponseCodeTypeErrorInvalidUserSid                                      = "ErrorInvalidUserSid"
	ResponseCodeTypeErrorInvalidUserSidMissingUPN                            = "ErrorInvalidUserSidMissingUPN"
	ResponseCodeTypeErrorInvalidValueForProperty                             = "ErrorInvalidValueForProperty"
	ResponseCodeTypeErrorInvalidWatermark                                    = "ErrorInvalidWatermark"
	ResponseCodeTypeErrorIPGatewayNotFound                                   = "ErrorIPGatewayNotFound"
	ResponseCodeTypeErrorIrresolvableConflict                                = "ErrorIrresolvableConflict"
	ResponseCodeTypeErrorItemCorrupt                                         = "ErrorItemCorrupt"
	ResponseCodeTypeErrorItemNotFound                                        = "ErrorItemNotFound"
	ResponseCodeTypeErrorItemPropertyRequestFailed                           = "ErrorItemPropertyRequestFailed"
	ResponseCodeTypeErrorItemSave                                            = "ErrorItemSave"
	ResponseCodeTypeErrorItemSavePropertyError                               = "ErrorItemSavePropertyError"
	ResponseCodeTypeErrorLegacyMailboxFreeBusyViewTypeNotMerged              = "ErrorLegacyMailboxFreeBusyViewTypeNotMerged"
	ResponseCodeTypeErrorLocalServerObjectNotFound                           = "ErrorLocalServerObjectNotFound"
	ResponseCodeTypeErrorLogonAsNetworkServiceFailed                         = "ErrorLogonAsNetworkServiceFailed"
	ResponseCodeTypeErrorMailboxConfiguration                                = "ErrorMailboxConfiguration"
	ResponseCodeTypeErrorMailboxDataArrayEmpty                               = "ErrorMailboxDataArrayEmpty"
	ResponseCodeTypeErrorMailboxDataArrayTooBig                              = "ErrorMailboxDataArrayTooBig"
	ResponseCodeTypeErrorMailboxHoldNotFound                                 = "ErrorMailboxHoldNotFound"
	ResponseCodeTypeErrorMailboxLogonFailed                                  = "ErrorMailboxLogonFailed"
	ResponseCodeTypeErrorMailboxMoveInProgress                               = "ErrorMailboxMoveInProgress"
	ResponseCodeTypeErrorMailboxStoreUnavailable                             = "ErrorMailboxStoreUnavailable"
	ResponseCodeTypeErrorMailRecipientNotFound                               = "ErrorMailRecipientNotFound"
	ResponseCodeTypeErrorMailTipsDisabled                                    = "ErrorMailTipsDisabled"
	ResponseCodeTypeErrorManagedFolderAlreadyExists                          = "ErrorManagedFolderAlreadyExists"
	ResponseCodeTypeErrorManagedFolderNotFound                               = "ErrorManagedFolderNotFound"
	ResponseCodeTypeErrorManagedFoldersRootFailure                           = "ErrorManagedFoldersRootFailure"
	ResponseCodeTypeErrorMeetingSuggestionGenerationFailed                   = "ErrorMeetingSuggestionGenerationFailed"
	ResponseCodeTypeErrorMessageDispositionRequired                          = "ErrorMessageDispositionRequired"
	ResponseCodeTypeErrorMessageSizeExceeded                                 = "ErrorMessageSizeExceeded"
	ResponseCodeTypeErrorMimeContentConversionFailed                         = "ErrorMimeContentConversionFailed"
	ResponseCodeTypeErrorMimeContentInvalid                                  = "ErrorMimeContentInvalid"
	ResponseCodeTypeErrorMimeContentInvalidBase64String                      = "ErrorMimeContentInvalidBase64String"
	ResponseCodeTypeErrorMissingArgument                                     = "ErrorMissingArgument"
	ResponseCodeTypeErrorMissingEmailAddress                                 = "ErrorMissingEmailAddress"
	ResponseCodeTypeErrorMissingEmailAddressForManagedFolder                 = "ErrorMissingEmailAddressForManagedFolder"
	ResponseCodeTypeErrorMissingInformationEmailAddress                      = "ErrorMissingInformationEmailAddress"
	ResponseCodeTypeErrorMissingInformationReferenceItemId                   = "ErrorMissingInformationReferenceItemId"
	ResponseCodeTypeErrorMissingItemForCreateItemAttachment                  = "ErrorMissingItemForCreateItemAttachment"
	ResponseCodeTypeErrorMissingManagedFolderId                              = "ErrorMissingManagedFolderId"
	ResponseCodeTypeErrorMissingRecipients                                   = "ErrorMissingRecipients"
	ResponseCodeTypeErrorMissingUserIdInformation                            = "ErrorMissingUserIdInformation"
	ResponseCodeTypeErrorMoreThanOneAccessModeSpecified                      = "ErrorMoreThanOneAccessModeSpecified"
	ResponseCodeTypeErrorMoveCopyFailed                                      = "ErrorMoveCopyFailed"
	ResponseCodeTypeErrorMoveDistinguishedFolder                             = "ErrorMoveDistinguishedFolder"
	ResponseCodeTypeErrorMoveUnifiedGroupPropertyFailed                      = "ErrorMoveUnifiedGroupPropertyFailed"
	ResponseCodeTypeErrorMultiLegacyMailboxAccess                            = "ErrorMultiLegacyMailboxAccess"
	ResponseCodeTypeErrorNameResolutionMultipleResults                       = "ErrorNameResolutionMultipleResults"
	ResponseCodeTypeErrorNameResolutionNoMailbox                             = "ErrorNameResolutionNoMailbox"
	ResponseCodeTypeErrorNameResolutionNoResults                             = "ErrorNameResolutionNoResults"
	ResponseCodeTypeErrorNoApplicableProxyCASServersAvailable                = "ErrorNoApplicableProxyCASServersAvailable"
	ResponseCodeTypeErrorNoCalendar                                          = "ErrorNoCalendar"
	ResponseCodeTypeErrorNoDestinationCASDueToKerberosRequirements           = "ErrorNoDestinationCASDueToKerberosRequirements"
	ResponseCodeTypeErrorNoDestinationCASDueToSSLRequirements                = "ErrorNoDestinationCASDueToSSLRequirements"
	ResponseCodeTypeErrorNoDestinationCASDueToVersionMismatch                = "ErrorNoDestinationCASDueToVersionMismatch"
	ResponseCodeTypeErrorNoFolderClassOverride                               = "ErrorNoFolderClassOverride"
	ResponseCodeTypeErrorNoFreeBusyAccess                                    = "ErrorNoFreeBusyAccess"
	ResponseCodeTypeErrorNonExistentMailbox                                  = "ErrorNonExistentMailbox"
	ResponseCodeTypeErrorNonPrimarySmtpAddress                               = "ErrorNonPrimarySmtpAddress"
	ResponseCodeTypeErrorNoPropertyTagForCustomProperties                    = "ErrorNoPropertyTagForCustomProperties"
	ResponseCodeTypeErrorNoPublicFolderReplicaAvailable                      = "ErrorNoPublicFolderReplicaAvailable"
	ResponseCodeTypeErrorNoPublicFolderServerAvailable                       = "ErrorNoPublicFolderServerAvailable"
	ResponseCodeTypeErrorNoRespondingCASInDestinationSite                    = "ErrorNoRespondingCASInDestinationSite"
	ResponseCodeTypeErrorNotDelegate                                         = "ErrorNotDelegate"
	ResponseCodeTypeErrorNotEnoughMemory                                     = "ErrorNotEnoughMemory"
	ResponseCodeTypeErrorObjectTypeChanged                                   = "ErrorObjectTypeChanged"
	ResponseCodeTypeErrorOccurrenceCrossingBoundary                          = "ErrorOccurrenceCrossingBoundary"
	ResponseCodeTypeErrorOccurrenceTimeSpanTooBig                            = "ErrorOccurrenceTimeSpanTooBig"
	ResponseCodeTypeErrorOperationNotAllowedWithPublicFolderRoot             = "ErrorOperationNotAllowedWithPublicFolderRoot"
	ResponseCodeTypeErrorParentFolderIdRequired                              = "ErrorParentFolderIdRequired"
	ResponseCodeTypeErrorParentFolderNotFound                                = "ErrorParentFolderNotFound"
	ResponseCodeTypeErrorPasswordChangeRequired                              = "ErrorPasswordChangeRequired"
	ResponseCodeTypeErrorPasswordExpired                                     = "ErrorPasswordExpired"
	ResponseCodeTypeErrorPhoneNumberNotDialable                              = "ErrorPhoneNumberNotDialable"
	ResponseCodeTypeErrorPropertyUpdate                                      = "ErrorPropertyUpdate"
	ResponseCodeTypeErrorPromptPublishingOperationFailed                     = "ErrorPromptPublishingOperationFailed"
	ResponseCodeTypeErrorPropertyValidationFailure                           = "ErrorPropertyValidationFailure"
	ResponseCodeTypeErrorProxiedSubscriptionCallFailure                      = "ErrorProxiedSubscriptionCallFailure"
	ResponseCodeTypeErrorProxyCallFailed                                     = "ErrorProxyCallFailed"
	ResponseCodeTypeErrorProxyGroupSidLimitExceeded                          = "ErrorProxyGroupSidLimitExceeded"
	ResponseCodeTypeErrorProxyRequestNotAllowed                              = "ErrorProxyRequestNotAllowed"
	ResponseCodeTypeErrorProxyRequestProcessingFailed                        = "ErrorProxyRequestProcessingFailed"
	ResponseCodeTypeErrorProxyServiceDiscoveryFailed                         = "ErrorProxyServiceDiscoveryFailed"
	ResponseCodeTypeErrorProxyTokenExpired                                   = "ErrorProxyTokenExpired"
	ResponseCodeTypeErrorPublicFolderMailboxDiscoveryFailed                  = "ErrorPublicFolderMailboxDiscoveryFailed"
	ResponseCodeTypeErrorPublicFolderOperationFailed                         = "ErrorPublicFolderOperationFailed"
	ResponseCodeTypeErrorPublicFolderRequestProcessingFailed                 = "ErrorPublicFolderRequestProcessingFailed"
	ResponseCodeTypeErrorPublicFolderServerNotFound                          = "ErrorPublicFolderServerNotFound"
	ResponseCodeTypeErrorPublicFolderSyncException                           = "ErrorPublicFolderSyncException"
	ResponseCodeTypeErrorQueryFilterTooLong                                  = "ErrorQueryFilterTooLong"
	ResponseCodeTypeErrorQuotaExceeded                                       = "ErrorQuotaExceeded"
	ResponseCodeTypeErrorReadEventsFailed                                    = "ErrorReadEventsFailed"
	ResponseCodeTypeErrorReadReceiptNotPending                               = "ErrorReadReceiptNotPending"
	ResponseCodeTypeErrorRecurrenceEndDateTooBig                             = "ErrorRecurrenceEndDateTooBig"
	ResponseCodeTypeErrorRecurrenceHasNoOccurrence                           = "ErrorRecurrenceHasNoOccurrence"
	ResponseCodeTypeErrorRemoveDelegatesFailed                               = "ErrorRemoveDelegatesFailed"
	ResponseCodeTypeErrorRequestAborted                                      = "ErrorRequestAborted"
	ResponseCodeTypeErrorRequestStreamTooBig                                 = "ErrorRequestStreamTooBig"
	ResponseCodeTypeErrorRequiredPropertyMissing                             = "ErrorRequiredPropertyMissing"
	ResponseCodeTypeErrorResolveNamesInvalidFolderType                       = "ErrorResolveNamesInvalidFolderType"
	ResponseCodeTypeErrorResolveNamesOnlyOneContactsFolderAllowed            = "ErrorResolveNamesOnlyOneContactsFolderAllowed"
	ResponseCodeTypeErrorResponseSchemaValidation                            = "ErrorResponseSchemaValidation"
	ResponseCodeTypeErrorRestrictionTooLong                                  = "ErrorRestrictionTooLong"
	ResponseCodeTypeErrorRestrictionTooComplex                               = "ErrorRestrictionTooComplex"
	ResponseCodeTypeErrorResultSetTooBig                                     = "ErrorResultSetTooBig"
	ResponseCodeTypeErrorInvalidExchangeImpersonationHeaderData              = "ErrorInvalidExchangeImpersonationHeaderData"
	ResponseCodeTypeErrorSavedItemFolderNotFound                             = "ErrorSavedItemFolderNotFound"
	ResponseCodeTypeErrorSchemaValidation                                    = "ErrorSchemaValidation"
	ResponseCodeTypeErrorSearchFolderNotInitialized                          = "ErrorSearchFolderNotInitialized"
	ResponseCodeTypeErrorSendAsDenied                                        = "ErrorSendAsDenied"
	ResponseCodeTypeErrorSendMeetingCancellationsRequired                    = "ErrorSendMeetingCancellationsRequired"
	ResponseCodeTypeErrorSendMeetingInvitationsOrCancellationsRequired       = "ErrorSendMeetingInvitationsOrCancellationsRequired"
	ResponseCodeTypeErrorSendMeetingInvitationsRequired                      = "ErrorSendMeetingInvitationsRequired"
	ResponseCodeTypeErrorSentMeetingRequestUpdate                            = "ErrorSentMeetingRequestUpdate"
	ResponseCodeTypeErrorSentTaskRequestUpdate                               = "ErrorSentTaskRequestUpdate"
	ResponseCodeTypeErrorServerBusy                                          = "ErrorServerBusy"
	ResponseCodeTypeErrorServiceDiscoveryFailed                              = "ErrorServiceDiscoveryFailed"
	ResponseCodeTypeErrorStaleObject                                         = "ErrorStaleObject"
	ResponseCodeTypeErrorSubmissionQuotaExceeded                             = "ErrorSubmissionQuotaExceeded"
	ResponseCodeTypeErrorSubscriptionAccessDenied                            = "ErrorSubscriptionAccessDenied"
	ResponseCodeTypeErrorSubscriptionDelegateAccessNotSupported              = "ErrorSubscriptionDelegateAccessNotSupported"
	ResponseCodeTypeErrorSubscriptionNotFound                                = "ErrorSubscriptionNotFound"
	ResponseCodeTypeErrorSubscriptionUnsubscribed                            = "ErrorSubscriptionUnsubscribed"
	ResponseCodeTypeErrorSyncFolderNotFound                                  = "ErrorSyncFolderNotFound"
	ResponseCodeTypeErrorTeamMailboxNotFound                                 = "ErrorTeamMailboxNotFound"
	ResponseCodeTypeErrorTeamMailboxNotLinkedToSharePoint                    = "ErrorTeamMailboxNotLinkedToSharePoint"
	ResponseCodeTypeErrorTeamMailboxUrlValidationFailed                      = "ErrorTeamMailboxUrlValidationFailed"
	ResponseCodeTypeErrorTeamMailboxNotAuthorizedOwner                       = "ErrorTeamMailboxNotAuthorizedOwner"
	ResponseCodeTypeErrorTeamMailboxActiveToPendingDelete                    = "ErrorTeamMailboxActiveToPendingDelete"
	ResponseCodeTypeErrorTeamMailboxFailedSendingNotifications               = "ErrorTeamMailboxFailedSendingNotifications"
	ResponseCodeTypeErrorTeamMailboxErrorUnknown                             = "ErrorTeamMailboxErrorUnknown"
	ResponseCodeTypeErrorTimeIntervalTooBig                                  = "ErrorTimeIntervalTooBig"
	ResponseCodeTypeErrorTimeoutExpired                                      = "ErrorTimeoutExpired"
	ResponseCodeTypeErrorTimeZone                                            = "ErrorTimeZone"
	ResponseCodeTypeErrorToFolderNotFound                                    = "ErrorToFolderNotFound"
	ResponseCodeTypeErrorTokenSerializationDenied                            = "ErrorTokenSerializationDenied"
	ResponseCodeTypeErrorTooManyObjectsOpened                                = "ErrorTooManyObjectsOpened"
	ResponseCodeTypeErrorUpdatePropertyMismatch                              = "ErrorUpdatePropertyMismatch"
	ResponseCodeTypeErrorAccessingPartialCreatedUnifiedGroup                 = "ErrorAccessingPartialCreatedUnifiedGroup"
	ResponseCodeTypeErrorUnifiedGroupMailboxAADCreationFailed                = "ErrorUnifiedGroupMailboxAADCreationFailed"
	ResponseCodeTypeErrorUnifiedGroupMailboxAADDeleteFailed                  = "ErrorUnifiedGroupMailboxAADDeleteFailed"
	ResponseCodeTypeErrorUnifiedGroupMailboxNamingPolicy                     = "ErrorUnifiedGroupMailboxNamingPolicy"
	ResponseCodeTypeErrorUnifiedGroupMailboxDeleteFailed                     = "ErrorUnifiedGroupMailboxDeleteFailed"
	ResponseCodeTypeErrorUnifiedGroupMailboxNotFound                         = "ErrorUnifiedGroupMailboxNotFound"
	ResponseCodeTypeErrorUnifiedGroupMailboxUpdateDelayed                    = "ErrorUnifiedGroupMailboxUpdateDelayed"
	ResponseCodeTypeErrorUnifiedGroupMailboxUpdatedPartialProperties         = "ErrorUnifiedGroupMailboxUpdatedPartialProperties"
	ResponseCodeTypeErrorUnifiedGroupMailboxUpdateFailed                     = "ErrorUnifiedGroupMailboxUpdateFailed"
	ResponseCodeTypeErrorUnifiedGroupMailboxProvisionFailed                  = "ErrorUnifiedGroupMailboxProvisionFailed"
	ResponseCodeTypeErrorUnifiedMessagingDialPlanNotFound                    = "ErrorUnifiedMessagingDialPlanNotFound"
	ResponseCodeTypeErrorUnifiedMessagingReportDataNotFound                  = "ErrorUnifiedMessagingReportDataNotFound"
	ResponseCodeTypeErrorUnifiedMessagingPromptNotFound                      = "ErrorUnifiedMessagingPromptNotFound"
	ResponseCodeTypeErrorUnifiedMessagingRequestFailed                       = "ErrorUnifiedMessagingRequestFailed"
	ResponseCodeTypeErrorUnifiedMessagingServerNotFound                      = "ErrorUnifiedMessagingServerNotFound"
	ResponseCodeTypeErrorUnableToGetUserOofSettings                          = "ErrorUnableToGetUserOofSettings"
	ResponseCodeTypeErrorUnableToRemoveImContactFromGroup                    = "ErrorUnableToRemoveImContactFromGroup"
	ResponseCodeTypeErrorUnsupportedSubFilter                                = "ErrorUnsupportedSubFilter"
	ResponseCodeTypeErrorUnsupportedCulture                                  = "ErrorUnsupportedCulture"
	ResponseCodeTypeErrorUnsupportedMapiPropertyType                         = "ErrorUnsupportedMapiPropertyType"
	ResponseCodeTypeErrorUnsupportedMimeConversion                           = "ErrorUnsupportedMimeConversion"
	ResponseCodeTypeErrorUnsupportedPathForQuery                             = "ErrorUnsupportedPathForQuery"
	ResponseCodeTypeErrorUnsupportedPathForSortGroup                         = "ErrorUnsupportedPathForSortGroup"
	ResponseCodeTypeErrorUnsupportedPropertyDefinition                       = "ErrorUnsupportedPropertyDefinition"
	ResponseCodeTypeErrorUnsupportedQueryFilter                              = "ErrorUnsupportedQueryFilter"
	ResponseCodeTypeErrorUnsupportedRecurrence                               = "ErrorUnsupportedRecurrence"
	ResponseCodeTypeErrorUnsupportedTypeForConversion                        = "ErrorUnsupportedTypeForConversion"
	ResponseCodeTypeErrorUpdateDelegatesFailed                               = "ErrorUpdateDelegatesFailed"
	ResponseCodeTypeErrorUserNotUnifiedMessagingEnabled                      = "ErrorUserNotUnifiedMessagingEnabled"
	ResponseCodeTypeErrorVoiceMailNotImplemented                             = "ErrorVoiceMailNotImplemented"
	ResponseCodeTypeErrorValueOutOfRange                                     = "ErrorValueOutOfRange"
	ResponseCodeTypeErrorVirusDetected                                       = "ErrorVirusDetected"
	ResponseCodeTypeErrorVirusMessageDeleted                                 = "ErrorVirusMessageDeleted"
	ResponseCodeTypeErrorWebRequestInInvalidState                            = "ErrorWebRequestInInvalidState"
	ResponseCodeTypeErrorWin32InteropError                                   = "ErrorWin32InteropError"
	ResponseCodeTypeErrorWorkingHoursSaveFailed                              = "ErrorWorkingHoursSaveFailed"
	ResponseCodeTypeErrorWorkingHoursXmlMalformed                            = "ErrorWorkingHoursXmlMalformed"
	ResponseCodeTypeErrorWrongServerVersion                                  = "ErrorWrongServerVersion"
	ResponseCodeTypeErrorWrongServerVersionDelegate                          = "ErrorWrongServerVersionDelegate"
	ResponseCodeTypeErrorMissingInformationSharingFolderId                   = "ErrorMissingInformationSharingFolderId"
	ResponseCodeTypeErrorDuplicateSOAPHeader                                 = "ErrorDuplicateSOAPHeader"
	ResponseCodeTypeErrorSharingSynchronizationFailed                        = "ErrorSharingSynchronizationFailed"
	ResponseCodeTypeErrorSharingNoExternalEwsAvailable                       = "ErrorSharingNoExternalEwsAvailable"
	ResponseCodeTypeErrorFreeBusyDLLimitReached                              = "ErrorFreeBusyDLLimitReached"
	ResponseCodeTypeErrorInvalidGetSharingFolderRequest                      = "ErrorInvalidGetSharingFolderRequest"
	ResponseCodeTypeErrorNotAllowedExternalSharingByPolicy                   = "ErrorNotAllowedExternalSharingByPolicy"
	ResponseCodeTypeErrorUserNotAllowedByPolicy                              = "ErrorUserNotAllowedByPolicy"
	ResponseCodeTypeErrorPermissionNotAllowedByPolicy                        = "ErrorPermissionNotAllowedByPolicy"
	ResponseCodeTypeErrorOrganizationNotFederated                            = "ErrorOrganizationNotFederated"
	ResponseCodeTypeErrorMailboxFailover                                     = "ErrorMailboxFailover"
	ResponseCodeTypeErrorInvalidExternalSharingInitiator                     = "ErrorInvalidExternalSharingInitiator"
	ResponseCodeTypeErrorMessageTrackingPermanentError                       = "ErrorMessageTrackingPermanentError"
	ResponseCodeTypeErrorMessageTrackingTransientError                       = "ErrorMessageTrackingTransientError"
	ResponseCodeTypeErrorMessageTrackingNoSuchDomain                         = "ErrorMessageTrackingNoSuchDomain"
	ResponseCodeTypeErrorUserWithoutFederatedProxyAddress                    = "ErrorUserWithoutFederatedProxyAddress"
	ResponseCodeTypeErrorInvalidOrganizationRelationshipForFreeBusy          = "ErrorInvalidOrganizationRelationshipForFreeBusy"
	ResponseCodeTypeErrorInvalidFederatedOrganizationId                      = "ErrorInvalidFederatedOrganizationId"
	ResponseCodeTypeErrorInvalidExternalSharingSubscriber                    = "ErrorInvalidExternalSharingSubscriber"
	ResponseCodeTypeErrorInvalidSharingData                                  = "ErrorInvalidSharingData"
	ResponseCodeTypeErrorInvalidSharingMessage                               = "ErrorInvalidSharingMessage"
	ResponseCodeTypeErrorNotSupportedSharingMessage                          = "ErrorNotSupportedSharingMessage"
	ResponseCodeTypeErrorApplyConversationActionFailed                       = "ErrorApplyConversationActionFailed"
	ResponseCodeTypeErrorInboxRulesValidationError                           = "ErrorInboxRulesValidationError"
	ResponseCodeTypeErrorOutlookRuleBlobExists                               = "ErrorOutlookRuleBlobExists"
	ResponseCodeTypeErrorRulesOverQuota                                      = "ErrorRulesOverQuota"
	ResponseCodeTypeErrorNewEventStreamConnectionOpened                      = "ErrorNewEventStreamConnectionOpened"
	ResponseCodeTypeErrorMissedNotificationEvents                            = "ErrorMissedNotificationEvents"
	ResponseCodeTypeErrorDuplicateLegacyDistinguishedName                    = "ErrorDuplicateLegacyDistinguishedName"
	ResponseCodeTypeErrorInvalidClientAccessTokenRequest                     = "ErrorInvalidClientAccessTokenRequest"
	ResponseCodeTypeErrorUnauthorizedClientAccessTokenRequest                = "ErrorUnauthorizedClientAccessTokenRequest"
	ResponseCodeTypeErrorNoSpeechDetected                                    = "ErrorNoSpeechDetected"
	ResponseCodeTypeErrorUMServerUnavailable                                 = "ErrorUMServerUnavailable"
	ResponseCodeTypeErrorRecipientNotFound                                   = "ErrorRecipientNotFound"
	ResponseCodeTypeErrorRecognizerNotInstalled                              = "ErrorRecognizerNotInstalled"
	ResponseCodeTypeErrorSpeechGrammarError                                  = "ErrorSpeechGrammarError"
	ResponseCodeTypeErrorInvalidManagementRoleHeader                         = "ErrorInvalidManagementRoleHeader"
	ResponseCodeTypeErrorLocationServicesDisabled                            = "ErrorLocationServicesDisabled"
	ResponseCodeTypeErrorLocationServicesRequestTimedOut                     = "ErrorLocationServicesRequestTimedOut"
	ResponseCodeTypeErrorLocationServicesRequestFailed                       = "ErrorLocationServicesRequestFailed"
	ResponseCodeTypeErrorLocationServicesInvalidRequest                      = "ErrorLocationServicesInvalidRequest"
	ResponseCodeTypeErrorWeatherServiceDisabled                              = "ErrorWeatherServiceDisabled"
	ResponseCodeTypeErrorMailboxScopeNotAllowedWithoutQueryString            = "ErrorMailboxScopeNotAllowedWithoutQueryString"
	ResponseCodeTypeErrorArchiveMailboxSearchFailed                          = "ErrorArchiveMailboxSearchFailed"
	ResponseCodeTypeErrorGetRemoteArchiveFolderFailed                        = "ErrorGetRemoteArchiveFolderFailed"
	ResponseCodeTypeErrorFindRemoteArchiveFolderFailed                       = "ErrorFindRemoteArchiveFolderFailed"
	ResponseCodeTypeErrorGetRemoteArchiveItemFailed                          = "ErrorGetRemoteArchiveItemFailed"
	ResponseCodeTypeErrorExportRemoteArchiveItemsFailed                      = "ErrorExportRemoteArchiveItemsFailed"
	ResponseCodeTypeErrorInvalidPhotoSize                                    = "ErrorInvalidPhotoSize"
	ResponseCodeTypeErrorSearchQueryHasTooManyKeywords                       = "ErrorSearchQueryHasTooManyKeywords"
	ResponseCodeTypeErrorSearchTooManyMailboxes                              = "ErrorSearchTooManyMailboxes"
	ResponseCodeTypeErrorInvalidRetentionTagNone                             = "ErrorInvalidRetentionTagNone"
	ResponseCodeTypeErrorDiscoverySearchesDisabled                           = "ErrorDiscoverySearchesDisabled"
	ResponseCodeTypeErrorCalendarSeekToConditionNotSupported                 = "ErrorCalendarSeekToConditionNotSupported"
	ResponseCodeTypeErrorCalendarIsGroupMailboxForAccept                     = "ErrorCalendarIsGroupMailboxForAccept"
	ResponseCodeTypeErrorCalendarIsGroupMailboxForDecline                    = "ErrorCalendarIsGroupMailboxForDecline"
	ResponseCodeTypeErrorCalendarIsGroupMailboxForTentative                  = "ErrorCalendarIsGroupMailboxForTentative"
	ResponseCodeTypeErrorCalendarIsGroupMailboxForSuppressReadReceipt        = "ErrorCalendarIsGroupMailboxForSuppressReadReceipt"
	ResponseCodeTypeErrorOrganizationAccessBlocked                           = "ErrorOrganizationAccessBlocked"
	ResponseCodeTypeErrorInvalidLicense                                      = "ErrorInvalidLicense"
	ResponseCodeTypeErrorMessagePerFolderCountReceiveQuotaExceeded           = "ErrorMessagePerFolderCountReceiveQuotaExceeded"
	ResponseCodeTypeErrorInvalidBulkActionType                               = "ErrorInvalidBulkActionType"
	ResponseCodeTypeErrorInvalidKeepNCount                                   = "ErrorInvalidKeepNCount"
	ResponseCodeTypeErrorInvalidKeepNType                                    = "ErrorInvalidKeepNType"
	ResponseCodeTypeErrorNoOAuthServerAvailableForRequest                    = "ErrorNoOAuthServerAvailableForRequest"
	ResponseCodeTypeErrorInstantSearchSessionExpired                         = "ErrorInstantSearchSessionExpired"
	ResponseCodeTypeErrorInstantSearchTimeout                                = "ErrorInstantSearchTimeout"
	ResponseCodeTypeErrorInstantSearchFailed                                 = "ErrorInstantSearchFailed"
	ResponseCodeTypeErrorUnsupportedUserForExecuteSearch                     = "ErrorUnsupportedUserForExecuteSearch"
	ResponseCodeTypeErrorDuplicateExtendedKeywordDefinition                  = "ErrorDuplicateExtendedKeywordDefinition"
	ResponseCodeTypeErrorMissingExchangePrincipal                            = "ErrorMissingExchangePrincipal"
	ResponseCodeTypeErrorUnexpectedUnifiedGroupsCount                        = "ErrorUnexpectedUnifiedGroupsCount"
	ResponseCodeTypeErrorParsingXMLResponse                                  = "ErrorParsingXMLResponse"
	ResponseCodeTypeErrorInvalidFederationOrganizationIdentifier             = "ErrorInvalidFederationOrganizationIdentifier"
	ResponseCodeTypeErrorInvalidSweepRule                                    = "ErrorInvalidSweepRule"
	ResponseCodeTypeErrorInvalidSweepRuleOperationType                       = "ErrorInvalidSweepRuleOperationType"
	ResponseCodeTypeErrorTargetDomainNotSupported                            = "ErrorTargetDomainNotSupported"
	ResponseCodeTypeErrorInvalidInternetWebProxyOnLocalServer                = "ErrorInvalidInternetWebProxyOnLocalServer"
	ResponseCodeTypeErrorNoSenderRestrictionsSettingsFoundInRequest          = "ErrorNoSenderRestrictionsSettingsFoundInRequest"
	ResponseCodeTypeErrorDuplicateSenderRestrictionsInputFound               = "ErrorDuplicateSenderRestrictionsInputFound"
	ResponseCodeTypeErrorSenderRestrictionsUpdateFailed                      = "ErrorSenderRestrictionsUpdateFailed"
	ResponseCodeTypeErrorMessageSubmissionBlocked                            = "ErrorMessageSubmissionBlocked"
	ResponseCodeTypeErrorExceededMessageLimit                                = "ErrorExceededMessageLimit"
	ResponseCodeTypeErrorExceededMaxRecipientLimitBlock                      = "ErrorExceededMaxRecipientLimitBlock"
	ResponseCodeTypeErrorAccountSuspend                                      = "ErrorAccountSuspend"
	ResponseCodeTypeErrorExceededMaxRecipientLimit                           = "ErrorExceededMaxRecipientLimit"
	ResponseCodeTypeErrorMessageBlocked                                      = "ErrorMessageBlocked"
	ResponseCodeTypeErrorAccountSuspendShowTierUpgrade                       = "ErrorAccountSuspendShowTierUpgrade"
	ResponseCodeTypeErrorExceededMessageLimitShowTierUpgrade                 = "ErrorExceededMessageLimitShowTierUpgrade"
	ResponseCodeTypeErrorExceededMaxRecipientLimitShowTierUpgrade            = "ErrorExceededMaxRecipientLimitShowTierUpgrade"
	ResponseCodeTypeErrorInvalidLongitude                                    = "ErrorInvalidLongitude"
	ResponseCodeTypeErrorInvalidLatitude                                     = "ErrorInvalidLatitude"
	ResponseCodeTypeErrorProxySoapException                                  = "ErrorProxySoapException"
	ResponseCodeTypeErrorUnifiedGroupAlreadyExists                           = "ErrorUnifiedGroupAlreadyExists"
	ResponseCodeTypeErrorUnifiedGroupAadAuthorizationRequestDenied           = "ErrorUnifiedGroupAadAuthorizationRequestDenied"
	ResponseCodeTypeErrorUnifiedGroupCreationDisabled                        = "ErrorUnifiedGroupCreationDisabled"
	ResponseCodeTypeErrorMarketPlaceExtensionAlreadyInstalledForOrg          = "ErrorMarketPlaceExtensionAlreadyInstalledForOrg"
	ResponseCodeTypeErrorExtensionAlreadyInstalledForOrg                     = "ErrorExtensionAlreadyInstalledForOrg"
	ResponseCodeTypeErrorNewerExtensionAlreadyInstalled                      = "ErrorNewerExtensionAlreadyInstalled"
	ResponseCodeTypeErrorNewerMarketPlaceExtensionAlreadyInstalled           = "ErrorNewerMarketPlaceExtensionAlreadyInstalled"
	ResponseCodeTypeErrorInvalidExtensionId                                  = "ErrorInvalidExtensionId"
	ResponseCodeTypeErrorCannotUninstallProvidedExtensions                   = "ErrorCannotUninstallProvidedExtensions"
	ResponseCodeTypeErrorNoRbacPermissionToInstallMarketPlaceExtensions      = "ErrorNoRbacPermissionToInstallMarketPlaceExtensions"
	ResponseCodeTypeErrorNoRbacPermissionToInstallReadWriteMailboxExtensions = "ErrorNoRbacPermissionToInstallReadWriteMailboxExtensions"
	ResponseCodeTypeErrorInvalidReportMessageActionType                      = "ErrorInvalidReportMessageActionType"
	ResponseCodeTypeErrorCannotDownloadExtensionManifest                     = "ErrorCannotDownloadExtensionManifest"
	ResponseCodeTypeErrorCalendarForwardActionNotAllowed                     = "ErrorCalendarForwardActionNotAllowed"
	ResponseCodeTypeErrorUnifiedGroupAliasNamingPolicy                       = "ErrorUnifiedGroupAliasNamingPolicy"
	ResponseCodeTypeErrorSubscriptionsDisabledForGroup                       = "ErrorSubscriptionsDisabledForGroup"
	ResponseCodeTypeErrorCannotFindFileAttachment                            = "ErrorCannotFindFileAttachment"
	ResponseCodeTypeErrorInvalidValueForFilter                               = "ErrorInvalidValueForFilter"
	ResponseCodeTypeErrorQuotaExceededOnDelete                               = "ErrorQuotaExceededOnDelete"
	ResponseCodeTypeErrorAccessDeniedDueToCompliance                         = "ErrorAccessDeniedDueToCompliance"
	ResponseCodeTypeErrorRecoverableItemsAccessDenied                        = "ErrorRecoverableItemsAccessDenied"
)

type ListOfExtensionIdsType []GuidType

type RequestServerVersionType struct {
	Version ExchangeVersionType `xml:"Version,attr"`
}

type MailboxGuidsType struct {
	MailboxGuid []GuidType `xml:"t:MailboxGuid,omitempty"`
}

type ServerVersionInfoType struct {
	MajorVersion     XsInt    `xml:"MajorVersion,attr"`
	MinorVersion     XsInt    `xml:"MinorVersion,attr"`
	MajorBuildNumber XsInt    `xml:"MajorBuildNumber,attr"`
	MinorBuildNumber XsInt    `xml:"MinorBuildNumber,attr"`
	Version          XsString `xml:"Version,attr"`
}

type SuggestionsType struct {
	Suggestion []*SuggestionType `xml:"t:Suggestion,omitempty"`
}

type ItemsType struct {
	Item []*ItemType `xml:"t:Item,omitempty"`
}

type ConversationsType struct {
	Conversation []*ConversationType `xml:"t:Conversation,omitempty"`
}

type PeopleType struct {
	Persona []*PersonaType `xml:"t:Persona,omitempty"`
}

type SearchRefinersTypeT struct {
	SearchRefiner []*SearchRefinerType `xml:"t:SearchRefiner,omitempty"`
}

type MailboxesInformationType struct {
	MailboxInformation []*MailboxInformationType `xml:"t:MailboxInformation,omitempty"`
}

type MessageXmlType struct {
}

type SearchRefinersTypeM struct {
	SearchRefiner []*DynamicRefinerQueryType `xml:"m:SearchRefiner,omitempty"`
}

type ExtendedKeywordsType struct {
	ExtendedKeywordDefinition []*ExtendedKeywordDefinitionType `xml:"m:ExtendedKeywordDefinition,omitempty"`
}

type ResolveNamesType struct {
	BaseRequestType       `xml:",omitempty"`
	ReturnFullContactData XsBoolean                         `xml:"ReturnFullContactData,attr"`
	SearchScope           ResolveNamesSearchScopeType       `xml:"SearchScope,attr"`
	ContactDataShape      DefaultShapeNamesType             `xml:"ContactDataShape,attr"`
	ParentFolderIds       *NonEmptyArrayOfBaseFolderIdsType `xml:"m:ParentFolderIds,omitempty"`
	UnresolvedEntry       NonEmptyStringType                `xml:"m:UnresolvedEntry,omitempty"`
}

type BaseRequestType struct {
}

type ResolveNamesResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type BaseResponseMessageType struct {
	ResponseMessages *ArrayOfResponseMessagesType `xml:"m:ResponseMessages,omitempty"`
}

type ArrayOfResponseMessagesType struct {
	CreateItemResponseMessage                      []*ItemInfoResponseMessageType                        `xml:"m:CreateItemResponseMessage,omitempty"`
	DeleteItemResponseMessage                      []*DeleteItemResponseMessageType                      `xml:"m:DeleteItemResponseMessage,omitempty"`
	GetItemResponseMessage                         []*ItemInfoResponseMessageType                        `xml:"m:GetItemResponseMessage,omitempty"`
	UpdateItemResponseMessage                      []*UpdateItemResponseMessageType                      `xml:"m:UpdateItemResponseMessage,omitempty"`
	UpdateItemInRecoverableItemsResponseMessage    []*UpdateItemInRecoverableItemsResponseMessageType    `xml:"m:UpdateItemInRecoverableItemsResponseMessage,omitempty"`
	SendItemResponseMessage                        []*ResponseMessageType                                `xml:"m:SendItemResponseMessage,omitempty"`
	DeleteFolderResponseMessage                    []*ResponseMessageType                                `xml:"m:DeleteFolderResponseMessage,omitempty"`
	EmptyFolderResponseMessage                     []*ResponseMessageType                                `xml:"m:EmptyFolderResponseMessage,omitempty"`
	CreateFolderResponseMessage                    []*FolderInfoResponseMessageType                      `xml:"m:CreateFolderResponseMessage,omitempty"`
	GetFolderResponseMessage                       []*FolderInfoResponseMessageType                      `xml:"m:GetFolderResponseMessage,omitempty"`
	FindFolderResponseMessage                      []*FindFolderResponseMessageType                      `xml:"m:FindFolderResponseMessage,omitempty"`
	UpdateFolderResponseMessage                    []*FolderInfoResponseMessageType                      `xml:"m:UpdateFolderResponseMessage,omitempty"`
	MoveFolderResponseMessage                      []*FolderInfoResponseMessageType                      `xml:"m:MoveFolderResponseMessage,omitempty"`
	CopyFolderResponseMessage                      []*FolderInfoResponseMessageType                      `xml:"m:CopyFolderResponseMessage,omitempty"`
	CreateFolderPathResponseMessage                []*FolderInfoResponseMessageType                      `xml:"m:CreateFolderPathResponseMessage,omitempty"`
	CreateAttachmentResponseMessage                []*AttachmentInfoResponseMessageType                  `xml:"m:CreateAttachmentResponseMessage,omitempty"`
	DeleteAttachmentResponseMessage                []*DeleteAttachmentResponseMessageType                `xml:"m:DeleteAttachmentResponseMessage,omitempty"`
	GetAttachmentResponseMessage                   []*AttachmentInfoResponseMessageType                  `xml:"m:GetAttachmentResponseMessage,omitempty"`
	UploadItemsResponseMessage                     []*UploadItemsResponseMessageType                     `xml:"m:UploadItemsResponseMessage,omitempty"`
	ExportItemsResponseMessage                     []*ExportItemsResponseMessageType                     `xml:"m:ExportItemsResponseMessage,omitempty"`
	MarkAllItemsAsReadResponseMessage              []*ResponseMessageType                                `xml:"m:MarkAllItemsAsReadResponseMessage,omitempty"`
	GetClientAccessTokenResponseMessage            []*GetClientAccessTokenResponseMessageType            `xml:"m:GetClientAccessTokenResponseMessage,omitempty"`
	GetAppManifestsResponseMessage                 []*ResponseMessageType                                `xml:"m:GetAppManifestsResponseMessage,omitempty"`
	SetClientExtensionResponseMessage              []*ResponseMessageType                                `xml:"m:SetClientExtensionResponseMessage,omitempty"`
	GetOMEConfigurationResponseMessage             []*ResponseMessageType                                `xml:"m:GetOMEConfigurationResponseMessage,omitempty"`
	SetOMEConfigurationResponseMessage             []*ResponseMessageType                                `xml:"m:SetOMEConfigurationResponseMessage,omitempty"`
	GetOMEMessageStatusResponseType                []*ResponseMessageType                                `xml:"m:GetOMEMessageStatusResponseType,omitempty"`
	SetOMEMessageStatusResponseType                []*ResponseMessageType                                `xml:"m:SetOMEMessageStatusResponseType,omitempty"`
	FindItemResponseMessage                        []*FindItemResponseMessageType                        `xml:"m:FindItemResponseMessage,omitempty"`
	MoveItemResponseMessage                        []*ItemInfoResponseMessageType                        `xml:"m:MoveItemResponseMessage,omitempty"`
	ArchiveItemResponseMessage                     []*ItemInfoResponseMessageType                        `xml:"m:ArchiveItemResponseMessage,omitempty"`
	CopyItemResponseMessage                        []*ItemInfoResponseMessageType                        `xml:"m:CopyItemResponseMessage,omitempty"`
	ResolveNamesResponseMessage                    []*ResolveNamesResponseMessageType                    `xml:"m:ResolveNamesResponseMessage,omitempty"`
	ExpandDLResponseMessage                        []*ExpandDLResponseMessageType                        `xml:"m:ExpandDLResponseMessage,omitempty"`
	GetServerTimeZonesResponseMessage              []*GetServerTimeZonesResponseMessageType              `xml:"m:GetServerTimeZonesResponseMessage,omitempty"`
	GetEventsResponseMessage                       []*GetEventsResponseMessageType                       `xml:"m:GetEventsResponseMessage,omitempty"`
	GetStreamingEventsResponseMessage              []*GetStreamingEventsResponseMessageType              `xml:"m:GetStreamingEventsResponseMessage,omitempty"`
	SubscribeResponseMessage                       []*SubscribeResponseMessageType                       `xml:"m:SubscribeResponseMessage,omitempty"`
	UnsubscribeResponseMessage                     []*ResponseMessageType                                `xml:"m:UnsubscribeResponseMessage,omitempty"`
	SendNotificationResponseMessage                []*SendNotificationResponseMessageType                `xml:"m:SendNotificationResponseMessage,omitempty"`
	SyncFolderHierarchyResponseMessage             []*SyncFolderHierarchyResponseMessageType             `xml:"m:SyncFolderHierarchyResponseMessage,omitempty"`
	SyncFolderItemsResponseMessage                 []*SyncFolderItemsResponseMessageType                 `xml:"m:SyncFolderItemsResponseMessage,omitempty"`
	CreateManagedFolderResponseMessage             []*FolderInfoResponseMessageType                      `xml:"m:CreateManagedFolderResponseMessage,omitempty"`
	ConvertIdResponseMessage                       []*ConvertIdResponseMessageType                       `xml:"m:ConvertIdResponseMessage,omitempty"`
	GetSharingMetadataResponseMessage              []*GetSharingMetadataResponseMessageType              `xml:"m:GetSharingMetadataResponseMessage,omitempty"`
	RefreshSharingFolderResponseMessage            []*RefreshSharingFolderResponseMessageType            `xml:"m:RefreshSharingFolderResponseMessage,omitempty"`
	GetSharingFolderResponseMessage                []*GetSharingFolderResponseMessageType                `xml:"m:GetSharingFolderResponseMessage,omitempty"`
	CreateUserConfigurationResponseMessage         []*ResponseMessageType                                `xml:"m:CreateUserConfigurationResponseMessage,omitempty"`
	DeleteUserConfigurationResponseMessage         []*ResponseMessageType                                `xml:"m:DeleteUserConfigurationResponseMessage,omitempty"`
	GetUserConfigurationResponseMessage            []*GetUserConfigurationResponseMessageType            `xml:"m:GetUserConfigurationResponseMessage,omitempty"`
	GetSpecificUserConfigurationResponseMessage    []*GetSpecificUserConfigurationResponseMessageType    `xml:"m:GetSpecificUserConfigurationResponseMessage,omitempty"`
	UpdateUserConfigurationResponseMessage         []*ResponseMessageType                                `xml:"m:UpdateUserConfigurationResponseMessage,omitempty"`
	GetRoomListsResponse                           []*GetRoomListsResponseMessageType                    `xml:"m:GetRoomListsResponse,omitempty"`
	GetRoomsResponse                               []*GetRoomsResponseMessageType                        `xml:"m:GetRoomsResponse,omitempty"`
	GetRemindersResponse                           []*GetRemindersResponseMessageType                    `xml:"m:GetRemindersResponse,omitempty"`
	PerformReminderActionResponse                  []*PerformReminderActionResponseMessageType           `xml:"m:PerformReminderActionResponse,omitempty"`
	ApplyConversationActionResponseMessage         []*ApplyConversationActionResponseMessageType         `xml:"m:ApplyConversationActionResponseMessage,omitempty"`
	FindMailboxStatisticsByKeywordsResponseMessage []*FindMailboxStatisticsByKeywordsResponseMessageType `xml:"m:FindMailboxStatisticsByKeywordsResponseMessage,omitempty"`
	GetSearchableMailboxesResponseMessage          []*GetSearchableMailboxesResponseMessageType          `xml:"m:GetSearchableMailboxesResponseMessage,omitempty"`
	SearchMailboxesResponseMessage                 []*SearchMailboxesResponseMessageType                 `xml:"m:SearchMailboxesResponseMessage,omitempty"`
	GetDiscoverySearchConfigurationResponseMessage []*GetDiscoverySearchConfigurationResponseMessageType `xml:"m:GetDiscoverySearchConfigurationResponseMessage,omitempty"`
	GetHoldOnMailboxesResponseMessage              []*GetHoldOnMailboxesResponseMessageType              `xml:"m:GetHoldOnMailboxesResponseMessage,omitempty"`
	SetHoldOnMailboxesResponseMessage              []*SetHoldOnMailboxesResponseMessageType              `xml:"m:SetHoldOnMailboxesResponseMessage,omitempty"`
	GetNonIndexableItemStatisticsResponseMessage   []*GetNonIndexableItemStatisticsResponseMessageType   `xml:"m:GetNonIndexableItemStatisticsResponseMessage,omitempty"`
	GetNonIndexableItemDetailsResponseMessage      []*GetNonIndexableItemDetailsResponseMessageType      `xml:"m:GetNonIndexableItemDetailsResponseMessage,omitempty"`
	FindPeopleResponseMessage                      []*FindPeopleResponseMessageType                      `xml:"m:FindPeopleResponseMessage,omitempty"`
	FindTagsResponseMessage                        []*FindTagsResponseMessageType                        `xml:"m:FindTagsResponseMessage,omitempty"`
	AddTagResponseMessage                          []*AddTagResponseMessageType                          `xml:"m:AddTagResponseMessage,omitempty"`
	HideTagResponseMessage                         []*HideTagResponseMessageType                         `xml:"m:HideTagResponseMessage,omitempty"`
	GetPasswordExpirationDateResponse              []*GetPasswordExpirationDateResponseMessageType       `xml:"m:GetPasswordExpirationDateResponse,omitempty"`
	GetPersonaResponseMessage                      []*GetPersonaResponseMessageType                      `xml:"m:GetPersonaResponseMessage,omitempty"`
	GetConversationItemsResponseMessage            []*GetConversationItemsResponseMessageType            `xml:"m:GetConversationItemsResponseMessage,omitempty"`
	GetUserRetentionPolicyTagsResponseMessage      []*GetUserRetentionPolicyTagsResponseMessageType      `xml:"m:GetUserRetentionPolicyTagsResponseMessage,omitempty"`
	GetUserPhotoResponseMessage                    []*GetUserPhotoResponseMessageType                    `xml:"m:GetUserPhotoResponseMessage,omitempty"`
	MarkAsJunkResponseMessage                      []*MarkAsJunkResponseMessageType                      `xml:"m:MarkAsJunkResponseMessage,omitempty"`
	MarkAsPhishingResponseMessage                  []*MarkAsPhishingResponseMessageType                  `xml:"m:MarkAsPhishingResponseMessage,omitempty"`
	ReportMessageResponseMessage                   []*ReportMessageResponseMessageType                   `xml:"m:ReportMessageResponseMessage,omitempty"`
	PostModernGroupItemResponseMessage             []*ItemInfoResponseMessageType                        `xml:"m:PostModernGroupItemResponseMessage,omitempty"`
	GetLastPrivateCatalogUpdateResponseMessage     []*ResponseMessageType                                `xml:"m:GetLastPrivateCatalogUpdateResponseMessage,omitempty"`
	GetPrivateCatalogAddInsResponseMessage         []*ResponseMessageType                                `xml:"m:GetPrivateCatalogAddInsResponseMessage,omitempty"`
}

type ItemInfoResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	Items               *ArrayOfRealItemsType `xml:"m:Items,omitempty"`
}

type ResponseMessageType struct {
	ResponseClass      ResponseClassType `xml:"ResponseClass,attr"`
	MessageText        XsString          `xml:"m:MessageText,omitempty"`
	ResponseCode       ResponseCodeType  `xml:"m:ResponseCode,omitempty"`
	DescriptiveLinkKey XsInt             `xml:"m:DescriptiveLinkKey,omitempty"`
	MessageXml         MessageXmlType    `xml:"m:MessageXml,omitempty"`
}

type DeleteItemResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
}

type UpdateItemResponseMessageType struct {
	ItemInfoResponseMessageType `xml:",omitempty"`
	ConflictResults             *ConflictResultsType `xml:"m:ConflictResults,omitempty"`
}

type UpdateItemInRecoverableItemsResponseMessageType struct {
	ItemInfoResponseMessageType `xml:",omitempty"`
	Attachments                 *ArrayOfAttachmentsType `xml:"m:Attachments,omitempty"`
	ConflictResults             *ConflictResultsType    `xml:"m:ConflictResults,omitempty"`
}

type FolderInfoResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	Folders             *ArrayOfFoldersType `xml:"m:Folders,omitempty"`
}

type FindFolderResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	RootFolder          *FindFolderParentType `xml:"m:RootFolder,omitempty"`
}

type AttachmentInfoResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	Attachments         *ArrayOfAttachmentsType `xml:"m:Attachments,omitempty"`
}

type DeleteAttachmentResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	RootItemId          *RootItemIdType `xml:"m:RootItemId,omitempty"`
}

type UploadItemsResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	ItemId              *ItemIdType `xml:"m:ItemId,omitempty"`
}

type ExportItemsResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	ItemId              *ItemIdType    `xml:"m:ItemId,omitempty"`
	Data                XsBase64Binary `xml:"m:Data,omitempty"`
}

type GetClientAccessTokenResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	Token               *ClientAccessTokenType `xml:"m:Token,omitempty"`
}

type FindItemResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	RootFolder          *FindItemParentType        `xml:"m:RootFolder,omitempty"`
	HighlightTerms      *ArrayOfHighlightTermsType `xml:"m:HighlightTerms,omitempty"`
}

type ResolveNamesResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	ResolutionSet       *ArrayOfResolutionType `xml:"m:ResolutionSet,omitempty"`
}

type ExpandDLResponseMessageType struct {
	ResponseMessageType     `xml:",omitempty"`
	IndexedPagingOffset     XsInt                   `xml:"IndexedPagingOffset,attr"`
	NumeratorOffset         XsInt                   `xml:"NumeratorOffset,attr"`
	AbsoluteDenominator     XsInt                   `xml:"AbsoluteDenominator,attr"`
	IncludesLastItemInRange XsBoolean               `xml:"IncludesLastItemInRange,attr"`
	TotalItemsInView        XsInt                   `xml:"TotalItemsInView,attr"`
	DLExpansion             *ArrayOfDLExpansionType `xml:"m:DLExpansion,omitempty"`
}

type GetServerTimeZonesResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	TimeZoneDefinitions *ArrayOfTimeZoneDefinitionType `xml:"m:TimeZoneDefinitions,omitempty"`
}

type GetEventsResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	Notification        *NotificationType `xml:"m:Notification,omitempty"`
}

type GetStreamingEventsResponseMessageType struct {
	ResponseMessageType  `xml:",omitempty"`
	Notifications        *NonEmptyArrayOfNotificationsType   `xml:"m:Notifications,omitempty"`
	ErrorSubscriptionIds *NonEmptyArrayOfSubscriptionIdsType `xml:"m:ErrorSubscriptionIds,omitempty"`
	ConnectionStatus     ConnectionStatusType                `xml:"m:ConnectionStatus,omitempty"`
}

type SubscribeResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	SubscriptionId      SubscriptionIdType `xml:"m:SubscriptionId,omitempty"`
	Watermark           WatermarkType      `xml:"m:Watermark,omitempty"`
}

type SendNotificationResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	Notification        *NotificationType `xml:"m:Notification,omitempty"`
}

type SyncFolderHierarchyResponseMessageType struct {
	ResponseMessageType       `xml:",omitempty"`
	SyncState                 XsString                        `xml:"m:SyncState,omitempty"`
	IncludesLastFolderInRange XsBoolean                       `xml:"m:IncludesLastFolderInRange,omitempty"`
	Changes                   *SyncFolderHierarchyChangesType `xml:"m:Changes,omitempty"`
}

type SyncFolderItemsResponseMessageType struct {
	ResponseMessageType     `xml:",omitempty"`
	SyncState               XsString                    `xml:"m:SyncState,omitempty"`
	IncludesLastItemInRange XsBoolean                   `xml:"m:IncludesLastItemInRange,omitempty"`
	Changes                 *SyncFolderItemsChangesType `xml:"m:Changes,omitempty"`
}

type ConvertIdResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	AlternateId         *AlternateIdBaseType `xml:"m:AlternateId,omitempty"`
}

type GetSharingMetadataResponseMessageType struct {
	ResponseMessageType                 `xml:",omitempty"`
	EncryptedSharedFolderDataCollection *ArrayOfEncryptedSharedFolderDataType `xml:"m:EncryptedSharedFolderDataCollection,omitempty"`
	InvalidRecipients                   *ArrayOfInvalidRecipientsType         `xml:"m:InvalidRecipients,omitempty"`
}

type RefreshSharingFolderResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
}

type GetSharingFolderResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	SharingFolderId     *FolderIdType `xml:"m:SharingFolderId,omitempty"`
}

type GetUserConfigurationResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	UserConfiguration   *UserConfigurationType `xml:"m:UserConfiguration,omitempty"`
}

type GetSpecificUserConfigurationResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	UserConfiguration   *UserConfigurationType `xml:"m:UserConfiguration,omitempty"`
}

type GetRoomListsResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	RoomLists           *ArrayOfEmailAddressesType `xml:"m:RoomLists,omitempty"`
}

type GetRoomsResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	Rooms               *ArrayOfRoomsType `xml:"m:Rooms,omitempty"`
}

type GetRemindersResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	Reminders           *ArrayOfRemindersType `xml:"m:Reminders,omitempty"`
}

type PerformReminderActionResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	UpdatedItemIds      *NonEmptyArrayOfItemIdsType `xml:"m:UpdatedItemIds,omitempty"`
}

type ApplyConversationActionResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
}

type FindMailboxStatisticsByKeywordsResponseMessageType struct {
	ResponseMessageType           `xml:",omitempty"`
	MailboxStatisticsSearchResult *MailboxStatisticsSearchResultType `xml:"m:MailboxStatisticsSearchResult,omitempty"`
}

type GetSearchableMailboxesResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	SearchableMailboxes *ArrayOfSearchableMailboxesType   `xml:"m:SearchableMailboxes,omitempty"`
	FailedMailboxes     *ArrayOfFailedSearchMailboxesType `xml:"m:FailedMailboxes,omitempty"`
}

type SearchMailboxesResponseMessageType struct {
	ResponseMessageType   `xml:",omitempty"`
	SearchMailboxesResult *SearchMailboxesResultType `xml:"m:SearchMailboxesResult,omitempty"`
}

type GetDiscoverySearchConfigurationResponseMessageType struct {
	ResponseMessageType           `xml:",omitempty"`
	DiscoverySearchConfigurations *ArrayOfDiscoverySearchConfigurationType `xml:"m:DiscoverySearchConfigurations,omitempty"`
}

type GetHoldOnMailboxesResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	MailboxHoldResult   *MailboxHoldResultType `xml:"m:MailboxHoldResult,omitempty"`
}

type SetHoldOnMailboxesResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	MailboxHoldResult   *MailboxHoldResultType `xml:"m:MailboxHoldResult,omitempty"`
}

type GetNonIndexableItemStatisticsResponseMessageType struct {
	ResponseMessageType        `xml:",omitempty"`
	NonIndexableItemStatistics *ArrayOfNonIndexableItemStatisticsType `xml:"m:NonIndexableItemStatistics,omitempty"`
}

type GetNonIndexableItemDetailsResponseMessageType struct {
	ResponseMessageType           `xml:",omitempty"`
	NonIndexableItemDetailsResult *NonIndexableItemDetailResultType `xml:"m:NonIndexableItemDetailsResult,omitempty"`
}

type FindPeopleResponseMessageType struct {
	ResponseMessageType       `xml:",omitempty"`
	People                    *ArrayOfPeopleType `xml:"m:People,omitempty"`
	TotalNumberOfPeopleInView XsInt              `xml:"m:TotalNumberOfPeopleInView,omitempty"`
	FirstMatchingRowIndex     XsInt              `xml:"m:FirstMatchingRowIndex,omitempty"`
	FirstLoadedRowIndex       XsInt              `xml:"m:FirstLoadedRowIndex,omitempty"`
	TransactionId             GuidType           `xml:"m:TransactionId,omitempty"`
}

type FindTagsResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	Tags                *ArrayOfStringsType `xml:"m:Tags,omitempty"`
}

type AddTagResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	WasSuccessful       XsBoolean `xml:"m:WasSuccessful,omitempty"`
}

type HideTagResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	WasSuccessful       XsBoolean `xml:"m:WasSuccessful,omitempty"`
}

type GetPasswordExpirationDateResponseMessageType struct {
	ResponseMessageType    `xml:",omitempty"`
	PasswordExpirationDate XsDateTime `xml:"m:PasswordExpirationDate,omitempty"`
}

type GetPersonaResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	Persona             *PersonaType `xml:"m:Persona,omitempty"`
}

type GetConversationItemsResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	Conversation        *ConversationResponseType `xml:"m:Conversation,omitempty"`
}

type GetUserRetentionPolicyTagsResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	RetentionPolicyTags *ArrayOfRetentionPolicyTagsType `xml:"m:RetentionPolicyTags,omitempty"`
}

type GetUserPhotoResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	HasChanged          XsBoolean      `xml:"m:HasChanged,omitempty"`
	PictureData         XsBase64Binary `xml:"m:PictureData,omitempty"`
}

type MarkAsJunkResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	MovedItemId         *ItemIdType `xml:"m:MovedItemId,omitempty"`
}

type MarkAsPhishingResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	MovedItemId         *ItemIdType `xml:"m:MovedItemId,omitempty"`
}

type ReportMessageResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	MovedItemId         *ItemIdType `xml:"m:MovedItemId,omitempty"`
	Policy              XsString    `xml:"m:Policy,omitempty"`
}

type ExpandDLType struct {
	BaseRequestType `xml:",omitempty"`
	Mailbox         *EmailAddressType `xml:"m:Mailbox,omitempty"`
}

type ExpandDLResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type GetServerTimeZonesType struct {
	BaseRequestType        `xml:",omitempty"`
	ReturnFullTimeZoneData XsBoolean                      `xml:"ReturnFullTimeZoneData,attr"`
	Ids                    *NonEmptyArrayOfTimeZoneIdType `xml:"m:Ids,omitempty"`
}

type GetServerTimeZonesResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type FindFolderType struct {
	BaseRequestType          `xml:",omitempty"`
	Traversal                FolderQueryTraversalType          `xml:"Traversal,attr"`
	FolderShape              *FolderResponseShapeType          `xml:"m:FolderShape,omitempty"`
	Restriction              *RestrictionType                  `xml:"m:Restriction,omitempty"`
	ParentFolderIds          *NonEmptyArrayOfBaseFolderIdsType `xml:"m:ParentFolderIds,omitempty"`
	IndexedPageFolderView    *IndexedPageViewType              `xml:"m:IndexedPageFolderView,omitempty"`
	FractionalPageFolderView *FractionalPageViewType           `xml:"m:FractionalPageFolderView,omitempty"`
}

type FindFolderResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type FindItemType struct {
	BaseRequestType             `xml:",omitempty"`
	Traversal                   ItemQueryTraversalType            `xml:"Traversal,attr"`
	ItemShape                   *ItemResponseShapeType            `xml:"m:ItemShape,omitempty"`
	Restriction                 *RestrictionType                  `xml:"m:Restriction,omitempty"`
	SortOrder                   *NonEmptyArrayOfFieldOrdersType   `xml:"m:SortOrder,omitempty"`
	ParentFolderIds             *NonEmptyArrayOfBaseFolderIdsType `xml:"m:ParentFolderIds,omitempty"`
	QueryString                 *QueryStringType                  `xml:"m:QueryString,omitempty"`
	IndexedPageItemView         *IndexedPageViewType              `xml:"m:IndexedPageItemView,omitempty"`
	FractionalPageItemView      *FractionalPageViewType           `xml:"m:FractionalPageItemView,omitempty"`
	SeekToConditionPageItemView *SeekToConditionPageViewType      `xml:"m:SeekToConditionPageItemView,omitempty"`
	CalendarView                *CalendarViewType                 `xml:"m:CalendarView,omitempty"`
	ContactsView                *ContactsViewType                 `xml:"m:ContactsView,omitempty"`
	GroupBy                     *GroupByType                      `xml:"m:GroupBy,omitempty"`
	DistinguishedGroupBy        *DistinguishedGroupByType         `xml:"m:DistinguishedGroupBy,omitempty"`
}

type QueryStringType struct {
	CharData             XsString  `xml:",chardata"`
	ResetCache           XsBoolean `xml:"ResetCache,attr"`
	ReturnHighlightTerms XsBoolean `xml:"ReturnHighlightTerms,attr"`
	ReturnDeletedItems   XsBoolean `xml:"ReturnDeletedItems,attr"`
}

type FindItemResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type GetFolderType struct {
	BaseRequestType `xml:",omitempty"`
	FolderShape     *FolderResponseShapeType          `xml:"m:FolderShape,omitempty"`
	FolderIds       *NonEmptyArrayOfBaseFolderIdsType `xml:"m:FolderIds,omitempty"`
}

type GetFolderResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type UploadItemsType struct {
	BaseRequestType `xml:",omitempty"`
	Items           *NonEmptyArrayOfUploadItemsType `xml:"m:Items,omitempty"`
}

type UploadItemsResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type ExportItemsType struct {
	BaseRequestType `xml:",omitempty"`
	ItemIds         *NonEmptyArrayOfItemIdsType `xml:"m:ItemIds,omitempty"`
}

type ExportItemsResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type ConvertIdType struct {
	BaseRequestType   `xml:",omitempty"`
	DestinationFormat IdFormatType                     `xml:"DestinationFormat,attr"`
	SourceIds         *NonEmptyArrayOfAlternateIdsType `xml:"m:SourceIds,omitempty"`
}

type ConvertIdResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type CreateFolderType struct {
	BaseRequestType `xml:",omitempty"`
	ParentFolderId  *TargetFolderIdType         `xml:"m:ParentFolderId,omitempty"`
	Folders         *NonEmptyArrayOfFoldersType `xml:"m:Folders,omitempty"`
}

type CreateFolderResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type CreateFolderPathType struct {
	BaseRequestType    `xml:",omitempty"`
	ParentFolderId     *TargetFolderIdType         `xml:"m:ParentFolderId,omitempty"`
	RelativeFolderPath *NonEmptyArrayOfFoldersType `xml:"m:RelativeFolderPath,omitempty"`
}

type CreateFolderPathResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type DeleteFolderType struct {
	BaseRequestType `xml:",omitempty"`
	DeleteType      DisposalType                      `xml:"DeleteType,attr"`
	FolderIds       *NonEmptyArrayOfBaseFolderIdsType `xml:"m:FolderIds,omitempty"`
}

type DeleteFolderResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type EmptyFolderType struct {
	BaseRequestType  `xml:",omitempty"`
	DeleteType       DisposalType                      `xml:"DeleteType,attr"`
	DeleteSubFolders XsBoolean                         `xml:"DeleteSubFolders,attr"`
	FolderIds        *NonEmptyArrayOfBaseFolderIdsType `xml:"m:FolderIds,omitempty"`
}

type EmptyFolderResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type UpdateFolderType struct {
	BaseRequestType `xml:",omitempty"`
	FolderChanges   *NonEmptyArrayOfFolderChangesType `xml:"m:FolderChanges,omitempty"`
}

type UpdateFolderResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type MoveFolderType struct {
	BaseMoveCopyFolderType `xml:",omitempty"`
}

type BaseMoveCopyFolderType struct {
	BaseRequestType `xml:",omitempty"`
	ToFolderId      *TargetFolderIdType               `xml:"m:ToFolderId,omitempty"`
	FolderIds       *NonEmptyArrayOfBaseFolderIdsType `xml:"m:FolderIds,omitempty"`
}

type MoveFolderResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type CopyFolderType struct {
	BaseMoveCopyFolderType `xml:",omitempty"`
}

type CopyFolderResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type SubscribeType struct {
	BaseRequestType              `xml:",omitempty"`
	PullSubscriptionRequest      *PullSubscriptionRequestType      `xml:"m:PullSubscriptionRequest,omitempty"`
	PushSubscriptionRequest      *PushSubscriptionRequestType      `xml:"m:PushSubscriptionRequest,omitempty"`
	StreamingSubscriptionRequest *StreamingSubscriptionRequestType `xml:"m:StreamingSubscriptionRequest,omitempty"`
}

type SubscribeResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type UnsubscribeType struct {
	BaseRequestType `xml:",omitempty"`
	SubscriptionId  SubscriptionIdType `xml:"m:SubscriptionId,omitempty"`
}

type UnsubscribeResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type GetEventsType struct {
	BaseRequestType `xml:",omitempty"`
	SubscriptionId  SubscriptionIdType `xml:"m:SubscriptionId,omitempty"`
	Watermark       WatermarkType      `xml:"m:Watermark,omitempty"`
}

type GetEventsResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type GetStreamingEventsType struct {
	BaseRequestType   `xml:",omitempty"`
	SubscriptionIds   *NonEmptyArrayOfSubscriptionIdsType        `xml:"m:SubscriptionIds,omitempty"`
	ConnectionTimeout StreamingSubscriptionConnectionTimeoutType `xml:"m:ConnectionTimeout,omitempty"`
}

type GetStreamingEventsResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type SyncFolderHierarchyType struct {
	BaseRequestType `xml:",omitempty"`
	FolderShape     *FolderResponseShapeType `xml:"m:FolderShape,omitempty"`
	SyncFolderId    *TargetFolderIdType      `xml:"m:SyncFolderId,omitempty"`
	SyncState       XsString                 `xml:"m:SyncState,omitempty"`
}

type SyncFolderHierarchyResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type SyncFolderItemsType struct {
	BaseRequestType    `xml:",omitempty"`
	ItemShape          *ItemResponseShapeType     `xml:"m:ItemShape,omitempty"`
	SyncFolderId       *TargetFolderIdType        `xml:"m:SyncFolderId,omitempty"`
	SyncState          XsString                   `xml:"m:SyncState,omitempty"`
	Ignore             *ArrayOfBaseItemIdsType    `xml:"m:Ignore,omitempty"`
	MaxChangesReturned MaxSyncChangesReturnedType `xml:"m:MaxChangesReturned,omitempty"`
	SyncScope          SyncFolderItemsScopeType   `xml:"m:SyncScope,omitempty"`
}

type SyncFolderItemsResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type CreateManagedFolderRequestType struct {
	BaseRequestType `xml:",omitempty"`
	FolderNames     *NonEmptyArrayOfFolderNamesType `xml:"m:FolderNames,omitempty"`
	Mailbox         *EmailAddressType               `xml:"m:Mailbox,omitempty"`
}

type CreateManagedFolderResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type GetItemType struct {
	BaseRequestType `xml:",omitempty"`
	ItemShape       *ItemResponseShapeType          `xml:"m:ItemShape,omitempty"`
	ItemIds         *NonEmptyArrayOfBaseItemIdsType `xml:"m:ItemIds,omitempty"`
}

type GetItemResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type CreateItemType struct {
	BaseRequestType        `xml:",omitempty"`
	MessageDisposition     MessageDispositionType                  `xml:"MessageDisposition,attr"`
	SendMeetingInvitations CalendarItemCreateOrDeleteOperationType `xml:"SendMeetingInvitations,attr"`
	SavedItemFolderId      *TargetFolderIdType                     `xml:"m:SavedItemFolderId,omitempty"`
	Items                  *NonEmptyArrayOfAllItemsType            `xml:"m:Items,omitempty"`
}

type CreateItemResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type DeleteItemType struct {
	BaseRequestType          `xml:",omitempty"`
	DeleteType               DisposalType                            `xml:"DeleteType,attr"`
	SendMeetingCancellations CalendarItemCreateOrDeleteOperationType `xml:"SendMeetingCancellations,attr"`
	AffectedTaskOccurrences  AffectedTaskOccurrencesType             `xml:"AffectedTaskOccurrences,attr"`
	SuppressReadReceipts     XsBoolean                               `xml:"SuppressReadReceipts,attr"`
	ItemIds                  *NonEmptyArrayOfBaseItemIdsType         `xml:"m:ItemIds,omitempty"`
}

type DeleteItemResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type UpdateItemType struct {
	BaseRequestType                       `xml:",omitempty"`
	ConflictResolution                    ConflictResolutionType          `xml:"ConflictResolution,attr"`
	MessageDisposition                    MessageDispositionType          `xml:"MessageDisposition,attr"`
	SendMeetingInvitationsOrCancellations CalendarItemUpdateOperationType `xml:"SendMeetingInvitationsOrCancellations,attr"`
	SuppressReadReceipts                  XsBoolean                       `xml:"SuppressReadReceipts,attr"`
	SavedItemFolderId                     *TargetFolderIdType             `xml:"m:SavedItemFolderId,omitempty"`
	ItemChanges                           *NonEmptyArrayOfItemChangesType `xml:"m:ItemChanges,omitempty"`
}

type UpdateItemResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type UpdateItemInRecoverableItemsType struct {
	BaseRequestType   `xml:",omitempty"`
	ItemId            *ItemIdType                                `xml:"m:ItemId,omitempty"`
	Updates           *NonEmptyArrayOfItemChangeDescriptionsType `xml:"m:Updates,omitempty"`
	Attachments       *NonEmptyArrayOfAttachmentsType            `xml:"m:Attachments,omitempty"`
	MakeItemImmutable XsBoolean                                  `xml:"m:MakeItemImmutable,omitempty"`
}

type UpdateItemInRecoverableItemsResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type SendItemType struct {
	BaseRequestType   `xml:",omitempty"`
	SaveItemToFolder  XsBoolean                       `xml:"SaveItemToFolder,attr"`
	ItemIds           *NonEmptyArrayOfBaseItemIdsType `xml:"m:ItemIds,omitempty"`
	SavedItemFolderId *TargetFolderIdType             `xml:"m:SavedItemFolderId,omitempty"`
}

type SendItemResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type MoveItemType struct {
	BaseMoveCopyItemType `xml:",omitempty"`
}

type BaseMoveCopyItemType struct {
	BaseRequestType  `xml:",omitempty"`
	ToFolderId       *TargetFolderIdType             `xml:"m:ToFolderId,omitempty"`
	ItemIds          *NonEmptyArrayOfBaseItemIdsType `xml:"m:ItemIds,omitempty"`
	ReturnNewItemIds XsBoolean                       `xml:"m:ReturnNewItemIds,omitempty"`
}

type MoveItemResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type CopyItemType struct {
	BaseMoveCopyItemType `xml:",omitempty"`
}

type CopyItemResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type ArchiveItemType struct {
	BaseRequestType       `xml:",omitempty"`
	ArchiveSourceFolderId *TargetFolderIdType             `xml:"m:ArchiveSourceFolderId,omitempty"`
	ItemIds               *NonEmptyArrayOfBaseItemIdsType `xml:"m:ItemIds,omitempty"`
}

type ArchiveItemResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type CreateAttachmentType struct {
	BaseRequestType `xml:",omitempty"`
	ParentItemId    *ItemIdType                     `xml:"m:ParentItemId,omitempty"`
	Attachments     *NonEmptyArrayOfAttachmentsType `xml:"m:Attachments,omitempty"`
}

type CreateAttachmentResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type DeleteAttachmentType struct {
	BaseRequestType `xml:",omitempty"`
	AttachmentIds   *NonEmptyArrayOfRequestAttachmentIdsType `xml:"m:AttachmentIds,omitempty"`
}

type DeleteAttachmentResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type GetAttachmentType struct {
	BaseRequestType `xml:",omitempty"`
	AttachmentShape *AttachmentResponseShapeType             `xml:"m:AttachmentShape,omitempty"`
	AttachmentIds   *NonEmptyArrayOfRequestAttachmentIdsType `xml:"m:AttachmentIds,omitempty"`
}

type GetAttachmentResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type GetClientAccessTokenType struct {
	BaseRequestType `xml:",omitempty"`
	TokenRequests   *NonEmptyArrayOfClientAccessTokenRequestsType `xml:"m:TokenRequests,omitempty"`
}

type GetClientAccessTokenResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type GetDelegateType struct {
	BaseDelegateType   `xml:",omitempty"`
	IncludePermissions XsBoolean          `xml:"IncludePermissions,attr"`
	UserIds            *ArrayOfUserIdType `xml:"m:UserIds,omitempty"`
}

type BaseDelegateType struct {
	BaseRequestType `xml:",omitempty"`
	Mailbox         *EmailAddressType `xml:"m:Mailbox,omitempty"`
}

type GetDelegateResponseMessageType struct {
	BaseDelegateResponseMessageType `xml:",omitempty"`
	DeliverMeetingRequests          DeliverMeetingRequestsType `xml:"m:DeliverMeetingRequests,omitempty"`
}

type BaseDelegateResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	ResponseMessages    *ArrayOfDelegateUserResponseMessageType `xml:"m:ResponseMessages,omitempty"`
}

type ArrayOfDelegateUserResponseMessageType struct {
	DelegateUserResponseMessageType []*DelegateUserResponseMessageType `xml:"m:DelegateUserResponseMessageType,omitempty"`
}

type DelegateUserResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	DelegateUser        *DelegateUserType `xml:"m:DelegateUser,omitempty"`
}

type AddDelegateType struct {
	BaseDelegateType       `xml:",omitempty"`
	DelegateUsers          *ArrayOfDelegateUserType   `xml:"m:DelegateUsers,omitempty"`
	DeliverMeetingRequests DeliverMeetingRequestsType `xml:"m:DeliverMeetingRequests,omitempty"`
}

type AddDelegateResponseMessageType struct {
	BaseDelegateResponseMessageType `xml:",omitempty"`
}

type RemoveDelegateType struct {
	BaseDelegateType `xml:",omitempty"`
	UserIds          *ArrayOfUserIdType `xml:"m:UserIds,omitempty"`
}

type RemoveDelegateResponseMessageType struct {
	BaseDelegateResponseMessageType `xml:",omitempty"`
}

type UpdateDelegateType struct {
	BaseDelegateType       `xml:",omitempty"`
	DelegateUsers          *ArrayOfDelegateUserType   `xml:"m:DelegateUsers,omitempty"`
	DeliverMeetingRequests DeliverMeetingRequestsType `xml:"m:DeliverMeetingRequests,omitempty"`
}

type UpdateDelegateResponseMessageType struct {
	BaseDelegateResponseMessageType `xml:",omitempty"`
}

type CreateUserConfigurationType struct {
	BaseRequestType   `xml:",omitempty"`
	UserConfiguration *UserConfigurationType `xml:"m:UserConfiguration,omitempty"`
}

type CreateUserConfigurationResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type DeleteUserConfigurationType struct {
	BaseRequestType       `xml:",omitempty"`
	UserConfigurationName *UserConfigurationNameType `xml:"m:UserConfigurationName,omitempty"`
}

type DeleteUserConfigurationResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type GetUserConfigurationType struct {
	BaseRequestType             `xml:",omitempty"`
	UserConfigurationName       *UserConfigurationNameType    `xml:"m:UserConfigurationName,omitempty"`
	UserConfigurationProperties UserConfigurationPropertyType `xml:"m:UserConfigurationProperties,omitempty"`
}

type GetUserConfigurationResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type GetSpecificUserConfigurationType struct {
	BaseRequestType             `xml:",omitempty"`
	UserConfigurationName       *UserConfigurationNameType    `xml:"m:UserConfigurationName,omitempty"`
	UserConfigurationProperties UserConfigurationPropertyType `xml:"m:UserConfigurationProperties,omitempty"`
}

type GetSpecificUserConfigurationResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type UpdateUserConfigurationType struct {
	BaseRequestType   `xml:",omitempty"`
	UserConfiguration *UserConfigurationType `xml:"m:UserConfiguration,omitempty"`
}

type UpdateUserConfigurationResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type GetUserAvailabilityRequestType struct {
	BaseRequestType        `xml:",omitempty"`
	TimeZone               *SerializableTimeZone       `xml:"t:TimeZone,omitempty"`
	MailboxDataArray       *ArrayOfMailboxData         `xml:"m:MailboxDataArray,omitempty"`
	FreeBusyViewOptions    *FreeBusyViewOptionsType    `xml:"t:FreeBusyViewOptions,omitempty"`
	SuggestionsViewOptions *SuggestionsViewOptionsType `xml:"t:SuggestionsViewOptions,omitempty"`
}

type GetUserAvailabilityResponseType struct {
	FreeBusyResponseArray *ArrayOfFreeBusyResponse `xml:"m:FreeBusyResponseArray,omitempty"`
	SuggestionsResponse   *SuggestionsResponseType `xml:"m:SuggestionsResponse,omitempty"`
}

type ArrayOfFreeBusyResponse struct {
	FreeBusyResponse []*FreeBusyResponseType `xml:"m:FreeBusyResponse,omitempty"`
}

type FreeBusyResponseType struct {
	ResponseMessage *ResponseMessageType `xml:"m:ResponseMessage,omitempty"`
	FreeBusyView    *FreeBusyView        `xml:"m:FreeBusyView,omitempty"`
}

type SuggestionsResponseType struct {
	ResponseMessage          *ResponseMessageType        `xml:"m:ResponseMessage,omitempty"`
	SuggestionDayResultArray *ArrayOfSuggestionDayResult `xml:"m:SuggestionDayResultArray,omitempty"`
}

type GetUserOofSettingsRequest struct {
	BaseRequestType `xml:",omitempty"`
	Mailbox         *EmailAddress `xml:"t:Mailbox,omitempty"`
}

type GetUserOofSettingsResponse struct {
	ResponseMessage  *ResponseMessageType `xml:"m:ResponseMessage,omitempty"`
	OofSettings      *UserOofSettings     `xml:"t:OofSettings,omitempty"`
	AllowExternalOof ExternalAudience     `xml:"m:AllowExternalOof,omitempty"`
}

type SetUserOofSettingsRequest struct {
	BaseRequestType `xml:",omitempty"`
	Mailbox         *EmailAddress    `xml:"t:Mailbox,omitempty"`
	UserOofSettings *UserOofSettings `xml:"t:UserOofSettings,omitempty"`
}

type SetUserOofSettingsResponse struct {
	ResponseMessage *ResponseMessageType `xml:"m:ResponseMessage,omitempty"`
}

type GetServiceConfigurationType struct {
	BaseRequestType             `xml:",omitempty"`
	ActingAs                    *EmailAddressType                `xml:"m:ActingAs,omitempty"`
	RequestedConfiguration      *ArrayOfServiceConfigurationType `xml:"m:RequestedConfiguration,omitempty"`
	ConfigurationRequestDetails ConfigurationRequestDetailsType  `xml:"m:ConfigurationRequestDetails,omitempty"`
}

type ArrayOfServiceConfigurationType struct {
	ConfigurationName []ServiceConfigurationType `xml:"m:ConfigurationName,omitempty"`
}

type GetServiceConfigurationResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	ResponseMessages    *ArrayOfServiceConfigurationResponseMessageType `xml:"m:ResponseMessages,omitempty"`
}

type ArrayOfServiceConfigurationResponseMessageType struct {
	ServiceConfigurationResponseMessageType []*ServiceConfigurationResponseMessageType `xml:"m:ServiceConfigurationResponseMessageType,omitempty"`
}

type ServiceConfigurationResponseMessageType struct {
	ResponseMessageType           `xml:",omitempty"`
	MailTipsConfiguration         *MailTipsServiceConfiguration        `xml:"m:MailTipsConfiguration,omitempty"`
	UnifiedMessagingConfiguration *UnifiedMessageServiceConfiguration  `xml:"m:UnifiedMessagingConfiguration,omitempty"`
	ProtectionRulesConfiguration  *ProtectionRulesServiceConfiguration `xml:"m:ProtectionRulesConfiguration,omitempty"`
	PolicyNudgeRulesConfiguration PolicyNudgeRulesServiceConfiguration `xml:"m:PolicyNudgeRulesConfiguration,omitempty"`
	SharePointURLsConfiguration   *SharePointURLsServiceConfiguration  `xml:"m:SharePointURLsConfiguration,omitempty"`
}

type GetMailTipsType struct {
	BaseRequestType   `xml:",omitempty"`
	SendingAs         *EmailAddressType      `xml:"m:SendingAs,omitempty"`
	Recipients        *ArrayOfRecipientsType `xml:"m:Recipients,omitempty"`
	MailTipsRequested MailTipTypes           `xml:"m:MailTipsRequested,omitempty"`
}

type GetMailTipsResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	ResponseMessages    *ArrayOfMailTipsResponseMessageType `xml:"m:ResponseMessages,omitempty"`
}

type ArrayOfMailTipsResponseMessageType struct {
	MailTipsResponseMessageType []*MailTipsResponseMessageType `xml:"m:MailTipsResponseMessageType,omitempty"`
}

type MailTipsResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	MailTips            *MailTips `xml:"m:MailTips,omitempty"`
}

type PlayOnPhoneType struct {
	BaseRequestType `xml:",omitempty"`
	ItemId          *ItemIdType `xml:"m:ItemId,omitempty"`
	DialString      XsString    `xml:"m:DialString,omitempty"`
}

type PlayOnPhoneResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	PhoneCallId         *PhoneCallIdType `xml:"m:PhoneCallId,omitempty"`
}

type GetPhoneCallInformationType struct {
	BaseRequestType `xml:",omitempty"`
	PhoneCallId     *PhoneCallIdType `xml:"m:PhoneCallId,omitempty"`
}

type GetPhoneCallInformationResponseMessageType struct {
	ResponseMessageType  `xml:",omitempty"`
	PhoneCallInformation *PhoneCallInformationType `xml:"m:PhoneCallInformation,omitempty"`
}

type DisconnectPhoneCallType struct {
	BaseRequestType `xml:",omitempty"`
	PhoneCallId     *PhoneCallIdType `xml:"m:PhoneCallId,omitempty"`
}

type DisconnectPhoneCallResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
}

type GetSharingMetadataType struct {
	BaseRequestType   `xml:",omitempty"`
	IdOfFolderToShare *FolderIdType           `xml:"m:IdOfFolderToShare,omitempty"`
	SenderSmtpAddress NonEmptyStringType      `xml:"m:SenderSmtpAddress,omitempty"`
	Recipients        *ArrayOfSmtpAddressType `xml:"m:Recipients,omitempty"`
}

type RefreshSharingFolderType struct {
	BaseRequestType `xml:",omitempty"`
	SharingFolderId *FolderIdType `xml:"m:SharingFolderId,omitempty"`
}

type GetSharingFolderType struct {
	BaseRequestType `xml:",omitempty"`
	SmtpAddress     NonEmptyStringType `xml:"m:SmtpAddress,omitempty"`
	DataType        SharingDataType    `xml:"m:DataType,omitempty"`
	SharedFolderId  NonEmptyStringType `xml:"m:SharedFolderId,omitempty"`
}

type SetTeamMailboxRequestType struct {
	BaseRequestType   `xml:",omitempty"`
	EmailAddress      *EmailAddressType             `xml:"m:EmailAddress,omitempty"`
	SharePointSiteUrl XsString                      `xml:"m:SharePointSiteUrl,omitempty"`
	State             TeamMailboxLifecycleStateType `xml:"m:State,omitempty"`
}

type SetTeamMailboxResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
}

type UnpinTeamMailboxRequestType struct {
	BaseRequestType `xml:",omitempty"`
	EmailAddress    *EmailAddressType `xml:"m:EmailAddress,omitempty"`
}

type UnpinTeamMailboxResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
}

type GetRoomListsType struct {
	BaseRequestType `xml:",omitempty"`
}

type GetRoomsType struct {
	BaseRequestType `xml:",omitempty"`
	RoomList        *EmailAddressType `xml:"m:RoomList,omitempty"`
}

type FindMessageTrackingReportRequestType struct {
	BaseRequestType `xml:",omitempty"`
}

type FindMessageTrackingReportResponseMessageType struct {
	ResponseMessageType          `xml:",omitempty"`
	Diagnostics                  *ArrayOfStringsType                         `xml:"m:Diagnostics,omitempty"`
	MessageTrackingSearchResults *ArrayOfFindMessageTrackingSearchResultType `xml:"m:MessageTrackingSearchResults,omitempty"`
	ExecutedSearchScope          XsString                                    `xml:"m:ExecutedSearchScope,omitempty"`
	Errors                       *ArrayOfArraysOfTrackingPropertiesType      `xml:"m:Errors,omitempty"`
	Properties                   *ArrayOfTrackingPropertiesType              `xml:"m:Properties,omitempty"`
}

type GetMessageTrackingReportRequestType struct {
	BaseRequestType `xml:",omitempty"`
}

type GetMessageTrackingReportResponseMessageType struct {
	ResponseMessageType   `xml:",omitempty"`
	MessageTrackingReport MessageTrackingReportType              `xml:"m:MessageTrackingReport,omitempty"`
	Diagnostics           *ArrayOfStringsType                    `xml:"m:Diagnostics,omitempty"`
	Errors                *ArrayOfArraysOfTrackingPropertiesType `xml:"m:Errors,omitempty"`
	Properties            *ArrayOfTrackingPropertiesType         `xml:"m:Properties,omitempty"`
}

type FindConversationType struct {
	BaseRequestType             `xml:",omitempty"`
	Traversal                   ConversationQueryTraversalType  `xml:"Traversal,attr"`
	ViewFilter                  ViewFilterType                  `xml:"ViewFilter,attr"`
	SortOrder                   *NonEmptyArrayOfFieldOrdersType `xml:"m:SortOrder,omitempty"`
	ParentFolderId              *TargetFolderIdType             `xml:"m:ParentFolderId,omitempty"`
	MailboxScope                MailboxSearchLocationType       `xml:"m:MailboxScope,omitempty"`
	QueryString                 *QueryStringType                `xml:"m:QueryString,omitempty"`
	ConversationShape           *ConversationResponseShapeType  `xml:"m:ConversationShape,omitempty"`
	IndexedPageItemView         *IndexedPageViewType            `xml:"m:IndexedPageItemView,omitempty"`
	SeekToConditionPageItemView *SeekToConditionPageViewType    `xml:"m:SeekToConditionPageItemView,omitempty"`
}

type FindConversationResponseMessageType struct {
	ResponseMessageType      `xml:",omitempty"`
	Conversations            *ArrayOfConversationsType  `xml:"m:Conversations,omitempty"`
	HighlightTerms           *ArrayOfHighlightTermsType `xml:"m:HighlightTerms,omitempty"`
	TotalConversationsInView XsInt                      `xml:"m:TotalConversationsInView,omitempty"`
	IndexedOffset            XsInt                      `xml:"m:IndexedOffset,omitempty"`
}

type ApplyConversationActionType struct {
	BaseRequestType     `xml:",omitempty"`
	ConversationActions *NonEmptyArrayOfApplyConversationActionType `xml:"m:ConversationActions,omitempty"`
}

type ApplyConversationActionResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type GetConversationItemsType struct {
	BaseRequestType  `xml:",omitempty"`
	ItemShape        *ItemResponseShapeType            `xml:"m:ItemShape,omitempty"`
	FoldersToIgnore  *NonEmptyArrayOfBaseFolderIdsType `xml:"m:FoldersToIgnore,omitempty"`
	MaxItemsToReturn XsInt                             `xml:"m:MaxItemsToReturn,omitempty"`
	SortOrder        ConversationNodeSortOrder         `xml:"m:SortOrder,omitempty"`
	MailboxScope     MailboxSearchLocationType         `xml:"m:MailboxScope,omitempty"`
	Conversations    *ArrayOfConversationRequestsType  `xml:"m:Conversations,omitempty"`
}

type GetConversationItemsResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type FindPeopleType struct {
	BaseRequestType             `xml:",omitempty"`
	PersonaShape                *PersonaResponseShapeType       `xml:"m:PersonaShape,omitempty"`
	IndexedPageItemView         *IndexedPageViewType            `xml:"m:IndexedPageItemView,omitempty"`
	Restriction                 *RestrictionType                `xml:"m:Restriction,omitempty"`
	AggregationRestriction      *RestrictionType                `xml:"m:AggregationRestriction,omitempty"`
	SortOrder                   *NonEmptyArrayOfFieldOrdersType `xml:"m:SortOrder,omitempty"`
	ParentFolderId              *TargetFolderIdType             `xml:"m:ParentFolderId,omitempty"`
	QueryString                 XsString                        `xml:"m:QueryString,omitempty"`
	SearchPeopleSuggestionIndex XsBoolean                       `xml:"m:SearchPeopleSuggestionIndex,omitempty"`
	TopicQueryString            XsString                        `xml:"m:TopicQueryString,omitempty"`
	Context                     *ArrayOfContextProperty         `xml:"m:Context,omitempty"`
	QuerySources                *ArrayOfPeopleQuerySource       `xml:"m:QuerySources,omitempty"`
	ReturnFlattenedResults      XsBoolean                       `xml:"m:ReturnFlattenedResults,omitempty"`
}

type FindTagsType struct {
	BaseRequestType     `xml:",omitempty"`
	IndexedPageItemView *IndexedPageViewType            `xml:"m:IndexedPageItemView,omitempty"`
	SortOrder           *NonEmptyArrayOfFieldOrdersType `xml:"m:SortOrder,omitempty"`
	QueryString         XsString                        `xml:"m:QueryString,omitempty"`
	Context             *ArrayOfContextProperty         `xml:"m:Context,omitempty"`
}

type AddTagType struct {
	BaseRequestType `xml:",omitempty"`
	Tag             XsString `xml:"m:Tag,omitempty"`
	AppName         XsString `xml:"m:AppName,omitempty"`
}

type HideTagType struct {
	BaseRequestType `xml:",omitempty"`
	Tag             XsString `xml:"m:Tag,omitempty"`
}

type GetPersonaType struct {
	BaseRequestType      `xml:",omitempty"`
	PersonaId            *ItemIdType                        `xml:"m:PersonaId,omitempty"`
	EmailAddress         *EmailAddressType                  `xml:"m:EmailAddress,omitempty"`
	ParentFolderId       *TargetFolderIdType                `xml:"m:ParentFolderId,omitempty"`
	ItemLinkId           XsString                           `xml:"m:ItemLinkId,omitempty"`
	AdditionalProperties *NonEmptyArrayOfPathsToElementType `xml:"m:AdditionalProperties,omitempty"`
}

type GetInboxRulesRequestType struct {
	BaseRequestType    `xml:",omitempty"`
	MailboxSmtpAddress XsString `xml:"m:MailboxSmtpAddress,omitempty"`
}

type GetInboxRulesResponseType struct {
	ResponseMessageType   `xml:",omitempty"`
	OutlookRuleBlobExists XsBoolean         `xml:"m:OutlookRuleBlobExists,omitempty"`
	InboxRules            *ArrayOfRulesType `xml:"m:InboxRules,omitempty"`
}

type UpdateInboxRulesRequestType struct {
	BaseRequestType       `xml:",omitempty"`
	MailboxSmtpAddress    XsString                   `xml:"m:MailboxSmtpAddress,omitempty"`
	RemoveOutlookRuleBlob XsBoolean                  `xml:"m:RemoveOutlookRuleBlob,omitempty"`
	Operations            *ArrayOfRuleOperationsType `xml:"m:Operations,omitempty"`
}

type UpdateInboxRulesResponseType struct {
	ResponseMessageType `xml:",omitempty"`
	RuleOperationErrors *ArrayOfRuleOperationErrorsType `xml:"m:RuleOperationErrors,omitempty"`
}

type GetPasswordExpirationDateType struct {
	BaseRequestType    `xml:",omitempty"`
	MailboxSmtpAddress XsString `xml:"m:MailboxSmtpAddress,omitempty"`
}

type GetSearchableMailboxesType struct {
	BaseRequestType       `xml:",omitempty"`
	SearchFilter          XsString  `xml:"m:SearchFilter,omitempty"`
	ExpandGroupMembership XsBoolean `xml:"m:ExpandGroupMembership,omitempty"`
}

type SearchMailboxesType struct {
	BaseRequestType          `xml:",omitempty"`
	SearchQueries            *NonEmptyArrayOfMailboxQueriesType `xml:"m:SearchQueries,omitempty"`
	ResultType               SearchResultType                   `xml:"m:ResultType,omitempty"`
	PreviewItemResponseShape *PreviewItemResponseShapeType      `xml:"m:PreviewItemResponseShape,omitempty"`
	SortBy                   *FieldOrderType                    `xml:"m:SortBy,omitempty"`
	Language                 XsString                           `xml:"m:Language,omitempty"`
	Deduplication            XsBoolean                          `xml:"m:Deduplication,omitempty"`
	PageSize                 XsInt                              `xml:"m:PageSize,omitempty"`
	PageItemReference        XsString                           `xml:"m:PageItemReference,omitempty"`
	PageDirection            SearchPageDirectionType            `xml:"m:PageDirection,omitempty"`
}

type SearchMailboxesResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type GetDiscoverySearchConfigurationType struct {
	BaseRequestType              `xml:",omitempty"`
	SearchId                     XsString  `xml:"m:SearchId,omitempty"`
	ExpandGroupMembership        XsBoolean `xml:"m:ExpandGroupMembership,omitempty"`
	InPlaceHoldConfigurationOnly XsBoolean `xml:"m:InPlaceHoldConfigurationOnly,omitempty"`
}

type GetHoldOnMailboxesType struct {
	BaseRequestType `xml:",omitempty"`
	HoldId          XsString `xml:"m:HoldId,omitempty"`
}

type SetHoldOnMailboxesType struct {
	BaseRequestType          `xml:",omitempty"`
	ActionType               HoldActionType      `xml:"m:ActionType,omitempty"`
	HoldId                   XsString            `xml:"m:HoldId,omitempty"`
	Query                    XsString            `xml:"m:Query,omitempty"`
	Mailboxes                *ArrayOfStringsType `xml:"m:Mailboxes,omitempty"`
	Language                 XsString            `xml:"m:Language,omitempty"`
	IncludeNonIndexableItems XsBoolean           `xml:"m:IncludeNonIndexableItems,omitempty"`
	Deduplication            XsBoolean           `xml:"m:Deduplication,omitempty"`
	InPlaceHoldIdentity      XsString            `xml:"m:InPlaceHoldIdentity,omitempty"`
	ItemHoldPeriod           XsString            `xml:"m:ItemHoldPeriod,omitempty"`
}

type GetNonIndexableItemStatisticsType struct {
	BaseRequestType   `xml:",omitempty"`
	Mailboxes         *NonEmptyArrayOfLegacyDNsType `xml:"m:Mailboxes,omitempty"`
	SearchArchiveOnly XsBoolean                     `xml:"m:SearchArchiveOnly,omitempty"`
}

type GetNonIndexableItemDetailsType struct {
	BaseRequestType   `xml:",omitempty"`
	Mailboxes         *NonEmptyArrayOfLegacyDNsType `xml:"m:Mailboxes,omitempty"`
	PageSize          XsInt                         `xml:"m:PageSize,omitempty"`
	PageItemReference XsString                      `xml:"m:PageItemReference,omitempty"`
	PageDirection     SearchPageDirectionType       `xml:"m:PageDirection,omitempty"`
	SearchArchiveOnly XsBoolean                     `xml:"m:SearchArchiveOnly,omitempty"`
}

type MarkAllItemsAsReadType struct {
	BaseRequestType      `xml:",omitempty"`
	ReadFlag             XsBoolean                         `xml:"m:ReadFlag,omitempty"`
	SuppressReadReceipts XsBoolean                         `xml:"m:SuppressReadReceipts,omitempty"`
	FolderIds            *NonEmptyArrayOfBaseFolderIdsType `xml:"m:FolderIds,omitempty"`
}

type MarkAllItemsAsReadResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type MarkAsJunkType struct {
	BaseRequestType `xml:",omitempty"`
	IsJunk          XsBoolean                       `xml:"IsJunk,attr"`
	MoveItem        XsBoolean                       `xml:"MoveItem,attr"`
	ItemIds         *NonEmptyArrayOfBaseItemIdsType `xml:"m:ItemIds,omitempty"`
}

type MarkAsJunkResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type ReportMessageType struct {
	BaseRequestType           `xml:",omitempty"`
	ReportAction              ReportMessageActionType   `xml:"ReportAction,attr"`
	ItemIds                   *ArrayOfBaseItemIdsType   `xml:"m:ItemIds,omitempty"`
	BlockReportingToMicrosoft XsBoolean                 `xml:"m:BlockReportingToMicrosoft,omitempty"`
	Platform                  ReportMessagePlatformType `xml:"m:Platform,omitempty"`
}

type ReportMessageResponseType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type GetAppManifestsType struct {
	BaseRequestType           `xml:",omitempty"`
	ApiVersionSupported       XsString                         `xml:"m:ApiVersionSupported,omitempty"`
	SchemaVersionSupported    XsString                         `xml:"m:SchemaVersionSupported,omitempty"`
	IncludeAllInstalledAddIns XsBoolean                        `xml:"m:IncludeAllInstalledAddIns,omitempty"`
	IncludeEntitlementData    XsBoolean                        `xml:"m:IncludeEntitlementData,omitempty"`
	IncludeManifestData       XsBoolean                        `xml:"m:IncludeManifestData,omitempty"`
	IncludeCustomAppsData     XsBoolean                        `xml:"m:IncludeCustomAppsData,omitempty"`
	ExtensionIds              ListOfExtensionIdsType           `xml:"m:ExtensionIds,omitempty"`
	AddIns                    *ArrayOfPrivateCatalogAddInsType `xml:"m:AddIns,omitempty"`
	IncludeExtensionMetaData  XsBoolean                        `xml:"m:IncludeExtensionMetaData,omitempty"`
}

type ArrayOfPrivateCatalogAddInsType struct {
	AddIn []*PrivateCatalogAddInsType `xml:"m:AddIn,omitempty"`
}

type PrivateCatalogAddInsType struct {
	ProductId            XsString                          `xml:"ProductId,attr"`
	State                AddInStateType                    `xml:"State,attr"`
	Version              VersionType                       `xml:"Version,attr"`
	DefaultEnabledStatus AADOfficeExtensionStatusType      `xml:"DefaultEnabledStatus,attr"`
	InstallTimeInTicks   XsLong                            `xml:"InstallTimeInTicks,attr"`
	StoreInfo            *PrivateCatalogAddInStoreInfoType `xml:"m:StoreInfo,omitempty"`
}

type PrivateCatalogAddInStoreInfoType struct {
	AssetId       XsString `xml:"AssetId,attr"`
	ContentMarket XsString `xml:"ContentMarket,attr"`
}

type GetAppManifestsResponseType struct {
	ResponseMessageType `xml:",omitempty"`
	Apps                *ArrayOfAppsType         `xml:"m:Apps,omitempty"`
	Manifests           *ArrayOfAppManifestsType `xml:"m:Manifests,omitempty"`
}

type ArrayOfAppManifestsType struct {
	Manifest []XsBase64Binary `xml:"m:Manifest,omitempty"`
}

type AddNewImContactToGroupType struct {
	BaseRequestType `xml:",omitempty"`
	ImAddress       NonEmptyStringType `xml:"m:ImAddress,omitempty"`
	DisplayName     NonEmptyStringType `xml:"m:DisplayName,omitempty"`
	GroupId         *ItemIdType        `xml:"m:GroupId,omitempty"`
}

type AddNewImContactToGroupResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	Persona             *PersonaType `xml:"m:Persona,omitempty"`
}

type AddNewTelUriContactToGroupType struct {
	BaseRequestType        `xml:",omitempty"`
	TelUriAddress          NonEmptyStringType `xml:"m:TelUriAddress,omitempty"`
	ImContactSipUriAddress NonEmptyStringType `xml:"m:ImContactSipUriAddress,omitempty"`
	ImTelephoneNumber      NonEmptyStringType `xml:"m:ImTelephoneNumber,omitempty"`
	GroupId                *ItemIdType        `xml:"m:GroupId,omitempty"`
}

type AddNewTelUriContactToGroupResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	Persona             *PersonaType `xml:"m:Persona,omitempty"`
}

type AddImContactToGroupType struct {
	BaseRequestType `xml:",omitempty"`
	ContactId       *ItemIdType `xml:"m:ContactId,omitempty"`
	GroupId         *ItemIdType `xml:"m:GroupId,omitempty"`
}

type AddImContactToGroupResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
}

type RemoveImContactFromGroupType struct {
	BaseRequestType `xml:",omitempty"`
	ContactId       *ItemIdType `xml:"m:ContactId,omitempty"`
	GroupId         *ItemIdType `xml:"m:GroupId,omitempty"`
}

type RemoveImContactFromGroupResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
}

type AddImGroupType struct {
	BaseRequestType `xml:",omitempty"`
	DisplayName     NonEmptyStringType `xml:"m:DisplayName,omitempty"`
}

type AddImGroupResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	ImGroup             *ImGroupType `xml:"m:ImGroup,omitempty"`
}

type AddDistributionGroupToImListType struct {
	BaseRequestType `xml:",omitempty"`
	SmtpAddress     NonEmptyStringType `xml:"m:SmtpAddress,omitempty"`
	DisplayName     NonEmptyStringType `xml:"m:DisplayName,omitempty"`
}

type AddDistributionGroupToImListResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	ImGroup             *ImGroupType `xml:"m:ImGroup,omitempty"`
}

type GetImItemListType struct {
	BaseRequestType    `xml:",omitempty"`
	ExtendedProperties *NonEmptyArrayOfExtendedFieldURIs `xml:"m:ExtendedProperties,omitempty"`
}

type GetImItemListResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	ImItemList          *ImItemListType `xml:"m:ImItemList,omitempty"`
}

type GetImItemsType struct {
	BaseRequestType    `xml:",omitempty"`
	ContactIds         *NonEmptyArrayOfBaseItemIdsType   `xml:"m:ContactIds,omitempty"`
	GroupIds           *NonEmptyArrayOfBaseItemIdsType   `xml:"m:GroupIds,omitempty"`
	ExtendedProperties *NonEmptyArrayOfExtendedFieldURIs `xml:"m:ExtendedProperties,omitempty"`
}

type GetImItemsResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	ImItemList          *ImItemListType `xml:"m:ImItemList,omitempty"`
}

type RemoveContactFromImListType struct {
	BaseRequestType `xml:",omitempty"`
	ContactId       *ItemIdType `xml:"m:ContactId,omitempty"`
}

type RemoveContactFromImListResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
}

type RemoveDistributionGroupFromImListType struct {
	BaseRequestType `xml:",omitempty"`
	GroupId         *ItemIdType `xml:"m:GroupId,omitempty"`
}

type RemoveDistributionGroupFromImListResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
}

type RemoveImGroupType struct {
	BaseRequestType `xml:",omitempty"`
	GroupId         *ItemIdType `xml:"m:GroupId,omitempty"`
}

type RemoveImGroupResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
}

type SetImGroupType struct {
	BaseRequestType `xml:",omitempty"`
	GroupId         *ItemIdType        `xml:"m:GroupId,omitempty"`
	NewDisplayName  NonEmptyStringType `xml:"m:NewDisplayName,omitempty"`
}

type SetImGroupResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
}

type SetImListMigrationCompletedType struct {
	BaseRequestType          `xml:",omitempty"`
	ImListMigrationCompleted XsBoolean `xml:"m:ImListMigrationCompleted,omitempty"`
}

type SetImListMigrationCompletedResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
}

type GetUserRetentionPolicyTagsType struct {
	BaseRequestType `xml:",omitempty"`
}

type InstallAppType struct {
	BaseRequestType          `xml:",omitempty"`
	Manifest                 XsBase64Binary `xml:"m:Manifest,omitempty"`
	MarketplaceAssetId       XsString       `xml:"m:MarketplaceAssetId,omitempty"`
	MarketplaceContentMarket XsString       `xml:"m:MarketplaceContentMarket,omitempty"`
	SendWelcomeEmail         XsBoolean      `xml:"m:SendWelcomeEmail,omitempty"`
	ManifestUrl              XsString       `xml:"m:ManifestUrl,omitempty"`
	MarketplaceCorrelationId XsString       `xml:"m:MarketplaceCorrelationId,omitempty"`
	CampaignId               XsString       `xml:"m:CampaignId,omitempty"`
	Id                       XsString       `xml:"m:Id,omitempty"`
	IsMetaOSApp              XsBoolean      `xml:"m:IsMetaOSApp,omitempty"`
	MetaOSSyncData           XsString       `xml:"m:MetaOSSyncData,omitempty"`
}

type InstallAppResponseType struct {
	ResponseMessageType `xml:",omitempty"`
	WasFirstInstall     XsBoolean         `xml:"m:WasFirstInstall,omitempty"`
	Extension           *InstalledAppType `xml:"m:Extension,omitempty"`
}

type UpdateExtensionUsageType struct {
	BaseRequestType `xml:",omitempty"`
	Client          XsString                             `xml:"m:Client,omitempty"`
	Extensions      *ArrayOfUpdateExtensionUsageItemType `xml:"m:Extensions,omitempty"`
}

type ArrayOfUpdateExtensionUsageItemType struct {
	ExtensionId XsString                                  `xml:"m:ExtensionId,omitempty"`
	Scenarios   *ArrayOfExtensionUsageScenarioCounterType `xml:"m:Scenarios,omitempty"`
}

type ArrayOfExtensionUsageScenarioCounterType struct {
	ScenarioName XsString `xml:"m:ScenarioName,omitempty"`
	Count        XsInt    `xml:"m:Count,omitempty"`
}

type UpdateExtensionUsageResponseType struct {
	ResponseMessageType `xml:",omitempty"`
}

type UninstallAppType struct {
	BaseRequestType `xml:",omitempty"`
	ID              XsString  `xml:"m:ID,omitempty"`
	IsMetaOSApp     XsBoolean `xml:"m:IsMetaOSApp,omitempty"`
}

type UninstallAppResponseType struct {
	ResponseMessageType `xml:",omitempty"`
}

type DisableAppType struct {
	BaseRequestType `xml:",omitempty"`
	ID              XsString          `xml:"m:ID,omitempty"`
	DisableReason   DisableReasonType `xml:"m:DisableReason,omitempty"`
	IsMetaOSApp     XsBoolean         `xml:"m:IsMetaOSApp,omitempty"`
}

type DisableAppResponseType struct {
	ResponseMessageType `xml:",omitempty"`
}

type GetAppMarketplaceUrlType struct {
	BaseRequestType `xml:",omitempty"`
}

type GetAppMarketplaceUrlResponseMessageType struct {
	ResponseMessageType     `xml:",omitempty"`
	AppMarketplaceUrl       XsString `xml:"m:AppMarketplaceUrl,omitempty"`
	ConnectorsManagementUrl XsString `xml:"m:ConnectorsManagementUrl,omitempty"`
}

type FindAvailableMeetingTimesType struct {
	BaseRequestType          `xml:",omitempty"`
	Attendees                *ArrayOfSmtpAddressType `xml:"m:Attendees,omitempty"`
	SearchWindowStart        XsDateTime              `xml:"m:SearchWindowStart,omitempty"`
	SearchWindowDuration     XsDuration              `xml:"m:SearchWindowDuration,omitempty"`
	MeetingDurationInMinutes XsInt                   `xml:"m:MeetingDurationInMinutes,omitempty"`
	Location                 XsString                `xml:"m:Location,omitempty"`
	MaxCandidates            XsInt                   `xml:"m:MaxCandidates,omitempty"`
	ActivityDomain           ActivityDomainType      `xml:"m:ActivityDomain,omitempty"`
}

type FindAvailableMeetingTimesResponseMessageType struct {
	ResponseMessageType   `xml:",omitempty"`
	MeetingTimeCandidates *ArrayOfMeetingTimeCandidate `xml:"m:MeetingTimeCandidates,omitempty"`
	EmptySuggestionsHint  EmptySuggestionReason        `xml:"m:EmptySuggestionsHint,omitempty"`
}

type FindMeetingTimeCandidatesType struct {
	BaseRequestType     `xml:",omitempty"`
	AttendeeConstraints *FindMeetingTimesAttendeeConstraints `xml:"m:AttendeeConstraints,omitempty"`
	LocationConstraints *FindMeetingTimesLocationConstraints `xml:"m:LocationConstraints,omitempty"`
	SearchConstraints   *FindMeetingTimesSearchConstraints   `xml:"m:SearchConstraints,omitempty"`
	Constraints         *FindMeetingTimesConstraints         `xml:"m:Constraints,omitempty"`
}

type FindMeetingTimeCandidatesResponseMessageType struct {
	ResponseMessageType   `xml:",omitempty"`
	MeetingTimeCandidates *ArrayOfMeetingTimeCandidate `xml:"m:MeetingTimeCandidates,omitempty"`
}

type GetUserPhotoType struct {
	BaseRequestType `xml:",omitempty"`
	Email           XsString          `xml:"m:Email,omitempty"`
	SizeRequested   UserPhotoSizeType `xml:"m:SizeRequested,omitempty"`
	TypeRequested   UserPhotoTypeType `xml:"m:TypeRequested,omitempty"`
}

type SetUserPhotoType struct {
	BaseRequestType `xml:",omitempty"`
	Email           NonEmptyStringType `xml:"m:Email,omitempty"`
	Content         XsString           `xml:"m:Content,omitempty"`
	TypeRequested   UserPhotoTypeType  `xml:"m:TypeRequested,omitempty"`
}

type SetUserPhotoResponseMessageType struct {
	BaseResponseMessageType `xml:",omitempty"`
}

type GetMeetingSpaceType struct {
	BaseRequestType `xml:",omitempty"`
	ItemId          *ItemIdType `xml:"m:ItemId,omitempty"`
}

type GetMeetingSpaceResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	MeetingSpace        *MeetingSpaceType `xml:"m:MeetingSpace,omitempty"`
}

type DeleteMeetingSpaceType struct {
	BaseRequestType `xml:",omitempty"`
	ItemId          *ItemIdType `xml:"m:ItemId,omitempty"`
}

type DeleteMeetingSpaceResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
}

type UpdateMeetingSpaceType struct {
	BaseRequestType `xml:",omitempty"`
	ItemId          *ItemIdType       `xml:"m:ItemId,omitempty"`
	MeetingSpace    *MeetingSpaceType `xml:"m:MeetingSpace,omitempty"`
}

type UpdateMeetingSpaceResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	MeetingSpace        *MeetingSpaceType `xml:"m:MeetingSpace,omitempty"`
}

type CreateMeetingSpaceType struct {
	BaseRequestType `xml:",omitempty"`
	MeetingSpace    *MeetingSpaceType `xml:"m:MeetingSpace,omitempty"`
}

type CreateMeetingSpaceResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	MeetingSpace        *MeetingSpaceType `xml:"m:MeetingSpace,omitempty"`
}

type FindMeetingSpaceByJoinUrlType struct {
	BaseRequestType `xml:",omitempty"`
	JoinUrl         XsString `xml:"m:JoinUrl,omitempty"`
}

type FindMeetingSpaceByJoinUrlResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	MeetingSpace        *MeetingSpaceType `xml:"m:MeetingSpace,omitempty"`
}

type GetMeetingInstanceRequestType struct {
	BaseRequestType `xml:",omitempty"`
	ItemId          *ItemIdType `xml:"m:ItemId,omitempty"`
}

type GetMeetingInstanceResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	MeetingInstance     *MeetingInstanceType `xml:"m:MeetingInstance,omitempty"`
}

type DeleteMeetingInstanceRequestType struct {
	BaseRequestType `xml:",omitempty"`
	ItemId          *ItemIdType `xml:"m:ItemId,omitempty"`
}

type DeleteMeetingInstanceResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
}

type UpdateMeetingInstanceRequestType struct {
	BaseRequestType            `xml:",omitempty"`
	ItemId                     *ItemIdType                           `xml:"m:ItemId,omitempty"`
	MeetingInstance            *MeetingInstanceType                  `xml:"m:MeetingInstance,omitempty"`
	ContentActivitiesToAdd     *NonEmptyArrayOfContentActivities     `xml:"m:ContentActivitiesToAdd,omitempty"`
	ParticipantActivitiesToAdd *NonEmptyArrayOfParticipantActivities `xml:"m:ParticipantActivitiesToAdd,omitempty"`
}

type UpdateMeetingInstanceResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	MeetingInstance     *MeetingInstanceType `xml:"m:MeetingInstance,omitempty"`
}

type CreateMeetingInstanceRequestType struct {
	BaseRequestType `xml:",omitempty"`
	MeetingInstance *MeetingInstanceType `xml:"m:MeetingInstance,omitempty"`
}

type CreateMeetingInstanceResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	MeetingInstance     *MeetingInstanceType `xml:"m:MeetingInstance,omitempty"`
}

type StartSearchSession struct {
	BaseRequestType `xml:",omitempty"`
	SearchSessionId GuidType                `xml:"m:SearchSessionId,omitempty"`
	WarmupOptions   WarmupOptionsType       `xml:"m:WarmupOptions,omitempty"`
	SuggestionTypes SuggestionKindType      `xml:"m:SuggestionTypes,omitempty"`
	SearchScope     *ArrayOfSearchScopeType `xml:"m:SearchScope,omitempty"`
	IdFormat        IdFormatType            `xml:"m:IdFormat,omitempty"`
	ApplicationId   XsString                `xml:"m:ApplicationId,omitempty"`
	Scenario        XsString                `xml:"m:Scenario,omitempty"`
}

type StartSearchSessionResponseMessage struct {
	ResponseMessageType `xml:",omitempty"`
}

type GetSearchSuggestions struct {
	BaseRequestType                      `xml:",omitempty"`
	SearchSessionId                      GuidType                `xml:"m:SearchSessionId,omitempty"`
	Query                                XsString                `xml:"m:Query,omitempty"`
	SuggestionTypes                      SuggestionKindType      `xml:"m:SuggestionTypes,omitempty"`
	SuggestionsPrimer                    XsBoolean               `xml:"m:SuggestionsPrimer,omitempty"`
	MaxSuggestionsCountPerSuggestionType XsLong                  `xml:"m:MaxSuggestionsCountPerSuggestionType,omitempty"`
	SearchScope                          *ArrayOfSearchScopeType `xml:"m:SearchScope,omitempty"`
	Scenario                             XsString                `xml:"m:Scenario,omitempty"`
}

type GetSearchSuggestionsResponseMessage struct {
	ResponseMessageType `xml:",omitempty"`
	SearchSuggestions   *SearchSuggestionsType `xml:"m:SearchSuggestions,omitempty"`
}

type DeleteSearchSuggestion struct {
	BaseRequestType `xml:",omitempty"`
	SearchSessionId GuidType                `xml:"m:SearchSessionId,omitempty"`
	Query           XsString                `xml:"m:Query,omitempty"`
	SuggestionTypes SuggestionKindType      `xml:"m:SuggestionTypes,omitempty"`
	SearchScope     *ArrayOfSearchScopeType `xml:"m:SearchScope,omitempty"`
	Scenario        XsString                `xml:"m:Scenario,omitempty"`
}

type DeleteSearchSuggestionResponseMessageType struct {
	ResponseMessageType `xml:",omitempty"`
	Response            *DeleteSearchSuggestionResponseType `xml:"m:Response,omitempty"`
}

type ExecuteSearch struct {
	BaseRequestType                `xml:",omitempty"`
	ApplicationId                  SearchApplicationIdType          `xml:"m:ApplicationId,omitempty"`
	Scenario                       XsString                         `xml:"m:Scenario,omitempty"`
	SearchSessionId                GuidType                         `xml:"m:SearchSessionId,omitempty"`
	SearchScope                    *ArrayOfSearchScopeType          `xml:"m:SearchScope,omitempty"`
	Query                          XsString                         `xml:"m:Query,omitempty"`
	AnalyzedQuery                  *AnalyzedQuery                   `xml:"m:AnalyzedQuery,omitempty"`
	ResultRowCount                 XsLong                           `xml:"m:ResultRowCount,omitempty"`
	ResultRowOffset                XsLong                           `xml:"m:ResultRowOffset,omitempty"`
	MaxResultsCountHint            XsLong                           `xml:"m:MaxResultsCountHint,omitempty"`
	MaxPreviewLength               XsLong                           `xml:"m:MaxPreviewLength,omitempty"`
	SearchRefiners                 *SearchRefinersTypeM             `xml:"m:SearchRefiners,omitempty"`
	ExtendedKeywords               *ExtendedKeywordsType            `xml:"m:ExtendedKeywords,omitempty"`
	RetrieveRefiners               XsBoolean                        `xml:"m:RetrieveRefiners,omitempty"`
	MaxRefinersCountPerRefinerType XsLong                           `xml:"m:MaxRefinersCountPerRefinerType,omitempty"`
	IdFormat                       IdFormatType                     `xml:"m:IdFormat,omitempty"`
	ItemTypes                      ItemTypesFilterType              `xml:"m:ItemTypes,omitempty"`
	PropertySetName                SearchResultsPropertySetNameType `xml:"m:PropertySetName,omitempty"`
	SearchRestrictions             *RestrictionType                 `xml:"m:SearchRestrictions,omitempty"`
	IncludeDeleted                 XsBoolean                        `xml:"m:IncludeDeleted,omitempty"`
	SortOrder                      ExecuteSearchSortOrderType       `xml:"m:SortOrder,omitempty"`
	KeywordMatchOption             MatchOptionsType                 `xml:"m:KeywordMatchOption,omitempty"`
	ReturnAdditionalIds            XsBoolean                        `xml:"m:ReturnAdditionalIds,omitempty"`
	RequestedProperties            *ArrayOfStringsType              `xml:"m:RequestedProperties,omitempty"`
}

type ExecuteSearchResponseMessage struct {
	ResponseMessageType `xml:",omitempty"`
	SearchResults       *SearchResultsType `xml:"m:SearchResults,omitempty"`
}

type EndSearchSession struct {
	BaseRequestType `xml:",omitempty"`
	SearchSessionId GuidType `xml:"m:SearchSessionId,omitempty"`
}

type EndSearchSessionResponseMessage struct {
	ResponseMessageType `xml:",omitempty"`
}

type GetLastPrivateCatalogUpdateType struct {
	BaseRequestType `xml:",omitempty"`
	Client          *OfficeClientType `xml:"m:Client,omitempty"`
}

type GetLastPrivateCatalogUpdateResponseType struct {
	ResponseMessageType `xml:",omitempty"`
	LastUpdate          XsDateTime `xml:"m:LastUpdate,omitempty"`
	CatalogHash         XsString   `xml:"m:CatalogHash,omitempty"`
}

type GetPrivateCatalogAddInsType struct {
	BaseRequestType `xml:",omitempty"`
	Client          *OfficeClientType `xml:"m:Client,omitempty"`
}

type GetPrivateCatalogAddInsResponseType struct {
	ResponseMessageType `xml:",omitempty"`
	AddIns              *ArrayOfPrivateCatalogAddInsType `xml:"m:AddIns,omitempty"`
}

type NonEmptyArrayOfBaseFolderIdsType struct {
	FolderId              []*FolderIdType              `xml:"t:FolderId,omitempty"`
	DistinguishedFolderId []*DistinguishedFolderIdType `xml:"t:DistinguishedFolderId,omitempty"`
}

type FolderIdType struct {
	BaseFolderIdType `xml:",omitempty"`
	Id               XsString `xml:"Id,attr"`
	ChangeKey        XsString `xml:"ChangeKey,attr"`
}

type BaseFolderIdType struct {
}

type DistinguishedFolderIdType struct {
	BaseFolderIdType `xml:",omitempty"`
	Id               DistinguishedFolderIdNameType `xml:"Id,attr"`
	ChangeKey        XsString                      `xml:"ChangeKey,attr"`
	Mailbox          *EmailAddressType             `xml:"t:Mailbox,omitempty"`
}

type EmailAddressType struct {
	BaseEmailAddressType `xml:",omitempty"`
	Name                 XsString           `xml:"t:Name,omitempty"`
	EmailAddress         NonEmptyStringType `xml:"t:EmailAddress,omitempty"`
	RoutingType          NonEmptyStringType `xml:"t:RoutingType,omitempty"`
	MailboxType          MailboxTypeType    `xml:"t:MailboxType,omitempty"`
	ItemId               *ItemIdType        `xml:"t:ItemId,omitempty"`
	OriginalDisplayName  XsString           `xml:"t:OriginalDisplayName,omitempty"`
}

type BaseEmailAddressType struct {
}

type ItemIdType struct {
	BaseItemIdType `xml:",omitempty"`
	Id             XsString `xml:"Id,attr"`
	ChangeKey      XsString `xml:"ChangeKey,attr"`
}

type BaseItemIdType struct {
}

type ExchangeImpersonationType struct {
	ConnectingSID *ConnectingSIDType `xml:"t:ConnectingSID,omitempty"`
}

type ConnectingSIDType struct {
	PrincipalName      *PrincipalNameType      `xml:"t:PrincipalName,omitempty"`
	SID                *SIDType                `xml:"t:SID,omitempty"`
	PrimarySmtpAddress *PrimarySmtpAddressType `xml:"t:PrimarySmtpAddress,omitempty"`
	SmtpAddress        *SmtpAddressType        `xml:"t:SmtpAddress,omitempty"`
}

type PrincipalNameType struct {
	CharData NonEmptyStringType `xml:",chardata"`
}

type SIDType struct {
	CharData NonEmptyStringType `xml:",chardata"`
}

type PrimarySmtpAddressType struct {
	CharData NonEmptyStringType `xml:",chardata"`
}

type SmtpAddressType struct {
	CharData NonEmptyStringType `xml:",chardata"`
}

type MailboxCultureType struct {
	CharData XsLanguage `xml:",chardata"`
}

type ArrayOfRealItemsType struct {
	Item                []*ItemType                       `xml:"t:Item,omitempty"`
	Message             []*MessageType                    `xml:"t:Message,omitempty"`
	SharingMessage      []*SharingMessageType             `xml:"t:SharingMessage,omitempty"`
	CalendarItem        []*CalendarItemType               `xml:"t:CalendarItem,omitempty"`
	Contact             []*ContactItemType                `xml:"t:Contact,omitempty"`
	DistributionList    []*DistributionListType           `xml:"t:DistributionList,omitempty"`
	MeetingMessage      []*MeetingMessageType             `xml:"t:MeetingMessage,omitempty"`
	MeetingRequest      []*MeetingRequestMessageType      `xml:"t:MeetingRequest,omitempty"`
	MeetingResponse     []*MeetingResponseMessageType     `xml:"t:MeetingResponse,omitempty"`
	MeetingCancellation []*MeetingCancellationMessageType `xml:"t:MeetingCancellation,omitempty"`
	Task                []*TaskType                       `xml:"t:Task,omitempty"`
	PostItem            []*PostItemType                   `xml:"t:PostItem,omitempty"`
	RoleMember          []*RoleMemberItemType             `xml:"t:RoleMember,omitempty"`
	Network             []*NetworkItemType                `xml:"t:Network,omitempty"`
	Person              []*AbchPersonItemType             `xml:"t:Person,omitempty"`
}

type ItemType struct {
	MimeContent                  *MimeContentType                          `xml:"t:MimeContent,omitempty"`
	ItemId                       *ItemIdType                               `xml:"t:ItemId,omitempty"`
	ParentFolderId               *FolderIdType                             `xml:"t:ParentFolderId,omitempty"`
	ItemClass                    ItemClassType                             `xml:"t:ItemClass,omitempty"`
	Subject                      XsString                                  `xml:"t:Subject,omitempty"`
	Sensitivity                  SensitivityChoicesType                    `xml:"t:Sensitivity,omitempty"`
	Body                         *BodyType                                 `xml:"t:Body,omitempty"`
	Attachments                  *NonEmptyArrayOfAttachmentsType           `xml:"t:Attachments,omitempty"`
	DateTimeReceived             XsDateTime                                `xml:"t:DateTimeReceived,omitempty"`
	Size                         XsInt                                     `xml:"t:Size,omitempty"`
	Categories                   *ArrayOfStringsType                       `xml:"t:Categories,omitempty"`
	Importance                   ImportanceChoicesType                     `xml:"t:Importance,omitempty"`
	InReplyTo                    XsString                                  `xml:"t:InReplyTo,omitempty"`
	IsSubmitted                  XsBoolean                                 `xml:"t:IsSubmitted,omitempty"`
	IsDraft                      XsBoolean                                 `xml:"t:IsDraft,omitempty"`
	IsFromMe                     XsBoolean                                 `xml:"t:IsFromMe,omitempty"`
	IsResend                     XsBoolean                                 `xml:"t:IsResend,omitempty"`
	IsUnmodified                 XsBoolean                                 `xml:"t:IsUnmodified,omitempty"`
	InternetMessageHeaders       *NonEmptyArrayOfInternetHeadersType       `xml:"t:InternetMessageHeaders,omitempty"`
	DateTimeSent                 XsDateTime                                `xml:"t:DateTimeSent,omitempty"`
	DateTimeCreated              XsDateTime                                `xml:"t:DateTimeCreated,omitempty"`
	ResponseObjects              *NonEmptyArrayOfResponseObjectsType       `xml:"t:ResponseObjects,omitempty"`
	ReminderDueBy                XsDateTime                                `xml:"t:ReminderDueBy,omitempty"`
	ReminderIsSet                XsBoolean                                 `xml:"t:ReminderIsSet,omitempty"`
	ReminderNextTime             XsDateTime                                `xml:"t:ReminderNextTime,omitempty"`
	ReminderMinutesBeforeStart   ReminderMinutesBeforeStartType            `xml:"t:ReminderMinutesBeforeStart,omitempty"`
	DisplayCc                    XsString                                  `xml:"t:DisplayCc,omitempty"`
	DisplayTo                    XsString                                  `xml:"t:DisplayTo,omitempty"`
	DisplayBcc                   XsString                                  `xml:"t:DisplayBcc,omitempty"`
	HasAttachments               XsBoolean                                 `xml:"t:HasAttachments,omitempty"`
	ExtendedProperty             []*ExtendedPropertyType                   `xml:"t:ExtendedProperty,omitempty"`
	Culture                      XsLanguage                                `xml:"t:Culture,omitempty"`
	EffectiveRights              *EffectiveRightsType                      `xml:"t:EffectiveRights,omitempty"`
	LastModifiedName             XsString                                  `xml:"t:LastModifiedName,omitempty"`
	LastModifiedTime             XsDateTime                                `xml:"t:LastModifiedTime,omitempty"`
	IsAssociated                 XsBoolean                                 `xml:"t:IsAssociated,omitempty"`
	WebClientReadFormQueryString XsString                                  `xml:"t:WebClientReadFormQueryString,omitempty"`
	WebClientEditFormQueryString XsString                                  `xml:"t:WebClientEditFormQueryString,omitempty"`
	ConversationId               *ItemIdType                               `xml:"t:ConversationId,omitempty"`
	UniqueBody                   *BodyType                                 `xml:"t:UniqueBody,omitempty"`
	Flag                         *FlagType                                 `xml:"t:Flag,omitempty"`
	StoreEntryId                 XsBase64Binary                            `xml:"t:StoreEntryId,omitempty"`
	InstanceKey                  XsBase64Binary                            `xml:"t:InstanceKey,omitempty"`
	NormalizedBody               *BodyType                                 `xml:"t:NormalizedBody,omitempty"`
	EntityExtractionResult       *EntityExtractionResultType               `xml:"t:EntityExtractionResult,omitempty"`
	PolicyTag                    *RetentionTagType                         `xml:"t:PolicyTag,omitempty"`
	ArchiveTag                   *RetentionTagType                         `xml:"t:ArchiveTag,omitempty"`
	RetentionDate                XsDateTime                                `xml:"t:RetentionDate,omitempty"`
	Preview                      XsString                                  `xml:"t:Preview,omitempty"`
	RightsManagementLicenseData  *RightsManagementLicenseDataType          `xml:"t:RightsManagementLicenseData,omitempty"`
	PredictedActionReasons       *NonEmptyArrayOfPredictedActionReasonType `xml:"t:PredictedActionReasons,omitempty"`
	IsClutter                    XsBoolean                                 `xml:"t:IsClutter,omitempty"`
	BlockStatus                  XsBoolean                                 `xml:"t:BlockStatus,omitempty"`
	HasBlockedImages             XsBoolean                                 `xml:"t:HasBlockedImages,omitempty"`
	TextBody                     *BodyType                                 `xml:"t:TextBody,omitempty"`
	IconIndex                    IconIndexType                             `xml:"t:IconIndex,omitempty"`
	SearchKey                    XsBase64Binary                            `xml:"t:SearchKey,omitempty"`
	SortKey                      XsLong                                    `xml:"t:SortKey,omitempty"`
	Hashtags                     *ArrayOfStringsType                       `xml:"t:Hashtags,omitempty"`
	Mentions                     *ArrayOfRecipientsType                    `xml:"t:Mentions,omitempty"`
	MentionedMe                  XsBoolean                                 `xml:"t:MentionedMe,omitempty"`
	MentionsPreview              *MentionsPreviewType                      `xml:"t:MentionsPreview,omitempty"`
	MentionsEx                   *NonEmptyArrayOfMentionActionsType        `xml:"t:MentionsEx,omitempty"`
	AppliedHashtags              *NonEmptyArrayOfAppliedHashtagType        `xml:"t:AppliedHashtags,omitempty"`
	AppliedHashtagsPreview       *AppliedHashtagsPreviewType               `xml:"t:AppliedHashtagsPreview,omitempty"`
	Likes                        *NonEmptyArrayOfLikeType                  `xml:"t:Likes,omitempty"`
	LikesPreview                 *LikesPreviewType                         `xml:"t:LikesPreview,omitempty"`
	PendingSocialActivityTagIds  *ArrayOfStringsType                       `xml:"t:PendingSocialActivityTagIds,omitempty"`
	AtAllMention                 XsBoolean                                 `xml:"t:AtAllMention,omitempty"`
	CanDelete                    XsBoolean                                 `xml:"t:CanDelete,omitempty"`
	InferenceClassification      InferenceClassificationType               `xml:"t:InferenceClassification,omitempty"`
}

type MimeContentType struct {
	CharData     XsString `xml:",chardata"`
	CharacterSet XsString `xml:"CharacterSet,attr"`
}

type BodyType struct {
	CharData    XsString     `xml:",chardata"`
	BodyType    BodyTypeType `xml:"BodyType,attr"`
	IsTruncated XsBoolean    `xml:"IsTruncated,attr"`
}

type NonEmptyArrayOfAttachmentsType struct {
	ItemAttachment      []*ItemAttachmentType      `xml:"t:ItemAttachment,omitempty"`
	FileAttachment      []*FileAttachmentType      `xml:"t:FileAttachment,omitempty"`
	ReferenceAttachment []*ReferenceAttachmentType `xml:"t:ReferenceAttachment,omitempty"`
}

type ItemAttachmentType struct {
	AttachmentType      `xml:",omitempty"`
	Item                *ItemType                       `xml:"t:Item,omitempty"`
	Message             *MessageType                    `xml:"t:Message,omitempty"`
	SharingMessage      *SharingMessageType             `xml:"t:SharingMessage,omitempty"`
	CalendarItem        *CalendarItemType               `xml:"t:CalendarItem,omitempty"`
	Contact             *ContactItemType                `xml:"t:Contact,omitempty"`
	MeetingMessage      *MeetingMessageType             `xml:"t:MeetingMessage,omitempty"`
	MeetingRequest      *MeetingRequestMessageType      `xml:"t:MeetingRequest,omitempty"`
	MeetingResponse     *MeetingResponseMessageType     `xml:"t:MeetingResponse,omitempty"`
	MeetingCancellation *MeetingCancellationMessageType `xml:"t:MeetingCancellation,omitempty"`
	Task                *TaskType                       `xml:"t:Task,omitempty"`
	PostItem            *PostItemType                   `xml:"t:PostItem,omitempty"`
	RoleMember          *RoleMemberItemType             `xml:"t:RoleMember,omitempty"`
	Network             *NetworkItemType                `xml:"t:Network,omitempty"`
	Person              *AbchPersonItemType             `xml:"t:Person,omitempty"`
}

type AttachmentType struct {
	AttachmentId          *AttachmentIdType `xml:"t:AttachmentId,omitempty"`
	Name                  XsString          `xml:"t:Name,omitempty"`
	ContentType           XsString          `xml:"t:ContentType,omitempty"`
	ContentId             XsString          `xml:"t:ContentId,omitempty"`
	ContentLocation       XsString          `xml:"t:ContentLocation,omitempty"`
	AttachmentOriginalUrl XsString          `xml:"t:AttachmentOriginalUrl,omitempty"`
	Size                  XsInt             `xml:"t:Size,omitempty"`
	LastModifiedTime      XsDateTime        `xml:"t:LastModifiedTime,omitempty"`
	IsInline              XsBoolean         `xml:"t:IsInline,omitempty"`
}

type AttachmentIdType struct {
	RequestAttachmentIdType `xml:",omitempty"`
	RootItemId              XsString `xml:"RootItemId,attr"`
	RootItemChangeKey       XsString `xml:"RootItemChangeKey,attr"`
}

type RequestAttachmentIdType struct {
	BaseItemIdType `xml:",omitempty"`
	Id             XsString `xml:"Id,attr"`
}

type MessageType struct {
	ItemType                   `xml:",omitempty"`
	Sender                     *SingleRecipientType     `xml:"t:Sender,omitempty"`
	ToRecipients               *ArrayOfRecipientsType   `xml:"t:ToRecipients,omitempty"`
	CcRecipients               *ArrayOfRecipientsType   `xml:"t:CcRecipients,omitempty"`
	BccRecipients              *ArrayOfRecipientsType   `xml:"t:BccRecipients,omitempty"`
	IsReadReceiptRequested     XsBoolean                `xml:"t:IsReadReceiptRequested,omitempty"`
	IsDeliveryReceiptRequested XsBoolean                `xml:"t:IsDeliveryReceiptRequested,omitempty"`
	ConversationIndex          XsBase64Binary           `xml:"t:ConversationIndex,omitempty"`
	ConversationTopic          XsString                 `xml:"t:ConversationTopic,omitempty"`
	From                       *SingleRecipientType     `xml:"t:From,omitempty"`
	InternetMessageId          XsString                 `xml:"t:InternetMessageId,omitempty"`
	IsRead                     XsBoolean                `xml:"t:IsRead,omitempty"`
	IsResponseRequested        XsBoolean                `xml:"t:IsResponseRequested,omitempty"`
	References                 XsString                 `xml:"t:References,omitempty"`
	ReplyTo                    *ArrayOfRecipientsType   `xml:"t:ReplyTo,omitempty"`
	ReceivedBy                 *SingleRecipientType     `xml:"t:ReceivedBy,omitempty"`
	ReceivedRepresenting       *SingleRecipientType     `xml:"t:ReceivedRepresenting,omitempty"`
	ApprovalRequestData        *ApprovalRequestDataType `xml:"t:ApprovalRequestData,omitempty"`
	VotingInformation          *VotingInformationType   `xml:"t:VotingInformation,omitempty"`
	ReminderMessageData        *ReminderMessageDataType `xml:"t:ReminderMessageData,omitempty"`
	MessageSafety              *MessageSafetyType       `xml:"t:MessageSafety,omitempty"`
	SenderSMTPAddress          *SmtpAddressType         `xml:"t:SenderSMTPAddress,omitempty"`
	MailboxGuids               *MailboxGuidsType        `xml:"t:MailboxGuids,omitempty"`
	PublishedCalendarItemIcs   XsString                 `xml:"t:PublishedCalendarItemIcs,omitempty"`
	PublishedCalendarItemName  XsString                 `xml:"t:PublishedCalendarItemName,omitempty"`
}

type SingleRecipientType struct {
	Mailbox *EmailAddressType `xml:"t:Mailbox,omitempty"`
}

type ArrayOfRecipientsType struct {
	Mailbox []*EmailAddressType `xml:"t:Mailbox,omitempty"`
}

type ApprovalRequestDataType struct {
	IsUndecidedApprovalRequest XsBoolean  `xml:"t:IsUndecidedApprovalRequest,omitempty"`
	ApprovalDecision           XsInt      `xml:"t:ApprovalDecision,omitempty"`
	ApprovalDecisionMaker      XsString   `xml:"t:ApprovalDecisionMaker,omitempty"`
	ApprovalDecisionTime       XsDateTime `xml:"t:ApprovalDecisionTime,omitempty"`
}

type VotingInformationType struct {
	UserOptions    *ArrayOfVotingOptionDataType `xml:"t:UserOptions,omitempty"`
	VotingResponse XsString                     `xml:"t:VotingResponse,omitempty"`
}

type ArrayOfVotingOptionDataType struct {
	VotingOptionData []*VotingOptionDataType `xml:"t:VotingOptionData,omitempty"`
}

type VotingOptionDataType struct {
	DisplayName XsString       `xml:"t:DisplayName,omitempty"`
	SendPrompt  SendPromptType `xml:"t:SendPrompt,omitempty"`
}

type ReminderMessageDataType struct {
	ReminderText             XsString    `xml:"t:ReminderText,omitempty"`
	Location                 XsString    `xml:"t:Location,omitempty"`
	StartTime                XsDateTime  `xml:"t:StartTime,omitempty"`
	EndTime                  XsDateTime  `xml:"t:EndTime,omitempty"`
	AssociatedCalendarItemId *ItemIdType `xml:"t:AssociatedCalendarItemId,omitempty"`
}

type MessageSafetyType struct {
	MessageSafetyLevel  XsInt    `xml:"t:MessageSafetyLevel,omitempty"`
	MessageSafetyReason XsInt    `xml:"t:MessageSafetyReason,omitempty"`
	Info                XsString `xml:"t:Info,omitempty"`
}

type SharingMessageType struct {
	MessageType           `xml:",omitempty"`
	SharingMessageAction  *SharingMessageActionType        `xml:"t:SharingMessageAction,omitempty"`
	SharingMessageActions *ArrayOfSharingMessageActionType `xml:"t:SharingMessageActions,omitempty"`
}

type SharingMessageActionType struct {
	Importance SharingActionImportance `xml:"t:Importance,omitempty"`
	ActionType SharingActionType       `xml:"t:ActionType,omitempty"`
	Action     SharingAction           `xml:"t:Action,omitempty"`
}

type ArrayOfSharingMessageActionType struct {
	SharingMessageAction *SharingMessageActionType `xml:"t:SharingMessageAction,omitempty"`
}

type CalendarItemType struct {
	ItemType                  `xml:",omitempty"`
	UID                       XsString                               `xml:"t:UID,omitempty"`
	RecurrenceId              XsDateTime                             `xml:"t:RecurrenceId,omitempty"`
	DateTimeStamp             XsDateTime                             `xml:"t:DateTimeStamp,omitempty"`
	Start                     XsDateTime                             `xml:"t:Start,omitempty"`
	End                       XsDateTime                             `xml:"t:End,omitempty"`
	OriginalStart             XsDateTime                             `xml:"t:OriginalStart,omitempty"`
	IsAllDayEvent             XsBoolean                              `xml:"t:IsAllDayEvent,omitempty"`
	LegacyFreeBusyStatus      LegacyFreeBusyType                     `xml:"t:LegacyFreeBusyStatus,omitempty"`
	Location                  XsString                               `xml:"t:Location,omitempty"`
	When                      XsString                               `xml:"t:When,omitempty"`
	IsMeeting                 XsBoolean                              `xml:"t:IsMeeting,omitempty"`
	IsCancelled               XsBoolean                              `xml:"t:IsCancelled,omitempty"`
	IsRecurring               XsBoolean                              `xml:"t:IsRecurring,omitempty"`
	MeetingRequestWasSent     XsBoolean                              `xml:"t:MeetingRequestWasSent,omitempty"`
	IsResponseRequested       XsBoolean                              `xml:"t:IsResponseRequested,omitempty"`
	CalendarItemType          CalendarItemTypeType                   `xml:"t:CalendarItemType,omitempty"`
	MyResponseType            ResponseTypeType                       `xml:"t:MyResponseType,omitempty"`
	Organizer                 *SingleRecipientType                   `xml:"t:Organizer,omitempty"`
	RequiredAttendees         *NonEmptyArrayOfAttendeesType          `xml:"t:RequiredAttendees,omitempty"`
	OptionalAttendees         *NonEmptyArrayOfAttendeesType          `xml:"t:OptionalAttendees,omitempty"`
	Resources                 *NonEmptyArrayOfAttendeesType          `xml:"t:Resources,omitempty"`
	InboxReminders            *ArrayOfInboxReminderType              `xml:"t:InboxReminders,omitempty"`
	ConflictingMeetingCount   XsInt                                  `xml:"t:ConflictingMeetingCount,omitempty"`
	AdjacentMeetingCount      XsInt                                  `xml:"t:AdjacentMeetingCount,omitempty"`
	ConflictingMeetings       *NonEmptyArrayOfAllItemsType           `xml:"t:ConflictingMeetings,omitempty"`
	AdjacentMeetings          *NonEmptyArrayOfAllItemsType           `xml:"t:AdjacentMeetings,omitempty"`
	Duration                  XsString                               `xml:"t:Duration,omitempty"`
	TimeZone                  XsString                               `xml:"t:TimeZone,omitempty"`
	AppointmentReplyTime      XsDateTime                             `xml:"t:AppointmentReplyTime,omitempty"`
	AppointmentSequenceNumber XsInt                                  `xml:"t:AppointmentSequenceNumber,omitempty"`
	AppointmentState          XsInt                                  `xml:"t:AppointmentState,omitempty"`
	Recurrence                *RecurrenceType                        `xml:"t:Recurrence,omitempty"`
	FirstOccurrence           *OccurrenceInfoType                    `xml:"t:FirstOccurrence,omitempty"`
	LastOccurrence            *OccurrenceInfoType                    `xml:"t:LastOccurrence,omitempty"`
	ModifiedOccurrences       *NonEmptyArrayOfOccurrenceInfoType     `xml:"t:ModifiedOccurrences,omitempty"`
	DeletedOccurrences        *NonEmptyArrayOfDeletedOccurrencesType `xml:"t:DeletedOccurrences,omitempty"`
	MeetingTimeZone           *TimeZoneType                          `xml:"t:MeetingTimeZone,omitempty"`
	StartTimeZone             *TimeZoneDefinitionType                `xml:"t:StartTimeZone,omitempty"`
	EndTimeZone               *TimeZoneDefinitionType                `xml:"t:EndTimeZone,omitempty"`
	ConferenceType            XsInt                                  `xml:"t:ConferenceType,omitempty"`
	AllowNewTimeProposal      XsBoolean                              `xml:"t:AllowNewTimeProposal,omitempty"`
	IsOnlineMeeting           XsBoolean                              `xml:"t:IsOnlineMeeting,omitempty"`
	MeetingWorkspaceUrl       XsString                               `xml:"t:MeetingWorkspaceUrl,omitempty"`
	NetShowUrl                XsString                               `xml:"t:NetShowUrl,omitempty"`
	EnhancedLocation          *EnhancedLocationType                  `xml:"t:EnhancedLocation,omitempty"`
	StartWallClock            XsDateTime                             `xml:"t:StartWallClock,omitempty"`
	EndWallClock              XsDateTime                             `xml:"t:EndWallClock,omitempty"`
	StartTimeZoneId           XsString                               `xml:"t:StartTimeZoneId,omitempty"`
	EndTimeZoneId             XsString                               `xml:"t:EndTimeZoneId,omitempty"`
	IntendedFreeBusyStatus    LegacyFreeBusyType                     `xml:"t:IntendedFreeBusyStatus,omitempty"`
	JoinOnlineMeetingUrl      XsString                               `xml:"t:JoinOnlineMeetingUrl,omitempty"`
	OnlineMeetingSettings     *OnlineMeetingSettingsType             `xml:"t:OnlineMeetingSettings,omitempty"`
	IsOrganizer               XsBoolean                              `xml:"t:IsOrganizer,omitempty"`
	CalendarActivityData      *CalendarActivityDataType              `xml:"t:CalendarActivityData,omitempty"`
	DoNotForwardMeeting       XsBoolean                              `xml:"t:DoNotForwardMeeting,omitempty"`
}

type NonEmptyArrayOfAttendeesType struct {
	Attendee []*AttendeeType `xml:"t:Attendee,omitempty"`
}

type AttendeeType struct {
	Mailbox          *EmailAddressType `xml:"t:Mailbox,omitempty"`
	ResponseType     ResponseTypeType  `xml:"t:ResponseType,omitempty"`
	LastResponseTime XsDateTime        `xml:"t:LastResponseTime,omitempty"`
	ProposedStart    XsDateTime        `xml:"t:ProposedStart,omitempty"`
	ProposedEnd      XsDateTime        `xml:"t:ProposedEnd,omitempty"`
}

type ArrayOfInboxReminderType struct {
	InboxReminder []*InboxReminderType `xml:"t:InboxReminder,omitempty"`
}

type InboxReminderType struct {
	Id                  GuidType                `xml:"t:Id,omitempty"`
	ReminderOffset      XsInt                   `xml:"t:ReminderOffset,omitempty"`
	Message             XsString                `xml:"t:Message,omitempty"`
	IsOrganizerReminder XsBoolean               `xml:"t:IsOrganizerReminder,omitempty"`
	OccurrenceChange    EmailReminderChangeType `xml:"t:OccurrenceChange,omitempty"`
	SendOption          EmailReminderSendOption `xml:"t:SendOption,omitempty"`
	Subject             XsString                `xml:"t:Subject,omitempty"`
	RichTextMessage     XsString                `xml:"t:RichTextMessage,omitempty"`
}

type NonEmptyArrayOfAllItemsType struct {
	Item                    []*ItemType                       `xml:"t:Item,omitempty"`
	Message                 []*MessageType                    `xml:"t:Message,omitempty"`
	SharingMessage          []*SharingMessageType             `xml:"t:SharingMessage,omitempty"`
	CalendarItem            []*CalendarItemType               `xml:"t:CalendarItem,omitempty"`
	Contact                 []*ContactItemType                `xml:"t:Contact,omitempty"`
	DistributionList        []*DistributionListType           `xml:"t:DistributionList,omitempty"`
	MeetingMessage          []*MeetingMessageType             `xml:"t:MeetingMessage,omitempty"`
	MeetingRequest          []*MeetingRequestMessageType      `xml:"t:MeetingRequest,omitempty"`
	MeetingResponse         []*MeetingResponseMessageType     `xml:"t:MeetingResponse,omitempty"`
	MeetingCancellation     []*MeetingCancellationMessageType `xml:"t:MeetingCancellation,omitempty"`
	Task                    []*TaskType                       `xml:"t:Task,omitempty"`
	PostItem                []*PostItemType                   `xml:"t:PostItem,omitempty"`
	ReplyToItem             []*ReplyToItemType                `xml:"t:ReplyToItem,omitempty"`
	ForwardItem             []*ForwardItemType                `xml:"t:ForwardItem,omitempty"`
	ReplyAllToItem          []*ReplyAllToItemType             `xml:"t:ReplyAllToItem,omitempty"`
	AcceptItem              []*AcceptItemType                 `xml:"t:AcceptItem,omitempty"`
	TentativelyAcceptItem   []*TentativelyAcceptItemType      `xml:"t:TentativelyAcceptItem,omitempty"`
	DeclineItem             []*DeclineItemType                `xml:"t:DeclineItem,omitempty"`
	CancelCalendarItem      []*CancelCalendarItemType         `xml:"t:CancelCalendarItem,omitempty"`
	RemoveItem              []*RemoveItemType                 `xml:"t:RemoveItem,omitempty"`
	SuppressReadReceipt     []*SuppressReadReceiptType        `xml:"t:SuppressReadReceipt,omitempty"`
	PostReplyItem           []*PostReplyItemType              `xml:"t:PostReplyItem,omitempty"`
	AcceptSharingInvitation []*AcceptSharingInvitationType    `xml:"t:AcceptSharingInvitation,omitempty"`
	RoleMember              []*RoleMemberItemType             `xml:"t:RoleMember,omitempty"`
	Network                 []*NetworkItemType                `xml:"t:Network,omitempty"`
	Person                  []*AbchPersonItemType             `xml:"t:Person,omitempty"`
}

type ContactItemType struct {
	ItemType                        `xml:",omitempty"`
	FileAs                          XsString                        `xml:"t:FileAs,omitempty"`
	FileAsMapping                   FileAsMappingType               `xml:"t:FileAsMapping,omitempty"`
	DisplayName                     XsString                        `xml:"t:DisplayName,omitempty"`
	GivenName                       XsString                        `xml:"t:GivenName,omitempty"`
	Initials                        XsString                        `xml:"t:Initials,omitempty"`
	MiddleName                      XsString                        `xml:"t:MiddleName,omitempty"`
	Nickname                        XsString                        `xml:"t:Nickname,omitempty"`
	CompleteName                    *CompleteNameType               `xml:"t:CompleteName,omitempty"`
	CompanyName                     XsString                        `xml:"t:CompanyName,omitempty"`
	EmailAddresses                  *EmailAddressDictionaryType     `xml:"t:EmailAddresses,omitempty"`
	AbchEmailAddresses              *AbchEmailAddressDictionaryType `xml:"t:AbchEmailAddresses,omitempty"`
	PhysicalAddresses               *PhysicalAddressDictionaryType  `xml:"t:PhysicalAddresses,omitempty"`
	PhoneNumbers                    *PhoneNumberDictionaryType      `xml:"t:PhoneNumbers,omitempty"`
	AssistantName                   XsString                        `xml:"t:AssistantName,omitempty"`
	Birthday                        XsDateTime                      `xml:"t:Birthday,omitempty"`
	BusinessHomePage                XsAnyURI                        `xml:"t:BusinessHomePage,omitempty"`
	Children                        *ArrayOfStringsType             `xml:"t:Children,omitempty"`
	Companies                       *ArrayOfStringsType             `xml:"t:Companies,omitempty"`
	ContactSource                   ContactSourceType               `xml:"t:ContactSource,omitempty"`
	Department                      XsString                        `xml:"t:Department,omitempty"`
	Generation                      XsString                        `xml:"t:Generation,omitempty"`
	ImAddresses                     *ImAddressDictionaryType        `xml:"t:ImAddresses,omitempty"`
	JobTitle                        XsString                        `xml:"t:JobTitle,omitempty"`
	Manager                         XsString                        `xml:"t:Manager,omitempty"`
	Mileage                         XsString                        `xml:"t:Mileage,omitempty"`
	OfficeLocation                  XsString                        `xml:"t:OfficeLocation,omitempty"`
	PostalAddressIndex              PhysicalAddressIndexType        `xml:"t:PostalAddressIndex,omitempty"`
	Profession                      XsString                        `xml:"t:Profession,omitempty"`
	SpouseName                      XsString                        `xml:"t:SpouseName,omitempty"`
	Surname                         XsString                        `xml:"t:Surname,omitempty"`
	WeddingAnniversary              XsDateTime                      `xml:"t:WeddingAnniversary,omitempty"`
	HasPicture                      XsBoolean                       `xml:"t:HasPicture,omitempty"`
	PhoneticFullName                XsString                        `xml:"t:PhoneticFullName,omitempty"`
	PhoneticFirstName               XsString                        `xml:"t:PhoneticFirstName,omitempty"`
	PhoneticLastName                XsString                        `xml:"t:PhoneticLastName,omitempty"`
	Alias                           XsString                        `xml:"t:Alias,omitempty"`
	Notes                           XsString                        `xml:"t:Notes,omitempty"`
	Photo                           XsBase64Binary                  `xml:"t:Photo,omitempty"`
	UserSMIMECertificate            *ArrayOfBinaryType              `xml:"t:UserSMIMECertificate,omitempty"`
	MSExchangeCertificate           *ArrayOfBinaryType              `xml:"t:MSExchangeCertificate,omitempty"`
	DirectoryId                     XsString                        `xml:"t:DirectoryId,omitempty"`
	ManagerMailbox                  *SingleRecipientType            `xml:"t:ManagerMailbox,omitempty"`
	DirectReports                   *ArrayOfRecipientsType          `xml:"t:DirectReports,omitempty"`
	AccountName                     XsString                        `xml:"t:AccountName,omitempty"`
	IsAutoUpdateDisabled            XsBoolean                       `xml:"t:IsAutoUpdateDisabled,omitempty"`
	IsMessengerEnabled              XsBoolean                       `xml:"t:IsMessengerEnabled,omitempty"`
	Comment                         XsString                        `xml:"t:Comment,omitempty"`
	ContactShortId                  XsInt                           `xml:"t:ContactShortId,omitempty"`
	ContactType                     XsString                        `xml:"t:ContactType,omitempty"`
	Gender                          XsString                        `xml:"t:Gender,omitempty"`
	IsHidden                        XsBoolean                       `xml:"t:IsHidden,omitempty"`
	ObjectId                        XsString                        `xml:"t:ObjectId,omitempty"`
	PassportId                      XsLong                          `xml:"t:PassportId,omitempty"`
	IsPrivate                       XsBoolean                       `xml:"t:IsPrivate,omitempty"`
	SourceId                        XsString                        `xml:"t:SourceId,omitempty"`
	TrustLevel                      XsInt                           `xml:"t:TrustLevel,omitempty"`
	CreatedBy                       XsString                        `xml:"t:CreatedBy,omitempty"`
	Urls                            *ContactUrlDictionaryType       `xml:"t:Urls,omitempty"`
	Cid                             XsLong                          `xml:"t:Cid,omitempty"`
	SkypeAuthCertificate            XsString                        `xml:"t:SkypeAuthCertificate,omitempty"`
	SkypeContext                    XsString                        `xml:"t:SkypeContext,omitempty"`
	SkypeId                         XsString                        `xml:"t:SkypeId,omitempty"`
	SkypeRelationship               XsString                        `xml:"t:SkypeRelationship,omitempty"`
	YomiNickname                    XsString                        `xml:"t:YomiNickname,omitempty"`
	XboxLiveTag                     XsString                        `xml:"t:XboxLiveTag,omitempty"`
	InviteFree                      XsBoolean                       `xml:"t:InviteFree,omitempty"`
	HidePresenceAndProfile          XsBoolean                       `xml:"t:HidePresenceAndProfile,omitempty"`
	IsPendingOutbound               XsBoolean                       `xml:"t:IsPendingOutbound,omitempty"`
	SupportGroupFeeds               XsBoolean                       `xml:"t:SupportGroupFeeds,omitempty"`
	UserTileHash                    XsString                        `xml:"t:UserTileHash,omitempty"`
	UnifiedInbox                    XsBoolean                       `xml:"t:UnifiedInbox,omitempty"`
	Mris                            *ArrayOfStringsType             `xml:"t:Mris,omitempty"`
	Wlid                            XsString                        `xml:"t:Wlid,omitempty"`
	AbchContactId                   GuidType                        `xml:"t:AbchContactId,omitempty"`
	NotInBirthdayCalendar           XsBoolean                       `xml:"t:NotInBirthdayCalendar,omitempty"`
	ShellContactType                XsString                        `xml:"t:ShellContactType,omitempty"`
	ImMri                           XsString                        `xml:"t:ImMri,omitempty"`
	PresenceTrustLevel              XsInt                           `xml:"t:PresenceTrustLevel,omitempty"`
	OtherMri                        XsString                        `xml:"t:OtherMri,omitempty"`
	ProfileLastChanged              XsString                        `xml:"t:ProfileLastChanged,omitempty"`
	MobileIMEnabled                 XsBoolean                       `xml:"t:MobileIMEnabled,omitempty"`
	PartnerNetworkProfilePhotoUrl   XsString                        `xml:"t:PartnerNetworkProfilePhotoUrl,omitempty"`
	PartnerNetworkThumbnailPhotoUrl XsString                        `xml:"t:PartnerNetworkThumbnailPhotoUrl,omitempty"`
	PersonId                        XsString                        `xml:"t:PersonId,omitempty"`
	ConversationGuid                GuidType                        `xml:"t:ConversationGuid,omitempty"`
}

type CompleteNameType struct {
	Title         XsString `xml:"t:Title,omitempty"`
	FirstName     XsString `xml:"t:FirstName,omitempty"`
	MiddleName    XsString `xml:"t:MiddleName,omitempty"`
	LastName      XsString `xml:"t:LastName,omitempty"`
	Suffix        XsString `xml:"t:Suffix,omitempty"`
	Initials      XsString `xml:"t:Initials,omitempty"`
	FullName      XsString `xml:"t:FullName,omitempty"`
	Nickname      XsString `xml:"t:Nickname,omitempty"`
	YomiFirstName XsString `xml:"t:YomiFirstName,omitempty"`
	YomiLastName  XsString `xml:"t:YomiLastName,omitempty"`
}

type EmailAddressDictionaryType struct {
	Entry []*EmailAddressDictionaryEntryType `xml:"t:Entry,omitempty"`
}

type EmailAddressDictionaryEntryType struct {
	CharData    XsString            `xml:",chardata"`
	Key         EmailAddressKeyType `xml:"Key,attr"`
	Name        XsString            `xml:"Name,attr"`
	RoutingType XsString            `xml:"RoutingType,attr"`
	MailboxType MailboxTypeType     `xml:"MailboxType,attr"`
}

type AbchEmailAddressDictionaryType struct {
	Email []*AbchEmailAddressDictionaryEntryType `xml:"t:Email,omitempty"`
}

type AbchEmailAddressDictionaryEntryType struct {
	Type               AbchEmailAddressTypeType `xml:"t:Type,omitempty"`
	Address            XsString                 `xml:"t:Address,omitempty"`
	IsMessengerEnabled XsBoolean                `xml:"t:IsMessengerEnabled,omitempty"`
	Capabilities       XsLong                   `xml:"t:Capabilities,omitempty"`
}

type PhysicalAddressDictionaryType struct {
	Entry []*PhysicalAddressDictionaryEntryType `xml:"t:Entry,omitempty"`
}

type PhysicalAddressDictionaryEntryType struct {
	Key             PhysicalAddressKeyType `xml:"Key,attr"`
	Street          XsString               `xml:"t:Street,omitempty"`
	City            XsString               `xml:"t:City,omitempty"`
	State           XsString               `xml:"t:State,omitempty"`
	CountryOrRegion XsString               `xml:"t:CountryOrRegion,omitempty"`
	PostalCode      XsString               `xml:"t:PostalCode,omitempty"`
}

type PhoneNumberDictionaryType struct {
	Entry []*PhoneNumberDictionaryEntryType `xml:"t:Entry,omitempty"`
}

type PhoneNumberDictionaryEntryType struct {
	CharData XsString           `xml:",chardata"`
	Key      PhoneNumberKeyType `xml:"Key,attr"`
}

type ArrayOfStringsType struct {
	String []XsString `xml:"t:String,omitempty"`
}

type ImAddressDictionaryType struct {
	Entry []*ImAddressDictionaryEntryType `xml:"t:Entry,omitempty"`
}

type ImAddressDictionaryEntryType struct {
	CharData XsString         `xml:",chardata"`
	Key      ImAddressKeyType `xml:"Key,attr"`
}

type ArrayOfBinaryType struct {
	Base64Binary []XsBase64Binary `xml:"t:Base64Binary,omitempty"`
}

type ContactUrlDictionaryType struct {
	Url []*ContactUrlDictionaryEntryType `xml:"t:Url,omitempty"`
}

type ContactUrlDictionaryEntryType struct {
	Type    ContactUrlKeyType `xml:"t:Type,omitempty"`
	Name    XsString          `xml:"t:Name,omitempty"`
	Address XsString          `xml:"t:Address,omitempty"`
}

type DistributionListType struct {
	ItemType      `xml:",omitempty"`
	DisplayName   XsString          `xml:"t:DisplayName,omitempty"`
	FileAs        XsString          `xml:"t:FileAs,omitempty"`
	ContactSource ContactSourceType `xml:"t:ContactSource,omitempty"`
	Members       *MembersListType  `xml:"t:Members,omitempty"`
}

type MembersListType struct {
	Member []*MemberType `xml:"t:Member,omitempty"`
}

type MemberType struct {
	Key     XsString          `xml:"Key,attr"`
	Mailbox *EmailAddressType `xml:"t:Mailbox,omitempty"`
	Status  MemberStatusType  `xml:"t:Status,omitempty"`
}

type MeetingMessageType struct {
	MessageType              `xml:",omitempty"`
	AssociatedCalendarItemId *ItemIdType      `xml:"t:AssociatedCalendarItemId,omitempty"`
	IsDelegated              XsBoolean        `xml:"t:IsDelegated,omitempty"`
	IsOutOfDate              XsBoolean        `xml:"t:IsOutOfDate,omitempty"`
	HasBeenProcessed         XsBoolean        `xml:"t:HasBeenProcessed,omitempty"`
	ResponseType             ResponseTypeType `xml:"t:ResponseType,omitempty"`
	UID                      XsString         `xml:"t:UID,omitempty"`
	RecurrenceId             XsDateTime       `xml:"t:RecurrenceId,omitempty"`
	DateTimeStamp            XsDateTime       `xml:"t:DateTimeStamp,omitempty"`
	IsOrganizer              XsBoolean        `xml:"t:IsOrganizer,omitempty"`
}

type MeetingRequestMessageType struct {
	MeetingMessageType        `xml:",omitempty"`
	MeetingRequestType        MeetingRequestTypeType                 `xml:"t:MeetingRequestType,omitempty"`
	IntendedFreeBusyStatus    LegacyFreeBusyType                     `xml:"t:IntendedFreeBusyStatus,omitempty"`
	Start                     XsDateTime                             `xml:"t:Start,omitempty"`
	End                       XsDateTime                             `xml:"t:End,omitempty"`
	OriginalStart             XsDateTime                             `xml:"t:OriginalStart,omitempty"`
	IsAllDayEvent             XsBoolean                              `xml:"t:IsAllDayEvent,omitempty"`
	LegacyFreeBusyStatus      LegacyFreeBusyType                     `xml:"t:LegacyFreeBusyStatus,omitempty"`
	Location                  XsString                               `xml:"t:Location,omitempty"`
	When                      XsString                               `xml:"t:When,omitempty"`
	IsMeeting                 XsBoolean                              `xml:"t:IsMeeting,omitempty"`
	IsCancelled               XsBoolean                              `xml:"t:IsCancelled,omitempty"`
	IsRecurring               XsBoolean                              `xml:"t:IsRecurring,omitempty"`
	MeetingRequestWasSent     XsBoolean                              `xml:"t:MeetingRequestWasSent,omitempty"`
	CalendarItemType          CalendarItemTypeType                   `xml:"t:CalendarItemType,omitempty"`
	MyResponseType            ResponseTypeType                       `xml:"t:MyResponseType,omitempty"`
	Organizer                 *SingleRecipientType                   `xml:"t:Organizer,omitempty"`
	RequiredAttendees         *NonEmptyArrayOfAttendeesType          `xml:"t:RequiredAttendees,omitempty"`
	OptionalAttendees         *NonEmptyArrayOfAttendeesType          `xml:"t:OptionalAttendees,omitempty"`
	Resources                 *NonEmptyArrayOfAttendeesType          `xml:"t:Resources,omitempty"`
	ConflictingMeetingCount   XsInt                                  `xml:"t:ConflictingMeetingCount,omitempty"`
	AdjacentMeetingCount      XsInt                                  `xml:"t:AdjacentMeetingCount,omitempty"`
	ConflictingMeetings       *NonEmptyArrayOfAllItemsType           `xml:"t:ConflictingMeetings,omitempty"`
	AdjacentMeetings          *NonEmptyArrayOfAllItemsType           `xml:"t:AdjacentMeetings,omitempty"`
	Duration                  XsString                               `xml:"t:Duration,omitempty"`
	TimeZone                  XsString                               `xml:"t:TimeZone,omitempty"`
	AppointmentReplyTime      XsDateTime                             `xml:"t:AppointmentReplyTime,omitempty"`
	AppointmentSequenceNumber XsInt                                  `xml:"t:AppointmentSequenceNumber,omitempty"`
	AppointmentState          XsInt                                  `xml:"t:AppointmentState,omitempty"`
	Recurrence                *RecurrenceType                        `xml:"t:Recurrence,omitempty"`
	FirstOccurrence           *OccurrenceInfoType                    `xml:"t:FirstOccurrence,omitempty"`
	LastOccurrence            *OccurrenceInfoType                    `xml:"t:LastOccurrence,omitempty"`
	ModifiedOccurrences       *NonEmptyArrayOfOccurrenceInfoType     `xml:"t:ModifiedOccurrences,omitempty"`
	DeletedOccurrences        *NonEmptyArrayOfDeletedOccurrencesType `xml:"t:DeletedOccurrences,omitempty"`
	MeetingTimeZone           *TimeZoneType                          `xml:"t:MeetingTimeZone,omitempty"`
	StartTimeZone             *TimeZoneDefinitionType                `xml:"t:StartTimeZone,omitempty"`
	EndTimeZone               *TimeZoneDefinitionType                `xml:"t:EndTimeZone,omitempty"`
	ConferenceType            XsInt                                  `xml:"t:ConferenceType,omitempty"`
	AllowNewTimeProposal      XsBoolean                              `xml:"t:AllowNewTimeProposal,omitempty"`
	IsOnlineMeeting           XsBoolean                              `xml:"t:IsOnlineMeeting,omitempty"`
	MeetingWorkspaceUrl       XsString                               `xml:"t:MeetingWorkspaceUrl,omitempty"`
	NetShowUrl                XsString                               `xml:"t:NetShowUrl,omitempty"`
	EnhancedLocation          *EnhancedLocationType                  `xml:"t:EnhancedLocation,omitempty"`
	ChangeHighlights          *ChangeHighlightsType                  `xml:"t:ChangeHighlights,omitempty"`
	StartWallClock            XsDateTime                             `xml:"t:StartWallClock,omitempty"`
	EndWallClock              XsDateTime                             `xml:"t:EndWallClock,omitempty"`
	StartTimeZoneId           XsString                               `xml:"t:StartTimeZoneId,omitempty"`
	EndTimeZoneId             XsString                               `xml:"t:EndTimeZoneId,omitempty"`
	DoNotForwardMeeting       XsBoolean                              `xml:"t:DoNotForwardMeeting,omitempty"`
}

type RecurrenceType struct {
	RelativeYearlyRecurrence  *RelativeYearlyRecurrencePatternType  `xml:"t:RelativeYearlyRecurrence,omitempty"`
	AbsoluteYearlyRecurrence  *AbsoluteYearlyRecurrencePatternType  `xml:"t:AbsoluteYearlyRecurrence,omitempty"`
	RelativeMonthlyRecurrence *RelativeMonthlyRecurrencePatternType `xml:"t:RelativeMonthlyRecurrence,omitempty"`
	AbsoluteMonthlyRecurrence *AbsoluteMonthlyRecurrencePatternType `xml:"t:AbsoluteMonthlyRecurrence,omitempty"`
	WeeklyRecurrence          *WeeklyRecurrencePatternType          `xml:"t:WeeklyRecurrence,omitempty"`
	DailyRecurrence           *DailyRecurrencePatternType           `xml:"t:DailyRecurrence,omitempty"`
	NoEndRecurrence           *NoEndRecurrenceRangeType             `xml:"t:NoEndRecurrence,omitempty"`
	EndDateRecurrence         *EndDateRecurrenceRangeType           `xml:"t:EndDateRecurrence,omitempty"`
	NumberedRecurrence        *NumberedRecurrenceRangeType          `xml:"t:NumberedRecurrence,omitempty"`
}

type RelativeYearlyRecurrencePatternType struct {
	RecurrencePatternBaseType `xml:",omitempty"`
	DaysOfWeek                DayOfWeekType      `xml:"t:DaysOfWeek,omitempty"`
	DayOfWeekIndex            DayOfWeekIndexType `xml:"t:DayOfWeekIndex,omitempty"`
	Month                     MonthNamesType     `xml:"t:Month,omitempty"`
}

type RecurrencePatternBaseType struct {
}

type AbsoluteYearlyRecurrencePatternType struct {
	RecurrencePatternBaseType `xml:",omitempty"`
	DayOfMonth                XsInt          `xml:"t:DayOfMonth,omitempty"`
	Month                     MonthNamesType `xml:"t:Month,omitempty"`
}

type RelativeMonthlyRecurrencePatternType struct {
	IntervalRecurrencePatternBaseType `xml:",omitempty"`
	DaysOfWeek                        DayOfWeekType      `xml:"t:DaysOfWeek,omitempty"`
	DayOfWeekIndex                    DayOfWeekIndexType `xml:"t:DayOfWeekIndex,omitempty"`
}

type IntervalRecurrencePatternBaseType struct {
	RecurrencePatternBaseType `xml:",omitempty"`
	Interval                  XsInt `xml:"t:Interval,omitempty"`
}

type AbsoluteMonthlyRecurrencePatternType struct {
	IntervalRecurrencePatternBaseType `xml:",omitempty"`
	DayOfMonth                        XsInt `xml:"t:DayOfMonth,omitempty"`
}

type WeeklyRecurrencePatternType struct {
	IntervalRecurrencePatternBaseType `xml:",omitempty"`
	DaysOfWeek                        DaysOfWeekType `xml:"t:DaysOfWeek,omitempty"`
	FirstDayOfWeek                    DayOfWeekType  `xml:"t:FirstDayOfWeek,omitempty"`
}

type DailyRecurrencePatternType struct {
	IntervalRecurrencePatternBaseType `xml:",omitempty"`
}

type NoEndRecurrenceRangeType struct {
	RecurrenceRangeBaseType `xml:",omitempty"`
}

type RecurrenceRangeBaseType struct {
	StartDate XsDate `xml:"t:StartDate,omitempty"`
}

type EndDateRecurrenceRangeType struct {
	RecurrenceRangeBaseType `xml:",omitempty"`
	EndDate                 XsDate `xml:"t:EndDate,omitempty"`
}

type NumberedRecurrenceRangeType struct {
	RecurrenceRangeBaseType `xml:",omitempty"`
	NumberOfOccurrences     XsInt `xml:"t:NumberOfOccurrences,omitempty"`
}

type OccurrenceInfoType struct {
	ItemId        *ItemIdType `xml:"t:ItemId,omitempty"`
	Start         XsDateTime  `xml:"t:Start,omitempty"`
	End           XsDateTime  `xml:"t:End,omitempty"`
	OriginalStart XsDateTime  `xml:"t:OriginalStart,omitempty"`
}

type NonEmptyArrayOfOccurrenceInfoType struct {
	Occurrence []*OccurrenceInfoType `xml:"t:Occurrence,omitempty"`
}

type NonEmptyArrayOfDeletedOccurrencesType struct {
	DeletedOccurrence []*DeletedOccurrenceInfoType `xml:"t:DeletedOccurrence,omitempty"`
}

type DeletedOccurrenceInfoType struct {
	Start XsDateTime `xml:"t:Start,omitempty"`
}

type TimeZoneType struct {
	TimeZoneName XsString        `xml:"TimeZoneName,attr"`
	BaseOffset   XsDuration      `xml:"t:BaseOffset,omitempty"`
	Standard     *TimeChangeType `xml:"t:Standard,omitempty"`
	Daylight     *TimeChangeType `xml:"t:Daylight,omitempty"`
}

type TimeChangeType struct {
	TimeZoneName             XsString                             `xml:"TimeZoneName,attr"`
	Offset                   XsDuration                           `xml:"t:Offset,omitempty"`
	Time                     XsTime                               `xml:"t:Time,omitempty"`
	RelativeYearlyRecurrence *RelativeYearlyRecurrencePatternType `xml:"t:RelativeYearlyRecurrence,omitempty"`
	AbsoluteDate             XsDate                               `xml:"t:AbsoluteDate,omitempty"`
}

type TimeZoneDefinitionType struct {
	Id                XsString                      `xml:"Id,attr"`
	Name              XsString                      `xml:"Name,attr"`
	Periods           *NonEmptyArrayOfPeriodsType   `xml:"t:Periods,omitempty"`
	TransitionsGroups *ArrayOfTransitionsGroupsType `xml:"t:TransitionsGroups,omitempty"`
	Transitions       *ArrayOfTransitionsType       `xml:"t:Transitions,omitempty"`
}

type NonEmptyArrayOfPeriodsType struct {
	Period []*PeriodType `xml:"t:Period,omitempty"`
}

type PeriodType struct {
	Bias XsDuration `xml:"Bias,attr"`
	Name XsString   `xml:"Name,attr"`
	Id   XsString   `xml:"Id,attr"`
}

type ArrayOfTransitionsGroupsType struct {
	TransitionsGroup []*ArrayOfTransitionsType `xml:"t:TransitionsGroup,omitempty"`
}

type ArrayOfTransitionsType struct {
	Id         XsString        `xml:"Id,attr"`
	Transition *TransitionType `xml:"t:Transition,omitempty"`
}

type TransitionType struct {
	To *TransitionTargetType `xml:"t:To,omitempty"`
}

type TransitionTargetType struct {
	CharData XsString                 `xml:",chardata"`
	Kind     TransitionTargetKindType `xml:"Kind,attr"`
}

type EnhancedLocationType struct {
	DisplayName   XsString                  `xml:"t:DisplayName,omitempty"`
	Annotation    XsString                  `xml:"t:Annotation,omitempty"`
	PostalAddress *PersonaPostalAddressType `xml:"t:PostalAddress,omitempty"`
}

type PersonaPostalAddressType struct {
	Street           XsString           `xml:"t:Street,omitempty"`
	City             XsString           `xml:"t:City,omitempty"`
	State            XsString           `xml:"t:State,omitempty"`
	Country          XsString           `xml:"t:Country,omitempty"`
	PostalCode       XsString           `xml:"t:PostalCode,omitempty"`
	PostOfficeBox    XsString           `xml:"t:PostOfficeBox,omitempty"`
	Type             XsString           `xml:"t:Type,omitempty"`
	Latitude         XsDouble           `xml:"t:Latitude,omitempty"`
	Longitude        XsDouble           `xml:"t:Longitude,omitempty"`
	Accuracy         XsDouble           `xml:"t:Accuracy,omitempty"`
	Altitude         XsDouble           `xml:"t:Altitude,omitempty"`
	AltitudeAccuracy XsDouble           `xml:"t:AltitudeAccuracy,omitempty"`
	FormattedAddress XsString           `xml:"t:FormattedAddress,omitempty"`
	LocationUri      XsString           `xml:"t:LocationUri,omitempty"`
	LocationSource   LocationSourceType `xml:"t:LocationSource,omitempty"`
}

type ChangeHighlightsType struct {
	HasLocationChanged  XsBoolean  `xml:"t:HasLocationChanged,omitempty"`
	Location            XsString   `xml:"t:Location,omitempty"`
	HasStartTimeChanged XsBoolean  `xml:"t:HasStartTimeChanged,omitempty"`
	Start               XsDateTime `xml:"t:Start,omitempty"`
	HasEndTimeChanged   XsBoolean  `xml:"t:HasEndTimeChanged,omitempty"`
	End                 XsDateTime `xml:"t:End,omitempty"`
}

type MeetingResponseMessageType struct {
	MeetingMessageType `xml:",omitempty"`
	Start              XsDateTime            `xml:"t:Start,omitempty"`
	End                XsDateTime            `xml:"t:End,omitempty"`
	Location           XsString              `xml:"t:Location,omitempty"`
	Recurrence         *RecurrenceType       `xml:"t:Recurrence,omitempty"`
	CalendarItemType   XsString              `xml:"t:CalendarItemType,omitempty"`
	ProposedStart      XsDateTime            `xml:"t:ProposedStart,omitempty"`
	ProposedEnd        XsDateTime            `xml:"t:ProposedEnd,omitempty"`
	EnhancedLocation   *EnhancedLocationType `xml:"t:EnhancedLocation,omitempty"`
}

type MeetingCancellationMessageType struct {
	MeetingMessageType  `xml:",omitempty"`
	Start               XsDateTime            `xml:"t:Start,omitempty"`
	End                 XsDateTime            `xml:"t:End,omitempty"`
	Location            XsString              `xml:"t:Location,omitempty"`
	Recurrence          *RecurrenceType       `xml:"t:Recurrence,omitempty"`
	CalendarItemType    XsString              `xml:"t:CalendarItemType,omitempty"`
	EnhancedLocation    *EnhancedLocationType `xml:"t:EnhancedLocation,omitempty"`
	DoNotForwardMeeting XsBoolean             `xml:"t:DoNotForwardMeeting,omitempty"`
}

type TaskType struct {
	ItemType             `xml:",omitempty"`
	ActualWork           XsInt                 `xml:"t:ActualWork,omitempty"`
	AssignedTime         XsDateTime            `xml:"t:AssignedTime,omitempty"`
	BillingInformation   XsString              `xml:"t:BillingInformation,omitempty"`
	ChangeCount          XsInt                 `xml:"t:ChangeCount,omitempty"`
	Companies            *ArrayOfStringsType   `xml:"t:Companies,omitempty"`
	CompleteDate         XsDateTime            `xml:"t:CompleteDate,omitempty"`
	Contacts             *ArrayOfStringsType   `xml:"t:Contacts,omitempty"`
	DelegationState      TaskDelegateStateType `xml:"t:DelegationState,omitempty"`
	Delegator            XsString              `xml:"t:Delegator,omitempty"`
	DueDate              XsDateTime            `xml:"t:DueDate,omitempty"`
	IsAssignmentEditable XsInt                 `xml:"t:IsAssignmentEditable,omitempty"`
	IsComplete           XsBoolean             `xml:"t:IsComplete,omitempty"`
	IsRecurring          XsBoolean             `xml:"t:IsRecurring,omitempty"`
	IsTeamTask           XsBoolean             `xml:"t:IsTeamTask,omitempty"`
	Mileage              XsString              `xml:"t:Mileage,omitempty"`
	Owner                XsString              `xml:"t:Owner,omitempty"`
	PercentComplete      XsDouble              `xml:"t:PercentComplete,omitempty"`
	Recurrence           *TaskRecurrenceType   `xml:"t:Recurrence,omitempty"`
	StartDate            XsDateTime            `xml:"t:StartDate,omitempty"`
	Status               TaskStatusType        `xml:"t:Status,omitempty"`
	StatusDescription    XsString              `xml:"t:StatusDescription,omitempty"`
	TotalWork            XsInt                 `xml:"t:TotalWork,omitempty"`
}

type TaskRecurrenceType struct {
	RelativeYearlyRecurrence  *RelativeYearlyRecurrencePatternType  `xml:"t:RelativeYearlyRecurrence,omitempty"`
	AbsoluteYearlyRecurrence  *AbsoluteYearlyRecurrencePatternType  `xml:"t:AbsoluteYearlyRecurrence,omitempty"`
	RelativeMonthlyRecurrence *RelativeMonthlyRecurrencePatternType `xml:"t:RelativeMonthlyRecurrence,omitempty"`
	AbsoluteMonthlyRecurrence *AbsoluteMonthlyRecurrencePatternType `xml:"t:AbsoluteMonthlyRecurrence,omitempty"`
	WeeklyRecurrence          *WeeklyRecurrencePatternType          `xml:"t:WeeklyRecurrence,omitempty"`
	DailyRecurrence           *DailyRecurrencePatternType           `xml:"t:DailyRecurrence,omitempty"`
	DailyRegeneration         *DailyRegeneratingPatternType         `xml:"t:DailyRegeneration,omitempty"`
	WeeklyRegeneration        *WeeklyRegeneratingPatternType        `xml:"t:WeeklyRegeneration,omitempty"`
	MonthlyRegeneration       *MonthlyRegeneratingPatternType       `xml:"t:MonthlyRegeneration,omitempty"`
	YearlyRegeneration        *YearlyRegeneratingPatternType        `xml:"t:YearlyRegeneration,omitempty"`
	NoEndRecurrence           *NoEndRecurrenceRangeType             `xml:"t:NoEndRecurrence,omitempty"`
	EndDateRecurrence         *EndDateRecurrenceRangeType           `xml:"t:EndDateRecurrence,omitempty"`
	NumberedRecurrence        *NumberedRecurrenceRangeType          `xml:"t:NumberedRecurrence,omitempty"`
}

type DailyRegeneratingPatternType struct {
	RegeneratingPatternBaseType `xml:",omitempty"`
}

type RegeneratingPatternBaseType struct {
	IntervalRecurrencePatternBaseType `xml:",omitempty"`
}

type WeeklyRegeneratingPatternType struct {
	RegeneratingPatternBaseType `xml:",omitempty"`
}

type MonthlyRegeneratingPatternType struct {
	RegeneratingPatternBaseType `xml:",omitempty"`
}

type YearlyRegeneratingPatternType struct {
	RegeneratingPatternBaseType `xml:",omitempty"`
}

type PostItemType struct {
	ItemType          `xml:",omitempty"`
	ConversationIndex XsBase64Binary       `xml:"t:ConversationIndex,omitempty"`
	ConversationTopic XsString             `xml:"t:ConversationTopic,omitempty"`
	From              *SingleRecipientType `xml:"t:From,omitempty"`
	InternetMessageId XsString             `xml:"t:InternetMessageId,omitempty"`
	IsRead            XsBoolean            `xml:"t:IsRead,omitempty"`
	PostedTime        XsDateTime           `xml:"t:PostedTime,omitempty"`
	References        XsString             `xml:"t:References,omitempty"`
	Sender            *SingleRecipientType `xml:"t:Sender,omitempty"`
}

type ReplyToItemType struct {
	SmartResponseType `xml:",omitempty"`
}

type SmartResponseType struct {
	SmartResponseBaseType `xml:",omitempty"`
	NewBodyContent        *BodyType `xml:"t:NewBodyContent,omitempty"`
}

type SmartResponseBaseType struct {
	ResponseObjectType `xml:",omitempty"`
}

type ResponseObjectType struct {
	ResponseObjectCoreType `xml:",omitempty"`
	ObjectName             XsString `xml:"ObjectName,attr"`
}

type ResponseObjectCoreType struct {
	MessageType     `xml:",omitempty"`
	ReferenceItemId *ItemIdType `xml:"t:ReferenceItemId,omitempty"`
}

type ForwardItemType struct {
	SmartResponseType `xml:",omitempty"`
}

type ReplyAllToItemType struct {
	SmartResponseType      `xml:",omitempty"`
	IsSpecificMessageReply XsBoolean `xml:"t:IsSpecificMessageReply,omitempty"`
}

type AcceptItemType struct {
	MeetingRegistrationResponseObjectType `xml:",omitempty"`
}

type MeetingRegistrationResponseObjectType struct {
	WellKnownResponseObjectType `xml:",omitempty"`
	ProposedStart               XsDateTime `xml:"t:ProposedStart,omitempty"`
	ProposedEnd                 XsDateTime `xml:"t:ProposedEnd,omitempty"`
}

type WellKnownResponseObjectType struct {
	ResponseObjectType `xml:",omitempty"`
}

type TentativelyAcceptItemType struct {
	MeetingRegistrationResponseObjectType `xml:",omitempty"`
}

type DeclineItemType struct {
	MeetingRegistrationResponseObjectType `xml:",omitempty"`
}

type CancelCalendarItemType struct {
	SmartResponseType `xml:",omitempty"`
}

type RemoveItemType struct {
	ResponseObjectType `xml:",omitempty"`
}

type SuppressReadReceiptType struct {
	ReferenceItemResponseType `xml:",omitempty"`
}

type ReferenceItemResponseType struct {
	ResponseObjectType `xml:",omitempty"`
}

type PostReplyItemType struct {
	PostReplyItemBaseType `xml:",omitempty"`
	NewBodyContent        *BodyType `xml:"t:NewBodyContent,omitempty"`
}

type PostReplyItemBaseType struct {
	ResponseObjectType `xml:",omitempty"`
}

type AcceptSharingInvitationType struct {
	ReferenceItemResponseType `xml:",omitempty"`
}

type RoleMemberItemType struct {
	ItemType    `xml:",omitempty"`
	DisplayName XsString           `xml:"t:DisplayName,omitempty"`
	Type        RoleMemberTypeType `xml:"t:Type,omitempty"`
	MemberId    XsString           `xml:"t:MemberId,omitempty"`
}

type NetworkItemType struct {
	ItemType                  `xml:",omitempty"`
	DomainId                  XsInt          `xml:"t:DomainId,omitempty"`
	DomainTag                 XsString       `xml:"t:DomainTag,omitempty"`
	UserTileUrl               XsString       `xml:"t:UserTileUrl,omitempty"`
	ProfileUrl                XsString       `xml:"t:ProfileUrl,omitempty"`
	Settings                  XsInt          `xml:"t:Settings,omitempty"`
	IsDefault                 XsBoolean      `xml:"t:IsDefault,omitempty"`
	AutoLinkError             XsString       `xml:"t:AutoLinkError,omitempty"`
	AutoLinkSuccess           XsString       `xml:"t:AutoLinkSuccess,omitempty"`
	UserEmail                 XsString       `xml:"t:UserEmail,omitempty"`
	ClientPublishSecret       XsString       `xml:"t:ClientPublishSecret,omitempty"`
	ClientToken               XsString       `xml:"t:ClientToken,omitempty"`
	ClientToken2              XsString       `xml:"t:ClientToken2,omitempty"`
	ContactSyncError          XsString       `xml:"t:ContactSyncError,omitempty"`
	ContactSyncSuccess        XsString       `xml:"t:ContactSyncSuccess,omitempty"`
	ErrorOffers               XsInt          `xml:"t:ErrorOffers,omitempty"`
	FirstAuthErrorDates       XsString       `xml:"t:FirstAuthErrorDates,omitempty"`
	LastVersionSaved          XsInt          `xml:"t:LastVersionSaved,omitempty"`
	LastWelcomeContact        XsString       `xml:"t:LastWelcomeContact,omitempty"`
	Offers                    XsInt          `xml:"t:Offers,omitempty"`
	PsaLastChanged            XsDateTime     `xml:"t:PsaLastChanged,omitempty"`
	RefreshToken2             XsString       `xml:"t:RefreshToken2,omitempty"`
	RefreshTokenExpiry2       XsString       `xml:"t:RefreshTokenExpiry2,omitempty"`
	SessionHandle             XsString       `xml:"t:SessionHandle,omitempty"`
	RejectedOffers            XsInt          `xml:"t:RejectedOffers,omitempty"`
	SyncEnabled               XsBoolean      `xml:"t:SyncEnabled,omitempty"`
	TokenRefreshLastAttempted XsDateTime     `xml:"t:TokenRefreshLastAttempted,omitempty"`
	TokenRefreshLastCompleted XsDateTime     `xml:"t:TokenRefreshLastCompleted,omitempty"`
	PsaState                  XsString       `xml:"t:PsaState,omitempty"`
	SourceEntryID             XsBase64Binary `xml:"t:SourceEntryID,omitempty"`
	AccountName               XsString       `xml:"t:AccountName,omitempty"`
	LastSync                  XsDateTime     `xml:"t:LastSync,omitempty"`
}

type AbchPersonItemType struct {
	ItemType             `xml:",omitempty"`
	AntiLinkInfo         XsString                             `xml:"t:AntiLinkInfo,omitempty"`
	PersonId             GuidType                             `xml:"t:PersonId,omitempty"`
	ContactHandles       *ArrayOfAbchPersonContactHandlesType `xml:"t:ContactHandles,omitempty"`
	ContactCategories    *ArrayOfStringsType                  `xml:"t:ContactCategories,omitempty"`
	RelevanceOrder1      XsString                             `xml:"t:RelevanceOrder1,omitempty"`
	RelevanceOrder2      XsString                             `xml:"t:RelevanceOrder2,omitempty"`
	TrustLevel           XsInt                                `xml:"t:TrustLevel,omitempty"`
	FavoriteOrder        XsInt                                `xml:"t:FavoriteOrder,omitempty"`
	ExchangePersonIdGuid GuidType                             `xml:"t:ExchangePersonIdGuid,omitempty"`
}

type ArrayOfAbchPersonContactHandlesType struct {
	ContactHandle []*AbchPersonContactHandle `xml:"t:ContactHandle,omitempty"`
}

type AbchPersonContactHandle struct {
	SourceId    XsString `xml:"t:SourceId,omitempty"`
	ObjectId    XsString `xml:"t:ObjectId,omitempty"`
	AccountName XsString `xml:"t:AccountName,omitempty"`
}

type OnlineMeetingSettingsType struct {
	LobbyBypass LobbyBypassType              `xml:"t:LobbyBypass,omitempty"`
	AccessLevel OnlineMeetingAccessLevelType `xml:"t:AccessLevel,omitempty"`
	Presenters  PresentersType               `xml:"t:Presenters,omitempty"`
}

type CalendarActivityDataType struct {
	ActivityAction XsString `xml:"t:ActivityAction,omitempty"`
	ClientId       XsString `xml:"t:ClientId,omitempty"`
	CasRequestId   GuidType `xml:"t:CasRequestId,omitempty"`
	IndexSelected  XsInt    `xml:"t:IndexSelected,omitempty"`
}

type FileAttachmentType struct {
	AttachmentType `xml:",omitempty"`
	IsContactPhoto XsBoolean      `xml:"t:IsContactPhoto,omitempty"`
	Content        XsBase64Binary `xml:"t:Content,omitempty"`
}

type ReferenceAttachmentType struct {
	AttachmentType         `xml:",omitempty"`
	AttachLongPathName     XsString  `xml:"t:AttachLongPathName,omitempty"`
	ProviderType           XsString  `xml:"t:ProviderType,omitempty"`
	ProviderEndpointUrl    XsString  `xml:"t:ProviderEndpointUrl,omitempty"`
	AttachmentThumbnailUrl XsString  `xml:"t:AttachmentThumbnailUrl,omitempty"`
	AttachmentPreviewUrl   XsString  `xml:"t:AttachmentPreviewUrl,omitempty"`
	PermissionType         XsInt     `xml:"t:PermissionType,omitempty"`
	OriginalPermissionType XsInt     `xml:"t:OriginalPermissionType,omitempty"`
	AttachmentIsFolder     XsBoolean `xml:"t:AttachmentIsFolder,omitempty"`
}

type NonEmptyArrayOfInternetHeadersType struct {
	InternetMessageHeader []*InternetHeaderType `xml:"t:InternetMessageHeader,omitempty"`
}

type InternetHeaderType struct {
	CharData   XsString `xml:",chardata"`
	HeaderName XsString `xml:"HeaderName,attr"`
}

type NonEmptyArrayOfResponseObjectsType struct {
	AcceptItem              []*AcceptItemType              `xml:"t:AcceptItem,omitempty"`
	TentativelyAcceptItem   []*TentativelyAcceptItemType   `xml:"t:TentativelyAcceptItem,omitempty"`
	DeclineItem             []*DeclineItemType             `xml:"t:DeclineItem,omitempty"`
	ReplyToItem             []*ReplyToItemType             `xml:"t:ReplyToItem,omitempty"`
	ForwardItem             []*ForwardItemType             `xml:"t:ForwardItem,omitempty"`
	ReplyAllToItem          []*ReplyAllToItemType          `xml:"t:ReplyAllToItem,omitempty"`
	CancelCalendarItem      []*CancelCalendarItemType      `xml:"t:CancelCalendarItem,omitempty"`
	RemoveItem              []*RemoveItemType              `xml:"t:RemoveItem,omitempty"`
	SuppressReadReceipt     []*SuppressReadReceiptType     `xml:"t:SuppressReadReceipt,omitempty"`
	PostReplyItem           []*PostReplyItemType           `xml:"t:PostReplyItem,omitempty"`
	AcceptSharingInvitation []*AcceptSharingInvitationType `xml:"t:AcceptSharingInvitation,omitempty"`
	AddItemToMyCalendar     []*AddItemToMyCalendarType     `xml:"t:AddItemToMyCalendar,omitempty"`
	ProposeNewTime          []*ProposeNewTimeType          `xml:"t:ProposeNewTime,omitempty"`
}

type AddItemToMyCalendarType struct {
	ResponseObjectType `xml:",omitempty"`
}

type ProposeNewTimeType struct {
	ResponseObjectType `xml:",omitempty"`
}

type ExtendedPropertyType struct {
	ExtendedFieldURI *PathToExtendedFieldType           `xml:"t:ExtendedFieldURI,omitempty"`
	Value            XsString                           `xml:"t:Value,omitempty"`
	Values           *NonEmptyArrayOfPropertyValuesType `xml:"t:Values,omitempty"`
}

type PathToExtendedFieldType struct {
	BasePathToElementType      `xml:",omitempty"`
	DistinguishedPropertySetId DistinguishedPropertySetType `xml:"DistinguishedPropertySetId,attr"`
	PropertySetId              GuidType                     `xml:"PropertySetId,attr"`
	PropertyTag                PropertyTagType              `xml:"PropertyTag,attr"`
	PropertyName               XsString                     `xml:"PropertyName,attr"`
	PropertyId                 XsInt                        `xml:"PropertyId,attr"`
	PropertyType               MapiPropertyTypeType         `xml:"PropertyType,attr"`
}

type BasePathToElementType struct {
}

type NonEmptyArrayOfPropertyValuesType struct {
	Value []XsString `xml:"t:Value,omitempty"`
}

type EffectiveRightsType struct {
	CreateAssociated XsBoolean `xml:"t:CreateAssociated,omitempty"`
	CreateContents   XsBoolean `xml:"t:CreateContents,omitempty"`
	CreateHierarchy  XsBoolean `xml:"t:CreateHierarchy,omitempty"`
	Delete           XsBoolean `xml:"t:Delete,omitempty"`
	Modify           XsBoolean `xml:"t:Modify,omitempty"`
	Read             XsBoolean `xml:"t:Read,omitempty"`
	ViewPrivateItems XsBoolean `xml:"t:ViewPrivateItems,omitempty"`
}

type FlagType struct {
	FlagStatus   FlagStatusType `xml:"t:FlagStatus,omitempty"`
	StartDate    XsDateTime     `xml:"t:StartDate,omitempty"`
	DueDate      XsDateTime     `xml:"t:DueDate,omitempty"`
	CompleteDate XsDateTime     `xml:"t:CompleteDate,omitempty"`
}

type EntityExtractionResultType struct {
	Addresses          *ArrayOfAddressEntitiesType        `xml:"t:Addresses,omitempty"`
	MeetingSuggestions *ArrayOfMeetingSuggestionsType     `xml:"t:MeetingSuggestions,omitempty"`
	TaskSuggestions    *ArrayOfTaskSuggestionsType        `xml:"t:TaskSuggestions,omitempty"`
	EmailAddresses     *ArrayOfEmailAddressEntitiesType   `xml:"t:EmailAddresses,omitempty"`
	Contacts           *ArrayOfContactsType               `xml:"t:Contacts,omitempty"`
	Urls               *ArrayOfUrlEntitiesType            `xml:"t:Urls,omitempty"`
	PhoneNumbers       *ArrayOfPhoneEntitiesType          `xml:"t:PhoneNumbers,omitempty"`
	ParcelDeliveries   *ArrayOfParcelDeliveryEntitiesType `xml:"t:ParcelDeliveries,omitempty"`
	FlightReservations *ArrayOfFlightReservationsType     `xml:"t:FlightReservations,omitempty"`
	SenderAddIns       *ArrayOfSenderAddInsType           `xml:"t:SenderAddIns,omitempty"`
}

type ArrayOfAddressEntitiesType struct {
	AddressEntity []*AddressEntityType `xml:"t:AddressEntity,omitempty"`
}

type AddressEntityType struct {
	EntityType `xml:",omitempty"`
	Address    XsString `xml:"t:Address,omitempty"`
}

type EntityType struct {
	Position []EmailPositionType `xml:"t:Position,omitempty"`
}

type ArrayOfMeetingSuggestionsType struct {
	MeetingSuggestion []*MeetingSuggestionType `xml:"t:MeetingSuggestion,omitempty"`
}

type MeetingSuggestionType struct {
	EntityType           `xml:",omitempty"`
	Attendees            *ArrayOfEmailUsersType `xml:"t:Attendees,omitempty"`
	Location             XsString               `xml:"t:Location,omitempty"`
	Subject              XsString               `xml:"t:Subject,omitempty"`
	MeetingString        XsString               `xml:"t:MeetingString,omitempty"`
	StartTime            XsDateTime             `xml:"t:StartTime,omitempty"`
	EndTime              XsDateTime             `xml:"t:EndTime,omitempty"`
	TimeStringBeginIndex XsInteger              `xml:"t:TimeStringBeginIndex,omitempty"`
	TimeStringLength     XsInteger              `xml:"t:TimeStringLength,omitempty"`
	EntityId             XsString               `xml:"t:EntityId,omitempty"`
	ExtractionId         XsString               `xml:"t:ExtractionId,omitempty"`
}

type ArrayOfEmailUsersType struct {
	EmailUser []*EmailUserType `xml:"t:EmailUser,omitempty"`
}

type EmailUserType struct {
	Name   XsString `xml:"t:Name,omitempty"`
	UserId XsString `xml:"t:UserId,omitempty"`
}

type ArrayOfTaskSuggestionsType struct {
	TaskSuggestion []*TaskSuggestionType `xml:"t:TaskSuggestion,omitempty"`
}

type TaskSuggestionType struct {
	EntityType `xml:",omitempty"`
	TaskString XsString               `xml:"t:TaskString,omitempty"`
	Assignees  *ArrayOfEmailUsersType `xml:"t:Assignees,omitempty"`
}

type ArrayOfEmailAddressEntitiesType struct {
	EmailAddressEntity []*EmailAddressEntityType `xml:"t:EmailAddressEntity,omitempty"`
}

type EmailAddressEntityType struct {
	EntityType   `xml:",omitempty"`
	EmailAddress XsString `xml:"t:EmailAddress,omitempty"`
}

type ArrayOfContactsType struct {
	Contact []*ContactType `xml:"t:Contact,omitempty"`
}

type ContactType struct {
	EntityType     `xml:",omitempty"`
	PersonName     XsString                        `xml:"t:PersonName,omitempty"`
	BusinessName   XsString                        `xml:"t:BusinessName,omitempty"`
	PhoneNumbers   *ArrayOfPhonesType              `xml:"t:PhoneNumbers,omitempty"`
	Urls           *ArrayOfUrlsType                `xml:"t:Urls,omitempty"`
	EmailAddresses *ArrayOfExtractedEmailAddresses `xml:"t:EmailAddresses,omitempty"`
	Addresses      *ArrayOfAddressesType           `xml:"t:Addresses,omitempty"`
	ContactString  XsString                        `xml:"t:ContactString,omitempty"`
}

type ArrayOfPhonesType struct {
	Phone []*PhoneType `xml:"t:Phone,omitempty"`
}

type PhoneType struct {
	OriginalPhoneString XsString `xml:"t:OriginalPhoneString,omitempty"`
	PhoneString         XsString `xml:"t:PhoneString,omitempty"`
	Type                XsString `xml:"t:Type,omitempty"`
}

type ArrayOfUrlsType struct {
	Url []XsString `xml:"t:Url,omitempty"`
}

type ArrayOfExtractedEmailAddresses struct {
	EmailAddress []XsString `xml:"t:EmailAddress,omitempty"`
}

type ArrayOfAddressesType struct {
	Address []XsString `xml:"t:Address,omitempty"`
}

type ArrayOfUrlEntitiesType struct {
	UrlEntity []*UrlEntityType `xml:"t:UrlEntity,omitempty"`
}

type UrlEntityType struct {
	EntityType `xml:",omitempty"`
	Url        XsString `xml:"t:Url,omitempty"`
}

type ArrayOfPhoneEntitiesType struct {
	Phone []*PhoneEntityType `xml:"t:Phone,omitempty"`
}

type PhoneEntityType struct {
	EntityType          `xml:",omitempty"`
	OriginalPhoneString XsString `xml:"t:OriginalPhoneString,omitempty"`
	PhoneString         XsString `xml:"t:PhoneString,omitempty"`
	Type                XsString `xml:"t:Type,omitempty"`
}

type ArrayOfParcelDeliveryEntitiesType struct {
	ParcelDelivery []*ParcelDeliveryEntityType `xml:"t:ParcelDelivery,omitempty"`
}

type ParcelDeliveryEntityType struct {
	Carrier              XsString `xml:"t:Carrier,omitempty"`
	TrackingNumber       XsString `xml:"t:TrackingNumber,omitempty"`
	TrackingUrl          XsString `xml:"t:TrackingUrl,omitempty"`
	ExpectedArrivalFrom  XsString `xml:"t:ExpectedArrivalFrom,omitempty"`
	ExpectedArrivalUntil XsString `xml:"t:ExpectedArrivalUntil,omitempty"`
	Product              XsString `xml:"t:Product,omitempty"`
	ProductUrl           XsString `xml:"t:ProductUrl,omitempty"`
	ProductImage         XsString `xml:"t:ProductImage,omitempty"`
	ProductSku           XsString `xml:"t:ProductSku,omitempty"`
	ProductDescription   XsString `xml:"t:ProductDescription,omitempty"`
	ProductBrand         XsString `xml:"t:ProductBrand,omitempty"`
	ProductColor         XsString `xml:"t:ProductColor,omitempty"`
	OrderNumber          XsString `xml:"t:OrderNumber,omitempty"`
	Seller               XsString `xml:"t:Seller,omitempty"`
	OrderStatus          XsString `xml:"t:OrderStatus,omitempty"`
	AddressName          XsString `xml:"t:AddressName,omitempty"`
	StreetAddress        XsString `xml:"t:StreetAddress,omitempty"`
	AddressLocality      XsString `xml:"t:AddressLocality,omitempty"`
	AddressRegion        XsString `xml:"t:AddressRegion,omitempty"`
	AddressCountry       XsString `xml:"t:AddressCountry,omitempty"`
	PostalCode           XsString `xml:"t:PostalCode,omitempty"`
}

type ArrayOfFlightReservationsType struct {
	FlightReservation []*FlightReservationEntityType `xml:"t:FlightReservation,omitempty"`
}

type FlightReservationEntityType struct {
	ReservationId     XsString            `xml:"t:ReservationId,omitempty"`
	ReservationStatus XsString            `xml:"t:ReservationStatus,omitempty"`
	UnderName         XsString            `xml:"t:UnderName,omitempty"`
	BrokerName        XsString            `xml:"t:BrokerName,omitempty"`
	BrokerPhone       XsString            `xml:"t:BrokerPhone,omitempty"`
	Flights           *ArrayOfFlightsType `xml:"t:Flights,omitempty"`
}

type ArrayOfFlightsType struct {
	Flight []*FlightEntityType `xml:"t:Flight,omitempty"`
}

type FlightEntityType struct {
	FlightNumber             XsString `xml:"t:FlightNumber,omitempty"`
	AirlineIataCode          XsString `xml:"t:AirlineIataCode,omitempty"`
	DepartureTime            XsString `xml:"t:DepartureTime,omitempty"`
	WindowsTimeZoneName      XsString `xml:"t:WindowsTimeZoneName,omitempty"`
	DepartureAirportIataCode XsString `xml:"t:DepartureAirportIataCode,omitempty"`
	ArrivalAirportIataCode   XsString `xml:"t:ArrivalAirportIataCode,omitempty"`
}

type ArrayOfSenderAddInsType struct {
	MicrosoftOutlookServicesSenderApp []*SenderAddInEntityType `xml:"t:Microsoft.OutlookServices.SenderApp,omitempty"`
}

type SenderAddInEntityType struct {
	ExtensionId XsString `xml:"t:ExtensionId,omitempty"`
}

type RetentionTagType struct {
	CharData   GuidType  `xml:",chardata"`
	IsExplicit XsBoolean `xml:"IsExplicit,attr"`
}

type RightsManagementLicenseDataType struct {
	RightsManagedMessageDecryptionStatus XsInt     `xml:"t:RightsManagedMessageDecryptionStatus,omitempty"`
	RmsTemplateId                        XsString  `xml:"t:RmsTemplateId,omitempty"`
	TemplateName                         XsString  `xml:"t:TemplateName,omitempty"`
	TemplateDescription                  XsString  `xml:"t:TemplateDescription,omitempty"`
	EditAllowed                          XsBoolean `xml:"t:EditAllowed,omitempty"`
	ReplyAllowed                         XsBoolean `xml:"t:ReplyAllowed,omitempty"`
	ReplyAllAllowed                      XsBoolean `xml:"t:ReplyAllAllowed,omitempty"`
	ForwardAllowed                       XsBoolean `xml:"t:ForwardAllowed,omitempty"`
	ModifyRecipientsAllowed              XsBoolean `xml:"t:ModifyRecipientsAllowed,omitempty"`
	ExtractAllowed                       XsBoolean `xml:"t:ExtractAllowed,omitempty"`
	PrintAllowed                         XsBoolean `xml:"t:PrintAllowed,omitempty"`
	ExportAllowed                        XsBoolean `xml:"t:ExportAllowed,omitempty"`
	ProgrammaticAccessAllowed            XsBoolean `xml:"t:ProgrammaticAccessAllowed,omitempty"`
	IsOwner                              XsBoolean `xml:"t:IsOwner,omitempty"`
	ContentOwner                         XsString  `xml:"t:ContentOwner,omitempty"`
	ContentExpiryDate                    XsString  `xml:"t:ContentExpiryDate,omitempty"`
}

type NonEmptyArrayOfPredictedActionReasonType struct {
	PredictedActionReason []PredictedActionReasonType `xml:"t:PredictedActionReason,omitempty"`
}

type MentionsPreviewType struct {
	IsMentioned XsBoolean `xml:"t:IsMentioned,omitempty"`
}

type NonEmptyArrayOfMentionActionsType struct {
	MentionAction []*MentionActionType `xml:"t:MentionAction,omitempty"`
}

type MentionActionType struct {
	Id                    XsString                  `xml:"t:Id,omitempty"`
	CreatedBy             *EmailAddressExtendedType `xml:"t:CreatedBy,omitempty"`
	CreatedDateTime       XsString                  `xml:"t:CreatedDateTime,omitempty"`
	ServerCreatedDateTime XsString                  `xml:"t:ServerCreatedDateTime,omitempty"`
	DeepLink              XsString                  `xml:"t:DeepLink,omitempty"`
	Application           XsString                  `xml:"t:Application,omitempty"`
	Mentioned             *EmailAddressExtendedType `xml:"t:Mentioned,omitempty"`
	MentionText           XsString                  `xml:"t:MentionText,omitempty"`
	ClientReference       XsString                  `xml:"t:ClientReference,omitempty"`
}

type EmailAddressExtendedType struct {
	EmailAddressType    `xml:",omitempty"`
	ExternalObjectId    XsString           `xml:"t:ExternalObjectId,omitempty"`
	PrimaryEmailAddress NonEmptyStringType `xml:"t:PrimaryEmailAddress,omitempty"`
}

type NonEmptyArrayOfAppliedHashtagType struct {
	AppliedHashtag []*AppliedHashtagType `xml:"t:AppliedHashtag,omitempty"`
}

type AppliedHashtagType struct {
	Id                    XsString                  `xml:"t:Id,omitempty"`
	CreatedBy             *EmailAddressExtendedType `xml:"t:CreatedBy,omitempty"`
	CreatedDateTime       XsString                  `xml:"t:CreatedDateTime,omitempty"`
	ServerCreatedDateTime XsString                  `xml:"t:ServerCreatedDateTime,omitempty"`
	DeepLink              XsString                  `xml:"t:DeepLink,omitempty"`
	Application           XsString                  `xml:"t:Application,omitempty"`
	Tag                   XsString                  `xml:"t:Tag,omitempty"`
	IsAutoTagged          XsBoolean                 `xml:"t:IsAutoTagged,omitempty"`
	IsInlined             XsBoolean                 `xml:"t:IsInlined,omitempty"`
}

type AppliedHashtagsPreviewType struct {
	Hashtags *ArrayOfStringsType `xml:"t:Hashtags,omitempty"`
}

type NonEmptyArrayOfLikeType struct {
	Like []*LikeType `xml:"t:Like,omitempty"`
}

type LikeType struct {
	Id                    XsString                  `xml:"t:Id,omitempty"`
	CreatedBy             *EmailAddressExtendedType `xml:"t:CreatedBy,omitempty"`
	CreatedDateTime       XsString                  `xml:"t:CreatedDateTime,omitempty"`
	ServerCreatedDateTime XsString                  `xml:"t:ServerCreatedDateTime,omitempty"`
	DeepLink              XsString                  `xml:"t:DeepLink,omitempty"`
	Application           XsString                  `xml:"t:Application,omitempty"`
}

type LikesPreviewType struct {
	LikeCount XsInt `xml:"t:LikeCount,omitempty"`
}

type ConflictResultsType struct {
	Count XsInt `xml:"t:Count,omitempty"`
}

type ArrayOfAttachmentsType struct {
	ItemAttachment      []*ItemAttachmentType      `xml:"t:ItemAttachment,omitempty"`
	FileAttachment      []*FileAttachmentType      `xml:"t:FileAttachment,omitempty"`
	ReferenceAttachment []*ReferenceAttachmentType `xml:"t:ReferenceAttachment,omitempty"`
}

type ArrayOfFoldersType struct {
	Folder         []*FolderType         `xml:"t:Folder,omitempty"`
	CalendarFolder []*CalendarFolderType `xml:"t:CalendarFolder,omitempty"`
	ContactsFolder []*ContactsFolderType `xml:"t:ContactsFolder,omitempty"`
	SearchFolder   []*SearchFolderType   `xml:"t:SearchFolder,omitempty"`
	TasksFolder    []*TasksFolderType    `xml:"t:TasksFolder,omitempty"`
}

type FolderType struct {
	BaseFolderType `xml:",omitempty"`
	PermissionSet  *PermissionSetType `xml:"t:PermissionSet,omitempty"`
	UnreadCount    XsInt              `xml:"t:UnreadCount,omitempty"`
}

type BaseFolderType struct {
	FolderId                 *FolderIdType                 `xml:"t:FolderId,omitempty"`
	ParentFolderId           *FolderIdType                 `xml:"t:ParentFolderId,omitempty"`
	FolderClass              XsString                      `xml:"t:FolderClass,omitempty"`
	DisplayName              XsString                      `xml:"t:DisplayName,omitempty"`
	TotalCount               XsInt                         `xml:"t:TotalCount,omitempty"`
	ChildFolderCount         XsInt                         `xml:"t:ChildFolderCount,omitempty"`
	ExtendedProperty         []*ExtendedPropertyType       `xml:"t:ExtendedProperty,omitempty"`
	ManagedFolderInformation *ManagedFolderInformationType `xml:"t:ManagedFolderInformation,omitempty"`
	EffectiveRights          *EffectiveRightsType          `xml:"t:EffectiveRights,omitempty"`
	DistinguishedFolderId    DistinguishedFolderIdNameType `xml:"t:DistinguishedFolderId,omitempty"`
	PolicyTag                *RetentionTagType             `xml:"t:PolicyTag,omitempty"`
	ArchiveTag               *RetentionTagType             `xml:"t:ArchiveTag,omitempty"`
	ReplicaList              *ArrayOfStringsType           `xml:"t:ReplicaList,omitempty"`
}

type ManagedFolderInformationType struct {
	CanDelete            XsBoolean `xml:"t:CanDelete,omitempty"`
	CanRenameOrMove      XsBoolean `xml:"t:CanRenameOrMove,omitempty"`
	MustDisplayComment   XsBoolean `xml:"t:MustDisplayComment,omitempty"`
	HasQuota             XsBoolean `xml:"t:HasQuota,omitempty"`
	IsManagedFoldersRoot XsBoolean `xml:"t:IsManagedFoldersRoot,omitempty"`
	ManagedFolderId      XsString  `xml:"t:ManagedFolderId,omitempty"`
	Comment              XsString  `xml:"t:Comment,omitempty"`
	StorageQuota         XsInt     `xml:"t:StorageQuota,omitempty"`
	FolderSize           XsInt     `xml:"t:FolderSize,omitempty"`
	HomePage             XsString  `xml:"t:HomePage,omitempty"`
}

type PermissionSetType struct {
	Permissions    *ArrayOfPermissionsType    `xml:"t:Permissions,omitempty"`
	UnknownEntries *ArrayOfUnknownEntriesType `xml:"t:UnknownEntries,omitempty"`
}

type ArrayOfPermissionsType struct {
	Permission []*PermissionType `xml:"t:Permission,omitempty"`
}

type PermissionType struct {
	BasePermissionType `xml:",omitempty"`
	ReadItems          PermissionReadAccessType `xml:"t:ReadItems,omitempty"`
	PermissionLevel    PermissionLevelType      `xml:"t:PermissionLevel,omitempty"`
}

type BasePermissionType struct {
	UserId              *UserIdType          `xml:"t:UserId,omitempty"`
	CanCreateItems      XsBoolean            `xml:"t:CanCreateItems,omitempty"`
	CanCreateSubFolders XsBoolean            `xml:"t:CanCreateSubFolders,omitempty"`
	IsFolderOwner       XsBoolean            `xml:"t:IsFolderOwner,omitempty"`
	IsFolderVisible     XsBoolean            `xml:"t:IsFolderVisible,omitempty"`
	IsFolderContact     XsBoolean            `xml:"t:IsFolderContact,omitempty"`
	EditItems           PermissionActionType `xml:"t:EditItems,omitempty"`
	DeleteItems         PermissionActionType `xml:"t:DeleteItems,omitempty"`
}

type UserIdType struct {
	SID                  XsString              `xml:"t:SID,omitempty"`
	PrimarySmtpAddress   XsString              `xml:"t:PrimarySmtpAddress,omitempty"`
	DisplayName          XsString              `xml:"t:DisplayName,omitempty"`
	DistinguishedUser    DistinguishedUserType `xml:"t:DistinguishedUser,omitempty"`
	ExternalUserIdentity XsString              `xml:"t:ExternalUserIdentity,omitempty"`
}

type ArrayOfUnknownEntriesType struct {
	UnknownEntry []XsString `xml:"t:UnknownEntry,omitempty"`
}

type CalendarFolderType struct {
	BaseFolderType         `xml:",omitempty"`
	SharingEffectiveRights CalendarPermissionReadAccessType `xml:"t:SharingEffectiveRights,omitempty"`
	PermissionSet          *CalendarPermissionSetType       `xml:"t:PermissionSet,omitempty"`
}

type CalendarPermissionSetType struct {
	CalendarPermissions *ArrayOfCalendarPermissionsType `xml:"t:CalendarPermissions,omitempty"`
	UnknownEntries      *ArrayOfUnknownEntriesType      `xml:"t:UnknownEntries,omitempty"`
}

type ArrayOfCalendarPermissionsType struct {
	CalendarPermission []*CalendarPermissionType `xml:"t:CalendarPermission,omitempty"`
}

type CalendarPermissionType struct {
	BasePermissionType      `xml:",omitempty"`
	ReadItems               CalendarPermissionReadAccessType `xml:"t:ReadItems,omitempty"`
	CalendarPermissionLevel CalendarPermissionLevelType      `xml:"t:CalendarPermissionLevel,omitempty"`
}

type ContactsFolderType struct {
	BaseFolderType         `xml:",omitempty"`
	SharingEffectiveRights PermissionReadAccessType `xml:"t:SharingEffectiveRights,omitempty"`
	PermissionSet          *PermissionSetType       `xml:"t:PermissionSet,omitempty"`
	SourceId               XsString                 `xml:"t:SourceId,omitempty"`
	AccountName            XsString                 `xml:"t:AccountName,omitempty"`
}

type SearchFolderType struct {
	FolderType       `xml:",omitempty"`
	SearchParameters *SearchParametersType `xml:"t:SearchParameters,omitempty"`
}

type SearchParametersType struct {
	Traversal     SearchFolderTraversalType         `xml:"Traversal,attr"`
	Restriction   *RestrictionType                  `xml:"t:Restriction,omitempty"`
	BaseFolderIds *NonEmptyArrayOfBaseFolderIdsType `xml:"t:BaseFolderIds,omitempty"`
}

type RestrictionType struct {
	SearchExpression SearchExpressionType `xml:"t:SearchExpression,omitempty"`
}

type SearchExpressionType struct {
}

type TasksFolderType struct {
	FolderType `xml:",omitempty"`
}

type FindFolderParentType struct {
	IndexedPagingOffset     XsInt               `xml:"IndexedPagingOffset,attr"`
	NumeratorOffset         XsInt               `xml:"NumeratorOffset,attr"`
	AbsoluteDenominator     XsInt               `xml:"AbsoluteDenominator,attr"`
	IncludesLastItemInRange XsBoolean           `xml:"IncludesLastItemInRange,attr"`
	TotalItemsInView        XsInt               `xml:"TotalItemsInView,attr"`
	Folders                 *ArrayOfFoldersType `xml:"t:Folders,omitempty"`
}

type RootItemIdType struct {
	BaseItemIdType    `xml:",omitempty"`
	RootItemId        XsString `xml:"RootItemId,attr"`
	RootItemChangeKey XsString `xml:"RootItemChangeKey,attr"`
}

type ClientAccessTokenType struct {
	Id         XsString                  `xml:"t:Id,omitempty"`
	TokenType  ClientAccessTokenTypeType `xml:"t:TokenType,omitempty"`
	TokenValue XsString                  `xml:"t:TokenValue,omitempty"`
	TTL        XsInteger                 `xml:"t:TTL,omitempty"`
}

type FindItemParentType struct {
	IndexedPagingOffset     XsInt                    `xml:"IndexedPagingOffset,attr"`
	NumeratorOffset         XsInt                    `xml:"NumeratorOffset,attr"`
	AbsoluteDenominator     XsInt                    `xml:"AbsoluteDenominator,attr"`
	IncludesLastItemInRange XsBoolean                `xml:"IncludesLastItemInRange,attr"`
	TotalItemsInView        XsInt                    `xml:"TotalItemsInView,attr"`
	Items                   *ArrayOfRealItemsType    `xml:"t:Items,omitempty"`
	Groups                  *ArrayOfGroupedItemsType `xml:"t:Groups,omitempty"`
}

type ArrayOfGroupedItemsType struct {
	GroupedItems []*GroupedItemsType `xml:"t:GroupedItems,omitempty"`
}

type GroupedItemsType struct {
	GroupIndex   XsString              `xml:"t:GroupIndex,omitempty"`
	Items        *ArrayOfRealItemsType `xml:"t:Items,omitempty"`
	GroupSummary *GroupSummaryType     `xml:"t:GroupSummary,omitempty"`
}

type GroupSummaryType struct {
	GroupCount   XsInt          `xml:"t:GroupCount,omitempty"`
	UnreadCount  XsInt          `xml:"t:UnreadCount,omitempty"`
	InstanceKey  XsBase64Binary `xml:"t:InstanceKey,omitempty"`
	GroupByValue XsString       `xml:"t:GroupByValue,omitempty"`
}

type ArrayOfHighlightTermsType struct {
	Term []*HighlightTermType `xml:"t:Term,omitempty"`
}

type HighlightTermType struct {
	Scope XsString `xml:"t:Scope,omitempty"`
	Value XsString `xml:"t:Value,omitempty"`
}

type ArrayOfResolutionType struct {
	IndexedPagingOffset     XsInt             `xml:"IndexedPagingOffset,attr"`
	NumeratorOffset         XsInt             `xml:"NumeratorOffset,attr"`
	AbsoluteDenominator     XsInt             `xml:"AbsoluteDenominator,attr"`
	IncludesLastItemInRange XsBoolean         `xml:"IncludesLastItemInRange,attr"`
	TotalItemsInView        XsInt             `xml:"TotalItemsInView,attr"`
	Resolution              []*ResolutionType `xml:"t:Resolution,omitempty"`
}

type ResolutionType struct {
	Mailbox *EmailAddressType `xml:"t:Mailbox,omitempty"`
	Contact *ContactItemType  `xml:"t:Contact,omitempty"`
}

type ArrayOfDLExpansionType struct {
	IndexedPagingOffset     XsInt               `xml:"IndexedPagingOffset,attr"`
	NumeratorOffset         XsInt               `xml:"NumeratorOffset,attr"`
	AbsoluteDenominator     XsInt               `xml:"AbsoluteDenominator,attr"`
	IncludesLastItemInRange XsBoolean           `xml:"IncludesLastItemInRange,attr"`
	TotalItemsInView        XsInt               `xml:"TotalItemsInView,attr"`
	Mailbox                 []*EmailAddressType `xml:"t:Mailbox,omitempty"`
}

type ArrayOfTimeZoneDefinitionType struct {
	TimeZoneDefinition []*TimeZoneDefinitionType `xml:"t:TimeZoneDefinition,omitempty"`
}

type NotificationType struct {
	SubscriptionId       SubscriptionIdType            `xml:"t:SubscriptionId,omitempty"`
	PreviousWatermark    WatermarkType                 `xml:"t:PreviousWatermark,omitempty"`
	MoreEvents           XsBoolean                     `xml:"t:MoreEvents,omitempty"`
	CopiedEvent          []*MovedCopiedEventType       `xml:"t:CopiedEvent,omitempty"`
	CreatedEvent         []*BaseObjectChangedEventType `xml:"t:CreatedEvent,omitempty"`
	DeletedEvent         []*BaseObjectChangedEventType `xml:"t:DeletedEvent,omitempty"`
	ModifiedEvent        []*ModifiedEventType          `xml:"t:ModifiedEvent,omitempty"`
	MovedEvent           []*MovedCopiedEventType       `xml:"t:MovedEvent,omitempty"`
	NewMailEvent         []*BaseObjectChangedEventType `xml:"t:NewMailEvent,omitempty"`
	StatusEvent          []*BaseNotificationEventType  `xml:"t:StatusEvent,omitempty"`
	FreeBusyChangedEvent []*BaseObjectChangedEventType `xml:"t:FreeBusyChangedEvent,omitempty"`
}

type MovedCopiedEventType struct {
	BaseObjectChangedEventType `xml:",omitempty"`
	OldParentFolderId          *FolderIdType `xml:"t:OldParentFolderId,omitempty"`
	OldFolderId                *FolderIdType `xml:"t:OldFolderId,omitempty"`
	OldItemId                  *ItemIdType   `xml:"t:OldItemId,omitempty"`
}

type BaseObjectChangedEventType struct {
	BaseNotificationEventType `xml:",omitempty"`
	TimeStamp                 XsDateTime    `xml:"t:TimeStamp,omitempty"`
	ParentFolderId            *FolderIdType `xml:"t:ParentFolderId,omitempty"`
	FolderId                  *FolderIdType `xml:"t:FolderId,omitempty"`
	ItemId                    *ItemIdType   `xml:"t:ItemId,omitempty"`
}

type BaseNotificationEventType struct {
	Watermark WatermarkType `xml:"t:Watermark,omitempty"`
}

type ModifiedEventType struct {
	BaseObjectChangedEventType `xml:",omitempty"`
	UnreadCount                XsInt `xml:"t:UnreadCount,omitempty"`
}

type NonEmptyArrayOfNotificationsType struct {
	Notification []*NotificationType `xml:"t:Notification,omitempty"`
}

type NonEmptyArrayOfSubscriptionIdsType struct {
	SubscriptionId []SubscriptionIdType `xml:"t:SubscriptionId,omitempty"`
}

type SyncFolderHierarchyChangesType struct {
	Create []*SyncFolderHierarchyCreateOrUpdateType `xml:"t:Create,omitempty"`
	Update []*SyncFolderHierarchyCreateOrUpdateType `xml:"t:Update,omitempty"`
	Delete []*SyncFolderHierarchyDeleteType         `xml:"t:Delete,omitempty"`
}

type SyncFolderHierarchyCreateOrUpdateType struct {
	Folder         *FolderType         `xml:"t:Folder,omitempty"`
	CalendarFolder *CalendarFolderType `xml:"t:CalendarFolder,omitempty"`
	ContactsFolder *ContactsFolderType `xml:"t:ContactsFolder,omitempty"`
	SearchFolder   *SearchFolderType   `xml:"t:SearchFolder,omitempty"`
	TasksFolder    *TasksFolderType    `xml:"t:TasksFolder,omitempty"`
}

type SyncFolderHierarchyDeleteType struct {
	FolderId *FolderIdType `xml:"t:FolderId,omitempty"`
}

type SyncFolderItemsChangesType struct {
	Create         []*SyncFolderItemsCreateOrUpdateType `xml:"t:Create,omitempty"`
	Update         []*SyncFolderItemsCreateOrUpdateType `xml:"t:Update,omitempty"`
	Delete         []*SyncFolderItemsDeleteType         `xml:"t:Delete,omitempty"`
	ReadFlagChange []*SyncFolderItemsReadFlagType       `xml:"t:ReadFlagChange,omitempty"`
}

type SyncFolderItemsCreateOrUpdateType struct {
	Item                *ItemType                       `xml:"t:Item,omitempty"`
	Message             *MessageType                    `xml:"t:Message,omitempty"`
	SharingMessage      *SharingMessageType             `xml:"t:SharingMessage,omitempty"`
	CalendarItem        *CalendarItemType               `xml:"t:CalendarItem,omitempty"`
	Contact             *ContactItemType                `xml:"t:Contact,omitempty"`
	DistributionList    *DistributionListType           `xml:"t:DistributionList,omitempty"`
	MeetingMessage      *MeetingMessageType             `xml:"t:MeetingMessage,omitempty"`
	MeetingRequest      *MeetingRequestMessageType      `xml:"t:MeetingRequest,omitempty"`
	MeetingResponse     *MeetingResponseMessageType     `xml:"t:MeetingResponse,omitempty"`
	MeetingCancellation *MeetingCancellationMessageType `xml:"t:MeetingCancellation,omitempty"`
	Task                *TaskType                       `xml:"t:Task,omitempty"`
	PostItem            *PostItemType                   `xml:"t:PostItem,omitempty"`
	RoleMember          *RoleMemberItemType             `xml:"t:RoleMember,omitempty"`
	Network             *NetworkItemType                `xml:"t:Network,omitempty"`
	Person              *AbchPersonItemType             `xml:"t:Person,omitempty"`
}

type SyncFolderItemsDeleteType struct {
	ItemId *ItemIdType `xml:"t:ItemId,omitempty"`
}

type SyncFolderItemsReadFlagType struct {
	ItemId *ItemIdType `xml:"t:ItemId,omitempty"`
	IsRead XsBoolean   `xml:"t:IsRead,omitempty"`
}

type AlternateIdBaseType struct {
	Format IdFormatType `xml:"Format,attr"`
}

type ArrayOfEncryptedSharedFolderDataType struct {
	EncryptedSharedFolderData []*EncryptedSharedFolderDataType `xml:"t:EncryptedSharedFolderData,omitempty"`
}

type EncryptedSharedFolderDataType struct {
	Token EncryptedDataContainerType `xml:"t:Token,omitempty"`
	Data  EncryptedDataContainerType `xml:"t:Data,omitempty"`
}

type EncryptedDataContainerType struct {
}

type ArrayOfInvalidRecipientsType struct {
	InvalidRecipient []*InvalidRecipientType `xml:"t:InvalidRecipient,omitempty"`
}

type InvalidRecipientType struct {
	SmtpAddress  NonEmptyStringType               `xml:"t:SmtpAddress,omitempty"`
	ResponseCode InvalidRecipientResponseCodeType `xml:"t:ResponseCode,omitempty"`
	MessageText  XsString                         `xml:"t:MessageText,omitempty"`
}

type UserConfigurationType struct {
	UserConfigurationName *UserConfigurationNameType       `xml:"t:UserConfigurationName,omitempty"`
	ItemId                *ItemIdType                      `xml:"t:ItemId,omitempty"`
	Dictionary            *UserConfigurationDictionaryType `xml:"t:Dictionary,omitempty"`
	XmlData               XsBase64Binary                   `xml:"t:XmlData,omitempty"`
	BinaryData            XsBase64Binary                   `xml:"t:BinaryData,omitempty"`
}

type UserConfigurationNameType struct {
	TargetFolderIdType `xml:",omitempty"`
	Name               NonEmptyStringType `xml:"Name,attr"`
}

type TargetFolderIdType struct {
	FolderId              *FolderIdType              `xml:"t:FolderId,omitempty"`
	DistinguishedFolderId *DistinguishedFolderIdType `xml:"t:DistinguishedFolderId,omitempty"`
	AddressListId         *AddressListIdType         `xml:"t:AddressListId,omitempty"`
}

type AddressListIdType struct {
	BaseFolderIdType `xml:",omitempty"`
	Id               XsString `xml:"Id,attr"`
}

type UserConfigurationDictionaryType struct {
	DictionaryEntry []*UserConfigurationDictionaryEntryType `xml:"t:DictionaryEntry,omitempty"`
}

type UserConfigurationDictionaryEntryType struct {
	DictionaryKey   *UserConfigurationDictionaryObjectType `xml:"t:DictionaryKey,omitempty"`
	DictionaryValue *UserConfigurationDictionaryObjectType `xml:"t:DictionaryValue,omitempty"`
}

type UserConfigurationDictionaryObjectType struct {
	Type  UserConfigurationDictionaryObjectTypesType `xml:"t:Type,omitempty"`
	Value []XsString                                 `xml:"t:Value,omitempty"`
}

type ArrayOfEmailAddressesType struct {
	Address []*EmailAddressType `xml:"t:Address,omitempty"`
}

type ArrayOfRoomsType struct {
	Room []*RoomType `xml:"t:Room,omitempty"`
}

type RoomType struct {
	DirectoryEntryType `xml:",omitempty"`
}

type DirectoryEntryType struct {
	Id *EmailAddressType `xml:"t:Id,omitempty"`
}

type ArrayOfRemindersType struct {
	Reminder []*ReminderType `xml:"t:Reminder,omitempty"`
}

type ReminderType struct {
	Subject               XsString          `xml:"t:Subject,omitempty"`
	Location              XsString          `xml:"t:Location,omitempty"`
	ReminderTime          XsDateTime        `xml:"t:ReminderTime,omitempty"`
	StartDate             XsDateTime        `xml:"t:StartDate,omitempty"`
	EndDate               XsDateTime        `xml:"t:EndDate,omitempty"`
	ItemId                *ItemIdType       `xml:"t:ItemId,omitempty"`
	RecurringMasterItemId *ItemIdType       `xml:"t:RecurringMasterItemId,omitempty"`
	ReminderGroup         ReminderGroupType `xml:"t:ReminderGroup,omitempty"`
	UID                   XsString          `xml:"t:UID,omitempty"`
}

type NonEmptyArrayOfItemIdsType struct {
	ItemId []*ItemIdType `xml:"t:ItemId,omitempty"`
}

type MailboxStatisticsSearchResultType struct {
	UserMailbox                   *UserMailboxType                   `xml:"t:UserMailbox,omitempty"`
	KeywordStatisticsSearchResult *KeywordStatisticsSearchResultType `xml:"t:KeywordStatisticsSearchResult,omitempty"`
}

type UserMailboxType struct {
	Id        XsString  `xml:"Id,attr"`
	IsArchive XsBoolean `xml:"IsArchive,attr"`
}

type KeywordStatisticsSearchResultType struct {
	Keyword  XsString `xml:"t:Keyword,omitempty"`
	ItemHits XsInt    `xml:"t:ItemHits,omitempty"`
	Size     XsLong   `xml:"t:Size,omitempty"`
}

type ArrayOfSearchableMailboxesType struct {
	SearchableMailbox []*SearchableMailboxType `xml:"t:SearchableMailbox,omitempty"`
}

type SearchableMailboxType struct {
	Guid                 GuidType  `xml:"t:Guid,omitempty"`
	PrimarySmtpAddress   XsString  `xml:"t:PrimarySmtpAddress,omitempty"`
	IsExternalMailbox    XsBoolean `xml:"t:IsExternalMailbox,omitempty"`
	ExternalEmailAddress XsString  `xml:"t:ExternalEmailAddress,omitempty"`
	DisplayName          XsString  `xml:"t:DisplayName,omitempty"`
	IsMembershipGroup    XsBoolean `xml:"t:IsMembershipGroup,omitempty"`
	ReferenceId          XsString  `xml:"t:ReferenceId,omitempty"`
}

type ArrayOfFailedSearchMailboxesType struct {
	FailedMailbox []*FailedSearchMailboxType `xml:"t:FailedMailbox,omitempty"`
}

type FailedSearchMailboxType struct {
	Mailbox      XsString  `xml:"t:Mailbox,omitempty"`
	ErrorCode    XsInt     `xml:"t:ErrorCode,omitempty"`
	ErrorMessage XsString  `xml:"t:ErrorMessage,omitempty"`
	IsArchive    XsBoolean `xml:"t:IsArchive,omitempty"`
}

type SearchMailboxesResultType struct {
	SearchQueries   *NonEmptyArrayOfMailboxQueriesType         `xml:"t:SearchQueries,omitempty"`
	ResultType      SearchResultType                           `xml:"t:ResultType,omitempty"`
	ItemCount       XsLong                                     `xml:"t:ItemCount,omitempty"`
	Size            XsLong                                     `xml:"t:Size,omitempty"`
	PageItemCount   XsInt                                      `xml:"t:PageItemCount,omitempty"`
	PageItemSize    XsLong                                     `xml:"t:PageItemSize,omitempty"`
	KeywordStats    *ArrayOfKeywordStatisticsSearchResultsType `xml:"t:KeywordStats,omitempty"`
	Items           *ArrayOfSearchPreviewItemsType             `xml:"t:Items,omitempty"`
	FailedMailboxes *ArrayOfFailedSearchMailboxesType          `xml:"t:FailedMailboxes,omitempty"`
	Refiners        *ArrayOfSearchRefinerItemsType             `xml:"t:Refiners,omitempty"`
	MailboxStats    *ArrayOfMailboxStatisticsItemsType         `xml:"t:MailboxStats,omitempty"`
}

type NonEmptyArrayOfMailboxQueriesType struct {
	MailboxQuery []*MailboxQueryType `xml:"t:MailboxQuery,omitempty"`
}

type MailboxQueryType struct {
	Query               XsString                                `xml:"t:Query,omitempty"`
	MailboxSearchScopes *NonEmptyArrayOfMailboxSearchScopesType `xml:"t:MailboxSearchScopes,omitempty"`
}

type NonEmptyArrayOfMailboxSearchScopesType struct {
	MailboxSearchScope []*MailboxSearchScopeType `xml:"t:MailboxSearchScope,omitempty"`
}

type MailboxSearchScopeType struct {
	Mailbox            XsString                       `xml:"t:Mailbox,omitempty"`
	SearchScope        MailboxSearchLocationType      `xml:"t:SearchScope,omitempty"`
	ExtendedAttributes *ArrayOfExtendedAttributesType `xml:"t:ExtendedAttributes,omitempty"`
}

type ArrayOfExtendedAttributesType struct {
	ExtendedAttribute []*ExtendedAttributeType `xml:"t:ExtendedAttribute,omitempty"`
}

type ExtendedAttributeType struct {
	Name  XsString `xml:"t:Name,omitempty"`
	Value XsString `xml:"t:Value,omitempty"`
}

type ArrayOfKeywordStatisticsSearchResultsType struct {
	KeywordStat []*KeywordStatisticsSearchResultType `xml:"t:KeywordStat,omitempty"`
}

type ArrayOfSearchPreviewItemsType struct {
	SearchPreviewItem []*SearchPreviewItemType `xml:"t:SearchPreviewItem,omitempty"`
}

type SearchPreviewItemType struct {
	Id                 *ItemIdType                          `xml:"t:Id,omitempty"`
	Mailbox            *PreviewItemMailboxType              `xml:"t:Mailbox,omitempty"`
	ParentId           *ItemIdType                          `xml:"t:ParentId,omitempty"`
	ItemClass          ItemClassType                        `xml:"t:ItemClass,omitempty"`
	UniqueHash         XsString                             `xml:"t:UniqueHash,omitempty"`
	SortValue          XsString                             `xml:"t:SortValue,omitempty"`
	OwaLink            XsString                             `xml:"t:OwaLink,omitempty"`
	Sender             XsString                             `xml:"t:Sender,omitempty"`
	ToRecipients       *ArrayOfSmtpAddressType              `xml:"t:ToRecipients,omitempty"`
	CcRecipients       *ArrayOfSmtpAddressType              `xml:"t:CcRecipients,omitempty"`
	BccRecipients      *ArrayOfSmtpAddressType              `xml:"t:BccRecipients,omitempty"`
	CreatedTime        XsDateTime                           `xml:"t:CreatedTime,omitempty"`
	ReceivedTime       XsDateTime                           `xml:"t:ReceivedTime,omitempty"`
	SentTime           XsDateTime                           `xml:"t:SentTime,omitempty"`
	Subject            XsString                             `xml:"t:Subject,omitempty"`
	Size               XsLong                               `xml:"t:Size,omitempty"`
	Preview            XsString                             `xml:"t:Preview,omitempty"`
	Importance         ImportanceChoicesType                `xml:"t:Importance,omitempty"`
	Read               XsBoolean                            `xml:"t:Read,omitempty"`
	HasAttachment      XsBoolean                            `xml:"t:HasAttachment,omitempty"`
	ExtendedProperties *NonEmptyArrayOfExtendedPropertyType `xml:"t:ExtendedProperties,omitempty"`
}

type PreviewItemMailboxType struct {
	MailboxId          XsString `xml:"t:MailboxId,omitempty"`
	PrimarySmtpAddress XsString `xml:"t:PrimarySmtpAddress,omitempty"`
}

type ArrayOfSmtpAddressType struct {
	SmtpAddress []NonEmptyStringType `xml:"t:SmtpAddress,omitempty"`
}

type NonEmptyArrayOfExtendedPropertyType struct {
	ExtendedProperty []*ExtendedPropertyType `xml:"t:ExtendedProperty,omitempty"`
}

type ArrayOfSearchRefinerItemsType struct {
	Refiner []*SearchRefinerItemType `xml:"t:Refiner,omitempty"`
}

type SearchRefinerItemType struct {
	Name  XsString `xml:"t:Name,omitempty"`
	Value XsString `xml:"t:Value,omitempty"`
	Count XsLong   `xml:"t:Count,omitempty"`
	Token XsString `xml:"t:Token,omitempty"`
}

type ArrayOfMailboxStatisticsItemsType struct {
	MailboxStat []*MailboxStatisticsItemType `xml:"t:MailboxStat,omitempty"`
}

type MailboxStatisticsItemType struct {
	MailboxId   XsString `xml:"t:MailboxId,omitempty"`
	DisplayName XsString `xml:"t:DisplayName,omitempty"`
	ItemCount   XsLong   `xml:"t:ItemCount,omitempty"`
	Size        XsLong   `xml:"t:Size,omitempty"`
}

type ArrayOfDiscoverySearchConfigurationType struct {
	DiscoverySearchConfiguration []*DiscoverySearchConfigurationType `xml:"t:DiscoverySearchConfiguration,omitempty"`
}

type DiscoverySearchConfigurationType struct {
	SearchId              XsString                        `xml:"t:SearchId,omitempty"`
	SearchQuery           XsString                        `xml:"t:SearchQuery,omitempty"`
	SearchableMailboxes   *ArrayOfSearchableMailboxesType `xml:"t:SearchableMailboxes,omitempty"`
	InPlaceHoldIdentity   XsString                        `xml:"t:InPlaceHoldIdentity,omitempty"`
	ManagedByOrganization XsString                        `xml:"t:ManagedByOrganization,omitempty"`
	Language              XsString                        `xml:"t:Language,omitempty"`
}

type MailboxHoldResultType struct {
	HoldId              XsString                      `xml:"t:HoldId,omitempty"`
	Query               XsString                      `xml:"t:Query,omitempty"`
	MailboxHoldStatuses *ArrayOfMailboxHoldStatusType `xml:"t:MailboxHoldStatuses,omitempty"`
}

type ArrayOfMailboxHoldStatusType struct {
	MailboxHoldStatus []*MailboxHoldStatusType `xml:"t:MailboxHoldStatus,omitempty"`
}

type MailboxHoldStatusType struct {
	Mailbox        XsString       `xml:"t:Mailbox,omitempty"`
	Status         HoldStatusType `xml:"t:Status,omitempty"`
	AdditionalInfo XsString       `xml:"t:AdditionalInfo,omitempty"`
}

type ArrayOfNonIndexableItemStatisticsType struct {
	NonIndexableItemStatistic []*NonIndexableItemStatisticType `xml:"t:NonIndexableItemStatistic,omitempty"`
}

type NonIndexableItemStatisticType struct {
	Mailbox      XsString `xml:"t:Mailbox,omitempty"`
	ItemCount    XsLong   `xml:"t:ItemCount,omitempty"`
	ErrorMessage XsString `xml:"t:ErrorMessage,omitempty"`
}

type NonIndexableItemDetailResultType struct {
	Items           *ArrayOfNonIndexableItemDetailsType `xml:"t:Items,omitempty"`
	FailedMailboxes *ArrayOfFailedSearchMailboxesType   `xml:"t:FailedMailboxes,omitempty"`
}

type ArrayOfNonIndexableItemDetailsType struct {
	NonIndexableItemDetail []*NonIndexableItemDetailType `xml:"t:NonIndexableItemDetail,omitempty"`
}

type NonIndexableItemDetailType struct {
	ItemId             *ItemIdType        `xml:"t:ItemId,omitempty"`
	ErrorCode          ItemIndexErrorType `xml:"t:ErrorCode,omitempty"`
	ErrorDescription   XsString           `xml:"t:ErrorDescription,omitempty"`
	IsPartiallyIndexed XsBoolean          `xml:"t:IsPartiallyIndexed,omitempty"`
	IsPermanentFailure XsBoolean          `xml:"t:IsPermanentFailure,omitempty"`
	SortValue          XsString           `xml:"t:SortValue,omitempty"`
	AttemptCount       XsInt              `xml:"t:AttemptCount,omitempty"`
	LastAttemptTime    XsDateTime         `xml:"t:LastAttemptTime,omitempty"`
	AdditionalInfo     XsString           `xml:"t:AdditionalInfo,omitempty"`
}

type ArrayOfPeopleType struct {
	Persona []*PersonaType `xml:"t:Persona,omitempty"`
}

type PersonaType struct {
	PersonaId                    *ItemIdType                                 `xml:"t:PersonaId,omitempty"`
	PersonaType                  XsString                                    `xml:"t:PersonaType,omitempty"`
	PersonaObjectStatus          XsString                                    `xml:"t:PersonaObjectStatus,omitempty"`
	CreationTime                 XsDateTime                                  `xml:"t:CreationTime,omitempty"`
	Bodies                       *ArrayOfBodyContentAttributedValuesType     `xml:"t:Bodies,omitempty"`
	DisplayNameFirstLastSortKey  XsString                                    `xml:"t:DisplayNameFirstLastSortKey,omitempty"`
	DisplayNameLastFirstSortKey  XsString                                    `xml:"t:DisplayNameLastFirstSortKey,omitempty"`
	CompanyNameSortKey           XsString                                    `xml:"t:CompanyNameSortKey,omitempty"`
	HomeCitySortKey              XsString                                    `xml:"t:HomeCitySortKey,omitempty"`
	WorkCitySortKey              XsString                                    `xml:"t:WorkCitySortKey,omitempty"`
	DisplayNameFirstLastHeader   XsString                                    `xml:"t:DisplayNameFirstLastHeader,omitempty"`
	DisplayNameLastFirstHeader   XsString                                    `xml:"t:DisplayNameLastFirstHeader,omitempty"`
	DisplayName                  XsString                                    `xml:"t:DisplayName,omitempty"`
	DisplayNameFirstLast         XsString                                    `xml:"t:DisplayNameFirstLast,omitempty"`
	DisplayNameLastFirst         XsString                                    `xml:"t:DisplayNameLastFirst,omitempty"`
	FileAs                       XsString                                    `xml:"t:FileAs,omitempty"`
	FileAsId                     XsString                                    `xml:"t:FileAsId,omitempty"`
	DisplayNamePrefix            XsString                                    `xml:"t:DisplayNamePrefix,omitempty"`
	GivenName                    XsString                                    `xml:"t:GivenName,omitempty"`
	MiddleName                   XsString                                    `xml:"t:MiddleName,omitempty"`
	Surname                      XsString                                    `xml:"t:Surname,omitempty"`
	Generation                   XsString                                    `xml:"t:Generation,omitempty"`
	Nickname                     XsString                                    `xml:"t:Nickname,omitempty"`
	YomiCompanyName              XsString                                    `xml:"t:YomiCompanyName,omitempty"`
	YomiFirstName                XsString                                    `xml:"t:YomiFirstName,omitempty"`
	YomiLastName                 XsString                                    `xml:"t:YomiLastName,omitempty"`
	Title                        XsString                                    `xml:"t:Title,omitempty"`
	Department                   XsString                                    `xml:"t:Department,omitempty"`
	CompanyName                  XsString                                    `xml:"t:CompanyName,omitempty"`
	Location                     XsString                                    `xml:"t:Location,omitempty"`
	EmailAddress                 *EmailAddressType                           `xml:"t:EmailAddress,omitempty"`
	EmailAddresses               *ArrayOfEmailAddressesType                  `xml:"t:EmailAddresses,omitempty"`
	PhoneNumber                  *PersonaPhoneNumberType                     `xml:"t:PhoneNumber,omitempty"`
	ImAddress                    XsString                                    `xml:"t:ImAddress,omitempty"`
	HomeCity                     XsString                                    `xml:"t:HomeCity,omitempty"`
	WorkCity                     XsString                                    `xml:"t:WorkCity,omitempty"`
	RelevanceScore               XsInt                                       `xml:"t:RelevanceScore,omitempty"`
	FolderIds                    *ArrayOfFolderIdType                        `xml:"t:FolderIds,omitempty"`
	Attributions                 *ArrayOfPersonaAttributionsType             `xml:"t:Attributions,omitempty"`
	DisplayNames                 *ArrayOfStringAttributedValuesType          `xml:"t:DisplayNames,omitempty"`
	FileAses                     *ArrayOfStringAttributedValuesType          `xml:"t:FileAses,omitempty"`
	FileAsIds                    *ArrayOfStringAttributedValuesType          `xml:"t:FileAsIds,omitempty"`
	DisplayNamePrefixes          *ArrayOfStringAttributedValuesType          `xml:"t:DisplayNamePrefixes,omitempty"`
	GivenNames                   *ArrayOfStringAttributedValuesType          `xml:"t:GivenNames,omitempty"`
	MiddleNames                  *ArrayOfStringAttributedValuesType          `xml:"t:MiddleNames,omitempty"`
	Surnames                     *ArrayOfStringAttributedValuesType          `xml:"t:Surnames,omitempty"`
	Generations                  *ArrayOfStringAttributedValuesType          `xml:"t:Generations,omitempty"`
	Nicknames                    *ArrayOfStringAttributedValuesType          `xml:"t:Nicknames,omitempty"`
	Initials                     *ArrayOfStringAttributedValuesType          `xml:"t:Initials,omitempty"`
	YomiCompanyNames             *ArrayOfStringAttributedValuesType          `xml:"t:YomiCompanyNames,omitempty"`
	YomiFirstNames               *ArrayOfStringAttributedValuesType          `xml:"t:YomiFirstNames,omitempty"`
	YomiLastNames                *ArrayOfStringAttributedValuesType          `xml:"t:YomiLastNames,omitempty"`
	BusinessPhoneNumbers         *ArrayOfPhoneNumberAttributedValuesType     `xml:"t:BusinessPhoneNumbers,omitempty"`
	BusinessPhoneNumbers2        *ArrayOfPhoneNumberAttributedValuesType     `xml:"t:BusinessPhoneNumbers2,omitempty"`
	HomePhones                   *ArrayOfPhoneNumberAttributedValuesType     `xml:"t:HomePhones,omitempty"`
	HomePhones2                  *ArrayOfPhoneNumberAttributedValuesType     `xml:"t:HomePhones2,omitempty"`
	MobilePhones                 *ArrayOfPhoneNumberAttributedValuesType     `xml:"t:MobilePhones,omitempty"`
	MobilePhones2                *ArrayOfPhoneNumberAttributedValuesType     `xml:"t:MobilePhones2,omitempty"`
	AssistantPhoneNumbers        *ArrayOfPhoneNumberAttributedValuesType     `xml:"t:AssistantPhoneNumbers,omitempty"`
	CallbackPhones               *ArrayOfPhoneNumberAttributedValuesType     `xml:"t:CallbackPhones,omitempty"`
	CarPhones                    *ArrayOfPhoneNumberAttributedValuesType     `xml:"t:CarPhones,omitempty"`
	HomeFaxes                    *ArrayOfPhoneNumberAttributedValuesType     `xml:"t:HomeFaxes,omitempty"`
	OrganizationMainPhones       *ArrayOfPhoneNumberAttributedValuesType     `xml:"t:OrganizationMainPhones,omitempty"`
	OtherFaxes                   *ArrayOfPhoneNumberAttributedValuesType     `xml:"t:OtherFaxes,omitempty"`
	OtherTelephones              *ArrayOfPhoneNumberAttributedValuesType     `xml:"t:OtherTelephones,omitempty"`
	OtherPhones2                 *ArrayOfPhoneNumberAttributedValuesType     `xml:"t:OtherPhones2,omitempty"`
	Pagers                       *ArrayOfPhoneNumberAttributedValuesType     `xml:"t:Pagers,omitempty"`
	RadioPhones                  *ArrayOfPhoneNumberAttributedValuesType     `xml:"t:RadioPhones,omitempty"`
	TelexNumbers                 *ArrayOfPhoneNumberAttributedValuesType     `xml:"t:TelexNumbers,omitempty"`
	TTYTDDPhoneNumbers           *ArrayOfPhoneNumberAttributedValuesType     `xml:"t:TTYTDDPhoneNumbers,omitempty"`
	WorkFaxes                    *ArrayOfPhoneNumberAttributedValuesType     `xml:"t:WorkFaxes,omitempty"`
	Emails1                      *ArrayOfEmailAddressAttributedValuesType    `xml:"t:Emails1,omitempty"`
	Emails2                      *ArrayOfEmailAddressAttributedValuesType    `xml:"t:Emails2,omitempty"`
	Emails3                      *ArrayOfEmailAddressAttributedValuesType    `xml:"t:Emails3,omitempty"`
	BusinessHomePages            *ArrayOfStringAttributedValuesType          `xml:"t:BusinessHomePages,omitempty"`
	PersonalHomePages            *ArrayOfStringAttributedValuesType          `xml:"t:PersonalHomePages,omitempty"`
	OfficeLocations              *ArrayOfStringAttributedValuesType          `xml:"t:OfficeLocations,omitempty"`
	ImAddresses                  *ArrayOfStringAttributedValuesType          `xml:"t:ImAddresses,omitempty"`
	ImAddresses2                 *ArrayOfStringAttributedValuesType          `xml:"t:ImAddresses2,omitempty"`
	ImAddresses3                 *ArrayOfStringAttributedValuesType          `xml:"t:ImAddresses3,omitempty"`
	BusinessAddresses            *ArrayOfPostalAddressAttributedValuesType   `xml:"t:BusinessAddresses,omitempty"`
	HomeAddresses                *ArrayOfPostalAddressAttributedValuesType   `xml:"t:HomeAddresses,omitempty"`
	OtherAddresses               *ArrayOfPostalAddressAttributedValuesType   `xml:"t:OtherAddresses,omitempty"`
	Titles                       *ArrayOfStringAttributedValuesType          `xml:"t:Titles,omitempty"`
	Departments                  *ArrayOfStringAttributedValuesType          `xml:"t:Departments,omitempty"`
	CompanyNames                 *ArrayOfStringAttributedValuesType          `xml:"t:CompanyNames,omitempty"`
	Managers                     *ArrayOfStringAttributedValuesType          `xml:"t:Managers,omitempty"`
	AssistantNames               *ArrayOfStringAttributedValuesType          `xml:"t:AssistantNames,omitempty"`
	Professions                  *ArrayOfStringAttributedValuesType          `xml:"t:Professions,omitempty"`
	SpouseNames                  *ArrayOfStringAttributedValuesType          `xml:"t:SpouseNames,omitempty"`
	Children                     *ArrayOfStringArrayAttributedValuesType     `xml:"t:Children,omitempty"`
	Schools                      *ArrayOfStringAttributedValuesType          `xml:"t:Schools,omitempty"`
	Hobbies                      *ArrayOfStringAttributedValuesType          `xml:"t:Hobbies,omitempty"`
	WeddingAnniversaries         *ArrayOfStringAttributedValuesType          `xml:"t:WeddingAnniversaries,omitempty"`
	Birthdays                    *ArrayOfStringAttributedValuesType          `xml:"t:Birthdays,omitempty"`
	Locations                    *ArrayOfStringAttributedValuesType          `xml:"t:Locations,omitempty"`
	InlineLinks                  *ArrayOfStringAttributedValuesType          `xml:"t:InlineLinks,omitempty"`
	ItemLinkIds                  *ArrayOfStringArrayAttributedValuesType     `xml:"t:ItemLinkIds,omitempty"`
	HasActiveDeals               XsString                                    `xml:"t:HasActiveDeals,omitempty"`
	IsBusinessContact            XsString                                    `xml:"t:IsBusinessContact,omitempty"`
	AttributedHasActiveDeals     *ArrayOfStringAttributedValuesType          `xml:"t:AttributedHasActiveDeals,omitempty"`
	AttributedIsBusinessContact  *ArrayOfStringAttributedValuesType          `xml:"t:AttributedIsBusinessContact,omitempty"`
	SourceMailboxGuids           *ArrayOfStringAttributedValuesType          `xml:"t:SourceMailboxGuids,omitempty"`
	LastContactedDate            XsDateTime                                  `xml:"t:LastContactedDate,omitempty"`
	ExtendedProperties           *ArrayOfExtendedPropertyAttributedValueType `xml:"t:ExtendedProperties,omitempty"`
	ExternalDirectoryObjectId    XsString                                    `xml:"t:ExternalDirectoryObjectId,omitempty"`
	MapiEntryId                  XsString                                    `xml:"t:MapiEntryId,omitempty"`
	MapiEmailAddress             XsString                                    `xml:"t:MapiEmailAddress,omitempty"`
	MapiAddressType              XsString                                    `xml:"t:MapiAddressType,omitempty"`
	MapiSearchKey                XsString                                    `xml:"t:MapiSearchKey,omitempty"`
	MapiTransmittableDisplayName XsString                                    `xml:"t:MapiTransmittableDisplayName,omitempty"`
	MapiSendRichInfo             XsBoolean                                   `xml:"t:MapiSendRichInfo,omitempty"`
}

type ArrayOfBodyContentAttributedValuesType struct {
	BodyContentAttributedValue []*BodyContentAttributedValueType `xml:"t:BodyContentAttributedValue,omitempty"`
}

type BodyContentAttributedValueType struct {
	Value        *BodyContentType              `xml:"t:Value,omitempty"`
	Attributions *ArrayOfValueAttributionsType `xml:"t:Attributions,omitempty"`
}

type BodyContentType struct {
	Value    XsString     `xml:"t:Value,omitempty"`
	BodyType BodyTypeType `xml:"t:BodyType,omitempty"`
}

type ArrayOfValueAttributionsType struct {
	Attribution []XsString `xml:"t:Attribution,omitempty"`
}

type PersonaPhoneNumberType struct {
	Number XsString `xml:"t:Number,omitempty"`
	Type   XsString `xml:"t:Type,omitempty"`
}

type ArrayOfFolderIdType struct {
	FolderId []*FolderIdType `xml:"t:FolderId,omitempty"`
}

type ArrayOfPersonaAttributionsType struct {
	Attribution []*PersonaAttributionType `xml:"t:Attribution,omitempty"`
}

type PersonaAttributionType struct {
	Id             XsString      `xml:"t:Id,omitempty"`
	SourceId       *ItemIdType   `xml:"t:SourceId,omitempty"`
	DisplayName    XsString      `xml:"t:DisplayName,omitempty"`
	IsWritable     XsBoolean     `xml:"t:IsWritable,omitempty"`
	IsQuickContact XsBoolean     `xml:"t:IsQuickContact,omitempty"`
	IsHidden       XsBoolean     `xml:"t:IsHidden,omitempty"`
	FolderId       *FolderIdType `xml:"t:FolderId,omitempty"`
}

type ArrayOfStringAttributedValuesType struct {
	StringAttributedValue []*StringAttributedValueType `xml:"t:StringAttributedValue,omitempty"`
}

type StringAttributedValueType struct {
	Value        XsString                      `xml:"t:Value,omitempty"`
	Attributions *ArrayOfValueAttributionsType `xml:"t:Attributions,omitempty"`
}

type ArrayOfPhoneNumberAttributedValuesType struct {
	PhoneNumberAttributedValue []*PhoneNumberAttributedValueType `xml:"t:PhoneNumberAttributedValue,omitempty"`
}

type PhoneNumberAttributedValueType struct {
	Value        *PersonaPhoneNumberType       `xml:"t:Value,omitempty"`
	Attributions *ArrayOfValueAttributionsType `xml:"t:Attributions,omitempty"`
}

type ArrayOfEmailAddressAttributedValuesType struct {
	EmailAddressAttributedValue []*EmailAddressAttributedValueType `xml:"t:EmailAddressAttributedValue,omitempty"`
}

type EmailAddressAttributedValueType struct {
	Value        *EmailAddressType             `xml:"t:Value,omitempty"`
	Attributions *ArrayOfValueAttributionsType `xml:"t:Attributions,omitempty"`
}

type ArrayOfPostalAddressAttributedValuesType struct {
	PostalAddressAttributedValue []*PostalAddressAttributedValueType `xml:"t:PostalAddressAttributedValue,omitempty"`
}

type PostalAddressAttributedValueType struct {
	Value        *PersonaPostalAddressType     `xml:"t:Value,omitempty"`
	Attributions *ArrayOfValueAttributionsType `xml:"t:Attributions,omitempty"`
}

type ArrayOfStringArrayAttributedValuesType struct {
	StringArrayAttributedValue []*StringArrayAttributedValueType `xml:"t:StringArrayAttributedValue,omitempty"`
}

type StringArrayAttributedValueType struct {
	Values       *ArrayOfStringValueType       `xml:"t:Values,omitempty"`
	Attributions *ArrayOfValueAttributionsType `xml:"t:Attributions,omitempty"`
}

type ArrayOfStringValueType struct {
	Value []XsString `xml:"t:Value,omitempty"`
}

type ArrayOfExtendedPropertyAttributedValueType struct {
	ExtendedPropertyAttributedValue []*ExtendedPropertyAttributedValueType `xml:"t:ExtendedPropertyAttributedValue,omitempty"`
}

type ExtendedPropertyAttributedValueType struct {
	Value        *ExtendedPropertyType         `xml:"t:Value,omitempty"`
	Attributions *ArrayOfValueAttributionsType `xml:"t:Attributions,omitempty"`
}

type ConversationResponseType struct {
	ConversationId    *ItemIdType                   `xml:"t:ConversationId,omitempty"`
	SyncState         XsBase64Binary                `xml:"t:SyncState,omitempty"`
	ConversationNodes *ArrayOfConversationNodesType `xml:"t:ConversationNodes,omitempty"`
	CanDelete         XsBoolean                     `xml:"t:CanDelete,omitempty"`
}

type ArrayOfConversationNodesType struct {
	ConversationNode []*ConversationNodeType `xml:"t:ConversationNode,omitempty"`
}

type ConversationNodeType struct {
	InternetMessageId       XsString                     `xml:"t:InternetMessageId,omitempty"`
	ParentInternetMessageId XsString                     `xml:"t:ParentInternetMessageId,omitempty"`
	Items                   *NonEmptyArrayOfAllItemsType `xml:"t:Items,omitempty"`
}

type ArrayOfRetentionPolicyTagsType struct {
	RetentionPolicyTag []*RetentionPolicyTagType `xml:"t:RetentionPolicyTag,omitempty"`
}

type RetentionPolicyTagType struct {
	DisplayName         XsString            `xml:"t:DisplayName,omitempty"`
	RetentionId         GuidType            `xml:"t:RetentionId,omitempty"`
	RetentionPeriod     XsInt               `xml:"t:RetentionPeriod,omitempty"`
	Type                ElcFolderType       `xml:"t:Type,omitempty"`
	RetentionAction     RetentionActionType `xml:"t:RetentionAction,omitempty"`
	Description         XsString            `xml:"t:Description,omitempty"`
	IsVisible           XsBoolean           `xml:"t:IsVisible,omitempty"`
	OptedInto           XsBoolean           `xml:"t:OptedInto,omitempty"`
	IsArchive           XsBoolean           `xml:"t:IsArchive,omitempty"`
	ParentLabelIdentity XsString            `xml:"t:ParentLabelIdentity,omitempty"`
	Priority            XsInt               `xml:"t:Priority,omitempty"`
}

type NonEmptyArrayOfTimeZoneIdType struct {
	Id []XsString `xml:"t:Id,omitempty"`
}

type FolderResponseShapeType struct {
	BaseShape            DefaultShapeNamesType              `xml:"t:BaseShape,omitempty"`
	AdditionalProperties *NonEmptyArrayOfPathsToElementType `xml:"t:AdditionalProperties,omitempty"`
}

type NonEmptyArrayOfPathsToElementType struct {
	Path []BasePathToElementType `xml:"t:Path,omitempty"`
}

type IndexedPageViewType struct {
	BasePagingType `xml:",omitempty"`
	Offset         XsInt              `xml:"Offset,attr"`
	BasePoint      IndexBasePointType `xml:"BasePoint,attr"`
}

type BasePagingType struct {
	MaxEntriesReturned XsInt `xml:"MaxEntriesReturned,attr"`
}

type FractionalPageViewType struct {
	BasePagingType `xml:",omitempty"`
	Numerator      XsInt `xml:"Numerator,attr"`
	Denominator    XsInt `xml:"Denominator,attr"`
}

type TimeZoneContextType struct {
	TimeZoneDefinition *TimeZoneDefinitionType `xml:"t:TimeZoneDefinition,omitempty"`
}

type ManagementRoleType struct {
	UserRoles        *NonEmptyArrayOfRoleType `xml:"t:UserRoles,omitempty"`
	ApplicationRoles *NonEmptyArrayOfRoleType `xml:"t:ApplicationRoles,omitempty"`
}

type NonEmptyArrayOfRoleType struct {
	Role []XsString `xml:"t:Role,omitempty"`
}

type ItemResponseShapeType struct {
	BaseShape                 DefaultShapeNamesType              `xml:"t:BaseShape,omitempty"`
	IncludeMimeContent        XsBoolean                          `xml:"t:IncludeMimeContent,omitempty"`
	BodyType                  BodyTypeResponseType               `xml:"t:BodyType,omitempty"`
	UniqueBodyType            BodyTypeResponseType               `xml:"t:UniqueBodyType,omitempty"`
	NormalizedBodyType        BodyTypeResponseType               `xml:"t:NormalizedBodyType,omitempty"`
	FilterHtmlContent         XsBoolean                          `xml:"t:FilterHtmlContent,omitempty"`
	ConvertHtmlCodePageToUTF8 XsBoolean                          `xml:"t:ConvertHtmlCodePageToUTF8,omitempty"`
	InlineImageUrlTemplate    XsString                           `xml:"t:InlineImageUrlTemplate,omitempty"`
	BlockExternalImages       XsBoolean                          `xml:"t:BlockExternalImages,omitempty"`
	AddBlankTargetToLinks     XsBoolean                          `xml:"t:AddBlankTargetToLinks,omitempty"`
	MaximumBodySize           XsInt                              `xml:"t:MaximumBodySize,omitempty"`
	AdditionalProperties      *NonEmptyArrayOfPathsToElementType `xml:"t:AdditionalProperties,omitempty"`
}

type NonEmptyArrayOfFieldOrdersType struct {
	FieldOrder []*FieldOrderType `xml:"t:FieldOrder,omitempty"`
}

type FieldOrderType struct {
	Order SortDirectionType     `xml:"Order,attr"`
	Path  BasePathToElementType `xml:"t:Path,omitempty"`
}

type SeekToConditionPageViewType struct {
	BasePagingType `xml:",omitempty"`
	BasePoint      IndexBasePointType `xml:"BasePoint,attr"`
	Condition      *RestrictionType   `xml:"t:Condition,omitempty"`
}

type CalendarViewType struct {
	BasePagingType `xml:",omitempty"`
	StartDate      XsDateTime `xml:"StartDate,attr"`
	EndDate        XsDateTime `xml:"EndDate,attr"`
}

type ContactsViewType struct {
	BasePagingType `xml:",omitempty"`
	InitialName    XsString `xml:"InitialName,attr"`
	FinalName      XsString `xml:"FinalName,attr"`
}

type GroupByType struct {
	BaseGroupByType      `xml:",omitempty"`
	AggregateOn          *AggregateOnType          `xml:"t:AggregateOn,omitempty"`
	UseCollapsibleGroups XsBoolean                 `xml:"t:UseCollapsibleGroups,omitempty"`
	ItemsPerGroup        XsNonNegativeInteger      `xml:"t:ItemsPerGroup,omitempty"`
	MaxItemsPerGroup     XsNonNegativeInteger      `xml:"t:MaxItemsPerGroup,omitempty"`
	GroupsToExpand       *ArrayOfGroupIdType       `xml:"t:GroupsToExpand,omitempty"`
	FieldURI             *PathToUnindexedFieldType `xml:"t:FieldURI,omitempty"`
	IndexedFieldURI      *PathToIndexedFieldType   `xml:"t:IndexedFieldURI,omitempty"`
	ExtendedFieldURI     *PathToExtendedFieldType  `xml:"t:ExtendedFieldURI,omitempty"`
}

type BaseGroupByType struct {
	Order SortDirectionType `xml:"Order,attr"`
}

type AggregateOnType struct {
	Aggregate        AggregateType             `xml:"Aggregate,attr"`
	FieldURI         *PathToUnindexedFieldType `xml:"t:FieldURI,omitempty"`
	IndexedFieldURI  *PathToIndexedFieldType   `xml:"t:IndexedFieldURI,omitempty"`
	ExtendedFieldURI *PathToExtendedFieldType  `xml:"t:ExtendedFieldURI,omitempty"`
}

type PathToUnindexedFieldType struct {
	BasePathToElementType `xml:",omitempty"`
	FieldURI              UnindexedFieldURIType `xml:"FieldURI,attr"`
}

type PathToIndexedFieldType struct {
	BasePathToElementType `xml:",omitempty"`
	FieldURI              DictionaryURIType `xml:"FieldURI,attr"`
	FieldIndex            XsString          `xml:"FieldIndex,attr"`
}

type ArrayOfGroupIdType struct {
	GroupId []XsBase64Binary `xml:"t:GroupId,omitempty"`
}

type DistinguishedGroupByType struct {
	BaseGroupByType `xml:",omitempty"`
	StandardGroupBy StandardGroupByType `xml:"t:StandardGroupBy,omitempty"`
}

type NonEmptyArrayOfUploadItemsType struct {
	Item []*UploadItemType `xml:"t:Item,omitempty"`
}

type UploadItemType struct {
	CreateAction   CreateActionType `xml:"CreateAction,attr"`
	IsAssociated   XsBoolean        `xml:"IsAssociated,attr"`
	ParentFolderId *FolderIdType    `xml:"t:ParentFolderId,omitempty"`
	ItemId         *ItemIdType      `xml:"t:ItemId,omitempty"`
	Data           XsBase64Binary   `xml:"t:Data,omitempty"`
}

type NonEmptyArrayOfAlternateIdsType struct {
	AlternateId                 []*AlternateIdType                 `xml:"t:AlternateId,omitempty"`
	AlternatePublicFolderId     []*AlternatePublicFolderIdType     `xml:"t:AlternatePublicFolderId,omitempty"`
	AlternatePublicFolderItemId []*AlternatePublicFolderItemIdType `xml:"t:AlternatePublicFolderItemId,omitempty"`
}

type AlternateIdType struct {
	AlternateIdBaseType `xml:",omitempty"`
	Id                  XsString           `xml:"Id,attr"`
	Mailbox             NonEmptyStringType `xml:"Mailbox,attr"`
	IsArchive           XsBoolean          `xml:"IsArchive,attr"`
}

type AlternatePublicFolderIdType struct {
	AlternateIdBaseType `xml:",omitempty"`
	FolderId            XsString `xml:"FolderId,attr"`
}

type AlternatePublicFolderItemIdType struct {
	AlternatePublicFolderIdType `xml:",omitempty"`
	ItemId                      XsString `xml:"ItemId,attr"`
}

type NonEmptyArrayOfFoldersType struct {
	Folder         []*FolderType         `xml:"t:Folder,omitempty"`
	CalendarFolder []*CalendarFolderType `xml:"t:CalendarFolder,omitempty"`
	ContactsFolder []*ContactsFolderType `xml:"t:ContactsFolder,omitempty"`
	SearchFolder   []*SearchFolderType   `xml:"t:SearchFolder,omitempty"`
	TasksFolder    []*TasksFolderType    `xml:"t:TasksFolder,omitempty"`
}

type NonEmptyArrayOfFolderChangesType struct {
	FolderChange []*FolderChangeType `xml:"t:FolderChange,omitempty"`
}

type FolderChangeType struct {
	Updates               *NonEmptyArrayOfFolderChangeDescriptionsType `xml:"t:Updates,omitempty"`
	FolderId              *FolderIdType                                `xml:"t:FolderId,omitempty"`
	DistinguishedFolderId *DistinguishedFolderIdType                   `xml:"t:DistinguishedFolderId,omitempty"`
}

type NonEmptyArrayOfFolderChangeDescriptionsType struct {
	AppendToFolderField []*AppendToFolderFieldType `xml:"t:AppendToFolderField,omitempty"`
	SetFolderField      []*SetFolderFieldType      `xml:"t:SetFolderField,omitempty"`
	DeleteFolderField   []*DeleteFolderFieldType   `xml:"t:DeleteFolderField,omitempty"`
}

type AppendToFolderFieldType struct {
	FolderChangeDescriptionType `xml:",omitempty"`
	Folder                      *FolderType         `xml:"t:Folder,omitempty"`
	CalendarFolder              *CalendarFolderType `xml:"t:CalendarFolder,omitempty"`
	ContactsFolder              *ContactsFolderType `xml:"t:ContactsFolder,omitempty"`
	SearchFolder                *SearchFolderType   `xml:"t:SearchFolder,omitempty"`
	TasksFolder                 *TasksFolderType    `xml:"t:TasksFolder,omitempty"`
}

type FolderChangeDescriptionType struct {
	ChangeDescriptionType `xml:",omitempty"`
}

type ChangeDescriptionType struct {
	Path BasePathToElementType `xml:"t:Path,omitempty"`
}

type SetFolderFieldType struct {
	FolderChangeDescriptionType `xml:",omitempty"`
	Folder                      *FolderType         `xml:"t:Folder,omitempty"`
	CalendarFolder              *CalendarFolderType `xml:"t:CalendarFolder,omitempty"`
	ContactsFolder              *ContactsFolderType `xml:"t:ContactsFolder,omitempty"`
	SearchFolder                *SearchFolderType   `xml:"t:SearchFolder,omitempty"`
	TasksFolder                 *TasksFolderType    `xml:"t:TasksFolder,omitempty"`
}

type DeleteFolderFieldType struct {
	FolderChangeDescriptionType `xml:",omitempty"`
}

type PullSubscriptionRequestType struct {
	BaseSubscriptionRequestType `xml:",omitempty"`
	Timeout                     SubscriptionTimeoutType `xml:"t:Timeout,omitempty"`
}

type BaseSubscriptionRequestType struct {
	SubscribeToAllFolders XsBoolean                                  `xml:"SubscribeToAllFolders,attr"`
	FolderIds             *NonEmptyArrayOfBaseFolderIdsType          `xml:"t:FolderIds,omitempty"`
	EventTypes            *NonEmptyArrayOfNotificationEventTypesType `xml:"t:EventTypes,omitempty"`
	Watermark             WatermarkType                              `xml:"t:Watermark,omitempty"`
}

type NonEmptyArrayOfNotificationEventTypesType struct {
	EventType []NotificationEventTypeType `xml:"t:EventType,omitempty"`
}

type PushSubscriptionRequestType struct {
	BaseSubscriptionRequestType `xml:",omitempty"`
	StatusFrequency             SubscriptionStatusFrequencyType `xml:"t:StatusFrequency,omitempty"`
	URL                         XsString                        `xml:"t:URL,omitempty"`
	CallerData                  XsString                        `xml:"t:CallerData,omitempty"`
}

type StreamingSubscriptionRequestType struct {
	SubscribeToAllFolders XsBoolean                                  `xml:"SubscribeToAllFolders,attr"`
	FolderIds             *NonEmptyArrayOfBaseFolderIdsType          `xml:"t:FolderIds,omitempty"`
	EventTypes            *NonEmptyArrayOfNotificationEventTypesType `xml:"t:EventTypes,omitempty"`
}

type ArrayOfBaseItemIdsType struct {
	ItemId []*ItemIdType `xml:"t:ItemId,omitempty"`
}

type NonEmptyArrayOfFolderNamesType struct {
	FolderName []XsString `xml:"t:FolderName,omitempty"`
}

type NonEmptyArrayOfBaseItemIdsType struct {
	ItemId                      []*ItemIdType                      `xml:"t:ItemId,omitempty"`
	OccurrenceItemId            []*OccurrenceItemIdType            `xml:"t:OccurrenceItemId,omitempty"`
	RecurringMasterItemId       []*RecurringMasterItemIdType       `xml:"t:RecurringMasterItemId,omitempty"`
	RecurringMasterItemIdRanges []*RecurringMasterItemIdRangesType `xml:"t:RecurringMasterItemIdRanges,omitempty"`
}

type OccurrenceItemIdType struct {
	BaseItemIdType    `xml:",omitempty"`
	RecurringMasterId DerivedItemIdType `xml:"RecurringMasterId,attr"`
	ChangeKey         XsString          `xml:"ChangeKey,attr"`
	InstanceIndex     XsInt             `xml:"InstanceIndex,attr"`
}

type RecurringMasterItemIdType struct {
	BaseItemIdType `xml:",omitempty"`
	OccurrenceId   DerivedItemIdType `xml:"OccurrenceId,attr"`
	ChangeKey      XsString          `xml:"ChangeKey,attr"`
}

type RecurringMasterItemIdRangesType struct {
	ItemIdType `xml:",omitempty"`
	Ranges     *ArrayOfOccurrenceRangesType `xml:"t:Ranges,omitempty"`
}

type ArrayOfOccurrenceRangesType struct {
	Range []*OccurrencesRangeType `xml:"t:Range,omitempty"`
}

type OccurrencesRangeType struct {
	Start                    XsDateTime `xml:"Start,attr"`
	End                      XsDateTime `xml:"End,attr"`
	Count                    XsInt      `xml:"Count,attr"`
	CompareOriginalStartTime XsBoolean  `xml:"CompareOriginalStartTime,attr"`
}

type NonEmptyArrayOfItemChangesType struct {
	ItemChange []*ItemChangeType `xml:"t:ItemChange,omitempty"`
}

type ItemChangeType struct {
	Updates               *NonEmptyArrayOfItemChangeDescriptionsType `xml:"t:Updates,omitempty"`
	CalendarActivityData  *CalendarActivityDataType                  `xml:"t:CalendarActivityData,omitempty"`
	ItemId                *ItemIdType                                `xml:"t:ItemId,omitempty"`
	OccurrenceItemId      *OccurrenceItemIdType                      `xml:"t:OccurrenceItemId,omitempty"`
	RecurringMasterItemId *RecurringMasterItemIdType                 `xml:"t:RecurringMasterItemId,omitempty"`
}

type NonEmptyArrayOfItemChangeDescriptionsType struct {
	AppendToItemField []*AppendToItemFieldType `xml:"t:AppendToItemField,omitempty"`
	SetItemField      []*SetItemFieldType      `xml:"t:SetItemField,omitempty"`
	DeleteItemField   []*DeleteItemFieldType   `xml:"t:DeleteItemField,omitempty"`
}

type AppendToItemFieldType struct {
	ItemChangeDescriptionType `xml:",omitempty"`
	Item                      *ItemType                       `xml:"t:Item,omitempty"`
	Message                   *MessageType                    `xml:"t:Message,omitempty"`
	SharingMessage            *SharingMessageType             `xml:"t:SharingMessage,omitempty"`
	CalendarItem              *CalendarItemType               `xml:"t:CalendarItem,omitempty"`
	Contact                   *ContactItemType                `xml:"t:Contact,omitempty"`
	DistributionList          *DistributionListType           `xml:"t:DistributionList,omitempty"`
	MeetingMessage            *MeetingMessageType             `xml:"t:MeetingMessage,omitempty"`
	MeetingRequest            *MeetingRequestMessageType      `xml:"t:MeetingRequest,omitempty"`
	MeetingResponse           *MeetingResponseMessageType     `xml:"t:MeetingResponse,omitempty"`
	MeetingCancellation       *MeetingCancellationMessageType `xml:"t:MeetingCancellation,omitempty"`
	Task                      *TaskType                       `xml:"t:Task,omitempty"`
	PostItem                  *PostItemType                   `xml:"t:PostItem,omitempty"`
	RoleMember                *RoleMemberItemType             `xml:"t:RoleMember,omitempty"`
	Network                   *NetworkItemType                `xml:"t:Network,omitempty"`
	Person                    *AbchPersonItemType             `xml:"t:Person,omitempty"`
}

type ItemChangeDescriptionType struct {
	ChangeDescriptionType `xml:",omitempty"`
}

type SetItemFieldType struct {
	ItemChangeDescriptionType `xml:",omitempty"`
	Item                      *ItemType                       `xml:"t:Item,omitempty"`
	Message                   *MessageType                    `xml:"t:Message,omitempty"`
	SharingMessage            *SharingMessageType             `xml:"t:SharingMessage,omitempty"`
	CalendarItem              *CalendarItemType               `xml:"t:CalendarItem,omitempty"`
	Contact                   *ContactItemType                `xml:"t:Contact,omitempty"`
	DistributionList          *DistributionListType           `xml:"t:DistributionList,omitempty"`
	MeetingMessage            *MeetingMessageType             `xml:"t:MeetingMessage,omitempty"`
	MeetingRequest            *MeetingRequestMessageType      `xml:"t:MeetingRequest,omitempty"`
	MeetingResponse           *MeetingResponseMessageType     `xml:"t:MeetingResponse,omitempty"`
	MeetingCancellation       *MeetingCancellationMessageType `xml:"t:MeetingCancellation,omitempty"`
	Task                      *TaskType                       `xml:"t:Task,omitempty"`
	PostItem                  *PostItemType                   `xml:"t:PostItem,omitempty"`
	RoleMember                *RoleMemberItemType             `xml:"t:RoleMember,omitempty"`
	Network                   *NetworkItemType                `xml:"t:Network,omitempty"`
	Person                    *AbchPersonItemType             `xml:"t:Person,omitempty"`
}

type DeleteItemFieldType struct {
	ItemChangeDescriptionType `xml:",omitempty"`
}

type NonEmptyArrayOfRequestAttachmentIdsType struct {
	AttachmentId []*RequestAttachmentIdType `xml:"t:AttachmentId,omitempty"`
}

type AttachmentResponseShapeType struct {
	IncludeMimeContent   XsBoolean                          `xml:"t:IncludeMimeContent,omitempty"`
	BodyType             BodyTypeResponseType               `xml:"t:BodyType,omitempty"`
	FilterHtmlContent    XsBoolean                          `xml:"t:FilterHtmlContent,omitempty"`
	AdditionalProperties *NonEmptyArrayOfPathsToElementType `xml:"t:AdditionalProperties,omitempty"`
}

type NonEmptyArrayOfClientAccessTokenRequestsType struct {
	TokenRequest []*ClientAccessTokenRequestType `xml:"t:TokenRequest,omitempty"`
}

type ClientAccessTokenRequestType struct {
	Id                    XsString                  `xml:"t:Id,omitempty"`
	TokenType             ClientAccessTokenTypeType `xml:"t:TokenType,omitempty"`
	Scope                 XsString                  `xml:"t:Scope,omitempty"`
	ResourceUri           XsString                  `xml:"t:ResourceUri,omitempty"`
	IdentityAndEwsTokenId XsString                  `xml:"t:IdentityAndEwsTokenId,omitempty"`
	RequestedCapabilities XsString                  `xml:"t:RequestedCapabilities,omitempty"`
	SupportsSharedFolders XsBoolean                 `xml:"t:SupportsSharedFolders,omitempty"`
}

type ArrayOfUserIdType struct {
	UserId []*UserIdType `xml:"t:UserId,omitempty"`
}

type DelegateUserType struct {
	UserId                         *UserIdType              `xml:"t:UserId,omitempty"`
	DelegatePermissions            *DelegatePermissionsType `xml:"t:DelegatePermissions,omitempty"`
	ReceiveCopiesOfMeetingMessages XsBoolean                `xml:"t:ReceiveCopiesOfMeetingMessages,omitempty"`
	ViewPrivateItems               XsBoolean                `xml:"t:ViewPrivateItems,omitempty"`
}

type DelegatePermissionsType struct {
	CalendarFolderPermissionLevel DelegateFolderPermissionLevelType `xml:"t:CalendarFolderPermissionLevel,omitempty"`
	TasksFolderPermissionLevel    DelegateFolderPermissionLevelType `xml:"t:TasksFolderPermissionLevel,omitempty"`
	InboxFolderPermissionLevel    DelegateFolderPermissionLevelType `xml:"t:InboxFolderPermissionLevel,omitempty"`
	ContactsFolderPermissionLevel DelegateFolderPermissionLevelType `xml:"t:ContactsFolderPermissionLevel,omitempty"`
	NotesFolderPermissionLevel    DelegateFolderPermissionLevelType `xml:"t:NotesFolderPermissionLevel,omitempty"`
	JournalFolderPermissionLevel  DelegateFolderPermissionLevelType `xml:"t:JournalFolderPermissionLevel,omitempty"`
}

type ArrayOfDelegateUserType struct {
	DelegateUser []*DelegateUserType `xml:"t:DelegateUser,omitempty"`
}

type SerializableTimeZone struct {
	Bias         XsInt                     `xml:"t:Bias,omitempty"`
	StandardTime *SerializableTimeZoneTime `xml:"t:StandardTime,omitempty"`
	DaylightTime *SerializableTimeZoneTime `xml:"t:DaylightTime,omitempty"`
}

type SerializableTimeZoneTime struct {
	Bias      XsInt         `xml:"t:Bias,omitempty"`
	Time      XsString      `xml:"t:Time,omitempty"`
	DayOrder  XsShort       `xml:"t:DayOrder,omitempty"`
	Month     XsShort       `xml:"t:Month,omitempty"`
	DayOfWeek DayOfWeekType `xml:"t:DayOfWeek,omitempty"`
	Year      XsString      `xml:"t:Year,omitempty"`
}

type ArrayOfMailboxData struct {
	MailboxData []*MailboxData `xml:"t:MailboxData,omitempty"`
}

type MailboxData struct {
	Email            *EmailAddress       `xml:"t:Email,omitempty"`
	AttendeeType     MeetingAttendeeType `xml:"t:AttendeeType,omitempty"`
	ExcludeConflicts XsBoolean           `xml:"t:ExcludeConflicts,omitempty"`
}

type EmailAddress struct {
	Name        XsString `xml:"t:Name,omitempty"`
	Address     XsString `xml:"t:Address,omitempty"`
	RoutingType XsString `xml:"t:RoutingType,omitempty"`
}

type FreeBusyViewOptionsType struct {
	TimeWindow                      *Duration        `xml:"t:TimeWindow,omitempty"`
	MergedFreeBusyIntervalInMinutes XsInt            `xml:"t:MergedFreeBusyIntervalInMinutes,omitempty"`
	RequestedView                   FreeBusyViewType `xml:"t:RequestedView,omitempty"`
}

type Duration struct {
	StartTime XsDateTime `xml:"t:StartTime,omitempty"`
	EndTime   XsDateTime `xml:"t:EndTime,omitempty"`
}

type SuggestionsViewOptionsType struct {
	GoodThreshold                  XsInt             `xml:"t:GoodThreshold,omitempty"`
	MaximumResultsByDay            XsInt             `xml:"t:MaximumResultsByDay,omitempty"`
	MaximumNonWorkHourResultsByDay XsInt             `xml:"t:MaximumNonWorkHourResultsByDay,omitempty"`
	MeetingDurationInMinutes       XsInt             `xml:"t:MeetingDurationInMinutes,omitempty"`
	MinimumSuggestionQuality       SuggestionQuality `xml:"t:MinimumSuggestionQuality,omitempty"`
	DetailedSuggestionsWindow      *Duration         `xml:"t:DetailedSuggestionsWindow,omitempty"`
	CurrentMeetingTime             XsDateTime        `xml:"t:CurrentMeetingTime,omitempty"`
	GlobalObjectId                 XsString          `xml:"t:GlobalObjectId,omitempty"`
}

type FreeBusyView struct {
	FreeBusyViewType   FreeBusyViewType      `xml:"t:FreeBusyViewType,omitempty"`
	MergedFreeBusy     XsString              `xml:"t:MergedFreeBusy,omitempty"`
	CalendarEventArray *ArrayOfCalendarEvent `xml:"t:CalendarEventArray,omitempty"`
	WorkingHours       *WorkingHours         `xml:"t:WorkingHours,omitempty"`
}

type ArrayOfCalendarEvent struct {
	CalendarEvent []*CalendarEvent `xml:"t:CalendarEvent,omitempty"`
}

type CalendarEvent struct {
	StartTime            XsDateTime            `xml:"t:StartTime,omitempty"`
	EndTime              XsDateTime            `xml:"t:EndTime,omitempty"`
	BusyType             LegacyFreeBusyType    `xml:"t:BusyType,omitempty"`
	CalendarEventDetails *CalendarEventDetails `xml:"t:CalendarEventDetails,omitempty"`
}

type CalendarEventDetails struct {
	ID            XsString  `xml:"t:ID,omitempty"`
	Subject       XsString  `xml:"t:Subject,omitempty"`
	Location      XsString  `xml:"t:Location,omitempty"`
	IsMeeting     XsBoolean `xml:"t:IsMeeting,omitempty"`
	IsRecurring   XsBoolean `xml:"t:IsRecurring,omitempty"`
	IsException   XsBoolean `xml:"t:IsException,omitempty"`
	IsReminderSet XsBoolean `xml:"t:IsReminderSet,omitempty"`
	IsPrivate     XsBoolean `xml:"t:IsPrivate,omitempty"`
}

type WorkingHours struct {
	TimeZone           *SerializableTimeZone `xml:"t:TimeZone,omitempty"`
	WorkingPeriodArray *ArrayOfWorkingPeriod `xml:"t:WorkingPeriodArray,omitempty"`
}

type ArrayOfWorkingPeriod struct {
	WorkingPeriod []*WorkingPeriod `xml:"t:WorkingPeriod,omitempty"`
}

type WorkingPeriod struct {
	DayOfWeek          DaysOfWeekType `xml:"t:DayOfWeek,omitempty"`
	StartTimeInMinutes XsInt          `xml:"t:StartTimeInMinutes,omitempty"`
	EndTimeInMinutes   XsInt          `xml:"t:EndTimeInMinutes,omitempty"`
}

type ArrayOfSuggestionDayResult struct {
	SuggestionDayResult []*SuggestionDayResult `xml:"t:SuggestionDayResult,omitempty"`
}

type SuggestionDayResult struct {
	Date            XsDateTime         `xml:"t:Date,omitempty"`
	DayQuality      SuggestionQuality  `xml:"t:DayQuality,omitempty"`
	SuggestionArray *ArrayOfSuggestion `xml:"t:SuggestionArray,omitempty"`
}

type ArrayOfSuggestion struct {
	Suggestion []*Suggestion `xml:"t:Suggestion,omitempty"`
}

type Suggestion struct {
	MeetingTime               XsDateTime                   `xml:"t:MeetingTime,omitempty"`
	IsWorkTime                XsBoolean                    `xml:"t:IsWorkTime,omitempty"`
	SuggestionQuality         SuggestionQuality            `xml:"t:SuggestionQuality,omitempty"`
	AttendeeConflictDataArray *ArrayOfAttendeeConflictData `xml:"t:AttendeeConflictDataArray,omitempty"`
}

type ArrayOfAttendeeConflictData struct {
	UnknownAttendeeConflictData     *UnknownAttendeeConflictData     `xml:"t:UnknownAttendeeConflictData,omitempty"`
	IndividualAttendeeConflictData  *IndividualAttendeeConflictData  `xml:"t:IndividualAttendeeConflictData,omitempty"`
	TooBigGroupAttendeeConflictData *TooBigGroupAttendeeConflictData `xml:"t:TooBigGroupAttendeeConflictData,omitempty"`
	GroupAttendeeConflictData       *GroupAttendeeConflictData       `xml:"t:GroupAttendeeConflictData,omitempty"`
}

type UnknownAttendeeConflictData struct {
	AttendeeConflictData `xml:",omitempty"`
}

type AttendeeConflictData struct {
}

type IndividualAttendeeConflictData struct {
	AttendeeConflictData `xml:",omitempty"`
	BusyType             LegacyFreeBusyType `xml:"t:BusyType,omitempty"`
}

type TooBigGroupAttendeeConflictData struct {
	AttendeeConflictData `xml:",omitempty"`
}

type GroupAttendeeConflictData struct {
	AttendeeConflictData        `xml:",omitempty"`
	NumberOfMembers             XsInt `xml:"t:NumberOfMembers,omitempty"`
	NumberOfMembersAvailable    XsInt `xml:"t:NumberOfMembersAvailable,omitempty"`
	NumberOfMembersWithConflict XsInt `xml:"t:NumberOfMembersWithConflict,omitempty"`
	NumberOfMembersWithNoData   XsInt `xml:"t:NumberOfMembersWithNoData,omitempty"`
}

type UserOofSettings struct {
	OofState                         OofState            `xml:"t:OofState,omitempty"`
	ExternalAudience                 ExternalAudience    `xml:"t:ExternalAudience,omitempty"`
	Duration                         *Duration           `xml:"t:Duration,omitempty"`
	InternalReply                    *ReplyBody          `xml:"t:InternalReply,omitempty"`
	ExternalReply                    *ReplyBody          `xml:"t:ExternalReply,omitempty"`
	DeclineMeetingReply              *ReplyBody          `xml:"t:DeclineMeetingReply,omitempty"`
	DeclineEventsForScheduledOOF     XsBoolean           `xml:"t:DeclineEventsForScheduledOOF,omitempty"`
	DeclineAllEventsForScheduledOOF  XsBoolean           `xml:"t:DeclineAllEventsForScheduledOOF,omitempty"`
	CreateOOFEvent                   XsBoolean           `xml:"t:CreateOOFEvent,omitempty"`
	OOFEventSubject                  XsString            `xml:"t:OOFEventSubject,omitempty"`
	AutoDeclineFutureRequestsWhenOOF XsBoolean           `xml:"t:AutoDeclineFutureRequestsWhenOOF,omitempty"`
	OOFEventID                       XsString            `xml:"t:OOFEventID,omitempty"`
	EventsToDeleteIDs                *ArrayOfEventIDType `xml:"t:EventsToDeleteIDs,omitempty"`
}

type ReplyBody struct {
	Xxmllang XsString `xml:"xml:lang,attr"`
	Message  XsString `xml:"t:Message,omitempty"`
}

type ArrayOfEventIDType struct {
	EventToDeleteID []XsString `xml:"t:EventToDeleteID,omitempty"`
}

type ConfigurationRequestDetailsType struct {
}

type MailTipsServiceConfiguration struct {
	ServiceConfiguration               `xml:",omitempty"`
	MailTipsEnabled                    XsBoolean       `xml:"t:MailTipsEnabled,omitempty"`
	MaxRecipientsPerGetMailTipsRequest XsInt           `xml:"t:MaxRecipientsPerGetMailTipsRequest,omitempty"`
	MaxMessageSize                     XsInt           `xml:"t:MaxMessageSize,omitempty"`
	LargeAudienceThreshold             XsInt           `xml:"t:LargeAudienceThreshold,omitempty"`
	ShowExternalRecipientCount         XsBoolean       `xml:"t:ShowExternalRecipientCount,omitempty"`
	InternalDomains                    *SmtpDomainList `xml:"t:InternalDomains,omitempty"`
	PolicyTipsEnabled                  XsBoolean       `xml:"t:PolicyTipsEnabled,omitempty"`
	LargeAudienceCap                   XsInt           `xml:"t:LargeAudienceCap,omitempty"`
}

type ServiceConfiguration struct {
}

type SmtpDomainList struct {
	Domain []*SmtpDomain `xml:"t:Domain,omitempty"`
}

type SmtpDomain struct {
	Name              XsString  `xml:"Name,attr"`
	IncludeSubdomains XsBoolean `xml:"IncludeSubdomains,attr"`
}

type UnifiedMessageServiceConfiguration struct {
	ServiceConfiguration  `xml:",omitempty"`
	UmEnabled             XsBoolean `xml:"t:UmEnabled,omitempty"`
	PlayOnPhoneDialString XsString  `xml:"t:PlayOnPhoneDialString,omitempty"`
	PlayOnPhoneEnabled    XsBoolean `xml:"t:PlayOnPhoneEnabled,omitempty"`
}

type ProtectionRulesServiceConfiguration struct {
	ServiceConfiguration `xml:",omitempty"`
	RefreshInterval      RefreshIntervalType         `xml:"RefreshInterval,attr"`
	Rules                *ArrayOfProtectionRulesType `xml:"t:Rules,omitempty"`
	InternalDomains      *SmtpDomainList             `xml:"t:InternalDomains,omitempty"`
}

type ArrayOfProtectionRulesType struct {
	Rule []*ProtectionRuleType `xml:"t:Rule,omitempty"`
}

type ProtectionRuleType struct {
	Name            NameType                     `xml:"Name,attr"`
	UserOverridable XsBoolean                    `xml:"UserOverridable,attr"`
	Priority        PriorityType                 `xml:"Priority,attr"`
	Condition       *ProtectionRuleConditionType `xml:"t:Condition,omitempty"`
	Action          *ProtectionRuleActionType    `xml:"t:Action,omitempty"`
}

type ProtectionRuleConditionType struct {
	AllInternal       ProtectionRuleAllInternalType        `xml:"t:AllInternal,omitempty"`
	And               *ProtectionRuleAndType               `xml:"t:And,omitempty"`
	RecipientIs       *ProtectionRuleRecipientIsType       `xml:"t:RecipientIs,omitempty"`
	SenderDepartments *ProtectionRuleSenderDepartmentsType `xml:"t:SenderDepartments,omitempty"`
	True              ProtectionRuleTrueType               `xml:"t:True,omitempty"`
}

type ProtectionRuleAndType struct {
	AllInternal       []ProtectionRuleAllInternalType        `xml:"t:AllInternal,omitempty"`
	And               []*ProtectionRuleAndType               `xml:"t:And,omitempty"`
	RecipientIs       []*ProtectionRuleRecipientIsType       `xml:"t:RecipientIs,omitempty"`
	SenderDepartments []*ProtectionRuleSenderDepartmentsType `xml:"t:SenderDepartments,omitempty"`
	True              []ProtectionRuleTrueType               `xml:"t:True,omitempty"`
}

type ProtectionRuleRecipientIsType struct {
	Value []ProtectionRuleValueType `xml:"t:Value,omitempty"`
}

type ProtectionRuleSenderDepartmentsType struct {
	Value []ProtectionRuleValueType `xml:"t:Value,omitempty"`
}

type ProtectionRuleActionType struct {
	Name     ProtectionRuleActionKindType  `xml:"Name,attr"`
	Argument []*ProtectionRuleArgumentType `xml:"t:Argument,omitempty"`
}

type ProtectionRuleArgumentType struct {
	Value ValueType `xml:"Value,attr"`
}

type PolicyNudgeRulesServiceConfiguration struct {
}

type SharePointURLsServiceConfiguration struct {
	ServiceConfiguration    `xml:",omitempty"`
	InternalSPMySiteHostURL XsString `xml:"t:InternalSPMySiteHostURL,omitempty"`
	ExternalSPMySiteHostURL XsString `xml:"t:ExternalSPMySiteHostURL,omitempty"`
}

type MailTips struct {
	RecipientAddress        *EmailAddressType                `xml:"t:RecipientAddress,omitempty"`
	PendingMailTips         MailTipTypes                     `xml:"t:PendingMailTips,omitempty"`
	OutOfOffice             *OutOfOfficeMailTip              `xml:"t:OutOfOffice,omitempty"`
	MailboxFull             XsBoolean                        `xml:"t:MailboxFull,omitempty"`
	CustomMailTip           XsString                         `xml:"t:CustomMailTip,omitempty"`
	TotalMemberCount        XsInt                            `xml:"t:TotalMemberCount,omitempty"`
	ExternalMemberCount     XsInt                            `xml:"t:ExternalMemberCount,omitempty"`
	MaxMessageSize          XsInt                            `xml:"t:MaxMessageSize,omitempty"`
	DeliveryRestricted      XsBoolean                        `xml:"t:DeliveryRestricted,omitempty"`
	IsModerated             XsBoolean                        `xml:"t:IsModerated,omitempty"`
	InvalidRecipient        XsBoolean                        `xml:"t:InvalidRecipient,omitempty"`
	Scope                   XsInt                            `xml:"t:Scope,omitempty"`
	RecipientSuggestions    *ArrayOfRecipientSuggestionsType `xml:"t:RecipientSuggestions,omitempty"`
	PreferAccessibleContent XsBoolean                        `xml:"t:PreferAccessibleContent,omitempty"`
}

type OutOfOfficeMailTip struct {
	ReplyBody *ReplyBody `xml:"t:ReplyBody,omitempty"`
	Duration  *Duration  `xml:"t:Duration,omitempty"`
}

type ArrayOfRecipientSuggestionsType struct {
	RecipientSuggestion []*RecipientSuggestionType `xml:"t:RecipientSuggestion,omitempty"`
}

type RecipientSuggestionType struct {
	DisplayName  XsString `xml:"t:DisplayName,omitempty"`
	EmailAddress XsString `xml:"t:EmailAddress,omitempty"`
}

type PhoneCallIdType struct {
	Id XsString `xml:"Id,attr"`
}

type PhoneCallInformationType struct {
	PhoneCallState         PhoneCallStateType         `xml:"t:PhoneCallState,omitempty"`
	ConnectionFailureCause ConnectionFailureCauseType `xml:"t:ConnectionFailureCause,omitempty"`
	SIPResponseText        XsString                   `xml:"t:SIPResponseText,omitempty"`
	SIPResponseCode        XsInt                      `xml:"t:SIPResponseCode,omitempty"`
}

type ArrayOfFindMessageTrackingSearchResultType struct {
	MessageTrackingSearchResult []FindMessageTrackingSearchResultType `xml:"t:MessageTrackingSearchResult,omitempty"`
}

type FindMessageTrackingSearchResultType struct {
}

type ArrayOfArraysOfTrackingPropertiesType struct {
	ArrayOfTrackingPropertiesType []*ArrayOfTrackingPropertiesType `xml:"t:ArrayOfTrackingPropertiesType,omitempty"`
}

type ArrayOfTrackingPropertiesType struct {
	TrackingPropertyType []*TrackingPropertyType `xml:"t:TrackingPropertyType,omitempty"`
}

type TrackingPropertyType struct {
	Name  XsString `xml:"t:Name,omitempty"`
	Value XsString `xml:"t:Value,omitempty"`
}

type MessageTrackingReportType struct {
}

type ConversationResponseShapeType struct {
	BaseShape            DefaultShapeNamesType              `xml:"t:BaseShape,omitempty"`
	AdditionalProperties *NonEmptyArrayOfPathsToElementType `xml:"t:AdditionalProperties,omitempty"`
}

type ArrayOfConversationsType struct {
	Conversation []*ConversationType `xml:"t:Conversation,omitempty"`
}

type ConversationType struct {
	ConversationId            *ItemIdType                     `xml:"t:ConversationId,omitempty"`
	ConversationTopic         XsString                        `xml:"t:ConversationTopic,omitempty"`
	UniqueRecipients          *ArrayOfStringsType             `xml:"t:UniqueRecipients,omitempty"`
	GlobalUniqueRecipients    *ArrayOfStringsType             `xml:"t:GlobalUniqueRecipients,omitempty"`
	UniqueUnreadSenders       *ArrayOfStringsType             `xml:"t:UniqueUnreadSenders,omitempty"`
	GlobalUniqueUnreadSenders *ArrayOfStringsType             `xml:"t:GlobalUniqueUnreadSenders,omitempty"`
	UniqueSenders             *ArrayOfStringsType             `xml:"t:UniqueSenders,omitempty"`
	GlobalUniqueSenders       *ArrayOfStringsType             `xml:"t:GlobalUniqueSenders,omitempty"`
	LastDeliveryTime          XsDateTime                      `xml:"t:LastDeliveryTime,omitempty"`
	GlobalLastDeliveryTime    XsDateTime                      `xml:"t:GlobalLastDeliveryTime,omitempty"`
	Categories                *ArrayOfStringsType             `xml:"t:Categories,omitempty"`
	GlobalCategories          *ArrayOfStringsType             `xml:"t:GlobalCategories,omitempty"`
	FlagStatus                FlagStatusType                  `xml:"t:FlagStatus,omitempty"`
	GlobalFlagStatus          FlagStatusType                  `xml:"t:GlobalFlagStatus,omitempty"`
	HasAttachments            XsBoolean                       `xml:"t:HasAttachments,omitempty"`
	GlobalHasAttachments      XsBoolean                       `xml:"t:GlobalHasAttachments,omitempty"`
	MessageCount              XsInt                           `xml:"t:MessageCount,omitempty"`
	GlobalMessageCount        XsInt                           `xml:"t:GlobalMessageCount,omitempty"`
	UnreadCount               XsInt                           `xml:"t:UnreadCount,omitempty"`
	GlobalUnreadCount         XsInt                           `xml:"t:GlobalUnreadCount,omitempty"`
	Size                      XsInt                           `xml:"t:Size,omitempty"`
	GlobalSize                XsInt                           `xml:"t:GlobalSize,omitempty"`
	ItemClasses               *ArrayOfItemClassType           `xml:"t:ItemClasses,omitempty"`
	GlobalItemClasses         *ArrayOfItemClassType           `xml:"t:GlobalItemClasses,omitempty"`
	Importance                ImportanceChoicesType           `xml:"t:Importance,omitempty"`
	GlobalImportance          ImportanceChoicesType           `xml:"t:GlobalImportance,omitempty"`
	ItemIds                   *NonEmptyArrayOfBaseItemIdsType `xml:"t:ItemIds,omitempty"`
	GlobalItemIds             *NonEmptyArrayOfBaseItemIdsType `xml:"t:GlobalItemIds,omitempty"`
	LastModifiedTime          XsDateTime                      `xml:"t:LastModifiedTime,omitempty"`
	InstanceKey               XsBase64Binary                  `xml:"t:InstanceKey,omitempty"`
	Preview                   XsString                        `xml:"t:Preview,omitempty"`
	MailboxScope              MailboxSearchLocationType       `xml:"t:MailboxScope,omitempty"`
	IconIndex                 IconIndexType                   `xml:"t:IconIndex,omitempty"`
	GlobalIconIndex           IconIndexType                   `xml:"t:GlobalIconIndex,omitempty"`
	DraftItemIds              *NonEmptyArrayOfBaseItemIdsType `xml:"t:DraftItemIds,omitempty"`
	HasIrm                    XsBoolean                       `xml:"t:HasIrm,omitempty"`
	GlobalHasIrm              XsBoolean                       `xml:"t:GlobalHasIrm,omitempty"`
	InferenceClassification   InferenceClassificationType     `xml:"t:InferenceClassification,omitempty"`
	SortKey                   XsLong                          `xml:"t:SortKey,omitempty"`
	MentionedMe               XsBoolean                       `xml:"t:MentionedMe,omitempty"`
	GlobalMentionedMe         XsBoolean                       `xml:"t:GlobalMentionedMe,omitempty"`
	SenderSMTPAddress         *SmtpAddressType                `xml:"t:SenderSMTPAddress,omitempty"`
	MailboxGuids              *MailboxGuidsType               `xml:"t:MailboxGuids,omitempty"`
	From                      *SingleRecipientType            `xml:"t:From,omitempty"`
	AtAllMention              XsBoolean                       `xml:"t:AtAllMention,omitempty"`
	GlobalAtAllMention        XsBoolean                       `xml:"t:GlobalAtAllMention,omitempty"`
}

type ArrayOfItemClassType struct {
	ItemClass []ItemClassType `xml:"t:ItemClass,omitempty"`
}

type NonEmptyArrayOfApplyConversationActionType struct {
	ConversationAction []*ConversationActionType `xml:"t:ConversationAction,omitempty"`
}

type ConversationActionType struct {
	Action                   ConversationActionTypeType `xml:"t:Action,omitempty"`
	ConversationId           *ItemIdType                `xml:"t:ConversationId,omitempty"`
	ContextFolderId          *TargetFolderIdType        `xml:"t:ContextFolderId,omitempty"`
	ConversationLastSyncTime XsDateTime                 `xml:"t:ConversationLastSyncTime,omitempty"`
	ProcessRightAway         XsBoolean                  `xml:"t:ProcessRightAway,omitempty"`
	DestinationFolderId      *TargetFolderIdType        `xml:"t:DestinationFolderId,omitempty"`
	Categories               *ArrayOfStringsType        `xml:"t:Categories,omitempty"`
	EnableAlwaysDelete       XsBoolean                  `xml:"t:EnableAlwaysDelete,omitempty"`
	IsRead                   XsBoolean                  `xml:"t:IsRead,omitempty"`
	DeleteType               DisposalType               `xml:"t:DeleteType,omitempty"`
	RetentionPolicyType      RetentionType              `xml:"t:RetentionPolicyType,omitempty"`
	RetentionPolicyTagId     XsString                   `xml:"t:RetentionPolicyTagId,omitempty"`
	Flag                     *FlagType                  `xml:"t:Flag,omitempty"`
	SuppressReadReceipts     XsBoolean                  `xml:"t:SuppressReadReceipts,omitempty"`
}

type ArrayOfConversationRequestsType struct {
	Conversation []*ConversationRequestType `xml:"t:Conversation,omitempty"`
}

type ConversationRequestType struct {
	ConversationId *ItemIdType    `xml:"t:ConversationId,omitempty"`
	SyncState      XsBase64Binary `xml:"t:SyncState,omitempty"`
}

type PersonaResponseShapeType struct {
	BaseShape            DefaultShapeNamesType              `xml:"t:BaseShape,omitempty"`
	AdditionalProperties *NonEmptyArrayOfPathsToElementType `xml:"t:AdditionalProperties,omitempty"`
}

type ArrayOfContextProperty struct {
	ContextProperty []*ContextPropertyType `xml:"t:ContextProperty,omitempty"`
}

type ContextPropertyType struct {
	Key   XsString `xml:"t:Key,omitempty"`
	Value XsString `xml:"t:Value,omitempty"`
}

type ArrayOfPeopleQuerySource struct {
	Source []XsString `xml:"t:Source,omitempty"`
}

type ArrayOfRulesType struct {
	Rule []*RuleType `xml:"t:Rule,omitempty"`
}

type RuleType struct {
	RuleId         XsString            `xml:"t:RuleId,omitempty"`
	DisplayName    XsString            `xml:"t:DisplayName,omitempty"`
	Priority       XsInt               `xml:"t:Priority,omitempty"`
	IsEnabled      XsBoolean           `xml:"t:IsEnabled,omitempty"`
	IsNotSupported XsBoolean           `xml:"t:IsNotSupported,omitempty"`
	IsInError      XsBoolean           `xml:"t:IsInError,omitempty"`
	Conditions     *RulePredicatesType `xml:"t:Conditions,omitempty"`
	Exceptions     *RulePredicatesType `xml:"t:Exceptions,omitempty"`
	Actions        *RuleActionsType    `xml:"t:Actions,omitempty"`
}

type RulePredicatesType struct {
	Categories                   *ArrayOfStringsType         `xml:"t:Categories,omitempty"`
	ContainsBodyStrings          *ArrayOfStringsType         `xml:"t:ContainsBodyStrings,omitempty"`
	ContainsHeaderStrings        *ArrayOfStringsType         `xml:"t:ContainsHeaderStrings,omitempty"`
	ContainsRecipientStrings     *ArrayOfStringsType         `xml:"t:ContainsRecipientStrings,omitempty"`
	ContainsSenderStrings        *ArrayOfStringsType         `xml:"t:ContainsSenderStrings,omitempty"`
	ContainsSubjectOrBodyStrings *ArrayOfStringsType         `xml:"t:ContainsSubjectOrBodyStrings,omitempty"`
	ContainsSubjectStrings       *ArrayOfStringsType         `xml:"t:ContainsSubjectStrings,omitempty"`
	FlaggedForAction             FlaggedForActionType        `xml:"t:FlaggedForAction,omitempty"`
	FromAddresses                *ArrayOfEmailAddressesType  `xml:"t:FromAddresses,omitempty"`
	FromConnectedAccounts        *ArrayOfStringsType         `xml:"t:FromConnectedAccounts,omitempty"`
	HasAttachments               XsBoolean                   `xml:"t:HasAttachments,omitempty"`
	Importance                   ImportanceChoicesType       `xml:"t:Importance,omitempty"`
	IsApprovalRequest            XsBoolean                   `xml:"t:IsApprovalRequest,omitempty"`
	IsAutomaticForward           XsBoolean                   `xml:"t:IsAutomaticForward,omitempty"`
	IsAutomaticReply             XsBoolean                   `xml:"t:IsAutomaticReply,omitempty"`
	IsEncrypted                  XsBoolean                   `xml:"t:IsEncrypted,omitempty"`
	IsMeetingRequest             XsBoolean                   `xml:"t:IsMeetingRequest,omitempty"`
	IsMeetingResponse            XsBoolean                   `xml:"t:IsMeetingResponse,omitempty"`
	IsNDR                        XsBoolean                   `xml:"t:IsNDR,omitempty"`
	IsPermissionControlled       XsBoolean                   `xml:"t:IsPermissionControlled,omitempty"`
	IsReadReceipt                XsBoolean                   `xml:"t:IsReadReceipt,omitempty"`
	IsSigned                     XsBoolean                   `xml:"t:IsSigned,omitempty"`
	IsVoicemail                  XsBoolean                   `xml:"t:IsVoicemail,omitempty"`
	ItemClasses                  *ArrayOfStringsType         `xml:"t:ItemClasses,omitempty"`
	MessageClassifications       *ArrayOfStringsType         `xml:"t:MessageClassifications,omitempty"`
	NotSentToMe                  XsBoolean                   `xml:"t:NotSentToMe,omitempty"`
	SentCcMe                     XsBoolean                   `xml:"t:SentCcMe,omitempty"`
	SentOnlyToMe                 XsBoolean                   `xml:"t:SentOnlyToMe,omitempty"`
	SentToAddresses              *ArrayOfEmailAddressesType  `xml:"t:SentToAddresses,omitempty"`
	SentToMe                     XsBoolean                   `xml:"t:SentToMe,omitempty"`
	SentToOrCcMe                 XsBoolean                   `xml:"t:SentToOrCcMe,omitempty"`
	Sensitivity                  SensitivityChoicesType      `xml:"t:Sensitivity,omitempty"`
	WithinDateRange              *RulePredicateDateRangeType `xml:"t:WithinDateRange,omitempty"`
	WithinSizeRange              *RulePredicateSizeRangeType `xml:"t:WithinSizeRange,omitempty"`
}

type RulePredicateDateRangeType struct {
	StartDateTime XsDateTime `xml:"t:StartDateTime,omitempty"`
	EndDateTime   XsDateTime `xml:"t:EndDateTime,omitempty"`
}

type RulePredicateSizeRangeType struct {
	MinimumSize XsInt `xml:"t:MinimumSize,omitempty"`
	MaximumSize XsInt `xml:"t:MaximumSize,omitempty"`
}

type RuleActionsType struct {
	AssignCategories                *ArrayOfStringsType        `xml:"t:AssignCategories,omitempty"`
	CopyToFolder                    *TargetFolderIdType        `xml:"t:CopyToFolder,omitempty"`
	Delete                          XsBoolean                  `xml:"t:Delete,omitempty"`
	ForwardAsAttachmentToRecipients *ArrayOfEmailAddressesType `xml:"t:ForwardAsAttachmentToRecipients,omitempty"`
	ForwardToRecipients             *ArrayOfEmailAddressesType `xml:"t:ForwardToRecipients,omitempty"`
	MarkImportance                  ImportanceChoicesType      `xml:"t:MarkImportance,omitempty"`
	MarkAsRead                      XsBoolean                  `xml:"t:MarkAsRead,omitempty"`
	MoveToFolder                    *TargetFolderIdType        `xml:"t:MoveToFolder,omitempty"`
	PermanentDelete                 XsBoolean                  `xml:"t:PermanentDelete,omitempty"`
	RedirectToRecipients            *ArrayOfEmailAddressesType `xml:"t:RedirectToRecipients,omitempty"`
	SendSMSAlertToRecipients        *ArrayOfEmailAddressesType `xml:"t:SendSMSAlertToRecipients,omitempty"`
	ServerReplyWithMessage          *ItemIdType                `xml:"t:ServerReplyWithMessage,omitempty"`
	StopProcessingRules             XsBoolean                  `xml:"t:StopProcessingRules,omitempty"`
}

type ArrayOfRuleOperationsType struct {
	CreateRuleOperation []*CreateRuleOperationType `xml:"t:CreateRuleOperation,omitempty"`
	SetRuleOperation    []*SetRuleOperationType    `xml:"t:SetRuleOperation,omitempty"`
	DeleteRuleOperation []*DeleteRuleOperationType `xml:"t:DeleteRuleOperation,omitempty"`
}

type CreateRuleOperationType struct {
	RuleOperationType `xml:",omitempty"`
	Rule              *RuleType `xml:"t:Rule,omitempty"`
}

type RuleOperationType struct {
}

type SetRuleOperationType struct {
	RuleOperationType `xml:",omitempty"`
	Rule              *RuleType `xml:"t:Rule,omitempty"`
}

type DeleteRuleOperationType struct {
	RuleOperationType `xml:",omitempty"`
	RuleId            XsString `xml:"t:RuleId,omitempty"`
}

type ArrayOfRuleOperationErrorsType struct {
	RuleOperationError []*RuleOperationErrorType `xml:"t:RuleOperationError,omitempty"`
}

type RuleOperationErrorType struct {
	OperationIndex   XsInt                            `xml:"t:OperationIndex,omitempty"`
	ValidationErrors *ArrayOfRuleValidationErrorsType `xml:"t:ValidationErrors,omitempty"`
}

type ArrayOfRuleValidationErrorsType struct {
	Error []*RuleValidationErrorType `xml:"t:Error,omitempty"`
}

type RuleValidationErrorType struct {
	FieldURI     RuleFieldURIType            `xml:"t:FieldURI,omitempty"`
	ErrorCode    RuleValidationErrorCodeType `xml:"t:ErrorCode,omitempty"`
	ErrorMessage XsString                    `xml:"t:ErrorMessage,omitempty"`
	FieldValue   XsString                    `xml:"t:FieldValue,omitempty"`
}

type PreviewItemResponseShapeType struct {
	BaseShape            PreviewItemBaseShapeType              `xml:"t:BaseShape,omitempty"`
	AdditionalProperties *NonEmptyArrayOfExtendedFieldURIsType `xml:"t:AdditionalProperties,omitempty"`
}

type NonEmptyArrayOfExtendedFieldURIsType struct {
	ExtendedFieldURI []*PathToExtendedFieldType `xml:"t:ExtendedFieldURI,omitempty"`
}

type NonEmptyArrayOfLegacyDNsType struct {
	LegacyDN []XsString `xml:"t:LegacyDN,omitempty"`
}

type ArrayOfAppsType struct {
	App      []*AppType `xml:"t:App,omitempty"`
	Metadata *Metadata  `xml:"t:Metadata,omitempty"`
}

type AppType struct {
	Metadata *AppMetadata   `xml:"t:Metadata,omitempty"`
	Manifest XsBase64Binary `xml:"t:Manifest,omitempty"`
}

type AppMetadata struct {
	EndNodeUrl          XsString   `xml:"t:EndNodeUrl,omitempty"`
	AppStatus           XsString   `xml:"t:AppStatus,omitempty"`
	ActionUrl           XsString   `xml:"t:ActionUrl,omitempty"`
	ProductId           XsString   `xml:"t:ProductId,omitempty"`
	EnabledStatus       XsBoolean  `xml:"t:EnabledStatus,omitempty"`
	ConsentState        XsString   `xml:"t:ConsentState,omitempty"`
	ExtensionType       XsString   `xml:"t:ExtensionType,omitempty"`
	MarketplaceAssetId  XsString   `xml:"t:MarketplaceAssetId,omitempty"`
	LicenseStatus       XsString   `xml:"t:LicenseStatus,omitempty"`
	TrialExpirationDate XsDateTime `xml:"t:TrialExpirationDate,omitempty"`
	InstalledBy         XsString   `xml:"t:InstalledBy,omitempty"`
	IsMandatory         XsBoolean  `xml:"t:IsMandatory,omitempty"`
}

type Metadata struct {
	CustomApps  XsString `xml:"t:CustomApps,omitempty"`
	GenericInfo XsString `xml:"t:GenericInfo,omitempty"`
}

type ImGroupType struct {
	DisplayName          NonEmptyStringType                   `xml:"t:DisplayName,omitempty"`
	GroupType            NonEmptyStringType                   `xml:"t:GroupType,omitempty"`
	ExchangeStoreId      *ItemIdType                          `xml:"t:ExchangeStoreId,omitempty"`
	MemberCorrelationKey *NonEmptyArrayOfItemIdsType          `xml:"t:MemberCorrelationKey,omitempty"`
	ExtendedProperties   *NonEmptyArrayOfExtendedPropertyType `xml:"t:ExtendedProperties,omitempty"`
	SmtpAddress          XsString                             `xml:"t:SmtpAddress,omitempty"`
}

type NonEmptyArrayOfExtendedFieldURIs struct {
	ExtendedProperty []*PathToExtendedFieldType `xml:"t:ExtendedProperty,omitempty"`
}

type ImItemListType struct {
	Groups   *ArrayOfImGroupType `xml:"t:Groups,omitempty"`
	Personas *ArrayOfPeopleType  `xml:"t:Personas,omitempty"`
}

type ArrayOfImGroupType struct {
	ImGroup []*ImGroupType `xml:"t:ImGroup,omitempty"`
}

type InstalledAppType struct {
	Id                    XsString   `xml:"t:Id,omitempty"`
	MarketplaceAssetId    XsString   `xml:"t:MarketplaceAssetId,omitempty"`
	Enabled               XsBoolean  `xml:"t:Enabled,omitempty"`
	ConsentState          XsString   `xml:"t:ConsentState,omitempty"`
	Type                  XsString   `xml:"t:Type,omitempty"`
	LicenseStatus         XsString   `xml:"t:LicenseStatus,omitempty"`
	TrialExpirationDate   XsDateTime `xml:"t:TrialExpirationDate,omitempty"`
	ProviderName          XsString   `xml:"t:ProviderName,omitempty"`
	IconUrl               XsString   `xml:"t:IconUrl,omitempty"`
	HighResolutionIconUrl XsString   `xml:"t:HighResolutionIconUrl,omitempty"`
	DisplayName           XsString   `xml:"t:DisplayName,omitempty"`
	Description           XsString   `xml:"t:Description,omitempty"`
	Requirements          XsString   `xml:"t:Requirements,omitempty"`
	Version               XsString   `xml:"t:Version,omitempty"`
}

type ArrayOfMeetingTimeCandidate struct {
	MeetingTimeCandidate []*MeetingTimeCandidate `xml:"t:MeetingTimeCandidate,omitempty"`
}

type MeetingTimeCandidate struct {
	MeetingTimeslot        *TimeSlot                    `xml:"t:MeetingTimeslot,omitempty"`
	Confidence             XsDouble                     `xml:"t:Confidence,omitempty"`
	Score                  XsInt                        `xml:"t:Score,omitempty"`
	OrganizerAvailability  AvailabilityStatusType       `xml:"t:OrganizerAvailability,omitempty"`
	AttendeeAvailabilities *ArrayOfAttendeeAvailability `xml:"t:AttendeeAvailabilities,omitempty"`
	Locations              *ArrayOfMeetingLocation      `xml:"t:Locations,omitempty"`
	SuggestionHint         XsString                     `xml:"t:SuggestionHint,omitempty"`
}

type TimeSlot struct {
	StartTime         XsDateTime `xml:"t:StartTime,omitempty"`
	DurationInMinutes XsDouble   `xml:"t:DurationInMinutes,omitempty"`
}

type ArrayOfAttendeeAvailability struct {
	AttendeeAvailability []*AttendeeAvailability `xml:"t:AttendeeAvailability,omitempty"`
}

type AttendeeAvailability struct {
	EmailAddress XsString               `xml:"t:EmailAddress,omitempty"`
	Availability AvailabilityStatusType `xml:"t:Availability,omitempty"`
}

type ArrayOfMeetingLocation struct {
	MeetingLocation []*MeetingLocation `xml:"t:MeetingLocation,omitempty"`
}

type MeetingLocation struct {
	EmailAddress XsString `xml:"t:EmailAddress,omitempty"`
	DisplayName  XsString `xml:"t:DisplayName,omitempty"`
}

type FindMeetingTimesAttendeeConstraints struct {
	AttendeeEntries *ArrayOfAttendeeConstraintItems `xml:"t:AttendeeEntries,omitempty"`
}

type ArrayOfAttendeeConstraintItems struct {
	AttendeeItem []*AttendeeConstraintItem `xml:"t:AttendeeItem,omitempty"`
}

type AttendeeConstraintItem struct {
	MeetingTimeCandidatesConstraintItem `xml:",omitempty"`
	IsRequired                          XsBoolean `xml:"t:IsRequired,omitempty"`
}

type MeetingTimeCandidatesConstraintItem struct {
	Email XsString `xml:"t:Email,omitempty"`
}

type FindMeetingTimesLocationConstraints struct {
	LocationEntries *ArrayOfLocationConstraintItems `xml:"t:LocationEntries,omitempty"`
	IsRequired      XsBoolean                       `xml:"t:IsRequired,omitempty"`
	SuggestLocation XsBoolean                       `xml:"t:SuggestLocation,omitempty"`
}

type ArrayOfLocationConstraintItems struct {
	LocationItem []*LocationConstraintItem `xml:"t:LocationItem,omitempty"`
}

type LocationConstraintItem struct {
	MeetingTimeCandidatesConstraintItem `xml:",omitempty"`
	Name                                XsString  `xml:"t:Name,omitempty"`
	ResolveAvailability                 XsBoolean `xml:"t:ResolveAvailability,omitempty"`
}

type FindMeetingTimesSearchConstraints struct {
	SearchWindows            *ArrayOfTimeSlot   `xml:"t:SearchWindows,omitempty"`
	MeetingDurationInMinutes XsInt              `xml:"t:MeetingDurationInMinutes,omitempty"`
	ActivityDomain           ActivityDomainType `xml:"t:ActivityDomain,omitempty"`
}

type ArrayOfTimeSlot struct {
	TimeSlot []*TimeSlot `xml:"t:TimeSlot,omitempty"`
}

type FindMeetingTimesConstraints struct {
	MaxCandidates             XsInt     `xml:"t:MaxCandidates,omitempty"`
	IsOrganizerOptional       XsBoolean `xml:"t:IsOrganizerOptional,omitempty"`
	ReturnSuggestionHints     XsBoolean `xml:"t:ReturnSuggestionHints,omitempty"`
	AppName                   XsString  `xml:"t:AppName,omitempty"`
	AppScenario               XsString  `xml:"t:AppScenario,omitempty"`
	MinimumAttendeePercentage XsDouble  `xml:"t:MinimumAttendeePercentage,omitempty"`
}

type MeetingSpaceType struct {
	Id               XsString             `xml:"t:Id,omitempty"`
	ChangeKey        XsString             `xml:"t:ChangeKey,omitempty"`
	Type             MeetingSpaceTypeEnum `xml:"t:Type,omitempty"`
	Version          XsString             `xml:"t:Version,omitempty"`
	JoinUrl          XsString             `xml:"t:JoinUrl,omitempty"`
	DateTimeCreated  XsString             `xml:"t:DateTimeCreated,omitempty"`
	DateTimeModified XsString             `xml:"t:DateTimeModified,omitempty"`
	ExpiryTime       XsString             `xml:"t:ExpiryTime,omitempty"`
	Meadata          XsString             `xml:"t:Meadata,omitempty"`
	Tag              XsString             `xml:"t:Tag,omitempty"`
}

type MeetingInstanceType struct {
	Id                    XsString                              `xml:"t:Id,omitempty"`
	ChangeKey             XsString                              `xml:"t:ChangeKey,omitempty"`
	Version               XsString                              `xml:"t:Version,omitempty"`
	JoinUrl               XsString                              `xml:"t:JoinUrl,omitempty"`
	DateTimeCreated       XsString                              `xml:"t:DateTimeCreated,omitempty"`
	DateTimeModified      XsString                              `xml:"t:DateTimeModified,omitempty"`
	Meadata               XsString                              `xml:"t:Meadata,omitempty"`
	Tag                   XsString                              `xml:"t:Tag,omitempty"`
	ParentGoid            XsString                              `xml:"t:ParentGoid,omitempty"`
	ParticipantActivities *NonEmptyArrayOfParticipantActivities `xml:"t:ParticipantActivities,omitempty"`
	ContentActivities     *NonEmptyArrayOfContentActivities     `xml:"t:ContentActivities,omitempty"`
}

type NonEmptyArrayOfParticipantActivities struct {
	ParticipantActivity []*ParticipantActivity `xml:"t:ParticipantActivity,omitempty"`
}

type ParticipantActivity struct {
	Id            GuidType                     `xml:"t:Id,omitempty"`
	CreatedBy     XsString                     `xml:"t:CreatedBy,omitempty"`
	StartTime     XsString                     `xml:"t:StartTime,omitempty"`
	EndTime       XsString                     `xml:"t:EndTime,omitempty"`
	ClientVersion XsString                     `xml:"t:ClientVersion,omitempty"`
	Role          ParticipantActivityRole      `xml:"t:Role,omitempty"`
	MediaType     ParticipantActivityMediaType `xml:"t:MediaType,omitempty"`
	MediaDetails  XsString                     `xml:"t:MediaDetails,omitempty"`
}

type NonEmptyArrayOfContentActivities struct {
	ContentActivity []*ContentActivity `xml:"t:ContentActivity,omitempty"`
}

type ContentActivity struct {
	Id              GuidType                 `xml:"t:Id,omitempty"`
	SharedBy        XsString                 `xml:"t:SharedBy,omitempty"`
	ContentLocation XsString                 `xml:"t:ContentLocation,omitempty"`
	StartTime       XsString                 `xml:"t:StartTime,omitempty"`
	EndTime         XsString                 `xml:"t:EndTime,omitempty"`
	Type            ContentActivityType      `xml:"t:Type,omitempty"`
	MediaType       ContentActivityMediaType `xml:"t:MediaType,omitempty"`
	Acl             ContentActivityAcl       `xml:"t:Acl,omitempty"`
}

type ArrayOfSearchScopeType struct {
	PrimaryMailboxSearchScope     *PrimaryMailboxSearchScopeType     `xml:"t:PrimaryMailboxSearchScope,omitempty"`
	LargeArchiveSearchScope       *LargeArchiveSearchScopeType       `xml:"t:LargeArchiveSearchScope,omitempty"`
	GroupSearchScope              *GroupSearchScopeType              `xml:"t:GroupSearchScope,omitempty"`
	CustomSearchScope             []*CustomSearchScopeType           `xml:"t:CustomSearchScope,omitempty"`
	OneDriveSearchScope           *OneDriveSearchScopeType           `xml:"t:OneDriveSearchScope,omitempty"`
	SingleLargeArchiveSearchScope *SingleLargeArchiveSearchScopeType `xml:"t:SingleLargeArchiveSearchScope,omitempty"`
	DelveSearchScope              *DelveSearchScopeType              `xml:"t:DelveSearchScope,omitempty"`
}

type PrimaryMailboxSearchScopeType struct {
	FolderScope     *SearchFolderScopeType `xml:"t:FolderScope,omitempty"`
	IsDeepTraversal XsBoolean              `xml:"t:IsDeepTraversal,omitempty"`
}

type SearchFolderScopeType struct {
	FolderId        *FolderIdType              `xml:"t:FolderId,omitempty"`
	WellKnownFolder *DistinguishedFolderIdType `xml:"t:WellKnownFolder,omitempty"`
}

type LargeArchiveSearchScopeType struct {
	ArchiveTypes SearchScopeArchivesType `xml:"t:ArchiveTypes,omitempty"`
}

type GroupSearchScopeType struct {
	GroupTypes SearchScopeGroupsType `xml:"t:GroupTypes,omitempty"`
}

type CustomSearchScopeType struct {
	MailboxGuid     GuidType               `xml:"t:MailboxGuid,omitempty"`
	FolderScope     *SearchFolderScopeType `xml:"t:FolderScope,omitempty"`
	IsDeepTraversal XsBoolean              `xml:"t:IsDeepTraversal,omitempty"`
}

type OneDriveSearchScopeType struct {
	OneDriveView OneDriveViewType `xml:"t:OneDriveView,omitempty"`
}

type SingleLargeArchiveSearchScopeType struct {
	MailboxGuid     GuidType               `xml:"t:MailboxGuid,omitempty"`
	FolderScope     *SearchFolderScopeType `xml:"t:FolderScope,omitempty"`
	IsDeepTraversal XsBoolean              `xml:"t:IsDeepTraversal,omitempty"`
}

type DelveSearchScopeType struct {
	DelveView DelveViewType `xml:"t:DelveView,omitempty"`
}

type SearchSuggestionsType struct {
	TDSuggestionsBatchId    XsLong                `xml:"t:TDSuggestionsBatchId,omitempty"`
	TDSuggestionsInstanceId GuidType              `xml:"t:TDSuggestionsInstanceId,omitempty"`
	Suggestions             *SuggestionsType      `xml:"t:Suggestions,omitempty"`
	DiagnosticsData         SearchDiagnosticsType `xml:"t:DiagnosticsData,omitempty"`
}

type SuggestionType struct {
	SuggestedQuery XsString           `xml:"t:SuggestedQuery,omitempty"`
	DisplayText    XsString           `xml:"t:DisplayText,omitempty"`
	SuggestionType SuggestionKindType `xml:"t:SuggestionType,omitempty"`
	Trigger        XsString           `xml:"t:Trigger,omitempty"`
	TDSuggestionId XsInt              `xml:"t:TDSuggestionId,omitempty"`
	IsDeletable    XsBoolean          `xml:"t:IsDeletable,omitempty"`
}

type SearchDiagnosticsType struct {
}

type DeleteSearchSuggestionResponseType struct {
	Success         XsBoolean             `xml:"t:Success,omitempty"`
	StatusMessage   XsString              `xml:"t:StatusMessage,omitempty"`
	DiagnosticsData SearchDiagnosticsType `xml:"t:DiagnosticsData,omitempty"`
}

type AnalyzedQuery struct {
	QueryLanguage      XsString         `xml:"t:QueryLanguage,omitempty"`
	SearchRestrictions *RestrictionType `xml:"t:SearchRestrictions,omitempty"`
}

type DynamicRefinerQueryType struct {
	RefinerQuery XsString `xml:"t:RefinerQuery,omitempty"`
	TDRefinerId  XsInt    `xml:"t:TDRefinerId,omitempty"`
}

type ExtendedKeywordDefinitionType struct {
	Keyword    XsString                              `xml:"t:Keyword,omitempty"`
	Properties *NonEmptyArrayOfExtendedFieldURIsType `xml:"t:Properties,omitempty"`
}

type SearchResultsType struct {
	Items                   *ItemsType                `xml:"t:Items,omitempty"`
	Conversations           *ConversationsType        `xml:"t:Conversations,omitempty"`
	People                  *PeopleType               `xml:"t:People,omitempty"`
	MoreResultsAvailable    XsBoolean                 `xml:"t:MoreResultsAvailable,omitempty"`
	RefinerTelemetryBatchId XsInt                     `xml:"t:RefinerTelemetryBatchId,omitempty"`
	SearchRefiners          *SearchRefinersTypeT      `xml:"t:SearchRefiners,omitempty"`
	DiagnosticsData         SearchDiagnosticsType     `xml:"t:DiagnosticsData,omitempty"`
	SearchResultsCount      XsInt                     `xml:"t:SearchResultsCount,omitempty"`
	TotalResultsCount       XsInt                     `xml:"t:TotalResultsCount,omitempty"`
	SearchTerms             *ArrayOfStringsType       `xml:"t:SearchTerms,omitempty"`
	QueryId                 *ExecuteSearchQueryIdType `xml:"t:QueryId,omitempty"`
	MailboxesInformation    *MailboxesInformationType `xml:"t:MailboxesInformation,omitempty"`
}

type SearchRefinerType struct {
	RefinerType RefinerTypeType          `xml:"t:RefinerType,omitempty"`
	Refiner     *DynamicRefinerQueryType `xml:"t:Refiner,omitempty"`
	ResultCount XsLong                   `xml:"t:ResultCount,omitempty"`
}

type ExecuteSearchQueryIdType struct {
	Id GuidType `xml:"Id,attr"`
}

type MailboxInformationType struct {
	MailboxGuid        GuidType         `xml:"t:MailboxGuid,omitempty"`
	MailboxAddress     *SmtpAddressType `xml:"t:MailboxAddress,omitempty"`
	MailboxDisplayName XsString         `xml:"t:MailboxDisplayName,omitempty"`
}

type OfficeClientType struct {
	Code    OfficeClientCodeType `xml:"Code,attr"`
	Version VersionType          `xml:"Version,attr"`
}

type ResolveNamesSoapIn struct {
	ResolveNames          *ResolveNamesType          `xml:"m:ResolveNames,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type ResolveNamesSoapOut struct {
	ResolveNamesResponse *ResolveNamesResponseType `xml:"m:ResolveNamesResponse,omitempty"`
	ServerVersionInfo    *ServerVersionInfoType    `xml:"t:ServerVersionInfo,omitempty"`
}

type ExpandDLSoapIn struct {
	ExpandDL              *ExpandDLType              `xml:"m:ExpandDL,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type ExpandDLSoapOut struct {
	ExpandDLResponse  *ExpandDLResponseType  `xml:"m:ExpandDLResponse,omitempty"`
	ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
}

type GetServerTimeZonesSoapIn struct {
	GetServerTimeZones   *GetServerTimeZonesType   `xml:"m:GetServerTimeZones,omitempty"`
	MailboxCulture       *MailboxCultureType       `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
}

type GetServerTimeZonesSoapOut struct {
	GetServerTimeZonesResponse *GetServerTimeZonesResponseType `xml:"m:GetServerTimeZonesResponse,omitempty"`
	ServerVersionInfo          *ServerVersionInfoType          `xml:"t:ServerVersionInfo,omitempty"`
}

type FindFolderSoapIn struct {
	FindFolder            *FindFolderType            `xml:"m:FindFolder,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	TimeZoneContext       *TimeZoneContextType       `xml:"t:TimeZoneContext,omitempty"`
	ManagementRole        *ManagementRoleType        `xml:"t:ManagementRole,omitempty"`
}

type FindFolderSoapOut struct {
	FindFolderResponse *FindFolderResponseType `xml:"m:FindFolderResponse,omitempty"`
	ServerVersionInfo  *ServerVersionInfoType  `xml:"t:ServerVersionInfo,omitempty"`
}

type FindItemSoapIn struct {
	FindItem              *FindItemType              `xml:"m:FindItem,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	TimeZoneContext       *TimeZoneContextType       `xml:"t:TimeZoneContext,omitempty"`
	DateTimePrecision     DateTimePrecisionType      `xml:"t:DateTimePrecision,omitempty"`
	ManagementRole        *ManagementRoleType        `xml:"t:ManagementRole,omitempty"`
}

type FindItemSoapOut struct {
	FindItemResponse  *FindItemResponseType  `xml:"m:FindItemResponse,omitempty"`
	ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
}

type GetFolderSoapIn struct {
	GetFolder             *GetFolderType             `xml:"m:GetFolder,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	TimeZoneContext       *TimeZoneContextType       `xml:"t:TimeZoneContext,omitempty"`
	ManagementRole        *ManagementRoleType        `xml:"t:ManagementRole,omitempty"`
}

type GetFolderSoapOut struct {
	GetFolderResponse *GetFolderResponseType `xml:"m:GetFolderResponse,omitempty"`
	ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
}

type UploadItemsSoapIn struct {
	UploadItems           *UploadItemsType           `xml:"m:UploadItems,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type UploadItemsSoapOut struct {
	UploadItemsResponse *UploadItemsResponseType `xml:"m:UploadItemsResponse,omitempty"`
	ServerVersionInfo   *ServerVersionInfoType   `xml:"t:ServerVersionInfo,omitempty"`
}

type ExportItemsSoapIn struct {
	ExportItems           *ExportItemsType           `xml:"m:ExportItems,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	ManagementRole        *ManagementRoleType        `xml:"t:ManagementRole,omitempty"`
}

type ExportItemsSoapOut struct {
	ExportItemsResponse *ExportItemsResponseType `xml:"m:ExportItemsResponse,omitempty"`
	ServerVersionInfo   *ServerVersionInfoType   `xml:"t:ServerVersionInfo,omitempty"`
}

type ConvertIdSoapIn struct {
	ConvertId             *ConvertIdType             `xml:"m:ConvertId,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type ConvertIdSoapOut struct {
	ConvertIdResponse *ConvertIdResponseType `xml:"m:ConvertIdResponse,omitempty"`
	ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
}

type CreateFolderSoapIn struct {
	CreateFolder          *CreateFolderType          `xml:"m:CreateFolder,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	TimeZoneContext       *TimeZoneContextType       `xml:"t:TimeZoneContext,omitempty"`
}

type CreateFolderSoapOut struct {
	CreateFolderResponse *CreateFolderResponseType `xml:"m:CreateFolderResponse,omitempty"`
	ServerVersionInfo    *ServerVersionInfoType    `xml:"t:ServerVersionInfo,omitempty"`
}

type CreateFolderPathSoapIn struct {
	CreateFolderPath      *CreateFolderPathType      `xml:"m:CreateFolderPath,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	TimeZoneContext       *TimeZoneContextType       `xml:"t:TimeZoneContext,omitempty"`
}

type CreateFolderPathSoapOut struct {
	CreateFolderPathResponse *CreateFolderPathResponseType `xml:"m:CreateFolderPathResponse,omitempty"`
	ServerVersionInfo        *ServerVersionInfoType        `xml:"t:ServerVersionInfo,omitempty"`
}

type DeleteFolderSoapIn struct {
	DeleteFolder          *DeleteFolderType          `xml:"m:DeleteFolder,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type DeleteFolderSoapOut struct {
	DeleteFolderResponse *DeleteFolderResponseType `xml:"m:DeleteFolderResponse,omitempty"`
	ServerVersionInfo    *ServerVersionInfoType    `xml:"t:ServerVersionInfo,omitempty"`
}

type EmptyFolderSoapIn struct {
	EmptyFolder           *EmptyFolderType           `xml:"m:EmptyFolder,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type EmptyFolderSoapOut struct {
	EmptyFolderResponse *EmptyFolderResponseType `xml:"m:EmptyFolderResponse,omitempty"`
	ServerVersionInfo   *ServerVersionInfoType   `xml:"t:ServerVersionInfo,omitempty"`
}

type UpdateFolderSoapIn struct {
	UpdateFolder          *UpdateFolderType          `xml:"m:UpdateFolder,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	TimeZoneContext       *TimeZoneContextType       `xml:"t:TimeZoneContext,omitempty"`
}

type UpdateFolderSoapOut struct {
	UpdateFolderResponse *UpdateFolderResponseType `xml:"m:UpdateFolderResponse,omitempty"`
	ServerVersionInfo    *ServerVersionInfoType    `xml:"t:ServerVersionInfo,omitempty"`
}

type MoveFolderSoapIn struct {
	MoveFolder            *MoveFolderType            `xml:"m:MoveFolder,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type MoveFolderSoapOut struct {
	MoveFolderResponse *MoveFolderResponseType `xml:"m:MoveFolderResponse,omitempty"`
	ServerVersionInfo  *ServerVersionInfoType  `xml:"t:ServerVersionInfo,omitempty"`
}

type CopyFolderSoapIn struct {
	CopyFolder            *CopyFolderType            `xml:"m:CopyFolder,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type CopyFolderSoapOut struct {
	CopyFolderResponse *CopyFolderResponseType `xml:"m:CopyFolderResponse,omitempty"`
	ServerVersionInfo  *ServerVersionInfoType  `xml:"t:ServerVersionInfo,omitempty"`
}

type SubscribeSoapIn struct {
	Subscribe             *SubscribeType             `xml:"m:Subscribe,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type SubscribeSoapOut struct {
	SubscribeResponse *SubscribeResponseType `xml:"m:SubscribeResponse,omitempty"`
	ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
}

type UnsubscribeSoapIn struct {
	Unsubscribe           *UnsubscribeType           `xml:"m:Unsubscribe,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type UnsubscribeSoapOut struct {
	UnsubscribeResponse *UnsubscribeResponseType `xml:"m:UnsubscribeResponse,omitempty"`
	ServerVersionInfo   *ServerVersionInfoType   `xml:"t:ServerVersionInfo,omitempty"`
}

type GetEventsSoapIn struct {
	GetEvents             *GetEventsType             `xml:"m:GetEvents,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type GetEventsSoapOut struct {
	GetEventsResponse *GetEventsResponseType `xml:"m:GetEventsResponse,omitempty"`
	ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
}

type GetStreamingEventsSoapIn struct {
	GetStreamingEvents    *GetStreamingEventsType    `xml:"m:GetStreamingEvents,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type GetStreamingEventsSoapOut struct {
	GetStreamingEventsResponse *GetStreamingEventsResponseType `xml:"m:GetStreamingEventsResponse,omitempty"`
	ServerVersionInfo          *ServerVersionInfoType          `xml:"t:ServerVersionInfo,omitempty"`
}

type SyncFolderHierarchySoapIn struct {
	SyncFolderHierarchy   *SyncFolderHierarchyType   `xml:"m:SyncFolderHierarchy,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type SyncFolderHierarchySoapOut struct {
	SyncFolderHierarchyResponse *SyncFolderHierarchyResponseType `xml:"m:SyncFolderHierarchyResponse,omitempty"`
	ServerVersionInfo           *ServerVersionInfoType           `xml:"t:ServerVersionInfo,omitempty"`
}

type SyncFolderItemsSoapIn struct {
	SyncFolderItems       *SyncFolderItemsType       `xml:"m:SyncFolderItems,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type SyncFolderItemsSoapOut struct {
	SyncFolderItemsResponse *SyncFolderItemsResponseType `xml:"m:SyncFolderItemsResponse,omitempty"`
	ServerVersionInfo       *ServerVersionInfoType       `xml:"t:ServerVersionInfo,omitempty"`
}

type CreateManagedFolderSoapIn struct {
	CreateManagedFolder   *CreateManagedFolderRequestType `xml:"m:CreateManagedFolder,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType      `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType             `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType       `xml:"t:RequestServerVersion,omitempty"`
}

type CreateManagedFolderSoapOut struct {
	CreateManagedFolderResponse *CreateManagedFolderResponseType `xml:"m:CreateManagedFolderResponse,omitempty"`
	ServerVersionInfo           *ServerVersionInfoType           `xml:"t:ServerVersionInfo,omitempty"`
}

type GetItemSoapIn struct {
	GetItem               *GetItemType               `xml:"m:GetItem,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	TimeZoneContext       *TimeZoneContextType       `xml:"t:TimeZoneContext,omitempty"`
	DateTimePrecision     DateTimePrecisionType      `xml:"t:DateTimePrecision,omitempty"`
	ManagementRole        *ManagementRoleType        `xml:"t:ManagementRole,omitempty"`
}

type GetItemSoapOut struct {
	GetItemResponse   *GetItemResponseType   `xml:"m:GetItemResponse,omitempty"`
	ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
}

type CreateItemSoapIn struct {
	CreateItem            *CreateItemType            `xml:"m:CreateItem,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	TimeZoneContext       *TimeZoneContextType       `xml:"t:TimeZoneContext,omitempty"`
}

type CreateItemSoapOut struct {
	CreateItemResponse *CreateItemResponseType `xml:"m:CreateItemResponse,omitempty"`
	ServerVersionInfo  *ServerVersionInfoType  `xml:"t:ServerVersionInfo,omitempty"`
}

type DeleteItemSoapIn struct {
	DeleteItem            *DeleteItemType            `xml:"m:DeleteItem,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type DeleteItemSoapOut struct {
	DeleteItemResponse *DeleteItemResponseType `xml:"m:DeleteItemResponse,omitempty"`
	ServerVersionInfo  *ServerVersionInfoType  `xml:"t:ServerVersionInfo,omitempty"`
}

type UpdateItemSoapIn struct {
	UpdateItem            *UpdateItemType            `xml:"m:UpdateItem,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	TimeZoneContext       *TimeZoneContextType       `xml:"t:TimeZoneContext,omitempty"`
}

type UpdateItemSoapOut struct {
	UpdateItemResponse *UpdateItemResponseType `xml:"m:UpdateItemResponse,omitempty"`
	ServerVersionInfo  *ServerVersionInfoType  `xml:"t:ServerVersionInfo,omitempty"`
}

type UpdateItemInRecoverableItemsSoapIn struct {
	UpdateItemInRecoverableItems *UpdateItemInRecoverableItemsType `xml:"m:UpdateItemInRecoverableItems,omitempty"`
	ExchangeImpersonation        *ExchangeImpersonationType        `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture               *MailboxCultureType               `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion         *RequestServerVersionType         `xml:"t:RequestServerVersion,omitempty"`
	TimeZoneContext              *TimeZoneContextType              `xml:"t:TimeZoneContext,omitempty"`
	ManagementRole               *ManagementRoleType               `xml:"t:ManagementRole,omitempty"`
}

type UpdateItemInRecoverableItemsSoapOut struct {
	UpdateItemInRecoverableItemsResponse *UpdateItemInRecoverableItemsResponseType `xml:"m:UpdateItemInRecoverableItemsResponse,omitempty"`
	ServerVersionInfo                    *ServerVersionInfoType                    `xml:"t:ServerVersionInfo,omitempty"`
}

type SendItemSoapIn struct {
	SendItem              *SendItemType              `xml:"m:SendItem,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type SendItemSoapOut struct {
	SendItemResponse  *SendItemResponseType  `xml:"m:SendItemResponse,omitempty"`
	ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
}

type MoveItemSoapIn struct {
	MoveItem              *MoveItemType              `xml:"m:MoveItem,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type MoveItemSoapOut struct {
	MoveItemResponse  *MoveItemResponseType  `xml:"m:MoveItemResponse,omitempty"`
	ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
}

type CopyItemSoapIn struct {
	CopyItem              *CopyItemType              `xml:"m:CopyItem,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type CopyItemSoapOut struct {
	CopyItemResponse  *CopyItemResponseType  `xml:"m:CopyItemResponse,omitempty"`
	ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
}

type ArchiveItemSoapIn struct {
	ArchiveItem           *ArchiveItemType           `xml:"m:ArchiveItem,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type ArchiveItemSoapOut struct {
	ArchiveItemResponse *ArchiveItemResponseType `xml:"m:ArchiveItemResponse,omitempty"`
	ServerVersionInfo   *ServerVersionInfoType   `xml:"t:ServerVersionInfo,omitempty"`
}

type CreateAttachmentSoapIn struct {
	CreateAttachment      *CreateAttachmentType      `xml:"m:CreateAttachment,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	TimeZoneContext       *TimeZoneContextType       `xml:"t:TimeZoneContext,omitempty"`
}

type CreateAttachmentSoapOut struct {
	CreateAttachmentResponse *CreateAttachmentResponseType `xml:"m:CreateAttachmentResponse,omitempty"`
	ServerVersionInfo        *ServerVersionInfoType        `xml:"t:ServerVersionInfo,omitempty"`
}

type DeleteAttachmentSoapIn struct {
	DeleteAttachment      *DeleteAttachmentType      `xml:"m:DeleteAttachment,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type DeleteAttachmentSoapOut struct {
	DeleteAttachmentResponse *DeleteAttachmentResponseType `xml:"m:DeleteAttachmentResponse,omitempty"`
	ServerVersionInfo        *ServerVersionInfoType        `xml:"t:ServerVersionInfo,omitempty"`
}

type GetAttachmentSoapIn struct {
	GetAttachment         *GetAttachmentType         `xml:"m:GetAttachment,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	TimeZoneContext       *TimeZoneContextType       `xml:"t:TimeZoneContext,omitempty"`
}

type GetAttachmentSoapOut struct {
	GetAttachmentResponse *GetAttachmentResponseType `xml:"m:GetAttachmentResponse,omitempty"`
	ServerVersionInfo     *ServerVersionInfoType     `xml:"t:ServerVersionInfo,omitempty"`
}

type GetClientAccessTokenSoapIn struct {
	GetClientAccessToken *GetClientAccessTokenType `xml:"m:GetClientAccessToken,omitempty"`
	RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
}

type GetClientAccessTokenSoapOut struct {
	GetClientAccessTokenResponse *GetClientAccessTokenResponseType `xml:"m:GetClientAccessTokenResponse,omitempty"`
	ServerVersionInfo            *ServerVersionInfoType            `xml:"t:ServerVersionInfo,omitempty"`
}

type GetDelegateSoapIn struct {
	GetDelegate           *GetDelegateType           `xml:"m:GetDelegate,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type GetDelegateSoapOut struct {
	GetDelegateResponse *GetDelegateResponseMessageType `xml:"m:GetDelegateResponse,omitempty"`
	ServerVersionInfo   *ServerVersionInfoType          `xml:"t:ServerVersionInfo,omitempty"`
}

type AddDelegateSoapIn struct {
	AddDelegate           *AddDelegateType           `xml:"m:AddDelegate,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type AddDelegateSoapOut struct {
	AddDelegateResponse *AddDelegateResponseMessageType `xml:"m:AddDelegateResponse,omitempty"`
	ServerVersionInfo   *ServerVersionInfoType          `xml:"t:ServerVersionInfo,omitempty"`
}

type RemoveDelegateSoapIn struct {
	RemoveDelegate        *RemoveDelegateType        `xml:"m:RemoveDelegate,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type RemoveDelegateSoapOut struct {
	RemoveDelegateResponse *RemoveDelegateResponseMessageType `xml:"m:RemoveDelegateResponse,omitempty"`
	ServerVersionInfo      *ServerVersionInfoType             `xml:"t:ServerVersionInfo,omitempty"`
}

type UpdateDelegateSoapIn struct {
	UpdateDelegate        *UpdateDelegateType        `xml:"m:UpdateDelegate,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type UpdateDelegateSoapOut struct {
	UpdateDelegateResponse *UpdateDelegateResponseMessageType `xml:"m:UpdateDelegateResponse,omitempty"`
	ServerVersionInfo      *ServerVersionInfoType             `xml:"t:ServerVersionInfo,omitempty"`
}

type CreateUserConfigurationSoapIn struct {
	CreateUserConfiguration *CreateUserConfigurationType `xml:"m:CreateUserConfiguration,omitempty"`
	ExchangeImpersonation   *ExchangeImpersonationType   `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture          *MailboxCultureType          `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion    *RequestServerVersionType    `xml:"t:RequestServerVersion,omitempty"`
}

type CreateUserConfigurationSoapOut struct {
	CreateUserConfigurationResponse *CreateUserConfigurationResponseType `xml:"m:CreateUserConfigurationResponse,omitempty"`
	ServerVersionInfo               *ServerVersionInfoType               `xml:"t:ServerVersionInfo,omitempty"`
}

type DeleteUserConfigurationSoapIn struct {
	DeleteUserConfiguration *DeleteUserConfigurationType `xml:"m:DeleteUserConfiguration,omitempty"`
	ExchangeImpersonation   *ExchangeImpersonationType   `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture          *MailboxCultureType          `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion    *RequestServerVersionType    `xml:"t:RequestServerVersion,omitempty"`
}

type DeleteUserConfigurationSoapOut struct {
	DeleteUserConfigurationResponse *DeleteUserConfigurationResponseType `xml:"m:DeleteUserConfigurationResponse,omitempty"`
	ServerVersionInfo               *ServerVersionInfoType               `xml:"t:ServerVersionInfo,omitempty"`
}

type GetUserConfigurationSoapIn struct {
	GetUserConfiguration  *GetUserConfigurationType  `xml:"m:GetUserConfiguration,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type GetUserConfigurationSoapOut struct {
	GetUserConfigurationResponse *GetUserConfigurationResponseType `xml:"m:GetUserConfigurationResponse,omitempty"`
	ServerVersionInfo            *ServerVersionInfoType            `xml:"t:ServerVersionInfo,omitempty"`
}

type GetSpecificUserConfigurationSoapIn struct {
	GetSpecificUserConfiguration *GetSpecificUserConfigurationType `xml:"m:GetSpecificUserConfiguration,omitempty"`
	ExchangeImpersonation        *ExchangeImpersonationType        `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture               *MailboxCultureType               `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion         *RequestServerVersionType         `xml:"t:RequestServerVersion,omitempty"`
}

type GetSpecificUserConfigurationSoapOut struct {
	GetSpecificUserConfigurationResponse *GetSpecificUserConfigurationResponseType `xml:"m:GetSpecificUserConfigurationResponse,omitempty"`
	ServerVersionInfo                    *ServerVersionInfoType                    `xml:"t:ServerVersionInfo,omitempty"`
}

type UpdateUserConfigurationSoapIn struct {
	UpdateUserConfiguration *UpdateUserConfigurationType `xml:"m:UpdateUserConfiguration,omitempty"`
	ExchangeImpersonation   *ExchangeImpersonationType   `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture          *MailboxCultureType          `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion    *RequestServerVersionType    `xml:"t:RequestServerVersion,omitempty"`
}

type UpdateUserConfigurationSoapOut struct {
	UpdateUserConfigurationResponse *UpdateUserConfigurationResponseType `xml:"m:UpdateUserConfigurationResponse,omitempty"`
	ServerVersionInfo               *ServerVersionInfoType               `xml:"t:ServerVersionInfo,omitempty"`
}

type GetUserAvailabilitySoapIn struct {
	GetUserAvailabilityRequest *GetUserAvailabilityRequestType `xml:"m:GetUserAvailabilityRequest,omitempty"`
	ExchangeImpersonation      *ExchangeImpersonationType      `xml:"t:ExchangeImpersonation,omitempty"`
	TimeZoneContext            *TimeZoneContextType            `xml:"t:TimeZoneContext,omitempty"`
	RequestServerVersion       *RequestServerVersionType       `xml:"t:RequestServerVersion,omitempty"`
}

type GetUserAvailabilitySoapOut struct {
	GetUserAvailabilityResponse *GetUserAvailabilityResponseType `xml:"m:GetUserAvailabilityResponse,omitempty"`
	ServerVersionInfo           *ServerVersionInfoType           `xml:"t:ServerVersionInfo,omitempty"`
}

type GetUserOofSettingsSoapIn struct {
	GetUserOofSettingsRequest *GetUserOofSettingsRequest `xml:"m:GetUserOofSettingsRequest,omitempty"`
	ExchangeImpersonation     *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	RequestServerVersion      *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type GetUserOofSettingsSoapOut struct {
	GetUserOofSettingsResponse *GetUserOofSettingsResponse `xml:"m:GetUserOofSettingsResponse,omitempty"`
	ServerVersionInfo          *ServerVersionInfoType      `xml:"t:ServerVersionInfo,omitempty"`
}

type SetUserOofSettingsSoapIn struct {
	SetUserOofSettingsRequest *SetUserOofSettingsRequest `xml:"m:SetUserOofSettingsRequest,omitempty"`
	ExchangeImpersonation     *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	RequestServerVersion      *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type SetUserOofSettingsSoapOut struct {
	SetUserOofSettingsResponse *SetUserOofSettingsResponse `xml:"m:SetUserOofSettingsResponse,omitempty"`
	ServerVersionInfo          *ServerVersionInfoType      `xml:"t:ServerVersionInfo,omitempty"`
}

type GetServiceConfigurationSoapIn struct {
	GetServiceConfiguration *GetServiceConfigurationType `xml:"m:GetServiceConfiguration,omitempty"`
	ExchangeImpersonation   *ExchangeImpersonationType   `xml:"t:ExchangeImpersonation,omitempty"`
	RequestServerVersion    *RequestServerVersionType    `xml:"t:RequestServerVersion,omitempty"`
	MailboxCulture          *MailboxCultureType          `xml:"t:MailboxCulture,omitempty"`
}

type GetServiceConfigurationSoapOut struct {
	GetServiceConfigurationResponse *GetServiceConfigurationResponseMessageType `xml:"m:GetServiceConfigurationResponse,omitempty"`
	ServerVersionInfo               *ServerVersionInfoType                      `xml:"t:ServerVersionInfo,omitempty"`
}

type GetMailTipsSoapIn struct {
	GetMailTips          *GetMailTipsType          `xml:"m:GetMailTips,omitempty"`
	RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
	MailboxCulture       *MailboxCultureType       `xml:"t:MailboxCulture,omitempty"`
}

type GetMailTipsSoapOut struct {
	GetMailTipsResponse *GetMailTipsResponseMessageType `xml:"m:GetMailTipsResponse,omitempty"`
	ServerVersionInfo   *ServerVersionInfoType          `xml:"t:ServerVersionInfo,omitempty"`
}

type PlayOnPhoneSoapIn struct {
	PlayOnPhone           *PlayOnPhoneType           `xml:"m:PlayOnPhone,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type PlayOnPhoneSoapOut struct {
	PlayOnPhoneResponse *PlayOnPhoneResponseMessageType `xml:"m:PlayOnPhoneResponse,omitempty"`
	ServerVersionInfo   *ServerVersionInfoType          `xml:"t:ServerVersionInfo,omitempty"`
}

type GetPhoneCallInformationSoapIn struct {
	GetPhoneCallInformation *GetPhoneCallInformationType `xml:"m:GetPhoneCallInformation,omitempty"`
	ExchangeImpersonation   *ExchangeImpersonationType   `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture          *MailboxCultureType          `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion    *RequestServerVersionType    `xml:"t:RequestServerVersion,omitempty"`
}

type GetPhoneCallInformationSoapOut struct {
	GetPhoneCallInformationResponse *GetPhoneCallInformationResponseMessageType `xml:"m:GetPhoneCallInformationResponse,omitempty"`
	ServerVersionInfo               *ServerVersionInfoType                      `xml:"t:ServerVersionInfo,omitempty"`
}

type DisconnectPhoneCallSoapIn struct {
	DisconnectPhoneCall   *DisconnectPhoneCallType   `xml:"m:DisconnectPhoneCall,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type DisconnectPhoneCallSoapOut struct {
	DisconnectPhoneCallResponse *DisconnectPhoneCallResponseMessageType `xml:"m:DisconnectPhoneCallResponse,omitempty"`
	ServerVersionInfo           *ServerVersionInfoType                  `xml:"t:ServerVersionInfo,omitempty"`
}

type GetSharingMetadataSoapIn struct {
	GetSharingMetadata   *GetSharingMetadataType   `xml:"m:GetSharingMetadata,omitempty"`
	RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
}

type GetSharingMetadataSoapOut struct {
	GetSharingMetadataResponse *GetSharingMetadataResponseMessageType `xml:"m:GetSharingMetadataResponse,omitempty"`
	ServerVersionInfo          *ServerVersionInfoType                 `xml:"t:ServerVersionInfo,omitempty"`
}

type RefreshSharingFolderSoapIn struct {
	RefreshSharingFolder *RefreshSharingFolderType `xml:"m:RefreshSharingFolder,omitempty"`
	RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
}

type RefreshSharingFolderSoapOut struct {
	RefreshSharingFolderResponse *RefreshSharingFolderResponseMessageType `xml:"m:RefreshSharingFolderResponse,omitempty"`
	ServerVersionInfo            *ServerVersionInfoType                   `xml:"t:ServerVersionInfo,omitempty"`
}

type GetSharingFolderSoapIn struct {
	GetSharingFolder     *GetSharingFolderType     `xml:"m:GetSharingFolder,omitempty"`
	RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
}

type GetSharingFolderSoapOut struct {
	GetSharingFolderResponse *GetSharingFolderResponseMessageType `xml:"m:GetSharingFolderResponse,omitempty"`
	ServerVersionInfo        *ServerVersionInfoType               `xml:"t:ServerVersionInfo,omitempty"`
}

type SetTeamMailboxSoapIn struct {
	SetTeamMailbox       *SetTeamMailboxRequestType `xml:"m:SetTeamMailbox,omitempty"`
	RequestServerVersion *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	ManagementRole       *ManagementRoleType        `xml:"t:ManagementRole,omitempty"`
}

type SetTeamMailboxSoapOut struct {
	SetTeamMailboxResponse *SetTeamMailboxResponseMessageType `xml:"m:SetTeamMailboxResponse,omitempty"`
	ServerVersionInfo      *ServerVersionInfoType             `xml:"t:ServerVersionInfo,omitempty"`
}

type UnpinTeamMailboxSoapIn struct {
	UnpinTeamMailbox     *UnpinTeamMailboxRequestType `xml:"m:UnpinTeamMailbox,omitempty"`
	RequestServerVersion *RequestServerVersionType    `xml:"t:RequestServerVersion,omitempty"`
}

type UnpinTeamMailboxSoapOut struct {
	UnpinTeamMailboxResponse *UnpinTeamMailboxResponseMessageType `xml:"m:UnpinTeamMailboxResponse,omitempty"`
	ServerVersionInfo        *ServerVersionInfoType               `xml:"t:ServerVersionInfo,omitempty"`
}

type GetRoomListsSoapIn struct {
	GetRoomLists          *GetRoomListsType          `xml:"m:GetRoomLists,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type GetRoomListsSoapOut struct {
	GetRoomListsResponse *GetRoomListsResponseMessageType `xml:"m:GetRoomListsResponse,omitempty"`
	ServerVersionInfo    *ServerVersionInfoType           `xml:"t:ServerVersionInfo,omitempty"`
}

type GetRoomsSoapIn struct {
	GetRooms              *GetRoomsType              `xml:"m:GetRooms,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type GetRoomsSoapOut struct {
	GetRoomsResponse  *GetRoomsResponseMessageType `xml:"m:GetRoomsResponse,omitempty"`
	ServerVersionInfo *ServerVersionInfoType       `xml:"t:ServerVersionInfo,omitempty"`
}

type FindMessageTrackingReportSoapIn struct {
	FindMessageTrackingReport *FindMessageTrackingReportRequestType `xml:"m:FindMessageTrackingReport,omitempty"`
	RequestServerVersion      *RequestServerVersionType             `xml:"t:RequestServerVersion,omitempty"`
}

type FindMessageTrackingReportSoapOut struct {
	FindMessageTrackingReportResponse *FindMessageTrackingReportResponseMessageType `xml:"m:FindMessageTrackingReportResponse,omitempty"`
	ServerVersionInfo                 *ServerVersionInfoType                        `xml:"t:ServerVersionInfo,omitempty"`
}

type GetMessageTrackingReportSoapIn struct {
	GetMessageTrackingReport *GetMessageTrackingReportRequestType `xml:"m:GetMessageTrackingReport,omitempty"`
	RequestServerVersion     *RequestServerVersionType            `xml:"t:RequestServerVersion,omitempty"`
}

type GetMessageTrackingReportSoapOut struct {
	GetMessageTrackingReportResponse *GetMessageTrackingReportResponseMessageType `xml:"m:GetMessageTrackingReportResponse,omitempty"`
	ServerVersionInfo                *ServerVersionInfoType                       `xml:"t:ServerVersionInfo,omitempty"`
}

type FindConversationSoapIn struct {
	FindConversation      *FindConversationType      `xml:"m:FindConversation,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type FindConversationSoapOut struct {
	FindConversationResponse *FindConversationResponseMessageType `xml:"m:FindConversationResponse,omitempty"`
	ServerVersionInfo        *ServerVersionInfoType               `xml:"t:ServerVersionInfo,omitempty"`
}

type ApplyConversationActionSoapIn struct {
	ApplyConversationAction *ApplyConversationActionType `xml:"m:ApplyConversationAction,omitempty"`
	ExchangeImpersonation   *ExchangeImpersonationType   `xml:"t:ExchangeImpersonation,omitempty"`
	RequestServerVersion    *RequestServerVersionType    `xml:"t:RequestServerVersion,omitempty"`
}

type ApplyConversationActionSoapOut struct {
	ApplyConversationActionResponse *ApplyConversationActionResponseType `xml:"m:ApplyConversationActionResponse,omitempty"`
	ServerVersionInfo               *ServerVersionInfoType               `xml:"t:ServerVersionInfo,omitempty"`
}

type GetConversationItemsSoapIn struct {
	GetConversationItems  *GetConversationItemsType  `xml:"m:GetConversationItems,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type GetConversationItemsSoapOut struct {
	GetConversationItemsResponse *GetConversationItemsResponseType `xml:"m:GetConversationItemsResponse,omitempty"`
	ServerVersionInfo            *ServerVersionInfoType            `xml:"t:ServerVersionInfo,omitempty"`
}

type FindPeopleSoapIn struct {
	FindPeople            *FindPeopleType            `xml:"m:FindPeople,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type FindPeopleSoapOut struct {
	FindPeopleResponse *FindPeopleResponseMessageType `xml:"m:FindPeopleResponse,omitempty"`
	ServerVersionInfo  *ServerVersionInfoType         `xml:"t:ServerVersionInfo,omitempty"`
}

type FindTagsSoapIn struct {
	FindTags              *FindTagsType              `xml:"m:FindTags,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type FindTagsSoapOut struct {
	FindTagsResponse  *FindTagsResponseMessageType `xml:"m:FindTagsResponse,omitempty"`
	ServerVersionInfo *ServerVersionInfoType       `xml:"t:ServerVersionInfo,omitempty"`
}

type AddTagSoapIn struct {
	AddTag                *AddTagType                `xml:"m:AddTag,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type AddTagSoapOut struct {
	AddTagResponse    *AddTagResponseMessageType `xml:"m:AddTagResponse,omitempty"`
	ServerVersionInfo *ServerVersionInfoType     `xml:"t:ServerVersionInfo,omitempty"`
}

type HideTagSoapIn struct {
	HideTag               *HideTagType               `xml:"m:HideTag,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type HideTagSoapOut struct {
	HideTagResponse   *HideTagResponseMessageType `xml:"m:HideTagResponse,omitempty"`
	ServerVersionInfo *ServerVersionInfoType      `xml:"t:ServerVersionInfo,omitempty"`
}

type GetPersonaSoapIn struct {
	GetPersona            *GetPersonaType            `xml:"m:GetPersona,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type GetPersonaSoapOut struct {
	GetPersonaResponseMessage *GetPersonaResponseMessageType `xml:"m:GetPersonaResponseMessage,omitempty"`
	ServerVersionInfo         *ServerVersionInfoType         `xml:"t:ServerVersionInfo,omitempty"`
}

type GetInboxRulesSoapIn struct {
	GetInboxRules         *GetInboxRulesRequestType  `xml:"m:GetInboxRules,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	TimeZoneContext       *TimeZoneContextType       `xml:"t:TimeZoneContext,omitempty"`
}

type GetInboxRulesSoapOut struct {
	GetInboxRulesResponse *GetInboxRulesResponseType `xml:"m:GetInboxRulesResponse,omitempty"`
	ServerVersionInfo     *ServerVersionInfoType     `xml:"t:ServerVersionInfo,omitempty"`
}

type UpdateInboxRulesSoapIn struct {
	UpdateInboxRules      *UpdateInboxRulesRequestType `xml:"m:UpdateInboxRules,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType   `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType          `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType    `xml:"t:RequestServerVersion,omitempty"`
	TimeZoneContext       *TimeZoneContextType         `xml:"t:TimeZoneContext,omitempty"`
}

type UpdateInboxRulesSoapOut struct {
	UpdateInboxRulesResponse *UpdateInboxRulesResponseType `xml:"m:UpdateInboxRulesResponse,omitempty"`
	ServerVersionInfo        *ServerVersionInfoType        `xml:"t:ServerVersionInfo,omitempty"`
}

type GetPasswordExpirationDateSoapIn struct {
	GetPasswordExpirationDate *GetPasswordExpirationDateType `xml:"m:GetPasswordExpirationDate,omitempty"`
	MailboxCulture            *MailboxCultureType            `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion      *RequestServerVersionType      `xml:"t:RequestServerVersion,omitempty"`
}

type GetPasswordExpirationDateSoapOut struct {
	GetPasswordExpirationDateResponse *GetPasswordExpirationDateResponseMessageType `xml:"m:GetPasswordExpirationDateResponse,omitempty"`
	ServerVersionInfo                 *ServerVersionInfoType                        `xml:"t:ServerVersionInfo,omitempty"`
}

type GetSearchableMailboxesSoapIn struct {
	GetSearchableMailboxes *GetSearchableMailboxesType `xml:"m:GetSearchableMailboxes,omitempty"`
	RequestServerVersion   *RequestServerVersionType   `xml:"t:RequestServerVersion,omitempty"`
	ManagementRole         *ManagementRoleType         `xml:"t:ManagementRole,omitempty"`
}

type GetSearchableMailboxesSoapOut struct {
	GetSearchableMailboxesResponse *GetSearchableMailboxesResponseMessageType `xml:"m:GetSearchableMailboxesResponse,omitempty"`
	ServerVersionInfo              *ServerVersionInfoType                     `xml:"t:ServerVersionInfo,omitempty"`
}

type SearchMailboxesSoapIn struct {
	SearchMailboxes      *SearchMailboxesType      `xml:"m:SearchMailboxes,omitempty"`
	RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
	ManagementRole       *ManagementRoleType       `xml:"t:ManagementRole,omitempty"`
}

type SearchMailboxesSoapOut struct {
	SearchMailboxesResponse *SearchMailboxesResponseType `xml:"m:SearchMailboxesResponse,omitempty"`
	ServerVersionInfo       *ServerVersionInfoType       `xml:"t:ServerVersionInfo,omitempty"`
}

type GetDiscoverySearchConfigurationSoapIn struct {
	GetDiscoverySearchConfiguration *GetDiscoverySearchConfigurationType `xml:"m:GetDiscoverySearchConfiguration,omitempty"`
	RequestServerVersion            *RequestServerVersionType            `xml:"t:RequestServerVersion,omitempty"`
	ManagementRole                  *ManagementRoleType                  `xml:"t:ManagementRole,omitempty"`
}

type GetDiscoverySearchConfigurationSoapOut struct {
	GetDiscoverySearchConfigurationResponse *GetDiscoverySearchConfigurationResponseMessageType `xml:"m:GetDiscoverySearchConfigurationResponse,omitempty"`
	ServerVersionInfo                       *ServerVersionInfoType                              `xml:"t:ServerVersionInfo,omitempty"`
}

type GetHoldOnMailboxesSoapIn struct {
	GetHoldOnMailboxes   *GetHoldOnMailboxesType   `xml:"m:GetHoldOnMailboxes,omitempty"`
	RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
	ManagementRole       *ManagementRoleType       `xml:"t:ManagementRole,omitempty"`
}

type GetHoldOnMailboxesSoapOut struct {
	GetHoldOnMailboxesResponse *GetHoldOnMailboxesResponseMessageType `xml:"m:GetHoldOnMailboxesResponse,omitempty"`
	ServerVersionInfo          *ServerVersionInfoType                 `xml:"t:ServerVersionInfo,omitempty"`
}

type SetHoldOnMailboxesSoapIn struct {
	SetHoldOnMailboxes   *SetHoldOnMailboxesType   `xml:"m:SetHoldOnMailboxes,omitempty"`
	RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
	ManagementRole       *ManagementRoleType       `xml:"t:ManagementRole,omitempty"`
}

type SetHoldOnMailboxesSoapOut struct {
	SetHoldOnMailboxesResponse *SetHoldOnMailboxesResponseMessageType `xml:"m:SetHoldOnMailboxesResponse,omitempty"`
	ServerVersionInfo          *ServerVersionInfoType                 `xml:"t:ServerVersionInfo,omitempty"`
}

type GetNonIndexableItemStatisticsSoapIn struct {
	GetNonIndexableItemStatistics *GetNonIndexableItemStatisticsType `xml:"m:GetNonIndexableItemStatistics,omitempty"`
	RequestServerVersion          *RequestServerVersionType          `xml:"t:RequestServerVersion,omitempty"`
	ManagementRole                *ManagementRoleType                `xml:"t:ManagementRole,omitempty"`
}

type GetNonIndexableItemStatisticsSoapOut struct {
	GetNonIndexableItemStatisticsResponse *GetNonIndexableItemStatisticsResponseMessageType `xml:"m:GetNonIndexableItemStatisticsResponse,omitempty"`
	ServerVersionInfo                     *ServerVersionInfoType                            `xml:"t:ServerVersionInfo,omitempty"`
}

type GetNonIndexableItemDetailsSoapIn struct {
	GetNonIndexableItemDetails *GetNonIndexableItemDetailsType `xml:"m:GetNonIndexableItemDetails,omitempty"`
	RequestServerVersion       *RequestServerVersionType       `xml:"t:RequestServerVersion,omitempty"`
	ManagementRole             *ManagementRoleType             `xml:"t:ManagementRole,omitempty"`
}

type GetNonIndexableItemDetailsSoapOut struct {
	GetNonIndexableItemDetailsResponse *GetNonIndexableItemDetailsResponseMessageType `xml:"m:GetNonIndexableItemDetailsResponse,omitempty"`
	ServerVersionInfo                  *ServerVersionInfoType                         `xml:"t:ServerVersionInfo,omitempty"`
}

type MarkAllItemsAsReadSoapIn struct {
	MarkAllItemsAsRead    *MarkAllItemsAsReadType    `xml:"m:MarkAllItemsAsRead,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type MarkAllItemsAsReadSoapOut struct {
	MarkAllItemsAsReadResponse *MarkAllItemsAsReadResponseType `xml:"m:MarkAllItemsAsReadResponse,omitempty"`
	ServerVersionInfo          *ServerVersionInfoType          `xml:"t:ServerVersionInfo,omitempty"`
}

type MarkAsJunkSoapIn struct {
	MarkAsJunk            *MarkAsJunkType            `xml:"m:MarkAsJunk,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type MarkAsJunkSoapOut struct {
	MarkAsJunkResponse *MarkAsJunkResponseType `xml:"m:MarkAsJunkResponse,omitempty"`
	ServerVersionInfo  *ServerVersionInfoType  `xml:"t:ServerVersionInfo,omitempty"`
}

type ReportMessageSoapIn struct {
	ReportMessage         *ReportMessageType         `xml:"m:ReportMessage,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type ReportMessageSoapOut struct {
	ReportMessageResponse *ReportMessageResponseType `xml:"m:ReportMessageResponse,omitempty"`
	ServerVersionInfo     *ServerVersionInfoType     `xml:"t:ServerVersionInfo,omitempty"`
}

type GetAppManifestsSoapIn struct {
	GetAppManifests      *GetAppManifestsType      `xml:"m:GetAppManifests,omitempty"`
	RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
}

type GetAppManifestsSoapOut struct {
	GetAppManifestsResponse *GetAppManifestsResponseType `xml:"m:GetAppManifestsResponse,omitempty"`
	ServerVersionInfo       *ServerVersionInfoType       `xml:"t:ServerVersionInfo,omitempty"`
}

type AddNewImContactToGroupSoapIn struct {
	AddNewImContactToGroup *AddNewImContactToGroupType `xml:"m:AddNewImContactToGroup,omitempty"`
	ExchangeImpersonation  *ExchangeImpersonationType  `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture         *MailboxCultureType         `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion   *RequestServerVersionType   `xml:"t:RequestServerVersion,omitempty"`
}

type AddNewImContactToGroupSoapOut struct {
	AddNewImContactToGroupResponse *AddNewImContactToGroupResponseMessageType `xml:"m:AddNewImContactToGroupResponse,omitempty"`
	ServerVersionInfo              *ServerVersionInfoType                     `xml:"t:ServerVersionInfo,omitempty"`
}

type AddNewTelUriContactToGroupSoapIn struct {
	AddNewTelUriContactToGroup *AddNewTelUriContactToGroupType `xml:"m:AddNewTelUriContactToGroup,omitempty"`
	ExchangeImpersonation      *ExchangeImpersonationType      `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture             *MailboxCultureType             `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion       *RequestServerVersionType       `xml:"t:RequestServerVersion,omitempty"`
}

type AddNewTelUriContactToGroupSoapOut struct {
	AddNewTelUriContactToGroupResponse *AddNewTelUriContactToGroupResponseMessageType `xml:"m:AddNewTelUriContactToGroupResponse,omitempty"`
	ServerVersionInfo                  *ServerVersionInfoType                         `xml:"t:ServerVersionInfo,omitempty"`
}

type AddImContactToGroupSoapIn struct {
	AddImContactToGroup   *AddImContactToGroupType   `xml:"m:AddImContactToGroup,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type AddImContactToGroupSoapOut struct {
	AddImContactToGroupResponse *AddImContactToGroupResponseMessageType `xml:"m:AddImContactToGroupResponse,omitempty"`
	ServerVersionInfo           *ServerVersionInfoType                  `xml:"t:ServerVersionInfo,omitempty"`
}

type RemoveImContactFromGroupSoapIn struct {
	RemoveImContactFromGroup *RemoveImContactFromGroupType `xml:"m:RemoveImContactFromGroup,omitempty"`
	ExchangeImpersonation    *ExchangeImpersonationType    `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture           *MailboxCultureType           `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion     *RequestServerVersionType     `xml:"t:RequestServerVersion,omitempty"`
}

type RemoveImContactFromGroupSoapOut struct {
	RemoveImContactFromGroupResponse *RemoveImContactFromGroupResponseMessageType `xml:"m:RemoveImContactFromGroupResponse,omitempty"`
	ServerVersionInfo                *ServerVersionInfoType                       `xml:"t:ServerVersionInfo,omitempty"`
}

type AddImGroupSoapIn struct {
	AddImGroup            *AddImGroupType            `xml:"m:AddImGroup,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type AddImGroupSoapOut struct {
	AddImGroupResponse *AddImGroupResponseMessageType `xml:"m:AddImGroupResponse,omitempty"`
	ServerVersionInfo  *ServerVersionInfoType         `xml:"t:ServerVersionInfo,omitempty"`
}

type AddDistributionGroupToImListSoapIn struct {
	AddDistributionGroupToImList *AddDistributionGroupToImListType `xml:"m:AddDistributionGroupToImList,omitempty"`
	ExchangeImpersonation        *ExchangeImpersonationType        `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture               *MailboxCultureType               `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion         *RequestServerVersionType         `xml:"t:RequestServerVersion,omitempty"`
}

type AddDistributionGroupToImListSoapOut struct {
	AddDistributionGroupToImListResponse *AddDistributionGroupToImListResponseMessageType `xml:"m:AddDistributionGroupToImListResponse,omitempty"`
	ServerVersionInfo                    *ServerVersionInfoType                           `xml:"t:ServerVersionInfo,omitempty"`
}

type GetImItemListSoapIn struct {
	GetImItemList         *GetImItemListType         `xml:"m:GetImItemList,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type GetImItemListSoapOut struct {
	GetImItemListResponse *GetImItemListResponseMessageType `xml:"m:GetImItemListResponse,omitempty"`
	ServerVersionInfo     *ServerVersionInfoType            `xml:"t:ServerVersionInfo,omitempty"`
}

type GetImItemsSoapIn struct {
	GetImItems            *GetImItemsType            `xml:"m:GetImItems,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type GetImItemsSoapOut struct {
	GetImItemsResponse *GetImItemsResponseMessageType `xml:"m:GetImItemsResponse,omitempty"`
	ServerVersionInfo  *ServerVersionInfoType         `xml:"t:ServerVersionInfo,omitempty"`
}

type RemoveContactFromImListSoapIn struct {
	RemoveContactFromImList *RemoveContactFromImListType `xml:"m:RemoveContactFromImList,omitempty"`
	ExchangeImpersonation   *ExchangeImpersonationType   `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture          *MailboxCultureType          `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion    *RequestServerVersionType    `xml:"t:RequestServerVersion,omitempty"`
}

type RemoveContactFromImListSoapOut struct {
	RemoveContactFromImListResponse *RemoveContactFromImListResponseMessageType `xml:"m:RemoveContactFromImListResponse,omitempty"`
	ServerVersionInfo               *ServerVersionInfoType                      `xml:"t:ServerVersionInfo,omitempty"`
}

type RemoveDistributionGroupFromImListSoapIn struct {
	RemoveDistributionGroupFromImList *RemoveDistributionGroupFromImListType `xml:"m:RemoveDistributionGroupFromImList,omitempty"`
	ExchangeImpersonation             *ExchangeImpersonationType             `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture                    *MailboxCultureType                    `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion              *RequestServerVersionType              `xml:"t:RequestServerVersion,omitempty"`
}

type RemoveDistributionGroupFromImListSoapOut struct {
	RemoveDistributionGroupFromImListResponse *RemoveDistributionGroupFromImListResponseMessageType `xml:"m:RemoveDistributionGroupFromImListResponse,omitempty"`
	ServerVersionInfo                         *ServerVersionInfoType                                `xml:"t:ServerVersionInfo,omitempty"`
}

type RemoveImGroupSoapIn struct {
	RemoveImGroup         *RemoveImGroupType         `xml:"m:RemoveImGroup,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type RemoveImGroupSoapOut struct {
	RemoveImGroupResponse *RemoveImGroupResponseMessageType `xml:"m:RemoveImGroupResponse,omitempty"`
	ServerVersionInfo     *ServerVersionInfoType            `xml:"t:ServerVersionInfo,omitempty"`
}

type SetImGroupSoapIn struct {
	SetImGroup            *SetImGroupType            `xml:"m:SetImGroup,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type SetImGroupSoapOut struct {
	SetImGroupResponse *SetImGroupResponseMessageType `xml:"m:SetImGroupResponse,omitempty"`
	ServerVersionInfo  *ServerVersionInfoType         `xml:"t:ServerVersionInfo,omitempty"`
}

type SetImListMigrationCompletedSoapIn struct {
	SetImListMigrationCompleted *SetImListMigrationCompletedType `xml:"m:SetImListMigrationCompleted,omitempty"`
	ExchangeImpersonation       *ExchangeImpersonationType       `xml:"t:ExchangeImpersonation,omitempty"`
	MailboxCulture              *MailboxCultureType              `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion        *RequestServerVersionType        `xml:"t:RequestServerVersion,omitempty"`
}

type SetImListMigrationCompletedSoapOut struct {
	SetImListMigrationCompletedResponse *SetImListMigrationCompletedResponseMessageType `xml:"m:SetImListMigrationCompletedResponse,omitempty"`
	ServerVersionInfo                   *ServerVersionInfoType                          `xml:"t:ServerVersionInfo,omitempty"`
}

type GetUserRetentionPolicyTagsSoapIn struct {
	GetUserRetentionPolicyTags *GetUserRetentionPolicyTagsType `xml:"m:GetUserRetentionPolicyTags,omitempty"`
	RequestServerVersion       *RequestServerVersionType       `xml:"t:RequestServerVersion,omitempty"`
}

type GetUserRetentionPolicyTagsSoapOut struct {
	GetUserRetentionPolicyTagsResponse *GetUserRetentionPolicyTagsResponseMessageType `xml:"m:GetUserRetentionPolicyTagsResponse,omitempty"`
	ServerVersionInfo                  *ServerVersionInfoType                         `xml:"t:ServerVersionInfo,omitempty"`
}

type InstallAppSoapIn struct {
	InstallApp            *InstallAppType            `xml:"m:InstallApp,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type InstallAppSoapOut struct {
	InstallAppResponse *InstallAppResponseType `xml:"m:InstallAppResponse,omitempty"`
	ServerVersionInfo  *ServerVersionInfoType  `xml:"t:ServerVersionInfo,omitempty"`
}

type UpdateExtensionUsageSoapIn struct {
	UpdateExtensionUsage  *UpdateExtensionUsageType  `xml:"m:UpdateExtensionUsage,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type UpdateExtensionUsageSoapOut struct {
	UpdateExtensionUsageResponse *UpdateExtensionUsageResponseType `xml:"m:UpdateExtensionUsageResponse,omitempty"`
	ServerVersionInfo            *ServerVersionInfoType            `xml:"t:ServerVersionInfo,omitempty"`
}

type UninstallAppSoapIn struct {
	UninstallApp          *UninstallAppType          `xml:"m:UninstallApp,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type UninstallAppSoapOut struct {
	UninstallAppResponse *UninstallAppResponseType `xml:"m:UninstallAppResponse,omitempty"`
	ServerVersionInfo    *ServerVersionInfoType    `xml:"t:ServerVersionInfo,omitempty"`
}

type DisableAppSoapIn struct {
	DisableApp            *DisableAppType            `xml:"m:DisableApp,omitempty"`
	ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
}

type DisableAppSoapOut struct {
	DisableAppResponse *DisableAppResponseType `xml:"m:DisableAppResponse,omitempty"`
	ServerVersionInfo  *ServerVersionInfoType  `xml:"t:ServerVersionInfo,omitempty"`
}

type GetAppMarketplaceUrlSoapIn struct {
	GetAppMarketplaceUrl *GetAppMarketplaceUrlType `xml:"m:GetAppMarketplaceUrl,omitempty"`
	RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
}

type GetAppMarketplaceUrlSoapOut struct {
	GetAppMarketplaceUrlResponse *GetAppMarketplaceUrlResponseMessageType `xml:"m:GetAppMarketplaceUrlResponse,omitempty"`
	ServerVersionInfo            *ServerVersionInfoType                   `xml:"t:ServerVersionInfo,omitempty"`
}

type FindAvailableMeetingTimesSoapIn struct {
	FindAvailableMeetingTimes *FindAvailableMeetingTimesType `xml:"m:FindAvailableMeetingTimes,omitempty"`
	RequestServerVersion      *RequestServerVersionType      `xml:"t:RequestServerVersion,omitempty"`
}

type FindAvailableMeetingTimesSoapOut struct {
	FindAvailableMeetingTimesResponse *FindAvailableMeetingTimesResponseMessageType `xml:"m:FindAvailableMeetingTimesResponse,omitempty"`
	ServerVersionInfo                 *ServerVersionInfoType                        `xml:"t:ServerVersionInfo,omitempty"`
}

type FindMeetingTimeCandidatesSoapIn struct {
	FindMeetingTimeCandidates *FindMeetingTimeCandidatesType `xml:"m:FindMeetingTimeCandidates,omitempty"`
	RequestServerVersion      *RequestServerVersionType      `xml:"t:RequestServerVersion,omitempty"`
	TimeZoneContext           *TimeZoneContextType           `xml:"t:TimeZoneContext,omitempty"`
}

type FindMeetingTimeCandidatesSoapOut struct {
	FindMeetingTimeCandidatesResponse *FindMeetingTimeCandidatesResponseMessageType `xml:"m:FindMeetingTimeCandidatesResponse,omitempty"`
	ServerVersionInfo                 *ServerVersionInfoType                        `xml:"t:ServerVersionInfo,omitempty"`
}

type GetUserPhotoSoapIn struct {
	GetUserPhoto         *GetUserPhotoType         `xml:"m:GetUserPhoto,omitempty"`
	RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
}

type GetUserPhotoSoapOut struct {
	GetUserPhotoResponse *GetUserPhotoResponseMessageType `xml:"m:GetUserPhotoResponse,omitempty"`
	ServerVersionInfo    *ServerVersionInfoType           `xml:"t:ServerVersionInfo,omitempty"`
}

type SetUserPhotoSoapIn struct {
	SetUserPhoto         *SetUserPhotoType         `xml:"m:SetUserPhoto,omitempty"`
	RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
}

type SetUserPhotoSoapOut struct {
	SetUserPhotoResponse *SetUserPhotoResponseMessageType `xml:"m:SetUserPhotoResponse,omitempty"`
	ServerVersionInfo    *ServerVersionInfoType           `xml:"t:ServerVersionInfo,omitempty"`
}

type GetMeetingSpaceSoapIn struct {
	GetMeetingSpace      *GetMeetingSpaceType      `xml:"m:GetMeetingSpace,omitempty"`
	RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
	ManagementRole       *ManagementRoleType       `xml:"t:ManagementRole,omitempty"`
}

type GetMeetingSpaceSoapOut struct {
	GetMeetingSpaceResponseMessage *GetMeetingSpaceResponseMessageType `xml:"m:GetMeetingSpaceResponseMessage,omitempty"`
	ServerVersionInfo              *ServerVersionInfoType              `xml:"t:ServerVersionInfo,omitempty"`
}

type DeleteMeetingSpaceSoapIn struct {
	DeleteMeetingSpace   *DeleteMeetingSpaceType   `xml:"m:DeleteMeetingSpace,omitempty"`
	RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
	ManagementRole       *ManagementRoleType       `xml:"t:ManagementRole,omitempty"`
}

type DeleteMeetingSpaceSoapOut struct {
	DeleteMeetingSpaceResponseMessage *DeleteMeetingSpaceResponseMessageType `xml:"m:DeleteMeetingSpaceResponseMessage,omitempty"`
	ServerVersionInfo                 *ServerVersionInfoType                 `xml:"t:ServerVersionInfo,omitempty"`
}

type UpdateMeetingSpaceSoapIn struct {
	UpdateMeetingSpace   *UpdateMeetingSpaceType   `xml:"m:UpdateMeetingSpace,omitempty"`
	RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
	ManagementRole       *ManagementRoleType       `xml:"t:ManagementRole,omitempty"`
}

type UpdateMeetingSpaceSoapOut struct {
	UpdateMeetingSpaceResponseMessage *UpdateMeetingSpaceResponseMessageType `xml:"m:UpdateMeetingSpaceResponseMessage,omitempty"`
	ServerVersionInfo                 *ServerVersionInfoType                 `xml:"t:ServerVersionInfo,omitempty"`
}

type CreateMeetingSpaceSoapIn struct {
	CreateMeetingSpace   *CreateMeetingSpaceType   `xml:"m:CreateMeetingSpace,omitempty"`
	RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
	ManagementRole       *ManagementRoleType       `xml:"t:ManagementRole,omitempty"`
}

type CreateMeetingSpaceSoapOut struct {
	CreateMeetingSpaceResponseMessage *CreateMeetingSpaceResponseMessageType `xml:"m:CreateMeetingSpaceResponseMessage,omitempty"`
	ServerVersionInfo                 *ServerVersionInfoType                 `xml:"t:ServerVersionInfo,omitempty"`
}

type FindMeetingSpaceByJoinUrlSoapIn struct {
	FindMeetingSpaceByJoinUrl *FindMeetingSpaceByJoinUrlType `xml:"m:FindMeetingSpaceByJoinUrl,omitempty"`
	RequestServerVersion      *RequestServerVersionType      `xml:"t:RequestServerVersion,omitempty"`
	ManagementRole            *ManagementRoleType            `xml:"t:ManagementRole,omitempty"`
}

type FindMeetingSpaceByJoinUrlSoapOut struct {
	FindMeetingSpaceByJoinUrlResponseMessage *FindMeetingSpaceByJoinUrlResponseMessageType `xml:"m:FindMeetingSpaceByJoinUrlResponseMessage,omitempty"`
	ServerVersionInfo                        *ServerVersionInfoType                        `xml:"t:ServerVersionInfo,omitempty"`
}

type GetMeetingInstanceSoapIn struct {
	GetMeetingInstanceRequest *GetMeetingInstanceRequestType `xml:"m:GetMeetingInstanceRequest,omitempty"`
	RequestServerVersion      *RequestServerVersionType      `xml:"t:RequestServerVersion,omitempty"`
	ManagementRole            *ManagementRoleType            `xml:"t:ManagementRole,omitempty"`
}

type GetMeetingInstanceSoapOut struct {
	GetMeetingInstanceResponse *GetMeetingInstanceResponseMessageType `xml:"m:GetMeetingInstanceResponse,omitempty"`
	ServerVersionInfo          *ServerVersionInfoType                 `xml:"t:ServerVersionInfo,omitempty"`
}

type DeleteMeetingInstanceSoapIn struct {
	DeleteMeetingInstanceRequest *DeleteMeetingInstanceRequestType `xml:"m:DeleteMeetingInstanceRequest,omitempty"`
	RequestServerVersion         *RequestServerVersionType         `xml:"t:RequestServerVersion,omitempty"`
	ManagementRole               *ManagementRoleType               `xml:"t:ManagementRole,omitempty"`
}

type DeleteMeetingInstanceSoapOut struct {
	DeleteMeetingInstanceResponse *DeleteMeetingInstanceResponseMessageType `xml:"m:DeleteMeetingInstanceResponse,omitempty"`
	ServerVersionInfo             *ServerVersionInfoType                    `xml:"t:ServerVersionInfo,omitempty"`
}

type UpdateMeetingInstanceSoapIn struct {
	UpdateMeetingInstanceRequest *UpdateMeetingInstanceRequestType `xml:"m:UpdateMeetingInstanceRequest,omitempty"`
	RequestServerVersion         *RequestServerVersionType         `xml:"t:RequestServerVersion,omitempty"`
	ManagementRole               *ManagementRoleType               `xml:"t:ManagementRole,omitempty"`
}

type UpdateMeetingInstanceSoapOut struct {
	UpdateMeetingInstanceResponse *UpdateMeetingInstanceResponseMessageType `xml:"m:UpdateMeetingInstanceResponse,omitempty"`
	ServerVersionInfo             *ServerVersionInfoType                    `xml:"t:ServerVersionInfo,omitempty"`
}

type CreateMeetingInstanceSoapIn struct {
	CreateMeetingInstanceRequest *CreateMeetingInstanceRequestType `xml:"m:CreateMeetingInstanceRequest,omitempty"`
	RequestServerVersion         *RequestServerVersionType         `xml:"t:RequestServerVersion,omitempty"`
	ManagementRole               *ManagementRoleType               `xml:"t:ManagementRole,omitempty"`
}

type CreateMeetingInstanceSoapOut struct {
	CreateMeetingInstanceResponse *CreateMeetingInstanceResponseMessageType `xml:"m:CreateMeetingInstanceResponse,omitempty"`
	ServerVersionInfo             *ServerVersionInfoType                    `xml:"t:ServerVersionInfo,omitempty"`
}

type StartSearchSessionSoapIn struct {
	StartSearchSession   *StartSearchSession       `xml:"m:StartSearchSession,omitempty"`
	MailboxCulture       *MailboxCultureType       `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
}

type StartSearchSessionSoapOut struct {
	StartSearchSessionResponse *StartSearchSessionResponseMessage `xml:"m:StartSearchSessionResponse,omitempty"`
	ServerVersionInfo          *ServerVersionInfoType             `xml:"t:ServerVersionInfo,omitempty"`
}

type GetSearchSuggestionsSoapIn struct {
	GetSearchSuggestions *GetSearchSuggestions     `xml:"m:GetSearchSuggestions,omitempty"`
	MailboxCulture       *MailboxCultureType       `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
}

type GetSearchSuggestionsSoapOut struct {
	GetSearchSuggestionsResponse *GetSearchSuggestionsResponseMessage `xml:"m:GetSearchSuggestionsResponse,omitempty"`
	ServerVersionInfo            *ServerVersionInfoType               `xml:"t:ServerVersionInfo,omitempty"`
}

type DeleteSearchSuggestionSoapIn struct {
	DeleteSearchSuggestion *DeleteSearchSuggestion   `xml:"m:DeleteSearchSuggestion,omitempty"`
	MailboxCulture         *MailboxCultureType       `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion   *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
}

type DeleteSearchSuggestionSoapOut struct {
	DeleteSearchSuggestionResponse *DeleteSearchSuggestionResponseMessageType `xml:"m:DeleteSearchSuggestionResponse,omitempty"`
	ServerVersionInfo              *ServerVersionInfoType                     `xml:"t:ServerVersionInfo,omitempty"`
}

type ExecuteSearchSoapIn struct {
	ExecuteSearch        *ExecuteSearch            `xml:"m:ExecuteSearch,omitempty"`
	MailboxCulture       *MailboxCultureType       `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
}

type ExecuteSearchSoapOut struct {
	ExecuteSearchResponse *ExecuteSearchResponseMessage `xml:"m:ExecuteSearchResponse,omitempty"`
	ServerVersionInfo     *ServerVersionInfoType        `xml:"t:ServerVersionInfo,omitempty"`
}

type EndSearchSessionSoapIn struct {
	EndSearchSession     *EndSearchSession         `xml:"m:EndSearchSession,omitempty"`
	MailboxCulture       *MailboxCultureType       `xml:"t:MailboxCulture,omitempty"`
	RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
}

type EndSearchSessionSoapOut struct {
	EndSearchSessionResponse *EndSearchSessionResponseMessage `xml:"m:EndSearchSessionResponse,omitempty"`
	ServerVersionInfo        *ServerVersionInfoType           `xml:"t:ServerVersionInfo,omitempty"`
}

type GetLastPrivateCatalogUpdateSoapIn struct {
	GetLastPrivateCatalogUpdate *GetLastPrivateCatalogUpdateType `xml:"m:GetLastPrivateCatalogUpdate,omitempty"`
	RequestServerVersion        *RequestServerVersionType        `xml:"t:RequestServerVersion,omitempty"`
}

type GetLastPrivateCatalogUpdateSoapOut struct {
	GetLastPrivateCatalogUpdateResponse *GetLastPrivateCatalogUpdateResponseType `xml:"m:GetLastPrivateCatalogUpdateResponse,omitempty"`
	ServerVersionInfo                   *ServerVersionInfoType                   `xml:"t:ServerVersionInfo,omitempty"`
}

type GetPrivateCatalogAddInsSoapIn struct {
	GetPrivateCatalogAddIns *GetPrivateCatalogAddInsType `xml:"m:GetPrivateCatalogAddIns,omitempty"`
	RequestServerVersion    *RequestServerVersionType    `xml:"t:RequestServerVersion,omitempty"`
}

type GetPrivateCatalogAddInsSoapOut struct {
	GetPrivateCatalogAddInsResponse *GetPrivateCatalogAddInsResponseType `xml:"m:GetPrivateCatalogAddInsResponse,omitempty"`
	ServerVersionInfo               *ServerVersionInfoType               `xml:"t:ServerVersionInfo,omitempty"`
}

type ExchangeServicePortType interface {
	ResolveNames(ctx context.Context, in *ResolveNamesSoapIn) (*ResolveNamesSoapOut, error)
	ExpandDL(ctx context.Context, in *ExpandDLSoapIn) (*ExpandDLSoapOut, error)
	GetServerTimeZones(ctx context.Context, in *GetServerTimeZonesSoapIn) (*GetServerTimeZonesSoapOut, error)
	FindFolder(ctx context.Context, in *FindFolderSoapIn) (*FindFolderSoapOut, error)
	FindItem(ctx context.Context, in *FindItemSoapIn) (*FindItemSoapOut, error)
	GetFolder(ctx context.Context, in *GetFolderSoapIn) (*GetFolderSoapOut, error)
	UploadItems(ctx context.Context, in *UploadItemsSoapIn) (*UploadItemsSoapOut, error)
	ExportItems(ctx context.Context, in *ExportItemsSoapIn) (*ExportItemsSoapOut, error)
	ConvertId(ctx context.Context, in *ConvertIdSoapIn) (*ConvertIdSoapOut, error)
	CreateFolder(ctx context.Context, in *CreateFolderSoapIn) (*CreateFolderSoapOut, error)
	CreateFolderPath(ctx context.Context, in *CreateFolderPathSoapIn) (*CreateFolderPathSoapOut, error)
	DeleteFolder(ctx context.Context, in *DeleteFolderSoapIn) (*DeleteFolderSoapOut, error)
	EmptyFolder(ctx context.Context, in *EmptyFolderSoapIn) (*EmptyFolderSoapOut, error)
	UpdateFolder(ctx context.Context, in *UpdateFolderSoapIn) (*UpdateFolderSoapOut, error)
	MoveFolder(ctx context.Context, in *MoveFolderSoapIn) (*MoveFolderSoapOut, error)
	CopyFolder(ctx context.Context, in *CopyFolderSoapIn) (*CopyFolderSoapOut, error)
	Subscribe(ctx context.Context, in *SubscribeSoapIn) (*SubscribeSoapOut, error)
	Unsubscribe(ctx context.Context, in *UnsubscribeSoapIn) (*UnsubscribeSoapOut, error)
	GetEvents(ctx context.Context, in *GetEventsSoapIn) (*GetEventsSoapOut, error)
	GetStreamingEvents(ctx context.Context, in *GetStreamingEventsSoapIn) (*GetStreamingEventsSoapOut, error)
	SyncFolderHierarchy(ctx context.Context, in *SyncFolderHierarchySoapIn) (*SyncFolderHierarchySoapOut, error)
	SyncFolderItems(ctx context.Context, in *SyncFolderItemsSoapIn) (*SyncFolderItemsSoapOut, error)
	CreateManagedFolder(ctx context.Context, in *CreateManagedFolderSoapIn) (*CreateManagedFolderSoapOut, error)
	GetItem(ctx context.Context, in *GetItemSoapIn) (*GetItemSoapOut, error)
	CreateItem(ctx context.Context, in *CreateItemSoapIn) (*CreateItemSoapOut, error)
	DeleteItem(ctx context.Context, in *DeleteItemSoapIn) (*DeleteItemSoapOut, error)
	UpdateItem(ctx context.Context, in *UpdateItemSoapIn) (*UpdateItemSoapOut, error)
	UpdateItemInRecoverableItems(ctx context.Context, in *UpdateItemInRecoverableItemsSoapIn) (*UpdateItemInRecoverableItemsSoapOut, error)
	SendItem(ctx context.Context, in *SendItemSoapIn) (*SendItemSoapOut, error)
	MoveItem(ctx context.Context, in *MoveItemSoapIn) (*MoveItemSoapOut, error)
	CopyItem(ctx context.Context, in *CopyItemSoapIn) (*CopyItemSoapOut, error)
	ArchiveItem(ctx context.Context, in *ArchiveItemSoapIn) (*ArchiveItemSoapOut, error)
	CreateAttachment(ctx context.Context, in *CreateAttachmentSoapIn) (*CreateAttachmentSoapOut, error)
	DeleteAttachment(ctx context.Context, in *DeleteAttachmentSoapIn) (*DeleteAttachmentSoapOut, error)
	GetAttachment(ctx context.Context, in *GetAttachmentSoapIn) (*GetAttachmentSoapOut, error)
	GetClientAccessToken(ctx context.Context, in *GetClientAccessTokenSoapIn) (*GetClientAccessTokenSoapOut, error)
	GetDelegate(ctx context.Context, in *GetDelegateSoapIn) (*GetDelegateSoapOut, error)
	AddDelegate(ctx context.Context, in *AddDelegateSoapIn) (*AddDelegateSoapOut, error)
	RemoveDelegate(ctx context.Context, in *RemoveDelegateSoapIn) (*RemoveDelegateSoapOut, error)
	UpdateDelegate(ctx context.Context, in *UpdateDelegateSoapIn) (*UpdateDelegateSoapOut, error)
	CreateUserConfiguration(ctx context.Context, in *CreateUserConfigurationSoapIn) (*CreateUserConfigurationSoapOut, error)
	DeleteUserConfiguration(ctx context.Context, in *DeleteUserConfigurationSoapIn) (*DeleteUserConfigurationSoapOut, error)
	GetUserConfiguration(ctx context.Context, in *GetUserConfigurationSoapIn) (*GetUserConfigurationSoapOut, error)
	GetSpecificUserConfiguration(ctx context.Context, in *GetSpecificUserConfigurationSoapIn) (*GetSpecificUserConfigurationSoapOut, error)
	UpdateUserConfiguration(ctx context.Context, in *UpdateUserConfigurationSoapIn) (*UpdateUserConfigurationSoapOut, error)
	GetUserAvailability(ctx context.Context, in *GetUserAvailabilitySoapIn) (*GetUserAvailabilitySoapOut, error)
	GetUserOofSettings(ctx context.Context, in *GetUserOofSettingsSoapIn) (*GetUserOofSettingsSoapOut, error)
	SetUserOofSettings(ctx context.Context, in *SetUserOofSettingsSoapIn) (*SetUserOofSettingsSoapOut, error)
	GetServiceConfiguration(ctx context.Context, in *GetServiceConfigurationSoapIn) (*GetServiceConfigurationSoapOut, error)
	GetMailTips(ctx context.Context, in *GetMailTipsSoapIn) (*GetMailTipsSoapOut, error)
	PlayOnPhone(ctx context.Context, in *PlayOnPhoneSoapIn) (*PlayOnPhoneSoapOut, error)
	GetPhoneCallInformation(ctx context.Context, in *GetPhoneCallInformationSoapIn) (*GetPhoneCallInformationSoapOut, error)
	DisconnectPhoneCall(ctx context.Context, in *DisconnectPhoneCallSoapIn) (*DisconnectPhoneCallSoapOut, error)
	GetSharingMetadata(ctx context.Context, in *GetSharingMetadataSoapIn) (*GetSharingMetadataSoapOut, error)
	RefreshSharingFolder(ctx context.Context, in *RefreshSharingFolderSoapIn) (*RefreshSharingFolderSoapOut, error)
	GetSharingFolder(ctx context.Context, in *GetSharingFolderSoapIn) (*GetSharingFolderSoapOut, error)
	SetTeamMailbox(ctx context.Context, in *SetTeamMailboxSoapIn) (*SetTeamMailboxSoapOut, error)
	UnpinTeamMailbox(ctx context.Context, in *UnpinTeamMailboxSoapIn) (*UnpinTeamMailboxSoapOut, error)
	GetRoomLists(ctx context.Context, in *GetRoomListsSoapIn) (*GetRoomListsSoapOut, error)
	GetRooms(ctx context.Context, in *GetRoomsSoapIn) (*GetRoomsSoapOut, error)
	FindMessageTrackingReport(ctx context.Context, in *FindMessageTrackingReportSoapIn) (*FindMessageTrackingReportSoapOut, error)
	GetMessageTrackingReport(ctx context.Context, in *GetMessageTrackingReportSoapIn) (*GetMessageTrackingReportSoapOut, error)
	FindConversation(ctx context.Context, in *FindConversationSoapIn) (*FindConversationSoapOut, error)
	ApplyConversationAction(ctx context.Context, in *ApplyConversationActionSoapIn) (*ApplyConversationActionSoapOut, error)
	GetConversationItems(ctx context.Context, in *GetConversationItemsSoapIn) (*GetConversationItemsSoapOut, error)
	FindPeople(ctx context.Context, in *FindPeopleSoapIn) (*FindPeopleSoapOut, error)
	FindTags(ctx context.Context, in *FindTagsSoapIn) (*FindTagsSoapOut, error)
	AddTag(ctx context.Context, in *AddTagSoapIn) (*AddTagSoapOut, error)
	HideTag(ctx context.Context, in *HideTagSoapIn) (*HideTagSoapOut, error)
	GetPersona(ctx context.Context, in *GetPersonaSoapIn) (*GetPersonaSoapOut, error)
	GetInboxRules(ctx context.Context, in *GetInboxRulesSoapIn) (*GetInboxRulesSoapOut, error)
	UpdateInboxRules(ctx context.Context, in *UpdateInboxRulesSoapIn) (*UpdateInboxRulesSoapOut, error)
	GetPasswordExpirationDate(ctx context.Context, in *GetPasswordExpirationDateSoapIn) (*GetPasswordExpirationDateSoapOut, error)
	GetSearchableMailboxes(ctx context.Context, in *GetSearchableMailboxesSoapIn) (*GetSearchableMailboxesSoapOut, error)
	SearchMailboxes(ctx context.Context, in *SearchMailboxesSoapIn) (*SearchMailboxesSoapOut, error)
	GetDiscoverySearchConfiguration(ctx context.Context, in *GetDiscoverySearchConfigurationSoapIn) (*GetDiscoverySearchConfigurationSoapOut, error)
	GetHoldOnMailboxes(ctx context.Context, in *GetHoldOnMailboxesSoapIn) (*GetHoldOnMailboxesSoapOut, error)
	SetHoldOnMailboxes(ctx context.Context, in *SetHoldOnMailboxesSoapIn) (*SetHoldOnMailboxesSoapOut, error)
	GetNonIndexableItemStatistics(ctx context.Context, in *GetNonIndexableItemStatisticsSoapIn) (*GetNonIndexableItemStatisticsSoapOut, error)
	GetNonIndexableItemDetails(ctx context.Context, in *GetNonIndexableItemDetailsSoapIn) (*GetNonIndexableItemDetailsSoapOut, error)
	MarkAllItemsAsRead(ctx context.Context, in *MarkAllItemsAsReadSoapIn) (*MarkAllItemsAsReadSoapOut, error)
	MarkAsJunk(ctx context.Context, in *MarkAsJunkSoapIn) (*MarkAsJunkSoapOut, error)
	ReportMessage(ctx context.Context, in *ReportMessageSoapIn) (*ReportMessageSoapOut, error)
	GetAppManifests(ctx context.Context, in *GetAppManifestsSoapIn) (*GetAppManifestsSoapOut, error)
	AddNewImContactToGroup(ctx context.Context, in *AddNewImContactToGroupSoapIn) (*AddNewImContactToGroupSoapOut, error)
	AddNewTelUriContactToGroup(ctx context.Context, in *AddNewTelUriContactToGroupSoapIn) (*AddNewTelUriContactToGroupSoapOut, error)
	AddImContactToGroup(ctx context.Context, in *AddImContactToGroupSoapIn) (*AddImContactToGroupSoapOut, error)
	RemoveImContactFromGroup(ctx context.Context, in *RemoveImContactFromGroupSoapIn) (*RemoveImContactFromGroupSoapOut, error)
	AddImGroup(ctx context.Context, in *AddImGroupSoapIn) (*AddImGroupSoapOut, error)
	AddDistributionGroupToImList(ctx context.Context, in *AddDistributionGroupToImListSoapIn) (*AddDistributionGroupToImListSoapOut, error)
	GetImItemList(ctx context.Context, in *GetImItemListSoapIn) (*GetImItemListSoapOut, error)
	GetImItems(ctx context.Context, in *GetImItemsSoapIn) (*GetImItemsSoapOut, error)
	RemoveContactFromImList(ctx context.Context, in *RemoveContactFromImListSoapIn) (*RemoveContactFromImListSoapOut, error)
	RemoveDistributionGroupFromImList(ctx context.Context, in *RemoveDistributionGroupFromImListSoapIn) (*RemoveDistributionGroupFromImListSoapOut, error)
	RemoveImGroup(ctx context.Context, in *RemoveImGroupSoapIn) (*RemoveImGroupSoapOut, error)
	SetImGroup(ctx context.Context, in *SetImGroupSoapIn) (*SetImGroupSoapOut, error)
	SetImListMigrationCompleted(ctx context.Context, in *SetImListMigrationCompletedSoapIn) (*SetImListMigrationCompletedSoapOut, error)
	GetUserRetentionPolicyTags(ctx context.Context, in *GetUserRetentionPolicyTagsSoapIn) (*GetUserRetentionPolicyTagsSoapOut, error)
	InstallApp(ctx context.Context, in *InstallAppSoapIn) (*InstallAppSoapOut, error)
	UpdateExtensionUsage(ctx context.Context, in *UpdateExtensionUsageSoapIn) (*UpdateExtensionUsageSoapOut, error)
	UninstallApp(ctx context.Context, in *UninstallAppSoapIn) (*UninstallAppSoapOut, error)
	DisableApp(ctx context.Context, in *DisableAppSoapIn) (*DisableAppSoapOut, error)
	GetAppMarketplaceUrl(ctx context.Context, in *GetAppMarketplaceUrlSoapIn) (*GetAppMarketplaceUrlSoapOut, error)
	FindAvailableMeetingTimes(ctx context.Context, in *FindAvailableMeetingTimesSoapIn) (*FindAvailableMeetingTimesSoapOut, error)
	FindMeetingTimeCandidates(ctx context.Context, in *FindMeetingTimeCandidatesSoapIn) (*FindMeetingTimeCandidatesSoapOut, error)
	GetUserPhoto(ctx context.Context, in *GetUserPhotoSoapIn) (*GetUserPhotoSoapOut, error)
	SetUserPhoto(ctx context.Context, in *SetUserPhotoSoapIn) (*SetUserPhotoSoapOut, error)
	GetMeetingSpace(ctx context.Context, in *GetMeetingSpaceSoapIn) (*GetMeetingSpaceSoapOut, error)
	DeleteMeetingSpace(ctx context.Context, in *DeleteMeetingSpaceSoapIn) (*DeleteMeetingSpaceSoapOut, error)
	UpdateMeetingSpace(ctx context.Context, in *UpdateMeetingSpaceSoapIn) (*UpdateMeetingSpaceSoapOut, error)
	CreateMeetingSpace(ctx context.Context, in *CreateMeetingSpaceSoapIn) (*CreateMeetingSpaceSoapOut, error)
	FindMeetingSpaceByJoinUrl(ctx context.Context, in *FindMeetingSpaceByJoinUrlSoapIn) (*FindMeetingSpaceByJoinUrlSoapOut, error)
	GetMeetingInstance(ctx context.Context, in *GetMeetingInstanceSoapIn) (*GetMeetingInstanceSoapOut, error)
	DeleteMeetingInstance(ctx context.Context, in *DeleteMeetingInstanceSoapIn) (*DeleteMeetingInstanceSoapOut, error)
	UpdateMeetingInstance(ctx context.Context, in *UpdateMeetingInstanceSoapIn) (*UpdateMeetingInstanceSoapOut, error)
	CreateMeetingInstance(ctx context.Context, in *CreateMeetingInstanceSoapIn) (*CreateMeetingInstanceSoapOut, error)
	StartSearchSession(ctx context.Context, in *StartSearchSessionSoapIn) (*StartSearchSessionSoapOut, error)
	GetSearchSuggestions(ctx context.Context, in *GetSearchSuggestionsSoapIn) (*GetSearchSuggestionsSoapOut, error)
	DeleteSearchSuggestion(ctx context.Context, in *DeleteSearchSuggestionSoapIn) (*DeleteSearchSuggestionSoapOut, error)
	ExecuteSearch(ctx context.Context, in *ExecuteSearchSoapIn) (*ExecuteSearchSoapOut, error)
	EndSearchSession(ctx context.Context, in *EndSearchSessionSoapIn) (*EndSearchSessionSoapOut, error)
	GetLastPrivateCatalogUpdate(ctx context.Context, in *GetLastPrivateCatalogUpdateSoapIn) (*GetLastPrivateCatalogUpdateSoapOut, error)
	GetPrivateCatalogAddIns(ctx context.Context, in *GetPrivateCatalogAddInsSoapIn) (*GetPrivateCatalogAddInsSoapOut, error)
}

type ExchangeServiceBinding struct {
	client SOAPClient
}

func (b *ExchangeServiceBinding) ResolveNames(ctx context.Context, input *ResolveNamesSoapIn, detail any) (*ResolveNamesSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		ResolveNames *ResolveNamesType `xml:"m:ResolveNames,omitempty"`
	}{
		ResolveNames: input.ResolveNames,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		ResolveNamesResponse *ResolveNamesResponseType `xml:"m:ResolveNamesResponse,omitempty"`
		Fault                *Fault                    `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/ResolveNames", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &ResolveNamesSoapOut{
		ServerVersionInfo:    outputHeader.ServerVersionInfo,
		ResolveNamesResponse: outputBody.ResolveNamesResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) ExpandDL(ctx context.Context, input *ExpandDLSoapIn, detail any) (*ExpandDLSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		ExpandDL *ExpandDLType `xml:"m:ExpandDL,omitempty"`
	}{
		ExpandDL: input.ExpandDL,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		ExpandDLResponse *ExpandDLResponseType `xml:"m:ExpandDLResponse,omitempty"`
		Fault            *Fault                `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/ExpandDL", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &ExpandDLSoapOut{
		ServerVersionInfo: outputHeader.ServerVersionInfo,
		ExpandDLResponse:  outputBody.ExpandDLResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetServerTimeZones(ctx context.Context, input *GetServerTimeZonesSoapIn, detail any) (*GetServerTimeZonesSoapOut, error) {
	inputHeader := &struct {
		MailboxCulture       *MailboxCultureType       `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
	}{
		MailboxCulture:       input.MailboxCulture,
		RequestServerVersion: input.RequestServerVersion,
	}
	inputBody := &struct {
		GetServerTimeZones *GetServerTimeZonesType `xml:"m:GetServerTimeZones,omitempty"`
	}{
		GetServerTimeZones: input.GetServerTimeZones,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetServerTimeZonesResponse *GetServerTimeZonesResponseType `xml:"m:GetServerTimeZonesResponse,omitempty"`
		Fault                      *Fault                          `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetServerTimeZones", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetServerTimeZonesSoapOut{
		ServerVersionInfo:          outputHeader.ServerVersionInfo,
		GetServerTimeZonesResponse: outputBody.GetServerTimeZonesResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) FindFolder(ctx context.Context, input *FindFolderSoapIn, detail any) (*FindFolderSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
		TimeZoneContext       *TimeZoneContextType       `xml:"t:TimeZoneContext,omitempty"`
		ManagementRole        *ManagementRoleType        `xml:"t:ManagementRole,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
		TimeZoneContext:       input.TimeZoneContext,
		ManagementRole:        input.ManagementRole,
	}
	inputBody := &struct {
		FindFolder *FindFolderType `xml:"m:FindFolder,omitempty"`
	}{
		FindFolder: input.FindFolder,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		FindFolderResponse *FindFolderResponseType `xml:"m:FindFolderResponse,omitempty"`
		Fault              *Fault                  `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/FindFolder", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &FindFolderSoapOut{
		ServerVersionInfo:  outputHeader.ServerVersionInfo,
		FindFolderResponse: outputBody.FindFolderResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) FindItem(ctx context.Context, input *FindItemSoapIn, detail any) (*FindItemSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
		TimeZoneContext       *TimeZoneContextType       `xml:"t:TimeZoneContext,omitempty"`
		DateTimePrecision     DateTimePrecisionType      `xml:"t:DateTimePrecision,omitempty"`
		ManagementRole        *ManagementRoleType        `xml:"t:ManagementRole,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
		TimeZoneContext:       input.TimeZoneContext,
		DateTimePrecision:     input.DateTimePrecision,
		ManagementRole:        input.ManagementRole,
	}
	inputBody := &struct {
		FindItem *FindItemType `xml:"m:FindItem,omitempty"`
	}{
		FindItem: input.FindItem,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		FindItemResponse *FindItemResponseType `xml:"m:FindItemResponse,omitempty"`
		Fault            *Fault                `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/FindItem", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &FindItemSoapOut{
		ServerVersionInfo: outputHeader.ServerVersionInfo,
		FindItemResponse:  outputBody.FindItemResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetFolder(ctx context.Context, input *GetFolderSoapIn, detail any) (*GetFolderSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
		TimeZoneContext       *TimeZoneContextType       `xml:"t:TimeZoneContext,omitempty"`
		ManagementRole        *ManagementRoleType        `xml:"t:ManagementRole,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
		TimeZoneContext:       input.TimeZoneContext,
		ManagementRole:        input.ManagementRole,
	}
	inputBody := &struct {
		GetFolder *GetFolderType `xml:"m:GetFolder,omitempty"`
	}{
		GetFolder: input.GetFolder,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetFolderResponse *GetFolderResponseType `xml:"m:GetFolderResponse,omitempty"`
		Fault             *Fault                 `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetFolder", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetFolderSoapOut{
		ServerVersionInfo: outputHeader.ServerVersionInfo,
		GetFolderResponse: outputBody.GetFolderResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) ConvertId(ctx context.Context, input *ConvertIdSoapIn, detail any) (*ConvertIdSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		ConvertId *ConvertIdType `xml:"m:ConvertId,omitempty"`
	}{
		ConvertId: input.ConvertId,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		ConvertIdResponse *ConvertIdResponseType `xml:"m:ConvertIdResponse,omitempty"`
		Fault             *Fault                 `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/ConvertId", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &ConvertIdSoapOut{
		ServerVersionInfo: outputHeader.ServerVersionInfo,
		ConvertIdResponse: outputBody.ConvertIdResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) UploadItems(ctx context.Context, input *UploadItemsSoapIn, detail any) (*UploadItemsSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		UploadItems *UploadItemsType `xml:"m:UploadItems,omitempty"`
	}{
		UploadItems: input.UploadItems,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		UploadItemsResponse *UploadItemsResponseType `xml:"m:UploadItemsResponse,omitempty"`
		Fault               *Fault                   `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/UploadItems", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &UploadItemsSoapOut{
		ServerVersionInfo:   outputHeader.ServerVersionInfo,
		UploadItemsResponse: outputBody.UploadItemsResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) ExportItems(ctx context.Context, input *ExportItemsSoapIn, detail any) (*ExportItemsSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
		ManagementRole        *ManagementRoleType        `xml:"t:ManagementRole,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
		ManagementRole:        input.ManagementRole,
	}
	inputBody := &struct {
		ExportItems *ExportItemsType `xml:"m:ExportItems,omitempty"`
	}{
		ExportItems: input.ExportItems,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		ExportItemsResponse *ExportItemsResponseType `xml:"m:ExportItemsResponse,omitempty"`
		Fault               *Fault                   `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/ExportItems", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &ExportItemsSoapOut{
		ServerVersionInfo:   outputHeader.ServerVersionInfo,
		ExportItemsResponse: outputBody.ExportItemsResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) CreateFolderPath(ctx context.Context, input *CreateFolderPathSoapIn, detail any) (*CreateFolderPathSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
		TimeZoneContext       *TimeZoneContextType       `xml:"t:TimeZoneContext,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
		TimeZoneContext:       input.TimeZoneContext,
	}
	inputBody := &struct {
		CreateFolderPath *CreateFolderPathType `xml:"m:CreateFolderPath,omitempty"`
	}{
		CreateFolderPath: input.CreateFolderPath,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		CreateFolderPathResponse *CreateFolderPathResponseType `xml:"m:CreateFolderPathResponse,omitempty"`
		Fault                    *Fault                        `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/CreateFolderPath", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &CreateFolderPathSoapOut{
		ServerVersionInfo:        outputHeader.ServerVersionInfo,
		CreateFolderPathResponse: outputBody.CreateFolderPathResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) CreateFolder(ctx context.Context, input *CreateFolderSoapIn, detail any) (*CreateFolderSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
		TimeZoneContext       *TimeZoneContextType       `xml:"t:TimeZoneContext,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
		TimeZoneContext:       input.TimeZoneContext,
	}
	inputBody := &struct {
		CreateFolder *CreateFolderType `xml:"m:CreateFolder,omitempty"`
	}{
		CreateFolder: input.CreateFolder,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		CreateFolderResponse *CreateFolderResponseType `xml:"m:CreateFolderResponse,omitempty"`
		Fault                *Fault                    `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/CreateFolder", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &CreateFolderSoapOut{
		ServerVersionInfo:    outputHeader.ServerVersionInfo,
		CreateFolderResponse: outputBody.CreateFolderResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) DeleteFolder(ctx context.Context, input *DeleteFolderSoapIn, detail any) (*DeleteFolderSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		DeleteFolder *DeleteFolderType `xml:"m:DeleteFolder,omitempty"`
	}{
		DeleteFolder: input.DeleteFolder,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		DeleteFolderResponse *DeleteFolderResponseType `xml:"m:DeleteFolderResponse,omitempty"`
		Fault                *Fault                    `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/DeleteFolder", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &DeleteFolderSoapOut{
		ServerVersionInfo:    outputHeader.ServerVersionInfo,
		DeleteFolderResponse: outputBody.DeleteFolderResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) EmptyFolder(ctx context.Context, input *EmptyFolderSoapIn, detail any) (*EmptyFolderSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		EmptyFolder *EmptyFolderType `xml:"m:EmptyFolder,omitempty"`
	}{
		EmptyFolder: input.EmptyFolder,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		EmptyFolderResponse *EmptyFolderResponseType `xml:"m:EmptyFolderResponse,omitempty"`
		Fault               *Fault                   `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/EmptyFolder", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &EmptyFolderSoapOut{
		ServerVersionInfo:   outputHeader.ServerVersionInfo,
		EmptyFolderResponse: outputBody.EmptyFolderResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) UpdateFolder(ctx context.Context, input *UpdateFolderSoapIn, detail any) (*UpdateFolderSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
		TimeZoneContext       *TimeZoneContextType       `xml:"t:TimeZoneContext,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
		TimeZoneContext:       input.TimeZoneContext,
	}
	inputBody := &struct {
		UpdateFolder *UpdateFolderType `xml:"m:UpdateFolder,omitempty"`
	}{
		UpdateFolder: input.UpdateFolder,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		UpdateFolderResponse *UpdateFolderResponseType `xml:"m:UpdateFolderResponse,omitempty"`
		Fault                *Fault                    `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/UpdateFolder", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &UpdateFolderSoapOut{
		ServerVersionInfo:    outputHeader.ServerVersionInfo,
		UpdateFolderResponse: outputBody.UpdateFolderResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) MoveFolder(ctx context.Context, input *MoveFolderSoapIn, detail any) (*MoveFolderSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		MoveFolder *MoveFolderType `xml:"m:MoveFolder,omitempty"`
	}{
		MoveFolder: input.MoveFolder,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		MoveFolderResponse *MoveFolderResponseType `xml:"m:MoveFolderResponse,omitempty"`
		Fault              *Fault                  `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/MoveFolder", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &MoveFolderSoapOut{
		ServerVersionInfo:  outputHeader.ServerVersionInfo,
		MoveFolderResponse: outputBody.MoveFolderResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) CopyFolder(ctx context.Context, input *CopyFolderSoapIn, detail any) (*CopyFolderSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		CopyFolder *CopyFolderType `xml:"m:CopyFolder,omitempty"`
	}{
		CopyFolder: input.CopyFolder,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		CopyFolderResponse *CopyFolderResponseType `xml:"m:CopyFolderResponse,omitempty"`
		Fault              *Fault                  `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/CopyFolder", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &CopyFolderSoapOut{
		ServerVersionInfo:  outputHeader.ServerVersionInfo,
		CopyFolderResponse: outputBody.CopyFolderResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) Subscribe(ctx context.Context, input *SubscribeSoapIn, detail any) (*SubscribeSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		Subscribe *SubscribeType `xml:"m:Subscribe,omitempty"`
	}{
		Subscribe: input.Subscribe,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		SubscribeResponse *SubscribeResponseType `xml:"m:SubscribeResponse,omitempty"`
		Fault             *Fault                 `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/Subscribe", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &SubscribeSoapOut{
		ServerVersionInfo: outputHeader.ServerVersionInfo,
		SubscribeResponse: outputBody.SubscribeResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) Unsubscribe(ctx context.Context, input *UnsubscribeSoapIn, detail any) (*UnsubscribeSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		Unsubscribe *UnsubscribeType `xml:"m:Unsubscribe,omitempty"`
	}{
		Unsubscribe: input.Unsubscribe,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		UnsubscribeResponse *UnsubscribeResponseType `xml:"m:UnsubscribeResponse,omitempty"`
		Fault               *Fault                   `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/Unsubscribe", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &UnsubscribeSoapOut{
		ServerVersionInfo:   outputHeader.ServerVersionInfo,
		UnsubscribeResponse: outputBody.UnsubscribeResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetEvents(ctx context.Context, input *GetEventsSoapIn, detail any) (*GetEventsSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		GetEvents *GetEventsType `xml:"m:GetEvents,omitempty"`
	}{
		GetEvents: input.GetEvents,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetEventsResponse *GetEventsResponseType `xml:"m:GetEventsResponse,omitempty"`
		Fault             *Fault                 `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetEvents", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetEventsSoapOut{
		ServerVersionInfo: outputHeader.ServerVersionInfo,
		GetEventsResponse: outputBody.GetEventsResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetStreamingEvents(ctx context.Context, input *GetStreamingEventsSoapIn, detail any) (*GetStreamingEventsSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		GetStreamingEvents *GetStreamingEventsType `xml:"m:GetStreamingEvents,omitempty"`
	}{
		GetStreamingEvents: input.GetStreamingEvents,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetStreamingEventsResponse *GetStreamingEventsResponseType `xml:"m:GetStreamingEventsResponse,omitempty"`
		Fault                      *Fault                          `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetEvents", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetStreamingEventsSoapOut{
		ServerVersionInfo:          outputHeader.ServerVersionInfo,
		GetStreamingEventsResponse: outputBody.GetStreamingEventsResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) SyncFolderHierarchy(ctx context.Context, input *SyncFolderHierarchySoapIn, detail any) (*SyncFolderHierarchySoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		SyncFolderHierarchy *SyncFolderHierarchyType `xml:"m:SyncFolderHierarchy,omitempty"`
	}{
		SyncFolderHierarchy: input.SyncFolderHierarchy,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		SyncFolderHierarchyResponse *SyncFolderHierarchyResponseType `xml:"m:SyncFolderHierarchyResponse,omitempty"`
		Fault                       *Fault                           `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/SyncFolderHierarchy", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &SyncFolderHierarchySoapOut{
		ServerVersionInfo:           outputHeader.ServerVersionInfo,
		SyncFolderHierarchyResponse: outputBody.SyncFolderHierarchyResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) SyncFolderItems(ctx context.Context, input *SyncFolderItemsSoapIn, detail any) (*SyncFolderItemsSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		SyncFolderItems *SyncFolderItemsType `xml:"m:SyncFolderItems,omitempty"`
	}{
		SyncFolderItems: input.SyncFolderItems,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		SyncFolderItemsResponse *SyncFolderItemsResponseType `xml:"m:SyncFolderItemsResponse,omitempty"`
		Fault                   *Fault                       `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/SyncFolderItems", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &SyncFolderItemsSoapOut{
		ServerVersionInfo:       outputHeader.ServerVersionInfo,
		SyncFolderItemsResponse: outputBody.SyncFolderItemsResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetItem(ctx context.Context, input *GetItemSoapIn, detail any) (*GetItemSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
		TimeZoneContext       *TimeZoneContextType       `xml:"t:TimeZoneContext,omitempty"`
		DateTimePrecision     DateTimePrecisionType      `xml:"t:DateTimePrecision,omitempty"`
		ManagementRole        *ManagementRoleType        `xml:"t:ManagementRole,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
		TimeZoneContext:       input.TimeZoneContext,
		DateTimePrecision:     input.DateTimePrecision,
		ManagementRole:        input.ManagementRole,
	}
	inputBody := &struct {
		GetItem *GetItemType `xml:"m:GetItem,omitempty"`
	}{
		GetItem: input.GetItem,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetItemResponse *GetItemResponseType `xml:"m:GetItemResponse,omitempty"`
		Fault           *Fault               `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetItem", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetItemSoapOut{
		ServerVersionInfo: outputHeader.ServerVersionInfo,
		GetItemResponse:   outputBody.GetItemResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) CreateItem(ctx context.Context, input *CreateItemSoapIn, detail any) (*CreateItemSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
		TimeZoneContext       *TimeZoneContextType       `xml:"t:TimeZoneContext,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
		TimeZoneContext:       input.TimeZoneContext,
	}
	inputBody := &struct {
		CreateItem *CreateItemType `xml:"m:CreateItem,omitempty"`
	}{
		CreateItem: input.CreateItem,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		CreateItemResponse *CreateItemResponseType `xml:"m:CreateItemResponse,omitempty"`
		Fault              *Fault                  `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/CreateItem", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &CreateItemSoapOut{
		ServerVersionInfo:  outputHeader.ServerVersionInfo,
		CreateItemResponse: outputBody.CreateItemResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) DeleteItem(ctx context.Context, input *DeleteItemSoapIn, detail any) (*DeleteItemSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		DeleteItem *DeleteItemType `xml:"m:DeleteItem,omitempty"`
	}{
		DeleteItem: input.DeleteItem,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		DeleteItemResponse *DeleteItemResponseType `xml:"m:DeleteItemResponse,omitempty"`
		Fault              *Fault                  `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/DeleteItem", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &DeleteItemSoapOut{
		ServerVersionInfo:  outputHeader.ServerVersionInfo,
		DeleteItemResponse: outputBody.DeleteItemResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) UpdateItem(ctx context.Context, input *UpdateItemSoapIn, detail any) (*UpdateItemSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
		TimeZoneContext       *TimeZoneContextType       `xml:"t:TimeZoneContext,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
		TimeZoneContext:       input.TimeZoneContext,
	}
	inputBody := &struct {
		UpdateItem *UpdateItemType `xml:"m:UpdateItem,omitempty"`
	}{
		UpdateItem: input.UpdateItem,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		UpdateItemResponse *UpdateItemResponseType `xml:"m:UpdateItemResponse,omitempty"`
		Fault              *Fault                  `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/UpdateItem", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &UpdateItemSoapOut{
		ServerVersionInfo:  outputHeader.ServerVersionInfo,
		UpdateItemResponse: outputBody.UpdateItemResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) UpdateItemInRecoverableItems(ctx context.Context, input *UpdateItemInRecoverableItemsSoapIn, detail any) (*UpdateItemInRecoverableItemsSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
		TimeZoneContext       *TimeZoneContextType       `xml:"t:TimeZoneContext,omitempty"`
		ManagementRole        *ManagementRoleType        `xml:"t:ManagementRole,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
		TimeZoneContext:       input.TimeZoneContext,
		ManagementRole:        input.ManagementRole,
	}
	inputBody := &struct {
		UpdateItemInRecoverableItems *UpdateItemInRecoverableItemsType `xml:"m:UpdateItemInRecoverableItems,omitempty"`
	}{
		UpdateItemInRecoverableItems: input.UpdateItemInRecoverableItems,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		UpdateItemInRecoverableItemsResponse *UpdateItemInRecoverableItemsResponseType `xml:"m:UpdateItemInRecoverableItemsResponse,omitempty"`
		Fault                                *Fault                                    `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/UpdateItemInRecoverableItems", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &UpdateItemInRecoverableItemsSoapOut{
		ServerVersionInfo:                    outputHeader.ServerVersionInfo,
		UpdateItemInRecoverableItemsResponse: outputBody.UpdateItemInRecoverableItemsResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) SendItem(ctx context.Context, input *SendItemSoapIn, detail any) (*SendItemSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		SendItem *SendItemType `xml:"m:SendItem,omitempty"`
	}{
		SendItem: input.SendItem,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		SendItemResponse *SendItemResponseType `xml:"m:SendItemResponse,omitempty"`
		Fault            *Fault                `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/SendItem", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &SendItemSoapOut{
		ServerVersionInfo: outputHeader.ServerVersionInfo,
		SendItemResponse:  outputBody.SendItemResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) MoveItem(ctx context.Context, input *MoveItemSoapIn, detail any) (*MoveItemSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		MoveItem *MoveItemType `xml:"m:MoveItem,omitempty"`
	}{
		MoveItem: input.MoveItem,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		MoveItemResponse *MoveItemResponseType `xml:"m:MoveItemResponse,omitempty"`
		Fault            *Fault                `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/MoveItem", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &MoveItemSoapOut{
		ServerVersionInfo: outputHeader.ServerVersionInfo,
		MoveItemResponse:  outputBody.MoveItemResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) CopyItem(ctx context.Context, input *CopyItemSoapIn, detail any) (*CopyItemSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		CopyItem *CopyItemType `xml:"m:CopyItem,omitempty"`
	}{
		CopyItem: input.CopyItem,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		CopyItemResponse *CopyItemResponseType `xml:"m:CopyItemResponse,omitempty"`
		Fault            *Fault                `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/CopyItem", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &CopyItemSoapOut{
		ServerVersionInfo: outputHeader.ServerVersionInfo,
		CopyItemResponse:  outputBody.CopyItemResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) ArchiveItem(ctx context.Context, input *ArchiveItemSoapIn, detail any) (*ArchiveItemSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		ArchiveItem *ArchiveItemType `xml:"m:ArchiveItem,omitempty"`
	}{
		ArchiveItem: input.ArchiveItem,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		ArchiveItemResponse *ArchiveItemResponseType `xml:"m:ArchiveItemResponse,omitempty"`
		Fault               *Fault                   `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/ArchiveItem", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &ArchiveItemSoapOut{
		ServerVersionInfo:   outputHeader.ServerVersionInfo,
		ArchiveItemResponse: outputBody.ArchiveItemResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) CreateAttachment(ctx context.Context, input *CreateAttachmentSoapIn, detail any) (*CreateAttachmentSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
		TimeZoneContext       *TimeZoneContextType       `xml:"t:TimeZoneContext,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
		TimeZoneContext:       input.TimeZoneContext,
	}
	inputBody := &struct {
		CreateAttachment *CreateAttachmentType `xml:"m:CreateAttachment,omitempty"`
	}{
		CreateAttachment: input.CreateAttachment,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		CreateAttachmentResponse *CreateAttachmentResponseType `xml:"m:CreateAttachmentResponse,omitempty"`
		Fault                    *Fault                        `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/CreateAttachment", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &CreateAttachmentSoapOut{
		ServerVersionInfo:        outputHeader.ServerVersionInfo,
		CreateAttachmentResponse: outputBody.CreateAttachmentResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) DeleteAttachment(ctx context.Context, input *DeleteAttachmentSoapIn, detail any) (*DeleteAttachmentSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		DeleteAttachment *DeleteAttachmentType `xml:"m:DeleteAttachment,omitempty"`
	}{
		DeleteAttachment: input.DeleteAttachment,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		DeleteAttachmentResponse *DeleteAttachmentResponseType `xml:"m:DeleteAttachmentResponse,omitempty"`
		Fault                    *Fault                        `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/DeleteAttachment", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &DeleteAttachmentSoapOut{
		ServerVersionInfo:        outputHeader.ServerVersionInfo,
		DeleteAttachmentResponse: outputBody.DeleteAttachmentResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetAttachment(ctx context.Context, input *GetAttachmentSoapIn, detail any) (*GetAttachmentSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
		TimeZoneContext       *TimeZoneContextType       `xml:"t:TimeZoneContext,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
		TimeZoneContext:       input.TimeZoneContext,
	}
	inputBody := &struct {
		GetAttachment *GetAttachmentType `xml:"m:GetAttachment,omitempty"`
	}{
		GetAttachment: input.GetAttachment,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetAttachmentResponse *GetAttachmentResponseType `xml:"m:GetAttachmentResponse,omitempty"`
		Fault                 *Fault                     `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetAttachment", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetAttachmentSoapOut{
		ServerVersionInfo:     outputHeader.ServerVersionInfo,
		GetAttachmentResponse: outputBody.GetAttachmentResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetClientAccessToken(ctx context.Context, input *GetClientAccessTokenSoapIn, detail any) (*GetClientAccessTokenSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
	}
	inputBody := &struct {
		GetClientAccessToken *GetClientAccessTokenType `xml:"m:GetClientAccessToken,omitempty"`
	}{
		GetClientAccessToken: input.GetClientAccessToken,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetClientAccessTokenResponse *GetClientAccessTokenResponseType `xml:"m:GetClientAccessTokenResponse,omitempty"`
		Fault                        *Fault                            `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetClientAccessToken", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetClientAccessTokenSoapOut{
		ServerVersionInfo:            outputHeader.ServerVersionInfo,
		GetClientAccessTokenResponse: outputBody.GetClientAccessTokenResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) CreateManagedFolder(ctx context.Context, input *CreateManagedFolderSoapIn, detail any) (*CreateManagedFolderSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		CreateManagedFolder *CreateManagedFolderRequestType `xml:"m:CreateManagedFolder,omitempty"`
	}{
		CreateManagedFolder: input.CreateManagedFolder,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		CreateManagedFolderResponse *CreateManagedFolderResponseType `xml:"m:CreateManagedFolderResponse,omitempty"`
		Fault                       *Fault                           `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/CreateManagedFolder", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &CreateManagedFolderSoapOut{
		ServerVersionInfo:           outputHeader.ServerVersionInfo,
		CreateManagedFolderResponse: outputBody.CreateManagedFolderResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetDelegate(ctx context.Context, input *GetDelegateSoapIn, detail any) (*GetDelegateSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		GetDelegate *GetDelegateType `xml:"m:GetDelegate,omitempty"`
	}{
		GetDelegate: input.GetDelegate,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetDelegateResponse *GetDelegateResponseMessageType `xml:"m:GetDelegateResponse,omitempty"`
		Fault               *Fault                          `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetDelegate", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetDelegateSoapOut{
		ServerVersionInfo:   outputHeader.ServerVersionInfo,
		GetDelegateResponse: outputBody.GetDelegateResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) AddDelegate(ctx context.Context, input *AddDelegateSoapIn, detail any) (*AddDelegateSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		AddDelegate *AddDelegateType `xml:"m:AddDelegate,omitempty"`
	}{
		AddDelegate: input.AddDelegate,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		AddDelegateResponse *AddDelegateResponseMessageType `xml:"m:AddDelegateResponse,omitempty"`
		Fault               *Fault                          `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/AddDelegate", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &AddDelegateSoapOut{
		ServerVersionInfo:   outputHeader.ServerVersionInfo,
		AddDelegateResponse: outputBody.AddDelegateResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) RemoveDelegate(ctx context.Context, input *RemoveDelegateSoapIn, detail any) (*RemoveDelegateSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		RemoveDelegate *RemoveDelegateType `xml:"m:RemoveDelegate,omitempty"`
	}{
		RemoveDelegate: input.RemoveDelegate,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		RemoveDelegateResponse *RemoveDelegateResponseMessageType `xml:"m:RemoveDelegateResponse,omitempty"`
		Fault                  *Fault                             `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/RemoveDelegate", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &RemoveDelegateSoapOut{
		ServerVersionInfo:      outputHeader.ServerVersionInfo,
		RemoveDelegateResponse: outputBody.RemoveDelegateResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) UpdateDelegate(ctx context.Context, input *UpdateDelegateSoapIn, detail any) (*UpdateDelegateSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		UpdateDelegate *UpdateDelegateType `xml:"m:UpdateDelegate,omitempty"`
	}{
		UpdateDelegate: input.UpdateDelegate,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		UpdateDelegateResponse *UpdateDelegateResponseMessageType `xml:"m:UpdateDelegateResponse,omitempty"`
		Fault                  *Fault                             `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/UpdateDelegate", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &UpdateDelegateSoapOut{
		ServerVersionInfo:      outputHeader.ServerVersionInfo,
		UpdateDelegateResponse: outputBody.UpdateDelegateResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) CreateUserConfiguration(ctx context.Context, input *CreateUserConfigurationSoapIn, detail any) (*CreateUserConfigurationSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		CreateUserConfiguration *CreateUserConfigurationType `xml:"m:CreateUserConfiguration,omitempty"`
	}{
		CreateUserConfiguration: input.CreateUserConfiguration,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		CreateUserConfigurationResponse *CreateUserConfigurationResponseType `xml:"m:CreateUserConfigurationResponse,omitempty"`
		Fault                           *Fault                               `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/CreateUserConfiguration", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &CreateUserConfigurationSoapOut{
		ServerVersionInfo:               outputHeader.ServerVersionInfo,
		CreateUserConfigurationResponse: outputBody.CreateUserConfigurationResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) DeleteUserConfiguration(ctx context.Context, input *DeleteUserConfigurationSoapIn, detail any) (*DeleteUserConfigurationSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		DeleteUserConfiguration *DeleteUserConfigurationType `xml:"m:DeleteUserConfiguration,omitempty"`
	}{
		DeleteUserConfiguration: input.DeleteUserConfiguration,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		DeleteUserConfigurationResponse *DeleteUserConfigurationResponseType `xml:"m:DeleteUserConfigurationResponse,omitempty"`
		Fault                           *Fault                               `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/DeleteUserConfiguration", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &DeleteUserConfigurationSoapOut{
		ServerVersionInfo:               outputHeader.ServerVersionInfo,
		DeleteUserConfigurationResponse: outputBody.DeleteUserConfigurationResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetUserConfiguration(ctx context.Context, input *GetUserConfigurationSoapIn, detail any) (*GetUserConfigurationSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		GetUserConfiguration *GetUserConfigurationType `xml:"m:GetUserConfiguration,omitempty"`
	}{
		GetUserConfiguration: input.GetUserConfiguration,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetUserConfigurationResponse *GetUserConfigurationResponseType `xml:"m:GetUserConfigurationResponse,omitempty"`
		Fault                        *Fault                            `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetUserConfiguration", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetUserConfigurationSoapOut{
		ServerVersionInfo:            outputHeader.ServerVersionInfo,
		GetUserConfigurationResponse: outputBody.GetUserConfigurationResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetSpecificUserConfiguration(ctx context.Context, input *GetSpecificUserConfigurationSoapIn, detail any) (*GetSpecificUserConfigurationSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		GetSpecificUserConfiguration *GetSpecificUserConfigurationType `xml:"m:GetSpecificUserConfiguration,omitempty"`
	}{
		GetSpecificUserConfiguration: input.GetSpecificUserConfiguration,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetSpecificUserConfigurationResponse *GetSpecificUserConfigurationResponseType `xml:"m:GetSpecificUserConfigurationResponse,omitempty"`
		Fault                                *Fault                                    `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetSpecificUserConfiguration", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetSpecificUserConfigurationSoapOut{
		ServerVersionInfo:                    outputHeader.ServerVersionInfo,
		GetSpecificUserConfigurationResponse: outputBody.GetSpecificUserConfigurationResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) UpdateUserConfiguration(ctx context.Context, input *UpdateUserConfigurationSoapIn, detail any) (*UpdateUserConfigurationSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		UpdateUserConfiguration *UpdateUserConfigurationType `xml:"m:UpdateUserConfiguration,omitempty"`
	}{
		UpdateUserConfiguration: input.UpdateUserConfiguration,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		UpdateUserConfigurationResponse *UpdateUserConfigurationResponseType `xml:"m:UpdateUserConfigurationResponse,omitempty"`
		Fault                           *Fault                               `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/UpdateUserConfiguration", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &UpdateUserConfigurationSoapOut{
		ServerVersionInfo:               outputHeader.ServerVersionInfo,
		UpdateUserConfigurationResponse: outputBody.UpdateUserConfigurationResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetUserAvailability(ctx context.Context, input *GetUserAvailabilitySoapIn, detail any) (*GetUserAvailabilitySoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		TimeZoneContext       *TimeZoneContextType       `xml:"t:TimeZoneContext,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		TimeZoneContext:       input.TimeZoneContext,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		GetUserAvailabilityRequest *GetUserAvailabilityRequestType `xml:"m:GetUserAvailabilityRequest,omitempty"`
	}{
		GetUserAvailabilityRequest: input.GetUserAvailabilityRequest,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetUserAvailabilityResponse *GetUserAvailabilityResponseType `xml:"m:GetUserAvailabilityResponse,omitempty"`
		Fault                       *Fault                           `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetUserAvailability", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetUserAvailabilitySoapOut{
		ServerVersionInfo:           outputHeader.ServerVersionInfo,
		GetUserAvailabilityResponse: outputBody.GetUserAvailabilityResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetUserOofSettings(ctx context.Context, input *GetUserOofSettingsSoapIn, detail any) (*GetUserOofSettingsSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
	}
	inputBody := &struct {
		GetUserOofSettingsRequest *GetUserOofSettingsRequest `xml:"m:GetUserOofSettingsRequest,omitempty"`
	}{
		GetUserOofSettingsRequest: input.GetUserOofSettingsRequest,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetUserOofSettingsResponse *GetUserOofSettingsResponse `xml:"m:GetUserOofSettingsResponse,omitempty"`
		Fault                      *Fault                      `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetUserOofSettings", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetUserOofSettingsSoapOut{
		ServerVersionInfo:          outputHeader.ServerVersionInfo,
		GetUserOofSettingsResponse: outputBody.GetUserOofSettingsResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) SetUserOofSettings(ctx context.Context, input *SetUserOofSettingsSoapIn, detail any) (*SetUserOofSettingsSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
	}
	inputBody := &struct {
		SetUserOofSettingsRequest *SetUserOofSettingsRequest `xml:"m:SetUserOofSettingsRequest,omitempty"`
	}{
		SetUserOofSettingsRequest: input.SetUserOofSettingsRequest,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		SetUserOofSettingsResponse *SetUserOofSettingsResponse `xml:"m:SetUserOofSettingsResponse,omitempty"`
		Fault                      *Fault                      `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/SetUserOofSettings", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &SetUserOofSettingsSoapOut{
		ServerVersionInfo:          outputHeader.ServerVersionInfo,
		SetUserOofSettingsResponse: outputBody.SetUserOofSettingsResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetServiceConfiguration(ctx context.Context, input *GetServiceConfigurationSoapIn, detail any) (*GetServiceConfigurationSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		RequestServerVersion:  input.RequestServerVersion,
		MailboxCulture:        input.MailboxCulture,
	}
	inputBody := &struct {
		GetServiceConfiguration *GetServiceConfigurationType `xml:"m:GetServiceConfiguration,omitempty"`
	}{
		GetServiceConfiguration: input.GetServiceConfiguration,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetServiceConfigurationResponse *GetServiceConfigurationResponseMessageType `xml:"m:GetServiceConfigurationResponse,omitempty"`
		Fault                           *Fault                                      `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetServiceConfiguration", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetServiceConfigurationSoapOut{
		ServerVersionInfo:               outputHeader.ServerVersionInfo,
		GetServiceConfigurationResponse: outputBody.GetServiceConfigurationResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetMailTips(ctx context.Context, input *GetMailTipsSoapIn, detail any) (*GetMailTipsSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
		MailboxCulture       *MailboxCultureType       `xml:"t:MailboxCulture,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
		MailboxCulture:       input.MailboxCulture,
	}
	inputBody := &struct {
		GetMailTips *GetMailTipsType `xml:"m:GetMailTips,omitempty"`
	}{
		GetMailTips: input.GetMailTips,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetMailTipsResponse *GetMailTipsResponseMessageType `xml:"m:GetMailTipsResponse,omitempty"`
		Fault               *Fault                          `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetMailTips", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetMailTipsSoapOut{
		ServerVersionInfo:   outputHeader.ServerVersionInfo,
		GetMailTipsResponse: outputBody.GetMailTipsResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) PlayOnPhone(ctx context.Context, input *PlayOnPhoneSoapIn, detail any) (*PlayOnPhoneSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		PlayOnPhone *PlayOnPhoneType `xml:"m:PlayOnPhone,omitempty"`
	}{
		PlayOnPhone: input.PlayOnPhone,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		PlayOnPhoneResponse *PlayOnPhoneResponseMessageType `xml:"m:PlayOnPhoneResponse,omitempty"`
		Fault               *Fault                          `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/PlayOnPhone", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &PlayOnPhoneSoapOut{
		ServerVersionInfo:   outputHeader.ServerVersionInfo,
		PlayOnPhoneResponse: outputBody.PlayOnPhoneResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetPhoneCallInformation(ctx context.Context, input *GetPhoneCallInformationSoapIn, detail any) (*GetPhoneCallInformationSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		GetPhoneCallInformation *GetPhoneCallInformationType `xml:"m:GetPhoneCallInformation,omitempty"`
	}{
		GetPhoneCallInformation: input.GetPhoneCallInformation,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetPhoneCallInformationResponse *GetPhoneCallInformationResponseMessageType `xml:"m:GetPhoneCallInformationResponse,omitempty"`
		Fault                           *Fault                                      `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetPhoneCallInformation", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetPhoneCallInformationSoapOut{
		ServerVersionInfo:               outputHeader.ServerVersionInfo,
		GetPhoneCallInformationResponse: outputBody.GetPhoneCallInformationResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) DisconnectPhoneCall(ctx context.Context, input *DisconnectPhoneCallSoapIn, detail any) (*DisconnectPhoneCallSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		DisconnectPhoneCall *DisconnectPhoneCallType `xml:"m:DisconnectPhoneCall,omitempty"`
	}{
		DisconnectPhoneCall: input.DisconnectPhoneCall,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		DisconnectPhoneCallResponse *DisconnectPhoneCallResponseMessageType `xml:"m:DisconnectPhoneCallResponse,omitempty"`
		Fault                       *Fault                                  `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/DisconnectPhoneCall", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &DisconnectPhoneCallSoapOut{
		ServerVersionInfo:           outputHeader.ServerVersionInfo,
		DisconnectPhoneCallResponse: outputBody.DisconnectPhoneCallResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetSharingMetadata(ctx context.Context, input *GetSharingMetadataSoapIn, detail any) (*GetSharingMetadataSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
	}
	inputBody := &struct {
		GetSharingMetadata *GetSharingMetadataType `xml:"m:GetSharingMetadata,omitempty"`
	}{
		GetSharingMetadata: input.GetSharingMetadata,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetSharingMetadataResponse *GetSharingMetadataResponseMessageType `xml:"m:GetSharingMetadataResponse,omitempty"`
		Fault                      *Fault                                 `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetSharingMetadata", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetSharingMetadataSoapOut{
		ServerVersionInfo:          outputHeader.ServerVersionInfo,
		GetSharingMetadataResponse: outputBody.GetSharingMetadataResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) RefreshSharingFolder(ctx context.Context, input *RefreshSharingFolderSoapIn, detail any) (*RefreshSharingFolderSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
	}
	inputBody := &struct {
		RefreshSharingFolder *RefreshSharingFolderType `xml:"m:RefreshSharingFolder,omitempty"`
	}{
		RefreshSharingFolder: input.RefreshSharingFolder,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		RefreshSharingFolderResponse *RefreshSharingFolderResponseMessageType `xml:"m:RefreshSharingFolderResponse,omitempty"`
		Fault                        *Fault                                   `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/RefreshSharingFolder", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &RefreshSharingFolderSoapOut{
		ServerVersionInfo:            outputHeader.ServerVersionInfo,
		RefreshSharingFolderResponse: outputBody.RefreshSharingFolderResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetSharingFolder(ctx context.Context, input *GetSharingFolderSoapIn, detail any) (*GetSharingFolderSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
	}
	inputBody := &struct {
		GetSharingFolder *GetSharingFolderType `xml:"m:GetSharingFolder,omitempty"`
	}{
		GetSharingFolder: input.GetSharingFolder,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetSharingFolderResponse *GetSharingFolderResponseMessageType `xml:"m:GetSharingFolderResponse,omitempty"`
		Fault                    *Fault                               `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetSharingFolder", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetSharingFolderSoapOut{
		ServerVersionInfo:        outputHeader.ServerVersionInfo,
		GetSharingFolderResponse: outputBody.GetSharingFolderResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) SetTeamMailbox(ctx context.Context, input *SetTeamMailboxSoapIn, detail any) (*SetTeamMailboxSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
		ManagementRole       *ManagementRoleType       `xml:"t:ManagementRole,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
		ManagementRole:       input.ManagementRole,
	}
	inputBody := &struct {
		SetTeamMailbox *SetTeamMailboxRequestType `xml:"m:SetTeamMailbox,omitempty"`
	}{
		SetTeamMailbox: input.SetTeamMailbox,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		SetTeamMailboxResponse *SetTeamMailboxResponseMessageType `xml:"m:SetTeamMailboxResponse,omitempty"`
		Fault                  *Fault                             `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/SetTeamMailbox", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &SetTeamMailboxSoapOut{
		ServerVersionInfo:      outputHeader.ServerVersionInfo,
		SetTeamMailboxResponse: outputBody.SetTeamMailboxResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) UnpinTeamMailbox(ctx context.Context, input *UnpinTeamMailboxSoapIn, detail any) (*UnpinTeamMailboxSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
	}
	inputBody := &struct {
		UnpinTeamMailbox *UnpinTeamMailboxRequestType `xml:"m:UnpinTeamMailbox,omitempty"`
	}{
		UnpinTeamMailbox: input.UnpinTeamMailbox,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		UnpinTeamMailboxResponse *UnpinTeamMailboxResponseMessageType `xml:"m:UnpinTeamMailboxResponse,omitempty"`
		Fault                    *Fault                               `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/UnpinTeamMailbox", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &UnpinTeamMailboxSoapOut{
		ServerVersionInfo:        outputHeader.ServerVersionInfo,
		UnpinTeamMailboxResponse: outputBody.UnpinTeamMailboxResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetRoomLists(ctx context.Context, input *GetRoomListsSoapIn, detail any) (*GetRoomListsSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		GetRoomLists *GetRoomListsType `xml:"m:GetRoomLists,omitempty"`
	}{
		GetRoomLists: input.GetRoomLists,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetRoomListsResponse *GetRoomListsResponseMessageType `xml:"m:GetRoomListsResponse,omitempty"`
		Fault                *Fault                           `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetRoomLists", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetRoomListsSoapOut{
		ServerVersionInfo:    outputHeader.ServerVersionInfo,
		GetRoomListsResponse: outputBody.GetRoomListsResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetRooms(ctx context.Context, input *GetRoomsSoapIn, detail any) (*GetRoomsSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		GetRooms *GetRoomsType `xml:"m:GetRooms,omitempty"`
	}{
		GetRooms: input.GetRooms,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetRoomsResponse *GetRoomsResponseMessageType `xml:"m:GetRoomsResponse,omitempty"`
		Fault            *Fault                       `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetRooms", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetRoomsSoapOut{
		ServerVersionInfo: outputHeader.ServerVersionInfo,
		GetRoomsResponse:  outputBody.GetRoomsResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) FindMessageTrackingReport(ctx context.Context, input *FindMessageTrackingReportSoapIn, detail any) (*FindMessageTrackingReportSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
	}
	inputBody := &struct {
		FindMessageTrackingReport *FindMessageTrackingReportRequestType `xml:"m:FindMessageTrackingReport,omitempty"`
	}{
		FindMessageTrackingReport: input.FindMessageTrackingReport,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		FindMessageTrackingReportResponse *FindMessageTrackingReportResponseMessageType `xml:"m:FindMessageTrackingReportResponse,omitempty"`
		Fault                             *Fault                                        `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/FindMessageTrackingReport", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &FindMessageTrackingReportSoapOut{
		ServerVersionInfo:                 outputHeader.ServerVersionInfo,
		FindMessageTrackingReportResponse: outputBody.FindMessageTrackingReportResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetMessageTrackingReport(ctx context.Context, input *GetMessageTrackingReportSoapIn, detail any) (*GetMessageTrackingReportSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
	}
	inputBody := &struct {
		GetMessageTrackingReport *GetMessageTrackingReportRequestType `xml:"m:GetMessageTrackingReport,omitempty"`
	}{
		GetMessageTrackingReport: input.GetMessageTrackingReport,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetMessageTrackingReportResponse *GetMessageTrackingReportResponseMessageType `xml:"m:GetMessageTrackingReportResponse,omitempty"`
		Fault                            *Fault                                       `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetMessageTrackingReport", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetMessageTrackingReportSoapOut{
		ServerVersionInfo:                outputHeader.ServerVersionInfo,
		GetMessageTrackingReportResponse: outputBody.GetMessageTrackingReportResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) FindConversation(ctx context.Context, input *FindConversationSoapIn, detail any) (*FindConversationSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	}{
		RequestServerVersion:  input.RequestServerVersion,
		ExchangeImpersonation: input.ExchangeImpersonation,
	}
	inputBody := &struct {
		FindConversation *FindConversationType `xml:"m:FindConversation,omitempty"`
	}{
		FindConversation: input.FindConversation,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		FindConversationResponse *FindConversationResponseMessageType `xml:"m:FindConversationResponse,omitempty"`
		Fault                    *Fault                               `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/FindConversation", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &FindConversationSoapOut{
		ServerVersionInfo:        outputHeader.ServerVersionInfo,
		FindConversationResponse: outputBody.FindConversationResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) ApplyConversationAction(ctx context.Context, input *ApplyConversationActionSoapIn, detail any) (*ApplyConversationActionSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	}{
		RequestServerVersion:  input.RequestServerVersion,
		ExchangeImpersonation: input.ExchangeImpersonation,
	}
	inputBody := &struct {
		ApplyConversationAction *ApplyConversationActionType `xml:"m:ApplyConversationAction,omitempty"`
	}{
		ApplyConversationAction: input.ApplyConversationAction,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		ApplyConversationActionResponse *ApplyConversationActionResponseType `xml:"m:ApplyConversationActionResponse,omitempty"`
		Fault                           *Fault                               `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/ApplyConversationAction", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &ApplyConversationActionSoapOut{
		ServerVersionInfo:               outputHeader.ServerVersionInfo,
		ApplyConversationActionResponse: outputBody.ApplyConversationActionResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetConversationItems(ctx context.Context, input *GetConversationItemsSoapIn, detail any) (*GetConversationItemsSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	}{
		RequestServerVersion:  input.RequestServerVersion,
		ExchangeImpersonation: input.ExchangeImpersonation,
	}
	inputBody := &struct {
		GetConversationItems *GetConversationItemsType `xml:"m:GetConversationItems,omitempty"`
	}{
		GetConversationItems: input.GetConversationItems,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetConversationItemsResponse *GetConversationItemsResponseType `xml:"m:GetConversationItemsResponse,omitempty"`
		Fault                        *Fault                            `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetConversationItems", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetConversationItemsSoapOut{
		ServerVersionInfo:            outputHeader.ServerVersionInfo,
		GetConversationItemsResponse: outputBody.GetConversationItemsResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) FindPeople(ctx context.Context, input *FindPeopleSoapIn, detail any) (*FindPeopleSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	}{
		RequestServerVersion:  input.RequestServerVersion,
		ExchangeImpersonation: input.ExchangeImpersonation,
	}
	inputBody := &struct {
		FindPeople *FindPeopleType `xml:"m:FindPeople,omitempty"`
	}{
		FindPeople: input.FindPeople,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		FindPeopleResponse *FindPeopleResponseMessageType `xml:"m:FindPeopleResponse,omitempty"`
		Fault              *Fault                         `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/FindPeople", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &FindPeopleSoapOut{
		ServerVersionInfo:  outputHeader.ServerVersionInfo,
		FindPeopleResponse: outputBody.FindPeopleResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) FindTags(ctx context.Context, input *FindTagsSoapIn, detail any) (*FindTagsSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	}{
		RequestServerVersion:  input.RequestServerVersion,
		ExchangeImpersonation: input.ExchangeImpersonation,
	}
	inputBody := &struct {
		FindTags *FindTagsType `xml:"m:FindTags,omitempty"`
	}{
		FindTags: input.FindTags,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		FindTagsResponse *FindTagsResponseMessageType `xml:"m:FindTagsResponse,omitempty"`
		Fault            *Fault                       `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/FindTags", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &FindTagsSoapOut{
		ServerVersionInfo: outputHeader.ServerVersionInfo,
		FindTagsResponse:  outputBody.FindTagsResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) AddTag(ctx context.Context, input *AddTagSoapIn, detail any) (*AddTagSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	}{
		RequestServerVersion:  input.RequestServerVersion,
		ExchangeImpersonation: input.ExchangeImpersonation,
	}
	inputBody := &struct {
		AddTag *AddTagType `xml:"m:AddTag,omitempty"`
	}{
		AddTag: input.AddTag,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		AddTagResponse *AddTagResponseMessageType `xml:"m:AddTagResponse,omitempty"`
		Fault          *Fault                     `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/AddTag", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &AddTagSoapOut{
		ServerVersionInfo: outputHeader.ServerVersionInfo,
		AddTagResponse:    outputBody.AddTagResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) HideTag(ctx context.Context, input *HideTagSoapIn, detail any) (*HideTagSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	}{
		RequestServerVersion:  input.RequestServerVersion,
		ExchangeImpersonation: input.ExchangeImpersonation,
	}
	inputBody := &struct {
		HideTag *HideTagType `xml:"m:HideTag,omitempty"`
	}{
		HideTag: input.HideTag,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		HideTagResponse *HideTagResponseMessageType `xml:"m:HideTagResponse,omitempty"`
		Fault           *Fault                      `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/HideTag", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &HideTagSoapOut{
		ServerVersionInfo: outputHeader.ServerVersionInfo,
		HideTagResponse:   outputBody.HideTagResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetPersona(ctx context.Context, input *GetPersonaSoapIn, detail any) (*GetPersonaSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
	}{
		RequestServerVersion:  input.RequestServerVersion,
		ExchangeImpersonation: input.ExchangeImpersonation,
	}
	inputBody := &struct {
		GetPersona *GetPersonaType `xml:"m:GetPersona,omitempty"`
	}{
		GetPersona: input.GetPersona,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetPersonaResponseMessage *GetPersonaResponseMessageType `xml:"m:GetPersonaResponseMessage,omitempty"`
		Fault                     *Fault                         `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetPersona", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetPersonaSoapOut{
		ServerVersionInfo:         outputHeader.ServerVersionInfo,
		GetPersonaResponseMessage: outputBody.GetPersonaResponseMessage,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetInboxRules(ctx context.Context, input *GetInboxRulesSoapIn, detail any) (*GetInboxRulesSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
		TimeZoneContext       *TimeZoneContextType       `xml:"t:TimeZoneContext,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
		TimeZoneContext:       input.TimeZoneContext,
	}
	inputBody := &struct {
		GetInboxRules *GetInboxRulesRequestType `xml:"m:GetInboxRules,omitempty"`
	}{
		GetInboxRules: input.GetInboxRules,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetInboxRulesResponse *GetInboxRulesResponseType `xml:"m:GetInboxRulesResponse,omitempty"`
		Fault                 *Fault                     `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetInboxRules", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetInboxRulesSoapOut{
		ServerVersionInfo:     outputHeader.ServerVersionInfo,
		GetInboxRulesResponse: outputBody.GetInboxRulesResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) UpdateInboxRules(ctx context.Context, input *UpdateInboxRulesSoapIn, detail any) (*UpdateInboxRulesSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
		TimeZoneContext       *TimeZoneContextType       `xml:"t:TimeZoneContext,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
		TimeZoneContext:       input.TimeZoneContext,
	}
	inputBody := &struct {
		UpdateInboxRules *UpdateInboxRulesRequestType `xml:"m:UpdateInboxRules,omitempty"`
	}{
		UpdateInboxRules: input.UpdateInboxRules,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		UpdateInboxRulesResponse *UpdateInboxRulesResponseType `xml:"m:UpdateInboxRulesResponse,omitempty"`
		Fault                    *Fault                        `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/UpdateInboxRules", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &UpdateInboxRulesSoapOut{
		ServerVersionInfo:        outputHeader.ServerVersionInfo,
		UpdateInboxRulesResponse: outputBody.UpdateInboxRulesResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetPasswordExpirationDate(ctx context.Context, input *GetPasswordExpirationDateSoapIn, detail any) (*GetPasswordExpirationDateSoapOut, error) {
	inputHeader := &struct {
		MailboxCulture       *MailboxCultureType       `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
	}{
		MailboxCulture:       input.MailboxCulture,
		RequestServerVersion: input.RequestServerVersion,
	}
	inputBody := &struct {
		GetPasswordExpirationDate *GetPasswordExpirationDateType `xml:"m:GetPasswordExpirationDate,omitempty"`
	}{
		GetPasswordExpirationDate: input.GetPasswordExpirationDate,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetPasswordExpirationDateResponse *GetPasswordExpirationDateResponseMessageType `xml:"m:GetPasswordExpirationDateResponse,omitempty"`
		Fault                             *Fault                                        `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetPasswordExpirationDate", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetPasswordExpirationDateSoapOut{
		ServerVersionInfo:                 outputHeader.ServerVersionInfo,
		GetPasswordExpirationDateResponse: outputBody.GetPasswordExpirationDateResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetDiscoverySearchConfiguration(ctx context.Context, input *GetDiscoverySearchConfigurationSoapIn, detail any) (*GetDiscoverySearchConfigurationSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
		ManagementRole       *ManagementRoleType       `xml:"t:ManagementRole,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
		ManagementRole:       input.ManagementRole,
	}
	inputBody := &struct {
		GetDiscoverySearchConfiguration *GetDiscoverySearchConfigurationType `xml:"m:GetDiscoverySearchConfiguration,omitempty"`
	}{
		GetDiscoverySearchConfiguration: input.GetDiscoverySearchConfiguration,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetDiscoverySearchConfigurationResponse *GetDiscoverySearchConfigurationResponseMessageType `xml:"m:GetDiscoverySearchConfigurationResponse,omitempty"`
		Fault                                   *Fault                                              `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetDiscoverySearchConfiguration", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetDiscoverySearchConfigurationSoapOut{
		ServerVersionInfo:                       outputHeader.ServerVersionInfo,
		GetDiscoverySearchConfigurationResponse: outputBody.GetDiscoverySearchConfigurationResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetSearchableMailboxes(ctx context.Context, input *GetSearchableMailboxesSoapIn, detail any) (*GetSearchableMailboxesSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
		ManagementRole       *ManagementRoleType       `xml:"t:ManagementRole,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
		ManagementRole:       input.ManagementRole,
	}
	inputBody := &struct {
		GetSearchableMailboxes *GetSearchableMailboxesType `xml:"m:GetSearchableMailboxes,omitempty"`
	}{
		GetSearchableMailboxes: input.GetSearchableMailboxes,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetSearchableMailboxesResponse *GetSearchableMailboxesResponseMessageType `xml:"m:GetSearchableMailboxesResponse,omitempty"`
		Fault                          *Fault                                     `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetSearchableMailboxes", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetSearchableMailboxesSoapOut{
		ServerVersionInfo:              outputHeader.ServerVersionInfo,
		GetSearchableMailboxesResponse: outputBody.GetSearchableMailboxesResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) SearchMailboxes(ctx context.Context, input *SearchMailboxesSoapIn, detail any) (*SearchMailboxesSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
		ManagementRole       *ManagementRoleType       `xml:"t:ManagementRole,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
		ManagementRole:       input.ManagementRole,
	}
	inputBody := &struct {
		SearchMailboxes *SearchMailboxesType `xml:"m:SearchMailboxes,omitempty"`
	}{
		SearchMailboxes: input.SearchMailboxes,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		SearchMailboxesResponse *SearchMailboxesResponseType `xml:"m:SearchMailboxesResponse,omitempty"`
		Fault                   *Fault                       `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/SearchMailboxes", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &SearchMailboxesSoapOut{
		ServerVersionInfo:       outputHeader.ServerVersionInfo,
		SearchMailboxesResponse: outputBody.SearchMailboxesResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetHoldOnMailboxes(ctx context.Context, input *GetHoldOnMailboxesSoapIn, detail any) (*GetHoldOnMailboxesSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
		ManagementRole       *ManagementRoleType       `xml:"t:ManagementRole,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
		ManagementRole:       input.ManagementRole,
	}
	inputBody := &struct {
		GetHoldOnMailboxes *GetHoldOnMailboxesType `xml:"m:GetHoldOnMailboxes,omitempty"`
	}{
		GetHoldOnMailboxes: input.GetHoldOnMailboxes,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetHoldOnMailboxesResponse *GetHoldOnMailboxesResponseMessageType `xml:"m:GetHoldOnMailboxesResponse,omitempty"`
		Fault                      *Fault                                 `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetHoldOnMailboxes", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetHoldOnMailboxesSoapOut{
		ServerVersionInfo:          outputHeader.ServerVersionInfo,
		GetHoldOnMailboxesResponse: outputBody.GetHoldOnMailboxesResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) SetHoldOnMailboxes(ctx context.Context, input *SetHoldOnMailboxesSoapIn, detail any) (*SetHoldOnMailboxesSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
		ManagementRole       *ManagementRoleType       `xml:"t:ManagementRole,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
		ManagementRole:       input.ManagementRole,
	}
	inputBody := &struct {
		SetHoldOnMailboxes *SetHoldOnMailboxesType `xml:"m:SetHoldOnMailboxes,omitempty"`
	}{
		SetHoldOnMailboxes: input.SetHoldOnMailboxes,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		SetHoldOnMailboxesResponse *SetHoldOnMailboxesResponseMessageType `xml:"m:SetHoldOnMailboxesResponse,omitempty"`
		Fault                      *Fault                                 `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/SetHoldOnMailboxes", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &SetHoldOnMailboxesSoapOut{
		ServerVersionInfo:          outputHeader.ServerVersionInfo,
		SetHoldOnMailboxesResponse: outputBody.SetHoldOnMailboxesResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetNonIndexableItemStatistics(ctx context.Context, input *GetNonIndexableItemStatisticsSoapIn, detail any) (*GetNonIndexableItemStatisticsSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
		ManagementRole       *ManagementRoleType       `xml:"t:ManagementRole,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
		ManagementRole:       input.ManagementRole,
	}
	inputBody := &struct {
		GetNonIndexableItemStatistics *GetNonIndexableItemStatisticsType `xml:"m:GetNonIndexableItemStatistics,omitempty"`
	}{
		GetNonIndexableItemStatistics: input.GetNonIndexableItemStatistics,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetNonIndexableItemStatisticsResponse *GetNonIndexableItemStatisticsResponseMessageType `xml:"m:GetNonIndexableItemStatisticsResponse,omitempty"`
		Fault                                 *Fault                                            `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetNonIndexableItemStatistics", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetNonIndexableItemStatisticsSoapOut{
		ServerVersionInfo:                     outputHeader.ServerVersionInfo,
		GetNonIndexableItemStatisticsResponse: outputBody.GetNonIndexableItemStatisticsResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetNonIndexableItemDetails(ctx context.Context, input *GetNonIndexableItemDetailsSoapIn, detail any) (*GetNonIndexableItemDetailsSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
		ManagementRole       *ManagementRoleType       `xml:"t:ManagementRole,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
		ManagementRole:       input.ManagementRole,
	}
	inputBody := &struct {
		GetNonIndexableItemDetails *GetNonIndexableItemDetailsType `xml:"m:GetNonIndexableItemDetails,omitempty"`
	}{
		GetNonIndexableItemDetails: input.GetNonIndexableItemDetails,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetNonIndexableItemDetailsResponse *GetNonIndexableItemDetailsResponseMessageType `xml:"m:GetNonIndexableItemDetailsResponse,omitempty"`
		Fault                              *Fault                                         `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetNonIndexableItemDetails", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetNonIndexableItemDetailsSoapOut{
		ServerVersionInfo:                  outputHeader.ServerVersionInfo,
		GetNonIndexableItemDetailsResponse: outputBody.GetNonIndexableItemDetailsResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) MarkAllItemsAsRead(ctx context.Context, input *MarkAllItemsAsReadSoapIn, detail any) (*MarkAllItemsAsReadSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		MarkAllItemsAsRead *MarkAllItemsAsReadType `xml:"m:MarkAllItemsAsRead,omitempty"`
	}{
		MarkAllItemsAsRead: input.MarkAllItemsAsRead,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		MarkAllItemsAsReadResponse *MarkAllItemsAsReadResponseType `xml:"m:MarkAllItemsAsReadResponse,omitempty"`
		Fault                      *Fault                          `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/MarkAllItemsAsRead", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &MarkAllItemsAsReadSoapOut{
		ServerVersionInfo:          outputHeader.ServerVersionInfo,
		MarkAllItemsAsReadResponse: outputBody.MarkAllItemsAsReadResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) MarkAsJunk(ctx context.Context, input *MarkAsJunkSoapIn, detail any) (*MarkAsJunkSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		MarkAsJunk *MarkAsJunkType `xml:"m:MarkAsJunk,omitempty"`
	}{
		MarkAsJunk: input.MarkAsJunk,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		MarkAsJunkResponse *MarkAsJunkResponseType `xml:"m:MarkAsJunkResponse,omitempty"`
		Fault              *Fault                  `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/MarkAsJunk", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &MarkAsJunkSoapOut{
		ServerVersionInfo:  outputHeader.ServerVersionInfo,
		MarkAsJunkResponse: outputBody.MarkAsJunkResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) ReportMessage(ctx context.Context, input *ReportMessageSoapIn, detail any) (*ReportMessageSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		ReportMessage *ReportMessageType `xml:"m:ReportMessage,omitempty"`
	}{
		ReportMessage: input.ReportMessage,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		ReportMessageResponse *ReportMessageResponseType `xml:"m:ReportMessageResponse,omitempty"`
		Fault                 *Fault                     `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/ReportMessage", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &ReportMessageSoapOut{
		ServerVersionInfo:     outputHeader.ServerVersionInfo,
		ReportMessageResponse: outputBody.ReportMessageResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetAppManifests(ctx context.Context, input *GetAppManifestsSoapIn, detail any) (*GetAppManifestsSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
	}
	inputBody := &struct {
		GetAppManifests *GetAppManifestsType `xml:"m:GetAppManifests,omitempty"`
	}{
		GetAppManifests: input.GetAppManifests,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetAppManifestsResponse *GetAppManifestsResponseType `xml:"m:GetAppManifestsResponse,omitempty"`
		Fault                   *Fault                       `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetAppManifests", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetAppManifestsSoapOut{
		ServerVersionInfo:       outputHeader.ServerVersionInfo,
		GetAppManifestsResponse: outputBody.GetAppManifestsResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) AddNewImContactToGroup(ctx context.Context, input *AddNewImContactToGroupSoapIn, detail any) (*AddNewImContactToGroupSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		AddNewImContactToGroup *AddNewImContactToGroupType `xml:"m:AddNewImContactToGroup,omitempty"`
	}{
		AddNewImContactToGroup: input.AddNewImContactToGroup,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		AddNewImContactToGroupResponse *AddNewImContactToGroupResponseMessageType `xml:"m:AddNewImContactToGroupResponse,omitempty"`
		Fault                          *Fault                                     `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/AddNewImContactToGroup", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &AddNewImContactToGroupSoapOut{
		ServerVersionInfo:              outputHeader.ServerVersionInfo,
		AddNewImContactToGroupResponse: outputBody.AddNewImContactToGroupResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) AddNewTelUriContactToGroup(ctx context.Context, input *AddNewTelUriContactToGroupSoapIn, detail any) (*AddNewTelUriContactToGroupSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		AddNewTelUriContactToGroup *AddNewTelUriContactToGroupType `xml:"m:AddNewTelUriContactToGroup,omitempty"`
	}{
		AddNewTelUriContactToGroup: input.AddNewTelUriContactToGroup,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		AddNewTelUriContactToGroupResponse *AddNewTelUriContactToGroupResponseMessageType `xml:"m:AddNewTelUriContactToGroupResponse,omitempty"`
		Fault                              *Fault                                         `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/AddNewTelUriContactToGroup", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &AddNewTelUriContactToGroupSoapOut{
		ServerVersionInfo:                  outputHeader.ServerVersionInfo,
		AddNewTelUriContactToGroupResponse: outputBody.AddNewTelUriContactToGroupResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) AddImContactToGroup(ctx context.Context, input *AddImContactToGroupSoapIn, detail any) (*AddImContactToGroupSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		AddImContactToGroup *AddImContactToGroupType `xml:"m:AddImContactToGroup,omitempty"`
	}{
		AddImContactToGroup: input.AddImContactToGroup,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		AddImContactToGroupResponse *AddImContactToGroupResponseMessageType `xml:"m:AddImContactToGroupResponse,omitempty"`
		Fault                       *Fault                                  `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/AddImContactToGroup", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &AddImContactToGroupSoapOut{
		ServerVersionInfo:           outputHeader.ServerVersionInfo,
		AddImContactToGroupResponse: outputBody.AddImContactToGroupResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) RemoveImContactFromGroup(ctx context.Context, input *RemoveImContactFromGroupSoapIn, detail any) (*RemoveImContactFromGroupSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		RemoveImContactFromGroup *RemoveImContactFromGroupType `xml:"m:RemoveImContactFromGroup,omitempty"`
	}{
		RemoveImContactFromGroup: input.RemoveImContactFromGroup,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		RemoveImContactFromGroupResponse *RemoveImContactFromGroupResponseMessageType `xml:"m:RemoveImContactFromGroupResponse,omitempty"`
		Fault                            *Fault                                       `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/RemoveImContactFromGroup", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &RemoveImContactFromGroupSoapOut{
		ServerVersionInfo:                outputHeader.ServerVersionInfo,
		RemoveImContactFromGroupResponse: outputBody.RemoveImContactFromGroupResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) AddImGroup(ctx context.Context, input *AddImGroupSoapIn, detail any) (*AddImGroupSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		AddImGroup *AddImGroupType `xml:"m:AddImGroup,omitempty"`
	}{
		AddImGroup: input.AddImGroup,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		AddImGroupResponse *AddImGroupResponseMessageType `xml:"m:AddImGroupResponse,omitempty"`
		Fault              *Fault                         `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/AddImGroup", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &AddImGroupSoapOut{
		ServerVersionInfo:  outputHeader.ServerVersionInfo,
		AddImGroupResponse: outputBody.AddImGroupResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) AddDistributionGroupToImList(ctx context.Context, input *AddDistributionGroupToImListSoapIn, detail any) (*AddDistributionGroupToImListSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		AddDistributionGroupToImList *AddDistributionGroupToImListType `xml:"m:AddDistributionGroupToImList,omitempty"`
	}{
		AddDistributionGroupToImList: input.AddDistributionGroupToImList,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		AddDistributionGroupToImListResponse *AddDistributionGroupToImListResponseMessageType `xml:"m:AddDistributionGroupToImListResponse,omitempty"`
		Fault                                *Fault                                           `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/AddDistributionGroupToImList", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &AddDistributionGroupToImListSoapOut{
		ServerVersionInfo:                    outputHeader.ServerVersionInfo,
		AddDistributionGroupToImListResponse: outputBody.AddDistributionGroupToImListResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetImItemList(ctx context.Context, input *GetImItemListSoapIn, detail any) (*GetImItemListSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		GetImItemList *GetImItemListType `xml:"m:GetImItemList,omitempty"`
	}{
		GetImItemList: input.GetImItemList,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetImItemListResponse *GetImItemListResponseMessageType `xml:"m:GetImItemListResponse,omitempty"`
		Fault                 *Fault                            `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetImItemList", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetImItemListSoapOut{
		ServerVersionInfo:     outputHeader.ServerVersionInfo,
		GetImItemListResponse: outputBody.GetImItemListResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetImItems(ctx context.Context, input *GetImItemsSoapIn, detail any) (*GetImItemsSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		GetImItems *GetImItemsType `xml:"m:GetImItems,omitempty"`
	}{
		GetImItems: input.GetImItems,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetImItemsResponse *GetImItemsResponseMessageType `xml:"m:GetImItemsResponse,omitempty"`
		Fault              *Fault                         `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetImItems", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetImItemsSoapOut{
		ServerVersionInfo:  outputHeader.ServerVersionInfo,
		GetImItemsResponse: outputBody.GetImItemsResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) RemoveContactFromImList(ctx context.Context, input *RemoveContactFromImListSoapIn, detail any) (*RemoveContactFromImListSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		RemoveContactFromImList *RemoveContactFromImListType `xml:"m:RemoveContactFromImList,omitempty"`
	}{
		RemoveContactFromImList: input.RemoveContactFromImList,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		RemoveContactFromImListResponse *RemoveContactFromImListResponseMessageType `xml:"m:RemoveContactFromImListResponse,omitempty"`
		Fault                           *Fault                                      `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/RemoveContactFromImList", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &RemoveContactFromImListSoapOut{
		ServerVersionInfo:               outputHeader.ServerVersionInfo,
		RemoveContactFromImListResponse: outputBody.RemoveContactFromImListResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) RemoveDistributionGroupFromImList(ctx context.Context, input *RemoveDistributionGroupFromImListSoapIn, detail any) (*RemoveDistributionGroupFromImListSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		RemoveDistributionGroupFromImList *RemoveDistributionGroupFromImListType `xml:"m:RemoveDistributionGroupFromImList,omitempty"`
	}{
		RemoveDistributionGroupFromImList: input.RemoveDistributionGroupFromImList,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		RemoveDistributionGroupFromImListResponse *RemoveDistributionGroupFromImListResponseMessageType `xml:"m:RemoveDistributionGroupFromImListResponse,omitempty"`
		Fault                                     *Fault                                                `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/RemoveDistributionGroupFromImList", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &RemoveDistributionGroupFromImListSoapOut{
		ServerVersionInfo:                         outputHeader.ServerVersionInfo,
		RemoveDistributionGroupFromImListResponse: outputBody.RemoveDistributionGroupFromImListResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) RemoveImGroup(ctx context.Context, input *RemoveImGroupSoapIn, detail any) (*RemoveImGroupSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		RemoveImGroup *RemoveImGroupType `xml:"m:RemoveImGroup,omitempty"`
	}{
		RemoveImGroup: input.RemoveImGroup,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		RemoveImGroupResponse *RemoveImGroupResponseMessageType `xml:"m:RemoveImGroupResponse,omitempty"`
		Fault                 *Fault                            `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/RemoveImGroup", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &RemoveImGroupSoapOut{
		ServerVersionInfo:     outputHeader.ServerVersionInfo,
		RemoveImGroupResponse: outputBody.RemoveImGroupResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) SetImGroup(ctx context.Context, input *SetImGroupSoapIn, detail any) (*SetImGroupSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		SetImGroup *SetImGroupType `xml:"m:SetImGroup,omitempty"`
	}{
		SetImGroup: input.SetImGroup,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		SetImGroupResponse *SetImGroupResponseMessageType `xml:"m:SetImGroupResponse,omitempty"`
		Fault              *Fault                         `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/SetImGroup", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &SetImGroupSoapOut{
		ServerVersionInfo:  outputHeader.ServerVersionInfo,
		SetImGroupResponse: outputBody.SetImGroupResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) SetImListMigrationCompleted(ctx context.Context, input *SetImListMigrationCompletedSoapIn, detail any) (*SetImListMigrationCompletedSoapOut, error) {
	inputHeader := &struct {
		ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
		MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
		RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
	}{
		ExchangeImpersonation: input.ExchangeImpersonation,
		MailboxCulture:        input.MailboxCulture,
		RequestServerVersion:  input.RequestServerVersion,
	}
	inputBody := &struct {
		SetImListMigrationCompleted *SetImListMigrationCompletedType `xml:"m:SetImListMigrationCompleted,omitempty"`
	}{
		SetImListMigrationCompleted: input.SetImListMigrationCompleted,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		SetImListMigrationCompletedResponse *SetImListMigrationCompletedResponseMessageType `xml:"m:SetImListMigrationCompletedResponse,omitempty"`
		Fault                               *Fault                                          `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/SetImListMigrationCompleted", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &SetImListMigrationCompletedSoapOut{
		ServerVersionInfo:                   outputHeader.ServerVersionInfo,
		SetImListMigrationCompletedResponse: outputBody.SetImListMigrationCompletedResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetUserRetentionPolicyTags(ctx context.Context, input *GetUserRetentionPolicyTagsSoapIn, detail any) (*GetUserRetentionPolicyTagsSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
	}
	inputBody := &struct {
		GetUserRetentionPolicyTags *GetUserRetentionPolicyTagsType `xml:"m:GetUserRetentionPolicyTags,omitempty"`
	}{
		GetUserRetentionPolicyTags: input.GetUserRetentionPolicyTags,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetUserRetentionPolicyTagsResponse *GetUserRetentionPolicyTagsResponseMessageType `xml:"m:GetUserRetentionPolicyTagsResponse,omitempty"`
		Fault                              *Fault                                         `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetUserRetentionPolicyTags", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetUserRetentionPolicyTagsSoapOut{
		ServerVersionInfo:                  outputHeader.ServerVersionInfo,
		GetUserRetentionPolicyTagsResponse: outputBody.GetUserRetentionPolicyTagsResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) DisableApp(ctx context.Context, input *DisableAppSoapIn, detail any) (*DisableAppSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
	}
	inputBody := &struct {
		DisableApp *DisableAppType `xml:"m:DisableApp,omitempty"`
	}{
		DisableApp: input.DisableApp,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		DisableAppResponse *DisableAppResponseType `xml:"m:DisableAppResponse,omitempty"`
		Fault              *Fault                  `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/DisableApp", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &DisableAppSoapOut{
		ServerVersionInfo:  outputHeader.ServerVersionInfo,
		DisableAppResponse: outputBody.DisableAppResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) InstallApp(ctx context.Context, input *InstallAppSoapIn, detail any) (*InstallAppSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
	}
	inputBody := &struct {
		InstallApp *InstallAppType `xml:"m:InstallApp,omitempty"`
	}{
		InstallApp: input.InstallApp,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		InstallAppResponse *InstallAppResponseType `xml:"m:InstallAppResponse,omitempty"`
		Fault              *Fault                  `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/InstallApp", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &InstallAppSoapOut{
		ServerVersionInfo:  outputHeader.ServerVersionInfo,
		InstallAppResponse: outputBody.InstallAppResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) UpdateExtensionUsage(ctx context.Context, input *UpdateExtensionUsageSoapIn, detail any) (*UpdateExtensionUsageSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
	}
	inputBody := &struct {
		UpdateExtensionUsage *UpdateExtensionUsageType `xml:"m:UpdateExtensionUsage,omitempty"`
	}{
		UpdateExtensionUsage: input.UpdateExtensionUsage,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		UpdateExtensionUsageResponse *UpdateExtensionUsageResponseType `xml:"m:UpdateExtensionUsageResponse,omitempty"`
		Fault                        *Fault                            `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/UpdateExtensionUsage", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &UpdateExtensionUsageSoapOut{
		ServerVersionInfo:            outputHeader.ServerVersionInfo,
		UpdateExtensionUsageResponse: outputBody.UpdateExtensionUsageResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) UninstallApp(ctx context.Context, input *UninstallAppSoapIn, detail any) (*UninstallAppSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
	}
	inputBody := &struct {
		UninstallApp *UninstallAppType `xml:"m:UninstallApp,omitempty"`
	}{
		UninstallApp: input.UninstallApp,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		UninstallAppResponse *UninstallAppResponseType `xml:"m:UninstallAppResponse,omitempty"`
		Fault                *Fault                    `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/UninstallApp", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &UninstallAppSoapOut{
		ServerVersionInfo:    outputHeader.ServerVersionInfo,
		UninstallAppResponse: outputBody.UninstallAppResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetAppMarketplaceUrl(ctx context.Context, input *GetAppMarketplaceUrlSoapIn, detail any) (*GetAppMarketplaceUrlSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
	}
	inputBody := &struct {
		GetAppMarketplaceUrl *GetAppMarketplaceUrlType `xml:"m:GetAppMarketplaceUrl,omitempty"`
	}{
		GetAppMarketplaceUrl: input.GetAppMarketplaceUrl,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetAppMarketplaceUrlResponse *GetAppMarketplaceUrlResponseMessageType `xml:"m:GetAppMarketplaceUrlResponse,omitempty"`
		Fault                        *Fault                                   `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetAppMarketplaceUrl", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetAppMarketplaceUrlSoapOut{
		ServerVersionInfo:            outputHeader.ServerVersionInfo,
		GetAppMarketplaceUrlResponse: outputBody.GetAppMarketplaceUrlResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) FindAvailableMeetingTimes(ctx context.Context, input *FindAvailableMeetingTimesSoapIn, detail any) (*FindAvailableMeetingTimesSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
	}
	inputBody := &struct {
		FindAvailableMeetingTimes *FindAvailableMeetingTimesType `xml:"m:FindAvailableMeetingTimes,omitempty"`
	}{
		FindAvailableMeetingTimes: input.FindAvailableMeetingTimes,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		FindAvailableMeetingTimesResponse *FindAvailableMeetingTimesResponseMessageType `xml:"m:FindAvailableMeetingTimesResponse,omitempty"`
		Fault                             *Fault                                        `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/FindAvailableMeetingTimes", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &FindAvailableMeetingTimesSoapOut{
		ServerVersionInfo:                 outputHeader.ServerVersionInfo,
		FindAvailableMeetingTimesResponse: outputBody.FindAvailableMeetingTimesResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) FindMeetingTimeCandidates(ctx context.Context, input *FindMeetingTimeCandidatesSoapIn, detail any) (*FindMeetingTimeCandidatesSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
	}
	inputBody := &struct {
		FindMeetingTimeCandidates *FindMeetingTimeCandidatesType `xml:"m:FindMeetingTimeCandidates,omitempty"`
	}{
		FindMeetingTimeCandidates: input.FindMeetingTimeCandidates,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		FindMeetingTimeCandidatesResponse *FindMeetingTimeCandidatesResponseMessageType `xml:"m:FindMeetingTimeCandidatesResponse,omitempty"`
		Fault                             *Fault                                        `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/FindMeetingTimeCandidates", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &FindMeetingTimeCandidatesSoapOut{
		ServerVersionInfo:                 outputHeader.ServerVersionInfo,
		FindMeetingTimeCandidatesResponse: outputBody.FindMeetingTimeCandidatesResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetUserPhoto(ctx context.Context, input *GetUserPhotoSoapIn, detail any) (*GetUserPhotoSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
	}
	inputBody := &struct {
		GetUserPhoto *GetUserPhotoType `xml:"m:GetUserPhoto,omitempty"`
	}{
		GetUserPhoto: input.GetUserPhoto,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetUserPhotoResponse *GetUserPhotoResponseMessageType `xml:"m:GetUserPhotoResponse,omitempty"`
		Fault                *Fault                           `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetUserPhoto", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetUserPhotoSoapOut{
		ServerVersionInfo:    outputHeader.ServerVersionInfo,
		GetUserPhotoResponse: outputBody.GetUserPhotoResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) SetUserPhoto(ctx context.Context, input *SetUserPhotoSoapIn, detail any) (*SetUserPhotoSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
	}
	inputBody := &struct {
		SetUserPhoto *SetUserPhotoType `xml:"m:SetUserPhoto,omitempty"`
	}{
		SetUserPhoto: input.SetUserPhoto,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		SetUserPhotoResponse *SetUserPhotoResponseMessageType `xml:"m:SetUserPhotoResponse,omitempty"`
		Fault                *Fault                           `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/SetUserPhoto", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &SetUserPhotoSoapOut{
		ServerVersionInfo:    outputHeader.ServerVersionInfo,
		SetUserPhotoResponse: outputBody.SetUserPhotoResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetMeetingSpace(ctx context.Context, input *GetMeetingSpaceSoapIn, detail any) (*GetMeetingSpaceSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
		ManagementRole       *ManagementRoleType       `xml:"t:ManagementRole,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
		ManagementRole:       input.ManagementRole,
	}
	inputBody := &struct {
		GetMeetingSpace *GetMeetingSpaceType `xml:"m:GetMeetingSpace,omitempty"`
	}{
		GetMeetingSpace: input.GetMeetingSpace,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetMeetingSpaceResponseMessage *GetMeetingSpaceResponseMessageType `xml:"m:GetMeetingSpaceResponseMessage,omitempty"`
		Fault                          *Fault                              `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetMeetingSpace", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetMeetingSpaceSoapOut{
		ServerVersionInfo:              outputHeader.ServerVersionInfo,
		GetMeetingSpaceResponseMessage: outputBody.GetMeetingSpaceResponseMessage,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) DeleteMeetingSpace(ctx context.Context, input *DeleteMeetingSpaceSoapIn, detail any) (*DeleteMeetingSpaceSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
		ManagementRole       *ManagementRoleType       `xml:"t:ManagementRole,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
		ManagementRole:       input.ManagementRole,
	}
	inputBody := &struct {
		DeleteMeetingSpace *DeleteMeetingSpaceType `xml:"m:DeleteMeetingSpace,omitempty"`
	}{
		DeleteMeetingSpace: input.DeleteMeetingSpace,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		DeleteMeetingSpaceResponseMessage *DeleteMeetingSpaceResponseMessageType `xml:"m:DeleteMeetingSpaceResponseMessage,omitempty"`
		Fault                             *Fault                                 `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/DeleteMeetingSpace", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &DeleteMeetingSpaceSoapOut{
		ServerVersionInfo:                 outputHeader.ServerVersionInfo,
		DeleteMeetingSpaceResponseMessage: outputBody.DeleteMeetingSpaceResponseMessage,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) UpdateMeetingSpace(ctx context.Context, input *UpdateMeetingSpaceSoapIn, detail any) (*UpdateMeetingSpaceSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
		ManagementRole       *ManagementRoleType       `xml:"t:ManagementRole,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
		ManagementRole:       input.ManagementRole,
	}
	inputBody := &struct {
		UpdateMeetingSpace *UpdateMeetingSpaceType `xml:"m:UpdateMeetingSpace,omitempty"`
	}{
		UpdateMeetingSpace: input.UpdateMeetingSpace,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		UpdateMeetingSpaceResponseMessage *UpdateMeetingSpaceResponseMessageType `xml:"m:UpdateMeetingSpaceResponseMessage,omitempty"`
		Fault                             *Fault                                 `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/UpdateMeetingSpace", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &UpdateMeetingSpaceSoapOut{
		ServerVersionInfo:                 outputHeader.ServerVersionInfo,
		UpdateMeetingSpaceResponseMessage: outputBody.UpdateMeetingSpaceResponseMessage,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) CreateMeetingSpace(ctx context.Context, input *CreateMeetingSpaceSoapIn, detail any) (*CreateMeetingSpaceSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
		ManagementRole       *ManagementRoleType       `xml:"t:ManagementRole,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
		ManagementRole:       input.ManagementRole,
	}
	inputBody := &struct {
		CreateMeetingSpace *CreateMeetingSpaceType `xml:"m:CreateMeetingSpace,omitempty"`
	}{
		CreateMeetingSpace: input.CreateMeetingSpace,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		CreateMeetingSpaceResponseMessage *CreateMeetingSpaceResponseMessageType `xml:"m:CreateMeetingSpaceResponseMessage,omitempty"`
		Fault                             *Fault                                 `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/CreateMeetingSpace", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &CreateMeetingSpaceSoapOut{
		ServerVersionInfo:                 outputHeader.ServerVersionInfo,
		CreateMeetingSpaceResponseMessage: outputBody.CreateMeetingSpaceResponseMessage,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) FindMeetingSpaceByJoinUrl(ctx context.Context, input *FindMeetingSpaceByJoinUrlSoapIn, detail any) (*FindMeetingSpaceByJoinUrlSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
		ManagementRole       *ManagementRoleType       `xml:"t:ManagementRole,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
		ManagementRole:       input.ManagementRole,
	}
	inputBody := &struct {
		FindMeetingSpaceByJoinUrl *FindMeetingSpaceByJoinUrlType `xml:"m:FindMeetingSpaceByJoinUrl,omitempty"`
	}{
		FindMeetingSpaceByJoinUrl: input.FindMeetingSpaceByJoinUrl,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		FindMeetingSpaceByJoinUrlResponseMessage *FindMeetingSpaceByJoinUrlResponseMessageType `xml:"m:FindMeetingSpaceByJoinUrlResponseMessage,omitempty"`
		Fault                                    *Fault                                        `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/FindMeetingSpaceByJoinUrl", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &FindMeetingSpaceByJoinUrlSoapOut{
		ServerVersionInfo:                        outputHeader.ServerVersionInfo,
		FindMeetingSpaceByJoinUrlResponseMessage: outputBody.FindMeetingSpaceByJoinUrlResponseMessage,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetMeetingInstance(ctx context.Context, input *GetMeetingInstanceSoapIn, detail any) (*GetMeetingInstanceSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
		ManagementRole       *ManagementRoleType       `xml:"t:ManagementRole,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
		ManagementRole:       input.ManagementRole,
	}
	inputBody := &struct {
		GetMeetingInstanceRequest *GetMeetingInstanceRequestType `xml:"m:GetMeetingInstanceRequest,omitempty"`
	}{
		GetMeetingInstanceRequest: input.GetMeetingInstanceRequest,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetMeetingInstanceResponse *GetMeetingInstanceResponseMessageType `xml:"m:GetMeetingInstanceResponse,omitempty"`
		Fault                      *Fault                                 `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetMeetingInstanceRequest", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetMeetingInstanceSoapOut{
		ServerVersionInfo:          outputHeader.ServerVersionInfo,
		GetMeetingInstanceResponse: outputBody.GetMeetingInstanceResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) DeleteMeetingInstance(ctx context.Context, input *DeleteMeetingInstanceSoapIn, detail any) (*DeleteMeetingInstanceSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
		ManagementRole       *ManagementRoleType       `xml:"t:ManagementRole,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
		ManagementRole:       input.ManagementRole,
	}
	inputBody := &struct {
		DeleteMeetingInstanceRequest *DeleteMeetingInstanceRequestType `xml:"m:DeleteMeetingInstanceRequest,omitempty"`
	}{
		DeleteMeetingInstanceRequest: input.DeleteMeetingInstanceRequest,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		DeleteMeetingInstanceResponse *DeleteMeetingInstanceResponseMessageType `xml:"m:DeleteMeetingInstanceResponse,omitempty"`
		Fault                         *Fault                                    `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/DeleteMeetingInstanceRequest", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &DeleteMeetingInstanceSoapOut{
		ServerVersionInfo:             outputHeader.ServerVersionInfo,
		DeleteMeetingInstanceResponse: outputBody.DeleteMeetingInstanceResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) UpdateMeetingInstance(ctx context.Context, input *UpdateMeetingInstanceSoapIn, detail any) (*UpdateMeetingInstanceSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
		ManagementRole       *ManagementRoleType       `xml:"t:ManagementRole,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
		ManagementRole:       input.ManagementRole,
	}
	inputBody := &struct {
		UpdateMeetingInstanceRequest *UpdateMeetingInstanceRequestType `xml:"m:UpdateMeetingInstanceRequest,omitempty"`
	}{
		UpdateMeetingInstanceRequest: input.UpdateMeetingInstanceRequest,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		UpdateMeetingInstanceResponse *UpdateMeetingInstanceResponseMessageType `xml:"m:UpdateMeetingInstanceResponse,omitempty"`
		Fault                         *Fault                                    `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/UpdateMeetingInstanceRequest", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &UpdateMeetingInstanceSoapOut{
		ServerVersionInfo:             outputHeader.ServerVersionInfo,
		UpdateMeetingInstanceResponse: outputBody.UpdateMeetingInstanceResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) CreateMeetingInstance(ctx context.Context, input *CreateMeetingInstanceSoapIn, detail any) (*CreateMeetingInstanceSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
		ManagementRole       *ManagementRoleType       `xml:"t:ManagementRole,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
		ManagementRole:       input.ManagementRole,
	}
	inputBody := &struct {
		CreateMeetingInstanceRequest *CreateMeetingInstanceRequestType `xml:"m:CreateMeetingInstanceRequest,omitempty"`
	}{
		CreateMeetingInstanceRequest: input.CreateMeetingInstanceRequest,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		CreateMeetingInstanceResponse *CreateMeetingInstanceResponseMessageType `xml:"m:CreateMeetingInstanceResponse,omitempty"`
		Fault                         *Fault                                    `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/CreateMeetingInstanceRequest", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &CreateMeetingInstanceSoapOut{
		ServerVersionInfo:             outputHeader.ServerVersionInfo,
		CreateMeetingInstanceResponse: outputBody.CreateMeetingInstanceResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) StartSearchSession(ctx context.Context, input *StartSearchSessionSoapIn, detail any) (*StartSearchSessionSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
		MailboxCulture       *MailboxCultureType       `xml:"t:MailboxCulture,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
		MailboxCulture:       input.MailboxCulture,
	}
	inputBody := &struct {
		StartSearchSession *StartSearchSession `xml:"m:StartSearchSession,omitempty"`
	}{
		StartSearchSession: input.StartSearchSession,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		StartSearchSessionResponse *StartSearchSessionResponseMessage `xml:"m:StartSearchSessionResponse,omitempty"`
		Fault                      *Fault                             `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/StartSearchSession", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &StartSearchSessionSoapOut{
		ServerVersionInfo:          outputHeader.ServerVersionInfo,
		StartSearchSessionResponse: outputBody.StartSearchSessionResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) ExecuteSearch(ctx context.Context, input *ExecuteSearchSoapIn, detail any) (*ExecuteSearchSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
		MailboxCulture       *MailboxCultureType       `xml:"t:MailboxCulture,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
		MailboxCulture:       input.MailboxCulture,
	}
	inputBody := &struct {
		ExecuteSearch *ExecuteSearch `xml:"m:ExecuteSearch,omitempty"`
	}{
		ExecuteSearch: input.ExecuteSearch,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		ExecuteSearchResponse *ExecuteSearchResponseMessage `xml:"m:ExecuteSearchResponse,omitempty"`
		Fault                 *Fault                        `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/ExecuteSearch", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &ExecuteSearchSoapOut{
		ServerVersionInfo:     outputHeader.ServerVersionInfo,
		ExecuteSearchResponse: outputBody.ExecuteSearchResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetSearchSuggestions(ctx context.Context, input *GetSearchSuggestionsSoapIn, detail any) (*GetSearchSuggestionsSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
		MailboxCulture       *MailboxCultureType       `xml:"t:MailboxCulture,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
		MailboxCulture:       input.MailboxCulture,
	}
	inputBody := &struct {
		GetSearchSuggestions *GetSearchSuggestions `xml:"m:GetSearchSuggestions,omitempty"`
	}{
		GetSearchSuggestions: input.GetSearchSuggestions,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetSearchSuggestionsResponse *GetSearchSuggestionsResponseMessage `xml:"m:GetSearchSuggestionsResponse,omitempty"`
		Fault                        *Fault                               `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetSearchSuggestions", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetSearchSuggestionsSoapOut{
		ServerVersionInfo:            outputHeader.ServerVersionInfo,
		GetSearchSuggestionsResponse: outputBody.GetSearchSuggestionsResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) DeleteSearchSuggestion(ctx context.Context, input *DeleteSearchSuggestionSoapIn, detail any) (*DeleteSearchSuggestionSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
		MailboxCulture       *MailboxCultureType       `xml:"t:MailboxCulture,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
		MailboxCulture:       input.MailboxCulture,
	}
	inputBody := &struct {
		DeleteSearchSuggestion *DeleteSearchSuggestion `xml:"m:DeleteSearchSuggestion,omitempty"`
	}{
		DeleteSearchSuggestion: input.DeleteSearchSuggestion,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		DeleteSearchSuggestionResponse *DeleteSearchSuggestionResponseMessageType `xml:"m:DeleteSearchSuggestionResponse,omitempty"`
		Fault                          *Fault                                     `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/DeleteSearchSuggestion", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &DeleteSearchSuggestionSoapOut{
		ServerVersionInfo:              outputHeader.ServerVersionInfo,
		DeleteSearchSuggestionResponse: outputBody.DeleteSearchSuggestionResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) EndSearchSession(ctx context.Context, input *EndSearchSessionSoapIn, detail any) (*EndSearchSessionSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
		MailboxCulture       *MailboxCultureType       `xml:"t:MailboxCulture,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
		MailboxCulture:       input.MailboxCulture,
	}
	inputBody := &struct {
		EndSearchSession *EndSearchSession `xml:"m:EndSearchSession,omitempty"`
	}{
		EndSearchSession: input.EndSearchSession,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		EndSearchSessionResponse *EndSearchSessionResponseMessage `xml:"m:EndSearchSessionResponse,omitempty"`
		Fault                    *Fault                           `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/EndSearchSession", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &EndSearchSessionSoapOut{
		ServerVersionInfo:        outputHeader.ServerVersionInfo,
		EndSearchSessionResponse: outputBody.EndSearchSessionResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetLastPrivateCatalogUpdate(ctx context.Context, input *GetLastPrivateCatalogUpdateSoapIn, detail any) (*GetLastPrivateCatalogUpdateSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
	}
	inputBody := &struct {
		GetLastPrivateCatalogUpdate *GetLastPrivateCatalogUpdateType `xml:"m:GetLastPrivateCatalogUpdate,omitempty"`
	}{
		GetLastPrivateCatalogUpdate: input.GetLastPrivateCatalogUpdate,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetLastPrivateCatalogUpdateResponse *GetLastPrivateCatalogUpdateResponseType `xml:"m:GetLastPrivateCatalogUpdateResponse,omitempty"`
		Fault                               *Fault                                   `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetLastPrivateCatalogUpdate", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetLastPrivateCatalogUpdateSoapOut{
		ServerVersionInfo:                   outputHeader.ServerVersionInfo,
		GetLastPrivateCatalogUpdateResponse: outputBody.GetLastPrivateCatalogUpdateResponse,
	}
	return output, nil
}

func (b *ExchangeServiceBinding) GetPrivateCatalogAddIns(ctx context.Context, input *GetPrivateCatalogAddInsSoapIn, detail any) (*GetPrivateCatalogAddInsSoapOut, error) {
	inputHeader := &struct {
		RequestServerVersion *RequestServerVersionType `xml:"t:RequestServerVersion,omitempty"`
	}{
		RequestServerVersion: input.RequestServerVersion,
	}
	inputBody := &struct {
		GetPrivateCatalogAddIns *GetPrivateCatalogAddInsType `xml:"m:GetPrivateCatalogAddIns,omitempty"`
	}{
		GetPrivateCatalogAddIns: input.GetPrivateCatalogAddIns,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetPrivateCatalogAddInsResponse *GetPrivateCatalogAddInsResponseType `xml:"m:GetPrivateCatalogAddInsResponse,omitempty"`
		Fault                           *Fault                               `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	err := b.client.Call(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetPrivateCatalogAddIns", inputHeader, inputBody, outputHeader, outputBody)
	if err != nil {
		return nil, err
	}
	output := &GetPrivateCatalogAddInsSoapOut{
		ServerVersionInfo:               outputHeader.ServerVersionInfo,
		GetPrivateCatalogAddInsResponse: outputBody.GetPrivateCatalogAddInsResponse,
	}
	return output, nil
}

func NewExchangeServicePortType(client SOAPClient) *ExchangeServiceBinding {
	return &ExchangeServiceBinding{client: client}
}
