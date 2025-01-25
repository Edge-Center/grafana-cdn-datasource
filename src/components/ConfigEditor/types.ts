import { DataSourcePluginOptionsEditorProps } from '@grafana/data';
import { DataSourceOptions, SecureJsonData } from "../../types";

export interface EditorProps extends DataSourcePluginOptionsEditorProps<DataSourceOptions, SecureJsonData> {}