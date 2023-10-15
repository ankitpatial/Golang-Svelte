import { gql } from '@urql/svelte';

export const QrySurveys = gql`
  query Surveys($search: String, $page: PageInput!, $orderBy: SurveyOrder) {
    surveys(search: $search, page: $page, orderBy: $orderBy) {
      totalCount
      edges {
        node {
          id
          date
          slot
          from
          to
          name
          phone
          address
          notes
          status
        }
      }
    }
  }
`;

export const QrySurveySlotAvailability = gql`
  query SurveySlotAvailability($date: String!) {
    surveySlotAvailability(date: $date) {
      id
      name
      available
    }
  }
`;

export const MutSurveyRequest = gql`
  mutation SurveyRequest($date: String!, $slotID: ID!) {
    surveyRequest(date: $date, slotID: $slotID)
  }
`;

export const MutSurveyReserve = gql`
  mutation SurveyReserve($input: SurveyInput!) {
    surveyReserve(input: $input) {
      id
      date
      slot
      from
      to
      name
      phone
      address
      notes
      status
    }
  }
`;

export const QrySurveyDetails = gql`
  query SurveyDetails($id: ID!) {
    surveyDetails(id: $id) {
      id
      date
      slot
      from
      to
      name
      phone
      address
      notes
      status
    }
  }
`;
