import React, {FC} from 'react';
import { Alert, InlineField, Input, SecretInput } from '@grafana/ui';
import { useChangeOptions } from "./useChangeOptions";
import { useChangeSecureOptions } from "./useChangeSecureOptions";
import { useResetSecureOptions } from "./useResetSecureOptions";
import { EditorProps } from "./types";

export const ConfigEditor: FC<EditorProps> = (props) => {
  const { jsonData, secureJsonData, secureJsonFields } = props.options;
  const onApiUrlChange = useChangeOptions(props, 'apiUrl');
  const onApiKeyChange = useChangeSecureOptions(props, 'apiKey');
  const onResetApiKey = useResetSecureOptions(props, 'apiKey');

  return (
    <>
      <InlineField label="API URL" labelWidth={14} interactive tooltip="The URL for the API endpoint, e.g., https://api.edgecenter.ru">
        <Input
            id="config-editor-api-url"
            onChange={onApiUrlChange}
            value={jsonData.apiUrl}
            placeholder="Enter the full API URL, e.g., https://api.edgecenter.ru"
            width={40}
        />
      </InlineField>

      <InlineField label="API Key" labelWidth={14} interactive tooltip={'Secure json field (backend only)'}>
        <SecretInput
          required
          id="config-editor-api-key"
          isConfigured={secureJsonFields.apiKey}
          value={secureJsonData?.apiKey}
          placeholder="Enter your API key"
          width={40}
          onReset={onResetApiKey}
          onChange={onApiKeyChange}
        />
      </InlineField>

      <Alert severity={"info"} title="How to create a API token?">
        <a
            href="https://edgecenter.ru/en/knowledge-base/account-management/get-a-permanent-token"
            target="_blank"
            rel="noreferrer"
        >
          https://edgecenter.ru/en/knowledge-base/account-management/get-a-permanent-token
        </a>
      </Alert>
    </>
  );
};
