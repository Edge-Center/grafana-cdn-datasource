import { DataSourceJsonData, SelectableValue } from '@grafana/data';
import { DataQuery } from '@grafana/schema';

export interface Query extends DataQuery {
  queryType: 'table' | 'timeSeries';

  regionsStr: string;
  hostsStr: string;
  resourcesStr: string;
  clientsStr: string;
  countriesStr: string;


  metrics: string[];
  regions: string[];
  hosts: string[];
  countries: string[];
  resources: number[];
  clients: number[];

  groupby: string[];
  granularity: string;
  legendFormat: string;
}

export const DEFAULT_QUERY: Partial<Query> = {
  queryType: 'timeSeries',
  metrics: ['total_bytes'],
  groupby: ['resource'],
  granularity: '1h',
};

export enum Variable {
  Resource = 'resource',
  Client = 'client',
  Host = 'host',
  Region = 'region',
  Country = 'country',
}

export interface VariableQuery {
  selector: SelectableValue<Variable>;
}

export interface DataSourceOptions extends DataSourceJsonData {
  apiUrl?: string;
}

export interface SecureJsonData {
  apiKey?: string;
}

export type CdnResourcesResponse = {
  resources: Array<{
    id: number;
    cname: string;
    client: number;
  }>;
};

export type ResourcesResponse = {
  resources: number[];
};

export type ClientsResponse = {
  clients: number[];
};

export type HostsResponse = {
  hosts: string[];
};

export type MetricsResponse = {
  metrics: string[];
};

export type CountriesResponse = {
  countries: string[];
};

export type RegionsResponse = {
  regions: string[];
};

export type GroupsResponse = {
  groups: string[];
};

export type GranularityResponse = {
  granularity: string[];
};

export type StringDetail = {
  label: string;
  desc: string;
};

export type Strings = {
  [key: string]: StringDetail;
};

export type StringsResponse = {
  metrics: Strings;
  pluginMetrics: Strings;
  groupBy: Strings;
  granularity: Strings;
};