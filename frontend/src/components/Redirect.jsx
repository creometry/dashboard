import axios from 'axios'
import React, { useEffect, useState } from 'react'
import { useSearchParams, useNavigate } from 'react-router-dom'
import { useCookies } from 'react-cookie';
import useStore from '../zustand/state';
import { checkPayment } from '../payment/payment';
import { getUserData } from '../github/github';


export const Redirect = () => {
    const [searchParams] = useSearchParams()
    const [loading, setLoading] = useState(true)
    const [onlyLogin, setOnlyLogin] = useState(false)
    const [cookies, setCookie] = useCookies(["access_token"]);
    const { user, setUser } = useStore();
    const navigate = useNavigate()
    useEffect(() => {
        const token = searchParams.get('payment_token')
        const onlyLogin = searchParams.get('onlyLogin')
        if (onlyLogin !== null && onlyLogin === "true") {
            setOnlyLogin(true)
            localStorage.setItem('plan', "not_specified")
        } else {

            if (!token) {
                alert("Could not validate payment, token not provided")
                localStorage.clear()
                navigate('/steps')
                return
            }
            // check if access token is present
            if (!cookies.access_token) {
                alert("Could not validate payment")
                localStorage.clear()
                navigate('/steps')
            }

            const payment = async () => {
                const resp = await checkPayment(token, process.env.REACT_APP_TOKEN)
                console.log(resp)
                if (resp !== "Success") {
                    localStorage.removeItem('user_data')
                    navigate('/paymenterror')
                    return
                }
            }
            payment()
        }

        // get user data from github access_token and set it in localStorage
        const resp = getUserData(cookies.access_token)
        resp.then(res => {
            const user = res.user
            if (user) {
                setUser(
                    {
                        id: user.id,
                        login: user.login,
                        name: user.name,
                        access_token: cookies.access_token,
                        avatar_url: user.avatar_url,
                        email: user.email,
                    }
                )
            } else {
                console.log(res.error)
            }
        })
        setLoading(false)

        // eslint-disable-next-line
        // create project (this should return the namespace name and the kubeconfig) then redirect user to dashboard

        // redirect to dashboard after 3 seconds
        setTimeout(() => {
            navigate('/')
        }
            , 3000)
    }, [])
    return (
        <div className='flex flex-col justify-center items-center h-screen'>
            {loading === false &&
                <div>
                    {onlyLogin === false && <div className='text-3xl text-green-500 font-bold'>
                        Successful Payment! <span className='ml-1'>Creating your project...</span>
                    </div>}
                    {user.access_token !== "" &&

                        <div className='flex items-center'>
                            <span>
                                {user.name} You will be redirected to the dashboard in a few seconds
                            </span>
                            <span>
                                <svg width='60px' height='60px' xmlns="http://www.w3.org/2000/svg" viewBox="0 0 100 100" preserveAspectRatio="xMidYMid" className="uil-ring"><rect x="0" y="0" width="100" height="100" fill="none" className="bk"></rect><defs><filter id="uil-ring-shadow" x="-100%" y="-100%" width="300%" height="300%"><feOffset result="offOut" in="SourceGraphic" dx="0" dy="0"></feOffset><feGaussianBlur result="blurOut" in="offOut" stdDeviation="0"></feGaussianBlur><feBlend in="SourceGraphic" in2="blurOut" mode="normal"></feBlend></filter></defs><path d="M10,50c0,0,0,0.5,0.1,1.4c0,0.5,0.1,1,0.2,1.7c0,0.3,0.1,0.7,0.1,1.1c0.1,0.4,0.1,0.8,0.2,1.2c0.2,0.8,0.3,1.8,0.5,2.8 c0.3,1,0.6,2.1,0.9,3.2c0.3,1.1,0.9,2.3,1.4,3.5c0.5,1.2,1.2,2.4,1.8,3.7c0.3,0.6,0.8,1.2,1.2,1.9c0.4,0.6,0.8,1.3,1.3,1.9 c1,1.2,1.9,2.6,3.1,3.7c2.2,2.5,5,4.7,7.9,6.7c3,2,6.5,3.4,10.1,4.6c3.6,1.1,7.5,1.5,11.2,1.6c4-0.1,7.7-0.6,11.3-1.6 c3.6-1.2,7-2.6,10-4.6c3-2,5.8-4.2,7.9-6.7c1.2-1.2,2.1-2.5,3.1-3.7c0.5-0.6,0.9-1.3,1.3-1.9c0.4-0.6,0.8-1.3,1.2-1.9 c0.6-1.3,1.3-2.5,1.8-3.7c0.5-1.2,1-2.4,1.4-3.5c0.3-1.1,0.6-2.2,0.9-3.2c0.2-1,0.4-1.9,0.5-2.8c0.1-0.4,0.1-0.8,0.2-1.2 c0-0.4,0.1-0.7,0.1-1.1c0.1-0.7,0.1-1.2,0.2-1.7C90,50.5,90,50,90,50s0,0.5,0,1.4c0,0.5,0,1,0,1.7c0,0.3,0,0.7,0,1.1 c0,0.4-0.1,0.8-0.1,1.2c-0.1,0.9-0.2,1.8-0.4,2.8c-0.2,1-0.5,2.1-0.7,3.3c-0.3,1.2-0.8,2.4-1.2,3.7c-0.2,0.7-0.5,1.3-0.8,1.9 c-0.3,0.7-0.6,1.3-0.9,2c-0.3,0.7-0.7,1.3-1.1,2c-0.4,0.7-0.7,1.4-1.2,2c-1,1.3-1.9,2.7-3.1,4c-2.2,2.7-5,5-8.1,7.1 c-0.8,0.5-1.6,1-2.4,1.5c-0.8,0.5-1.7,0.9-2.6,1.3L66,87.7l-1.4,0.5c-0.9,0.3-1.8,0.7-2.8,1c-3.8,1.1-7.9,1.7-11.8,1.8L47,90.8 c-1,0-2-0.2-3-0.3l-1.5-0.2l-0.7-0.1L41.1,90c-1-0.3-1.9-0.5-2.9-0.7c-0.9-0.3-1.9-0.7-2.8-1L34,87.7l-1.3-0.6 c-0.9-0.4-1.8-0.8-2.6-1.3c-0.8-0.5-1.6-1-2.4-1.5c-3.1-2.1-5.9-4.5-8.1-7.1c-1.2-1.2-2.1-2.7-3.1-4c-0.5-0.6-0.8-1.4-1.2-2 c-0.4-0.7-0.8-1.3-1.1-2c-0.3-0.7-0.6-1.3-0.9-2c-0.3-0.7-0.6-1.3-0.8-1.9c-0.4-1.3-0.9-2.5-1.2-3.7c-0.3-1.2-0.5-2.3-0.7-3.3 c-0.2-1-0.3-2-0.4-2.8c-0.1-0.4-0.1-0.8-0.1-1.2c0-0.4,0-0.7,0-1.1c0-0.7,0-1.2,0-1.7C10,50.5,10,50,10,50z" fill="#337ab7" filter="url(#uil-ring-shadow)"><animateTransform attributeName="transform" type="rotate" from="0 50 50" to="360 50 50" repeatCount="indefinite" dur="1s"></animateTransform></path></svg>
                            </span>
                        </div>

                    }
                </div>
            }
        </div>
    )
}
