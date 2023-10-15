/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

import {gql} from '@urql/svelte';

export const SaveApiAccess = gql`
  mutation SaveApiAccess($input: ApiAccessInput!){
    saveApiAccess(input: $input)
  }`;


export const UpdateApiAccessKey = gql`
  mutation UpdateApiAccessKey($id: ID!, $key: String!){
    updateApiAccessKey(id: $id, key: $key)
  }`;

export const UpdateApiAccessSecret = gql`
    mutation UpdateApiAccessSecret($id: ID!, $secret: String!){
        updateAPIAccessSecret(id: $id, secret: $secret)
    }`;


export const QryApiAccess = gql`
  query ApiAccess($search: String, $page: PageInput!) {
    apiAccess(search: $search, page: $page) {
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
          url
          username
        }
      }
    }
  }`;


