<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->
<script lang='ts'>
  import {QryMyCompanyUsers} from '$lib/gql/user';
  import type Entity from '$lib/models/Entity';
  import {createEventDispatcher} from 'svelte';
  import InputChkWrapper from "$lib/components/form/InputChkWrapper.svelte";
  import DataList from "$lib/components/data/DataList.svelte";

  export let value: Entity;

  const dispatch = createEventDispatcher();
</script>


<DataList
  dataQry={QryMyCompanyUsers}
  dataProp='myCompanyUsers'
  searchable
>
  <svelte:fragment slot='row' let:item>
    <InputChkWrapper label="{item.firstName} {item.lastName} ({item.email})">
      <input
        type='radio'
        name='prods'
        class='radio radio-primary'
        on:input={() => dispatch('change', item)}
      />
    </InputChkWrapper>
  </svelte:fragment>
</DataList>
