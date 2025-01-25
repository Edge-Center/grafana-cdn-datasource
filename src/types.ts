import {DataSourceJsonData, SelectableValue} from '@grafana/data';
import { DataQuery } from '@grafana/schema';

export interface Query extends DataQuery {
  queryType: 'table' | 'timeSeries'

  metricsStr: string
  regionsStr: string
  vhostsStr: string
  resourcesStr: string
  clientsStr: string
  countriesStr: string

  metrics: Array<string>
  regions: Array<string>
  vhosts: Array<string>
  countries: Array<string>
  resources: Array<number>
  clients: Array<number>

  groupby: Array<string>
  granularity: string
  legendFormat: string
}

export const DEFAULT_QUERY: Partial<Query> = {
  queryType: 'timeSeries',
  metricsStr: 'total_bytes',
  groupby: ['resource'],
  granularity: '1h'
};

export enum Variable {
  Resource = "resource",
  Client = "client",
  Vhost = "vhost",
  Region = "region",
}

export interface VariableQuery {
  selector: SelectableValue<Variable>;
}

export const defaultVariableQuery: Partial<VariableQuery> = {
  selector: { value: Variable.Resource, label: "resourceID" },
};

export interface GCVariableQuery {

}

export interface DataSourceOptions extends DataSourceJsonData {
  apiUrl?: string;
}

export interface SecureJsonData {
  apiKey?: string;
}

export type ResourcesResponse = {
  resources: Array<{
    id: number;
    cname: string;
    client: number;
  }>;
}

export type MetricsResponse = {
  metrics: Array<string>;
}

export type RegionsResponse = {
  regions: Array<string>;
}

export type GroupsResponse = {
  groups: Array<string>;
}

export type GranularityResponse = {
  granularity: Array<string>;
}
