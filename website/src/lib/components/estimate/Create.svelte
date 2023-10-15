<script lang='ts'>
  import alerts from '$lib/stores/alerts';
  import {goto} from '$app/navigation';
  import {getContextClient} from '@urql/svelte';
  import Form from './Form.svelte';
  import PricingCalc from '$lib/components/PricingCalc.svelte';
  import {authUser} from '$lib/stores/auth';
  import {page} from '$app/stores';
  import {type CreateEstimateInput} from "$lib/graph/graphql";
  import {CreateEstimate} from "$lib/gql/estimate";

  export let dashboardRoute

  const client = getContextClient();
  let estimate = {}, mounted = false;

  $:if ($authUser && !mounted) {
    estimate = {
      ownerFirstName: '',
      ownerLastName: '',
      streetNumber: '',
      streetName: '',
      city: '',
      state: '',
      zip: '',
      latitude: undefined,
      longitude: undefined,
      companyName: $authUser?.partner?.partnerName || 'RFX',
      repFirstName: $authUser?.firstName || '',
      repLastName: $authUser?.lastName || '',
      repEmail: $authUser?.email || '',
      repMobile: $authUser?.phone || '',
      currentMaterial: '',
      newRoofingMaterial: '',
      lowSlope: false,
      redeck: false,
      layers: 1,
      measurementType: 'PrimaryStructureOnly',
      partial: 0
    };

    mounted = true;
  }

  async function save(input: CreateEstimateInput) {
    const title = 'Save Job';
    const res = await client.mutation(CreateEstimate, {input}).toPromise();
    if (res.error) {
      // push error to alerts hub
      alerts.push({type: 'error', title, body: res.error.message});
      return;
    }

    await goto(dashboardRoute);
    alerts.push({type: 'success', title, body: 'Successfully Saved'});
  }
</script>

<Form bind:estimate {save}/>

{#if Number($page.url.searchParams.get('calc')) === 1}
  <PricingCalc {job}/>
{/if}
