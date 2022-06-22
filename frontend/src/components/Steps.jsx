import React, { useState, useEffect } from 'react'
import { useSearchParams, useNavigate } from 'react-router-dom'
import LoginGithub from 'react-login-github';
import axios from 'axios';
import { useCookies } from 'react-cookie';
import useStore from '../zustand/state';



export const Steps = () => {
    const [searchParams] = useSearchParams()
    const plan = searchParams.get('plan')
    const [cookies, setCookie] = useCookies(["access_token"]);
    const { user, setUser } = useStore();
    const navigate = useNavigate()
    const {
        REACT_APP_GITHUB_CLIENT_ID,
        REACT_APP_URL,
        REACT_APP_VENDOR,
        REACT_APP_TOKEN,
        REACT_APP_SUCCESS_URL,
        REACT_APP_ERROR_URL,
        REACT_APP_CREATE_PAYMENT_URL,
        REACT_APP_CREATE_PAYMENT_GATEWAY } = process.env
    const [step, setStep] = useState(1)
    const [projectName, setProjectName] = useState(localStorage.getItem('projectName') || '')
    const [repoName, setRepoName] = useState(localStorage.getItem('repoName') || '')
    const [repoUrl, setRepoUrl] = useState(localStorage.getItem('repoUrl') || '')
    const [repoBranch, setRepoBranch] = useState(localStorage.getItem('repoBranch') || '')
    const [namespace, setNamespace] = useState(localStorage.getItem('namespace') || '')

    const handleSubmit = (e) => {
        e.preventDefault()
        localStorage.setItem('projectName', projectName)
        localStorage.setItem('repoName', repoName)
        localStorage.setItem('repoUrl', repoUrl)
        localStorage.setItem('repoBranch', repoBranch)
        localStorage.setItem('plan', plan)
        localStorage.setItem('namespace', namespace)
        setStep(2)
    }

    const exchange = async (code) => {
        try {
            const resp = await axios.get(`${REACT_APP_URL}/api/v1/github/exchange/${code}`
            )
            console.log("resp : ", resp.data)
            const access_token = resp.data.access_token
            console.log("access_token: ", access_token)
            // set access_token in cookie
            setCookie("access_token", access_token, { path: "/" })
        } catch (err) {
            console.log("get acces_token err : " + err)
        }
    }

    const onSuccess = async (response) => {
        // extract code from response
        const code = response.code;
        console.log("code: ", code)
        // exchange code for github access_token
        exchange(code)
        // calculate price based on plan
        const price = plan === 'Starter' ? 49 : plan === 'Pro' ? 99 : plan === 'Elite' ? 199 : 0

        try {
            const resp = await fetch(
                REACT_APP_CREATE_PAYMENT_URL,
                {
                    method: "POST",
                    body: JSON.stringify({ vendor: REACT_APP_VENDOR, amount: price, note: "test" }),
                    headers: {
                        "Content-Type": "application/json",
                        Authorization: `Token ${REACT_APP_TOKEN}`,
                    },
                }
            );
            // get the response as json
            const data = await resp.json();

            if (data.message !== "Success") {
                navigate('/paymenterror')
            }

            const token = data.data.token;

            // submit a form with the token as a hidden field
            const form = document.createElement("form");
            form.action = REACT_APP_CREATE_PAYMENT_GATEWAY;
            form.method = "POST";
            form.innerHTML = `<input type="hidden" name="payment_token" value="${token}">`;
            form.innerHTML += `<input type="hidden" name="url_ok" value="${REACT_APP_SUCCESS_URL}">`;
            form.innerHTML += `<input type="hidden" name="url_ko" value="${REACT_APP_ERROR_URL}">`;

            document.body.appendChild(form);
            form.submit();
        } catch (err) {
            console.log(err)
            navigate('/paymenterror')
        }

    };
    const onFailure = response => {
        alert('Failed to login with GitHub')
        navigate('/steps?plan=' + plan)
    };

    useEffect(() => {
        if (user.access_token !== "") {
            // the user is logged in so we should redirect to "/"
            navigate('/')
            return
        }
        if ((!plan) || (plan !== "Starter" && plan !== "Pro" && plan !== "Elite")) {
            const code = searchParams.get('code')
            if (code) { console.log("closing...") }
            else {
                alert('Invalid plan, redirecting to starter plan')
                localStorage.clear()
                navigate('/steps?plan=Starter')
            }
        }
    }, [])

    return (
        <div
            className='flex flex-col justify-center items-center h-screen bg-gray-100'
        >
            {step === 1 && <div className='flex flex-col items-center'>

                <div className='text-2xl font-bold text-gray-700'>
                    Fill this form to continue
                </div>
                <form className="bg-white shadow-md rounded-lg px-8 pt-6 pb-8  mt-4 flex flex-col items-center w-96" onSubmit={handleSubmit}>
                    <div className="mb-4 w-full">
                        <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="projectName">
                            Project Name
                        </label>
                        <input className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline focus:border-creo
                    " id="projectName" type="text" placeholder="Project Name" required
                            value={projectName}
                            onChange={(e) => setProjectName(e.target.value)}
                        />
                    </div>
                    <div className="mb-4 w-full">
                        <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="namespace">
                            Namespace
                        </label>
                        <input className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline focus:border-creo
                    " id="namespace" type="text" placeholder="Namespace" required
                            value={namespace}
                            onChange={(e) => setNamespace(e.target.value)}
                        />
                    </div>
                    <div className="mb-4 w-full">
                        <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="RepositoryName">
                            Repository Name
                        </label>
                        <input className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline focus:border-creo" id="RepositoryName" type="text" placeholder="Repository Name" required
                            value={repoName}
                            onChange={(e) => setRepoName(e.target.value)}
                        />
                    </div>
                    <div className="mb-2 w-full">
                        <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="RepositoryURL">
                            Repository URL
                        </label>
                        <input className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline focus:border-creo" id="RepositoryURL" type="text" placeholder="Repository URL" required
                            value={repoUrl}
                            onChange={(e) => setRepoUrl(e.target.value)}
                        />
                    </div>

                    <div className="mb-4 w-full">
                        <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="Branch">
                            Branch
                        </label>
                        <input className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline focus:border-creo" id="Branch" type="text" placeholder="Branch" required
                            value={repoBranch}
                            onChange={(e) => setRepoBranch(e.target.value)}
                        />
                    </div>

                    <button type="submit" className='py-2 px-6 border rounded-md bg-creo text-white'>Next</button>

                </form>
            </div>}
            {step === 2 &&

                <LoginGithub
                    clientId={REACT_APP_GITHUB_CLIENT_ID}
                    onSuccess={onSuccess}
                    onFailure={onFailure}
                    className="flex items-center bg-gray-900 rounded-md py-3 px-2 cursor-pointer hover:bg-gray-700 text-gray-100"
                    buttonText="Login with GitHub"
                    scopes="user"
                />

            }

        </div>
    )
}
