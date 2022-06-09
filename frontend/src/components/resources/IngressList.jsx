import React, { useEffect, useState } from 'react'
import axios from 'axios'
import { useNavigate } from 'react-router-dom'

export const IngressList = () => {
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
        getResourceData('ingresses');
    }, [])
    return (
        <div className='p-8'>
            <div className='text-2xl text-gray-400 mb-4'>Ingresses</div>
            <div className='bg-zinc-600 rounded-sm shadow-md'>
                <table className='min-w-full'>
                    <thead className='border-b border-gray-700 bg-zinc-500'>
                        <th className="text-sm font-bold text-gray-900 px-6 py-4 text-left">Name</th>
                        <th className="text-sm font-bold text-gray-900 px-6 py-4 text-left">Namespace</th>
                        <th className="text-sm font-bold text-gray-900 px-6 py-4 text-left">Load Balancers</th>
                        <th className="text-sm font-bold text-gray-900 px-6 py-4 text-left">Rules</th>
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
        </div>

    )
}
