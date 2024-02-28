import { getUsers, login } from "@/api/user";
import { defineStore } from "pinia";

const useUserStore = defineStore({
    id: "user",
    state: () => ({
        userInfo: {
            ID: null,
            username: "John Doe",
            age: 30,
        },
        token: "",
        isCollapseMenu: false,
    }),
    actions: {
        async fetchUser() {
            const data = await getUsers()
        },
        async login({ username, password }: { username: string, password: string }) {
            const data = await login({ username, password });
            this.userInfo = data.user;
            this.token = data.token;
            localStorage.setItem("token", data.token);
            return data
        },
        toggleCollapsed() {
            this.isCollapseMenu = !this.isCollapseMenu;
        }
    },
});


export { useUserStore };