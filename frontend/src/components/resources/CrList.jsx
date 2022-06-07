import React from 'react'

export const CrList = ({ data }) => {
    console.log(data)
    return (
        <div className='bg-gray-500 rounded-sm shadow-md'>
            <table className='min-w-full'>
                <thead className='border-b border-gray-700'>
                    <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Name</th>
                    <th className="text-sm font-medium text-gray-900 px-6 py-4 text-left">Group</th>
                </thead>
                <tbody>
                    {data.map(el => (
                        <tr className='border-b border-gray-700'>
                            <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>{el.kind}</td>
                            <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>{el.apiVersion}</td>
                        </tr>
                    ))}

                </tbody>
            </table>

        </div>
    )
}
