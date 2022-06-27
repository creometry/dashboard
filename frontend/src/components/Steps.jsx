import React, { useState, useEffect } from 'react'
import { useSearchParams, useNavigate } from 'react-router-dom'
import LoginGithub from 'react-login-github';
import { useCookies } from 'react-cookie';
import useStore from '../zustand/state';
import { createPayment } from '../payment/payment';
import { exchange } from '../github/github';



export const Steps = ({ onlyLogin }) => {
    const [searchParams] = useSearchParams()
    const plan = searchParams.get('plan')
    const [cookies, setCookie] = useCookies(["access_token"]);
    const { user } = useStore();
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
    const [error, setError] = useState("")

    const handleSubmit = (e) => {
        e.preventDefault()
        localStorage.setItem('projectName', projectName)
        localStorage.setItem('repoName', repoName)
        localStorage.setItem('repoUrl', repoUrl)
        localStorage.setItem('repoBranch', repoBranch)
        localStorage.setItem('plan', plan)
        setStep(2)
    }

    const handleProjectNameChange = (value) => {
        // the project name should only contain lowercase letters and -, and should not start with a -

        const projectNameRegex = /^[a-z-]+$/
        if (!projectNameRegex.test(value) && value) {
            setError("Project name should only contain lowercase letters and -, and should not start with a -")
        } else {
            setError("")
        }
        setProjectName(value)
    }

    const onSuccess = async (response) => {
        // extract code from response
        const code = response.code;
        console.log("code: ", code)
        // exchange code for github access_token
        const data = await exchange(code, REACT_APP_URL)
        if (data.error) {
            console.log("error: ", data.error)
            return
        } else {
            console.log("data: ", data)
            setCookie("access_token", data.access_token, { path: "/" })
        }

        // calculate price based on plan
        if (onlyLogin !== true) {
            const price = plan === 'Starter' ? 49 : plan === 'Pro' ? 99 : plan === 'Elite' ? 199 : 0

            // create payment
            const data = await createPayment(REACT_APP_CREATE_PAYMENT_URL, REACT_APP_VENDOR, price, REACT_APP_TOKEN)

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

        } else {
            navigate("/success?onlyLogin=true")
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
        if (onlyLogin !== true) {
            if ((!plan) || (plan !== "Starter" && plan !== "Pro" && plan !== "Elite")) {
                const code = searchParams.get('code')
                if (code) { console.log("closing...") }
                else {
                    alert('Invalid plan, redirecting to starter plan')
                    localStorage.clear()
                    navigate('/steps?plan=Starter')
                }
            }
        }
    }, [])

    return (
        <div
            className='flex flex-col justify-center items-center h-screen bg-gray-100'
        >
            {(step === 1 && onlyLogin === false) && <div className='flex flex-col items-center'>

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
                            onChange={(e) => handleProjectNameChange(e.target.value)}
                        />
                    </div>
                    <div className='text-red-500'>{error}</div>
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

                    <button type="submit" className='py-2 px-6 border rounded-md bg-creo text-white' disabled={error !== ""}>Next</button>

                </form>
                <a
                    href='/login'
                    className='underline cursor-pointer text-lg mt-3 font-bold text-gray-700'>
                    Or login with github and choose a plan and pay for it later
                </a>
            </div>}
            {(step === 2 || onlyLogin) === true &&

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
