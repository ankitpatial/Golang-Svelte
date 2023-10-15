<script>
  import { onMount } from 'svelte';
  import { page } from '$app/stores';
  import { getContextClient } from '@urql/svelte';
  import alerts from '$lib/stores/alerts.ts';
  import { MutSurveyReserve } from '$lib/gql/survey.ts';
  import Input from '$lib/components/form/Input.svelte';
  import TextArea from '$lib/components/form/TextArea.svelte';
  import BtnSubmit from '$lib/components/form/BtnSubmit.svelte';
  import Page from '$lib/components/Page.svelte';

  let id;
  let name = '';
  let phoneNumber = '';
  let address = '';
  let notes = '';

  const client = getContextClient();

  async function fillSurvey() {
    const result = await client
      .mutation(MutSurveyReserve, {
        input: {
          id: $page.params.id,
          name,
          phoneNumber,
          address,
          notes
        }
      })
      .toPromise();

    if (result.error) {
      alerts.error('Error filling survey', result.error.message);
    } else {
      alerts.success('Survey successfully filled');
    }
  }

  onMount(async () => {
    const { id: surveyId } = $page.params.id;
    id = surveyId;
  });
</script>

<Page title='Survey Request'>
  <form on:submit|preventDefault={fillSurvey}>
    <Input label='Name' bind:value={name} />
    <Input label='Phone Number' bind:value={phoneNumber} />
    <Input label='Address' bind:value={address} />
    <TextArea label='Notes' bind:value={notes} />
    <BtnSubmit>Submit</BtnSubmit>
  </form>
</Page>
