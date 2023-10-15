<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import {createEventDispatcher, onMount} from 'svelte';
  import mapMount from "$lib/components/maps/mapMount";

  const dispatch = createEventDispatcher();
  let map;

  onMount(mapMount(initMap));

  // initMap
  function initMap(): void {
    try {
      map = new google.maps.Map(
        document.getElementById('map') as HTMLElement,
        {
          center: {lat: 40.749933, lng: -73.98633},
          zoom: 15,
          mapTypeControl: false,
          mapTypeId: 'satellite',
          streetViewControl: false,
          fullscreenControl: false
        }
      );
      map.setTilt(0);

      const card = document.getElementById('pac-card') as HTMLElement;
      const input = document.getElementById('map-inp') as HTMLElement;
      const options = {
        fields: ['formatted_address', 'address_components', 'geometry', 'name'],
        strictBounds: false,
        types: [] // address, establishment, geocode, regions, cities
      };

      input.focus();

      map.controls[google.maps.ControlPosition.TOP_LEFT].push(card);
      const autocomplete = new google.maps.places.Autocomplete(input, options);
      // Bind the map's bounds (viewport) property to the autocomplete object,
      // so that the autocomplete requests use the current map bounds for the
      // bounds option in the request.
      autocomplete.bindTo('bounds', map);

      const infoWin = new google.maps.InfoWindow();
      const infoWinContent = document.getElementById(
        'map-inf-win'
      ) as HTMLElement;

      infoWin.setContent(infoWinContent);

      let place;
      const marker = new google.maps.Marker({
        map,
        anchorPoint: new google.maps.Point(0, -29),
        draggable: true
      });

      // ref: https://developers.google.com/maps/documentation/javascript/reference/marker
      marker.addListener('dragend', () => {
        map.setCenter(marker.getPosition());
        change(place, marker);
      });

      autocomplete.addListener('place_changed', () => {
        infoWin.close();
        marker.setVisible(false);
        place = autocomplete.getPlace();
        placeChange(map, place, marker, infoWinContent);
        infoWin.open(map, marker);
      });
    } catch (ex) {
      console.trace(ex);
    }
  }

  function placeChange(map, place, marker, infoWinContent) {
    if (!place.geometry || !place.geometry.location) {
      // User entered the name of a Place that was not suggested and
      // pressed the Enter key, or the Place Details request failed.
      window.alert('No details available for input: \'' + place.name + '\'');
      return;
    }

    // If the place has a geometry, then present it on a map.
    if (place.geometry.viewport) {
      map.fitBounds(place.geometry.viewport);
    } else {
      map.setCenter(place.geometry.location);
      map.setZoom(17);
    }

    marker.setPosition(place.geometry.location);
    marker.setVisible(true);

    change(place, marker);

    const plcName = infoWinContent.children['place-name'];
    if (plcName) {
      plcName.textContent = place.name;
    }

    const plcAddress = infoWinContent.children['place-address'];
    if (plcAddress) {
      plcAddress.textContent = place.formatted_address;
    }
  }

  function change(place, marker) {
    let stNo, st, city, state, zip, country;
      (place.address_components || []).forEach((cmp) => {
      if (cmp.types.includes('street_number')) {
        stNo = cmp.long_name;
      } else if (cmp.types.includes('route')) {
        st = cmp.long_name;
      } else if (cmp.types.includes('locality') || cmp.types.includes('neighborhood')) {
        city = cmp.long_name;
      } else if (cmp.types.includes('administrative_area_level_1')) {
        state = cmp.long_name;
      } else if (cmp.types.includes('postal_code')) {
        zip = cmp.long_name;
      } else if (cmp.types.includes('country')) {
        country = cmp.short_name;
      }
    });

    // dispatch 'change' event
    dispatch('change', {
      name: place.name,
      address: place.formatted_address,
      streetNumber: stNo || '',
      street: st || '',
      city: city || '',
      state: state || '',
      zip: zip || '',
      country: country || '',
      latitude: marker.position.lat().toFixed(8), // place.geometry.location.lat()
      longitude: marker.position.lng().toFixed(8) // place.geometry.location.lng()
    });
  }
</script>
<div id='pac-card' class='pac-card float-left'>
  <div id='pac-container' class='p-2'>
    <input id='map-inp' type='text' placeholder='search a location...' class='input input-ghost'/>
  </div>
</div>

<div id='map'></div>

<div id='map-inf-win'>
  <span id='place-name' class='title'></span><br/>
  <span id='place-address'></span>
</div>

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
