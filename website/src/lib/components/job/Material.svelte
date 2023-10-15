<script>
  import CardWithBorderTitle from '$lib/components/CardWithBorderTitle.svelte';
  import InputReadonly from '$lib/components/form/InputReadonly.svelte';

  export let data;

  $: isCurrentLowSlopeOnly = data?.currentMaterial === 'Low Slope Only';
</script>

<CardWithBorderTitle title='Job Details'>
  <InputReadonly label='Current Material - Existing Steep Slope Material'
                 value={data?.currentMaterial || 'N/A'} />
  {#if isCurrentLowSlopeOnly}

    <InputReadonly
      label='New Low Slope Material - Low Slope Products'
      value={data?.newRoofingMaterialLowSlope || 'N/A'}
    />

  {:else }
    <InputReadonly
      label='New Roofing Material - Steep Slope Products'
      value={data?.newRoofingMaterial || 'N/A'}
    />
  {/if}

  <InputReadonly label='Redeck' value={(data?.redeck ?'Yes': 'No')} />
  <InputReadonly label='Layers' value={data?.layers} />
  {#if (data?.layers > 1)}
    <InputReadonly label='Layer 2 Material' value={data?.layer2Material || 'N/A'} />
  {/if}
  {#if (data?.layers > 2)}
    <InputReadonly label='Layer 3 Material' value={data?.layer3Material || 'N/A'} />
  {/if}
  <InputReadonly label='Measurement Type' value={data?.measurementType || 'N/A'} />
  {#if (data?.partial > 0)}
    <InputReadonly label='Partial Percentage' value={data?.partial || 0} />
  {/if}
</CardWithBorderTitle>
