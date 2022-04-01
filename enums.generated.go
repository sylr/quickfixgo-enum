package enum

// AccountType field enumeration values.
type AccountType string

const (
	AccountType_ACCOUNT_IS_CARRIED_ON_CUSTOMER_SIDE_OF_THE_BOOKS                       AccountType = "1"
	AccountType_OPTIONS_MARKET_MAKER                                                   AccountType = "10"
	AccountType_OPTIONS_FIRM_ACCOUNT                                                   AccountType = "11"
	AccountType_ACCOUNT_FOR_CUSTOMER_AND_NON_CUSTOMER_ORDERS                           AccountType = "12"
	AccountType_ACCOUNT_FOR_ORDERS_FROM_MULTIPLE_CUSTOMERS                             AccountType = "13"
	AccountType_ACCOUNT_IS_CARRIED_ON_NON_CUSTOMER_SIDE_OF_BOOKS                       AccountType = "2"
	AccountType_HOUSE_TRADER                                                           AccountType = "3"
	AccountType_FLOOR_TRADER                                                           AccountType = "4"
	AccountType_ACCOUNT_IS_CARRIED_ON_NON_CUSTOMER_SIDE_OF_BOOKS_AND_IS_CROSS_MARGINED AccountType = "6"
	AccountType_ACCOUNT_IS_HOUSE_TRADER_AND_IS_CROSS_MARGINED                          AccountType = "7"
	AccountType_JOINT_BACK_OFFICE_ACCOUNT                                              AccountType = "8"
	AccountType_EQUITIES_SPECIALIST                                                    AccountType = "9"
)

// AcctIDSource field enumeration values.
type AcctIDSource string

const (
	AcctIDSource_BIC                           AcctIDSource = "1"
	AcctIDSource_SID_CODE                      AcctIDSource = "2"
	AcctIDSource_TFM                           AcctIDSource = "3"
	AcctIDSource_OMGEO                         AcctIDSource = "4"
	AcctIDSource_DTCC_CODE                     AcctIDSource = "5"
	AcctIDSource_SPECIAL_SEGREGATED_ACCOUNT_ID AcctIDSource = "6"
	AcctIDSource_OTHER                         AcctIDSource = "99"
)

// Adjustment field enumeration values.
type Adjustment string

const (
	Adjustment_CANCEL     Adjustment = "1"
	Adjustment_ERROR      Adjustment = "2"
	Adjustment_CORRECTION Adjustment = "3"
)

// AdjustmentType field enumeration values.
type AdjustmentType string

const (
	AdjustmentType_PROCESS_REQUEST_AS_MARGIN_DISPOSITION AdjustmentType = "0"
	AdjustmentType_DELTA_PLUS                            AdjustmentType = "1"
	AdjustmentType_DELTA_MINUS                           AdjustmentType = "2"
	AdjustmentType_FINAL                                 AdjustmentType = "3"
	AdjustmentType_CUSTOMER_SPECIFIC_POSITION            AdjustmentType = "4"
)

// AdvSide field enumeration values.
type AdvSide string

const (
	AdvSide_BUY   AdvSide = "B"
	AdvSide_SELL  AdvSide = "S"
	AdvSide_TRADE AdvSide = "T"
	AdvSide_CROSS AdvSide = "X"
)

// AdvTransType field enumeration values.
type AdvTransType string

const (
	AdvTransType_CANCEL  AdvTransType = "C"
	AdvTransType_NEW     AdvTransType = "N"
	AdvTransType_REPLACE AdvTransType = "R"
)

// AffirmStatus field enumeration values.
type AffirmStatus string

const (
	AffirmStatus_RECEIVED                         AffirmStatus = "1"
	AffirmStatus_CONFIRM_REJECTED_IE_NOT_AFFIRMED AffirmStatus = "2"
	AffirmStatus_AFFIRMED                         AffirmStatus = "3"
)

// AggregatedBook field enumeration values.
type AggregatedBook string

const (
	AggregatedBook_NO  AggregatedBook = "N"
	AggregatedBook_YES AggregatedBook = "Y"
)

// AggressorIndicator field enumeration values.
type AggressorIndicator string

const (
	AggressorIndicator_NO  AggressorIndicator = "N"
	AggressorIndicator_YES AggressorIndicator = "Y"
)

// AlgorithmicTradeIndicator field enumeration values.
type AlgorithmicTradeIndicator string

const (
	AlgorithmicTradeIndicator_NON_ALGORITHMIC_TRADE AlgorithmicTradeIndicator = "0"
	AlgorithmicTradeIndicator_ALGORITHMIC_TRADE     AlgorithmicTradeIndicator = "1"
)

// AllocAccountType field enumeration values.
type AllocAccountType string

const (
	AllocAccountType_ACCOUNT_IS_CARRIED_PN_CUSTOMER_SIDE_OF_BOOKS                           AllocAccountType = "1"
	AllocAccountType_ACCOUNT_IS_CARRIED_ON_NON_CUSTOMER_SIDE_OF_BOOKS                       AllocAccountType = "2"
	AllocAccountType_HOUSE_TRADER                                                           AllocAccountType = "3"
	AllocAccountType_FLOOR_TRADER                                                           AllocAccountType = "4"
	AllocAccountType_ACCOUNT_IS_CARRIED_ON_NON_CUSTOMER_SIDE_OF_BOOKS_AND_IS_CROSS_MARGINED AllocAccountType = "6"
	AllocAccountType_ACCOUNT_IS_HOUSE_TRADER_AND_IS_CROSS_MARGINED                          AllocAccountType = "7"
	AllocAccountType_JOINT_BACK_OFFICE_ACCOUNT                                              AllocAccountType = "8"
)

// AllocCancReplaceReason field enumeration values.
type AllocCancReplaceReason string

const (
	AllocCancReplaceReason_ORIGINAL_DETAILS_INCOMPLETE_INCORRECT AllocCancReplaceReason = "1"
	AllocCancReplaceReason_CHANGE_IN_UNDERLYING_ORDER_DETAILS    AllocCancReplaceReason = "2"
	AllocCancReplaceReason_CANCELLED_BY_GIVE_UP_FIRM             AllocCancReplaceReason = "3"
	AllocCancReplaceReason_OTHER                                 AllocCancReplaceReason = "99"
)

// AllocGroupStatus field enumeration values.
type AllocGroupStatus string

const (
	AllocGroupStatus_ADDED    AllocGroupStatus = "0"
	AllocGroupStatus_CANCELED AllocGroupStatus = "1"
	AllocGroupStatus_REPLACED AllocGroupStatus = "2"
	AllocGroupStatus_CHANGED  AllocGroupStatus = "3"
	AllocGroupStatus_PENDING  AllocGroupStatus = "4"
)

// AllocHandlInst field enumeration values.
type AllocHandlInst string

const (
	AllocHandlInst_MATCH              AllocHandlInst = "1"
	AllocHandlInst_FORWARD            AllocHandlInst = "2"
	AllocHandlInst_FORWARD_AND_MATCH  AllocHandlInst = "3"
	AllocHandlInst_AUTO_CLAIM_GIVE_UP AllocHandlInst = "4"
)

// AllocIntermedReqType field enumeration values.
type AllocIntermedReqType string

const (
	AllocIntermedReqType_PENDING_ACCEPT       AllocIntermedReqType = "1"
	AllocIntermedReqType_PENDING_RELEASE      AllocIntermedReqType = "2"
	AllocIntermedReqType_PENDING_REVERSAL     AllocIntermedReqType = "3"
	AllocIntermedReqType_ACCEPT               AllocIntermedReqType = "4"
	AllocIntermedReqType_BLOCK_LEVEL_REJECT   AllocIntermedReqType = "5"
	AllocIntermedReqType_ACCOUNT_LEVEL_REJECT AllocIntermedReqType = "6"
)

// AllocLinkType field enumeration values.
type AllocLinkType string

const (
	AllocLinkType_FX_NETTING AllocLinkType = "0"
	AllocLinkType_FX_SWAP    AllocLinkType = "1"
)

// AllocMethod field enumeration values.
type AllocMethod string

const (
	AllocMethod_AUTOMATIC       AllocMethod = "1"
	AllocMethod_GUARANTOR       AllocMethod = "2"
	AllocMethod_MANUAL          AllocMethod = "3"
	AllocMethod_BROKER_ASSIGNED AllocMethod = "4"
)

// AllocNoOrdersType field enumeration values.
type AllocNoOrdersType string

const (
	AllocNoOrdersType_NOT_SPECIFIED          AllocNoOrdersType = "0"
	AllocNoOrdersType_EXPLICIT_LIST_PROVIDED AllocNoOrdersType = "1"
)

// AllocPositionEffect field enumeration values.
type AllocPositionEffect string

const (
	AllocPositionEffect_CLOSE  AllocPositionEffect = "C"
	AllocPositionEffect_FIFO   AllocPositionEffect = "F"
	AllocPositionEffect_OPEN   AllocPositionEffect = "O"
	AllocPositionEffect_ROLLED AllocPositionEffect = "R"
)

// AllocRejCode field enumeration values.
type AllocRejCode string

const (
	AllocRejCode_UNKNOWN_OR_MISSING_ACCOUNT                            AllocRejCode = "0"
	AllocRejCode_INCORRECT_OR_MISSING_BLOCK_QUANTITY                   AllocRejCode = "1"
	AllocRejCode_UNKNOWN_OR_STALE_EXECID                               AllocRejCode = "10"
	AllocRejCode_MISMATCHED_DATA                                       AllocRejCode = "11"
	AllocRejCode_UNKNOWN_CLORDID                                       AllocRejCode = "12"
	AllocRejCode_WAREHOUSE_REQUEST_REJECTED                            AllocRejCode = "13"
	AllocRejCode_DUPLICATE_OR_MISSING_INDIVIDUALALLOCID                AllocRejCode = "14"
	AllocRejCode_TRADE_NOT_RECOGNIZED                                  AllocRejCode = "15"
	AllocRejCode_TRADE_PREVIOUSLY_ALLOCATED                            AllocRejCode = "16"
	AllocRejCode_INCORRECT_OR_MISSING_INSTRUMENT                       AllocRejCode = "17"
	AllocRejCode_INCORRECT_OR_MISSING_SETTLEMENT_DATE                  AllocRejCode = "18"
	AllocRejCode_INCORRECT_OR_MISSING_FUND_ID_OR_FUND_NAME             AllocRejCode = "19"
	AllocRejCode_INCORRECT_OR_MISSING_AVERAGE_PRICE                    AllocRejCode = "2"
	AllocRejCode_INCORRECT_OR_MISSING_SETTLEMENT_INSTRUCTIONS          AllocRejCode = "20"
	AllocRejCode_INCORRECT_OR_MISSING_FEES                             AllocRejCode = "21"
	AllocRejCode_INCORRECT_OR_MISSING_TAX                              AllocRejCode = "22"
	AllocRejCode_UNKNOWN_OR_MISSING_PARTY                              AllocRejCode = "23"
	AllocRejCode_INCORRECT_OR_MISSING_SIDE                             AllocRejCode = "24"
	AllocRejCode_INCORRECT_OR_MISSING_NET_MONEY                        AllocRejCode = "25"
	AllocRejCode_INCORRECT_OR_MISSING_TRADE_DATE                       AllocRejCode = "26"
	AllocRejCode_INCORRECT_OR_MISSING_SETTLEMENT_CURRENCY_INSTRUCTIONS AllocRejCode = "27"
	AllocRejCode_INCORRRECT_OR_MISSING_PROCESSCODE                     AllocRejCode = "28"
	AllocRejCode_UNKNOWN_EXECUTING_BROKER_MNEMONIC                     AllocRejCode = "3"
	AllocRejCode_INCORRECT_OR_MISSING_COMMISSION                       AllocRejCode = "4"
	AllocRejCode_UNKNOWN_ORDERID                                       AllocRejCode = "5"
	AllocRejCode_UNKNOWN_LISTID                                        AllocRejCode = "6"
	AllocRejCode_OTHER_7                                               AllocRejCode = "7"
	AllocRejCode_INCORRECT_OR_MISSING_ALLOCATED_QUANTITY               AllocRejCode = "8"
	AllocRejCode_CALCULATION_DIFFERENCE                                AllocRejCode = "9"
	AllocRejCode_OTHER_99                                              AllocRejCode = "99"
)

// AllocReportType field enumeration values.
type AllocReportType string

const (
	AllocReportType_REJECT                                  AllocReportType = "10"
	AllocReportType_ACCEPT_PENDING                          AllocReportType = "11"
	AllocReportType_COMPLETE                                AllocReportType = "12"
	AllocReportType_REVERSE_PENDING                         AllocReportType = "14"
	AllocReportType_GIVE_UP                                 AllocReportType = "15"
	AllocReportType_TAKE_UP                                 AllocReportType = "16"
	AllocReportType_REVERSAL                                AllocReportType = "17"
	AllocReportType_ALLEGED_REVERSAL                        AllocReportType = "18"
	AllocReportType_SUB_ALLOCATION_GIVE_UP                  AllocReportType = "19"
	AllocReportType_PRELIMINARY_REQUEST_TO_INTERMEDIARY     AllocReportType = "2"
	AllocReportType_SELLSIDE_CALCULATED_USING_PRELIMINARY   AllocReportType = "3"
	AllocReportType_SELLSIDE_CALCULATED_WITHOUT_PRELIMINARY AllocReportType = "4"
	AllocReportType_WAREHOUSE_RECAP                         AllocReportType = "5"
	AllocReportType_REQUEST_TO_INTERMEDIARY                 AllocReportType = "8"
	AllocReportType_ACCEPT                                  AllocReportType = "9"
)

// AllocRequestStatus field enumeration values.
type AllocRequestStatus string

const (
	AllocRequestStatus_ACCEPTED AllocRequestStatus = "0"
	AllocRequestStatus_REJECTED AllocRequestStatus = "1"
)

// AllocReversalStatus field enumeration values.
type AllocReversalStatus string

const (
	AllocReversalStatus_COMPLETED AllocReversalStatus = "0"
	AllocReversalStatus_REFUSED   AllocReversalStatus = "1"
	AllocReversalStatus_CANCELLED AllocReversalStatus = "2"
)

// AllocSettlInstType field enumeration values.
type AllocSettlInstType string

const (
	AllocSettlInstType_USE_DEFAULT_INSTRUCTIONS        AllocSettlInstType = "0"
	AllocSettlInstType_DERIVE_FROM_PARAMETERS_PROVIDED AllocSettlInstType = "1"
	AllocSettlInstType_FULL_DETAILS_PROVIDED           AllocSettlInstType = "2"
	AllocSettlInstType_SSI_DB_IDS_PROVIDED             AllocSettlInstType = "3"
	AllocSettlInstType_PHONE_FOR_INSTRUCTIONS          AllocSettlInstType = "4"
)

// AllocStatus field enumeration values.
type AllocStatus string

const (
	AllocStatus_ACCEPTED                  AllocStatus = "0"
	AllocStatus_BLOCK_LEVEL_REJECT        AllocStatus = "1"
	AllocStatus_REFUSED                   AllocStatus = "10"
	AllocStatus_PENDING_GIVE_UP_APPROVAL  AllocStatus = "11"
	AllocStatus_CANCELLED                 AllocStatus = "12"
	AllocStatus_PENDING_TAKE_UP_APPROVAL  AllocStatus = "13"
	AllocStatus_REVERSAL_PENDING          AllocStatus = "14"
	AllocStatus_ACCOUNT_LEVEL_REJECT      AllocStatus = "2"
	AllocStatus_RECEIVED                  AllocStatus = "3"
	AllocStatus_INCOMPLETE                AllocStatus = "4"
	AllocStatus_REJECTED_BY_INTERMEDIARY  AllocStatus = "5"
	AllocStatus_ALLOCATION_PENDING        AllocStatus = "6"
	AllocStatus_REVERSED                  AllocStatus = "7"
	AllocStatus_CANCELLED_BY_INTERMEDIARY AllocStatus = "8"
	AllocStatus_CLAIMED                   AllocStatus = "9"
)

// AllocTransType field enumeration values.
type AllocTransType string

const (
	AllocTransType_NEW                            AllocTransType = "0"
	AllocTransType_REPLACE                        AllocTransType = "1"
	AllocTransType_CANCEL                         AllocTransType = "2"
	AllocTransType_PRELIMINARY                    AllocTransType = "3"
	AllocTransType_CALCULATED                     AllocTransType = "4"
	AllocTransType_CALCULATED_WITHOUT_PRELIMINARY AllocTransType = "5"
	AllocTransType_REVERSAL                       AllocTransType = "6"
)

// AllocType field enumeration values.
type AllocType string

const (
	AllocType_CALCULATED                                    AllocType = "1"
	AllocType_REJECT                                        AllocType = "10"
	AllocType_ACCEPT_PENDING                                AllocType = "11"
	AllocType_INCOMPLETE_GROUP                              AllocType = "12"
	AllocType_COMPLETE_GROUP                                AllocType = "13"
	AllocType_REVERSAL_PENDING                              AllocType = "14"
	AllocType_REOPEN_GROUP                                  AllocType = "15"
	AllocType_CANCEL_GROUP                                  AllocType = "16"
	AllocType_GIVE_UP                                       AllocType = "17"
	AllocType_TAKE_UP                                       AllocType = "18"
	AllocType_REFUSE_TAKE_UP                                AllocType = "19"
	AllocType_PRELIMINARY                                   AllocType = "2"
	AllocType_INITIATE_REVERSAL                             AllocType = "20"
	AllocType_REVERSE                                       AllocType = "21"
	AllocType_REFUSE_REVERSAL                               AllocType = "22"
	AllocType_SUB_ALLOCATION_GIVE_UP                        AllocType = "23"
	AllocType_APPROVE_GIVE_UP                               AllocType = "24"
	AllocType_APPROVE_TAKE_UP                               AllocType = "25"
	AllocType_NOTIONAL_VALUE_AVERAGE_PRICE_GROUP_ALLOCATION AllocType = "26"
	AllocType_SELLSIDE_CALCULATED_USING_PRELIMINARY         AllocType = "3"
	AllocType_SELLSIDE_CALCULATEDD_WITHOUT_PRELIMINARY      AllocType = "4"
	AllocType_READY_TO_BOOK_SINGLE_ORDER                    AllocType = "5"
	AllocType_BUYSIDE_READY_TO_BOOK                         AllocType = "6"
	AllocType_WAREHOUSE_INSTRUCTION                         AllocType = "7"
	AllocType_REQUEST_TO_INTERMEDIARY                       AllocType = "8"
	AllocType_ACCEPT                                        AllocType = "9"
)

// AllocationRollupInstruction field enumeration values.
type AllocationRollupInstruction string

const (
	AllocationRollupInstruction_ROLL_UP        AllocationRollupInstruction = "0"
	AllocationRollupInstruction_DO_NOT_ROLL_UP AllocationRollupInstruction = "1"
)

// ApplLevelRecoveryIndicator field enumeration values.
type ApplLevelRecoveryIndicator string

const (
	ApplLevelRecoveryIndicator_APPLICATION_LEVEL_RECOVERY_IS_NOT_NEEDED ApplLevelRecoveryIndicator = "0"
	ApplLevelRecoveryIndicator_APPLICATION_LEVEL_RECOVERY_IS_NEEDED     ApplLevelRecoveryIndicator = "1"
)

// ApplQueueAction field enumeration values.
type ApplQueueAction string

const (
	ApplQueueAction_NO_ACTION_TAKEN ApplQueueAction = "0"
	ApplQueueAction_QUEUE_FLUSHED   ApplQueueAction = "1"
	ApplQueueAction_OVERLAY_LAST    ApplQueueAction = "2"
	ApplQueueAction_END_SESSION     ApplQueueAction = "3"
)

// ApplQueueResolution field enumeration values.
type ApplQueueResolution string

const (
	ApplQueueResolution_NO_ACTION_TAKEN ApplQueueResolution = "0"
	ApplQueueResolution_QUEUE_FLUSHED   ApplQueueResolution = "1"
	ApplQueueResolution_OVERLAY_LAST    ApplQueueResolution = "2"
	ApplQueueResolution_END_SESSION     ApplQueueResolution = "3"
)

// ApplReportType field enumeration values.
type ApplReportType string

const (
	ApplReportType_RESET_APPLSEQNUM_TO_NEW_VALUE_SPECIFIED_IN_APPLNEWSEQNUM                               ApplReportType = "0"
	ApplReportType_REPORTS_THAT_THE_LAST_MESSAGE_HAS_BEEN_SENT_FOR_THE_APPLIDS_REFER_TO_REFAPPLLASTSEQNUM ApplReportType = "1"
	ApplReportType_HEARTBEAT_MESSAGE_INDICATING_THAT_APPLICATION_IDENTIFIED_BY_REFAPPLID                  ApplReportType = "2"
	ApplReportType_APPLICATION_MESSAGE_RE_SEND_COMPLETED                                                  ApplReportType = "3"
)

// ApplReqType field enumeration values.
type ApplReqType string

const (
	ApplReqType_RETRANSMISSION_OF_APPLICATION_MESSAGES_FOR_THE_SPECIFIED_APPLICATIONS        ApplReqType = "0"
	ApplReqType_SUBSCRIPTION_TO_THE_SPECIFIED_APPLICATIONS                                   ApplReqType = "1"
	ApplReqType_REQUEST_FOR_THE_LAST_APPLLASTSEQNUM_PUBLISHED_FOR_THE_SPECIFIED_APPLICATIONS ApplReqType = "2"
	ApplReqType_REQUEST_VALID_SET_OF_APPLICATIONS                                            ApplReqType = "3"
	ApplReqType_UNSUBSCRIBE_TO_THE_SPECIFIED_APPLICATIONS                                    ApplReqType = "4"
	ApplReqType_CANCEL_RETRANSMISSION                                                        ApplReqType = "5"
	ApplReqType_CANCEL_RETRANSMISSION_AND_UNSUBSCRIBE_TO_THE_SPECIFIED_APPLICATIONS          ApplReqType = "6"
)

// ApplResponseError field enumeration values.
type ApplResponseError string

const (
	ApplResponseError_APPLICATION_DOES_NOT_EXIST           ApplResponseError = "0"
	ApplResponseError_MESSAGES_REQUESTED_ARE_NOT_AVAILABLE ApplResponseError = "1"
	ApplResponseError_USER_NOT_AUTHORIZED_FOR_APPLICATION  ApplResponseError = "2"
)

// ApplResponseType field enumeration values.
type ApplResponseType string

const (
	ApplResponseType_REQUEST_SUCCESSFULLY_PROCESSED ApplResponseType = "0"
	ApplResponseType_APPLICATION_DOES_NOT_EXIST     ApplResponseType = "1"
	ApplResponseType_MESSAGES_NOT_AVAILABLE         ApplResponseType = "2"
)

// ApplVerID field enumeration values.
type ApplVerID string

const (
	ApplVerID_FIX27     ApplVerID = "0"
	ApplVerID_FIX30     ApplVerID = "1"
	ApplVerID_FIXLATEST ApplVerID = "10"
	ApplVerID_FIX40     ApplVerID = "2"
	ApplVerID_FIX41     ApplVerID = "3"
	ApplVerID_FIX42     ApplVerID = "4"
	ApplVerID_FIX43     ApplVerID = "5"
	ApplVerID_FIX44     ApplVerID = "6"
	ApplVerID_FIX50     ApplVerID = "7"
	ApplVerID_FIX50SP1  ApplVerID = "8"
	ApplVerID_FIX50SP2  ApplVerID = "9"
)

// AsOfIndicator field enumeration values.
type AsOfIndicator string

const (
	AsOfIndicator_FALSE AsOfIndicator = "0"
	AsOfIndicator_TRUE  AsOfIndicator = "1"
)

// AssetClass field enumeration values.
type AssetClass string

const (
	AssetClass_INTEREST_RATE AssetClass = "1"
	AssetClass_LOAN_FACILITY AssetClass = "10"
	AssetClass_INDEX         AssetClass = "11"
	AssetClass_CURRENCY      AssetClass = "2"
	AssetClass_CREDIT        AssetClass = "3"
	AssetClass_EQUITY        AssetClass = "4"
	AssetClass_COMMODITY     AssetClass = "5"
	AssetClass_OTHER         AssetClass = "6"
	AssetClass_CASH          AssetClass = "7"
	AssetClass_DEBT          AssetClass = "8"
	AssetClass_FUND          AssetClass = "9"
)

// AssetGroup field enumeration values.
type AssetGroup string

const (
	AssetGroup_FINANCIALS              AssetGroup = "1"
	AssetGroup_COMMODITIES             AssetGroup = "2"
	AssetGroup_ALTERNATIVE_INVESTMENTS AssetGroup = "3"
)

// AssetSubClass field enumeration values.
type AssetSubClass string

const (
	AssetSubClass_SINGLE_CURRENCY               AssetSubClass = "1"
	AssetSubClass_PREFERRED                     AssetSubClass = "10"
	AssetSubClass_EQUITY_INDEX                  AssetSubClass = "11"
	AssetSubClass_EQUITY_BASKET                 AssetSubClass = "12"
	AssetSubClass_METALS                        AssetSubClass = "13"
	AssetSubClass_BULLION                       AssetSubClass = "14"
	AssetSubClass_ENERGY                        AssetSubClass = "15"
	AssetSubClass_COMMODITY_INDEX               AssetSubClass = "16"
	AssetSubClass_AGRICULTURAL                  AssetSubClass = "17"
	AssetSubClass_ENVIRONMENTAL                 AssetSubClass = "18"
	AssetSubClass_FREIGHT                       AssetSubClass = "19"
	AssetSubClass_CROSS_CURRENCY                AssetSubClass = "2"
	AssetSubClass_GOVERNMENT                    AssetSubClass = "20"
	AssetSubClass_AGENCY                        AssetSubClass = "21"
	AssetSubClass_CORPORATE                     AssetSubClass = "22"
	AssetSubClass_FINANCING                     AssetSubClass = "23"
	AssetSubClass_MONEY_MARKET                  AssetSubClass = "24"
	AssetSubClass_MORTGAGE                      AssetSubClass = "25"
	AssetSubClass_MUNICIPAL                     AssetSubClass = "26"
	AssetSubClass_MUTUAL_FUND                   AssetSubClass = "27"
	AssetSubClass_COLLECTIVE_INVESTMENT_VEHICLE AssetSubClass = "28"
	AssetSubClass_INVESTMENT_PROGRAM            AssetSubClass = "29"
	AssetSubClass_BASKET_FOR_MULTI_CURRENCY     AssetSubClass = "3"
	AssetSubClass_SPECIALIZED_ACCOUNT_PROGRAM   AssetSubClass = "30"
	AssetSubClass_TERM_LOAN                     AssetSubClass = "31"
	AssetSubClass_BRIDGE_LOAN                   AssetSubClass = "32"
	AssetSubClass_LETTER_OF_CREDIT              AssetSubClass = "33"
	AssetSubClass_DIVIDEND_INDEX                AssetSubClass = "34"
	AssetSubClass_STOCK_DIVIDEND                AssetSubClass = "35"
	AssetSubClass_EXCHANGE_TRADED_FUND          AssetSubClass = "36"
	AssetSubClass_VOLATILITY_INDEX              AssetSubClass = "37"
	AssetSubClass_FX_CROSS_RATES                AssetSubClass = "38"
	AssetSubClass_FX_EMERGING_MARKETS           AssetSubClass = "39"
	AssetSubClass_SINGLE_NAME                   AssetSubClass = "4"
	AssetSubClass_FX_MAJORS                     AssetSubClass = "40"
	AssetSubClass_FERTILIZER                    AssetSubClass = "41"
	AssetSubClass_INDUSTRIAL_PRODUCT            AssetSubClass = "42"
	AssetSubClass_INFLATION                     AssetSubClass = "43"
	AssetSubClass_PAPER                         AssetSubClass = "44"
	AssetSubClass_POLYPROPYLENE                 AssetSubClass = "45"
	AssetSubClass_OFFICIAL_ECONOMIC_STATISTICS  AssetSubClass = "46"
	AssetSubClass_OTHER_C10                     AssetSubClass = "47"
	AssetSubClass_OTHER                         AssetSubClass = "48"
	AssetSubClass_CREDIT_INDEX                  AssetSubClass = "5"
	AssetSubClass_INDEX_TRANCHE                 AssetSubClass = "6"
	AssetSubClass_CREDIT_BASKET                 AssetSubClass = "7"
	AssetSubClass_EXOTIC                        AssetSubClass = "8"
	AssetSubClass_COMMON                        AssetSubClass = "9"
)

// AssignmentMethod field enumeration values.
type AssignmentMethod string

const (
	AssignmentMethod_PRO_RATA AssignmentMethod = "P"
	AssignmentMethod_RANDOM   AssignmentMethod = "R"
)

// AttachmentEncodingType field enumeration values.
type AttachmentEncodingType string

const (
	AttachmentEncodingType_BASE64_ENCODING          AttachmentEncodingType = "0"
	AttachmentEncodingType_UNENCODED_BINARY_CONTENT AttachmentEncodingType = "1"
)

// AuctionInstruction field enumeration values.
type AuctionInstruction string

const (
	AuctionInstruction_AUTOMATIC_AUCTION_PERMITTED     AuctionInstruction = "0"
	AuctionInstruction_AUTOMATIC_AUCTION_NOT_PERMITTED AuctionInstruction = "1"
)

// AuctionType field enumeration values.
type AuctionType string

const (
	AuctionType_NONE                                       AuctionType = "0"
	AuctionType_BLOCK_ORDER_AUCTION                        AuctionType = "1"
	AuctionType_DIRECTED_ORDER_AUCTION                     AuctionType = "2"
	AuctionType_EXPOSURE_ORDER_AUCTION                     AuctionType = "3"
	AuctionType_FLASH_ORDER_AUCTION                        AuctionType = "4"
	AuctionType_FACILITATION_ORDER_AUCTION                 AuctionType = "5"
	AuctionType_SOLICITATION_ORDER_AUCTION                 AuctionType = "6"
	AuctionType_PRICE_IMPROVEMENT_MECHANISM                AuctionType = "7"
	AuctionType_DIRECTED_ORDER_PRICE_IMPROVEMENT_MECHANISM AuctionType = "8"
)

// AveragePriceType field enumeration values.
type AveragePriceType string

const (
	AveragePriceType_TIME_WEIGHTED_AVERAGE_PRICE     AveragePriceType = "0"
	AveragePriceType_VOLUME_WEIGHTED_AVERAGE_PRICE   AveragePriceType = "1"
	AveragePriceType_PERCENT_OF_VOLUME_AVERAGE_PRICE AveragePriceType = "2"
	AveragePriceType_LIMIT_ORDER_AVERAGE_PRICE       AveragePriceType = "3"
)

// AvgPxIndicator field enumeration values.
type AvgPxIndicator string

const (
	AvgPxIndicator_NO_AVERAGE_PRICING                                                     AvgPxIndicator = "0"
	AvgPxIndicator_TRADE_IS_PART_OF_AN_AVERAGE_PRICE_GROUP_IDENTIFIED_BY_THE_AVGPXGROUPID AvgPxIndicator = "1"
	AvgPxIndicator_LAST_TRADE_OF_THE_AVERAGE_PRICE_GROUP_IDENTIFIED_BY_THE_AVGPXGROUPID   AvgPxIndicator = "2"
	AvgPxIndicator_TRADE_IS_PART_OF_A_NOTIONAL_VALUE_AVERAGE_PRICE_GROUP                  AvgPxIndicator = "3"
	AvgPxIndicator_TRADE_IS_AVERAGE_PRICED                                                AvgPxIndicator = "4"
)

// BasisPxType field enumeration values.
type BasisPxType string

const (
	BasisPxType_CLOSING_PRICE_AT_MORNING_SESSION              BasisPxType = "2"
	BasisPxType_CLOSING_PRICE                                 BasisPxType = "3"
	BasisPxType_CURRENT_PRICE                                 BasisPxType = "4"
	BasisPxType_SQ                                            BasisPxType = "5"
	BasisPxType_VWAP_THROUGH_A_DAY                            BasisPxType = "6"
	BasisPxType_VWAP_THROUGH_A_MORNING_SESSION                BasisPxType = "7"
	BasisPxType_VWAP_THROUGH_AN_AFTERNOON_SESSION             BasisPxType = "8"
	BasisPxType_VWAP_THROUGH_A_DAY_EXCEPT_YORI                BasisPxType = "9"
	BasisPxType_VWAP_THROUGH_A_MORNING_SESSION_EXCEPT_YORI    BasisPxType = "A"
	BasisPxType_VWAP_THROUGH_AN_AFTERNOON_SESSION_EXCEPT_YORI BasisPxType = "B"
	BasisPxType_STRIKE                                        BasisPxType = "C"
	BasisPxType_OPEN                                          BasisPxType = "D"
	BasisPxType_OTHERS                                        BasisPxType = "Z"
)

// BatchProcessMode field enumeration values.
type BatchProcessMode string

const (
	BatchProcessMode_UPDATE_INCREMENTAL BatchProcessMode = "0"
	BatchProcessMode_SNAPSHOT           BatchProcessMode = "1"
)

// Benchmark field enumeration values.
type Benchmark string

const (
	Benchmark_CURVE    Benchmark = "1"
	Benchmark_5YR      Benchmark = "2"
	Benchmark_OLD5     Benchmark = "3"
	Benchmark_10YR     Benchmark = "4"
	Benchmark_OLD10    Benchmark = "5"
	Benchmark_30YR     Benchmark = "6"
	Benchmark_OLD30    Benchmark = "7"
	Benchmark_3MOLIBOR Benchmark = "8"
	Benchmark_6MOLIBOR Benchmark = "9"
)

// BenchmarkCurveName field enumeration values.
type BenchmarkCurveName string

const (
	BenchmarkCurveName_AUSTRALIAN_BANK_BILL_SWAP_RATE              BenchmarkCurveName = "AUBSW"
	BenchmarkCurveName_BUDAPEST_BANK_OFFERED_RATE                  BenchmarkCurveName = "BUBOR"
	BenchmarkCurveName_CANADIAN_DOLLAR_OFFERED_RATE                BenchmarkCurveName = "CDOR"
	BenchmarkCurveName_COPENHAGEN_INTERBANK_OFFERED_RATE           BenchmarkCurveName = "CIBOR"
	BenchmarkCurveName_EONIA                                       BenchmarkCurveName = "EONIA"
	BenchmarkCurveName_EURO_OVERNIGHT_INDEX_AVERAGE_SWAP_RATE      BenchmarkCurveName = "EONIASWAP"
	BenchmarkCurveName_EURO_SHORT_TERM_RATE                        BenchmarkCurveName = "ESTR"
	BenchmarkCurveName_EUREPO                                      BenchmarkCurveName = "EUREPO"
	BenchmarkCurveName_EURO_INTERBANK_OFFER_RATE                   BenchmarkCurveName = "EURIBOR"
	BenchmarkCurveName_EURO_DOLLAR_RATE                            BenchmarkCurveName = "EURODOLLAR"
	BenchmarkCurveName_EURO_SWISS_FRANC_RATE                       BenchmarkCurveName = "EUROSWISS"
	BenchmarkCurveName_EURIBOR                                     BenchmarkCurveName = "Euribor"
	BenchmarkCurveName_US_FEDERAL_RESERVE_FED_FUNDS_EFFECTIVE_RATE BenchmarkCurveName = "FEDEFF"
	BenchmarkCurveName_US_FED_FUNDS_TARGET_RATE                    BenchmarkCurveName = "FEDOPEN"
	BenchmarkCurveName_FUTURESWAP                                  BenchmarkCurveName = "FutureSWAP"
	BenchmarkCurveName_DTCC_GENERAL_COLLATERAL_FINANCE_REPO_INDEX  BenchmarkCurveName = "GCFREPO"
	BenchmarkCurveName_ICE_SWAP_RATE                               BenchmarkCurveName = "ISDAFIX"
	BenchmarkCurveName_JOHANNESBURG_INTERBANK_AGREED_RATE          BenchmarkCurveName = "JIBAR"
	BenchmarkCurveName_LIBID                                       BenchmarkCurveName = "LIBID"
	BenchmarkCurveName_LIBOR                                       BenchmarkCurveName = "LIBOR"
	BenchmarkCurveName_MOSCOW_PRIME_OFFERED_RATE                   BenchmarkCurveName = "MOSPRIM"
	BenchmarkCurveName_MUNIAAA                                     BenchmarkCurveName = "MuniAAA"
	BenchmarkCurveName_NIGERIA_THREE_MONTH_INTERBANK_RATE          BenchmarkCurveName = "NIBOR"
	BenchmarkCurveName_OTHER                                       BenchmarkCurveName = "OTHER"
	BenchmarkCurveName_CZECH_REPUBLIC_INTERBANK_OFFERED_RATE       BenchmarkCurveName = "PRIBOR"
	BenchmarkCurveName_PFANDBRIEFE                                 BenchmarkCurveName = "Pfandbriefe"
	BenchmarkCurveName_SECURED_OVERNIGHT_FINANCING_RATE            BenchmarkCurveName = "SOFR"
	BenchmarkCurveName_SONIA                                       BenchmarkCurveName = "SONIA"
	BenchmarkCurveName_STOCKHOLM_INTERBANK_OFFERED_RATE            BenchmarkCurveName = "STIBOR"
	BenchmarkCurveName_SWAP                                        BenchmarkCurveName = "SWAP"
	BenchmarkCurveName_BANK_OF_ISRAEL_INTERBANK_OFFERED_RATE       BenchmarkCurveName = "TELBOR"
	BenchmarkCurveName_TOKYO_INTERBANK_OFFERED_RATE                BenchmarkCurveName = "TIBOR"
	BenchmarkCurveName_TREASURY                                    BenchmarkCurveName = "Treasury"
	BenchmarkCurveName_WARSAW_INTERBANK_OFFERED_RATE               BenchmarkCurveName = "WIBOR"
)

// BidDescriptorType field enumeration values.
type BidDescriptorType string

const (
	BidDescriptorType_SECTOR  BidDescriptorType = "1"
	BidDescriptorType_COUNTRY BidDescriptorType = "2"
	BidDescriptorType_INDEX   BidDescriptorType = "3"
)

// BidRequestTransType field enumeration values.
type BidRequestTransType string

const (
	BidRequestTransType_CANCEL BidRequestTransType = "C"
	BidRequestTransType_NO     BidRequestTransType = "N"
)

// BidTradeType field enumeration values.
type BidTradeType string

const (
	BidTradeType_AGENCY           BidTradeType = "A"
	BidTradeType_VWAP_GUARANTEE   BidTradeType = "G"
	BidTradeType_GUARANTEED_CLOSE BidTradeType = "J"
	BidTradeType_RISK_TRADE       BidTradeType = "R"
)

// BidType field enumeration values.
type BidType string

const (
	BidType_NON_DISCLOSED_STYLE BidType = "1"
	BidType_DISCLOSED_SYTLE     BidType = "2"
	BidType_NO_BIDDING_PROCESS  BidType = "3"
)

// BlockTrdAllocIndicator field enumeration values.
type BlockTrdAllocIndicator string

const (
	BlockTrdAllocIndicator_BLOCK_TO_BE_ALLOCATED     BlockTrdAllocIndicator = "0"
	BlockTrdAllocIndicator_BLOCK_NOT_TO_BE_ALLOCATED BlockTrdAllocIndicator = "1"
	BlockTrdAllocIndicator_ALLOCATED_TRADE           BlockTrdAllocIndicator = "2"
)

// BookingType field enumeration values.
type BookingType string

const (
	BookingType_REGULAR_BOOKING   BookingType = "0"
	BookingType_CFD               BookingType = "1"
	BookingType_TOTAL_RETURN_SWAP BookingType = "2"
)

// BookingUnit field enumeration values.
type BookingUnit string

const (
	BookingUnit_EACH_PARTIAL_EXECUTION_IS_A_BOOKABLE_UNIT                               BookingUnit = "0"
	BookingUnit_AGGREGATE_PARTIAL_EXECUTIONS_ON_THIS_ORDER_AND_BOOK_ONE_TRADE_PER_ORDER BookingUnit = "1"
	BookingUnit_AGGREGATE_EXECUTIONS_FOR_THIS_SYMBOL_SIDE_AND_SETTLEMENT_DATE           BookingUnit = "2"
)

// BusinessDayConvention field enumeration values.
type BusinessDayConvention string

const (
	BusinessDayConvention_NOT_APPLICABLE         BusinessDayConvention = "0"
	BusinessDayConvention_NONE                   BusinessDayConvention = "1"
	BusinessDayConvention_FOLLOWING_DAY          BusinessDayConvention = "2"
	BusinessDayConvention_FLOATING_RATE_NOTE     BusinessDayConvention = "3"
	BusinessDayConvention_MODIFIED_FOLLOWING_DAY BusinessDayConvention = "4"
	BusinessDayConvention_PRECEDING_DAY          BusinessDayConvention = "5"
	BusinessDayConvention_MODIFIED_PRECEDING_DAY BusinessDayConvention = "6"
	BusinessDayConvention_NEAREST_DAY            BusinessDayConvention = "7"
)

// BusinessRejectReason field enumeration values.
type BusinessRejectReason string

const (
	BusinessRejectReason_OTHER                                                BusinessRejectReason = "0"
	BusinessRejectReason_UNKNOWN_ID                                           BusinessRejectReason = "1"
	BusinessRejectReason_THROTTLED_MESSAGES_REJECTED_ON_REQUEST               BusinessRejectReason = "10"
	BusinessRejectReason_INVALID_PRICE_INCREMENT                              BusinessRejectReason = "18"
	BusinessRejectReason_UNKNOWN_SECURITY                                     BusinessRejectReason = "2"
	BusinessRejectReason_UNSUPPORTED_MESSAGE_TYPE                             BusinessRejectReason = "3"
	BusinessRejectReason_APPLICATION_NOT_AVAILABLE                            BusinessRejectReason = "4"
	BusinessRejectReason_CONDITIONALLY_REQUIRED_FIELD_MISSING                 BusinessRejectReason = "5"
	BusinessRejectReason_NOT_AUTHORIZED                                       BusinessRejectReason = "6"
	BusinessRejectReason_DELIVERTO_FIRM_NOT_AVAILABLE_AT_THIS_TIME            BusinessRejectReason = "7"
	BusinessRejectReason_THROTTLE_LIMIT_EXCEEDED                              BusinessRejectReason = "8"
	BusinessRejectReason_THROTTLE_LIMIT_EXCEEDED_SESSION_WILL_BE_DISCONNECTED BusinessRejectReason = "9"
)

// CPProgram field enumeration values.
type CPProgram string

const (
	CPProgram_3_1   CPProgram = "1"
	CPProgram_4     CPProgram = "2"
	CPProgram_3_3   CPProgram = "3"
	CPProgram_3_4   CPProgram = "4"
	CPProgram_3_5   CPProgram = "5"
	CPProgram_3_6   CPProgram = "6"
	CPProgram_3_7   CPProgram = "7"
	CPProgram_3_8   CPProgram = "8"
	CPProgram_OTHER CPProgram = "99"
)

// CalculationMethod field enumeration values.
type CalculationMethod string

const (
	CalculationMethod_AUTOMATIC CalculationMethod = "0"
	CalculationMethod_MANUAL    CalculationMethod = "1"
)

// CancellationRights field enumeration values.
type CancellationRights string

const (
	CancellationRights_NO_M CancellationRights = "M"
	CancellationRights_NO_N CancellationRights = "N"
	CancellationRights_NO_O CancellationRights = "O"
	CancellationRights_YES  CancellationRights = "Y"
)

// CashMargin field enumeration values.
type CashMargin string

const (
	CashMargin_CASH         CashMargin = "1"
	CashMargin_MARGIN_OPEN  CashMargin = "2"
	CashMargin_MARGIN_CLOSE CashMargin = "3"
)

// CashSettlPriceDefault field enumeration values.
type CashSettlPriceDefault string

const (
	CashSettlPriceDefault_CLOSE CashSettlPriceDefault = "0"
	CashSettlPriceDefault_HEDGE CashSettlPriceDefault = "1"
)

// CashSettlQuoteMethod field enumeration values.
type CashSettlQuoteMethod string

const (
	CashSettlQuoteMethod_BID   CashSettlQuoteMethod = "0"
	CashSettlQuoteMethod_MID   CashSettlQuoteMethod = "1"
	CashSettlQuoteMethod_OFFER CashSettlQuoteMethod = "2"
)

// CashSettlValuationMethod field enumeration values.
type CashSettlValuationMethod string

const (
	CashSettlValuationMethod_MARKET                  CashSettlValuationMethod = "0"
	CashSettlValuationMethod_HIGHEST                 CashSettlValuationMethod = "1"
	CashSettlValuationMethod_AVERAGE_MARKET          CashSettlValuationMethod = "2"
	CashSettlValuationMethod_AVERAGE_HIGHEST         CashSettlValuationMethod = "3"
	CashSettlValuationMethod_BLENDED_MARKET          CashSettlValuationMethod = "4"
	CashSettlValuationMethod_BLENDED_HIGHEST         CashSettlValuationMethod = "5"
	CashSettlValuationMethod_AVERAGE_BLENDED_MARKET  CashSettlValuationMethod = "6"
	CashSettlValuationMethod_AVERAGE_BLENDED_HIGHEST CashSettlValuationMethod = "7"
)

// ClearedIndicator field enumeration values.
type ClearedIndicator string

const (
	ClearedIndicator_NOT_CLEARED ClearedIndicator = "0"
	ClearedIndicator_CLEARED     ClearedIndicator = "1"
	ClearedIndicator_SUBMITTED   ClearedIndicator = "2"
	ClearedIndicator_REJECTED    ClearedIndicator = "3"
)

// ClearingAccountType field enumeration values.
type ClearingAccountType string

const (
	ClearingAccountType_CUSTOMER     ClearingAccountType = "1"
	ClearingAccountType_FIRM         ClearingAccountType = "2"
	ClearingAccountType_MARKET_MAKER ClearingAccountType = "3"
)

// ClearingFeeIndicator field enumeration values.
type ClearingFeeIndicator string

const (
	ClearingFeeIndicator_1ST_YEAR_DELEGATE_TRADING_FOR_OWN_ACCOUNT                              ClearingFeeIndicator = "1"
	ClearingFeeIndicator_2ND_YEAR_DELEGATE_TRADING_FOR_OWN_ACCOUNT                              ClearingFeeIndicator = "2"
	ClearingFeeIndicator_3RD_YEAR_DELEGATE_TRADING_FOR_OWN_ACCOUNT                              ClearingFeeIndicator = "3"
	ClearingFeeIndicator_4TH_YEAR_DELEGATE_TRADING_FOR_OWN_ACCOUNT                              ClearingFeeIndicator = "4"
	ClearingFeeIndicator_5TH_YEAR_DELEGATE_TRADING_FOR_OWN_ACCOUNT                              ClearingFeeIndicator = "5"
	ClearingFeeIndicator_6TH_YEAR_DELEGATE_TRADING_FOR_OWN_ACCOUNT                              ClearingFeeIndicator = "9"
	ClearingFeeIndicator_CBOE_MEMBER                                                            ClearingFeeIndicator = "B"
	ClearingFeeIndicator_NON_MEMBER_AND_CUSTOMER                                                ClearingFeeIndicator = "C"
	ClearingFeeIndicator_EQUITY_MEMBER_AND_CLEARING_MEMBER                                      ClearingFeeIndicator = "E"
	ClearingFeeIndicator_FULL_AND_ASSOCIATE_MEMBER_TRADING_FOR_OWN_ACCOUNT_AND_AS_FLOOR_BROKERS ClearingFeeIndicator = "F"
	ClearingFeeIndicator_106H_AND_106J_FIRMS                                                    ClearingFeeIndicator = "H"
	ClearingFeeIndicator_GIM_IDEM_AND_COM_MEMBERSHIP_INTEREST_HOLDERS                           ClearingFeeIndicator = "I"
	ClearingFeeIndicator_LESSEE_106F_EMPLOYEES                                                  ClearingFeeIndicator = "L"
	ClearingFeeIndicator_ALL_OTHER_OWNERSHIP_TYPES                                              ClearingFeeIndicator = "M"
)

// ClearingInstruction field enumeration values.
type ClearingInstruction string

const (
	ClearingInstruction_PROCESS_NORMALLY                     ClearingInstruction = "0"
	ClearingInstruction_EXCLUDE_FROM_ALL_NETTING             ClearingInstruction = "1"
	ClearingInstruction_AUTOMATIC_GIVE_UP_MODE               ClearingInstruction = "10"
	ClearingInstruction_QUALIFIED_SERVICE_REPRESENTATIVE_QSR ClearingInstruction = "11"
	ClearingInstruction_CUSTOMER_TRADE                       ClearingInstruction = "12"
	ClearingInstruction_SELF_CLEARING                        ClearingInstruction = "13"
	ClearingInstruction_BUY_IN                               ClearingInstruction = "14"
	ClearingInstruction_BILATERAL_NETTING_ONLY               ClearingInstruction = "2"
	ClearingInstruction_EX_CLEARING                          ClearingInstruction = "3"
	ClearingInstruction_SPECIAL_TRADE                        ClearingInstruction = "4"
	ClearingInstruction_MULTILATERAL_NETTING                 ClearingInstruction = "5"
	ClearingInstruction_CLEAR_AGAINST_CENTRAL_COUNTERPARTY   ClearingInstruction = "6"
	ClearingInstruction_EXCLUDE_FROM_CENTRAL_COUNTERPARTY    ClearingInstruction = "7"
	ClearingInstruction_MANUAL_MODE                          ClearingInstruction = "8"
	ClearingInstruction_AUTOMATIC_POSTING_MODE               ClearingInstruction = "9"
)

// ClearingIntention field enumeration values.
type ClearingIntention string

const (
	ClearingIntention_DO_NOT_INTEND_TO_CLEAR ClearingIntention = "0"
	ClearingIntention_INTEND_TO_CLEAR        ClearingIntention = "1"
)

// ClearingRequirementException field enumeration values.
type ClearingRequirementException string

const (
	ClearingRequirementException_NO_EXCEPTION                 ClearingRequirementException = "0"
	ClearingRequirementException_EXCEPTION                    ClearingRequirementException = "1"
	ClearingRequirementException_END_USER_EXCEPTION           ClearingRequirementException = "2"
	ClearingRequirementException_INTER_AFFILIATE_EXCEPTION    ClearingRequirementException = "3"
	ClearingRequirementException_TREASURY_AFFILIATE_EXCEPTION ClearingRequirementException = "4"
	ClearingRequirementException_COOPERATIVE_EXCEPTION        ClearingRequirementException = "5"
)

// CollAction field enumeration values.
type CollAction string

const (
	CollAction_RETAIN CollAction = "0"
	CollAction_ADD    CollAction = "1"
	CollAction_REMOVE CollAction = "2"
)

// CollApplType field enumeration values.
type CollApplType string

const (
	CollApplType_SPECIFIC_DEPOSIT CollApplType = "0"
	CollApplType_GENERAL          CollApplType = "1"
)

// CollAsgnReason field enumeration values.
type CollAsgnReason string

const (
	CollAsgnReason_INITIAL                   CollAsgnReason = "0"
	CollAsgnReason_SCHEDULED                 CollAsgnReason = "1"
	CollAsgnReason_PLEDGE                    CollAsgnReason = "10"
	CollAsgnReason_TIME_WARNING              CollAsgnReason = "2"
	CollAsgnReason_MARGIN_DEFICIENCY         CollAsgnReason = "3"
	CollAsgnReason_MARGIN_EXCESS             CollAsgnReason = "4"
	CollAsgnReason_FORWARD_COLLATERAL_DEMAND CollAsgnReason = "5"
	CollAsgnReason_EVENT_OF_DEFAULT          CollAsgnReason = "6"
	CollAsgnReason_ADVERSE_TAX_EVENT         CollAsgnReason = "7"
	CollAsgnReason_TRANSFER_DEPOSIT          CollAsgnReason = "8"
	CollAsgnReason_TRANSFER_WITHDRAWAL       CollAsgnReason = "9"
)

// CollAsgnRejectReason field enumeration values.
type CollAsgnRejectReason string

const (
	CollAsgnRejectReason_UNKNOWN_DEAL                  CollAsgnRejectReason = "0"
	CollAsgnRejectReason_UNKNOWN_OR_INVALID_INSTRUMENT CollAsgnRejectReason = "1"
	CollAsgnRejectReason_UNAUTHORIZED_TRANSACTION      CollAsgnRejectReason = "2"
	CollAsgnRejectReason_INSUFFICIENT_COLLATERAL       CollAsgnRejectReason = "3"
	CollAsgnRejectReason_INVALID_TYPE_OF_COLLATERAL    CollAsgnRejectReason = "4"
	CollAsgnRejectReason_EXCESSIVE_SUBSTITUTION        CollAsgnRejectReason = "5"
	CollAsgnRejectReason_OTHER                         CollAsgnRejectReason = "99"
)

// CollAsgnRespType field enumeration values.
type CollAsgnRespType string

const (
	CollAsgnRespType_RECEIVED                           CollAsgnRespType = "0"
	CollAsgnRespType_ACCEPTED                           CollAsgnRespType = "1"
	CollAsgnRespType_DECLINED                           CollAsgnRespType = "2"
	CollAsgnRespType_REJECTED                           CollAsgnRespType = "3"
	CollAsgnRespType_TRANSACTION_PENDING                CollAsgnRespType = "4"
	CollAsgnRespType_TRANSACTION_COMPLETED_WITH_WARNING CollAsgnRespType = "5"
)

// CollAsgnTransType field enumeration values.
type CollAsgnTransType string

const (
	CollAsgnTransType_NEW     CollAsgnTransType = "0"
	CollAsgnTransType_REPLACE CollAsgnTransType = "1"
	CollAsgnTransType_CANCEL  CollAsgnTransType = "2"
	CollAsgnTransType_RELEASE CollAsgnTransType = "3"
	CollAsgnTransType_REVERSE CollAsgnTransType = "4"
)

// CollInquiryQualifier field enumeration values.
type CollInquiryQualifier string

const (
	CollInquiryQualifier_TRADE_DATE            CollInquiryQualifier = "0"
	CollInquiryQualifier_GC_INSTRUMENT         CollInquiryQualifier = "1"
	CollInquiryQualifier_COLLATERAL_INSTRUMENT CollInquiryQualifier = "2"
	CollInquiryQualifier_SUBSTITUTION_ELIGIBLE CollInquiryQualifier = "3"
	CollInquiryQualifier_NOT_ASSIGNED          CollInquiryQualifier = "4"
	CollInquiryQualifier_PARTIALLY_ASSIGNED    CollInquiryQualifier = "5"
	CollInquiryQualifier_FULLY_ASSIGNED        CollInquiryQualifier = "6"
	CollInquiryQualifier_OUTSTANDING_TRADES    CollInquiryQualifier = "7"
)

// CollInquiryResult field enumeration values.
type CollInquiryResult string

const (
	CollInquiryResult_SUCCESSFUL                                  CollInquiryResult = "0"
	CollInquiryResult_INVALID_OR_UNKNOWN_INSTRUMENT               CollInquiryResult = "1"
	CollInquiryResult_INVALID_OR_UNKNOWN_COLLATERAL_TYPE          CollInquiryResult = "2"
	CollInquiryResult_INVALID_PARTIES                             CollInquiryResult = "3"
	CollInquiryResult_INVALID_TRANSPORT_TYPE_REQUESTED            CollInquiryResult = "4"
	CollInquiryResult_INVALID_DESTINATION_REQUESTED               CollInquiryResult = "5"
	CollInquiryResult_NO_COLLATERAL_FOUND_FOR_THE_TRADE_SPECIFIED CollInquiryResult = "6"
	CollInquiryResult_NO_COLLATERAL_FOUND_FOR_THE_ORDER_SPECIFIED CollInquiryResult = "7"
	CollInquiryResult_COLLATERAL_INQUIRY_TYPE_NOT_SUPPORTED       CollInquiryResult = "8"
	CollInquiryResult_UNAUTHORIZED_FOR_COLLATERAL_INQUIRY         CollInquiryResult = "9"
	CollInquiryResult_OTHER                                       CollInquiryResult = "99"
)

// CollInquiryStatus field enumeration values.
type CollInquiryStatus string

const (
	CollInquiryStatus_ACCEPTED                CollInquiryStatus = "0"
	CollInquiryStatus_ACCEPTED_WITH_WARNINGS  CollInquiryStatus = "1"
	CollInquiryStatus_COMPLETED               CollInquiryStatus = "2"
	CollInquiryStatus_COMPLETED_WITH_WARNINGS CollInquiryStatus = "3"
	CollInquiryStatus_REJECTED                CollInquiryStatus = "4"
)

// CollRptRejectReason field enumeration values.
type CollRptRejectReason string

const (
	CollRptRejectReason_UNKNOWN_TRADE_OR_TRANSACTION               CollRptRejectReason = "0"
	CollRptRejectReason_UNKNOWN_OR_INVALID_INSTRUMENT              CollRptRejectReason = "1"
	CollRptRejectReason_UNKNOWN_OR_INVALID_COUNTERPARTY            CollRptRejectReason = "2"
	CollRptRejectReason_UNKNOWN_OR_INVALID_POSITION                CollRptRejectReason = "3"
	CollRptRejectReason_UNACCEPTABLE_OR_INVALID_TYPE_OF_COLLATERAL CollRptRejectReason = "4"
	CollRptRejectReason_OTHER                                      CollRptRejectReason = "99"
)

// CollRptStatus field enumeration values.
type CollRptStatus string

const (
	CollRptStatus_ACCEPTED CollRptStatus = "0"
	CollRptStatus_RECEIVED CollRptStatus = "1"
	CollRptStatus_REJECTED CollRptStatus = "2"
)

// CollStatus field enumeration values.
type CollStatus string

const (
	CollStatus_UNASSIGNED          CollStatus = "0"
	CollStatus_PARTIALLY_ASSIGNED  CollStatus = "1"
	CollStatus_ASSIGNMENT_PROPOSED CollStatus = "2"
	CollStatus_ASSIGNED            CollStatus = "3"
	CollStatus_CHALLENGED          CollStatus = "4"
	CollStatus_REUSED              CollStatus = "5"
)

// CollateralAmountType field enumeration values.
type CollateralAmountType string

const (
	CollateralAmountType_MARKET_VALUATION                                              CollateralAmountType = "0"
	CollateralAmountType_PORTFOLIO_VALUE_BEFORE_PROCESSING_PLEDGE_REQUEST              CollateralAmountType = "1"
	CollateralAmountType_VALUE_CONFIRMED_AS_LOCKED_UP_FOR_PROCESSING_A_PLEDGE_REQUEST  CollateralAmountType = "2"
	CollateralAmountType_CREDIT_VALUE_OF_COLLATERAL_AT_CCP_PROCESSING_A_PLEDGE_REQUEST CollateralAmountType = "3"
	CollateralAmountType_ADDITIONAL_COLLATERAL_VALUE                                   CollateralAmountType = "4"
	CollateralAmountType_ESTIMATED_MARKET_VALUATION                                    CollateralAmountType = "5"
)

// CollateralReinvestmentType field enumeration values.
type CollateralReinvestmentType string

const (
	CollateralReinvestmentType_MONEY_MARKET_FUND             CollateralReinvestmentType = "0"
	CollateralReinvestmentType_OTHER_COMINGLED_POOL          CollateralReinvestmentType = "1"
	CollateralReinvestmentType_REPO_MARKET                   CollateralReinvestmentType = "2"
	CollateralReinvestmentType_DIRECT_PURCHASE_OF_SECURITIES CollateralReinvestmentType = "3"
	CollateralReinvestmentType_OTHER_INVESTMENTS             CollateralReinvestmentType = "4"
)

// CommType field enumeration values.
type CommType string

const (
	CommType_AMOUNT_PER_UNIT                        CommType = "1"
	CommType_PERCENT                                CommType = "2"
	CommType_ABSOLUTE                               CommType = "3"
	CommType_PERCENTAGE_WAIVED_CASH_DISCOUNT_BASIS  CommType = "4"
	CommType_PERCENTAGE_WAIVED_ENHANCED_UNITS_BASIS CommType = "5"
	CommType_POINTS_PER_BOND_OR_CONTRACT            CommType = "6"
	CommType_BASIS_POINTS                           CommType = "7"
	CommType_AMOUNT_PER_CONTRACT                    CommType = "8"
)

// CommissionAmountSubType field enumeration values.
type CommissionAmountSubType string

const (
	CommissionAmountSubType_RESEARCH_PAYMENT_ACCOUNT       CommissionAmountSubType = "0"
	CommissionAmountSubType_COMISSION_SHARING_AGREEMENT    CommissionAmountSubType = "1"
	CommissionAmountSubType_OTHER_TYPE_OF_RESEARCH_PAYMENT CommissionAmountSubType = "2"
)

// CommissionAmountType field enumeration values.
type CommissionAmountType string

const (
	CommissionAmountType_UNSPECIFIED      CommissionAmountType = "0"
	CommissionAmountType_ACCEPTANCE       CommissionAmountType = "1"
	CommissionAmountType_BROKER           CommissionAmountType = "2"
	CommissionAmountType_CLEARING_BROKER  CommissionAmountType = "3"
	CommissionAmountType_RETAIL           CommissionAmountType = "4"
	CommissionAmountType_SALES_COMMISSION CommissionAmountType = "5"
	CommissionAmountType_LOCAL_COMMISSION CommissionAmountType = "6"
	CommissionAmountType_RESEARCH_PAYMENT CommissionAmountType = "7"
)

// CommodityFinalPriceType field enumeration values.
type CommodityFinalPriceType string

const (
	CommodityFinalPriceType_ARGUS_MCCLOSKEY CommodityFinalPriceType = "0"
	CommodityFinalPriceType_BALTIC          CommodityFinalPriceType = "1"
	CommodityFinalPriceType_EXCHANGE        CommodityFinalPriceType = "2"
	CommodityFinalPriceType_GLOBAL_COAL     CommodityFinalPriceType = "3"
	CommodityFinalPriceType_IHS_MCCLOSKEY   CommodityFinalPriceType = "4"
	CommodityFinalPriceType_PLATTS          CommodityFinalPriceType = "5"
	CommodityFinalPriceType_OTHER           CommodityFinalPriceType = "99"
)

// ComplexEventCondition field enumeration values.
type ComplexEventCondition string

const (
	ComplexEventCondition_AND ComplexEventCondition = "1"
	ComplexEventCondition_OR  ComplexEventCondition = "2"
)

// ComplexEventCreditEventNotifyingParty field enumeration values.
type ComplexEventCreditEventNotifyingParty string

const (
	ComplexEventCreditEventNotifyingParty_SELLER_NOTIFIES          ComplexEventCreditEventNotifyingParty = "0"
	ComplexEventCreditEventNotifyingParty_BUYER_NOTIFIES           ComplexEventCreditEventNotifyingParty = "1"
	ComplexEventCreditEventNotifyingParty_SELLER_OR_BUYER_NOTIFIES ComplexEventCreditEventNotifyingParty = "2"
)

// ComplexEventDateOffsetDayType field enumeration values.
type ComplexEventDateOffsetDayType string

const (
	ComplexEventDateOffsetDayType_BUSINESS              ComplexEventDateOffsetDayType = "0"
	ComplexEventDateOffsetDayType_CALENDAR              ComplexEventDateOffsetDayType = "1"
	ComplexEventDateOffsetDayType_COMMODITY_BUSINESS    ComplexEventDateOffsetDayType = "2"
	ComplexEventDateOffsetDayType_CURRENCY_BUSINESS     ComplexEventDateOffsetDayType = "3"
	ComplexEventDateOffsetDayType_EXCHANGE_BUSINESS     ComplexEventDateOffsetDayType = "4"
	ComplexEventDateOffsetDayType_SCHEDULED_TRADING_DAY ComplexEventDateOffsetDayType = "5"
)

// ComplexEventPVFinalPriceElectionFallback field enumeration values.
type ComplexEventPVFinalPriceElectionFallback string

const (
	ComplexEventPVFinalPriceElectionFallback_CLOSE          ComplexEventPVFinalPriceElectionFallback = "0"
	ComplexEventPVFinalPriceElectionFallback_HEDGE_ELECTION ComplexEventPVFinalPriceElectionFallback = "1"
)

// ComplexEventPeriodType field enumeration values.
type ComplexEventPeriodType string

const (
	ComplexEventPeriodType_ASIAN_OUT     ComplexEventPeriodType = "0"
	ComplexEventPeriodType_ASIAN_IN      ComplexEventPeriodType = "1"
	ComplexEventPeriodType_BARRIER_CAP   ComplexEventPeriodType = "2"
	ComplexEventPeriodType_BARRIER_FLOOR ComplexEventPeriodType = "3"
	ComplexEventPeriodType_KNOCK_OUT     ComplexEventPeriodType = "4"
	ComplexEventPeriodType_KNOCK_IN      ComplexEventPeriodType = "5"
)

// ComplexEventPriceBoundaryMethod field enumeration values.
type ComplexEventPriceBoundaryMethod string

const (
	ComplexEventPriceBoundaryMethod_LESS_THAN_COMPLEXEVENTPRICE                ComplexEventPriceBoundaryMethod = "1"
	ComplexEventPriceBoundaryMethod_LESS_THAN_OR_EQUAL_TO_COMPLEXEVENTPRICE    ComplexEventPriceBoundaryMethod = "2"
	ComplexEventPriceBoundaryMethod_EQUAL_TO_COMPLEXEVENTPRICE                 ComplexEventPriceBoundaryMethod = "3"
	ComplexEventPriceBoundaryMethod_GREATER_THAN_OR_EQUAL_TO_COMPLEXEVENTPRICE ComplexEventPriceBoundaryMethod = "4"
	ComplexEventPriceBoundaryMethod_GREATER_THAN_COMPLEXEVENTPRICE             ComplexEventPriceBoundaryMethod = "5"
)

// ComplexEventPriceTimeType field enumeration values.
type ComplexEventPriceTimeType string

const (
	ComplexEventPriceTimeType_EXPIRATION                          ComplexEventPriceTimeType = "1"
	ComplexEventPriceTimeType_IMMEDIATE                           ComplexEventPriceTimeType = "2"
	ComplexEventPriceTimeType_SPECIFIED_DATE_TIME                 ComplexEventPriceTimeType = "3"
	ComplexEventPriceTimeType_CLOSE                               ComplexEventPriceTimeType = "4"
	ComplexEventPriceTimeType_OPEN                                ComplexEventPriceTimeType = "5"
	ComplexEventPriceTimeType_OFFICIAL_SETTLEMENT_PRICE           ComplexEventPriceTimeType = "6"
	ComplexEventPriceTimeType_DERIVATIVES_CLOSE                   ComplexEventPriceTimeType = "7"
	ComplexEventPriceTimeType_AS_SPECIFIED_IN_MASTER_CONFIRMATION ComplexEventPriceTimeType = "8"
)

// ComplexEventQuoteBasis field enumeration values.
type ComplexEventQuoteBasis string

const (
	ComplexEventQuoteBasis_CURRENCY_1_PER_CURRENCY_2 ComplexEventQuoteBasis = "0"
	ComplexEventQuoteBasis_CURRENCY_2_PER_CURRENCY_1 ComplexEventQuoteBasis = "1"
)

// ComplexEventType field enumeration values.
type ComplexEventType string

const (
	ComplexEventType_CAPPED                          ComplexEventType = "1"
	ComplexEventType_ONE_TOUCH                       ComplexEventType = "10"
	ComplexEventType_NO_TOUCH                        ComplexEventType = "11"
	ComplexEventType_DOUBLE_ONE_TOUCH                ComplexEventType = "12"
	ComplexEventType_DOUBLE_NO_TOUCH                 ComplexEventType = "13"
	ComplexEventType_FOREIGN_EXCHANGE_COMPOSITE      ComplexEventType = "14"
	ComplexEventType_FOREIGN_EXCHANGE_QUANTO         ComplexEventType = "15"
	ComplexEventType_FOREIGN_EXCHANGE_CROSS_CURRENCY ComplexEventType = "16"
	ComplexEventType_STRIKE_SPREAD                   ComplexEventType = "17"
	ComplexEventType_CALENDAR_SPREAD                 ComplexEventType = "18"
	ComplexEventType_PRICE_OBSERVATION               ComplexEventType = "19"
	ComplexEventType_TRIGGER                         ComplexEventType = "2"
	ComplexEventType_PASS_THROUGH                    ComplexEventType = "20"
	ComplexEventType_STRIKE_SCHEDULE                 ComplexEventType = "21"
	ComplexEventType_EQUITY_VALUATION                ComplexEventType = "22"
	ComplexEventType_DIVIDEND_VALUATION              ComplexEventType = "23"
	ComplexEventType_KNOCK_IN_UP                     ComplexEventType = "3"
	ComplexEventType_KNOCK_IN_DOWN                   ComplexEventType = "4"
	ComplexEventType_KNOCK_OUT_UP                    ComplexEventType = "5"
	ComplexEventType_KNOCK_OUT_DOWN                  ComplexEventType = "6"
	ComplexEventType_UNDERLYING                      ComplexEventType = "7"
	ComplexEventType_RESET_BARRIER                   ComplexEventType = "8"
	ComplexEventType_ROLLING_BARRIER                 ComplexEventType = "9"
)

// ComplexOptPayoutTime field enumeration values.
type ComplexOptPayoutTime string

const (
	ComplexOptPayoutTime_CLOSE                               ComplexOptPayoutTime = "0"
	ComplexOptPayoutTime_OPEN                                ComplexOptPayoutTime = "1"
	ComplexOptPayoutTime_OFFICIAL_SETTLEMENT                 ComplexOptPayoutTime = "2"
	ComplexOptPayoutTime_VALUATION_TIME                      ComplexOptPayoutTime = "3"
	ComplexOptPayoutTime_EXCHANGE_SETTLEMENT_TIME            ComplexOptPayoutTime = "4"
	ComplexOptPayoutTime_DERIVATIVES_CLOSE                   ComplexOptPayoutTime = "5"
	ComplexOptPayoutTime_AS_SPECIFIED_IN_MASTER_CONFIRMATION ComplexOptPayoutTime = "6"
)

// ConfirmRejReason field enumeration values.
type ConfirmRejReason string

const (
	ConfirmRejReason_INCORRECT_OR_MISSING_ACCOUNT                          ConfirmRejReason = "1"
	ConfirmRejReason_INCORRECT_OR_MISSING_FUND_ID_OR_FUND_NAME             ConfirmRejReason = "10"
	ConfirmRejReason_INCORRECT_OR_MISSING_QUANTITY                         ConfirmRejReason = "11"
	ConfirmRejReason_INCORRECT_OR_MISSING_FEES                             ConfirmRejReason = "12"
	ConfirmRejReason_INCORRECT_OR_MISSING_TAX                              ConfirmRejReason = "13"
	ConfirmRejReason_INCORRECT_OR_MISSING_PARTY                            ConfirmRejReason = "14"
	ConfirmRejReason_INCORRECT_OR_MISSING_SIDE                             ConfirmRejReason = "15"
	ConfirmRejReason_INCORRECT_OR_MISSING_NET_MONEY                        ConfirmRejReason = "16"
	ConfirmRejReason_INCORRECT_OR_MISSING_TRADE_DATE                       ConfirmRejReason = "17"
	ConfirmRejReason_INCORRECT_OR_MISSING_SETTLEMENT_CURRENCY_INSTRUCTIONS ConfirmRejReason = "18"
	ConfirmRejReason_INCORRECT_OR_MISSING_CAPACITY                         ConfirmRejReason = "19"
	ConfirmRejReason_INCORRECT_OR_MISSING_SETTLEMENT_INSTRUCTIONS          ConfirmRejReason = "2"
	ConfirmRejReason_UNKNOWN_OR_MISSING_INDIVIDUALALLOCID                  ConfirmRejReason = "3"
	ConfirmRejReason_TRANSACTION_NOT_RECOGNIZED                            ConfirmRejReason = "4"
	ConfirmRejReason_DUPLICATE_TRANSACTION                                 ConfirmRejReason = "5"
	ConfirmRejReason_INCORRECT_OR_MISSING_INSTRUMENT                       ConfirmRejReason = "6"
	ConfirmRejReason_INCORRECT_OR_MISSING_PRICE                            ConfirmRejReason = "7"
	ConfirmRejReason_INCORRECT_OR_MISSING_COMMISSION                       ConfirmRejReason = "8"
	ConfirmRejReason_INCORRECT_OR_MISSING_SETTLEMENT_DATE                  ConfirmRejReason = "9"
	ConfirmRejReason_OTHER                                                 ConfirmRejReason = "99"
)

// ConfirmStatus field enumeration values.
type ConfirmStatus string

const (
	ConfirmStatus_RECEIVED                        ConfirmStatus = "1"
	ConfirmStatus_MISMATCHED_ACCOUNT              ConfirmStatus = "2"
	ConfirmStatus_MISSING_SETTLEMENT_INSTRUCTIONS ConfirmStatus = "3"
	ConfirmStatus_CONFIRMED                       ConfirmStatus = "4"
	ConfirmStatus_REQUEST_REJECTED                ConfirmStatus = "5"
)

// ConfirmTransType field enumeration values.
type ConfirmTransType string

const (
	ConfirmTransType_NEW     ConfirmTransType = "0"
	ConfirmTransType_REPLACE ConfirmTransType = "1"
	ConfirmTransType_CANCEL  ConfirmTransType = "2"
)

// ConfirmType field enumeration values.
type ConfirmType string

const (
	ConfirmType_STATUS                        ConfirmType = "1"
	ConfirmType_CONFIRMATION                  ConfirmType = "2"
	ConfirmType_CONFIRMATION_REQUEST_REJECTED ConfirmType = "3"
)

// ConfirmationMethod field enumeration values.
type ConfirmationMethod string

const (
	ConfirmationMethod_NON_ELECTRONIC ConfirmationMethod = "0"
	ConfirmationMethod_ELECTRONIC     ConfirmationMethod = "1"
	ConfirmationMethod_UNCONFIRMED    ConfirmationMethod = "2"
)

// ContAmtType field enumeration values.
type ContAmtType string

const (
	ContAmtType_COMMISSION_AMOUNT                       ContAmtType = "1"
	ContAmtType_EXIT_CHARGE_PERCENT                     ContAmtType = "10"
	ContAmtType_FUND_BASED_RENEWAL_COMMISSION_PERCENT   ContAmtType = "11"
	ContAmtType_PROJECTED_FUND_VALUE                    ContAmtType = "12"
	ContAmtType_FUND_BASED_RENEWAL_COMMISSION_AMOUNT_13 ContAmtType = "13"
	ContAmtType_FUND_BASED_RENEWAL_COMMISSION_AMOUNT_14 ContAmtType = "14"
	ContAmtType_NET_SETTLEMENT_AMOUNT                   ContAmtType = "15"
	ContAmtType_COMMISSION_PERCENT                      ContAmtType = "2"
	ContAmtType_INITIAL_CHARGE_AMOUNT                   ContAmtType = "3"
	ContAmtType_INITIAL_CHARGE_PERCENT                  ContAmtType = "4"
	ContAmtType_DISCOUNT_AMOUNT                         ContAmtType = "5"
	ContAmtType_DISCOUNT_PERCENT                        ContAmtType = "6"
	ContAmtType_DILUTION_LEVY_AMOUNT                    ContAmtType = "7"
	ContAmtType_DILUTION_LEVY_PERCENT                   ContAmtType = "8"
	ContAmtType_EXIT_CHARGE_AMOUNT                      ContAmtType = "9"
)

// ContingencyType field enumeration values.
type ContingencyType string

const (
	ContingencyType_ONE_CANCELS_THE_OTHER   ContingencyType = "1"
	ContingencyType_ONE_TRIGGERS_THE_OTHER  ContingencyType = "2"
	ContingencyType_ONE_UPDATES_THE_OTHER_3 ContingencyType = "3"
	ContingencyType_ONE_UPDATES_THE_OTHER_4 ContingencyType = "4"
	ContingencyType_BID_AND_OFFER           ContingencyType = "5"
	ContingencyType_BID_AND_OFFER_OCO       ContingencyType = "6"
)

// ContractMultiplierUnit field enumeration values.
type ContractMultiplierUnit string

const (
	ContractMultiplierUnit_SHARES ContractMultiplierUnit = "0"
	ContractMultiplierUnit_HOURS  ContractMultiplierUnit = "1"
	ContractMultiplierUnit_DAYS   ContractMultiplierUnit = "2"
)

// ContractRefPosType field enumeration values.
type ContractRefPosType string

const (
	ContractRefPosType_TWO_COMPONENT_INTERCOMMODITY_SPREAD ContractRefPosType = "0"
	ContractRefPosType_INDEX_OR_BASKET                     ContractRefPosType = "1"
	ContractRefPosType_TWO_COMPONENT_LOCATIONAL_BASIS      ContractRefPosType = "2"
	ContractRefPosType_OTHER                               ContractRefPosType = "99"
)

// CorporateAction field enumeration values.
type CorporateAction string

const (
	CorporateAction_EX_DIVIDEND                  CorporateAction = "A"
	CorporateAction_EX_DISTRIBUTION              CorporateAction = "B"
	CorporateAction_EX_RIGHTS                    CorporateAction = "C"
	CorporateAction_NEW                          CorporateAction = "D"
	CorporateAction_EX_INTEREST                  CorporateAction = "E"
	CorporateAction_CASH_DIVIDEND                CorporateAction = "F"
	CorporateAction_STOCK_DIVIDEND               CorporateAction = "G"
	CorporateAction_NON_INTEGER_STOCK_SPLIT      CorporateAction = "H"
	CorporateAction_REVERSE_STOCK_SPLIT          CorporateAction = "I"
	CorporateAction_STANDARD_INTEGER_STOCK_SPLIT CorporateAction = "J"
	CorporateAction_POSITION_CONSOLIDATION       CorporateAction = "K"
	CorporateAction_LIQUIDATION_REORGANIZATION   CorporateAction = "L"
	CorporateAction_MERGER_REORGANIZATION        CorporateAction = "M"
	CorporateAction_RIGHTS_OFFERING              CorporateAction = "N"
	CorporateAction_SHAREHOLDER_MEETING          CorporateAction = "O"
	CorporateAction_SPINOFF                      CorporateAction = "P"
	CorporateAction_TENDER_OFFER                 CorporateAction = "Q"
	CorporateAction_WARRANT                      CorporateAction = "R"
	CorporateAction_SPECIAL_ACTION               CorporateAction = "S"
	CorporateAction_SYMBOL_CONVERSION            CorporateAction = "T"
	CorporateAction_CUSIP                        CorporateAction = "U"
	CorporateAction_LEAP_ROLLOVER                CorporateAction = "V"
	CorporateAction_SUCCESSION_EVENT             CorporateAction = "W"
)

// CouponDayCount field enumeration values.
type CouponDayCount string

const (
	CouponDayCount_1_1          CouponDayCount = "0"
	CouponDayCount_30_360_1     CouponDayCount = "1"
	CouponDayCount_ACT_ACT_10   CouponDayCount = "10"
	CouponDayCount_ACT_ACT_11   CouponDayCount = "11"
	CouponDayCount_BUS_252      CouponDayCount = "12"
	CouponDayCount_30E_PLUS_360 CouponDayCount = "13"
	CouponDayCount_ACT_365L     CouponDayCount = "14"
	CouponDayCount_NL365        CouponDayCount = "15"
	CouponDayCount_NL360        CouponDayCount = "16"
	CouponDayCount_ACT_364      CouponDayCount = "17"
	CouponDayCount_30_365       CouponDayCount = "18"
	CouponDayCount_30_ACTUAL    CouponDayCount = "19"
	CouponDayCount_30_360_2     CouponDayCount = "2"
	CouponDayCount_30_360_20    CouponDayCount = "20"
	CouponDayCount_30E2_360     CouponDayCount = "21"
	CouponDayCount_30E3_360     CouponDayCount = "22"
	CouponDayCount_30_360M      CouponDayCount = "3"
	CouponDayCount_30E_360_4    CouponDayCount = "4"
	CouponDayCount_30E_360_5    CouponDayCount = "5"
	CouponDayCount_ACT_360      CouponDayCount = "6"
	CouponDayCount_ACT_365      CouponDayCount = "7"
	CouponDayCount_ACT_ACT_8    CouponDayCount = "8"
	CouponDayCount_ACT_ACT_9    CouponDayCount = "9"
	CouponDayCount_OTHER        CouponDayCount = "99"
)

// CouponFrequencyUnit field enumeration values.
type CouponFrequencyUnit string

const (
	CouponFrequencyUnit_DAY    CouponFrequencyUnit = "D"
	CouponFrequencyUnit_HOUR   CouponFrequencyUnit = "H"
	CouponFrequencyUnit_MINUTE CouponFrequencyUnit = "Min"
	CouponFrequencyUnit_MONTH  CouponFrequencyUnit = "Mo"
	CouponFrequencyUnit_SECOND CouponFrequencyUnit = "S"
	CouponFrequencyUnit_TERM   CouponFrequencyUnit = "T"
	CouponFrequencyUnit_WEEK   CouponFrequencyUnit = "Wk"
	CouponFrequencyUnit_YEAR   CouponFrequencyUnit = "Yr"
)

// CouponType field enumeration values.
type CouponType string

const (
	CouponType_ZERO          CouponType = "0"
	CouponType_FIXED_RATE    CouponType = "1"
	CouponType_FLOATING_RATE CouponType = "2"
	CouponType_STRUCTURED    CouponType = "3"
)

// CoveredOrUncovered field enumeration values.
type CoveredOrUncovered string

const (
	CoveredOrUncovered_COVERED   CoveredOrUncovered = "0"
	CoveredOrUncovered_UNCOVERED CoveredOrUncovered = "1"
)

// CrossPrioritization field enumeration values.
type CrossPrioritization string

const (
	CrossPrioritization_NONE                     CrossPrioritization = "0"
	CrossPrioritization_BUY_SIDE_IS_PRIORITIZED  CrossPrioritization = "1"
	CrossPrioritization_SELL_SIDE_IS_PRIORITIZED CrossPrioritization = "2"
)

// CrossType field enumeration values.
type CrossType string

const (
	CrossType_ALL_OR_NONE_CROSS             CrossType = "1"
	CrossType_IMMEDIATE_OR_CANCEL_CROSS     CrossType = "2"
	CrossType_ONE_SIDED_CROSS               CrossType = "3"
	CrossType_CROSS_EXECUTED_AGAINST_BOOK   CrossType = "4"
	CrossType_BASIS_CROSS                   CrossType = "5"
	CrossType_CONTINGENT_CROSS              CrossType = "6"
	CrossType_VOLUME_WEIGHTED_AVERAGE_PRICE CrossType = "7"
	CrossType_SPECIAL_TRADING_SESSION_CROSS CrossType = "8"
	CrossType_CUSTOMER_TO_CUSTOMER_CROSS    CrossType = "9"
)

// CrossedIndicator field enumeration values.
type CrossedIndicator string

const (
	CrossedIndicator_NO_CROSS       CrossedIndicator = "0"
	CrossedIndicator_CROSS_REJECTED CrossedIndicator = "1"
	CrossedIndicator_CROSS_ACCEPTED CrossedIndicator = "2"
)

// CustOrderCapacity field enumeration values.
type CustOrderCapacity string

const (
	CustOrderCapacity_MEMBER_TRADING_FOR_THEIR_OWN_ACCOUNT              CustOrderCapacity = "1"
	CustOrderCapacity_CLEARING_FIRM_TRADING_FOR_ITS_PROPRIETARY_ACCOUNT CustOrderCapacity = "2"
	CustOrderCapacity_MEMBER_TRADING_FOR_ANOTHER_MEMBER                 CustOrderCapacity = "3"
	CustOrderCapacity_ALL_OTHER                                         CustOrderCapacity = "4"
	CustOrderCapacity_RETAIL_CUSTOMER                                   CustOrderCapacity = "5"
)

// CustOrderHandlingInst field enumeration values.
type CustOrderHandlingInst string

const (
	CustOrderHandlingInst_PHONE_SIMPLE                                          CustOrderHandlingInst = "A"
	CustOrderHandlingInst_ADD_ON_ORDER                                          CustOrderHandlingInst = "ADD"
	CustOrderHandlingInst_ALL_OR_NONE                                           CustOrderHandlingInst = "AON"
	CustOrderHandlingInst_PHONE_COMPLEX                                         CustOrderHandlingInst = "B"
	CustOrderHandlingInst_FCM_PROVIDED_SCREEN                                   CustOrderHandlingInst = "C"
	CustOrderHandlingInst_CONDITIONAL_ORDER                                     CustOrderHandlingInst = "CND"
	CustOrderHandlingInst_CASH_NOT_HELD                                         CustOrderHandlingInst = "CNH"
	CustOrderHandlingInst_DELIVERY_INSTRUCTIONS_CSH                             CustOrderHandlingInst = "CSH"
	CustOrderHandlingInst_OTHER_PROVIDED_SCREEN                                 CustOrderHandlingInst = "D"
	CustOrderHandlingInst_DIRECTED_ORDER                                        CustOrderHandlingInst = "DIR"
	CustOrderHandlingInst_DISCRETIONARY_LIMIT_ORDER                             CustOrderHandlingInst = "DLO"
	CustOrderHandlingInst_CLIENT_PROVIDED_PLATFORM_CONTROLLED_BY_FCM            CustOrderHandlingInst = "E"
	CustOrderHandlingInst_EXCHANGE_FOR_PHYSICAL_TRANSACTION                     CustOrderHandlingInst = "E.W"
	CustOrderHandlingInst_CLIENT_PROVIDED_PLATFORM_DIRECT_TO_EXCHANGE           CustOrderHandlingInst = "F"
	CustOrderHandlingInst_STAY_ON_OFFERSIDE                                     CustOrderHandlingInst = "F0"
	CustOrderHandlingInst_GO_ALONG                                              CustOrderHandlingInst = "F3"
	CustOrderHandlingInst_PARTICIPATE_DO_NOT_INITIATE                           CustOrderHandlingInst = "F6"
	CustOrderHandlingInst_STRICT_SCALE                                          CustOrderHandlingInst = "F7"
	CustOrderHandlingInst_TRY_TO_SCALE                                          CustOrderHandlingInst = "F8"
	CustOrderHandlingInst_STAY_ON_BIDSIDE                                       CustOrderHandlingInst = "F9"
	CustOrderHandlingInst_NO_CROSS                                              CustOrderHandlingInst = "FA"
	CustOrderHandlingInst_OK_TO_CROSS                                           CustOrderHandlingInst = "FB"
	CustOrderHandlingInst_CALL_FIRST                                            CustOrderHandlingInst = "FC"
	CustOrderHandlingInst_PERCENT_OF_VOLUME                                     CustOrderHandlingInst = "FD"
	CustOrderHandlingInst_REINSTATE_ON_SYSTEM_FAILURE                           CustOrderHandlingInst = "FH"
	CustOrderHandlingInst_INSTITUTION_ONLY                                      CustOrderHandlingInst = "FI"
	CustOrderHandlingInst_REINSTATE_ON_TRADING_HALT                             CustOrderHandlingInst = "FJ"
	CustOrderHandlingInst_CANCEL_ON_TRADING_HALF                                CustOrderHandlingInst = "FK"
	CustOrderHandlingInst_LAST_PEG                                              CustOrderHandlingInst = "FL"
	CustOrderHandlingInst_MID_PRICE_PEG                                         CustOrderHandlingInst = "FM"
	CustOrderHandlingInst_NON_NEGOTIABLE                                        CustOrderHandlingInst = "FN"
	CustOrderHandlingInst_OPENING_PEG                                           CustOrderHandlingInst = "FO"
	CustOrderHandlingInst_FILL_OR_KILL                                          CustOrderHandlingInst = "FOK"
	CustOrderHandlingInst_MARKET_PEG                                            CustOrderHandlingInst = "FP"
	CustOrderHandlingInst_CANCEL_ON_SYSTEM_FAILURE                              CustOrderHandlingInst = "FQ"
	CustOrderHandlingInst_PRIMARY_PEG                                           CustOrderHandlingInst = "FR"
	CustOrderHandlingInst_SUSPEND                                               CustOrderHandlingInst = "FS"
	CustOrderHandlingInst_FIXED_PEG_TO_LOCAL_BEST_BID_OR_OFFER_AT_TIME_OF_ORDER CustOrderHandlingInst = "FT"
	CustOrderHandlingInst_PEG_TO_VWAP                                           CustOrderHandlingInst = "FW"
	CustOrderHandlingInst_TRADE_ALONG                                           CustOrderHandlingInst = "FX"
	CustOrderHandlingInst_TRY_TO_STOP                                           CustOrderHandlingInst = "FY"
	CustOrderHandlingInst_CANCEL_IF_NOT_BEST                                    CustOrderHandlingInst = "FZ"
	CustOrderHandlingInst_STRICT_LIMIT                                          CustOrderHandlingInst = "Fb"
	CustOrderHandlingInst_IGNORE_PRICE_VALIDITY_CHECKS                          CustOrderHandlingInst = "Fc"
	CustOrderHandlingInst_PEG_TO_LIMIT_PRICE                                    CustOrderHandlingInst = "Fd"
	CustOrderHandlingInst_WORK_TO_TARGET_STRATEGY                               CustOrderHandlingInst = "Fe"
	CustOrderHandlingInst_G_ORDER                                               CustOrderHandlingInst = "G"
	CustOrderHandlingInst_ALGO_ENGINE                                           CustOrderHandlingInst = "H"
	CustOrderHandlingInst_INTRADAY_CROSS                                        CustOrderHandlingInst = "IDX"
	CustOrderHandlingInst_IMBALANCE_ONLY                                        CustOrderHandlingInst = "IO"
	CustOrderHandlingInst_IMMEDIATE_OR_CANCEL                                   CustOrderHandlingInst = "IOC"
	CustOrderHandlingInst_INTERMARKET_SWEEP_ORDER                               CustOrderHandlingInst = "ISO"
	CustOrderHandlingInst_PRICE_AT_EXECUTION                                    CustOrderHandlingInst = "J"
	CustOrderHandlingInst_LIMIT_ON_CLOSE                                        CustOrderHandlingInst = "LOC"
	CustOrderHandlingInst_LIMIT_ON_OPEN                                         CustOrderHandlingInst = "LOO"
	CustOrderHandlingInst_MARKET_AT_CLOSE                                       CustOrderHandlingInst = "MAC"
	CustOrderHandlingInst_MARKET_AT_OPEN                                        CustOrderHandlingInst = "MAO"
	CustOrderHandlingInst_MARKET_ON_CLOSE                                       CustOrderHandlingInst = "MOC"
	CustOrderHandlingInst_MARKET_ON_OPEN                                        CustOrderHandlingInst = "MOO"
	CustOrderHandlingInst_MERGER_RELATED_TRANSFER_POSITION                      CustOrderHandlingInst = "MPT"
	CustOrderHandlingInst_MINIMUM_QUANTITY                                      CustOrderHandlingInst = "MQT"
	CustOrderHandlingInst_MARKET_TO_LIMIT                                       CustOrderHandlingInst = "MTL"
	CustOrderHandlingInst_DELIVERY_INSTRUCTIONS_ND                              CustOrderHandlingInst = "ND"
	CustOrderHandlingInst_NOT_HELD                                              CustOrderHandlingInst = "NH"
	CustOrderHandlingInst_OPTIONS_RELATED_TRANSACTION                           CustOrderHandlingInst = "OPT"
	CustOrderHandlingInst_OVER_THE_DAY                                          CustOrderHandlingInst = "OVD"
	CustOrderHandlingInst_PEGGED                                                CustOrderHandlingInst = "PEG"
	CustOrderHandlingInst_RESERVE_SIZE_ORDER                                    CustOrderHandlingInst = "RSV"
	CustOrderHandlingInst_STOP_STOCK_TRANSACTION                                CustOrderHandlingInst = "S.W"
	CustOrderHandlingInst_SCALE                                                 CustOrderHandlingInst = "SCL"
	CustOrderHandlingInst_DELIVERY_INSTRUCTIONS_SLR                             CustOrderHandlingInst = "SLR"
	CustOrderHandlingInst_TIME_ORDER                                            CustOrderHandlingInst = "TMO"
	CustOrderHandlingInst_TRAILING_STOP                                         CustOrderHandlingInst = "TS"
	CustOrderHandlingInst_DESK_W                                                CustOrderHandlingInst = "W"
	CustOrderHandlingInst_WORK                                                  CustOrderHandlingInst = "WRK"
	CustOrderHandlingInst_DESK_X                                                CustOrderHandlingInst = "X"
	CustOrderHandlingInst_CLIENT_Y                                              CustOrderHandlingInst = "Y"
	CustOrderHandlingInst_CLIENT_Z                                              CustOrderHandlingInst = "Z"
)

// CustomerOrFirm field enumeration values.
type CustomerOrFirm string

const (
	CustomerOrFirm_CUSTOMER CustomerOrFirm = "0"
	CustomerOrFirm_FIRM     CustomerOrFirm = "1"
)

// CustomerPriority field enumeration values.
type CustomerPriority string

const (
	CustomerPriority_NO_PRIORITY            CustomerPriority = "0"
	CustomerPriority_UNCONDITIONAL_PRIORITY CustomerPriority = "1"
)

// CxlRejReason field enumeration values.
type CxlRejReason string

const (
	CxlRejReason_TOO_LATE_TO_CANCEL                                        CxlRejReason = "0"
	CxlRejReason_UNKNOWN_ORDER                                             CxlRejReason = "1"
	CxlRejReason_INVALID_PRICE_INCREMENT                                   CxlRejReason = "18"
	CxlRejReason_BROKER                                                    CxlRejReason = "2"
	CxlRejReason_ORDER_ALREADY_IN_PENDING_CANCEL_OR_PENDING_REPLACE_STATUS CxlRejReason = "3"
	CxlRejReason_UNABLE_TO_PROCESS_ORDER_MASS_CANCEL_REQUEST               CxlRejReason = "4"
	CxlRejReason_ORIGORDMODTIME                                            CxlRejReason = "5"
	CxlRejReason_DUPLICATE_CLORDID                                         CxlRejReason = "6"
	CxlRejReason_PRICE_EXCEEDS_CURRENT_PRICE                               CxlRejReason = "7"
	CxlRejReason_PRICE_EXCEEDS_CURRENT_PRICE_BAND                          CxlRejReason = "8"
	CxlRejReason_OTHER                                                     CxlRejReason = "99"
)

// CxlRejResponseTo field enumeration values.
type CxlRejResponseTo string

const (
	CxlRejResponseTo_ORDER_CANCEL_REQUEST         CxlRejResponseTo = "1"
	CxlRejResponseTo_ORDER_CANCEL_REPLACE_REQUEST CxlRejResponseTo = "2"
)

// CxlType field enumeration values.
type CxlType string

const (
	CxlType_FULL_REMAINING_QUANTITY CxlType = "F"
	CxlType_PARTIAL_CANCEL          CxlType = "P"
)

// DKReason field enumeration values.
type DKReason string

const (
	DKReason_UNKNOWN_SECURITY            DKReason = "A"
	DKReason_WRONG_SIDE                  DKReason = "B"
	DKReason_QUANTITY_EXCEEDS_ORDER      DKReason = "C"
	DKReason_NO_MATCHING_ORDER           DKReason = "D"
	DKReason_PRICE_EXCEEDS_LIMIT         DKReason = "E"
	DKReason_CALCULATION_DIFFERENCE      DKReason = "F"
	DKReason_NO_MATCHING_EXECUTIONREPORT DKReason = "G"
	DKReason_OTHER                       DKReason = "Z"
)

// DateRollConvention field enumeration values.
type DateRollConvention string

const (
	DateRollConvention_1ST_DAY_OF_THE_MONTH                                                                                      DateRollConvention = "1"
	DateRollConvention_10TH_DAY_OF_THE_MONTH                                                                                     DateRollConvention = "10"
	DateRollConvention_11TH_DAY_OF_THE_MONTH                                                                                     DateRollConvention = "11"
	DateRollConvention_12TH_DAY_OF_THE_MONTH                                                                                     DateRollConvention = "12"
	DateRollConvention_13TH_DAY_OF_THE_MONTH                                                                                     DateRollConvention = "13"
	DateRollConvention_14TH_DAY_OF_THE_MONTH                                                                                     DateRollConvention = "14"
	DateRollConvention_15TH_DAY_OF_THE_MONTH                                                                                     DateRollConvention = "15"
	DateRollConvention_16TH_DAY_OF_THE_MONTH                                                                                     DateRollConvention = "16"
	DateRollConvention_17TH_DAY_OF_THE_MONTH                                                                                     DateRollConvention = "17"
	DateRollConvention_18TH_DAY_OF_THE_MONTH                                                                                     DateRollConvention = "18"
	DateRollConvention_19TH_DAY_OF_THE_MONTH                                                                                     DateRollConvention = "19"
	DateRollConvention_2ND_DAY_OF_THE_MONTH                                                                                      DateRollConvention = "2"
	DateRollConvention_20TH_DAY_OF_THE_MONTH                                                                                     DateRollConvention = "20"
	DateRollConvention_21ST_DAY_OF_THE_MONTH                                                                                     DateRollConvention = "21"
	DateRollConvention_22ND_DAY_OF_THE_MONTH                                                                                     DateRollConvention = "22"
	DateRollConvention_23RD_DAY_OF_THE_MONTH                                                                                     DateRollConvention = "23"
	DateRollConvention_24TH_DAY_OF_THE_MONTH                                                                                     DateRollConvention = "24"
	DateRollConvention_25TH_DAY_OF_THE_MONTH                                                                                     DateRollConvention = "25"
	DateRollConvention_26TH_DAY_OF_THE_MONTH                                                                                     DateRollConvention = "26"
	DateRollConvention_27TH_DAY_OF_THE_MONTH                                                                                     DateRollConvention = "27"
	DateRollConvention_28TH_DAY_OF_THE_MONTH                                                                                     DateRollConvention = "28"
	DateRollConvention_29TH_DAY_OF_THE_MONTH                                                                                     DateRollConvention = "29"
	DateRollConvention_3RD_DAY_OF_THE_MONTH                                                                                      DateRollConvention = "3"
	DateRollConvention_30TH_DAY_OF_THE_MONTH                                                                                     DateRollConvention = "30"
	DateRollConvention_4TH_DAY_OF_THE_MONTH                                                                                      DateRollConvention = "4"
	DateRollConvention_5TH_DAY_OF_THE_MONTH                                                                                      DateRollConvention = "5"
	DateRollConvention_6THD_DAY_OF_THE_MONTH                                                                                     DateRollConvention = "6"
	DateRollConvention_7TH_DAY_OF_THE_MONTH                                                                                      DateRollConvention = "7"
	DateRollConvention_8TH_DAY_OF_THE_MONTH                                                                                      DateRollConvention = "8"
	DateRollConvention_9TH_DAY_OF_THE_MONTH                                                                                      DateRollConvention = "9"
	DateRollConvention_THE_END_OF_THE_MONTH                                                                                      DateRollConvention = "EOM"
	DateRollConvention_FRIDAY                                                                                                    DateRollConvention = "FRI"
	DateRollConvention_THE_FLOATING_RATE_NOTE_CONVENTION_OR_EURODOLLAR_CONVENTION                                                DateRollConvention = "FRN"
	DateRollConvention_THE_INTERNATIONAL_MONEY_MARKET_SETTLEMENT_DATE_IE_THE_3RD_WEDNESDAY_OF_THE_MONTH                          DateRollConvention = "IMM"
	DateRollConvention_THE_LAST_TRADING_DAY_OF_THE_SYDNEY_FUTURES_EXCHANGE_AUSTRALIAN_90_DAY_BANK_ACCEPTED_BILL_FUTURES_CONTRACT DateRollConvention = "IMMAUD"
	DateRollConvention_THE_LAST_TRADING_DAY_EXPIRATION_DAY_OF_THE_CANADIAN_DERIVATIVES_EXCHANGE                                  DateRollConvention = "IMMCAD"
	DateRollConvention_THE_LAST_TRADING_DAY_OF_THE_SYDNEY_FUTURES_EXCHANGE_NEW_ZEALAND_90_DAY_BANK_BILL_FUTURES_CONTRACT         DateRollConvention = "IMMNZD"
	DateRollConvention_MONDAY                                                                                                    DateRollConvention = "MON"
	DateRollConvention_NO_ADJUSTMENT                                                                                             DateRollConvention = "NONE"
	DateRollConvention_SATURDAY                                                                                                  DateRollConvention = "SAT"
	DateRollConvention_THE_SYDNEY_FUTURES_EXCHANGE_90_DAY_BANK_ACCEPTED_BILL_FUTURES_SETTLEMENT_DATES                            DateRollConvention = "SFE"
	DateRollConvention_SUNDAY                                                                                                    DateRollConvention = "SUN"
	DateRollConvention_THE_13_WEEK_AND_26_WEEK_US_TREASURY_BILL_AUCTION_DATES                                                    DateRollConvention = "TBILL"
	DateRollConvention_THURSDAY                                                                                                  DateRollConvention = "THU"
	DateRollConvention_TUESDAY                                                                                                   DateRollConvention = "TUE"
	DateRollConvention_WEDNESDAY                                                                                                 DateRollConvention = "WED"
)

// DayBookingInst field enumeration values.
type DayBookingInst string

const (
	DayBookingInst_CAN_TRIGGER_BOOKING_WITHOUT_REFERENCE_TO_THE_ORDER_INITIATOR DayBookingInst = "0"
	DayBookingInst_SPEAK_WITH_ORDER_INITIATOR_BEFORE_BOOKING                    DayBookingInst = "1"
	DayBookingInst_ACCUMULATE                                                   DayBookingInst = "2"
)

// DealingCapacity field enumeration values.
type DealingCapacity string

const (
	DealingCapacity_AGENT              DealingCapacity = "A"
	DealingCapacity_PRINCIPAL          DealingCapacity = "P"
	DealingCapacity_RISKLESS_PRINCIPAL DealingCapacity = "R"
)

// DeleteReason field enumeration values.
type DeleteReason string

const (
	DeleteReason_CANCELLATION DeleteReason = "0"
	DeleteReason_ERROR        DeleteReason = "1"
)

// DeliveryForm field enumeration values.
type DeliveryForm string

const (
	DeliveryForm_BOOK_ENTRY DeliveryForm = "1"
	DeliveryForm_BEARER     DeliveryForm = "2"
)

// DeliveryScheduleSettlDay field enumeration values.
type DeliveryScheduleSettlDay string

const (
	DeliveryScheduleSettlDay_MONDAY       DeliveryScheduleSettlDay = "1"
	DeliveryScheduleSettlDay_ALL_WEEKENDS DeliveryScheduleSettlDay = "10"
	DeliveryScheduleSettlDay_TUESDAY      DeliveryScheduleSettlDay = "2"
	DeliveryScheduleSettlDay_WEDNESDAY    DeliveryScheduleSettlDay = "3"
	DeliveryScheduleSettlDay_THURSDAY     DeliveryScheduleSettlDay = "4"
	DeliveryScheduleSettlDay_FRIDAY       DeliveryScheduleSettlDay = "5"
	DeliveryScheduleSettlDay_SATURDAY     DeliveryScheduleSettlDay = "6"
	DeliveryScheduleSettlDay_SUNDAY       DeliveryScheduleSettlDay = "7"
	DeliveryScheduleSettlDay_ALL_WEEKDAYS DeliveryScheduleSettlDay = "8"
	DeliveryScheduleSettlDay_ALL_DAYS     DeliveryScheduleSettlDay = "9"
)

// DeliveryScheduleSettlFlowType field enumeration values.
type DeliveryScheduleSettlFlowType string

const (
	DeliveryScheduleSettlFlowType_ALL_TIMES   DeliveryScheduleSettlFlowType = "0"
	DeliveryScheduleSettlFlowType_ON_PEAK     DeliveryScheduleSettlFlowType = "1"
	DeliveryScheduleSettlFlowType_OFF_PEAK    DeliveryScheduleSettlFlowType = "2"
	DeliveryScheduleSettlFlowType_BASE        DeliveryScheduleSettlFlowType = "3"
	DeliveryScheduleSettlFlowType_BLOCK_HOURS DeliveryScheduleSettlFlowType = "4"
	DeliveryScheduleSettlFlowType_OTHER       DeliveryScheduleSettlFlowType = "5"
)

// DeliveryScheduleSettlHolidaysProcessingInstruction field enumeration values.
type DeliveryScheduleSettlHolidaysProcessingInstruction string

const (
	DeliveryScheduleSettlHolidaysProcessingInstruction_DO_NOT_INCLUDE_HOLIDAYS DeliveryScheduleSettlHolidaysProcessingInstruction = "0"
	DeliveryScheduleSettlHolidaysProcessingInstruction_INCLUDE_HOLIDAYS        DeliveryScheduleSettlHolidaysProcessingInstruction = "1"
)

// DeliveryScheduleSettlTimeType field enumeration values.
type DeliveryScheduleSettlTimeType string

const (
	DeliveryScheduleSettlTimeType_HOUR_OF_THE_DAY  DeliveryScheduleSettlTimeType = "0"
	DeliveryScheduleSettlTimeType_HHMM_TIME_FORMAT DeliveryScheduleSettlTimeType = "1"
)

// DeliveryScheduleToleranceType field enumeration values.
type DeliveryScheduleToleranceType string

const (
	DeliveryScheduleToleranceType_ABSOLUTE   DeliveryScheduleToleranceType = "0"
	DeliveryScheduleToleranceType_PERCENTAGE DeliveryScheduleToleranceType = "1"
)

// DeliveryScheduleType field enumeration values.
type DeliveryScheduleType string

const (
	DeliveryScheduleType_NOTIONAL                   DeliveryScheduleType = "0"
	DeliveryScheduleType_DELIVERY                   DeliveryScheduleType = "1"
	DeliveryScheduleType_PHYSICAL_SETTLEMENT_PERIOD DeliveryScheduleType = "2"
)

// DeliveryStreamDeliveryPointSource field enumeration values.
type DeliveryStreamDeliveryPointSource string

const (
	DeliveryStreamDeliveryPointSource_PROPRIETARY                DeliveryStreamDeliveryPointSource = "0"
	DeliveryStreamDeliveryPointSource_ENERGY_IDENTIFICATION_CODE DeliveryStreamDeliveryPointSource = "1"
)

// DeliveryStreamDeliveryRestriction field enumeration values.
type DeliveryStreamDeliveryRestriction string

const (
	DeliveryStreamDeliveryRestriction_FIRM                      DeliveryStreamDeliveryRestriction = "1"
	DeliveryStreamDeliveryRestriction_INTERRUPTABLE_OR_NON_FIRM DeliveryStreamDeliveryRestriction = "2"
	DeliveryStreamDeliveryRestriction_FORCE_MAJEURE             DeliveryStreamDeliveryRestriction = "3"
	DeliveryStreamDeliveryRestriction_SYSTEM_FIRM               DeliveryStreamDeliveryRestriction = "4"
	DeliveryStreamDeliveryRestriction_UNIT_FIRM                 DeliveryStreamDeliveryRestriction = "5"
)

// DeliveryStreamElectingPartySide field enumeration values.
type DeliveryStreamElectingPartySide string

const (
	DeliveryStreamElectingPartySide_BUYER  DeliveryStreamElectingPartySide = "0"
	DeliveryStreamElectingPartySide_SELLER DeliveryStreamElectingPartySide = "1"
)

// DeliveryStreamTitleTransferCondition field enumeration values.
type DeliveryStreamTitleTransferCondition string

const (
	DeliveryStreamTitleTransferCondition_TRANSFERS_WITH_RISK_OF_LOSS         DeliveryStreamTitleTransferCondition = "0"
	DeliveryStreamTitleTransferCondition_DOES_NOT_TRANSFER_WITH_RISK_OF_LOSS DeliveryStreamTitleTransferCondition = "1"
)

// DeliveryStreamToleranceOptionSide field enumeration values.
type DeliveryStreamToleranceOptionSide string

const (
	DeliveryStreamToleranceOptionSide_BUYER  DeliveryStreamToleranceOptionSide = "1"
	DeliveryStreamToleranceOptionSide_SELLER DeliveryStreamToleranceOptionSide = "2"
)

// DeliveryStreamType field enumeration values.
type DeliveryStreamType string

const (
	DeliveryStreamType_PERIODIC DeliveryStreamType = "0"
	DeliveryStreamType_INITIAL  DeliveryStreamType = "1"
	DeliveryStreamType_SINGLE   DeliveryStreamType = "2"
)

// DeliveryType field enumeration values.
type DeliveryType string

const (
	DeliveryType_VERSUS_PAYMENT_DELIVER DeliveryType = "0"
	DeliveryType_FREE_DELIVER           DeliveryType = "1"
	DeliveryType_TRI_PARTY              DeliveryType = "2"
	DeliveryType_HOLD_IN_CUSTODY        DeliveryType = "3"
	DeliveryType_DELIVER_BY_VALUE       DeliveryType = "4"
)

// DerivativeSecurityListRequestType field enumeration values.
type DerivativeSecurityListRequestType string

const (
	DerivativeSecurityListRequestType_SYMBOL                                    DerivativeSecurityListRequestType = "0"
	DerivativeSecurityListRequestType_SECURITYTYPE_AND_OR_CFICODE               DerivativeSecurityListRequestType = "1"
	DerivativeSecurityListRequestType_PRODUCT                                   DerivativeSecurityListRequestType = "2"
	DerivativeSecurityListRequestType_TRADINGSESSIONID                          DerivativeSecurityListRequestType = "3"
	DerivativeSecurityListRequestType_ALL_SECURITIES                            DerivativeSecurityListRequestType = "4"
	DerivativeSecurityListRequestType_UNDELYINGSYMBOL                           DerivativeSecurityListRequestType = "5"
	DerivativeSecurityListRequestType_UNDERLYING_SECURITYTYPE_AND_OR_CFICODE    DerivativeSecurityListRequestType = "6"
	DerivativeSecurityListRequestType_UNDERLYING_PRODUCT                        DerivativeSecurityListRequestType = "7"
	DerivativeSecurityListRequestType_MARKETID_OR_MARKETID_PLUS_MARKETSEGMENTID DerivativeSecurityListRequestType = "8"
)

// DeskOrderHandlingInst field enumeration values.
type DeskOrderHandlingInst string

const (
	DeskOrderHandlingInst_ADD_ON_ORDER                      DeskOrderHandlingInst = "ADD"
	DeskOrderHandlingInst_ALL_OR_NONE                       DeskOrderHandlingInst = "AON"
	DeskOrderHandlingInst_CASH_NOT_HELD                     DeskOrderHandlingInst = "CNH"
	DeskOrderHandlingInst_DIRECTED_ORDER                    DeskOrderHandlingInst = "DIR"
	DeskOrderHandlingInst_EXCHANGE_FOR_PHYSICAL_TRANSACTION DeskOrderHandlingInst = "E.W"
	DeskOrderHandlingInst_FILL_OR_KILL                      DeskOrderHandlingInst = "FOK"
	DeskOrderHandlingInst_IMBALANCE_ONLY                    DeskOrderHandlingInst = "IO"
	DeskOrderHandlingInst_IMMEDIATE_OR_CANCEL               DeskOrderHandlingInst = "IOC"
	DeskOrderHandlingInst_LIMIT_ON_CLOSE                    DeskOrderHandlingInst = "LOC"
	DeskOrderHandlingInst_LIMIT_ON_OPEN                     DeskOrderHandlingInst = "LOO"
	DeskOrderHandlingInst_MARKET_AT_CLOSE                   DeskOrderHandlingInst = "MAC"
	DeskOrderHandlingInst_MARKET_AT_OPEN                    DeskOrderHandlingInst = "MAO"
	DeskOrderHandlingInst_MARKET_ON_CLOSE                   DeskOrderHandlingInst = "MOC"
	DeskOrderHandlingInst_MARKET_ON_OPEN                    DeskOrderHandlingInst = "MOO"
	DeskOrderHandlingInst_MINIMUM_QUANTITY                  DeskOrderHandlingInst = "MQT"
	DeskOrderHandlingInst_NOT_HELD                          DeskOrderHandlingInst = "NH"
	DeskOrderHandlingInst_OVER_THE_DAY                      DeskOrderHandlingInst = "OVD"
	DeskOrderHandlingInst_PEGGED                            DeskOrderHandlingInst = "PEG"
	DeskOrderHandlingInst_RESERVE_SIZE_ORDER                DeskOrderHandlingInst = "RSV"
	DeskOrderHandlingInst_STOP_STOCK_TRANSACTION            DeskOrderHandlingInst = "S.W"
	DeskOrderHandlingInst_SCALE                             DeskOrderHandlingInst = "SCL"
	DeskOrderHandlingInst_TIME_ORDER                        DeskOrderHandlingInst = "TMO"
	DeskOrderHandlingInst_TRAILING_STOP                     DeskOrderHandlingInst = "TS"
	DeskOrderHandlingInst_WORK                              DeskOrderHandlingInst = "WRK"
)

// DeskType field enumeration values.
type DeskType string

const (
	DeskType_AGENCY                                  DeskType = "A"
	DeskType_ARBITRAGE                               DeskType = "AR"
	DeskType_BLOCK_TRADING                           DeskType = "B"
	DeskType_CONVERTIBLE_DESK                        DeskType = "C"
	DeskType_CENTRAL_RISK_BOOKS                      DeskType = "CR"
	DeskType_DERIVATIVES                             DeskType = "D"
	DeskType_EQUITY_CAPITAL_MARKETS                  DeskType = "EC"
	DeskType_FLOOR_BROKER                            DeskType = "FB"
	DeskType_INTERNATIONAL                           DeskType = "IN"
	DeskType_INSTITUTIONAL                           DeskType = "IS"
	DeskType_OTHER                                   DeskType = "O"
	DeskType_PREFERRED_TRADING                       DeskType = "PF"
	DeskType_PROPRIETARY                             DeskType = "PR"
	DeskType_PROGRAM_TRADING                         DeskType = "PT"
	DeskType_SALES                                   DeskType = "S"
	DeskType_SWAPS                                   DeskType = "SW"
	DeskType_TRADING_DESK_OR_SYSTEM_NON_MARKET_MAKER DeskType = "T"
	DeskType_TREASURY                                DeskType = "TR"
)

// DeskTypeSource field enumeration values.
type DeskTypeSource string

const (
	DeskTypeSource_FINRA_OATS DeskTypeSource = "1"
)

// DisclosureInstruction field enumeration values.
type DisclosureInstruction string

const (
	DisclosureInstruction_NO                  DisclosureInstruction = "0"
	DisclosureInstruction_YES                 DisclosureInstruction = "1"
	DisclosureInstruction_USE_DEFAULT_SETTING DisclosureInstruction = "2"
)

// DisclosureType field enumeration values.
type DisclosureType string

const (
	DisclosureType_VOLUME           DisclosureType = "1"
	DisclosureType_PRICE            DisclosureType = "2"
	DisclosureType_SIDE             DisclosureType = "3"
	DisclosureType_AON              DisclosureType = "4"
	DisclosureType_GENERAL          DisclosureType = "5"
	DisclosureType_CLEARING_ACCOUNT DisclosureType = "6"
	DisclosureType_CMTA_ACCOUNT     DisclosureType = "7"
)

// DiscretionInst field enumeration values.
type DiscretionInst string

const (
	DiscretionInst_RELATED_TO_DISPLAYED_PRICE     DiscretionInst = "0"
	DiscretionInst_RELATED_TO_MARKET_PRICE        DiscretionInst = "1"
	DiscretionInst_RELATED_TO_PRIMARY_PRICE       DiscretionInst = "2"
	DiscretionInst_RELATED_TO_LOCAL_PRIMARY_PRICE DiscretionInst = "3"
	DiscretionInst_RELATED_TO_MIDPOINT_PRICE      DiscretionInst = "4"
	DiscretionInst_RELATED_TO_LAST_TRADE_PRICE    DiscretionInst = "5"
	DiscretionInst_RELATED_TO_VWAP                DiscretionInst = "6"
	DiscretionInst_AVERAGE_PRICE_GUARANTEE        DiscretionInst = "7"
)

// DiscretionLimitType field enumeration values.
type DiscretionLimitType string

const (
	DiscretionLimitType_OR_BETTER DiscretionLimitType = "0"
	DiscretionLimitType_STRICT    DiscretionLimitType = "1"
	DiscretionLimitType_OR_WORSE  DiscretionLimitType = "2"
)

// DiscretionMoveType field enumeration values.
type DiscretionMoveType string

const (
	DiscretionMoveType_FLOATING DiscretionMoveType = "0"
	DiscretionMoveType_FIXED    DiscretionMoveType = "1"
)

// DiscretionOffsetType field enumeration values.
type DiscretionOffsetType string

const (
	DiscretionOffsetType_PRICE        DiscretionOffsetType = "0"
	DiscretionOffsetType_BASIS_POINTS DiscretionOffsetType = "1"
	DiscretionOffsetType_TICKS        DiscretionOffsetType = "2"
	DiscretionOffsetType_PRICE_TIER   DiscretionOffsetType = "3"
)

// DiscretionRoundDirection field enumeration values.
type DiscretionRoundDirection string

const (
	DiscretionRoundDirection_MORE_AGGRESSIVE DiscretionRoundDirection = "1"
	DiscretionRoundDirection_MORE_PASSIVE    DiscretionRoundDirection = "2"
)

// DiscretionScope field enumeration values.
type DiscretionScope string

const (
	DiscretionScope_LOCAL                    DiscretionScope = "1"
	DiscretionScope_NATIONAL                 DiscretionScope = "2"
	DiscretionScope_GLOBAL                   DiscretionScope = "3"
	DiscretionScope_NATIONAL_EXCLUDING_LOCAL DiscretionScope = "4"
)

// DisplayMethod field enumeration values.
type DisplayMethod string

const (
	DisplayMethod_INITIAL     DisplayMethod = "1"
	DisplayMethod_NEW         DisplayMethod = "2"
	DisplayMethod_RANDOM      DisplayMethod = "3"
	DisplayMethod_UNDISCLOSED DisplayMethod = "4"
)

// DisplayWhen field enumeration values.
type DisplayWhen string

const (
	DisplayWhen_IMMEDIATE DisplayWhen = "1"
	DisplayWhen_EXHAUST   DisplayWhen = "2"
)

// DistribPaymentMethod field enumeration values.
type DistribPaymentMethod string

const (
	DistribPaymentMethod_CREST                            DistribPaymentMethod = "1"
	DistribPaymentMethod_BPAY                             DistribPaymentMethod = "10"
	DistribPaymentMethod_HIGH_VALUE_CLEARING_SYSTEM_HVACS DistribPaymentMethod = "11"
	DistribPaymentMethod_REINVEST_IN_FUND                 DistribPaymentMethod = "12"
	DistribPaymentMethod_NSCC                             DistribPaymentMethod = "2"
	DistribPaymentMethod_EUROCLEAR                        DistribPaymentMethod = "3"
	DistribPaymentMethod_CLEARSTREAM                      DistribPaymentMethod = "4"
	DistribPaymentMethod_CHEQUE                           DistribPaymentMethod = "5"
	DistribPaymentMethod_TELEGRAPHIC_TRANSFER             DistribPaymentMethod = "6"
	DistribPaymentMethod_FED_WIRE                         DistribPaymentMethod = "7"
	DistribPaymentMethod_DIRECT_CREDIT                    DistribPaymentMethod = "8"
	DistribPaymentMethod_ACH_CREDIT                       DistribPaymentMethod = "9"
)

// DividendAmountType field enumeration values.
type DividendAmountType string

const (
	DividendAmountType_RECORD_AMOUNT                       DividendAmountType = "0"
	DividendAmountType_EX_AMOUNT                           DividendAmountType = "1"
	DividendAmountType_PAID_AMOUNT                         DividendAmountType = "2"
	DividendAmountType_AS_SPECIFIED_IN_MASTER_CONFIRMATION DividendAmountType = "3"
)

// DividendComposition field enumeration values.
type DividendComposition string

const (
	DividendComposition_EQUITY_AMOUNT_RECEIVER_ELECTION DividendComposition = "0"
	DividendComposition_CALCULATION_AGENT_ELECTION      DividendComposition = "1"
)

// DividendEntitlementEvent field enumeration values.
type DividendEntitlementEvent string

const (
	DividendEntitlementEvent_EX_DATE     DividendEntitlementEvent = "0"
	DividendEntitlementEvent_RECORD_DATE DividendEntitlementEvent = "1"
)

// DlvyInstType field enumeration values.
type DlvyInstType string

const (
	DlvyInstType_CASH       DlvyInstType = "C"
	DlvyInstType_SECURITIES DlvyInstType = "S"
)

// DueToRelated field enumeration values.
type DueToRelated string

const (
	DueToRelated_NO  DueToRelated = "N"
	DueToRelated_YES DueToRelated = "Y"
)

// DuplicateClOrdIDIndicator field enumeration values.
type DuplicateClOrdIDIndicator string

const (
	DuplicateClOrdIDIndicator_NO  DuplicateClOrdIDIndicator = "N"
	DuplicateClOrdIDIndicator_YES DuplicateClOrdIDIndicator = "Y"
)

// EmailType field enumeration values.
type EmailType string

const (
	EmailType_NEW         EmailType = "0"
	EmailType_REPLY       EmailType = "1"
	EmailType_ADMIN_REPLY EmailType = "2"
)

// EncryptMethod field enumeration values.
type EncryptMethod string

const (
	EncryptMethod_NONE_OTHER  EncryptMethod = "0"
	EncryptMethod_PKCS        EncryptMethod = "1"
	EncryptMethod_DES         EncryptMethod = "2"
	EncryptMethod_PKCS_DES    EncryptMethod = "3"
	EncryptMethod_PGP_DES     EncryptMethod = "4"
	EncryptMethod_PGP_DES_MD5 EncryptMethod = "5"
	EncryptMethod_PEM_DES_MD5 EncryptMethod = "6"
)

// EntitlementAttribDatatype field enumeration values.
type EntitlementAttribDatatype string

const (
	EntitlementAttribDatatype_INT                 EntitlementAttribDatatype = "1"
	EntitlementAttribDatatype_AMT                 EntitlementAttribDatatype = "10"
	EntitlementAttribDatatype_PERCENTAGE          EntitlementAttribDatatype = "11"
	EntitlementAttribDatatype_CHAR                EntitlementAttribDatatype = "12"
	EntitlementAttribDatatype_BOOLEAN             EntitlementAttribDatatype = "13"
	EntitlementAttribDatatype_STRING              EntitlementAttribDatatype = "14"
	EntitlementAttribDatatype_MULTIPLECHARVALUE   EntitlementAttribDatatype = "15"
	EntitlementAttribDatatype_CURRENCY            EntitlementAttribDatatype = "16"
	EntitlementAttribDatatype_EXCHANGE            EntitlementAttribDatatype = "17"
	EntitlementAttribDatatype_MONTHYEAR           EntitlementAttribDatatype = "18"
	EntitlementAttribDatatype_UTCTIMESTAMP        EntitlementAttribDatatype = "19"
	EntitlementAttribDatatype_LENGTH              EntitlementAttribDatatype = "2"
	EntitlementAttribDatatype_UTCTIMEONLY         EntitlementAttribDatatype = "20"
	EntitlementAttribDatatype_LOCALMKTDATE        EntitlementAttribDatatype = "21"
	EntitlementAttribDatatype_UTCDATEONLY         EntitlementAttribDatatype = "22"
	EntitlementAttribDatatype_DATA                EntitlementAttribDatatype = "23"
	EntitlementAttribDatatype_MULTIPLESTRINGVALUE EntitlementAttribDatatype = "24"
	EntitlementAttribDatatype_COUNTRY             EntitlementAttribDatatype = "25"
	EntitlementAttribDatatype_LANGUAGE            EntitlementAttribDatatype = "26"
	EntitlementAttribDatatype_TZTIMEONLY          EntitlementAttribDatatype = "27"
	EntitlementAttribDatatype_TZTIMESTAMP         EntitlementAttribDatatype = "28"
	EntitlementAttribDatatype_TENOR               EntitlementAttribDatatype = "29"
	EntitlementAttribDatatype_NUMINGROUP          EntitlementAttribDatatype = "3"
	EntitlementAttribDatatype_DAYOFMONTH          EntitlementAttribDatatype = "30"
	EntitlementAttribDatatype_XMLDATA             EntitlementAttribDatatype = "31"
	EntitlementAttribDatatype_PATTERN             EntitlementAttribDatatype = "32"
	EntitlementAttribDatatype_RESERVED100PLUS     EntitlementAttribDatatype = "33"
	EntitlementAttribDatatype_RESERVED1000PLUS    EntitlementAttribDatatype = "34"
	EntitlementAttribDatatype_RESERVED4000PLUS    EntitlementAttribDatatype = "35"
	EntitlementAttribDatatype_SEQNUM              EntitlementAttribDatatype = "4"
	EntitlementAttribDatatype_TAGNUM              EntitlementAttribDatatype = "5"
	EntitlementAttribDatatype_FLOAT               EntitlementAttribDatatype = "6"
	EntitlementAttribDatatype_QTY                 EntitlementAttribDatatype = "7"
	EntitlementAttribDatatype_PRICE               EntitlementAttribDatatype = "8"
	EntitlementAttribDatatype_PRICEOFFSET         EntitlementAttribDatatype = "9"
)

// EntitlementRequestResult field enumeration values.
type EntitlementRequestResult string

const (
	EntitlementRequestResult_SUCCESSFUL                            EntitlementRequestResult = "0"
	EntitlementRequestResult_INVALID_PARTY                         EntitlementRequestResult = "1"
	EntitlementRequestResult_INSTRUMENT_SCOPE_NOT_SUPPORTED        EntitlementRequestResult = "10"
	EntitlementRequestResult_MARKET_SEGMENT_SCOPE_NOT_SUPPORTED    EntitlementRequestResult = "11"
	EntitlementRequestResult_ENTITLEMENT_NOT_APPROVED_FOR_PARTY    EntitlementRequestResult = "12"
	EntitlementRequestResult_ENTITLEMENT_ALREADY_DEFINED_FOR_PARTY EntitlementRequestResult = "13"
	EntitlementRequestResult_INSTRUMENT_NOT_APPROVED_FOR_PARTY     EntitlementRequestResult = "14"
	EntitlementRequestResult_INVALID_RELATED_PARTY                 EntitlementRequestResult = "2"
	EntitlementRequestResult_INVALID_ENTITLEMENT_TYPE              EntitlementRequestResult = "3"
	EntitlementRequestResult_INVALID_ENTITLEMENT_ID                EntitlementRequestResult = "4"
	EntitlementRequestResult_INVALID_ENTITLEMENT_ATTRIBUTE         EntitlementRequestResult = "5"
	EntitlementRequestResult_INVALID_INSTRUMENT_SCOPE              EntitlementRequestResult = "6"
	EntitlementRequestResult_INVALID_MARKET_SEGMENT_SCOPE          EntitlementRequestResult = "7"
	EntitlementRequestResult_INVALID_START_DATE                    EntitlementRequestResult = "8"
	EntitlementRequestResult_INVALID_END_DATE                      EntitlementRequestResult = "9"
	EntitlementRequestResult_NOT_AUTHORIZED                        EntitlementRequestResult = "98"
	EntitlementRequestResult_OTHER                                 EntitlementRequestResult = "99"
)

// EntitlementStatus field enumeration values.
type EntitlementStatus string

const (
	EntitlementStatus_ACCEPTED              EntitlementStatus = "0"
	EntitlementStatus_ACCEPTED_WITH_CHANGES EntitlementStatus = "1"
	EntitlementStatus_REJECTED              EntitlementStatus = "2"
	EntitlementStatus_PENDING               EntitlementStatus = "3"
	EntitlementStatus_REQUESTED             EntitlementStatus = "4"
	EntitlementStatus_DEFERRED              EntitlementStatus = "5"
)

// EntitlementSubType field enumeration values.
type EntitlementSubType string

const (
	EntitlementSubType_ORDER_ENTRY            EntitlementSubType = "1"
	EntitlementSubType_HIT_LIFT               EntitlementSubType = "2"
	EntitlementSubType_VIEW_INDICATIVE_PRICES EntitlementSubType = "3"
	EntitlementSubType_VIEW_EXECUTABLE_PRICES EntitlementSubType = "4"
	EntitlementSubType_SINGLE_QUOTE           EntitlementSubType = "5"
	EntitlementSubType_STREAMING_QUOTES       EntitlementSubType = "6"
	EntitlementSubType_SINGLE_BROKER          EntitlementSubType = "7"
	EntitlementSubType_MULTI_BROKERS          EntitlementSubType = "8"
)

// EntitlementType field enumeration values.
type EntitlementType string

const (
	EntitlementType_TRADE                          EntitlementType = "0"
	EntitlementType_MAKE_MARKETS                   EntitlementType = "1"
	EntitlementType_HOLD_POSITIONS                 EntitlementType = "2"
	EntitlementType_PERFORM_GIVE_UPS               EntitlementType = "3"
	EntitlementType_SUBMIT_INDICATIONS_OF_INTEREST EntitlementType = "4"
	EntitlementType_SUBSCRIBE_TO_MARKET_DATA       EntitlementType = "5"
	EntitlementType_SHORT_WITH_PRE_BORROW          EntitlementType = "6"
	EntitlementType_SUBMIT_QUOTE_REQUESTS          EntitlementType = "7"
	EntitlementType_RESPOND_TO_QUOTE_REQUESTS      EntitlementType = "8"
)

// EventInitiatorType field enumeration values.
type EventInitiatorType string

const (
	EventInitiatorType_CUSTOMER_OR_CLIENT          EventInitiatorType = "C"
	EventInitiatorType_EXCHANGE_OR_EXECUTION_VENUE EventInitiatorType = "E"
	EventInitiatorType_FIRM_OR_BROKER              EventInitiatorType = "F"
)

// EventTimeUnit field enumeration values.
type EventTimeUnit string

const (
	EventTimeUnit_DAY    EventTimeUnit = "D"
	EventTimeUnit_HOUR   EventTimeUnit = "H"
	EventTimeUnit_MINUTE EventTimeUnit = "Min"
	EventTimeUnit_MONTH  EventTimeUnit = "Mo"
	EventTimeUnit_SECOND EventTimeUnit = "S"
	EventTimeUnit_WEEK   EventTimeUnit = "Wk"
	EventTimeUnit_YEAR   EventTimeUnit = "Yr"
)

// EventType field enumeration values.
type EventType string

const (
	EventType_PUT                               EventType = "1"
	EventType_SWAP_ROLL_DATE                    EventType = "10"
	EventType_SWAP_NEXT_START_DATE              EventType = "11"
	EventType_SWAP_NEXT_ROLL_DATE               EventType = "12"
	EventType_FIRST_DELIVERY_DATE               EventType = "13"
	EventType_LAST_DELIVERY_DATE                EventType = "14"
	EventType_INITIAL_INVENTORY_DUE_DATE        EventType = "15"
	EventType_FINAL_INVENTORY_DUE_DATE          EventType = "16"
	EventType_FIRST_INTENT_DATE                 EventType = "17"
	EventType_LAST_INTENT_DATE                  EventType = "18"
	EventType_POSITION_REMOVAL_DATE             EventType = "19"
	EventType_CALL                              EventType = "2"
	EventType_MINIMUM_NOTICE                    EventType = "20"
	EventType_DELIVERY_START_TIME               EventType = "21"
	EventType_DELIVERY_END_TIME                 EventType = "22"
	EventType_FIRST_NOTICE_DATE                 EventType = "23"
	EventType_LAST_NOTICE_DATE                  EventType = "24"
	EventType_FIRST_EXERCISE_DATE               EventType = "25"
	EventType_REDEMPTION_DATE                   EventType = "26"
	EventType_TRADE_CONTINUATION_EFFECTIVE_DATE EventType = "27"
	EventType_TENDER                            EventType = "3"
	EventType_SINKING_FUND_CALL                 EventType = "4"
	EventType_ACTIVATION                        EventType = "5"
	EventType_INACTIVATION                      EventType = "6"
	EventType_LAST_ELIGIBLE_TRADE_DATE          EventType = "7"
	EventType_SWAP_START_DATE                   EventType = "8"
	EventType_SWAP_END_DATE                     EventType = "9"
	EventType_OTHER                             EventType = "99"
)

// ExDestination field enumeration values.
type ExDestination string

const (
	ExDestination_NONE  ExDestination = "0"
	ExDestination_POSIT ExDestination = "4"
)

// ExDestinationIDSource field enumeration values.
type ExDestinationIDSource string

const (
	ExDestinationIDSource_BIC                                              ExDestinationIDSource = "B"
	ExDestinationIDSource_GENERALLY_ACCEPTED_MARKET_PARTICIPANT_IDENTIFIER ExDestinationIDSource = "C"
	ExDestinationIDSource_PROPRIETARY                                      ExDestinationIDSource = "D"
	ExDestinationIDSource_ISO_COUNTRY_CODE                                 ExDestinationIDSource = "E"
	ExDestinationIDSource_MIC                                              ExDestinationIDSource = "G"
)

// ExDestinationType field enumeration values.
type ExDestinationType string

const (
	ExDestinationType_NO_RESTRICTION                                              ExDestinationType = "0"
	ExDestinationType_CAN_BE_TRADED_ONLY_ON_A_TRADING_VENUE                       ExDestinationType = "1"
	ExDestinationType_CAN_BE_TRADED_ONLY_ON_A_SYSTEMATIC_INTERNALISER             ExDestinationType = "2"
	ExDestinationType_CAN_BE_TRADED_ON_A_TRADING_VENUE_OR_SYSTEMATIC_INTERNALISER ExDestinationType = "3"
)

// ExchangeForPhysical field enumeration values.
type ExchangeForPhysical string

const (
	ExchangeForPhysical_NO  ExchangeForPhysical = "N"
	ExchangeForPhysical_YES ExchangeForPhysical = "Y"
)

// ExecAckStatus field enumeration values.
type ExecAckStatus string

const (
	ExecAckStatus_RECEIVED_NOT_YET_PROCESSED ExecAckStatus = "0"
	ExecAckStatus_ACCEPTED                   ExecAckStatus = "1"
	ExecAckStatus_DONT_KNOW                  ExecAckStatus = "2"
)

// ExecInst field enumeration values.
type ExecInst string

const (
	ExecInst_STAY_ON_OFFER_SIDE                                    ExecInst = "0"
	ExecInst_NOT_HELD                                              ExecInst = "1"
	ExecInst_WORK                                                  ExecInst = "2"
	ExecInst_GO_ALONG                                              ExecInst = "3"
	ExecInst_OVER_THE_DAY                                          ExecInst = "4"
	ExecInst_HELD                                                  ExecInst = "5"
	ExecInst_PARTICIPATE_DONT_INITIATE                             ExecInst = "6"
	ExecInst_STRICT_SCALE                                          ExecInst = "7"
	ExecInst_TRY_TO_SCALE                                          ExecInst = "8"
	ExecInst_STAY_ON_BID_SIDE                                      ExecInst = "9"
	ExecInst_NO_CROSS                                              ExecInst = "A"
	ExecInst_OK_TO_CROSS                                           ExecInst = "B"
	ExecInst_CALL_FIRST                                            ExecInst = "C"
	ExecInst_PERCENT_OF_VOLUME                                     ExecInst = "D"
	ExecInst_DO_NOT_INCREASE                                       ExecInst = "E"
	ExecInst_DO_NOT_REDUCE                                         ExecInst = "F"
	ExecInst_ALL_OR_NONE                                           ExecInst = "G"
	ExecInst_REINSTATE_ON_SYSTEM_FAILURE                           ExecInst = "H"
	ExecInst_INSTITUTIONS_ONLY                                     ExecInst = "I"
	ExecInst_REINSTATE_ON_TRADING_HALT                             ExecInst = "J"
	ExecInst_CANCEL_ON_TRADING_HALT                                ExecInst = "K"
	ExecInst_LAST_PEG                                              ExecInst = "L"
	ExecInst_MID_PRICE_PEG                                         ExecInst = "M"
	ExecInst_NON_NEGOTIABLE                                        ExecInst = "N"
	ExecInst_OPENING_PEG                                           ExecInst = "O"
	ExecInst_MARKET_PEG                                            ExecInst = "P"
	ExecInst_CANCEL_ON_SYSTEM_FAILURE                              ExecInst = "Q"
	ExecInst_PRIMARY_PEG                                           ExecInst = "R"
	ExecInst_SUSPEND                                               ExecInst = "S"
	ExecInst_FIXED_PEG_TO_LOCAL_BEST_BID_OR_OFFER_AT_TIME_OF_ORDER ExecInst = "T"
	ExecInst_CUSTOMER_DISPLAY_INSTRUCTION                          ExecInst = "U"
	ExecInst_NETTING                                               ExecInst = "V"
	ExecInst_PEG_TO_VWAP                                           ExecInst = "W"
	ExecInst_TRADE_ALONG                                           ExecInst = "X"
	ExecInst_TRY_TO_STOP                                           ExecInst = "Y"
	ExecInst_CANCEL_IF_NOT_BEST                                    ExecInst = "Z"
	ExecInst_TRAILING_STOP_PEG                                     ExecInst = "a"
	ExecInst_STRICT_LIMIT                                          ExecInst = "b"
	ExecInst_IGNORE_PRICE_VALIDITY_CHECKS                          ExecInst = "c"
	ExecInst_PEG_TO_LIMIT_PRICE                                    ExecInst = "d"
	ExecInst_WORK_TO_TARGET_STRATEGY                               ExecInst = "e"
	ExecInst_INTERMARKET_SWEEP                                     ExecInst = "f"
	ExecInst_EXTERNAL_ROUTING_ALLOWED                              ExecInst = "g"
	ExecInst_EXTERNAL_ROUTING_NOT_ALLOWED                          ExecInst = "h"
	ExecInst_IMBALANCE_ONLY                                        ExecInst = "i"
	ExecInst_SINGLE_EXECUTION_REQUESTED_FOR_BLOCK_TRADE            ExecInst = "j"
	ExecInst_BEST_EXECUTION                                        ExecInst = "k"
	ExecInst_SUSPEND_ON_SYSTEM_FAILURE                             ExecInst = "l"
	ExecInst_SUSPEND_ON_TRADING_HALT                               ExecInst = "m"
	ExecInst_REINSTATE_ON_CONNECTION_LOSS                          ExecInst = "n"
	ExecInst_CANCEL_ON_CONNECTION_LOSS                             ExecInst = "o"
	ExecInst_SUSPEND_ON_CONNECTION_LOSS                            ExecInst = "p"
	ExecInst_RELEASE                                               ExecInst = "q"
	ExecInst_EXECUTE_AS_DELTA_NEUTRAL_USING_VOLATILITY_PROVIDED    ExecInst = "r"
	ExecInst_EXECUTE_AS_DURATION_NEUTRAL                           ExecInst = "s"
	ExecInst_EXECUTE_AS_FX_NEUTRAL                                 ExecInst = "t"
	ExecInst_MINIMUM_GUARANTEED_FILL_ELIGIBLE                      ExecInst = "u"
	ExecInst_BYPASS_NON_DISPLAYED_LIQUIDITY                        ExecInst = "v"
	ExecInst_LOCK                                                  ExecInst = "w"
	ExecInst_IGNORE_NOTIONAL_VALUE_CHECKS                          ExecInst = "x"
	ExecInst_TRADE_AT_REFERENCE_PRICE                              ExecInst = "y"
	ExecInst_ALLOW_FACILITATION                                    ExecInst = "z"
)

// ExecMethod field enumeration values.
type ExecMethod string

const (
	ExecMethod_UNDEFINED_UNSPECIFIED ExecMethod = "0"
	ExecMethod_MANUAL                ExecMethod = "1"
	ExecMethod_AUTOMATED             ExecMethod = "2"
	ExecMethod_VOICE_BROKERED        ExecMethod = "3"
)

// ExecPriceType field enumeration values.
type ExecPriceType string

const (
	ExecPriceType_BID_PRICE                              ExecPriceType = "B"
	ExecPriceType_CREATION_PRICE                         ExecPriceType = "C"
	ExecPriceType_CREATION_PRICE_PLUS_ADJUSTMENT_PERCENT ExecPriceType = "D"
	ExecPriceType_CREATION_PRICE_PLUS_ADJUSTMENT_AMOUNT  ExecPriceType = "E"
	ExecPriceType_OFFER_PRICE                            ExecPriceType = "O"
	ExecPriceType_OFFER_PRICE_MINUS_ADJUSTMENT_PERCENT   ExecPriceType = "P"
	ExecPriceType_OFFER_PRICE_MINUS_ADJUSTMENT_AMOUNT    ExecPriceType = "Q"
	ExecPriceType_SINGLE_PRICE                           ExecPriceType = "S"
)

// ExecRestatementReason field enumeration values.
type ExecRestatementReason string

const (
	ExecRestatementReason_GT_CORPORATE_ACTION             ExecRestatementReason = "0"
	ExecRestatementReason_GT_RENEWAL                      ExecRestatementReason = "1"
	ExecRestatementReason_WAREHOUSE_RECAP                 ExecRestatementReason = "10"
	ExecRestatementReason_PEG_REFRESH                     ExecRestatementReason = "11"
	ExecRestatementReason_CANCEL_ON_CONNECTION_LOSS       ExecRestatementReason = "12"
	ExecRestatementReason_CANCEL_ON_LOGOUT                ExecRestatementReason = "13"
	ExecRestatementReason_ASSIGN_TIME_PRIORITY            ExecRestatementReason = "14"
	ExecRestatementReason_CANCELLED_TRADE_PRICE_VIOLATION ExecRestatementReason = "15"
	ExecRestatementReason_CANCELLED_CROSS_IMBALANCE       ExecRestatementReason = "16"
	ExecRestatementReason_VERBAL_CHANGE                   ExecRestatementReason = "2"
	ExecRestatementReason_REPRICING_OF_ORDER              ExecRestatementReason = "3"
	ExecRestatementReason_BROKER_OPTION                   ExecRestatementReason = "4"
	ExecRestatementReason_PARTIAL_DECLINE_OF_ORDERQTY     ExecRestatementReason = "5"
	ExecRestatementReason_CANCEL_ON_TRADING_HALT          ExecRestatementReason = "6"
	ExecRestatementReason_CANCEL_ON_SYSTEM_FAILURE        ExecRestatementReason = "7"
	ExecRestatementReason_MARKET                          ExecRestatementReason = "8"
	ExecRestatementReason_CANCELED_NOT_BEST               ExecRestatementReason = "9"
	ExecRestatementReason_OTHER                           ExecRestatementReason = "99"
)

// ExecTransType field enumeration values.
type ExecTransType string

const (
	ExecTransType_NEW     ExecTransType = "0"
	ExecTransType_CANCEL  ExecTransType = "1"
	ExecTransType_CORRECT ExecTransType = "2"
	ExecTransType_STATUS  ExecTransType = "3"
)

// ExecType field enumeration values.
type ExecType string

const (
	ExecType_NEW                                 ExecType = "0"
	ExecType_PARTIAL_FILL                        ExecType = "1"
	ExecType_FILL                                ExecType = "2"
	ExecType_DONE_FOR_DAY                        ExecType = "3"
	ExecType_CANCELED                            ExecType = "4"
	ExecType_REPLACED                            ExecType = "5"
	ExecType_PENDING_CANCEL                      ExecType = "6"
	ExecType_STOPPED                             ExecType = "7"
	ExecType_REJECTED                            ExecType = "8"
	ExecType_SUSPENDED                           ExecType = "9"
	ExecType_PENDING_NEW                         ExecType = "A"
	ExecType_CALCULATED                          ExecType = "B"
	ExecType_EXPIRED                             ExecType = "C"
	ExecType_RESTATED                            ExecType = "D"
	ExecType_PENDING_REPLACE                     ExecType = "E"
	ExecType_TRADE                               ExecType = "F"
	ExecType_TRADE_CORRECT                       ExecType = "G"
	ExecType_TRADE_CANCEL                        ExecType = "H"
	ExecType_ORDER_STATUS                        ExecType = "I"
	ExecType_TRADE_IN_A_CLEARING_HOLD            ExecType = "J"
	ExecType_TRADE_HAS_BEEN_RELEASED_TO_CLEARING ExecType = "K"
	ExecType_TRIGGERED_OR_ACTIVATED_BY_SYSTEM    ExecType = "L"
	ExecType_LOCKED                              ExecType = "M"
	ExecType_RELEASED                            ExecType = "N"
)

// ExecTypeReason field enumeration values.
type ExecTypeReason string

const (
	ExecTypeReason_ORDER_ADDED_UPON_REQUEST                           ExecTypeReason = "1"
	ExecTypeReason_ORDER_CANCELLATION_PENDING                         ExecTypeReason = "10"
	ExecTypeReason_PENDING_CANCELLATION_EXECUTED                      ExecTypeReason = "11"
	ExecTypeReason_RESTING_ORDER_TRIGGERED                            ExecTypeReason = "12"
	ExecTypeReason_SUSPENDED_ORDER_ACTIVATED                          ExecTypeReason = "13"
	ExecTypeReason_ACTIVE_ORDER_SUSPENDED                             ExecTypeReason = "14"
	ExecTypeReason_ORDER_EXPIRED                                      ExecTypeReason = "15"
	ExecTypeReason_ORDER_REPLACED_UPON_REQUEST                        ExecTypeReason = "2"
	ExecTypeReason_ORDER_CANCELLED_UPON_REQUEST                       ExecTypeReason = "3"
	ExecTypeReason_UNSOLICITED_ORDER_CANCELLATION                     ExecTypeReason = "4"
	ExecTypeReason_NON_RESTING_ORDER_ADDED_UPON_REQUEST               ExecTypeReason = "5"
	ExecTypeReason_ORDER_REPLACED_WITH_NON_RESTING_ORDER_UPON_REQUEST ExecTypeReason = "6"
	ExecTypeReason_TRIGGER_ORDER_REPLACED_UPON_REQUEST                ExecTypeReason = "7"
	ExecTypeReason_SUSPENDED_ORDER_REPLACED_UPON_REQUEST              ExecTypeReason = "8"
	ExecTypeReason_SUSPENDED_ORDER_CANCELED_UPON_REQUEST              ExecTypeReason = "9"
)

// ExerciseConfirmationMethod field enumeration values.
type ExerciseConfirmationMethod string

const (
	ExerciseConfirmationMethod_NOT_REQUIRED              ExerciseConfirmationMethod = "0"
	ExerciseConfirmationMethod_NON_ELECTRONIC            ExerciseConfirmationMethod = "1"
	ExerciseConfirmationMethod_ELECTRONIC                ExerciseConfirmationMethod = "2"
	ExerciseConfirmationMethod_UNKNOWN_AT_TIME_OF_REPORT ExerciseConfirmationMethod = "3"
)

// ExerciseMethod field enumeration values.
type ExerciseMethod string

const (
	ExerciseMethod_AUTOMATIC ExerciseMethod = "A"
	ExerciseMethod_MANUAL    ExerciseMethod = "M"
)

// ExerciseStyle field enumeration values.
type ExerciseStyle string

const (
	ExerciseStyle_EUROPEAN ExerciseStyle = "0"
	ExerciseStyle_AMERICAN ExerciseStyle = "1"
	ExerciseStyle_BERMUDA  ExerciseStyle = "2"
	ExerciseStyle_OTHER    ExerciseStyle = "99"
)

// ExpType field enumeration values.
type ExpType string

const (
	ExpType_AUTO_EXERCISE           ExpType = "1"
	ExpType_NON_AUTO_EXERCISE       ExpType = "2"
	ExpType_FINAL_WILL_BE_EXERCISED ExpType = "3"
	ExpType_CONTRARY_INTENTION      ExpType = "4"
	ExpType_DIFFERENCE              ExpType = "5"
)

// ExpirationCycle field enumeration values.
type ExpirationCycle string

const (
	ExpirationCycle_EXPIRE_ON_TRADING_SESSION_CLOSE                                                ExpirationCycle = "0"
	ExpirationCycle_EXPIRE_ON_TRADING_SESSION_OPEN                                                 ExpirationCycle = "1"
	ExpirationCycle_TRADING_ELIGIBILITY_EXPIRATION_SPECIFIED_IN_THE_DATE_AND_TIME_FIELDS_EVENTDATE ExpirationCycle = "2"
)

// ExpirationQtyType field enumeration values.
type ExpirationQtyType string

const (
	ExpirationQtyType_AUTO_EXERCISE           ExpirationQtyType = "1"
	ExpirationQtyType_NON_AUTO_EXERCISE       ExpirationQtyType = "2"
	ExpirationQtyType_FINAL_WILL_BE_EXERCISED ExpirationQtyType = "3"
	ExpirationQtyType_CONTRARY_INTENTION      ExpirationQtyType = "4"
	ExpirationQtyType_DIFFERENCE              ExpirationQtyType = "5"
)

// ExtraordinaryEventAdjustmentMethod field enumeration values.
type ExtraordinaryEventAdjustmentMethod string

const (
	ExtraordinaryEventAdjustmentMethod_CALCULATION_AGENT ExtraordinaryEventAdjustmentMethod = "0"
	ExtraordinaryEventAdjustmentMethod_OPTIONS_EXCHANGE  ExtraordinaryEventAdjustmentMethod = "1"
)

// FinancialStatus field enumeration values.
type FinancialStatus string

const (
	FinancialStatus_BANKRUPT          FinancialStatus = "1"
	FinancialStatus_PENDING_DELISTING FinancialStatus = "2"
	FinancialStatus_RESTRICTED        FinancialStatus = "3"
)

// FlowScheduleType field enumeration values.
type FlowScheduleType string

const (
	FlowScheduleType_NERC_EASTERN_OFF_PEAK           FlowScheduleType = "0"
	FlowScheduleType_NERC_WESTERN_OFF_PEAK           FlowScheduleType = "1"
	FlowScheduleType_NERC_CALENDAR_ALL_DAYS_IN_MONTH FlowScheduleType = "2"
	FlowScheduleType_NERC_EASTERN_PEAK               FlowScheduleType = "3"
	FlowScheduleType_NERC_WESTERN_PEAK               FlowScheduleType = "4"
	FlowScheduleType_ALL_TIMES                       FlowScheduleType = "5"
	FlowScheduleType_ON_PEAK                         FlowScheduleType = "6"
	FlowScheduleType_OFF_PEAK                        FlowScheduleType = "7"
	FlowScheduleType_BASE                            FlowScheduleType = "8"
	FlowScheduleType_BLOCK                           FlowScheduleType = "9"
	FlowScheduleType_OTHER                           FlowScheduleType = "99"
)

// ForexReq field enumeration values.
type ForexReq string

const (
	ForexReq_NO  ForexReq = "N"
	ForexReq_YES ForexReq = "Y"
)

// FundRenewWaiv field enumeration values.
type FundRenewWaiv string

const (
	FundRenewWaiv_NO  FundRenewWaiv = "N"
	FundRenewWaiv_YES FundRenewWaiv = "Y"
)

// FundingSource field enumeration values.
type FundingSource string

const (
	FundingSource_REPURCHASE_AGREEMENT FundingSource = "0"
	FundingSource_CASH                 FundingSource = "1"
	FundingSource_FREE_CREDITS         FundingSource = "2"
	FundingSource_CUSTOMER_SHORT_SALES FundingSource = "3"
	FundingSource_BROKER_SHORT_SALES   FundingSource = "4"
	FundingSource_UNSECURED_BORROWING  FundingSource = "5"
	FundingSource_OTHER                FundingSource = "99"
)

// FuturesValuationMethod field enumeration values.
type FuturesValuationMethod string

const (
	FuturesValuationMethod_PREMIUM_STYLE                                  FuturesValuationMethod = "EQTY"
	FuturesValuationMethod_FUTURES_STYLE_MARK_TO_MARKET                   FuturesValuationMethod = "FUT"
	FuturesValuationMethod_FUTURES_STYLE_WITH_AN_ATTACHED_CASH_ADJUSTMENT FuturesValuationMethod = "FUTDA"
)

// GTBookingInst field enumeration values.
type GTBookingInst string

const (
	GTBookingInst_BOOK_OUT_ALL_TRADES_ON_DAY_OF_EXECUTION                GTBookingInst = "0"
	GTBookingInst_ACCUMULATE_EXECUTIONS_UNTIL_ORDER_IS_FILLED_OR_EXPIRES GTBookingInst = "1"
	GTBookingInst_ACCUMULATE_UNTIL_VERBALLY_NOTIFIED_OTHERWISE           GTBookingInst = "2"
)

// GapFillFlag field enumeration values.
type GapFillFlag string

const (
	GapFillFlag_NO  GapFillFlag = "N"
	GapFillFlag_YES GapFillFlag = "Y"
)

// HaltReasonChar field enumeration values.
type HaltReasonChar string

const (
	HaltReasonChar_NEWS_DISSEMINATION     HaltReasonChar = "D"
	HaltReasonChar_ORDER_INFLUX           HaltReasonChar = "E"
	HaltReasonChar_ORDER_IMBALANCE        HaltReasonChar = "I"
	HaltReasonChar_ADDITIONAL_INFORMATION HaltReasonChar = "M"
	HaltReasonChar_NEW_PENDING            HaltReasonChar = "P"
	HaltReasonChar_EQUIPMENT_CHANGEOVER   HaltReasonChar = "X"
)

// HaltReasonInt field enumeration values.
type HaltReasonInt string

const (
	HaltReasonInt_NEWS_DISSEMINATION     HaltReasonInt = "0"
	HaltReasonInt_ORDER_INFLUX           HaltReasonInt = "1"
	HaltReasonInt_ORDER_IMBALANCE        HaltReasonInt = "2"
	HaltReasonInt_ADDITIONAL_INFORMATION HaltReasonInt = "3"
	HaltReasonInt_NEWS_PENDING           HaltReasonInt = "4"
	HaltReasonInt_EQUIPMENT_CHANGEOVER   HaltReasonInt = "5"
)

// HandlInst field enumeration values.
type HandlInst string

const (
	HandlInst_AUTOMATED_EXECUTION_ORDER_PRIVATE_NO_BROKER_INTERVENTION HandlInst = "1"
	HandlInst_AUTOMATED_EXECUTION_ORDER_PUBLIC_BROKER_INTERVENTION_OK  HandlInst = "2"
	HandlInst_MANUAL_ORDER_BEST_EXECUTION                              HandlInst = "3"
)

// IDSource field enumeration values.
type IDSource string

const (
	IDSource_CUSIP                         IDSource = "1"
	IDSource_SEDOL                         IDSource = "2"
	IDSource_QUIK                          IDSource = "3"
	IDSource_ISIN_NUMBER                   IDSource = "4"
	IDSource_RIC_CODE                      IDSource = "5"
	IDSource_ISO_CURRENCY_CODE             IDSource = "6"
	IDSource_ISO_COUNTRY_CODE              IDSource = "7"
	IDSource_EXCHANGE_SYMBOL               IDSource = "8"
	IDSource_CONSOLIDATED_TAPE_ASSOCIATION IDSource = "9"
)

// IOINaturalFlag field enumeration values.
type IOINaturalFlag string

const (
	IOINaturalFlag_NO  IOINaturalFlag = "N"
	IOINaturalFlag_YES IOINaturalFlag = "Y"
)

// IOIOthSvc field enumeration values.
type IOIOthSvc string

const (
	IOIOthSvc_AUTEX  IOIOthSvc = "A"
	IOIOthSvc_BRIDGE IOIOthSvc = "B"
)

// IOIQltyInd field enumeration values.
type IOIQltyInd string

const (
	IOIQltyInd_HIGH   IOIQltyInd = "H"
	IOIQltyInd_LOW    IOIQltyInd = "L"
	IOIQltyInd_MEDIUM IOIQltyInd = "M"
)

// IOIQty field enumeration values.
type IOIQty string

const (
	IOIQty_1000000000           IOIQty = "0"
	IOIQty_LARGE                IOIQty = "L"
	IOIQty_MEDIUM               IOIQty = "M"
	IOIQty_SMALL                IOIQty = "S"
	IOIQty_UNDISCLOSED_QUANTITY IOIQty = "U"
)

// IOIQualifier field enumeration values.
type IOIQualifier string

const (
	IOIQualifier_QUANTITY_IS_NEGOTIABLE       IOIQualifier = "1"
	IOIQualifier_ALLOW_LATE_BIDS              IOIQualifier = "2"
	IOIQualifier_IMMEDIATE_OR_COUNTER         IOIQualifier = "3"
	IOIQualifier_AUTO_TRADE                   IOIQualifier = "4"
	IOIQualifier_ALL_OR_NONE                  IOIQualifier = "A"
	IOIQualifier_MARKET_ON_CLOSE              IOIQualifier = "B"
	IOIQualifier_AT_THE_CLOSE                 IOIQualifier = "C"
	IOIQualifier_VWAP                         IOIQualifier = "D"
	IOIQualifier_AXE                          IOIQualifier = "E"
	IOIQualifier_AXE_ON_BID                   IOIQualifier = "F"
	IOIQualifier_AXE_ON_OFFER                 IOIQualifier = "G"
	IOIQualifier_CLIENT_NATURAL_WORKING       IOIQualifier = "H"
	IOIQualifier_IN_TOUCH_WITH                IOIQualifier = "I"
	IOIQualifier_POSITION_WANTED              IOIQualifier = "J"
	IOIQualifier_MARKET_MAKING                IOIQualifier = "K"
	IOIQualifier_LIMIT                        IOIQualifier = "L"
	IOIQualifier_MORE_BEHIND                  IOIQualifier = "M"
	IOIQualifier_CLIENT_NATURAL_BLOCK         IOIQualifier = "N"
	IOIQualifier_AT_THE_OPEN                  IOIQualifier = "O"
	IOIQualifier_TAKING_A_POSITION            IOIQualifier = "P"
	IOIQualifier_AT_THE_MARKET                IOIQualifier = "Q"
	IOIQualifier_READY_TO_TRADE               IOIQualifier = "R"
	IOIQualifier_INVENTORY_OR_PORTFOLIO_SHOWN IOIQualifier = "S"
	IOIQualifier_THROUGH_THE_DAY              IOIQualifier = "T"
	IOIQualifier_UNWIND                       IOIQualifier = "U"
	IOIQualifier_VERSUS                       IOIQualifier = "V"
	IOIQualifier_INDICATION                   IOIQualifier = "W"
	IOIQualifier_CROSSING_OPPORTUNITY         IOIQualifier = "X"
	IOIQualifier_AT_THE_MIDPOINT              IOIQualifier = "Y"
	IOIQualifier_PRE_OPEN                     IOIQualifier = "Z"
	IOIQualifier_AUTOMATIC_SPOT               IOIQualifier = "a"
	IOIQualifier_PLATFORM_CALCULATED_SPOT     IOIQualifier = "b"
	IOIQualifier_OUTSIDE_SPREAD               IOIQualifier = "c"
	IOIQualifier_DEFERRED_SPOT                IOIQualifier = "d"
	IOIQualifier_NEGOTIATED_SPOT              IOIQualifier = "n"
)

// IOIShares field enumeration values.
type IOIShares string

const (
	IOIShares_LARGE  IOIShares = "L"
	IOIShares_MEDIUM IOIShares = "M"
	IOIShares_SMALL  IOIShares = "S"
)

// IOITransType field enumeration values.
type IOITransType string

const (
	IOITransType_CANCEL  IOITransType = "C"
	IOITransType_NEW     IOITransType = "N"
	IOITransType_REPLACE IOITransType = "R"
)

// IRSDirection field enumeration values.
type IRSDirection string

const (
	IRSDirection_SWAP_IS_FLOAT_FLOAT_OR_FIXED_FIXED IRSDirection = "NA"
	IRSDirection_PRINCIPAL_IS_PAYING_FIXED_RATE     IRSDirection = "PAY"
	IRSDirection_PRINCIPAL_IS_RECEIVING_FIXED_RATE  IRSDirection = "RCV"
)

// ImpliedMarketIndicator field enumeration values.
type ImpliedMarketIndicator string

const (
	ImpliedMarketIndicator_NOT_IMPLIED                     ImpliedMarketIndicator = "0"
	ImpliedMarketIndicator_IMPLIED_IN                      ImpliedMarketIndicator = "1"
	ImpliedMarketIndicator_IMPLIED_OUT                     ImpliedMarketIndicator = "2"
	ImpliedMarketIndicator_BOTH_IMPLIED_IN_AND_IMPLIED_OUT ImpliedMarketIndicator = "3"
)

// InTheMoneyCondition field enumeration values.
type InTheMoneyCondition string

const (
	InTheMoneyCondition_STANDARD_IN_THE_MONEY             InTheMoneyCondition = "0"
	InTheMoneyCondition_AT_THE_MONEY_IS_IN_THE_MONEY      InTheMoneyCondition = "1"
	InTheMoneyCondition_AT_THE_MONEY_CALL_IS_IN_THE_MONEY InTheMoneyCondition = "2"
	InTheMoneyCondition_AT_THE_MONEY_PUT_IS_IN_THE_MONEY  InTheMoneyCondition = "3"
)

// InViewOfCommon field enumeration values.
type InViewOfCommon string

const (
	InViewOfCommon_NO  InViewOfCommon = "N"
	InViewOfCommon_YES InViewOfCommon = "Y"
)

// IncTaxInd field enumeration values.
type IncTaxInd string

const (
	IncTaxInd_NET   IncTaxInd = "1"
	IncTaxInd_GROSS IncTaxInd = "2"
)

// IndividualAllocType field enumeration values.
type IndividualAllocType string

const (
	IndividualAllocType_SUB_ALLOCATE           IndividualAllocType = "1"
	IndividualAllocType_THIRD_PARTY_ALLOCATION IndividualAllocType = "2"
)

// InstrAttribType field enumeration values.
type InstrAttribType string

const (
	InstrAttribType_FLAT                                                        InstrAttribType = "1"
	InstrAttribType_ORIGINAL_ISSUE_DISCOUNT                                     InstrAttribType = "10"
	InstrAttribType_CALLABLE_PUTTABLE                                           InstrAttribType = "11"
	InstrAttribType_ESCROWED_TO_MATURITY                                        InstrAttribType = "12"
	InstrAttribType_ESCROWED_TO_REDEMPTION_DATE                                 InstrAttribType = "13"
	InstrAttribType_PRE_REFUNDED                                                InstrAttribType = "14"
	InstrAttribType_IN_DEFAULT                                                  InstrAttribType = "15"
	InstrAttribType_UNRATED                                                     InstrAttribType = "16"
	InstrAttribType_TAXABLE                                                     InstrAttribType = "17"
	InstrAttribType_INDEXED                                                     InstrAttribType = "18"
	InstrAttribType_SUBJECT_TO_ALTERNATIVE_MINIMUM_TAX                          InstrAttribType = "19"
	InstrAttribType_ZERO_COUPON                                                 InstrAttribType = "2"
	InstrAttribType_ORIGINAL_ISSUE_DISCOUNT_PRICE                               InstrAttribType = "20"
	InstrAttribType_CALLABLE_BELOW_MATURITY_VALUE                               InstrAttribType = "21"
	InstrAttribType_CALLABLE_WITHOUT_NOTICE_BY_MAIL_TO_HOLDER_UNLESS_REGISTERED InstrAttribType = "22"
	InstrAttribType_PRICE_TICK_RULES_FOR_SECURITY                               InstrAttribType = "23"
	InstrAttribType_TRADE_TYPE_ELIGIBILITY_DETAILS_FOR_SECURITY                 InstrAttribType = "24"
	InstrAttribType_INSTRUMENT_DENOMINATOR                                      InstrAttribType = "25"
	InstrAttribType_INSTRUMENT_NUMERATOR                                        InstrAttribType = "26"
	InstrAttribType_INSTRUMENT_PRICE_PRECISION                                  InstrAttribType = "27"
	InstrAttribType_INSTRUMENT_STRIKE_PRICE                                     InstrAttribType = "28"
	InstrAttribType_TRADEABLE_INDICATOR                                         InstrAttribType = "29"
	InstrAttribType_INTEREST_BEARING                                            InstrAttribType = "3"
	InstrAttribType_INSTRUMENT_IS_ELIGIBLE_TO_ACCEPT_ANONYMOUS_ORDERS           InstrAttribType = "30"
	InstrAttribType_MINIMUM_GUARANTEED_FILL_VOLUME                              InstrAttribType = "31"
	InstrAttribType_MINIMUM_GUARANTEED_FILL_STATUS                              InstrAttribType = "32"
	InstrAttribType_TRADE_AT_SETTLEMENT                                         InstrAttribType = "33"
	InstrAttribType_TEST_INSTRUMENT                                             InstrAttribType = "34"
	InstrAttribType_DUMMY_INSTRUMENT                                            InstrAttribType = "35"
	InstrAttribType_NEGATIVE_SETTLEMENT_PRICE_ELIGIBILITY                       InstrAttribType = "36"
	InstrAttribType_NEGATIVE_STRIKE_PRICE_ELIGIBILITY                           InstrAttribType = "37"
	InstrAttribType_US_STANDARD_CONTRACT_INDICATOR                              InstrAttribType = "38"
	InstrAttribType_ADMITTED_TO_TRADING_ON_A_TRADING_VENUE                      InstrAttribType = "39"
	InstrAttribType_NO_PERIODIC_PAYMENTS                                        InstrAttribType = "4"
	InstrAttribType_AVERAGE_DAILY_NOTIONAL_AMOUNT                               InstrAttribType = "40"
	InstrAttribType_AVERAGE_DAILY_NUMBER_OF_TRADES                              InstrAttribType = "41"
	InstrAttribType_VARIABLE_RATE                                               InstrAttribType = "5"
	InstrAttribType_LESS_FEE_FOR_PUT                                            InstrAttribType = "6"
	InstrAttribType_STEPPED_COUPON                                              InstrAttribType = "7"
	InstrAttribType_COUPON_PERIOD                                               InstrAttribType = "8"
	InstrAttribType_WHEN_AND_IF_ISSUED                                          InstrAttribType = "9"
	InstrAttribType_TEXT                                                        InstrAttribType = "99"
)

// InstrRegistry field enumeration values.
type InstrRegistry string

const (
	InstrRegistry_CUSTODIAN InstrRegistry = "BIC"
	InstrRegistry_COUNTRY   InstrRegistry = "ISO"
	InstrRegistry_PHYSICAL  InstrRegistry = "ZZ"
)

// InstrmtAssignmentMethod field enumeration values.
type InstrmtAssignmentMethod string

const (
	InstrmtAssignmentMethod_PRO_RATA InstrmtAssignmentMethod = "P"
	InstrmtAssignmentMethod_RANDOM   InstrmtAssignmentMethod = "R"
)

// InstrumentScopeOperator field enumeration values.
type InstrumentScopeOperator string

const (
	InstrumentScopeOperator_INCLUDE InstrumentScopeOperator = "1"
	InstrumentScopeOperator_EXCLUDE InstrumentScopeOperator = "2"
)

// LastCapacity field enumeration values.
type LastCapacity string

const (
	LastCapacity_AGENT              LastCapacity = "1"
	LastCapacity_CROSS_AS_AGENT     LastCapacity = "2"
	LastCapacity_CROSS_AS_PRINCIPAL LastCapacity = "3"
	LastCapacity_PRINCIPAL          LastCapacity = "4"
	LastCapacity_RISKLESS_PRINCIPAL LastCapacity = "5"
)

// LastFragment field enumeration values.
type LastFragment string

const (
	LastFragment_NO  LastFragment = "N"
	LastFragment_YES LastFragment = "Y"
)

// LastLiquidityInd field enumeration values.
type LastLiquidityInd string

const (
	LastLiquidityInd_NEITHER_ADDED_NOR_REMOVED_LIQUIDITY           LastLiquidityInd = "0"
	LastLiquidityInd_ADDED_LIQUIDITY                               LastLiquidityInd = "1"
	LastLiquidityInd_REMOVED_LIQUIDITY                             LastLiquidityInd = "2"
	LastLiquidityInd_LIQUIDITY_ROUTED_OUT                          LastLiquidityInd = "3"
	LastLiquidityInd_AUCTION_EXECUTION                             LastLiquidityInd = "4"
	LastLiquidityInd_TRIGGERED_STOP_ORDER                          LastLiquidityInd = "5"
	LastLiquidityInd_TRIGGERED_CONTINGENCY_ORDER                   LastLiquidityInd = "6"
	LastLiquidityInd_TRIGGERED_MARKET_ORDER                        LastLiquidityInd = "7"
	LastLiquidityInd_REMOVED_LIQUIDITY_AFTER_FIRM_ORDER_COMMITMENT LastLiquidityInd = "8"
	LastLiquidityInd_AUCTION_EXECUTION_AFTER_FIRM_ORDER_COMMITMENT LastLiquidityInd = "9"
)

// LastRptRequested field enumeration values.
type LastRptRequested string

const (
	LastRptRequested_NO  LastRptRequested = "N"
	LastRptRequested_YES LastRptRequested = "Y"
)

// LegSwapType field enumeration values.
type LegSwapType string

const (
	LegSwapType_PAR_FOR_PAR       LegSwapType = "1"
	LegSwapType_MODIFIED_DURATION LegSwapType = "2"
	LegSwapType_RISK              LegSwapType = "4"
	LegSwapType_PROCEEDS          LegSwapType = "5"
)

// LegalConfirm field enumeration values.
type LegalConfirm string

const (
	LegalConfirm_NO  LegalConfirm = "N"
	LegalConfirm_YES LegalConfirm = "Y"
)

// LienSeniority field enumeration values.
type LienSeniority string

const (
	LienSeniority_UNKNOWN     LienSeniority = "0"
	LienSeniority_FIRST_LIEN  LienSeniority = "1"
	LienSeniority_SECOND_LIEN LienSeniority = "2"
	LienSeniority_THIRD_LIEN  LienSeniority = "3"
)

// LimitAmtType field enumeration values.
type LimitAmtType string

const (
	LimitAmtType_CREDIT_LIMIT         LimitAmtType = "0"
	LimitAmtType_GROSS_POSITION_LIMIT LimitAmtType = "1"
	LimitAmtType_NET_POSITION_LIMIT   LimitAmtType = "2"
	LimitAmtType_RISK_EXPOSURE_LIMIT  LimitAmtType = "3"
	LimitAmtType_LONG_POSITION_LIMIT  LimitAmtType = "4"
	LimitAmtType_SHORT_POSITION_LIMIT LimitAmtType = "5"
)

// LiquidityIndType field enumeration values.
type LiquidityIndType string

const (
	LiquidityIndType_5_DAY_MOVING_AVERAGE  LiquidityIndType = "1"
	LiquidityIndType_20_DAY_MOVING_AVERAGE LiquidityIndType = "2"
	LiquidityIndType_NORMAL_MARKET_SIZE    LiquidityIndType = "3"
	LiquidityIndType_OTHER                 LiquidityIndType = "4"
)

// ListExecInstType field enumeration values.
type ListExecInstType string

const (
	ListExecInstType_IMMEDIATE                   ListExecInstType = "1"
	ListExecInstType_WAIT_FOR_EXECUT_INSTRUCTION ListExecInstType = "2"
	ListExecInstType_EXCHANGE_SWITCH_CIV_ORDER_3 ListExecInstType = "3"
	ListExecInstType_EXCHANGE_SWITCH_CIV_ORDER_4 ListExecInstType = "4"
	ListExecInstType_EXCHANGE_SWITCH_CIV_ORDER_5 ListExecInstType = "5"
)

// ListMethod field enumeration values.
type ListMethod string

const (
	ListMethod_PRE_LISTED_ONLY ListMethod = "0"
	ListMethod_USER_REQUESTED  ListMethod = "1"
)

// ListOrderStatus field enumeration values.
type ListOrderStatus string

const (
	ListOrderStatus_IN_BIDDING_PROCESS     ListOrderStatus = "1"
	ListOrderStatus_RECEIVED_FOR_EXECUTION ListOrderStatus = "2"
	ListOrderStatus_EXECUTING              ListOrderStatus = "3"
	ListOrderStatus_CANCELLING             ListOrderStatus = "4"
	ListOrderStatus_ALERT                  ListOrderStatus = "5"
	ListOrderStatus_ALL_DONE               ListOrderStatus = "6"
	ListOrderStatus_REJECT                 ListOrderStatus = "7"
)

// ListRejectReason field enumeration values.
type ListRejectReason string

const (
	ListRejectReason_BROKER                           ListRejectReason = "0"
	ListRejectReason_UNSUPPORTED_ORDER_CHARACTERISTIC ListRejectReason = "11"
	ListRejectReason_EXCHANGE_CLOSED                  ListRejectReason = "2"
	ListRejectReason_TOO_LATE_TO_ENTER                ListRejectReason = "4"
	ListRejectReason_UNKNOWN_ORDER                    ListRejectReason = "5"
	ListRejectReason_DUPLICATE_ORDER                  ListRejectReason = "6"
	ListRejectReason_OTHER                            ListRejectReason = "99"
)

// ListStatusType field enumeration values.
type ListStatusType string

const (
	ListStatusType_ACK          ListStatusType = "1"
	ListStatusType_RESPONSE     ListStatusType = "2"
	ListStatusType_TIMED        ListStatusType = "3"
	ListStatusType_EXEC_STARTED ListStatusType = "4"
	ListStatusType_ALL_DONE     ListStatusType = "5"
	ListStatusType_ALERT        ListStatusType = "6"
)

// ListUpdateAction field enumeration values.
type ListUpdateAction string

const (
	ListUpdateAction_ADD      ListUpdateAction = "A"
	ListUpdateAction_DELETE   ListUpdateAction = "D"
	ListUpdateAction_MODIFY   ListUpdateAction = "M"
	ListUpdateAction_SNAPSHOT ListUpdateAction = "S"
)

// LoanFacility field enumeration values.
type LoanFacility string

const (
	LoanFacility_BRIDGE_LOAN       LoanFacility = "0"
	LoanFacility_LETTER_OF_CREDIT  LoanFacility = "1"
	LoanFacility_REVOLVING_LOAN    LoanFacility = "2"
	LoanFacility_SWINGLINE_FUNDING LoanFacility = "3"
	LoanFacility_TERM_LOAN         LoanFacility = "4"
	LoanFacility_TRADE_CLAIM       LoanFacility = "5"
)

// LocateReqd field enumeration values.
type LocateReqd string

const (
	LocateReqd_NO  LocateReqd = "N"
	LocateReqd_YES LocateReqd = "Y"
)

// LockType field enumeration values.
type LockType string

const (
	LockType_NOT_LOCKED             LockType = "0"
	LockType_AWAY_MARKET_BETTER     LockType = "1"
	LockType_THREE_TICK_LOCKED      LockType = "2"
	LockType_LOCKED_BY_MARKET_MAKER LockType = "3"
	LockType_DIRECTED_ORDER_LOCK    LockType = "4"
	LockType_MULTILEG_LOCK          LockType = "5"
	LockType_MARKET_ORDER_LOCK      LockType = "6"
	LockType_PRE_ASSIGNMENT_LOCK    LockType = "7"
)

// LotType field enumeration values.
type LotType string

const (
	LotType_ODD_LOT                            LotType = "1"
	LotType_ROUND_LOT                          LotType = "2"
	LotType_BLOCK_LOT                          LotType = "3"
	LotType_ROUND_LOT_BASED_UPON_UNITOFMEASURE LotType = "4"
)

// MDBookType field enumeration values.
type MDBookType string

const (
	MDBookType_TOP_OF_BOOK MDBookType = "1"
	MDBookType_PRICE_DEPTH MDBookType = "2"
	MDBookType_ORDER_DEPTH MDBookType = "3"
)

// MDEntryType field enumeration values.
type MDEntryType string

const (
	MDEntryType_BID                                             MDEntryType = "0"
	MDEntryType_OFFER                                           MDEntryType = "1"
	MDEntryType_TRADE                                           MDEntryType = "2"
	MDEntryType_INDEX_VALUE                                     MDEntryType = "3"
	MDEntryType_OPENING_PRICE                                   MDEntryType = "4"
	MDEntryType_CLOSING_PRICE                                   MDEntryType = "5"
	MDEntryType_SETTLEMENT_PRICE                                MDEntryType = "6"
	MDEntryType_TRADING_SESSION_HIGH_PRICE                      MDEntryType = "7"
	MDEntryType_TRADING_SESSION_LOW_PRICE                       MDEntryType = "8"
	MDEntryType_VOLUME_WEIGHTED_AVERAGE_PRICE                   MDEntryType = "9"
	MDEntryType_IMBALANCE                                       MDEntryType = "A"
	MDEntryType_TRADE_VOLUME                                    MDEntryType = "B"
	MDEntryType_OPEN_INTEREST                                   MDEntryType = "C"
	MDEntryType_COMPOSITE_UNDERLYING_PRICE                      MDEntryType = "D"
	MDEntryType_SIMULATED_SELL_PRICE                            MDEntryType = "E"
	MDEntryType_SIMULATED_BUY_PRICE                             MDEntryType = "F"
	MDEntryType_MARGIN_RATE                                     MDEntryType = "G"
	MDEntryType_MID_PRICE                                       MDEntryType = "H"
	MDEntryType_EMPTY_BOOK                                      MDEntryType = "J"
	MDEntryType_SETTLE_HIGH_PRICE                               MDEntryType = "K"
	MDEntryType_SETTLE_LOW_PRICE                                MDEntryType = "L"
	MDEntryType_PRIOR_SETTLE_PRICE                              MDEntryType = "M"
	MDEntryType_SESSION_HIGH_BID                                MDEntryType = "N"
	MDEntryType_SESSION_LOW_OFFER                               MDEntryType = "O"
	MDEntryType_EARLY_PRICES                                    MDEntryType = "P"
	MDEntryType_AUCTION_CLEARING_PRICE                          MDEntryType = "Q"
	MDEntryType_DAILY_VALUE_ADJUSTMENT_FOR_LONG_POSITIONS       MDEntryType = "R"
	MDEntryType_SWAP_VALUE_FACTOR                               MDEntryType = "S"
	MDEntryType_CUMULATIVE_VALUE_ADJUSTMENT_FOR_LONG_POSITIONS  MDEntryType = "T"
	MDEntryType_DAILY_VALUE_ADJUSTMENT_FOR_SHORT_POSITIONS      MDEntryType = "U"
	MDEntryType_CUMULATIVE_VALUE_ADJUSTMENT_FOR_SHORT_POSITIONS MDEntryType = "V"
	MDEntryType_FIXING_PRICE                                    MDEntryType = "W"
	MDEntryType_CASH_RATE                                       MDEntryType = "X"
	MDEntryType_RECOVERY_RATE                                   MDEntryType = "Y"
	MDEntryType_RECOVERY_RATE_FOR_LONG_POSITIONS                MDEntryType = "Z"
	MDEntryType_RECOVERY_RATE_FOR_SHORT_POSITIONS               MDEntryType = "a"
	MDEntryType_MARKET_BID                                      MDEntryType = "b"
	MDEntryType_MARKET_OFFER                                    MDEntryType = "c"
	MDEntryType_SHORT_SALE_MINIMUM_PRICE                        MDEntryType = "d"
	MDEntryType_PREVIOUS_CLOSING_PRICE                          MDEntryType = "e"
	MDEntryType_THRESHOLD_LIMITS_AND_PRICE_BANDING              MDEntryType = "g"
	MDEntryType_DAILY_FINANCING_VALUE                           MDEntryType = "h"
	MDEntryType_ACCRUED_FINANCING_VALUE                         MDEntryType = "i"
	MDEntryType_TIME_WEIGHTED_AVERAGE_PRICE                     MDEntryType = "t"
)

// MDImplicitDelete field enumeration values.
type MDImplicitDelete string

const (
	MDImplicitDelete_NO  MDImplicitDelete = "N"
	MDImplicitDelete_YES MDImplicitDelete = "Y"
)

// MDOriginType field enumeration values.
type MDOriginType string

const (
	MDOriginType_BOOK                  MDOriginType = "0"
	MDOriginType_OFF_BOOK              MDOriginType = "1"
	MDOriginType_CROSS                 MDOriginType = "2"
	MDOriginType_QUOTE_DRIVEN_MARKET   MDOriginType = "3"
	MDOriginType_DARK_ORDER_BOOK       MDOriginType = "4"
	MDOriginType_AUCTION_DRIVEN_MARKET MDOriginType = "5"
	MDOriginType_QUOTE_NEGOTIATION     MDOriginType = "6"
	MDOriginType_VOICE_NEGOTIATION     MDOriginType = "7"
	MDOriginType_HYBRID_MARKET         MDOriginType = "8"
)

// MDQuoteType field enumeration values.
type MDQuoteType string

const (
	MDQuoteType_INDICATIVE               MDQuoteType = "0"
	MDQuoteType_TRADEABLE                MDQuoteType = "1"
	MDQuoteType_RESTRICTED_TRADEABLE     MDQuoteType = "2"
	MDQuoteType_COUNTER                  MDQuoteType = "3"
	MDQuoteType_INDICATIVE_AND_TRADEABLE MDQuoteType = "4"
)

// MDReportEvent field enumeration values.
type MDReportEvent string

const (
	MDReportEvent_START_OF_INSTRUMENT_REFERENCE_DATA MDReportEvent = "1"
	MDReportEvent_END_OF_SETTLEMENT_PRICES           MDReportEvent = "10"
	MDReportEvent_START_OF_STATISTICS_REFERENCE_DATA MDReportEvent = "11"
	MDReportEvent_END_OF_STATISTICS_REFERENCE_DATA   MDReportEvent = "12"
	MDReportEvent_START_OF_STATISTICS                MDReportEvent = "13"
	MDReportEvent_END_OF_STATISTICS                  MDReportEvent = "14"
	MDReportEvent_END_OF_INSTRUMENT_REFERENCE_DATA   MDReportEvent = "2"
	MDReportEvent_START_OF_OFF_MARKET_TRADES         MDReportEvent = "3"
	MDReportEvent_END_OF_OFF_MARKET_TRADES           MDReportEvent = "4"
	MDReportEvent_START_OF_ORDER_BOOK_TRADES         MDReportEvent = "5"
	MDReportEvent_END_OF_ORDER_BOOK_TRADES           MDReportEvent = "6"
	MDReportEvent_START_OF_OPEN_INTEREST             MDReportEvent = "7"
	MDReportEvent_END_OF_OPEN_INTEREST               MDReportEvent = "8"
	MDReportEvent_START_OF_SETTLEMENT_PRICES         MDReportEvent = "9"
)

// MDReqRejReason field enumeration values.
type MDReqRejReason string

const (
	MDReqRejReason_UNKNOWN_SYMBOL                      MDReqRejReason = "0"
	MDReqRejReason_DUPLICATE_MDREQID                   MDReqRejReason = "1"
	MDReqRejReason_INSUFFICIENT_BANDWIDTH              MDReqRejReason = "2"
	MDReqRejReason_INSUFFICIENT_PERMISSIONS            MDReqRejReason = "3"
	MDReqRejReason_UNSUPPORTED_SUBSCRIPTIONREQUESTTYPE MDReqRejReason = "4"
	MDReqRejReason_UNSUPPORTED_MARKETDEPTH             MDReqRejReason = "5"
	MDReqRejReason_UNSUPPORTED_MDUPDATETYPE            MDReqRejReason = "6"
	MDReqRejReason_UNSUPPORTED_AGGREGATEDBOOK          MDReqRejReason = "7"
	MDReqRejReason_UNSUPPORTED_MDENTRYTYPE             MDReqRejReason = "8"
	MDReqRejReason_UNSUPPORTED_TRADINGSESSIONID        MDReqRejReason = "9"
	MDReqRejReason_UNSUPPORTED_SCOPE                   MDReqRejReason = "A"
	MDReqRejReason_UNSUPPORTED_OPENCLOSESETTLEFLAG     MDReqRejReason = "B"
	MDReqRejReason_UNSUPPORTED_MDIMPLICITDELETE        MDReqRejReason = "C"
	MDReqRejReason_INSUFFICIENT_CREDIT                 MDReqRejReason = "D"
)

// MDSecSizeType field enumeration values.
type MDSecSizeType string

const (
	MDSecSizeType_CUSTOMER              MDSecSizeType = "1"
	MDSecSizeType_CUSTOMER_PROFESSIONAL MDSecSizeType = "2"
	MDSecSizeType_DO_NOT_TRADE_THROUGH  MDSecSizeType = "3"
)

// MDStatisticIntervalType field enumeration values.
type MDStatisticIntervalType string

const (
	MDStatisticIntervalType_SLIDING_WINDOW                         MDStatisticIntervalType = "1"
	MDStatisticIntervalType_SLIDING_WINDOW_PEAK                    MDStatisticIntervalType = "2"
	MDStatisticIntervalType_FIXED_DATE_RANGE                       MDStatisticIntervalType = "3"
	MDStatisticIntervalType_FIXED_TIME_RANGE                       MDStatisticIntervalType = "4"
	MDStatisticIntervalType_CURRENT_TIME_UNIT                      MDStatisticIntervalType = "5"
	MDStatisticIntervalType_PREVIOUS_TIME_UNIT                     MDStatisticIntervalType = "6"
	MDStatisticIntervalType_MAXIMUM_RANGE                          MDStatisticIntervalType = "7"
	MDStatisticIntervalType_MAXIMUM_RANGE_UP_TO_PREVIOUS_TIME_UNIT MDStatisticIntervalType = "8"
)

// MDStatisticRatioType field enumeration values.
type MDStatisticRatioType string

const (
	MDStatisticRatioType_BUYERS_TO_SELLERS                   MDStatisticRatioType = "1"
	MDStatisticRatioType_FEES_TO_TOTAL_TRADED_VALUE          MDStatisticRatioType = "10"
	MDStatisticRatioType_TRADE_VOLUME_TO_TOTAL_TRADED_VOLUME MDStatisticRatioType = "11"
	MDStatisticRatioType_ORDERS_TO_TOTAL_NUMBER_OF_ORDERS    MDStatisticRatioType = "12"
	MDStatisticRatioType_UPTICKS_TO_DOWNTICKS                MDStatisticRatioType = "2"
	MDStatisticRatioType_MARKET_MAKER_TO_NON_MARKET_MAKER    MDStatisticRatioType = "3"
	MDStatisticRatioType_AUTOMATED_TO_NON_AUTOMATED          MDStatisticRatioType = "4"
	MDStatisticRatioType_ORDERS_TO_TRADES                    MDStatisticRatioType = "5"
	MDStatisticRatioType_QUOTES_TO_TRADES                    MDStatisticRatioType = "6"
	MDStatisticRatioType_ORDERS_AND_QUOTES_TO_TRADES         MDStatisticRatioType = "7"
	MDStatisticRatioType_FAILED_TO_TOTAL_TRADED_VALUE        MDStatisticRatioType = "8"
	MDStatisticRatioType_BENEFITS_TO_TOTAL_TRADED_VALUE      MDStatisticRatioType = "9"
)

// MDStatisticRequestResult field enumeration values.
type MDStatisticRequestResult string

const (
	MDStatisticRequestResult_SUCCESSFUL                            MDStatisticRequestResult = "0"
	MDStatisticRequestResult_INVALID_OR_UNKNOWN_MARKET             MDStatisticRequestResult = "1"
	MDStatisticRequestResult_MARKET_DEPTH_NOT_SUPPORTED            MDStatisticRequestResult = "10"
	MDStatisticRequestResult_FREQUENCY_NOT_SUPPORTED               MDStatisticRequestResult = "11"
	MDStatisticRequestResult_STATISTIC_INTERVAL_NOT_SUPPORTED      MDStatisticRequestResult = "12"
	MDStatisticRequestResult_STATISTIC_DATE_RANGE_NOT_SUPPORTED    MDStatisticRequestResult = "13"
	MDStatisticRequestResult_STATISTIC_TIME_RANGE_NOT_SUPPORTED    MDStatisticRequestResult = "14"
	MDStatisticRequestResult_RATIO_TYPE_NOT_SUPPORTED              MDStatisticRequestResult = "15"
	MDStatisticRequestResult_INVALID_OR_UNKNOWN_TRADE_INPUT_SOURCE MDStatisticRequestResult = "16"
	MDStatisticRequestResult_INVALID_OR_UNKNOWN_TRADING_SESSION    MDStatisticRequestResult = "17"
	MDStatisticRequestResult_UNAUTHORIZED_FOR_STATISTIC_REQUEST    MDStatisticRequestResult = "18"
	MDStatisticRequestResult_INVALID_OR_UNKNOWN_MARKET_SEGMENT     MDStatisticRequestResult = "2"
	MDStatisticRequestResult_INVALID_OR_UNKNOWN_SECURITY_LIST      MDStatisticRequestResult = "3"
	MDStatisticRequestResult_INVALID_OR_UNKNOWN_INSTRUMENT         MDStatisticRequestResult = "4"
	MDStatisticRequestResult_INVALID_PARTIES                       MDStatisticRequestResult = "5"
	MDStatisticRequestResult_TRADE_DATE_OUT_OF_SUPPORTED_RANGE     MDStatisticRequestResult = "6"
	MDStatisticRequestResult_STATISTIC_TYPE_NOT_SUPPORTED          MDStatisticRequestResult = "7"
	MDStatisticRequestResult_SCOPE_OR_SUB_SCOPE_NOT_SUPPORTED      MDStatisticRequestResult = "8"
	MDStatisticRequestResult_SCOPE_TYPE_NOT_SUPPORTED              MDStatisticRequestResult = "9"
	MDStatisticRequestResult_OTHER                                 MDStatisticRequestResult = "99"
)

// MDStatisticScope field enumeration values.
type MDStatisticScope string

const (
	MDStatisticScope_BID_PRICES            MDStatisticScope = "1"
	MDStatisticScope_AUCTION_PRICES        MDStatisticScope = "10"
	MDStatisticScope_OPENING_PRICES        MDStatisticScope = "11"
	MDStatisticScope_CLOSING_PRICES        MDStatisticScope = "12"
	MDStatisticScope_SETTLEMENT_PRICES     MDStatisticScope = "13"
	MDStatisticScope_UNDERLYING_PRICES     MDStatisticScope = "14"
	MDStatisticScope_OPEN_INTEREST         MDStatisticScope = "15"
	MDStatisticScope_INDEX_VALUES          MDStatisticScope = "16"
	MDStatisticScope_MARGIN_RATES          MDStatisticScope = "17"
	MDStatisticScope_OUTAGES               MDStatisticScope = "18"
	MDStatisticScope_SCHEDULED_AUCTIONS    MDStatisticScope = "19"
	MDStatisticScope_OFFER_PRICES          MDStatisticScope = "2"
	MDStatisticScope_REFERENCE_PRICES      MDStatisticScope = "20"
	MDStatisticScope_TRADE_VALUE           MDStatisticScope = "21"
	MDStatisticScope_MARKET_DATA_FEE_ITEMS MDStatisticScope = "22"
	MDStatisticScope_REBATES               MDStatisticScope = "23"
	MDStatisticScope_DISCOUNTS             MDStatisticScope = "24"
	MDStatisticScope_PAYMENTS              MDStatisticScope = "25"
	MDStatisticScope_TAXES                 MDStatisticScope = "26"
	MDStatisticScope_LEVIES                MDStatisticScope = "27"
	MDStatisticScope_BENEFITS              MDStatisticScope = "28"
	MDStatisticScope_FEES                  MDStatisticScope = "29"
	MDStatisticScope_BID_DEPTH             MDStatisticScope = "3"
	MDStatisticScope_ORDERS_AND_RFQS       MDStatisticScope = "30"
	MDStatisticScope_MARKET_MAKERS         MDStatisticScope = "31"
	MDStatisticScope_TRADING_INTERRUPTIONS MDStatisticScope = "32"
	MDStatisticScope_TRADING_SUSPENSIONS   MDStatisticScope = "33"
	MDStatisticScope_NO_QUOTES             MDStatisticScope = "34"
	MDStatisticScope_REQUEST_FOR_QUOTES    MDStatisticScope = "35"
	MDStatisticScope_TRADE_VOLUME          MDStatisticScope = "36"
	MDStatisticScope_OFFER_DEPTH           MDStatisticScope = "4"
	MDStatisticScope_ORDERS                MDStatisticScope = "5"
	MDStatisticScope_QUOTES                MDStatisticScope = "6"
	MDStatisticScope_ORDERS_AND_QUOTES     MDStatisticScope = "7"
	MDStatisticScope_TRADES                MDStatisticScope = "8"
	MDStatisticScope_TRADE_PRICES          MDStatisticScope = "9"
)

// MDStatisticScopeType field enumeration values.
type MDStatisticScopeType string

const (
	MDStatisticScopeType_ENTRY_RATE        MDStatisticScopeType = "1"
	MDStatisticScopeType_MODIFICATION_RATE MDStatisticScopeType = "2"
	MDStatisticScopeType_CANCEL_RATE       MDStatisticScopeType = "3"
	MDStatisticScopeType_DOWNWARD_MOVE     MDStatisticScopeType = "4"
	MDStatisticScopeType_UPWARD_MOVE       MDStatisticScopeType = "5"
)

// MDStatisticStatus field enumeration values.
type MDStatisticStatus string

const (
	MDStatisticStatus_ACTIVE   MDStatisticStatus = "1"
	MDStatisticStatus_INACTIVE MDStatisticStatus = "2"
)

// MDStatisticSubScope field enumeration values.
type MDStatisticSubScope string

const (
	MDStatisticSubScope_VISIBLE                           MDStatisticSubScope = "1"
	MDStatisticSubScope_NETWORK_ERROR                     MDStatisticSubScope = "10"
	MDStatisticSubScope_FAILED                            MDStatisticSubScope = "11"
	MDStatisticSubScope_EXECUTED                          MDStatisticSubScope = "12"
	MDStatisticSubScope_ENTERED                           MDStatisticSubScope = "13"
	MDStatisticSubScope_MODIFIED                          MDStatisticSubScope = "14"
	MDStatisticSubScope_CANCELLED                         MDStatisticSubScope = "15"
	MDStatisticSubScope_MARKET_DATA_ACCESS                MDStatisticSubScope = "16"
	MDStatisticSubScope_TERMINAL_ACCESS                   MDStatisticSubScope = "17"
	MDStatisticSubScope_VOLUME                            MDStatisticSubScope = "18"
	MDStatisticSubScope_CLEARED                           MDStatisticSubScope = "19"
	MDStatisticSubScope_HIDDEN                            MDStatisticSubScope = "2"
	MDStatisticSubScope_SETTLED                           MDStatisticSubScope = "20"
	MDStatisticSubScope_OTHER                             MDStatisticSubScope = "21"
	MDStatisticSubScope_MONETARY                          MDStatisticSubScope = "22"
	MDStatisticSubScope_NON_MONETARY                      MDStatisticSubScope = "23"
	MDStatisticSubScope_GROSS                             MDStatisticSubScope = "24"
	MDStatisticSubScope_LARGE_IN_SCALE                    MDStatisticSubScope = "25"
	MDStatisticSubScope_NEITHER_HIDDEN_NOR_LARGE_IN_SCALE MDStatisticSubScope = "26"
	MDStatisticSubScope_CORPORATE_ACTION                  MDStatisticSubScope = "27"
	MDStatisticSubScope_VENUE_DECISION                    MDStatisticSubScope = "28"
	MDStatisticSubScope_MINIMUM_TIME_PERIOD               MDStatisticSubScope = "29"
	MDStatisticSubScope_INDICATIVE                        MDStatisticSubScope = "3"
	MDStatisticSubScope_OPEN                              MDStatisticSubScope = "30"
	MDStatisticSubScope_NOT_EXECUTED                      MDStatisticSubScope = "31"
	MDStatisticSubScope_AGGRESSIVE                        MDStatisticSubScope = "32"
	MDStatisticSubScope_DIRECTED                          MDStatisticSubScope = "33"
	MDStatisticSubScope_TRADEABLE                         MDStatisticSubScope = "4"
	MDStatisticSubScope_PASSIVE                           MDStatisticSubScope = "5"
	MDStatisticSubScope_MARKET_CONSENSUS                  MDStatisticSubScope = "6"
	MDStatisticSubScope_POWER                             MDStatisticSubScope = "7"
	MDStatisticSubScope_HARDWARE_ERROR                    MDStatisticSubScope = "8"
	MDStatisticSubScope_SOFTWARE_ERROR                    MDStatisticSubScope = "9"
)

// MDStatisticType field enumeration values.
type MDStatisticType string

const (
	MDStatisticType_COUNT                         MDStatisticType = "1"
	MDStatisticType_TICK                          MDStatisticType = "10"
	MDStatisticType_AVERAGE_VALUE                 MDStatisticType = "11"
	MDStatisticType_TOTAL_VALUE                   MDStatisticType = "12"
	MDStatisticType_HIGH                          MDStatisticType = "13"
	MDStatisticType_LOW                           MDStatisticType = "14"
	MDStatisticType_MIDPOINT                      MDStatisticType = "15"
	MDStatisticType_FIRST                         MDStatisticType = "16"
	MDStatisticType_LAST                          MDStatisticType = "17"
	MDStatisticType_FINAL                         MDStatisticType = "18"
	MDStatisticType_EXCHANGE_BEST                 MDStatisticType = "19"
	MDStatisticType_AVERAGE_VOLUME                MDStatisticType = "2"
	MDStatisticType_EXCHANGE_BEST_WITH_VOLUME     MDStatisticType = "20"
	MDStatisticType_CONSOLIDATED_BEST             MDStatisticType = "21"
	MDStatisticType_CONSOLIDATED_BEST_WITH_VOLUME MDStatisticType = "22"
	MDStatisticType_TIME_WEIGHTED_AVERAGE_PRICE   MDStatisticType = "23"
	MDStatisticType_AVERAGE_DURATION              MDStatisticType = "24"
	MDStatisticType_AVERAGE_PRICE                 MDStatisticType = "25"
	MDStatisticType_TOTAL_FEES                    MDStatisticType = "26"
	MDStatisticType_TOTAL_BENEFITS                MDStatisticType = "27"
	MDStatisticType_MEDIAN_VALUE                  MDStatisticType = "28"
	MDStatisticType_AVERAGE_LIQUIDITY             MDStatisticType = "29"
	MDStatisticType_TOTAL_VOLUME                  MDStatisticType = "3"
	MDStatisticType_MEDIAN_DURATION               MDStatisticType = "30"
	MDStatisticType_DISTRIBUTION                  MDStatisticType = "4"
	MDStatisticType_RATIO                         MDStatisticType = "5"
	MDStatisticType_LIQUIDITY                     MDStatisticType = "6"
	MDStatisticType_VOLUME_WEIGHTED_AVERAGE_PRICE MDStatisticType = "7"
	MDStatisticType_VOLATILITY                    MDStatisticType = "8"
	MDStatisticType_DURATION                      MDStatisticType = "9"
)

// MDStatisticValueType field enumeration values.
type MDStatisticValueType string

const (
	MDStatisticValueType_ABSOLUTE   MDStatisticValueType = "1"
	MDStatisticValueType_PERCENTAGE MDStatisticValueType = "2"
)

// MDUpdateAction field enumeration values.
type MDUpdateAction string

const (
	MDUpdateAction_NEW         MDUpdateAction = "0"
	MDUpdateAction_CHANGE      MDUpdateAction = "1"
	MDUpdateAction_DELETE      MDUpdateAction = "2"
	MDUpdateAction_DELETE_THRU MDUpdateAction = "3"
	MDUpdateAction_DELETE_FROM MDUpdateAction = "4"
	MDUpdateAction_OVERLAY     MDUpdateAction = "5"
)

// MDUpdateType field enumeration values.
type MDUpdateType string

const (
	MDUpdateType_FULL_REFRESH        MDUpdateType = "0"
	MDUpdateType_INCREMENTAL_REFRESH MDUpdateType = "1"
)

// MDValueTier field enumeration values.
type MDValueTier string

const (
	MDValueTier_RANGE_1 MDValueTier = "1"
	MDValueTier_RANGE_2 MDValueTier = "2"
	MDValueTier_RANGE_3 MDValueTier = "3"
)

// MarginAmtType field enumeration values.
type MarginAmtType string

const (
	MarginAmtType_ADDITIONAL_MARGIN           MarginAmtType = "1"
	MarginAmtType_FUTURES_SPREAD_MARGIN       MarginAmtType = "10"
	MarginAmtType_INITIAL_MARGIN              MarginAmtType = "11"
	MarginAmtType_LIQUIDATING_MARGIN          MarginAmtType = "12"
	MarginAmtType_MARGIN_CALL_AMOUNT          MarginAmtType = "13"
	MarginAmtType_MARGIN_DEFICIT_AMOUNT       MarginAmtType = "14"
	MarginAmtType_MARGIN_EXCESS_AMOUNT        MarginAmtType = "15"
	MarginAmtType_OPTION_PREMIUM_AMOUNT       MarginAmtType = "16"
	MarginAmtType_PREMIUM_MARGIN              MarginAmtType = "17"
	MarginAmtType_RESERVE_MARGIN              MarginAmtType = "18"
	MarginAmtType_SECURITY_COLLATERAL_AMOUNT  MarginAmtType = "19"
	MarginAmtType_ADJUSTED_MARGIN             MarginAmtType = "2"
	MarginAmtType_STRESS_TEST_ADD_ON_AMOUNT   MarginAmtType = "20"
	MarginAmtType_SUPER_MARGIN                MarginAmtType = "21"
	MarginAmtType_TOTAL_MARGIN                MarginAmtType = "22"
	MarginAmtType_VARIATION_MARGIN            MarginAmtType = "23"
	MarginAmtType_SECONDARY_VARIATION_MARGIN  MarginAmtType = "24"
	MarginAmtType_ROLLED_UP_MARGIN_DEFICIT    MarginAmtType = "25"
	MarginAmtType_SPREAD_RESPONSE_MARGIN      MarginAmtType = "26"
	MarginAmtType_SYSTEMIC_RISK_MARGIN        MarginAmtType = "27"
	MarginAmtType_CURVE_RISK_MARGIN           MarginAmtType = "28"
	MarginAmtType_INDEX_SPREAD_RISK_MARGIN    MarginAmtType = "29"
	MarginAmtType_UNADJUSTED_MARGIN           MarginAmtType = "3"
	MarginAmtType_SECTOR_RISK_MARGIN          MarginAmtType = "30"
	MarginAmtType_JUMP_TO_DEFAULT_RISK_MARGIN MarginAmtType = "31"
	MarginAmtType_BASIS_RISK_MARGIN           MarginAmtType = "32"
	MarginAmtType_INTEREST_RATE_RISK_MARGIN   MarginAmtType = "33"
	MarginAmtType_JUMP_TO_HEALTH_RISK_MARGIN  MarginAmtType = "34"
	MarginAmtType_OTHER_RISK_MARGIN           MarginAmtType = "35"
	MarginAmtType_BINARY_ADD_ON_AMOUNT        MarginAmtType = "4"
	MarginAmtType_CASH_BALANCE_AMOUNT         MarginAmtType = "5"
	MarginAmtType_CONCENTRATION_MARGIN        MarginAmtType = "6"
	MarginAmtType_CORE_MARGIN                 MarginAmtType = "7"
	MarginAmtType_DELIVERY_MARGIN             MarginAmtType = "8"
	MarginAmtType_DISCRETIONARY_MARGIN        MarginAmtType = "9"
)

// MarginDirection field enumeration values.
type MarginDirection string

const (
	MarginDirection_POSTED   MarginDirection = "0"
	MarginDirection_RECEIVED MarginDirection = "1"
)

// MarginReqmtInqQualifier field enumeration values.
type MarginReqmtInqQualifier string

const (
	MarginReqmtInqQualifier_SUMMARY        MarginReqmtInqQualifier = "0"
	MarginReqmtInqQualifier_DETAIL         MarginReqmtInqQualifier = "1"
	MarginReqmtInqQualifier_EXCESS_DEFICIT MarginReqmtInqQualifier = "2"
	MarginReqmtInqQualifier_NET_POSITION   MarginReqmtInqQualifier = "3"
)

// MarginReqmtInqResult field enumeration values.
type MarginReqmtInqResult string

const (
	MarginReqmtInqResult_SUCCESSFUL                                         MarginReqmtInqResult = "0"
	MarginReqmtInqResult_INVALID_OR_UNKNOWN_INSTRUMENT                      MarginReqmtInqResult = "1"
	MarginReqmtInqResult_INVALID_OR_UNKNOWN_MARGIN_CLASS                    MarginReqmtInqResult = "2"
	MarginReqmtInqResult_INVALID_PARTIES                                    MarginReqmtInqResult = "3"
	MarginReqmtInqResult_INVALID_TRANSPORT_TYPE_REQUESTED                   MarginReqmtInqResult = "4"
	MarginReqmtInqResult_INVALID_DESTINATION_REQUESTED                      MarginReqmtInqResult = "5"
	MarginReqmtInqResult_NO_MARGIN_REQUIREMENT_FOUND                        MarginReqmtInqResult = "6"
	MarginReqmtInqResult_MARGIN_REQUIREMENT_INQUIRY_QUALIFIER_NOT_SUPPORTED MarginReqmtInqResult = "7"
	MarginReqmtInqResult_UNAUTHORIZED_FOR_MARGIN_REQUIREMENT_INQUIRY        MarginReqmtInqResult = "8"
	MarginReqmtInqResult_OTHER                                              MarginReqmtInqResult = "99"
)

// MarginReqmtRptType field enumeration values.
type MarginReqmtRptType string

const (
	MarginReqmtRptType_SUMMARY        MarginReqmtRptType = "0"
	MarginReqmtRptType_DETAIL         MarginReqmtRptType = "1"
	MarginReqmtRptType_EXCESS_DEFICIT MarginReqmtRptType = "2"
)

// MarketCondition field enumeration values.
type MarketCondition string

const (
	MarketCondition_NORMAL      MarketCondition = "0"
	MarketCondition_STRESSED    MarketCondition = "1"
	MarketCondition_EXCEPTIONAL MarketCondition = "2"
)

// MarketDisruptionFallbackProvision field enumeration values.
type MarketDisruptionFallbackProvision string

const (
	MarketDisruptionFallbackProvision_AS_SPECIFIED_IN_MASTER_AGREEMENT MarketDisruptionFallbackProvision = "0"
	MarketDisruptionFallbackProvision_AS_SPECIFIED_IN_CONFIRMATION     MarketDisruptionFallbackProvision = "1"
)

// MarketDisruptionFallbackUnderlierType field enumeration values.
type MarketDisruptionFallbackUnderlierType string

const (
	MarketDisruptionFallbackUnderlierType_BASKET               MarketDisruptionFallbackUnderlierType = "0"
	MarketDisruptionFallbackUnderlierType_BOND                 MarketDisruptionFallbackUnderlierType = "1"
	MarketDisruptionFallbackUnderlierType_MORTGAGE             MarketDisruptionFallbackUnderlierType = "10"
	MarketDisruptionFallbackUnderlierType_MUTUAL_FUND          MarketDisruptionFallbackUnderlierType = "11"
	MarketDisruptionFallbackUnderlierType_CASH                 MarketDisruptionFallbackUnderlierType = "2"
	MarketDisruptionFallbackUnderlierType_COMMODITY            MarketDisruptionFallbackUnderlierType = "3"
	MarketDisruptionFallbackUnderlierType_CONVERTIBLE_BOND     MarketDisruptionFallbackUnderlierType = "4"
	MarketDisruptionFallbackUnderlierType_EQUITY               MarketDisruptionFallbackUnderlierType = "5"
	MarketDisruptionFallbackUnderlierType_EXCHANGE_TRADED_FUND MarketDisruptionFallbackUnderlierType = "6"
	MarketDisruptionFallbackUnderlierType_FUTURE               MarketDisruptionFallbackUnderlierType = "7"
	MarketDisruptionFallbackUnderlierType_INDEX                MarketDisruptionFallbackUnderlierType = "8"
	MarketDisruptionFallbackUnderlierType_LOAN                 MarketDisruptionFallbackUnderlierType = "9"
)

// MarketDisruptionProvision field enumeration values.
type MarketDisruptionProvision string

const (
	MarketDisruptionProvision_NOT_APPLICABLE                   MarketDisruptionProvision = "0"
	MarketDisruptionProvision_APPLICABLE                       MarketDisruptionProvision = "1"
	MarketDisruptionProvision_AS_SPECIFIED_IN_MASTER_AGREEMENT MarketDisruptionProvision = "2"
	MarketDisruptionProvision_AS_SPECIFIED_IN_CONFIRMATION     MarketDisruptionProvision = "3"
)

// MarketMakerActivity field enumeration values.
type MarketMakerActivity string

const (
	MarketMakerActivity_NO_PARTICIPATION                MarketMakerActivity = "0"
	MarketMakerActivity_BUY_PARTICIPATION               MarketMakerActivity = "1"
	MarketMakerActivity_SELL_PARTICIPATION              MarketMakerActivity = "2"
	MarketMakerActivity_BOTH_BUY_AND_SELL_PARTICIPATION MarketMakerActivity = "3"
)

// MarketSegmentRelationship field enumeration values.
type MarketSegmentRelationship string

const (
	MarketSegmentRelationship_MARKET_SEGMENT_POOL_MEMBER MarketSegmentRelationship = "1"
	MarketSegmentRelationship_RETAIL_SEGMENT             MarketSegmentRelationship = "2"
	MarketSegmentRelationship_WHOLESALE_SEGMENT          MarketSegmentRelationship = "3"
)

// MarketSegmentStatus field enumeration values.
type MarketSegmentStatus string

const (
	MarketSegmentStatus_ACTIVE    MarketSegmentStatus = "1"
	MarketSegmentStatus_INACTIVE  MarketSegmentStatus = "2"
	MarketSegmentStatus_PUBLISHED MarketSegmentStatus = "3"
)

// MarketSegmentSubType field enumeration values.
type MarketSegmentSubType string

const (
	MarketSegmentSubType_INTER_PRODUCT_SPREAD MarketSegmentSubType = "1"
)

// MarketSegmentType field enumeration values.
type MarketSegmentType string

const (
	MarketSegmentType_POOL      MarketSegmentType = "1"
	MarketSegmentType_RETAIL    MarketSegmentType = "2"
	MarketSegmentType_WHOLESALE MarketSegmentType = "3"
)

// MarketUpdateAction field enumeration values.
type MarketUpdateAction string

const (
	MarketUpdateAction_ADD    MarketUpdateAction = "A"
	MarketUpdateAction_DELETE MarketUpdateAction = "D"
	MarketUpdateAction_MODIFY MarketUpdateAction = "M"
)

// MassActionReason field enumeration values.
type MassActionReason string

const (
	MassActionReason_NO_SPECIAL_REASON          MassActionReason = "0"
	MassActionReason_TRADING_RISK_CONTROL       MassActionReason = "1"
	MassActionReason_COMPLEX_INSTRUMENT_DELETED MassActionReason = "10"
	MassActionReason_CIRCUIT_BREAKER_ACTIVATED  MassActionReason = "11"
	MassActionReason_CLEARING_RISK_CONTROL      MassActionReason = "2"
	MassActionReason_MARKET_MAKER_PROTECTION    MassActionReason = "3"
	MassActionReason_STOP_TRADING               MassActionReason = "4"
	MassActionReason_EMERGENCY_ACTION           MassActionReason = "5"
	MassActionReason_SESSION_LOSS_OR_LOGOUT     MassActionReason = "6"
	MassActionReason_DUPLICATE_LOGIN            MassActionReason = "7"
	MassActionReason_PRODUCT_NOT_TRADED         MassActionReason = "8"
	MassActionReason_INSTRUMENT_NOT_TRADED      MassActionReason = "9"
	MassActionReason_OTHER                      MassActionReason = "99"
)

// MassActionRejectReason field enumeration values.
type MassActionRejectReason string

const (
	MassActionRejectReason_MASS_ACTION_NOT_SUPPORTED                        MassActionRejectReason = "0"
	MassActionRejectReason_INVALID_OR_UNKNOWN_SECURITY                      MassActionRejectReason = "1"
	MassActionRejectReason_INVALID_OR_UNKNOWN_SECURITY_ISSUER               MassActionRejectReason = "10"
	MassActionRejectReason_INVALID_OR_UNKNOWN_ISSUER_OF_UNDERLYING_SECURITY MassActionRejectReason = "11"
	MassActionRejectReason_INVALID_OR_UNKNOWN_UNDERLYING_SECURITY           MassActionRejectReason = "2"
	MassActionRejectReason_INVALID_OR_UNKNOWN_PRODUCT                       MassActionRejectReason = "3"
	MassActionRejectReason_INVALID_OR_UNKNOWN_CFICODE                       MassActionRejectReason = "4"
	MassActionRejectReason_INVALID_OR_UNKNOWN_SECURITYTYPE                  MassActionRejectReason = "5"
	MassActionRejectReason_INVALID_OR_UNKNOWN_TRADING_SESSION               MassActionRejectReason = "6"
	MassActionRejectReason_INVALID_OR_UNKNOWN_MARKET                        MassActionRejectReason = "7"
	MassActionRejectReason_INVALID_OR_UNKNOWN_MARKET_SEGMENT                MassActionRejectReason = "8"
	MassActionRejectReason_INVALID_OR_UNKNOWN_SECURITY_GROUP                MassActionRejectReason = "9"
	MassActionRejectReason_OTHER                                            MassActionRejectReason = "99"
)

// MassActionResponse field enumeration values.
type MassActionResponse string

const (
	MassActionResponse_REJECTED  MassActionResponse = "0"
	MassActionResponse_ACCEPTED  MassActionResponse = "1"
	MassActionResponse_COMPLETED MassActionResponse = "2"
)

// MassActionScope field enumeration values.
type MassActionScope string

const (
	MassActionScope_ALL_ORDERS_FOR_A_SECURITY                MassActionScope = "1"
	MassActionScope_ALL_ORDERS_FOR_A_SECURITY_GROUP          MassActionScope = "10"
	MassActionScope_CANCEL_FOR_SECURITY_ISSUER               MassActionScope = "11"
	MassActionScope_CANCEL_FOR_ISSUER_OF_UNDERLYING_SECURITY MassActionScope = "12"
	MassActionScope_ALL_ORDERS_FOR_AN_UNDERLYING_SECURITY    MassActionScope = "2"
	MassActionScope_ALL_ORDERS_FOR_A_PRODUCT                 MassActionScope = "3"
	MassActionScope_ALL_ORDERS_FOR_A_CFICODE                 MassActionScope = "4"
	MassActionScope_ALL_ORDERS_FOR_A_SECURITYTYPE            MassActionScope = "5"
	MassActionScope_ALL_ORDERS_FOR_A_TRADING_SESSION         MassActionScope = "6"
	MassActionScope_ALL_ORDERS                               MassActionScope = "7"
	MassActionScope_ALL_ORDERS_FOR_A_MARKET                  MassActionScope = "8"
	MassActionScope_ALL_ORDERS_FOR_A_MARKET_SEGMENT          MassActionScope = "9"
)

// MassActionType field enumeration values.
type MassActionType string

const (
	MassActionType_SUSPEND_ORDERS                 MassActionType = "1"
	MassActionType_RELEASE_ORDERS_FROM_SUSPENSION MassActionType = "2"
	MassActionType_CANCEL_ORDERS                  MassActionType = "3"
)

// MassCancelRejectReason field enumeration values.
type MassCancelRejectReason string

const (
	MassCancelRejectReason_MASS_CANCEL_NOT_SUPPORTED                        MassCancelRejectReason = "0"
	MassCancelRejectReason_INVALID_OR_UNKNOWN_SECURITY                      MassCancelRejectReason = "1"
	MassCancelRejectReason_INVALID_OR_UNKNOWN_SECURITY_ISSUER               MassCancelRejectReason = "10"
	MassCancelRejectReason_INVALID_OR_UNKNOWN_ISSUER_OF_UNDERLYING_SECURITY MassCancelRejectReason = "11"
	MassCancelRejectReason_INVALID_OR_UNKOWN_UNDERLYING_SECURITY            MassCancelRejectReason = "2"
	MassCancelRejectReason_INVALID_OR_UNKNOWN_PRODUCT                       MassCancelRejectReason = "3"
	MassCancelRejectReason_INVALID_OR_UNKNOWN_CFICODE                       MassCancelRejectReason = "4"
	MassCancelRejectReason_INVALID_OR_UNKNOWN_SECURITYTYPE                  MassCancelRejectReason = "5"
	MassCancelRejectReason_INVALID_OR_UNKNOWN_TRADING_SESSION               MassCancelRejectReason = "6"
	MassCancelRejectReason_INVALID_OR_UNKNOWN_MARKET                        MassCancelRejectReason = "7"
	MassCancelRejectReason_INVALID_OR_UNKOWN_MARKET_SEGMENT                 MassCancelRejectReason = "8"
	MassCancelRejectReason_INVALID_OR_UNKNOWN_SECURITY_GROUP                MassCancelRejectReason = "9"
	MassCancelRejectReason_OTHER                                            MassCancelRejectReason = "99"
)

// MassCancelRequestType field enumeration values.
type MassCancelRequestType string

const (
	MassCancelRequestType_CANCEL_ORDERS_FOR_A_SECURITY             MassCancelRequestType = "1"
	MassCancelRequestType_CANCEL_ORDERS_FOR_AN_UNDERLYING_SECURITY MassCancelRequestType = "2"
	MassCancelRequestType_CANCEL_ORDERS_FOR_A_PRODUCT              MassCancelRequestType = "3"
	MassCancelRequestType_CANCEL_ORDERS_FOR_A_CFICODE              MassCancelRequestType = "4"
	MassCancelRequestType_CANCEL_ORDERS_FOR_A_SECURITYTYPE         MassCancelRequestType = "5"
	MassCancelRequestType_CANCEL_ORDERS_FOR_A_TRADING_SESSION      MassCancelRequestType = "6"
	MassCancelRequestType_CANCEL_ALL_ORDERS                        MassCancelRequestType = "7"
	MassCancelRequestType_CANCEL_ORDERS_FOR_A_MARKET               MassCancelRequestType = "8"
	MassCancelRequestType_CANCEL_ORDERS_FOR_A_MARKET_SEGMENT       MassCancelRequestType = "9"
	MassCancelRequestType_CANCEL_ORDERS_FOR_A_SECURITY_GROUP       MassCancelRequestType = "A"
	MassCancelRequestType_CANCEL_FOR_SECURITY_ISSUER               MassCancelRequestType = "B"
	MassCancelRequestType_CANCEL_FOR_ISSUER_OF_UNDERLYING_SECURITY MassCancelRequestType = "C"
)

// MassCancelResponse field enumeration values.
type MassCancelResponse string

const (
	MassCancelResponse_CANCEL_REQUEST_REJECTED                         MassCancelResponse = "0"
	MassCancelResponse_CANCEL_ORDERS_FOR_A_SECURITY                    MassCancelResponse = "1"
	MassCancelResponse_CANCEL_ORDERS_FOR_AN_UNDERLYING_SECURITY        MassCancelResponse = "2"
	MassCancelResponse_CANCEL_ORDERS_FOR_A_PRODUCT                     MassCancelResponse = "3"
	MassCancelResponse_CANCEL_ORDERS_FOR_A_CFICODE                     MassCancelResponse = "4"
	MassCancelResponse_CANCEL_ORDERS_FOR_A_SECURITYTYPE                MassCancelResponse = "5"
	MassCancelResponse_CANCEL_ORDERS_FOR_A_TRADING_SESSION             MassCancelResponse = "6"
	MassCancelResponse_CANCEL_ALL_ORDERS                               MassCancelResponse = "7"
	MassCancelResponse_CANCEL_ORDERS_FOR_A_MARKET                      MassCancelResponse = "8"
	MassCancelResponse_CANCEL_ORDERS_FOR_A_MARKET_SEGMENT              MassCancelResponse = "9"
	MassCancelResponse_CANCEL_ORDERS_FOR_A_SECURITY_GROUP              MassCancelResponse = "A"
	MassCancelResponse_CANCEL_ORDERS_FOR_A_SECURITIES_ISSUER           MassCancelResponse = "B"
	MassCancelResponse_CANCEL_ORDERS_FOR_ISSUER_OF_UNDERLYING_SECURITY MassCancelResponse = "C"
)

// MassOrderRequestResult field enumeration values.
type MassOrderRequestResult string

const (
	MassOrderRequestResult_SUCCESSFUL                   MassOrderRequestResult = "0"
	MassOrderRequestResult_RESPONSE_LEVEL_NOT_SUPPORTED MassOrderRequestResult = "1"
	MassOrderRequestResult_INVALID_MARKET               MassOrderRequestResult = "2"
	MassOrderRequestResult_INVALID_MARKET_SEGMENT       MassOrderRequestResult = "3"
	MassOrderRequestResult_OTHER                        MassOrderRequestResult = "99"
)

// MassOrderRequestStatus field enumeration values.
type MassOrderRequestStatus string

const (
	MassOrderRequestStatus_ACCEPTED                        MassOrderRequestStatus = "1"
	MassOrderRequestStatus_ACCEPTED_WITH_ADDITIONAL_EVENTS MassOrderRequestStatus = "2"
	MassOrderRequestStatus_REJECTED                        MassOrderRequestStatus = "3"
)

// MassStatusReqType field enumeration values.
type MassStatusReqType string

const (
	MassStatusReqType_STATUS_FOR_ORDERS_FOR_A_SECURITY             MassStatusReqType = "1"
	MassStatusReqType_STATUS_FOR_ISSUER_OF_UNDERLYING_SECURITY     MassStatusReqType = "10"
	MassStatusReqType_STATUS_FOR_ORDERS_FOR_AN_UNDERLYING_SECURITY MassStatusReqType = "2"
	MassStatusReqType_STATUS_FOR_ORDERS_FOR_A_PRODUCT              MassStatusReqType = "3"
	MassStatusReqType_STATUS_FOR_ORDERS_FOR_A_CFICODE              MassStatusReqType = "4"
	MassStatusReqType_STATUS_FOR_ORDERS_FOR_A_SECURITYTYPE         MassStatusReqType = "5"
	MassStatusReqType_STATUS_FOR_ORDERS_FOR_A_TRADING_SESSION      MassStatusReqType = "6"
	MassStatusReqType_STATUS_FOR_ALL_ORDERS                        MassStatusReqType = "7"
	MassStatusReqType_STATUS_FOR_ORDERS_FOR_A_PARTYID              MassStatusReqType = "8"
	MassStatusReqType_STATUS_FOR_SECURITY_ISSUER                   MassStatusReqType = "9"
)

// MatchExceptionElementType field enumeration values.
type MatchExceptionElementType string

const (
	MatchExceptionElementType_ACCRUED_INTEREST               MatchExceptionElementType = "1"
	MatchExceptionElementType_INVESTMENT_MANAGER_ID          MatchExceptionElementType = "10"
	MatchExceptionElementType_NET_AMOUNT                     MatchExceptionElementType = "11"
	MatchExceptionElementType_PLACE_OF_SETTLEMENT            MatchExceptionElementType = "12"
	MatchExceptionElementType_COMMISSIONS                    MatchExceptionElementType = "13"
	MatchExceptionElementType_SECURITY_IDENTIFIER            MatchExceptionElementType = "14"
	MatchExceptionElementType_QUANTITY_ALLOCATED             MatchExceptionElementType = "15"
	MatchExceptionElementType_PRINCIPAL                      MatchExceptionElementType = "16"
	MatchExceptionElementType_FEES                           MatchExceptionElementType = "17"
	MatchExceptionElementType_TAX                            MatchExceptionElementType = "18"
	MatchExceptionElementType_DEAL_PRICE                     MatchExceptionElementType = "2"
	MatchExceptionElementType_TRADE_DATE                     MatchExceptionElementType = "3"
	MatchExceptionElementType_SETTLEMENT_DATE                MatchExceptionElementType = "4"
	MatchExceptionElementType_SIDE_INDICATOR                 MatchExceptionElementType = "5"
	MatchExceptionElementType_TRADED_CURRENCY                MatchExceptionElementType = "6"
	MatchExceptionElementType_ACCOUNT_ID                     MatchExceptionElementType = "7"
	MatchExceptionElementType_EXECUTING_BROKER_ID            MatchExceptionElementType = "8"
	MatchExceptionElementType_SETTLEMENT_CURRENCY_AND_AMOUNT MatchExceptionElementType = "9"
)

// MatchExceptionToleranceValueType field enumeration values.
type MatchExceptionToleranceValueType string

const (
	MatchExceptionToleranceValueType_FIXED_AMOUNT MatchExceptionToleranceValueType = "1"
	MatchExceptionToleranceValueType_PERCENTAGE   MatchExceptionToleranceValueType = "2"
)

// MatchExceptionType field enumeration values.
type MatchExceptionType string

const (
	MatchExceptionType_NO_MATCHING_CONFIRMATION             MatchExceptionType = "0"
	MatchExceptionType_NO_MATCHING_ALLOCATION               MatchExceptionType = "1"
	MatchExceptionType_ALLOCATION_DATA_ELEMENT_MISSING      MatchExceptionType = "2"
	MatchExceptionType_CONFIRMATION_DATA_ELEMENT_MISSING    MatchExceptionType = "3"
	MatchExceptionType_DATA_DIFFERENCE_NOT_WITHIN_TOLERANCE MatchExceptionType = "4"
	MatchExceptionType_MATCH_WITHIN_TOLERANCE               MatchExceptionType = "5"
	MatchExceptionType_OTHER                                MatchExceptionType = "99"
)

// MatchInst field enumeration values.
type MatchInst string

const (
	MatchInst_MATCH        MatchInst = "1"
	MatchInst_DO_NOT_MATCH MatchInst = "2"
)

// MatchStatus field enumeration values.
type MatchStatus string

const (
	MatchStatus_COMPARED_MATCHED_OR_AFFIRMED       MatchStatus = "0"
	MatchStatus_UNCOMPARED_UNMATCHED_OR_UNAFFIRMED MatchStatus = "1"
	MatchStatus_ADVISORY_OR_ALERT                  MatchStatus = "2"
	MatchStatus_MISMATCHED                         MatchStatus = "3"
)

// MatchType field enumeration values.
type MatchType string

const (
	MatchType_ONE_PARTY_TRADE_REPORT                                                                                                           MatchType = "1"
	MatchType_AUTO_MATCH_WITH_LAST_LOOK                                                                                                        MatchType = "10"
	MatchType_CROSS_AUCTION_WITH_LAST_LOOK                                                                                                     MatchType = "11"
	MatchType_TWO_PARTY_TRADE_REPORT                                                                                                           MatchType = "2"
	MatchType_CONFIRMED_TRADE_REPORT                                                                                                           MatchType = "3"
	MatchType_AUTO_MATCH                                                                                                                       MatchType = "4"
	MatchType_CROSS_AUCTION                                                                                                                    MatchType = "5"
	MatchType_COUNTER_ORDER_SELECTION                                                                                                          MatchType = "6"
	MatchType_ONE_PARTY_PRIVATELY_NEGOTIATED_TRADE_REPORT                                                                                      MatchType = "60"
	MatchType_TWO_PARTY_PRIVATELY_NEGOTIATED_TRADE_REPORT                                                                                      MatchType = "61"
	MatchType_CONTINUOUS_AUTO_MATCH                                                                                                            MatchType = "62"
	MatchType_CROSS_AUCTION_63                                                                                                                 MatchType = "63"
	MatchType_COUNTER_ORDER_SELECTION_64                                                                                                       MatchType = "64"
	MatchType_CALL_AUCTION_65                                                                                                                  MatchType = "65"
	MatchType_CALL_AUCTION                                                                                                                     MatchType = "7"
	MatchType_ISSUING_BUY_BACK_AUCTION                                                                                                         MatchType = "8"
	MatchType_SYSTEMATIC_INTERNALISER                                                                                                          MatchType = "9"
	MatchType_EXACT_MATCH_ON_TRADE_DATE_STOCK_SYMBOL_QUANTITY_PRICE_TRADE_TYPE_AND_SPECIAL_TRADE_INDICATOR_PLUS_FOUR_BADGES_AND_EXECUTION_TIME MatchType = "A1"
	MatchType_EXACT_MATCH_ON_TRADE_DATE_STOCK_SYMBOL_QUANTITY_PRICE_TRADE_TYPE_AND_SPECIAL_TRADE_INDICATOR_PLUS_FOUR_BADGES                    MatchType = "A2"
	MatchType_EXACT_MATCH_ON_TRADE_DATE_STOCK_SYMBOL_QUANTITY_PRICE_TRADE_TYPE_AND_SPECIAL_TRADE_INDICATOR_PLUS_TWO_BADGES_AND_EXECUTION_TIME  MatchType = "A3"
	MatchType_EXACT_MATCH_ON_TRADE_DATE_STOCK_SYMBOL_QUANTITY_PRICE_TRADE_TYPE_AND_SPECIAL_TRADE_INDICATOR_PLUS_TWO_BADGES                     MatchType = "A4"
	MatchType_EXACT_MATCH_ON_TRADE_DATE_STOCK_SYMBOL_QUANTITY_PRICE_TRADETYPE_AND_SPECIAL_TRADE_INDICATOR_PLUS_EXECUTION_TIME                  MatchType = "A5"
	MatchType_NASDAQACTM1MATCH                                                                                                                 MatchType = "ACTM1"
	MatchType_NASDAQACTM2MATCH                                                                                                                 MatchType = "ACTM2"
	MatchType_NASDAQACTACCEPTEDTRADE                                                                                                           MatchType = "ACTM3"
	MatchType_NASDAQACTDEFAULTTRADE                                                                                                            MatchType = "ACTM4"
	MatchType_NASDAQACTDEFAULTAFTERM2                                                                                                          MatchType = "ACTM5"
	MatchType_NASDAQACTM6MATCH                                                                                                                 MatchType = "ACTM6"
	MatchType_NASDAQNONACT                                                                                                                     MatchType = "ACTMT"
	MatchType_COMPARED_RECORDS_RESULTING_FROM_STAMPED_ADVISORIES_OR_SPECIALIST_ACCEPTS_PAIR_OFFS                                               MatchType = "AQ"
	MatchType_EXACT_MATCH_ON_TRADE_DATE_STOCK_SYMBOL_QUANTITY_PRICE_TRADE_TYPE_AND_SPECIAL_TRADE_INDICATOR_MINUS_BADGES_AND_TIMES_ACT_M1_MATCH MatchType = "M1"
	MatchType_SUMMARIZED_MATCH_MINUS_BADGES_AND_TIMES_ACT_M2_MATCH                                                                             MatchType = "M2"
	MatchType_ACT_ACCEPTED_TRADE                                                                                                               MatchType = "M3"
	MatchType_ACT_DEFAULT_TRADE                                                                                                                MatchType = "M4"
	MatchType_ACT_DEFAULT_AFTER_M2                                                                                                             MatchType = "M5"
	MatchType_ACT_M6_MATCH                                                                                                                     MatchType = "M6"
	MatchType_OCS_LOCKED_IN_NON_ACT                                                                                                            MatchType = "MT"
	MatchType_SUMMARIZED_MATCH_USING_A1_EXACT_MATCH_CRITERIA_EXCEPT_QUANTITY_IS_SUMMARIED                                                      MatchType = "S1"
	MatchType_SUMMARIZED_MATCH_USING_A2_EXACT_MATCH_CRITERIA_EXCEPT_QUANTITY_IS_SUMMARIZED                                                     MatchType = "S2"
	MatchType_SUMMARIZED_MATCH_USING_A3_EXACT_MATCH_CRITERIA_EXCEPT_QUANTITY_IS_SUMMARIZED                                                     MatchType = "S3"
	MatchType_SUMMARIZED_MATCH_USING_A4_EXACT_MATCH_CRITERIA_EXCEPT_QUANTITY_IS_SUMMARIZED                                                     MatchType = "S4"
	MatchType_SUMMARIZED_MATCH_USING_A5_EXACT_MATCH_CRITERIA_EXCEPT_QUANTITY_IS_SUMMARIZED                                                     MatchType = "S5"
)

// MatchingDataPointIndicator field enumeration values.
type MatchingDataPointIndicator string

const (
	MatchingDataPointIndicator_MANDATORY MatchingDataPointIndicator = "1"
	MatchingDataPointIndicator_OPTIONAL  MatchingDataPointIndicator = "2"
)

// MaturityMonthYearFormat field enumeration values.
type MaturityMonthYearFormat string

const (
	MaturityMonthYearFormat_YEARMONTH_ONLY MaturityMonthYearFormat = "0"
	MaturityMonthYearFormat_YEARMONTHDAY   MaturityMonthYearFormat = "1"
	MaturityMonthYearFormat_YEARMONTHWEEK  MaturityMonthYearFormat = "2"
)

// MaturityMonthYearIncrementUnits field enumeration values.
type MaturityMonthYearIncrementUnits string

const (
	MaturityMonthYearIncrementUnits_MONTHS MaturityMonthYearIncrementUnits = "0"
	MaturityMonthYearIncrementUnits_DAYS   MaturityMonthYearIncrementUnits = "1"
	MaturityMonthYearIncrementUnits_WEEKS  MaturityMonthYearIncrementUnits = "2"
	MaturityMonthYearIncrementUnits_YEARS  MaturityMonthYearIncrementUnits = "3"
)

// MessageEncoding field enumeration values.
type MessageEncoding string

const (
	MessageEncoding_EUC_JP      MessageEncoding = "EUC-JP"
	MessageEncoding_ISO_2022_JP MessageEncoding = "ISO-2022-JP"
	MessageEncoding_SHIFT_JIS   MessageEncoding = "SHIFT_JIS"
	MessageEncoding_UTF_8       MessageEncoding = "UTF-8"
)

// MinQtyMethod field enumeration values.
type MinQtyMethod string

const (
	MinQtyMethod_ONCE     MinQtyMethod = "1"
	MinQtyMethod_MULTIPLE MinQtyMethod = "2"
)

// MiscFeeBasis field enumeration values.
type MiscFeeBasis string

const (
	MiscFeeBasis_ABSOLUTE   MiscFeeBasis = "0"
	MiscFeeBasis_PER_UNIT   MiscFeeBasis = "1"
	MiscFeeBasis_PERCENTAGE MiscFeeBasis = "2"
)

// MiscFeeQualifier field enumeration values.
type MiscFeeQualifier string

const (
	MiscFeeQualifier_CONTRIBUTES         MiscFeeQualifier = "0"
	MiscFeeQualifier_DOES_NOT_CONTRIBUTE MiscFeeQualifier = "1"
)

// MiscFeeType field enumeration values.
type MiscFeeType string

const (
	MiscFeeType_REGULATORY                       MiscFeeType = "1"
	MiscFeeType_PER_TRANSACTION                  MiscFeeType = "10"
	MiscFeeType_CONVERSION                       MiscFeeType = "11"
	MiscFeeType_AGENT                            MiscFeeType = "12"
	MiscFeeType_TRANSFER_FEE                     MiscFeeType = "13"
	MiscFeeType_SECURITY_LENDING                 MiscFeeType = "14"
	MiscFeeType_TRADE_REPORTING                  MiscFeeType = "15"
	MiscFeeType_TAX_ON_PRINCIPAL_AMOUNT          MiscFeeType = "16"
	MiscFeeType_TAX_ON_ACCRUED_INTEREST_AMOUNT   MiscFeeType = "17"
	MiscFeeType_NEW_ISSUANCE_FEE                 MiscFeeType = "18"
	MiscFeeType_SERVICE_FEE                      MiscFeeType = "19"
	MiscFeeType_TAX                              MiscFeeType = "2"
	MiscFeeType_ODD_LOT_FEE                      MiscFeeType = "20"
	MiscFeeType_AUCTION_FEE                      MiscFeeType = "21"
	MiscFeeType_VALUE_ADDED_TAX                  MiscFeeType = "22"
	MiscFeeType_SALES_TAX                        MiscFeeType = "23"
	MiscFeeType_EXECUTION_VENUE_FEE              MiscFeeType = "24"
	MiscFeeType_ORDER_OR_QUOTE_ENTRY_FEE         MiscFeeType = "25"
	MiscFeeType_ORDER_OR_QUOTE_MODIFICATION_FEE  MiscFeeType = "26"
	MiscFeeType_ORDERS_OR_QUOTE_CANCELLATION_FEE MiscFeeType = "27"
	MiscFeeType_MARKET_DATA_ACCESS_FEE           MiscFeeType = "28"
	MiscFeeType_MARKET_DATA_TERMINAL_FEE         MiscFeeType = "29"
	MiscFeeType_LOCAL_COMMISSION                 MiscFeeType = "3"
	MiscFeeType_MARKET_DATA_VOLUME_FEE           MiscFeeType = "30"
	MiscFeeType_CLEARING_FEE                     MiscFeeType = "31"
	MiscFeeType_SETTLEMENT_FEE                   MiscFeeType = "32"
	MiscFeeType_REBATES                          MiscFeeType = "33"
	MiscFeeType_DISCOUNTS                        MiscFeeType = "34"
	MiscFeeType_PAYMENTS                         MiscFeeType = "35"
	MiscFeeType_NON_MONETARY_PAYMENTS            MiscFeeType = "36"
	MiscFeeType_EXCHANGE_FEES                    MiscFeeType = "4"
	MiscFeeType_STAMP                            MiscFeeType = "5"
	MiscFeeType_LEVY                             MiscFeeType = "6"
	MiscFeeType_OTHER                            MiscFeeType = "7"
	MiscFeeType_MARKUP                           MiscFeeType = "8"
	MiscFeeType_CONSUMPTION_TAX                  MiscFeeType = "9"
)

// ModelType field enumeration values.
type ModelType string

const (
	ModelType_UTILITY_PROVIDED_STANDARD_MODEL ModelType = "0"
	ModelType_PROPRIETARY                     ModelType = "1"
)

// MoneyLaunderingStatus field enumeration values.
type MoneyLaunderingStatus string

const (
	MoneyLaunderingStatus_EXEMPT_1    MoneyLaunderingStatus = "1"
	MoneyLaunderingStatus_EXEMPT_2    MoneyLaunderingStatus = "2"
	MoneyLaunderingStatus_EXEMPT_3    MoneyLaunderingStatus = "3"
	MoneyLaunderingStatus_NOT_CHECKED MoneyLaunderingStatus = "N"
	MoneyLaunderingStatus_PASSED      MoneyLaunderingStatus = "Y"
)

// MsgDirection field enumeration values.
type MsgDirection string

const (
	MsgDirection_RECEIVE MsgDirection = "R"
	MsgDirection_SEND    MsgDirection = "S"
)

// MsgType field enumeration values.
type MsgType string

const (
	MsgType_HEARTBEAT                             MsgType = "0"
	MsgType_TEST_REQUEST                          MsgType = "1"
	MsgType_RESEND_REQUEST                        MsgType = "2"
	MsgType_REJECT                                MsgType = "3"
	MsgType_SEQUENCE_RESET                        MsgType = "4"
	MsgType_LOGOUT                                MsgType = "5"
	MsgType_INDICATION_OF_INTEREST                MsgType = "6"
	MsgType_ADVERTISEMENT                         MsgType = "7"
	MsgType_EXECUTION_REPORT                      MsgType = "8"
	MsgType_ORDER_CANCEL_REJECT                   MsgType = "9"
	MsgType_LOGON                                 MsgType = "A"
	MsgType_DERIVATIVE_SECURITY_LIST              MsgType = "AA"
	MsgType_NEW_ORDER_MULTILEG                    MsgType = "AB"
	MsgType_MULTILEG_ORDER_CANCEL_REPLACE         MsgType = "AC"
	MsgType_TRADE_CAPTURE_REPORT_REQUEST          MsgType = "AD"
	MsgType_TRADE_CAPTURE_REPORT                  MsgType = "AE"
	MsgType_ORDER_MASS_STATUS_REQUEST             MsgType = "AF"
	MsgType_QUOTE_REQUEST_REJECT                  MsgType = "AG"
	MsgType_RFQ_REQUEST                           MsgType = "AH"
	MsgType_QUOTE_STATUS_REPORT                   MsgType = "AI"
	MsgType_QUOTE_RESPONSE                        MsgType = "AJ"
	MsgType_CONFIRMATION                          MsgType = "AK"
	MsgType_POSITION_MAINTENANCE_REQUEST          MsgType = "AL"
	MsgType_POSITION_MAINTENANCE_REPORT           MsgType = "AM"
	MsgType_REQUEST_FOR_POSITIONS                 MsgType = "AN"
	MsgType_REQUEST_FOR_POSITIONS_ACK             MsgType = "AO"
	MsgType_POSITION_REPORT                       MsgType = "AP"
	MsgType_TRADE_CAPTURE_REPORT_REQUEST_ACK      MsgType = "AQ"
	MsgType_TRADE_CAPTURE_REPORT_ACK              MsgType = "AR"
	MsgType_ALLOCATION_REPORT                     MsgType = "AS"
	MsgType_ALLOCATION_REPORT_ACK                 MsgType = "AT"
	MsgType_CONFIRMATION_ACK                      MsgType = "AU"
	MsgType_SETTLEMENT_INSTRUCTION_REQUEST        MsgType = "AV"
	MsgType_ASSIGNMENT_REPORT                     MsgType = "AW"
	MsgType_COLLATERAL_REQUEST                    MsgType = "AX"
	MsgType_COLLATERAL_ASSIGNMENT                 MsgType = "AY"
	MsgType_COLLATERAL_RESPONSE                   MsgType = "AZ"
	MsgType_NEWS                                  MsgType = "B"
	MsgType_COLLATERAL_REPORT                     MsgType = "BA"
	MsgType_COLLATERAL_INQUIRY                    MsgType = "BB"
	MsgType_NETWORK_STATUS_REQUEST                MsgType = "BC"
	MsgType_NETWORK_STATUS_RESPONSE               MsgType = "BD"
	MsgType_USER_REQUEST                          MsgType = "BE"
	MsgType_USER_RESPONSE                         MsgType = "BF"
	MsgType_COLLATERAL_INQUIRY_ACK                MsgType = "BG"
	MsgType_CONFIRMATION_REQUEST                  MsgType = "BH"
	MsgType_TRADING_SESSION_LIST_REQUEST          MsgType = "BI"
	MsgType_TRADING_SESSION_LIST                  MsgType = "BJ"
	MsgType_SECURITY_LIST_UPDATE_REPORT           MsgType = "BK"
	MsgType_ADJUSTED_POSITION_REPORT              MsgType = "BL"
	MsgType_ALLOCATION_INSTRUCTION_ALERT          MsgType = "BM"
	MsgType_EXECUTION_ACKNOWLEDGEMENT             MsgType = "BN"
	MsgType_CONTRARY_INTENTION_REPORT             MsgType = "BO"
	MsgType_SECURITY_DEFINITION_UPDATE_REPORT     MsgType = "BP"
	MsgType_SETTLEMENTOBLIGATIONREPORT            MsgType = "BQ"
	MsgType_DERIVATIVESECURITYLISTUPDATEREPORT    MsgType = "BR"
	MsgType_TRADINGSESSIONLISTUPDATEREPORT        MsgType = "BS"
	MsgType_MARKETDEFINITIONREQUEST               MsgType = "BT"
	MsgType_MARKETDEFINITION                      MsgType = "BU"
	MsgType_MARKETDEFINITIONUPDATEREPORT          MsgType = "BV"
	MsgType_APPLICATIONMESSAGEREQUEST             MsgType = "BW"
	MsgType_APPLICATIONMESSAGEREQUESTACK          MsgType = "BX"
	MsgType_APPLICATIONMESSAGEREPORT              MsgType = "BY"
	MsgType_ORDERMASSACTIONREPORT                 MsgType = "BZ"
	MsgType_EMAIL                                 MsgType = "C"
	MsgType_ORDERMASSACTIONREQUEST                MsgType = "CA"
	MsgType_USERNOTIFICATION                      MsgType = "CB"
	MsgType_STREAMASSIGNMENTREQUEST               MsgType = "CC"
	MsgType_STREAMASSIGNMENTREPORT                MsgType = "CD"
	MsgType_STREAMASSIGNMENTREPORTACK             MsgType = "CE"
	MsgType_PARTYDETAILSLISTREQUEST               MsgType = "CF"
	MsgType_PARTYDETAILSLISTREPORT                MsgType = "CG"
	MsgType_MARGINREQUIREMENTINQUIRY              MsgType = "CH"
	MsgType_MARGINREQUIREMENTINQUIRYACK           MsgType = "CI"
	MsgType_MARGINREQUIREMENTREPORT               MsgType = "CJ"
	MsgType_PARTYDETAILSLISTUPDATEREPORT          MsgType = "CK"
	MsgType_PARTYRISKLIMITSREQUEST                MsgType = "CL"
	MsgType_PARTYRISKLIMITSREPORT                 MsgType = "CM"
	MsgType_SECURITYMASSSTATUSREQUEST             MsgType = "CN"
	MsgType_SECURITYMASSSTATUS                    MsgType = "CO"
	MsgType_ACCOUNTSUMMARYREPORT                  MsgType = "CQ"
	MsgType_PARTYRISKLIMITSUPDATEREPORT           MsgType = "CR"
	MsgType_PARTYRISKLIMITSDEFINITIONREQUEST      MsgType = "CS"
	MsgType_PARTYRISKLIMITSDEFINITIONREQUESTACK   MsgType = "CT"
	MsgType_PARTYENTITLEMENTSREQUEST              MsgType = "CU"
	MsgType_PARTYENTITLEMENTSREPORT               MsgType = "CV"
	MsgType_QUOTEACK                              MsgType = "CW"
	MsgType_PARTYDETAILSDEFINITIONREQUEST         MsgType = "CX"
	MsgType_PARTYDETAILSDEFINITIONREQUESTACK      MsgType = "CY"
	MsgType_PARTYENTITLEMENTSUPDATEREPORT         MsgType = "CZ"
	MsgType_ORDER_SINGLE                          MsgType = "D"
	MsgType_PARTYENTITLEMENTSDEFINITIONREQUEST    MsgType = "DA"
	MsgType_PARTYENTITLEMENTSDEFINITIONREQUESTACK MsgType = "DB"
	MsgType_TRADEMATCHREPORT                      MsgType = "DC"
	MsgType_TRADEMATCHREPORTACK                   MsgType = "DD"
	MsgType_PARTYRISKLIMITSREPORTACK              MsgType = "DE"
	MsgType_PARTYRISKLIMITCHECKREQUEST            MsgType = "DF"
	MsgType_PARTYRISKLIMITCHECKREQUESTACK         MsgType = "DG"
	MsgType_PARTYACTIONREQUEST                    MsgType = "DH"
	MsgType_PARTYACTIONREPORT                     MsgType = "DI"
	MsgType_MASSORDER                             MsgType = "DJ"
	MsgType_MASSORDERACK                          MsgType = "DK"
	MsgType_POSITIONTRANSFERINSTRUCTION           MsgType = "DL"
	MsgType_POSITIONTRANSFERINSTRUCTIONACK        MsgType = "DM"
	MsgType_POSITIONTRANSFERREPORT                MsgType = "DN"
	MsgType_MARKETDATASTATISTICSREQUEST           MsgType = "DO"
	MsgType_MARKETDATASTATISTICSREPORT            MsgType = "DP"
	MsgType_COLLATERALREPORTACK                   MsgType = "DQ"
	MsgType_MARKETDATAREPORT                      MsgType = "DR"
	MsgType_CROSSREQUEST                          MsgType = "DS"
	MsgType_CROSSREQUESTACK                       MsgType = "DT"
	MsgType_ALLOCATIONINSTRUCTIONALERTREQUEST     MsgType = "DU"
	MsgType_ALLOCATIONINSTRUCTIONALERTREQUESTACK  MsgType = "DV"
	MsgType_TRADEAGGREGATIONREQUEST               MsgType = "DW"
	MsgType_TRADEAGGREGATIONREPORT                MsgType = "DX"
	MsgType_PAYMANAGEMENTREQUEST                  MsgType = "DY"
	MsgType_PAYMANAGEMENTREQUESTACK               MsgType = "DZ"
	MsgType_ORDER_LIST                            MsgType = "E"
	MsgType_PAYMANAGEMENTREPORT                   MsgType = "EA"
	MsgType_PAYMANAGEMENTREPORTACK                MsgType = "EB"
	MsgType_ORDER_CANCEL_REQUEST                  MsgType = "F"
	MsgType_ORDER_CANCEL_REPLACE_REQUEST          MsgType = "G"
	MsgType_ORDER_STATUS_REQUEST                  MsgType = "H"
	MsgType_ALLOCATION_INSTRUCTION                MsgType = "J"
	MsgType_LIST_CANCEL_REQUEST                   MsgType = "K"
	MsgType_LIST_EXECUTE                          MsgType = "L"
	MsgType_LIST_STATUS_REQUEST                   MsgType = "M"
	MsgType_LIST_STATUS                           MsgType = "N"
	MsgType_ALLOCATION_INSTRUCTION_ACK            MsgType = "P"
	MsgType_DONT_KNOW_TRADE                       MsgType = "Q"
	MsgType_QUOTE_REQUEST                         MsgType = "R"
	MsgType_QUOTE                                 MsgType = "S"
	MsgType_SETTLEMENT_INSTRUCTIONS               MsgType = "T"
	MsgType_MARKET_DATA_REQUEST                   MsgType = "V"
	MsgType_MARKET_DATA_SNAPSHOT_FULL_REFRESH     MsgType = "W"
	MsgType_MARKET_DATA_INCREMENTAL_REFRESH       MsgType = "X"
	MsgType_MARKET_DATA_REQUEST_REJECT            MsgType = "Y"
	MsgType_QUOTE_CANCEL                          MsgType = "Z"
	MsgType_QUOTE_STATUS_REQUEST                  MsgType = "a"
	MsgType_MASS_QUOTE_ACKNOWLEDGEMENT            MsgType = "b"
	MsgType_SECURITY_DEFINITION_REQUEST           MsgType = "c"
	MsgType_SECURITY_DEFINITION                   MsgType = "d"
	MsgType_SECURITY_STATUS_REQUEST               MsgType = "e"
	MsgType_SECURITY_STATUS                       MsgType = "f"
	MsgType_TRADING_SESSION_STATUS_REQUEST        MsgType = "g"
	MsgType_TRADING_SESSION_STATUS                MsgType = "h"
	MsgType_MASS_QUOTE                            MsgType = "i"
	MsgType_BUSINESS_MESSAGE_REJECT               MsgType = "j"
	MsgType_BID_REQUEST                           MsgType = "k"
	MsgType_BID_RESPONSE                          MsgType = "l"
	MsgType_LIST_STRIKE_PRICE                     MsgType = "m"
	MsgType_XML_MESSAGE                           MsgType = "n"
	MsgType_REGISTRATION_INSTRUCTIONS             MsgType = "o"
	MsgType_REGISTRATION_INSTRUCTIONS_RESPONSE    MsgType = "p"
	MsgType_ORDER_MASS_CANCEL_REQUEST             MsgType = "q"
	MsgType_ORDER_MASS_CANCEL_REPORT              MsgType = "r"
	MsgType_NEW_ORDER_CROSS                       MsgType = "s"
	MsgType_CROSS_ORDER_CANCEL_REPLACE_REQUEST    MsgType = "t"
	MsgType_CROSS_ORDER_CANCEL_REQUEST            MsgType = "u"
	MsgType_SECURITY_TYPE_REQUEST                 MsgType = "v"
	MsgType_SECURITY_TYPES                        MsgType = "w"
	MsgType_SECURITY_LIST_REQUEST                 MsgType = "x"
	MsgType_SECURITY_LIST                         MsgType = "y"
	MsgType_DERIVATIVE_SECURITY_LIST_REQUEST      MsgType = "z"
)

// MultiLegReportingType field enumeration values.
type MultiLegReportingType string

const (
	MultiLegReportingType_SINGLE_SECURITY                        MultiLegReportingType = "1"
	MultiLegReportingType_INDIVIDUAL_LEG_OF_A_MULTI_LEG_SECURITY MultiLegReportingType = "2"
	MultiLegReportingType_MULTI_LEG_SECURITY                     MultiLegReportingType = "3"
)

// MultiLegRptTypeReq field enumeration values.
type MultiLegRptTypeReq string

const (
	MultiLegRptTypeReq_REPORT_BY_MULITLEG_SECURITY_ONLY                                                      MultiLegRptTypeReq = "0"
	MultiLegRptTypeReq_REPORT_BY_MULTILEG_SECURITY_AND_BY_INSTRUMENT_LEGS_BELONGING_TO_THE_MULTILEG_SECURITY MultiLegRptTypeReq = "1"
	MultiLegRptTypeReq_REPORT_BY_INSTRUMENT_LEGS_BELONGING_TO_THE_MULTILEG_SECURITY_ONLY                     MultiLegRptTypeReq = "2"
)

// MultilegModel field enumeration values.
type MultilegModel string

const (
	MultilegModel_PREDEFINED_MULTILEG_SECURITY          MultilegModel = "0"
	MultilegModel_USER_DEFINED_MULTILEG_SECURITY        MultilegModel = "1"
	MultilegModel_USER_DEFINED_NON_SECURITIZED_MULTILEG MultilegModel = "2"
)

// MultilegPriceMethod field enumeration values.
type MultilegPriceMethod string

const (
	MultilegPriceMethod_NET_PRICE                       MultilegPriceMethod = "0"
	MultilegPriceMethod_REVERSED_NET_PRICE              MultilegPriceMethod = "1"
	MultilegPriceMethod_YIELD_DIFFERENCE                MultilegPriceMethod = "2"
	MultilegPriceMethod_INDIVIDUAL                      MultilegPriceMethod = "3"
	MultilegPriceMethod_CONTRACT_WEIGHTED_AVERAGE_PRICE MultilegPriceMethod = "4"
	MultilegPriceMethod_MULTIPLIED_PRICE                MultilegPriceMethod = "5"
)

// NBBOEntryType field enumeration values.
type NBBOEntryType string

const (
	NBBOEntryType_BID       NBBOEntryType = "0"
	NBBOEntryType_OFFER     NBBOEntryType = "1"
	NBBOEntryType_MID_PRICE NBBOEntryType = "2"
)

// NBBOSource field enumeration values.
type NBBOSource string

const (
	NBBOSource_NOT_APPLICABLE                   NBBOSource = "0"
	NBBOSource_DIRECT                           NBBOSource = "1"
	NBBOSource_SECURITIES_INFORMATION_PROCESSOR NBBOSource = "2"
	NBBOSource_HYBRID                           NBBOSource = "3"
)

// NegotiationMethod field enumeration values.
type NegotiationMethod string

const (
	NegotiationMethod_AUTO_SPOT                                                                                     NegotiationMethod = "0"
	NegotiationMethod_NEGOTIATED_SPOT                                                                               NegotiationMethod = "1"
	NegotiationMethod_THE_SPOT_PRICE_FOR_THE_REFERENCE_OR_BENCHMARK_SECURITY_IS_TO_BE_NEGOTIATED_VIA_PHONE_OR_VOICE NegotiationMethod = "2"
)

// NetGrossInd field enumeration values.
type NetGrossInd string

const (
	NetGrossInd_NET   NetGrossInd = "1"
	NetGrossInd_GROSS NetGrossInd = "2"
)

// NetworkRequestType field enumeration values.
type NetworkRequestType string

const (
	NetworkRequestType_SNAPSHOT                                        NetworkRequestType = "1"
	NetworkRequestType_SUBSCRIBE                                       NetworkRequestType = "2"
	NetworkRequestType_STOP_SUBSCRIBING                                NetworkRequestType = "4"
	NetworkRequestType_LEVEL_OF_DETAIL_THEN_NOCOMPIDS_BECOMES_REQUIRED NetworkRequestType = "8"
)

// NetworkStatusResponseType field enumeration values.
type NetworkStatusResponseType string

const (
	NetworkStatusResponseType_FULL               NetworkStatusResponseType = "1"
	NetworkStatusResponseType_INCREMENTAL_UPDATE NetworkStatusResponseType = "2"
)

// NewsCategory field enumeration values.
type NewsCategory string

const (
	NewsCategory_COMPANY_NEWS          NewsCategory = "0"
	NewsCategory_MARKETPLACE_NEWS      NewsCategory = "1"
	NewsCategory_FINANCIAL_MARKET_NEWS NewsCategory = "2"
	NewsCategory_TECHNICAL_NEWS        NewsCategory = "3"
	NewsCategory_OTHER_NEWS            NewsCategory = "99"
)

// NewsRefType field enumeration values.
type NewsRefType string

const (
	NewsRefType_REPLACEMENT    NewsRefType = "0"
	NewsRefType_OTHER_LANGUAGE NewsRefType = "1"
	NewsRefType_COMPLIMENTARY  NewsRefType = "2"
	NewsRefType_WITHDRAWAL     NewsRefType = "3"
)

// NoSides field enumeration values.
type NoSides string

const (
	NoSides_ONE_SIDE   NoSides = "1"
	NoSides_BOTH_SIDES NoSides = "2"
)

// NonCashDividendTreatment field enumeration values.
type NonCashDividendTreatment string

const (
	NonCashDividendTreatment_POTENTIAL_ADJUSTMENT_EVENT NonCashDividendTreatment = "0"
	NonCashDividendTreatment_CASH_EQUIVALENT            NonCashDividendTreatment = "1"
)

// NonDeliverableFixingDateType field enumeration values.
type NonDeliverableFixingDateType string

const (
	NonDeliverableFixingDateType_UNADJUSTED NonDeliverableFixingDateType = "0"
	NonDeliverableFixingDateType_ADJUSTED   NonDeliverableFixingDateType = "1"
)

// NotAffectedReason field enumeration values.
type NotAffectedReason string

const (
	NotAffectedReason_ORDER_SUSPENDED      NotAffectedReason = "0"
	NotAffectedReason_INSTRUMENT_SUSPENDED NotAffectedReason = "1"
)

// NotifyBrokerOfCredit field enumeration values.
type NotifyBrokerOfCredit string

const (
	NotifyBrokerOfCredit_NO  NotifyBrokerOfCredit = "N"
	NotifyBrokerOfCredit_YES NotifyBrokerOfCredit = "Y"
)

// ObligationType field enumeration values.
type ObligationType string

const (
	ObligationType_BOND             ObligationType = "0"
	ObligationType_CONVERTIBLE_BOND ObligationType = "1"
	ObligationType_MORTGAGE         ObligationType = "2"
	ObligationType_LOAN             ObligationType = "3"
)

// OddLot field enumeration values.
type OddLot string

const (
	OddLot_NO  OddLot = "N"
	OddLot_YES OddLot = "Y"
)

// OffsetInstruction field enumeration values.
type OffsetInstruction string

const (
	OffsetInstruction_OFFSET OffsetInstruction = "0"
	OffsetInstruction_ONSET  OffsetInstruction = "1"
)

// OffshoreIndicator field enumeration values.
type OffshoreIndicator string

const (
	OffshoreIndicator_REGULAR  OffshoreIndicator = "0"
	OffshoreIndicator_OFFSHORE OffshoreIndicator = "1"
	OffshoreIndicator_ONSHORE  OffshoreIndicator = "2"
)

// OpenClose field enumeration values.
type OpenClose string

const (
	OpenClose_CLOSE OpenClose = "C"
	OpenClose_OPEN  OpenClose = "O"
)

// OpenCloseSettlFlag field enumeration values.
type OpenCloseSettlFlag string

const (
	OpenCloseSettlFlag_DAILY_OPEN                       OpenCloseSettlFlag = "0"
	OpenCloseSettlFlag_SESSION_OPEN                     OpenCloseSettlFlag = "1"
	OpenCloseSettlFlag_DELIVERY_SETTLEMENT_ENTRY        OpenCloseSettlFlag = "2"
	OpenCloseSettlFlag_EXPECTED_ENTRY                   OpenCloseSettlFlag = "3"
	OpenCloseSettlFlag_ENTRY_FROM_PREVIOUS_BUSINESS_DAY OpenCloseSettlFlag = "4"
	OpenCloseSettlFlag_THEORETICAL_PRICE_VALUE          OpenCloseSettlFlag = "5"
)

// OpenCloseSettleFlag field enumeration values.
type OpenCloseSettleFlag string

const (
	OpenCloseSettleFlag_DAILY_OPEN                       OpenCloseSettleFlag = "0"
	OpenCloseSettleFlag_SESSION_OPEN                     OpenCloseSettleFlag = "1"
	OpenCloseSettleFlag_DELIVERY_SETTLEMENT_PRICE        OpenCloseSettleFlag = "2"
	OpenCloseSettleFlag_EXPECTED_PRICE                   OpenCloseSettleFlag = "3"
	OpenCloseSettleFlag_PRICE_FROM_PREVIOUS_BUSINESS_DAY OpenCloseSettleFlag = "4"
)

// OptPayoutType field enumeration values.
type OptPayoutType string

const (
	OptPayoutType_VANILLA              OptPayoutType = "1"
	OptPayoutType_CAPPED               OptPayoutType = "2"
	OptPayoutType_DIGITAL              OptPayoutType = "3"
	OptPayoutType_ASIAN                OptPayoutType = "4"
	OptPayoutType_BARRIER              OptPayoutType = "5"
	OptPayoutType_DIGITAL_BARRIER      OptPayoutType = "6"
	OptPayoutType_LOOKBACK             OptPayoutType = "7"
	OptPayoutType_OTHER_PATH_DEPENDENT OptPayoutType = "8"
	OptPayoutType_OTHER                OptPayoutType = "99"
)

// OptionExerciseDateType field enumeration values.
type OptionExerciseDateType string

const (
	OptionExerciseDateType_UNADJUSTED OptionExerciseDateType = "0"
	OptionExerciseDateType_ADJUSTED   OptionExerciseDateType = "1"
)

// OrdRejReason field enumeration values.
type OrdRejReason string

const (
	OrdRejReason_BROKER                                                     OrdRejReason = "0"
	OrdRejReason_UNKNOWN_SYMBOL                                             OrdRejReason = "1"
	OrdRejReason_INVALID_INVESTOR_ID                                        OrdRejReason = "10"
	OrdRejReason_UNSUPPORTED_ORDER_CHARACTERISTIC                           OrdRejReason = "11"
	OrdRejReason_SURVEILLANCE_OPTION                                        OrdRejReason = "12"
	OrdRejReason_INCORRECT_QUANTITY                                         OrdRejReason = "13"
	OrdRejReason_INCORRECT_ALLOCATED_QUANTITY                               OrdRejReason = "14"
	OrdRejReason_UNKNOWN_ACCOUNT                                            OrdRejReason = "15"
	OrdRejReason_PRICE_EXCEEDS_CURRENT_PRICE_BAND                           OrdRejReason = "16"
	OrdRejReason_INVALID_PRICE_INCREMENT                                    OrdRejReason = "18"
	OrdRejReason_REFERENCE_PRICE_NOT_AVAILABLE                              OrdRejReason = "19"
	OrdRejReason_EXCHANGE_CLOSED                                            OrdRejReason = "2"
	OrdRejReason_NOTIONAL_VALUE_EXCEEDS_THRESHOLD                           OrdRejReason = "20"
	OrdRejReason_ALGORITHM_RISK_THRESHOLD_BREACHED                          OrdRejReason = "21"
	OrdRejReason_SHORT_SELL_NOT_PERMITTED                                   OrdRejReason = "22"
	OrdRejReason_SHORT_SELL_REJECTED_DUE_TO_SECURITY_PRE_BORROW_RESTRICTION OrdRejReason = "23"
	OrdRejReason_SHORT_SELL_REJECTED_DUE_TO_ACCOUNT_PRE_BORROW_RESTRICTION  OrdRejReason = "24"
	OrdRejReason_INSUFFICIENT_CREDIT_LIMIT                                  OrdRejReason = "25"
	OrdRejReason_EXCEEDED_CLIP_SIZE_LIMIT                                   OrdRejReason = "26"
	OrdRejReason_EXCEEDED_MAXIMUM_NOTIONAL_ORDER_AMOUNT                     OrdRejReason = "27"
	OrdRejReason_EXCEEDED_DV01_PV01_LIMIT                                   OrdRejReason = "28"
	OrdRejReason_EXCEEDED_CS01_LIMIT                                        OrdRejReason = "29"
	OrdRejReason_ORDER_EXCEEDS_LIMIT                                        OrdRejReason = "3"
	OrdRejReason_TOO_LATE_TO_ENTER                                          OrdRejReason = "4"
	OrdRejReason_UNKNOWN_ORDER                                              OrdRejReason = "5"
	OrdRejReason_DUPLICATE_ORDER                                            OrdRejReason = "6"
	OrdRejReason_DUPLICATE_OF_A_VERBALLY_COMMUNICATED_ORDER                 OrdRejReason = "7"
	OrdRejReason_STALE_ORDER                                                OrdRejReason = "8"
	OrdRejReason_TRADE_ALONG_REQUIRED                                       OrdRejReason = "9"
	OrdRejReason_OTHER                                                      OrdRejReason = "99"
)

// OrdStatus field enumeration values.
type OrdStatus string

const (
	OrdStatus_NEW                  OrdStatus = "0"
	OrdStatus_PARTIALLY_FILLED     OrdStatus = "1"
	OrdStatus_FILLED               OrdStatus = "2"
	OrdStatus_DONE_FOR_DAY         OrdStatus = "3"
	OrdStatus_CANCELED             OrdStatus = "4"
	OrdStatus_REPLACED             OrdStatus = "5"
	OrdStatus_PENDING_CANCEL       OrdStatus = "6"
	OrdStatus_STOPPED              OrdStatus = "7"
	OrdStatus_REJECTED             OrdStatus = "8"
	OrdStatus_SUSPENDED            OrdStatus = "9"
	OrdStatus_PENDING_NEW          OrdStatus = "A"
	OrdStatus_CALCULATED           OrdStatus = "B"
	OrdStatus_EXPIRED              OrdStatus = "C"
	OrdStatus_ACCEPTED_FOR_BIDDING OrdStatus = "D"
	OrdStatus_PENDING_REPLACE      OrdStatus = "E"
)

// OrdType field enumeration values.
type OrdType string

const (
	OrdType_MARKET                         OrdType = "1"
	OrdType_LIMIT                          OrdType = "2"
	OrdType_STOP_STOP_LOSS                 OrdType = "3"
	OrdType_STOP_LIMIT                     OrdType = "4"
	OrdType_MARKET_ON_CLOSE                OrdType = "5"
	OrdType_WITH_OR_WITHOUT                OrdType = "6"
	OrdType_LIMIT_OR_BETTER                OrdType = "7"
	OrdType_LIMIT_WITH_OR_WITHOUT          OrdType = "8"
	OrdType_ON_BASIS                       OrdType = "9"
	OrdType_ON_CLOSE                       OrdType = "A"
	OrdType_LIMIT_ON_CLOSE                 OrdType = "B"
	OrdType_FOREX_MARKET                   OrdType = "C"
	OrdType_PREVIOUSLY_QUOTED              OrdType = "D"
	OrdType_PREVIOUSLY_INDICATED           OrdType = "E"
	OrdType_FOREX_LIMIT                    OrdType = "F"
	OrdType_FOREX_SWAP                     OrdType = "G"
	OrdType_FOREX_PREVIOUSLY_QUOTED        OrdType = "H"
	OrdType_FUNARI                         OrdType = "I"
	OrdType_MARKET_IF_TOUCHED              OrdType = "J"
	OrdType_MARKET_WITH_LEFT_OVER_AS_LIMIT OrdType = "K"
	OrdType_PREVIOUS_FUND_VALUATION_POINT  OrdType = "L"
	OrdType_NEXT_FUND_VALUATION_POINT      OrdType = "M"
	OrdType_PEGGED                         OrdType = "P"
	OrdType_COUNTER_ORDER_SELECTION        OrdType = "Q"
	OrdType_STOP_ON_BID_OR_OFFER           OrdType = "R"
	OrdType_STOP_LIMIT_ON_BID_OR_OFFER     OrdType = "S"
)

// OrderAttributeType field enumeration values.
type OrderAttributeType string

const (
	OrderAttributeType_AGGREGATED_ORDER                                           OrderAttributeType = "0"
	OrderAttributeType_ORDER_PENDING_ALLOCATION                                   OrderAttributeType = "1"
	OrderAttributeType_SUBJECT_TO_EU_SHARE_TRADING_OBLIGATION                     OrderAttributeType = "10"
	OrderAttributeType_SUBJECT_TO_UK_SHARE_TRADING_OBLIGATION                     OrderAttributeType = "11"
	OrderAttributeType_REPRESENTATIVE_ORDER                                       OrderAttributeType = "12"
	OrderAttributeType_LINKAGE_TYPE                                               OrderAttributeType = "13"
	OrderAttributeType_EXEMPT_FROM_SHARE_TRADING_OBLIGATION                       OrderAttributeType = "14"
	OrderAttributeType_LIQUIDITY_PROVISION_ACTIVITY_ORDER                         OrderAttributeType = "2"
	OrderAttributeType_RISK_REDUCTION_ORDER                                       OrderAttributeType = "3"
	OrderAttributeType_ALGORITHMIC_ORDER                                          OrderAttributeType = "4"
	OrderAttributeType_SYSTEMIC_INTERNALISER_ORDER                                OrderAttributeType = "5"
	OrderAttributeType_ALL_EXECUTIONS_FOR_THE_ORDER_ARE_TO_BE_SUBMITTED_TO_AN_APA OrderAttributeType = "6"
	OrderAttributeType_ORDER_EXECUTION_INSTRUCTED_BY_CLIENT                       OrderAttributeType = "7"
	OrderAttributeType_LARGE_IN_SCALE_ORDER                                       OrderAttributeType = "8"
	OrderAttributeType_HIDDEN_ORDER                                               OrderAttributeType = "9"
)

// OrderCapacity field enumeration values.
type OrderCapacity string

const (
	OrderCapacity_AGENCY                 OrderCapacity = "A"
	OrderCapacity_PROPRIETARY            OrderCapacity = "G"
	OrderCapacity_INDIVIDUAL             OrderCapacity = "I"
	OrderCapacity_MIXED_CAPACITY         OrderCapacity = "M"
	OrderCapacity_PRINCIPAL              OrderCapacity = "P"
	OrderCapacity_RISKLESS_PRINCIPAL     OrderCapacity = "R"
	OrderCapacity_AGENT_FOR_OTHER_MEMBER OrderCapacity = "W"
)

// OrderCategory field enumeration values.
type OrderCategory string

const (
	OrderCategory_ORDER                      OrderCategory = "1"
	OrderCategory_QUOTE                      OrderCategory = "2"
	OrderCategory_PRIVATELY_NEGOTIATED_TRADE OrderCategory = "3"
	OrderCategory_MULTILEG_ORDER             OrderCategory = "4"
	OrderCategory_LINKED_ORDER               OrderCategory = "5"
	OrderCategory_QUOTE_REQUEST              OrderCategory = "6"
	OrderCategory_IMPLIED_ORDER              OrderCategory = "7"
	OrderCategory_CROSS_ORDER                OrderCategory = "8"
	OrderCategory_STREAMING_PRICE            OrderCategory = "9"
	OrderCategory_INTERNAL_CROSS_ORDER       OrderCategory = "A"
)

// OrderDelayUnit field enumeration values.
type OrderDelayUnit string

const (
	OrderDelayUnit_SECONDS                OrderDelayUnit = "0"
	OrderDelayUnit_TENTHS_OF_A_SECOND     OrderDelayUnit = "1"
	OrderDelayUnit_MINUTES                OrderDelayUnit = "10"
	OrderDelayUnit_HOURS                  OrderDelayUnit = "11"
	OrderDelayUnit_DAYS                   OrderDelayUnit = "12"
	OrderDelayUnit_WEEKS                  OrderDelayUnit = "13"
	OrderDelayUnit_MONTHS                 OrderDelayUnit = "14"
	OrderDelayUnit_YEARS                  OrderDelayUnit = "15"
	OrderDelayUnit_HUNDREDTHS_OF_A_SECOND OrderDelayUnit = "2"
	OrderDelayUnit_MILLISECONDS           OrderDelayUnit = "3"
	OrderDelayUnit_MICROSECONDS           OrderDelayUnit = "4"
	OrderDelayUnit_NANOSECONDS            OrderDelayUnit = "5"
)

// OrderEntryAction field enumeration values.
type OrderEntryAction string

const (
	OrderEntryAction_ADD     OrderEntryAction = "1"
	OrderEntryAction_MODIFY  OrderEntryAction = "2"
	OrderEntryAction_DELETE  OrderEntryAction = "3"
	OrderEntryAction_SUSPEND OrderEntryAction = "4"
	OrderEntryAction_RELEASE OrderEntryAction = "5"
)

// OrderEventReason field enumeration values.
type OrderEventReason string

const (
	OrderEventReason_ADD_ORDER_REQUEST            OrderEventReason = "1"
	OrderEventReason_AWAY_MARKET_BETTER           OrderEventReason = "10"
	OrderEventReason_CORPORATE_ACTION             OrderEventReason = "11"
	OrderEventReason_START_OF_DAY                 OrderEventReason = "12"
	OrderEventReason_END_OF_DAY                   OrderEventReason = "13"
	OrderEventReason_MODIFY_ORDER_REQUEST         OrderEventReason = "2"
	OrderEventReason_DELETE_ORDER_REQUEST         OrderEventReason = "3"
	OrderEventReason_ORDER_ENTERED_OUT_OF_BAND    OrderEventReason = "4"
	OrderEventReason_ORDER_MODIFIED_OUT_OF_BAND   OrderEventReason = "5"
	OrderEventReason_ORDER_DELETED_OUT_OF_BAND    OrderEventReason = "6"
	OrderEventReason_ORDER_ACTIVATED_OR_TRIGGERED OrderEventReason = "7"
	OrderEventReason_ORDER_EXPIRED                OrderEventReason = "8"
	OrderEventReason_RESERVE_ORDER_REFRESHED      OrderEventReason = "9"
)

// OrderEventType field enumeration values.
type OrderEventType string

const (
	OrderEventType_ADDED            OrderEventType = "1"
	OrderEventType_TRIGGERED        OrderEventType = "10"
	OrderEventType_ACTIVATED        OrderEventType = "11"
	OrderEventType_MODIFIED         OrderEventType = "2"
	OrderEventType_DELETED          OrderEventType = "3"
	OrderEventType_PARTIALLY_FILLED OrderEventType = "4"
	OrderEventType_FILLED           OrderEventType = "5"
	OrderEventType_SUSPENDED        OrderEventType = "6"
	OrderEventType_RELEASED         OrderEventType = "7"
	OrderEventType_RESTATED         OrderEventType = "8"
	OrderEventType_LOCKED           OrderEventType = "9"
)

// OrderHandlingInstSource field enumeration values.
type OrderHandlingInstSource string

const (
	OrderHandlingInstSource_FINRA_OATS                OrderHandlingInstSource = "1"
	OrderHandlingInstSource_FIA_EXECUTION_SOURCE_CODE OrderHandlingInstSource = "2"
)

// OrderOrigination field enumeration values.
type OrderOrigination string

const (
	OrderOrigination_ORDER_RECEIVED_FROM_A_CUSTOMER                                    OrderOrigination = "1"
	OrderOrigination_ORDER_RECEIVED_FROM_WITHIN_THE_FIRM                               OrderOrigination = "2"
	OrderOrigination_ORDER_RECEIVED_FROM_ANOTHER_BROKER_DEALER                         OrderOrigination = "3"
	OrderOrigination_ORDER_RECEIVED_FROM_A_CUSTOMER_OR_ORIGINATED_FROM_WITHIN_THE_FIRM OrderOrigination = "4"
	OrderOrigination_ORDER_RECEIVED_FROM_A_DIRECT_ACCESS_OR_SPONSORED_ACCESS_CUSTOMER  OrderOrigination = "5"
	OrderOrigination_ORDER_RECEIVED_FROM_A_FOREIGN_DEALER_EQUIVALENT                   OrderOrigination = "6"
	OrderOrigination_ORDER_RECEIVED_FROM_AN_EXECUTION_ONLY_SERVICE                     OrderOrigination = "7"
)

// OrderOwnershipIndicator field enumeration values.
type OrderOwnershipIndicator string

const (
	OrderOwnershipIndicator_NO_CHANGE_OF_OWNERSHIP                 OrderOwnershipIndicator = "0"
	OrderOwnershipIndicator_CHANGE_OF_OWNERSHIP_TO_EXECUTING_PARTY OrderOwnershipIndicator = "1"
	OrderOwnershipIndicator_CHANGE_OF_OWNERSHIP_TO_ENTERING_PARTY  OrderOwnershipIndicator = "2"
	OrderOwnershipIndicator_CHANGE_OF_OWNERSHIP_TO_SPECIFIED_PARTY OrderOwnershipIndicator = "3"
)

// OrderRelationship field enumeration values.
type OrderRelationship string

const (
	OrderRelationship_NOT_SPECIFIED     OrderRelationship = "0"
	OrderRelationship_ORDER_AGGREGATION OrderRelationship = "1"
	OrderRelationship_ORDER_SPLIT       OrderRelationship = "2"
)

// OrderResponseLevel field enumeration values.
type OrderResponseLevel string

const (
	OrderResponseLevel_NO_ACKNOWLEDGEMENT      OrderResponseLevel = "0"
	OrderResponseLevel_MINIMUM_ACKNOWLEDGEMENT OrderResponseLevel = "1"
	OrderResponseLevel_ACKNOWLEDGE_EACH_ORDER  OrderResponseLevel = "2"
	OrderResponseLevel_SUMMARY_ACKNOWLEDGEMENT OrderResponseLevel = "3"
)

// OrderRestrictions field enumeration values.
type OrderRestrictions string

const (
	OrderRestrictions_PROGRAM_TRADE                                                                            OrderRestrictions = "1"
	OrderRestrictions_INDEX_ARBITRAGE                                                                          OrderRestrictions = "2"
	OrderRestrictions_NON_INDEX_ARBITRAGE                                                                      OrderRestrictions = "3"
	OrderRestrictions_COMPETING_MARKET_MAKER                                                                   OrderRestrictions = "4"
	OrderRestrictions_ACTING_AS_MARKET_MAKER_OR_SPECIALIST_IN_THE_SECURITY                                     OrderRestrictions = "5"
	OrderRestrictions_ACTING_AS_MARKET_MAKER_OR_SPECIALIST_IN_THE_UNDERLYING_SECURITY_OF_A_DERIVATIVE_SECURITY OrderRestrictions = "6"
	OrderRestrictions_FOREIGN_ENTITY                                                                           OrderRestrictions = "7"
	OrderRestrictions_EXTERNAL_MARKET_PARTICIPANT                                                              OrderRestrictions = "8"
	OrderRestrictions_EXTERNAL_INTER_CONNECTED_MARKET_LINKAGE                                                  OrderRestrictions = "9"
	OrderRestrictions_RISKLESS_ARBITRAGE                                                                       OrderRestrictions = "A"
	OrderRestrictions_ISSUER_HOLDING                                                                           OrderRestrictions = "B"
	OrderRestrictions_ISSUE_PRICE_STABILIZATION                                                                OrderRestrictions = "C"
	OrderRestrictions_NON_ALGORITHMIC                                                                          OrderRestrictions = "D"
	OrderRestrictions_ALGORITHMIC                                                                              OrderRestrictions = "E"
	OrderRestrictions_CROSS                                                                                    OrderRestrictions = "F"
	OrderRestrictions_INSIDER_ACCOUNT                                                                          OrderRestrictions = "G"
	OrderRestrictions_SIGNIFICANT_SHAREHOLDER                                                                  OrderRestrictions = "H"
	OrderRestrictions_NORMAL_COURSE_ISSUER_BID                                                                 OrderRestrictions = "I"
)

// OrigCustOrderCapacity field enumeration values.
type OrigCustOrderCapacity string

const (
	OrigCustOrderCapacity_MEMBER_TRADING_FOR_THEIR_OWN_ACCOUNT              OrigCustOrderCapacity = "1"
	OrigCustOrderCapacity_CLEARING_FIRM_TRADING_FOR_ITS_PROPRIETARY_ACCOUNT OrigCustOrderCapacity = "2"
	OrigCustOrderCapacity_MEMBER_TRADING_FOR_ANOTHER_MEMBER                 OrigCustOrderCapacity = "3"
	OrigCustOrderCapacity_ALL_OTHER                                         OrigCustOrderCapacity = "4"
)

// OwnerType field enumeration values.
type OwnerType string

const (
	OwnerType_INDIVIDUAL_INVESTOR                       OwnerType = "1"
	OwnerType_NETWORKING_SUB_ACCOUNT                    OwnerType = "10"
	OwnerType_NON_PROFIT_ORGANIZATION                   OwnerType = "11"
	OwnerType_CORPORATE_BODY                            OwnerType = "12"
	OwnerType_NOMINEE                                   OwnerType = "13"
	OwnerType_INSTITUTIONAL_CUSTOMER                    OwnerType = "14"
	OwnerType_COMBINED                                  OwnerType = "15"
	OwnerType_MEMBER_FIRM_EMPLOYEE_OR_ASSOCIATED_PERSON OwnerType = "16"
	OwnerType_MARKET_MAKING_ACCOUNT                     OwnerType = "17"
	OwnerType_PROPRIETARY_ACCOUNT                       OwnerType = "18"
	OwnerType_NON_BROKER_DEALER                         OwnerType = "19"
	OwnerType_PUBLIC_COMPANY                            OwnerType = "2"
	OwnerType_UNKNOWN_BENEFICIAL_OWNER_TYPE             OwnerType = "20"
	OwnerType_ERROR_ACCOUNT_OF_FIRM                     OwnerType = "21"
	OwnerType_FIRM_AGENCY_AVERAGE_PRICE_ACCOUNT         OwnerType = "22"
	OwnerType_PRIVATE_COMPANY                           OwnerType = "3"
	OwnerType_INDIVIDUAL_TRUSTEE                        OwnerType = "4"
	OwnerType_COMPANY_TRUSTEE                           OwnerType = "5"
	OwnerType_PENSION_PLAN                              OwnerType = "6"
	OwnerType_CUSTODIAN_UNDER_GIFTS_TO_MINORS_ACT       OwnerType = "7"
	OwnerType_TRUSTS                                    OwnerType = "8"
	OwnerType_FIDUCIARIES                               OwnerType = "9"
)

// OwnershipType field enumeration values.
type OwnershipType string

const (
	OwnershipType_JOINT_TRUSTEES    OwnershipType = "2"
	OwnershipType_JOINT_INVESTORS   OwnershipType = "J"
	OwnershipType_TENANTS_IN_COMMON OwnershipType = "T"
)

// PartyActionRejectReason field enumeration values.
type PartyActionRejectReason string

const (
	PartyActionRejectReason_INVALID_PARTY_OR_PARTIES PartyActionRejectReason = "0"
	PartyActionRejectReason_UNKNOWN_REQUESTING_PARTY PartyActionRejectReason = "1"
	PartyActionRejectReason_NOT_AUTHORIZED           PartyActionRejectReason = "98"
	PartyActionRejectReason_OTHER                    PartyActionRejectReason = "99"
)

// PartyActionResponse field enumeration values.
type PartyActionResponse string

const (
	PartyActionResponse_ACCEPTED  PartyActionResponse = "0"
	PartyActionResponse_COMPLETED PartyActionResponse = "1"
	PartyActionResponse_REJECTED  PartyActionResponse = "2"
)

// PartyActionType field enumeration values.
type PartyActionType string

const (
	PartyActionType_SUSPEND      PartyActionType = "0"
	PartyActionType_HALT_TRADING PartyActionType = "1"
	PartyActionType_REINSTATE    PartyActionType = "2"
)

// PartyDetailDefinitionStatus field enumeration values.
type PartyDetailDefinitionStatus string

const (
	PartyDetailDefinitionStatus_ACCEPTED              PartyDetailDefinitionStatus = "0"
	PartyDetailDefinitionStatus_ACCEPTED_WITH_CHANGES PartyDetailDefinitionStatus = "1"
	PartyDetailDefinitionStatus_REJECTED              PartyDetailDefinitionStatus = "2"
)

// PartyDetailRequestResult field enumeration values.
type PartyDetailRequestResult string

const (
	PartyDetailRequestResult_SUCCESSFUL            PartyDetailRequestResult = "0"
	PartyDetailRequestResult_INVALID_PARTY         PartyDetailRequestResult = "1"
	PartyDetailRequestResult_INVALID_RELATED_PARTY PartyDetailRequestResult = "2"
	PartyDetailRequestResult_INVALID_PARTY_STATUS  PartyDetailRequestResult = "3"
	PartyDetailRequestResult_NOT_AUTHORIZED        PartyDetailRequestResult = "98"
	PartyDetailRequestResult_OTHER                 PartyDetailRequestResult = "99"
)

// PartyDetailRequestStatus field enumeration values.
type PartyDetailRequestStatus string

const (
	PartyDetailRequestStatus_ACCEPTED              PartyDetailRequestStatus = "0"
	PartyDetailRequestStatus_ACCEPTED_WITH_CHANGES PartyDetailRequestStatus = "1"
	PartyDetailRequestStatus_REJECTED              PartyDetailRequestStatus = "2"
	PartyDetailRequestStatus_ACCEPTANCE_PENDING    PartyDetailRequestStatus = "3"
)

// PartyDetailRoleQualifier field enumeration values.
type PartyDetailRoleQualifier string

const (
	PartyDetailRoleQualifier_AGENCY                                    PartyDetailRoleQualifier = "0"
	PartyDetailRoleQualifier_PRINCIPAL                                 PartyDetailRoleQualifier = "1"
	PartyDetailRoleQualifier_ORIGINAL_TRADE_REPOSITORY                 PartyDetailRoleQualifier = "10"
	PartyDetailRoleQualifier_ADDITIONAL_INTERNATIONAL_TRADE_REPOSITORY PartyDetailRoleQualifier = "11"
	PartyDetailRoleQualifier_ADDITIONAL_DOMESTIC_TRADE_REPOSITORY      PartyDetailRoleQualifier = "12"
	PartyDetailRoleQualifier_RELATED_EXCHANGE                          PartyDetailRoleQualifier = "13"
	PartyDetailRoleQualifier_OPTIONS_EXCHANGE                          PartyDetailRoleQualifier = "14"
	PartyDetailRoleQualifier_SPECIFIED_EXCHANGE                        PartyDetailRoleQualifier = "15"
	PartyDetailRoleQualifier_CONSTITUENT_EXCHANGE                      PartyDetailRoleQualifier = "16"
	PartyDetailRoleQualifier_EXEMPT_FROM_TRADE_REPORTING               PartyDetailRoleQualifier = "17"
	PartyDetailRoleQualifier_CURRENT                                   PartyDetailRoleQualifier = "18"
	PartyDetailRoleQualifier_NEW                                       PartyDetailRoleQualifier = "19"
	PartyDetailRoleQualifier_RISKLESS_PRINCIPAL                        PartyDetailRoleQualifier = "2"
	PartyDetailRoleQualifier_DESIGNATED_SPONSOR                        PartyDetailRoleQualifier = "20"
	PartyDetailRoleQualifier_SPECIALIST                                PartyDetailRoleQualifier = "21"
	PartyDetailRoleQualifier_ALGORITHM                                 PartyDetailRoleQualifier = "22"
	PartyDetailRoleQualifier_FIRM_OR_LEGAL_ENTITY                      PartyDetailRoleQualifier = "23"
	PartyDetailRoleQualifier_NATURAL_PERSON                            PartyDetailRoleQualifier = "24"
	PartyDetailRoleQualifier_REGULAR_TRADER                            PartyDetailRoleQualifier = "25"
	PartyDetailRoleQualifier_HEAD_TRADER                               PartyDetailRoleQualifier = "26"
	PartyDetailRoleQualifier_SUPERVISOR                                PartyDetailRoleQualifier = "27"
	PartyDetailRoleQualifier_TRI_PARTY                                 PartyDetailRoleQualifier = "28"
	PartyDetailRoleQualifier_LENDER                                    PartyDetailRoleQualifier = "29"
	PartyDetailRoleQualifier_GENERAL_CLEARING_MEMBER                   PartyDetailRoleQualifier = "3"
	PartyDetailRoleQualifier_INDIVIDUAL_CLEARING_MEMBER                PartyDetailRoleQualifier = "4"
	PartyDetailRoleQualifier_PREFERRED_MARKET_MAKER                    PartyDetailRoleQualifier = "5"
	PartyDetailRoleQualifier_DIRECTED_MARKET_MAKER                     PartyDetailRoleQualifier = "6"
	PartyDetailRoleQualifier_BANK                                      PartyDetailRoleQualifier = "7"
	PartyDetailRoleQualifier_HUB                                       PartyDetailRoleQualifier = "8"
	PartyDetailRoleQualifier_PRIMARY_TRADE_REPOSITORY                  PartyDetailRoleQualifier = "9"
)

// PartyDetailStatus field enumeration values.
type PartyDetailStatus string

const (
	PartyDetailStatus_ACTIVE    PartyDetailStatus = "0"
	PartyDetailStatus_SUSPENDED PartyDetailStatus = "1"
	PartyDetailStatus_HALTED    PartyDetailStatus = "2"
)

// PartyIDSource field enumeration values.
type PartyIDSource string

const (
	PartyIDSource_KOREAN_INVESTOR_ID                                                                                PartyIDSource = "1"
	PartyIDSource_TAIWANESE_QUALIFIED_FOREIGN_INVESTOR_ID_QFII_FID                                                  PartyIDSource = "2"
	PartyIDSource_TAIWANESE_TRADING_ACCT                                                                            PartyIDSource = "3"
	PartyIDSource_MALAYSIAN_CENTRAL_DEPOSITORY                                                                      PartyIDSource = "4"
	PartyIDSource_CHINESE_INVESTOR_ID                                                                               PartyIDSource = "5"
	PartyIDSource_UK_NATIONAL_INSURANCE_OR_PENSION_NUMBER                                                           PartyIDSource = "6"
	PartyIDSource_US_SOCIAL_SECURITY_NUMBER                                                                         PartyIDSource = "7"
	PartyIDSource_US_EMPLOYER_OR_TAX_ID_NUMBER                                                                      PartyIDSource = "8"
	PartyIDSource_AUSTRALIAN_BUSINESS_NUMBER                                                                        PartyIDSource = "9"
	PartyIDSource_AUSTRALIAN_TAX_FILE_NUMBER                                                                        PartyIDSource = "A"
	PartyIDSource_BIC                                                                                               PartyIDSource = "B"
	PartyIDSource_GENERALLY_ACCEPTED_MARKET_PARTICIPANT_IDENTIFIER                                                  PartyIDSource = "C"
	PartyIDSource_PROPRIETARY                                                                                       PartyIDSource = "D"
	PartyIDSource_ISO_COUNTRY_CODE                                                                                  PartyIDSource = "E"
	PartyIDSource_SETTLEMENT_ENTITY_LOCATION                                                                        PartyIDSource = "F"
	PartyIDSource_MARKET_IDENTIFIER_CODE                                                                            PartyIDSource = "G"
	PartyIDSource_CSD_PARTICIPANT_MEMBER_CODE                                                                       PartyIDSource = "H"
	PartyIDSource_DIRECTED_BROKER_THREE_CHARACTER_ACRONYM_AS_DEFINED_IN_ISITC_ETC_BEST_PRACTICE_GUIDELINES_DOCUMENT PartyIDSource = "I"
	PartyIDSource_TAX_ID                                                                                            PartyIDSource = "J"
	PartyIDSource_AUSTRALIAN_COMPANY_NUMBER                                                                         PartyIDSource = "K"
	PartyIDSource_AUSTRALIAN_REGISTERED_BODY_NUMBER                                                                 PartyIDSource = "L"
	PartyIDSource_CFTC_REPORTING_FIRM_IDENTIFIER                                                                    PartyIDSource = "M"
	PartyIDSource_LEGAL_ENTITY_IDENTIFIER                                                                           PartyIDSource = "N"
	PartyIDSource_INTERIM_IDENTIFIER                                                                                PartyIDSource = "O"
	PartyIDSource_SHORT_CODE_IDENTIFIER                                                                             PartyIDSource = "P"
	PartyIDSource_NATIONAL_ID_OF_NATURAL_PERSON                                                                     PartyIDSource = "Q"
	PartyIDSource_INDIA_PERMANENT_ACCOUNT_NUMBER                                                                    PartyIDSource = "R"
	PartyIDSource_FIRM_DESIGNATED_IDENTIFIER                                                                        PartyIDSource = "S"
	PartyIDSource_SPECIAL_SEGREGATED_ACCOUNT_ID                                                                     PartyIDSource = "T"
	PartyIDSource_MASTER_SPECIAL_SEGREGATED_ACCOUNT_ID                                                              PartyIDSource = "U"
)

// PartyRelationship field enumeration values.
type PartyRelationship string

const (
	PartyRelationship_IS_ALSO                       PartyRelationship = "0"
	PartyRelationship_CLEARS_FOR                    PartyRelationship = "1"
	PartyRelationship_HAS_MEMBERS                   PartyRelationship = "10"
	PartyRelationship_PROVIDES_MARKETPLACE_FOR      PartyRelationship = "11"
	PartyRelationship_PARTICIPANT_OF_MARKETPLACE    PartyRelationship = "12"
	PartyRelationship_CARRIES_POSITIONS_FOR         PartyRelationship = "13"
	PartyRelationship_POSTS_TRADES_TO               PartyRelationship = "14"
	PartyRelationship_ENTERS_TRADES_FOR             PartyRelationship = "15"
	PartyRelationship_ENTERS_TRADES_THROUGH         PartyRelationship = "16"
	PartyRelationship_PROVIDES_QUOTES_TO            PartyRelationship = "17"
	PartyRelationship_REQUESTS_QUOTES_FROM          PartyRelationship = "18"
	PartyRelationship_INVESTS_FOR                   PartyRelationship = "19"
	PartyRelationship_CLEARS_THROUGH                PartyRelationship = "2"
	PartyRelationship_INVESTS_THROUGH               PartyRelationship = "20"
	PartyRelationship_BROKERS_TRADES_FOR            PartyRelationship = "21"
	PartyRelationship_BROKERS_TRADES_THROUGH        PartyRelationship = "22"
	PartyRelationship_PROVIDES_TRADING_SERVICES_FOR PartyRelationship = "23"
	PartyRelationship_USES_TRADING_SERVICES_OF      PartyRelationship = "24"
	PartyRelationship_APPROVES_OF                   PartyRelationship = "25"
	PartyRelationship_APPROVED_BY                   PartyRelationship = "26"
	PartyRelationship_PARENT_FIRM_FOR               PartyRelationship = "27"
	PartyRelationship_SUBSIDIARY_OF                 PartyRelationship = "28"
	PartyRelationship_REGULATORY_OWNER_OF           PartyRelationship = "29"
	PartyRelationship_TRADES_FOR                    PartyRelationship = "3"
	PartyRelationship_OWNED_BY_30                   PartyRelationship = "30"
	PartyRelationship_CONTROLS                      PartyRelationship = "31"
	PartyRelationship_IS_CONTROLLED_BY              PartyRelationship = "32"
	PartyRelationship_LEGAL                         PartyRelationship = "33"
	PartyRelationship_OWNED_BY_34                   PartyRelationship = "34"
	PartyRelationship_BENEFICIAL_OWNER_OF           PartyRelationship = "35"
	PartyRelationship_OWNED_BY_36                   PartyRelationship = "36"
	PartyRelationship_SETTLES_FOR                   PartyRelationship = "37"
	PartyRelationship_SETTLES_THROUGH               PartyRelationship = "38"
	PartyRelationship_TRADES_THROUGH                PartyRelationship = "4"
	PartyRelationship_SPONSORS                      PartyRelationship = "5"
	PartyRelationship_SPONSORED_THROUGH             PartyRelationship = "6"
	PartyRelationship_PROVIDES_GUARANTEE_FOR        PartyRelationship = "7"
	PartyRelationship_IS_GUARANTEED_BY              PartyRelationship = "8"
	PartyRelationship_MEMBER_OF                     PartyRelationship = "9"
)

// PartyRiskLimitStatus field enumeration values.
type PartyRiskLimitStatus string

const (
	PartyRiskLimitStatus_DISABLED PartyRiskLimitStatus = "0"
	PartyRiskLimitStatus_ENABLED  PartyRiskLimitStatus = "1"
)

// PartyRole field enumeration values.
type PartyRole string

const (
	PartyRole_EXECUTING_FIRM                                                        PartyRole = "1"
	PartyRole_SETTLEMENT_LOCATION                                                   PartyRole = "10"
	PartyRole_MARGIN_ACCOUNT                                                        PartyRole = "100"
	PartyRole_COLLATERAL_ASSET_ACCOUNT                                              PartyRole = "101"
	PartyRole_DATA_REPOSITORY                                                       PartyRole = "102"
	PartyRole_CALCULATION_AGENT                                                     PartyRole = "103"
	PartyRole_SENDER_OF_EXERCISE_NOTICE                                             PartyRole = "104"
	PartyRole_RECEIVER_OF_EXERCISE_NOTICE                                           PartyRole = "105"
	PartyRole_RATE_REFERENCE_BANK                                                   PartyRole = "106"
	PartyRole_CORRESPONDENT                                                         PartyRole = "107"
	PartyRole_BENEFICIARYS_BANK_OR_DEPOSITORY_INSTITUTION                           PartyRole = "109"
	PartyRole_ORDER_ORIGINATION_TRADER                                              PartyRole = "11"
	PartyRole_BORROWER                                                              PartyRole = "110"
	PartyRole_PRIMARY_OBLIGATOR                                                     PartyRole = "111"
	PartyRole_GUARANTOR                                                             PartyRole = "112"
	PartyRole_EXCLUDED_REFERENCE_ENTITY                                             PartyRole = "113"
	PartyRole_DETERMINING_PARTY                                                     PartyRole = "114"
	PartyRole_HEDGING_PARTY                                                         PartyRole = "115"
	PartyRole_REPORTING_ENTITY                                                      PartyRole = "116"
	PartyRole_SALES_PERSON                                                          PartyRole = "117"
	PartyRole_OPERATOR                                                              PartyRole = "118"
	PartyRole_CENTRAL_SECURITIES_DEPOSITORY_119                                     PartyRole = "119"
	PartyRole_EXECUTING_TRADER                                                      PartyRole = "12"
	PartyRole_INTERNATIONAL_CENTRAL_SECURITIES_DEPOSITORY                           PartyRole = "120"
	PartyRole_TRADING_SUB_ACCOUNT                                                   PartyRole = "121"
	PartyRole_INVESTMENT_DECISION_MAKER                                             PartyRole = "122"
	PartyRole_PUBLISHING_INTERMEDIARY                                               PartyRole = "123"
	PartyRole_CENTRAL_SECURITIES_DEPOSITORY_124                                     PartyRole = "124"
	PartyRole_ISSUER                                                                PartyRole = "125"
	PartyRole_CONTRA_CUSTOMER_ACCOUNT                                               PartyRole = "126"
	PartyRole_CONTRA_INVESTMENT_DECISION_MAKER                                      PartyRole = "127"
	PartyRole_ORDER_ORIGINATION_FIRM                                                PartyRole = "13"
	PartyRole_GIVEUP_CLEARING_FIRM                                                  PartyRole = "14"
	PartyRole_CORRESPONDANT_CLEARING_FIRM                                           PartyRole = "15"
	PartyRole_EXECUTING_SYSTEM                                                      PartyRole = "16"
	PartyRole_CONTRA_FIRM                                                           PartyRole = "17"
	PartyRole_CONTRA_CLEARING_FIRM                                                  PartyRole = "18"
	PartyRole_SPONSORING_FIRM                                                       PartyRole = "19"
	PartyRole_BROKER_OF_CREDIT                                                      PartyRole = "2"
	PartyRole_UNDERLYING_CONTRA_FIRM                                                PartyRole = "20"
	PartyRole_CLEARING_ORGANIZATION                                                 PartyRole = "21"
	PartyRole_EXCHANGE                                                              PartyRole = "22"
	PartyRole_CUSTOMER_ACCOUNT                                                      PartyRole = "24"
	PartyRole_CORRESPONDENT_CLEARING_ORGANIZATION                                   PartyRole = "25"
	PartyRole_CORRESPONDENT_BROKER                                                  PartyRole = "26"
	PartyRole_BUYER_SELLER                                                          PartyRole = "27"
	PartyRole_CUSTODIAN                                                             PartyRole = "28"
	PartyRole_INTERMEDIARY                                                          PartyRole = "29"
	PartyRole_CLIENT_ID                                                             PartyRole = "3"
	PartyRole_AGENT                                                                 PartyRole = "30"
	PartyRole_SUB_CUSTODIAN                                                         PartyRole = "31"
	PartyRole_BENEFICIARY                                                           PartyRole = "32"
	PartyRole_INTERESTED_PARTY                                                      PartyRole = "33"
	PartyRole_REGULATORY_BODY                                                       PartyRole = "34"
	PartyRole_LIQUIDITY_PROVIDER                                                    PartyRole = "35"
	PartyRole_ENTERING_TRADER                                                       PartyRole = "36"
	PartyRole_CONTRA_TRADER                                                         PartyRole = "37"
	PartyRole_POSITION_ACCOUNT                                                      PartyRole = "38"
	PartyRole_CONTRA_INVESTOR_ID                                                    PartyRole = "39"
	PartyRole_CLEARING_FIRM                                                         PartyRole = "4"
	PartyRole_TRANSFER_TO_FIRM                                                      PartyRole = "40"
	PartyRole_CONTRA_POSITION_ACCOUNT                                               PartyRole = "41"
	PartyRole_CONTRA_EXCHANGE                                                       PartyRole = "42"
	PartyRole_INTERNAL_CARRY_ACCOUNT                                                PartyRole = "43"
	PartyRole_ORDER_ENTRY_OPERATOR_ID                                               PartyRole = "44"
	PartyRole_SECONDARY_ACCOUNT_NUMBER                                              PartyRole = "45"
	PartyRole_FOREIGN_FIRM                                                          PartyRole = "46"
	PartyRole_THIRD_PARTY_ALLOCATION_FIRM                                           PartyRole = "47"
	PartyRole_CLAIMING_ACCOUNT                                                      PartyRole = "48"
	PartyRole_ASSET_MANAGER                                                         PartyRole = "49"
	PartyRole_INVESTOR_ID                                                           PartyRole = "5"
	PartyRole_PLEDGOR_ACCOUNT                                                       PartyRole = "50"
	PartyRole_PLEDGEE_ACCOUNT                                                       PartyRole = "51"
	PartyRole_LARGE_TRADER_REPORTABLE_ACCOUNT                                       PartyRole = "52"
	PartyRole_TRADER_MNEMONIC                                                       PartyRole = "53"
	PartyRole_SENDER_LOCATION                                                       PartyRole = "54"
	PartyRole_SESSION_ID                                                            PartyRole = "55"
	PartyRole_ACCEPTABLE_COUNTERPARTY                                               PartyRole = "56"
	PartyRole_UNACCEPTABLE_COUNTERPARTY                                             PartyRole = "57"
	PartyRole_ENTERING_UNIT                                                         PartyRole = "58"
	PartyRole_EXECUTING_UNIT                                                        PartyRole = "59"
	PartyRole_INTRODUCING_FIRM                                                      PartyRole = "6"
	PartyRole_INTRODUCING_BROKER                                                    PartyRole = "60"
	PartyRole_QUOTE_ORIGINATOR                                                      PartyRole = "61"
	PartyRole_REPORT_ORIGINATOR                                                     PartyRole = "62"
	PartyRole_SYSTEMATIC_INTERNALISER                                               PartyRole = "63"
	PartyRole_MULTILATERAL_TRADING_FACILITY                                         PartyRole = "64"
	PartyRole_REGULATED_MARKET                                                      PartyRole = "65"
	PartyRole_MARKET_MAKER                                                          PartyRole = "66"
	PartyRole_INVESTMENT_FIRM                                                       PartyRole = "67"
	PartyRole_HOST_COMPETENT_AUTHORITY                                              PartyRole = "68"
	PartyRole_HOME_COMPETENT_AUTHORITY                                              PartyRole = "69"
	PartyRole_ENTERING_FIRM                                                         PartyRole = "7"
	PartyRole_COMPETENT_AUTHORITY_OF_THE_MOST_RELEVANT_MARKET_IN_TERMS_OF_LIQUIDITY PartyRole = "70"
	PartyRole_COMPETENT_AUTHORITY_OF_THE_TRANSACTION                                PartyRole = "71"
	PartyRole_REPORTING_INTERMEDIARY                                                PartyRole = "72"
	PartyRole_EXECUTION_VENUE                                                       PartyRole = "73"
	PartyRole_MARKET_DATA_ENTRY_ORIGINATOR                                          PartyRole = "74"
	PartyRole_LOCATION_ID                                                           PartyRole = "75"
	PartyRole_DESK_ID                                                               PartyRole = "76"
	PartyRole_MARKET_DATA_MARKET                                                    PartyRole = "77"
	PartyRole_ALLOCATION_ENTITY                                                     PartyRole = "78"
	PartyRole_PRIME_BROKER_PROVIDING_GENERAL_TRADE_SERVICES                         PartyRole = "79"
	PartyRole_LOCATE                                                                PartyRole = "8"
	PartyRole_STEP_OUT_FIRM                                                         PartyRole = "80"
	PartyRole_BROKER_CIENT_ID                                                       PartyRole = "81"
	PartyRole_CENTRAL_REGISTRATION_DEPOSITORY                                       PartyRole = "82"
	PartyRole_CLEARING_ACCOUNT                                                      PartyRole = "83"
	PartyRole_ACCEPTABLE_SETTLING_COUNTERPARTY                                      PartyRole = "84"
	PartyRole_UNACCEPTABLE_SETTLING_COUNTERPARTY                                    PartyRole = "85"
	PartyRole_CLS_MEMBER_BANK                                                       PartyRole = "86"
	PartyRole_IN_CONCERT_GROUP                                                      PartyRole = "87"
	PartyRole_IN_CONCERT_CONTROLLING_ENTITY                                         PartyRole = "88"
	PartyRole_LARGE_POSITIONS_REPORTING_ACCOUNT                                     PartyRole = "89"
	PartyRole_FUND_MANAGER_CLIENT_ID                                                PartyRole = "9"
	PartyRole_SETTLEMENT_FIRM                                                       PartyRole = "90"
	PartyRole_SETTLEMENT_ACCOUNT                                                    PartyRole = "91"
	PartyRole_REPORTING_MARKET_CENTER                                               PartyRole = "92"
	PartyRole_RELATED_REPORTING_MARKET_CENTER                                       PartyRole = "93"
	PartyRole_AWAY_MARKET                                                           PartyRole = "94"
	PartyRole_GIVE_UP                                                               PartyRole = "95"
	PartyRole_TAKE_UP                                                               PartyRole = "96"
	PartyRole_GIVE_UP_CLEARING_FIRM                                                 PartyRole = "97"
	PartyRole_TAKE_UP_CLEARING_FIRM                                                 PartyRole = "98"
	PartyRole_ORIGINATING_MARKET                                                    PartyRole = "99"
)

// PartySubIDType field enumeration values.
type PartySubIDType string

const (
	PartySubIDType_FIRM                                                                      PartySubIDType = "1"
	PartySubIDType_SECURITIES_ACCOUNT_NUMBER                                                 PartySubIDType = "10"
	PartySubIDType_REGISTRATION_NUMBER                                                       PartySubIDType = "11"
	PartySubIDType_REGISTERED_ADDRESS_12                                                     PartySubIDType = "12"
	PartySubIDType_REGULATORY_STATUS                                                         PartySubIDType = "13"
	PartySubIDType_REGISTRATION_NAME                                                         PartySubIDType = "14"
	PartySubIDType_CASH_ACCOUNT_NUMBER                                                       PartySubIDType = "15"
	PartySubIDType_BIC                                                                       PartySubIDType = "16"
	PartySubIDType_CSD_PARTICIPANT_MEMBER_CODE                                               PartySubIDType = "17"
	PartySubIDType_REGISTERED_ADDRESS_18                                                     PartySubIDType = "18"
	PartySubIDType_FUND_ACCOUNT_NAME                                                         PartySubIDType = "19"
	PartySubIDType_PERSON                                                                    PartySubIDType = "2"
	PartySubIDType_TELEX_NUMBER                                                              PartySubIDType = "20"
	PartySubIDType_FAX_NUMBER                                                                PartySubIDType = "21"
	PartySubIDType_SECURITIES_ACCOUNT_NAME                                                   PartySubIDType = "22"
	PartySubIDType_CASH_ACCOUNT_NAME                                                         PartySubIDType = "23"
	PartySubIDType_DEPARTMENT                                                                PartySubIDType = "24"
	PartySubIDType_LOCATION_DESK                                                             PartySubIDType = "25"
	PartySubIDType_POSITION_ACCOUNT_TYPE                                                     PartySubIDType = "26"
	PartySubIDType_SECURITY_LOCATE_ID                                                        PartySubIDType = "27"
	PartySubIDType_MARKET_MAKER                                                              PartySubIDType = "28"
	PartySubIDType_ELIGIBLE_COUNTERPARTY                                                     PartySubIDType = "29"
	PartySubIDType_SYSTEM                                                                    PartySubIDType = "3"
	PartySubIDType_PROFESSIONAL_CLIENT                                                       PartySubIDType = "30"
	PartySubIDType_LOCATION                                                                  PartySubIDType = "31"
	PartySubIDType_EXECUTION_VENUE                                                           PartySubIDType = "32"
	PartySubIDType_CURRENCY_DELIVERY_IDENTIFIER                                              PartySubIDType = "33"
	PartySubIDType_ADDRESS_CITY                                                              PartySubIDType = "34"
	PartySubIDType_ADDRESS_STATE_PROVINCE                                                    PartySubIDType = "35"
	PartySubIDType_ADDRESS_POSTAL_CODE                                                       PartySubIDType = "36"
	PartySubIDType_ADDRESS_STREET                                                            PartySubIDType = "37"
	PartySubIDType_ADDRESS_COUNTRY                                                           PartySubIDType = "38"
	PartySubIDType_ISO_COUNTRY_CODE                                                          PartySubIDType = "39"
	PartySubIDType_APPLICATION                                                               PartySubIDType = "4"
	PartySubIDType_MARKET_SEGMENT                                                            PartySubIDType = "40"
	PartySubIDType_RESERVEDANDAVAILABLEFORBILATERALLYAGREEDUPONUSERDEFINEDVALUES             PartySubIDType = "4000"
	PartySubIDType_CUSTOMER_ACCOUNT_TYPE                                                     PartySubIDType = "41"
	PartySubIDType_OMNIBUS_ACCOUNT                                                           PartySubIDType = "42"
	PartySubIDType_FUNDS_SEGREGATION_TYPE                                                    PartySubIDType = "43"
	PartySubIDType_GUARANTEE_FUND                                                            PartySubIDType = "44"
	PartySubIDType_SWAP_DEALER                                                               PartySubIDType = "45"
	PartySubIDType_MAJOR_PARTICIPANT                                                         PartySubIDType = "46"
	PartySubIDType_FINANCIAL_ENTITY                                                          PartySubIDType = "47"
	PartySubIDType_US_PERSON                                                                 PartySubIDType = "48"
	PartySubIDType_REPORTING_ENTITY_INDICATOR                                                PartySubIDType = "49"
	PartySubIDType_FULL_LEGAL_NAME_OF_FIRM                                                   PartySubIDType = "5"
	PartySubIDType_ELECTED_CLEARING_REQUIREMENT_EXCEPTION                                    PartySubIDType = "50"
	PartySubIDType_BUSINESS_CENTER                                                           PartySubIDType = "51"
	PartySubIDType_REFERENCE_TEXT                                                            PartySubIDType = "52"
	PartySubIDType_SHORT_MARKING_EXEMPT_ACCOUNT                                              PartySubIDType = "53"
	PartySubIDType_PARENT_FIRM_IDENTIFIER                                                    PartySubIDType = "54"
	PartySubIDType_PARENT_FIRM_NAME                                                          PartySubIDType = "55"
	PartySubIDType_DEAL_IDENTIFIER                                                           PartySubIDType = "56"
	PartySubIDType_SYSTEM_TRADE_IDENTIFIER                                                   PartySubIDType = "57"
	PartySubIDType_SYSTEM_TRADE_SUB_IDENTIFIER                                               PartySubIDType = "58"
	PartySubIDType_FUTURES_COMMISSION_MERCHANT                                               PartySubIDType = "59"
	PartySubIDType_POSTAL_ADDRESS                                                            PartySubIDType = "6"
	PartySubIDType_DELIVERY_TERMINAL_CUSTOMER_ACCOUNT_CODE                                   PartySubIDType = "60"
	PartySubIDType_VOLUNTARY_REPORTING_ENTITY                                                PartySubIDType = "61"
	PartySubIDType_REPORTING_OBLIGATION_JURISDICTION                                         PartySubIDType = "62"
	PartySubIDType_VOLUNTARY_REPORTING_JURISDICTION                                          PartySubIDType = "63"
	PartySubIDType_COMPANY_ACTIVITIES                                                        PartySubIDType = "64"
	PartySubIDType_EUROPEAN_ECONOMIC_AREA_DOMICILED                                          PartySubIDType = "65"
	PartySubIDType_CONTRACT_LINKED_TO_COMMERCIAL_OR_TREASURY_FINANCING_FOR_THIS_COUNTERPARTY PartySubIDType = "66"
	PartySubIDType_CONTRACT_ABOVE_CLEARING_THRESHOLD_FOR_THIS_COUNTERPARTY                   PartySubIDType = "67"
	PartySubIDType_VOLUNTARY_REPORTING_PARTY                                                 PartySubIDType = "68"
	PartySubIDType_END_USER                                                                  PartySubIDType = "69"
	PartySubIDType_PHONE_NUMBER                                                              PartySubIDType = "7"
	PartySubIDType_LOCATION_OR_JURISDICTION                                                  PartySubIDType = "70"
	PartySubIDType_DERIVATIVES_DEALER                                                        PartySubIDType = "71"
	PartySubIDType_DOMICILE                                                                  PartySubIDType = "72"
	PartySubIDType_EXEMPT_FROM_RECOGNITION                                                   PartySubIDType = "73"
	PartySubIDType_PAYER                                                                     PartySubIDType = "74"
	PartySubIDType_RECEIVER                                                                  PartySubIDType = "75"
	PartySubIDType_SYSTEMATIC_INTERNALISER                                                   PartySubIDType = "76"
	PartySubIDType_PUBLISHING_ENTITY_INDICATOR                                               PartySubIDType = "77"
	PartySubIDType_FIRST_NAME                                                                PartySubIDType = "78"
	PartySubIDType_SURNAME                                                                   PartySubIDType = "79"
	PartySubIDType_EMAIL_ADDRESS                                                             PartySubIDType = "8"
	PartySubIDType_DATE_OF_BIRTH                                                             PartySubIDType = "80"
	PartySubIDType_ORDER_TRANSMITTING_FIRM                                                   PartySubIDType = "81"
	PartySubIDType_ORDER_TRANSMITTING_FIRM_FOR_BUYER                                         PartySubIDType = "82"
	PartySubIDType_ORDER_TRANSMITTER_FOR_SELLER                                              PartySubIDType = "83"
	PartySubIDType_LEGAL_ENTITY_IDENTIFIER                                                   PartySubIDType = "84"
	PartySubIDType_SUB_SECTOR_CLASSIFICATION                                                 PartySubIDType = "85"
	PartySubIDType_PARTY_SIDE                                                                PartySubIDType = "86"
	PartySubIDType_LEGAL_REGISTRATION_COUNTRY                                                PartySubIDType = "87"
	PartySubIDType_CONTACT_NAME                                                              PartySubIDType = "9"
)

// PayReportStatus field enumeration values.
type PayReportStatus string

const (
	PayReportStatus_RECEIVED_NOT_YET_PROCESSED PayReportStatus = "0"
	PayReportStatus_ACCEPTED                   PayReportStatus = "1"
	PayReportStatus_REJECTED                   PayReportStatus = "2"
	PayReportStatus_DISPUTED                   PayReportStatus = "3"
)

// PayReportTransType field enumeration values.
type PayReportTransType string

const (
	PayReportTransType_NEW     PayReportTransType = "0"
	PayReportTransType_REPLACE PayReportTransType = "1"
	PayReportTransType_STATUS  PayReportTransType = "2"
)

// PayRequestStatus field enumeration values.
type PayRequestStatus string

const (
	PayRequestStatus_RECEIVED_NOT_YET_PROCESSED PayRequestStatus = "0"
	PayRequestStatus_ACCEPTED                   PayRequestStatus = "1"
	PayRequestStatus_REJECTED                   PayRequestStatus = "2"
	PayRequestStatus_DISPUTED                   PayRequestStatus = "3"
)

// PayRequestTransType field enumeration values.
type PayRequestTransType string

const (
	PayRequestTransType_NEW    PayRequestTransType = "0"
	PayRequestTransType_CANCEL PayRequestTransType = "1"
)

// PaymentDateOffsetDayType field enumeration values.
type PaymentDateOffsetDayType string

const (
	PaymentDateOffsetDayType_BUSINESS              PaymentDateOffsetDayType = "0"
	PaymentDateOffsetDayType_CALENDAR              PaymentDateOffsetDayType = "1"
	PaymentDateOffsetDayType_COMMODITY_BUSINESS    PaymentDateOffsetDayType = "2"
	PaymentDateOffsetDayType_CURRENCY_BUSINESS     PaymentDateOffsetDayType = "3"
	PaymentDateOffsetDayType_EXCHANGE_BUSINESS     PaymentDateOffsetDayType = "4"
	PaymentDateOffsetDayType_SCHEDULED_TRADING_DAY PaymentDateOffsetDayType = "5"
)

// PaymentForwardStartType field enumeration values.
type PaymentForwardStartType string

const (
	PaymentForwardStartType_PREPAID   PaymentForwardStartType = "0"
	PaymentForwardStartType_POST_PAID PaymentForwardStartType = "1"
	PaymentForwardStartType_VARIABLE  PaymentForwardStartType = "2"
	PaymentForwardStartType_FIXED     PaymentForwardStartType = "3"
)

// PaymentMethod field enumeration values.
type PaymentMethod string

const (
	PaymentMethod_CREST                      PaymentMethod = "1"
	PaymentMethod_DIRECT_CREDIT              PaymentMethod = "10"
	PaymentMethod_CREDIT_CARD                PaymentMethod = "11"
	PaymentMethod_ACH_DEBIT                  PaymentMethod = "12"
	PaymentMethod_ACH_CREDIT                 PaymentMethod = "13"
	PaymentMethod_BPAY                       PaymentMethod = "14"
	PaymentMethod_HIGH_VALUE_CLEARING_SYSTEM PaymentMethod = "15"
	PaymentMethod_CHIPS                      PaymentMethod = "16"
	PaymentMethod_SWIFT                      PaymentMethod = "17"
	PaymentMethod_CHAPS                      PaymentMethod = "18"
	PaymentMethod_SIC                        PaymentMethod = "19"
	PaymentMethod_NSCC                       PaymentMethod = "2"
	PaymentMethod_EUROSIC                    PaymentMethod = "20"
	PaymentMethod_EUROCLEAR                  PaymentMethod = "3"
	PaymentMethod_CLEARSTREAM                PaymentMethod = "4"
	PaymentMethod_CHEQUE                     PaymentMethod = "5"
	PaymentMethod_TELEGRAPHIC_TRANSFER       PaymentMethod = "6"
	PaymentMethod_FED_WIRE                   PaymentMethod = "7"
	PaymentMethod_DEBIT_CARD                 PaymentMethod = "8"
	PaymentMethod_DIRECT_DEBIT               PaymentMethod = "9"
)

// PaymentPaySide field enumeration values.
type PaymentPaySide string

const (
	PaymentPaySide_BUY  PaymentPaySide = "1"
	PaymentPaySide_SELL PaymentPaySide = "2"
)

// PaymentScheduleStepRelativeTo field enumeration values.
type PaymentScheduleStepRelativeTo string

const (
	PaymentScheduleStepRelativeTo_INITIAL  PaymentScheduleStepRelativeTo = "0"
	PaymentScheduleStepRelativeTo_PREVIOUS PaymentScheduleStepRelativeTo = "1"
)

// PaymentScheduleType field enumeration values.
type PaymentScheduleType string

const (
	PaymentScheduleType_NOTIONAL                                     PaymentScheduleType = "0"
	PaymentScheduleType_CASH_FLOW                                    PaymentScheduleType = "1"
	PaymentScheduleType_NON_DELIVERABLE_SETTLEMENT_PAYMENT_DATES     PaymentScheduleType = "10"
	PaymentScheduleType_NON_DELIVERABLE_SETTLEMENT_CALCULATION_DATES PaymentScheduleType = "11"
	PaymentScheduleType_NON_DELIVERABLE_FIXING_DATES                 PaymentScheduleType = "12"
	PaymentScheduleType_SETTLEMENT_PERIOD_NOTIONAL                   PaymentScheduleType = "13"
	PaymentScheduleType_SETTLEMENT_PERIOD_PRICE                      PaymentScheduleType = "14"
	PaymentScheduleType_CALCULATION_PERIOD                           PaymentScheduleType = "15"
	PaymentScheduleType_DIVIDEND_ACCRUAL_RATE_MULTIPLIER             PaymentScheduleType = "16"
	PaymentScheduleType_DIVIDEND_ACCRUAL_RATE_SPREAD                 PaymentScheduleType = "17"
	PaymentScheduleType_DIVIDEND_ACCRUAL_CAP_RATE                    PaymentScheduleType = "18"
	PaymentScheduleType_DIVIDEND_ACCRUAL_FLOOR_RATE                  PaymentScheduleType = "19"
	PaymentScheduleType_FX_LINKED_NOTIONAL                           PaymentScheduleType = "2"
	PaymentScheduleType_COMPOUNDING_RATE_MULTIPLIER                  PaymentScheduleType = "20"
	PaymentScheduleType_COMPOUNDING_RATE_SPREAD                      PaymentScheduleType = "21"
	PaymentScheduleType_COMPOUNDING_CAP_RATE                         PaymentScheduleType = "22"
	PaymentScheduleType_COMPOUNDING_FLOOR_RATE                       PaymentScheduleType = "23"
	PaymentScheduleType_FIXED_RATE                                   PaymentScheduleType = "3"
	PaymentScheduleType_FUTURE_VALUE_NOTIONAL                        PaymentScheduleType = "4"
	PaymentScheduleType_KNOWN_AMOUNT                                 PaymentScheduleType = "5"
	PaymentScheduleType_FLOATING_RATE_MULTIPLIER                     PaymentScheduleType = "6"
	PaymentScheduleType_SPREAD                                       PaymentScheduleType = "7"
	PaymentScheduleType_CAP_RATE                                     PaymentScheduleType = "8"
	PaymentScheduleType_FLOOR_RATE                                   PaymentScheduleType = "9"
)

// PaymentSettlStyle field enumeration values.
type PaymentSettlStyle string

const (
	PaymentSettlStyle_STANDARD         PaymentSettlStyle = "0"
	PaymentSettlStyle_NET              PaymentSettlStyle = "1"
	PaymentSettlStyle_STANDARD_AND_NET PaymentSettlStyle = "2"
)

// PaymentStreamAveragingMethod field enumeration values.
type PaymentStreamAveragingMethod string

const (
	PaymentStreamAveragingMethod_UNWEIGHTED PaymentStreamAveragingMethod = "0"
	PaymentStreamAveragingMethod_WEIGHTED   PaymentStreamAveragingMethod = "1"
)

// PaymentStreamCapRateBuySide field enumeration values.
type PaymentStreamCapRateBuySide string

const (
	PaymentStreamCapRateBuySide_BUYER_OF_THE_TRADE  PaymentStreamCapRateBuySide = "1"
	PaymentStreamCapRateBuySide_SELLER_OF_THE_TRADE PaymentStreamCapRateBuySide = "2"
)

// PaymentStreamCompoundingMethod field enumeration values.
type PaymentStreamCompoundingMethod string

const (
	PaymentStreamCompoundingMethod_NONE             PaymentStreamCompoundingMethod = "0"
	PaymentStreamCompoundingMethod_FLAT             PaymentStreamCompoundingMethod = "1"
	PaymentStreamCompoundingMethod_STRAIGHT         PaymentStreamCompoundingMethod = "2"
	PaymentStreamCompoundingMethod_SPREAD_EXCLUSIVE PaymentStreamCompoundingMethod = "3"
)

// PaymentStreamDiscountType field enumeration values.
type PaymentStreamDiscountType string

const (
	PaymentStreamDiscountType_STANDARD               PaymentStreamDiscountType = "0"
	PaymentStreamDiscountType_FORWARD_RATE_AGREEMENT PaymentStreamDiscountType = "1"
)

// PaymentStreamFRADiscounting field enumeration values.
type PaymentStreamFRADiscounting string

const (
	PaymentStreamFRADiscounting_NONE                                            PaymentStreamFRADiscounting = "0"
	PaymentStreamFRADiscounting_INTERNATIONAL_SWAPS_AND_DERIVATIVES_ASSOCIATION PaymentStreamFRADiscounting = "1"
	PaymentStreamFRADiscounting_AUSTRALIAN_FINANCIAL_MARKETS_ASSOCIATION        PaymentStreamFRADiscounting = "2"
)

// PaymentStreamFloorRateBuySide field enumeration values.
type PaymentStreamFloorRateBuySide string

const (
	PaymentStreamFloorRateBuySide_BUYER_OF_THE_TRADE  PaymentStreamFloorRateBuySide = "1"
	PaymentStreamFloorRateBuySide_SELLER_OF_THE_TRADE PaymentStreamFloorRateBuySide = "2"
)

// PaymentStreamInflationInterpolationMethod field enumeration values.
type PaymentStreamInflationInterpolationMethod string

const (
	PaymentStreamInflationInterpolationMethod_NONE              PaymentStreamInflationInterpolationMethod = "0"
	PaymentStreamInflationInterpolationMethod_LINEAR_ZERO_YIELD PaymentStreamInflationInterpolationMethod = "1"
)

// PaymentStreamInflationLagDayType field enumeration values.
type PaymentStreamInflationLagDayType string

const (
	PaymentStreamInflationLagDayType_BUSINESS              PaymentStreamInflationLagDayType = "0"
	PaymentStreamInflationLagDayType_CALENDAR              PaymentStreamInflationLagDayType = "1"
	PaymentStreamInflationLagDayType_COMMODITY_BUSINESS    PaymentStreamInflationLagDayType = "2"
	PaymentStreamInflationLagDayType_CURRENCY_BUSINESS     PaymentStreamInflationLagDayType = "3"
	PaymentStreamInflationLagDayType_EXCHANGE_BUSINESS     PaymentStreamInflationLagDayType = "4"
	PaymentStreamInflationLagDayType_SCHEDULED_TRADING_DAY PaymentStreamInflationLagDayType = "5"
)

// PaymentStreamInflationLagUnit field enumeration values.
type PaymentStreamInflationLagUnit string

const (
	PaymentStreamInflationLagUnit_DAY   PaymentStreamInflationLagUnit = "D"
	PaymentStreamInflationLagUnit_MONTH PaymentStreamInflationLagUnit = "Mo"
	PaymentStreamInflationLagUnit_WEEK  PaymentStreamInflationLagUnit = "Wk"
	PaymentStreamInflationLagUnit_YEAR  PaymentStreamInflationLagUnit = "Yr"
)

// PaymentStreamInterpolationPeriod field enumeration values.
type PaymentStreamInterpolationPeriod string

const (
	PaymentStreamInterpolationPeriod_INITIAL           PaymentStreamInterpolationPeriod = "0"
	PaymentStreamInterpolationPeriod_INITIAL_AND_FINAL PaymentStreamInterpolationPeriod = "1"
	PaymentStreamInterpolationPeriod_FINAL             PaymentStreamInterpolationPeriod = "2"
	PaymentStreamInterpolationPeriod_ANY_PERIOD        PaymentStreamInterpolationPeriod = "3"
)

// PaymentStreamLinkStrikePriceType field enumeration values.
type PaymentStreamLinkStrikePriceType string

const (
	PaymentStreamLinkStrikePriceType_VOLATILITY PaymentStreamLinkStrikePriceType = "0"
	PaymentStreamLinkStrikePriceType_VARIANCE   PaymentStreamLinkStrikePriceType = "1"
)

// PaymentStreamNegativeRateTreatment field enumeration values.
type PaymentStreamNegativeRateTreatment string

const (
	PaymentStreamNegativeRateTreatment_ZERO_INTEREST_RATE_METHOD     PaymentStreamNegativeRateTreatment = "0"
	PaymentStreamNegativeRateTreatment_NEGATIVE_INTEREST_RATE_METHOD PaymentStreamNegativeRateTreatment = "1"
)

// PaymentStreamPaymentDateOffsetDayType field enumeration values.
type PaymentStreamPaymentDateOffsetDayType string

const (
	PaymentStreamPaymentDateOffsetDayType_BUSINESS              PaymentStreamPaymentDateOffsetDayType = "0"
	PaymentStreamPaymentDateOffsetDayType_CALENDAR              PaymentStreamPaymentDateOffsetDayType = "1"
	PaymentStreamPaymentDateOffsetDayType_COMMODITY_BUSINESS    PaymentStreamPaymentDateOffsetDayType = "2"
	PaymentStreamPaymentDateOffsetDayType_CURRENCY_BUSINESS     PaymentStreamPaymentDateOffsetDayType = "3"
	PaymentStreamPaymentDateOffsetDayType_EXCHANGE_BUSINESS     PaymentStreamPaymentDateOffsetDayType = "4"
	PaymentStreamPaymentDateOffsetDayType_SCHEDULED_TRADING_DAY PaymentStreamPaymentDateOffsetDayType = "5"
)

// PaymentStreamPaymentDateOffsetUnit field enumeration values.
type PaymentStreamPaymentDateOffsetUnit string

const (
	PaymentStreamPaymentDateOffsetUnit_DAY   PaymentStreamPaymentDateOffsetUnit = "D"
	PaymentStreamPaymentDateOffsetUnit_MONTH PaymentStreamPaymentDateOffsetUnit = "Mo"
	PaymentStreamPaymentDateOffsetUnit_WEEK  PaymentStreamPaymentDateOffsetUnit = "Wk"
	PaymentStreamPaymentDateOffsetUnit_YEAR  PaymentStreamPaymentDateOffsetUnit = "Yr"
)

// PaymentStreamPaymentFrequencyUnit field enumeration values.
type PaymentStreamPaymentFrequencyUnit string

const (
	PaymentStreamPaymentFrequencyUnit_DAY   PaymentStreamPaymentFrequencyUnit = "D"
	PaymentStreamPaymentFrequencyUnit_MONTH PaymentStreamPaymentFrequencyUnit = "Mo"
	PaymentStreamPaymentFrequencyUnit_TERM  PaymentStreamPaymentFrequencyUnit = "T"
	PaymentStreamPaymentFrequencyUnit_WEEK  PaymentStreamPaymentFrequencyUnit = "Wk"
	PaymentStreamPaymentFrequencyUnit_YEAR  PaymentStreamPaymentFrequencyUnit = "Yr"
)

// PaymentStreamPricingDayDistribution field enumeration values.
type PaymentStreamPricingDayDistribution string

const (
	PaymentStreamPricingDayDistribution_ALL         PaymentStreamPricingDayDistribution = "0"
	PaymentStreamPricingDayDistribution_FIRST       PaymentStreamPricingDayDistribution = "1"
	PaymentStreamPricingDayDistribution_LAST        PaymentStreamPricingDayDistribution = "2"
	PaymentStreamPricingDayDistribution_PENULTIMATE PaymentStreamPricingDayDistribution = "3"
)

// PaymentStreamPricingDayOfWeek field enumeration values.
type PaymentStreamPricingDayOfWeek string

const (
	PaymentStreamPricingDayOfWeek_EVERY_DAY PaymentStreamPricingDayOfWeek = "0"
	PaymentStreamPricingDayOfWeek_MONDAY    PaymentStreamPricingDayOfWeek = "1"
	PaymentStreamPricingDayOfWeek_TUESDAY   PaymentStreamPricingDayOfWeek = "2"
	PaymentStreamPricingDayOfWeek_WEDNESDAY PaymentStreamPricingDayOfWeek = "3"
	PaymentStreamPricingDayOfWeek_THURSDAY  PaymentStreamPricingDayOfWeek = "4"
	PaymentStreamPricingDayOfWeek_FRIDAY    PaymentStreamPricingDayOfWeek = "5"
	PaymentStreamPricingDayOfWeek_SATURDAY  PaymentStreamPricingDayOfWeek = "6"
	PaymentStreamPricingDayOfWeek_SUNDAY    PaymentStreamPricingDayOfWeek = "7"
)

// PaymentStreamRateIndexCurveUnit field enumeration values.
type PaymentStreamRateIndexCurveUnit string

const (
	PaymentStreamRateIndexCurveUnit_DAY   PaymentStreamRateIndexCurveUnit = "D"
	PaymentStreamRateIndexCurveUnit_MONTH PaymentStreamRateIndexCurveUnit = "Mo"
	PaymentStreamRateIndexCurveUnit_WEEK  PaymentStreamRateIndexCurveUnit = "Wk"
	PaymentStreamRateIndexCurveUnit_YEAR  PaymentStreamRateIndexCurveUnit = "Yr"
)

// PaymentStreamRateIndexSource field enumeration values.
type PaymentStreamRateIndexSource string

const (
	PaymentStreamRateIndexSource_BLOOMBERG PaymentStreamRateIndexSource = "0"
	PaymentStreamRateIndexSource_REUTERS   PaymentStreamRateIndexSource = "1"
	PaymentStreamRateIndexSource_TELERATE  PaymentStreamRateIndexSource = "2"
	PaymentStreamRateIndexSource_OTHER     PaymentStreamRateIndexSource = "99"
)

// PaymentStreamRateSpreadPositionType field enumeration values.
type PaymentStreamRateSpreadPositionType string

const (
	PaymentStreamRateSpreadPositionType_SHORT PaymentStreamRateSpreadPositionType = "0"
	PaymentStreamRateSpreadPositionType_LONG  PaymentStreamRateSpreadPositionType = "1"
)

// PaymentStreamRateSpreadType field enumeration values.
type PaymentStreamRateSpreadType string

const (
	PaymentStreamRateSpreadType_ABSOLUTE   PaymentStreamRateSpreadType = "0"
	PaymentStreamRateSpreadType_PERCENTAGE PaymentStreamRateSpreadType = "1"
)

// PaymentStreamRateTreatment field enumeration values.
type PaymentStreamRateTreatment string

const (
	PaymentStreamRateTreatment_BOND_EQUIVALENT_YIELD PaymentStreamRateTreatment = "0"
	PaymentStreamRateTreatment_MONEY_MARKET_YIELD    PaymentStreamRateTreatment = "1"
)

// PaymentStreamRealizedVarianceMethod field enumeration values.
type PaymentStreamRealizedVarianceMethod string

const (
	PaymentStreamRealizedVarianceMethod_PREVIOUS PaymentStreamRealizedVarianceMethod = "0"
	PaymentStreamRealizedVarianceMethod_LAST     PaymentStreamRealizedVarianceMethod = "1"
	PaymentStreamRealizedVarianceMethod_BOTH     PaymentStreamRealizedVarianceMethod = "2"
)

// PaymentStreamResetWeeklyRollConvention field enumeration values.
type PaymentStreamResetWeeklyRollConvention string

const (
	PaymentStreamResetWeeklyRollConvention_FRIDAY    PaymentStreamResetWeeklyRollConvention = "FRI"
	PaymentStreamResetWeeklyRollConvention_MONDAY    PaymentStreamResetWeeklyRollConvention = "MON"
	PaymentStreamResetWeeklyRollConvention_SATURDAY  PaymentStreamResetWeeklyRollConvention = "SAT"
	PaymentStreamResetWeeklyRollConvention_SUNDAY    PaymentStreamResetWeeklyRollConvention = "SUN"
	PaymentStreamResetWeeklyRollConvention_THURSDAY  PaymentStreamResetWeeklyRollConvention = "THU"
	PaymentStreamResetWeeklyRollConvention_TUESDAY   PaymentStreamResetWeeklyRollConvention = "TUE"
	PaymentStreamResetWeeklyRollConvention_WEDNESDAY PaymentStreamResetWeeklyRollConvention = "WED"
)

// PaymentStreamSettlLevel field enumeration values.
type PaymentStreamSettlLevel string

const (
	PaymentStreamSettlLevel_AVERAGE    PaymentStreamSettlLevel = "0"
	PaymentStreamSettlLevel_MAXIMUM    PaymentStreamSettlLevel = "1"
	PaymentStreamSettlLevel_MINIMUM    PaymentStreamSettlLevel = "2"
	PaymentStreamSettlLevel_CUMULATIVE PaymentStreamSettlLevel = "3"
)

// PaymentStreamType field enumeration values.
type PaymentStreamType string

const (
	PaymentStreamType_PERIODIC        PaymentStreamType = "0"
	PaymentStreamType_INITIAL         PaymentStreamType = "1"
	PaymentStreamType_SINGLE          PaymentStreamType = "2"
	PaymentStreamType_DIVIDEND        PaymentStreamType = "3"
	PaymentStreamType_INTEREST        PaymentStreamType = "4"
	PaymentStreamType_DIVIDEND_RETURN PaymentStreamType = "5"
	PaymentStreamType_PRICE_RETURN    PaymentStreamType = "6"
	PaymentStreamType_TOTAL_RETURN    PaymentStreamType = "7"
	PaymentStreamType_VARIANCE        PaymentStreamType = "8"
	PaymentStreamType_CORRELATION     PaymentStreamType = "9"
)

// PaymentStubLength field enumeration values.
type PaymentStubLength string

const (
	PaymentStubLength_SHORT PaymentStubLength = "0"
	PaymentStubLength_LONG  PaymentStubLength = "1"
)

// PaymentStubType field enumeration values.
type PaymentStubType string

const (
	PaymentStubType_INITIAL             PaymentStubType = "0"
	PaymentStubType_FINAL               PaymentStubType = "1"
	PaymentStubType_COMPOUNDING_INITIAL PaymentStubType = "2"
	PaymentStubType_COMPOUNDING_FINAL   PaymentStubType = "3"
)

// PaymentSubType field enumeration values.
type PaymentSubType string

const (
	PaymentSubType_INITIAL       PaymentSubType = "0"
	PaymentSubType_INTERMEDIATE  PaymentSubType = "1"
	PaymentSubType_FLOATING_RATE PaymentSubType = "10"
	PaymentSubType_FINAL         PaymentSubType = "2"
	PaymentSubType_PREPAID       PaymentSubType = "3"
	PaymentSubType_POSTPAID      PaymentSubType = "4"
	PaymentSubType_VARIABLE      PaymentSubType = "5"
	PaymentSubType_FIXED         PaymentSubType = "6"
	PaymentSubType_SWAP          PaymentSubType = "7"
	PaymentSubType_CONDITIONAL   PaymentSubType = "8"
	PaymentSubType_FIXED_RATE    PaymentSubType = "9"
)

// PaymentType field enumeration values.
type PaymentType string

const (
	PaymentType_BROKERAGE                   PaymentType = "0"
	PaymentType_UPFRONT_FEE                 PaymentType = "1"
	PaymentType_OPTION_PREMIUM              PaymentType = "10"
	PaymentType_SETTLEMENT_PAYMENT          PaymentType = "11"
	PaymentType_CASH_SETTLEMENT             PaymentType = "12"
	PaymentType_SECURITY_LENDING            PaymentType = "13"
	PaymentType_REBATE                      PaymentType = "14"
	PaymentType_INDEPENDENT_AMOUNT          PaymentType = "2"
	PaymentType_PRINCIPAL_EXCHANGE          PaymentType = "3"
	PaymentType_NOVATION                    PaymentType = "4"
	PaymentType_EARLY_TERMINATION_PROVISION PaymentType = "5"
	PaymentType_CANCELABLE_PROVISION        PaymentType = "6"
	PaymentType_EXTENDIBLE_PROVISION        PaymentType = "7"
	PaymentType_CAP_RATE_PROVISION          PaymentType = "8"
	PaymentType_FLOOR_RATE_PROVISION        PaymentType = "9"
	PaymentType_OTHER                       PaymentType = "99"
)

// PegLimitType field enumeration values.
type PegLimitType string

const (
	PegLimitType_OR_BETTER PegLimitType = "0"
	PegLimitType_STRICT    PegLimitType = "1"
	PegLimitType_OR_WORSE  PegLimitType = "2"
)

// PegMoveType field enumeration values.
type PegMoveType string

const (
	PegMoveType_FLOATING PegMoveType = "0"
	PegMoveType_FIXED    PegMoveType = "1"
)

// PegOffsetType field enumeration values.
type PegOffsetType string

const (
	PegOffsetType_PRICE        PegOffsetType = "0"
	PegOffsetType_BASIS_POINTS PegOffsetType = "1"
	PegOffsetType_TICKS        PegOffsetType = "2"
	PegOffsetType_PRICE_TIER   PegOffsetType = "3"
	PegOffsetType_PERCENTAGE   PegOffsetType = "4"
)

// PegPriceType field enumeration values.
type PegPriceType string

const (
	PegPriceType_LAST_PEG                                              PegPriceType = "1"
	PegPriceType_SHORT_SALE_MINIMUM_PRICE_PEG                          PegPriceType = "10"
	PegPriceType_MID_PRICE_PEG                                         PegPriceType = "2"
	PegPriceType_OPENING_PEG                                           PegPriceType = "3"
	PegPriceType_MARKET_PEG                                            PegPriceType = "4"
	PegPriceType_PRIMARY_PEG                                           PegPriceType = "5"
	PegPriceType_FIXED_PEG_TO_LOCAL_BEST_BID_OR_OFFER_AT_TIME_OF_ORDER PegPriceType = "6"
	PegPriceType_PEG_TO_VWAP                                           PegPriceType = "7"
	PegPriceType_TRAILING_STOP_PEG                                     PegPriceType = "8"
	PegPriceType_PEG_TO_LIMIT_PRICE                                    PegPriceType = "9"
)

// PegRoundDirection field enumeration values.
type PegRoundDirection string

const (
	PegRoundDirection_MORE_AGGRESSIVE PegRoundDirection = "1"
	PegRoundDirection_MORE_PASSIVE    PegRoundDirection = "2"
)

// PegScope field enumeration values.
type PegScope string

const (
	PegScope_LOCAL                    PegScope = "1"
	PegScope_NATIONAL                 PegScope = "2"
	PegScope_GLOBAL                   PegScope = "3"
	PegScope_NATIONAL_EXCLUDING_LOCAL PegScope = "4"
)

// PosAmtReason field enumeration values.
type PosAmtReason string

const (
	PosAmtReason_OPTIONS_SETTLEMENT         PosAmtReason = "0"
	PosAmtReason_PENDING_EROSION_ADJUSTMENT PosAmtReason = "1"
	PosAmtReason_FINAL_EROSION_ADJUSTMENT   PosAmtReason = "2"
	PosAmtReason_TEAR_UP_COUPON_AMOUNT      PosAmtReason = "3"
	PosAmtReason_PRICE_ALIGNMENT_INTEREST   PosAmtReason = "4"
	PosAmtReason_DELIVERY_INVOICE_CHARGES   PosAmtReason = "5"
	PosAmtReason_DELIVERY_STORAGE_CHARGES   PosAmtReason = "6"
)

// PosAmtType field enumeration values.
type PosAmtType string

const (
	PosAmtType_THE_FIVE_YEAR_EQUIVALENT_NOTIONAL_AMOUNT     PosAmtType = "5YREN"
	PosAmtType_ACCRUED_COUPON_AMOUNT                        PosAmtType = "ACPN"
	PosAmtType_TOTAL_BANKED_AMOUNT                          PosAmtType = "BANK"
	PosAmtType_CASH_AMOUNT                                  PosAmtType = "CASH"
	PosAmtType_COLLATERALIZED_MARK_TO_MARKET                PosAmtType = "CMTM"
	PosAmtType_TOTAL_COLLATERALIZED_AMOUNT                  PosAmtType = "COLAT"
	PosAmtType_COUPON_AMOUNT                                PosAmtType = "CPN"
	PosAmtType_CASH_RESIDUAL_AMOUNT                         PosAmtType = "CRES"
	PosAmtType_COMPENSATION_AMOUNT                          PosAmtType = "DLV"
	PosAmtType_END_VALUE                                    PosAmtType = "ENDV"
	PosAmtType_FINAL_MARK_TO_MARKET_AMOUNT                  PosAmtType = "FMTM"
	PosAmtType_INCREMENTAL_ACCRUED_COUPON                   PosAmtType = "IACPN"
	PosAmtType_INCREMENTAL_COLLATERALIZED_MARK_TO_MARKET    PosAmtType = "ICMTM"
	PosAmtType_INITIAL_TRADE_COUPON_AMOUNT                  PosAmtType = "ICPN"
	PosAmtType_INCREMENTAL_MARK_TO_MARKET                   PosAmtType = "IMTM"
	PosAmtType_LOAN_VALUE                                   PosAmtType = "LNVL"
	PosAmtType_LONG_PAIRED_SWAP_OR_SWAPTION_NOTIONAL_VALUE  PosAmtType = "LSNV"
	PosAmtType_OUTSTANDING_MARGIN_LOAN                      PosAmtType = "MGNLN"
	PosAmtType_MARK_TO_MODEL                                PosAmtType = "MTD"
	PosAmtType_NET_CASH_FLOW                                PosAmtType = "NCF"
	PosAmtType_NET_PRESENT_VALUE                            PosAmtType = "NPV"
	PosAmtType_PREMIUM_AMOUNT                               PosAmtType = "PREM"
	PosAmtType_PRESENT_VALUE_OF_ONE_BASIS_POINTS            PosAmtType = "PV01"
	PosAmtType_PRESENT_VALUE_OF_ALL_FEES                    PosAmtType = "PVFEES"
	PosAmtType_START_OF_DAY_ACCRUED_COUPON                  PosAmtType = "SACPN"
	PosAmtType_SETTLEMENT_VALUE                             PosAmtType = "SETL"
	PosAmtType_START_OF_DAY_MARK_TO_MARKET                  PosAmtType = "SMTM"
	PosAmtType_START_OF_DAY_NET_PRESENT_VALUE               PosAmtType = "SNPV"
	PosAmtType_SHORT_PAIRED_SWAP_OR_SWAPTION_NOTIONAL_VALUE PosAmtType = "SSNV"
	PosAmtType_TRADE_VARIATION_AMOUNT                       PosAmtType = "TVAR"
	PosAmtType_UNDISCOUNTED_MARK_TO_MARKET                  PosAmtType = "UMTM"
	PosAmtType_UPFRONT_PAYMENT                              PosAmtType = "UPFRNT"
	PosAmtType_VALUE_ADJUSTED_AMOUNT                        PosAmtType = "VADJ"
	PosAmtType_MARK_TO_MODEL_VARIANCE                       PosAmtType = "VMTD"
	PosAmtType_MARK_TO_MARKET_VARIANCE                      PosAmtType = "VMTM"
)

// PosMaintAction field enumeration values.
type PosMaintAction string

const (
	PosMaintAction_NEW     PosMaintAction = "1"
	PosMaintAction_REPLACE PosMaintAction = "2"
	PosMaintAction_CANCEL  PosMaintAction = "3"
	PosMaintAction_REVERSE PosMaintAction = "4"
)

// PosMaintResult field enumeration values.
type PosMaintResult string

const (
	PosMaintResult_SUCCESSFUL_COMPLETION PosMaintResult = "0"
	PosMaintResult_REJECTED              PosMaintResult = "1"
	PosMaintResult_OTHER                 PosMaintResult = "99"
)

// PosMaintStatus field enumeration values.
type PosMaintStatus string

const (
	PosMaintStatus_ACCEPTED                PosMaintStatus = "0"
	PosMaintStatus_ACCEPTED_WITH_WARNINGS  PosMaintStatus = "1"
	PosMaintStatus_REJECTED                PosMaintStatus = "2"
	PosMaintStatus_COMPLETED               PosMaintStatus = "3"
	PosMaintStatus_COMPLETED_WITH_WARNINGS PosMaintStatus = "4"
)

// PosQtyStatus field enumeration values.
type PosQtyStatus string

const (
	PosQtyStatus_SUBMITTED PosQtyStatus = "0"
	PosQtyStatus_ACCEPTED  PosQtyStatus = "1"
	PosQtyStatus_REJECTED  PosQtyStatus = "2"
)

// PosReqResult field enumeration values.
type PosReqResult string

const (
	PosReqResult_VALID_REQUEST                          PosReqResult = "0"
	PosReqResult_INVALID_OR_UNSUPPORTED_REQUEST         PosReqResult = "1"
	PosReqResult_NO_POSITIONS_FOUND_THAT_MATCH_CRITERIA PosReqResult = "2"
	PosReqResult_NOT_AUTHORIZED_TO_REQUEST_POSITIONS    PosReqResult = "3"
	PosReqResult_REQUEST_FOR_POSITION_NOT_SUPPORTED     PosReqResult = "4"
	PosReqResult_OTHER                                  PosReqResult = "99"
)

// PosReqStatus field enumeration values.
type PosReqStatus string

const (
	PosReqStatus_COMPLETED               PosReqStatus = "0"
	PosReqStatus_COMPLETED_WITH_WARNINGS PosReqStatus = "1"
	PosReqStatus_REJECTED                PosReqStatus = "2"
)

// PosReqType field enumeration values.
type PosReqType string

const (
	PosReqType_POSITIONS                              PosReqType = "0"
	PosReqType_TRADES                                 PosReqType = "1"
	PosReqType_POSITION_LIMIT_REPORTING_SUBMISSION    PosReqType = "10"
	PosReqType_EXERCISES                              PosReqType = "2"
	PosReqType_ASSIGNMENTS                            PosReqType = "3"
	PosReqType_SETTLEMENT_ACTIVITY                    PosReqType = "4"
	PosReqType_BACKOUT_MESSAGE                        PosReqType = "5"
	PosReqType_DELTA_POSITIONS                        PosReqType = "6"
	PosReqType_NET_POSITION                           PosReqType = "7"
	PosReqType_LARGE_POSITIONS_REPORTING              PosReqType = "8"
	PosReqType_EXERCISE_POSITION_REPORTING_SUBMISSION PosReqType = "9"
)

// PosTransType field enumeration values.
type PosTransType string

const (
	PosTransType_EXERCISE                             PosTransType = "1"
	PosTransType_TRANSFER_OF_FIRM                     PosTransType = "10"
	PosTransType_EXTERNAL_TRANSFER                    PosTransType = "11"
	PosTransType_CORPORATE_ACTION                     PosTransType = "12"
	PosTransType_NOTIFICATION                         PosTransType = "13"
	PosTransType_POSITION_CREATION                    PosTransType = "14"
	PosTransType_CLOSE_OUT                            PosTransType = "15"
	PosTransType_REOPEN                               PosTransType = "16"
	PosTransType_DO_NOT_EXERCISE                      PosTransType = "2"
	PosTransType_POSITION_ADJUSTMENT                  PosTransType = "3"
	PosTransType_POSITION_CHANGE_SUBMISSION           PosTransType = "4"
	PosTransType_PLEDGE                               PosTransType = "5"
	PosTransType_LARGE_TRADER_SUBMISSION              PosTransType = "6"
	PosTransType_LARGE_POSITIONS_REPORTING_SUBMISSION PosTransType = "7"
	PosTransType_LONG_HOLDINGS                        PosTransType = "8"
	PosTransType_INTERNAL_TRANSFER                    PosTransType = "9"
)

// PosType field enumeration values.
type PosType string

const (
	PosType_ALLOCATION_TRADE_QTY                       PosType = "ALC"
	PosType_OPTION_ASSIGNMENT                          PosType = "AS"
	PosType_AS_OF_TRADE_QTY                            PosType = "ASF"
	PosType_CORPORATE_ACTION_ADJUSTMENT                PosType = "CAA"
	PosType_CREDIT_EVENT_ADJUSTMENT                    PosType = "CEA"
	PosType_CASH_FUTURES_EQUIVALENT_QUANTITY           PosType = "CFE"
	PosType_DELTA_ADJUSTED_PAIRED_SWAPTION_POSITION    PosType = "DAS"
	PosType_NET_DELTA_QTY                              PosType = "DLT"
	PosType_DELIVERY_QTY                               PosType = "DLV"
	PosType_DELIVERY_NOTICE_QTY                        PosType = "DN"
	PosType_EXCHANGE_FOR_PHYSICAL_QTY                  PosType = "EP"
	PosType_ELECTRONIC_TRADE_QTY                       PosType = "ETR"
	PosType_OPTION_EXERCISE_QTY                        PosType = "EX"
	PosType_EXPIRING_QUANTITY                          PosType = "EXP"
	PosType_END_OF_DAY_QTY                             PosType = "FIN"
	PosType_GROSS_QTY                                  PosType = "GRS"
	PosType_INTRA_SPREAD_QTY                           PosType = "IAS"
	PosType_INTER_SPREAD_QTY                           PosType = "IES"
	PosType_INTRADAY_QTY                               PosType = "ITD"
	PosType_GROSS_NON_DELTA_ADJUSTED_SWAPTION_POSITION PosType = "NDAS"
	PosType_NET_QTY                                    PosType = "NET"
	PosType_ADJUSTMENT_QTY                             PosType = "PA"
	PosType_PIT_TRADE_QTY                              PosType = "PIT"
	PosType_PRIVATELY_NEGOTIATED_TRADE_QTY             PosType = "PNTN"
	PosType_RECEIVE_QUANTITY                           PosType = "RCV"
	PosType_REQUESTED_EXERCISE_QUANTITY                PosType = "REQ"
	PosType_SUCCESSION_EVENT_ADJUSTMENT                PosType = "SEA"
	PosType_LOAN_OR_BORROWED_QUANTITY                  PosType = "SECLN"
	PosType_START_OF_DAY_QTY                           PosType = "SOD"
	PosType_INTEGRAL_SPLIT                             PosType = "SPL"
	PosType_TRANSACTION_FROM_ASSIGNMENT                PosType = "TA"
	PosType_TOTAL_TRANSACTION_QTY                      PosType = "TOT"
	PosType_TRANSACTION_QUANTITY                       PosType = "TQ"
	PosType_TRANSFER_TRADE_QTY                         PosType = "TRF"
	PosType_TRANSACTION_FROM_EXERCISE                  PosType = "TX"
	PosType_QUANTITY_NOT_EXERCISED                     PosType = "UNEX"
	PosType_CROSS_MARGIN_QTY                           PosType = "XM"
)

// PositionCapacity field enumeration values.
type PositionCapacity string

const (
	PositionCapacity_PRINCIPAL    PositionCapacity = "0"
	PositionCapacity_AGENT        PositionCapacity = "1"
	PositionCapacity_CUSTOMER     PositionCapacity = "2"
	PositionCapacity_COUNTERPARTY PositionCapacity = "3"
)

// PositionEffect field enumeration values.
type PositionEffect string

const (
	PositionEffect_CLOSE                    PositionEffect = "C"
	PositionEffect_DEFAULT                  PositionEffect = "D"
	PositionEffect_FIFO                     PositionEffect = "F"
	PositionEffect_CLOSE_BUT_NOTIFY_ON_OPEN PositionEffect = "N"
	PositionEffect_OPEN                     PositionEffect = "O"
	PositionEffect_ROLLED                   PositionEffect = "R"
)

// PossDupFlag field enumeration values.
type PossDupFlag string

const (
	PossDupFlag_NO  PossDupFlag = "N"
	PossDupFlag_YES PossDupFlag = "Y"
)

// PossResend field enumeration values.
type PossResend string

const (
	PossResend_NO  PossResend = "N"
	PossResend_YES PossResend = "Y"
)

// PostTradePaymentDebitOrCredit field enumeration values.
type PostTradePaymentDebitOrCredit string

const (
	PostTradePaymentDebitOrCredit_DEBIT  PostTradePaymentDebitOrCredit = "0"
	PostTradePaymentDebitOrCredit_CREDIT PostTradePaymentDebitOrCredit = "1"
)

// PostTradePaymentStatus field enumeration values.
type PostTradePaymentStatus string

const (
	PostTradePaymentStatus_NEW       PostTradePaymentStatus = "0"
	PostTradePaymentStatus_INITIATED PostTradePaymentStatus = "1"
	PostTradePaymentStatus_PENDING   PostTradePaymentStatus = "2"
	PostTradePaymentStatus_CONFIRMED PostTradePaymentStatus = "3"
	PostTradePaymentStatus_REJECTED  PostTradePaymentStatus = "4"
)

// PreallocMethod field enumeration values.
type PreallocMethod string

const (
	PreallocMethod_PRO_RATA        PreallocMethod = "0"
	PreallocMethod_DO_NOT_PRO_RATA PreallocMethod = "1"
)

// PreviouslyReported field enumeration values.
type PreviouslyReported string

const (
	PreviouslyReported_NO  PreviouslyReported = "N"
	PreviouslyReported_YES PreviouslyReported = "Y"
)

// PriceLimitType field enumeration values.
type PriceLimitType string

const (
	PriceLimitType_PRICE      PriceLimitType = "0"
	PriceLimitType_TICKS      PriceLimitType = "1"
	PriceLimitType_PERCENTAGE PriceLimitType = "2"
)

// PriceMovementType field enumeration values.
type PriceMovementType string

const (
	PriceMovementType_AMOUNT     PriceMovementType = "0"
	PriceMovementType_PERCENTAGE PriceMovementType = "1"
)

// PriceProtectionScope field enumeration values.
type PriceProtectionScope string

const (
	PriceProtectionScope_NONE     PriceProtectionScope = "0"
	PriceProtectionScope_LOCAL    PriceProtectionScope = "1"
	PriceProtectionScope_NATIONAL PriceProtectionScope = "2"
	PriceProtectionScope_GLOBAL   PriceProtectionScope = "3"
)

// PriceQualifier field enumeration values.
type PriceQualifier string

const (
	PriceQualifier_ACCRUED_INTEREST                                                                             PriceQualifier = "0"
	PriceQualifier_TAX_IS_FACTORED_INTO_THE_PRICE                                                               PriceQualifier = "1"
	PriceQualifier_THE_EFFECT_OF_BOND_AMORTIZATION_OR_THE_FLOATING_RATE_INDEX_OFFSET_IS_FACTORED_INTO_THE_PRICE PriceQualifier = "2"
)

// PriceQuoteMethod field enumeration values.
type PriceQuoteMethod string

const (
	PriceQuoteMethod_INTEREST_RATE_INDEX                   PriceQuoteMethod = "INT"
	PriceQuoteMethod_INDEX                                 PriceQuoteMethod = "INX"
	PriceQuoteMethod_PERCENT_OF_PAR                        PriceQuoteMethod = "PCTPAR"
	PriceQuoteMethod_STANDARD_MONEY_PER_UNIT_OF_A_PHYSICAL PriceQuoteMethod = "STD"
)

// PriceType field enumeration values.
type PriceType string

const (
	PriceType_PERCENTAGE                          PriceType = "1"
	PriceType_FIXED_CABINET_TRADE_PRICE           PriceType = "10"
	PriceType_VARIABLE_CABINET_TRADE_PRICE        PriceType = "11"
	PriceType_PRICE_SPREAD                        PriceType = "12"
	PriceType_PRODUCT_TICKS_IN_HALVES             PriceType = "13"
	PriceType_PRODUCT_TICKS_IN_FOURTHS            PriceType = "14"
	PriceType_PRODUCT_TICKS_IN_EIGHTHS            PriceType = "15"
	PriceType_PRODUCT_TICKS_IN_SIXTEENTHS         PriceType = "16"
	PriceType_PRODUCT_TICKS_IN_THIRTY_SECONDS     PriceType = "17"
	PriceType_PRODUCT_TICKS_IN_SIXTY_FOURTHS      PriceType = "18"
	PriceType_PRODUCT_TICKS_IN_ONE_TWENTY_EIGHTHS PriceType = "19"
	PriceType_PER_UNIT                            PriceType = "2"
	PriceType_NORMAL_RATE_REPRESENTATION          PriceType = "20"
	PriceType_INVERSE_RATE_REPRESENTATION         PriceType = "21"
	PriceType_BASIS_POINTS                        PriceType = "22"
	PriceType_UP_FRONT_POINTS                     PriceType = "23"
	PriceType_INTEREST_RATE                       PriceType = "24"
	PriceType_PERCENTAGE_OF_NOTIONAL              PriceType = "25"
	PriceType_FIXED_AMOUNT                        PriceType = "3"
	PriceType_DISCOUNT                            PriceType = "4"
	PriceType_PREMIUM                             PriceType = "5"
	PriceType_SPREAD                              PriceType = "6"
	PriceType_TED_PRICE                           PriceType = "7"
	PriceType_TED_YIELD                           PriceType = "8"
	PriceType_YIELD                               PriceType = "9"
)

// PriorityIndicator field enumeration values.
type PriorityIndicator string

const (
	PriorityIndicator_PRIORITY_UNCHANGED                      PriorityIndicator = "0"
	PriorityIndicator_LOST_PRIORITY_AS_RESULT_OF_ORDER_CHANGE PriorityIndicator = "1"
)

// PrivateQuote field enumeration values.
type PrivateQuote string

const (
	PrivateQuote_NO  PrivateQuote = "N"
	PrivateQuote_YES PrivateQuote = "Y"
)

// ProcessCode field enumeration values.
type ProcessCode string

const (
	ProcessCode_REGULAR              ProcessCode = "0"
	ProcessCode_SOFT_DOLLAR          ProcessCode = "1"
	ProcessCode_STEP_IN              ProcessCode = "2"
	ProcessCode_STEP_OUT             ProcessCode = "3"
	ProcessCode_SOFT_DOLLAR_STEP_IN  ProcessCode = "4"
	ProcessCode_SOFT_DOLLAR_STEP_OUT ProcessCode = "5"
	ProcessCode_PLAN_SPONSOR         ProcessCode = "6"
)

// Product field enumeration values.
type Product string

const (
	Product_AGENCY      Product = "1"
	Product_MORTGAGE    Product = "10"
	Product_MUNICIPAL   Product = "11"
	Product_OTHER       Product = "12"
	Product_FINANCING   Product = "13"
	Product_COMMODITY   Product = "2"
	Product_CORPORATE   Product = "3"
	Product_CURRENCY    Product = "4"
	Product_EQUITY      Product = "5"
	Product_GOVERNMENT  Product = "6"
	Product_INDEX       Product = "7"
	Product_LOAN        Product = "8"
	Product_MONEYMARKET Product = "9"
)

// ProgRptReqs field enumeration values.
type ProgRptReqs string

const (
	ProgRptReqs_BUY_SIDE_EXPLICITLY_REQUESTS_STATUS_USING_STATUE_REQUEST                                            ProgRptReqs = "1"
	ProgRptReqs_SELL_SIDE_PERIODICALLY_SENDS_STATUS_USING_LIST_STATUS_PERIOD_OPTIONALLY_SPECIFIED_IN_PROGRESSPERIOD ProgRptReqs = "2"
	ProgRptReqs_REAL_TIME_EXECUTION_REPORTS                                                                         ProgRptReqs = "3"
)

// ProtectionTermEventDayType field enumeration values.
type ProtectionTermEventDayType string

const (
	ProtectionTermEventDayType_BUSINESS              ProtectionTermEventDayType = "0"
	ProtectionTermEventDayType_CALENDAR              ProtectionTermEventDayType = "1"
	ProtectionTermEventDayType_COMMODITY_BUSINESS    ProtectionTermEventDayType = "2"
	ProtectionTermEventDayType_CURRENCY_BUSINESS     ProtectionTermEventDayType = "3"
	ProtectionTermEventDayType_EXCHANGE_BUSINESS     ProtectionTermEventDayType = "4"
	ProtectionTermEventDayType_SCHEDULED_TRADING_DAY ProtectionTermEventDayType = "5"
)

// ProtectionTermEventQualifier field enumeration values.
type ProtectionTermEventQualifier string

const (
	ProtectionTermEventQualifier_FLOATING_RATE_INTEREST_SHORTFALL ProtectionTermEventQualifier = "C"
	ProtectionTermEventQualifier_RESTRUCTURING                    ProtectionTermEventQualifier = "E"
	ProtectionTermEventQualifier_RETRUCTURING                     ProtectionTermEventQualifier = "H"
)

// ProtectionTermEventUnit field enumeration values.
type ProtectionTermEventUnit string

const (
	ProtectionTermEventUnit_DAY   ProtectionTermEventUnit = "D"
	ProtectionTermEventUnit_MONTH ProtectionTermEventUnit = "Mo"
	ProtectionTermEventUnit_WEEK  ProtectionTermEventUnit = "Wk"
	ProtectionTermEventUnit_YEAR  ProtectionTermEventUnit = "Yr"
)

// ProvisionBreakFeeElection field enumeration values.
type ProvisionBreakFeeElection string

const (
	ProvisionBreakFeeElection_FLAT_FEE                      ProvisionBreakFeeElection = "0"
	ProvisionBreakFeeElection_AMORTIZED_FEE                 ProvisionBreakFeeElection = "1"
	ProvisionBreakFeeElection_FUNDING_FEE                   ProvisionBreakFeeElection = "2"
	ProvisionBreakFeeElection_FLAT_FEE_AND_FUNDING_FEE      ProvisionBreakFeeElection = "3"
	ProvisionBreakFeeElection_AMORTIZED_FEE_AND_FUNDING_FEE ProvisionBreakFeeElection = "4"
)

// ProvisionCalculationAgent field enumeration values.
type ProvisionCalculationAgent string

const (
	ProvisionCalculationAgent_EXERCISING_PARTY                              ProvisionCalculationAgent = "0"
	ProvisionCalculationAgent_NON_EXERCISING_PARTY                          ProvisionCalculationAgent = "1"
	ProvisionCalculationAgent_AS_SPECIFIED_IN_THE_MASTER_AGREEMENT          ProvisionCalculationAgent = "2"
	ProvisionCalculationAgent_AS_SPECIFIED_IN_THE_STANDARD_TERMS_SUPPLEMENT ProvisionCalculationAgent = "3"
)

// ProvisionCashSettlMethod field enumeration values.
type ProvisionCashSettlMethod string

const (
	ProvisionCashSettlMethod_CASH_PRICE                       ProvisionCashSettlMethod = "0"
	ProvisionCashSettlMethod_CASH_PRICE_ALTERNATE             ProvisionCashSettlMethod = "1"
	ProvisionCashSettlMethod_PAR_YIELD_CURVE_ADJUSTED         ProvisionCashSettlMethod = "2"
	ProvisionCashSettlMethod_ZERO_COUPON_YIELD_CURVE_ADJUSTED ProvisionCashSettlMethod = "3"
	ProvisionCashSettlMethod_PAR_YIELD_CURVE_UNADJUSTED       ProvisionCashSettlMethod = "4"
	ProvisionCashSettlMethod_CROSS_CURRENCY                   ProvisionCashSettlMethod = "5"
	ProvisionCashSettlMethod_COLLATERALIZED_PRICE             ProvisionCashSettlMethod = "6"
)

// ProvisionCashSettlPaymentDateType field enumeration values.
type ProvisionCashSettlPaymentDateType string

const (
	ProvisionCashSettlPaymentDateType_UNADJUSTED ProvisionCashSettlPaymentDateType = "0"
	ProvisionCashSettlPaymentDateType_ADJUSTED   ProvisionCashSettlPaymentDateType = "1"
)

// ProvisionCashSettlQuoteType field enumeration values.
type ProvisionCashSettlQuoteType string

const (
	ProvisionCashSettlQuoteType_BID                   ProvisionCashSettlQuoteType = "0"
	ProvisionCashSettlQuoteType_MID                   ProvisionCashSettlQuoteType = "1"
	ProvisionCashSettlQuoteType_OFFER                 ProvisionCashSettlQuoteType = "2"
	ProvisionCashSettlQuoteType_EXERCISING_PARTY_PAYS ProvisionCashSettlQuoteType = "3"
)

// ProvisionDateTenorUnit field enumeration values.
type ProvisionDateTenorUnit string

const (
	ProvisionDateTenorUnit_DAY   ProvisionDateTenorUnit = "D"
	ProvisionDateTenorUnit_MONTH ProvisionDateTenorUnit = "Mo"
	ProvisionDateTenorUnit_WEEK  ProvisionDateTenorUnit = "Wk"
	ProvisionDateTenorUnit_YEAR  ProvisionDateTenorUnit = "Yr"
)

// ProvisionOptionExerciseEarliestDateOffsetUnit field enumeration values.
type ProvisionOptionExerciseEarliestDateOffsetUnit string

const (
	ProvisionOptionExerciseEarliestDateOffsetUnit_DAY   ProvisionOptionExerciseEarliestDateOffsetUnit = "D"
	ProvisionOptionExerciseEarliestDateOffsetUnit_MONTH ProvisionOptionExerciseEarliestDateOffsetUnit = "Mo"
	ProvisionOptionExerciseEarliestDateOffsetUnit_WEEK  ProvisionOptionExerciseEarliestDateOffsetUnit = "Wk"
	ProvisionOptionExerciseEarliestDateOffsetUnit_YEAR  ProvisionOptionExerciseEarliestDateOffsetUnit = "Yr"
)

// ProvisionOptionExerciseFixedDateType field enumeration values.
type ProvisionOptionExerciseFixedDateType string

const (
	ProvisionOptionExerciseFixedDateType_UNADJUSTED ProvisionOptionExerciseFixedDateType = "0"
	ProvisionOptionExerciseFixedDateType_ADJUSTED   ProvisionOptionExerciseFixedDateType = "1"
)

// ProvisionOptionSinglePartyBuyerSide field enumeration values.
type ProvisionOptionSinglePartyBuyerSide string

const (
	ProvisionOptionSinglePartyBuyerSide_BUY  ProvisionOptionSinglePartyBuyerSide = "1"
	ProvisionOptionSinglePartyBuyerSide_SELL ProvisionOptionSinglePartyBuyerSide = "2"
)

// ProvisionType field enumeration values.
type ProvisionType string

const (
	ProvisionType_MANDATORY_EARLY_TERMINATION ProvisionType = "0"
	ProvisionType_OPTIONAL_EARLY_TERMINATION  ProvisionType = "1"
	ProvisionType_CANCELABLE                  ProvisionType = "2"
	ProvisionType_EXTENDABLE                  ProvisionType = "3"
	ProvisionType_MUTUAL_EARLY_TERMINATION    ProvisionType = "4"
	ProvisionType_EVERGREEN                   ProvisionType = "5"
	ProvisionType_CALLABLE                    ProvisionType = "6"
	ProvisionType_PUTTABLE                    ProvisionType = "7"
)

// PublishTrdIndicator field enumeration values.
type PublishTrdIndicator string

const (
	PublishTrdIndicator_NO  PublishTrdIndicator = "N"
	PublishTrdIndicator_YES PublishTrdIndicator = "Y"
)

// PutOrCall field enumeration values.
type PutOrCall string

const (
	PutOrCall_PUT     PutOrCall = "0"
	PutOrCall_CALL    PutOrCall = "1"
	PutOrCall_OTHER   PutOrCall = "2"
	PutOrCall_CHOOSER PutOrCall = "3"
)

// QtyType field enumeration values.
type QtyType string

const (
	QtyType_UNITS                         QtyType = "0"
	QtyType_CONTRACTS                     QtyType = "1"
	QtyType_UNIT_OF_MEASURE_PER_TIME_UNIT QtyType = "2"
)

// QuantityType field enumeration values.
type QuantityType string

const (
	QuantityType_SHARES       QuantityType = "1"
	QuantityType_BONDS        QuantityType = "2"
	QuantityType_CURRENTFACE  QuantityType = "3"
	QuantityType_ORIGINALFACE QuantityType = "4"
	QuantityType_CURRENCY     QuantityType = "5"
	QuantityType_CONTRACTS    QuantityType = "6"
	QuantityType_OTHER        QuantityType = "7"
	QuantityType_PAR          QuantityType = "8"
)

// QuoteAckStatus field enumeration values.
type QuoteAckStatus string

const (
	QuoteAckStatus_RECEIVED_NOT_YET_PROCESSED QuoteAckStatus = "0"
	QuoteAckStatus_ACCEPTED                   QuoteAckStatus = "1"
	QuoteAckStatus_REJECTED                   QuoteAckStatus = "2"
	QuoteAckStatus_CANCELED_FOR_UNDERLYING    QuoteAckStatus = "3"
	QuoteAckStatus_CANCELED_ALL               QuoteAckStatus = "4"
)

// QuoteAttributeType field enumeration values.
type QuoteAttributeType string

const (
	QuoteAttributeType_QUOTE_IS_ABOVE_STANDARD_MARKET_SIZE               QuoteAttributeType = "0"
	QuoteAttributeType_QUOTE_IS_ABOVE_SIZE_SPECIFIC_TO_THE_INSTRUMENT    QuoteAttributeType = "1"
	QuoteAttributeType_QUOTE_APPLICABLE_FOR_LIQUIDITY_PROVISION_ACTIVITY QuoteAttributeType = "2"
	QuoteAttributeType_QUOTE_ISSUER_STATUS                               QuoteAttributeType = "3"
	QuoteAttributeType_BID_OR_ASK_REQUEST                                QuoteAttributeType = "4"
)

// QuoteCancelType field enumeration values.
type QuoteCancelType string

const (
	QuoteCancelType_CANCEL_FOR_ONE_OR_MORE_SECURITIES        QuoteCancelType = "1"
	QuoteCancelType_CANCEL_FOR_SECURITY_TYPE                 QuoteCancelType = "2"
	QuoteCancelType_CANCEL_FOR_UNDERLYING_SECURITY           QuoteCancelType = "3"
	QuoteCancelType_CANCEL_ALL_QUOTES                        QuoteCancelType = "4"
	QuoteCancelType_CANCEL_SPECIFIED_SINGLE_QUOTE            QuoteCancelType = "5"
	QuoteCancelType_CANCEL_BY_TYPE_OF_QUOTE                  QuoteCancelType = "6"
	QuoteCancelType_CANCEL_FOR_SECURITY_ISSUER               QuoteCancelType = "7"
	QuoteCancelType_CANCEL_FOR_ISSUER_OF_UNDERLYING_SECURITY QuoteCancelType = "8"
)

// QuoteCondition field enumeration values.
type QuoteCondition string

const (
	QuoteCondition_RESERVED_SAM                        QuoteCondition = "0"
	QuoteCondition_NO_ACTIVE_SAM                       QuoteCondition = "1"
	QuoteCondition_RESTRICTED                          QuoteCondition = "2"
	QuoteCondition_REST_OF_BOOK_VWAP                   QuoteCondition = "3"
	QuoteCondition_BETTER_PRICES_IN_CONDITIONAL_ORDERS QuoteCondition = "4"
	QuoteCondition_MEDIAN_PRICE                        QuoteCondition = "5"
	QuoteCondition_FULL_CURVE                          QuoteCondition = "6"
	QuoteCondition_FLAT_CURVE                          QuoteCondition = "7"
	QuoteCondition_OPEN_ACTIVE                         QuoteCondition = "A"
	QuoteCondition_CLOSED_INACTIVE                     QuoteCondition = "B"
	QuoteCondition_EXCHANGE_BEST                       QuoteCondition = "C"
	QuoteCondition_CONSOLIDATED_BEST                   QuoteCondition = "D"
	QuoteCondition_LOCKED                              QuoteCondition = "E"
	QuoteCondition_CROSSED                             QuoteCondition = "F"
	QuoteCondition_DEPTH                               QuoteCondition = "G"
	QuoteCondition_FAST_TRADING                        QuoteCondition = "H"
	QuoteCondition_NON_FIRM                            QuoteCondition = "I"
	QuoteCondition_OUTRIGHT_PRICE                      QuoteCondition = "J"
	QuoteCondition_IMPLIED_PRICE                       QuoteCondition = "K"
	QuoteCondition_MANUAL_SLOW_QUOTE                   QuoteCondition = "L"
	QuoteCondition_DEPTH_ON_OFFER                      QuoteCondition = "M"
	QuoteCondition_DEPTH_ON_BID                        QuoteCondition = "N"
	QuoteCondition_CLOSING                             QuoteCondition = "O"
	QuoteCondition_NEWS_DISSEMINATION                  QuoteCondition = "P"
	QuoteCondition_TRADING_RANGE                       QuoteCondition = "Q"
	QuoteCondition_ORDER_INFLUX                        QuoteCondition = "R"
	QuoteCondition_DUE_TO_RELATED                      QuoteCondition = "S"
	QuoteCondition_NEWS_PENDING                        QuoteCondition = "T"
	QuoteCondition_ADDITIONAL_INFO                     QuoteCondition = "U"
	QuoteCondition_ADDITIONAL_INFO_DUE_TO_RELATED      QuoteCondition = "V"
	QuoteCondition_RESUME                              QuoteCondition = "W"
	QuoteCondition_VIEW_OF_COMMON                      QuoteCondition = "X"
	QuoteCondition_VOLUME_ALERT                        QuoteCondition = "Y"
	QuoteCondition_ORDER_IMBALANCE                     QuoteCondition = "Z"
	QuoteCondition_EQUIPMENT_CHANGEOVER                QuoteCondition = "a"
	QuoteCondition_NO_OPEN                             QuoteCondition = "b"
	QuoteCondition_REGULAR_ETH                         QuoteCondition = "c"
	QuoteCondition_AUTOMATIC_EXECUTION                 QuoteCondition = "d"
	QuoteCondition_AUTOMATIC_EXECUTION_ETH             QuoteCondition = "e"
	QuoteCondition_FAST_MARKET_ETH                     QuoteCondition = "f"
	QuoteCondition_INACTIVE_ETH                        QuoteCondition = "g"
	QuoteCondition_ROTATION                            QuoteCondition = "h"
	QuoteCondition_ROTATION_ETH                        QuoteCondition = "i"
	QuoteCondition_HALT                                QuoteCondition = "j"
	QuoteCondition_HALT_ETH                            QuoteCondition = "k"
	QuoteCondition_DUE_TO_NEWS_DISSEMINATION           QuoteCondition = "l"
	QuoteCondition_DUE_TO_NEWS_PENDING                 QuoteCondition = "m"
	QuoteCondition_TRADING_RESUME                      QuoteCondition = "n"
	QuoteCondition_OUT_OF_SEQUENCE                     QuoteCondition = "o"
	QuoteCondition_BID_SPECIALIST                      QuoteCondition = "p"
	QuoteCondition_OFFER_SPECIALIST                    QuoteCondition = "q"
	QuoteCondition_BID_OFFER_SPECIALIST                QuoteCondition = "r"
	QuoteCondition_END_OF_DAY_SAM                      QuoteCondition = "s"
	QuoteCondition_FORBIDDEN_SAM                       QuoteCondition = "t"
	QuoteCondition_FROZEN_SAM                          QuoteCondition = "u"
	QuoteCondition_PREOPENING_SAM                      QuoteCondition = "v"
	QuoteCondition_OPENING_SAM                         QuoteCondition = "w"
	QuoteCondition_OPEN_SAM                            QuoteCondition = "x"
	QuoteCondition_SURVEILLANCE_SAM                    QuoteCondition = "y"
	QuoteCondition_SUSPENDED_SAM                       QuoteCondition = "z"
)

// QuoteEntryRejectReason field enumeration values.
type QuoteEntryRejectReason string

const (
	QuoteEntryRejectReason_UNKNOWN_SYMBOL                   QuoteEntryRejectReason = "1"
	QuoteEntryRejectReason_EXHCNAGE                         QuoteEntryRejectReason = "2"
	QuoteEntryRejectReason_QUOTE_EXCEEDS_LIMIT              QuoteEntryRejectReason = "3"
	QuoteEntryRejectReason_TOO_LATE_TO_ENTER                QuoteEntryRejectReason = "4"
	QuoteEntryRejectReason_UNKNOWN_QUOTE                    QuoteEntryRejectReason = "5"
	QuoteEntryRejectReason_DUPLICATE_QUOTE                  QuoteEntryRejectReason = "6"
	QuoteEntryRejectReason_INVALID_BID_ASK_SPREAD           QuoteEntryRejectReason = "7"
	QuoteEntryRejectReason_INVALID_PRICE                    QuoteEntryRejectReason = "8"
	QuoteEntryRejectReason_NOT_AUTHORIZED_TO_QUOTE_SECURITY QuoteEntryRejectReason = "9"
	QuoteEntryRejectReason_OTHER                            QuoteEntryRejectReason = "99"
)

// QuoteEntryStatus field enumeration values.
type QuoteEntryStatus string

const (
	QuoteEntryStatus_ACCEPTED                     QuoteEntryStatus = "0"
	QuoteEntryStatus_LOCKED_MARKET_WARNING        QuoteEntryStatus = "12"
	QuoteEntryStatus_CROSS_MARKET_WARNING         QuoteEntryStatus = "13"
	QuoteEntryStatus_CANCELED_DUE_TO_LOCK_MARKET  QuoteEntryStatus = "14"
	QuoteEntryStatus_CANCELED_DUE_TO_CROSS_MARKET QuoteEntryStatus = "15"
	QuoteEntryStatus_ACTIVE                       QuoteEntryStatus = "16"
	QuoteEntryStatus_REJECTED                     QuoteEntryStatus = "5"
	QuoteEntryStatus_REMOVED_FROM_MARKET          QuoteEntryStatus = "6"
	QuoteEntryStatus_EXPIRED                      QuoteEntryStatus = "7"
)

// QuoteModelType field enumeration values.
type QuoteModelType string

const (
	QuoteModelType_QUOTE_ENTRY        QuoteModelType = "1"
	QuoteModelType_QUOTE_MODIFICATION QuoteModelType = "2"
)

// QuotePriceType field enumeration values.
type QuotePriceType string

const (
	QuotePriceType_PERCENTAGE                          QuotePriceType = "1"
	QuotePriceType_YIELD                               QuotePriceType = "10"
	QuotePriceType_PRICE_SPREAD                        QuotePriceType = "12"
	QuotePriceType_PRODUCT_TICKS_IN_HALVES             QuotePriceType = "13"
	QuotePriceType_PRODUCT_TICKS_IN_FOURTHS            QuotePriceType = "14"
	QuotePriceType_PRODUCT_TICKS_IN_EIGHTHS            QuotePriceType = "15"
	QuotePriceType_PRODUCT_TICKS_IN_SIXTEENTHS         QuotePriceType = "16"
	QuotePriceType_PRODUCT_TICKS_IN_THIRTY_SECONDS     QuotePriceType = "17"
	QuotePriceType_PRODUCT_TICKS_IN_SIXTY_FOURTHS      QuotePriceType = "18"
	QuotePriceType_PRODUCT_TICKS_IN_ONE_TWENTY_EIGHTHS QuotePriceType = "19"
	QuotePriceType_PER_UNIT                            QuotePriceType = "2"
	QuotePriceType_NORMAL_RATE_REPRESENTATION          QuotePriceType = "20"
	QuotePriceType_INVERSE_RATE_REPRESENTATION         QuotePriceType = "21"
	QuotePriceType_BASIS_POINTS                        QuotePriceType = "22"
	QuotePriceType_UP_FRONT_POINTS                     QuotePriceType = "23"
	QuotePriceType_INTEREST_RATE                       QuotePriceType = "24"
	QuotePriceType_PERCENTAGE_OF_NOTIONAL              QuotePriceType = "25"
	QuotePriceType_FIXED_AMOUNT                        QuotePriceType = "3"
	QuotePriceType_DISCOUNT                            QuotePriceType = "4"
	QuotePriceType_PREMIUM                             QuotePriceType = "5"
	QuotePriceType_SPREAD                              QuotePriceType = "6"
	QuotePriceType_TED_PRICE                           QuotePriceType = "7"
	QuotePriceType_TED_YIELD                           QuotePriceType = "8"
	QuotePriceType_YIELD_SPREAD                        QuotePriceType = "9"
)

// QuoteRejectReason field enumeration values.
type QuoteRejectReason string

const (
	QuoteRejectReason_UNKNOWN_SYMBOL                                   QuoteRejectReason = "1"
	QuoteRejectReason_PRICE_EXCEEDS_CURRENT_PRICE_BAND_10              QuoteRejectReason = "10"
	QuoteRejectReason_QUOTE_LOCKED                                     QuoteRejectReason = "11"
	QuoteRejectReason_INVALID_OR_UNKNOWN_SECURITY_ISSUER               QuoteRejectReason = "12"
	QuoteRejectReason_INVALID_OR_UNKNOWN_ISSUER_OF_UNDERLYING_SECURITY QuoteRejectReason = "13"
	QuoteRejectReason_NOTIONAL_VALUE_EXCEEDS_THRESHOLD                 QuoteRejectReason = "14"
	QuoteRejectReason_PRICE_EXCEEDS_CURRENT_PRICE_BAND_15              QuoteRejectReason = "15"
	QuoteRejectReason_REFERENCE_PRICE_NOT_AVAILABLE                    QuoteRejectReason = "16"
	QuoteRejectReason_INSUFFICIENT_CREDIT_LIMIT                        QuoteRejectReason = "17"
	QuoteRejectReason_EXCEEDED_CLIP_SIZE_LIMIT                         QuoteRejectReason = "18"
	QuoteRejectReason_EXCEEDED_MAXIMUM_NOTIONAL_ORDER_AMOUNT           QuoteRejectReason = "19"
	QuoteRejectReason_EXCHANGE                                         QuoteRejectReason = "2"
	QuoteRejectReason_EXCEEDED_DV01_PV01_LIMIT                         QuoteRejectReason = "20"
	QuoteRejectReason_EXCEEDED_CS01_LIMIT                              QuoteRejectReason = "21"
	QuoteRejectReason_QUOTE_REQUEST_EXCEEDS_LIMIT                      QuoteRejectReason = "3"
	QuoteRejectReason_TOO_LATE_TO_ENTER                                QuoteRejectReason = "4"
	QuoteRejectReason_UNKNOWN_QUOTE                                    QuoteRejectReason = "5"
	QuoteRejectReason_DUPLICATE_QUOTE                                  QuoteRejectReason = "6"
	QuoteRejectReason_INVALID_BID_ASK_SPREAD                           QuoteRejectReason = "7"
	QuoteRejectReason_INVALID_PRICE                                    QuoteRejectReason = "8"
	QuoteRejectReason_NOT_AUTHORIZED_TO_QUOTE_SECURITY                 QuoteRejectReason = "9"
	QuoteRejectReason_OTHER                                            QuoteRejectReason = "99"
)

// QuoteRequestRejectReason field enumeration values.
type QuoteRequestRejectReason string

const (
	QuoteRequestRejectReason_UNKNOWN_SYMBOL                         QuoteRequestRejectReason = "1"
	QuoteRequestRejectReason_PASS                                   QuoteRequestRejectReason = "10"
	QuoteRequestRejectReason_INSUFFICIENT_CREDIT                    QuoteRequestRejectReason = "11"
	QuoteRequestRejectReason_EXCEEDED_CLIP_SIZE_LIMIT               QuoteRequestRejectReason = "12"
	QuoteRequestRejectReason_EXCEEDED_MAXIMUM_NOTIONAL_ORDER_AMOUNT QuoteRequestRejectReason = "13"
	QuoteRequestRejectReason_EXCEEDED_DV01_PV01_LIMIT               QuoteRequestRejectReason = "14"
	QuoteRequestRejectReason_EXCEEDED_CS01_LIMIT                    QuoteRequestRejectReason = "15"
	QuoteRequestRejectReason_EXCHANGE                               QuoteRequestRejectReason = "2"
	QuoteRequestRejectReason_QUOTE_REQUEST_EXCEEDS_LIMIT            QuoteRequestRejectReason = "3"
	QuoteRequestRejectReason_TOO_LATE_TO_ENTER                      QuoteRequestRejectReason = "4"
	QuoteRequestRejectReason_INVALID_PRICE                          QuoteRequestRejectReason = "5"
	QuoteRequestRejectReason_NOT_AUTHORIZED_TO_REQUEST_QUOTE        QuoteRequestRejectReason = "6"
	QuoteRequestRejectReason_NO_MATCH_FOR_INQUIRY                   QuoteRequestRejectReason = "7"
	QuoteRequestRejectReason_NO_MARKET_FOR_INSTRUMENT               QuoteRequestRejectReason = "8"
	QuoteRequestRejectReason_NO_INVENTORY                           QuoteRequestRejectReason = "9"
	QuoteRequestRejectReason_OTHER                                  QuoteRequestRejectReason = "99"
)

// QuoteRequestType field enumeration values.
type QuoteRequestType string

const (
	QuoteRequestType_MANUAL        QuoteRequestType = "1"
	QuoteRequestType_AUTOMATIC     QuoteRequestType = "2"
	QuoteRequestType_CONFIRM_QUOTE QuoteRequestType = "3"
)

// QuoteRespType field enumeration values.
type QuoteRespType string

const (
	QuoteRespType_HIT_LIFT           QuoteRespType = "1"
	QuoteRespType_TIED_COVER         QuoteRespType = "10"
	QuoteRespType_ACCEPT             QuoteRespType = "11"
	QuoteRespType_TERMINATE_CONTRACT QuoteRespType = "12"
	QuoteRespType_COUNTER            QuoteRespType = "2"
	QuoteRespType_EXPIRED            QuoteRespType = "3"
	QuoteRespType_COVER              QuoteRespType = "4"
	QuoteRespType_DONE_AWAY          QuoteRespType = "5"
	QuoteRespType_PASS               QuoteRespType = "6"
	QuoteRespType_END_TRADE          QuoteRespType = "7"
	QuoteRespType_TIMED_OUT          QuoteRespType = "8"
	QuoteRespType_TIED               QuoteRespType = "9"
)

// QuoteResponseLevel field enumeration values.
type QuoteResponseLevel string

const (
	QuoteResponseLevel_NO_ACKNOWLEDGEMENT                            QuoteResponseLevel = "0"
	QuoteResponseLevel_ACKNOWLEDGE_ONLY_NEGATIVE_OR_ERRONEOUS_QUOTES QuoteResponseLevel = "1"
	QuoteResponseLevel_ACKNOWLEDGE_EACH_QUOTE_MESSAGE                QuoteResponseLevel = "2"
	QuoteResponseLevel_SUMMARY_ACKNOWLEDGEMENT                       QuoteResponseLevel = "3"
)

// QuoteSideIndicator field enumeration values.
type QuoteSideIndicator string

const (
	QuoteSideIndicator_NO  QuoteSideIndicator = "N"
	QuoteSideIndicator_YES QuoteSideIndicator = "Y"
)

// QuoteStatus field enumeration values.
type QuoteStatus string

const (
	QuoteStatus_ACCEPTED                            QuoteStatus = "0"
	QuoteStatus_CANCELED_FOR_SPECIFIC_SECURITIES    QuoteStatus = "1"
	QuoteStatus_PENDING                             QuoteStatus = "10"
	QuoteStatus_PASS                                QuoteStatus = "11"
	QuoteStatus_LOCKED_MARKET_WARNING               QuoteStatus = "12"
	QuoteStatus_CROSSED_MARKET_WARNING              QuoteStatus = "13"
	QuoteStatus_CANCELED_DUE_TO_LOCKED_MARKET       QuoteStatus = "14"
	QuoteStatus_CANCELED_DUE_TO_CROSSED_MARKET      QuoteStatus = "15"
	QuoteStatus_ACTIVE                              QuoteStatus = "16"
	QuoteStatus_CANCELED                            QuoteStatus = "17"
	QuoteStatus_UNSOLICITED_QUOTE_REPLENISHMENT     QuoteStatus = "18"
	QuoteStatus_PENDING_END_TRADE                   QuoteStatus = "19"
	QuoteStatus_CANCELED_FOR_SPECIFIC_SECURITYTYPES QuoteStatus = "2"
	QuoteStatus_TOO_LATE_TO_END                     QuoteStatus = "20"
	QuoteStatus_TRADED                              QuoteStatus = "21"
	QuoteStatus_TRADED_AND_REMOVED                  QuoteStatus = "22"
	QuoteStatus_CONTRACT_TERMINATED                 QuoteStatus = "23"
	QuoteStatus_CANCELED_FOR_UNDERLYING             QuoteStatus = "3"
	QuoteStatus_CANCELED_ALL                        QuoteStatus = "4"
	QuoteStatus_REJECTED                            QuoteStatus = "5"
	QuoteStatus_REMOVED_FROM_MARKET                 QuoteStatus = "6"
	QuoteStatus_EXPIRED                             QuoteStatus = "7"
	QuoteStatus_QUERY                               QuoteStatus = "8"
	QuoteStatus_QUOTE_NOT_FOUND                     QuoteStatus = "9"
)

// QuoteType field enumeration values.
type QuoteType string

const (
	QuoteType_INDICATIVE           QuoteType = "0"
	QuoteType_TRADEABLE            QuoteType = "1"
	QuoteType_RESTRICTED_TRADEABLE QuoteType = "2"
	QuoteType_COUNTER              QuoteType = "3"
	QuoteType_INITIALLY_TRADEABLE  QuoteType = "4"
)

// RateSource field enumeration values.
type RateSource string

const (
	RateSource_BLOOMBERG                   RateSource = "0"
	RateSource_REUTERS                     RateSource = "1"
	RateSource_TELERATE                    RateSource = "2"
	RateSource_ISDA_SETTLEMENT_RATE_OPTION RateSource = "3"
	RateSource_OTHER                       RateSource = "99"
)

// RateSourceType field enumeration values.
type RateSourceType string

const (
	RateSourceType_PRIMARY   RateSourceType = "0"
	RateSourceType_SECONDARY RateSourceType = "1"
)

// RefOrdIDReason field enumeration values.
type RefOrdIDReason string

const (
	RefOrdIDReason_GTC_FROM_PREVIOUS_DAY  RefOrdIDReason = "0"
	RefOrdIDReason_PARTIAL_FILL_REMAINING RefOrdIDReason = "1"
	RefOrdIDReason_ORDER_CHANGED          RefOrdIDReason = "2"
)

// RefOrderIDSource field enumeration values.
type RefOrderIDSource string

const (
	RefOrderIDSource_SECONDARY_ORDER_ID        RefOrderIDSource = "0"
	RefOrderIDSource_ORDER_ID                  RefOrderIDSource = "1"
	RefOrderIDSource_MARKET_DATA_ENTRY_ID      RefOrderIDSource = "2"
	RefOrderIDSource_QUOTE_ENTRY_ID            RefOrderIDSource = "3"
	RefOrderIDSource_ORIGINAL_ORDER_ID         RefOrderIDSource = "4"
	RefOrderIDSource_QUOTE_ID                  RefOrderIDSource = "5"
	RefOrderIDSource_QUOTE_REQUEST_ID          RefOrderIDSource = "6"
	RefOrderIDSource_PREVIOUS_ORDER_IDENTIFIER RefOrderIDSource = "7"
	RefOrderIDSource_PREVIOUS_QUOTE_IDENTIFIER RefOrderIDSource = "8"
	RefOrderIDSource_PARENT_ORDER_IDENTIFIER   RefOrderIDSource = "9"
	RefOrderIDSource_MANUAL_ORDER_IDENTIFIER   RefOrderIDSource = "A"
)

// RefRiskLimitCheckIDType field enumeration values.
type RefRiskLimitCheckIDType string

const (
	RefRiskLimitCheckIDType_RISKLIMITREQUESTID     RefRiskLimitCheckIDType = "0"
	RefRiskLimitCheckIDType_RISKLIMITCHECKID       RefRiskLimitCheckIDType = "1"
	RefRiskLimitCheckIDType_OUT_OF_BAND_IDENTIFIER RefRiskLimitCheckIDType = "3"
)

// ReferenceDataDateType field enumeration values.
type ReferenceDataDateType string

const (
	ReferenceDataDateType_DATE_OF_REQUEST_FOR_ADMISSION_TO_TRADING            ReferenceDataDateType = "0"
	ReferenceDataDateType_DATE_OF_APPROVAL_OF_ADMISSION_TO_TRADING            ReferenceDataDateType = "1"
	ReferenceDataDateType_DATE_OF_ADMISSION_TO_TRADING_OR_DATE_OF_FIRST_TRADE ReferenceDataDateType = "2"
	ReferenceDataDateType_TERMINATION_DATE                                    ReferenceDataDateType = "3"
)

// ReferenceEntityType field enumeration values.
type ReferenceEntityType string

const (
	ReferenceEntityType_ASIAN                           ReferenceEntityType = "1"
	ReferenceEntityType_WESTERN_EUROPEAN_INSURANCE      ReferenceEntityType = "10"
	ReferenceEntityType_AUSTRALIAN_AND_NEW_ZEALAND      ReferenceEntityType = "2"
	ReferenceEntityType_EUROPEAN_EMERGING_MARKETS       ReferenceEntityType = "3"
	ReferenceEntityType_JAPANESE                        ReferenceEntityType = "4"
	ReferenceEntityType_NORTH_AMERICAN_HIGH_YIELD       ReferenceEntityType = "5"
	ReferenceEntityType_NORTH_AMERICAN_INSURANCE        ReferenceEntityType = "6"
	ReferenceEntityType_NORTH_AMERICAN_INVESTMENT_GRADE ReferenceEntityType = "7"
	ReferenceEntityType_SINGAPOREAN                     ReferenceEntityType = "8"
	ReferenceEntityType_WESTERN_EUROPEAN                ReferenceEntityType = "9"
)

// RegistRejReasonCode field enumeration values.
type RegistRejReasonCode string

const (
	RegistRejReasonCode_INVALID_UNACCEPTABLE_ACCOUNT_TYPE                  RegistRejReasonCode = "1"
	RegistRejReasonCode_INVALID_UNACEEPTABLE_INVESTOR_ID_SOURCE            RegistRejReasonCode = "10"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_DATE_OF_BIRTH                 RegistRejReasonCode = "11"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_INVESTOR_COUNTRY_OF_RESIDENCE RegistRejReasonCode = "12"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_NO_DISTRIB_INSTNS             RegistRejReasonCode = "13"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_DISTRIB_PERCENTAGE            RegistRejReasonCode = "14"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_DISTRIB_PAYMENT_METHOD        RegistRejReasonCode = "15"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_CASH_DISTRIB_AGENT_ACCT_NAME  RegistRejReasonCode = "16"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_CASH_DISTRIB_AGENT_CODE       RegistRejReasonCode = "17"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_CASH_DISTRIB_AGENT_ACCT_NUM   RegistRejReasonCode = "18"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_TAX_EXEMPT_TYPE               RegistRejReasonCode = "2"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_OWNERSHIP_TYPE                RegistRejReasonCode = "3"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_NO_REG_DETAILS                RegistRejReasonCode = "4"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_REG_SEQ_NO                    RegistRejReasonCode = "5"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_REG_DETAILS                   RegistRejReasonCode = "6"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_MAILING_DETAILS               RegistRejReasonCode = "7"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_MAILING_INSTRUCTIONS          RegistRejReasonCode = "8"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_INVESTOR_ID                   RegistRejReasonCode = "9"
	RegistRejReasonCode_OTHER                                              RegistRejReasonCode = "99"
)

// RegistStatus field enumeration values.
type RegistStatus string

const (
	RegistStatus_ACCEPTED RegistStatus = "A"
	RegistStatus_HELD     RegistStatus = "H"
	RegistStatus_REMINDER RegistStatus = "N"
	RegistStatus_REJECTED RegistStatus = "R"
)

// RegistTransType field enumeration values.
type RegistTransType string

const (
	RegistTransType_NEW     RegistTransType = "0"
	RegistTransType_REPLACE RegistTransType = "1"
	RegistTransType_CANCEL  RegistTransType = "2"
)

// RegulatoryReportType field enumeration values.
type RegulatoryReportType string

const (
	RegulatoryReportType_REAL_TIME                                                                                              RegulatoryReportType = "0"
	RegulatoryReportType_PRIMARY_ECONOMIC_TERMS                                                                                 RegulatoryReportType = "1"
	RegulatoryReportType_POST_TRADE_EVENT_RT_REPORTABLE                                                                         RegulatoryReportType = "10"
	RegulatoryReportType_LIMITED_DETAILS_TRADE                                                                                  RegulatoryReportType = "11"
	RegulatoryReportType_DAILY_AGGREGATED_TRADE                                                                                 RegulatoryReportType = "12"
	RegulatoryReportType_VOLUME_OMISSION_TRADE                                                                                  RegulatoryReportType = "13"
	RegulatoryReportType_FOUR_WEEKS_AGGREGATION_TRADE                                                                           RegulatoryReportType = "14"
	RegulatoryReportType_INDEFINITE_AGGREGATION_TRADE                                                                           RegulatoryReportType = "15"
	RegulatoryReportType_VOLUME_OMISSION_TRADE_ELIGIBLE_FOR_SUBSEQUENT_AGGREGATED_ENRICHMENT                                    RegulatoryReportType = "16"
	RegulatoryReportType_FULL_DETAILS_TRADE_OF_LIMITED_DETAILS_TRADE                                                            RegulatoryReportType = "17"
	RegulatoryReportType_FULL_DETAILS_OF_DAILY_AGGREGATED_TRADE                                                                 RegulatoryReportType = "18"
	RegulatoryReportType_FULL_DETAILS_OF_VOLUME_OMISSION_TRADE                                                                  RegulatoryReportType = "19"
	RegulatoryReportType_SNAPSHOT                                                                                               RegulatoryReportType = "2"
	RegulatoryReportType_FULL_DETAILS_OF_FOUR_WEEKS_AGGREGATION_TRADE                                                           RegulatoryReportType = "20"
	RegulatoryReportType_FULL_DETAILS_IN_AGGREGATED_FORM_OF_VOLUME_OMISSION_TRADE_ELIGIBLE_FOR_SUBSEQUENT_AGGREGATED_ENRICHMENT RegulatoryReportType = "21"
	RegulatoryReportType_ORDER                                                                                                  RegulatoryReportType = "22"
	RegulatoryReportType_CHILD_ORDER                                                                                            RegulatoryReportType = "23"
	RegulatoryReportType_ORDER_ROUTE                                                                                            RegulatoryReportType = "24"
	RegulatoryReportType_TRADE                                                                                                  RegulatoryReportType = "25"
	RegulatoryReportType_QUOTE                                                                                                  RegulatoryReportType = "26"
	RegulatoryReportType_SUPPLEMENT                                                                                             RegulatoryReportType = "27"
	RegulatoryReportType_NEW_TRANSACTION                                                                                        RegulatoryReportType = "28"
	RegulatoryReportType_TRANSACTION_CORRECTION                                                                                 RegulatoryReportType = "29"
	RegulatoryReportType_CONFIRMATION                                                                                           RegulatoryReportType = "3"
	RegulatoryReportType_TRANSACTION_MODIFICATION                                                                               RegulatoryReportType = "30"
	RegulatoryReportType_COLLATERAL_UPDATE                                                                                      RegulatoryReportType = "31"
	RegulatoryReportType_MARGIN_UPDATE                                                                                          RegulatoryReportType = "32"
	RegulatoryReportType_TRANSACTION_REPORTED_IN_ERROR                                                                          RegulatoryReportType = "33"
	RegulatoryReportType_TERMINATION                                                                                            RegulatoryReportType = "34"
	RegulatoryReportType_COMBINATION_OF_RT_AND_PET                                                                              RegulatoryReportType = "4"
	RegulatoryReportType_COMBINATION_OF_PET_AND_CONFIRMATION                                                                    RegulatoryReportType = "5"
	RegulatoryReportType_COMBINATION_OF_RT_PET_AND_CONFIRMATION                                                                 RegulatoryReportType = "6"
	RegulatoryReportType_POST_TRADE_VALUATION                                                                                   RegulatoryReportType = "7"
	RegulatoryReportType_VERIFICATION                                                                                           RegulatoryReportType = "8"
	RegulatoryReportType_POST_TRADE_EVENT                                                                                       RegulatoryReportType = "9"
)

// RegulatoryTradeIDEvent field enumeration values.
type RegulatoryTradeIDEvent string

const (
	RegulatoryTradeIDEvent_INITIAL_BLOCK_TRADE  RegulatoryTradeIDEvent = "0"
	RegulatoryTradeIDEvent_ALLOCATION           RegulatoryTradeIDEvent = "1"
	RegulatoryTradeIDEvent_CLEARING             RegulatoryTradeIDEvent = "2"
	RegulatoryTradeIDEvent_COMPRESSION          RegulatoryTradeIDEvent = "3"
	RegulatoryTradeIDEvent_NOVATION             RegulatoryTradeIDEvent = "4"
	RegulatoryTradeIDEvent_TERMINATION          RegulatoryTradeIDEvent = "5"
	RegulatoryTradeIDEvent_POST_TRADE_VALUATION RegulatoryTradeIDEvent = "6"
)

// RegulatoryTradeIDScope field enumeration values.
type RegulatoryTradeIDScope string

const (
	RegulatoryTradeIDScope_CLEARING_MEMBER RegulatoryTradeIDScope = "1"
	RegulatoryTradeIDScope_CLIENT          RegulatoryTradeIDScope = "2"
)

// RegulatoryTradeIDType field enumeration values.
type RegulatoryTradeIDType string

const (
	RegulatoryTradeIDType_CURRENT                              RegulatoryTradeIDType = "0"
	RegulatoryTradeIDType_PREVIOUS                             RegulatoryTradeIDType = "1"
	RegulatoryTradeIDType_BLOCK                                RegulatoryTradeIDType = "2"
	RegulatoryTradeIDType_RELATED                              RegulatoryTradeIDType = "3"
	RegulatoryTradeIDType_CLEARED_BLOCK_TRADE                  RegulatoryTradeIDType = "4"
	RegulatoryTradeIDType_TRADING_VENUE_TRANSACTION_IDENTIFIER RegulatoryTradeIDType = "5"
)

// RegulatoryTransactionType field enumeration values.
type RegulatoryTransactionType string

const (
	RegulatoryTransactionType_NONE                      RegulatoryTransactionType = "0"
	RegulatoryTransactionType_SWAP_EXECUTION_FACILITY_1 RegulatoryTransactionType = "1"
	RegulatoryTransactionType_SWAP_EXECUTION_FACILITY_2 RegulatoryTransactionType = "2"
)

// RelatedInstrumentType field enumeration values.
type RelatedInstrumentType string

const (
	RelatedInstrumentType_HEDGES_FOR_INSTRUMENT                     RelatedInstrumentType = "1"
	RelatedInstrumentType_UNDERLIER                                 RelatedInstrumentType = "2"
	RelatedInstrumentType_EQUITY_EQUIVALENT                         RelatedInstrumentType = "3"
	RelatedInstrumentType_NEAREST_EXCHANGE_TRADED_CONTRACT          RelatedInstrumentType = "4"
	RelatedInstrumentType_RETAIL_EQUIVALENT_OF_WHOLESALE_INSTRUMENT RelatedInstrumentType = "5"
	RelatedInstrumentType_LEG                                       RelatedInstrumentType = "6"
)

// RelatedOrderIDSource field enumeration values.
type RelatedOrderIDSource string

const (
	RelatedOrderIDSource_NON_FIX_SOURCE                    RelatedOrderIDSource = "0"
	RelatedOrderIDSource_ORDER_IDENTIFIER                  RelatedOrderIDSource = "1"
	RelatedOrderIDSource_CLIENT_ORDER_IDENTIFIER           RelatedOrderIDSource = "2"
	RelatedOrderIDSource_SECONDARY_ORDER_IDENTIFIER        RelatedOrderIDSource = "3"
	RelatedOrderIDSource_SECONDARY_CLIENT_ORDER_IDENTIFIER RelatedOrderIDSource = "4"
)

// RelatedPositionIDSource field enumeration values.
type RelatedPositionIDSource string

const (
	RelatedPositionIDSource_POSITION_MAINTENANCE_REPORT_ID RelatedPositionIDSource = "1"
	RelatedPositionIDSource_POSITION_TRANSFER_ID           RelatedPositionIDSource = "2"
	RelatedPositionIDSource_POSITION_ENTITY_ID             RelatedPositionIDSource = "3"
)

// RelatedPriceSource field enumeration values.
type RelatedPriceSource string

const (
	RelatedPriceSource_NBB RelatedPriceSource = "1"
	RelatedPriceSource_NBO RelatedPriceSource = "2"
)

// RelatedTradeIDSource field enumeration values.
type RelatedTradeIDSource string

const (
	RelatedTradeIDSource_NON_FIX_SOURCE          RelatedTradeIDSource = "0"
	RelatedTradeIDSource_TRADE_ID                RelatedTradeIDSource = "1"
	RelatedTradeIDSource_SECONDARY_TRADE_ID      RelatedTradeIDSource = "2"
	RelatedTradeIDSource_TRADE_REPORT_ID         RelatedTradeIDSource = "3"
	RelatedTradeIDSource_FIRM_TRADE_ID           RelatedTradeIDSource = "4"
	RelatedTradeIDSource_SECONDARY_FIRM_TRADE_ID RelatedTradeIDSource = "5"
	RelatedTradeIDSource_REGULATORY_TRADE_ID     RelatedTradeIDSource = "6"
)

// RelativeValueSide field enumeration values.
type RelativeValueSide string

const (
	RelativeValueSide_BID   RelativeValueSide = "1"
	RelativeValueSide_MID   RelativeValueSide = "2"
	RelativeValueSide_OFFER RelativeValueSide = "3"
)

// RelativeValueType field enumeration values.
type RelativeValueType string

const (
	RelativeValueType_ASSET_SWAP_SPREAD             RelativeValueType = "1"
	RelativeValueType_OVERNIGHT_INDEXED_SWAP_SPREAD RelativeValueType = "2"
	RelativeValueType_ZERO_VOLATILITY_SPREAD        RelativeValueType = "3"
	RelativeValueType_DISCOUNT_MARGIN               RelativeValueType = "4"
	RelativeValueType_INTERPOLATED_SPREAD           RelativeValueType = "5"
	RelativeValueType_OPTION_ADJUSTED_SPREAD        RelativeValueType = "6"
	RelativeValueType_G_SPREAD                      RelativeValueType = "7"
	RelativeValueType_CDS_BASIS                     RelativeValueType = "8"
	RelativeValueType_CDS_INTERPOLATED_BASIS        RelativeValueType = "9"
)

// ReleaseInstruction field enumeration values.
type ReleaseInstruction string

const (
	ReleaseInstruction_INTERMARKET_SWEEP_ORDER     ReleaseInstruction = "1"
	ReleaseInstruction_NO_AWAY_MARKET_BETTER_CHECK ReleaseInstruction = "2"
)

// RemunerationIndicator field enumeration values.
type RemunerationIndicator string

const (
	RemunerationIndicator_NO_REMUNERATION_PAID RemunerationIndicator = "0"
	RemunerationIndicator_REMUNERATION_PAID    RemunerationIndicator = "1"
)

// ReportToExch field enumeration values.
type ReportToExch string

const (
	ReportToExch_NO  ReportToExch = "N"
	ReportToExch_YES ReportToExch = "Y"
)

// RequestResult field enumeration values.
type RequestResult string

const (
	RequestResult_VALID_REQUEST                               RequestResult = "0"
	RequestResult_INVALID_OR_UNSUPPORTED_REQUEST              RequestResult = "1"
	RequestResult_NO_DATA_FOUND_THAT_MATCH_SELECTION_CRITERIA RequestResult = "2"
	RequestResult_NOT_AUTHORIZED_TO_RETRIEVE_DATA             RequestResult = "3"
	RequestResult_DATA_TEMPORARILY_UNAVAILABLE                RequestResult = "4"
	RequestResult_REQUEST_FOR_DATA_NOT_SUPPORTED              RequestResult = "5"
	RequestResult_OTHER                                       RequestResult = "99"
)

// ResetSeqNumFlag field enumeration values.
type ResetSeqNumFlag string

const (
	ResetSeqNumFlag_NO  ResetSeqNumFlag = "N"
	ResetSeqNumFlag_YES ResetSeqNumFlag = "Y"
)

// RespondentType field enumeration values.
type RespondentType string

const (
	RespondentType_ALL_MARKET_PARTICIPANTS       RespondentType = "1"
	RespondentType_SPECIFIED_MARKET_PARTICIPANTS RespondentType = "2"
	RespondentType_ALL_MARKET_MAKERS             RespondentType = "3"
	RespondentType_PRIMARY_MARKET_MAKER          RespondentType = "4"
)

// ResponseTransportType field enumeration values.
type ResponseTransportType string

const (
	ResponseTransportType_IN_BAND     ResponseTransportType = "0"
	ResponseTransportType_OUT_OF_BAND ResponseTransportType = "1"
)

// RestructuringType field enumeration values.
type RestructuringType string

const (
	RestructuringType_FULL_RESTRUCTURING         RestructuringType = "FR"
	RestructuringType_MODIFIED_MOD_RESTRUCTURING RestructuringType = "MM"
	RestructuringType_MODIFIED_RESTRUCTURING     RestructuringType = "MR"
	RestructuringType_NO_RESTRUCTURING_SPECIFIED RestructuringType = "XR"
)

// ReturnRateDateMode field enumeration values.
type ReturnRateDateMode string

const (
	ReturnRateDateMode_PRICE_VALUATION    ReturnRateDateMode = "0"
	ReturnRateDateMode_DIVIDEND_VALUATION ReturnRateDateMode = "1"
)

// ReturnRatePriceBasis field enumeration values.
type ReturnRatePriceBasis string

const (
	ReturnRatePriceBasis_GROSS     ReturnRatePriceBasis = "0"
	ReturnRatePriceBasis_NET       ReturnRatePriceBasis = "1"
	ReturnRatePriceBasis_ACCRUED   ReturnRatePriceBasis = "2"
	ReturnRatePriceBasis_CLEAN_NET ReturnRatePriceBasis = "3"
)

// ReturnRatePriceSequence field enumeration values.
type ReturnRatePriceSequence string

const (
	ReturnRatePriceSequence_INITIAL ReturnRatePriceSequence = "0"
	ReturnRatePriceSequence_INTERIM ReturnRatePriceSequence = "1"
	ReturnRatePriceSequence_FINAL   ReturnRatePriceSequence = "2"
)

// ReturnRatePriceType field enumeration values.
type ReturnRatePriceType string

const (
	ReturnRatePriceType_ABSOLUTE_TERMS         ReturnRatePriceType = "0"
	ReturnRatePriceType_PERCENTAGE_OF_NOTIONAL ReturnRatePriceType = "1"
)

// ReturnRateQuoteTimeType field enumeration values.
type ReturnRateQuoteTimeType string

const (
	ReturnRateQuoteTimeType_OPEN                                    ReturnRateQuoteTimeType = "0"
	ReturnRateQuoteTimeType_OFFICIAL_SETTLEMENT_PRICE_TIME          ReturnRateQuoteTimeType = "1"
	ReturnRateQuoteTimeType_XETRA                                   ReturnRateQuoteTimeType = "2"
	ReturnRateQuoteTimeType_CLOSE                                   ReturnRateQuoteTimeType = "3"
	ReturnRateQuoteTimeType_DERIVATIVES_CLOSE                       ReturnRateQuoteTimeType = "4"
	ReturnRateQuoteTimeType_HIGH                                    ReturnRateQuoteTimeType = "5"
	ReturnRateQuoteTimeType_LOW                                     ReturnRateQuoteTimeType = "6"
	ReturnRateQuoteTimeType_AS_SPECIFIED_IN_THE_MASTER_CONFIRMATION ReturnRateQuoteTimeType = "7"
)

// ReturnRateValuationPriceOption field enumeration values.
type ReturnRateValuationPriceOption string

const (
	ReturnRateValuationPriceOption_NONE          ReturnRateValuationPriceOption = "0"
	ReturnRateValuationPriceOption_FUTURES_PRICE ReturnRateValuationPriceOption = "1"
	ReturnRateValuationPriceOption_OPTIONS_PRICE ReturnRateValuationPriceOption = "2"
)

// ReturnTrigger field enumeration values.
type ReturnTrigger string

const (
	ReturnTrigger_DIVIDEND                               ReturnTrigger = "1"
	ReturnTrigger_VARIANCE                               ReturnTrigger = "2"
	ReturnTrigger_VOLATILITY                             ReturnTrigger = "3"
	ReturnTrigger_TOTAL_RETURN                           ReturnTrigger = "4"
	ReturnTrigger_CONTRACT_FOR_DIFFERENCE                ReturnTrigger = "5"
	ReturnTrigger_CREDIT_DEFAULT                         ReturnTrigger = "6"
	ReturnTrigger_SPREAD_BET                             ReturnTrigger = "7"
	ReturnTrigger_PRICE                                  ReturnTrigger = "8"
	ReturnTrigger_FORWARD_PRICE_OF_UNDERLYING_INSTRUMENT ReturnTrigger = "9"
	ReturnTrigger_OTHER                                  ReturnTrigger = "99"
)

// RiskLimitAction field enumeration values.
type RiskLimitAction string

const (
	RiskLimitAction_QUEUE_INBOUND                                RiskLimitAction = "0"
	RiskLimitAction_QUEUE_OUTBOUND                               RiskLimitAction = "1"
	RiskLimitAction_HALT_TRADING                                 RiskLimitAction = "10"
	RiskLimitAction_REJECT                                       RiskLimitAction = "2"
	RiskLimitAction_DISCONNECT                                   RiskLimitAction = "3"
	RiskLimitAction_WARNING                                      RiskLimitAction = "4"
	RiskLimitAction_PING_CREDIT_CHECK_MODEL_WITH_REVALIDATION    RiskLimitAction = "5"
	RiskLimitAction_PING_CREDIT_CHECK_MODEL_WITHOUT_REVALIDATION RiskLimitAction = "6"
	RiskLimitAction_PUSH_CREDIT_CHECK_MODEL_WITH_REVALIDATION    RiskLimitAction = "7"
	RiskLimitAction_PUSH_CREDIT_CHECK_MODEL_WITHOUT_REVALIDATION RiskLimitAction = "8"
	RiskLimitAction_SUSPEND                                      RiskLimitAction = "9"
)

// RiskLimitCheckModelType field enumeration values.
type RiskLimitCheckModelType string

const (
	RiskLimitCheckModelType_NONE          RiskLimitCheckModelType = "0"
	RiskLimitCheckModelType_PLUSONE_MODEL RiskLimitCheckModelType = "1"
	RiskLimitCheckModelType_PING_MODEL    RiskLimitCheckModelType = "2"
	RiskLimitCheckModelType_PUSH_MODEL    RiskLimitCheckModelType = "3"
)

// RiskLimitCheckRequestResult field enumeration values.
type RiskLimitCheckRequestResult string

const (
	RiskLimitCheckRequestResult_SUCCESSFUL                                    RiskLimitCheckRequestResult = "0"
	RiskLimitCheckRequestResult_INVALID_PARTY                                 RiskLimitCheckRequestResult = "1"
	RiskLimitCheckRequestResult_REQUESTED_AMOUNT_EXCEEDS_CREDIT_LIMIT         RiskLimitCheckRequestResult = "2"
	RiskLimitCheckRequestResult_REQUESTED_AMOUNT_EXCEEDS_CLIP_SIZE_LIMIT      RiskLimitCheckRequestResult = "3"
	RiskLimitCheckRequestResult_REQUEST_EXCEEDS_MAXIMUM_NOTIONAL_ORDER_AMOUNT RiskLimitCheckRequestResult = "4"
	RiskLimitCheckRequestResult_OTHER                                         RiskLimitCheckRequestResult = "99"
)

// RiskLimitCheckRequestStatus field enumeration values.
type RiskLimitCheckRequestStatus string

const (
	RiskLimitCheckRequestStatus_APPROVED           RiskLimitCheckRequestStatus = "0"
	RiskLimitCheckRequestStatus_PARTIALLY_APPROVED RiskLimitCheckRequestStatus = "1"
	RiskLimitCheckRequestStatus_REJECTED           RiskLimitCheckRequestStatus = "2"
	RiskLimitCheckRequestStatus_APPROVAL_PENDING   RiskLimitCheckRequestStatus = "3"
	RiskLimitCheckRequestStatus_CANCELLED          RiskLimitCheckRequestStatus = "4"
)

// RiskLimitCheckRequestType field enumeration values.
type RiskLimitCheckRequestType string

const (
	RiskLimitCheckRequestType_ALL_OR_NONE RiskLimitCheckRequestType = "0"
	RiskLimitCheckRequestType_PARTIAL     RiskLimitCheckRequestType = "1"
)

// RiskLimitCheckStatus field enumeration values.
type RiskLimitCheckStatus string

const (
	RiskLimitCheckStatus_ACCEPTED                             RiskLimitCheckStatus = "0"
	RiskLimitCheckStatus_REJECTED                             RiskLimitCheckStatus = "1"
	RiskLimitCheckStatus_ACCEPTED_BY_CREDIT_HUB               RiskLimitCheckStatus = "10"
	RiskLimitCheckStatus_REJECTED_BY_CREDIT_HUB               RiskLimitCheckStatus = "11"
	RiskLimitCheckStatus_PENDING_CREDIT_HUB_CHECK             RiskLimitCheckStatus = "12"
	RiskLimitCheckStatus_ACCEPTED_BY_EXECUTION_VENUE          RiskLimitCheckStatus = "13"
	RiskLimitCheckStatus_REJECTED_BY_EXECUTION_VENUE          RiskLimitCheckStatus = "14"
	RiskLimitCheckStatus_CLAIM_REQUIRED                       RiskLimitCheckStatus = "2"
	RiskLimitCheckStatus_PRE_DEFINED_LIMIT_CHECK_SUCCEEDED    RiskLimitCheckStatus = "3"
	RiskLimitCheckStatus_PRE_DEFINED_LIMIT_CHECK_FAILED       RiskLimitCheckStatus = "4"
	RiskLimitCheckStatus_PRE_DEFINED_AUTO_ACCEPT_RULE_INVOKED RiskLimitCheckStatus = "5"
	RiskLimitCheckStatus_PRE_DEFINED_AUTO_REJECT_RULE_INVOKED RiskLimitCheckStatus = "6"
	RiskLimitCheckStatus_ACCEPTED_BY_CLEARING_FIRM            RiskLimitCheckStatus = "7"
	RiskLimitCheckStatus_REJECTED_BY_CLEARING_FIRM            RiskLimitCheckStatus = "8"
	RiskLimitCheckStatus_PENDING                              RiskLimitCheckStatus = "9"
)

// RiskLimitCheckTransType field enumeration values.
type RiskLimitCheckTransType string

const (
	RiskLimitCheckTransType_NEW     RiskLimitCheckTransType = "0"
	RiskLimitCheckTransType_CANCEL  RiskLimitCheckTransType = "1"
	RiskLimitCheckTransType_REPLACE RiskLimitCheckTransType = "2"
)

// RiskLimitCheckType field enumeration values.
type RiskLimitCheckType string

const (
	RiskLimitCheckType_SUBMIT         RiskLimitCheckType = "0"
	RiskLimitCheckType_LIMIT_CONSUMED RiskLimitCheckType = "1"
)

// RiskLimitReportRejectReason field enumeration values.
type RiskLimitReportRejectReason string

const (
	RiskLimitReportRejectReason_UNKNOWN_RISKLIMITREPORTID RiskLimitReportRejectReason = "0"
	RiskLimitReportRejectReason_UNKNOWN_PARTY             RiskLimitReportRejectReason = "1"
	RiskLimitReportRejectReason_OTHER                     RiskLimitReportRejectReason = "99"
)

// RiskLimitReportStatus field enumeration values.
type RiskLimitReportStatus string

const (
	RiskLimitReportStatus_ACCEPTED RiskLimitReportStatus = "0"
	RiskLimitReportStatus_REJECTED RiskLimitReportStatus = "1"
)

// RiskLimitRequestResult field enumeration values.
type RiskLimitRequestResult string

const (
	RiskLimitRequestResult_SUCCESSFUL                           RiskLimitRequestResult = "0"
	RiskLimitRequestResult_INVALID_PARTY                        RiskLimitRequestResult = "1"
	RiskLimitRequestResult_WARNING_LEVEL_ACTIONS_NOT_SUPPORTED  RiskLimitRequestResult = "10"
	RiskLimitRequestResult_RISK_INSTRUMENT_SCOPE_NOT_SUPPORTED  RiskLimitRequestResult = "11"
	RiskLimitRequestResult_RISK_LIMIT_NOT_APPROVED_FOR_PARTY    RiskLimitRequestResult = "12"
	RiskLimitRequestResult_RISK_LIMIT_ALREADY_DEFINED_FOR_PARTY RiskLimitRequestResult = "13"
	RiskLimitRequestResult_INSTRUMENT_NOT_APPROVED_FOR_PARTY    RiskLimitRequestResult = "14"
	RiskLimitRequestResult_INVALID_RELATED_PARTY                RiskLimitRequestResult = "2"
	RiskLimitRequestResult_INVALID_RISK_LIMIT_TYPE              RiskLimitRequestResult = "3"
	RiskLimitRequestResult_INVALID_RISK_LIMIT_ID                RiskLimitRequestResult = "4"
	RiskLimitRequestResult_INVALID_RISK_LIMIT_AMOUNT            RiskLimitRequestResult = "5"
	RiskLimitRequestResult_INVALID_RISK_WARNING_LEVEL_ACTION    RiskLimitRequestResult = "6"
	RiskLimitRequestResult_INVALID_RISK_INSTRUMENT_SCOPE        RiskLimitRequestResult = "7"
	RiskLimitRequestResult_RISK_LIMIT_ACTIONS_NOT_SUPPORTED     RiskLimitRequestResult = "8"
	RiskLimitRequestResult_WARNING_LEVELS_NOT_SUPPORTED         RiskLimitRequestResult = "9"
	RiskLimitRequestResult_NOT_AUTHORIZED                       RiskLimitRequestResult = "98"
	RiskLimitRequestResult_OTHER                                RiskLimitRequestResult = "99"
)

// RiskLimitRequestType field enumeration values.
type RiskLimitRequestType string

const (
	RiskLimitRequestType_DEFINITIONS                 RiskLimitRequestType = "1"
	RiskLimitRequestType_UTILIZATION                 RiskLimitRequestType = "2"
	RiskLimitRequestType_DEFINITIONS_AND_UTILIZATION RiskLimitRequestType = "3"
)

// RiskLimitStatus field enumeration values.
type RiskLimitStatus string

const (
	RiskLimitStatus_ACCEPTED              RiskLimitStatus = "0"
	RiskLimitStatus_ACCEPTED_WITH_CHANGES RiskLimitStatus = "1"
	RiskLimitStatus_REJECTED              RiskLimitStatus = "2"
)

// RiskLimitType field enumeration values.
type RiskLimitType string

const (
	RiskLimitType_CREDIT_LIMIT                                               RiskLimitType = "0"
	RiskLimitType_GROSS_LIMIT                                                RiskLimitType = "1"
	RiskLimitType_CLIP_SIZE_NOTIONAL_LIMIT_PER_TIME_PERIOD                   RiskLimitType = "10"
	RiskLimitType_MAXIMUM_NOTIONAL_ORDER_SIZE                                RiskLimitType = "11"
	RiskLimitType_DV01_PV01_LIMIT                                            RiskLimitType = "12"
	RiskLimitType_CS01_LIMIT                                                 RiskLimitType = "13"
	RiskLimitType_VOLUME_LIMIT_PER_TIME_PERIOD                               RiskLimitType = "14"
	RiskLimitType_VOLUME_FILLED_AS_PERCENT_OF_ORDERED_VOLUME_PER_TIME_PERIOD RiskLimitType = "15"
	RiskLimitType_NOTIONAL_FILLED_AS_PERCENT_OF_NOTIONAL_PER_TIME_PERIOD     RiskLimitType = "16"
	RiskLimitType_TRANSACTION_EXECUTION_LIMIT_PER_TIME_PERIOD                RiskLimitType = "17"
	RiskLimitType_NET_LIMIT                                                  RiskLimitType = "2"
	RiskLimitType_EXPOSURE                                                   RiskLimitType = "3"
	RiskLimitType_LONG_LIMIT                                                 RiskLimitType = "4"
	RiskLimitType_SHORT_LIMIT                                                RiskLimitType = "5"
	RiskLimitType_CASH_MARGIN                                                RiskLimitType = "6"
	RiskLimitType_ADDITIONAL_MARGIN                                          RiskLimitType = "7"
	RiskLimitType_TOTAL_MARGIN                                               RiskLimitType = "8"
	RiskLimitType_LIMIT_CONSUMED                                             RiskLimitType = "9"
)

// RoundingDirection field enumeration values.
type RoundingDirection string

const (
	RoundingDirection_ROUND_TO_NEAREST RoundingDirection = "0"
	RoundingDirection_ROUND_DOWN       RoundingDirection = "1"
	RoundingDirection_ROUND_UP         RoundingDirection = "2"
)

// RoutingArrangmentIndicator field enumeration values.
type RoutingArrangmentIndicator string

const (
	RoutingArrangmentIndicator_NO_ROUTING_ARRANGEMENT_IN_PLACE RoutingArrangmentIndicator = "0"
	RoutingArrangmentIndicator_ROUTING_ARRANGEMENT_IN_PLACE    RoutingArrangmentIndicator = "1"
)

// RoutingType field enumeration values.
type RoutingType string

const (
	RoutingType_TARGET_FIRM   RoutingType = "1"
	RoutingType_TARGET_LIST   RoutingType = "2"
	RoutingType_BLOCK_FIRM    RoutingType = "3"
	RoutingType_BLOCK_LIST    RoutingType = "4"
	RoutingType_TARGET_PERSON RoutingType = "5"
	RoutingType_BLOCK_PERSON  RoutingType = "6"
)

// Rule80A field enumeration values.
type Rule80A string

const (
	Rule80A_AGENCY_SINGLE_ORDER                                                                                        Rule80A = "A"
	Rule80A_SHORT_EXEMPT_TRANSACTION_B                                                                                 Rule80A = "B"
	Rule80A_PROPRIETARY_NON_ALGORITHMIC_PROGRAM_TRADE                                                                  Rule80A = "C"
	Rule80A_PROGRAM_ORDER_INDEX_ARB_FOR_MEMBER_FIRM_ORG                                                                Rule80A = "D"
	Rule80A_SHORT_EXEMPT_TRANSACTION_FOR_PRINCIPAL                                                                     Rule80A = "E"
	Rule80A_SHORT_EXEMPT_TRANSACTION_F                                                                                 Rule80A = "F"
	Rule80A_SHORT_EXEMPT_TRANSACTION_H                                                                                 Rule80A = "H"
	Rule80A_INDIVIDUAL_INVESTOR_SINGLE_ORDER                                                                           Rule80A = "I"
	Rule80A_PROPRIETARY_ALGORITHMIC_PROGRAM_TRADING                                                                    Rule80A = "J"
	Rule80A_AGENCY_ALGORITHMIC_PROGRAM_TRADING                                                                         Rule80A = "K"
	Rule80A_SHORT_EXEMPT_TRANSACTION_FOR_MEMBER_COMPETING_MARKET_MAKER_AFFLIATED_WITH_THE_FIRM_CLEARING_THE_TRADE      Rule80A = "L"
	Rule80A_PROGRAM_ORDER_INDEX_ARB_FOR_OTHER_MEMBER                                                                   Rule80A = "M"
	Rule80A_AGENT_FOR_OTHER_MEMBER_NON_ALGORITHMIC_PROGRAM_TRADE                                                       Rule80A = "N"
	Rule80A_PROPRIETARY_TRANSACTIONS_FOR_COMPETING_MARKET_MAKER_THAT_IS_AFFILIATED_WITH_THE_CLEARING_MEMBER            Rule80A = "O"
	Rule80A_PRINCIPAL                                                                                                  Rule80A = "P"
	Rule80A_TRANSACTIONS_FOR_THE_ACCOUNT_OF_A_NON_MEMBER_COMPTING_MARKET_MAKER                                         Rule80A = "R"
	Rule80A_SPECIALIST_TRADES                                                                                          Rule80A = "S"
	Rule80A_TRANSACTIONS_FOR_THE_ACCOUNT_OF_AN_UNAFFILIATED_MEMBERS_COMPETING_MARKET_MAKER                             Rule80A = "T"
	Rule80A_AGENCY_INDEX_ARBITRAGE                                                                                     Rule80A = "U"
	Rule80A_ALL_OTHER_ORDERS_AS_AGENT_FOR_OTHER_MEMBER                                                                 Rule80A = "W"
	Rule80A_SHORT_EXEMPT_TRANSACTION_FOR_MEMBER_COMPETING_MARKET_MAKER_NOT_AFFILIATED_WITH_THE_FIRM_CLEARING_THE_TRADE Rule80A = "X"
	Rule80A_AGENCY_NON_ALGORITHMIC_PROGRAM_TRADE                                                                       Rule80A = "Y"
	Rule80A_SHORT_EXEMPT_TRANSACTION_FOR_NON_MEMBER_COMPETING_MARKET_MAKER                                             Rule80A = "Z"
)

// Scope field enumeration values.
type Scope string

const (
	Scope_LOCAL_MARKET Scope = "1"
	Scope_NATIONAL     Scope = "2"
	Scope_GLOBAL       Scope = "3"
)

// SecDefStatus field enumeration values.
type SecDefStatus string

const (
	SecDefStatus_PENDING_APPROVAL           SecDefStatus = "0"
	SecDefStatus_APPROVED                   SecDefStatus = "1"
	SecDefStatus_REJECTED                   SecDefStatus = "2"
	SecDefStatus_UNAUTHORIZED_REQUEST       SecDefStatus = "3"
	SecDefStatus_INVALID_DEFINITION_REQUEST SecDefStatus = "4"
)

// SecurityClassificationReason field enumeration values.
type SecurityClassificationReason string

const (
	SecurityClassificationReason_FEE               SecurityClassificationReason = "0"
	SecurityClassificationReason_CREDIT_CONTROLS   SecurityClassificationReason = "1"
	SecurityClassificationReason_MARGIN            SecurityClassificationReason = "2"
	SecurityClassificationReason_ENTITLEMENT       SecurityClassificationReason = "3"
	SecurityClassificationReason_MARKET_DATA       SecurityClassificationReason = "4"
	SecurityClassificationReason_ACCOUNT_SELECTION SecurityClassificationReason = "5"
	SecurityClassificationReason_DELIVERY_PROCESS  SecurityClassificationReason = "6"
	SecurityClassificationReason_SECTOR            SecurityClassificationReason = "7"
)

// SecurityIDSource field enumeration values.
type SecurityIDSource string

const (
	SecurityIDSource_CUSIP                                  SecurityIDSource = "1"
	SecurityIDSource_SEDOL                                  SecurityIDSource = "2"
	SecurityIDSource_QUIK                                   SecurityIDSource = "3"
	SecurityIDSource_ISIN                                   SecurityIDSource = "4"
	SecurityIDSource_RIC                                    SecurityIDSource = "5"
	SecurityIDSource_ISO_CURRENCY_CODE                      SecurityIDSource = "6"
	SecurityIDSource_ISO_COUNTRY_CODE                       SecurityIDSource = "7"
	SecurityIDSource_EXCHANGE_SYMBOL                        SecurityIDSource = "8"
	SecurityIDSource_CONSOLIDATED_TAPE_ASSOCIATION          SecurityIDSource = "9"
	SecurityIDSource_BLOOMBERG_SYMBOL                       SecurityIDSource = "A"
	SecurityIDSource_WERTPAPIER                             SecurityIDSource = "B"
	SecurityIDSource_DUTCH                                  SecurityIDSource = "C"
	SecurityIDSource_VALOREN                                SecurityIDSource = "D"
	SecurityIDSource_SICOVAM                                SecurityIDSource = "E"
	SecurityIDSource_BELGIAN                                SecurityIDSource = "F"
	SecurityIDSource_COMMON                                 SecurityIDSource = "G"
	SecurityIDSource_CLEARING_HOUSE                         SecurityIDSource = "H"
	SecurityIDSource_ISDA_FPML_PRODUCT_SPECIFICATION        SecurityIDSource = "I"
	SecurityIDSource_OPTION_PRICE_REPORTING_AUTHORITY       SecurityIDSource = "J"
	SecurityIDSource_ISDA_FPML_PRODUCT_URL                  SecurityIDSource = "K"
	SecurityIDSource_LETTER_OF_CREDIT                       SecurityIDSource = "L"
	SecurityIDSource_MARKETPLACE_ASSIGNED_IDENTIFIER        SecurityIDSource = "M"
	SecurityIDSource_MARKIT_RED_ENTITY_CLIP                 SecurityIDSource = "N"
	SecurityIDSource_MARKIT_RED_PAIR_CLIP                   SecurityIDSource = "P"
	SecurityIDSource_CFTC_COMMODITY_CODE                    SecurityIDSource = "Q"
	SecurityIDSource_ISDA_COMMODITY_REFERENCE_PRICE         SecurityIDSource = "R"
	SecurityIDSource_FINANCIAL_INSTRUMENT_GLOBAL_IDENTIFIER SecurityIDSource = "S"
	SecurityIDSource_LEGAL_ENTITY_IDENTIFIER                SecurityIDSource = "T"
	SecurityIDSource_SYNTHETIC                              SecurityIDSource = "U"
	SecurityIDSource_FIDESSA_INSTRUMENT_MNEMONIC            SecurityIDSource = "V"
	SecurityIDSource_INDEX_NAME                             SecurityIDSource = "W"
	SecurityIDSource_UNIFORM_SYMBOL                         SecurityIDSource = "X"
)

// SecurityListRequestType field enumeration values.
type SecurityListRequestType string

const (
	SecurityListRequestType_SYMBOL                                    SecurityListRequestType = "0"
	SecurityListRequestType_SECURITYTYPE_AND_OR_CFICODE               SecurityListRequestType = "1"
	SecurityListRequestType_PRODUCT                                   SecurityListRequestType = "2"
	SecurityListRequestType_TRADINGSESSIONID                          SecurityListRequestType = "3"
	SecurityListRequestType_ALL_SECURITIES                            SecurityListRequestType = "4"
	SecurityListRequestType_MARKETID_OR_MARKETID_PLUS_MARKETSEGMENTID SecurityListRequestType = "5"
)

// SecurityListType field enumeration values.
type SecurityListType string

const (
	SecurityListType_INDUSTRY_CLASSIFICATION SecurityListType = "1"
	SecurityListType_TRADING_LIST            SecurityListType = "2"
	SecurityListType_MARKET                  SecurityListType = "3"
	SecurityListType_NEWSPAPER_LIST          SecurityListType = "4"
)

// SecurityListTypeSource field enumeration values.
type SecurityListTypeSource string

const (
	SecurityListTypeSource_ICB   SecurityListTypeSource = "1"
	SecurityListTypeSource_NAICS SecurityListTypeSource = "2"
	SecurityListTypeSource_GICS  SecurityListTypeSource = "3"
)

// SecurityRejectReason field enumeration values.
type SecurityRejectReason string

const (
	SecurityRejectReason_INVALID_INSTRUMENT_REQUESTED               SecurityRejectReason = "1"
	SecurityRejectReason_INVALID_OR_MISSING_DATA_ON_FX_LEG          SecurityRejectReason = "10"
	SecurityRejectReason_INVALID_LEG_PRICE_SPECIFIED                SecurityRejectReason = "11"
	SecurityRejectReason_INVALID_INSTRUMENT_STRUCTURE_SPECIFIED     SecurityRejectReason = "12"
	SecurityRejectReason_INSTRUMENT_ALREADY_EXISTS                  SecurityRejectReason = "2"
	SecurityRejectReason_REQUEST_TYPE_NOT_SUPPORTED                 SecurityRejectReason = "3"
	SecurityRejectReason_SYSTEM_UNAVAILABLE_FOR_INSTRUMENT_CREATION SecurityRejectReason = "4"
	SecurityRejectReason_INELIGIBLE_INSTRUMENT_GROUP                SecurityRejectReason = "5"
	SecurityRejectReason_INSTRUMENT_ID_UNAVAILABLE                  SecurityRejectReason = "6"
	SecurityRejectReason_INVALID_OR_MISSING_DATA_ON_OPTION_LEG      SecurityRejectReason = "7"
	SecurityRejectReason_INVALID_OR_MISSING_DATA_ON_FUTURE_LEG      SecurityRejectReason = "8"
)

// SecurityRequestResult field enumeration values.
type SecurityRequestResult string

const (
	SecurityRequestResult_VALID_REQUEST                                      SecurityRequestResult = "0"
	SecurityRequestResult_INVALID_OR_UNSUPPORTED_REQUEST                     SecurityRequestResult = "1"
	SecurityRequestResult_NO_INSTRUMENTS_FOUND_THAT_MATCH_SELECTION_CRITERIA SecurityRequestResult = "2"
	SecurityRequestResult_NOT_AUTHORIZED_TO_RETRIEVE_INSTRUMENT_DATA         SecurityRequestResult = "3"
	SecurityRequestResult_INSTRUMENT_DATA_TEMPORARILY_UNAVAILABLE            SecurityRequestResult = "4"
	SecurityRequestResult_REQUEST_FOR_INSTRUMENT_DATA_NOT_SUPPORTED          SecurityRequestResult = "5"
)

// SecurityRequestType field enumeration values.
type SecurityRequestType string

const (
	SecurityRequestType_REQUEST_SECURITY_IDENTITY_AND_SPECIFICATIONS              SecurityRequestType = "0"
	SecurityRequestType_REQUEST_SECURITY_IDENTITY_FOR_THE_SPECIFICATIONS_PROVIDED SecurityRequestType = "1"
	SecurityRequestType_REQUEST_LIST_SECURITY_TYPES                               SecurityRequestType = "2"
	SecurityRequestType_REQUEST_LIST_SECURITIES                                   SecurityRequestType = "3"
	SecurityRequestType_SYMBOL                                                    SecurityRequestType = "4"
	SecurityRequestType_SECURITYTYPE_AND_OR_CFICODE                               SecurityRequestType = "5"
	SecurityRequestType_PRODUCT                                                   SecurityRequestType = "6"
	SecurityRequestType_TRADINGSESSIONID                                          SecurityRequestType = "7"
	SecurityRequestType_ALL_SECURITIES                                            SecurityRequestType = "8"
	SecurityRequestType_MARKETID_OR_MARKETID_PLUS_MARKETSEGMENTID                 SecurityRequestType = "9"
)

// SecurityResponseType field enumeration values.
type SecurityResponseType string

const (
	SecurityResponseType_ACCEPT_SECURITY_PROPOSAL_AS_IS                                      SecurityResponseType = "1"
	SecurityResponseType_ACCEPT_SECURITY_PROPOSAL_WITH_REVISIONS_AS_INDICATED_IN_THE_MESSAGE SecurityResponseType = "2"
	SecurityResponseType_LIST_OF_SECURITY_TYPES_RETURNED_PER_REQUEST                         SecurityResponseType = "3"
	SecurityResponseType_LIST_OF_SECURITIES_RETURNED_PER_REQUEST                             SecurityResponseType = "4"
	SecurityResponseType_REJECT_SECURITY_PROPOSAL                                            SecurityResponseType = "5"
	SecurityResponseType_CANNOT_MATCH_SELECTION_CRITERIA                                     SecurityResponseType = "6"
)

// SecurityStatus field enumeration values.
type SecurityStatus string

const (
	SecurityStatus_ACTIVE                     SecurityStatus = "1"
	SecurityStatus_PUBLISHED                  SecurityStatus = "10"
	SecurityStatus_PENDING_DELETION           SecurityStatus = "11"
	SecurityStatus_INACTIVE                   SecurityStatus = "2"
	SecurityStatus_ACTIVE_CLOSING_ORDERS_ONLY SecurityStatus = "3"
	SecurityStatus_EXPIRED                    SecurityStatus = "4"
	SecurityStatus_DELISTED                   SecurityStatus = "5"
	SecurityStatus_KNOCKED_OUT                SecurityStatus = "6"
	SecurityStatus_KNOCK_OUT_REVOKED          SecurityStatus = "7"
	SecurityStatus_PENDING_EXPIRY             SecurityStatus = "8"
	SecurityStatus_SUSPENDED                  SecurityStatus = "9"
)

// SecurityTradingEvent field enumeration values.
type SecurityTradingEvent string

const (
	SecurityTradingEvent_ORDER_IMBALANCE_AUCTION_IS_EXTENDED SecurityTradingEvent = "1"
	SecurityTradingEvent_TRADING_RESUMES                     SecurityTradingEvent = "2"
	SecurityTradingEvent_PRICE_VOLATILITY_INTERRUPTION       SecurityTradingEvent = "3"
	SecurityTradingEvent_CHANGE_OF_TRADING_SESSION           SecurityTradingEvent = "4"
	SecurityTradingEvent_CHANGE_OF_TRADING_SUBSESSION        SecurityTradingEvent = "5"
	SecurityTradingEvent_CHANGE_OF_SECURITY_TRADING_STATUS   SecurityTradingEvent = "6"
	SecurityTradingEvent_CHANGE_OF_BOOK_TYPE                 SecurityTradingEvent = "7"
	SecurityTradingEvent_CHANGE_OF_MARKET_DEPTH              SecurityTradingEvent = "8"
	SecurityTradingEvent_CORPORATE_ACTION                    SecurityTradingEvent = "9"
)

// SecurityTradingStatus field enumeration values.
type SecurityTradingStatus string

const (
	SecurityTradingStatus_OPENING_DELAY                  SecurityTradingStatus = "1"
	SecurityTradingStatus_MARKET_ON_CLOSE_IMBALANCE_SELL SecurityTradingStatus = "10"
	SecurityTradingStatus_11                             SecurityTradingStatus = "11"
	SecurityTradingStatus_NO_MARKET_IMBALANCE            SecurityTradingStatus = "12"
	SecurityTradingStatus_NO_MARKET_ON_CLOSE_IMBALANCE   SecurityTradingStatus = "13"
	SecurityTradingStatus_ITS_PRE_OPENING                SecurityTradingStatus = "14"
	SecurityTradingStatus_NEW_PRICE_INDICATION           SecurityTradingStatus = "15"
	SecurityTradingStatus_TRADE_DISSEMINATION_TIME       SecurityTradingStatus = "16"
	SecurityTradingStatus_READY_TO_TRADE                 SecurityTradingStatus = "17"
	SecurityTradingStatus_NOT_AVAILABLE_FOR_TRADING      SecurityTradingStatus = "18"
	SecurityTradingStatus_NOT_TRADED_ON_THIS_MARKET      SecurityTradingStatus = "19"
	SecurityTradingStatus_TRADING_HALT                   SecurityTradingStatus = "2"
	SecurityTradingStatus_UNKNOWN_OR_INVALID             SecurityTradingStatus = "20"
	SecurityTradingStatus_PRE_OPEN                       SecurityTradingStatus = "21"
	SecurityTradingStatus_OPENING_ROTATION               SecurityTradingStatus = "22"
	SecurityTradingStatus_FAST_MARKET                    SecurityTradingStatus = "23"
	SecurityTradingStatus_PRE_CROSS                      SecurityTradingStatus = "24"
	SecurityTradingStatus_CROSS                          SecurityTradingStatus = "25"
	SecurityTradingStatus_POST_CLOSE                     SecurityTradingStatus = "26"
	SecurityTradingStatus_NO_CANCEL                      SecurityTradingStatus = "27"
	SecurityTradingStatus_RESUME                         SecurityTradingStatus = "3"
	SecurityTradingStatus_NO_OPEN                        SecurityTradingStatus = "4"
	SecurityTradingStatus_PRICE_INDICATION               SecurityTradingStatus = "5"
	SecurityTradingStatus_TRADING_RANGE_INDICATION       SecurityTradingStatus = "6"
	SecurityTradingStatus_MARKET_IMBALANCE_BUY           SecurityTradingStatus = "7"
	SecurityTradingStatus_MARKET_IMBALANCE_SELL          SecurityTradingStatus = "8"
	SecurityTradingStatus_MARKET_ON_CLOSE_IMBALANCE_BUY  SecurityTradingStatus = "9"
)

// SecurityType field enumeration values.
type SecurityType string

const (
	SecurityType_WILDCARD_ENTRY_FOR_USE_ON_SECURITY_DEFINITION_REQUEST     SecurityType = "?"
	SecurityType_ASSET_BACKED_SECURITIES                                   SecurityType = "ABS"
	SecurityType_AMENDED_RESTATED                                          SecurityType = "AMENDED"
	SecurityType_OTHER_ANTICIPATION_NOTES                                  SecurityType = "AN"
	SecurityType_BANKERS_ACCEPTANCE                                        SecurityType = "BA"
	SecurityType_BOND_BASKET                                               SecurityType = "BDBSKT"
	SecurityType_BANK_DEPOSITORY_NOTE                                      SecurityType = "BDN"
	SecurityType_BANK_NOTES                                                SecurityType = "BN"
	SecurityType_BILL_OF_EXCHANGES                                         SecurityType = "BOX"
	SecurityType_BRADY_BOND                                                SecurityType = "BRADY"
	SecurityType_BRIDGE_LOAN                                               SecurityType = "BRIDGE"
	SecurityType_BUY_SELLBACK                                              SecurityType = "BUYSELL"
	SecurityType_CANADIAN_MONEY_MARKETS                                    SecurityType = "CAMM"
	SecurityType_CANADIAN_TREASURY_NOTES                                   SecurityType = "CAN"
	SecurityType_CAP                                                       SecurityType = "CAP"
	SecurityType_CASH                                                      SecurityType = "CASH"
	SecurityType_CONVERTIBLE_BOND                                          SecurityType = "CB"
	SecurityType_CERTIFICATE_OF_DEPOSIT                                    SecurityType = "CD"
	SecurityType_CREDIT_DEFAULT_SWAP                                       SecurityType = "CDS"
	SecurityType_CONTRACT_FOR_DIFFERENCE                                   SecurityType = "CFD"
	SecurityType_CALL_LOANS                                                SecurityType = "CL"
	SecurityType_COLLAR                                                    SecurityType = "CLLR"
	SecurityType_CANADIAN_MORTGAGE_BONDS                                   SecurityType = "CMB"
	SecurityType_CORP_MORTGAGE_BACKED_SECURITIES                           SecurityType = "CMBS"
	SecurityType_COMMODITY_SWAP                                            SecurityType = "CMDTYSWAP"
	SecurityType_COLLATERALIZED_MORTGAGE_OBLIGATION                        SecurityType = "CMO"
	SecurityType_CERTIFICATE_OF_OBLIGATION                                 SecurityType = "COFO"
	SecurityType_CERTIFICATE_OF_PARTICIPATION                              SecurityType = "COFP"
	SecurityType_COLLATERAL_BASKET                                         SecurityType = "COLLBSKT"
	SecurityType_CORPORATE_BOND                                            SecurityType = "CORP"
	SecurityType_COMMERCIAL_PAPER                                          SecurityType = "CP"
	SecurityType_CORPORATE_PRIVATE_PLACEMENT                               SecurityType = "CPP"
	SecurityType_CORRELATION_SWAP                                          SecurityType = "CRLTNSWAP"
	SecurityType_COMMON_STOCK                                              SecurityType = "CS"
	SecurityType_CANADIAN_TREASURY_BILLS                                   SecurityType = "CTB"
	SecurityType_DEFAULTED                                                 SecurityType = "DEFLTED"
	SecurityType_DEBTOR_IN_POSSESSION                                      SecurityType = "DINP"
	SecurityType_DEPOSIT_NOTES                                             SecurityType = "DN"
	SecurityType_DEPOSITORY_RECEIPTS                                       SecurityType = "DR"
	SecurityType_DUAL_CURRENCY                                             SecurityType = "DUAL"
	SecurityType_DIVIDEND_SWAP                                             SecurityType = "DVDNDSWAP"
	SecurityType_DELIVERY_VERSUS_PLEDGE                                    SecurityType = "DVPLDG"
	SecurityType_EQUITY_BASKET                                             SecurityType = "EQBSKT"
	SecurityType_EQUITY_FORWARD                                            SecurityType = "EQFWD"
	SecurityType_EXCHANGE_TRADED_COMMODITY                                 SecurityType = "ETC"
	SecurityType_EXCHANGE_TRADED_NOTE                                      SecurityType = "ETN"
	SecurityType_EURO_CERTIFICATE_OF_DEPOSIT                               SecurityType = "EUCD"
	SecurityType_EURO_CORPORATE_BOND                                       SecurityType = "EUCORP"
	SecurityType_EURO_COMMERCIAL_PAPER                                     SecurityType = "EUCP"
	SecurityType_EURO_CORPORATE_FLOATING_RATE_NOTES                        SecurityType = "EUFRN"
	SecurityType_EURO_SOVEREIGNS                                           SecurityType = "EUSOV"
	SecurityType_EURO_SUPRANATIONAL_COUPONS                                SecurityType = "EUSUPRA"
	SecurityType_EXOTIC                                                    SecurityType = "EXOTIC"
	SecurityType_FEDERAL_AGENCY_COUPON                                     SecurityType = "FAC"
	SecurityType_FEDERAL_AGENCY_DISCOUNT_NOTE                              SecurityType = "FADN"
	SecurityType_FEDERAL_HOUSING_AUTHORITY                                 SecurityType = "FHA"
	SecurityType_FEDERAL_HOME_LOAN                                         SecurityType = "FHL"
	SecurityType_FLOOR                                                     SecurityType = "FLR"
	SecurityType_FEDERAL_NATIONAL_MORTGAGE_ASSOCIATION                     SecurityType = "FN"
	SecurityType_FOREIGN_EXCHANGE_CONTRACT                                 SecurityType = "FOR"
	SecurityType_FORWARD                                                   SecurityType = "FORWARD"
	SecurityType_FORWARD_RATE_AGREEMENT                                    SecurityType = "FRA"
	SecurityType_US_CORPORATE_FLOATING_RATE_NOTES                          SecurityType = "FRN"
	SecurityType_FUTURE                                                    SecurityType = "FUT"
	SecurityType_FUTURES_ON_A_SWAP                                         SecurityType = "FUTSWAP"
	SecurityType_DERIVATIVE_FORWARD                                        SecurityType = "FWD"
	SecurityType_FORWARD_FREIGHT_AGREEMENT                                 SecurityType = "FWDFRTAGMT"
	SecurityType_FORWARDS_ON_A_SWAP                                        SecurityType = "FWDSWAP"
	SecurityType_FX_FORWARD                                                SecurityType = "FXFWD"
	SecurityType_NON_DELIVERABLE_FORWARD                                   SecurityType = "FXNDF"
	SecurityType_NON_DELIVERABLE_SWAP                                      SecurityType = "FXNDS"
	SecurityType_FX_SPOT                                                   SecurityType = "FXSPOT"
	SecurityType_FX_SWAP                                                   SecurityType = "FXSWAP"
	SecurityType_GOVERNMENT_NATIONAL_MORTGAGE_ASSOCIATION                  SecurityType = "GN"
	SecurityType_GENERAL_OBLIGATION_BONDS                                  SecurityType = "GO"
	SecurityType_TREASURIES_PLUS_AGENCY_DEBENTURE                          SecurityType = "GOVT"
	SecurityType_IOETTE_MORTGAGE                                           SecurityType = "IET"
	SecurityType_GENERAL_TYPE_FOR_A_CONTRACT_BASED_ON_AN_ESTABLISHED_INDEX SecurityType = "INDEX"
	SecurityType_INTEREST_RATE_SWAP                                        SecurityType = "IRS"
	SecurityType_LOAN_LEASE                                                SecurityType = "LOANLEASE"
	SecurityType_LETTER_OF_CREDIT                                          SecurityType = "LOFC"
	SecurityType_LIQUIDITY_NOTE                                            SecurityType = "LQN"
	SecurityType_MATURED                                                   SecurityType = "MATURED"
	SecurityType_MORTGAGE_BACKED_SECURITIES                                SecurityType = "MBS"
	SecurityType_MUTUAL_FUND                                               SecurityType = "MF"
	SecurityType_MORTGAGE_INTEREST_ONLY                                    SecurityType = "MIO"
	SecurityType_MULTILEG_INSTRUMENT                                       SecurityType = "MLEG"
	SecurityType_MORTGAGE_PRINCIPAL_ONLY                                   SecurityType = "MPO"
	SecurityType_MORTGAGE_PRIVATE_PLACEMENT                                SecurityType = "MPP"
	SecurityType_MISCELLANEOUS_PASS_THROUGH                                SecurityType = "MPT"
	SecurityType_MARGIN_LOAN                                               SecurityType = "MRGNLOAN"
	SecurityType_MANDATORY_TENDER                                          SecurityType = "MT"
	SecurityType_MEDIUM_TERM_NOTES                                         SecurityType = "MTN"
	SecurityType_MUNICIPAL_BOND                                            SecurityType = "MUNI"
	SecurityType_NO_SECURITY_TYPE                                          SecurityType = "NONE"
	SecurityType_OVERNIGHT                                                 SecurityType = "ONITE"
	SecurityType_OPTIONS_ON_COMBO                                          SecurityType = "OOC"
	SecurityType_OPTIONS_ON_FUTURES                                        SecurityType = "OOF"
	SecurityType_OPTIONS_ON_PHYSICAL                                       SecurityType = "OOP"
	SecurityType_OPTION                                                    SecurityType = "OPT"
	SecurityType_OTHER                                                     SecurityType = "Other"
	SecurityType_PRIVATE_EXPORT_FUNDING                                    SecurityType = "PEF"
	SecurityType_PFANDBRIEFE                                               SecurityType = "PFAND"
	SecurityType_PROMISSORY_NOTE                                           SecurityType = "PN"
	SecurityType_AGENCY_POOLS                                              SecurityType = "POOL"
	SecurityType_CANADIAN_PROVINCIAL_BONDS                                 SecurityType = "PROV"
	SecurityType_PORTFOLIO_SWAP                                            SecurityType = "PRTFLIOSWAP"
	SecurityType_PREFERRED_STOCK                                           SecurityType = "PS"
	SecurityType_PLAZOS_FIJOS                                              SecurityType = "PZFJ"
	SecurityType_REVENUE_ANTICIPATION_NOTE                                 SecurityType = "RAN"
	SecurityType_REPLACED                                                  SecurityType = "REPLACD"
	SecurityType_REPURCHASE                                                SecurityType = "REPO"
	SecurityType_RETIRED                                                   SecurityType = "RETIRED"
	SecurityType_REVENUE_BONDS                                             SecurityType = "REV"
	SecurityType_REPURCHASE_AGREEMENT                                      SecurityType = "RP"
	SecurityType_RETURN_SWAP                                               SecurityType = "RTRNSWAP"
	SecurityType_REVOLVER_LOAN                                             SecurityType = "RVLV"
	SecurityType_REVOLVER_TERM_LOAN                                        SecurityType = "RVLVTRM"
	SecurityType_REVERSE_REPURCHASE_AGREEMENT                              SecurityType = "RVRP"
	SecurityType_SECURITIZED_DERIVATIVE                                    SecurityType = "SECDERIV"
	SecurityType_SECURITIES_LOAN                                           SecurityType = "SECLOAN"
	SecurityType_SECURITIES_PLEDGE                                         SecurityType = "SECPLEDGE"
	SecurityType_STRUCTURED_FINANCE_PRODUCT                                SecurityType = "SFP"
	SecurityType_STUDENT_LOAN_MARKETING_ASSOCIATION                        SecurityType = "SL"
	SecurityType_SECURED_LIQUIDITY_NOTE                                    SecurityType = "SLQN"
	SecurityType_SPECIAL_ASSESSMENT                                        SecurityType = "SPCLA"
	SecurityType_SPECIAL_OBLIGATION                                        SecurityType = "SPCLO"
	SecurityType_SPECIAL_TAX                                               SecurityType = "SPCLT"
	SecurityType_SPOT_FORWARD                                              SecurityType = "SPOTFWD"
	SecurityType_SPREAD_BETTING                                            SecurityType = "SPREADBET"
	SecurityType_SHORT_TERM_LOAN_NOTE                                      SecurityType = "STN"
	SecurityType_STRUCTURED_NOTES                                          SecurityType = "STRUCT"
	SecurityType_USD_SUPRANATIONAL_COUPONS                                 SecurityType = "SUPRA"
	SecurityType_SWAP_OPTION                                               SecurityType = "SWAPTION"
	SecurityType_SWING_LINE_FACILITY                                       SecurityType = "SWING"
	SecurityType_TAX_ANTICIPATION_NOTE                                     SecurityType = "TAN"
	SecurityType_TAX_ALLOCATION                                            SecurityType = "TAXA"
	SecurityType_TREASURY_BILL                                             SecurityType = "TB"
	SecurityType_TO_BE_ANNOUNCED                                           SecurityType = "TBA"
	SecurityType_US_TREASURY_BILL_TBILL                                    SecurityType = "TBILL"
	SecurityType_US_TREASURY_BOND                                          SecurityType = "TBOND"
	SecurityType_PRINCIPAL_STRIP_OF_A_CALLABLE_BOND_OR_NOTE                SecurityType = "TCAL"
	SecurityType_TIME_DEPOSIT                                              SecurityType = "TD"
	SecurityType_TAX_EXEMPT_COMMERCIAL_PAPER                               SecurityType = "TECP"
	SecurityType_TERM_LOAN                                                 SecurityType = "TERM"
	SecurityType_INTEREST_STRIP_FROM_ANY_BOND_OR_NOTE                      SecurityType = "TINT"
	SecurityType_TREASURY_INFLATION_PROTECTED_SECURITIES                   SecurityType = "TIPS"
	SecurityType_TERM_LIQUIDITY_NOTE                                       SecurityType = "TLQN"
	SecurityType_TAXABLE_MUNICIPAL_CP                                      SecurityType = "TMCP"
	SecurityType_US_TREASURY_NOTE_TNOTE                                    SecurityType = "TNOTE"
	SecurityType_PRINCIPAL_STRIP_FROM_A_NON_CALLABLE_BOND_OR_NOTE          SecurityType = "TPRN"
	SecurityType_TAX_REVENUE_ANTICIPATION_NOTE                             SecurityType = "TRAN"
	SecurityType_TOTAL_RETURN_SWAP                                         SecurityType = "TRS"
	SecurityType_US_TREASURY_NOTE_UST                                      SecurityType = "UST"
	SecurityType_US_TREASURY_BILL_USTB                                     SecurityType = "USTB"
	SecurityType_VARIANCE_SWAP                                             SecurityType = "VARSWAP"
	SecurityType_VARIABLE_RATE_DEMAND_NOTE                                 SecurityType = "VRDN"
	SecurityType_WARRANT                                                   SecurityType = "WAR"
	SecurityType_WITHDRAWN                                                 SecurityType = "WITHDRN"
	SecurityType_WILDCARD_ENTRY                                            SecurityType = "WLD"
	SecurityType_EXTENDED_COMM_NOTE                                        SecurityType = "XCN"
	SecurityType_INDEXED_LINKED                                            SecurityType = "XLINKD"
	SecurityType_TRANSMISSION                                              SecurityType = "XMISSION"
	SecurityType_YANKEE_CORPORATE_BOND                                     SecurityType = "YANK"
	SecurityType_YANKEE_CERTIFICATE_OF_DEPOSIT                             SecurityType = "YCD"
	SecurityType_CATS_TIGERS_LIONS                                         SecurityType = "ZOO"
)

// SecurityUpdateAction field enumeration values.
type SecurityUpdateAction string

const (
	SecurityUpdateAction_ADD    SecurityUpdateAction = "A"
	SecurityUpdateAction_DELETE SecurityUpdateAction = "D"
	SecurityUpdateAction_MODIFY SecurityUpdateAction = "M"
)

// Seniority field enumeration values.
type Seniority string

const (
	Seniority_JUNIOR               Seniority = "JR"
	Seniority_MEZZANINE            Seniority = "MZ"
	Seniority_SUBORDINATED         Seniority = "SB"
	Seniority_SENIOR_SECURED       Seniority = "SD"
	Seniority_SENIOR_NON_PREFERRED Seniority = "SN"
	Seniority_SENIOR               Seniority = "SR"
)

// SessionRejectReason field enumeration values.
type SessionRejectReason string

const (
	SessionRejectReason_INVALID_TAG_NUMBER                             SessionRejectReason = "0"
	SessionRejectReason_REQUIRED_TAG_MISSING                           SessionRejectReason = "1"
	SessionRejectReason_SENDINGTIME_ACCURACY_PROBLEM                   SessionRejectReason = "10"
	SessionRejectReason_INVALID_MSGTYPE                                SessionRejectReason = "11"
	SessionRejectReason_XML_VALIDATION_ERROR                           SessionRejectReason = "12"
	SessionRejectReason_TAG_APPEARS_MORE_THAN_ONCE                     SessionRejectReason = "13"
	SessionRejectReason_TAG_SPECIFIED_OUT_OF_REQUIRED_ORDER            SessionRejectReason = "14"
	SessionRejectReason_REPEATING_GROUP_FIELDS_OUT_OF_ORDER            SessionRejectReason = "15"
	SessionRejectReason_INCORRECT_NUMINGROUP_COUNT_FOR_REPEATING_GROUP SessionRejectReason = "16"
	SessionRejectReason_NON_DATA_VALUE_INCLUDES_FIELD_DELIMITER        SessionRejectReason = "17"
	SessionRejectReason_INVALID_UNSUPPORTED_APPLICATION_VERSION        SessionRejectReason = "18"
	SessionRejectReason_TAG_NOT_DEFINED_FOR_THIS_MESSAGE_TYPE          SessionRejectReason = "2"
	SessionRejectReason_UNDEFINED_TAG                                  SessionRejectReason = "3"
	SessionRejectReason_TAG_SPECIFIED_WITHOUT_A_VALUE                  SessionRejectReason = "4"
	SessionRejectReason_VALUE_IS_INCORRECT                             SessionRejectReason = "5"
	SessionRejectReason_INCORRECT_DATA_FORMAT_FOR_VALUE                SessionRejectReason = "6"
	SessionRejectReason_DECRYPTION_PROBLEM                             SessionRejectReason = "7"
	SessionRejectReason_SIGNATURE_PROBLEM                              SessionRejectReason = "8"
	SessionRejectReason_COMPID_PROBLEM                                 SessionRejectReason = "9"
	SessionRejectReason_OTHER                                          SessionRejectReason = "99"
)

// SessionStatus field enumeration values.
type SessionStatus string

const (
	SessionStatus_SESSION_ACTIVE                                   SessionStatus = "0"
	SessionStatus_SESSION_PASSWORD_CHANGED                         SessionStatus = "1"
	SessionStatus_RECEIVED_NEXTEXPECTEDMSGSEQNUM                   SessionStatus = "10"
	SessionStatus_SESSION_PASSWORD_DUE_TO_EXPIRE                   SessionStatus = "2"
	SessionStatus_NEW_SESSION_PASSWORD_DOES_NOT_COMPLY_WITH_POLICY SessionStatus = "3"
	SessionStatus_SESSION_LOGOUT_COMPLETE                          SessionStatus = "4"
	SessionStatus_INVALID_USERNAME_OR_PASSWORD                     SessionStatus = "5"
	SessionStatus_ACCOUNT_LOCKED                                   SessionStatus = "6"
	SessionStatus_LOGONS_ARE_NOT_ALLOWED_AT_THIS_TIME              SessionStatus = "7"
	SessionStatus_PASSWORD_EXPIRED                                 SessionStatus = "8"
	SessionStatus_RECEIVED_MSGSEQNUM                               SessionStatus = "9"
)

// SettlCurrFxRateCalc field enumeration values.
type SettlCurrFxRateCalc string

const (
	SettlCurrFxRateCalc_DIVIDE   SettlCurrFxRateCalc = "D"
	SettlCurrFxRateCalc_MULTIPLY SettlCurrFxRateCalc = "M"
)

// SettlDeliveryType field enumeration values.
type SettlDeliveryType string

const (
	SettlDeliveryType_VERSUS_PAYMENT_DELIVER SettlDeliveryType = "0"
	SettlDeliveryType_FREE_DELIVER           SettlDeliveryType = "1"
	SettlDeliveryType_TRI_PARTY              SettlDeliveryType = "2"
	SettlDeliveryType_HOLD_IN_CUSTODY        SettlDeliveryType = "3"
)

// SettlDisruptionProvision field enumeration values.
type SettlDisruptionProvision string

const (
	SettlDisruptionProvision_NEGOTIATION              SettlDisruptionProvision = "1"
	SettlDisruptionProvision_CANCELLATION_AND_PAYMENT SettlDisruptionProvision = "2"
)

// SettlInstMode field enumeration values.
type SettlInstMode string

const (
	SettlInstMode_DEFAULT                                SettlInstMode = "0"
	SettlInstMode_STANDING_INSTRUCTIONS_PROVIDED         SettlInstMode = "1"
	SettlInstMode_SPECIFIC_ALLOCATION_ACCOUNT_OVERRIDING SettlInstMode = "2"
	SettlInstMode_SPECIFIC_ALLOCATION_ACCOUNT_STANDING   SettlInstMode = "3"
	SettlInstMode_SPECIFIC_ORDER_FOR_A_SINGLE_ACCOUNT    SettlInstMode = "4"
	SettlInstMode_REQUEST_REJECT                         SettlInstMode = "5"
)

// SettlInstReqRejCode field enumeration values.
type SettlInstReqRejCode string

const (
	SettlInstReqRejCode_UNABLE_TO_PROCESS_REQUEST                 SettlInstReqRejCode = "0"
	SettlInstReqRejCode_UNKNOWN_ACCOUNT                           SettlInstReqRejCode = "1"
	SettlInstReqRejCode_NO_MATCHING_SETTLEMENT_INSTRUCTIONS_FOUND SettlInstReqRejCode = "2"
	SettlInstReqRejCode_OTHER                                     SettlInstReqRejCode = "99"
)

// SettlInstSource field enumeration values.
type SettlInstSource string

const (
	SettlInstSource_BROKERS_INSTRUCTIONS      SettlInstSource = "1"
	SettlInstSource_INSTITUTIONS_INSTRUCTIONS SettlInstSource = "2"
	SettlInstSource_INVESTOR                  SettlInstSource = "3"
)

// SettlInstTransType field enumeration values.
type SettlInstTransType string

const (
	SettlInstTransType_CANCEL  SettlInstTransType = "C"
	SettlInstTransType_NEW     SettlInstTransType = "N"
	SettlInstTransType_REPLACE SettlInstTransType = "R"
	SettlInstTransType_RESTATE SettlInstTransType = "T"
)

// SettlLocation field enumeration values.
type SettlLocation string

const (
	SettlLocation_CEDEL                        SettlLocation = "CED"
	SettlLocation_DEPOSITORY_TRUST_COMPANY     SettlLocation = "DTC"
	SettlLocation_EURO_CLEAR                   SettlLocation = "EUR"
	SettlLocation_FEDERAL_BOOK_ENTRY           SettlLocation = "FED"
	SettlLocation_LOCAL_MARKET_SETTLE_LOCATION SettlLocation = "ISO_Country_Code"
	SettlLocation_PHYSICAL                     SettlLocation = "PNY"
	SettlLocation_PARTICIPANT_TRUST_COMPANY    SettlLocation = "PTC"
)

// SettlMethod field enumeration values.
type SettlMethod string

const (
	SettlMethod_CASH_SETTLEMENT_REQUIRED     SettlMethod = "C"
	SettlMethod_ELECTION_AT_EXERCISE         SettlMethod = "E"
	SettlMethod_PHYSICAL_SETTLEMENT_REQUIRED SettlMethod = "P"
)

// SettlObligMode field enumeration values.
type SettlObligMode string

const (
	SettlObligMode_PRELIMINARY SettlObligMode = "1"
	SettlObligMode_FINAL       SettlObligMode = "2"
)

// SettlObligSource field enumeration values.
type SettlObligSource string

const (
	SettlObligSource_INSTRUCTIONS_OF_BROKER          SettlObligSource = "1"
	SettlObligSource_INSTRUCTIONS_FOR_INSTITUTION    SettlObligSource = "2"
	SettlObligSource_INVESTOR                        SettlObligSource = "3"
	SettlObligSource_BUYERS_SETTLEMENT_INSTRUCTIONS  SettlObligSource = "4"
	SettlObligSource_SELLERS_SETTLEMENT_INSTRUCTIONS SettlObligSource = "5"
)

// SettlObligTransType field enumeration values.
type SettlObligTransType string

const (
	SettlObligTransType_CANCEL  SettlObligTransType = "C"
	SettlObligTransType_NEW     SettlObligTransType = "N"
	SettlObligTransType_REPLACE SettlObligTransType = "R"
	SettlObligTransType_RESTATE SettlObligTransType = "T"
)

// SettlPriceType field enumeration values.
type SettlPriceType string

const (
	SettlPriceType_FINAL       SettlPriceType = "1"
	SettlPriceType_THEORETICAL SettlPriceType = "2"
)

// SettlSessID field enumeration values.
type SettlSessID string

const (
	SettlSessID_END_OF_DAY               SettlSessID = "EOD"
	SettlSessID_ELECTRONIC_TRADING_HOURS SettlSessID = "ETH"
	SettlSessID_INTRADAY                 SettlSessID = "ITD"
	SettlSessID_REGULAR_TRADING_HOURS    SettlSessID = "RTH"
)

// SettlSubMethod field enumeration values.
type SettlSubMethod string

const (
	SettlSubMethod_SHARES             SettlSubMethod = "1"
	SettlSubMethod_DERIVATIVES        SettlSubMethod = "2"
	SettlSubMethod_PAYMENT_VS_PAYMENT SettlSubMethod = "3"
	SettlSubMethod_NOTIONAL           SettlSubMethod = "4"
	SettlSubMethod_CASCADE            SettlSubMethod = "5"
	SettlSubMethod_REPURCHASE         SettlSubMethod = "6"
	SettlSubMethod_OTHER              SettlSubMethod = "99"
)

// SettlType field enumeration values.
type SettlType string

const (
	SettlType_REGULAR                 SettlType = "0"
	SettlType_CASH                    SettlType = "1"
	SettlType_NEXT_DAY                SettlType = "2"
	SettlType_T_PLUS_2                SettlType = "3"
	SettlType_T_PLUS_3                SettlType = "4"
	SettlType_T_PLUS_4                SettlType = "5"
	SettlType_FUTURE                  SettlType = "6"
	SettlType_WHEN_AND_IF_ISSUED      SettlType = "7"
	SettlType_SELLERS_OPTION          SettlType = "8"
	SettlType_T_PLUS_5                SettlType = "9"
	SettlType_BROKEN_DATE             SettlType = "B"
	SettlType_FX_SPOT_NEXT_SETTLEMENT SettlType = "C"
)

// SettlmntTyp field enumeration values.
type SettlmntTyp string

const (
	SettlmntTyp_REGULAR            SettlmntTyp = "0"
	SettlmntTyp_CASH               SettlmntTyp = "1"
	SettlmntTyp_NEXT_DAY           SettlmntTyp = "2"
	SettlmntTyp_T_PLUS_2           SettlmntTyp = "3"
	SettlmntTyp_T_PLUS_3           SettlmntTyp = "4"
	SettlmntTyp_T_PLUS_4           SettlmntTyp = "5"
	SettlmntTyp_FUTURE             SettlmntTyp = "6"
	SettlmntTyp_WHEN_AND_IF_ISSUED SettlmntTyp = "7"
	SettlmntTyp_SELLERS_OPTION     SettlmntTyp = "8"
	SettlmntTyp_T_PLUS_5           SettlmntTyp = "9"
	SettlmntTyp_T_PLUS_1           SettlmntTyp = "A"
)

// ShortSaleExemptionReason field enumeration values.
type ShortSaleExemptionReason string

const (
	ShortSaleExemptionReason_EXEMPTION_REASON_UNKNOWN              ShortSaleExemptionReason = "0"
	ShortSaleExemptionReason_INCOME_SELL_SHORT_EXEMPT              ShortSaleExemptionReason = "1"
	ShortSaleExemptionReason_ABOVE_NATIONAL_BEST_BID               ShortSaleExemptionReason = "2"
	ShortSaleExemptionReason_DELAYED_DELIVERY                      ShortSaleExemptionReason = "3"
	ShortSaleExemptionReason_ODD_LOT                               ShortSaleExemptionReason = "4"
	ShortSaleExemptionReason_DOMESTIC_ARBITRAGE                    ShortSaleExemptionReason = "5"
	ShortSaleExemptionReason_INTERNATIONAL_ARBITRAGE               ShortSaleExemptionReason = "6"
	ShortSaleExemptionReason_UNDERWRITER_OR_SYNDICATE_DISTRIBUTION ShortSaleExemptionReason = "7"
	ShortSaleExemptionReason_RISKLESS_PRINCIPAL                    ShortSaleExemptionReason = "8"
	ShortSaleExemptionReason_VWAP                                  ShortSaleExemptionReason = "9"
)

// ShortSaleReason field enumeration values.
type ShortSaleReason string

const (
	ShortSaleReason_DEALER_SOLD_SHORT                        ShortSaleReason = "0"
	ShortSaleReason_DEALER_SOLD_SHORT_EXEMPT                 ShortSaleReason = "1"
	ShortSaleReason_SELLING_CUSTOMER_SOLD_SHORT              ShortSaleReason = "2"
	ShortSaleReason_SELLING_CUSTOMER_SOLD_SHORT_EXEMPT       ShortSaleReason = "3"
	ShortSaleReason_QUALIFIED_SERVICE_REPRESENTATIVE         ShortSaleReason = "4"
	ShortSaleReason_QSR_OR_AGU_CONTRA_SIDE_SOLD_SHORT_EXEMPT ShortSaleReason = "5"
)

// ShortSaleRestriction field enumeration values.
type ShortSaleRestriction string

const (
	ShortSaleRestriction_NO_RESTRICTIONS                                 ShortSaleRestriction = "0"
	ShortSaleRestriction_SECURITY_IS_NOT_SHORTABLE                       ShortSaleRestriction = "1"
	ShortSaleRestriction_SECURITY_NOT_SHORTABLE_AT_OR_BELOW_THE_BEST_BID ShortSaleRestriction = "2"
	ShortSaleRestriction_SECURITY_IS_NOT_SHORTABLE_WITHOUT_PRE_BORROW    ShortSaleRestriction = "3"
)

// Side field enumeration values.
type Side string

const (
	Side_BUY                Side = "1"
	Side_SELL               Side = "2"
	Side_BUY_MINUS          Side = "3"
	Side_SELL_PLUS          Side = "4"
	Side_SELL_SHORT         Side = "5"
	Side_SELL_SHORT_EXEMPT  Side = "6"
	Side_UNDISCLOSED        Side = "7"
	Side_CROSS              Side = "8"
	Side_CROSS_SHORT        Side = "9"
	Side_CROSS_SHORT_EXEMPT Side = "A"
	Side_AS_DEFINED         Side = "B"
	Side_OPPOSITE           Side = "C"
	Side_SUBSCRIBE          Side = "D"
	Side_REDEEM             Side = "E"
	Side_LEND               Side = "F"
	Side_BORROW             Side = "G"
	Side_SELL_UNDISCLOSED   Side = "H"
)

// SideAvgPxIndicator field enumeration values.
type SideAvgPxIndicator string

const (
	SideAvgPxIndicator_NO_AVERAGE_PRICING                                                          SideAvgPxIndicator = "0"
	SideAvgPxIndicator_TRADE_IS_PART_OF_THE_AVERAGE_PRICE_GROUP_IDENTIFIED_BY_THE_SIDEAVGPXGROUPID SideAvgPxIndicator = "1"
	SideAvgPxIndicator_LAST_TRADE_IS_THE_AVERAGE_PRICE_GROUP_IDENTIFIED_BY_THE_SIDEAVGPXGROUPID    SideAvgPxIndicator = "2"
)

// SideClearingTradePriceType field enumeration values.
type SideClearingTradePriceType string

const (
	SideClearingTradePriceType_TRADE_CLEARING_AT_EXECUTION_PRICE          SideClearingTradePriceType = "0"
	SideClearingTradePriceType_TRADE_CLEARING_AT_ALTERNATE_CLEARING_PRICE SideClearingTradePriceType = "1"
)

// SideMultiLegReportingType field enumeration values.
type SideMultiLegReportingType string

const (
	SideMultiLegReportingType_SINGLE_SECURITY                       SideMultiLegReportingType = "1"
	SideMultiLegReportingType_INDIVIDUAL_LEG_OF_A_MULTILEG_SECURITY SideMultiLegReportingType = "2"
	SideMultiLegReportingType_MULTILEG_SECURITY                     SideMultiLegReportingType = "3"
)

// SideTrdSubTyp field enumeration values.
type SideTrdSubTyp string

const (
	SideTrdSubTyp_CMTA                                            SideTrdSubTyp = "0"
	SideTrdSubTyp_INTERNAL_TRANSFER                               SideTrdSubTyp = "1"
	SideTrdSubTyp_TRANSACTION_FROM_ASSIGNMENT                     SideTrdSubTyp = "10"
	SideTrdSubTyp_EXTERNAL_TRANSFER                               SideTrdSubTyp = "2"
	SideTrdSubTyp_REJECT_FOR_SUBMITTING_TRADE                     SideTrdSubTyp = "3"
	SideTrdSubTyp_ADVISORY_FOR_CONTRA_SIDE                        SideTrdSubTyp = "4"
	SideTrdSubTyp_OFFSET_DUE_TO_AN_ALLOCATION                     SideTrdSubTyp = "5"
	SideTrdSubTyp_ONSET_DUE_TO_AN_ALLOCATION                      SideTrdSubTyp = "6"
	SideTrdSubTyp_DIFFERENTIAL_SPREAD                             SideTrdSubTyp = "7"
	SideTrdSubTyp_IMPLIED_SPREAD_LEG_EXECUTED_AGAINST_AN_OUTRIGHT SideTrdSubTyp = "8"
	SideTrdSubTyp_TRANSACTION_FROM_EXERCISE                       SideTrdSubTyp = "9"
)

// SideValueInd field enumeration values.
type SideValueInd string

const (
	SideValueInd_SIDE_VALUE_1 SideValueInd = "1"
	SideValueInd_SIDE_VALUE_2 SideValueInd = "2"
)

// SingleQuoteIndicator field enumeration values.
type SingleQuoteIndicator string

const (
	SingleQuoteIndicator_NO  SingleQuoteIndicator = "N"
	SingleQuoteIndicator_YES SingleQuoteIndicator = "Y"
)

// SolicitedFlag field enumeration values.
type SolicitedFlag string

const (
	SolicitedFlag_NO  SolicitedFlag = "N"
	SolicitedFlag_YES SolicitedFlag = "Y"
)

// StandInstDbType field enumeration values.
type StandInstDbType string

const (
	StandInstDbType_OTHER              StandInstDbType = "0"
	StandInstDbType_DTC_SID            StandInstDbType = "1"
	StandInstDbType_THOMSON_ALERT      StandInstDbType = "2"
	StandInstDbType_A_GLOBAL_CUSTODIAN StandInstDbType = "3"
	StandInstDbType_ACCOUNTNET         StandInstDbType = "4"
)

// StatsType field enumeration values.
type StatsType string

const (
	StatsType_EXCHANGE_LAST StatsType = "1"
	StatsType_HIGH          StatsType = "2"
	StatsType_AVERAGE_PRICE StatsType = "3"
	StatsType_TURNOVER      StatsType = "4"
)

// StatusValue field enumeration values.
type StatusValue string

const (
	StatusValue_CONNECTED       StatusValue = "1"
	StatusValue_NOT_CONNECTED_2 StatusValue = "2"
	StatusValue_NOT_CONNECTED_3 StatusValue = "3"
	StatusValue_IN_PROCESS      StatusValue = "4"
)

// StipulationType field enumeration values.
type StipulationType string

const (
	StipulationType_ABSOLUTE_PREPAYMENT_SPEED                          StipulationType = "ABS"
	StipulationType_ADDITIONAL_TERM                                    StipulationType = "ADDTRM"
	StipulationType_ALL_GUARANTEES                                     StipulationType = "ALLGUARANTEES"
	StipulationType_ALTERNATIVE_MINIMUM_TAX                            StipulationType = "AMT"
	StipulationType_AUTO_REINVESTMENT_AT_RATE_OR_BETTER                StipulationType = "AUTOREINV"
	StipulationType_AVAILABLE_OFFER_QUANTITY_TO_BE_SHOWN_TO_THE_STREET StipulationType = "AVAILQTY"
	StipulationType_AVERAGE_FICO_SCORE                                 StipulationType = "AVFICO"
	StipulationType_AVERAGE_LOAN_SIZE                                  StipulationType = "AVSIZE"
	StipulationType_BANK_QUALIFIED                                     StipulationType = "BANKQUAL"
	StipulationType_BARGAIN_CONDITIONS                                 StipulationType = "BGNCON"
	StipulationType_BROKERS_SALES_CREDIT                               StipulationType = "BROKERCREDIT"
	StipulationType_COMPONENT_SECURITY_FALLBACK                        StipulationType = "COMPSECFLLBCK"
	StipulationType_COUPON_RANGE                                       StipulationType = "COUPON"
	StipulationType_CONSTANT_PREPAYMENT_PENALTY                        StipulationType = "CPP"
	StipulationType_CONSTANT_PREPAYMENT_RATE                           StipulationType = "CPR"
	StipulationType_CONSTANT_PREPAYMENT_YIELD                          StipulationType = "CPY"
	StipulationType_ISO_CURRENCY_CODE                                  StipulationType = "CURRENCY"
	StipulationType_CUSTOM_START_END_DATE                              StipulationType = "CUSTOMDATE"
	StipulationType_DISCOUNT_RATE                                      StipulationType = "DISCOUNT"
	StipulationType_GEOGRAPHICS_AND_RANGE                              StipulationType = "GEOG"
	StipulationType_VALUATION_DISCOUNT                                 StipulationType = "HAIRCUT"
	StipulationType_FINAL_CPR_OF_HOME_EQUITY_PREPAYMENT_CURVE          StipulationType = "HEP"
	StipulationType_INCURRED_RECOVERY                                  StipulationType = "INCURRCVY"
	StipulationType_INSURED                                            StipulationType = "INSURED"
	StipulationType_OFFER_PRICE_TO_BE_SHOWN_TO_INTERNAL_BROKERS        StipulationType = "INTERNALPX"
	StipulationType_OFFER_QUANTITY_TO_BE_SHOWN_TO_INTERNAL_BROKERS     StipulationType = "INTERNALQTY"
	StipulationType_YEAR_OR_YEAR_MONTH_OF_ISSUE                        StipulationType = "ISSUE"
	StipulationType_ISSUERS_TICKER                                     StipulationType = "ISSUER"
	StipulationType_ISSUE_SIZE_RANGE                                   StipulationType = "ISSUESIZE"
	StipulationType_THE_MINIMUM_RESIDUAL_OFFER_QUANTITY                StipulationType = "LEAVEQTY"
	StipulationType_LOCAL_JURISDICTION                                 StipulationType = "LOCLJRSDCTN"
	StipulationType_LOOKBACK_DAYS                                      StipulationType = "LOOKBACK"
	StipulationType_EXPLICIT_LOT_IDENTIFIER                            StipulationType = "LOT"
	StipulationType_LOT_VARIANCE                                       StipulationType = "LOTVAR"
	StipulationType_MATURITY_YEAR_AND_MONTH                            StipulationType = "MAT"
	StipulationType_MATURITY_RANGE                                     StipulationType = "MATURITY"
	StipulationType_MAXIMUM_LOAN_BALANCE                               StipulationType = "MAXBAL"
	StipulationType_MAXIMUMDENOMINATION                                StipulationType = "MAXDNOM"
	StipulationType_MAXIMUM_ORDER_SIZE                                 StipulationType = "MAXORDQTY"
	StipulationType_MAXIMUM_SUBSTITUTIONS                              StipulationType = "MAXSUBS"
	StipulationType_PERCENT_OF_MANUFACTURED_HOUSING_PREPAYMENT_CURVE   StipulationType = "MHP"
	StipulationType_MINIMUM_DENOMINATION                               StipulationType = "MINDNOM"
	StipulationType_MINIMUM_INCREMENT                                  StipulationType = "MININCR"
	StipulationType_MINIMUM_QUANTITY                                   StipulationType = "MINQTY"
	StipulationType_MODIFIED_EQUITY_DELIVERY                           StipulationType = "MODEQTYDLVY"
	StipulationType_MONTHLY_PREPAYMENT_RATE                            StipulationType = "MPR"
	StipulationType_MULTIPLE_EXCHANGE_FALLBACK                         StipulationType = "MULTEXCHFLLBCK"
	StipulationType_NO_REFERENCE_OBLIGATION                            StipulationType = "NOREFOBLIG"
	StipulationType_ORDER_QUANTITY_INCREMENT                           StipulationType = "ORDRINCR"
	StipulationType_ORIGINAL_AMOUNT                                    StipulationType = "ORIGAMT"
	StipulationType_PAYMENT_FREQUENCY_CALENDAR                         StipulationType = "PAYFREQ"
	StipulationType_INTEREST_PAYOFF_OF_ROLLING_OR_AMENDING_TRADE       StipulationType = "PAYOFF"
	StipulationType_NUMBER_OF_PIECES                                   StipulationType = "PIECES"
	StipulationType_POOLS_MAXIMUM                                      StipulationType = "PMAX"
	StipulationType_POOLSMINIMUM                                       StipulationType = "PMIN"
	StipulationType_POOL_IDENTIFIER                                    StipulationType = "POOL"
	StipulationType_POOL_EFFECTIVE_DATE                                StipulationType = "POOLEFFDT"
	StipulationType_POOL_INITIAL_FACTOR                                StipulationType = "POOLINITFCTR"
	StipulationType_PERCENT_OF_PROSPECTUS_PREPAYMENT_CURVE             StipulationType = "PPC"
	StipulationType_POOLS_PER_LOT                                      StipulationType = "PPL"
	StipulationType_POOLS_PER_MILLION                                  StipulationType = "PPM"
	StipulationType_POOLS_PER_TRADE                                    StipulationType = "PPT"
	StipulationType_PRICE_RANGE                                        StipulationType = "PRICE"
	StipulationType_PRICING_FREQUENCY                                  StipulationType = "PRICEFREQ"
	StipulationType_PRIMARY_OR_SECONDARY_MARKET_INDICATOR              StipulationType = "PRIMARY"
	StipulationType_PRODUCTION_YEAR                                    StipulationType = "PROD"
	StipulationType_CALL_PROTECTION                                    StipulationType = "PROTECT"
	StipulationType_PERCENT_OF_BMA_PREPAYMENT_CURVE                    StipulationType = "PSA"
	StipulationType_PURPOSE                                            StipulationType = "PURPOSE"
	StipulationType_BENCHMARK_PRICE_SOURCE                             StipulationType = "PXSOURCE"
	StipulationType_RATING_SOURCE_AND_RANGE                            StipulationType = "RATING"
	StipulationType_TYPE_OF_REDEMPTION                                 StipulationType = "REDEMPTION"
	StipulationType_INTEREST_OF_ROLLING_OR_CLOSING_TRADE               StipulationType = "REFINT"
	StipulationType_REFERENCE_POLICY                                   StipulationType = "REFPOLICY"
	StipulationType_PRINCIPAL_TO_ROLLING_OR_CLOSING_TRADE              StipulationType = "REFPRIN"
	StipulationType_REFERENCE_PRICE                                    StipulationType = "REFPX"
	StipulationType_REFERENCE_TO_ROLLING_OR_CLOSING_TRADE              StipulationType = "REFTRADE"
	StipulationType_RELEVANT_JURISDICTION                              StipulationType = "RELVJRSDCTN"
	StipulationType_RESTRICTED                                         StipulationType = "RESTRICTED"
	StipulationType_TYPE_OF_ROLL_TRADE                                 StipulationType = "ROLLTYPE"
	StipulationType_BROKER_SALES_CREDIT_OVERRIDE                       StipulationType = "SALESCREDITOVR"
	StipulationType_SECURED_LIST                                       StipulationType = "SECRDLIST"
	StipulationType_MARKET_SECTOR                                      StipulationType = "SECTOR"
	StipulationType_SECURITY_TYPE_INCLUDED_OR_EXCLUDED                 StipulationType = "SECTYPE"
	StipulationType_SINGLE_MONTHLY_MORTALITY                           StipulationType = "SMM"
	StipulationType_STRUCTURE                                          StipulationType = "STRUCT"
	StipulationType_SUBSTITUTIONS_FREQUENCY                            StipulationType = "SUBSFREQ"
	StipulationType_SUBSTITUTIONS_LEFT                                 StipulationType = "SUBSLEFT"
	StipulationType_SUBSTITUTION                                       StipulationType = "SUBSTITUTION"
	StipulationType_FREEFORM_TEXT                                      StipulationType = "TEXT"
	StipulationType_TRADERS_CREDIT                                     StipulationType = "TRADERCREDIT"
	StipulationType_TRANCHE_IDENTIFIER                                 StipulationType = "TRANCHE"
	StipulationType_TRADE_VARIANCE                                     StipulationType = "TRDVAR"
	StipulationType_UNKNOWN_REFERENCE_OBLIGATION                       StipulationType = "UNKREFOBLIG"
	StipulationType_WEIGHTED_AVERAGE_COUPON                            StipulationType = "WAC"
	StipulationType_WEIGHTED_AVERAGE_LIFE_COUPON                       StipulationType = "WAL"
	StipulationType_WEIGHTED_AVERAGE_LOAN_AGE                          StipulationType = "WALA"
	StipulationType_WEIGHTED_AVERAGE_MATURITY                          StipulationType = "WAM"
	StipulationType_WHOLE_POOL                                         StipulationType = "WHOLE"
	StipulationType_YIELD_RANGE                                        StipulationType = "YIELD"
	StipulationType_YIELD_TO_MATURITY                                  StipulationType = "YTM"
)

// StrategyParameterType field enumeration values.
type StrategyParameterType string

const (
	StrategyParameterType_INT                 StrategyParameterType = "1"
	StrategyParameterType_AMT                 StrategyParameterType = "10"
	StrategyParameterType_PERCENTAGE          StrategyParameterType = "11"
	StrategyParameterType_CHAR                StrategyParameterType = "12"
	StrategyParameterType_BOOLEAN             StrategyParameterType = "13"
	StrategyParameterType_STRING              StrategyParameterType = "14"
	StrategyParameterType_MULTIPLECHARVALUE   StrategyParameterType = "15"
	StrategyParameterType_CURRENCY            StrategyParameterType = "16"
	StrategyParameterType_EXCHANGE            StrategyParameterType = "17"
	StrategyParameterType_MONTHYEAR           StrategyParameterType = "18"
	StrategyParameterType_UTCTIMESTAMP        StrategyParameterType = "19"
	StrategyParameterType_LENGTH              StrategyParameterType = "2"
	StrategyParameterType_UTCTIMEONLY         StrategyParameterType = "20"
	StrategyParameterType_LOCALMKTDATE        StrategyParameterType = "21"
	StrategyParameterType_UTCDATEONLY         StrategyParameterType = "22"
	StrategyParameterType_DATA                StrategyParameterType = "23"
	StrategyParameterType_MULTIPLESTRINGVALUE StrategyParameterType = "24"
	StrategyParameterType_COUNTRY             StrategyParameterType = "25"
	StrategyParameterType_LANGUAGE            StrategyParameterType = "26"
	StrategyParameterType_TZTIMEONLY          StrategyParameterType = "27"
	StrategyParameterType_TZTIMESTAMP         StrategyParameterType = "28"
	StrategyParameterType_TENOR               StrategyParameterType = "29"
	StrategyParameterType_NUMINGROUP          StrategyParameterType = "3"
	StrategyParameterType_SEQNUM              StrategyParameterType = "4"
	StrategyParameterType_TAGNUM              StrategyParameterType = "5"
	StrategyParameterType_FLOAT               StrategyParameterType = "6"
	StrategyParameterType_QTY                 StrategyParameterType = "7"
	StrategyParameterType_PRICE               StrategyParameterType = "8"
	StrategyParameterType_PRICEOFFSET         StrategyParameterType = "9"
)

// StrategyType field enumeration values.
type StrategyType string

const (
	StrategyType_BUTTERFLY                    StrategyType = "BF"
	StrategyType_CALLABLE_INVERSIBLE_SNOWBALL StrategyType = "CISN"
	StrategyType_CONDOR                       StrategyType = "CNDR"
	StrategyType_OTHER                        StrategyType = "OTHER"
	StrategyType_STRADDLE                     StrategyType = "STD"
	StrategyType_STRANGLE                     StrategyType = "STG"
)

// StreamAsgnAckType field enumeration values.
type StreamAsgnAckType string

const (
	StreamAsgnAckType_ASSIGNMENT_ACCEPTED StreamAsgnAckType = "0"
	StreamAsgnAckType_ASSIGNMENT_REJECTED StreamAsgnAckType = "1"
)

// StreamAsgnRejReason field enumeration values.
type StreamAsgnRejReason string

const (
	StreamAsgnRejReason_UNKNOWN_CLIENT                   StreamAsgnRejReason = "0"
	StreamAsgnRejReason_EXCEEDS_MAXIMUM_SIZE             StreamAsgnRejReason = "1"
	StreamAsgnRejReason_UNKNOWN_OR_INVALID_CURRENCY_PAIR StreamAsgnRejReason = "2"
	StreamAsgnRejReason_NO_AVAILABLE_STREAM              StreamAsgnRejReason = "3"
	StreamAsgnRejReason_OTHER                            StreamAsgnRejReason = "99"
)

// StreamAsgnReqType field enumeration values.
type StreamAsgnReqType string

const (
	StreamAsgnReqType_STREAM_ASSIGNMENT_FOR_NEW_CUSTOMER      StreamAsgnReqType = "1"
	StreamAsgnReqType_STREAM_ASSIGNMENT_FOR_EXISTING_CUSTOMER StreamAsgnReqType = "2"
)

// StreamAsgnType field enumeration values.
type StreamAsgnType string

const (
	StreamAsgnType_ASSIGNMENT         StreamAsgnType = "1"
	StreamAsgnType_REJECTED           StreamAsgnType = "2"
	StreamAsgnType_TERMINATE_UNASSIGN StreamAsgnType = "3"
)

// StreamCommodityDataSourceIDType field enumeration values.
type StreamCommodityDataSourceIDType string

const (
	StreamCommodityDataSourceIDType_CITY                 StreamCommodityDataSourceIDType = "0"
	StreamCommodityDataSourceIDType_AIRPORT              StreamCommodityDataSourceIDType = "1"
	StreamCommodityDataSourceIDType_WEATHER_STATION_WBAN StreamCommodityDataSourceIDType = "2"
	StreamCommodityDataSourceIDType_WEATHER_INDEX_WMO    StreamCommodityDataSourceIDType = "3"
)

// StreamCommodityNearbySettlDayUnit field enumeration values.
type StreamCommodityNearbySettlDayUnit string

const (
	StreamCommodityNearbySettlDayUnit_MONTH StreamCommodityNearbySettlDayUnit = "Mo"
	StreamCommodityNearbySettlDayUnit_WEEK  StreamCommodityNearbySettlDayUnit = "Wk"
)

// StreamCommoditySettlDateRollUnit field enumeration values.
type StreamCommoditySettlDateRollUnit string

const (
	StreamCommoditySettlDateRollUnit_DAY StreamCommoditySettlDateRollUnit = "D"
)

// StreamNotionalAdjustments field enumeration values.
type StreamNotionalAdjustments string

const (
	StreamNotionalAdjustments_EXECUTION             StreamNotionalAdjustments = "0"
	StreamNotionalAdjustments_PORTFOLIO_REBALANCING StreamNotionalAdjustments = "1"
	StreamNotionalAdjustments_STANDARD              StreamNotionalAdjustments = "2"
)

// StreamNotionalCommodityFrequency field enumeration values.
type StreamNotionalCommodityFrequency string

const (
	StreamNotionalCommodityFrequency_TERM                   StreamNotionalCommodityFrequency = "0"
	StreamNotionalCommodityFrequency_PER_BUSINESS_DAY       StreamNotionalCommodityFrequency = "1"
	StreamNotionalCommodityFrequency_PER_CALCULATION_PERIOD StreamNotionalCommodityFrequency = "2"
	StreamNotionalCommodityFrequency_PER_SETTLEMENT_PERIOD  StreamNotionalCommodityFrequency = "3"
	StreamNotionalCommodityFrequency_PER_CALENDAR_DAY       StreamNotionalCommodityFrequency = "4"
	StreamNotionalCommodityFrequency_PER_HOUR               StreamNotionalCommodityFrequency = "5"
	StreamNotionalCommodityFrequency_PER_MONTH              StreamNotionalCommodityFrequency = "6"
)

// StreamType field enumeration values.
type StreamType string

const (
	StreamType_PAYMENT           StreamType = "0"
	StreamType_PHYSICAL_DELIVERY StreamType = "1"
)

// StrikeIndexQuote field enumeration values.
type StrikeIndexQuote string

const (
	StrikeIndexQuote_BID   StrikeIndexQuote = "0"
	StrikeIndexQuote_MID   StrikeIndexQuote = "1"
	StrikeIndexQuote_OFFER StrikeIndexQuote = "2"
)

// StrikePriceBoundaryMethod field enumeration values.
type StrikePriceBoundaryMethod string

const (
	StrikePriceBoundaryMethod_LESS_THAN_UNDERLYING_PRICE_IS_IN_THE_MONEY                 StrikePriceBoundaryMethod = "1"
	StrikePriceBoundaryMethod_LESS_THAN_OR_EQUAL_TO_THE_UNDERLYING_PRICE_IS_IN_THE_MONEY StrikePriceBoundaryMethod = "2"
	StrikePriceBoundaryMethod_EQUAL_TO_THE_UNDERLYING_PRICE_IS_IN_THE_MONEY              StrikePriceBoundaryMethod = "3"
	StrikePriceBoundaryMethod_GREATER_THAN_OR_EQUAL_TO_UNDERLYING_PRICE_IS_IN_THE_MONEY  StrikePriceBoundaryMethod = "4"
	StrikePriceBoundaryMethod_GREATER_THAN_UNDERLYING_IS_IN_THE_MONEY                    StrikePriceBoundaryMethod = "5"
)

// StrikePriceDeterminationMethod field enumeration values.
type StrikePriceDeterminationMethod string

const (
	StrikePriceDeterminationMethod_FIXED_STRIKE                                                                       StrikePriceDeterminationMethod = "1"
	StrikePriceDeterminationMethod_STRIKE_SET_AT_EXPIRATION_TO_UNDERLYING_OR_OTHER_VALUE                              StrikePriceDeterminationMethod = "2"
	StrikePriceDeterminationMethod_STRIKE_SET_TO_AVERAGE_OF_UNDERLYING_SETTLEMENT_PRICE_ACROSS_THE_LIFE_OF_THE_OPTION StrikePriceDeterminationMethod = "3"
	StrikePriceDeterminationMethod_STRIKE_SET_TO_OPTIMAL_VALUE                                                        StrikePriceDeterminationMethod = "4"
)

// SubscriptionRequestType field enumeration values.
type SubscriptionRequestType string

const (
	SubscriptionRequestType_SNAPSHOT                                      SubscriptionRequestType = "0"
	SubscriptionRequestType_SNAPSHOT_PLUS_UPDATES                         SubscriptionRequestType = "1"
	SubscriptionRequestType_DISABLE_PREVIOUS_SNAPSHOT_PLUS_UPDATE_REQUEST SubscriptionRequestType = "2"
)

// SwapClass field enumeration values.
type SwapClass string

const (
	SwapClass_BROAD_BASED_SECURITY_SWAP SwapClass = "BB"
	SwapClass_BASIS_SWAP                SwapClass = "BS"
	SwapClass_INDEX_SWAP                SwapClass = "IX"
	SwapClass_BASKET_SWAP               SwapClass = "SK"
)

// SwapSubClass field enumeration values.
type SwapSubClass string

const (
	SwapSubClass_ACCRETING_NOTIONAL_SCHEDULE  SwapSubClass = "ACRT"
	SwapSubClass_AMORTIZING_NOTIONAL_SCHEDULE SwapSubClass = "AMTZ"
	SwapSubClass_CONSTANT_NOTIONAL_SCHEDULE   SwapSubClass = "CNST"
	SwapSubClass_COMPOUNDING                  SwapSubClass = "COMP"
	SwapSubClass_CUSTOM_NOTIONAL_SCHEDULE     SwapSubClass = "CUST"
)

// SymbolSfx field enumeration values.
type SymbolSfx string

const (
	SymbolSfx_EUCP_WITH_LUMP_SUM_INTEREST_RATHER_THAN_DISCOUNT_PRICE               SymbolSfx = "CD"
	SymbolSfx_WHEN_ISSUED_FOR_A_SECURITY_TO_BE_REISSUED_UNDER_AN_OLD_CUSIP_OR_ISIN SymbolSfx = "WI"
)

// TargetStrategy field enumeration values.
type TargetStrategy string

const (
	TargetStrategy_VWAP                                                          TargetStrategy = "1"
	TargetStrategy_RESERVEDANDAVAILABLEFORBILATERALLYAGREEDUPONUSERDEFINEDVALUES TargetStrategy = "1000"
	TargetStrategy_PARTICIPATE                                                   TargetStrategy = "2"
	TargetStrategy_MININIZE_MARKET_IMPACT                                        TargetStrategy = "3"
)

// TaxAdvantageType field enumeration values.
type TaxAdvantageType string

const (
	TaxAdvantageType_NONE_NOT_APPLICABLE              TaxAdvantageType = "0"
	TaxAdvantageType_MAXI_ISA                         TaxAdvantageType = "1"
	TaxAdvantageType_EMPLOYEE_10                      TaxAdvantageType = "10"
	TaxAdvantageType_EMPLOYER_11                      TaxAdvantageType = "11"
	TaxAdvantageType_EMPLOYER_12                      TaxAdvantageType = "12"
	TaxAdvantageType_NON_FUND_PROTOTYPE_IRA           TaxAdvantageType = "13"
	TaxAdvantageType_NON_FUND_QUALIFIED_PLAN          TaxAdvantageType = "14"
	TaxAdvantageType_DEFINED_CONTRIBUTION_PLAN        TaxAdvantageType = "15"
	TaxAdvantageType_INDIVIDUAL_RETIREMENT_ACCOUNT_16 TaxAdvantageType = "16"
	TaxAdvantageType_INDIVIDUAL_RETIREMENT_ACCOUNT_17 TaxAdvantageType = "17"
	TaxAdvantageType_KEOGH                            TaxAdvantageType = "18"
	TaxAdvantageType_PROFIT_SHARING_PLAN              TaxAdvantageType = "19"
	TaxAdvantageType_TESSA                            TaxAdvantageType = "2"
	TaxAdvantageType_401                              TaxAdvantageType = "20"
	TaxAdvantageType_SELF_DIRECTED_IRA                TaxAdvantageType = "21"
	TaxAdvantageType_403                              TaxAdvantageType = "22"
	TaxAdvantageType_457                              TaxAdvantageType = "23"
	TaxAdvantageType_ROTH_IRA_24                      TaxAdvantageType = "24"
	TaxAdvantageType_ROTH_IRA_25                      TaxAdvantageType = "25"
	TaxAdvantageType_ROTH_CONVERSION_IRA_26           TaxAdvantageType = "26"
	TaxAdvantageType_ROTH_CONVERSION_IRA_27           TaxAdvantageType = "27"
	TaxAdvantageType_EDUCATION_IRA_28                 TaxAdvantageType = "28"
	TaxAdvantageType_EDUCATION_IRA_29                 TaxAdvantageType = "29"
	TaxAdvantageType_MINI_CASH_ISA                    TaxAdvantageType = "3"
	TaxAdvantageType_MINI_STOCKS_AND_SHARES_ISA       TaxAdvantageType = "4"
	TaxAdvantageType_MINI_INSURANCE_ISA               TaxAdvantageType = "5"
	TaxAdvantageType_CURRENT_YEAR_PAYMENT             TaxAdvantageType = "6"
	TaxAdvantageType_PRIOR_YEAR_PAYMENT               TaxAdvantageType = "7"
	TaxAdvantageType_ASSET_TRANSFER                   TaxAdvantageType = "8"
	TaxAdvantageType_EMPLOYEE_9                       TaxAdvantageType = "9"
	TaxAdvantageType_OTHER                            TaxAdvantageType = "999"
)

// TaxonomyType field enumeration values.
type TaxonomyType string

const (
	TaxonomyType_INTERIM_TAXONOMY                                 TaxonomyType = "E"
	TaxonomyType_ISIN_OR_ALTERNATE_INSTRUMENT_IDENTIFIER_PLUS_CFI TaxonomyType = "I"
)

// TerminationType field enumeration values.
type TerminationType string

const (
	TerminationType_OVERNIGHT TerminationType = "1"
	TerminationType_TERM      TerminationType = "2"
	TerminationType_FLEXIBLE  TerminationType = "3"
	TerminationType_OPEN      TerminationType = "4"
)

// TestMessageIndicator field enumeration values.
type TestMessageIndicator string

const (
	TestMessageIndicator_NO  TestMessageIndicator = "N"
	TestMessageIndicator_YES TestMessageIndicator = "Y"
)

// ThrottleAction field enumeration values.
type ThrottleAction string

const (
	ThrottleAction_QUEUE_INBOUND  ThrottleAction = "0"
	ThrottleAction_QUEUE_OUTBOUND ThrottleAction = "1"
	ThrottleAction_REJECT         ThrottleAction = "2"
	ThrottleAction_DISCONNECT     ThrottleAction = "3"
	ThrottleAction_WARNING        ThrottleAction = "4"
)

// ThrottleCountIndicator field enumeration values.
type ThrottleCountIndicator string

const (
	ThrottleCountIndicator_OUTSTANDING_REQUESTS_UNCHANGED ThrottleCountIndicator = "0"
	ThrottleCountIndicator_OUTSTANDING_REQUESTS_DECREASED ThrottleCountIndicator = "1"
)

// ThrottleInst field enumeration values.
type ThrottleInst string

const (
	ThrottleInst_REJECT_IF_THROTTLE_LIMIT_EXCEEDED ThrottleInst = "0"
	ThrottleInst_QUEUE_IF_THROTTLE_LIMIT_EXCEEDED  ThrottleInst = "1"
)

// ThrottleStatus field enumeration values.
type ThrottleStatus string

const (
	ThrottleStatus_THROTTLE_LIMIT_NOT_EXCEEDED_NOT_QUEUED ThrottleStatus = "0"
	ThrottleStatus_QUEUED_DUE_TO_THROTTLE_LIMIT_EXCEEDED  ThrottleStatus = "1"
)

// ThrottleType field enumeration values.
type ThrottleType string

const (
	ThrottleType_INBOUND_RATE         ThrottleType = "0"
	ThrottleType_OUTSTANDING_REQUESTS ThrottleType = "1"
)

// TickDirection field enumeration values.
type TickDirection string

const (
	TickDirection_PLUS_TICK       TickDirection = "0"
	TickDirection_ZERO_PLUS_TICK  TickDirection = "1"
	TickDirection_MINUS_TICK      TickDirection = "2"
	TickDirection_ZERO_MINUS_TICK TickDirection = "3"
)

// TickRuleType field enumeration values.
type TickRuleType string

const (
	TickRuleType_REGULAR_TRADING         TickRuleType = "0"
	TickRuleType_VARIABLE_CABINET        TickRuleType = "1"
	TickRuleType_FIXED_CABINET           TickRuleType = "2"
	TickRuleType_TRADED_AS_A_SPREAD_LEG  TickRuleType = "3"
	TickRuleType_SETTLED_AS_A_SPREAD_LEG TickRuleType = "4"
	TickRuleType_TRADED_AS_SPREAD        TickRuleType = "5"
)

// TimeInForce field enumeration values.
type TimeInForce string

const (
	TimeInForce_DAY                   TimeInForce = "0"
	TimeInForce_GOOD_TILL_CANCEL      TimeInForce = "1"
	TimeInForce_AT_THE_OPENING        TimeInForce = "2"
	TimeInForce_IMMEDIATE_OR_CANCEL   TimeInForce = "3"
	TimeInForce_FILL_OR_KILL          TimeInForce = "4"
	TimeInForce_GOOD_TILL_CROSSING    TimeInForce = "5"
	TimeInForce_GOOD_TILL_DATE        TimeInForce = "6"
	TimeInForce_AT_THE_CLOSE          TimeInForce = "7"
	TimeInForce_GOOD_THROUGH_CROSSING TimeInForce = "8"
	TimeInForce_AT_CROSSING           TimeInForce = "9"
	TimeInForce_GOOD_FOR_TIME         TimeInForce = "A"
	TimeInForce_GOOD_FOR_AUCTION      TimeInForce = "B"
	TimeInForce_GOOD_FOR_THIS_MONTH   TimeInForce = "C"
)

// TimeUnit field enumeration values.
type TimeUnit string

const (
	TimeUnit_DAY     TimeUnit = "D"
	TimeUnit_HOUR    TimeUnit = "H"
	TimeUnit_MINUTE  TimeUnit = "Min"
	TimeUnit_MONTH   TimeUnit = "Mo"
	TimeUnit_QUARTER TimeUnit = "Q"
	TimeUnit_SECOND  TimeUnit = "S"
	TimeUnit_WEEK    TimeUnit = "Wk"
	TimeUnit_YEAR    TimeUnit = "Yr"
)

// TradSesControl field enumeration values.
type TradSesControl string

const (
	TradSesControl_AUTOMATIC TradSesControl = "0"
	TradSesControl_MANUAL    TradSesControl = "1"
)

// TradSesEvent field enumeration values.
type TradSesEvent string

const (
	TradSesEvent_TRADING_RESUMES              TradSesEvent = "0"
	TradSesEvent_CHANGE_OF_TRADING_SESSION    TradSesEvent = "1"
	TradSesEvent_CHANGE_OF_TRADING_SUBSESSION TradSesEvent = "2"
	TradSesEvent_CHANGE_OF_TRADING_STATUS     TradSesEvent = "3"
)

// TradSesMethod field enumeration values.
type TradSesMethod string

const (
	TradSesMethod_ELECTRONIC  TradSesMethod = "1"
	TradSesMethod_OPEN_OUTCRY TradSesMethod = "2"
	TradSesMethod_TWO_PARTY   TradSesMethod = "3"
	TradSesMethod_VOICE       TradSesMethod = "4"
)

// TradSesMode field enumeration values.
type TradSesMode string

const (
	TradSesMode_TESTING    TradSesMode = "1"
	TradSesMode_SIMULATED  TradSesMode = "2"
	TradSesMode_PRODUCTION TradSesMode = "3"
)

// TradSesStatus field enumeration values.
type TradSesStatus string

const (
	TradSesStatus_UNKNOWN          TradSesStatus = "0"
	TradSesStatus_HALTED           TradSesStatus = "1"
	TradSesStatus_OPEN             TradSesStatus = "2"
	TradSesStatus_CLOSED           TradSesStatus = "3"
	TradSesStatus_PRE_OPEN         TradSesStatus = "4"
	TradSesStatus_PRE_CLOSE        TradSesStatus = "5"
	TradSesStatus_REQUEST_REJECTED TradSesStatus = "6"
)

// TradSesStatusRejReason field enumeration values.
type TradSesStatusRejReason string

const (
	TradSesStatusRejReason_UNKNOWN_OR_INVALID_TRADINGSESSIONID TradSesStatusRejReason = "1"
	TradSesStatusRejReason_OTHER                               TradSesStatusRejReason = "99"
)

// TradeAggregationRejectReason field enumeration values.
type TradeAggregationRejectReason string

const (
	TradeAggregationRejectReason_UNKNOWN_ORDER          TradeAggregationRejectReason = "0"
	TradeAggregationRejectReason_UNKNOWN_EXECUTION_FILL TradeAggregationRejectReason = "1"
	TradeAggregationRejectReason_OTHER                  TradeAggregationRejectReason = "99"
)

// TradeAggregationRequestStatus field enumeration values.
type TradeAggregationRequestStatus string

const (
	TradeAggregationRequestStatus_ACCEPTED TradeAggregationRequestStatus = "0"
	TradeAggregationRequestStatus_REJECTED TradeAggregationRequestStatus = "1"
)

// TradeAggregationTransType field enumeration values.
type TradeAggregationTransType string

const (
	TradeAggregationTransType_NEW     TradeAggregationTransType = "0"
	TradeAggregationTransType_CANCEL  TradeAggregationTransType = "1"
	TradeAggregationTransType_REPLACE TradeAggregationTransType = "2"
)

// TradeAllocGroupInstruction field enumeration values.
type TradeAllocGroupInstruction string

const (
	TradeAllocGroupInstruction_ADD_TO_AN_EXISTING_ALLOCATION_GROUP_IF_ONE_EXISTS TradeAllocGroupInstruction = "0"
	TradeAllocGroupInstruction_DO_NOT_ADD_THE_TRADE_TO_AN_ALLOCATION_GROUP       TradeAllocGroupInstruction = "1"
)

// TradeAllocIndicator field enumeration values.
type TradeAllocIndicator string

const (
	TradeAllocIndicator_ALLOCATION_NOT_REQUIRED                TradeAllocIndicator = "0"
	TradeAllocIndicator_ALLOCATION_REQUIRED                    TradeAllocIndicator = "1"
	TradeAllocIndicator_USE_ALLOCATION_PROVIDED_WITH_THE_TRADE TradeAllocIndicator = "2"
	TradeAllocIndicator_ALLOCATION_GIVE_UP_EXECUTOR            TradeAllocIndicator = "3"
	TradeAllocIndicator_ALLOCATION_FROM_EXECUTOR               TradeAllocIndicator = "4"
	TradeAllocIndicator_ALLOCATION_TO_CLAIM_ACCOUNT            TradeAllocIndicator = "5"
	TradeAllocIndicator_TRADE_SPLIT                            TradeAllocIndicator = "6"
)

// TradeAllocStatus field enumeration values.
type TradeAllocStatus string

const (
	TradeAllocStatus_PENDING_CLEAR TradeAllocStatus = "0"
	TradeAllocStatus_CLAIMED       TradeAllocStatus = "1"
	TradeAllocStatus_CLEARED       TradeAllocStatus = "2"
	TradeAllocStatus_REJECTED      TradeAllocStatus = "3"
)

// TradeCollateralization field enumeration values.
type TradeCollateralization string

const (
	TradeCollateralization_UNCOLLATERALIZED           TradeCollateralization = "0"
	TradeCollateralization_PARTIALLY_COLLATERALIZED   TradeCollateralization = "1"
	TradeCollateralization_ONE_WAY_COLLATERALLIZATION TradeCollateralization = "2"
	TradeCollateralization_FULLY_COLLATERALIZED       TradeCollateralization = "3"
	TradeCollateralization_NET_EXPOSURE               TradeCollateralization = "4"
)

// TradeCondition field enumeration values.
type TradeCondition string

const (
	TradeCondition_CANCEL                                  TradeCondition = "0"
	TradeCondition_IMPLIED_TRADE                           TradeCondition = "1"
	TradeCondition_MARKETPLACE_ENTERED_TRADE               TradeCondition = "2"
	TradeCondition_MULTI_ASSET_CLASS_MULTILEG_TRADE        TradeCondition = "3"
	TradeCondition_MULTILEG_TO_MULTILEG_TRADE              TradeCondition = "4"
	TradeCondition_SHORT_SALE_MINIMUM_PRICE                TradeCondition = "5"
	TradeCondition_BENCHMARK                               TradeCondition = "6"
	TradeCondition_CASH                                    TradeCondition = "A"
	TradeCondition_SPREAD                                  TradeCondition = "AA"
	TradeCondition_SPREAD_ETH                              TradeCondition = "AB"
	TradeCondition_STRADDLE                                TradeCondition = "AC"
	TradeCondition_STRADDLE_ETH                            TradeCondition = "AD"
	TradeCondition_STOPPED                                 TradeCondition = "AE"
	TradeCondition_STOPPED_ETH                             TradeCondition = "AF"
	TradeCondition_REGULAR_ETH                             TradeCondition = "AG"
	TradeCondition_COMBO                                   TradeCondition = "AH"
	TradeCondition_COMBO_ETH                               TradeCondition = "AI"
	TradeCondition_OFFICIAL_CLOSING_PRICE                  TradeCondition = "AJ"
	TradeCondition_PRIOR_REFERENCE_PRICE                   TradeCondition = "AK"
	TradeCondition_STOPPED_SOLD_LAST                       TradeCondition = "AL"
	TradeCondition_STOPPED_OUT_OF_SEQUENCE                 TradeCondition = "AM"
	TradeCondition_OFFICAL_CLOSING_PRICE                   TradeCondition = "AN"
	TradeCondition_CROSSED_AO                              TradeCondition = "AO"
	TradeCondition_FAST_MARKET                             TradeCondition = "AP"
	TradeCondition_AUTOMATIC_EXECUTION                     TradeCondition = "AQ"
	TradeCondition_FORM_T                                  TradeCondition = "AR"
	TradeCondition_BASKET_INDEX                            TradeCondition = "AS"
	TradeCondition_BURST_BASKET                            TradeCondition = "AT"
	TradeCondition_TRADE_THROUGH_EXEMPT                    TradeCondition = "AU"
	TradeCondition_QUOTE_SPREAD                            TradeCondition = "AV"
	TradeCondition_LAST_AUCTION_PRICE                      TradeCondition = "AW"
	TradeCondition_HIGH_PRICE                              TradeCondition = "AX"
	TradeCondition_LOW_PRICE                               TradeCondition = "AY"
	TradeCondition_SYSTEMATIC_INTERNALISER                 TradeCondition = "AZ"
	TradeCondition_AVERAGE_PRICE_TRADE                     TradeCondition = "B"
	TradeCondition_AWAY_MARKET                             TradeCondition = "BA"
	TradeCondition_MID_POINT_PRICE                         TradeCondition = "BB"
	TradeCondition_TRADED_BEFORE_ISSUE_DATE                TradeCondition = "BC"
	TradeCondition_PREVIOUS_CLOSING_PRICE                  TradeCondition = "BD"
	TradeCondition_NATIONAL_BEST_BID_AND_OFFER             TradeCondition = "BE"
	TradeCondition_CASH_TRADE                              TradeCondition = "C"
	TradeCondition_NEXT_DAY                                TradeCondition = "D"
	TradeCondition_OPENING_REOPENING_TRADE_DETAIL          TradeCondition = "E"
	TradeCondition_INTRADAY_TRADE_DETAIL                   TradeCondition = "F"
	TradeCondition_RULE_127_TRADE                          TradeCondition = "G"
	TradeCondition_RULE_155_TRADE                          TradeCondition = "H"
	TradeCondition_SOLD_LAST                               TradeCondition = "I"
	TradeCondition_NEXT_DAY_TRADE                          TradeCondition = "J"
	TradeCondition_OPENED                                  TradeCondition = "K"
	TradeCondition_SELLER                                  TradeCondition = "L"
	TradeCondition_SOLD                                    TradeCondition = "M"
	TradeCondition_STOPPED_STOCK                           TradeCondition = "N"
	TradeCondition_IMBALANCE_MORE_BUYERS                   TradeCondition = "P"
	TradeCondition_IMBALANCE_MORE_SELLERS                  TradeCondition = "Q"
	TradeCondition_OPENING_PRICE                           TradeCondition = "R"
	TradeCondition_BARGAIN_CONDITION                       TradeCondition = "S"
	TradeCondition_CONVERTED_PRICE_INDICATOR               TradeCondition = "T"
	TradeCondition_EXCHANGE_LAST                           TradeCondition = "U"
	TradeCondition_FINAL_PRICE_OF_SESSION                  TradeCondition = "V"
	TradeCondition_EX_PIT                                  TradeCondition = "W"
	TradeCondition_CROSSED_X                               TradeCondition = "X"
	TradeCondition_TRADES_RESULTING_FROM_MANUAL_SLOW_QUOTE TradeCondition = "Y"
	TradeCondition_TRADES_RESULTING_FROM_INTERMARKET_SWEEP TradeCondition = "Z"
	TradeCondition_VOLUME_ONLY                             TradeCondition = "a"
	TradeCondition_DIRECT_PLUS                             TradeCondition = "b"
	TradeCondition_ACQUISITION                             TradeCondition = "c"
	TradeCondition_BUNCHED                                 TradeCondition = "d"
	TradeCondition_DISTRIBUTION                            TradeCondition = "e"
	TradeCondition_BUNCHED_SALE                            TradeCondition = "f"
	TradeCondition_SPLIT_TRADE                             TradeCondition = "g"
	TradeCondition_CANCEL_STOPPED                          TradeCondition = "h"
	TradeCondition_CANCEL_ETH                              TradeCondition = "i"
	TradeCondition_CANCEL_STOPPED_ETH                      TradeCondition = "j"
	TradeCondition_OUT_OF_SEQUENCE_ETH                     TradeCondition = "k"
	TradeCondition_CANCEL_LAST_ETH                         TradeCondition = "l"
	TradeCondition_SOLD_LAST_SALE_ETH                      TradeCondition = "m"
	TradeCondition_CANCEL_LAST                             TradeCondition = "n"
	TradeCondition_SOLD_LAST_SALE                          TradeCondition = "o"
	TradeCondition_CANCEL_OPEN                             TradeCondition = "p"
	TradeCondition_CANCEL_OPEN_ETH                         TradeCondition = "q"
	TradeCondition_OPENED_SALE_ETH                         TradeCondition = "r"
	TradeCondition_CANCEL_ONLY                             TradeCondition = "s"
	TradeCondition_CANCEL_ONLY_ETH                         TradeCondition = "t"
	TradeCondition_LATE_OPEN_ETH                           TradeCondition = "u"
	TradeCondition_AUTO_EXECUTION_ETH                      TradeCondition = "v"
	TradeCondition_REOPEN                                  TradeCondition = "w"
	TradeCondition_REOPEN_ETH                              TradeCondition = "x"
	TradeCondition_ADJUSTED                                TradeCondition = "y"
	TradeCondition_ADJUSTED_ETH                            TradeCondition = "z"
)

// TradeContingency field enumeration values.
type TradeContingency string

const (
	TradeContingency_DOES_NOT_APPLY       TradeContingency = "0"
	TradeContingency_CONTINGENT_TRADE     TradeContingency = "1"
	TradeContingency_NON_CONTINGENT_TRADE TradeContingency = "2"
)

// TradeContinuation field enumeration values.
type TradeContinuation string

const (
	TradeContinuation_NOVATION                              TradeContinuation = "0"
	TradeContinuation_PARTIAL_NOVATION                      TradeContinuation = "1"
	TradeContinuation_CREDIT_EVENT                          TradeContinuation = "10"
	TradeContinuation_STRATEGIC_RESTRUCTURING               TradeContinuation = "11"
	TradeContinuation_SUCCESSION_EVENT_REORGANIZATION       TradeContinuation = "12"
	TradeContinuation_SUCCESSION_EVENT_RENAMING             TradeContinuation = "13"
	TradeContinuation_PORTING                               TradeContinuation = "14"
	TradeContinuation_WITHDRAWAL                            TradeContinuation = "15"
	TradeContinuation_VOID                                  TradeContinuation = "16"
	TradeContinuation_ACCOUNT_TRANSFER                      TradeContinuation = "17"
	TradeContinuation_GIVE_UP                               TradeContinuation = "18"
	TradeContinuation_TAKEUP                                TradeContinuation = "19"
	TradeContinuation_TRADE_UNWIND                          TradeContinuation = "2"
	TradeContinuation_AVERAGE_PRICING                       TradeContinuation = "20"
	TradeContinuation_REVERSAL                              TradeContinuation = "21"
	TradeContinuation_ALLOCATION_TRADE_POSTING              TradeContinuation = "22"
	TradeContinuation_CASCADE                               TradeContinuation = "23"
	TradeContinuation_DELIVERY                              TradeContinuation = "24"
	TradeContinuation_OPTION_ASSIGNMENT                     TradeContinuation = "25"
	TradeContinuation_EXPIRATION                            TradeContinuation = "26"
	TradeContinuation_MATURITY                              TradeContinuation = "27"
	TradeContinuation_EQUAL_POSITION_ADJUSTMENT             TradeContinuation = "28"
	TradeContinuation_UNEQUAL_POSITION_ADJUSTMENT           TradeContinuation = "29"
	TradeContinuation_PARTIAL_TRADE_UNWIND                  TradeContinuation = "3"
	TradeContinuation_CORRECTION                            TradeContinuation = "30"
	TradeContinuation_EARLY_TERMINATION                     TradeContinuation = "31"
	TradeContinuation_RERATE                                TradeContinuation = "32"
	TradeContinuation_EXERCISE                              TradeContinuation = "4"
	TradeContinuation_COMPRESSION_NETTING                   TradeContinuation = "5"
	TradeContinuation_FULL_NETTING                          TradeContinuation = "6"
	TradeContinuation_PARTIAL_NETTING                       TradeContinuation = "7"
	TradeContinuation_AMENDMENT                             TradeContinuation = "8"
	TradeContinuation_INCREASE                              TradeContinuation = "9"
	TradeContinuation_OTHER_PRICE_FORMING_CONTINUATION_DATA TradeContinuation = "99"
)

// TradeHandlingInstr field enumeration values.
type TradeHandlingInstr string

const (
	TradeHandlingInstr_TRADE_CONFIRMATION                  TradeHandlingInstr = "0"
	TradeHandlingInstr_TWO_PARTY_REPORT                    TradeHandlingInstr = "1"
	TradeHandlingInstr_ONE_PARTY_REPORT_FOR_MATCHING       TradeHandlingInstr = "2"
	TradeHandlingInstr_ONE_PARTY_REPORT_FOR_PASS_THROUGH   TradeHandlingInstr = "3"
	TradeHandlingInstr_AUTOMATED_FLOOR_ORDER_ROUTING       TradeHandlingInstr = "4"
	TradeHandlingInstr_TWO_PARTY_REPORT_FOR_CLAIM          TradeHandlingInstr = "5"
	TradeHandlingInstr_ONE_PARTY_REPORT                    TradeHandlingInstr = "6"
	TradeHandlingInstr_THIRD_PARTY_REPORT_FOR_PASS_THROUGH TradeHandlingInstr = "7"
	TradeHandlingInstr_ONE_PARTY_REPORT_FOR_AUTO_MATCH     TradeHandlingInstr = "8"
)

// TradeMatchAckStatus field enumeration values.
type TradeMatchAckStatus string

const (
	TradeMatchAckStatus_RECEIVED_NOT_YET_PROCESSED TradeMatchAckStatus = "0"
	TradeMatchAckStatus_ACCEPTED                   TradeMatchAckStatus = "1"
	TradeMatchAckStatus_REJECTED                   TradeMatchAckStatus = "2"
)

// TradeMatchRejectReason field enumeration values.
type TradeMatchRejectReason string

const (
	TradeMatchRejectReason_SUCCESSFUL                      TradeMatchRejectReason = "0"
	TradeMatchRejectReason_INVALID_PARTY_INFORMATION       TradeMatchRejectReason = "1"
	TradeMatchRejectReason_UNKNOWN_INSTRUMENT              TradeMatchRejectReason = "2"
	TradeMatchRejectReason_NOT_AUTHORIZED_TO_REPORT_TRADES TradeMatchRejectReason = "3"
	TradeMatchRejectReason_INVALID_TRADE_TYPE              TradeMatchRejectReason = "4"
	TradeMatchRejectReason_OTHER                           TradeMatchRejectReason = "99"
)

// TradePriceCondition field enumeration values.
type TradePriceCondition string

const (
	TradePriceCondition_SPECIAL_CUM_DIVIDEND                   TradePriceCondition = "0"
	TradePriceCondition_SPECIAL_CUM_RIGHTS                     TradePriceCondition = "1"
	TradePriceCondition_SPECIAL_PRICE                          TradePriceCondition = "10"
	TradePriceCondition_SPECIAL_EX_BONUS                       TradePriceCondition = "11"
	TradePriceCondition_GUARANTEED_DELIVERY                    TradePriceCondition = "12"
	TradePriceCondition_SPECIAL_DIVIDEND                       TradePriceCondition = "13"
	TradePriceCondition_PRICE_IMPROVEMENT                      TradePriceCondition = "14"
	TradePriceCondition_NON_PRICE_FORMING_TRADE                TradePriceCondition = "15"
	TradePriceCondition_TRADE_EXEMPTED_FROM_TRADING_OBLIGATION TradePriceCondition = "16"
	TradePriceCondition_PRICE_OR_STRIKE_PRICE_IS_PENDING       TradePriceCondition = "17"
	TradePriceCondition_PRICE_IS_NOT_APPLICABLE                TradePriceCondition = "18"
	TradePriceCondition_SPECIAL_EX_DIVIDEND                    TradePriceCondition = "2"
	TradePriceCondition_SPECIAL_EX_RIGHTS                      TradePriceCondition = "3"
	TradePriceCondition_SPECIAL_CUM_COUPON                     TradePriceCondition = "4"
	TradePriceCondition_SPECIAL_CUM_CAPITAL_REPAYMENTS         TradePriceCondition = "5"
	TradePriceCondition_SPECIAL_EX_COUPON                      TradePriceCondition = "6"
	TradePriceCondition_SPECIAL_EX_CAPITAL_REPAYMENTS          TradePriceCondition = "7"
	TradePriceCondition_CASH_SETTLEMENT                        TradePriceCondition = "8"
	TradePriceCondition_SPECIAL_CUM_BONUS                      TradePriceCondition = "9"
)

// TradePriceNegotiationMethod field enumeration values.
type TradePriceNegotiationMethod string

const (
	TradePriceNegotiationMethod_PERCENT_OF_PAR                    TradePriceNegotiationMethod = "0"
	TradePriceNegotiationMethod_DEAL_SPREAD                       TradePriceNegotiationMethod = "1"
	TradePriceNegotiationMethod_UPFRONT_POINTS                    TradePriceNegotiationMethod = "2"
	TradePriceNegotiationMethod_UPFRONT_AMOUNT                    TradePriceNegotiationMethod = "3"
	TradePriceNegotiationMethod_PERCENT_OF_PAR_AND_UPFRONT_AMOUNT TradePriceNegotiationMethod = "4"
	TradePriceNegotiationMethod_DEAL_SPREAD_AND_UPFRONT_AMOUNT    TradePriceNegotiationMethod = "5"
	TradePriceNegotiationMethod_UPFRONT_POINTS_AND_UPFRONT_AMOUNT TradePriceNegotiationMethod = "6"
)

// TradePublishIndicator field enumeration values.
type TradePublishIndicator string

const (
	TradePublishIndicator_DO_NOT_PUBLISH_TRADE TradePublishIndicator = "0"
	TradePublishIndicator_PUBLISH_TRADE        TradePublishIndicator = "1"
	TradePublishIndicator_DEFERRED_PUBLICATION TradePublishIndicator = "2"
	TradePublishIndicator_PUBLISHED            TradePublishIndicator = "3"
)

// TradeQtyType field enumeration values.
type TradeQtyType string

const (
	TradeQtyType_CLEARED_QUANTITY                  TradeQtyType = "0"
	TradeQtyType_LONG_SIDE_CLAIMED_QUANTITY        TradeQtyType = "1"
	TradeQtyType_SHORT_SIDE_CLAIMED_QUANTITY       TradeQtyType = "2"
	TradeQtyType_LONG_SIDE_REJECTED_QUANTITY       TradeQtyType = "3"
	TradeQtyType_SHORT_SIDE_REJECTED_QUANTITY      TradeQtyType = "4"
	TradeQtyType_PENDING_QUANTITY                  TradeQtyType = "5"
	TradeQtyType_TRANSACTION_QUANTITY              TradeQtyType = "6"
	TradeQtyType_REMAINING_TRADE_QUANTITY          TradeQtyType = "7"
	TradeQtyType_PREVIOUS_REMAINING_TRADE_QUANTITY TradeQtyType = "8"
)

// TradeReportRejectReason field enumeration values.
type TradeReportRejectReason string

const (
	TradeReportRejectReason_SUCCESSFUL                       TradeReportRejectReason = "0"
	TradeReportRejectReason_INVALID_PARTY_INFORMATION        TradeReportRejectReason = "1"
	TradeReportRejectReason_UNKNOWN_INSTRUMENT               TradeReportRejectReason = "2"
	TradeReportRejectReason_UNAUTHORIZED_TO_REPORT_TRADES    TradeReportRejectReason = "3"
	TradeReportRejectReason_INVALID_TRADE_TYPE               TradeReportRejectReason = "4"
	TradeReportRejectReason_PRICE_EXCEEDS_CURRENT_PRICE_BAND TradeReportRejectReason = "5"
	TradeReportRejectReason_REFERENCE_PRICE_NOT_AVAILABLE    TradeReportRejectReason = "6"
	TradeReportRejectReason_NOTIONAL_VALUE_EXCEEDS_THRESHOLD TradeReportRejectReason = "7"
	TradeReportRejectReason_OTHER                            TradeReportRejectReason = "99"
)

// TradeReportTransType field enumeration values.
type TradeReportTransType string

const (
	TradeReportTransType_NEW                             TradeReportTransType = "0"
	TradeReportTransType_CANCEL                          TradeReportTransType = "1"
	TradeReportTransType_REPLACE                         TradeReportTransType = "2"
	TradeReportTransType_RELEASE                         TradeReportTransType = "3"
	TradeReportTransType_REVERSE                         TradeReportTransType = "4"
	TradeReportTransType_CANCEL_DUE_TO_BACK_OUT_OF_TRADE TradeReportTransType = "5"
)

// TradeReportType field enumeration values.
type TradeReportType string

const (
	TradeReportType_SUBMIT                      TradeReportType = "0"
	TradeReportType_ALLEGED_1                   TradeReportType = "1"
	TradeReportType_PENDED                      TradeReportType = "10"
	TradeReportType_ALLEGED_NEW                 TradeReportType = "11"
	TradeReportType_ALLEGED_ADDENDUM            TradeReportType = "12"
	TradeReportType_ALLEGED_NO_WAS              TradeReportType = "13"
	TradeReportType_ALLEGED_TRADE_REPORT_CANCEL TradeReportType = "14"
	TradeReportType_ALLEGED_15                  TradeReportType = "15"
	TradeReportType_VERIFY                      TradeReportType = "16"
	TradeReportType_DISPUTE                     TradeReportType = "17"
	TradeReportType_NON_MATERIAL_UPDATE         TradeReportType = "18"
	TradeReportType_ACCEPT                      TradeReportType = "2"
	TradeReportType_DECLINE                     TradeReportType = "3"
	TradeReportType_ADDENDUM                    TradeReportType = "4"
	TradeReportType_NO_WAS                      TradeReportType = "5"
	TradeReportType_TRADE_REPORT_CANCEL         TradeReportType = "6"
	TradeReportType_7                           TradeReportType = "7"
	TradeReportType_DEFAULTED                   TradeReportType = "8"
	TradeReportType_INVALID_CMTA                TradeReportType = "9"
)

// TradeReportingIndicator field enumeration values.
type TradeReportingIndicator string

const (
	TradeReportingIndicator_TRADE_HAS_NOT                                                                                                       TradeReportingIndicator = "0"
	TradeReportingIndicator_TRADE_HAS_BEEN_OR_WILL_BE_REPORTED_BY_A_TRADING_VENUE_AS_AN_ON_BOOK_TRADE                                           TradeReportingIndicator = "1"
	TradeReportingIndicator_TRADE_HAS_BEEN_OR_WILL_BE_REPORTED_AS_A_SYSTEMATIC_INTERNALISER_SELLER_TRADE                                        TradeReportingIndicator = "2"
	TradeReportingIndicator_TRADE_HAS_BEEN_OR_WILL_BE_REPORTED_AS_A_SYSTEMATIC_INTERNALISER_BUYER_TRADE                                         TradeReportingIndicator = "3"
	TradeReportingIndicator_TRADE_HAS_BEEN_OR_WILL_BE_REPORTED_AS_A_NON_SYSTEMATIC_INTERNALISER_SELLER_TRADE                                    TradeReportingIndicator = "4"
	TradeReportingIndicator_TRADE_HAS_BEEN_OR_WILL_BE_REPORTED_UNDER_A_SUB_DELEGATION_ARRANGEMENT_BY_AN_INVESTMENT_FIRM_TO_A_REPORTING_FACILITY TradeReportingIndicator = "5"
	TradeReportingIndicator_TRADE_HAS_BEEN_OR_WILL_BE_REPORTED                                                                                  TradeReportingIndicator = "6"
	TradeReportingIndicator_TRADE_HAS_BEEN_OR_WILL_BE_REPORTED_AS_A_NON_SYSTEMATIC_INTERNALISER_BUYER_TRADE                                     TradeReportingIndicator = "7"
	TradeReportingIndicator_TRADE_HAS_BEEN_OR_WILL_BE_REPORTED_BY_A_TRADING_VENUE_AS_AN_OFF_BOOK_TRADE                                          TradeReportingIndicator = "8"
	TradeReportingIndicator_TRADE_IS_NOT_REPORTABLE                                                                                             TradeReportingIndicator = "9"
)

// TradeRequestResult field enumeration values.
type TradeRequestResult string

const (
	TradeRequestResult_SUCCESSFUL                       TradeRequestResult = "0"
	TradeRequestResult_INVALID_OR_UNKNOWN_INSTRUMENT    TradeRequestResult = "1"
	TradeRequestResult_INVALID_TYPE_OF_TRADE_REQUESTED  TradeRequestResult = "2"
	TradeRequestResult_INVALID_PARTIES                  TradeRequestResult = "3"
	TradeRequestResult_INVALID_TRANSPORT_TYPE_REQUESTED TradeRequestResult = "4"
	TradeRequestResult_INVALID_DESTINATION_REQUESTED    TradeRequestResult = "5"
	TradeRequestResult_TRADEREQUESTTYPE_NOT_SUPPORTED   TradeRequestResult = "8"
	TradeRequestResult_NOT_AUTHORIZED                   TradeRequestResult = "9"
	TradeRequestResult_OTHER                            TradeRequestResult = "99"
)

// TradeRequestStatus field enumeration values.
type TradeRequestStatus string

const (
	TradeRequestStatus_ACCEPTED  TradeRequestStatus = "0"
	TradeRequestStatus_COMPLETED TradeRequestStatus = "1"
	TradeRequestStatus_REJECTED  TradeRequestStatus = "2"
)

// TradeRequestType field enumeration values.
type TradeRequestType string

const (
	TradeRequestType_ALL_TRADES                                           TradeRequestType = "0"
	TradeRequestType_MATCHED_TRADES_MATCHING_CRITERIA_PROVIDED_ON_REQUEST TradeRequestType = "1"
	TradeRequestType_UNMATCHED_TRADES_THAT_MATCH_CRITERIA                 TradeRequestType = "2"
	TradeRequestType_UNREPORTED_TRADES_THAT_MATCH_CRITERIA                TradeRequestType = "3"
	TradeRequestType_ADVISORIES_THAT_MATCH_CRITERIA                       TradeRequestType = "4"
)

// TradeType field enumeration values.
type TradeType string

const (
	TradeType_AGENCY           TradeType = "A"
	TradeType_VWAP_GUARANTEE   TradeType = "G"
	TradeType_GUARANTEED_CLOSE TradeType = "J"
	TradeType_RISK_TRADE       TradeType = "R"
)

// TradeVolType field enumeration values.
type TradeVolType string

const (
	TradeVolType_NUMBER_OF_UNITS      TradeVolType = "0"
	TradeVolType_NUMBER_OF_ROUND_LOTS TradeVolType = "1"
)

// TradedFlatSwitch field enumeration values.
type TradedFlatSwitch string

const (
	TradedFlatSwitch_NO  TradedFlatSwitch = "N"
	TradedFlatSwitch_YES TradedFlatSwitch = "Y"
)

// TradingCapacity field enumeration values.
type TradingCapacity string

const (
	TradingCapacity_CUSTOMER                TradingCapacity = "1"
	TradingCapacity_CUSTOMER_PROFESSIONAL   TradingCapacity = "2"
	TradingCapacity_BROKER_DEALER           TradingCapacity = "3"
	TradingCapacity_CUSTOMER_BROKER_DEALER  TradingCapacity = "4"
	TradingCapacity_PRINCIPAL               TradingCapacity = "5"
	TradingCapacity_MARKET_MAKER            TradingCapacity = "6"
	TradingCapacity_AWAY_MARKET_MAKER       TradingCapacity = "7"
	TradingCapacity_SYSTEMATIC_INTERNALISER TradingCapacity = "8"
)

// TradingSessionID field enumeration values.
type TradingSessionID string

const (
	TradingSessionID_DAY         TradingSessionID = "1"
	TradingSessionID_HALFDAY     TradingSessionID = "2"
	TradingSessionID_MORNING     TradingSessionID = "3"
	TradingSessionID_AFTERNOON   TradingSessionID = "4"
	TradingSessionID_EVENING     TradingSessionID = "5"
	TradingSessionID_AFTER_HOURS TradingSessionID = "6"
	TradingSessionID_HOLIDAY     TradingSessionID = "7"
)

// TradingSessionSubID field enumeration values.
type TradingSessionSubID string

const (
	TradingSessionSubID_PRE_TRADING                  TradingSessionSubID = "1"
	TradingSessionSubID_OUT_OF_MAIN_SESSION_TRADING  TradingSessionSubID = "10"
	TradingSessionSubID_PRIVATE_AUCTION              TradingSessionSubID = "11"
	TradingSessionSubID_PUBLIC_AUCTION               TradingSessionSubID = "12"
	TradingSessionSubID_GROUP_AUCTION                TradingSessionSubID = "13"
	TradingSessionSubID_OPENING_OR_OPENING_AUCTION   TradingSessionSubID = "2"
	TradingSessionSubID_3                            TradingSessionSubID = "3"
	TradingSessionSubID_CLOSING_OR_CLOSING_AUCTION   TradingSessionSubID = "4"
	TradingSessionSubID_POST_TRADING                 TradingSessionSubID = "5"
	TradingSessionSubID_SCHEDULED_INTRADAY_AUCTION   TradingSessionSubID = "6"
	TradingSessionSubID_QUIESCENT                    TradingSessionSubID = "7"
	TradingSessionSubID_ANY_AUCTION                  TradingSessionSubID = "8"
	TradingSessionSubID_UNSCHEDULED_INTRADAY_AUCTION TradingSessionSubID = "9"
)

// TransactionAttributeType field enumeration values.
type TransactionAttributeType string

const (
	TransactionAttributeType_EXCLUSIVE_ARRANGEMENT       TransactionAttributeType = "0"
	TransactionAttributeType_COLLATERAL_REUSE            TransactionAttributeType = "1"
	TransactionAttributeType_COLLATERAL_ARRANGEMENT_TYPE TransactionAttributeType = "2"
)

// TransferRejectReason field enumeration values.
type TransferRejectReason string

const (
	TransferRejectReason_SUCCESS                            TransferRejectReason = "0"
	TransferRejectReason_INVALID_PARTY                      TransferRejectReason = "1"
	TransferRejectReason_UNKNOWN_INSTRUMENT                 TransferRejectReason = "2"
	TransferRejectReason_NOT_AUTHORIZED_TO_SUBMIT_TRANSFERS TransferRejectReason = "3"
	TransferRejectReason_UNKNOWN_POSITION                   TransferRejectReason = "4"
	TransferRejectReason_OTHER                              TransferRejectReason = "99"
)

// TransferReportType field enumeration values.
type TransferReportType string

const (
	TransferReportType_SUBMIT  TransferReportType = "0"
	TransferReportType_ALLEGED TransferReportType = "1"
)

// TransferScope field enumeration values.
type TransferScope string

const (
	TransferScope_INTER_FIRM_TRANSFER              TransferScope = "0"
	TransferScope_INTRA_FIRM_TRANSFER              TransferScope = "1"
	TransferScope_CLEARING_MEMBER_TRADE_ASSIGNMENT TransferScope = "2"
)

// TransferStatus field enumeration values.
type TransferStatus string

const (
	TransferStatus_RECEIVED                 TransferStatus = "0"
	TransferStatus_REJECTED_BY_INTERMEDIARY TransferStatus = "1"
	TransferStatus_ACCEPT_PENDING           TransferStatus = "2"
	TransferStatus_ACCEPTED                 TransferStatus = "3"
	TransferStatus_DECLINED                 TransferStatus = "4"
	TransferStatus_CANCELLED                TransferStatus = "5"
)

// TransferTransType field enumeration values.
type TransferTransType string

const (
	TransferTransType_NEW     TransferTransType = "0"
	TransferTransType_REPLACE TransferTransType = "1"
	TransferTransType_CANCEL  TransferTransType = "2"
)

// TransferType field enumeration values.
type TransferType string

const (
	TransferType_REQUEST_TRANSFER TransferType = "0"
	TransferType_ACCEPT_TRANSFER  TransferType = "1"
	TransferType_DECLINE_TRANSFER TransferType = "2"
)

// TrdAckStatus field enumeration values.
type TrdAckStatus string

const (
	TrdAckStatus_ACCEPTED TrdAckStatus = "0"
	TrdAckStatus_REJECTED TrdAckStatus = "1"
	TrdAckStatus_RECEIVED TrdAckStatus = "2"
)

// TrdRegPublicationReason field enumeration values.
type TrdRegPublicationReason string

const (
	TrdRegPublicationReason_NO_PRECEDING_ORDER_IN_BOOK_AS_TRANSACTION_PRICE_SET_WITHIN_AVERAGE_SPREAD_OF_A_LIQUID_INSTRUMENT                         TrdRegPublicationReason = "0"
	TrdRegPublicationReason_NO_PRECEDING_ORDER_IN_BOOK_AS_TRANSACTION_PRICE_DEPENDS_ON_SYSTEM_SET_REFERENCE_PRICE_FOR_AN_ILLIQUID_INSTRUMENT         TrdRegPublicationReason = "1"
	TrdRegPublicationReason_NO_PUBLIC_PRICE_AND_OR_SIZE_QUOTED_DUE_TO_ORDER_BEING_HIDDEN                                                             TrdRegPublicationReason = "10"
	TrdRegPublicationReason_EXEMPTED_DUE_TO_SECURITIES_FINANCING_TRANSACTION                                                                         TrdRegPublicationReason = "11"
	TrdRegPublicationReason_EXEMPTED_DUE_TO_EUROPEAN_SYSTEM_OF_CENTRAL_BANKS                                                                         TrdRegPublicationReason = "12"
	TrdRegPublicationReason_EXCEPTION_DUE_TO_REPORT_BY_PAPER                                                                                         TrdRegPublicationReason = "13"
	TrdRegPublicationReason_EXCEPTION_DUE_TO_TRADE_WITH_NON_REPORTING_PARTY                                                                          TrdRegPublicationReason = "14"
	TrdRegPublicationReason_EXCEPTION_DUE_TO_INTRA_FIRM_ORDER                                                                                        TrdRegPublicationReason = "15"
	TrdRegPublicationReason_REPORTED_OUTSIDE_OF_REPORTING_HOURS                                                                                      TrdRegPublicationReason = "16"
	TrdRegPublicationReason_NO_PRECEDING_ORDER_IN_BOOK_AS_TRANSACTION_PRICE_IS_FOR_TRANSACTION_SUBJECT_TO_CONDITIONS_OTHER_THAN_CURRENT_MARKET_PRICE TrdRegPublicationReason = "2"
	TrdRegPublicationReason_NO_PUBLIC_PRICE_FOR_PRECEDING_ORDER_AS_PUBLIC_REFERENCE_PRICE_WAS_USED_FOR_MATCHING_ORDERS                               TrdRegPublicationReason = "3"
	TrdRegPublicationReason_NO_PUBLIC_PRICE_QUOTED_AS_INSTRUMENT_IS_ILLIQUID                                                                         TrdRegPublicationReason = "4"
	TrdRegPublicationReason_NO_PUBLIC_PRICE_QUOTED_DUE_TO_SIZE                                                                                       TrdRegPublicationReason = "5"
	TrdRegPublicationReason_DEFERRAL_DUE_TO_LARGE_IN_SCALE                                                                                           TrdRegPublicationReason = "6"
	TrdRegPublicationReason_DEFERRAL_DUE_TO_ILLIQUID_INSTRUMENT                                                                                      TrdRegPublicationReason = "7"
	TrdRegPublicationReason_DEFERRAL_DUE_TO_SIZE_SPECIFIC                                                                                            TrdRegPublicationReason = "8"
	TrdRegPublicationReason_NO_PUBLIC_PRICE_AND_OR_SIZE_QUOTED_AS_TRANSACTION_IS_LARGE_IN_SCALE                                                      TrdRegPublicationReason = "9"
)

// TrdRegPublicationType field enumeration values.
type TrdRegPublicationType string

const (
	TrdRegPublicationType_PRE_TRADE_TRANSPARENCY_WAIVER           TrdRegPublicationType = "0"
	TrdRegPublicationType_POST_TRADE_DEFERRAL                     TrdRegPublicationType = "1"
	TrdRegPublicationType_EXEMPT_FROM_PUBLICATION                 TrdRegPublicationType = "2"
	TrdRegPublicationType_ORDER_LEVEL_PUBLICATION_TO_SUBSCRIBERS  TrdRegPublicationType = "3"
	TrdRegPublicationType_PRICE_LEVEL_PUBLICATION_TO_SUBSCRIBERS  TrdRegPublicationType = "4"
	TrdRegPublicationType_ORDER_LEVEL_PUBLICATION_TO_THE_PUBLIC   TrdRegPublicationType = "5"
	TrdRegPublicationType_PUBLICATION_INTERNAL_TO_EXECUTION_VENUE TrdRegPublicationType = "6"
)

// TrdRegTimestampManualIndicator field enumeration values.
type TrdRegTimestampManualIndicator string

const (
	TrdRegTimestampManualIndicator_NO  TrdRegTimestampManualIndicator = "N"
	TrdRegTimestampManualIndicator_YES TrdRegTimestampManualIndicator = "Y"
)

// TrdRegTimestampType field enumeration values.
type TrdRegTimestampType string

const (
	TrdRegTimestampType_EXECUTION_TIME                TrdRegTimestampType = "1"
	TrdRegTimestampType_ORDER_SUBMISSION_TIME         TrdRegTimestampType = "10"
	TrdRegTimestampType_PUBLICLY_REPORTED             TrdRegTimestampType = "11"
	TrdRegTimestampType_PUBLIC_REPORT_UPDATED         TrdRegTimestampType = "12"
	TrdRegTimestampType_NON_PUBLICLY_REPORTED         TrdRegTimestampType = "13"
	TrdRegTimestampType_NON_PUBLIC_REPORT_UPDATED     TrdRegTimestampType = "14"
	TrdRegTimestampType_SUBMITTED_FOR_CONFIRMATION    TrdRegTimestampType = "15"
	TrdRegTimestampType_UPDATED_FOR_CONFIRMATION      TrdRegTimestampType = "16"
	TrdRegTimestampType_CONFIRMED                     TrdRegTimestampType = "17"
	TrdRegTimestampType_UPDATED_FOR_CLEARING          TrdRegTimestampType = "18"
	TrdRegTimestampType_CLEARED                       TrdRegTimestampType = "19"
	TrdRegTimestampType_TIME_IN                       TrdRegTimestampType = "2"
	TrdRegTimestampType_ALLOCATIONS_SUBMITTED         TrdRegTimestampType = "20"
	TrdRegTimestampType_ALLOCATIONS_UPDATED           TrdRegTimestampType = "21"
	TrdRegTimestampType_ALLOCATIONS_COMPLETED         TrdRegTimestampType = "22"
	TrdRegTimestampType_SUBMITTED_TO_REPOSITORY       TrdRegTimestampType = "23"
	TrdRegTimestampType_POST_TRADE_CONTINUATION_EVENT TrdRegTimestampType = "24"
	TrdRegTimestampType_POST_TRADE_VALUATION          TrdRegTimestampType = "25"
	TrdRegTimestampType_PREVIOUS_TIME_PRIORITY        TrdRegTimestampType = "26"
	TrdRegTimestampType_IDENTIFIER_ASSIGNED           TrdRegTimestampType = "27"
	TrdRegTimestampType_PREVIOUS_IDENTIFIER_ASSIGNED  TrdRegTimestampType = "28"
	TrdRegTimestampType_ORDER_CANCELLATION_TIME       TrdRegTimestampType = "29"
	TrdRegTimestampType_TIME_OUT                      TrdRegTimestampType = "3"
	TrdRegTimestampType_ORDER_MODIFICATION_TIME       TrdRegTimestampType = "30"
	TrdRegTimestampType_ORDER_ROUTING_TIME            TrdRegTimestampType = "31"
	TrdRegTimestampType_TRADE_CANCELLATION_TIME       TrdRegTimestampType = "32"
	TrdRegTimestampType_TRADE_MODIFICATION_TIME       TrdRegTimestampType = "33"
	TrdRegTimestampType_REFERENCE_TIME_FOR_NBBO       TrdRegTimestampType = "34"
	TrdRegTimestampType_BROKER_RECEIPT                TrdRegTimestampType = "4"
	TrdRegTimestampType_BROKER_EXECUTION              TrdRegTimestampType = "5"
	TrdRegTimestampType_DESK_RECEIPT                  TrdRegTimestampType = "6"
	TrdRegTimestampType_SUBMISSION_TO_CLEARING        TrdRegTimestampType = "7"
	TrdRegTimestampType_TIME_PRIORITY                 TrdRegTimestampType = "8"
	TrdRegTimestampType_ORDERBOOK_ENTRY_TIME          TrdRegTimestampType = "9"
)

// TrdRptStatus field enumeration values.
type TrdRptStatus string

const (
	TrdRptStatus_ACCEPTED             TrdRptStatus = "0"
	TrdRptStatus_REJECTED             TrdRptStatus = "1"
	TrdRptStatus_VERIFIED             TrdRptStatus = "10"
	TrdRptStatus_DISPUTED             TrdRptStatus = "11"
	TrdRptStatus_CANCELLED            TrdRptStatus = "2"
	TrdRptStatus_ACCEPTED_WITH_ERRORS TrdRptStatus = "3"
	TrdRptStatus_PENDING_NEW          TrdRptStatus = "4"
	TrdRptStatus_PENDING_CANCEL       TrdRptStatus = "5"
	TrdRptStatus_PENDING_REPLACE      TrdRptStatus = "6"
	TrdRptStatus_TERMINATED           TrdRptStatus = "7"
	TrdRptStatus_PENDING_VERIFICATION TrdRptStatus = "8"
	TrdRptStatus_DEEMED_VERIFIED      TrdRptStatus = "9"
)

// TrdSubType field enumeration values.
type TrdSubType string

const (
	TrdSubType_CMTA                                                 TrdSubType = "0"
	TrdSubType_INTERNAL_TRANSFER_OR_ADJUSTMENT                      TrdSubType = "1"
	TrdSubType_TRANSACTION_FROM_ASSIGNMENT                          TrdSubType = "10"
	TrdSubType_ACATS                                                TrdSubType = "11"
	TrdSubType_AI                                                   TrdSubType = "14"
	TrdSubType_B                                                    TrdSubType = "15"
	TrdSubType_K                                                    TrdSubType = "16"
	TrdSubType_LC                                                   TrdSubType = "17"
	TrdSubType_M                                                    TrdSubType = "18"
	TrdSubType_N                                                    TrdSubType = "19"
	TrdSubType_EXTERNAL_TRANSFER_OR_TRANSFER_OF_ACCOUNT             TrdSubType = "2"
	TrdSubType_NM                                                   TrdSubType = "20"
	TrdSubType_NR                                                   TrdSubType = "21"
	TrdSubType_P                                                    TrdSubType = "22"
	TrdSubType_PA                                                   TrdSubType = "23"
	TrdSubType_PC                                                   TrdSubType = "24"
	TrdSubType_PN                                                   TrdSubType = "25"
	TrdSubType_R                                                    TrdSubType = "26"
	TrdSubType_RO                                                   TrdSubType = "27"
	TrdSubType_RT                                                   TrdSubType = "28"
	TrdSubType_SW                                                   TrdSubType = "29"
	TrdSubType_REJECT_FOR_SUBMITTING_SIDE                           TrdSubType = "3"
	TrdSubType_T                                                    TrdSubType = "30"
	TrdSubType_WN                                                   TrdSubType = "31"
	TrdSubType_WT                                                   TrdSubType = "32"
	TrdSubType_OFF_HOURS_TRADE                                      TrdSubType = "33"
	TrdSubType_ON_HOURS_TRADE                                       TrdSubType = "34"
	TrdSubType_OTC_QUOTE                                            TrdSubType = "35"
	TrdSubType_CONVERTED_SWAP                                       TrdSubType = "36"
	TrdSubType_CROSSED_TRADE                                        TrdSubType = "37"
	TrdSubType_INTERIM_PROTECTED_TRADE                              TrdSubType = "38"
	TrdSubType_LARGE_IN_SCALE                                       TrdSubType = "39"
	TrdSubType_ADVISORY_FOR_CONTRA_SIDE                             TrdSubType = "4"
	TrdSubType_WASH_TRADE                                           TrdSubType = "40"
	TrdSubType_TRADE_AT_SETTLEMENT                                  TrdSubType = "41"
	TrdSubType_AUCTION_TRADE                                        TrdSubType = "42"
	TrdSubType_TRADE_AT_MARKER                                      TrdSubType = "43"
	TrdSubType_DEFAULT                                              TrdSubType = "44"
	TrdSubType_RESTRUCTURING                                        TrdSubType = "45"
	TrdSubType_MERGER                                               TrdSubType = "46"
	TrdSubType_SPIN_OFF                                             TrdSubType = "47"
	TrdSubType_MULTILATERAL_COMPRESSION                             TrdSubType = "48"
	TrdSubType_OFFSET_DUE_TO_AN_ALLOCATION                          TrdSubType = "5"
	TrdSubType_BALANCING                                            TrdSubType = "50"
	TrdSubType_BASIS_TRADE_INDEX_CLOSE                              TrdSubType = "51"
	TrdSubType_TRADE_AT_CASH_OPEN                                   TrdSubType = "52"
	TrdSubType_TRADE_SUBMITTED_TO_VENUE_FOR_CLEARING_AND_SETTLEMENT TrdSubType = "53"
	TrdSubType_BILATERAL_COMPRESSION                                TrdSubType = "54"
	TrdSubType_ONSET_DUE_TO_AN_ALLOCATION                           TrdSubType = "6"
	TrdSubType_DIFFERENTIAL_SPREAD                                  TrdSubType = "7"
	TrdSubType_IMPLIED_SPREAD_LEG_EXECUTED_AGAINST_AN_OUTRIGHT      TrdSubType = "8"
	TrdSubType_TRANSACTION_FROM_EXERCISE                            TrdSubType = "9"
)

// TrdType field enumeration values.
type TrdType string

const (
	TrdType_REGULAR_TRADE                                   TrdType = "0"
	TrdType_BLOCK_TRADE_1                                   TrdType = "1"
	TrdType_AFTER_HOURS_TRADE                               TrdType = "10"
	TrdType_EXCHANGE_FOR_RISK                               TrdType = "11"
	TrdType_EXCHANGE_FOR_SWAP                               TrdType = "12"
	TrdType_EXCHANGE_OF_FUTURES_FOR_IN_MARKET_FUTURES       TrdType = "13"
	TrdType_EXCHANGE_OF_OPTIONS_FOR_OPTIONS                 TrdType = "14"
	TrdType_TRADING_AT_SETTLEMENT                           TrdType = "15"
	TrdType_ALL_OR_NONE                                     TrdType = "16"
	TrdType_FUTURES_LARGE_ORDER_EXECUTION                   TrdType = "17"
	TrdType_EXCHANGE_OF_FUTURES_FOR_EXTERNAL_MARKET_FUTURES TrdType = "18"
	TrdType_OPTION_INTERIM_TRADE                            TrdType = "19"
	TrdType_EXCHANGE_FOR_PHYSICAL                           TrdType = "2"
	TrdType_OPTION_CABINET_TRADE                            TrdType = "20"
	TrdType_PRIVATELY_NEGOTIATED_TRADE                      TrdType = "22"
	TrdType_SUBSTITUTION_OF_FUTURES_FOR_FORWARDS            TrdType = "23"
	TrdType_ERROR_TRADE                                     TrdType = "24"
	TrdType_SPECIAL_CUM_DIVIDEND                            TrdType = "25"
	TrdType_SPECIAL_EX_DIVIDEND                             TrdType = "26"
	TrdType_SPECIAL_CUM_COUPON                              TrdType = "27"
	TrdType_SPECIAL_EX_COUPON                               TrdType = "28"
	TrdType_CASH_SETTLEMENT                                 TrdType = "29"
	TrdType_TRANSFER                                        TrdType = "3"
	TrdType_SPECIAL_PRICE                                   TrdType = "30"
	TrdType_GUARANTEED_DELIVERY                             TrdType = "31"
	TrdType_SPECIAL_CUM_RIGHTS                              TrdType = "32"
	TrdType_SPECIAL_EX_RIGHTS                               TrdType = "33"
	TrdType_SPECIAL_CUM_CAPITAL_REPAYMENTS                  TrdType = "34"
	TrdType_SPECIAL_EX_CAPITAL_REPAYMENTS                   TrdType = "35"
	TrdType_SPECIAL_CUM_BONUS                               TrdType = "36"
	TrdType_SPECIAL_EX_BONUS                                TrdType = "37"
	TrdType_BLOCK_TRADE_38                                  TrdType = "38"
	TrdType_WORKED_PRINCIPAL_TRADE                          TrdType = "39"
	TrdType_LATE_TRADE                                      TrdType = "4"
	TrdType_BLOCK_TRADES                                    TrdType = "40"
	TrdType_NAME_CHANGE                                     TrdType = "41"
	TrdType_PORTFOLIO_TRANSFER                              TrdType = "42"
	TrdType_PROROGATION_BUY                                 TrdType = "43"
	TrdType_PROROGATION_SELL                                TrdType = "44"
	TrdType_OPTION_EXERCISE                                 TrdType = "45"
	TrdType_DELTA_NEUTRAL_TRANSACTION                       TrdType = "46"
	TrdType_FINANCING_TRANSACTION                           TrdType = "47"
	TrdType_NON_STANDARD_SETTLEMENT                         TrdType = "48"
	TrdType_DERIVATIVE_RELATED_TRANSACTION                  TrdType = "49"
	TrdType_T_TRADE                                         TrdType = "5"
	TrdType_PORTFOLIO_TRADE                                 TrdType = "50"
	TrdType_VOLUME_WEIGHTED_AVERAGE_TRADE                   TrdType = "51"
	TrdType_EXCHANGE_GRANTED_TRADE                          TrdType = "52"
	TrdType_REPURCHASE_AGREEMENT                            TrdType = "53"
	TrdType_OTC                                             TrdType = "54"
	TrdType_EXCHANGE_BASIS_FACILITY                         TrdType = "55"
	TrdType_OPENING_TRADE                                   TrdType = "56"
	TrdType_NETTED_TRADE                                    TrdType = "57"
	TrdType_BLOCK_SWAP_TRADE                                TrdType = "58"
	TrdType_CREDIT_EVENT_TRADE                              TrdType = "59"
	TrdType_WEIGHTED_AVERAGE_PRICE_TRADE                    TrdType = "6"
	TrdType_SUCCESSION_EVENT_TRADE                          TrdType = "60"
	TrdType_GIVE_UP_GIVE_IN_TRADE                           TrdType = "61"
	TrdType_DARK_TRADE                                      TrdType = "62"
	TrdType_TECHNICAL_TRADE                                 TrdType = "63"
	TrdType_BENCHMARK                                       TrdType = "64"
	TrdType_PACKAGE_TRADE                                   TrdType = "65"
	TrdType_ROLL_TRADE                                      TrdType = "66"
	TrdType_BUNCHED_TRADE                                   TrdType = "7"
	TrdType_LATE_BUNCHED_TRADE                              TrdType = "8"
	TrdType_PRIOR_REFERENCE_PRICE_TRADE                     TrdType = "9"
)

// TriggerAction field enumeration values.
type TriggerAction string

const (
	TriggerAction_ACTIVATE TriggerAction = "1"
	TriggerAction_MODIFY   TriggerAction = "2"
	TriggerAction_CANCEL   TriggerAction = "3"
)

// TriggerOrderType field enumeration values.
type TriggerOrderType string

const (
	TriggerOrderType_MARKET TriggerOrderType = "1"
	TriggerOrderType_LIMIT  TriggerOrderType = "2"
)

// TriggerPriceDirection field enumeration values.
type TriggerPriceDirection string

const (
	TriggerPriceDirection_TRIGGER_IF_THE_PRICE_OF_THE_SPECIFIED_TYPE_GOES_DOWN_TO_OR_THROUGH_THE_SPECIFIED_TRIGGER_PRICE TriggerPriceDirection = "D"
	TriggerPriceDirection_TRIGGER_IF_THE_PRICE_OF_THE_SPECIFIED_TYPE_GOES_UP_TO_OR_THROUGH_THE_SPECIFIED_TRIGGER_PRICE   TriggerPriceDirection = "U"
)

// TriggerPriceType field enumeration values.
type TriggerPriceType string

const (
	TriggerPriceType_BEST_OFFER               TriggerPriceType = "1"
	TriggerPriceType_LAST_TRADE               TriggerPriceType = "2"
	TriggerPriceType_BEST_BID                 TriggerPriceType = "3"
	TriggerPriceType_BEST_BID_OR_LAST_TRADE   TriggerPriceType = "4"
	TriggerPriceType_BEST_OFFER_OR_LAST_TRADE TriggerPriceType = "5"
	TriggerPriceType_BEST_MID                 TriggerPriceType = "6"
)

// TriggerPriceTypeScope field enumeration values.
type TriggerPriceTypeScope string

const (
	TriggerPriceTypeScope_NONE     TriggerPriceTypeScope = "0"
	TriggerPriceTypeScope_LOCAL    TriggerPriceTypeScope = "1"
	TriggerPriceTypeScope_NATIONAL TriggerPriceTypeScope = "2"
	TriggerPriceTypeScope_GLOBAL   TriggerPriceTypeScope = "3"
)

// TriggerScope field enumeration values.
type TriggerScope string

const (
	TriggerScope_THIS_ORDER                                             TriggerScope = "0"
	TriggerScope_OTHER_ORDER                                            TriggerScope = "1"
	TriggerScope_ALL_OTHER_ORDERS_FOR_THE_GIVEN_SECURITY                TriggerScope = "2"
	TriggerScope_ALL_OTHER_ORDERS_FOR_THE_GIVEN_SECURITY_AND_PRICE      TriggerScope = "3"
	TriggerScope_ALL_OTHER_ORDERS_FOR_THE_GIVEN_SECURITY_AND_SIDE       TriggerScope = "4"
	TriggerScope_ALL_OTHER_ORDERS_FOR_THE_GIVEN_SECURITY_PRICE_AND_SIDE TriggerScope = "5"
)

// TriggerType field enumeration values.
type TriggerType string

const (
	TriggerType_PARTIAL_EXECUTION                          TriggerType = "1"
	TriggerType_SPECIFIED_TRADING_SESSION                  TriggerType = "2"
	TriggerType_NEXT_AUCTION                               TriggerType = "3"
	TriggerType_PRICE_MOVEMENT                             TriggerType = "4"
	TriggerType_ON_ORDER_ENTRY_OR_ORDER_MODIFICATION_ENTRY TriggerType = "5"
)

// Triggered field enumeration values.
type Triggered string

const (
	Triggered_NOT_TRIGGERED          Triggered = "0"
	Triggered_TRIGGERED              Triggered = "1"
	Triggered_STOP_ORDER_TRIGGERED   Triggered = "2"
	Triggered_ONE_CANCELS_THE_OTHER  Triggered = "3"
	Triggered_ONE_TRIGGERS_THE_OTHER Triggered = "4"
	Triggered_ONE_UPDATES_THE_OTHER  Triggered = "5"
)

// UnderlyingCashType field enumeration values.
type UnderlyingCashType string

const (
	UnderlyingCashType_DIFF  UnderlyingCashType = "DIFF"
	UnderlyingCashType_FIXED UnderlyingCashType = "FIXED"
)

// UnderlyingFXRateCalc field enumeration values.
type UnderlyingFXRateCalc string

const (
	UnderlyingFXRateCalc_DIVIDE   UnderlyingFXRateCalc = "D"
	UnderlyingFXRateCalc_MULTIPLY UnderlyingFXRateCalc = "M"
)

// UnderlyingNotionalAdjustments field enumeration values.
type UnderlyingNotionalAdjustments string

const (
	UnderlyingNotionalAdjustments_EXECUTION             UnderlyingNotionalAdjustments = "0"
	UnderlyingNotionalAdjustments_PORTFOLIO_REBALANCING UnderlyingNotionalAdjustments = "1"
	UnderlyingNotionalAdjustments_STANDRD               UnderlyingNotionalAdjustments = "2"
)

// UnderlyingObligationType field enumeration values.
type UnderlyingObligationType string

const (
	UnderlyingObligationType_BOND             UnderlyingObligationType = "0"
	UnderlyingObligationType_CONVERTIBLE_BOND UnderlyingObligationType = "1"
	UnderlyingObligationType_MORTGAGE         UnderlyingObligationType = "2"
	UnderlyingObligationType_LOAN             UnderlyingObligationType = "3"
)

// UnderlyingPriceDeterminationMethod field enumeration values.
type UnderlyingPriceDeterminationMethod string

const (
	UnderlyingPriceDeterminationMethod_REGULAR           UnderlyingPriceDeterminationMethod = "1"
	UnderlyingPriceDeterminationMethod_SPECIAL_REFERENCE UnderlyingPriceDeterminationMethod = "2"
	UnderlyingPriceDeterminationMethod_OPTIMAL_VALUE     UnderlyingPriceDeterminationMethod = "3"
	UnderlyingPriceDeterminationMethod_AVERAGE_VALUE     UnderlyingPriceDeterminationMethod = "4"
)

// UnderlyingSettlementType field enumeration values.
type UnderlyingSettlementType string

const (
	UnderlyingSettlementType_T_PLUS_1 UnderlyingSettlementType = "2"
	UnderlyingSettlementType_T_PLUS_3 UnderlyingSettlementType = "4"
	UnderlyingSettlementType_T_PLUS_4 UnderlyingSettlementType = "5"
)

// UnitOfMeasure field enumeration values.
type UnitOfMeasure string

const (
	UnitOfMeasure_ALLOWANCES                                 UnitOfMeasure = "Alw"
	UnitOfMeasure_BOARD_FEET                                 UnitOfMeasure = "BDFT"
	UnitOfMeasure_BARRELS                                    UnitOfMeasure = "Bbl"
	UnitOfMeasure_BILLION_CUBIC_FEET                         UnitOfMeasure = "Bcf"
	UnitOfMeasure_BUSHELS                                    UnitOfMeasure = "Bu"
	UnitOfMeasure_CUBIC_METERS                               UnitOfMeasure = "CBM"
	UnitOfMeasure_COOLING_DEGREE_DAY                         UnitOfMeasure = "CDD"
	UnitOfMeasure_CERTIFIED_EMISSIONS_REDUCTION              UnitOfMeasure = "CER"
	UnitOfMeasure_CRITICAL_PRECIPITATION_DAY                 UnitOfMeasure = "CPD"
	UnitOfMeasure_CLIMATE_RESERVE_TONNES                     UnitOfMeasure = "CRT"
	UnitOfMeasure_AMOUNT_OF_CURRENCY                         UnitOfMeasure = "Ccy"
	UnitOfMeasure_DIESEL_GALLON_EQUIVALENT                   UnitOfMeasure = "DGE"
	UnitOfMeasure_ENVIRONMENTAL_ALLOWANCE_CERTIFICATES       UnitOfMeasure = "EnvAllwnc"
	UnitOfMeasure_ENVIRONMENTAL_CREDIT                       UnitOfMeasure = "EnvCrd"
	UnitOfMeasure_ENVIRONMENTAL_OFFSET                       UnitOfMeasure = "EnvOfst"
	UnitOfMeasure_GASONLINE_GALLON_EQUIVALENT                UnitOfMeasure = "GGE"
	UnitOfMeasure_GIGAJOULES                                 UnitOfMeasure = "GJ"
	UnitOfMeasure_GROSS_TONS                                 UnitOfMeasure = "GT"
	UnitOfMeasure_GALLONS                                    UnitOfMeasure = "Gal"
	UnitOfMeasure_GB_GALLON                                  UnitOfMeasure = "Gal_gb"
	UnitOfMeasure_HEATING_DEGREE_DAY                         UnitOfMeasure = "HDD"
	UnitOfMeasure_INDEX_POINT                                UnitOfMeasure = "IPNT"
	UnitOfMeasure_LITERS                                     UnitOfMeasure = "L"
	UnitOfMeasure_METER                                      UnitOfMeasure = "M"
	UnitOfMeasure_MEGA_HEAT_RATE                             UnitOfMeasure = "MHR"
	UnitOfMeasure_ONE_MILLION_BTU                            UnitOfMeasure = "MMBtu"
	UnitOfMeasure_MILLION_BARRELS                            UnitOfMeasure = "MMbbl"
	UnitOfMeasure_MEGAWATT_MONTH                             UnitOfMeasure = "MW-M"
	UnitOfMeasure_MEGAWATT_YEAR                              UnitOfMeasure = "MW-a"
	UnitOfMeasure_MEGAWATT_DAY                               UnitOfMeasure = "MW-d"
	UnitOfMeasure_MEGAWATT_HOUR                              UnitOfMeasure = "MW-h"
	UnitOfMeasure_MEGAWATT_MINUTE                            UnitOfMeasure = "MW-min"
	UnitOfMeasure_MEGAWATT_HOURS                             UnitOfMeasure = "MWh"
	UnitOfMeasure_PRINCIPAL_WITH_RELATION_TO_DEBT_INSTRUMENT UnitOfMeasure = "PRINC"
	UnitOfMeasure_SQUARE_METER                               UnitOfMeasure = "SqM"
	UnitOfMeasure_SQUARE_CENTIMETER                          UnitOfMeasure = "SqcM"
	UnitOfMeasure_SQUARE_FOOT                                UnitOfMeasure = "Sqft"
	UnitOfMeasure_SQUARE_INCH                                UnitOfMeasure = "Sqin"
	UnitOfMeasure_SQUARE_KILOMETER                           UnitOfMeasure = "SqkM"
	UnitOfMeasure_SQUARE_MILLIMETER                          UnitOfMeasure = "SqmM"
	UnitOfMeasure_SQUARE_MILE                                UnitOfMeasure = "Sqmi"
	UnitOfMeasure_SQUARE_YARD                                UnitOfMeasure = "Sqyd"
	UnitOfMeasure_US_DOLLARS                                 UnitOfMeasure = "USD"
	UnitOfMeasure_ARE                                        UnitOfMeasure = "a"
	UnitOfMeasure_ACRE                                       UnitOfMeasure = "ac"
	UnitOfMeasure_CENTILITER                                 UnitOfMeasure = "cL"
	UnitOfMeasure_CENTIMETER                                 UnitOfMeasure = "cM"
	UnitOfMeasure_HUNDREDWEIGHT                              UnitOfMeasure = "cwt"
	UnitOfMeasure_DAYS                                       UnitOfMeasure = "day"
	UnitOfMeasure_DRY_METRIC_TONS                            UnitOfMeasure = "dt"
	UnitOfMeasure_FOOT                                       UnitOfMeasure = "ft"
	UnitOfMeasure_GRAMS                                      UnitOfMeasure = "g"
	UnitOfMeasure_HECTARE                                    UnitOfMeasure = "ha"
	UnitOfMeasure_INCH                                       UnitOfMeasure = "in"
	UnitOfMeasure_HEAT_RATE                                  UnitOfMeasure = "kHR"
	UnitOfMeasure_KILOLITERS                                 UnitOfMeasure = "kL"
	UnitOfMeasure_KILOMETER                                  UnitOfMeasure = "kM"
	UnitOfMeasure_KILOWATT_MONTH                             UnitOfMeasure = "kW-M"
	UnitOfMeasure_KILOWATT_YEAR                              UnitOfMeasure = "kW-a"
	UnitOfMeasure_KILOWATT_DAY                               UnitOfMeasure = "kW-d"
	UnitOfMeasure_KILOWATT_HOUR                              UnitOfMeasure = "kW-h"
	UnitOfMeasure_KILOWATT_MINUTE                            UnitOfMeasure = "kW-min"
	UnitOfMeasure_KILOWATT_HOURS                             UnitOfMeasure = "kWh"
	UnitOfMeasure_KILOGRAMS                                  UnitOfMeasure = "kg"
	UnitOfMeasure_POUNDS                                     UnitOfMeasure = "lbs"
	UnitOfMeasure_MILLILITER                                 UnitOfMeasure = "mL"
	UnitOfMeasure_MILLIMETER                                 UnitOfMeasure = "mM"
	UnitOfMeasure_MILE                                       UnitOfMeasure = "mi"
	UnitOfMeasure_US_OUNCE                                   UnitOfMeasure = "oz"
	UnitOfMeasure_TROY_OUNCES                                UnitOfMeasure = "oz_tr"
	UnitOfMeasure_PIECE                                      UnitOfMeasure = "pc"
	UnitOfMeasure_US_PINT                                    UnitOfMeasure = "pt"
	UnitOfMeasure_GB_PINT                                    UnitOfMeasure = "pt_gb"
	UnitOfMeasure_US_QUART                                   UnitOfMeasure = "qt"
	UnitOfMeasure_GB_QUART                                   UnitOfMeasure = "qt_gb"
	UnitOfMeasure_METRIC_TONS                                UnitOfMeasure = "t"
	UnitOfMeasure_THERMS                                     UnitOfMeasure = "thm"
	UnitOfMeasure_TONS                                       UnitOfMeasure = "tn"
	UnitOfMeasure_TONS_OF_CARBON_DIOXIDE                     UnitOfMeasure = "tnCO2"
	UnitOfMeasure_YARD                                       UnitOfMeasure = "yd"
)

// UnsolicitedIndicator field enumeration values.
type UnsolicitedIndicator string

const (
	UnsolicitedIndicator_NO  UnsolicitedIndicator = "N"
	UnsolicitedIndicator_YES UnsolicitedIndicator = "Y"
)

// UpfrontPriceType field enumeration values.
type UpfrontPriceType string

const (
	UpfrontPriceType_PERCENTAGE   UpfrontPriceType = "1"
	UpfrontPriceType_FIXED_AMOUNT UpfrontPriceType = "3"
)

// Urgency field enumeration values.
type Urgency string

const (
	Urgency_NORMAL     Urgency = "0"
	Urgency_FLASH      Urgency = "1"
	Urgency_BACKGROUND Urgency = "2"
)

// UserRequestType field enumeration values.
type UserRequestType string

const (
	UserRequestType_LOG_ON_USER                    UserRequestType = "1"
	UserRequestType_LOG_OFF_USER                   UserRequestType = "2"
	UserRequestType_CHANGE_PASSWORD_FOR_USER       UserRequestType = "3"
	UserRequestType_REQUEST_INDIVIDUAL_USER_STATUS UserRequestType = "4"
	UserRequestType_REQUEST_THROTTLE_LIMIT         UserRequestType = "5"
)

// UserStatus field enumeration values.
type UserStatus string

const (
	UserStatus_LOGGED_IN                      UserStatus = "1"
	UserStatus_NOT_LOGGED_IN                  UserStatus = "2"
	UserStatus_USER_NOT_RECOGNISED            UserStatus = "3"
	UserStatus_PASSWORD_INCORRECT             UserStatus = "4"
	UserStatus_PASSWORD_CHANGED               UserStatus = "5"
	UserStatus_OTHER                          UserStatus = "6"
	UserStatus_FORCED_USER_LOGOUT_BY_EXCHANGE UserStatus = "7"
	UserStatus_SESSION_SHUTDOWN_WARNING       UserStatus = "8"
	UserStatus_THROTTLE_PARAMETERS_CHANGED    UserStatus = "9"
)

// ValuationMethod field enumeration values.
type ValuationMethod string

const (
	ValuationMethod_CDS_STYLE_COLLATERALIZATION_OF_MARKET_TO_MARKET_AND_COUPON ValuationMethod = "CDS"
	ValuationMethod_CDS_IN_DELIVERY                                            ValuationMethod = "CDSD"
	ValuationMethod_PREMIUM_STYLE                                              ValuationMethod = "EQTY"
	ValuationMethod_FUTURES_STYLE_MARK_TO_MARKET                               ValuationMethod = "FUT"
	ValuationMethod_FUTURES_STYLE_WITH_AN_ATTACHED_CASH_ADJUSTMENT             ValuationMethod = "FUTDA"
)

// ValueCheckAction field enumeration values.
type ValueCheckAction string

const (
	ValueCheckAction_DO_NOT_CHECK ValueCheckAction = "0"
	ValueCheckAction_CHECK        ValueCheckAction = "1"
	ValueCheckAction_BEST_EFFORT  ValueCheckAction = "2"
)

// ValueCheckType field enumeration values.
type ValueCheckType string

const (
	ValueCheckType_PRICE_CHECK          ValueCheckType = "1"
	ValueCheckType_NOTIONAL_VALUE_CHECK ValueCheckType = "2"
	ValueCheckType_QUANTITY_CHECK       ValueCheckType = "3"
)

// VenueType field enumeration values.
type VenueType string

const (
	VenueType_AUCTION_DRIVEN_MARKET    VenueType = "A"
	VenueType_CENTRAL_LIMIT_ORDER_BOOK VenueType = "B"
	VenueType_CLEARINGHOUSE            VenueType = "C"
	VenueType_DARK_ORDER_BOOK          VenueType = "D"
	VenueType_ELECTRONIC_EXCHANGE      VenueType = "E"
	VenueType_HYBRID_MARKET            VenueType = "H"
	VenueType_QUOTE_NEGOTIATION        VenueType = "N"
	VenueType_OFF_MARKET               VenueType = "O"
	VenueType_PIT                      VenueType = "P"
	VenueType_QUOTE_DRIVEN_MARKET      VenueType = "Q"
	VenueType_REGISTERED_MARKET        VenueType = "R"
	VenueType_VOICE_NEOTIATION         VenueType = "V"
	VenueType_EX_PIT                   VenueType = "X"
)

// VerificationMethod field enumeration values.
type VerificationMethod string

const (
	VerificationMethod_NON_ELECTRONIC VerificationMethod = "0"
	VerificationMethod_ELECTRONIC     VerificationMethod = "1"
)

// WorkingIndicator field enumeration values.
type WorkingIndicator string

const (
	WorkingIndicator_NO  WorkingIndicator = "N"
	WorkingIndicator_YES WorkingIndicator = "Y"
)

// YieldType field enumeration values.
type YieldType string

const (
	YieldType_AFTER_TAX_YIELD                                                                        YieldType = "AFTERTAX"
	YieldType_ANNUAL_YIELD                                                                           YieldType = "ANNUAL"
	YieldType_YIELD_AT_ISSUE                                                                         YieldType = "ATISSUE"
	YieldType_YIELD_TO_AVERAGE_LIFE_THE_YIELD_ASSUMING_THAT_ALL_SINKS                                YieldType = "AVGLIFE"
	YieldType_YIELD_TO_AVG_MATURITY                                                                  YieldType = "AVGMATURITY"
	YieldType_BOOK_YIELD                                                                             YieldType = "BOOK"
	YieldType_YIELD_TO_NEXT_CALL                                                                     YieldType = "CALL"
	YieldType_YIELD_CHANGE_SINCE_CLOSE                                                               YieldType = "CHANGE"
	YieldType_CLOSING_YIELD                                                                          YieldType = "CLOSE"
	YieldType_COMPOUND_YIELD                                                                         YieldType = "COMPOUND"
	YieldType_CURRENT_YIELD                                                                          YieldType = "CURRENT"
	YieldType_GVNT_EQUIVALENT_YIELD                                                                  YieldType = "GOVTEQUIV"
	YieldType_TRUE_GROSS_YIELD                                                                       YieldType = "GROSS"
	YieldType_YIELD_WITH_INFLATION_ASSUMPTION                                                        YieldType = "INFLATION"
	YieldType_INVERSE_FLOATER_BOND_YIELD                                                             YieldType = "INVERSEFLOATER"
	YieldType_MOST_RECENT_CLOSING_YIELD                                                              YieldType = "LASTCLOSE"
	YieldType_CLOSING_YIELD_MOST_RECENT_MONTH                                                        YieldType = "LASTMONTH"
	YieldType_CLOSING_YIELD_MOST_RECENT_QUARTER                                                      YieldType = "LASTQUARTER"
	YieldType_CLOSING_YIELD_MOST_RECENT_YEAR                                                         YieldType = "LASTYEAR"
	YieldType_YIELD_TO_LONGEST_AVERAGE_LIFE                                                          YieldType = "LONGAVGLIFE"
	YieldType_YIELD_TO_LONGEST_AVERAGE                                                               YieldType = "LONGEST"
	YieldType_MARK_TO_MARKET_YIELD                                                                   YieldType = "MARK"
	YieldType_YIELD_TO_MATURITY                                                                      YieldType = "MATURITY"
	YieldType_YIELD_TO_NEXT_REFUND                                                                   YieldType = "NEXTREFUND"
	YieldType_OPEN_AVERAGE_YIELD                                                                     YieldType = "OPENAVG"
	YieldType_PREVIOUS_CLOSE_YIELD                                                                   YieldType = "PREVCLOSE"
	YieldType_PROCEEDS_YIELD                                                                         YieldType = "PROCEEDS"
	YieldType_YIELD_TO_NEXT_PUT                                                                      YieldType = "PUT"
	YieldType_SEMI_ANNUAL_YIELD                                                                      YieldType = "SEMIANNUAL"
	YieldType_YIELD_TO_SHORTEST_AVERAGE_LIFE                                                         YieldType = "SHORTAVGLIFE"
	YieldType_YIELD_TO_SHORTEST_AVERAGE                                                              YieldType = "SHORTEST"
	YieldType_SIMPLE_YIELD                                                                           YieldType = "SIMPLE"
	YieldType_TAX_EQUIVALENT_YIELD                                                                   YieldType = "TAXEQUIV"
	YieldType_YIELD_TO_TENDER_DATE                                                                   YieldType = "TENDER"
	YieldType_TRUE_YIELD                                                                             YieldType = "TRUE"
	YieldType_YIELD_VALUE_OF_1_32_THE_AMOUNT_THAT_THE_YIELD_WILL_CHANGE_FOR_A_1_32ND_CHANGE_IN_PRICE YieldType = "VALUE1/32"
	YieldType_YIELD_VALUE_OF_1_32                                                                    YieldType = "VALUE1_32"
	YieldType_YIELD_TO_WORST                                                                         YieldType = "WORST"
)
