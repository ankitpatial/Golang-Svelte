/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

import { gql } from '@urql/svelte';
import { document } from '$lib/gql/document';
import { homeOwner, material, salesRep } from '$lib/gql/fragments';

const partner = gql`
	fragment partner on Partner {
		id
		createdAt
		type
		name
		address
		website
		isNationWide
		status
	}
`;

const contact = gql`
	fragment contact on PartnerContact {
		id
		userID
		type
		firstName
		lastName
		email
		phone
		otherEmail
		title
		description
	}
`;

export const SavePartner = gql`
	mutation SavePartner($input: PartnerInput!) {
		savePartner(input: $input)
	}
`;

export const InvitePartner = gql`
	mutation invitePartner($input: InvitePartnerInput!) {
		invitePartner(input: $input) {
			id
			type
			companyName
			contactID
			userID
			firstName
			lastName
			email
			phone
			createdAt
		}
	}
`;

export const SavePartnerCompletedSteps = gql`
	mutation SavePartnerCompletedSteps($partnerID: ID!, $step: Int!, $done: Boolean) {
		savePartnerCompletedSteps(partnerID: $partnerID, step: $step, done: $done)
	}
`;

export const SetPartnerActive = gql`
	mutation SetPartnerActive($partnerID: ID!, $value: Boolean!) {
		setPartnerActive(partnerID: $partnerID, value: $value)
	}
`;

export const SavePartnerContacts = gql`
	mutation savePartnerContacts($partnerID: ID!, $contacts: [PartnerContactInput!]!) {
		savePartnerContacts(partnerID: $partnerID, contacts: $contacts) {
			...contact
		}
	}
	${contact}
`;

export const SavePartnerOperations = gql`
	mutation savePartnerOperations($partnerID: ID!, $input: PartnerOperationInput!) {
		savePartnerOperations(partnerID: $partnerID, inp: $input)
	}
`;

export const SaveService = gql`
	mutation SaveService($id: ID!, $partnerID: ID!, $service: Service!, $active: Boolean!) {
		saveService(id: $id, partnerID: $partnerID, service: $service, active: $active)
	}
`;

export const SaveLeadTime = gql`
	mutation SaveLeadTime($partnerID: ID!, $asphalt: String, $metal: String, $tile: String) {
		saveLeadTime(partnerID: $partnerID, asphalt: $asphalt, metal: $metal, tile: $tile)
	}
`;

export const SaveServiceState = gql`
	mutation SaveServiceState($partnerID: ID!, $state: String!, $licNo: String, $expDate: Time, $proofDocID: String) {
		saveServiceState(partnerID: $partnerID, state: $state, licNo: $licNo, expDate: $expDate, proofDocID: $proofDocID)
	}
`;

export const SaveServiceCity = gql`
	mutation SaveServiceCity($partnerID: ID!, $postalID: ID!, $active: Boolean, $licNo: String, $proofDocID: String) {
		saveServiceCity(partnerID: $partnerID, postalID: $postalID, active: $active, licNo: $licNo, proofDocID: $proofDocID)
	}
`;

export const QryPartnerNameAvailable = gql`
	query PartnerNameAvailable($id: ID!, $type: PartnerType!, $name: String!) {
		partnerNameAvailable(id: $id, type: $type, name: $name)
	}
`;

export const QryPartner = gql`
	query Partner($id: ID!) {
		partner(id: $id) {
			...partner
			crewCount
			jobCapacity
			setupStepsCompleted
			asphaltLeadT
			metalLeadT
			tileLeadT
			contacts {
				...contact
				otherEmail
				title
				description
			}
		}
	}
	${partner}
	${contact}
`;

export const QrySolarPartner = gql`
	query SolarPartner($id: ID!) {
		partner(id: $id, type: SOLAR) {
			id
			type
			status
			name
			address
			website
			yearsInBusiness
			salesVolume
			financeOptions
			downPayment
			pifDate
			installInHouse
			epcOptions
			setupStepsCompleted
			contacts {
				...contact
				otherEmail
				title
				description
			}
		}
	}
	${contact}
`;

export const QryPartnerOptions = gql`
	query PartnerOptions($partnerID: ID!) {
		partnerOptions(partnerID: $partnerID) {
			type
			options {
				id
				name
			}
		}
	}
`;

export const QryPartners = gql`
	query Partners($search: String, $partnerType: PartnerType, $status: String, $page: PageInput!) {
		partners(search: $search, partnerType: $partnerType, status: $status, page: $page) {
			totalCount
			pageInfo {
				hasPreviousPage
				hasNextPage
				startCursor
				endCursor
			}
			edges {
				node {
					...partner
					crewCount
					jobCapacity
					isActive
				}
			}
		}
	}
	${partner}
`;

export const QryMyJob = gql`
	query MyJob($id: ID!) {
		myJob(id: $id) {
			id
			createdAt
			updatedAt
			companyName
			homeOwner {
				...homeOwner
			}
			salesRep {
				...salesRep
			}
			estimate {
				...material
				measurementType
			}
			price
			workOrderPrice
			contractor {
				id
				name
			}
		}
	}
	${homeOwner}
	${salesRep}
	${material}
`;

export const QryPartnerDocs = gql`
	query PartnerDocs($partnerID: ID!, $section: DocumentSection!) {
		partnerDocs(partnerID: $partnerID, section: $section) {
			...document
		}
	}
	${document}
`;

export const QryPartnerServiceStates = gql`
	query PartnerServiceStates($partnerID: ID!) {
		partnerServiceStates(partnerID: $partnerID) {
			id
			name
			licenseNo
			licenseExpDate
			expand
			cities {
				id
				active
				licenseNo
				licenseProof
				cityZip
				cityName
			}
		}
		partnerDocs(partnerID: $partnerID, section: Proof) {
			...document
		}
	}
	${document}
`;

export const QryPartnerServices = gql`
	query PartnerServices($partnerID: ID!) {
		partnerServices(partnerID: $partnerID) {
			id
			service
			description
			active
		}
	}
`;
