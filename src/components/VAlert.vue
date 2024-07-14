<script setup>
import { VDialog,VCard,VCardText, VCardTitle, VCardActions, VBtn} from 'vuetify/components';
import { dialog,dialogMessage } from '../utils/alert.js';
</script>

<template>
    <v-dialog v-model="dialog" persistent max-width="290">
        <v-card rounded="xl">
        <v-card-title class="text-h5">提示</v-card-title>
        <v-card-text>{{ dialogMessage }}</v-card-text>
        <v-card-actions>
            <v-btn color="blue darken-1" @click="dialog = false" rounded="xl">好的</v-btn>
        </v-card-actions>
        </v-card>
    </v-dialog>
</template>

<style scoped>

</style>