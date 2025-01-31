import React, { FC } from 'react';
import { Stack, useStyles2 } from '@grafana/ui';
import { EditorProps } from './types';
import { useChangeString } from './useChangeString';
import { useChangeSelectableValue } from './useChangeSelectableValue';
import { useSelectableValue } from './useSelectableValue';
import { ResourcesPicker } from './fields/ResourcesPicker';
import { HostsPicker } from './fields/HostsPicker';
import { RegionsPicker } from './fields/RegionsPicker';
import { MetricPicker } from './fields/MetricPicker';
import { GranularityPicker } from './fields/GranularityPicker';
import { GroupByPicker } from './fields/GroupByPicker';
import { LegendFormatInput } from './fields/LegendFormatInput';
import { ClientsPicker } from './fields/ClientsPicker';
import { useChangeSelectableValues } from './useChangeSelectableValues';
import { useSelectableValues } from './useSelectableValues';
import { CountriesPicker } from './fields/CountriesPicker';
import { useStrings } from './useStrings';
import { css } from '@emotion/css';
import {GrafanaTheme2, SelectableValue} from '@grafana/data';
import {getMetricStrings} from "./utils";
import {QueryTypePicker} from "./fields/QueryTypePicker";
import {QueryType} from "../../types";

export const QueryEditor: FC<EditorProps> = (props) => {
  const { datasource, query } = props;

  const strings = useStrings(datasource);
  const groupBy = useSelectableValues(query.groupby, strings.value?.groupBy);
  const metrics = useSelectableValues(query.metrics, getMetricStrings(strings.value));
  const granularity = useSelectableValue(query.granularity, strings.value?.granularity);
  const queryType = useSelectableValue(query.queryType, strings.value?.queryTypes);
  const legendFormat = query.legendFormat;
  const resources = query.resourcesStr;
  const clients = query.clientsStr;
  const countries = query.countriesStr;
  const hosts = query.hostsStr;
  const regions = query.regionsStr;


  const onChangeQueryType = useChangeSelectableValue(props, {
    propertyName: 'queryType',
    runQuery: true,
  });

  const onChangeGroupBy = useChangeSelectableValues(props, {
    propertyName: 'groupby',
    runQuery: true,
  });

  const onChangeMetrics = useChangeSelectableValues(props, {
    propertyName: 'metrics',
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

  const onChangeHosts = useChangeString(props, {
    propertyName: 'hostsStr',
    runQuery: true,
  });

  const onChangeLegendFormat = useChangeString(props, {
    propertyName: 'legendFormat',
    runQuery: true,
  });

  const styles = useStyles2(getStyles);

  return (
    <div className={styles.editor}>
      <Stack direction={'column'} gap={2}>
        <Stack direction={'row'} gap={2}>
          <Stack direction={'column'} gap={1}>
            <ResourcesPicker value={resources} onChange={onChangeResources} />
            <HostsPicker value={hosts} onChange={onChangeHosts} datasource={datasource} />
            <RegionsPicker value={regions} onChange={onChangeRegions} />
            <ClientsPicker value={clients} onChange={onChangeClients} />
            <CountriesPicker value={countries} onChange={onChangeCountries} />
          </Stack>

          <Stack direction={'column'} gap={1}>
            <QueryTypePicker value={queryType} onChange={onChangeQueryType} datasource={datasource} />
            <MetricPicker value={metrics} onChange={onChangeMetrics} datasource={datasource} />
            <GroupByPicker value={groupBy} onChange={onChangeGroupBy} datasource={datasource} />
            {
                shouldRenderGranularityPicker(queryType) &&
              <GranularityPicker value={granularity} onChange={onChangeGranularity} datasource={datasource} />
            }
          </Stack>
        </Stack>

        <LegendFormatInput value={legendFormat} onChange={onChangeLegendFormat} />
      </Stack>
    </div>
  );
};

const shouldRenderGranularityPicker = (queryType: SelectableValue<string> | undefined) => {
  if (!queryType) {
    return true
  }
  return queryType.value === QueryType.TimeSeries;
};

function getStyles(theme: GrafanaTheme2) {
  return {
    editor: css`
      margin: ${theme.spacing(2, 0.5, 0.5, 0)};
    `,
  };
}
