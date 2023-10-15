<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->
<script lang='ts'>
  import { page } from '$app/stores';
  import InputReadonly from '$lib/components/form/InputReadonly.svelte';
  import { getContextClient } from '@urql/svelte';
  import { QryEstimate } from '$lib/gql/estimate';
  import alerts from '$lib/stores/alerts';
  import CardWithBorderTitle from '$lib/components/CardWithBorderTitle.svelte';
  import { createEventDispatcher } from 'svelte';
  import { usdFormat } from '$lib/utils/currency.js';
  import Polygon from '$lib/components/maps/Polygon.svelte';
  import { type Estimate } from '$lib/graph/graphql';
  import H3WithBorder from '$lib/components/H3WithBorder.svelte';
  import Doc from '$lib/components/Doc.svelte';
  import PriceSummary from '$lib/components/estimate/PriceSummary.svelte';

  export let id: string;

  // pull job info along with measurement orders
  const client = getContextClient();
  const dispatch = createEventDispatcher();

  let estimate: Estimate;
  let busy = false;
  let boundary;
  let viewSummary = false;
  let priceSummary = '';

  busy = true;
  client.query(QryEstimate, { id }).toPromise().then((res) => {
    busy = false;
    // check for error and alert
    if (res.error) {
      alerts.error('Job Info', res.error.message);
      return;
    }
    estimate = res.data.estimate;
    dispatch<Estimate>('load', estimate);
  });

  $: pathname = $page.url.pathname;
  $: isCurrentLowSlopeOnly = estimate?.currentMaterial === 'Low Slope Only';

  // onMount(() => {
  //   return wsMessage.subscribe(msg => {
  //     if (!msg) return;
  //     if (Channel.Estimate === msg.channel && Topic.StatusChange === msg.topic && msg.data.id === id) {
  //     }
  //   })
  // })

</script>

<div class='flex gap-4 flex-col sm:flex-row flex-wrap' class:blur-sm={busy}>
  <CardWithBorderTitle title='Home Owner'>
    <InputReadonly label='First Name' value={estimate?.homeOwner?.firstName} />
    <InputReadonly label='Last Name' value={estimate?.homeOwner?.lastName} />
    <InputReadonly label='Street Number' value={estimate?.homeOwner?.streetNumber} />
    <InputReadonly label='Street Name' value={estimate?.homeOwner?.streetName} />
    <InputReadonly label='City' value={estimate?.homeOwner?.city} />
    <InputReadonly label='State' value={estimate?.homeOwner?.state} />
    <InputReadonly label='Zip' value={estimate?.homeOwner?.zip} />

  </CardWithBorderTitle>

  <CardWithBorderTitle title='Sales Rep'>
    <InputReadonly label='Company Name' value={estimate?.companyName || 'N/A'} />
    <InputReadonly label='First Name' value={estimate?.salesRep?.firstName} />
    <InputReadonly label='Last Name' value={estimate?.salesRep?.lastName} />
    <InputReadonly label='Email' value={estimate?.salesRep?.email} />
    <InputReadonly label='Mobile' value={estimate?.salesRep?.phone} />
  </CardWithBorderTitle>

  <CardWithBorderTitle title='Job Details'>
    <InputReadonly label='Current Material - Existing Steep Slope Material'
                   value={estimate?.currentMaterial || 'N/A'} />
    {#if isCurrentLowSlopeOnly}
      <InputReadonly
        label='New Low Slope Material - Low Slope Products'
        value={estimate?.newRoofingMaterialLowSlope || 'N/A'}
      />
    {:else }
      <InputReadonly label='New Roofing Material - Steep Slope Products'
                     value={estimate?.newRoofingMaterial || 'N/A'} />
    {/if}
    <InputReadonly label='Redeck' value={(estimate?.redeck ?'Yes': 'No')} />
    <InputReadonly label='Layers' value={estimate?.layers} />
    {#if (estimate?.layers > 1)}
      <InputReadonly label='Layer 2 Material' value={estimate?.layer2Material || 'N/A'} />
    {/if}
    {#if (estimate?.layers > 2)}
      <InputReadonly label='Layer 3 Material' value={estimate?.layer3Material || 'N/A'} />
    {/if}
    <InputReadonly label='Measurement Type' value={estimate?.measurementType || 'N/A'} />
    {#if (estimate?.partial > 0)}
      <InputReadonly label='Partial Percentage' value={estimate?.partial || 0} />
    {/if}
  </CardWithBorderTitle>

  <CardWithBorderTitle title='Price Detail'>
    <InputReadonly label='Price' value={usdFormat(estimate?.price)} />
    <InputReadonly label='Total SQ' value={estimate?.totalSquares} />
    <InputReadonly label='Primary Pitch' value={estimate?.primaryPitch} />
    <PriceSummary summary={estimate?.priceSummary} />
    <Doc data={estimate?.pdf} />
  </CardWithBorderTitle>
</div>

{#if !!estimate?.failureReason }
  <div class='card w-full sm:w-11/12 mt-4'>
    <div class='card-body'>
      <H3WithBorder>Failure Reason</H3WithBorder>
      <p class='whitespace-pre'>
        {estimate?.failureReason}
      </p>
    </div>
  </div>
{/if}
{#if estimate?.homeOwner?.latitude && estimate?.homeOwner?.longitude}
  <div class='card w-full sm:w-11/12 mt-4'>
    <div class='card-body'>
      <H3WithBorder> Map</H3WithBorder>
      <Polygon position={{lat:estimate?.homeOwner?.latitude, lng: estimate?.homeOwner?.longitude }} readonly />
    </div>
  </div>
{/if}

<div class='clear-both mt-10 text-sm'>
  Requester: {estimate?.creatorName}
</div>



