/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */


import {gql} from '@urql/svelte';

export const QryPricing = gql`
  query GetPricing{
    getPricing {
      id
      items {
        id
        country
        state
        stateAbr
        city
        zip
        productId
        price
        pricePer
      }
      products {
        id
        name
      }

    }
  }
`;
