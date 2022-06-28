import React, { useState, useEffect } from 'react'
import { Link } from "react-router-dom";
import creometryLogo from "../../src/creo.png";
import { usePopup } from '../zustand/state';

export const Sidebar = () => {
    const [workloads, setWorkloads] = useState(false);
    const [configuration, setConfiguration] = useState(false);
    const [network, setNetwork] = useState(false);
    const [storage, setStorage] = useState(false);
    const [hpa, setHpa] = useState(false);
    const [customResource, setCustomResource] = useState(false);
    const [observability, setObservability] = useState(false);
    const [planPaid, setPlanPaid] = useState(true);
    const { setIsOpen } = usePopup()
    useEffect(() => {
        const ns = localStorage.getItem('namespace')
        if (ns === null || ns === undefined || ns === "") {
            setPlanPaid(false)
        } else {
            setPlanPaid(true)
        }
    }, [])
    return (
        <div className="lg:w-1/6 w-1/4 h-7/6 text-gray-400 bg-zinc-700 shadow-md flex flex-col items-center overflow-y-scroll scrollbar-thin scrollbar-thumb-creo scrollbar-track-gray-500">
            <div className="h-12 w-12">
                <img src={creometryLogo} alt="logo" />
            </div>
            {planPaid === false &&
                <div className='mt-3 font-bold bg-red-600 text-white rounded-md px-1 py-1 hover:bg-red-500 cursor-pointer' onClick={() => setIsOpen(true)}>
                    Choose a plan continue
                </div>}
            <ul className="font-semibold text-lg pt-4 w-full flex flex-col">
                <li
                    className="p-2 w-full rounded-md hover:bg-creo hover:text-gray-100 cursor-pointer flex items-center justify-between"
                    onClick={() => setWorkloads(!workloads)}
                >
                    <div className="flex items-center justify-between w-full">
                        <span>Workloads</span>
                        <div>
                            {!workloads ? (
                                <svg
                                    xmlns="http://www.w3.org/2000/svg"
                                    className="h-6 w-6"
                                    fill="none"
                                    viewBox="0 0 24 24"
                                    stroke="currentColor"
                                    strokeWidth={2}
                                >
                                    <path
                                        strokeLinecap="round"
                                        strokeLinejoin="round"
                                        d="M19 9l-7 7-7-7"
                                    />
                                </svg>
                            ) : (
                                <svg
                                    xmlns="http://www.w3.org/2000/svg"
                                    className="h-6 w-6"
                                    fill="none"
                                    viewBox="0 0 24 24"
                                    stroke="currentColor"
                                    strokeWidth={2}
                                >
                                    <path
                                        strokeLinecap="round"
                                        strokeLinejoin="round"
                                        d="M5 15l7-7 7 7"
                                    />
                                </svg>
                            )}
                        </div>
                    </div>
                </li>
                {workloads && (
                    <div className="ml-8 flex flex-col">
                        <Link to="/pods">
                            <div className="p-2 w-full rounded-md hover:bg-creo hover:text-gray-100 cursor-pointer">
                                Pods
                            </div>
                        </Link>
                        <Link to="/deployments">
                            <div className="p-2 w-full rounded-md hover:bg-creo hover:text-gray-100 cursor-pointer">
                                Deployments
                            </div>
                        </Link>
                        <Link to="/statefulsets">
                            <div className="p-2 w-full rounded-md hover:bg-creo hover:text-gray-100 cursor-pointer">
                                Stateful sets
                            </div>
                        </Link>
                        <Link to="/jobs">
                            <div className="p-2 w-full rounded-md hover:bg-creo hover:text-gray-100 cursor-pointer">
                                Jobs
                            </div>
                        </Link>
                        <Link to="/cronjobs">
                            <div className="p-2 w-full rounded-md hover:bg-creo hover:text-gray-100 cursor-pointer">
                                Cronjobs
                            </div>
                        </Link>
                    </div>
                )}
                <li
                    className="p-2 w-full rounded-md hover:bg-creo hover:text-gray-100 cursor-pointer flex items-center justify-between"
                    onClick={() => setConfiguration(!configuration)}
                >
                    <div className="flex items-center justify-between w-full">
                        <span>Configuration</span>
                        <div>
                            {!configuration ? (
                                <svg
                                    xmlns="http://www.w3.org/2000/svg"
                                    className="h-6 w-6"
                                    fill="none"
                                    viewBox="0 0 24 24"
                                    stroke="currentColor"
                                    strokeWidth={2}
                                >
                                    <path
                                        strokeLinecap="round"
                                        strokeLinejoin="round"
                                        d="M19 9l-7 7-7-7"
                                    />
                                </svg>
                            ) : (
                                <svg
                                    xmlns="http://www.w3.org/2000/svg"
                                    className="h-6 w-6"
                                    fill="none"
                                    viewBox="0 0 24 24"
                                    stroke="currentColor"
                                    strokeWidth={2}
                                >
                                    <path
                                        strokeLinecap="round"
                                        strokeLinejoin="round"
                                        d="M5 15l7-7 7 7"
                                    />
                                </svg>
                            )}
                        </div>
                    </div>
                </li>
                {configuration && (
                    <div className="ml-8">
                        <Link to="/configmaps">
                            <div className="p-2 w-full rounded-md hover:bg-creo hover:text-gray-100 cursor-pointer">
                                Configmaps
                            </div>
                        </Link>
                        <Link to="/secrets">
                            <div className="p-2 w-full rounded-md hover:bg-creo hover:text-gray-100 cursor-pointer">
                                Secrets
                            </div>
                        </Link>
                    </div>
                )}
                <li
                    className="p-2 w-full rounded-md hover:bg-creo hover:text-gray-100 cursor-pointer flex items-center justify-between"
                    onClick={() => setNetwork(!network)}
                >
                    <div className="flex items-center justify-between w-full">
                        <span>Network</span>
                        <div>
                            {!network ? (
                                <svg
                                    xmlns="http://www.w3.org/2000/svg"
                                    className="h-6 w-6"
                                    fill="none"
                                    viewBox="0 0 24 24"
                                    stroke="currentColor"
                                    strokeWidth={2}
                                >
                                    <path
                                        strokeLinecap="round"
                                        strokeLinejoin="round"
                                        d="M19 9l-7 7-7-7"
                                    />
                                </svg>
                            ) : (
                                <svg
                                    xmlns="http://www.w3.org/2000/svg"
                                    className="h-6 w-6"
                                    fill="none"
                                    viewBox="0 0 24 24"
                                    stroke="currentColor"
                                    strokeWidth={2}
                                >
                                    <path
                                        strokeLinecap="round"
                                        strokeLinejoin="round"
                                        d="M5 15l7-7 7 7"
                                    />
                                </svg>
                            )}
                        </div>
                    </div>
                </li>
                {network && (
                    <div className="ml-8">
                        <Link to="/services">
                            <div className="p-2 w-full rounded-md hover:bg-creo hover:text-gray-100 cursor-pointer">
                                Services
                            </div>
                        </Link>
                        <Link to="/ingresses">
                            <div className="p-2 w-full rounded-md hover:bg-creo hover:text-gray-100 cursor-pointer">
                                Ingresses
                            </div>
                        </Link>
                        <Link to="/endpoints">
                            <div className="p-2 w-full rounded-md hover:bg-creo hover:text-gray-100 cursor-pointer">
                                Endpoints
                            </div>
                        </Link>
                    </div>
                )}
                <li
                    className="p-2 w-full rounded-md hover:bg-creo hover:text-gray-100 cursor-pointer flex items-center justify-between"
                    onClick={() => setStorage(!storage)}
                >
                    <div className="flex items-center justify-between w-full">
                        <span>Storage</span>
                        <div>
                            {!storage ? (
                                <svg
                                    xmlns="http://www.w3.org/2000/svg"
                                    className="h-6 w-6"
                                    fill="none"
                                    viewBox="0 0 24 24"
                                    stroke="currentColor"
                                    strokeWidth={2}
                                >
                                    <path
                                        strokeLinecap="round"
                                        strokeLinejoin="round"
                                        d="M19 9l-7 7-7-7"
                                    />
                                </svg>
                            ) : (
                                <svg
                                    xmlns="http://www.w3.org/2000/svg"
                                    className="h-6 w-6"
                                    fill="none"
                                    viewBox="0 0 24 24"
                                    stroke="currentColor"
                                    strokeWidth={2}
                                >
                                    <path
                                        strokeLinecap="round"
                                        strokeLinejoin="round"
                                        d="M5 15l7-7 7 7"
                                    />
                                </svg>
                            )}
                        </div>
                    </div>
                </li>
                {storage && (
                    <div className="ml-8">
                        <Link to="/pvcs">
                            <div className="p-2 w-full rounded-md hover:bg-creo hover:text-gray-100 cursor-pointer">
                                Persistent volume claims
                            </div>
                        </Link>
                    </div>
                )}
                <li
                    className="p-2 w-full rounded-md hover:bg-creo hover:text-gray-100 cursor-pointer flex items-center justify-between"
                    onClick={() => setCustomResource(!customResource)}
                >
                    <div className="flex items-center justify-between w-full">
                        <span>Custom Resources</span>
                        <div>
                            {!customResource ? (
                                <svg
                                    xmlns="http://www.w3.org/2000/svg"
                                    className="h-6 w-6"
                                    fill="none"
                                    viewBox="0 0 24 24"
                                    stroke="currentColor"
                                    strokeWidth={2}
                                >
                                    <path
                                        strokeLinecap="round"
                                        strokeLinejoin="round"
                                        d="M19 9l-7 7-7-7"
                                    />
                                </svg>
                            ) : (
                                <svg
                                    xmlns="http://www.w3.org/2000/svg"
                                    className="h-6 w-6"
                                    fill="none"
                                    viewBox="0 0 24 24"
                                    stroke="currentColor"
                                    strokeWidth={2}
                                >
                                    <path
                                        strokeLinecap="round"
                                        strokeLinejoin="round"
                                        d="M5 15l7-7 7 7"
                                    />
                                </svg>
                            )}
                        </div>
                    </div>
                </li>
                {customResource && (
                    <div className="ml-8">
                        <Link to="/customresources">
                            <div className="p-2 w-full rounded-md hover:bg-creo hover:text-gray-100 cursor-pointer">
                                Custom resources
                            </div>
                        </Link>

                    </div>
                )}
                <li
                    className="p-2 w-full rounded-md hover:bg-creo hover:text-gray-100 cursor-pointer flex items-center justify-between"
                    onClick={() => {
                        setObservability(!observability);
                    }}
                >
                    <div className="flex items-center justify-between w-full">
                        <span>Observability</span>
                        <div>
                            {!observability ? (
                                <svg
                                    xmlns="http://www.w3.org/2000/svg"
                                    className="h-6 w-6"
                                    fill="none"
                                    viewBox="0 0 24 24"
                                    stroke="currentColor"
                                    strokeWidth={2}
                                >
                                    <path
                                        strokeLinecap="round"
                                        strokeLinejoin="round"
                                        d="M19 9l-7 7-7-7"
                                    />
                                </svg>
                            ) : (
                                <svg
                                    xmlns="http://www.w3.org/2000/svg"
                                    className="h-6 w-6"
                                    fill="none"
                                    viewBox="0 0 24 24"
                                    stroke="currentColor"
                                    strokeWidth={2}
                                >
                                    <path
                                        strokeLinecap="round"
                                        strokeLinejoin="round"
                                        d="M5 15l7-7 7 7"
                                    />
                                </svg>
                            )}
                        </div>
                    </div>
                </li>
                {observability && (
                    <div className="ml-8">
                        <Link to="/logs">
                            <div className="p-2 w-full rounded-md hover:bg-creo hover:text-gray-100 cursor-pointer">
                                Logs
                            </div>
                        </Link>
                        <Link to="/metrics">
                            <div className="p-2 w-full rounded-md hover:bg-creo hover:text-gray-100 cursor-pointer">
                                Metrics
                            </div>
                        </Link>
                        <Link to="/monitoring">
                            <div className="p-2 w-full rounded-md hover:bg-creo hover:text-gray-100 cursor-pointer">
                                Monitoring
                            </div>
                        </Link>
                    </div>
                )}

                <li
                    className="p-2 w-full rounded-md hover:bg-creo hover:text-gray-100 cursor-pointer flex items-center justify-between"
                    onClick={() => {
                        setHpa(!hpa);
                    }}
                >
                    <div className="flex items-center justify-between w-full">
                        <span>Scaling</span>
                        <div>
                            {!hpa ? (
                                <svg
                                    xmlns="http://www.w3.org/2000/svg"
                                    className="h-6 w-6"
                                    fill="none"
                                    viewBox="0 0 24 24"
                                    stroke="currentColor"
                                    strokeWidth={2}
                                >
                                    <path
                                        strokeLinecap="round"
                                        strokeLinejoin="round"
                                        d="M19 9l-7 7-7-7"
                                    />
                                </svg>
                            ) : (
                                <svg
                                    xmlns="http://www.w3.org/2000/svg"
                                    className="h-6 w-6"
                                    fill="none"
                                    viewBox="0 0 24 24"
                                    stroke="currentColor"
                                    strokeWidth={2}
                                >
                                    <path
                                        strokeLinecap="round"
                                        strokeLinejoin="round"
                                        d="M5 15l7-7 7 7"
                                    />
                                </svg>
                            )}
                        </div>
                    </div>
                </li>
                {hpa && (
                    <div className="ml-8">
                        <Link to="/hpas">
                            <div className="p-2 w-full rounded-md hover:bg-creo hover:text-gray-100 cursor-pointer">
                                Horizontal pod autoscalers
                            </div>
                        </Link>

                    </div>
                )}
                <div className="ml-1">
                    <Link to="/appstore">
                        <div className="p-2 w-full rounded-md hover:bg-creo hover:text-gray-100 cursor-pointer">
                            App store
                        </div>
                    </Link>

                </div>
                <div className="ml-1 mt-1">
                    <Link to="/billing">
                        <div className="p-2 w-full rounded-md hover:bg-creo hover:text-gray-100 cursor-pointer">
                            Billing
                        </div>
                    </Link>
                </div>
                <div className="ml-1 mt-1">
                    <Link to="/events">
                        <div className="p-2 w-full rounded-md hover:bg-creo hover:text-gray-100 cursor-pointer">
                            Events
                        </div>
                    </Link>
                </div>

                <div className="ml-1 mt-1">
                    <Link to="/accesscontrol">
                        <div className="p-2 w-full rounded-md hover:bg-creo hover:text-gray-100 cursor-pointer">
                            Access control
                        </div>
                    </Link>
                </div>

            </ul>
        </div>
    )
}
