<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import InputReadonly from '$lib/components/form/InputReadonly.svelte';
  import { usdFormat } from '$lib/utils/currency.js';
  import BtnBackSmall from '$lib/components/form/BtnBackSmall.svelte';
  import { goto } from '$app/navigation';
  import { PATH } from '$lib/enum.js';
  import { getContextClient } from '@urql/svelte';
  import { AssignPartnerToJob, QryJob } from '$lib/gql/job';
  import { page } from '$app/stores';
  import CardWithBorderTitle from '$lib/components/CardWithBorderTitle.svelte';
  import SearchPartner from '$lib/components/search/SearchPartner.svelte';
  import alerts from '$lib/stores/alerts';
  import Page from '$lib/components/Page.svelte';
  import type Job from '$lib/models/Job';

  const client = getContextClient();
  let open = false, loading = true,
    job: Job,
    data: Pick<Job, 'id' | 'partner' | 'materialDate' | 'removeDate' | 'installDate' | 'completionDate'> = {
      id: '',
      partner: undefined,
      materialDate: undefined,
      removeDate: undefined,
      installDate: undefined,
      completionDate: undefined
    },
    saving = false, tmp;

  client.query(QryJob, { id: $page.params.id }, { requestPolicy: 'cache-first' }).toPromise()
    .then((res) => {
      if (res.error) {
        alerts.push({ type: 'error', title: 'Load Data Error', body: res.error.message });
        return;
      }

      job = res.data.job;
      data.materialDate = job?.materialDate;
      data.removeDate = job?.removeDate;
      data.installDate = job?.installDate;
      data.completionDate = job?.completionDate;
    })
    .finally(() => {
      loading = false;
    });

  function toggle() {
    open = !open;
  }

  async function save() {
    if (saving) {
      return;
    }

    let title = 'Assign Partner';
    let partnerID = data.partner?.id;
    if (!partnerID) {
      alerts.push({ type: 'warning', title, body: 'Please select a partner.' });
      return;
    }

    const res = await client.mutation(AssignPartnerToJob, {
      jobID: job.id,
      partnerID: partnerID
    }).toPromise();

    if (res.error) {
      alerts.push({ type: 'error', title: `Error: ${title}`, body: res.error.message });
      return;
    }

    alerts.push({ type: 'success', title, body: 'Assigned Successfully' });
    await back();
  }

  async function back() {
    return goto(PATH.JOBS_UNASSIGNED + $page.url.search);
  }
</script>


<Page title='Assign Partner To Job'>
  <svelte:fragment slot='buttons'>
    <BtnBackSmall on:click={back} tooltip='Back to Unassigned Jobs' />
  </svelte:fragment>

  <div class='flex gap-4 flex-col sm:flex-row flex-wrap' class:blur-sm={loading}>
    <CardWithBorderTitle title='Job Detail'>
      <InputReadonly label='Job Owner' value={`${job?.ownerFirstName} ${job?.ownerLastName}`} />
      <InputReadonly
        label='Job Address'
        value={`${job?.streetNumber} ${job?.streetName}, ${job?.city}, ${job?.state} ${job?.zip}`}
      />
      <InputReadonly label='Job Price' value={job?.price ? usdFormat(job?.price) : 'N/A'} />
    </CardWithBorderTitle>

    <CardWithBorderTitle title='Partner'>
      <SearchPartner partnerType='ROOFING' bind:value={data.partner} />
    </CardWithBorderTitle>
  </div>

  <div class='mt-5'>
    <button class='btn btn-primary' class:loading={saving} on:click={save}>
      ASSIGN
    </button>
  </div>
</Page>
