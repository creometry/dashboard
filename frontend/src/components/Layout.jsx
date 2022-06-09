import { useState } from "react";
import { Link } from "react-router-dom";
import creometryLogo from "../../src/creo.png";

export const Layout = ({ children }) => {
    const { REACT_APP_NAMESPACE } = process.env;
    const [workloads, setWorkloads] = useState(false);
    const [configuration, setConfiguration] = useState(false);
    const [network, setNetwork] = useState(false);
    const [storage, setStorage] = useState(false);
    const [hpa, setHpa] = useState(false);
    const [customResource, setCustomResource] = useState(false);
    const [observability, setObservability] = useState(false);

    return (
        <div className="bg-zinc-900 w-full flex h-screen">
            <div className="lg:w-1/6 w-1/4 h-7/6 text-gray-400 bg-zinc-700 shadow-md flex flex-col items-center overflow-y-scroll scrollbar-thin scrollbar-thumb-creo scrollbar-track-gray-500">
                <div className="h-12 w-12">
                    <img src={creometryLogo} alt="logo" />
                </div>
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
                            <Link to="/events">
                                <div className="p-2 w-full rounded-md hover:bg-creo hover:text-gray-100 cursor-pointer">
                                    Events
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
                </ul>
            </div>
            <div className="lg:w-5/6 md:w-4/5 xs:w-3/4 w-full">
                <div className="bg-zinc-800 h-12 flex justify-between items-center  px-12">
                    <div className="bg-creo text-gray-200 px-3 py-1 rounded-md cursor-pointer font-bold">
                        Namespace : {REACT_APP_NAMESPACE}
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
                        <div className="ml-4">
                            <button
                                type="button"
                                className="bg-gray-800 flex text-sm rounded-full focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-gray-800 focus:ring-white"
                                id="user-menu-button"
                                aria-expanded="false"
                                aria-haspopup="true"
                            >
                                <span className="sr-only">Open user menu</span>
                                <img
                                    className="h-8 w-8 rounded-full"
                                    src="https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80"
                                    alt=""
                                />
                            </button>
                        </div>
                    </div>
                </div>
                <div className="h-5/6 overflow-scroll overscroll-contain">
                    {children}
                </div>

            </div>
        </div>)
}

