<script lang='ts'>
  import { goto } from '$app/navigation';
  import Badge from '$lib/components/Badge.svelte';
  import IconArrowLongRight from '$lib/assets/svg/IconArrowLongRight.svelte';
  import { PATH } from '$lib/enum';
  import { getContextClient } from '@urql/svelte';
  import { QryOverview } from '$lib/gql/dashboard';
  import Chart from '$lib/components/Chart.svelte';
  import Page from '$lib/components/Page.svelte';


  const client = getContextClient();
  let filter = 'BY_REGION', data, chartData, loading = false, barChart = false;

  $: {
    load(filter);
  }

  async function load(by) {
    loading = true;
    const res = await client.query(QryOverview, { filter: by }).toPromise();
    loading = false;
    if (res.error) {
      console.error(res.error);
      return;
    }

    data = res.data.overview;
    renderChart(data.items);
  }

  function jobs() {
    goto(PATH.JOBS_UNASSIGNED);
  }

  function renderChart(items) {
    let data = items || [];
    chartData = {
      labels: data.map(i => i.name || 'N/A'),
      datasets: [
        { label: '# Of Jobs', data: data.map(i => i.count), borderWidth: 1 }
      ]
    };
  }

</script>

<Page title='Company Overview' subTitle='({loading ? "loading" : data?.total || 0})'>
  <svelte:fragment slot='filters'>
    <select class='select select-bordered' bind:value={filter}>
      <option value='BY_REGION'>By Region</option>
      <option value='BY_STATE'>By State</option>
    </select>
  </svelte:fragment>

  <div class='flex justify-start'>
    <div class='form-control'>
      <label class='label cursor-pointer'>
        <input type='checkbox' class='toggle toggle-sm toggle-primary mr-2' bind:checked={barChart} />
        <span class='label-text '>Bar Chart</span>
      </label>
    </div>
  </div>

  <!-- Chart Data View -->
  {#if (barChart)}

    <div class='card shadow-md p-4'>
      <div class='flex justify-between'>
        <div>
          <h2 class='text-primary inline-flex gap-x-3 place-items-center group' on:click={jobs}>
            <span> New Unassigned Jobs </span>
            <span class='btn btn-ghost btn-circle btn-sm text-primary group-hover:bg-primary group-hover:text-white'>
            <IconArrowLongRight />
          </span>
          </h2>
        </div>
      </div>

      <div class='flex flex-col sm:flex-row justify-between'>
        <div class='grow'>
          <div class='overflow-visible'>
            {#if (chartData)}
              <Chart data={chartData} type='bar' />
            {/if}
          </div>
        </div>
      </div>

      <div class='flex flex-col md:flex-row justify-between mt-2'>
        <div class='grow'>
        </div>
        <div class='flex sm:flex-row flex-col gap-2 mt-4 sm:mt-0'>
          <Badge type='success'>PAST DUE: SCHEDULE (0)</Badge>
          <Badge>PAST DUE: Completion Doc (0)</Badge>
        </div>
      </div>
    </div>


    <!-- DoNut Data View -->
  {:else }
    <div class='card shadow-md p-4 w-96'>
      <div class='flex justify-between'>
        <div>
          <h2 class='text-primary inline-flex gap-x-3 place-items-center group' on:click={jobs}>
            <span> New Unassigned Jobs </span>
            <span class='btn btn-ghost btn-circle btn-sm text-primary group-hover:bg-primary group-hover:text-white'>
            <IconArrowLongRight />
          </span>
          </h2>
        </div>
      </div>

      <div class='flex flex-col sm:flex-row justify-between'>
        <div class='grow'>
          <div class='overflow-visible'>
            {#if (chartData)}
              <Chart data={chartData} type='doughnut' />
            {/if}
          </div>
        </div>
      </div>

      <div class='flex flex-col md:flex-row justify-between mt-2'>
        <div class='grow'>
        </div>
        <div class='flex sm:flex-row flex-col gap-2 mt-4 sm:mt-0'>
          <Badge type='success'>PAST DUE: SCHEDULE (0)</Badge>
          <Badge>PAST DUE: Completion Doc (0)</Badge>
        </div>
      </div>
    </div>
  {/if}
</Page>






