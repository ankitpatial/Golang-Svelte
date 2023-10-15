<script lang='ts'>
  import BtnPlusSmall from '$lib/components/form/BtnPlusSmall.svelte';
  import { goto } from '$app/navigation';
  import { PATH } from '$lib/enum';
  import DataTable from '$lib/components/DataTable/DataTable.svelte';
  import { QryProducts } from '$lib/gql/product';
  import DataColumn from '$lib/components/DataTable/DataColumn.svelte';
  import { humanize } from '$lib/utils/string';
  import { onMount } from 'svelte';
  import BtnIcon from '$lib/components/form/BtnIcon.svelte';
  import IconEdit from '$lib/assets/svg/IconEdit.svelte';
  import type { ProductInfo } from '$lib/graph/graphql';
  import { editProduct } from './store';
  import Page from '$lib/components/Page.svelte';
  import BtnBackSmall from '$lib/components/form/BtnBackSmall.svelte';
  import { usdFormat } from '$lib/utils/currency';
  import DropDown from '$lib/components/form/DropDown.svelte';
  import { ProductType } from '$lib/graph/graphql';

  const title = 'Manage Products';
  let category;
  let variables;
  $: variables = {
    search: '',
    category
  }

  onMount(() => {
    editProduct.set(undefined);
  });

  function back() {
    goto(PATH.PRODUCT_PACKAGES);
  }

  function gotoAdd() {
    goto(PATH.PRODUCTS_ADD);
  }

  function edit(item: ProductInfo) {
    editProduct.set(item);
    goto(PATH.PRODUCTS_EDIT);
  }

</script>

<Page {title}>
  <svelte:fragment slot='buttons'>
    <BtnBackSmall on:click={back} tooltip='View Packages' />
    <BtnPlusSmall on:click={gotoAdd} tooltip='Add New Product' />
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
    query={QryProducts}
    dataProp='products'
    {variables}
    headers={[
    "CATEGORY",
    "IMAGE",
    "NAME",
    "DESCRIPTION",
    {label: "PRICE", align: 'right'},
    "EDIT"
  ]}
  >
    <svelte:fragment slot='row' let:item let:headers>
      <DataColumn header={headers[0]}>
        {humanize(item.category)}
      </DataColumn>
      <DataColumn header={headers[1]}>
        <img class='w-24' src={item.image.publicUrl} alt='product' />
      </DataColumn>
      <DataColumn header={headers[2]}>
        {item.name}
      </DataColumn>
      <DataColumn header={headers[3]}>
        {item.description}
      </DataColumn>
      <DataColumn header={headers[4]} align='right'>
        {usdFormat(item.price)}
      </DataColumn>
      <DataColumn header={headers[5]}>
        <BtnIcon on:click={() => edit(item)}>
          <IconEdit />
        </BtnIcon>
      </DataColumn>
    </svelte:fragment>
  </DataTable>
</Page>
