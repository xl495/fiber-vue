/// <reference types="vite/client" />


interface MenuMate {
    hidden: boolean,
    keepAlive: boolean,
    icon: string,
    title: string

}

interface MenuItem {
    path: string,
    name: string,
    component: string,
    parentId: number,
    sort: number,
    mate: MenuMate
}