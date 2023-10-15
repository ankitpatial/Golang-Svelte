<div class="fixed w-full h-0.5 z-10 text-center">
  <div class="progress progress-striped">
    <div class="progress-bar">
    </div>
  </div>
</div>

<style>
  .progress {
    background: #EFECEC10;
  }

  .progress-bar {
    height: 3px;
    transition: 0.4s linear;
    transition-property: width, background-color;
  }

  .progress-striped .progress-bar {
    background-color: #4B99FF;
    width: 100%;
    animation: progressAnimationStrike 10s;
  }

  @keyframes progressAnimationStrike {
    from {
      width: 0
    }
    to {
      width: 100%
    }
  }

  @keyframes progressAnimation {
    0% {
      width: 5%;
    }
    100% {
      width: 85%;
    }
  }
</style>
