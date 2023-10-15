import { gql } from '@urql/svelte';

export const NotificationRead = gql`
  mutation NotificationRead($messageID: ID!) {
    notificationRead(messageID: $messageID)
  }`;


export const QryMyUnreadNotificationsCount = gql`
  query MyUnreadNotificationsCount {
    myUnreadNotificationsCount
  }`;

export const QryMyNotifications = gql`
  query MyNotifications($page: PageInput!) {
    myNotifications(page: $page) {
      totalCount
      pageInfo {
        hasPreviousPage
        startCursor
        hasNextPage
        endCursor
      }
      edges {
        node {
          id
          channel
          topic
          refID
          title
          message
          from
          unread
          createdAt
        }
      }
    }
  }`;
