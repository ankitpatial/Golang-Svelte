<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import {getContextClient} from '@urql/svelte';
  import {goto} from '$app/navigation';
  import {page} from '$app/stores';
  import alerts from '$lib/stores/alerts';
  import {QryPartnerServices, SaveLeadTime, SavePartnerCompletedSteps, SaveService} from '$lib/gql/partner';
  import CardWithBorderTitle from '$lib/components/CardWithBorderTitle.svelte';
  import InputToggle from '$lib/components/form/InputToggle.svelte';
  import InputTime from '$lib/components/form/InputTime.svelte';
  import BtnSubmit from '$lib/components/form/BtnSubmit.svelte';
  import H3WithBorder from '$lib/components/H3WithBorder.svelte';
  import {authUser} from '$lib/stores/auth';
  import {isProfilePage, partner, stepName} from './index';

  const client = getContextClient();
  const stepNo = 5;

  let services, busy = false;
  $:partnerID = $page.params.id;
  $:asProfile = $isProfilePage;
  $: if (partnerID) {
    client.query(QryPartnerServices, {partnerID}).toPromise()
      .then((res) => {
        services = res.data.partnerServices;
      });
  }

  async function save(item: object, active: boolean) {
    const res = await client.mutation(SaveService, {
      id: item.id,
      partnerID,
      service: item.service,
      active
    }).toPromise();

    if (res.error) {
      alerts.error(item.description, 'failed with error: ' + res.error.message);
      const idx = services.findIndex(s => s.id === item.id);
      services[idx].active = !active;
      return;
    }

    alerts.success($stepName, 'Successfully Saved Trades/Work Capabilities');
  }

  async function saveTime(field: string, val: string) {
    const inp = {partnerID};
    inp[field] = val;
    const res = await client.mutation(SaveLeadTime, inp).toPromise();
    if (res.error) {
      alerts.error(`Error: ${$stepName}`, res.error.message);
      return;
    }

    alerts.success($stepName, 'Successfully Saved Lead Times');
  }

  async function done() {
    if (busy) return;
    busy = true;
    const res = await client.mutation(SavePartnerCompletedSteps, {partnerID, step: stepNo, done: true}).toPromise();
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

    if ($authUser?.partner && !asProfile) {
      await goto(`/${$authUser?.partner.type.toLowerCase()}/onboarding/done`);
    }
  }
</script>

<H3WithBorder>{$stepName}</H3WithBorder>
<CardWithBorderTitle cls='xs:w-full sm:w-8/12'>
  <table class='table table-compact border-b' class:blur-sm={busy}>
    <thead>
    <tr>
      <th></th>
      <th>Receive Bid Requests</th>
    </tr>
    </thead>
    <tbody>
    {#each services || [] as s}
      <tr>
        <td>
          {s.description}
        </td>
        <td>
          <InputToggle
            bind:value={s.active}
            on:change={({detail}) => save(s, detail)}
          />
        </td>
      </tr>
    {/each}
    </tbody>
  </table>
  <p class='mt-5'>
    Please provide the following lead times to completely complete a project with the listed materials. <br>
    When providing this lead time assume that material ordering is not an issue:
  </p>
  <div class='flex flex-row flex-wrap space-x-5'>
    <InputTime
      label='Asphalt Lead Time'
      w='w-52'
      bind:value={$partner.asphaltLeadT}
      saveOnChange={(val) => saveTime('asphalt', val)}
    />
    <InputTime
      label='Metal Lead Time'
      w='w-52'
      bind:value={$partner.metalLeadT}
      saveOnChange={(val) => saveTime('metal', val)}
    />
    <InputTime
      label='Tile Lead Time'
      w='w-52'
      bind:value={$partner.tileLeadT}
      saveOnChange={(val) => saveTime('tile', val)}
    />
  </div>
</CardWithBorderTitle>

<BtnSubmit cls='mt-5' loading={busy} on:click={done}>
  {asProfile ? 'SAVE' : 'DONE'}
</BtnSubmit>
