<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import { createEventDispatcher } from 'svelte';
  import { fade } from 'svelte/transition';
  import Modal from '$lib/components/Modal.svelte';
  import { modalAlert } from '$lib/stores/alerts';
  import Btn from '$lib/components/form/Btn.svelte';

  export let title = 'Please Confirm';
  export let message = 'Are you sure you want to perform this action?';
  export let type: 'link' | 'toggle' | 'btn' = 'btn';
  export let okText = 'Yes';
  export let cls = '';
  export let noteRequired = false;
  export let noteTitle = 'Note';
  export let value = undefined;
  export let disabled = false;
  export let tooltip = '';
  export let size: 'sm' | 'md' | 'lg' = 'md';
  export let confirm: (arg: string | undefined) => Promise<void> = () => Promise.resolve();

  const dispatch = createEventDispatcher();
  let show = false, note = '', noteErr, busy = false;

  $: if (type === 'btn' && cls === '') {
    cls = 'btn btn-ghost btn-sm';
  }

  function showModal() {
    show = true;
  }

  function hideModal() {
    show = false;
  }

  async function handleOK() {
    if (busy) return;
    if (noteRequired && !note) {
      noteErr = 'Required';
      return;
    }

    try {
      busy = true;
      await confirm(noteRequired ? note : undefined);
      hideModal();
    } catch (ex) {
      modalAlert.error(title || 'Error on confirm', ex.message || ex);
    } finally {
      busy = false;
    }
  }
</script>


{#if type === 'link'}
  <div class='tooltip-top' class:tooltip={!!tooltip} data-tip={tooltip} transition:fade>
    <a href='#' on:click|preventDefault={showModal} class='link {cls}' {disabled}>
      <slot></slot>
    </a>
  </div>
{:else if type === 'toggle'}
  <div class='tooltip-top' class:tooltip={!!tooltip} data-tip={tooltip} transition:fade>
    <input
      type='checkbox'
      class='toggle toggle-primary'
      bind:checked={value}
      on:click|preventDefault={showModal}
      {disabled}
    />
  </div>
{:else}
  <div transition:fade>
    <Btn small ghost on:click={showModal} {disabled} {tooltip}>
      <slot></slot>
    </Btn>
  </div>
{/if}

{#if show}
  <Modal size={size} open title={title} {okText} on:ok={handleOK} on:cancel={hideModal} {busy}>
    <slot name='message'>
      {@html message}
    </slot>
    {#if noteRequired}
      <div class='form-control'>
        <label class='label'>
          <span class='label-text'>{noteTitle}</span>
        </label>
        <textarea
          class='textarea textarea-bordered  w-full'
          class:textarea-error={!!noteErr}
          bind:value={note}
          maxlength='500'
        ></textarea>
      </div>
    {/if}
  </Modal>
{/if}


