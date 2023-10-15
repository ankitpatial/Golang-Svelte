<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import Create from '$lib/components/estimate/Create.svelte';
  import { PATH } from '$lib/enum';
  import Page from '$lib/components/Page.svelte';
</script>

<Page title='Request An Estimate'>
  <Create dashboardRoute={PATH.ESTIMATES_VIEW} />
</Page>


