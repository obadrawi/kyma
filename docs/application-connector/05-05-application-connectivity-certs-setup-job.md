---
title: Application Connectivity Certs Setup Job
type: Configuration
---

To configure the Application Connectivity Certs Setup Job, override the default values of its `values.yaml` file. This document describes parameters that you can configure.

>**TIP:** To learn more about how to use overrides in Kyma, see the following documents: 
>* [Helm overrides for Kyma installation](/root/kyma/#configuration-helm-overrides-for-kyma-installation)
>* [Sub-charts overrides](/root/kyma/#configuration-helm-overrides-for-kyma-installation-sub-chart-overrides)


## Configurable parameters

This table lists the configurable parameters, their descriptions, and default values:

| Parameter | Description | Default value |
|-----------|-------------|---------------|
| **global.applicationConnectorCaKey** | Specifies the base64-encoded private key for the Application Connector. If you don't provide it, a private key is generated automatically. | `` |
| **global.applicationConnectorCa** | Specifies the base64-encoded certificate for the Application Connector. If you don't provide it, the certificate is generated automatically. | `` |
| **application_connectivity_certs_setup_job.certificate.validityTime** | Specifies how long the generated certificate is valid. | `92d` | 

