<script lang='ts'>
  import BtnPlusSmall from '$lib/components/form/BtnPlusSmall.svelte';
  import { QryNewULID } from '$lib/gql/index.js';
  import alerts from '$lib/stores/alerts.js';
  import { goto } from '$app/navigation';
  import { getContextClient } from '@urql/svelte';

  export let baseURL: string;

  const client = getContextClient();

  async function createSurvey() {
    return to(`${baseURL}/:id/save`);
  }

  async function to(path) {
    const res = await client.query(QryNewULID, {}).toPromise();
    if (res.error) {
      alerts.error('New ULID Error', res.error.message);
      return;
    }
    await goto(path.replace(':id', res.data.newULID));
  }

</script>

<BtnPlusSmall on:click={createSurvey} />
<button class='btn btn-outline uppercase btn-sm' on:click={createSurvey}>
  Create a Survey
</button>
