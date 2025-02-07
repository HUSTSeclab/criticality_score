// This file is auto-generated by @hey-api/openapi-ts

import type { Options as ClientOptions, TDataShape, Client } from '@hey-api/client-fetch';
import type { GetHistoriesData, GetHistoriesResponse, GetResultsData, GetResultsResponse, GetResultsByScoreidData, GetResultsByScoreidResponse } from './types.gen';
import { client as _heyApiClient } from './client.gen';

export type Options<TData extends TDataShape = TDataShape, ThrowOnError extends boolean = boolean> = ClientOptions<TData, ThrowOnError> & {
    /**
     * You can provide a client instance returned by `createClient()` instead of
     * individual options. This might be also useful if you want to implement a
     * custom client.
     */
    client?: Client;
};

/**
 * Get score histories
 * Get score histories by git link
 */
export const getHistories = <ThrowOnError extends boolean = false>(options: Options<GetHistoriesData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).get<GetHistoriesResponse, unknown, ThrowOnError>({
        url: '/histories',
        ...options
    });
};

/**
 * Search score results by git link
 * Search score results by git link
 * NOTE: All details are ignored, should use /results/:scoreid to get details
 * NOTE: Maxium take count is 1000
 */
export const getResults = <ThrowOnError extends boolean = false>(options: Options<GetResultsData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).get<GetResultsResponse, unknown, ThrowOnError>({
        url: '/results',
        ...options
    });
};

/**
 * Get score results
 * Get score results, including all details by scoreid
 */
export const getResultsByScoreid = <ThrowOnError extends boolean = false>(options: Options<GetResultsByScoreidData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).get<GetResultsByScoreidResponse, unknown, ThrowOnError>({
        url: '/results/{scoreid}',
        ...options
    });
};