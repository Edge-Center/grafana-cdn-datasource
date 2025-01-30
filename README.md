# Grafana Data Source Plugin: EdgeCenter CDN Datasource

## Overview

The **CDN Datasource** is a backend plugin for Grafana that provides real-time and historical CDN performance data from EdgeCenter, including traffic, latency, cache hit ratio, and other key performance indicators.

## Features

- **Visualize data**: Display data from the CDN stats API.
- **Filters**: Use resource and client IDs to filter data.
- **Group data**: Group data for better organization.
- **Variables**: Use variables for flexible queries.

## What is Grafana Data Source Backend Plugin?
A Backend plugin is a type of data-source plugin that runs on the server.
You will need access to the Edgecenter API from the Grafana server to properly interact with the plugin.

## Installation

### Prerequisites

- Grafana v7.0 or later.
- EdgeCenter API credentials.

## Getting Started

To begin using this plugin, download it and place it in the `grafana/plugins` directory.

Note: This plugin is currently unsigned, so Grafana will block it from loading by default. To enable it, you need to explicitly allow unsigned plugins. If you're using Grafana with containers, you can do this by adding the following environment variable:

```bash
-e "GF_PLUGINS_ALLOW_LOADING_UNSIGNED_PLUGINS=edgecenter-cdn-datasource"
```

This will allow Grafana to load the plugin and start using it.

## Configuration

To configure the data source, update the Grafana `datasources.yaml` file:

```yaml
apiVersion: 1

datasources:
  - name: 'cdn-datasource'
    type: 'edgecenter-cdn-datasource'
    access: proxy
    isDefault: false
    orgId: 1
    version: 1
    editable: true
    jsonData:
      apiUrl: 'https://api.edgecenter.ru'
    secureJsonData:
      apiKey: '####'
```

Alternatively, use the Grafana API to add the data source:

```sh
curl -XPOST -i http://admin:$ADMIN_PASSWORD@localhost:3000/api/datasources \
     --data-binary '{"name": "cdn-datasource","type": "edgecenter-cdn-datasource", \
     "orgId": 1,"access":"proxy", "jsonData":{"apiUrl": "https://api.edgecenter.ru"}, \
     "secureJsonData":{"apiKey": "####"}}' -H "Content-Type: application/json"
```

## Usage

- Create a new dashboard in Grafana.
- Add a panel and select **CDN Datasource** as the data source.
- Configure the query parameters to retrieve specific CDN metrics.
- Apply visualizations and filters as needed.