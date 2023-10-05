import {defineStore} from "pinia";
import {ref} from "vue";
import type { Ref } from 'vue';

export enum ENotificationsType {
    'error' = 'error',
    'warning' = 'warning',
    'success' = 'success',
    'info' = 'info',
}

export enum EButtonsType {
    'retry' = 'retry',
    'more_information' = 'more_information',
}

type TNotificationButton = {
    type: EButtonsType;
    callback: () => void;
}

type NotificationActions = TNotificationButton[];

type Notification = {
    type: ENotificationsType;
    text: string;
    id: number;
    duration?: number;
}

type NotificationProps = {
    text: string;
    duration?: number; 
    actions: NotificationActions;
}

export interface NotificationService {
    notifications: Notification[];
    showNotification:  {
        error: (props: NotificationProps) => void;
        warning: (props: NotificationProps) => void;
        success: (props: NotificationProps) => void;
    }
    hideNotification: (id: number) => void;
}

export const useNotificationsStore = defineStore('notifications', () => {
    const notifications: Ref<Notification[]> = ref([]);

    const showNotification: NotificationService['showNotification'] = {
        error: ({text, duration, actions}) => {
            notifications.value.push({
                id: notifications.value.length + 1,
                text,
                type: ENotificationsType.error,
                duration,
                actions,
            });
            if (duration) {
                setTimeout(() => {
                    hideNotification(notifications.value.length);
                }, duration)
            }
        },
        warning: ({text, duration, actions}) => {
            notifications.value.push({
                id: notifications.value.length + 1,
                text,
                type: ENotificationsType.warning,
                duration,
                actions,
            });
            if (duration) {
                setTimeout(() => {
                    hideNotification(notifications.value.length);
                }, duration)
            }
        },
        success: ({text, duration, actions}) => {
            notifications.value.push({
                id: notifications.value.length + 1,
                text,
                type: ENotificationsType.success,
                duration,
                actions,
            });
            if (duration) {
                setTimeout(() => {
                    hideNotification(notifications.value.length);
                }, duration)
            }
        }
    }

    const hideNotification: NotificationService['hideNotification'] = (id) =>  {
        const itemIndex = notifications.value.findIndex((item: Notification) => item.id === id);
        notifications.value.splice(itemIndex, 1);
    }

    return {
        notifications,
        showNotification,
        hideNotification,
    }
});