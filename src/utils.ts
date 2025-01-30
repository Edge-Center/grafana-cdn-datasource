import { getTemplateSrv } from '@grafana/runtime';
import { ScopedVars } from '@grafana/data';

export const parseNumbers = (target?: string, scopedVars?: ScopedVars, format?: string | Function): number[] => {
  return parseStrings(target, scopedVars, format).map((x) => +x);
};

export const parseStrings = (target?: string, scopedVars?: ScopedVars, format?: string | Function): string[] => {
  return getTemplateSrv()
    .replace(target || '', scopedVars, format)
    .split(',')
    .filter(Boolean);
};
