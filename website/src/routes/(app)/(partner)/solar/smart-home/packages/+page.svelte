<script lang='ts'>
  import Page from '$lib/components/Page.svelte';
  import { QrySmartHomePackages } from '$lib/gql/product';
  import PackageList from '$lib/components/product/PackageList.svelte';
  import { PATH } from '$lib/enum';
  import { InstallationType } from '$lib/graph/graphql';
</script>

<Page title='Smart Home Packages'>
  <PackageList
    dataQry={QrySmartHomePackages}
    dataProp='smartHomePackages'
    installationType={InstallationType.SmartHome}
    pendingJobsPath={PATH.SOLAR_PARTNER_SMART_HOME_PENDING}
  />
</Page>
