<!--
  - Copyright (c) 2022. SimSaw Software Private Limited.
  - Unauthorized copying of this file, via any medium is strictly prohibited. Proprietary and confidential.
  - Author: Ankit Patial
  - Dated:  26/04/22, 1:37 PM
  -->

<script lang='ts'>
  import type Entity from '$lib/models/Entity';
  import IconSpin from '$lib/assets/svg/IconSpin.svelte';
  import IconSearch from '$lib/assets/svg/IconSearch.svelte';
  import Chip from '../Chip.svelte';
  import {createEventDispatcher} from 'svelte';
  import Label from "$lib/components/form/Label.svelte";

  export let label: string = '';
  export let placeholder: string = 'Search...';
  export let value: [Entity] | Entity | null = null;
  export let multiple = false;
  export let noSelection = false;
  export let cls = 'w-full mb-4';
  // search fn is required to perform search
  export let searchHandler: (text: string) => Promise<[Entity]>;
  export let render: (i: any) => string = (i: Entity) => {
    return i.name;
  };

  const watchInterval = 400; // ms
  const dispatch = createEventDispatcher();

  let id = crypto.randomUUID();
  let inputRef;
  let ul;
  let opIdx = -1;
  let searchText;
  let timeOut;
  let isFocus = false;
  let listFocus = false;
  let loading = false;
  let options: [Entity] = [];
  let optionsToShow;
  let selection: [Entity] = [];
  let isSelected = false;

  $: selection = value ? [].concat(value) : [];
  $: optionsToShow = !selection ? options : options.filter(o => selection.findIndex(p => p.id === o.id) === -1);
  $: if (!isFocus || !listFocus) {
    opIdx = -1;
  }

  // debounce type input
  function watch() {
    if (timeOut) {
      clearTimeout(timeOut);
    }

    timeOut = setTimeout(() => {
      search(searchText);
    }, watchInterval);
  }

  async function search(q: string) {
    isSelected = false;
    loading = true;
    try {
      options = await searchHandler(decode(q)) || [];
    } catch (err) {
      console.error(err);
    } finally {
      loading = false;
    }
  }

  function decode(str) {
    let txt = new DOMParser().parseFromString(str, 'text/html');
    return txt.documentElement.textContent;
  }

  // option click
  function handleSelection(val) {
    searchText = '';
    isSelected = true;
    const op = options.find(o => o.id == val);
    if (noSelection) {
      dispatch('change', op);
      return;
    }

    selection = multiple ? selection.concat(op) : [op];
    value = multiple ? selection : op;
    dispatch('change', value);
  }

  // option, on enter
  function inpKeyDown(e) {
    const key = e.code;
    // handle enter key;
    if (key === 'Enter') {
      e.preventDefault();
      if (opIdx === -1) {
        return; // no option selected
      }
      // set selection
      const key = optionsToShow[opIdx].id;
      handleSelection(key);
      opIdx = -1; // clear option selected
      return;
    }

    if (['ArrowDown', 'ArrowUp'].indexOf(key) === -1) return;

    if (key === 'ArrowDown' && opIdx < optionsToShow.length - 1) {
      opIdx += 1;
    }

    if (key === 'ArrowUp' && opIdx > 0) {
      opIdx -= 1;
    }
  }

  // remove selection
  function handleRemoveSelection(id) {
    selection = (selection || []).filter(v => v.id !== id);
    value = multiple ? selection : selection[0];
    dispatch('change', value);
    inputRef.focus();
  }

  function inpFocus() {
    isFocus = true;
    optionsToShow = [];
    if (searchText) {
      search(searchText);
    }
  }

  function inpBlur() {
    isFocus = false;
  }

  export function setFocus() {
    inputRef?.focus();
  }

</script>

<div class='form-control relative {cls}' on:focus={setFocus} on:click={setFocus}>
  {#if (label)}
    <Label uppercase {id} content={label} />
  {/if}

  <div class='flex border rounded-md relative pl-4 pb-2 max-h-fit '>
		<span class='top-0 -ml-2 mt-2 mr-2'>
			{#if loading}
				<IconSpin cls='h-5 w-5'/>
			{:else }
      	<IconSearch/>
			{/if}
    </span>
    <div class='mt-0 flex flex-row'>
      {#if (!noSelection)}
        {#each selection as op}
          <Chip content={render(op)} on:cancel={() => handleRemoveSelection(op.id)}/>
        {/each}
      {/if}
    </div>

    {#if (!isFocus && !searchText)}
      <div class='mt-2 text-gray-400 w-full'>
        {placeholder}
      </div>
    {/if}

    <div
      id={id}
      bind:this={inputRef}
      class='focus:outline-0 focus:border-primary w-full mt-1.5'
      contenteditable='true'
      bind:innerHTML={searchText}
      on:focus={inpFocus}
      on:blur={inpBlur}
      on:keydown={inpKeyDown}
      on:input={watch}
    ></div>

  </div>

  {#if (isFocus || listFocus) && !isSelected && ($$slots.create || optionsToShow.length)}
    <div
      bind:this={ul}
      class='menu bg-base-100 rounded-b-box border w-full mt-12 absolute z-50 '
      on:mouseover={() => {listFocus = true}}
      on:mouseout={() => {listFocus = false}}
    >
      <ul class="menu dropdown-content p-2 shadow-md bg-base-100 rounded-box">
        {#if optionsToShow.length}
          {#each optionsToShow as option, idx}
            <li id='op-{idx}-{id}' on:click={() => { handleSelection(option.id)}}>
							<span class:active={opIdx === idx}>
								{render(option)}
							</span>
            </li>
          {/each}
        {/if}

        {#if $$slots.create}
          <li>
            <slot name='create'></slot>
          </li>
        {/if}
      </ul>
    </div>
  {/if}
</div>
