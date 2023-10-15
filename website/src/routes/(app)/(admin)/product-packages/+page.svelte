<script lang='ts'>
  import { goto } from '$app/navigation';
  import { PATH } from '$lib/enum';
  import BtnPlusSmall from '$lib/components/form/BtnPlusSmall.svelte';
  import Page from '$lib/components/Page.svelte';
  import IconGear from '$lib/assets/svg/IconGear.svelte';
  import Btn from '$lib/components/form/Btn.svelte';
  import { QryProductPackages } from '$lib/gql/product';
  import { humanize } from '$lib/utils/string';
  import BtnIcon from '$lib/components/form/BtnIcon.svelte';
  import IconEdit from '$lib/assets/svg/IconEdit.svelte';
  import DataTable from '$lib/components/DataTable/DataTable.svelte';
  import DataColumn from '$lib/components/DataTable/DataColumn.svelte';
  import { type ProductPackage, ProductType } from '$lib/graph/graphql';
  import { onMount } from 'svelte';
  import DropDown from '$lib/components/form/DropDown.svelte';
  import { usdFormat } from '$lib/utils/currency';
  import { editPackage } from './store';

  const title = 'Product Packages';
  let category;
  let variables;
  $: variables = {
    search: '',
    category
  };
  onMount(() => {
    editPackage.set(undefined);
  });

  function add() {
    goto(PATH.PRODUCT_PACKAGES_ADD);
  }

  function edit(item: ProductPackage) {
    editPackage.set(item);
    goto(PATH.PRODUCT_PACKAGES_EDIT);
  }

  function gotoProducts() {
    goto(PATH.PRODUCTS);
  }
</script>

<Page {title}>
  <svelte:fragment slot='buttons'>
    <BtnPlusSmall on:click={add} tooltip='Add New Package' />
    <Btn small on:click={gotoProducts}>
      <IconGear />
      Manage Products
    </Btn>
  </svelte:fragment>
  <svelte:fragment slot='filters'>
    <DropDown
      bind:value={category}
      options={[
        {id: undefined, name: 'All Categories'},
        {id: ProductType.SmartHome, name: 'Smart Home'},
        {id: ProductType.Hvac, name: 'HVAC'},
      ]}
    />
  </svelte:fragment>
  <DataTable
    query={QryProductPackages}
    dataProp='productPackages'
    headers={[
    "CATEGORY",
    "NAME",
    "DESCRIPTION",
    {label: 'PRICE', align: 'right'},
    "EDIT"
  ]}
    {variables}
  >
    <svelte:fragment slot='row' let:item let:headers>
      <DataColumn header={headers[0]}>
        {humanize(item.category)}
      </DataColumn>
      <DataColumn header={headers[1]}>
        {item.name}
      </DataColumn>
      <DataColumn header={headers[2]}>
        {item.description}
      </DataColumn>
      <DataColumn header={headers[3]} align='right'>
        {usdFormat(item.price)}
      </DataColumn>
      <DataColumn header={headers[4]}>
        <BtnIcon on:click={() => edit(item)}>
          <IconEdit />
        </BtnIcon>
      </DataColumn>
    </svelte:fragment>
  </DataTable>
</Page>
