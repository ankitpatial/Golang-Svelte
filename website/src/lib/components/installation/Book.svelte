<script lang='ts'>
  import { createForm } from 'svelte-forms-lib';
  import { number, object, string } from 'yup';
  import Input from '$lib/components/form/Input.svelte';
  import type { Package } from '$lib/graph/graphql';
  import { InstallationType } from '$lib/graph/graphql';
  import Btn from '$lib/components/form/Btn.svelte';
  import { createEventDispatcher } from 'svelte';
  import { usdFormat } from '$lib/utils/currency';
  import PlacesModal from '$lib/components/maps/PlacesModal.svelte';
  import InputPhone from '$lib/components/form/InputPhone.svelte';
  import TextArea from '$lib/components/form/TextArea.svelte';
  import { getContextClient } from '@urql/svelte';
  import { BookInstallation } from '$lib/gql/installation';
  import alerts from '$lib/stores/alerts';

  export let installationType: InstallationType;
  export let pkg: Package;
  export let productID = undefined;

  const client = getContextClient();
  const dispatch = createEventDispatcher();
  const { form, errors, isSubmitting, handleChange, handleSubmit } = createForm({
    initialValues: {
      name: '',
      address: '',
      email: '',
      phone: '',
      specialNote: '',
      latitude: undefined,
      longitude: undefined
    },
    validationSchema: object().shape({
      name: string().required('Name is required').max(100),
      address: string().required('Address is required').max(250),
      email: string().required('Email is required').email('Invalid email address').max(150),
      phone: string().required('Phone is required').max(15),
      specialNote: string().optional().max(500),
      latitude: number().optional(),
      longitude: number().optional()
    }),
    onSubmit: async (values) => {
      const title = 'Book An Installation';
      const res = await client.mutation(BookInstallation, {
        type: installationType,
        pkgID: pkg.id,
        productID,
        owner: {
          name: values.name,
          email: values.email,
          phone: values.phone,
          address: values.address,
          latitude: values.latitude,
          longitude: values.longitude,
          specialNote: values.specialNote
        }
      }).toPromise();
      if (res.error) {
        alerts.error(title, res.error.message);
        return;
      }

      alerts.success(title, 'Successfully Booked An Installation');
      dispatch('saved');
    }
  });

  function cancel() {
    dispatch('cancel');
  }

  function setAddress({ detail }) {
    if (!detail) return;

    $form.address = detail.address;
    $form.latitude = detail.latitude;
    $form.longitude = detail.longitude;
  }
</script>

<div class='max-w-2xl'>
  <form on:submit={handleSubmit}>
    <Input
      label='Package'
      value={pkg.name}
      disabled
    />
    <Input
      label='Price'
      value={usdFormat(pkg.price)}
      disabled
    />
    <Input
      label='Customer Full Name'
      name='name'
      bind:value={$form.name}
      err={$errors.name}
      {handleChange}
    />
    <Input
      label='Customer Address'
      name='address'
      bind:value={$form.address}
      err={$errors.address}
      {handleChange}
    >
      <PlacesModal label='find address on map' on:ok={setAddress} slot='alt' />
    </Input>

    <Input
      label='Customer Email'
      name='email'
      bind:value={$form.email}
      err={$errors.email}
      {handleChange}
    />
    <InputPhone
      label='Customer Phone Number'
      name='phone'
      bind:value={$form.phone}
      err={$errors.phone}
      {handleChange}
    />

    <TextArea
      label='Special Note'
      name='specialNote'
      bind:value={$form.specialNote}
      err={$errors.specialNote}
      {handleChange}
    />

    <div class='flex flex-row gap-5'>
      <div>
        <Btn type='submit' primary busy={$isSubmitting}>
          BOOK
        </Btn>
      </div>
      <div>
        <Btn on:click={cancel} disabled={$isSubmitting}>
          Cancel
        </Btn>
      </div>
    </div>
  </form>
</div>
