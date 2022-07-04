import create from "zustand";
import { persist } from "zustand/middleware";

const useStore = create(
  persist(
    (set, get) => ({
      user: {
        id: "",
        login: "",
        name: "",
        access_token: "",
        avatar_url: "",
        email: "",
      },
      setUser: (user) =>
        set((state) => ({
          user: {
            ...state.user,
            ...user,
          },
        })),
    }),
    {
      name: "user_data",
    }
  )
);

export const usePopup = create((set, get) => ({
  isOpen: false,
  setIsOpen: (isOpen) => set((state) => ({ isOpen })),
}));

export const useAddMemberPopup = create((set, get) => ({
  isPopupOpen: false,
  setIsPopupOpen: (isPopupOpen) => set((state) => ({ isPopupOpen })),
}));

export default useStore;
