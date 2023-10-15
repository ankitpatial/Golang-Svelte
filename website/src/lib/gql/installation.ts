import { gql } from '@urql/svelte';
import { page } from '$lib/gql/fragments';

export const BookInstallation = gql`
  mutation BookInstallation($type: InstallationType!, $pkgID: ID!, $productID: ID,$owner: InstallationOwnerInput!) {
    bookInstallation(type: $type, pkgID: $pkgID, productID: $productID, owner: $owner)
  }`;

export const ApproveInstallation = gql`
  mutation ApproveInstallation($input: InstallationApproveInput!) {
    approveInstallation(input: $input)
  }`;

export const DenyInstallation = gql`
  mutation DenyInstallation($id: ID!, $reason: String!) {
    denyInstallation(id: $id, reason: $reason)
  }`;

export const UndoDenyInstallation = gql`
  mutation UndoDenyInstallation($id: ID!) {
    undoDenyInstallation(id: $id)
  }`;

export const QryPendingInstallations = gql`
  query PendingInstallations($type: InstallationType!, $approval: Approval, $search: String, $betweenDates: [String!], $page: PageInput!){
    pendingInstallations(type: $type, approval: $approval, search:  $search, betweenDates: $betweenDates, page: $page) {
      totalCount
      pageInfo {
        ...page
      },
      edges {
        node {
          id
          ownerName
          ownerAddress
          ownerEmail
          ownerPhone
          pkg
          price
          approval
          status
          createdAt
        }
      }
    }
  }${page}`;

export const QryApprovedInstallations = gql`
  query ApprovedInstallations($type: InstallationType!, $status: InstallationStatus, $search: String, $betweenDates: [String!], $page: PageInput!){
    approvedInstallations(type: $type, status: $status, search: $search, betweenDates: $betweenDates,  page: $page) {
      totalCount
      pageInfo {
        ...page
      },
      edges {
        node {
          id
          ownerName
          ownerAddress
          pkg
          price
          approval
          status
          approvalAt
        }
      }
    }
  }${page}`;

