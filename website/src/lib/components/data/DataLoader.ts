/*
 *  Copyright (c) 2022 - 2022.  Ankit Patial.
 *  Unauthorized copying of this file, via any medium is strictly prohibited. Proprietary and confidential.
 *  Author: Ankit Patial
 */

import { writable } from 'svelte/store';

export interface Data {
  count: number;
  items: Array<any>;
  isLoadMore: boolean;
  hasPrevious: boolean;
  hasNext: boolean;
  startCursor: undefined | string;
  endCursor: undefined | string;
  error: undefined | string;
}

export interface Params {
  [key: string]: any;
}

export default class DataLoader {
  private after: string | undefined;
  private first: number | undefined;
  private before: string | undefined;
  private last: number | undefined;

  private readonly client: any;
  private readonly qry: any;
  private readonly limit: number;
  private readonly dataFieldName: string;
  private readonly data: Data;
  private readonly params: Params;

  public where: object;
  public isAsc: boolean;
  public orderField: string;

  public store: any;
  public loading: any;

  constructor(client: any, qry: any, dataFieldName: string, limit: number) {
    this.client = client;
    this.qry = qry;
    this.limit = limit;
    this.first = limit;
    this.dataFieldName = dataFieldName;
    this.where = {};
    this.isAsc = true;
    this.orderField = 'CREATED';

    this.params = {};
    this.data = {
      isLoadMore: false,
      count: 0,
      items: [],
      hasPrevious: false,
      hasNext: false,
      startCursor: undefined,
      endCursor: undefined,
      error: undefined
    };

    // store
    this.store = writable(this.data);
    this.loading = writable(false);
  }

  async load(resetPage = true) {
    this.loading.set(true);

    const p = {
      ...this.params,
      page: (
        resetPage
          ? { first: this.limit }
          : {
            after: this.after,
            first: this.first,
            before: this.before,
            last: this.last
          }
      ),
      where: this.where,
      orderBy: { direction: this.isAsc ? 'ASC' : 'DESC', field: this.orderField }
    };

    const r = await this.client.query(this.qry, p).toPromise();

    this.loading.set(false);

    if (r.err) {
      this.data.error = r.err.message;
      return;
    }

    const { totalCount, edges, pageInfo } = (r.data || {})[this.dataFieldName] || {};

    this.data.count = totalCount || 0;
    this.data.startCursor = pageInfo?.startCursor;
    this.data.endCursor = pageInfo?.endCursor;
    this.data.hasPrevious = pageInfo?.hasPreviousPage;
    this.data.hasNext = pageInfo?.hasNextPage;

    if (this.data.isLoadMore) {
      this.data.isLoadMore = false;
      this.data.items = this.data.items.concat((edges || []).map((o: any) => o.node));
    } else {
      this.data.items = (edges || []).map((o: any) => o.node);
    }

    this.store.set(this.data);
  }

  nextPage = () => {
    this.after = this.data.endCursor;
    this.first = this.limit;
    this.before = undefined;
    this.last = undefined;
    return this.load(false);
  };

  previousPage = () => {
    this.after = undefined;
    this.first = undefined;
    this.before = this.data.startCursor;
    this.last = this.limit;
    return this.load(false);
  };

  loadMore = () => {
    this.data.isLoadMore = true;

    this.after = this.data.endCursor;
    this.first = this.limit;
    this.before = undefined;
    this.last = undefined;

    return this.load(false);
  };

  setParam = (name: string, value: any) => {
    this.params[name] = value;
  };
}
