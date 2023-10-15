<script lang='ts'>
  import { InstallationType } from '$lib/graph/graphql';
  import { PATH } from '$lib/enum';
  import InstallationApprovedJobs from '$lib/components/installation/InstallationApprovedJobs.svelte';
</script>

<InstallationApprovedJobs
  title='Smart Home Approved Jobs'
  installationType={InstallationType.SmartHome}
  basePath={PATH.SMART_HOME_APPROVED}
/>
