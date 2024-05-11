<script setup>
import { ref } from 'vue'
import { Grid as TinyGrid, GridColumn as TinyGridColumn } from '@opentiny/vue'

const pagerConfig = ref({
  attrs: {
    currentPage: 1,
    pageSize: 40,
    total: 0,
    align: 'center',
    layout: 'prev, pager, next'
  }
})
const fetchData = ref({
  api: getData
})

const tableData = ref([
  
])

function getData({ page }) {
  const { currentPage, pageSize } = page
  const offset = (currentPage - 1) * pageSize

  return Promise.resolve({
    result: tableData.value.slice(offset, offset + pageSize),
    page: { total: tableData.value.length }
  })
}
</script>

<template>
    <div>
        <h1>题目</h1>
    </div>
    <div class="parent">
        <div class="restrictmaincomponents">
            <tiny-grid :fetch-data="fetchData" seq-serial :pager="pagerConfig" min-height="800" align="center" height="20">
                <tiny-grid-column field="question" title="题目"></tiny-grid-column>
                <tiny-grid-column field="numbers" title="通过人数" width="15%"></tiny-grid-column>
            </tiny-grid>
        </div>
    </div>
</template>

<style scoped>

</style>