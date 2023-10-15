<script lang='ts'>
  import Btn from '$lib/components/form/Btn.svelte';
  import IconEye from '$lib/assets/svg/IconEye.svelte';
  import Modal from '$lib/components/Modal.svelte';
  import MDView from '$lib/components/markdown/View.svelte';

  export let summary = '';
  let viewSummary = false;
</script>

{#if (!!summary)}
  <Btn
    cls='mt-2 w-full'
    small
    tooltip='View Summary'
    on:click={() => {viewSummary = true;}}
  >
    <IconEye />
    View Price Summary
  </Btn>
{/if}

{#if (viewSummary)}
  <Modal open title='Price Summary' on:ok={()=> {viewSummary = false}} hideCancelBtn>
    {#if summary.indexOf('| Description |') > -1}
      <MDView content={summary} />
    {:else}
      <p class='whitespace-pre'>
        {summary}
      </p>
    {/if}
  </Modal>
{/if}
