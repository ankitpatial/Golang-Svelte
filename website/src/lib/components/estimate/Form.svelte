<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->
<script lang='ts'>
  import CardWithBorderTitle from '$lib/components/CardWithBorderTitle.svelte';
  import Input from '$lib/components/form/Input.svelte';
  import InputToggle from '$lib/components/form/InputToggle.svelte';
  import PlacesModal from '$lib/components/maps/PlacesModal.svelte';
  import alerts from '$lib/stores/alerts';
  import {
    layerMaterial,
    lowSlopeMaterial,
    lowSlopeProduct,
    measurementTypes,
    partialRoof,
    steepSlopeMaterial,
    steepSlopeProduct
  } from './d';
  import DropDown from '$lib/components/form/DropDown.svelte';
  import BtnSubmit from '$lib/components/form/BtnSubmit.svelte';
  import InputPhone from '$lib/components/form/InputPhone.svelte';
  import {authUser} from '$lib/stores/auth';
  import Modal from '$lib/components/Modal.svelte';
  import SearchPartnerUsers from "$lib/components/search/SearchPartnerUsers.svelte";
  import {type CreateEstimateInput, Role} from "$lib/graph/graphql";

  export let save: (data: CreateEstimateInput) => Promise<void>;
  export let estimate: CreateEstimateInput;
  export let hideRep = false;

  const currentMaterial = [''].concat(steepSlopeMaterial).map((item) => ({
    id: item,
    name: item
  }));
  const prodSteepSlope = [''].concat(steepSlopeProduct).map((item) => ({id: item, name: item}));
  // const materialLowSlope = [''].concat(lowSlopeMaterial).map((item) => ({id: item, name: item}));
  const layerMaterials = [''].concat(layerMaterial).map((item) => ({id: item, name: item}));
  const prodLowSlope = [''].concat(lowSlopeProduct).map((item) => ({id: item, name: item}));
  const layerOpt = [1, 2, 3].map((item) => ({id: item, name: item}));

  let busy, coNameEl, showModal = false;

  $: isCurrentLowSlopeOnly = estimate?.currentMaterial === 'Low Slope Only';
  $: isCurrentWoodshake = estimate?.currentMaterial === 'Wood Shakes';

  function setAddress({detail}) {
    if (!detail) return;

    estimate.streetNumber = detail.streetNumber;
    estimate.streetName = detail.street;
    estimate.city = detail.city;
    estimate.state = detail.state;
    estimate.zip = detail.zip;
    estimate.latitude = detail.latitude;
    estimate.longitude = detail.longitude;

    if (coNameEl?.focus) {
      setTimeout(() => coNameEl.focus(), 100);
    }
  }

  async function submit() {
    if (busy) return;

    try {
      busy = true;
      await save(estimate).then(() => busy = false);
    } catch (ex) {
      alerts.push({type: 'error', title: 'Error', body: ex.message || ex});
    } finally {
      busy = false;
    }
  }

  function layerChanged() {
    if (estimate.layers === 2) {
      estimate.layer3Material = '';
    } else if (estimate.layers === 1) {
      estimate.layer2Material = '';
      estimate.layer3Material = '';
    }
  }

  function currentMaterialChanged() {
    if (isCurrentLowSlopeOnly) {
      estimate.newRoofingMaterial = '';
      estimate.newRoofingMaterialLowSlope = '';
      estimate.layers = 1;
      layerChanged();
    } else if (isCurrentWoodshake) {
      estimate.redeck = true;
    }
  }

  function showAssignRepModal() {
    showModal = true;
  }

  function assignRep({detail}) {
    showModal = false;
    estimate.repFirstName = detail.firstName;
    estimate.repLastName = detail.lastName;
    estimate.repEmail = detail.email;
    estimate.repMobile = detail.phone || '';
  }

</script>

