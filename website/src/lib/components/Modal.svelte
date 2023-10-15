<!--
  - Copyright (c) 2022. SimSaw Software Private Limited.
  - Unauthorized copying of this file, via any medium is strictly prohibited. Proprietary and confidential.
  - Author: Ankit Patial
  - Dated:  02/05/22, 6:07 PM
  -->

<script lang='ts'>
  import { createEventDispatcher } from 'svelte';
  import { portal } from '$lib/actions/portal.js';
  import { modalAlert } from '$lib/stores/alerts';
  import Alert from '$lib/components/AlertHub/Alerts.svelte';
  import Btn from '$lib/components/form/Btn.svelte';

  export let open = false;
  export let busy = false;
  export let title;
  export let hideOkBtn = false;
  export let disableOkBtn = false;
  export let hideCancelBtn = false;
  export let okText = 'OK';
  export let cancelText = 'CANCEL';
  export let size: 'sm' | 'md' | 'lg' = 'md';
  export let cls = '';

  const dispatch = createEventDispatcher();
  let w;

  $:{
    switch (size) {
      case 'sm':
        w = '';
        break;
      case 'md':
        w = 'max-w-2xl';
        break;
      case 'lg':
        w = 'w-11/12 max-w-5xl';
    }
  }

  function ok() {
    open && dispatch('ok');
  }

  function cancel() {
    open && dispatch('cancel');
  }

  function escKeyPress({ code }) {
    if (open && code === 'Escape') {
      dispatch('escape');
    }
  }
</script>

<svelte:window on:keyup={escKeyPress} />

{#if open}
  <div use:portal={'modals'} class='modal' class:modal-open={open}>
    <div class='modal-box {w} {cls} pb-1'>
      <h2>{@html title}</h2>
      <div class='py-4'>
        <slot></slot>
      </div>
      <div class='modal-action border-t py-3'>
        <slot name='actions'>
          {#if !hideOkBtn}
            <Btn primary on:click={ok} {busy} disabled={busy || disableOkBtn}>
              {okText}
            </Btn>
          {/if}
          {#if !hideCancelBtn}
            <Btn on:click={cancel} disabled={busy}>
              {cancelText}
            </Btn>
          {/if}
        </slot>
      </div>
    </div>
    <Alert store={modalAlert} />
  </div>
{/if}
