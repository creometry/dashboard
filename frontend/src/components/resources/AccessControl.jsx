import React, { useState, useEffect } from 'react'
import axios from 'axios'
import { useAddMemberPopup } from '../../zustand/state';


export const AccessControl = () => {
    const [data, setData] = useState([]);
    const { setIsPopupOpen } = useAddMemberPopup();
    const { REACT_APP_GET_TEAM_MEMBERS_URL } = process.env
    useEffect(() => {
        const getAccessControlData = async () => {
            try {
                const resp = await axios.get(`${REACT_APP_GET_TEAM_MEMBERS_URL}${localStorage.getItem('project_id') || ""}`);
                if (resp.data.members !== null) {
                    setData(resp.data.members)
                }
            } catch (err) {
                console.log(err)
            }
        }
        getAccessControlData()
    }, [])
    return (
        <div className='p-8'>
            <div className='flex items-center justify-between'>
                <div className='text-2xl text-gray-400 mb-4'>Access control</div>
                <div className='text-gray-100 font-bold bg-creo py-2 px-1 rounded-md cursor-pointer'
                    onClick={() => setIsPopupOpen(true)}
                >Add a team member</div>
            </div>
            <div className='bg-zinc-600 rounded-sm shadow-md overflow-y-scroll  overflow-x-scroll h-4/5'>
                <table className='min-w-full'>
                    <thead className='border-b border-gray-700 bg-zinc-500'>
                        <th className="text-sm font-bold text-gray-900 px-6 py-4 text-left">Username</th>
                        <th className="text-sm font-bold text-gray-900 px-6 py-4 text-left">Id</th>
                        <th className="text-sm font-bold text-gray-900 px-6 py-4 text-left">Role</th>
                    </thead>
                    <tbody>
                        {data.map(el => (
                            <tr className='border-b border-gray-700' key={el.id}>
                                <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>{el.name || el.username}</td>
                                <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>{el.id}</td>
                                <td className='text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap'>{el.type}</td>
                            </tr>
                        ))}

                    </tbody>
                </table>

            </div>
        </div>
    )
}
