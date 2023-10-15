/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2023.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package feedback

type Event struct {
	Type     string        `json:"eventType"` // Send | Delivery | Bounce
	Mail     MailInfo      `json:"mail"`
	Bounce   *BounceInfo   `json:"bounce"`
	Delivery *DeliveryInfo `json:"delivery"`
}

type MailInfo struct {
	To           []string `json:"destination"`
	Sender       string   `json:"source"`
	SendingAccId string   `json:"sendingAccountId"`
	MessageID    string   `json:"messageId"`
}

type BounceInfo struct {
	Type         string             `json:"bounceType"` // Permanent
	SubType      string             `json:"bounceSubType"`
	Recipients   []*BounceRecipient `json:"bouncedRecipients"`
	TimeStamp    string             `json:"timestamp"`
	ReportingMTA string             `json:"reportingMTA"`
}

type BounceRecipient struct {
	Email      string `json:"emailAddress"`
	Action     string `json:"action"`
	Status     string `json:"status"`
	Diagnostic string `json:"diagnosticCode"`
}

type DeliveryInfo struct {
	Recipients     []string `json:"recipients"`
	ProcessingTime uint     `json:"processingTimeMillis"`
	TimeStamp      string   `json:"timestamp"`
	ReportingMTA   string   `json:"reportingMTA"`
}
