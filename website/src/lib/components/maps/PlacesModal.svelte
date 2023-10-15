<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import {createEventDispatcher} from 'svelte';
  import Modal from '$lib/components/Modal.svelte';
  import Places from './Places.svelte';
  import type {Address} from './type';

  export let label;
  export let cls = '';

  const dispatch = createEventDispatcher();

  let show = false, address: Address;

  function open() {
    show = true;
  }

  function close() {
    show = false;
  }

  function change({detail}) {
    address = detail;
  }

  function ok() {
    dispatch('ok', address);
    close();
  }

</script>

<div class='link text-sm {cls} mt-0' on:click={open}>
  {label}
</div>

{#if show}
  <Modal open size='lg' title='Search Places' okText='Use This Location' on:ok={ok} on:cancel={close}>
    <Places on:change={change}/>
  </Modal>
{/if}
