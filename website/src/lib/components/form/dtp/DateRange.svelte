<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  - ROOF-24: chane  Yesterday l  Today l  This Week  l  Last Week  l  Last Month  l  MTD l  YTD
  -->

<script lang='ts'>
  import {createEventDispatcher, onDestroy} from 'svelte';
  import {default as flatpickr} from 'flatpickr';

  import './InputDate.css';
  import {
    dtSubtractDays,
    endOfDay,
    lastMonth,
    lastWeek,
    localDate,
    startOfDay,
    thisMonth,
    thisWeek,
    thisYear
  } from '$lib/utils/time';

  export let label = '';
  export let name = '';
  export let placeholder = '';
  export let required = false;
  export let value: Array<Date>;
  export let disabled = false;
  export let w = 'w-full';
  export let m = 'mb-3';
  export let maxDate = endOfDay(new Date());

  const dateFormat = 'm/d/Y';
  const dispatch = createEventDispatcher();

  let ref = undefined;
  let id = crypto.randomUUID();
  let pickr, busy = false, dirty = false, err, ddlVal, display;

  $:isTip = !!err;
  $:isTipErr = !!err;
  $:tip = err;
  $: if (ref && !pickr) {
    pickr = flatpickr(ref, {
      mode: 'range',
      dateFormat,
      defaultDate: value,
      maxDate,
      onClose: (selectedDates) => {
        if (!checkDate(value, selectedDates)) {
          value = selectedDates;
        }
      }
    });
  }
  $: {
    check(value);
    setDisplay(value);
  }

  onDestroy(() => {
    if (pickr) {
      pickr.destroy();
    }
  });

  // check the selected date range and set the quick matching option
  function changeVal(e: any) {
    let r: Array<Date>;
    let t: Date;
    switch (e.target.value) {
      case 'ALL':
        r = [];
        break;
      case 'YESTERDAY':
        t = dtSubtractDays(new Date(), 1);
        r = [startOfDay(new Date(t.getTime())), endOfDay(new Date(t.getTime()))];
        break;
      case 'TODAY':
        r = [startOfDay(new Date()), endOfDay(new Date())];
        break;
      case 'THIS_WEEK':
        r = thisWeek();
        break;
      case 'LAST_WEEK':
        r = lastWeek();
        break;
      case 'LAST_MONTH':
        r = lastMonth();
        break;
      case 'THIS_MONTH':
        r = thisMonth();
        break;
      case 'THIS_YEAR':
        r = thisYear();
        break;
    }

    value = r;
    pickr.selectedDates = r;
    pickr.redraw();

    dispatch('change', r)
  }

  function check(v: Array<Date>) {
    let check, t: Date;
    ddlVal = '';

    if (!v || !Array.isArray(v) || v.length < 2) {
      ddlVal = 'ALL';
      return;
    }

    // YESTERDAY
    t = dtSubtractDays(new Date(), 1);
    check = [t, t];
    if (checkDate(v, check)) {
      ddlVal = 'YESTERDAY';
      return;
    }

    // TODAY
    check = [new Date(), new Date()];
    if (checkDate(v, check)) {
      ddlVal = 'TODAY';
      return;
    }

    // THIS_WEEK
    check = thisWeek();
    if (checkDate(v, check)) {
      ddlVal = 'THIS_WEEK';
      return;
    }

    // LAST_WEEK
    check = lastWeek();
    if (checkDate(v, check)) {
      ddlVal = 'LAST_WEEK';
      return;
    }

    // LAST_MONTH
    check = lastMonth();
    if (checkDate(v, check)) {
      ddlVal = 'LAST_MONTH';
      return;
    }

    // THIS_MONTH
    check = thisMonth();
    if (checkDate(v, check)) {
      ddlVal = 'THIS_MONTH';
      return;
    }

    // THIS_YEAR
    check = thisYear();
    if (checkDate(v, check)) {
      ddlVal = 'THIS_YEAR';
      return;
    }
  }

  function setDisplay(v: Array<Date>) {
    if (!Array.isArray(v) || v.length < 2) {
      // localDate(v, false);
      display = '';
      return '';
    }

    const f = localDate(v[0], false);
    const t = localDate(v[1], false);
    display = `${f} - ${t}`;
  }

  function checkDate(a, b) {
    return localDate(a[0], false) == localDate(b[0], false) && localDate(a[1], false) == localDate(b[1], false);
  }

  function focus() {
    pickr.jumpToDate(value || new Date());
  }

</script>

<div class='form-control'>
  <div class='input-group flex justify-between'>
    <select class='select select-bordered grow' value={ddlVal} on:change={changeVal}>
      <option value='' disabled>{display || ''}</option>
      <option value='ALL'>All</option>
      <option value='YESTERDAY'>Yesterday</option>
      <option value='TODAY'>Today</option>
      <option value='THIS_WEEK'>This Week</option>
      <option value='LAST_WEEK'>Last Week</option>
      <option value='LAST_MONTH'>Last Month</option>
      <option value='THIS_MONTH'>MTD</option>
      <option value='THIS_YEAR'>YTD</option>
    </select>

    <div class='dropdown dropdown-end' id={id} bind:this={ref}>
      <div class='btn btn-outline btn-square bg-white' on:click={ref.focus}>
        <svg xmlns='http://www.w3.org/2000/svg' fill='none' viewBox='0 0 24 24' stroke-width='1.5' stroke='currentColor'
             class='w-6 h-6'>
          <path stroke-linecap='round' stroke-linejoin='round'
                d='M6.75 3v2.25M17.25 3v2.25M3 18.75V7.5a2.25 2.25 0 012.25-2.25h13.5A2.25 2.25 0 0121 7.5v11.25m-18 0A2.25 2.25 0 005.25 21h13.5A2.25 2.25 0 0021 18.75m-18 0v-7.5A2.25 2.25 0 015.25 9h13.5A2.25 2.25 0 0121 11.25v7.5'/>
        </svg>
      </div>
    </div>
  </div>

</div>

<div class='form-control {w} {m} hidden'>
  <label class='label p-0' for={id}>
    <span class='text-xs uppercase'>
      {label}
    </span>
  </label>

</div>
