/*
 * Copyright (c) 2022. SimSaw Software Private Limited.
 * Unauthorized copying of this file, via any medium is strictly prohibited. Proprietary and confidential.
 * Author: Ankit Patial
 * Dated:  17/05/22, 4:07 PM
 */

package infra

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewCertificate(stack constructs.Construct, certID string, domain string, zone awsroute53.IHostedZone) awscertificatemanager.Certificate {
	// api https certificate, email validation will sent over
	cert := &awscertificatemanager.CertificateProps{
		DomainName: &domain,
	}

	if zone != nil {
		cert.Validation = awscertificatemanager.CertificateValidation_FromDns(zone)
	}

	return awscertificatemanager.NewCertificate(stack, jsii.String(certID), cert)
}
