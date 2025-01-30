import { DataSourceInstanceSettings, CoreApp, ScopedVars, MetricFindValue } from '@grafana/data';
import { DataSourceWithBackend } from '@grafana/runtime';
import {
  Query,
  DataSourceOptions,
  DEFAULT_QUERY,
  MetricsResponse,
  RegionsResponse,
  GroupsResponse,
  GranularityResponse,
  VariableQuery,
  ClientsResponse,
  HostsResponse,
  CdnResourcesResponse,
  CountriesResponse,
  StringsResponse,
} from './types';
import { VariableSupport } from './services/VariableSupport';
import { parseNumbers, parseStrings } from './utils';
import { countries } from './data/countries';

export class DataSource extends DataSourceWithBackend<Query, DataSourceOptions> {
  public variableSupport: VariableSupport;

  constructor(instanceSettings: DataSourceInstanceSettings<DataSourceOptions>) {
    super(instanceSettings);

    this.variableSupport = new VariableSupport(this);
  }

  public metricFindQuery(query: VariableQuery): Promise<MetricFindValue[]> {
    return this.variableSupport.query(query);
  }

  public getDefaultQuery(_: CoreApp): Partial<Query> {
    return DEFAULT_QUERY;
  }

  public applyTemplateVariables(query: Query, scopedVars: ScopedVars): Query {
    const q: Query = {
      ...query,
      hosts: parseStrings(query.hostsStr, scopedVars, 'csv'),
      clients: parseNumbers(query.clientsStr, scopedVars, 'csv'),
      regions: parseStrings(query.regionsStr, scopedVars, 'csv'),
      countries: parseStrings(query.countriesStr, scopedVars, 'csv'),
      resources: parseNumbers(query.resourcesStr, scopedVars, 'csv'),
      granularity: query.granularity || '1h',
    };
    console.log(q, scopedVars);
    return q;
  }

  public getAvailableResources(): Promise<CdnResourcesResponse> {
    return this.getResource<CdnResourcesResponse>('resources');
  }

  public getAvailableClients(): Promise<ClientsResponse> {
    return this.getResource<CdnResourcesResponse>('resources').then(({ resources }) => {
      return {
        clients: resources.map(({ client }) => client),
      };
    });
  }

  public getAvailableHosts(): Promise<HostsResponse> {
    return this.getResource<CdnResourcesResponse>('resources').then(({ resources }) => {
      return {
        hosts: resources.map(({ cname }) => cname),
      };
    });
  }

  public getAvailableMetrics(): Promise<MetricsResponse> {
    return this.getResource('metrics');
  }

  public getAvailableCountries(): Promise<CountriesResponse> {
    return Promise.resolve({ countries });
  }

  public getAvailableRegions(): Promise<RegionsResponse> {
    return this.getResource('regions');
  }

  public getAvailableGroups(): Promise<GroupsResponse> {
    return this.getResource('groups');
  }

  public getAvailableGranularity(): Promise<GranularityResponse> {
    return this.getResource('granularity');
  }

  public getStrings(): Promise<StringsResponse> {
    return this.getResource('strings');
  }
}
