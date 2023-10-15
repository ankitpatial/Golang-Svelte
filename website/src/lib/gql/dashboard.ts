/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */


import {gql} from '@urql/svelte';

export const QryOverview = gql`
  query Overview($filter: Filter!){
    overview(f: $filter) {
      id
      total
      items {
        id
        name
        count
      }
    }
  }
`
