import React from 'react'

export const IngressList = ({ data }) => {
    return (
        <div className='bg-gray-500 rounded-sm shadow-md'>
            <table className='min-w-full'>
                <thead className='border-b border-gray-700'>
                    <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Name</th>
                    <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Namespace</th>
                    <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Load Balancers</th>
                    <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Rules</th>
                </thead>
                <tbody>
                    {Array.isArray(data) ? data.map(el => (
                        <tr className='border-b border-gray-700'>
                            <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>{el.metadata.name}</td>
                            <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>{el.metadata.namespace}</td>
                            <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>-</td>
                            <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>
                                {el.spec.rules.map(r => {
                                    return `${r.host}`
                                })}
                            </td>

                        </tr>
                    )) :
                        <tr>
                            <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>{data.metadata.resourceVersion}</td>
                        </tr>
                    }

                </tbody>
            </table>

        </div>
    )
}
