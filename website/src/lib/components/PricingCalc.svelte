<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->
<script lang='ts'>
  import {getContextClient} from '@urql/svelte';
  import type Job from "$lib/models/Job";
  import Modal from "$lib/components/Modal.svelte";
  import {DataColumn, Table} from "$lib/components/DataTable";
  import BtnIcon from "$lib/components/form/BtnIcon.svelte";
  import IconCancel from "$lib/assets/svg/IconCancel.svelte";
  import {TestPricing} from "$lib/gql/estimate";

  export let job: Job;
  const client = getContextClient();

  let showModal = false, showPricing, sq, pitch, err, pricing, amt, summary;
  let measurements: Array<{ squares: number, pitch: string }> = [{squares: 18, pitch: '7/12'}];

  // Test Pricing modal
  function openModal() {
    measurements = [{squares: 18, pitch: '7/12'}];
    summary = '';
    showModal = true;
    showPricing = false;
    sq = undefined;
    pitch = undefined;
    err = undefined;
  }

  // Test Pricing calculation
  async function getPricing() {
    err = undefined;
    if (!measurements.length === 0) {
      err = 'add a measurements';
      return;
    }

    showPricing = false;
    pricing = true;
    const res = await client.mutation(TestPricing, {
      job: {
        ownerFirstName: job.ownerFirstName,
        ownerLastName: job.ownerLastName,
        streetNumber: job.streetNumber,
        streetName: job.streetName,
        city: job.city,
        state: job.state,
        zip: job.zip,
        latitude: job.latitude,
        longitude: job.longitude,
        repFirstName: job.repFirstName,
        repLastName: job.repLastName,
        repEmail: job.repEmail,
        repMobile: job.repMobile,
        companyName: job.companyName,
        currentMaterial: job.currentMaterial,
        newRoofingMaterial: job.newRoofingMaterial,
        lowSlope: job.lowSlope,
        currentMaterialLowSlope: job.currentMaterialLowSlope,
        newRoofingMaterialLowSlope: job.newRoofingMaterialLowSlope,
        redeck: job.redeck,
        layers: job.layers,
        layer2Material: job.layer2Material,
        layer3Material: job.layer3Material,
        webhookURL: job.webhookURL,
        measurementType: job.measurementType,
      }, measure: measurements
    }).toPromise();
    pricing = false;
    if (res.error) {
      err = res.error.message;
      return;
    }

    amt = res.data.testJobPricing.total;
    summary = res.data.testJobPricing.summary;
  }
</script>

<button class='btn btn-outline w-28 mt-10' on:click={openModal}>
  Price Calculator
</button>

{#if (showModal)}
  <Modal open size="lg" title='Test Pricing' okText="Calculate" on:ok={getPricing} on:cancel={()=> {showModal = false}}>
    <Table
      headers={["SQ", "PITCH", ""]}
      items={measurements}
    >
      <svelte:fragment slot="row" let:idx>
        <DataColumn header="SQ">
          <input type="number" class="input input-bordered input-sm" bind:value={measurements[idx].squares}/>
        </DataColumn>
        <DataColumn header="PITCH">
          <select class="select select-bordered select-sm" bind:value={measurements[idx].pitch}>
            <option value="1/12">1/12</option>
            <option value="2/12">2/12</option>
            <option value="3/12">3/12</option>
            <option value="4/12">4/12</option>
            <option value="5/12">5/12</option>
            <option value="6/12">6/12</option>
            <option value="7/12">7/12</option>
            <option value="8/12">8/12</option>
            <option value="9/12">9/12</option>
            <option value="10/12">10/12</option>
            <option value="11/12">11/12</option>
            <option value="12/12">12/12</option>
            <option value="13/12">13/12 Or Grater</option>
          </select>
        </DataColumn>
        <DataColumn>
          {#if idx > 0}
            <BtnIcon on:click={() => {
              measurements = measurements.filter((_, i) => i !== idx)
            }}>
              <IconCancel/>
            </BtnIcon>
          {/if}
        </DataColumn>
      </svelte:fragment>
    </Table>
    <button class="btn btn-outline mt-5" on:click={()=> {
      measurements = measurements.concat({squares: 20, pitch: '10/12'})
    }}>
      ADD
    </button>

    {#if (err)}
      <div class='alert alert-error'>
        {err}
      </div>
    {/if}

    <div class="whitespace-pre mt-5">
      {summary || ''}
    </div>
  </Modal>
{/if}
