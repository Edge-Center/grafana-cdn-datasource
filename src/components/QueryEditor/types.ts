import {QueryEditorProps} from "@grafana/data";
import {DataSource} from "../../datasource";
import {DataSourceOptions, Query} from "../../types";

export  type EditorProps = QueryEditorProps<DataSource, Query, DataSourceOptions>;

export type ChangeOptions<T> = {
    propertyName: keyof T;
    runQuery: boolean;
};