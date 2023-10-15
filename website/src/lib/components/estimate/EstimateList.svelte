<script lang='ts'>
  import { page } from '$app/stores';
  import { QryEstimates } from '$lib/gql/estimate';
  import { usdFormat } from '$lib/utils/currency';
  import { localDate } from '$lib/utils/time';
  import { wsMessage } from '$lib/stores/socket';
  import { DataColumn, DataTable } from '$lib/components/DataTable';
  import EstimateActions from './EstimateActions.svelte';
  import { onMount } from 'svelte';
  import { authUser } from '$lib/stores/auth';
  import { QryPartnerOptions } from '$lib/gql/partner';
  import { getContextClient } from '@urql/svelte';
  import { Channel, EstimateStatus, Topic } from '$lib/graph/graphql';

  export let dt;
  export let status;
  export let totalCount = 0;
  export let ignore = [];
  export let basePath = '';

  const client = getContextClient();
  let store;
  let variables;
  let epcs = [];
  $:variables = {
    status: status || undefined,
    search: $page.url.searchParams.get('q'),
    dates: dt
  };

  onMount(() => {
    if ($authUser?.isCompanyAdmin) {
      client.query(QryPartnerOptions, { partnerID: $authUser.partner?.id }).toPromise()
        .then((res) => {
          if (res.error) {
            console.log(res.error.message);
            return;
          }

          const d = res.data.partnerOptions || [];
          d.forEach(o => {
            switch (o.type) {
              case 'EPC':
                epcs = o.options;
                break;
            }
          });
        });
    }

    return wsMessage.subscribe(msg => {
      if (!msg) return;
      switch (msg.channel) {
        case Channel.Estimate:
          if (Topic.StatusChange === msg.topic) {
            switch (msg.data.status) {
              case EstimateStatus.Pending:
                variables._v = Date.now(); // reload to fetch pricing
                break;
              default:
                changeStatus({ detail: { id: msg.data.id, status: msg.data.status } });
            }
          }
          break;
        case Channel.Task:
          if (msg.data.done && msg.data.task === 'estimate') {
            variables._v = Date.now();
          }
          break;
      }
    });
  });

  function changeStatus({ detail }) {
    if (!store) return;

    store.update(s => {
      const i = s.find(i => i.id === detail.id);
      if (i) { // update
        i.status = detail.status;
      } else { // reload
        variables._v = Date.now();
      }
      return s;
    });
  }
</script>

<div class='w-full sm:w-full'>
  <div class='card-body p-0'>
    <DataTable
      headers={['HOME OWNER', 'ADDRESS', 'SOLAR COMPANY', 'REQUESTER', 'PRICE', 'STATUS', 'CREATED' , 'ACTIONS']}
      query={QryEstimates}
      dataProp='estimates'
      {variables}
      searchable
      bind:totalCount
      cbSetStore={(s) => {store = s;}}
    >
      <svelte:fragment slot='row' let:headers let:item>
        <DataColumn header={headers[0]}>
          {item.homeOwner.firstName} {item.homeOwner.lastName}
        </DataColumn>
        <DataColumn header={headers[1]} whitespaceWrap>
          {item.homeOwner.streetNumber} {item.homeOwner.streetName} <br>
          {item.homeOwner.city} {item.homeOwner.state}, {item.homeOwner.zip}
        </DataColumn>
        <DataColumn header={headers[2]}>
          {item.companyName}
        </DataColumn>
        <DataColumn header={headers[3]}>
          {item.salesRep?.firstName} {item.salesRep?.lastName}
        </DataColumn>
        <DataColumn header={headers[4]}>
          {item.price > 0 ? usdFormat(item.price) : 'N/A'}
        </DataColumn>
        <DataColumn header={headers[5]}>
          {#if EstimateStatus.Approved === item.status}
            Approved - Moved to production
          {:else if EstimateStatus.Failed === item.status}
            <span class='badge badge-error tooltip tooltip-top tooltip-error' data-tip={item.failureReason}>
              {item.status}
            </span>
          {:else }
            {item.status}
          {/if}
        </DataColumn>
        <DataColumn header={headers[6]}>
          {localDate(item.createdAt)}
        </DataColumn>
        <DataColumn header={headers[7]}>
          <EstimateActions {item} {basePath} {epcs} on:changeStatus={changeStatus} />
        </DataColumn>
      </svelte:fragment>
    </DataTable>
  </div>
</div>
