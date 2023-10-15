import {gql} from '@urql/svelte';

export const page = gql`
  fragment page on PageInfo {
    startCursor
    hasNextPage
    endCursor
    hasPreviousPage
  }`

export const userBasic = gql`
  fragment userBasic on User {
    id
    email
    name
    phone
    role
    status
    picture
  }`;

export const homeOwner = gql`
  fragment homeOwner on HomeOwner {
    id
    firstName
    lastName
    email
    phone
    streetNumber
    streetName
    city
    state
    zip
  }`

// sales rep fragment
export const salesRep = gql`
  fragment salesRep on UserInfo {
    id
    firstName
    lastName
    email
    phone
  }`

export const createdBy = gql`
  fragment createdBy on UserInfo {
    id
    firstName
    lastName
  }`

// material fragment on Estimate
export const material = gql`
  fragment material on Estimate {
    currentMaterial
    newRoofingMaterial
    currentMaterialLowSlope
    newRoofingMaterialLowSlope
    redeck
    layers
    layer2Material
    layer3Material
  }`
