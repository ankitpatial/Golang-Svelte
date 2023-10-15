<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import Input from '$lib/components/form/Input.svelte';
  import BtnSubmit from '$lib/components/form/BtnSubmit.svelte';
  import { getContextClient } from '@urql/svelte';
  import { SavePartner } from '$lib/gql/partner';
  import { goto } from '$app/navigation';
  import { page } from '$app/stores';
  import CardWithBorderTitle from '$lib/components/CardWithBorderTitle.svelte';
  import alerts from '$lib/stores/alerts';
  import InputNumber from '$lib/components/form/InputNumber.svelte';
  import DDLPartnerType from '$lib/components/form/DDLPartnerType.svelte';
  import InputPartnerName from '$lib/components/form/InputPartnerName.svelte';
  import H3WithBorder from '$lib/components/H3WithBorder.svelte';
  import {activeStep, isProfilePage, partner, stepName} from './index';

  const client = getContextClient();

  let data, busy: boolean, readOnly: boolean;
  $: partnerID = $page.params.id;
  $:  {
    data = $partner || {};
  }
  $:asProfile = $isProfilePage;

  async function submit() {
    if (busy) return;

    // save to db
    busy = true;
    const res = await client.mutation(SavePartner, {
      input: {
        id: partnerID,
        type: data.type,
        name: data.name,
        address: data.address,
        latitude: data.latitude,
        longitude: data.longitude,
        website: data.website,
        crewCount: Number(data.crewCount),
        jobCapacity: Number(data.jobCapacity)
      }
    }).toPromise();
    busy = false;

    if (res.error) {
      alerts.error($stepName, res.error.message);
      return;
    }

    if (asProfile) {
      alerts.success($stepName, 'Saved Successfully');
      return;
    }

    // move to next
    $page.url.searchParams.set('step', `${$activeStep + 1}`);

    await goto($page.url.pathname + $page.url.search);

    if ($partner.setupStepsCompleted < 1) {
      partner.update(p => {
        p.setupStepsCompleted = 1;
        return p;
      });
    }
  }
</script>

<H3WithBorder>{$stepName}</H3WithBorder>
<form on:submit|preventDefault={submit}>
  <CardWithBorderTitle cls='sm:w-1/2'>
    <DDLPartnerType bind:value={data.type} readOnly />
    <InputPartnerName label='Company Name' partnerType={data.type} {partnerID} bind:value={data.name} />
    <Input label='Company Address' bind:value={data.address} />
    <Input label='Website' bind:value={data.website} maxlength='150' />
    <div class='flex flex-row justify-between space-x-4 '>
      <InputNumber label='Total Crew' bind:value={data.crewCount} min={0} max={100} />
      <InputNumber label='New Jobs Capacity (weekly)' bind:value={data.jobCapacity} min={0} max={100} />
    </div>
  </CardWithBorderTitle>

  <BtnSubmit cls='mt-5' loading={busy}>
    {asProfile ? 'SAVE' : 'SAVE & NEXT'}
  </BtnSubmit>
</form>


