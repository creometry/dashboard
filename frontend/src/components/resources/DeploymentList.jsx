import React from 'react'

export const DeploymentList = ({ data, isSts }) => {
    return (
        <div className='bg-gray-500 rounded-sm shadow-md'>
            <table className='min-w-full'>
                <thead className='border-b border-gray-700'>
                    <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Name</th>
                    <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Namespace</th>
                    <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Pods</th>
                    <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Replicas</th>
                    {!isSts && <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Conditions</th>}
                </thead>
                <tbody>
                    {data.map(el => (
                        <tr className='border-b border-gray-700'>
                            <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>{el.metadata.name}</td>
                            <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>{el.metadata.namespace}</td>
                            <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>{`${el.status.readyReplicas}/${el.status.availableReplicas}`}</td>
                            <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>
                                {el.spec.replicas}
                            </td>
                            {!isSts && <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>
                                {el.status.conditions.map(c => {
                                    return (
                                        <span className={c.type === 'Available' ? "ml-1 font-bold text-green-600" : "ml-1 font-bold text-blue-700"}>{c.type}</span>
                                    )
                                })}
                            </td>}
                        </tr>
                    ))}

                </tbody>
            </table>

        </div>
    )
}
