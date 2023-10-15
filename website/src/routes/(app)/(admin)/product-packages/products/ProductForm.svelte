<script lang='ts'>

  import { createForm } from 'svelte-forms-lib';
  import { array, number, object, string } from 'yup';
  import Input from '$lib/components/form/Input.svelte';
  import Btn from '$lib/components/form/Btn.svelte';
  import TextArea from '$lib/components/form/TextArea.svelte';
  import IconPlusSmall from '$lib/assets/svg/IconPlusSmall.svelte';
  import BtnIcon from '$lib/components/form/BtnIcon.svelte';
  import IconTrash from '$lib/assets/svg/IconTrash.svelte';
  import InputNumber from '$lib/components/form/InputNumber.svelte';
  import InputPublicImage from '$lib/components/form/InputPublicImage.svelte';
  import DropDown from '$lib/components/form/DropDown.svelte';
  import { ProductType } from '$lib/graph/graphql';
  import { getContextClient } from '@urql/svelte';
  import { SaveProduct } from '$lib/gql/product';
  import alerts from '$lib/stores/alerts';
  import { createEventDispatcher } from 'svelte';

  export let id = '';
  export let isEdit = false;
  export let data = undefined;

  const client = getContextClient();
  const dispatch = createEventDispatcher();
  const title = 'Save Product';
  const maxPrice = 99999;
  const { form, errors, handleChange, handleSubmit, isSubmitting } = createForm({
    initialValues: !!data
      ? {
        id: data.id,
        imageID: data.image.id,
        category: data.category,
        name: data.name,
        description: data.description,
        price: data.price,
        features: data.features,
        specialNote: data.specialNote
      }
      : {
        id,
        imageID: '',
        category: ProductType.SmartHome,
        name: '',
        description: '',
        features: [''],
        specialNote: ''
      },
    validationSchema: object().shape({
      id: string().required('id is required'),
      imageID: string().required('image is required'),
      category: string().required('category is required'),
      name: string().required('name is required'),
      description: string().required('description is required'),
      price: number().required('price is required').min(0).max(maxPrice),
      features: array().of(string().required('required')),
      specialNote: string().optional()
    }),
    onSubmit: async (input) => {
      const res = await client.mutation(SaveProduct, { input }).toPromise();
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

  $:console.log(data);
</script>

<form on:submit={handleSubmit}>
  <div class='flex flex-col sm:flex-row gap-2'>
    <div class='w-full'>
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
      <Input
        label='Product Name'
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
      {#each $form.features as feature, idx}
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
      <div class='mt-5'>
        <Input
          label='Special Note'
          name='specialNote'
          bind:value={$form.specialNote}
          err={$errors.specialNote}
          cls=''
          {handleChange}
        />
      </div>
    </div>
    <div class='w-full sm:w-1/2'>
      <div class='card'>
        <div class='card-body'>
          <InputPublicImage
            uploadBtnLabel='Upload Poster'
            docDir={id}
            docURL={data?.image.publicUrl}
            docName='product.jpeg'
            docSection='Image'
            aspectRatio={1}
            on:complete={({detail}) => {
              $form.imageID = detail.id;
            }}
          />
          {#if ($errors.imageID)}
            <p class='badge badge-error'>{$errors.imageID}</p>
          {/if}
        </div>
      </div>
    </div>
  </div>
  <Btn type='submit' cls='mt-10' busy={$isSubmitting} primary>
    SAVE
  </Btn>
</form>
