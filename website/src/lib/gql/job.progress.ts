import { gql } from '@urql/svelte';

//
// Mutation ==>
//
export const JobProgressUpdate = gql`
  mutation JobProgressUpdate(
    $id: ID!, $step: JobProgress!, $stepComplete: Boolean!, $note: String!,$data: ProgressInput
  ) {
    jobProgressUpdate(id: $id, step: $step, stepComplete: $stepComplete, note: $note, data: $data)
  }`;

export const JobDelay = gql`
  mutation JobDelay($id: ID!, $flag: Boolean!, $reason: String!) {
    jobDelay(id: $id, flag: $flag, reason: $reason)
  }`;

//
// Query ==>
//
export const QryJobProgress = gql`
  query JobProgress($id: ID!, $page: PageInput!) {
    jobProgress(id: $id, page: $page) {
      totalCount
      pageInfo {
        startCursor
        endCursor
        hasNextPage
        hasPreviousPage
      }
      edges {
        node {
          id
          status
          statusAt
          note
        }
      }
    }
  }`;

export const QryJobCompletedProgress = gql`
  query JobCompletedProgress($id: ID!) {
    jobCompletedProgress(id: $id) {
      id
      status
      statusAt
      complete
      note
    }
  }`;



