import {Strings, StringsResponse} from '../../types';

export const resolveLabel = (item: string, strings?: Strings): string => {
  let label = item;
  if (strings && strings[item]) {
    return strings[item].label;
  }

  return label;
};

export const resolveDesc = (item: string, strings?: Strings): string => {
  if (strings && strings[item]) {
    return strings[item].desc;
  }

  return '';
};


export const getMetricStrings = (strings: StringsResponse | undefined): Strings => {
  if (!strings) {
    return {};
  }
  return {...strings.metrics, ...strings.pluginMetrics};
}
