<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import { PATH } from '$lib/enum.js';
  import { goto } from '$app/navigation';
  import Form from '../Form.svelte';
  import BtnBackSmall from '$lib/components/form/BtnBackSmall.svelte';
  import Page from '$lib/components/Page.svelte';

  function back() {
    goto(PATH.API_ACCESS);
  }
</script>

<Page title='New API Access'>
  <svelte:fragment slot='buttons'>
    <BtnBackSmall on:click={back} />
  </svelte:fragment>

  <Form />
</Page>
