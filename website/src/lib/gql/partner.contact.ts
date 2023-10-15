import {gql} from "@urql/svelte";
import {page} from "./fragments";

export const SavePartnerContact = gql`
  mutation SavePartnerContact($partnerID: ID!, $contact: PartnerContactInput!){
    savePartnerContact(
      partnerID: $partnerID
      contact: $contact
    )}`

export const PartnerContactSendPwdResetEmail = gql`
    mutation PartnerContactSendPwdResetEmail($partnerID: ID! $userID: ID!){
        partnerContactSendPwdResetEmail(
            partnerID: $partnerID
            userID: $userID
        )}`

export const QryPartnerContacts = gql`
  query partnerContacts(
    $partnerID: ID!
    $search: String
    $page: PageInput!
  ) {
    partnerContacts(
      partnerID: $partnerID
      search: $search
      page: $page
    ) {
      totalCount
      pageInfo {
        ...page
      }
      edges {
        node {
          id
          userID
          firstName
          lastName
          email
          phone
          role
          accountStatus
          picture
          type
          title
          description
        }
      }
    }
  } ${page}`;
