import React, { useState } from 'react'
import { useOutsideAlerter } from '../hooks/useOutsideAlerter'
import { useAddMemberPopup } from '../zustand/state';
import axios from "axios";

export const AddMemberPopup = () => {
    const { setIsPopupOpen } = useAddMemberPopup();
    const { ref } = useOutsideAlerter(setIsPopupOpen)
    const { REACT_APP_GET_TEAM_MEMBERS_URL } = process.env
    const [username, setUsername] = useState('')
    const [loading, setLoading] = useState(false)

    const handleSubmit = async (e) => {
        e.preventDefault()
        setLoading(true)
        try {
            const resp = await axios.post(`${REACT_APP_GET_TEAM_MEMBERS_URL}${localStorage.getItem("project_id")}/${username}`)
            console.log(resp.data)
            if (resp.data && resp.data.data.type !== "error") {
                console.log("Success")
                window.location.reload()
            } else {
                console.log("Error")
                setIsPopupOpen(false)
            }
        } catch (err) {
            console.log(err)
            setIsPopupOpen(false)
        }
    }

    return (
        <div className="min-h-screen  fixed  left-0 top-0  flex justify-center items-center inset-0 z-50 outline-none focus:outline-none">
            <div className="absolute bg-black opacity-60 inset-0 z-0"></div>
            <div
                ref={ref}
                className="relative bg-white rounded-lg shadow-lg z-50 p-5">
                {/** create a form with  a username input*/}
                <label htmlFor="username" className="text-lg text-left">Username</label>
                <form className='flex flex-col items-center space-y-5 mt-3' onSubmit={handleSubmit}>
                    <input type="text" id="username" className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline focus:border-creo"
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}
                        required />
                    <button type="submit" className={`${loading === true ? "bg-gray-600" : "bg-creo"} text-white py-1 px-2 text-bold rounded-md cursor-pointer text-lg`} disabled={loading}>{loading === true ? "Loading..." : "Add member"}</button>
                </form>

            </div>
        </div>
    )
}


