<!--
- Copyright (c) Simsaw Software Pvt. Ltd. 2023.
- Author: Ankit Patial <ankit@simsaw.com>
-->

<script lang="ts">
  import {PATH} from '$lib/enum'
  import Estimates from "$lib/components/estimate/Estimates.svelte";

  const title = 'Completed Estimates';
  const basePath = PATH.ESTIMATES
</script>

<Estimates {title} {basePath}/>
