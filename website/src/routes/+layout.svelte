<script lang='ts'>
  import { setContextClient } from '@urql/svelte';
  import { authUser } from '$lib/stores/auth';
  import AlertHub from '$lib/components/AlertHub';
  import { navigating } from '$app/stores';
  import { createPortal } from '$lib/actions/portal.js';
  import Progress from '$lib/components/Progress.svelte';
  import client from '$lib/gql/client';
  import '../app.css';
  import { resolution } from '$lib/stores/responsive';
  import { online } from '$lib/stores/app';

  export let data;

  setContextClient(client());

  authUser.set(data.authUser);
</script>

<svelte:window bind:innerWidth={$resolution} bind:online={$online} />

{#if $navigating}
  <Progress />
{/if}
<slot></slot>
<AlertHub />
<div use:createPortal={'modals'}></div>
