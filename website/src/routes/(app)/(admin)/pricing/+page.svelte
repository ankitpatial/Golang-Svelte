<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->
<script lang='ts'>
  import { getContextClient } from '@urql/svelte';
  import { QryPricing } from '$lib/gql/pricing';
  import alerts from '$lib/stores/alerts';
  import Table from '$lib/components/DataTable/Table.svelte';
  import DataColumn from '$lib/components/DataTable/DataColumn.svelte';
  import { usdFormat } from '$lib/utils/currency.js';
  import Page from '$lib/components/Page.svelte';

  const client = getContextClient();
  let headers = [
      'State, City - Zip'
    ],
    cities,
    products,
    loading = true;

  client
    .query(QryPricing, {})
    .toPromise()
    .then((res) => {
      if (res.error) {
        alerts.push({ type: 'error', title: 'Get Pricing', body: res.error.message });
        return;
      }

      const d = res.data.getPricing;
      products = d.products;
      headers = [].concat(headers, products.map(o => o.name));

      let list = [];
      (d.items || []).forEach((i) => {
        let city = list.find(c => c.zip === i.zip);
        if (!city) {
          city = {
            zip: i.zip,
            name: i.city,
            state: i.state
          };
          list.push(city);
        }

        let name = `p${i.productId}`;
        let prod = city[name];
        if (!prod) {
          prod = {
            id: i.id,
            price: i.price
          };
          city[name] = prod;
        }
      });

      cities = list;
    })
    .catch((ex) => console.error(ex))
    .finally(() => {
      loading = false;
    });

</script>

<Page title='Pricing'>
  <Table headers={headers} {loading} items={cities}>
    <svelte:fragment slot='row' let:headers let:item>
      <DataColumn header={headers[0]}>
        {item.state}, {item.name} - {item.zip}
      </DataColumn>
      {#each products as prod, idx}
        <DataColumn header={headers[idx + 1]}>
          <div class='text-right'>
            {usdFormat(item[`p${prod.id}`]?.price || 0)}
          </div>
        </DataColumn>
      {/each}
    </svelte:fragment>
  </Table>
</Page>
