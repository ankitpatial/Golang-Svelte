import {derived, get, writable} from 'svelte/store';
import {humanize} from '$lib/utils/string';
import type {Client} from '@urql/svelte/dist/urql-svelte';
import type {Job} from "$lib/graph/graphql";
import {localDate} from '$lib/utils/time';
import type {Progress} from '$lib/components/JobProgress/d';
import {JobProgressUpdate, QryJobCompletedProgress} from '$lib/gql/job.progress';
import alerts from '$lib/stores/alerts';
import type JobProgressHistory from '$lib/models/JobProgressHistory';
import {JobProgress} from "$lib/graph/graphql";
import Step01New from './Step01New.svelte';
import Step02ContactCustomer from './Step02ContactCustomer.svelte';
import Step03JobDetailConfirmed from './Step03JobDetailConfirmed.svelte';
import Step04Permitting from './Step04Permitting.svelte';
import Step05Scheduled from './Step05Scheduled.svelte';
import Step06InProgress from './Step06InProgress.svelte';
import Step07Installed from './Step07Installed.svelte';
import Step08Invoiced from './Step08Invoiced.svelte';

const stepNames = [
  JobProgress.New,
  JobProgress.CustomerContacted,
  JobProgress.JobDetailsConfirmed,
  JobProgress.Permitting,
  JobProgress.Scheduled,
  JobProgress.InProgress,
  JobProgress.Installed,
  JobProgress.Invoiced
];

const components = [
  Step01New,
  Step02ContactCustomer,
  Step03JobDetailConfirmed,
  Step04Permitting,
  Step05Scheduled,
  Step06InProgress,
  Step07Installed,
  Step08Invoiced
];


function createStore() {
  const {subscribe, set, update} = writable<Progress>();

  let jobID: string;
  let client: Client;

  async function pullHistory() {
    update(p => {
      p.loading = true;
      return p;
    });

    // pull job progress history
    const res = await client.query(QryJobCompletedProgress, {id: jobID}).toPromise();
    if (res.error) {
      alerts.error('Data Load Error', res.error.message);
      update(p => {
        p.loading = false;
        return p;
      });
      return;
    }

    update(p => {
      p.loading = false;
      p.history = res.data.jobCompletedProgress || [];
      p.delayed = p.job.progress === JobProgress.Delayed;

      const lp = p.history.findLast(p => p.status !== JobProgress.Delayed);
      if (lp) {
        p.activeIdx = stepNames.indexOf(lp.status) + 1;
      } else {
        p.activeIdx = 1;
      }
      return p;
    });
  }

  return {
    subscribe,
    //
    // init state with gql context & job data
    init: async (cl: Client, job: Job, readonly: boolean) => {
      jobID = job.id || '';
      client = cl;

      // set state
      set({
        job: job,
        loading: true,
        saving: false,
        activeIdx: 0,
        history: [],
        delayed: false,
        isDirty: false,
        readonly
      });

      await pullHistory();
    },

    pullHistory,

    setJobDelay: (flag: boolean) => {
      if (flag) {
        update(p => {
          p.delayed = true;
          p.isDirty = true;
          if (p.job) {
            p.job.progress = JobProgress.Delayed;
          }
          return p;
        });
      } else {
        update(p => {
          p.delayed = false;
          p.isDirty = true;
          if (p.job) {
            const lp = p.history.findLast(p => p.status !== JobProgress.Delayed);
            p.job.progress = lp?.status || JobProgress.New;
          }
          return p;
        });
      }
    },

    //
    // save progress
    save: async (note: string, data: object | undefined = undefined): Promise<boolean> => {
      // saving = true
      update((s: Progress) => {
        s.saving = true;
        s.isDirty = true;
        return s;
      });
      // save
      const res = await saveProgress(client, false, note, data);
      // saving = false
      update((s: Progress) => {
        s.saving = false;
        return s;
      });
      return res;
    },

    //
    // move state to next step
    saveAndNext: async (note: string, data: object | undefined = undefined): Promise<boolean> => {
      // saving = true
      update((s: Progress) => {
        s.saving = true;
        s.isDirty = true;
        return s;
      });
      // save
      const res = await saveProgress(client, true, note, data);
      if (res) {
        update((s: Progress) => {
          const idx = s.activeIdx;
          const step = stepNames[idx];

          s.saving = false;
          s.isDirty = true;
          s.activeIdx = idx + 1;

          // change job status
          s.job.progress = step;
          s.job.progressAt = new Date();
          s.job.progressFlagged = false;

          // add entry to progressHistory
          s.history = (s.history || []).concat({
            id: '',
            status: step,
            statusAt: new Date(),
            complete: true
          });

          if (data) {
            Object.assign(s.job, data);
          }
          return s;
        });
      } else {
        // saving = false
        update((s: Progress) => {
          s.saving = false;
          return s;
        });
      }

      return res;
    }
  };
}

export const progress = createStore();

export const steps = derived(progress, (p) => {
  if (!p?.job) {
    return [];
  }
  const activeIdx = p.activeIdx;
  const complete: Array<JobProgressHistory> = [{
    id: '',
    status: JobProgress.New,
    statusAt: p.job.progressAt || new Date(),
    complete: true
  }].concat(p.history);


  return stepNames.map((s, idx) => {
    const h = complete.find(o => o.status === s);
    return {
      name: humanize(s).toUpperCase(),
      active: idx === activeIdx,
      completed: !!h,
      completedAt: localDate(h?.statusAt),
      component: components[idx]
    };
  });
});

export const saving = writable<boolean>(false);

async function saveProgress(
  client: Client, done: boolean, note: string, data: object | undefined = undefined
): Promise<boolean> {
  try {
    const p = get(progress);
    const idx = p.activeIdx;
    const step = stepNames[idx];

    saving.set(true);
    const r = await client.mutation(
      JobProgressUpdate,
      {
        id: p.job.id,
        step,
        stepComplete: done,
        note,
        data
      }
    ).toPromise();

    if (r.error) {
      alerts.error('Save Job Progress', r.error.message);
      return false;
    }

    return true;
  } catch (ex) {
    console.error(ex);
    alerts.error('Save Job Progress', 'unexpected error');
    return false;
  } finally {
    saving.set(false);
  }
}

