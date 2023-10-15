<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import {createEventDispatcher, onDestroy} from 'svelte';
  import {default as flatpickr} from 'flatpickr';
  import IconCalendar from './InputDateIcon.svelte';
  import IconSpin from '$lib/assets/svg/IconSpin.svelte';

  import './InputDate.css';
  import Label from "$lib/components/form/Label.svelte";

  export let label = '';
  export let name = '';
  export let placeholder = '';
  export let required = false;
  export let value: Date | undefined = undefined;
  export let disabled = false;
  export let w = 'w-full';
  export let m = 'mb-3';

  // saveOnChange handler, will be used to invoke auto save
  export let saveOnChange: ((val: any) => Promise<void>) | undefined = undefined;

  const inputCls = 'input input-bordered w-full';
  const dateFormat = 'm/d/Y';
  const dispatch = createEventDispatcher();

  let ref = undefined;
  let id = crypto.randomUUID();
  let pickr;
  let busy = false;
  let dirty = false
  let err;

  $:isTip = !!err;
  $:isTipErr = !!err;
  $:tip = err;
  $: if (ref) {
    pickr = flatpickr(ref, {
      dateFormat,
      defaultDate: value
    });
  }

  function changeVal(e) {
    if (disabled) return;

    value = flatpickr.parseDate(e.target.value, dateFormat);
    dispatch('change', value);
    save(value);
  }

  function save(v) {
    if (typeof saveOnChange !== 'function') {
      return;
    }

    dirty = true;
    err = undefined;
    busy = true;
    saveOnChange(v)
      .then(() => {
        busy = false;
        dirty = false;
      })
      .catch((ex) => {
        busy = false;
        err = ex;
      });
  }

  onDestroy(() => {
    if (pickr) {
      pickr.destroy();
    }
  });

  function focus() {
    pickr.jumpToDate(value || new Date());
  }

</script>

<div class='form-control {w} {m}'>
  <Label content={label} {id} uppercase/>

  <div class='relative text-right' class:tooltip={isTip} class:tooltip-error={isTipErr} data-tip={tip}>
		<span class='absolute right-0 mr-3 mt-[7px] bg-white z-20'>
			<IconCalendar/>
		</span>
    <input
      id={id}
      class='{inputCls}'
      class:input-error={isTipErr}
      bind:this={ref}
      {name}
      {required}
      {placeholder}
      {disabled}
      on:input={changeVal}
      on:focus={focus}
    />
    {#if (busy)}
      <div class='relative float-right -mt-8 mr-1'>
        <IconSpin/>
      </div>
    {/if}
  </div>
</div>
