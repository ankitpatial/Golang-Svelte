import {browser} from '$app/environment';
import client from "$lib/gql/client";
import {QryMe} from "$lib/gql/account";

export const ssr = false;
export const prerender = false;

export async function load() {
  if (!browser) {
    return {authUser: undefined};
  }

  // in browser
  const res = await client().query(QryMe, {}).toPromise()
  return {
    authUser: res.data?.me
  }
}
