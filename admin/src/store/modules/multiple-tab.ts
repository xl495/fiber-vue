import { cloneDeep } from 'lodash-es'
import { defineStore } from 'pinia'
import { RouteLocationNormalized } from 'vue-router'
import router from '@/router'
export interface TabRouteLocationNormalized extends RouteLocationNormalized {
    KEY?: string
}
export interface MultipleTabState {
    tabList: Map<string, TabRouteLocationNormalized>
    currentKey: string
    contextMenuKey?: string | null
}

export const MULTIPLE_TABS_KEY = 'MULTIPLE_TABS__KEY'

export const MULTIPLE_TABS_CURRENT_KEY = 'MULTIPLE_TABS_CURRENT_KEY'

const getRawTaKey = (route: TabRouteLocationNormalized): string => {
    // 拼接一个唯一的key
    const key = `${String(route.name)}-${JSON.stringify(route.query)}-${JSON.stringify(route.params)}`
    return key
}

export const useMultipleTabStore = defineStore({
    id: 'app-multiple-tab',
    state: (): MultipleTabState => ({
        tabList: new Map(),
        currentKey: '',
        contextMenuKey: '',
    }),
    getters: {
        getTabList(state): TabRouteLocationNormalized[] {
            return Array.from(state.tabList.values())
        },
    },
    actions: {
        setTab(route: TabRouteLocationNormalized) {
            const key = getRawTaKey(route)
            if (key === this.currentKey) return
            if (!this.tabList.get(key)) {
                const r = cloneDeep({ ...route, KEY: key, matched: [] })
                this.tabList.set(key, r)
            }
            this.currentKey = key
        },
        setContextMenuKey(key: string) {
            this.contextMenuKey = key
        },
        removeTab(key: string) {
            this.tabList.delete(key)
            if (!this.tabList.get(this.currentKey) && this.getTabList.length) {
                router.push(this.getTabList[this.getTabList.length - 1])
            } else if (this.getTabList.length === 0) {
                router.push({ path: '/' })
            }
        },
        changeTab(key: string) {
            console.log(this.currentKey,key);
            
            if (key === this.currentKey) return
            const route = this.tabList.get(key) as TabRouteLocationNormalized
            router.push(route)
        },
        // 设置当前页面标题
        setCurrentTitle(title: string) {
            const route = this.tabList.get(this.currentKey)
            if (route) {
                this.tabList.set(this.currentKey, {
                    ...route,
                    meta: {
                        ...route.meta,
                        title,
                    },

                })
            }
        },
        removeLeftTabs() {
            // 从左到右删除
            for (const [key] of this.tabList) {
                if (key !== this.contextMenuKey) {
                    this.tabList.delete(key)
                } else {
                    break
                }
            }
            if (!this.tabList.get(this.currentKey)) {
                router.push(this.getTabList[0])
            }
        },
        removeRightTabs() {
            // 删除右侧
            let start = false
            for (const [key] of this.tabList) {
                if (start) {
                    this.tabList.delete(key)
                }
                if (key === this.contextMenuKey) {
                    start = true
                }
            }
            if (!this.tabList.get(this.currentKey)) {
                router.push(this.getTabList[this.getTabList.length - 1])
            }
        },
        removeAllTabs() {
            for (const iterator of this.tabList.keys()) {
                this.tabList.delete(iterator)
            }
            router.push({ path: '/' })
        },
        removeOtherTabs() {
            for (const iterator of this.tabList.keys()) {
                if (iterator !== this.contextMenuKey) {
                    this.tabList.delete(iterator)
                }
            }
            router.push(this.getTabList[0])
        },
    },
})

const store = sessionStorage

const multipleTabStore = useMultipleTabStore()

const init = () => {
    const tabsList = sessionStorage.getItem(MULTIPLE_TABS_KEY)
    const tabCurrent = sessionStorage.getItem(MULTIPLE_TABS_CURRENT_KEY)

    if (tabsList) {
        const tabs = JSON.parse(tabsList)
        multipleTabStore.tabList = new Map(tabs.map((item: TabRouteLocationNormalized) => [item.KEY, item]))
    }

    if (tabCurrent) {
        multipleTabStore.currentKey = tabCurrent
    }
}
// 初始化
init()

multipleTabStore.$subscribe((mutation, state) => {
    store.setItem(MULTIPLE_TABS_CURRENT_KEY, state.currentKey)
    store.setItem(MULTIPLE_TABS_KEY, JSON.stringify(Array.from(state.tabList.values())))
})