import React, { useState } from 'react'
import { useOutsideAlerter } from '../hooks/useOutsideAlerter'
import { usePopup } from '../zustand/state';
import { createPayment } from '../payment/payment';
import { useNavigate } from 'react-router-dom';



export const PlansPopup = () => {
    const { setIsOpen } = usePopup();
    const { ref } = useOutsideAlerter(setIsOpen)
    const [step, setStep] = useState(1)
    const [plan, setPlan] = useState(null)
    const navigate = useNavigate()
    const {
        REACT_APP_VENDOR,
        REACT_APP_TOKEN,
        REACT_APP_SUCCESS_URL_2,
        REACT_APP_ERROR_URL,
        REACT_APP_CREATE_PAYMENT_URL,
        REACT_APP_CREATE_PAYMENT_GATEWAY } = process.env
    const [projectName, setProjectName] = useState(localStorage.getItem('projectName') || '')
    const [repoName, setRepoName] = useState(localStorage.getItem('repoName') || '')
    const [repoUrl, setRepoUrl] = useState(localStorage.getItem('repoUrl') || '')
    const [repoBranch, setRepoBranch] = useState(localStorage.getItem('repoBranch') || '')
    const [namespace, setNamespace] = useState(localStorage.getItem('namespace') || '')
    const choosePlanAndContinue = (plan) => {
        setPlan(plan)
        setStep(2)
    }

    const goBack = () => {
        setStep(1)
    }

    const handleSubmit = async (e) => {
        e.preventDefault()
        localStorage.setItem('projectName', projectName)
        localStorage.setItem('repoName', repoName)
        localStorage.setItem('repoUrl', repoUrl)
        localStorage.setItem('repoBranch', repoBranch)
        localStorage.setItem('plan', plan)
        localStorage.setItem('namespace', namespace)

        // proceed with payment 
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
        form.innerHTML += `<input type="hidden" name="url_ok" value="${REACT_APP_SUCCESS_URL_2}">`;
        form.innerHTML += `<input type="hidden" name="url_ko" value="${REACT_APP_ERROR_URL}">`;

        document.body.appendChild(form);
        form.submit();
    }

    return (
        <div className="min-h-screen  fixed  left-0 top-0  flex justify-center items-center inset-0 z-50 outline-none focus:outline-none">
            <div className="absolute bg-black opacity-60 inset-0 z-0"></div>
            <div
                ref={ref}
                className="relative bg-white rounded-lg shadow-lg z-50 p-5">
                {step === 1 && <div className="flex justify-between items-center">
                    <div className='flex flex-col items-center justify-center border-2  text-gray-800 p-2 rounded-md shadow-md space-y-4 w-52'>
                        <span className='text-3xl font-bold '>Starter</span>
                        <span className='text-lg'>49TND <span className='text-sm text-gray-500'>/month</span></span>
                        <span>1 CPU</span>
                        <span>2G RAM</span>
                        <span>30G Storage</span>
                        <span className='bg-creo text-white py-1 px-2 text-bold rounded-md cursor-pointer' onClick={() => choosePlanAndContinue("Starter")}>Get Started</span>
                    </div>
                    <div className='flex flex-col items-center justify-center border-2  text-gray-800 p-2 rounded-md shadow-md space-y-4 w-52 ml-2'>
                        <span className='text-3xl font-bold '>Pro</span>
                        <span className='text-lg'>99TND <span className='text-sm text-gray-500'>/month</span></span>
                        <span>2 CPU</span>
                        <span>4G RAM</span>
                        <span>80G Storage</span>
                        <span className='bg-creo text-white py-1 px-2 text-bold rounded-md cursor-pointer'
                            onClick={() => choosePlanAndContinue("Pro")}
                        >Get Started</span>
                    </div>
                    <div className='flex flex-col items-center justify-center border-2  text-gray-800 p-2 rounded-md shadow-md space-y-4 w-52 ml-2'>
                        <span className='text-3xl font-bold '>Elite</span>
                        <span className='text-lg'>49TND <span className='text-sm text-gray-500'>/month</span></span>
                        <span>4 CPU</span>
                        <span>8G RAM</span>
                        <span>120G Storage</span>
                        <span className='bg-creo text-white py-1 px-2 text-bold rounded-md cursor-pointer'
                            onClick={() => choosePlanAndContinue("Elite")}
                        >Get Started</span>
                    </div>
                </div>}
                {
                    step === 2 &&
                    <div className='flex flex-col items-center'>

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
                            <div className='flex items-center justify-center'>
                                <div
                                    className='py-2 px-6 border rounded-md mr-1 bg-gray-700 text-white cursor-pointer hover:bg-gray-600'
                                    onClick={() => goBack()}>Go back</div>
                                <button type="submit" className='py-2 px-6 border rounded-md bg-creo text-white'>Proceed with payment</button>

                            </div>

                        </form>
                    </div>
                }
            </div>
        </div>
    )
}
