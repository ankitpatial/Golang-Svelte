<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import Input from '$lib/components/form/Input.svelte';
  import BtnSubmit from '$lib/components/form/BtnSubmit.svelte';
  import {getContextClient} from '@urql/svelte';
  import {SavePartner} from '$lib/gql/partner';
  import {goto} from '$app/navigation';
  import {page} from '$app/stores';
  import CardWithBorderTitle from '$lib/components/CardWithBorderTitle.svelte';
  import alerts from '$lib/stores/alerts';
  import InputPartnerName from '$lib/components/form/InputPartnerName.svelte';
  import H3WithBorder from '$lib/components/H3WithBorder.svelte';
  import {activeStep, isProfilePage, partner, stepName} from './index';
  import {PartnerType} from "$lib/graph/graphql";

  const client = getContextClient();

  let data, busy: boolean, readOnly: boolean;
  $:partnerID = $page.params.id;
  $:{
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
        type: PartnerType.Solar,
        name: data.name,
        address: data.address,
        website: data.website,
        yearsInBusiness: data.yearsInBusiness,
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
    <InputPartnerName label='Solar Company Name' partnerType={data.type} {partnerID} bind:value={data.name}/>
    <Input label='Company HQ Address' bind:value={data.address}/>
    <Input label='Company Website' bind:value={data.website} maxlength='150'/>
    <Input label='Years in Business' type="number" bind:value={data.yearsInBusiness} maxlength='150'/>
  </CardWithBorderTitle>

  <BtnSubmit cls='mt-5' loading={busy}>
    {asProfile ? 'SAVE' : 'SAVE & NEXT'}
  </BtnSubmit>
</form>


