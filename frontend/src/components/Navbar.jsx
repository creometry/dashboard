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
            <div className="flex items-center">
                <div>
                    <button
                        type="button"
                        className="bg-zinc-700 p-1 rounded-full text-gray-400 hover:text-white focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-gray-800 focus:ring-white"
                    >
                        <span className="sr-only">View notifications</span>
                        <svg
                            className="h-6 w-6"
                            xmlns="http://www.w3.org/2000/svg"
                            fill="none"
                            viewBox="0 0 24 24"
                            strokeWidth="2"
                            stroke="currentColor"
                            aria-hidden="true"
                        >
                            <path
                                strokeLinecap="round"
                                strokeLinejoin="round"
                                d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"
                            />
                        </svg>
                    </button>
                </div>

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
