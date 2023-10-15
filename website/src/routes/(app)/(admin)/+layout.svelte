<script lang='ts'>
  import Authorize from '$lib/components/Authorize.svelte';
  import {Role} from "$lib/graph/graphql";

  export let data; // will come from root layout's load()
</script>

<Authorize role={Role.Admin}>
  <slot></slot>
</Authorize>
