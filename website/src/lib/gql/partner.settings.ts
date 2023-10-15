import {gql} from '@urql/svelte';


export const SavePartnerMobileSettings = gql`
  mutation SavePartnerMobileSettings($id: ID!, $inp: InputMobileAppSettings!) {
    savePartnerMobileSettings( id: $id, inp: $inp)
  }`;


export const QryPartnerMobileSettings = gql`
  query PartnerMobileSettings($id: ID!) {
    partnerMobileSettings( id: $id) {
      logoURL
      primaryColor
      hideTabs
    }
  }`;
