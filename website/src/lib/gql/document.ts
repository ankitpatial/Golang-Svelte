import {gql} from '@urql/svelte';

export const document = gql`
  fragment document on Document {
    id
    key
    name
    contentType
    contentSize
    ready
  }
`;

export const DocumentUploadUrl = gql`
  mutation documentUploadUrl(
    $folder: DocumentFolder!
    $dir: String!
    $section: DocumentSection!
    $name: String!
    $fileName: String!
    $fileType: String!
    $fileSize: Int64!
    $overwrite: Boolean!
  ) {
    documentUploadUrl(
      doc: {
        folder: $folder
        dir: $dir
        section: $section
        name: $name,
        fileName: $fileName
        contentType: $fileType,
        contentSize: $fileSize
        overwrite: $overwrite
      }
    ) {
      ...document
      uploadUrl
      meta
    }
  } ${document}`;


export const PublicDataUploadUrl = gql`
  mutation PublicDataUploadUrl(
    $dir: ID!
    $name: String!
    $section: PublicDataSection!
    $fileName: String!
    $fileType: String!
    $fileSize: Int64!
  ) {
    publicDataUploadUrl(
      entityID: $dir,
      section: $section
      doc: { name: $name, fileName: $fileName, contentType: $fileType, contentSize: $fileSize}
    ) {
      id
      key
      meta
      publicUrl
      uploadUrl
    }
  }`;

export const PartnerDocUploadUrl = gql`
  mutation PartnerDocUploadUrl(
    $partnerID: ID!
    $section: DocumentSection!
    $name: String!
    $fileName: String!
    $fileType: String!
    $fileSize: Int64!
  ) {
    partnerDocUploadUrl(
      partnerID: $partnerID
      section: $section
      doc: { name: $name, fileName: $fileName, contentType: $fileType, contentSize: $fileSize}
    ) {
      ...document
      uploadUrl
      meta
    }
  } ${document}`;

export const JobDocUploadUrl = gql`
  mutation JobDocUploadUrl(
    $jobID: ID!
    $section: DocumentSection!
    $name: String!
    $fileName: String!
    $fileType: String!
    $fileSize: Int64!
  ) {
    jobDocUploadUrl(
      jobID: $jobID
      section: $section
      doc: {name: $name, fileName: $fileName, contentType: $fileType, contentSize: $fileSize}
    ) {
      ...document
      uploadUrl
      meta
    }
  } ${document}`;


export const DeleteDoc = gql`
  mutation  DeleteDoc($id: ID!){
    deleteDoc(id: $id)
  }`;

export const QryJobDocs = gql`
  query JobDocs($jobID: ID!) {
    jobDocs(jobID: $jobID) {
      id
      name
      filename
      section
      meta
    }
  }`;
