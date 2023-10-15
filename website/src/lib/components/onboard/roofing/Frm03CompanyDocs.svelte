<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import {getContextClient} from '@urql/svelte';
  import {page} from '$app/stores';
  import {PartnerDocUploadUrl} from '$lib/gql/document';
  import {QryPartnerDocs, SavePartnerCompletedSteps} from '$lib/gql/partner';
  import {goto} from '$app/navigation';
  import alerts from '$lib/stores/alerts';
  import CardWithBorderTitle from '$lib/components/CardWithBorderTitle.svelte';
  import RfxAddress from '$lib/components/RfxAddress.svelte';
  import FileUpload from '$lib/components/FileUpload';
  import type {Document, FileInfo} from '$lib/components/FileUpload/d';
  import Doc from '$lib/components/Doc.svelte';
  import H3WithBorder from '$lib/components/H3WithBorder.svelte';
  import BtnSubmit from '$lib/components/form/BtnSubmit.svelte';
  import FormField from './FormField.svelte';
  import {activeStep, isProfilePage, partner, stepName} from './index';

  const client = getContextClient();
  const stepNo = 3;
  let busy = false,
    docs = [];

  $:partnerID = $page.params.id;
  $:asProfile = $isProfilePage;
  $: if (partnerID) {
    busy = true;
    client.query(QryPartnerDocs, {partnerID, section: 'Docs'}).toPromise()
      .then((res) => {
        docs = res.data?.partnerDocs || [];
      })
      .finally(() => {
        busy = false;
      });
  }

  async function submit() {
    if (busy) return;
    busy = true;
    const res = await client.mutation(SavePartnerCompletedSteps, {partnerID, step: stepNo}).toPromise();
    busy = false;
    if (res.error) {
      alerts.error(`Error: ${$stepName}`, res.error.message);
      return;
    }

    partner.update(p => {
      if (p.setupStepsCompleted < stepNo) {
        p.setupStepsCompleted = stepNo;
      }
      return p;
    });
    // move to next
    $page.url.searchParams.set('step', `${$activeStep + 1}`);
    await goto($page.url.pathname + $page.url.search);
  }

  async function uploadUrl(name: string, f: FileInfo) {
    const res = await client.mutation(
      PartnerDocUploadUrl,
      {partnerID, section: 'Docs', name, fileName: f.name, fileType: f.type, fileSize: f.size}
    ).toPromise();

    if (res.error) {
      return Promise.reject(res.error.message);
    }

    // data =>  {id, key, name, contentType, contentSize, ready, uploadUrl}
    return res.data.partnerDocUploadUrl;
  }

  function uploadComplete(inpName: string, files: Array<Document>) {
    const {uploadUrl, ...data} = files[0];
    const idx = (docs || []).findIndex(o => o.key.endsWith(`/${inpName}`));
    if (idx == -1) {
      docs = (docs || []).concat(data);
    } else {
      docs[idx] = data;
    }
    alerts.success($stepName, `Saved Successfully`);
  }

</script>

<H3WithBorder>{$stepName}</H3WithBorder>

<form on:submit|preventDefault={submit}>
  <CardWithBorderTitle cls='sm:w-1/2' {busy}>
    <FormField label='Completed W-9'>
      <Doc data={(docs || []).find(o => o.ready && o.key.endsWith('/W9'))}/>
      <FileUpload
        name='W9'
        accept='image/*,application/pdf'
        h='h-14'
        {uploadUrl}
        {uploadComplete}
      />
    </FormField>

    <FormField label='General Liability'>
      <Doc data={(docs || []).find(o => o.ready && o.key.endsWith('/GENERAL_LIABILITY'))}/>
      <FileUpload
        name='GENERAL_LIABILITY'
        accept='image/*,application/pdf'
        h='h-14'
        {uploadUrl}
        {uploadComplete}
      />
    </FormField>

    <FormField label='Workers Compensation'>
      <Doc data={(docs || []).find(o => o.ready && o.key.endsWith('/WORKERS_COMPENSATION'))}/>
      <FileUpload
        name='WORKERS_COMPENSATION'
        accept='image/*,application/pdf'
        h='h-14'
        {uploadUrl}
        {uploadComplete}
      />
    </FormField>
    <p class='text-sm text-gray-400 italic'>
      Note: Both general liability and workers compensation must have RFX Home LLC listed as the additionally insured
      using
      the following information:
    </p>
    <RfxAddress/>
  </CardWithBorderTitle>

  {#if !asProfile}
    <BtnSubmit cls='mt-5' loading={busy}>
      Next
    </BtnSubmit>
  {/if}
</form>


