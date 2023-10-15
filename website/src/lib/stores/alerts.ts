import {writable} from 'svelte/store';
import type Alert from '$lib/models/Alert';

function createStore() {
  const {subscribe, update, set} = writable<[Alert]>();
  const QueueMaxLen = 5;

  let list: [Alert] | undefined;

  subscribe((arr) => {
    list = arr;
  });

  function remove(id: number) {
    if (!Array.isArray(list)) return;

    const idx = list.findIndex((p) => p.id == id);
    if (idx == -1) return;

    list.splice(idx, 1);
    set(list);
  }

  function add(msg: Alert, clearInSeconds = 5) {
    msg.id = Date.now();
    update((p: Array<Alert>): any => {
      const arr = (p || []).concat(msg);
      arr.sort(sort);

      // remove for queue if queue is grown beyond maxlength
      const over = arr.length - QueueMaxLen;
      if (over > 0) {
        arr.splice(0, over).forEach(m => {
          remove(m.id || 0)
        });
      }
      return arr;
    });

    setTimeout(() => {
      remove(msg.id || 0);
    }, 1000 * clearInSeconds);
  }

  function sort(a: Alert, b: Alert) {
    if (Number(a.id) < Number(b.id)) {
      return -1
    }
    if (Number(a.id) > Number(b.id)) {
      return 1
    }
    return 0
  }

  return {
    subscribe,
    push: add,
    info: (title: string, message: string, clearInSeconds = 5) => {
      add({type: 'info', title, body: message}, clearInSeconds);
    },
    success: (title: string, message: string, clearInSeconds = 5) => {
      add({type: 'success', title, body: message}, clearInSeconds);
    },
    warning: (title: string, message: string, clearInSeconds = 5) => {
      add({type: 'warning', title, body: message}, clearInSeconds);
    },
    error: (title: string, message: string, clearInSeconds = 5) => {
      add({type: 'error', title, body: message}, clearInSeconds);
    },
    remove
  };
}

export default createStore();

export const modalAlert = createStore();
