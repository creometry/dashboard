import axios from 'axios'
import React, { useEffect, useState } from 'react'
import { useSearchParams, useNavigate } from 'react-router-dom'
import GithubLogo from '../../src/GitHub.png'

export const GithubAuth = () => {
    const [searchParams] = useSearchParams()
    const [loading, setLoading] = useState(true)
    const navigate = useNavigate()
    useEffect(() => {
        const token = searchParams.get('payment_token')
        if (!token) {
            localStorage.removeItem('user_data')
            navigate('/plans')
            return
        }
        const checkPayment = async () => {
            // make axios request to check if the payment is valid or not with authorization header
            try {
                const resp = await axios.get(`https://sandbox.paymee.tn/api/v1/payments/${token}/check`, {
                    headers: {
                        Authorization: `Token ${process.env.REACT_APP_TOKEN}`
                    }
                })

                if (resp.data.message !== "Success") {
                    localStorage.removeItem('user_data')
                    navigate('/paymenterror')
                    return
                }

                const userData = {
                    plan: localStorage.getItem('plan'),
                }

                localStorage.setItem('user_data', JSON.stringify(userData))

                setLoading(false)

            } catch (err) {
                localStorage.removeItem('user_data')
                navigate('/paymenterror')
                return
            }


        }
        checkPayment()
        // eslint-disable-next-line
    }, [])
    return (
        <div className='flex flex-col justify-center items-center h-screen'>
            {loading === false && <div className='text-3xl text-green-500 font-bold'>
                Successful Payment!
            </div>}
            {loading === false && <div className='flex items-center bg-gray-900 rounded-md py-3 px-2 cursor-pointer hover:bg-gray-700 mt-4'>
                <div>
                    <img src={GithubLogo} alt="github_logo" height={32} width={32} />
                </div>
                <div className='text-gray-100 ml-3'>Login with GitHub</div>
            </div>}

        </div>
    )
}
