<template>
    <h1>FeasOJ</h1>
    <div class="parent">
        <div class="restrictmaincomponents">
            <tiny-carousel height="400px" autoplay>
                <!-- TODO:轮播IMG -->
                <tiny-carousel-item class="carousel-item" v-for="item in 2" :key="item">
                <h3>{{ item }}</h3>
                </tiny-carousel-item>
            </tiny-carousel>
            <tiny-collapse v-model="activeNames">
                <tiny-collapse-item title="公告" name="1">
                    <h2>现在处于开发阶段</h2>
                </tiny-collapse-item>
            </tiny-collapse>
        </div>
        <!-- TODO:组件待扩展 -->
    </div>
</template>

<script setup>
import { ref } from 'vue'
import { Collapse as TinyCollapse, CollapseItem as TinyCollapseItem } from '@opentiny/vue'
import { Carousel as TinyCarousel, CarouselItem as TinyCarouselItem } from '@opentiny/vue'

const activeNames = ref(['1'])
</script>

<style scoped>
.carousel-item {
  background-color: #edf0f3;
}
</style>