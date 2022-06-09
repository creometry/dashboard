import React, { useEffect, useState } from 'react'
import axios from 'axios'
import { useNavigate } from 'react-router-dom'

export const ServiceList = () => {
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
        getResourceData('services');
    }, [])
    return (
        <div className='p-8'>
            <div className='text-2xl text-gray-400 mb-4'>Services</div>
            <div className='bg-zinc-600 rounded-sm shadow-md'>
                <table className='min-w-full'>
                    <thead className='border-b border-gray-700 bg-zinc-500'>
                        <th className="text-sm font-bold text-gray-900 px-6 py-4 text-left">Name</th>
                        <th className="text-sm font-bold text-gray-900 px-6 py-4 text-left">Namespace</th>
                        <th className="text-sm font-bold text-gray-900 px-6 py-4 text-left">Type</th>
                        <th className="text-sm font-bold text-gray-900 px-6 py-4 text-left">Cluster IP</th>
                        <th className="text-sm font-bold text-gray-900 px-6 py-4 text-left">Ports</th>
                        <th className="text-sm font-bold text-gray-900 px-6 py-4 text-left">External IP</th>
                        <th className="text-sm font-bold text-gray-900 px-6 py-4 text-left">Selector</th>
                        <th className="text-sm font-bold text-gray-900 px-6 py-4 text-left">Status</th>

                    </thead>
                    <tbody>
                        {data.map(el => (
                            <tr className='border-b border-gray-700'>
                                <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>{el.metadata.name}</td>
                                <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>{el.metadata.namespace}</td>
                                <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>{el.spec.type}</td>
                                <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>
                                    {el.spec.clusterIP}
                                </td>
                                <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>
                                    {el.spec.ports.map(p => {
                                        return `${p.port}/${p.protocol}`
                                    })}
                                </td>
                                <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>
                                    {el.spec?.externalIP ? el.spec.externalIP : '-'}
                                </td>
                                <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>
                                    {`${Object.keys(el.spec.selector)}=${el.spec.selector[Object.keys(el.spec.selector)]}`}
                                </td>
                                <td className='text-sm text-green-600 font-bold px-6 py-4 whitespace-nowrap'>
                                    Active
                                </td>
                            </tr>
                        ))}

                    </tbody>
                </table>

            </div>
        </div>

    )
}
