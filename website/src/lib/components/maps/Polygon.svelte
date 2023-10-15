<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import {createEventDispatcher, onMount} from 'svelte';
  import mapMount from "$lib/components/maps/mapMount";

  export let position = {lat: 40.749933, lng: -73.98633}
  export let bounds = [];
  export let readonly = false

  const dispatch = createEventDispatcher();
  let map, drawMgr, selectedShape;

  onMount(mapMount(initMap));

  // initMap
  function initMap(): void {
    try {
      map = new google.maps.Map(
        document.getElementById('map') as HTMLElement,
        {
          center: position,
          zoom: 20,
          mapTypeControl: false,
          mapTypeId: 'satellite',
          streetViewControl: false,
          fullscreenControl: false
        }
      );

      map.setTilt(0);

      const marker = new google.maps.Marker({map, anchorPoint: position,});
      marker.setPosition(position);
      marker.setVisible(true);

      draw();
    } catch (ex) {
      console.trace(ex);
    }
  }

  function draw() {
    if (!Array.isArray(bounds) || bounds.length === 0) return;
    selectedShape = new google.maps.Polygon({paths: bounds});
    if (!readonly) {
      google.maps.event.addListener(selectedShape, 'click', function () {
        setSelection(selectedShape, true);
      });
    }
    selectedShape.setMap(map);
    change();
  }

  function startDrawing() {
    drawMgr = new google.maps.drawing.DrawingManager({
      drawingMode: google.maps.drawing.OverlayType.POLYGON,
      drawingControl: false,
      drawingControlOptions: {
        position: google.maps.ControlPosition.TOP_CENTER,
        drawingModes: ['polygon']
      },
      polygonOptions: {
        editable: true
      }

    });
    drawMgr.setMap(map);

    google.maps.event.addListener(drawMgr, 'overlaycomplete', (e) => {
      e.overlay.set('editable', false);
      drawMgr.setMap(null);

      let newShape = e.overlay;
      newShape.type = e.type;
      google.maps.event.addListener(newShape, 'click', function () {
        setSelection(newShape, true);
      });

      setSelection(newShape);
    });
  }

  // set selected shape
  function setSelection(shape, edit = false) {
    clearSelection();
    selectedShape = shape;
    shape.setEditable(edit);
    change();
  }

  function clearSelection() {
    if (selectedShape) {
      selectedShape.setEditable(false);
      selectedShape = null;
    }
  }

  function deleteSelectedShape() {
    if (selectedShape) {
      selectedShape.setMap(null);
    }
    change();
  }

  function change() {
    dispatch('change', selectedShape);
  }

</script>

<div class="absolute top-0 right-0 mt-5 mr-10 " class:hidden={readonly}>
  <button class="btn btn-outline btn-sm" on:click={startDrawing}>
    draw
  </button>
  <button class="btn btn-outline btn-sm" on:click={deleteSelectedShape}>
    clear
  </button>
</div>

<div id='pac-card' class='pac-card float-left'>
</div>

<div id='map'></div>

<style>
  #map {
    height: 50vh;
  }

  #map #map-inf-win {
    display: inline;
  }

  .pac-card {
    background-color: #fff;
    border: 0;
    border-radius: 2px;
    box-shadow: 0 1px 4px -1px rgba(0, 0, 0, 0.3);
    margin: 10px;
    padding: 0 0.5em;
    font: 400 18px Roboto, Arial, sans-serif;
    overflow: hidden;
  }

  #pac-container {
    padding-bottom: 12px;
    margin-right: 12px;
  }
</style>
