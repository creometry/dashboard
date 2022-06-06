import React, { useEffect } from 'react'

const data = new Map([
    ['Pods', 'pods'],
    ['Deployments', 'deployments'],
    ['Statefulsets', 'sts'],
    ['Jobs', 'jobs'],
    ['Cronjobs', 'cronjobs'],
    ['Services', 'services'],
    ['Ingresses', 'ingresses'],
    ['Secrets', 'secrets'],
    ['Configmaps', 'configmaps'],
    ['Events', 'events'],
    ['Endpoints', 'endpoints'],
    ['Persistent volume claims', 'pvcs'],
    ['Horizontal pod autoscalers', 'horizontalpodautoscalers'],
    ['Logs', 'logs'],
    ['Metrics', 'metrics'],
    ['Monitoring', 'monitoring'],
    ['App store', 'appstore'],
    ['Billing', 'billing'],
]);

export const Resource = ({ resourceName, getResourceData }) => {
    const [resource, setResource] = React.useState("")

    useEffect(() => {
        setResource(data.get(resourceName))
    }, [resourceName])
    return (
        <div
            className="hover:bg-creo p-1 cursor-pointer hover:text-gray-100 rounded-md"
            onClick={() => getResourceData(resource)}
        >
            {resourceName}
        </div>
    )
}
