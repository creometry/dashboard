import React from 'react'

export const PvcList = ({ data }) => {
    return (
        <div className='bg-gray-500 rounded-sm shadow-md'>
            <table className='min-w-full'>
                <thead className='border-b border-gray-700'>
                    <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Name</th>
                    <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Storage Class</th>
                    <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Capacity</th>
                    <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Claim</th>
                    <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Status</th>
                </thead>
                <tbody>
                    {data.map(el => (
                        <tr className='border-b border-gray-700'>
                            <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>{el.metadata.name}</td>
                            <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>{el.spec.storageClassName}</td>
                            <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>{el.spec.resources.requests.storage}</td>
                            <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'></td>
                            <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>{el.status.phase}</td>

                        </tr>
                    ))
                    }

                </tbody>
            </table>

        </div>
    )
}
