<script lang='ts'>
  import { goto } from '$app/navigation';
  import { PATH } from '$lib/enum';
  import { page } from '$app/stores';
  import TrainingKinds from '$lib/components/training/TrainingKinds.svelte';
  import { QryMyTrainingVideoKinds } from '$lib/gql/training';
  import Page from '$lib/components/Page.svelte';

  const title = 'Training Center';

  function back() {
    goto(PATH.ROOFING_PARTNER + $page.url.search);
  }
</script>

<Page {title}>
  <TrainingKinds qry={QryMyTrainingVideoKinds} dataProp='myTrainingVideoKinds' />
</Page>

