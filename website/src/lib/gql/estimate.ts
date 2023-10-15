import {gql} from "@urql/svelte";
import {homeOwner, material, page, salesRep} from "./fragments";

export const CreateEstimate = gql`
  mutation createEstimate($input: CreateEstimateInput!) {
    createEstimate(input: $input)
  }`;

export const ApproveEstimate = gql`
  mutation ApproveEstimate($input: ApproveEstimateInput!) {
    approveEstimate(input: $input)
  }`;


export const DenyEstimate = gql`
  mutation DenyEstimate($input: DenyEstimateInput!) {
    denyEstimate(input: $input)
  }`;

export const RemoveDenied = gql`
  mutation RemoveDenied($id: ID!) {
    removeDenied(id: $id)
  }`;


export const TestPricing = gql`
  mutation TestPricing($job: CreateEstimateInput!, $measure: [Measurement!]!) {
    testPricing(job: $job, measure: $measure) {
      total
      summary
    }
  }`;

export const QryEstimate = gql`
  query estimate($id: ID!) {
    estimate(id: $id) {
      id
      createdAt
      status
      companyName
      ...material
      partial
      measurementType
      price
      priceSummary
      totalSquares
      primaryPitch
      failureReason
      bounds {
        lat
        lng
      }
      homeOwner {
        ...homeOwner
        latitude
        longitude
      }
      salesRep {
        ...salesRep
      }
      pdf {
        id
        key
        name
        contentType
        contentSize
        ready
      }
      creatorName
    }
  } ${material} ${homeOwner} ${salesRep}`


export const QryEstimates = gql`
  query Estimates(
    $status: EstimateStatus, $search: String, $dates: [String!], $page: PageInput!,
  ) {
    estimates(status: $status, search: $search, dtRange: $dates, page: $page) {
      totalCount
      pageInfo {
        ...page
      }
      edges {
        node {
          id
          createdAt
          status
          companyName
          measurementType
          price
          failureReason
          homeOwner {
            ...homeOwner
          }
          salesRep {
            ...salesRep
          }
        }
      }
    }
  }
  ${page} ${homeOwner} ${salesRep}`
