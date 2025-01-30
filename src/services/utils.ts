import { MetricFindValue } from '@grafana/data';

export const createMetricFindValues = (target: Array<string | number>): MetricFindValue[] => {
  return target.filter(unique).map((value) => ({ text: `${value}`, value: value }));
};

const unique = <T>(v: T, idx: number, a: T[]) => a.indexOf(v) === idx;
