import { BrowserRouter, Routes, Route } from "react-router-dom";
import { Layout } from "./components/Layout";
import { ConfigmapList } from "./components/resources/ConfigmapList";
import { CrList } from "./components/resources/CrList";
import { CronjobList } from "./components/resources/CronjobList";
import { DeploymentList } from "./components/resources/DeploymentList";
import { HpaList } from "./components/resources/HpaList";
import { IngressList } from "./components/resources/IngressList";
import { JobList } from "./components/resources/JobList";
import { PodList } from "./components/resources/PodList";
import { PvcList } from "./components/resources/PvcList";
import { SecretList } from "./components/resources/SecretList";
import { ServiceList } from "./components/resources/ServiceList";
import { EndpointList } from "./components/resources/EndpointList";
import { EventList } from "./components/resources/EventList";
import { NotFound } from "./components/NotFound";
import { PaymentError } from "./components/PaymentError";
import { GithubAuth } from "./components/GithubAuth";
import { Plan } from "./components/plans/Plan";
import { Starter } from "./components/plans/Starter";
import { Dev } from "./components/plans/Dev";
import { Pro } from "./components/plans/Pro";
import { Steps } from "./components/Steps";

export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/plans" element={<Plan />} />
        <Route path="/plans/starter" element={<Starter />} />
        <Route path="/plans/dev" element={<Dev />} />
        <Route path="/plans/pro" element={<Pro />} />
        <Route path="/paymenterror" element={<PaymentError />} />
        <Route path="/success" element={<GithubAuth />} />
        <Route path="/steps" element={<Steps />} />
        <Route
          path="/"
          element={
            <Layout>
              <PodList />
            </Layout>
          }
        />
        <Route
          path="/pods"
          element={
            <Layout>
              <PodList />
            </Layout>
          }
        />
        <Route
          path="/deployments"
          element={
            <Layout>
              <DeploymentList isSts={false} />
            </Layout>
          }
        />
        <Route
          path="/services"
          element={
            <Layout>
              <ServiceList />
            </Layout>
          }
        />
        <Route
          path="/statefulsets"
          element={
            <Layout>
              <DeploymentList isSts={true} />
            </Layout>
          }
        />
        <Route
          path="/jobs"
          element={
            <Layout>
              <JobList />
            </Layout>
          }
        />
        <Route
          path="/cronjobs"
          element={
            <Layout>
              <CronjobList />
            </Layout>
          }
        />
        <Route
          path="/hpas"
          element={
            <Layout>
              <HpaList />
            </Layout>
          }
        />
        <Route
          path="/pvcs"
          element={
            <Layout>
              <PvcList />
            </Layout>
          }
        />
        <Route
          path="/secrets"
          element={
            <Layout>
              <SecretList />
            </Layout>
          }
        />
        <Route
          path="/configmaps"
          element={
            <Layout>
              <ConfigmapList />
            </Layout>
          }
        />
        <Route
          path="/customresources"
          element={
            <Layout>
              <CrList />
            </Layout>
          }
        />
        <Route
          path="/ingresses"
          element={
            <Layout>
              <IngressList />
            </Layout>
          }
        />
        <Route
          path="/endpoints"
          element={
            <Layout>
              <EndpointList />
            </Layout>
          }
        />
        <Route
          path="/events"
          element={
            <Layout>
              <EventList />
            </Layout>
          }
        />
        <Route path="*" element={<NotFound />} />
      </Routes>
    </BrowserRouter>
  );
}
