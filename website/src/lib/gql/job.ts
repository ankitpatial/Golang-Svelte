import { gql } from '@urql/svelte';
import { homeOwner, material, page, salesRep } from './fragments';

//
// Mutation ==>
//
export const AssignPartnerToJob = gql`
	mutation AssignPartnerToJob($jobID: ID!, $partnerID: ID!) {
		assignPartnerToJob(jobID: $jobID, partnerID: $partnerID)
	}
`;

//
// Query ==>
//

export const QryJob = gql`
	query Job($id: ID!) {
		job(id: $id) {
			id
			createdAt
			updatedAt
			homeOwner {
				...homeOwner
        latitude
        longitude
			}
			companyName
			salesRep {
				...salesRep
			}
			estimate {
				...material
				partial
				measurementType
				priceSummary
			}
			price
			workOrderPrice
			creatorName
		}
	}
	${homeOwner}
	${salesRep}
	${material}
`;

export const QryUnassignedJobs = gql`
	query UnassignedJobs($page: PageInput!, $search: String, $betweenDates: [String!], $orderBy: JobOrder) {
		unassignedJobs(search: $search, betweenDates: $betweenDates, page: $page, orderBy: $orderBy) {
			totalCount
			pageInfo {
				...page
			}
			edges {
				node {
					id
					createdAt
					homeOwner {
						...homeOwner
					}
					companyName
					epcName
					price
					workOrderPrice
				}
			}
		}
	}
	${page}
	${homeOwner}
`;

export const QryAssignedJobs = gql`
	query AssignedJobs(
		$progress: JobProgress
		$search: String
		$betweenDates: [String!]
		$page: PageInput!
		$orderBy: JobOrder
	) {
		assignedJobs(progress: $progress, page: $page, search: $search, betweenDates: $betweenDates, orderBy: $orderBy) {
			totalCount
			pageInfo {
				...page
			}
			edges {
				node {
					id
					createdAt
					homeOwner {
						...homeOwner
					}
					companyName
					epcName
					contractor {
						id
						name
					}
					price
					workOrderPrice
					progress
					progressFlagged
				}
			}
		}
	}
	${page}
	${homeOwner}
`;

export const QryPendingInvoiceJobs = gql`
	query PendingInvoiceJobs($page: PageInput!, $search: String, $betweenDates: [String!], $orderBy: JobOrder) {
		jobsByProgress(status: Invoiced, page: $page, search: $search, betweenDates: $betweenDates, orderBy: $orderBy) {
			totalCount
			pageInfo {
				...page
			}
			edges {
				node {
					id
					progress
					progressAt
					homeOwner {
						...homeOwner
					}
					salesRep {
						...salesRep
					}
					companyName
					price
					installDate
					inspectionDate
					contractor {
						id
						name
					}
				}
			}
		}
	}
	${page}
	${homeOwner}
	${salesRep}
`;

export const QryPartnerJobStats = gql`
	query PartnerJobStats($partnerType: PartnerType!, $search: String, $skip: Int!, $take: Int!) {
		partnerJobStats(partnerType: $partnerType, search: $search, skip: $skip, take: $take) {
			id
			name
			status
			newCount
			newCountFlagged
			contactedCount
			contactedCountFlagged
			confirmedCount
			confirmedCountFlagged
			permittingCount
			permittingCountFlagged
			scheduledCount
			scheduledCountFlagged
			inProgressCount
			inProgressCountFlagged
			installedCount
			installedCountFlagged
			invoicedCount
			invoicedCountFlagged
			delayedCount
			total
			totalFlagged
		}
	}
`;

export const QryPartnerJobs = gql`
	query PartnerJobs(
		$partnerID: ID!
		$search: String
		$flagged: Boolean
		$progress: JobProgress
		$dates: [String!]
		$page: PageInput!
	) {
		partnerJobs(
			partnerID: $partnerID
			search: $search
			progress: $progress
			flagged: $flagged
			dates: $dates
			page: $page
		) {
			totalCount
			pageInfo {
				...page
			}
			edges {
				node {
					id
					homeOwner {
						...homeOwner
					}
					salesRep {
						...salesRep
					}
					companyName
					price
					progress
					progressAt
					progressFlagged
					permitRequired
					installDate
					inspectionDate
					inspectionRequired
					workOrderPrice
					contractor {
						id
						name
					}
					epcName
					estimate {
						...material
					}
				}
			}
		}
	}
	${homeOwner}
	${salesRep}
	${material}
	${page}
`;

export const QryJobEstimates = gql`
	query JobEstimates($jobID: ID!) {
		jobEstimates(jobID: $jobID) {
			id
			status
			estimates {
				id
				totalSquares
				primaryPitch
				price
				priceSummary
				bounds {
					lat
					lng
				}
			}
		}
	}
`;

export const QryNearmapResponse = gql`
	query NearmapResponse($id: ID!, $respID: ID!) {
		nearmapResponse(id: $id, respID: $respID) {
			raw
			detail {
				price
				primaryPitchInDegrees
				primaryPitch
				tileArea
				tileRatio
				shingleArea
				shingleRatio
				metalArea
				metalRatio
				flatArea
				flatRatio
				gableArea
				gableRatio
				hipArea
				hipRatio
				dutchGableArea
				dutchGableRatio
				turretArea
				turretRatio
				dominantRoofMaterial
				dominantRoofMaterialID
				roofCount
				totalUnclippedArea
				roofMaterialRatioTotal
				roofTypeRatioTotal
				roofMaterialSurfaceAreaTotal
				roofTypeSurfaceAreaTotal
				treeOverhangRatioPrimaryRoof
				treeOverhangConfidencePrimaryRoof
				treeOverhangPresenceConfidence
				treeOverhangAreaPrimaryRoof
				treeOverhangTotalClippedArea
				treeOverhangTotalUnClippedArea
				treeOverhangPresent
				treeOverhangCount
			}
		}
	}
`;
