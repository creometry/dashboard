import React, { useEffect, useState } from 'react'
import axios from 'axios'
import { useNavigate } from 'react-router-dom'

export const EventList = () => {
    const { REACT_APP_URL, REACT_APP_NAMESPACE } = process.env;
    const [data, setData] = useState([]);
    const navigate = useNavigate()
    useEffect(() => {
        const getResourceData = async (resource) => {
            try {
                const resp = await axios.get(
                    `${REACT_APP_URL}/api/v1/${resource}/${REACT_APP_NAMESPACE}`
                );
                setData(resp.data.data);
            } catch (error) {
                navigate("/")
            }
        };
        getResourceData('events');
    }, [])
    return (
        <div className='p-8'>
            <div className='text-2xl text-gray-400 mb-4'>Events</div>
            <div className='bg-zinc-600 rounded-sm shadow-md overflow-y-scroll  overflow-x-scroll h-4/5'>
                <table className='min-w-full'>
                    <thead className='border-b border-gray-700 bg-zinc-500'>
                        <th className="text-sm font-bold text-gray-900 px-6 py-4 text-left">Name</th>
                        <th className="text-sm font-bold text-gray-900 px-6 py-4 text-left">Type</th>
                        <th className="text-sm font-bold text-gray-900 px-6 py-4 text-left">Involved Object</th>
                        <th className="text-sm font-bold text-gray-900 px-6 py-4 text-left">Message</th>
                        <th className="text-sm font-bold text-gray-900 px-6 py-4 text-left">Source</th>
                        <th className="text-sm font-bold text-gray-900 px-6 py-4 text-left">Count</th>
                    </thead>
                    <tbody>
                        {data.map(el => (
                            <tr className='border-b border-gray-700'>
                                <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>{el.metadata.name}</td>
                                <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>{el.type}</td>
                                <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>{`${el.involvedObject.kind}: ${el.involvedObject.name}`}</td>
                                <td className='text-sm text-gray-900 font-light px-6 py-4'>
                                    {el.message}
                                </td>
                                <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>{el.source.component}</td>
                                <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>{el.count}</td>
                            </tr>
                        ))}

                    </tbody>
                </table>

            </div>
        </div>
    )
}
