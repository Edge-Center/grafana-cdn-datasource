import {
  DataSourceInstanceSettings,
  CoreApp,
  ScopedVars,
  MetricFindValue
} from '@grafana/data';
import { DataSourceWithBackend } from '@grafana/runtime';
import {
  Query,
  DataSourceOptions,
  DEFAULT_QUERY,
  ResourcesResponse,
  MetricsResponse,
  RegionsResponse,
  GroupsResponse,
  GranularityResponse,
  VariableQuery
} from './types';
import { VariableSupport } from "./services/VariableSupport";
import { parseNumbers, parseStrings } from "./utils";

export class DataSource extends DataSourceWithBackend<Query, DataSourceOptions> {
  public variableSupport: VariableSupport;

  constructor(instanceSettings: DataSourceInstanceSettings<DataSourceOptions>) {
    super(instanceSettings);

    this.variableSupport = new VariableSupport(this)
  }

  public metricFindQuery(query: VariableQuery): Promise<Array<MetricFindValue>> {
    return this.variableSupport.query(query)
  }

  public getDefaultQuery(_: CoreApp): Partial<Query> {
    return DEFAULT_QUERY;
  }

  public applyTemplateVariables(query: Query, scopedVars: ScopedVars): Query {
    const q = {
      ...query,
      groupby: [].concat(query.groupby),
      vhosts: parseStrings(query.vhostsStr, scopedVars, "csv"),
      clients: parseNumbers(query.clientsStr, scopedVars, "csv"),
      regions: parseStrings(query.regionsStr, scopedVars, "csv"),
      countries: parseStrings(query.countriesStr, scopedVars, "csv"),
      resources: parseNumbers(query.resourcesStr, scopedVars, "csv"),
      metrics: parseStrings(query.metricsStr, scopedVars, "csv"),
    }
    console.log(q);
    return q;
  }

  public getAvailableResources(): Promise<ResourcesResponse> {
    return this.getResource('resources');
  }

  public getAvailableMetrics(): Promise<MetricsResponse> {
    return this.getResource('metrics');
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
}