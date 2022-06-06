import { useState } from "react";
import axios from "axios";
import loadingSpinner from "../src/loading.gif";
import creometryLogo from "../src/creo.png";
import { Resource } from "./components/Resource";

export default function App() {
  const { REACT_APP_URL } = process.env;
  const NAMESPACE = "colibris";
  const [loading, setLoading] = useState(false);
  const [err, setError] = useState("");
  const [resource, setResource] = useState("pods");
  const [data, setData] = useState([]);
  const [workloads, setWorkloads] = useState(false);
  const [configuration, setConfiguration] = useState(false);
  const [network, setNetwork] = useState(false);
  const [storage, setStorage] = useState(false);
  const [hpa, setHpa] = useState(false);
  const [observability, setObservability] = useState(false);
  const getResourceData = async (rs) => {
    setLoading(true);
    setResource(rs);
    try {
      const resp = await axios.get(
        `${REACT_APP_URL}/api/v1/${rs}/${NAMESPACE}`
      );
      setData(resp.data);
      setLoading(false);
    } catch (error) {
      setError(error);
      setLoading(false);
    }
  };
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
            <div className="ml-8">
              <Resource resourceName="Pods" getResourceData={getResourceData} />
              <Resource
                resourceName="Deployments"
                getResourceData={getResourceData}
              />
              <Resource
                resourceName="Statefulsets"
                getResourceData={getResourceData}
              />
              <Resource resourceName="Jobs" getResourceData={getResourceData} />
              <Resource
                resourceName="Cronjobs"
                getResourceData={getResourceData}
              />
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
              <Resource
                resourceName="Configmaps"
                getResourceData={getResourceData}
              />
              <Resource
                resourceName="Secrets"
                getResourceData={getResourceData}
              />
              <Resource
                resourceName="Events"
                getResourceData={getResourceData}
              />
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
              <Resource
                resourceName="Services"
                getResourceData={getResourceData}
              />
              <Resource
                resourceName="Ingresses"
                getResourceData={getResourceData}
              />
              <Resource
                resourceName="Endpoints"
                getResourceData={getResourceData}
              />
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
              <Resource
                resourceName="Persistent volume claims"
                getResourceData={getResourceData}
              />
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
              <Resource resourceName="Logs" getResourceData={getResourceData} />
              <Resource
                resourceName="Metrics"
                getResourceData={getResourceData}
              />
              <Resource
                resourceName="Monitoring"
                getResourceData={getResourceData}
              />
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
              <Resource
                resourceName="Horizontal pod autoscalers"
                getResourceData={getResourceData}
              />
            </div>
          )}
          <div className="ml-1">
            <Resource
              resourceName="App store"
              getResourceData={getResourceData}
            />
          </div>
          <div className="ml-1 mt-1">
            <Resource
              resourceName="Billing"
              getResourceData={getResourceData}
            />
          </div>
        </ul>
      </div>
      <div className="lg:w-5/6 md:w-4/5 xs:w-3/4 w-full">
        <div className="bg-zinc-800 h-12 flex justify-between items-center  px-12">
          <div className="bg-creo text-gray-100 px-2 rounded-md cursor-pointer">
            Namespace : {NAMESPACE}
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
        <div
          className={`m-6 p-2  bg-neutral-800 shadow-lg rounded-sm text-gray-100 h-5/6 overflow-scroll overscroll-contain ${
            loading && "flex justify-center items-center"
          }`}
        >
          {" "}
          {loading ? (
            <div className="h-12 w-12">
              <img src={loadingSpinner} alt="" />
            </div>
          ) : err === "" ? (
            <div className="flex flex-col">
              <div>{resource}</div>
              <div>{JSON.stringify(data)}</div>
            </div>
          ) : (
            <div className="text-red-600">{JSON.stringify(err)}</div>
          )}
        </div>
      </div>
    </div>
  );
}
