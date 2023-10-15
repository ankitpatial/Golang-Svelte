<script lang='ts'>
  import Create from '$lib/components/estimate/Create.svelte';
  import { PATH } from '$lib/enum';
  import Page from '$lib/components/Page.svelte';
</script>

<Page title='Request An Estimate'>
  <Create dashboardRoute={PATH.SOLAR_PARTNER_EST_COMPLETED} />
</Page>
