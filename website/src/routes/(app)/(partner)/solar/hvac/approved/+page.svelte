<script lang='ts'>
  import { PATH } from '$lib/enum';
  import { InstallationType } from '$lib/graph/graphql';
  import InstallationApprovedJobs from '$lib/components/installation/InstallationApprovedJobs.svelte';
</script>

<InstallationApprovedJobs
  title='HVAC Approved Jobs'
  installationType={InstallationType.Hvac}
  basePath={PATH.SOLAR_PARTNER_HVAC_APPROVED}
/>
