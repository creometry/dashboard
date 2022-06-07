import React from 'react'

export const ServiceList = ({ data }) => {
    return (
        <div className='bg-gray-500 rounded-sm shadow-md'>
            <table className='min-w-full'>
                <thead className='border-b border-gray-700'>
                    <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Name</th>
                    <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Namespace</th>
                    <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Type</th>
                    <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Cluster IP</th>
                    <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Ports</th>
                    <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">External IP</th>
                    <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Selector</th>
                    <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Status</th>

                </thead>
                <tbody>
                    {data.map(el => (
                        <tr className='border-b border-gray-700'>
                            <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>{el.metadata.name}</td>
                            <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>{el.metadata.namespace}</td>
                            <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>{el.spec.type}</td>
                            <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>
                                {el.spec.clusterIP}
                            </td>
                            <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>
                                {el.spec.ports.map(p => {
                                    return `${p.port}/${p.protocol}`
                                })}
                            </td>
                            <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>
                                {el.spec?.externalIP ? el.spec.externalIP : '-'}
                            </td>
                            <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>
                                {`${Object.keys(el.spec.selector)}=${el.spec.selector[Object.keys(el.spec.selector)]}`}
                            </td>
                            <td className='text-sm text-green-600 font-bold px-6 py-4 whitespace-nowrap'>
                                Active
                            </td>
                        </tr>
                    ))}

                </tbody>
            </table>

        </div>
    )
}
