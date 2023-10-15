<script lang='ts'>
  import {authBusy, authUser} from '$lib/stores/auth.js';
  import PlaceholderBasic from "$lib/components/placeholder/PlaceholderBasic.svelte";
</script>

{#if ($authBusy)}
  <PlaceholderBasic/>
{:else if $authUser?.isCompanyAdmin}
  <slot></slot>
{:else}
  <p>
    Only a Company Admin has access to this page.
  </p>
{/if}


