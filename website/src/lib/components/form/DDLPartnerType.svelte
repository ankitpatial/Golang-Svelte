<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->
<script lang='ts'>
  import InputReadonly from '$lib/components/form/InputReadonly.svelte';
  import DropDown from '$lib/components/form/DropDown.svelte';
  import {PartnerType} from "$lib/graph/graphql";

  export let label = 'Type';
  export let value = '';
  export let readOnly = false;

  const partnerTypes = [
    {id: PartnerType.Roofing, name: 'Roofing'},
    {id: PartnerType.Solar, name: 'Solar'},
    {id: PartnerType.Epc, name: 'EPC'},
    {id: PartnerType.Integration, name: 'Integration'},
  ];
</script>

{#if readOnly}
  <InputReadonly {label} value={value}/>
{:else}
  <DropDown {label} bind:value={value} options={partnerTypes}/>
{/if}
