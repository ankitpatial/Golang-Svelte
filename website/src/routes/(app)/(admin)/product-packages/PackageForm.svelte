<script lang='ts'>

  import { createForm } from 'svelte-forms-lib';
  import { array, number, object, string } from 'yup';
  import Input from '$lib/components/form/Input.svelte';
  import Btn from '$lib/components/form/Btn.svelte';
  import TextArea from '$lib/components/form/TextArea.svelte';
  import IconPlusSmall from '$lib/assets/svg/IconPlusSmall.svelte';
  import BtnIcon from '$lib/components/form/BtnIcon.svelte';
  import IconTrash from '$lib/assets/svg/IconTrash.svelte';
  import DropDown from '$lib/components/form/DropDown.svelte';
  import { type ProductsQuery, ProductType, SoldAs } from '$lib/graph/graphql';
  import { getContextClient } from '@urql/svelte';
  import { QryProducts, SaveProductPackage } from '$lib/gql/product';
  import alerts, { modalAlert } from '$lib/stores/alerts';
  import { createEventDispatcher } from 'svelte';
  import Label from '$lib/components/form/Label.svelte';
  import Modal from '$lib/components/Modal.svelte';
  import DataList from '$lib/components/data/DataList.svelte';
  import InputChkWrapper from '$lib/components/form/InputChkWrapper.svelte';

  export let id = '';
  export let isEdit = false;
  export let data = undefined;

  let showAddProduct = false;
  let prods;

  const client = getContextClient();
  const dispatch = createEventDispatcher();
  const title = 'Save Product';
  const maxPrice = 99999;
  const { form, errors, isValid, handleChange, handleSubmit, isSubmitting } = createForm({
    initialValues: !!data
      ? {
        id: data.id,
        category: data.category,
        soldAs: data.soldAs,
        name: data.name,
        description: data.description,
        price: data.price,
        features: data.features,
        items: (data.items || []).map(i => ({ id: i.id, name: i.name, imageURL: i.image.publicUrl }))
      }
      : {
        id,
        imageID: '',
        category: ProductType.SmartHome,
        soldAs: SoldAs.Package,
        name: '',
        description: '',
        features: [''],
        items: []
      },
    validationSchema: object().shape({
      id: string().required('ID is required'),
      category: string().required('Category is required'),
      soldAs: string().required('Sold As is required'),
      name: string().required('Name is required'),
      description: string().when('soldAs', {
        is: (v) => v === SoldAs.Package,
        then: (s) => s.required('Description is required')
      }),
      price: number().when('soldAs', {
        is: (v) => v === SoldAs.Package,
        then: (s) => s.required('Price is required').min(0).max(maxPrice)
      }),
      features: array().of(string().optional()).when('soldAs', {
        is: (v) => v === SoldAs.Package,
        then: (s) => s.of(string().required('required'))
      }),
      items: array().min(1, 'add a product').of(object().shape({
        id: string().required(),
        name: string().optional(),
        imageURL: string().optional()
      }))
    }),
    onSubmit: async (values) => {
      const res = await client.mutation(SaveProductPackage, {
        input: {
          id: values.id,
          category: values.category,
          soldAs: values.soldAs,
          name: values.name,
          description: values.description,
          price: values.price || 0,
          features: values.features,
          productIDs: values.items.map(p => p.id)
        }
      }).toPromise();
      if (res.error) {
        alerts.error(title, res.error.message);
        return;
      }

      alerts.success(title, 'Successfully saved');
      dispatch('saved');
    }
  });

  function addFeature() {
    form.update(f => {
      f.features = f.features.concat('');
      return f;
    });
  }

  function removeFeature(idx) {
    if (idx <= 0) return;

    form.update(f => {
      f.features.splice(idx, 1);
      return f;
    });
  }

  function addItems() {
    form.update(f => {
      f.items = f.items.concat(prods.map(p => ({id: p.id, name: p.name, imageURL: p.image.publicUrl})));
      return f;
    });
    toggleAddProduct();
  }

  function removeItem(idx) {
    form.update(f => {
      f.items.splice(idx, 1);
      return f;
    });
  }

  function toggleAddProduct() {
    prods = [];
    showAddProduct = !showAddProduct;
  }

  async function searchHandler(search: string) {
    const res = await client.query<ProductsQuery>(QryProducts, {
      category: $form.category,
      search,
      page: { first: 10 }
    });
    if (res.error) {
      modalAlert.error('Search Product', res.error.message);
      return;
    }

    return res.data?.products?.edges?.map(e => {
      return ({
        id: e.node?.id,
        name: e.node?.name,
        imageURL: e.node?.image?.publicUrl
      });
    });
  }
