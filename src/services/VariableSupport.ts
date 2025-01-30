import { MetricFindValue } from '@grafana/data';
import { DataSource } from '../datasource';
import { Variable, VariableQuery } from '../types';
import { createMetricFindValues } from './utils';

export class VariableSupport {
  constructor(private datasource: DataSource) {}

  public async query(query: VariableQuery): Promise<MetricFindValue[]> {
    if (!query.selector) {
      return [];
    }

    const selector = query.selector.value!;

    switch (selector) {
      case Variable.Client:
        const { clients } = await this.datasource.getAvailableClients();
        return createMetricFindValues(clients);
      case Variable.Host:
        const { hosts } = await this.datasource.getAvailableHosts();
        return createMetricFindValues(hosts);
      case Variable.Resource:
        const { resources } = await this.datasource.getAvailableResources();
        return resources.map(({id, cname}) =>
            ({
              text: `${id} (${cname})`,
              value: id
            }));
      case Variable.Region:
        const { regions } = await this.datasource.getAvailableRegions();
        return createMetricFindValues(regions);
      case Variable.Country:
        const { countries } = await this.datasource.getAvailableCountries();
        return createMetricFindValues(countries);
    }

    return [];
  }
}
