<script>
  import { onMount } from 'svelte';
  import { getContextClient } from '@urql/svelte';
  import { QrySurveyDetails } from '$lib/gql/survey';
  import { page } from '$app/stores';
  import Page from '$lib/components/Page.svelte';

  let surveyDetails = null;
  const client = getContextClient();

  onMount(async () => {
    const id = $page.params.id; // Get the survey id from the URL
    const result = await client.query(QrySurveyDetails, { id }).toPromise();

    if (result.error) {
      console.error('Error:', result.error);
    } else {
      surveyDetails = result.data.surveyDetails;
    }
  });
</script>

<Page title='Survey Detail'>


  {#if surveyDetails}
    <div>
      <h2>Name</h2>
      <p>{surveyDetails.name}</p>
    </div>
    <div>
      <h2>Date:</h2>
      <p>{surveyDetails.date}</p>
    </div>
    <div>
      <h2>Slot:</h2>
      <p>{surveyDetails.slot}</p>
    </div>
    <div>
      <h2>Status:</h2>
      <p>{surveyDetails.status}</p>
    </div>
    <div>
      <h2>Address:</h2>
      <p>{surveyDetails.address}</p>
    </div>
    <div>
      <h2>Phone Number:</h2>
      <p>{surveyDetails.phone}</p>
    </div>
  {:else}
    <p>Loading survey details...</p>
  {/if}
</Page>