<form class='flex gap-4 flex-col sm:flex-row flex-wrap' on:submit|preventDefault={submit}>
  <CardWithBorderTitle title='Home Owner'>
    <Input label='First Name' bind:value={estimate.ownerFirstName} required/>
    <Input label='Last Name' bind:value={estimate.ownerLastName} required/>
    <Input label='Street Number' bind:value={estimate.streetNumber} required>
      <PlacesModal label='find address on map' on:ok={setAddress} slot='alt'/>
    </Input>
    <Input label='Street Name' bind:value={estimate.streetName} required/>
    <Input label='City' bind:value={estimate.city} required/>
    <Input label='State' bind:value={estimate.state} required/>
    <Input label='Zip' bind:value={estimate.zip} required/>
  </CardWithBorderTitle>

  <CardWithBorderTitle title='Sales Rep'>
    <svelte:fragment slot='btn'>
      {#if ($authUser.role === Role.SiteUser)}
        <div type='button' class='btn btn-outline btn-xs uppercase' on:click|preventDefault={showAssignRepModal}>
          Assign
        </div>
      {/if}
    </svelte:fragment>
    <Input label='Company Name' bind:value={estimate.companyName} required bind:ref={coNameEl} disabled/>
    <Input label='First Name' bind:value={estimate.repFirstName} required disabled/>
    <Input label='Last Name' bind:value={estimate.repLastName} required disabled/>
    <Input label='Email' bind:value={estimate.repEmail} type='email' required disabled/>
    <InputPhone label='Mobile' bind:value={estimate.repMobile} disabled/>
  </CardWithBorderTitle>
  <CardWithBorderTitle title='Job Details'>
    <DropDown
      label='What type of material is currently on the roof?'
      tooltipForLabel='What type of material is currently on the roof?'
      options={currentMaterial}
      on:change={currentMaterialChanged}
      bind:value={estimate.currentMaterial}
      required
    />

    {#if isCurrentLowSlopeOnly}
      <DropDown
        label='What type of material will go on the new roof?'
        tooltipForLabel='What type of material will go on the new roof?'
        options={prodLowSlope}
        bind:value={estimate.newRoofingMaterialLowSlope}
        required
      />
    {:else }
      <DropDown
        label='What type of material will go on the new roof?'
        tooltipForLabel='What type of material will go on the new roof?'
        options={prodSteepSlope}
        bind:value={estimate.newRoofingMaterial}
        disabled={isCurrentLowSlopeOnly}
        required={!isCurrentLowSlopeOnly}
      />
    {/if}

    <InputToggle
      label='Will new decking or sheathing be required?'
      wrapperCls="w-full mb-3"
      tooltipForLabel='Will new decking or sheathing be required?'
      bind:value={estimate.redeck}
      disabled={isCurrentWoodshake}
      tooltip={isCurrentWoodshake ? 'Wood Shakes requires redecking' : ''}
    />
    <DropDown
      label='How many layers of roofing material are there?'
      tooltipForLabel='How many layers of roofing material are there?'
      options={layerOpt}
      bind:value={estimate.layers}
      on:change={layerChanged}
      disabled={isCurrentLowSlopeOnly}
      tooltip={isCurrentLowSlopeOnly ? 'Current Material is "Low Slope Only"' : ''}
    />
    {#if (estimate.layers > 1)}
      <DropDown
        label='Layer 2 Material'
        options={layerMaterials}
        bind:value={estimate.layer2Material}
        required
      />
    {/if}
    {#if (estimate.layers > 2)}
      <DropDown
        label='Layer 3 Material'
        options={layerMaterials}
        bind:value={estimate.layer3Material}
        required
      />
    {/if}

    <DropDown
      label='Is this estimate for a partial roof replacement?'
      tooltipForLabel='Is this estimate for a partial roof replacement?'
      options={partialRoof}
      bind:value={estimate.partial}
    />

    <DropDown
      label='Measurement Type'
      options={measurementTypes}
      bind:value={estimate.measurementType}
    />

  </CardWithBorderTitle>

  <p class="w-full"></p>
  <BtnSubmit loading={busy}>
    Submit
  </BtnSubmit>
</form>

<Modal open={showModal} hideOkBtn title='Sales Rep Search' on:cancel={()=> {showModal = false;}} size='sm'>
  <SearchPartnerUsers on:change={assignRep}/>
</Modal>
