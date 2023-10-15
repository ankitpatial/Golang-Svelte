<script lang='ts'>
  import Page from '$lib/components/Page.svelte';
  import { QeyHVACPackages } from '$lib/gql/product';
  import PackageList from '$lib/components/product/PackageList.svelte';
  import { PATH } from '$lib/enum';
  import { InstallationType } from '$lib/graph/graphql';
</script>

<Page title='HVAC Packages'>
  <PackageList
    dataQry={QeyHVACPackages}
    dataProp='hvacPackages'
    installationType={InstallationType.Hvac}
    pendingJobsPath={PATH.HVAC_PENDING}
  />
</Page>
