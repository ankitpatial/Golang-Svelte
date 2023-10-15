import { gql } from '@urql/svelte';

export const ForgotPassword = gql`
  mutation ForgotPassword($email: String!){
    accForgotPwd(email: $email)
  }`;

export const SetUserPwd = gql`
  mutation SetUserPwd($userID: ID!, $pwd: String!, $confirmPwd: String!){
    setUserPwd(userID: $userID, pwd: $pwd, confirmPwd: $confirmPwd)
  }`;

export const UpdateProfile = gql`
  mutation UpdateProfile($inp: InputUserProfile!) {
    updateProfile(input: $inp)
  }`;

export const QryMe = gql`
  query Me {
    me {
      id
      email
      firstName
      lastName
      status
      role
      phone
      picture
      partner {
        id
        type
        partnerName
        status
        contactType
        role
        mobileAppSettings {
          hideTabs
        }
      }
      token
      isAdmin
      isCompanyAdmin
    }
  }`;

export const QryUserEmailAvailable = gql`
  query UserEmailAvailable($id: String!, $email: String!){
    userEmailAvailable(id: $id, email: $email)
  }`;
