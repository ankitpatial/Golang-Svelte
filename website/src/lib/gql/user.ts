import {gql} from '@urql/svelte';
import {userBasic} from './fragments';

export const QryUser = gql`
  query User($id: ID!) {
    user(id: $id) {
      id
      email
      firstName
      lastName
      phone
      role
      status
      note
    }
  }`;

export const QryUsers = gql`
  query Users($page: PageInput! $where: UserWhereInput) {
    users(page: $page where: $where) {
      totalCount
      pageInfo {
        startCursor
        hasNextPage
        endCursor
        hasPreviousPage
      }
      edges {
        node {
          ...userBasic
        }
      }
    }
  }${userBasic}`;

export const QryUsersSearch = gql`
  query usersSearch($search: String) {
    usersSearch(search: $search) {
      id
      firstName
      lastName
      email
      phone
      partnerID
      partnerName
      partnerContactTypeID
      partnerContactTitle
    }
  }`;

export const QryMyCompanyUsers = gql`
  query myCompanyUsers($search: String, $page: PageInput!) {
    myCompanyUsers(search: $search, page: $page) {
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
          firstName
          lastName
          email
          phone
          picture
        }
      }
    }
  }`;

export const CreateUser = gql`
  mutation CreateUser($input: CreateUserInput!) {
    createUser(input: $input)
  }`;

export const UpdateUser = gql`
  mutation UpdateUser($input: UpdateUserInput!) {
    updateUser(input: $input)
  }`;

export const SaveNotifySettings = gql`
  mutation SaveNotifySettings($id: ID! $topicID: String! $email: Boolean!) {
    saveNotifySettings(userID: $id topicID: $topicID email: $email)
  }`;

export const QryUserNotifySettings = gql`
  query UserNotifySettings($id: ID!) {
    userNotifySettings(id: $id) {
      id
      topic
      receiveEmail
      receiveSMS
    }
  }`;
