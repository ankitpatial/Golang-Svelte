<!--
  - Copyright (c) 2022. SimSaw Software Private Limited.
  - Unauthorized copying of this file, via any medium is strictly prohibited. Proprietary and confidential.
  - Author: Ankit Patial
  -->
<script lang='ts'>
  import {createEventDispatcher, onDestroy, onMount} from 'svelte';
  import IconSpin from '$lib/assets/svg/IconSpin.svelte';
  import IMask from 'imask';
  import Label from "$lib/components/form/Label.svelte";

  type Validator = (val: any) => Promise<string>

  export let ref = undefined;
  export let id: string = crypto.randomUUID();
  export let label = '';
  export let name = '';
  export let type: 'text' | 'number' | 'range' | 'email' | 'password' = 'text';
  export let placeholder = '';
  export let required = false;
  export let value = undefined;
  export let disabled = false;
  export let focus = false;
  export let autocomplete = 'off';
  export let maskOptions: object | undefined = undefined;
  export let min = undefined;
  export let max = undefined;
  export let maxlength = undefined;
  export let err = '';
  export let w = 'w-full';
  export let m = 'mb-3';
  export let step=''
  export let validate: Validator | undefined = undefined;
  export let handleChange: (event: Event) => any = null;
  // saveOnChange handler, will be used to invoke auto save
  export let saveOnChange: ((val: any) => Promise<void>) | undefined = undefined;

  const dispatch = createEventDispatcher();
  const inputCls = 'input input-bordered w-full';

  let imask;
  let busy = false
  let timeOut;
  let saveTO;
  let mounted = false;

  $:isTip = !!err;
  $:isTipErr = !!err;
  $:tip = err;
  $:if (value && validate) {
    doValidate(value);
  }
  // $:if (imask ) imask.value = 'number' ? Number(value) : imask(value || '').toString();

  onMount(() => {
    mounted = true;
    if (focus) ref.focus();
    if (maskOptions) {
      imask = new IMask(ref, maskOptions);
      imask.on('complete', (e) => {  // accept | complete
        handleChange && handleChange(e);
        value = type === 'number' ? Number(imask.value) : imask.value;
      });
    }
  });

  onDestroy(() => {
    mounted = false
    imask && imask.destroy();
  });

  function changeVal(e) {
    if (disabled || maskOptions) return;

    const val = (type === 'number' || type === 'range') ? Number(e.target.value) : e.target.value;
    if (!maskOptions) {
      value = val;
      //
      handleChange && handleChange(e);
    }

    if (saveTO) clearTimeout(saveTO);
    saveTO = setTimeout(() => {
      if (typeof saveOnChange == 'function') {
        busy = true;
        saveOnChange(value)
          .then(() => {
            busy = false;
          })
          .catch((ex) => {
            busy = false;
            err = ex;
          });
      } else {
        dispatch('change', {name, value: val});
      }
    }, 700);
  }

  function doValidate(value) {
    if (timeOut) clearTimeout(timeOut);

    timeOut = setTimeout(() => {
      busy = true;
      err = '';
      validate && validate(value)
        .then((errMsg) => {
          busy = false;
          err = errMsg;
        })
        .catch((ex) => {
          busy = false;
          err = ex.message || ex;
        });
    }, 500);
  }

</script>

<div class='form-control {w} {m}'>
  {#if (label)}
    <Label uppercase {id} content={label}>
      {#if $$slots.alt}
        <slot name='alt'></slot>
      {/if}
    </Label>
  {/if}

  <div class='tooltip-top' class:tooltip={isTip} class:tooltip-error={isTipErr} data-tip={tip}>
    {#if $$slots.group}
      <label class='input-group'>
        <input
          class='{inputCls}'
          class:input-error={isTipErr}
          bind:this={ref}
          {id}
          {name}
          {type}
          {required}
          {placeholder}
          {value}
          {disabled}
          {autocomplete}
          {min}
          {max}
          {maxlength}
          {step}
          on:input={changeVal}
          on:keypress
        />
        <slot name='group'></slot>
      </label>
    {:else}
      <input
        class='{inputCls}'
        class:input-error={isTipErr}
        bind:this={ref}
        id={id}
        {name}
        {type}
        {required}
        {placeholder}
        {value}
        {disabled}
        {min}
        {max}
        {maxlength}
        {autocomplete}
        {step}
        on:input={changeVal}
        on:keypress
      />
    {/if}

    {#if $$slots.desc}
      <label for={id} class="ml-2">
        <slot name='desc'></slot>
      </label>
    {/if}

    {#if (busy)}
      <div class='relative float-right -mt-8 mr-1'>
        <IconSpin/>
      </div>
    {/if}
  </div>
</div>

