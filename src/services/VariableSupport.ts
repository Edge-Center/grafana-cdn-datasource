import { MetricFindValue } from '@grafana/data';
import { DataSource } from "../datasource";
import { Variable, VariableQuery } from "../types";
import { createMetricFindValues } from "./utils";


export class VariableSupport {
    constructor(private datasource: DataSource) {}

    public async query(query: VariableQuery): Promise<Array<MetricFindValue>> {
        if (!query.selector) {
            return [];
        }

        const selector = query.selector.value!;

        switch (selector) {
            case Variable.Client:
            case Variable.Resource:
            case Variable.Vhost:
                const {resources} = await this.datasource.getAvailableResources();

                switch (selector) {
                    case Variable.Vhost:
                        return createMetricFindValues(resources.map((item) => item.cname));
                    case Variable.Resource:
                        return createMetricFindValues(resources.map((item) => item.id));
                    case Variable.Client:
                        return createMetricFindValues(resources.map((item) => item.client));
                }
                break;
            case Variable.Region:
                const {regions} = await this.datasource.getAvailableRegions();
                return createMetricFindValues(regions);
        }

        return [];
    }
}
