import {gql} from '@urql/svelte';

export const QryApiUsers = gql`
  query ApiUsers($page: PageInput! $where: ApiUserWhereInput) {
    apiUsers(page: $page where: $where) {
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
          username
          active
        }
      }
    }
  }`;

export const QryApiUser = gql`
  query ApiUser($id: ID!) {
    apiUser(id: $id) {
      id
      username
      active
      cbApiAuth
      cbApiUrl
      cbApiUser
      cbApiPwd
      cbApiToken
      cbApiEndpoints
    }
  }`;


export const AddApiUser = gql`
  mutation AddApiUser($username: String!) {
    addApiUser(username: $username)
  }`;

export const EditApiUser = gql`
  mutation EditApiUser($input: ApiUserInput!) {
    editApiUser(input: $input)
  }`;

export const ChangeApiUserStatus = gql`
  mutation ChangeApiUserStatus($id: ID! $isActive: Boolean!) {
    changeApiUserStatus(id: $id isActive: $isActive)
  }`;

export const RefreshApiUserPwd = gql`
  mutation RefreshApiUserPwd($id: ID!) {
    refreshApiUserPwd(id: $id)
  }`;
