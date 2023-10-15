<script context="module">
  const elements = new Set();
</script>

<script lang='ts'>
  import {onMount} from "svelte";

  export let poster = 'poster';
  export let src = '';

  let video,
    currentTime = 0, // current time
    duration = 0,  // video duration
    paused = true,
    volume = 0.3, // 0-1 volume level
    muted = false;

  onMount(() => {
    elements.add(video);
    return () => elements.delete(video);
  })

  function stopOthers() {
    elements.forEach(element => {
      if (element !== video) {
        element.pause();
        // element.currentTime = 0;
      }
    });
  }
</script>

<video
  controls
  {poster}
  bind:this={video}
  bind:currentTime
  bind:duration
  bind:paused
  bind:volume
  bind:muted
  on:play={stopOthers}
  class="w-full"
  preload="none"
  controlsList="nodownload"
  oncontextmenu="return false;"
>
  <source src={src}>
  <p>
    Your browser doesn't support HTML video. Here is a
    <a href={src}>link to the video</a> instead.
  </p>
</video>
