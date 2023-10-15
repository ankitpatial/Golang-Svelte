<!--
  ref: https://github.com/iandevlin/iandevlin.github.io/tree/master/mdn/video-player
-->

<script context="module">
  const elements = new Set();
</script>

<script lang='ts'>
  import {onMount} from "svelte";
  import IconPlay from "$lib/assets/svg/IconPlay.svelte";
  import IconPause from "$lib/assets/svg/IconPause.svelte";
  import Btn from "$lib/components/form/Btn.svelte";
  import IconSpeakerMute from "$lib/assets/svg/IconSpeakerMute.svelte";
  import IconSpeaker from "$lib/assets/svg/IconSpeaker.svelte";
  import IconFullScreenUndo from "$lib/assets/svg/IconFullScreenUndo.svelte";
  import IconFullScreen from "$lib/assets/svg/IconFullScreen.svelte";

  export let poster = 'poster';
  export let src = '';

  let videoContainer, video, fullScreenEnabled,
    fullScreen = false,
    currentTime = 0, // current time
    duration = 0,  // video duration
    paused = true,
    volume = 0.3, // 0-1 volume level
    muted = false;

  onMount(() => {
    // Hide the default controls
    video.controls = false;
    // Check if the browser supports the Fullscreen API
    fullScreenEnabled = !!(document.fullscreenEnabled || document.mozFullScreenEnabled || document.msFullscreenEnabled
      || document.webkitSupportsFullscreen || document.webkitFullscreenEnabled
      || document.createElement('video').webkitRequestFullScreen
    );

    elements.add(video);
    return () => elements.delete(video);
  })

  function isFullScreen() {
    return !!(document.fullScreen || document.webkitIsFullScreen || document.mozFullScreen
      || document.msFullscreenElement || document.fullscreenElement
    );
  }

  function stopOthers() {
    elements.forEach(element => {
      if (element !== video) element.pause();
    });
  }

  // Set the video container's fullscreen state
  function fullscreenchange() {
    fullScreen = isFullScreen();
  }

  function handleFullscreen() {
    // If fullscreen mode is active...
    if (isFullScreen()) {
      // ...exit fullscreen mode
      // (Note: this can only be called on document)
      if (document.exitFullscreen) document.exitFullscreen();
      else if (document.mozCancelFullScreen) document.mozCancelFullScreen();
      else if (document.webkitCancelFullScreen) document.webkitCancelFullScreen();
      else if (document.msExitFullscreen) document.msExitFullscreen();
    } else {
      // ...otherwise enter fullscreen mode
      // (Note: can be called on document, but here the specific element is used as it will also ensure that the
      // element's children, e.g. the custom controls, go fullscreen also)
      if (videoContainer.requestFullscreen) videoContainer.requestFullscreen();
      else if (videoContainer.mozRequestFullScreen) videoContainer.mozRequestFullScreen();
      else if (videoContainer.webkitRequestFullScreen) {
        // Safari 5.1 only allows proper fullscreen on the video element. This also works fine on other WebKit browsers
        // as the following CSS (set in styles.css) hides the default controls that appear again, and
        // ensures that our custom controls are visible:
        // figure[data-fullscreen=true] video::-webkit-media-controls { display:none !important; }
        // figure[data-fullscreen=true] .controls { z-index:2147483647; }
        video.webkitRequestFullScreen();
      } else if (videoContainer.msRequestFullscreen) videoContainer.msRequestFullscreen();
    }
  }

  // React to the user clicking within the progress bar
  function handleProgress(e) {
    const pos = (e.pageX - this.offsetLeft) / this.offsetWidth;
    currentTime = pos * duration;
  }

  function playPause() {
    paused = !paused;
  }

  function mute() {
    muted = !muted;
  }
</script>

<svelte:window on:fullscreenchange={fullscreenchange}/>

<div bind:this={videoContainer} class="flex flex-col">
  <div>
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
    >
      <source src={src}>
      <p>
        Your browser doesn't support HTML video. Here is a
        <a href={src}>link to the video</a> instead.
      </p>
    </video>
  </div>
  <div>
    {#if video}
      <div class="flex flex-row absolute ml mr-1 -mt-8 bg-transparent/20 w-full space-x-1">
        <div class="w-10 ml-1">
          <Btn small circle transparent on:click={playPause}>
            {#if paused }
              <IconPlay/>
            {:else }
              <IconPause/>
            {/if}
          </Btn>
        </div>
        {#if (currentTime)}
          <div class="w-full">
            <progress class="progress" value={currentTime} max={duration} on:click={handleProgress}>
            </progress>
          </div>
          <div class="w-10">
            <Btn small circle transparent on:click={mute}>
              {#if muted }
                <IconSpeakerMute color="red"/>
              {:else }
                <IconSpeaker/>
              {/if}
            </Btn>
          </div>
          <div class="w-10" class:hidden={!fullScreenEnabled}>
            <Btn small circle transparent on:click={handleFullscreen}>
              {#if fullScreen }
                <IconFullScreenUndo/>
              {:else }
                <IconFullScreen/>
              {/if}
            </Btn>
          </div>
        {/if}
      </div>
    {/if}
  </div>
</div>
