import {gql} from "@urql/svelte";
import {page} from "./fragments";


const videoInfo = gql`
  fragment videoInfo on TrainingVideo {
    id
    title
    description
    posterURL
    videoURL
  }`


export const CreateTrainingCourse = gql`
  mutation CreateTrainingCourse($name: String!){
    createTrainingCourse(name: $name) {
      id
      name
    }
  }`;

export const PartnerTrainingVideoAccess = gql`
  mutation PartnerTrainingVideoAccess($partnerID: ID!, $videoID: ID!, $enabled: Boolean!) {
    partnerTrainingVideoAccess(partnerID: $partnerID, videoID: $videoID, enabled: $enabled)
  }`;


export const SaveTrainingVideo = gql`
  mutation SaveTrainingVideo($inp: InputTrainingVideo!) {
    saveTrainingVideo(inp: $inp)
  }`;


export const QryTrainingCourses = gql`
  query TrainingCourses($search: String, $page: PageInput!) {
    trainingCourses(search: $search, page: $page) {
      totalCount
      pageInfo {
        ...page
      }
      edges {
        cursor
        node {
          id
          name
        }
      }
    }
  }${page}`;

export const QryTrainingVideoKinds = gql`
  query TrainingVideoKinds {
    trainingVideoKinds {
      id
      name
    }
  }`

export const QryTrainingVideoCourses = gql`
  query TrainingVideoCourses($kind: TrainingType!, $partnerID: ID,  $pageSize: Int) {
    trainingVideoCourses(kind: $kind, partnerID: $partnerID, pageSize: $pageSize) {
      id
      name
      videos {
        totalCount
        pageInfo {
          ...page
        }
        edges {
          cursor
          node {
            ...videoInfo
            assigned
          }
        }
      }
    }
  } ${page} ${videoInfo}`;

export const QryTrainingVideos = gql`
  query TrainingVideos($kind: TrainingType!, $courseID: ID!, $search: String, $partnerID: ID, $page: PageInput!) {
    trainingVideos(kind: $kind, courseID: $courseID, search: $search, partnerID: $partnerID, page: $page) {
      totalCount
      pageInfo {
        ...page
      }
      edges {
        cursor
        node {
          ...videoInfo
          kind
          assigned
        }
      }
    }
  }${page} ${videoInfo}`;


export const QryMyTrainingVideoKinds = gql`
  query MyTrainingVideoKinds {
    myTrainingVideoKinds {
      id
      name
    }
  }`

export const QryMyTrainingVideoCourses = gql`
  query MyTrainingVideoCourses($kind: TrainingType!,  $pageSize: Int) {
    myTrainingVideoCourses(kind: $kind,  pageSize: $pageSize) {
      id
      name
      videos {
        totalCount
        pageInfo {
          ...page
        }
        edges {
          cursor
          node {
            ...videoInfo
          }
        }
      }
    }
  } ${page} ${videoInfo}`;

export const QryMyTrainingVideos = gql`
  query MyTrainingVideos($kind: TrainingType!, $courseID: ID!, $search: String, $page: PageInput!) {
    myTrainingVideos(kind: $kind, courseID: $courseID, search: $search, page: $page) {
      totalCount
      pageInfo {
        ...page
      }
      edges {
        cursor
        node {
          ...videoInfo
          kind
        }
      }
    }
  }${page} ${videoInfo}`;
