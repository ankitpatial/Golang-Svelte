<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import {getContextClient} from '@urql/svelte';
  import {goto} from '$app/navigation';
  import {page} from '$app/stores';
  import alerts from '$lib/stores/alerts';
  import CardWithBorderTitle from '$lib/components/CardWithBorderTitle.svelte';
  import Collapse from '$lib/components/ui/Collapse.svelte';
  import InputToggle from '$lib/components/form/InputToggle.svelte';
  import {QryPartnerServiceStates, SavePartnerCompletedSteps, SaveServiceCity} from '$lib/gql/partner';
  import BtnSubmit from '$lib/components/form/BtnSubmit.svelte';
  import H3WithBorder from '$lib/components/H3WithBorder.svelte';
  import type {ServiceState} from './index';
  import {activeStep, isProfilePage, partner, stepName} from './index';

  const client = getContextClient();
  const stepNo = 3;

  let data: Array<ServiceState> = [], busy = false;
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
      })
      .finally(() => {
        busy = false;
      });
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
      <table class='table table-compact table-fixed w-full'>
        <thead>
        <tr>
          <th>City</th>
          <th>Select at least One</th>
        </tr>
        </thead>
        <tbody>
        {#each state.cities || [] as sa (sa.id)}
          <tr>
            <td class='w-32'>{sa.cityName} - {sa.cityZip}</td>
            <td>
              <InputToggle bind:value={sa.active} on:change={({detail}) => saveCity(sa.id, 'active', detail)}/>
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
