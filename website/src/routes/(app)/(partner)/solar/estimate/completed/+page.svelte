<!--
- Copyright (c) Simsaw Software Pvt. Ltd. 2023.
- Author: Ankit Patial <ankit@simsaw.com>
-->

<script lang="ts">
  import {PATH} from '$lib/enum'
  import Estimates from "$lib/components/estimate/Estimates.svelte";

  const title = 'Completed Estimates';
  const ignore = ['Request'];
  const basePath = PATH.SOLAR_PARTNER_EST_COMPLETED
</script>

<Estimates {title} {basePath} {ignore}/>
