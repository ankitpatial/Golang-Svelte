import type {JobProgress} from "$lib/graph/graphql";

export default interface JobProgressHistory {
  id: string,
  status: JobProgress,
  statusAt: Date
  complete: boolean,
  response?: [{
    question: string,
    answer: string,
    order: number,
  }]
}
