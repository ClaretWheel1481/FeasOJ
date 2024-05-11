<script setup>
import { Grid as TinyGrid, GridColumn as TinyGridColumn } from '@opentiny/vue'
import { reactive } from 'vue'

const tableData = reactive([
//   TODO:数据库传入前50Score数据
    
])
</script>

<template>
    <div>
        <h1>Top50</h1>
    </div>
    <div class="parent">
        <div class="restrictmaincomponents">
            <tiny-grid :data="tableData" min-height="500" header-align="center" size="medium" align="center">
                <tiny-grid-column type="index" title="Rank"></tiny-grid-column>
                <tiny-grid-column field="username" title="用户名称"></tiny-grid-column>
                <tiny-grid-column field="score" title="Score"></tiny-grid-column>
            </tiny-grid>
        </div>
    </div>
</template>

<style scoped>

</style>