import {gql} from "@urql/svelte";

export const CreatJobNote = gql`
  mutation creatJobNote($jobID: ID!, $note: String!)  {
    creatJobNote(jobID: $jobID, note:  $note)
  }`

export const EditJobNote = gql`
  mutation editJobNote($jobID: ID!, $noteID: ID!, $note: String!)  {
    editJobNote(jobID: $jobID, noteID: $noteID, note:  $note)
  }`


export const QryJobNotes = gql`
  query JobNotes($jobID: ID!) {
    jobNotes(jobID: $jobID) {
      id
      note
      createdAt
      updatedAt
      creator {
        id
        firstName
        lastName
        picture
      }
    }
  }`;

