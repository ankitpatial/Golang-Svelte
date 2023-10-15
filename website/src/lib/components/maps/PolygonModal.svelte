<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import {createEventDispatcher} from 'svelte';
  import Modal from '$lib/components/Modal.svelte';
  import Polygon from './Polygon.svelte';

  export let title = "Draw Polygon"
  export let show = false;
  export let position = {lat: 40.749933, lng: -73.98633}
  export let bounds = [];
  export let readonly = false

  const dispatch = createEventDispatcher();
  let polygon;

  function open() {
    show = true;
  }

  function close() {
    show = false;
  }

  function change({detail}) {
    polygon = detail;
  }

  function getPolygonBounds() {
    if (!polygon) {
      return;
    }
    const bounds = []
    polygon.getPath().getArray().forEach(b => {
      bounds.push({lat: b.lat(), lng: b.lng()})
    })
    return bounds;
  }

  function ok() {
    const bounds = getPolygonBounds();
    dispatch('done', bounds);
  }
</script>

{#if show}
  <Modal
    open size='lg'
    {title}
    okText='Done'
    cancelText="Close"
    hideOkBtn={readonly}
    on:ok={ok}
    on:cancel={close}
  >
    <Polygon {position} on:change={change} {bounds} {readonly}/>
  </Modal>
{/if}
