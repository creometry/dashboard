import React from 'react'
import { PodList } from './resources/PodList'
import { DeploymentList } from './resources/DeploymentList'
import { JobList } from './resources/JobList'
import { CronjobList } from './resources/CronjobList'
import { ConfigmapList } from './resources/ConfigmapList'
import { SecretList } from './resources/SecretList'
import { EventList } from './resources/EventList'
import { ServiceList } from './resources/ServiceList'
import { IngressList } from './resources/IngressList'
import { EndpointList } from './resources/EndpointList'
import { PvcList } from './resources/PvcList'
import { HpaList } from './resources/HpaList'
import { CrList } from './resources/CrList'

export const Table = ({ data, resource }) => {
    if (resource === 'pods') {
        return <PodList data={data} />
    }
    if (resource === 'deployments') {
        return <DeploymentList data={data} isSts={false} />
    }

    if (resource === 'sts') {
        return <DeploymentList data={data} isSts={true} />
    }

    if (resource === 'jobs') {
        return <JobList data={data} />
    }

    if (resource === 'cronjobs') {
        return <CronjobList data={data} />
    }

    if (resource === 'configmaps') {
        return <ConfigmapList data={data} />
    }

    if (resource === 'secrets') {
        return <SecretList data={data} />
    }

    if (resource === 'events') {
        return <EventList data={data} />
    }

    if (resource === 'services') {
        return <ServiceList data={data} />
    }

    if (resource === 'ingresses') {
        return <IngressList data={data} />
    }

    if (resource === 'endpoints') {
        return <EndpointList data={data} />
    }

    if (resource === 'pvcs') {
        return <PvcList data={data} />
    }

    if (resource === 'horizontalpodautoscalers') {
        return <HpaList data={data} />
    }

    if (resource === 'customresources') {
        return <CrList data={data} />
    }

}
