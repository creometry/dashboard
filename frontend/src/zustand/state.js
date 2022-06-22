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

export default useStore;
