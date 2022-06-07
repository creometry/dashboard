import React from 'react'

export const SecretList = ({ data }) => {
    return (
        <div className='bg-gray-500 rounded-sm shadow-md'>
            <table className='min-w-full'>
                <thead className='border-b border-gray-700'>
                    <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Name</th>
                    <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Namespace</th>
                    <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Labels</th>
                    <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Keys</th>
                    <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Type</th>
                </thead>
                <tbody>
                    {data.map(el => (
                        <tr className='border-b border-gray-700'>
                            <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>{el.metadata.name}</td>
                            <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>{el.metadata.namespace}</td>
                            <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'></td>
                            <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>
                                {Object.keys(el.data).map(key => {
                                    return (
                                        <span className='ml-1 text-gray-900 font-light'>{key}</span>
                                    )
                                })}
                            </td>
                            <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>{el.type}</td>
                        </tr>
                    ))}

                </tbody>
            </table>

        </div>
    )
}
