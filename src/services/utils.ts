import {MetricFindValue} from "@grafana/data";

export const createMetricFindValues = (
    target: Array<string | number>
): Array<MetricFindValue> => {
    return target.filter(unique).map((value) => ({ text: `${value}`, value: value }));
};

const unique = <T>(v: T, idx: number, a: Array<T>) => a.indexOf(v) === idx;

