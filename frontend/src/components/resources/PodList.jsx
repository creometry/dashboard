import React, { useEffect, useState } from 'react'
import axios from 'axios'
import { useNavigate } from 'react-router-dom'

export const PodList = () => {
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
        getResourceData('pods');
    }, [])
    return (
        <div className='p-8'>
            <div className='text-2xl text-gray-400 mb-4'>Pods</div>
            <div className='bg-zinc-600 rounded-sm shadow-md'>
                <table className='min-w-full'>
                    <thead className='border-b border-gray-700 bg-zinc-500'>
                        <th className="text-sm font-bold text-gray-900 px-6 py-4 text-left">Name</th>
                        <th className="text-sm font-bold text-gray-900 px-6 py-4 text-left">Namespace</th>
                        <th className="text-sm font-bold text-gray-900 px-6 py-4 text-left">Containers</th>
                        <th className="text-sm font-bold text-gray-900 px-6 py-4 text-left">Status</th>
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
                                        <span className='bg-red-500 px-2 rounded-md py-1'>{el.status.phase}</span>
                                    }

                                </td>
                            </tr>
                        ))}

                    </tbody>
                </table>

            </div>
        </div>
    )
}
