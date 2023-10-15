import {gql} from '@urql/svelte';

export const MarkServiceArea = gql`
  mutation MarkServiceArea($id: ID!, $value: Boolean!) {
    markServiceArea(id: $id, value: $value)
  }`;

export const QryAllServiceAreas = gql`
  query AllServiceAreas {
    allServiceAreas {
      id
      name
      cities {
        id
        zip
        name
      }
    }
  }`;

export const QryServiceStates = gql`
  query ServiceStates($q: String!) {
    serviceStates(q: $q) {
      id
      name
      cities {
        id
        zip
        name
      }
    }
  }`;

export const QryStates = gql`
  query States($q: String!) {
    states(q: $q) {
      id
      name
      cities {
        id
        zip
        name
      }
    }
  }`;

export const QryCities = gql`
  query Cities($state: String!, $q: String!, $skip: Int!, $take: Int!) {
    cities(state: $state, q: $q, skip: $skip, take: $take) {
      id
      zip
      name
    }
  }`;

