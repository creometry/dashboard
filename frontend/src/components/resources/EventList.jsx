import React from 'react'

export const EventList = ({ data }) => {
    return (
        <div className='bg-gray-500 rounded-sm shadow-md'>
            <table className='min-w-full bg-gray-500'>
                <thead className='border-b border-gray-700'>
                    <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Name</th>
                    <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Type</th>
                    <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Involved Object</th>
                    <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Message</th>
                    <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Source</th>
                    <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Count</th>
                </thead>
                <tbody>
                    {data.map(el => (
                        <tr className='border-b border-gray-700'>
                            <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>{el.metadata.name}</td>
                            <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>{el.type}</td>
                            <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>{`${el.involvedObject.kind}: ${el.involvedObject.name}`}</td>
                            <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>
                                {el.message}
                            </td>
                            <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>{el.source.component}</td>
                            <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>{el.count}</td>
                        </tr>
                    ))}

                </tbody>
            </table>

        </div>
    )
}
