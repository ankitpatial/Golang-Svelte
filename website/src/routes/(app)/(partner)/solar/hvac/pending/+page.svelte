<script lang='ts'>
  import { InstallationType } from '$lib/graph/graphql';
  import InstallationPendingJobs from '$lib/components/installation/InstallationPendingJobs.svelte';
  import { PATH } from '$lib/enum';
</script>

<InstallationPendingJobs
  title='HVAC Pending Jobs'
  installationType={InstallationType.Hvac}
  basePath={PATH.SOLAR_PARTNER_HVAC_PENDING}
/>
