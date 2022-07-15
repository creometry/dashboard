import { Navbar } from "./Navbar";
import { Sidebar } from "./Sidebar";
import { PlansPopup } from "./PlansPopup";
import { AddMemberPopup } from "./AddMemberPopup";
import useStore, { usePopup, useAddMemberPopup } from "../zustand/state";
import { useNavigate } from "react-router-dom";
import { useEffect } from "react";

export const Layout = ({ children }) => {
    const { user } = useStore();
    const { isOpen } = usePopup();
    const { isPopupOpen } = useAddMemberPopup();
    const navigate = useNavigate()

    useEffect(() => {
        if (!user || user.access_token === "") {
            navigate("/login");
            return
        }
    }, [user]);

    return (
        <div className="bg-zinc-900 w-full flex h-screen">
            {
                isOpen === true && <PlansPopup />
            }
            {
                isPopupOpen === true && <AddMemberPopup />
            }
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

