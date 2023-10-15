export default function mapMount(initMap: () => void) {
  return () => {
    let needScript = true;
    document
      .querySelectorAll('script')
      .forEach((script) => {
        if (script.src.includes('maps.googleapis.com')) {
          needScript = false;
        }
      });

    if (needScript) {
      window.initMap = initMap;
      appendScript(`https://maps.googleapis.com/maps/api/js?key=${import.meta.env.VITE_GAPI_KEY}&callback=initMap&libraries=places,drawing`);
    } else {
      initMap();
    }
  }
}


// appendScript to head
function appendScript(src: string) {
  const ele = document.createElement('script');
  ele.setAttribute('type', 'text/javascript');
  ele.setAttribute('src', src);
  document.head.appendChild(ele);
}
