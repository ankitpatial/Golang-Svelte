import {gql} from "@urql/svelte";

export const QryCounts = gql`
  query {
    counts{
      estimateCount
      unassignedJobCount
      assignedJobCount
      paymentsPending
      paymentsApproved
      paymentsCompleted
    }
  }`;
