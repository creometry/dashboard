import React, { useEffect } from 'react'
import { useCookies } from 'react-cookie';
import useStore from '../zustand/state';
import axios from 'axios';

export const Navbar = () => {
    const { user, setUser } = useStore();
    const [cookies, setCookie, removeCookie] = useCookies(["access_token", "rancher_token"]);
    const { REACT_APP_GET_KUBECONFIG_URL } = process.env
    const handleLogout = () => {
        setUser({
            id: "",
            login: "",
            name: "",
            access_token: "",
            avatar_url: "",
            email: "",
        })
        localStorage.clear()
        removeCookie("access_token", { path: "/" })
        removeCookie("rancher_token", { path: "/" })
        // redirect to /login
        window.location.href = "/login"
    }
    const handleKubeconfig = async () => {
        try {

            const resp = await axios.post(REACT_APP_GET_KUBECONFIG_URL, {
                token: cookies.rancher_token
            })
            const { data } = resp
            if (data) {
                console.log(data.config)
                navigator.clipboard.writeText(data.config)
            } else {
                console.log("error")
            }
        } catch (e) {
            console.log(e)
        }
    }

    useEffect(() => {
        if (cookies.access_token === undefined) {
            alert("err")
            handleLogout()
        }
    }, [])
    return (
        <div className="bg-zinc-800 h-12 flex justify-between items-center  px-12">
            <div className='text-gray-200'>
                Hello {user.name}
            </div>
            <div className="bg-creo text-gray-200 px-3 py-1 rounded-md cursor-pointer font-bold">
                Namespace : {localStorage.getItem('namespace') || "Not Set"}
            </div>
            <div className='text-gray-200 font-semibold underline cursor-pointer' onClick={() => handleKubeconfig()}>
                Copy kubeconfig
            </div>
            <div className='text-gray-200 font-semibold underline cursor-pointer' onClick={() => handleLogout()}>
                Logout
            </div>
        </div>
    )
}
