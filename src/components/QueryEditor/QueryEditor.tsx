import React, { FC } from 'react';
import {Label, Stack} from '@grafana/ui';
import { EditorProps } from "./types";
import { useChangeString } from "./useChangeString";
import { useChangeSelectableValue } from "./useChangeSelectableValue";
import { useSelectableValue } from "./useSelectableValue";
import { ResourcesPicker } from "./fields/ResourcesPicker";
import { VHostsPicker } from "./fields/VHostsPicker";
import { RegionsPicker } from "./fields/RegionsPicker";
import { MetricPicker } from "./fields/MetricPicker";
import { GranularityPicker } from "./fields/GranularityPicker";
import { GroupByPicker } from "./fields/GroupByPicker";
import { LegendFormatInput } from "./fields/LegendFormatInput";
import { ClientsPicker } from "./fields/ClientsPicker";
import {useChangeSelectableValues} from "./useChangeSelectableValues";
import {useSelectableValues} from "./useSelectableValues";
import {CountriesPicker} from "./fields/CountriesPicker";


export const QueryEditor: FC<EditorProps> = (props) => {
    const { datasource, query } = props;

    const groupBy = useSelectableValues(query.groupby);
    const metrics = useSelectableValue(query.metricsStr);
    const granularity = useSelectableValue(query.granularity);
    const legendFormat = query.legendFormat;
    const resources = query.resourcesStr;
    const clients = query.clientsStr;
    const countries = query.countriesStr;
    const vhosts = query.vhostsStr;
    const regions = query.regionsStr;


    const onChangeGroupBy = useChangeSelectableValues(props, {
        propertyName: 'groupby',
        runQuery: true,
    });

    const onChangeMetrics = useChangeSelectableValue(props, {
        propertyName: 'metricsStr',
        runQuery: true,
    });

    const onChangeGranularity = useChangeSelectableValue(props, {
        propertyName: 'granularity',
        runQuery: true,
    });

    const onChangeResources = useChangeString(props, {
        propertyName: 'resourcesStr',
        runQuery: true,
    });

    const onChangeClients = useChangeString(props, {
        propertyName: 'clientsStr',
        runQuery: true,
    });

    const onChangeCountries = useChangeString(props, {
        propertyName: 'countriesStr',
        runQuery: true,
    });

    const onChangeRegions = useChangeString(props, {
        propertyName: 'regionsStr',
        runQuery: true,
    });

    const onChangeVhosts = useChangeString(props, {
        propertyName: 'vhostsStr',
        runQuery: true,
    });

    const onChangeLegendFormat = useChangeString(props, {
        propertyName: 'legendFormat',
        runQuery: true,
    });


    return (
        <Stack direction={'column'} gap={2}>
            <Stack direction={'row'} gap={2}>
                <Stack direction={'column'} gap={1}>
                    <Label>Filters (comma separated)</Label>
                    <ResourcesPicker value={resources} onChange={onChangeResources} />
                    <VHostsPicker value={vhosts} onChange={onChangeVhosts} />
                    <RegionsPicker value={regions} onChange={onChangeRegions} />
                    <ClientsPicker value={clients} onChange={onChangeClients} />
                    <CountriesPicker value={countries} onChange={onChangeCountries} />
                </Stack>

                <Stack direction={'column'} gap={1}>
                    <Label>Query Props</Label>
                    <MetricPicker value={metrics} onChange={onChangeMetrics} datasource={datasource} />
                    <GroupByPicker value={groupBy} onChange={onChangeGroupBy} datasource={datasource} />
                    <GranularityPicker value={granularity} onChange={onChangeGranularity} datasource={datasource} />
                </Stack>
            </Stack>

            <LegendFormatInput value={legendFormat} onChange={onChangeLegendFormat} />
        </Stack>
    );
};
