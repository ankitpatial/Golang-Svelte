/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

import {writable} from 'svelte/store';
import sleep from '$lib/utils/sleep';
import type {SocketMsg} from '$lib/models/SocketMsg';
import {Channel} from "$lib/graph/graphql";

export const socket = createSocket();

export const wsMessage = wsMsg();

function wsMsg() {
  const s = writable<SocketMsg | undefined>(undefined);
  let tOut: any;
  return {
    subscribe: s.subscribe,
    set: (d: SocketMsg) => {
      s.set(d)

      if (tOut) clearTimeout(tOut)
      tOut = setTimeout(() => {
        s.set(undefined)
      }, 1000)
    },
    // clear: () => s.set(undefined)
  };
}

let client: WebSocket | undefined;

function createSocket() {
  let attempt = 0,
    connecting = false,
    host: string,
    token: string,
    pingTimeOut: any;

  function ping() {
    // debounce ping for a second
    if (pingTimeOut) clearTimeout(pingTimeOut);
    pingTimeOut = setTimeout(() => {
      send(Channel.Ping, {}).catch(console.error);
    }, 800);
  }

  function attemptConnection() {
    if (connecting) return;

    connecting = true;
    client = undefined;
    sleep(2000)
      .then(() => {
        attempt++;
        connect(host, token, true);
      })
      .catch((ex) => {
        console.log('ws: failed to connect', ex);
      })
      .finally(() => {
        connecting = false;
      });
  }

  function setMsg(msg: string) {
    try {
      const obj = typeof msg === 'object' ? msg : JSON.parse(msg);
      wsMessage.set(obj);
    } catch (err) {
      console.error('ws: failed to parse msg', err, msg);
    }
  }

  function connect(wsHost: string, t: string, force = false) {
    if (!force && client) {
      return;
    }

    if (!t) {
      return;
    }

    host = wsHost;
    token = t
    client = newWebSocket(`${host}?token=${token}`);
    if (!client) return;

    client.onopen = () => {
      attempt = 0;
      console.log('ws: connected');
    };

    // on socket message
    client.onmessage = (evt: any) => {
      const msg = evt?.data;
      if (!msg) {
        return;
      }

      console.log('ws: msg', msg)
      setMsg(msg);
    };

    // on error
    client.onerror = (event: any) => {
      console.log('ws: error', event);
    };

    // on close
    client.onclose = (event: any) => {
      console.log(`ws: connection closed, code=${event.code} reason=${event.reason}`);
      attemptConnection();
    };
  }

  async function send(channel: Channel, data: object) {
    let sent = false;
    attempt = 0;

    while (!sent && attempt < 30) {
      attempt++;
      if (!client || connecting) {
        // will attempt in a second
        // console.log("ws: failed to send as socket connection is not ready");
        await sleep(1000);
        continue;
      }

      try {
        // web version
        await client.send(JSON.stringify({channel, data}));
        sent = true;
      } catch (ex: any) {
        // console.log("ws: failed to send message, error", ex);
        if (ex.toString().includes('connection not found')) {
          client = undefined;
          attemptConnection();
          continue;
        }
        await sleep(2);
      }
    }
  }

  return {
    connect,
    send,
    close: (reason: string) => {
      if (!client) return;
      host = '';
      token = '';
      client.close(3000, reason);
    },
    ping
  };
}

function newWebSocket(host: string) {
  if (typeof WebSocket !== 'undefined') {
    return new WebSocket(host);
  }

  // eslint-disable-next-line @typescript-eslint/ban-ts-comment
  // @ts-ignore
  if (typeof MozWebSocket !== 'undefined') {
    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    // @ts-ignore
    return new MozWebSocket(host);
  }

  if (typeof window !== 'undefined') {
    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    // @ts-ignore
    const WS = window.WebSocket || window.MozWebSocket;
    if (WS) {
      return new WS(host);
    }
  }

  return undefined;
}