</script>

<form on:submit={handleSubmit}>
  <DropDown
    label='Category'
    name='category'
    options={[
          {id: ProductType.SmartHome, name: 'Smart Home'},
          {id: ProductType.Hvac, name: 'HVAC'}
        ]}
    bind:value={$form.category}
    err={$errors.category}
    {handleChange}
  />
  <DropDown
    label='Sold As'
    name='soldAs'
    options={[
          {id: SoldAs.Package, name: 'Package'},
          {id: SoldAs.IndividualItem, name: 'Individual Item'}
        ]}
    bind:value={$form.soldAs}
    err={$errors.soldAs}
    {handleChange}
  />
  <Input
    label='Package Name'
    name='name'
    bind:value={$form.name}
    err={$errors.name}
    {handleChange}
  />
  <TextArea
    label='Description'
    name='description'
    bind:value={$form.description}
    err={$errors.description}
    {handleChange}
  />
  <Input
    type='number'
    step='any'
    label='Price'
    name='price'
    bind:value={$form.price}
    err={$errors.price}
    {handleChange}
    max={maxPrice}
    showRange={false}
  />
  {#each $form.features || [] as feature, idx}
    {@const label = idx ? '' : "Features"}
    <div class='flex flex-row gap-2'>
      <Input
        label={label}
        name='features[{idx}]'
        placeholder='enter text...'
        bind:value={$form.features[idx]}
        err={$errors.features[idx]}
        {handleChange}
      />
      {#if (idx) > 0}
        <BtnIcon on:click={()=>removeFeature(idx)}>
          <IconTrash color='red' />
        </BtnIcon>
      {/if}
    </div>
  {/each}
  <Btn small on:click={addFeature}>
    <IconPlusSmall />
    Feature
  </Btn>

  <div class='form-control'>
    <Label content='Products' cls='mt-3' />
  </div>
  {#each $form.items || [] as p, idx}
    <div class='flex flex-row gap-2'>
      <div class='w-32'>
        <img src={p.imageURL} alt='product' />
      </div>
      <div class='w-52'>
        {p.name}
      </div>
      <div>
        <BtnIcon on:click={()=>removeItem(idx)}>
          <IconTrash color='red' />
        </BtnIcon>
      </div>
    </div>
  {/each}
  {#if (!$isValid && $form.items.length === 0)}
    <p class='badge badge-error mt-3 mb-3'>Please add a product</p>
  {/if}
  <Btn small on:click={toggleAddProduct} cls='mt-5'>
    <IconPlusSmall />
    Product
  </Btn>
  <Btn type='submit' cls='mt-10' busy={$isSubmitting}>
    SAVE
  </Btn>
</form>

<Modal open={showAddProduct} title='Search A Product' on:ok={addItems} on:cancel={toggleAddProduct}>
  <DataList
    dataQry={QryProducts}
    dataProp='products'
    variables={{
      category: $form.category,
    }}
    searchable
    on:search={() => {
      prods = [];
    }}
  >
    <svelte:fragment slot='row' let:item>
      <InputChkWrapper label={item.name}>
        <input
          type='checkbox'
          name='prods'
          class='checkbox checkbox-primary'
          value={item}
          bind:group={prods}
        />
      </InputChkWrapper>
    </svelte:fragment>
  </DataList>
</Modal>
