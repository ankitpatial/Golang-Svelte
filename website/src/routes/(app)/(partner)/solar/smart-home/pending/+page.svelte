<script lang='ts'>
  import InstallationPendingJobs from '$lib/components/installation/InstallationPendingJobs.svelte';
  import { InstallationType } from '$lib/graph/graphql';
  import { PATH } from '$lib/enum';
</script>

<InstallationPendingJobs
  title='Smart Home Pending Jobs'
  installationType={InstallationType.SmartHome}
  basePath={PATH.SOLAR_PARTNER_SMART_HOME_PENDING}
/>
