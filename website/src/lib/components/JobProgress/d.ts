import type JobProgressHistory from '$lib/models/JobProgressHistory';
import type {Job, JobProgress} from "$lib/graph/graphql";

export type QType = 'buttons' | 'radio' | 'dtp' | 'note' | 'fileUpload'


export interface Progress {
  job: Job;
  activeIdx: number;
  loading: boolean;
  saving: boolean;
  history: Array<JobProgressHistory>;
  delayed: boolean;
  isDirty: boolean;
  readonly: boolean;
}

export interface Question {
  id: number;
  type: QType;
  text: string;
  description?: string;
  options?: Array<string>;
  on?: object;
  end?: boolean;
  inActive?: boolean;
  upload?: {
    name?: string
    accept: string
    multiple: boolean
    section: string
  };
}

export interface Answer {
  qID: number;
  qType: QType;
  question: string;
  answer: string;
  hasOtherTxt?: boolean;
  otherTxt?: string;
  order?: number;
  data?: object;
}

export interface JobProgressStep {
  active: boolean;
  completed: boolean;
  completedAt?: string;
  component: any;
  name: JobProgress;
}
