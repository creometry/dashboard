import { Navbar } from "./Navbar";
import { Sidebar } from "./Sidebar";

export const Layout = ({ children }) => {
    return (
        <div className="bg-zinc-900 w-full flex h-screen">
            <Sidebar />
            <div className="lg:w-5/6 md:w-4/5 xs:w-3/4 w-full">
                <Navbar />
                <div className="h-5/6 overflow-scroll overscroll-contain">
                    {children}
                </div>

            </div>
        </div>
    )
}

