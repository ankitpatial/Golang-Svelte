<script>
  import { onDestroy, onMount } from 'svelte';
  import Chart from 'chart.js/auto';

  export let type = 'line';
  export let data = {
    labels: [],
    datasets: [{
      label: '',
      data: [12, 19, 3, 5, 2, 3],
      borderWidth: 1
    }],
    yMarkers: {},
    yRegions: []
  };

  // const colors = ['light-blue', 'blue', 'violet', 'red', 'orange', 'yellow', 'green', 'light-green', 'purple',
  //   'magenta', 'light-grey', 'dark-grey', '#CD5C5C', '#F08080', '#FFA07A', '#FFBF00', '#DFFF00', '#6495ED',
  //   '#D8BFD8', '#DDA0DD', '#DA70D6', '#FF00FF', '#BA55D3', '#9370DB', '#663399', '#9ACD32',
  //   '#808000', '#556B2F', '#66CDAA', '#8FBC8B', '#008B8B', '#40E0D0', '#5F9EA0', '#B0C4DE', '#CD853F', '#D2691E',
  //   '#F5DEB3', '#CC99FF', '#FF99FF', '#FF99CC', '#FF9999', '#FF9966', '#CC6600', '#CC6666', '#CC6699', '#CC66CC',
  //   '#FF6666', '#FF3333', '#FFFF00', '#FF0099'];

  //  Update the chart when incoming data changes

  let chart = null, chartRef;

  onMount(() => {
    chart = new Chart(chartRef, {
      type,
      data: {
        labels: data.labels,
        datasets: data.datasets
      }
    });
  });

  onDestroy(() => {
    chart?.destroy();
  });
</script>

<div>
  <canvas bind:this={chartRef}></canvas>
</div>
