/*
 *  Copyright (c) 2022 - 2023.  Ankit Patial.
 *  Unauthorized copying of this file, via any medium is strictly prohibited. Proprietary and confidential.
 *  Author: Ankit Patial
 */

import { CombinedError, createClient, errorExchange, fetchExchange, type Operation } from '@urql/svelte';
import type { Client } from '@urql/svelte/dist/urql-svelte';
import { authUser } from '$lib/stores/auth';
import { goto } from '$app/navigation';
import { PATH } from '$lib/enum';
import alerts from '$lib/stores/alerts';

export default function(): Client {
  return createClient({
    url: '/query',
    requestPolicy: 'cache-and-network',
    exchanges: [
      errorExchange({
        onError: (error: CombinedError, operation: Operation) => {
          if (error.graphQLErrors) {
            if (error.message.includes('[GraphQL] ent')) {
              error.message = error.message.replace('[GraphQL] ent', 'DB Error');
            } else if (error.message.includes('[GraphQL] ')) {
              error.message = error.message.replace('[GraphQL] ', '');
            }
          }
          if (error.message.includes('401 Unauthorized')) {
            authUser.set(null);
            goto(PATH.LOGIN);
            alerts.push({
              type: 'warning',
              title: 'Session',
              body: 'Session is expired, please login again'
            });
          }
        }
      }),
      fetchExchange
    ]
  });
}
