<!--
  - Copyright (c) 2022. SimSaw Software Private Limited.
  - Unauthorized copying of this file, via any medium is strictly prohibited. Proprietary and confidential.
  - Author: Ankit Patial
  - Dated:  20/05/22, 2:51 PM
  -->

<script>
  import {createEventDispatcher, onMount} from 'svelte';

  export let once = false;
  export let top = 0;
  export let bottom = 0;
  export let left = 0;
  export let right = 0;

  const dispatch = createEventDispatcher();
  let intersecting = false;
  let container;

  onMount(() => {
    if (typeof IntersectionObserver !== 'undefined') {
      const rootMargin = `${bottom}px ${left}px ${top}px ${right}px`;

      const observer = new IntersectionObserver(entries => {
        intersecting = entries[0].isIntersecting;
        if (intersecting) {
          dispatch('intersect');
          if (once) observer.unobserve(container);
        }
      }, {
        rootMargin
      });

      observer.observe(container);
      return () => observer.unobserve(container);
    }
  });
</script>

<style>
  div {
    width: 100%;
    height: 100%;
  }
</style>

<div bind:this={container}>
  <slot {intersecting}></slot>
</div>
