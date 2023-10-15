/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2023.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package eagleview

type Status uint8
type SubStatus uint8

const ProviderName = "EagleView"

const (
	StatusCreated   Status = 1
	StatusInProcess Status = 2
	StatusPending   Status = 3
	StatusClosed    Status = 4
	StatusCompleted Status = 5

	SubStatusCreated               SubStatus = 1  // NOT MENTIONED in EagleView docs but sent over on GetReport
	SubStatusUnderReview           SubStatus = 5  // Report is being reviewed.Related Status: InProcess.
	SubStatusProcessStarted        SubStatus = 6  // Measurement process has been started.Related Status: InProcess.
	SubStatusNeedToID              SubStatus = 8  // Structure to measure needs to be identified.Related Status: Pending.
	SubStatusDuplicate             SubStatus = 9  // There exists another report on the customer's account for the same structure. Related Status: Closed.
	SubStatusCanceledByClient      SubStatus = 10 //  Report canceled by the client.Related Status: Closed, Completed.
	SubStatusPoorImages            SubStatus = 11 // Report cannot be generated from the images available.Related Status: Closed.
	SubStatusNoID                  SubStatus = 12 // Unable to correctly identify the structure. Related Status: Closed.
	SubStatusReportTypeChange      SubStatus = 14 // Report type needs to be/has been changed.Related Status: Pending, Closed.
	SubStatusNoImages              SubStatus = 15 // Report cannot be generated due to no images being available.Related Status: Closed.
	SubStatusOther                 SubStatus = 16 // Other.Related Status: Closed.
	SubStatusCardRejected          SubStatus = 18 // Credit card transaction has been rejected.Related Status: Closed.
	SubStatusSent                  SubStatus = 19 // Report has been delivered successfully.Related Status: Completed.
	SubStatusSent20                SubStatus = 20 // Report has been delivered successfully.Related Status: Completed.
	SubStatusCustomerResponse      SubStatus = 21 // Awaiting customer response.Related Status: Pending.
	SubStatusWrongHouse            SubStatus = 24 // The wrong house was measured for the report.Related Status: Closed.
	SubStatusReadyToSend           SubStatus = 26 // Report is ready to send to the customer.Related Status: InProcess.
	SubStatusCreditCardFailure     SubStatus = 27 // Credit card transaction is failing.Related Status: Pending.
	SubStatusUnSketchable          SubStatus = 28 // Unable to generate a sketch for this report.Related Status: Closed.
	SubStatusAddressVerified       SubStatus = 29 // Address has been identified.Related Status: InProcess.
	SubStatusNotSent               SubStatus = 30 // Report has not been sent.Related Status: Completed.
	SubStatusProcessing            SubStatus = 31 // Report is being processed. Related Status: Created.
	SubStatusMeasurementData       SubStatus = 32 // Report is being measured.Related Status: Pending.
	SubStatusUnsupportedStructure  SubStatus = 38 // Report type selected does not allow the structure to be measured (e.g.residential report cannot measure a commercial structure).Related Status: Closed.
	SubStatusSubmitted             SubStatus = 40 // Report has been submitted.Related Status: Completed.
	SubStatusReadyToCapturePayment SubStatus = 43 // Payment is being captured for the report.Related Status: InProcess.
	SubStatusPendingClosure        SubStatus = 49 // Report is pending closure.Related Status: InProcess.
	SubStatusClosureDenied         SubStatus = 50 // Report closure has been denied. Related Status: InProcess.
	SubStatusCancelled             SubStatus = 55
)

func (s Status) String() string {
	switch s {
	case StatusCreated:
		return "Created"
	case StatusInProcess:
		return "InProcess"
	case StatusPending:
		return "Pending"
	case StatusClosed:
		return "Closed"
	case StatusCompleted:
		return "Completed"
	default:
		return ""
	}
}

func (s SubStatus) String() string {
	switch s {
	case SubStatusCreated:
		return "Created"
	case SubStatusUnderReview:
		return "UnderReview"
	case SubStatusProcessStarted:
		return "ProcessStarted"
	case SubStatusNeedToID:
		return "NeedToID"
	case SubStatusDuplicate:
		return "Duplicate"
	case SubStatusCanceledByClient:
		return "CanceledByClient"
	case SubStatusPoorImages:
		return "PoorImages"
	case SubStatusNoID:
		return "NoID"
	case SubStatusReportTypeChange:
		return "ReportTypeChange"
	case SubStatusNoImages:
		return "NoImages"
	case SubStatusOther:
		return "Other"
	case SubStatusCardRejected:
		return "CardRejected"
	case SubStatusSent, SubStatusSent20:
		return "Sent"
	case SubStatusCustomerResponse:
		return "CustomerResponse"
	case SubStatusWrongHouse:
		return "WrongHouse"
	case SubStatusReadyToSend:
		return "ReadyToSend"
	case SubStatusCreditCardFailure:
		return "CreditCardFailure"
	case SubStatusUnSketchable:
		return "UnSketchable"
	case SubStatusAddressVerified:
		return "AddressVerified"
	case SubStatusNotSent:
		return "NotSent"
	case SubStatusProcessing:
		return "Processing"
	case SubStatusMeasurementData:
		return "MeasurementData"
	case SubStatusUnsupportedStructure:
		return "UnsupportedStructure"
	case SubStatusSubmitted:
		return "Submitted"
	case SubStatusReadyToCapturePayment:
		return "ReadyToCapturePayment"
	case SubStatusPendingClosure:
		return "PendingClosure"
	case SubStatusClosureDenied:
		return "ClosureDenied"
	case SubStatusCancelled:
		return "Cancelled"
	default:
		return ""
	}
}
