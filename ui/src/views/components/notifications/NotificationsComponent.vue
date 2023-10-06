<script setup lang="ts">
import {useNotificationsStore, ENotificationsType} from "../../../modules/store/notifications.store";

const notificationService = useNotificationsStore();

function hideNotification(id: number) {
    notificationService.hideNotification(id);
}

</script>

<template>
    <div class="notifications-container">
        <v-alert 
            v-for="notification of notificationService.notifications" 
            :title="(notification.type).toUpperCase()" 
            :type="notification.type"
            :text="notification.text"
            :icon="notification.type === ENotificationsType.error ? 'mdi-alert-outline' : false"
            class="notification"
            :style="(notification.actions && notification.actions.length) ? 'padding-bottom: 40px' : 'padding-bottom: 16px'"
        >
            <v-icon class="notification__close-btn" @click="hideNotification(notification.id)">mdi-close</v-icon>
            <template #append>
                <div v-if="notification.actions && notification.actions.length" class="notification__actions">
                    <v-btn dense variant="tonal" v-for="action of notification.actions">
                        {{ action.type.charAt(0).toUpperCase() + action.type.slice(1)}}
                    </v-btn>
                </div>
            </template>
        </v-alert>
    </div>
</template>

<style lang="scss" scoped>
    .notifications-container {
        position: fixed;
        top: 68px;
        right: 4px;
        min-width: 400px;
        max-width: 400px;
        z-index: 9999;
        .notification {
            position: relative;
            margin-top: 4px;
            &__close-btn {
                position: absolute;
                top: 4px;
                right: 4px;
                cursor: pointer;
            }
            &__actions {
                position: absolute;
                display: flex;
                gap: 4px;
                bottom: 8px;
                right: 8px;
            }
        }
    }
</style>