import {gql} from '@urql/svelte';

export const QryAuditLogs = gql`
  query AuditLogs($search: String, $page: PageInput!, $orderBy: AuditLogOrder) {
    auditLogs(search: $search, page: $page, orderBy: $orderBy) {
      totalCount
      pageInfo {
        startCursor
        hasNextPage
        endCursor
        hasPreviousPage
      }
      edges {
        node {
          id
          createdAt
          action
          description
          ip
          user {
            firstName
            lastName
          }
          apiUser {
            username
          }
        }
      }
    }
  }`;

export const QryApiUserAuditLogs = gql`
  query ApiUserAuditLogs($id: ID!, $search: String, $page: PageInput!, $orderBy: AuditLogOrder) {
    apiUserAuditLogs(id: $id, search: $search, page: $page, orderBy: $orderBy) {
      totalCount
      pageInfo {
        startCursor
        hasNextPage
        endCursor
        hasPreviousPage
      }
      edges {
        node {
          id
          createdAt
          action
          description
          ip
        }
      }
    }
  }`;

