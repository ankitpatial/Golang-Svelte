<script lang='ts'>
  import { QrySurveys } from '$lib/gql/survey';
  import { goto } from '$app/navigation';
  import { DataColumn, DataTable } from '$lib/components/DataTable';
  import BtnView from '$lib/components/form/BtnView.svelte';
  import BtnIcon from '$lib/components/form/BtnIcon.svelte';
  import IconUserPlus from '$lib/assets/svg/IconUserPlus.svelte';
  import { page } from '$app/stores';
  import DateRange from '$lib/components/form/dtp/DateRange.svelte';
  import { localDate, thisMonth } from '$lib/utils/time.js';
  import { PATH } from '$lib/enum';
  import Page from '$lib/components/Page.svelte';

  let count = 0,
    dt: Array<Date> = $page.url.searchParams.has('dt')
      ? JSON.parse($page.url.searchParams.get('dt'))
      : thisMonth(),
    partner: { id: '', name: '' } | undefined;
  $: variables = {
    search: $page.url.searchParams.get('q'),
    betweenDates: dt,
    orderBy: {
      direction: 'DESC',
      field: 'CREATED'
    }
  };

  function fillSurvey(id) {
    goto(`dashboard/${id}/fill-survey` + $page.url.search);
  }

  function gotoDetail(id) {
    goto(`${$page.url.pathname}/${id}` + $page.url.search);
  }

  function dtChange({ detail }) {
    dt = detail;
    $page.url.searchParams.set('dt', JSON.stringify(dt));
  }
</script>


<Page title='Solar Partner Surveys' subTitle='({count})'>
  <svelte:fragment slot='buttons'>
    <BtnIcon on:click={() => goto(PATH.SOLAR_PARTNER_SURVEY_REQUEST)} tooltip='Reserve Slot'>
      <IconUserPlus />
    </BtnIcon>
  </svelte:fragment>
  <svelte:fragment slot='filters'>
    <DateRange value={dt} on:change={dtChange} />
  </svelte:fragment>
  <div class='w-full sm:w-full'>
    <div class='card-body p-0'>
      <DataTable
        headers={['DATE', 'SLOT', 'STATUS', 'SURVEY DETAILS', 'ACTIONS']}
        query={QrySurveys}
        dataProp='surveys'
        {variables}
        searchable
        bind:totalCount={count}
      >
        <svelte:fragment slot='row' let:item let:headers>
          <DataColumn header={headers[0]}>
            {localDate(item.date, false)}
          </DataColumn>
          <DataColumn header={headers[1]}>
            {item.slot}
          </DataColumn>
          <DataColumn header={headers[2]}>
            {item.status}
          </DataColumn>
          <DataColumn header={headers[3]}>
            <BtnView on:click={() => gotoDetail(item.id)} tooltip='View Survey Details' />
          </DataColumn>
          <DataColumn header={headers[4]}>
            {#if item.status === 'REQUESTING' || item.status === 'REQUESTED'}
              <BtnIcon on:click={() => fillSurvey(item.id)} tooltip='Fill Survey'>
                <IconUserPlus />
              </BtnIcon>
            {/if}
          </DataColumn>
        </svelte:fragment>
      </DataTable>
    </div>
  </div>
</Page>
