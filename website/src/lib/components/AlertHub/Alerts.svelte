<!--
  - Author: Ankit Patial
  -->

<script lang="ts">
  import { slide } from "svelte/transition";
  import { quintOut } from "svelte/easing";
  import type Alert from "$lib/models/Alert";
  import { humanize } from "$lib/utils/string";
  import IconCheckCircle from "$lib/assets/svg/IconCheckCircle.svelte";
  import IconExclamationTriangle from "$lib/assets/svg/IconExclamationTriangle.svelte";
  import IconXCircle from "$lib/assets/svg/IconXCircle.svelte";
  import IconInformationCircle from "$lib/assets/svg/IconInformationCircle.svelte";
  import IconCancel from "$lib/assets/svg/IconCancel.svelte";

  export let store;
  let messages: [Alert] | undefined;

  $:messages = (store && $store || []) as [Alert];

</script>

{#if (true || messages.length > 0)}
  <div class="toast toast-center sm:w-1/2 w-full">
    {#each messages as msg (msg.id)}
      <div class="w-full">
        <div class="alert alert-{msg.type} shadow-lg flex" transition:slide="{{  easing: quintOut }}">
          <div>
            {#if msg.type === 'success'}
              <IconCheckCircle />
            {:else if msg.type === 'warning'}
              <IconExclamationTriangle />
            {:else if msg.type === 'error'}
              <IconXCircle />
            {:else }
              <IconInformationCircle />
            {/if}
          </div>
          <div class="flex-1 text-left">
            <p class="font-bold text-md">{msg.title || humanize(msg.type)}</p>
            <p class="text-md whitespace-pre-wrap">{msg.body}</p>
          </div>
          <div>
            <button class="btn btn-ghost btn-sm" on:click={() => store.remove(msg.id)}>
              <IconCancel />
            </button>
          </div>
        </div>
      </div>
    {/each}
  </div>
{/if}
