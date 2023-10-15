<script>
  import { onMount } from 'svelte';
  import { MutSurveyRequest, QrySurveySlotAvailability } from '$lib/gql/survey';
  import { goto } from '$app/navigation';
  import { getContextClient } from '@urql/svelte';
  import InputDate from '$lib/components/form/dtp/InputDate.svelte';
  import DropDown from '$lib/components/form/DropDown.svelte';
  import BtnSubmit from '$lib/components/form/BtnSubmit.svelte';
  import { PATH } from '$lib/enum';
  import Page from '$lib/components/Page.svelte';

  let date = '';
  let slotID = '';
  let slots = [];
  let formattedSlots = [];
  $: formattedSlots = slots.map(slot => ({ id: slot.id, name: slot.name }));

  const client = getContextClient();

  async function fetchSlots() {
    const formattedDate = new Date(date).toISOString().split('T')[0];
    const result = await client.query(QrySurveySlotAvailability, { date: formattedDate }).toPromise();
    if (result.error) {
      console.error('Error:', result.error);
    } else {
      slots = result.data.surveySlotAvailability;
    }
  }

  async function reserveSlot() {
    const formattedDate = new Date(date).toISOString().split('T')[0];
    const result = await client.mutation(MutSurveyRequest, {
      date: formattedDate,
      slotID
    }).toPromise();

    if (result.error) {
      console.error('Error:', result.error);
    } else {
      console.log('Reserved slot:', result.data.surveyRequest);
      await goto(PATH.SOLAR_PARTNER_SURVEYS);
    }
  }

  onMount(async () => {
    // Fetch available slots for the selected date
    await fetchSlots();
  });

</script>

<Page title='Reserve Slot'>
  <form on:submit|preventDefault={reserveSlot}>
    <label>Date:</label>
    <InputDate bind:value={date} on:change={fetchSlots} required />

    <label>Slot:</label>
    <DropDown bind:value={slotID} options={formattedSlots} required on:change={(e) => (slotID = e.detail)} />

    <BtnSubmit>Reserve Slot</BtnSubmit>
  </form>
</Page>
