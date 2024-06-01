<script setup>
import { VCard, VCarousel,VCarouselItem } from 'vuetify/lib/components/index.mjs';
</script>

<template>
    <div class="title">
        <h1>FeasOJ</h1>
    </div>
    <!-- TODO:轮播图组件待处理 -->
    <v-card rounded="xl" style="margin: 50px;" elevation="10">
        <v-carousel show-arrows="hover">
            <v-carousel-item
                src="https://cdn.vuetifyjs.com/images/cards/docks.jpg"
                cover
            ></v-carousel-item>

            <v-carousel-item
                src="https://cdn.vuetifyjs.com/images/cards/hotel.jpg"
                cover
            ></v-carousel-item>
        </v-carousel>
    </v-card>
    
</template>