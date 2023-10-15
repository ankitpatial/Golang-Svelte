<!--
 Copyright (c) 2022. Ankit Patial.
 Unauthorized copying of this file, via any medium is strictly prohibited. Proprietary and confidential.
 Author: Ankit Patial
  -->

<script>
  import Image from '$lib/components/Image.svelte';

  export let name = '';
  export let src = '';
  export let imgClass = 'w-8 mask mask-circle';
  export let placeHolderClass = 'w-auto bg-slate-200 text-gray-700 mask mask-circle';
  export let className = 'mask mask-square';
  export let txtClass = 'text-xl p-2 font-bold';

  let hasImg = false;
  let placeHolder = 'N/A';
  let nameArr;

  $: {
    nameArr = name && name.split(' ');
    if (nameArr && nameArr.length === 1) {
      placeHolder = nameArr[0].substring(0, 2);
    } else if (nameArr) {
      placeHolder = '';
      nameArr.forEach(n => {
        placeHolder += n.substring(0, 1);
      });
    }
  }
</script>

<div class='mx-auto flex-none {className}'>
  {#if src}
    <Image src={src} alt={(name || 'image')} cls={imgClass} on:done={(e) => { hasImg = e.detail.success }}/>
  {:else}
    <div class='avatar placeholder'>
      <div class='{placeHolderClass}'>
        <span class='{txtClass}'>{(placeHolder || 'N/A').toUpperCase()}</span>
      </div>
    </div>
  {/if}
</div>
