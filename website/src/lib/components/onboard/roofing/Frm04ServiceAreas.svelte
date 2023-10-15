<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import {getContextClient} from '@urql/svelte';
  import {goto} from '$app/navigation';
  import {page} from '$app/stores';
  import {QryServiceStates} from '$lib/gql/postalCode';
  import {PartnerDocUploadUrl} from '$lib/gql/document';
  import alerts from '$lib/stores/alerts';
  import CardWithBorderTitle from '$lib/components/CardWithBorderTitle.svelte';
  import SearchBox from '$lib/components/form/SearchBox.svelte';
  import Collapse from '$lib/components/ui/Collapse.svelte';
  import InputToggle from '$lib/components/form/InputToggle.svelte';
  import Input from '$lib/components/form/Input.svelte';
  import InputDate from '$lib/components/form/dtp/InputDate.svelte';
  import FileUpload from '$lib/components/FileUpload/FileUpload.svelte';
  import {
    QryPartnerServiceStates,
    SavePartnerCompletedSteps,
    SaveServiceCity,
    SaveServiceState
  } from '$lib/gql/partner';
  import type {Document, FileInfo} from '$lib/components/FileUpload/d';
  import Doc from '$lib/components/Doc.svelte';
  import BtnSubmit from '$lib/components/form/BtnSubmit.svelte';
  import H3WithBorder from '$lib/components/H3WithBorder.svelte';
  import type {ServiceState} from './index';
  import {activeStep, isProfilePage, partner, stepName} from './index';
  import Label from "$lib/components/form/Label.svelte";

  const client = getContextClient();
  const stepNo = 4;

  let data: Array<ServiceState> = [], docs = [], busy = false;
  $:partnerID = $page.params.id;
  $:asProfile = $isProfilePage;
  $: if (partnerID) {
    busy = true;
    client.query(QryPartnerServiceStates, {partnerID}).toPromise()
      .then((res) => {
        if (res.error) {
          console.error(res.error);
          return;
        }

        data = res.data.partnerServiceStates;
        docs = res.data.partnerDocs;
      })
      .finally(() => {
        busy = false;
      });
  }

  async function uploadUrl(inpName: string, f: FileInfo) {
    const res = await client.mutation(PartnerDocUploadUrl,
      {partnerID, section: 'Proof', name: inpName, fileName: f.name, fileType: f.type, fileSize: f.size}
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
    if (idx === -1) {
      docs = (docs || []).concat(data);
    } else {
      docs[idx] = data;
    }
  }

  function stateUploadComplete(stateName: string, inpName: string, files: Array<Document>) {
    uploadComplete(inpName, files);
    saveState(stateName, 'proofDocID', files[0].id);
  }

  async function saveState(stateName, field, value) {
    const inp = {partnerID, state: stateName};
    inp[field] = value;
    return client.mutation(SaveServiceState, inp).toPromise()
      .then((res) => {
        if (res.error) {
          alerts.error('Error: Save State', `failed to save '${field}' with error: ${res.error.message}`);
          return;
        }

        alerts.success('Save State', 'Successfully saved');
      });
  }

  function cityUploadComplete(postalID: string, inpName: string, files: Array<Document>) {
    uploadComplete(inpName, files);
    saveCity(postalID, 'proofDocID', files[0].id);
  }

  async function saveCity(postalID, field, value) {
    const inp = {partnerID, postalID};
    inp[field] = value;

    return client.mutation(SaveServiceCity, inp).toPromise()
      .then((res) => {
        if (res.error) {
          alerts.error('Error: Save City', `failed to save '${field}' with error: ${res.error.message}`);
          return;
        }

        alerts.success('Save City', 'Successfully saved');
      });
  }

  async function next() {
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
</script>

<H3WithBorder>{$stepName}</H3WithBorder>

<CardWithBorderTitle cls='xs:w-full sm:w-8/12'>
  {#each data as state (state.name)}
    <Collapse title={state.name} bind:open={state.expand}>
      <div class='flex flex-row flex-wrap gap-x-2'>
        <div class='w-52'>
          <Input
            label='License No'
            required
            maxlength='50'
            bind:value={state.licenseNo}
            saveOnChange={(v) => saveState(state.name, 'licNo', v)}
          />
        </div>
        <div class='w-52'>
          <InputDate
            label='License Expiry Date'
            required
            w='w-52'
            bind:value={state.licenseExpDate}
            saveOnChange={(v) => saveState(state.name, 'expDate', v)}
          />
        </div>
        <div class='w-52'>
          <Label content="PROOF"/>
          <div class="mt-1">
            <FileUpload
              w='w-52'
              mt="-mt-1"
              name='{state.name}-lic'
              accept='image/*,application/pdf'
              {uploadUrl}
              uploadComplete={(inpName, files) => stateUploadComplete(state.name, inpName, files)}
            />
          </div>
        </div>
        <div class=' mt-4'>
          <Doc data={(docs || []).find(o => o.ready && o.key.endsWith(`/${state.name}-lic`))}/>
        </div>
      </div>
      <table class='table table-compact table-fixed w-full'>
        <thead>
        <tr>
          <th>City</th>
          <th>Select at least One</th>
          <th>Registration Number(if applicable)</th>
          <th>Registration Proof(if applicable)</th>
          <th></th>
        </tr>
        </thead>
        <tbody>
        {#each state.cities || [] as sa (sa.id)}
          {@const city = `${sa.cityName}-${sa.cityZip}`}
          <tr>
            <td class='w-32'>{sa.cityName} - {sa.cityZip}</td>
            <td>
              <InputToggle
                bind:value={sa.active}
                on:change={({detail}) => saveCity(sa.id, 'active', detail)}
              />
            </td>
            <td>
              <Input
                bind:value={sa.licenseNo}
                disabled={!sa.active}
                saveOnChange={(v) => saveCity(sa.id, 'licNo', v)}
              />
            </td>
            <td>
              <FileUpload
                w='w-52'
                mt="-mt-3"
                name='{city}-lic'
                accept='image/*,application/pdf'
                {uploadUrl}
                uploadComplete={(inpName, files) => cityUploadComplete(sa.id, inpName, files)}
                disabled={!sa.active}
              />
            </td>
            <td>
              <Doc data={(docs || []).find(o => o.ready && o.key.endsWith(`/${city}-lic`))}/>
            </td>
          </tr>
        {/each}
        </tbody>
      </table>
    </Collapse>
  {/each}
</CardWithBorderTitle>

{#if (!asProfile)}
  <BtnSubmit cls='mt-5' loading={busy} on:click={next}>
    Next
  </BtnSubmit>
{/if}
