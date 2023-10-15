<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->
<script lang='ts'>
  import { getContextClient } from '@urql/svelte';
  import SearchCities from '$lib/components/search/SearchCities.svelte';
  import SearchStates from '$lib/components/search/SearchStates.svelte';
  import Chip from '$lib/components/Chip.svelte';
  import { MarkServiceArea, QryAllServiceAreas } from '$lib/gql/postalCode';
  import alerts from '$lib/stores/alerts';
  import H3WithBorder from '../../../../lib/components/H3WithBorder.svelte';
  import Page from '$lib/components/Page.svelte';

  const client = getContextClient();

  let states = {},
    busy = false;

  client.query(QryAllServiceAreas, {}).toPromise()
    .then((res) => {
      if (res.error) {
        return;
      }
      // set states
      const data = res.data.allServiceAreas || [];
      const stateNames = data.map(i => i.name);
      stateNames.sort();
      stateNames.forEach(s => {
        states[s] = data.find(i => i.name === s).cities;
      });
    });


  function addState({ detail }) {
    const stateName = detail.name;
    if (states[stateName] !== undefined) return;

    const stateNames = Object.keys(states).concat(stateName);
    stateNames.sort();
    const changed = {};
    stateNames.forEach(s => {
      changed[s] = states[s] || [];
    });

    states = changed;
  }

  function addCity(stateName, city) {
    const cities = states[stateName] || [];
    if (cities.findIndex(c => c.id === city.id) > -1) { // exclude already added one
      return;
    }

    addStateCity(stateName, city);

    // mark service area = true
    client
      .mutation(MarkServiceArea, { id: city.id, value: true })
      .toPromise()
      .then((res) => {
        if (res.error) {
          // remove from queue
          alerts.error('Failed To Add City ', `adding city '${city.name}' failed with error: ${res.error.message}`);
          removeStateCity(stateName, city.id);
        }
      });
  }

  function removeCity(stateName, cityID) {
    const removed = removeStateCity(stateName, cityID);
    client
      .mutation(MarkServiceArea, { id: cityID, value: false })
      .toPromise()
      .then((res) => {
        if (res.error) {
          // remove from queue
          alerts.error('Failed To Remove City ', `removing city '${removed.name}' failed with error: ${res.error.message}`);
          addStateCity(stateName, removed);
        }
      });
  }

  function addStateCity(stateName, city) {
    const cities = states[stateName];
    cities.push(city);
    cities.sort((a, b) => {
      if (a.name < b.name) {
        return -1;
      }
      if (a.name > b.name) {
        return 1;
      }

      return 0;
    });

    states[stateName] = cities;
  }

  function removeStateCity(stateName, cityID) {
    const cities = states[stateName];
    const idx = cities.findIndex(c => c.id === cityID);
    if (idx === -1) return;

    const removed = cities.splice(idx, 1);
    states[stateName] = cities;
    return removed[0];
  }
</script>

<Page title='Service Areas'>
  <svelte:fragment slot='buttons'>
    <SearchStates
      placeholder='States...'
      cls='w-52'
      on:change={addState}
    />
  </svelte:fragment>

  <div class='w-8/12 ml-4' class:blur-sm={busy}>
    {#each Object.entries(states) || [] as [stateName, cities]}
      <div class='flex flex-row place-items-center justify-between mb-2 mt-4'>
        <H3WithBorder>{stateName}</H3WithBorder>
        <SearchCities
          placeholder='City...'
          state={stateName}
          on:change={({detail})=> addCity(stateName, detail)}
          ignore={(cities || []).map(c=>({id: c.id}))}
          cls='w-52 bg-white'
        />
      </div>
      <div class='flex flex-row flex-wrap mb-8 ml-4 border-t border-slate-200'>
        {#each cities || [] as city (city.id)}
          <Chip content='{city.name} - {city.zip}' on:cancel={() => removeCity(stateName, city.id)} />
        {/each}
      </div>
    {/each}
  </div>
</Page>
