import React from 'react'

export const Table = ({ data, resource }) => {
    return (
        resource !== "ingresses" ?
            <div>{
                <div className='bg-gray-500 rounded-sm shadow-md'>
                    <table className='min-w-full'>
                        <thead className='border-b border-gray-700'>
                            <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Name</th>
                            <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Namespace</th>
                            <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Containers</th>
                            <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Status</th>
                        </thead>
                        <tbody>
                            {data.map(el => (
                                <tr className='border-b border-gray-700'>
                                    <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>{el.metadata.name}</td>
                                    <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>{el.metadata.namespace}</td>
                                    <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>{el.spec?.containers?.length}</td>
                                    <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>
                                        {el.status.phase === 'Running' ?
                                            <span className='bg-green-500 px-2 rounded-md py-1'>{el.status.phase}</span>
                                            :
                                            <span className='text-red-500'>{el.status.phase}</span>
                                        }

                                    </td>
                                </tr>
                            ))}

                        </tbody>
                    </table>

                </div>
            }</div>
            :
            <div>{
                data.items.map(el => {
                    return (
                        <div>
                            <h1>{el.metadata.name}</h1>
                        </div>
                    )
                })
            }</div>
    )
}
