<script lang="ts">
  import ConfirmAction from "$lib/components/ConfirmAction.svelte";
  import {SetPartnerActive} from "$lib/gql/partner";
  import alerts from "$lib/stores/alerts";
  import {getContextClient} from "@urql/svelte";
  import {PartnerStatus} from "$lib/graph/graphql";

  export let partnerID: string;
  export let name: string;
  export let status: string;
  export let store: any;
  let isActive: boolean;

  const client = getContextClient();
  let busy = false;

  $: isActive = (status == PartnerStatus.Active)

  async function changeActive() {
    busy = true;
    const value = !isActive;
    const res = await client.mutation(SetPartnerActive, {partnerID, value}).toPromise();
    busy = false;
    if (res.error) {
      alerts.error('Error', `failed to ${value ? 'activate' : 'deactivate'} partner`);
      return;
    }

    status = value ? PartnerStatus.Active : PartnerStatus.InActive;
  }
</script>

<ConfirmAction
  type='toggle'
  title="{isActive ? 'Deactivate' : 'Activate'} Partner <strong>{name}</strong>"
  message={isActive
                            ? 'Deactivating a partner will remove their access to use the website. Are you sure you want to proceed?'
                            : 'Activate a partner will give them access to use the website. Are you sure you want to proceed?'
                            }
  value={isActive}
  confirm={changeActive}
  disabled={(status === PartnerStatus.Onboarding)}
  tooltip={isActive ? 'Deactivate' : 'Activate'}
/>
